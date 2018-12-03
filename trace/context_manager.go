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
	"github.com/WithLin/skywalking-go/config"
)

type ctxKey struct{}

var swSpanKey = ctxKey{}

// Create or get the existed SkyWalking context from go context.
func GetOrCreateContext(ctx context.Context,operationName string,forceSampling bool)(context.Context,TracerContext){
	var cont context.Context
	var swCtx TracerContext
	var ok bool
	if swCtx,ok =ctx.Value(swSpanKey).(TracerContext); ok {
		if len(operationName) <=0 {
			cont=context.WithValue(ctx,swSpanKey,&IgnoredTracerContext{})
		}else {
			//todo nitialized => ApplicationId.HasValue && ApplicationInstanceId.HasValue
			if config.Conf.ApplicationInstanceId >0 {
				sampler :=new(defaultSampler)
				if forceSampling || sampler.sampled() {
					cont=context.WithValue(ctx,swSpanKey,&TracingContext{})
				}else {
					cont=context.WithValue(ctx,swSpanKey,&TracingContext{})
				}
			}else {
				cont=context.WithValue(ctx,swSpanKey,&TracingContext{})
			}
		}
	}
	return cont,swCtx
}



func CreateEntrySpan(ctx context.Context,operationName string,carrier *ContextCarrier) Span{
	samplingService :=new(defaultSampler)
	if carrier !=nil && carrier.IsValid() {
		samplingService.forceSampled()
		_,traceContext:=GetOrCreateContext(ctx,operationName,true)
		span:=traceContext.createEntrySpan(operationName)
		traceContext.extract(*carrier)
		return  span
	}else {
		_,traceContext:=GetOrCreateContext(ctx,operationName,false)
		return traceContext.createEntrySpan(operationName)
	}
}


func CreateLocalSpan(ctx context.Context,operationName string)Span {
	_, traceContext := GetOrCreateContext(ctx, operationName, false)
	return traceContext.createLocalSpan(operationName)
}


func CreateExitSpan(ctx context.Context,operationName string,carrier *ContextCarrier,remotePeer string)Span {
	_, traceContext := GetOrCreateContext(ctx, operationName, false)
	span:= traceContext.createExitSpan(operationName,remotePeer)

	traceContext.inject(*carrier)
	return  span
}

func Context(ctx context.Context) *TracingContext{
	return ctx.Value(swSpanKey).(*TracingContext)
}

func Inject(ctx context.Context,carrier ContextCarrier){
	if Context(ctx) != nil {
		Context(ctx).inject(carrier)
	}
}


func Extract(ctx context.Context,carrier ContextCarrier){
	if Context(ctx) != nil {
		Context(ctx).extract(carrier)
	}
}

func Continued(ctx context.Context,snapshot ContextSnapshot){
	if snapshot.IsValid() && !snapshot.IsFromCurrent() {
		if Context(ctx) != nil {
			Context(ctx).continued(snapshot)
		}
	}
}


func StopSpan(ctx context.Context){
	if Context(ctx) != nil {
	    Context(ctx).stopSpan(ActiveSpan(ctx))
    }

}



func ActiveSpan(ctx context.Context) *TracingSpan{
	if Context(ctx) != nil {
	 return 	Context(ctx).activeSpanStacks.peek().(*TracingSpan)
	}
	return  nil

}







