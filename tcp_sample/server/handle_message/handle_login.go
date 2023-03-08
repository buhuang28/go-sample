package handle_message

import (
	"go-sample/tcp_sample/cst"
	"go-sample/tcp_sample/pb/pb"
	"go-sample/tcp_sample/server/pool"
	"go-sample/tcp_sample/tlv"
	"go-sample/utils"
	"google.golang.org/protobuf/proto"
	"time"
)

func HandleLogin(t tlv.TLVer) ([]byte, error) {
	v := t.(*tlv.LoginTLV)
	defer func() {
		pool.LoginPool.Put(t)
	}()
	resp := new(pb.LoginResp)
	if time.Now().Unix()-int64(utils.LitBytes2Int(v.Time)) > 300 {
		resp.Status = cst.ERROR_CODE
		resp.Result = "time is error"
		marshal, _ := proto.Marshal(resp)
		v.AddVal(marshal)
		return v.ToBytes(), nil
	}

	if len(v.Val) == 0 {
		resp.Status = cst.ERROR_CODE
		resp.Result = "null value"
		marshal, _ := proto.Marshal(resp)
		v.AddVal(marshal)
		return v.ToBytes(), nil
	}
	account := new(pb.LoginReq)
	err := proto.Unmarshal(v.Val, account)
	if err != nil {
		resp.Status = cst.ERROR_CODE
		resp.Result = err.Error()
		marshal, _ := proto.Marshal(resp)
		v.AddVal(marshal)
		return v.ToBytes(), nil
	}
	if account.Account != "root" && account.Password != "123456" {
		resp.Status = cst.ERROR_CODE
		resp.Result = "password or user error"
		marshal, _ := proto.Marshal(resp)
		v.AddVal(marshal)
		return v.ToBytes(), nil
	}
	resp.Status = cst.SUCCESS_CODE
	resp.Result = cst.SUCCESS_MSG
	marshal, _ := proto.Marshal(resp)
	v.AddVal(marshal)
	return v.ToBytes(), nil
}
