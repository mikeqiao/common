package p

import (
	pb "github.com/mikeqiao/ant/net/protobuf"
	"github.com/mikeqiao/ant/network"
	ftype "github.com/mikeqiao/common/moudle/functype"
	"github.com/mikeqiao/common/msg"
	"github.com/mikeqiao/common/msgtype"
)

func SetP() {
	p := pb.NewProcessor()
	network.P.AddData(1, p)
	//base
	p.Register(&msg.DataUserAccount{}, msgtype.DataUserAccount)
	p.Register(&msg.DataUserBase{}, msgtype.DataUserBase)
	p.Register(&msg.DataUserAdd{}, msgtype.DataUserAdd)
	p.Register(&msg.DataUserBag{}, msgtype.DataUserBag)

	//array
	p.Register(&msg.ArrayDataUserBag{}, msgtype.ArrayDataUserBag)

	//config
	p.Register(&msg.DataBaseConfig{}, msgtype.DataBaseConfig)

	//conf array
	p.Register(&msg.ArrayDataBaseConfig{}, msgtype.ArrayDataBaseConfig)

	//msg
	p.Register(&msg.CUserLoginRQ{}, msgtype.CUserLoginRQ)
	p.Register(&msg.CUserLoginRS{}, msgtype.CUserLoginRS)
	p.Register(&msg.CCreateAccountRQ{}, msgtype.CCreateAccountRQ)
	p.Register(&msg.CCreateAccountRS{}, msgtype.CCreateAccountRS)
	p.Register(&msg.CCreateUserRQ{}, msgtype.CCreateUserRQ)
	p.Register(&msg.CCreateUserRS{}, msgtype.CCreateUserRS)

	//set SetRouter
	p.SetRouter(msgtype.CUserLoginRQ, ftype.HandleUserLogin)
	p.SetRouter(msgtype.CCreateAccountRQ, ftype.HandCreateAccount)
	p.SetRouter(msgtype.CCreateUserRQ, ftype.HandCreateUser)
}
