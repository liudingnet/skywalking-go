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
	"context"
	"fmt"
	"github.com/WithLin/skywalking-go/config"
	pb "github.com/WithLin/skywalking-go/proto"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

type AgentOsInfoRequest struct {
	osName string
	hostName string
	processNo int32
	ipAddress []string
}


type SkyWalkingClient interface {
	RegisterApplication(applicationCode string)
	RegisterApplicationInstance(applicationId int,agentUUID string,registerTime int64,
		osInfoRequest AgentOsInfoRequest)
	Heartbeat(applicationInstance int,heartbeatTime int64)
	Collect(request []TraceSegmentRequest)
}


type GrpcClient struct {
	conn  *grpc.ClientConn
}


var client *GrpcClient
var one sync.Once


func GetGrpcClientInstance() *GrpcClient{
	one.Do(func() {
		var err error
		client,err=newGrpcClient()
		if err !=nil {
			log.Println(fmt.Sprintf("GetInstance GrpcClient err: ---->%s",err))
		}
	})
	return client
}


func newGrpcClient()(*GrpcClient,error){
	client :=new(GrpcClient)
	conn,err :=grpc.Dial(config.Conf.Servers,grpc.WithInsecure())
	if err !=nil {
		return  nil,err
	}
	client.conn=conn
	return client,nil

}



