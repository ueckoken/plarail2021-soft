package clientSync

type ClientSync struct {
	State []SingleState
}

type SingleState struct {
  Name  string `json:"name"`
  OnOff bool `json:"onOff"`
}
