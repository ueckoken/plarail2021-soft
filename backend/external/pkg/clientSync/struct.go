package clientSync

type ClientSync struct {
	State []SingleState
}

type SingleState struct {
	name  string
	onOff bool
}
