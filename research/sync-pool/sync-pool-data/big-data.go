package sync_pool_data

import "sync"

var bigDataSyncPool = sync.Pool{
	New: func() interface{} {
		return &BigData{}
	},
}

func GetData() *BigData {
	data := bigDataSyncPool.Get()
	if data == nil {
		return &BigData{}
	}

	return bigDataSyncPool.Get().(*BigData)
}

func PutData(data *BigData) {
	bigDataSyncPool.Put(data)
}

func NewData() *BigData {
	return &BigData{}
}

type BigData struct {
	B1     bool
	BB1    bool
	BB2    bool
	I1     myInts
	Mix1   mixPtrs
	S1     myStrs
	B2     bool
	Mix2   mixPtrs
	I2     myInts
	B3     bool
	S2     myStrs
	MInts1 mixInts
	MInts2 mixInts
	MInts3 mixInts
	MInts4 mixInts
	MInts5 mixInts
	MInts6 mixInts
	MInts7 mixInts
	MInts8 mixInts
	MInts9 mixInts
}

type mixPtrs struct {
	PtrInts *myInts
	MStrInt map[string]int64
	PtrStrs *myStrs
}

type mixInts struct {
	M1 myInts
	M2 myInts
	M3 myInts
	M4 myInts
	M5 myInts
	M6 myInts
	M7 myInts
	M8 myInts
	M9 myInts
}

type myInts struct {
	N1  uint64
	N2  uint64
	N3  uint64
	N4  uint64
	N5  uint64
	N6  uint64
	N7  uint64
	N8  uint64
	N9  uint64
	N19 uint64
	N29 uint64
	N39 uint64
	N49 uint64
	N59 uint64
	N69 uint64
	N79 uint64
	N89 uint64
	N99 uint64
}

type myStrs struct {
	S1 string
	S2 string
	S3 string
	S4 string
	S5 string
	S6 string
	S7 string
	S8 string
	S9 string
}
