package cache

type Error string

func (e Error) Error() string { return string(e) }

var NotFound = Error("Not Found")
