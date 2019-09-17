package msgtype

// base
const (
	_Datatype uint32 = 1000 + iota
	DataUserAccount
	DataUserBase
	DataUserAdd
	DataUserBag
)

//array
const (
	_Arraytype uint32 = 2000 + iota
	ArrayDataUserBag
)

//config
const (
	_Configtype uint32 = 10000 + iota
	DataBaseConfig
)

//confarray
const (
	_ConfArraytype uint32 = 20000 + iota
	ArrayDataBaseConfig
)

//msg c-s
const (
	_CStype uint32 = 30000 + iota
	CUserLoginRQ
	CUserLoginRS
	CCreateAccountRQ
	CCreateAccountRS
	CCreateUserRQ
	CCreateUserRS
)
