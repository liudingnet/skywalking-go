/*
 * Licensed to the OpenSkywalking under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenSkywalking licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package trace



type TracerContext interface {
	inject(carrier ContextCarrier)
	extract(carrier ContextCarrier)
	capture() *ContextSnapshot
	continued(snapshot ContextSnapshot)
	getReadableGlobalTraceId() string
	createEntrySpan(operationName string) Span
	createLocalSpan(operationName string) Span
	createExitSpan(operationName string,remotePeer string) Span
	stopSpan(span Span)
}

type TracingContext struct {
	lastWarningTimestamp int64
	sampler defaultSampler
	segment TraceSegment
	spanIdGenerator int32
	activeSpanStacks stack
}

var _ TracerContext=(*TracingContext)(nil)

//Inject  Inject the context into the given carrier, only when the active span is an exit one.
func (c *TracingContext)inject(carrier ContextCarrier){
	span :=c.activeSpanStacks.peek().(TracingSpan)
	if &span == nil{
		panic("Inject can be done only in Exit Span")
	}

	carrier.traceSegmentId=c.segment.traceSegmentId
	carrier.spanId=span.spanId
	carrier.parentApplicationInstanceId=c.segment.applicationId

	if span.peerId ==0 {
		carrier.peerHost=span.peer
	}else {
		carrier.peerId=span.peerId
	}

	refs :=c.segment.refs

	firstSpan :=c.activeSpanStacks.getLast().(TracingSpan)
	var operationId  int32
	var  operationName string
	var entryApplicationInstanceId  int32
	if refs != nil && len(refs)>0{
		ref :=refs[0]
		operationId=ref.entryOperationId
		operationName=ref.entryOperationName
		entryApplicationInstanceId=ref.entryApplicationInstanceId
	}else {
		s := firstSpan
		operationId=s.operationId
		operationName=s.operationName
		entryApplicationInstanceId=c.segment.applicationInstanceId
	}

	carrier.entryApplicationInstanceId=entryApplicationInstanceId

	if operationId==0 {
		if len(operationName)>0 {
			carrier.SetEntryOperationName(operationName)
		}

	}else {
		carrier.entryOperationId=operationId
	}

	parentOperationId :=firstSpan.operationId
	if parentOperationId ==0 {
		carrier.SetEntryOperationName(firstSpan.operationName)
	}else {
		carrier.entryOperationId=parentOperationId
	}

	carrier.primaryDistributedTraceId=c.segment.relatedGlobalTraces.relatedGlobalTraces[0]

}


//Extract Extract the carrier to build the reference for the pre segment.
func(c *TracingContext)extract(carrier ContextCarrier){
	traceSegmentRef,err :=NewTraceSegmentRefByContextCarrier(carrier)
	if err !=nil{
		return
	}
	c.segment.Ref(traceSegmentRef)
	c.segment.RelatedGlobalTrace(carrier.primaryDistributedTraceId)
	span :=c.activeSpanStacks.peek().(TracingSpan)
	if &span == nil{
		panic("Inject can be done only in Exit Span")
	}

	span.Ref(*traceSegmentRef)
}

//capture  capture the snapshot of current context.
func (c *TracingContext)capture() *ContextSnapshot  {
	refs :=c.segment.refs
	span :=c.activeSpanStacks.peek().(TracingSpan)
	if &span == nil{
		panic("Inject can be done only in Exit Span")
	}
	snapshot :=NewContextSnapshot(c.segment.traceSegmentId,span.spanId,c.segment.relatedGlobalTraces.relatedGlobalTraces)

	var  entryOperationId int32
	var entryOperationName string
	var entryApplicationInstanceId int32
	firstSpan := c.activeSpanStacks.getLast().(TracingSpan)

	if refs != nil && len(refs)>0 {
		ref :=refs[0]
		entryOperationId=ref.entryOperationId
		entryOperationName=ref.entryOperationName
		entryApplicationInstanceId=ref.entryApplicationInstanceId
	}else {
		entryOperationId = firstSpan.operationId
		entryOperationName = firstSpan.operationName
		entryApplicationInstanceId = c.segment.applicationInstanceId
	}
	snapshot.SetEntryApplicationInstanceId(entryApplicationInstanceId)

	if entryOperationId ==0 {
		if len(entryOperationName)>0 {
			snapshot.SetEntryOperationName(entryOperationName)
		}
	}else{
		snapshot.SetEntryOperationId(entryOperationId)
	}

	if firstSpan.operationId == 0 {
		snapshot.SetParentOperationName(firstSpan.operationName)
	}else {
		snapshot.SetParentOperationId(firstSpan.operationId)
	}

	return snapshot
}

func (c *TracingContext)continued(snapshot ContextSnapshot) {
	segmentRef,err :=NewTraceSegmentRefByContextSnapshot(snapshot)
	if err != nil {
		return
	}

	c.segment.Ref(segmentRef)
	span :=c.activeSpanStacks.peek().(TracingSpan)
	if &span == nil{
		panic("Inject can be done only in Exit Span")
	}
	span.Ref(*segmentRef)
	c.segment.RelatedGlobalTrace(snapshot.primaryDistributedTraceId)
}

func (c *TracingContext)getReadableGlobalTraceId() string{
	return c.segment.relatedGlobalTraces.relatedGlobalTraces[0].ToString()
}


func (c *TracingContext)createEntrySpan(operationName string) Span{
	if c.isLimitMechanismWorking() {
		span :=NewNoopEntrySpan()
		return  span
	}
	parentSpan :=c.activeSpanStacks.peek().(TracingSpan)
	var parentSpanId int32
	if &parentSpan == nil  {
		parentSpanId = -1
	}else{
		parentSpanId = parentSpan.spanId
	}

	if &parentSpan != nil && parentSpan.IsEntry{
		parentSpan.operationName=operationName
		parentSpan.Start(0)
		return  &parentSpan
	}else {
		entrySpan :=newTracingSpan(&SpanOptions{spanId:c.spanIdGenerator+1,parentSpanId:parentSpanId,operationName:operationName,spanType:EntrySpan})
	    entrySpan.Start(0)
		c.activeSpanStacks.push(entrySpan)
		return  entrySpan
	}


}


func (c *TracingContext)createLocalSpan(operationName string) Span{
	if ! c.isLimitMechanismWorking() {
		span :=NewNoopEntrySpan()
		return  span
	}

	parentSpan :=c.activeSpanStacks.peek().(TracingSpan)
	var parentSpanId int32
	if &parentSpan == nil  {
		parentSpanId = -1
	}else{
		parentSpanId = parentSpan.spanId
	}

	span :=newTracingSpan(&SpanOptions{spanId:c.spanIdGenerator+1,parentSpanId:parentSpanId,operationName:operationName,peerId:0,spanType:LocalSpan})

	span.Start(0)

	c.activeSpanStacks.push(span)

	return  span

}


func (c *TracingContext)createExitSpan(operationName string,remotePeer string) Span{
	parentSpan :=c.activeSpanStacks.peek().(TracingSpan)

	if &parentSpan != nil &&  parentSpan.IsExit  {
		parentSpan.Start(0)
		return &parentSpan
	}else {
		var parentSpanId int32
		if &parentSpan == nil  {
			parentSpanId = -1
		}else{
			parentSpanId = parentSpan.spanId
		}

		if c.isLimitMechanismWorking() {
			span :=NewNoopExitSpan(0,remotePeer)
			return  span
		}else {

			exitSpan := newTracingSpan(&SpanOptions{spanId:c.spanIdGenerator+1,parentSpanId:parentSpanId,
			operationName:operationName,peer:remotePeer,spanType:ExitSpan})
			c.activeSpanStacks.push(exitSpan)
			return  exitSpan
		}
	}

}

func (c *TracingContext) stopSpan(span Span){
	lastSpan :=c.activeSpanStacks.peek().(TracingSpan)
	if lastSpan.Finish(c.segment) {
		c.activeSpanStacks.pop()
	}

	if  c.activeSpanStacks.Size==0{
		c.finish()
	}

}



func (c *TracingContext)finish(){
	c.segment.Finish(c.isLimitMechanismWorking())
	if !c.segment.hasRef && c.segment.isSingleSpanSegment {
		if !c.sampler.sampled() {
			c.segment.isIgnore=true
		}
	}

	afterFinished(c.segment)

}



func (c *TracingContext)isLimitMechanismWorking() bool{
	if c.spanIdGenerator >=300 {
		currentTimeMillis :=getMillis()
		if currentTimeMillis -c.lastWarningTimestamp > 30 * 1000 {
			c.lastWarningTimestamp = currentTimeMillis
		}
		return true
	}else {
		return false
	}
}



