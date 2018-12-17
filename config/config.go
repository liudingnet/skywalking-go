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

package config

//config   agent config
type config struct {
	Namespace  string
	ApplicationCode string
	SpanLimitPerSegment int32
	SamplePer3Secs int32
	PendingSegmentLimit int32
	Interval int32
	PendingSegmentTimeout int32
	ApplicationInstanceId int32
	GrpcConfig
	RuntimeEnvironment
}

type GrpcConfig struct {
	Servers string
	ConnectTimeout int32
	Timeout int32
}

type RuntimeEnvironment struct {
	ApplicationId int32
	ApplicationInstanceId int32
	AgentUUID string

}

var Conf  config


func init(){
	Conf.Servers="localhost:11800"
	Conf.ApplicationId=1233
	Conf.ApplicationInstanceId=123232
	Conf.PendingSegmentLimit=30000
	Conf.Interval=3000
	Conf.PendingSegmentTimeout=1000
	Conf.SamplePer3Secs=-1
	Conf.SpanLimitPerSegment=300
}







