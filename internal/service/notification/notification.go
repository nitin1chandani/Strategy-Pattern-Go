package notification

import "fmt"

//raw manner
// func Notify(notificationType, message string) {
// 	if notificationType == string(constants.MAIL) {
// 		fmt.Println(message)
// 	}
// 	if notificationType == string(constants.SMS) {
// 		fmt.Println(message)
// 	}
// 	if notificationType == string(constants.PUSH_NOTIFICATION) {
// 		fmt.Println(message)
// 	}
// }

// Strategy
// common contract for each notification type
type NotificationStrategy interface {
	Send(message string)
}

// concrete strategy
type Email struct{}

func (e Email) Send(message string) {
	fmt.Println(message)
}

// concrete strategy
type SMS struct{}

func (s SMS) Send(message string) {
	fmt.Println(message)
}

// concrete strategy
type Push struct{}

func (p Push) Send(message string) {
	fmt.Println(message)
}

type Notify struct {
	Strategy NotificationStrategy
}

func (n Notify) Send(message string) {
	n.Strategy.Send(message)
}
