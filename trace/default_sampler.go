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
	"sync"
	"sync/atomic"
)

type Sampler interface {
	sampled() bool
	forceSampled()
	reset()
}

type defaultSampler struct {
	sync.RWMutex
	samplePer3Secs int
	sampleOn bool
}

var _ Sampler=(*defaultSampler)(nil)
var ops uint64

var instance *defaultSampler
var once sync.Once

func GetInstance() *defaultSampler{
	once.Do(func() {
		instance=&defaultSampler{}
	})
	return instance
}

func(c *defaultSampler) sampled() bool{
	c.Lock()
	defer c.Unlock()
	if !c.sampleOn {
		return  true
	}
	return  atomic.AddUint64(&ops,1) <uint64(c.samplePer3Secs)

}

func (c *defaultSampler)setSamplePer3Secs(samplePer3Secs int){
	c.samplePer3Secs=samplePer3Secs
	c.sampleOn=samplePer3Secs>-1
}


func(c *defaultSampler) forceSampled(){
	c.Lock()
	defer c.Unlock()
	if c.sampleOn {
		atomic.AddUint64(&ops,1)
	}
}

func(c *defaultSampler) reset(){
	c.Lock()
	defer c.Unlock()
	ops=0
}
