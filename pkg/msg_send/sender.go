package msg_send

type EmailSender struct {
}

type SMSSender struct {
}

type Sender interface {
	Send(msg interface{}) error
}
