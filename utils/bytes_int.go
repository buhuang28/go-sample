package utils

import (
	"bytes"
	"encoding/binary"
)

// 大端字节序：高位数据存储在低地址
// 小端字节序：低位数据存储在低地址

// 大端：字节转Int
func BigBytes2Int(bys []byte) int {
	switch len(bys) {
	case 8:
		buf := bytes.NewBuffer(bys)
		var data uint64
		binary.Read(buf, binary.BigEndian, &data)
		return int(data)
	case 4:
		buf := bytes.NewBuffer(bys)
		var data uint32
		binary.Read(buf, binary.BigEndian, &data)
		return int(data)
	case 3:
		return int(uint32(bys[2]) | uint32(bys[1])<<8 | uint32(bys[0])<<16)
	case 2:
		buf := bytes.NewBuffer(bys)
		var data uint16
		binary.Read(buf, binary.BigEndian, &data)
		return int(data)
	case 1:
		buf := bytes.NewBuffer(bys)
		var data uint8
		binary.Read(buf, binary.BigEndian, &data)
		return int(data)
	}
	return 0
}

// 小端：字节转Int
func LitBytes2Int(bys []byte) int {
	bytebuffer := bytes.NewBuffer(bys)
	switch len(bys) {
	case 4:
		var data uint32
		binary.Read(bytebuffer, binary.LittleEndian, &data)
		return int(data)
	case 3:
		return int(uint32(bys[0]) | uint32(bys[1])<<8 + uint32((bys[2]))<<16)
	case 2:
		var data uint16
		binary.Read(bytebuffer, binary.LittleEndian, &data)
		return int(data)
	case 1:
		var data uint8
		binary.Read(bytebuffer, binary.LittleEndian, &data)
		return int(data)
	}
	return 0
}

func BigInt2Bytes(i int) []byte {
	data := uint16(i)
	bytebuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}

func BigInt23Bytes(i int) []byte {
	b := make([]byte, 3)
	b[0] = byte(i >> 16)
	b[1] = byte(i >> 8)
	b[2] = byte(i)
	return b
}

func BigInt24Bytes(i int) []byte {
	data := uint32(i)
	bytebuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}

func BigInt28Bytes(i int64) []byte {
	data := uint64(i)
	bytebuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}

func LitInt22Bytes(i int) []byte {
	data := uint16(i)
	bytebuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuffer, binary.LittleEndian, data)
	return bytebuffer.Bytes()
}

func LitInt23Bytes(i int) []byte {
	b := make([]byte, 3)
	b[0] = byte(i)
	b[1] = byte(i >> 8)
	b[2] = byte(i >> 16)
	return b
}

func LitInt24Bytes(i uint32) []byte {
	bytebuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuffer, binary.LittleEndian, i)
	return bytebuffer.Bytes()
}

func LitInt28Bytes(i int64) []byte {
	data := uint64(i)
	bytebuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuffer, binary.LittleEndian, data)
	return bytebuffer.Bytes()
}
