package models

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"managerStudent/myThrift/gen-go/apiservice"
	"sync"
)

type ConThrift *apiservice.ManagerStudentClient

var client ConThrift

var doOnce sync.Once

func DoSomething() {
	doOnce.Do(func() {
		fmt.Println("Run once - first time, loading...")
		client = Connect()
	})
	fmt.Println("Run this every time")
}

func Connect() ConThrift{
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("127.0.0.1:7777")
	if err != nil {
		log.Fatal("Error opening socket:", err)
	}
	transportFactory := thrift.NewTTransportFactory()
	transport, err = transportFactory.GetTransport(transport)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	if err != nil {
		log.Fatal(err)
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return apiservice.NewManagerStudentClient(thrift.NewTStandardClient(iprot, oprot))
}
