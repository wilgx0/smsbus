package smsbus

type SmsDevice interface {
	Send(phone string, message string) error
}

type Log interface {
	Error(message string)
}
