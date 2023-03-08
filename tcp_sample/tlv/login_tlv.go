package tlv

import (
	"errors"
	"go-sample/utils"
)

type LoginTLV struct {
	Header
	Val []byte //pb信息
}

func (l *LoginTLV) GetMid() int {
	return utils.LitBytes2Int(l.Mid)
}

func (l *LoginTLV) ToBytes() []byte {
	data := make([]byte, DEF_HEADE_LEN+utils.LitBytes2Int(l.Len)+2)
	copy(data[0:2], utils.LitInt22Bytes(LOGIN_TAG))
	copy(data[2:4], l.Mid)
	copy(data[4:8], l.Time)
	copy(data[8:10], l.Len)
	copy(data[10:14], l.CRC32)
	copy(data[14:], l.Val)
	return data
}

func (l *LoginTLV) Bytes2Header(data []byte) error {
	if len(data) != DEF_HEADE_LEN {
		return errors.New("header length error")
	}
	l.Mid = data[0:2]
	l.Time = data[2:6]
	l.Len = data[6:8]
	l.CRC32 = data[8:12]
	return nil
}

func (l *LoginTLV) Bytes2Body(v []byte) error {
	vLen := len(v)
	if utils.LitBytes2Int(l.Len) != vLen {
		return errors.New("data length error")
	}
	data := make([]byte, DEF_HEADE_LEN+vLen-4)
	copy(data[0:2], l.Mid)
	copy(data[2:6], l.Time)
	copy(data[6:8], l.Len)
	copy(data[8:], v)
	crc32 := utils.CRC32(data)
	if crc32 != uint32(utils.LitBytes2Int(l.CRC32)) {
		return errors.New("login crc error")
	}
	l.Val = v
	return nil
}

func (l *LoginTLV) GetHeaderLen() int {
	return DEF_HEADE_LEN
}

func (l *LoginTLV) GetBodyLen() int {
	return utils.LitBytes2Int(l.Len)
}

func (l *LoginTLV) AddVal(val []byte) {
	l.Len = utils.LitInt22Bytes(len(val))
	data := make([]byte, DEF_HEADE_LEN+len(val)-4)
	copy(data[0:2], l.Mid)
	copy(data[2:6], l.Time)
	copy(data[6:8], l.Len)
	copy(data[8:], val)
	crc32 := utils.CRC32(data)
	l.CRC32 = utils.LitInt24Bytes(crc32)
	l.Val = val
}

func NewLoginTLV(id int, time int64) *LoginTLV {
	return &LoginTLV{
		Header: Header{
			Mid:  utils.LitInt22Bytes(id),
			Time: utils.LitInt24Bytes(uint32(time)),
		},
	}
}
