package tlv

type TLVer interface {
	ToBytes() []byte
	Bytes2Header([]byte) error
	Bytes2Body(v []byte) error
	GetHeaderLen() int
	GetBodyLen() int
	GetMid() int
}
