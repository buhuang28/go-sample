package tlv

import (
	"errors"
	"go-sample/utils"
)

type MessageTLV struct {
	Header
	Frag byte   //分段ID，当消息分段发送的时候确定是否为同一段，不分段为0
	Cont byte   //是否有后续标识符 1有后续，0无后续
	Val  []byte //消息内容
}

const (
	MSG_TLV_HEAD_LEN = DEF_HEADE_LEN + 2
)

func (m *MessageTLV) ToBytes() []byte {
	data := make([]byte, MSG_TLV_HEAD_LEN+utils.LitBytes2Int(m.Len)+2)
	copy(data[0:2], utils.LitInt22Bytes(MESSAGE_TAG))
	copy(data[2:4], m.Mid)
	copy(data[4:8], m.Time)
	copy(data[8:10], m.Len)
	copy(data[10:14], m.CRC32)
	data[14] = m.Frag
	data[15] = m.Cont
	copy(data[16:], m.Val)
	return data
}

func (m *MessageTLV) Bytes2Header(data []byte) error {
	if len(data) != MSG_TLV_HEAD_LEN {
		return errors.New("header length error")
	}
	m.Mid = data[0:2]
	m.Time = data[2:6]
	m.Len = data[6:8]
	m.CRC32 = data[8:12]
	m.Frag = data[12]
	m.Cont = data[13]
	return nil
}

func (m *MessageTLV) Bytes2Body(v []byte) error {
	vLen := len(v)
	if utils.LitBytes2Int(m.Len) != vLen {
		return errors.New("data length error")
	}
	data := make([]byte, MSG_TLV_HEAD_LEN+vLen-4) //-4是减少crc32
	copy(data[0:2], m.Mid)
	copy(data[2:6], m.Time)
	copy(data[6:8], m.Len)
	data[8] = m.Frag
	data[9] = m.Cont
	copy(data[10:], v)
	crc32 := utils.CRC32(data)
	if crc32 != uint32(utils.LitBytes2Int(m.CRC32)) {
		return errors.New("message crc error")
	}
	m.Val = v
	return nil
}

func (m *MessageTLV) GetHeaderLen() int {
	return MSG_TLV_HEAD_LEN
}

func (m *MessageTLV) GetBodyLen() int {
	return utils.LitBytes2Int(m.Len)
}

func (m *MessageTLV) AddVal(val []byte) {
	m.Len = utils.LitInt22Bytes(len(val))
	data := make([]byte, MSG_TLV_HEAD_LEN+len(val)-4)
	copy(data[0:2], m.Mid)
	copy(data[2:6], m.Time)
	copy(data[6:8], m.Len)
	data[8] = m.Frag
	data[9] = m.Cont
	copy(data[10:], val)
	crc32 := utils.CRC32(data)
	m.CRC32 = utils.LitInt24Bytes(crc32)
	m.Val = val
}

func (m *MessageTLV) GetMid() int {
	return utils.LitBytes2Int(m.Mid)
}

func NewMessageTLV(id int, time int64, fid, cont byte) *MessageTLV {
	return &MessageTLV{
		Header: Header{
			Mid:  utils.LitInt22Bytes(id),
			Time: utils.LitInt24Bytes(uint32(time)),
		},
		Frag: fid,
		Cont: cont,
	}
}
