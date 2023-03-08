package handle_message

import (
	"errors"
	"fmt"
	"go-sample/tcp_sample/cst"
	"go-sample/tcp_sample/pb/pb"
	"go-sample/tcp_sample/server/pool"
	"go-sample/tcp_sample/tlv"
	"go-sample/utils"
	"google.golang.org/protobuf/proto"
	"time"
)

func HandleMessage(t tlv.TLVer) ([]byte, error) {
	v := t.(*tlv.MessageTLV)
	defer func() {
		pool.MessagePool.Put(t)
	}()
	resp := new(pb.MessageResp)
	if time.Now().Unix()-int64(utils.LitBytes2Int(v.Time)) > 300 {
		resp.Status = cst.ERROR_CODE
		resp.Result = "time is error"
		marshal, _ := proto.Marshal(resp)
		v.AddVal(marshal)
		return v.ToBytes(), errors.New("null value")
	}

	if len(v.Val) == 0 {
		resp.Status = cst.ERROR_CODE
		resp.Result = "null value"
		marshal, _ := proto.Marshal(resp)
		v.AddVal(marshal)
		return v.ToBytes(), errors.New("null value")
	}

	data := new(pb.MessageReq)
	err := proto.Unmarshal(v.Val, data)
	if err != nil {
		resp.Status = cst.ERROR_CODE
		resp.Result = err.Error()
		marshal, _ := proto.Marshal(resp)
		v.AddVal(marshal)
		return v.ToBytes(), errors.New("error data format")
	}

	if data.Data == "" {
		resp.Status = cst.ERROR_CODE
		resp.Result = "data is null"
		marshal, _ := proto.Marshal(resp)
		v.AddVal(marshal)
		return v.ToBytes(), errors.New("data is null")
	}

	resp.Status = cst.SUCCESS_CODE
	resp.Result = cst.SUCCESS_MSG
	resp.Message = fmt.Sprintf("This is your data:%s", data.Data)
	marshal, _ := proto.Marshal(resp)
	v.AddVal(marshal)
	return v.ToBytes(), nil
}
