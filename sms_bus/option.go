package smsbus

type Option func(*SmsBus)

func WithChNum(i int) Option {
	return func(bus *SmsBus) {
		bus.ch = make(chan SmsItem, i)
	}
}

func WithSmsDevice(d SmsDevice) Option {
	return func(bus *SmsBus) {
		bus.smsDevice = d
	}
}

func WithLog(l Log) Option {
	return func(bus *SmsBus) {
		bus.log = l
	}
}

func WithWorkNum(i int) Option {
	return func(bus *SmsBus) {
		bus.workNum = i
	}
}
