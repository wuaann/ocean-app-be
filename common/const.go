package common

import "log"

const (
	DbTypePost = 1
	DbTypeUser = 2
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recover error:", err)
	}
}

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
