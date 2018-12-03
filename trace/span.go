/*
 *
 *  * Licensed to the OpenSkywalking under one or more
 *  * contributor license agreements.  See the NOTICE file distributed with
 *  * this work for additional information regarding copyright ownership.
 *  * The ASF licenses this file to You under the Apache License, Version 2.0
 *  * (the "License"); you may not use this file except in compliance with
 *  * the License.  You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *  *
 *
 */

package trace

import (
	"reflect"
	"sync"
)

type SpanLayer int

const (
	DB            SpanLayer = iota
	RPC_FRAMEWORK
	HTTP
	MQ
	CACHE
)

type SpanType int

const (
	EntrySpan SpanType =iota
	ExitSpan
	LocalSpan
)


type Span interface {
	SetComponent(componentName string)
	Tag(key string, value string)
	SetLayer(layer *SpanLayer)
	Log(err error)
	ErrorOccurred()
	Start(timestamp int64)
	Ref(traceSegmentRef TraceSegmentRef)
	SetOperationName(operationName string)
	SetOperationId(operationId  int32)
}



type TracingSpan struct {
	sync.RWMutex
	spanId        int32
	parentSpanId  int32
	tags          map[string]string
	operationName string
	operationId   int32
	layer         *SpanLayer
	startTime     int64
	endTime       int64
	errorOccurred bool
	componentId   int32
	componentName string
	logs          []LogDataEntity
	refs          []TraceSegmentRef
	IsExit        bool
	IsEntry       bool
	stackDepth    int32
	peer          string
	peerId         int32
	spanType SpanType
	currentMaxDepth int32
}


type SpanOptions struct {
	spanId int32
	parentSpanId int32
	operationName string
	peer string
	peerId int32
	spanType SpanType
}

var _ Span  = (*TracingSpan)(nil)


func newTracingSpan(spanOptions *SpanOptions) *TracingSpan{
	span :=&TracingSpan{
		spanId:spanOptions.spanId,
		parentSpanId:spanOptions.parentSpanId,
		operationName:spanOptions.operationName,
		peer:spanOptions.peer,
		peerId:spanOptions.peerId,
		spanType:spanOptions.spanType,
	}

	if span.spanType ==EntrySpan {
		span.stackDepth=0
		span.IsExit=false
		span.IsEntry=true
	}else if span.spanType==ExitSpan {
		span.IsEntry = false
		span.IsExit = true
	}else {
		span.IsEntry=false
		span.IsExit=false
	}

	return  span
}


func (c *TracingSpan) SetOperationName(operationName string) {
	c.Lock()
	defer c.Unlock()
	if c.spanType==EntrySpan {

		if c.currentMaxDepth==c.stackDepth {
			c.operationName=operationName
		}
	}else if c.spanType==ExitSpan {
		if c.stackDepth==1 {
			c.operationName=operationName
		}
	}else{
		c.operationName=operationName
	}
}


func (c *TracingSpan) SetOperationId(operationId  int32) {
	c.Lock()
	defer c.Unlock()
	if c.spanType==EntrySpan {

		if c.currentMaxDepth==c.stackDepth {
			c.operationId=operationId
		}
	}else if c.spanType==ExitSpan {
		if c.stackDepth==1 {
			c.operationId=operationId
		}
	}else{
		c.operationId=operationId
	}
}


func (c *TracingSpan) getStart(timestamp int64) {
	if c.startTime == 0 {
		c.startTime = getMillis()
		return
	}
	c.startTime = timestamp
}

func (c *TracingSpan) clearWhenRestart() {
	c.componentId = 0
	c.componentName = ""
	c.layer = nil
	c.logs = nil
	c.tags = nil
}

func (c *TracingSpan) Start(timestamp int64) {
	c.Lock()
	defer c.Unlock()
	if c.spanType==EntrySpan {
		c.currentMaxDepth=c.stackDepth+1
		if c.currentMaxDepth==1 {
			c.getStart(timestamp)
		}
		c.clearWhenRestart()
	}else if c.spanType==ExitSpan {
		if c.stackDepth+1==1 {
			c.getStart(timestamp)
		}
	}else{
		c.getStart(timestamp)
	}

}


