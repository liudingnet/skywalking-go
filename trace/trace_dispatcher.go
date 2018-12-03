package trace

import (
	"fmt"
	"github.com/WithLin/skywalking-go/config"
)

import (
	"log"
)

type traceDispatcher interface {
	Dispatch(segment *TraceSegmentRequest) bool
	Flush()
	Close()
}



type traceingDispatcher struct {
	segment chan *TraceSegmentRequest
}

var _ traceDispatcher =(*traceingDispatcher)(nil)

func newtraceingDispatcher() *traceingDispatcher{
	c :=&traceingDispatcher{
	}

	c.segment=make(chan *TraceSegmentRequest,config.Conf.PendingSegmentLimit)

	return  c

}

func (c *traceingDispatcher) Dispatch(segment *TraceSegmentRequest) bool{
	if config.Conf.PendingSegmentLimit < int32(len(c.segment)) {
		return  false
	}
	c.segment<-segment


	log.Println(fmt.Sprintln("Dispatch trace segment. [SegmentId] = %d",segment.Segment.segmentId))

	return  true

}

func (c *traceingDispatcher) Flush(){
	//limit :=config.Conf.PendingSegmentLimit
	var segments []TraceSegmentRequest
	//index :=0


	client, err :=	newGrpcClient()
	if err != nil {

		log.Println(fmt.Sprintf("Flush error %s",err))

		return
	}

	for  x := range c.segment {
		segments=append(segments,*x)
		//todo if c.segment no data that it  will block
		//if index+1 < int(limit){
		//
		//}

		client.Collect(segments)

	}

}


func (c *traceingDispatcher) Close(){
	close(c.segment)
}