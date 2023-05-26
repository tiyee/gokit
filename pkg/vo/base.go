package vo

type Base struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}
