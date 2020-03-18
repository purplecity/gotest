package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"test/testrpc/hpgrpc"
	"time"
)

var RemoteAddrPort = ":55555"


func main() {
	conn, err := grpc.Dial(RemoteAddrPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("ERROR----did not connect grpc server: %v\n", err)
	}
	defer conn.Close()
	c := hpgrpc.NewHpgrpcserviceClient(conn)
	hpstream, err := c.QuotationSettleData(context.Background())

	go func() {
		for {
			data, err := hpstream.Recv()
			if err == io.EOF {
				// read done.
				log.Fatalf("err iof")
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Println("PCTEST----xmlfileread:%v\n",data)
		}
	}()

	tick := time.Tick(time.Second)
	for range tick {
		go func() {
			_ = hpstream.Send(&hpgrpc.QuotationReqData{Price:3.33,Timestamp:33333})
		}()
	}

}

