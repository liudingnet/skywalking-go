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


type NoopSpan struct {
	spanId        int32
	operationName string
	operationId   int32
	isEntry       bool
	isExit        bool
}

var _ Span =(*NoopEntrySpan)(nil)


func (c *NoopSpan) ErrorOccurred() {
	return
}

func (c *NoopSpan) Log(err error) {
	return
}



func (c *NoopSpan) Ref(traceSegmentRef TraceSegmentRef) {
	return
}

func (c *NoopSpan) SetComponent(componentName string) {
	return
}

func (c *NoopSpan) SetLayer(layer *SpanLayer) {
	return
}

func (c *NoopSpan) Start(timestamp int64) {
	return
}

func (c *NoopSpan) Tag(key string, value string) {
	return
}


func (c *NoopSpan) SetOperationName(operationName string) {

}


func (c *NoopSpan) SetOperationId(operationId  int32) {

}


type NoopEntrySpan struct {
	NoopSpan
}

var _ Span =(*NoopEntrySpan)(nil)


func NewNoopEntrySpan() *NoopEntrySpan {
	noopEntrySpan := new(NoopEntrySpan)
	noopEntrySpan.spanId = 0
	noopEntrySpan.operationName = ""
	noopEntrySpan.operationId = 0
	noopEntrySpan.isEntry = true
	noopEntrySpan.isExit = false
	return noopEntrySpan
}

type NoopExitSpan struct {
	peerId int
	peer   string
	NoopSpan
}

func NewNoopExitSpan(peerId int, peer string) *NoopExitSpan {
	noopExitSpan := new(NoopExitSpan)
	if len(peer) <= 0 {
		noopExitSpan.peerId = peerId
	} else {
		noopExitSpan.peer = peer
	}
	noopExitSpan.isExit = true
	return noopExitSpan
}
