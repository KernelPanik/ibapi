package main

import (
	"github.com/KernelPanik/ibapi"

	"log"
	"time"
)

func main() {
	e, err := ibapi.NewEngine(ibapi.DefaultGateway, 1)
	if err != nil {
		log.Printf("Rcvd error on open, %v\n", err)
		panic(err)
	}

	dumpReplies(e)

	writeRequest(e, &ibapi.MsgOutReqContractData{ReqId:100, Symbol:"DBK", SecurityType:"STK", Exchange:"FWB"})

	<-time.After(time.Second * 10)
	writeRequest(e, &ibapi.MsgOutReqContractData{ReqId:101, Symbol:"IBM", SecurityType:"STK", Exchange:"NYSE"})

	<-time.After(time.Second * 10)
	writeRequest(e, &ibapi.MsgOutReqMktData{
		TickerId:1,
		Symbol:"USD",
		SecurityType:"CASH",
		Exchange:"IDEALPRO",
		Currency:"JPY",
	})

	<-time.After(time.Second * 20)
	e.Stop()
}

func dumpReplies(e *ibapi.Engine) {
	go func() {
		for {
			rep, err := e.ReadReply()
			if err != nil {
				log.Printf("Rcvd error on read, %v\n", err)
				break
			}
			log.Printf("Rcvd message: %#v\n", rep)
		}
	}()
}

func writeRequest(e *ibapi.Engine, req interface{}) {
	log.Printf("Writing message: %#v\n", req)
	err := e.WriteRequest(req)
	if err != nil {
		log.Printf("Rcvd error on write, %v\n", err)
		panic("failed write")
	}
	log.Printf("Wrote message: %#v\n", req)
}
