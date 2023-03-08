package tlv

type Header struct {
	Mid   []byte //消息ID 2byte
	Time  []byte //时间   4byte
	Len   []byte //2byte   消息体长度
	CRC32 []byte //4byte 校验码
}

const (
	DEF_HEADE_LEN = 12
)