func (c *TracingSpan) getTag(key string, value string) {

	if c.tags == nil {
		c.tags = make(map[string]string)
	}
	c.tags[key] = value
}



func (c *TracingSpan) Tag(key string, value string) {
	c.Lock()
	defer c.Unlock()
	if c.spanType==EntrySpan {

		if c.stackDepth==c.currentMaxDepth {
			c.getTag(key,value)
		}
	}else if c.spanType==ExitSpan {
		if c.stackDepth==1 {
			c.getTag(key,value)
		}
	}else{
		c.getTag(key,value)
	}
}


func (c *TracingSpan) SetLayer(layer *SpanLayer) {
	c.Lock()
	defer c.Unlock()
	if c.spanType==EntrySpan {
		if c.stackDepth==c.currentMaxDepth {
			c.layer = layer
		}
	}else if c.spanType==ExitSpan {
		if c.stackDepth==1 {
			c.layer = layer
		}
	}else{
		c.layer = layer
	}
}




func (c *TracingSpan) SetComponent(componentName string)  {
	c.Lock()
	defer c.Unlock()
	if c.spanType==EntrySpan {
		if c.stackDepth==c.currentMaxDepth {
			c.componentName = componentName
		}
	}else if c.spanType==ExitSpan {
		if c.stackDepth==1 {
			c.componentName = componentName
		}
	}else{
		c.componentName = componentName
	}
}



func (c *TracingSpan) Log(err error) {
	c.Lock()
	defer c.Unlock()
	logs :=make(map[string]string)
	logs["event"]="error"
	logs["error.kind"]=reflect.TypeOf(err).Name()
	logs["message"]=err.Error()
	c.logs=append(c.logs,LogDataEntity{timestamp:getMillis(),logs:logs})
}


func (c *TracingSpan) ErrorOccurred()  {
	c.Lock()
	defer c.Unlock()
	c.errorOccurred=true
}



func (c *TracingSpan) Ref(traceSegmentRef TraceSegmentRef) {
	c.Lock()
	defer c.Unlock()
	if !TraceSegmentRefContains(traceSegmentRef, c.refs) {
		c.refs = append(c.refs, traceSegmentRef)
	}
}




func (c *TracingSpan) Finish(owner TraceSegment) bool {
	c.Lock()
	c.endTime = getMillis()
	owner.Archive(*c)
	return true
}



func (c *TracingSpan) setTransform() *SpanRequest {
	spanRequest := &SpanRequest{
		spanId:       c.spanId,
		parentSpanId: c.parentSpanId,
		startTime:    c.startTime,
		endTime:      c.endTime,
	}
	if c.operationId != 0 {
		spanRequest.operationName = string(c.operationId)
	} else {
		spanRequest.operationName = c.operationName
	}

	if c.IsEntry {
		spanRequest.spanType = 0
	} else if c.IsEntry {
		spanRequest.spanType = 1
	} else {
		spanRequest.spanType = 2
	}

	if c.layer != nil {
		spanRequest.spanLayer = int32(*c.layer)
	}

	if c.componentId != 0 {
		spanRequest.component = string(c.componentId)
	} else {
		if len(c.componentName) > 0 {
			spanRequest.component = c.componentName
		}
	}

	spanRequest.isError = c.errorOccurred

	spanRequest.tags = c.tags

	for _, log := range c.logs {
		spanRequest.logs = append(spanRequest.logs, *log.Transform())
	}

	if len(c.refs) <= 0 {
		return spanRequest
	}
	for _, ref := range c.refs {
		spanRequest.references = append(spanRequest.references, *ref.Transform())
	}

	return spanRequest

}


func (c *TracingSpan) Transform() *SpanRequest {
	c.Lock()
	defer c.Unlock()
	if c.spanType==EntrySpan {
		return c.setTransform()
	}else if c.spanType==ExitSpan {
		spanRequest :=c.setTransform()
		if c.peerId==0 {
			spanRequest.peer=c.peer
		}else{
			spanRequest.peerId=c.peerId
		}
		return  spanRequest
	}else{
		return c.setTransform()
	}
}

