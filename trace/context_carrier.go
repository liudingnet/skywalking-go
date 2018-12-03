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
"fmt"
"strconv"
"strings"
)

type ContextCarrier struct {
	traceSegmentId              *ID
	spanId                      int32
	parentApplicationInstanceId int32
	entryApplicationInstanceId  int32
	peerHost                    string
	peerId                      int32
	entryOperationName          string
	entryOperationId            int32
	parentOperationName         string
	namespace                   string
	primaryDistributedTraceId   *DistributedTraceId
}

func NewContextCarrierWithNameSpace(namespace string) *ContextCarrier {
	c := &ContextCarrier{
		namespace: namespace,
	}
	return c
}

func (c *ContextCarrier) SetEntryOperationName(entryOperationName string) {
	c.entryOperationName = "#" + entryOperationName
}

func (c *ContextCarrier) SetParentOperationName(parentOperationName string) {
	c.parentOperationName = "#" + parentOperationName
}

func (c *ContextCarrier) SetPeerHost(peerHost string) {
	c.peerHost = "#" + peerHost
}

func (c *ContextCarrier) IsValid() bool {
	flag := c.traceSegmentId != nil && c.spanId > -1 && c.parentApplicationInstanceId != 0 && c.entryApplicationInstanceId != 0 && len(c.peerHost) > 0 && len(c.parentOperationName) > 0 && len(c.entryOperationName) > 0 && c.primaryDistributedTraceId != nil
	return flag
}


func (c *ContextCarrier) Deserialize(text string) (*ContextCarrier, error) {
	parts := strings.SplitN(text, "|", 8)
	spanid, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	parentApplicationInstanceId, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}
	entryApplicationInstanceId, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}

	if len(parts) == 8 {
		id, err := newID(parts[0])
		if err != nil {
			return nil, err
		}
		c.traceSegmentId = id
		c.spanId = int32(spanid)
		c.parentApplicationInstanceId = int32(parentApplicationInstanceId)
		c.entryApplicationInstanceId = int32(entryApplicationInstanceId)
		c.peerHost = parts[4]
		c.entryOperationName = parts[5]
		c.parentOperationName = parts[6]
		prid, err := newDistributedTraceId(parts[7])
		if err != nil {
			return nil, err
		}
		c.primaryDistributedTraceId = prid
	}

	return c, nil
}

func (c *ContextCarrier) Serialize() *string {
	flag := c.IsValid()
	if !flag {
		return nil
	}

	s := fmt.Sprintf("%s|%d|%d|%d|%s|%s|%s|%s",
		c.traceSegmentId.encoding,
		c.spanId,
		c.parentApplicationInstanceId,
		c.entryApplicationInstanceId,
		c.peerHost,
		c.entryOperationName,
		c.parentOperationName,
		c.primaryDistributedTraceId.encoding,
	)
	return &s

}

func (c *ContextCarrier) GetAllItems() *CarrierItem {
	carrierItem := newSW3CarrierItem(*c, nil, c.namespace)
	head := newCarrierItemHead(carrierItem.carrier, c.namespace)
	return &head.carrierItem
}

