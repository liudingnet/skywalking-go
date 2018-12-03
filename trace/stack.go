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

import "sync"

type element struct {
	data interface{}
	next *element
}

type stack struct {
	sync.Mutex
	head *element
	first *element
	Size int
}

func (c *stack) push(data interface{}) {
	c.Lock()
	defer  c.Unlock()
	element := new(element)
	element.data = data
	temp := c.head
	element.next = temp
	c.head = element

	c.Size++
	if c.Size ==1 {
		c.first=element
	}

}

func (c *stack) pop() interface{} {
	c.Lock()
	defer  c.Unlock()
	if c.head == nil {
		return nil
	}

	r := c.head.data
	c.head = c.head.next
	c.Size--

	if c.Size ==0 {
		c.first=nil
	}

	return r
}

func (c *stack) peek() interface{}{

	c.Lock()
	defer c.Unlock()

	if c.head ==nil {
		return nil
	}

	return c.head.data

}

func (c *stack) getLast() interface{}{
	c.Lock()
	defer c.Unlock()
	if c.first == nil{
		return nil
	}
	return c.first.data
}


