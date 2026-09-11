package service

import (
	"fmt"
	constants "go-strategy-pattern/internal"
)

func Notify(notificationType, message string) {
	if notificationType == string(constants.MAIL) {
		fmt.Println(message)
	}
	if notificationType == string(constants.SMS) {
		fmt.Println(message)
	}
	if notificationType == string(constants.PUSH_NOTIFICATION) {
		fmt.Println(message)
	}
}
