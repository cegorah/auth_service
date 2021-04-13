package redis_cache

type EmptyResult struct {
	msg string
}

func (e *EmptyResult) Error() string {
	return e.msg
}
