package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"managerStudent/ver2/thrift/gen-go/datamanager"
)

func init() {
	fmt.Println("Run once - Connect to BigSet")
	client = Connect()
	_, err := client.TotalStringKeyCount(context.Background())
	if err != nil {
		log.Fatal(err, " myThrift/Server.go:16")
	}
	fmt.Println("Connected")
}

func main() {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket("127.0.0.1:7778")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", transport)
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTCompactProtocolFactory()
	handler := NewManagerStudentHandler()
	processor := datamanager.NewManagerStudentProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", "127.0.0.1:7778")
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
