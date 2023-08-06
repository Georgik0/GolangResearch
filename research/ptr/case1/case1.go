package case1

import "encoding/binary"

func Case() {
	x := []byte{0x0A, 0x15, 0x0e, 0x28, 0x05, 0x96, 0x0b, 0xd0, 0x0}

	a := x[0]
	b := binary.LittleEndian.Uint16(x[1:3])
	c := binary.LittleEndian.Uint16(x[3:5])
	d := binary.LittleEndian.Uint32(x[5:9])

	println(a, b, c, d)
	println(a, b, c, d)
}
