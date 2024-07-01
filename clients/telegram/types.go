package telegram

type UPdatesRresponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID       int    `json:"update_id"`
	Messsage string `json:"message"`
}
