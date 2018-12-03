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

import (
"strconv"

"github.com/WithLin/skywalking-go/config"
)

type SegmentRefType int32

const (
	CrossProcess SegmentRefType = iota
	CrossThread  SegmentRefType = 1
)

type TraceSegmentRef struct {
	refType                     SegmentRefType
	traceSegmentId              *ID
	spanId                      int32
	peerId                      int32
	peerHost                    string
	entryApplicationInstanceId  int32
	parentApplicationInstanceId int32
	entryOperationName          string
	entryOperationId            int32
	parentOperationName         string
	parentOperationId           int32
}

func NewTraceSegmentRefByContextCarrier(carrier ContextCarrier) (*TraceSegmentRef, error) {
	t := &TraceSegmentRef{}
	t.refType = CrossProcess
	t.traceSegmentId = carrier.traceSegmentId
	t.spanId = carrier.spanId
	t.parentApplicationInstanceId = carrier.parentApplicationInstanceId
	t.entryApplicationInstanceId = carrier.entryApplicationInstanceId
	host := []rune(carrier.peerHost)
	if host[0] == '#' {
		t.peerHost = string(host[1 : len(host)-2])
	} else {
		id, err := strconv.Atoi(carrier.peerHost)
		if err != nil {
			return nil, err
		}
		t.peerId = int32(id)
	}
	entryOperationName := []rune(carrier.entryOperationName)
	if entryOperationName[0] == '#' {
		t.entryOperationName = string(entryOperationName[1 : len(entryOperationName)-2])
	} else {
		entryOperationId, err := strconv.Atoi(carrier.entryOperationName)
		if err != nil {
			return nil, err
		}
		t.entryOperationId = int32(entryOperationId)
	}

	return t, nil

}

func NewTraceSegmentRefByContextSnapshot(contextSnapshot ContextSnapshot) (*TraceSegmentRef, error) {
	c := &TraceSegmentRef{}
	c.refType = CrossThread
	c.traceSegmentId = contextSnapshot.traceSegmentId
	c.spanId = contextSnapshot.spanId
	c.parentApplicationInstanceId = config.Conf.ApplicationInstanceId
	c.entryApplicationInstanceId = contextSnapshot.entryApplicationInstanceId
	entryOperationName := []rune(contextSnapshot.entryOperationName)

	if entryOperationName[0] == '#' {
		c.entryOperationName = string(entryOperationName[1 : len(entryOperationName)-2])
	} else {
		entryOperationId, err := strconv.Atoi(contextSnapshot.entryOperationName)
		if err != nil {
			return nil, err
		}
		c.entryOperationId = int32(entryOperationId)
	}

	parentOperationName := []rune(contextSnapshot.parentOperationName)

	if parentOperationName[0] == '#' {
		c.parentOperationName = string(parentOperationName[1 : len(parentOperationName)-2])
	} else {
		parentOperationId, err := strconv.Atoi(contextSnapshot.entryOperationName)
		if err != nil {
			return nil, err
		}
		c.parentOperationId =int32(parentOperationId)
	}
	return c, nil
}

type TraceSegmentReference struct {
	parentTraceSegmentId        *UniqueIdRequest
	parentApplicationInstanceId int32
	parentSpanId                int32
	entryApplicationInstanceId  int32
	refType                     int32
	parentServiceName           string
	entryServiceName            string
	networkAddress              string
	networkAddressId            int32
	entryServiceId              int32
	parentServiceId             int32
}

func (c *TraceSegmentRef) Transform() *TraceSegmentReference {
	t := &TraceSegmentReference{}
	if c.refType == CrossProcess {
		t.refType = int32(CrossProcess)
		if c.peerId == 0 {
			t.networkAddress = c.peerHost
		} else {
			t.networkAddressId = c.peerId
		}
	} else {
		t.refType = int32(CrossThread)
	}

	t.parentApplicationInstanceId = c.parentApplicationInstanceId
	t.entryApplicationInstanceId = c.entryApplicationInstanceId
	t.parentSpanId = c.spanId

	if c.entryOperationId == 0 {
		if len(c.entryOperationName) > 0 {
			t.entryServiceName = c.entryOperationName
		}
	} else {
		t.entryServiceId = c.entryOperationId
	}
	if c.parentApplicationInstanceId == 0 {
		if len(c.parentOperationName) > 0 {
			t.parentServiceName = c.parentOperationName
		}

	} else {
		t.parentServiceId = c.parentOperationId
	}

	return t
}
