package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
	"unsafe"

	sync_pool_data "GolangResearch/research/sync-pool/sync-pool-data"
)

func main() {
	sv1 := &starterV1{}

	http.Handle("/start-v1", sv1)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type starterV1 struct{}

var ReqCounter = atomic.Int64{}

func (sv1 *starterV1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ReqCounter.Add(1)

	log.Printf("counter: %v\n", ReqCounter.Load())

	//loggingAddr()
	speedAllocTest()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func speedAllocTest() {
	if ReqCounter.Load() == 1 {
		addr := make([]unsafe.Pointer, 0, 1000)
		start := time.Now()
		summTime := time.Time{}

		for i := 0; i < 10; i++ {
			data := sync_pool_data.GetData()
			sync_pool_data.PutData(data)

			finishTime := time.Since(start)
			summTime.Add(finishTime.Abs())

			addr = append(addr, unsafe.Pointer(data))
		}

		fmt.Printf("sunc.Pool allocTime: %v", summTime.Nanosecond())
	} else {
		addr := make([]unsafe.Pointer, 0, 1000)
		start := time.Now()
		summTime := time.Time{}

		for i := 0; i < 10; i++ {
			data := sync_pool_data.NewData()

			finishTime := time.Since(start)
			summTime.Add(finishTime.Abs())

			data.BB2 = true
			addr = append(addr, unsafe.Pointer(data))
		}

		fmt.Printf("stack allocTime: %v", summTime.Nanosecond())
	}

	fmt.Println()
}

func loggingAddr() {
	if ReqCounter.Load() == 1 {
		data := sync_pool_data.GetData()

		data.BB1 = true
		sync_pool_data.PutData(data)
		fmt.Printf("1: %v; addr: %p\n", *data, data)

		data.BB2 = true
		sync_pool_data.PutData(data)
		fmt.Printf("1: %v; addr: %p\n", *data, data)

		time.Sleep(2 * time.Second)

		data = sync_pool_data.GetData()
		fmt.Printf("1: %v; addr: %p\n", *data, data)
	} else {
		data := sync_pool_data.GetData()
		data.B1 = true
		fmt.Printf("2: %v; addr: %p\n", *data, data)
	}
}
