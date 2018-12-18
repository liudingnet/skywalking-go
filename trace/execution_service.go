package trace

import (
	"github.com/WithLin/skywalking-go/config"
	"os"
)

func DoExcuteService(){

}



func excuteSamplingRefreshService(){
	GetInstance().reset()
}

func RegisterApplication(){
	if !(config.Conf.ApplicationId ==0 ) {
		client :=GetGrpcClientInstance()
		client.RegisterApplication(config.Conf.ApplicationCode)

	}
}

func RegisterApplicationInstance(){
	if config.Conf.RuntimeEnvironment.ApplicationId >0 && !(config.Conf.RuntimeEnvironment.ApplicationInstanceId>0) {
		osInfoRequest :=&AgentOsInfoRequest{
			hostName:getHostName(),
			ipAddress:getIpV4(),
			osName:getOsName(),
			processNo:int32(os.Getegid()),
		}
		client :=GetGrpcClientInstance()


		c :=config.Conf.RuntimeEnvironment
		client.RegisterApplicationInstance(c.ApplicationInstanceId,c.AgentUUID,getMillis(),*osInfoRequest)
	}
}

func Heartbeat(){
	client :=GetGrpcClientInstance()
	//
	//c :=config.Conf.RuntimeEnvironment
	client.Heartbeat()
}

func TraceSegmentTransportService(){
	c :=newtraceingDispatcher()
	c.Flush()
}

func afterFinished(traceSegment TraceSegment){
	if !traceSegment.isIgnore {
		c :=newtraceingDispatcher()
		c.Dispatch(traceSegment.Transform())
	}
}




