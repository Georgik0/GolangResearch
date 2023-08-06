package only_slices

type SnapshotArray struct {
	length         int
	snapCount      int
	keyValueSource keyValueSource
}

type valueWithSnapshot struct {
	snap int
	val  int
}

func newValueSnap() valueWithSnapshot {
	return valueWithSnapshot{snap: -1}
}

type keyValueSource []*Data

func (v *valueWithSnapshot) setValueAndSnap(valSnap valueWithSnapshot) {
	*v = valSnap
}

func (v *valueWithSnapshot) incSnap() {
	v.snap++
}

type Data struct {
	isNeedAppend bool
	arr          []valueWithSnapshot
}

func (d *Data) initArrIfNIL() {
	if len(d.arr) == 0 {
		d.arr = append(d.arr, newValueSnap())
	}
}

func (d *Data) getElementBySnap(snapID int) int {
	for _, v := range d.arr {
		if v.snap >= snapID {
			return v.val
		}
	}

	return 0
}

func (d *Data) getLastArrElement() *valueWithSnapshot {
	return &d.arr[len(d.arr)-1]
}

func (d *Data) setLastValueWithSnap(valSnap valueWithSnapshot) {
	d.getLastArrElement().setValueAndSnap(valSnap)
}

func (d *Data) incrementSnap() {
	d.getLastArrElement().incSnap()
}

func Constructor(length int) SnapshotArray {
	snapArr := SnapshotArray{
		length:         length,
		keyValueSource: make(keyValueSource, length),
		snapCount:      -1,
	}

	for i := 0; i < length; i++ {
		snapArr.keyValueSource[i] = &Data{}
		snapArr.keyValueSource[i].initArrIfNIL()
	}

	return snapArr
}

func (this *SnapshotArray) Set(index int, val int) {
	if this.keyValueSource[index].getLastArrElement().val == val {
		return
	}

	valSnap := valueWithSnapshot{
		val:  val,
		snap: this.snapCount,
	}

	if this.keyValueSource[index].isNeedAppend {
		this.keyValueSource[index].arr = append(this.keyValueSource[index].arr, valSnap)
	} else {
		this.keyValueSource[index].setLastValueWithSnap(valSnap)
	}

}

func (this *SnapshotArray) Snap() int {
	for _, v := range this.keyValueSource {
		v.isNeedAppend = true
		v.incrementSnap()
	}

	this.snapCount++

	return this.snapCount
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	//if this.keyValueSource[index] == nil {
	//	this.keyValueSource[index] = &Data{}
	//}

	//this.keyValueSource[index].initArrIfNIL()

	return this.keyValueSource[index].getElementBySnap(snap_id)
}
