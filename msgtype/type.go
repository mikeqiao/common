package msgtype

// base
const (
	_Datatype uint16 = 1000 + iota
	DataUserAccount
	DataUserBase
	DataUserAdd
	DataUserBag
)

//array
const (
	_Arraytype uint16 = 2000 + iota
	ArrayDataUserBag
)

//config
const (
	_Configtype uint16 = 10000 + iota
	DataBaseConfig
)

//confarray
const (
	_ConfArraytype uint16 = 20000 + iota
	ArrayDataBaseConfig
)

//msg c-s
const (
	_CStype uint16 = 30000 + iota
	CUserLoginRQ
	CUserLoginRS
	CCreateAccountRQ
	CCreateAccountRS
	CCreateUserRQ
	CCreateUserRS
)
