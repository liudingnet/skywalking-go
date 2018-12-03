package main

import (
	"context"
	pb "github.com/WithLin/skywalking-go/proto"
	"google.golang.org/grpc"
	"log"
)


type test struct {

}

func (c *test)ApplicationCodeRegister(ctx context.Context, in *pb.Application)(*pb.ApplicationMapping,error){
	a :=new(pb.ApplicationMapping)
	b :=new(pb.KeyWithIntegerValue)
	b.Value=12345

	a.Application=b
	return  a,nil
}





func main()  {

	//go func() {
	//
	//	listen, err := net.Listen("tcp", "localhost:11111")
	//	if err != nil {
	//		grpclog.Fatalf("failed to listen: %v", err)
	//	}
	//
	//	s := grpc.NewServer()
	//
	//
	//	pb.RegisterApplicationRegisterServiceServer(s, &test{})
	//
	//	grpclog.Println("Listen on " + "localhost:11111")
	//
	//	s.Serve(listen)
	//}()

	conn, err := grpc.Dial("localhost:11800", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Can't connect: " + "localhost:11800")
	}
	defer conn.Close()
	for ; ;  {
		resp,err:=pb.NewApplicationRegisterServiceClient(conn).ApplicationCodeRegister(context.Background(),&pb.Application{ApplicationCode:"aaaaa"})
		if err != nil {
			log.Fatalln( err)
		}
		log.Println(resp)

	}



}

