package models

import (
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"managerStudent/ver2/thrift/gen-go/datamanager"
	"sync"
)

var client *datamanager.ManagerStudentClient
var doOnce = &sync.Once{}

func initOne() {
	log.Println("Run once - first time, Connect to thrift")
	client = Connect()
	log.Println("Connected")
}

func GetClient() *datamanager.ManagerStudentClient {
	doOnce.Do(initOne)
	return client
}

func Connect() *datamanager.ManagerStudentClient {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("127.0.0.1:7778")
	if err != nil {
		log.Fatal("Error opening socket:", err)
	}
	transportFactory := thrift.NewTTransportFactory()
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		log.Fatal(err, " src/beeApi/models/oneRun.go:35")
	}
	if err := transport.Open(); err != nil {
		log.Fatal(err)
	}
	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return datamanager.NewManagerStudentClient(thrift.NewTStandardClient(iprot, oprot))
}
