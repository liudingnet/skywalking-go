/*
 * Licensed to the OpenSkywalking under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
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

import (
	"fmt"
	"sync"
)

type TraceSegment struct {
	sync.RWMutex
	traceSegmentId        *ID
	refs                  []TraceSegmentRef
	spans                 []TracingSpan
	relatedGlobalTraces   *DistributedTraceIds
	applicationId         int32
	applicationInstanceId int32
	hasRef                bool
	isIgnore              bool
	isSingleSpanSegment   bool
	isSizeLimited         bool
}

type TraceSegmentRequest struct {
	uniqueIds []*UniqueIdRequest
	Segment   *TraceSegmentObjectRequest
}

type TraceSegmentObjectRequest struct {
	segmentId             *UniqueIdRequest
	applicationId         int32
	applicationInstanceId int32
	spans                 []SpanRequest
}

type SpanRequest struct {
	spanId          int32
	spanType        int32
	spanLayer       int32
	parentSpanId    int32
	componentId     int32
	operationNameId int32
	startTime       int64
	endTime         int64
	component       string
	operationName   string
	peer            string
	isError         bool
	references      []TraceSegmentReference
	tags            map[string]string
	logs            []LogMessage
	peerId          int32
}

func NewTraceSegment() *TraceSegment {

	var relatedGlobalTraces []*DistributedTraceId
	var spans []TracingSpan
	var refs []TraceSegmentRef
	globalId := new(GlobalIdGenerator)
	globalId.threadIdSequence = newIDContext(getMillis(), 0)
	t := &TraceSegment{
		traceSegmentId:      globalId.Generate(),
		relatedGlobalTraces: &DistributedTraceIds{relatedGlobalTraces: relatedGlobalTraces},
		spans:               spans,
		refs:                refs,
	}
	t.relatedGlobalTraces.append(&DistributedTraceId{})

	return t

}

func (c *TraceSegment) Archive(finishedSpan TracingSpan) {
	c.spans = append(c.spans, finishedSpan)
}

func (c *TraceSegment) Finish(isSizeLimited bool) {
	c.isSizeLimited = isSizeLimited
}

func (c *TraceSegment) Ref(refSegment *TraceSegmentRef) {
	if !TraceSegmentRefContains(*refSegment, c.refs) {
		c.refs = append(c.refs, *refSegment)
	}
}

func TraceSegmentRefContains(a TraceSegmentRef, list []TraceSegmentRef) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (c *TraceSegment) RelatedGlobalTrace(distributedTraceId *DistributedTraceId) {
	c.relatedGlobalTraces.append(distributedTraceId)
}

func (c *TraceSegment) Transform() *TraceSegmentRequest {
	var uids []*UniqueIdRequest
	d := c.relatedGlobalTraces.getRelatedGlobalTraces()
	for _, id := range d {
		uids = append(uids, id.ToUniqueId())
	}
	upstreamSegment := &TraceSegmentRequest{
		uniqueIds: uids,
	}

	var sp []SpanRequest
	for _, span := range c.spans {
		sp = append(sp, *span.Transform())
	}
	upstreamSegment.Segment = &TraceSegmentObjectRequest{
		segmentId:             c.traceSegmentId.Transform(),
		spans:                 sp,
		applicationId:         c.applicationId,
		applicationInstanceId: c.applicationInstanceId,
	}

	return upstreamSegment
}

func (c *TraceSegment) ToString() string {

	s := fmt.Sprintf("TraceSegment{traceSegmentId=%+v, refs=%+v,spans=%+v,relatedGlobalTraces=%+v}",
		c.traceSegmentId,
		c.refs,
		c.spans,
		c.relatedGlobalTraces)
	return s
}
