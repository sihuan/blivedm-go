package message

type User struct {
	Uid          int
	Uname        string
	Admin        bool
	Owner        bool
	Urank        int
	MobileVerify bool
	Medal        *Medal
	GuardLevel   int
}

type Medal struct {
	Name     string
	Level    int
	Color    int
	UpRoomId int
	UpUid    int
	UpName   string
}
