package smsbus

import "sync"

var instance *SmsBus
var once sync.Once

func GetInstance() *SmsBus {
	once.Do(func() {
		instance = NewSmsBus()
	})
	return instance
}
