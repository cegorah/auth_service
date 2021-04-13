package config

type JwtEmptyError struct {
	msg string
}

func (e *JwtEmptyError) Error() string{
	return e.msg
}
