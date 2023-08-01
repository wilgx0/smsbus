package smsbus

type SmsBus struct {
	ch        chan SmsItem
	smsDevice SmsDevice
	log       Log
	workNum   int
}

type SmsItem struct {
	Phone   string
	Message string
}

func NewSmsBus(Options ...Option) *SmsBus {
	sb := &SmsBus{}
	for _, fn := range Options {
		fn(sb)
	}

	return sb
}

func (sb *SmsBus) Option(Options ...Option) {
	for _, fn := range Options {
		fn(sb)
	}
}

func (sb *SmsBus) Send(phone string, message string) {
	if phone == "" {
		sb.log.Error("phone is empty")
		return
	}
	sb.ch <- SmsItem{phone, message}
}

func (sb *SmsBus) smsDeviceSend(item SmsItem) {
	err := sb.smsDevice.Send(item.Phone, item.Message)
	if err != nil {
		sb.log.Error(err.Error())
	}
}

func (sb *SmsBus) Stop() {
	close(sb.ch)
}

func (sb *SmsBus) Start() {
	if sb.workNum == 0 {
		sb.workNum = 1
	}
	for i := 0; i < sb.workNum; i++ {
		go func() {
			for item := range sb.ch {
				sb.smsDeviceSend(item)
			}
		}()
	}
}
