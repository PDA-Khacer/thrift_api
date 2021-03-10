package main

import (
	"log"
	"managerStudent/serverDB/thrift/gen-go/openstars/core/bigset/generic"

	"github.com/apache/thrift/lib/go/thrift"
)

func Connect() *generic.TStringBigSetKVServiceClient {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("127.0.0.1:18990")
	if err != nil {
		log.Fatal("Error opening socket:", err)
	}
	transportBuff := thrift.NewTBufferedTransportFactory(8192)
	transportFactory1 := thrift.NewTFramedTransportFactory(transportBuff)
	transport, err = transportFactory1.GetTransport(transport)
	if err != nil {
		log.Fatal(err)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactory(true, true)
	if err := transport.Open(); err != nil {
		log.Fatal(err)
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return generic.NewTStringBigSetKVServiceClient(thrift.NewTStandardClient(iprot, oprot))
}
