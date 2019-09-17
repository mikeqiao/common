package gate

import (
	"github.com/mikeqiao/common/msg"
	"github.com/mikeqiao/newant/log"
	"github.com/mikeqiao/newant/network"
	"github.com/mikeqiao/newant/rpc"
)

func HandleUserLogin(call *rpc.CallInfo) {
	log.Debug("msg:%v,", call.Args)
	in := call.Args.(*msg.CUserLoginRQ) //本服务需要的参数
	uid := call.Data.UId
	//		call.Data //附加信息
	/*
		cb := func(in interface{}, e error) {
			if nil != in {
				bm, _ := in.(*msg.CUserLoginRS)
				if errtype.OK == bm.GetResult() && nil != bm.GetUser() {
					network.CM.AddUserConn(uid, bm.GetUser().GetUserUID())
				}
			}
			call.SetResult(in, e)
		}
		mod.RPC.Route(ftype.HandleVerifyAccount, cb, in, call.Data)
	*/
	network.CM.AddUserConn(uid, uid)
	log.Debug("msg:%v, uid:%v", in, uid)
	bmsg := new(msg.CUserLoginRS)
	bmsg.Result = 1
	bmsg.AccountID = "get the request"

	call.SetResult(bmsg, nil)

}
