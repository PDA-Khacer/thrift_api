package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"managerStudent/myThrift/gen-go/apiservice"
)

func main() {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket("127.0.0.1:7777")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", transport)
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTCompactProtocolFactory()
	handler := NewManagerStudentHandler()
	processor := apiservice.NewManagerStudentProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", "127.0.0.1:7777")
	server.Serve()
}

