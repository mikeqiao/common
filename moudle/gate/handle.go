package gate

import (
	mod "github.com/mikeqiao/ant/module"
	"github.com/mikeqiao/ant/network"
	"github.com/mikeqiao/ant/rpc"
	"github.com/mikeqiao/common/errtype"
	ftype "github.com/mikeqiao/common/moudle/functype"
	"github.com/mikeqiao/common/msg"
)

func HandleUserLogin(call *rpc.CallInfo) {
	in := call.Args.(*msg.CUserLoginRQ) //本服务需要的参数
	uid := call.Data.UId
	//		call.Data //附加信息
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

}
