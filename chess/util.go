
package chess

import (
	bytes "bytes"
)

func writeUint32ToBuf(buf *bytes.Buffer, n uint32)(*bytes.Buffer){
	buf.WriteByte(byte(n >> 24 & 0xff))
	buf.WriteByte(byte(n >> 16 & 0xff))
	buf.WriteByte(byte(n >> 8  & 0xff))
	buf.WriteByte(byte(n & 0xff))
	return buf
}
func readUint32FromBuf(buf *bytes.Buffer) uint32 {
	n1, _ := buf.ReadByte()
	n2, _ := buf.ReadByte()
	n3, _ := buf.ReadByte()
	n4, _ := buf.ReadByte()
	return uint32(n1 << 24 | n2 << 16 | n3 << 8 | n4)
}

func checkCond(cond bool, truev, falsev interface{}) interface{} {
	if cond {
		return truev
	}
	return falsev
}


func zipByte2(arr *[2]byte) byte {
	return (arr[0] << 4) | (arr[1] & 0xf)
}

func uzipByte2(n byte)(arr *[2]byte){
	arr = new([2]byte)
	arr[0] = n >> 4
	arr[1] = n & 0xf
	return arr
}
