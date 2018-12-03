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

type CarrierItem struct {
	headKey   string
	headValue string
	next      *CarrierItem
}

func newCarrierItem(headKey string, headValue string, next *CarrierItem, namespace string) *CarrierItem {
	c := &CarrierItem{}
	if len(namespace) > 0 {
		s := fmt.Sprintf("%s-%s", namespace, headKey)
		c.headKey = s
	} else {
		c.headKey = headKey
	}
	c.headValue = headValue
	c.next = next
	return c
}

type SW3CarrierItem struct {
	carrier        *CarrierItem
	contextCarrier ContextCarrier
}

func newSW3CarrierItem(carrier ContextCarrier, next *CarrierItem, namespace string) *SW3CarrierItem {
	sw := &SW3CarrierItem{}

	sw.carrier = newCarrierItem("sw3", *carrier.Serialize(), next, namespace)
	sw.carrier = next
	return sw
}

func (c *SW3CarrierItem) getHeadValue() string {
	return c.carrier.headValue
}

func (c *SW3CarrierItem) setHeadValue(value string) error {
	_, err := c.contextCarrier.Deserialize(value)
	if err != nil {
		return err
	}
	return nil
}

type CarrierItemHead struct {
	carrierItem CarrierItem
}

func newCarrierItemHead(next *CarrierItem, namespace string) *CarrierItemHead {
	c := &CarrierItemHead{}
	c.carrierItem = *newCarrierItem("", "", next, namespace)
	return c
}
