package case1

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func Check() {
	rw := NewResponseWriter()

	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		iGo := i
		str := strconv.Itoa(iGo)
		go func() {
			rw.AddState(&str, int64(iGo))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(rw.ret)
	buf := make(map[string]struct{})
	for _, s := range rw.strArr {
		if i, ok := buf[s]; ok {
			fmt.Println("id = ", i, "; s = ", s, ";   ", rw.strArr)
			return
		}
		buf[s] = struct{}{}
	}
}

type ResponseWriter struct {
	mtx    sync.Mutex
	ret    []int64
	strArr []string
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		mtx:    sync.Mutex{},
		ret:    make([]int64, 0),
		strArr: make([]string, 0),
	}
}

func (rw *ResponseWriter) AddState(str *string, component int64) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)

	time.Sleep(time.Duration(randomNum + 50000))

	rw.strArr = append(rw.strArr, *str)
	rw.addStateWithId(component)
}

func (rw *ResponseWriter) addStateWithId(stateWithID int64) {
	//rw.mtx.Lock()
	//defer rw.mtx.Unlock()

	rw.ret = append(rw.ret, stateWithID)
}
