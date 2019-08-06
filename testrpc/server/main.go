package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"test/testrpc/hpgrpc"
)

var GRPCPORT="localhost:55555"

var testdata = hpgrpc.SettleResData{
	Uid :"555",
	Tid:"555",
	 Handletime:111,
	 Settletime:111,
	 Inputamount:5.5,
	 Outputamount:5.5,
	 Ordervalue:5.5,
	 Settlevalue:5.5,
	  Side:1,
	 Interval:30,
	 Symbol:"BTC",
	 Orderresult:5,
	 Balance:5,
	 Vitualbalance:4,
}
type server struct {

}

func (*server) QuotationSettleData(stream hpgrpc.Hpgrpcservice_QuotationSettleDataServer) error{
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}



		log.Printf("PCTEST----server data:%v\n",data)
		 if err := stream.Send(&testdata); err != nil {
			 log.Printf("PCTEST----server send err:%v\n",err)
		 }
	}
}


func main() {
	lis, _ := net.Listen("tcp",GRPCPORT)
	s := grpc.NewServer()
	hpgrpc.RegisterHpgrpcserviceServer(s,&server{})
	if err := s.Serve(lis); err != nil {
		log.Printf("PCTEST----err:%v\n",err)
	}
}