func (c *GrpcClient)RegisterApplication(applicationCode string) {
	client := pb.NewApplicationRegisterServiceClient(c.conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer  cancel()
	client.ApplicationCodeRegister(context.Background(),&pb.Application{ApplicationCode:applicationCode})
}


func (c *GrpcClient)RegisterApplicationInstance(applicationId int32,agentUUID string,registerTime int64,
osInfoRequest AgentOsInfoRequest){
	applicationInstance :=&pb.ApplicationInstance{
		ApplicationId:applicationId,
		AgentUUID:agentUUID,
		RegisterTime:registerTime,
		Osinfo:&pb.OSInfo{
			OsName:osInfoRequest.osName,
			Hostname:osInfoRequest.hostName,
			ProcessNo:osInfoRequest.processNo,
	},
	}

	applicationInstance.Osinfo.Ipv4S=append(applicationInstance.Osinfo.Ipv4S,osInfoRequest.ipAddress...)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer  cancel()

	client :=pb.NewInstanceDiscoveryServiceClient(c.conn)

	client.RegisterInstance(context.Background(),applicationInstance)

}

func (c *GrpcClient)Heartbeat(applicationInstance int32,heartbeatTime int64){
	heartbeat :=&pb.ApplicationInstanceHeartbeat{
		ApplicationInstanceId:applicationInstance,
		HeartbeatTime:heartbeatTime,
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer  cancel()

	client :=pb.NewInstanceDiscoveryServiceClient(c.conn)
	client.Heartbeat(context.Background(),heartbeat)
}


func (c *GrpcClient)Collect(requests []TraceSegmentRequest){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer  cancel()
	client :=pb.NewTraceSegmentServiceClient(c.conn)

	asyncClientStreamingCall,err:=client.Collect(ctx)
	if err  != nil {
		log.Println(fmt.Sprintf("Collect  err :---> %s",err))
	}

	for _, request := range requests {

		asyncClientStreamingCall.Send(traceSegmentRequestMap(request))
	}


}

func traceSegmentRequestMap(request TraceSegmentRequest) *pb.UpstreamSegment{
	upstreamSegment :=new(pb.UpstreamSegment)
	upstreamSegment.GlobalTraceIds=append(upstreamSegment.GlobalTraceIds,selectUniqueIds(request.uniqueIds)...)

	traceSegment :=&pb.TraceSegmentObject{
		TraceSegmentId:mapToUniqueId(request.Segment.segmentId),
		ApplicationId:request.Segment.applicationId,
		ApplicationInstanceId:request.Segment.applicationInstanceId,
		IsSizeLimited:false,
	}

	for _, span := range request.Segment.spans {
		traceSegment.Spans=append(traceSegment.Spans,mapToSpan(span))
	}


	segmentByte,err :=traceSegment.Marshal()
	if err !=nil {
		if err != nil {
			log.Println(fmt.Sprintf("traceSegment Marshal err :---> %s",err))
		}
	}
	upstreamSegment.Segment=segmentByte
	return upstreamSegment

}

func selectUniqueIds(request []*UniqueIdRequest) []*pb.UniqueId{
	var uniqueId  []*pb.UniqueId
	for _, uniqueidValue := range request {
		uniqueId=append(uniqueId,mapToUniqueId(uniqueidValue))
	}
	return  uniqueId
}


func mapToUniqueId(uniqueIdRequest *UniqueIdRequest) *pb.UniqueId{
	uniqueId := &pb.UniqueId{
	}
	uniqueId.IdParts=append(uniqueId.IdParts,uniqueIdRequest.part1)
	uniqueId.IdParts=append(uniqueId.IdParts,uniqueIdRequest.part2)
	uniqueId.IdParts=append(uniqueId.IdParts,uniqueIdRequest.part3)
	return uniqueId
}


func mapToSpan(request SpanRequest) *pb.SpanObject{
	spanObject :=&pb.SpanObject{
		SpanId:request.spanId,
		ParentSpanId:request.parentSpanId,
		StartTime:request.startTime,
		EndTime:request.endTime,
		SpanType:pb.SpanType(request.spanType),
		SpanLayer:pb.SpanLayer(request.spanLayer),
		IsError:request.isError,
	}

	if len(request.component) >0 {
		spanObject.Component=request.component
	}else{
		spanObject.ComponentId=request.componentId
	}

	if len(request.operationName)>0 {
		spanObject.OperationName=request.operationName
	}else{
		spanObject.OperationNameId=request.operationNameId
	}

	if len(request.peer) >0 {
		spanObject.Peer=request.peer
	}else{
		spanObject.PeerId=request.peerId
	}

	for key, value := range request.tags {
		mk :=new(pb.KeyWithStringValue)
		mk.Value=value
		mk.Key=key
		spanObject.Tags=append(spanObject.Tags,mk)
	}


	for _, referenceRequest := range request.references {
		ts := new(pb.TraceSegmentReference)
		ts.ParentApplicationInstanceId=referenceRequest.parentApplicationInstanceId
		ts.EntryApplicationInstanceId=referenceRequest.entryApplicationInstanceId
		ts.ParentSpanId=referenceRequest.parentSpanId
		ts.RefType=pb.RefType(referenceRequest.refType)
		ts.ParentTraceSegmentId=mapToUniqueId(referenceRequest.parentTraceSegmentId)

		if len(referenceRequest.networkAddress)>0 {
			ts.NetworkAddress=referenceRequest.networkAddress
		}else {
			ts.NetworkAddressId=referenceRequest.networkAddressId
		}

		if len(referenceRequest.entryServiceName) >0 {
			ts.EntryServiceName=referenceRequest.entryServiceName
		}else {
			ts.EntryServiceId=referenceRequest.entryServiceId
		}

		if len(referenceRequest.parentServiceName) >0 {
			ts.ParentServiceName=referenceRequest.parentServiceName
		}else{
			ts.ParentServiceId=referenceRequest.parentServiceId
		}

		spanObject.Refs=append(spanObject.Refs,ts)
	}

	for _, log := range request.logs {
		logMessage :=new(pb.LogMessage)
		ky :=new(pb.KeyWithStringValue)
		var kys []*pb.KeyWithStringValue
		logMessage.Time=log.timestamp
		for key, value := range log.data {
			ky.Key=key
			ky.Value=value
			kys=append(kys,ky)
		}
		logMessage.Data=kys
		spanObject.Logs=append(spanObject.Logs,logMessage)
	}

	return spanObject

}