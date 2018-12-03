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
"bytes"
"fmt"
"math/rand"
"reflect"
"runtime"
"strconv"
"strings"
"time"

"github.com/WithLin/skywalking-go/config"
)

type ID struct {
	part1    int64
	part2    int64
	part3    int64
	encoding string
	isValid  bool
}

//newID new a unique id
func newID(encodingString string) (*ID, error) {
	idParts := strings.SplitN(encodingString, "\\.", 3)

	id := &ID{}
	id.isValid = true
	for part := 0; part < 3; part++ {
		if part == 0 {
			i, err := strconv.ParseInt(idParts[part], 10, 64)
			if err != nil {
				id.isValid = false
				return id, err
			}
			id.part1 = i
		} else if part == 1 {
			i, err := strconv.ParseInt(idParts[part], 10, 64)
			if err != nil {
				id.isValid = false
				return id, err
			}
			id.part2 = i
		} else {
			i, err := strconv.ParseInt(idParts[part], 10, 64)
			if err != nil {
				id.isValid = false
				return id, err
			}
			id.part3 = i
		}
	}
	return id, nil

}

//Encode encode id toString
func (c *ID) Encode() string {
	if len(c.encoding) == 0 {
		s := fmt.Sprintf("%d.%d.%d", c.part1, c.part2, c.part3)
		c.encoding = s
	}
	return c.encoding
}

//IsValid Verify that it is valid
func (c *ID) IsValid() bool {
	return c.isValid
}

//Equals Equals
func (c *ID) Equals(other ID) bool {
	return reflect.DeepEqual(c, other)
}

//ToString format id
func (c *ID) ToString() string {
	s := fmt.Sprintf("%d.%d.%d", c.part1, c.part2, c.part3)
	return s
}

type UniqueIdRequest struct {
	part1 int64
	part2 int64
	part3 int64
}

//Transform
func (c *ID) Transform() *UniqueIdRequest {
	return &UniqueIdRequest{
		part1:c.part1,
		part2:c.part2,
		part3:c.part3,
	}
}

type DistributedTraceId struct {
	id       ID
	encoding string
}

//newDistributedTraceId new a DistributedTraceId instance
func newDistributedTraceId(id string) (*DistributedTraceId, error) {
	Id, err := newID(id)
	if err != nil {
		return nil, err
	}
	d := &DistributedTraceId{
		id:       *Id,
		encoding: Id.encoding,
	}
	return d, nil
}


//ToUniqueId generate a unique id
func (c *DistributedTraceId) ToUniqueId() *UniqueIdRequest {
	return c.id.Transform()
}

//ToString  id to string
func (c *DistributedTraceId) ToString() string {
	return c.id.ToString()
}

type PropagatedTraceId struct {
	DistributedTraceId
}

type DistributedTraceIds struct {
	relatedGlobalTraces []*DistributedTraceId
}

//getRelatedGlobalTraces get RelatedGlobalTraces
func (c *DistributedTraceIds) getRelatedGlobalTraces() []*DistributedTraceId {
	return c.relatedGlobalTraces
}

//append append distributedTraceId
func (c *DistributedTraceIds) append(distributedTraceId *DistributedTraceId) {
	if len(c.relatedGlobalTraces) > 0 {
		c.relatedGlobalTraces = append(c.relatedGlobalTraces[:0], c.relatedGlobalTraces[1:]...)
	}
	if !contains(c.relatedGlobalTraces, distributedTraceId) {
		c.relatedGlobalTraces = append(c.relatedGlobalTraces, distributedTraceId)
	}

}

//contains  determine if it is inside the current array
func contains(s []*DistributedTraceId, e *DistributedTraceId) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type GlobalIdGenerator struct {
	threadIdSequence *IDContext
}


//Generate generate a id
func (c *GlobalIdGenerator) Generate() *ID {
	if config.Conf.ApplicationInstanceId == 0 {
		panic("InvalidOperationException")
	}

	return &ID{
		part1:int64(config.Conf.ApplicationInstanceId),
		part2:getID(),
		part3:c.threadIdSequence.NextSeq(),
		encoding:"",
		isValid:true,
	}

}

//get goroutine id  todo:You should not do this, but follow the concept of skywalking
func getID() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseInt(string(b), 10, 64)
	return n
}

type IDContext struct {
	lastTimestamp      int64
	threadSeq          int
	runRandomTimestamp int64
	lastRandomValue    int
}

//newIDContext new a IDContext
func newIDContext(lastTimestamp int64, threadSeq int) *IDContext {
	context := &IDContext{
		lastTimestamp: lastTimestamp,
		threadSeq:     threadSeq,
	}
	return context
}

//getTimeStamp from IDContext get timestamp
func (c *IDContext) getTimeStamp() int64 {
	currentTimeMillis := getMillis()
	rand.Seed(time.Now().UnixNano())
	if currentTimeMillis < c.lastTimestamp {
		if c.runRandomTimestamp != currentTimeMillis {
			c.lastRandomValue = rand.Int()
			c.runRandomTimestamp = currentTimeMillis

		}
		return int64(c.lastRandomValue)
	} else {
		c.lastTimestamp = currentTimeMillis
		return c.lastTimestamp
	}
}

//nextThreadSeq
func (c *IDContext) nextThreadSeq() int {
	if c.threadSeq == 10000 {
		c.threadSeq = 0
	}
	return c.threadSeq + 1
}

//NextSeq
func (c *IDContext) NextSeq() int64 {
	return c.getTimeStamp()*10000 + int64(c.nextThreadSeq())
}


