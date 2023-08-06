package main

import (
	"fmt"
)

type SnapshotArray struct {
	length                    int
	snapCount                 int
	wasCommitted              bool
	snapshotArrWasInitialized bool
	keyValueSource            keyValueSource
}

type valueWithSnapshot struct {
	val     int
	snapTmp int
	snap    *int
}

func newValueSnap(snap *int) valueWithSnapshot {
	return valueWithSnapshot{
		snap:    snap, // snap: -1
		snapTmp: *snap,
	}
}

type keyValueSource map[int]*Data

func (v *valueWithSnapshot) setValueAndSnap(valSnap valueWithSnapshot) {
	*v = valSnap
}

func (v *valueWithSnapshot) setTmpSnap() {
	v.snapTmp = *v.snap
	v.snap = &v.snapTmp
}

type Data struct {
	arr []valueWithSnapshot
}

func (d *Data) initArr(snap *int) {
	if len(d.arr) == 0 {
		d.arr = append(d.arr, newValueSnap(snap))
	}
}

func (d *Data) getElementBySnap(snapID int, wasCommitted bool) int {
	lastIndx := len(d.arr) - 1
	penultimateIndx := lastIndx - 1

	for i := 0; i <= penultimateIndx; i++ {
		if wasCommitted && d.arr[i].snapTmp >= snapID { //  || *d.arr[i].snap >= snapID
			return d.arr[i].val
		}
	}

	if wasCommitted && (d.arr[lastIndx].snapTmp >= snapID || *d.arr[lastIndx].snap >= snapID) {
		return d.arr[lastIndx].val
	}

	return 0
}

func (d *Data) getLastArrElement() *valueWithSnapshot {
	return &d.arr[len(d.arr)-1]
}

func (d *Data) getPenultimateArrElement() *valueWithSnapshot {
	return &d.arr[len(d.arr)-2]
}

func (d *Data) setLastValueWithSnap(valSnap valueWithSnapshot) {
	d.getLastArrElement().setValueAndSnap(valSnap)
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		length:         length,
		keyValueSource: make(keyValueSource, length),
		snapCount:      -1,
	}
}

func (this *SnapshotArray) Initialize() {
	for i := 0; i < this.length; i++ {
		this.keyValueSource[i] = &Data{}
		this.keyValueSource[i].initArr(&this.snapCount)
	}

	this.snapshotArrWasInitialized = true
}

func (this *SnapshotArray) Set(index int, val int) {
	if !this.snapshotArrWasInitialized {
		this.Initialize()
	}

	if this.keyValueSource[index].getLastArrElement().val == val {
		return
	}

	valSnap := valueWithSnapshot{
		val:     val,
		snap:    &this.snapCount,
		snapTmp: this.snapCount,
	}

	if this.wasCommitted {
		this.keyValueSource[index].getLastArrElement().setTmpSnap()
		this.keyValueSource[index].arr = append(this.keyValueSource[index].arr, valSnap)
	} else {
		this.keyValueSource[index].setLastValueWithSnap(valSnap)
	}

}

func (this *SnapshotArray) Snap() int {
	if !this.snapshotArrWasInitialized {
		this.Initialize()
	}

	this.wasCommitted = true
	this.snapCount = this.snapCount + 1

	return this.snapCount
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	if !this.snapshotArrWasInitialized {
		this.Initialize()
	}

	return this.keyValueSource[index].getElementBySnap(snap_id, this.wasCommitted)
}

func main() {
	//case1()
	//case2()
	case3()
}

func case1() {
	snapshotArr := Constructor(100)

	for i := 0; i < 100000; i++ {
		for j := 0; j < 10; j++ {
			snapshotArr.Set(j, i)
			snapshotArr.Snap()
		}
	}

	for j := 0; j < 10; j++ {
		for i := 10; i < 100000; i++ {
			res := snapshotArr.Get(j, i)

			expectedRes := i / 10
			if i%10 < j {
				expectedRes--
			}

			if res != expectedRes {
				fmt.Println("Bad result")
				fmt.Println("i = ", i)
				fmt.Println("j = ", j)
				fmt.Println("res = ", res)
				return
			}
		}
	}
}

func case2() {
	var res int
	snapshotArr := Constructor(1)

	snapshotArr.Set(0, 15)
	snapshotArr.Snap()
	snapshotArr.Snap()
	snapshotArr.Snap()
	res = snapshotArr.Get(0, 2)
	fmt.Println(res)

	snapshotArr.Snap()
	snapshotArr.Snap()
	res = snapshotArr.Get(0, 0)
	fmt.Println(res)
}

func case3() {
	var res int
	snapshotArr := Constructor(1)

	snapshotArr.Snap()
	snapshotArr.Snap()
	snapshotArr.Set(0, 4)
	snapshotArr.Snap()

	res = snapshotArr.Get(0, 1)
	fmt.Println(res)
	snapshotArr.Set(0, 12)
	res = snapshotArr.Get(0, 1)
	fmt.Println(res)

	snapshotArr.Snap()
	res = snapshotArr.Get(0, 3)
	fmt.Println(res)
}
