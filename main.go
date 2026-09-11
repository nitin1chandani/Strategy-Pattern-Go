package main

import (
	"fmt"
	. "go-strategy-pattern/internal/service/notification"
)

func main() {
	fmt.Println("Practicing Strategy Pattern")
	// service.Notify("", "Hello Folks")

	notify := Notify{
		Strategy: Email{},
	}
	message := "hello folks"
	notify.Send(message)
}
