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

import "fmt"

type ContextSnapshot struct {
	traceSegmentId             *ID
	spanId                     int32
	entryOperationName         string
	parentOperationName        string
	primaryDistributedTraceId  *DistributedTraceId
	entryApplicationInstanceId int32
}

func NewContextSnapshot(traceSegmentId *ID, spanId int32, distributedTraceIds []*DistributedTraceId) *ContextSnapshot {
	c := &ContextSnapshot{}
	c.traceSegmentId = traceSegmentId
	c.spanId = spanId
	c.primaryDistributedTraceId = distributedTraceIds[0]
	return c
}

func (c *ContextSnapshot) SetEntryOperationName(entryOperationName string) {
	c.entryOperationName = "#" + entryOperationName
}

func (c *ContextSnapshot) SetEntryOperationId(entryOperationId int32) {
	c.entryOperationName = fmt.Sprintf("%d%s", entryOperationId, "")
}

func (c *ContextSnapshot) SetParentOperationName(parentOperationName string) {
	c.parentOperationName = "#" + parentOperationName
}

func (c *ContextSnapshot) SetParentOperationId(parentOperationId int32) {
	c.parentOperationName = fmt.Sprintf("%d%s", parentOperationId, "")
}

func (c *ContextSnapshot) IsValid() bool {
	return c.traceSegmentId != nil && c.spanId > -1 && c.entryApplicationInstanceId != 0 && c.primaryDistributedTraceId != nil
}

func (c *ContextSnapshot) SetEntryApplicationInstanceId(entryApplicationInstanceId int32) {
	c.entryApplicationInstanceId = entryApplicationInstanceId
}

func (c *ContextSnapshot) IsFromCurrent() bool {
	// _traceSegmentId.Equals(ContextManager.Capture.TraceSegmentId)
	return false
}
