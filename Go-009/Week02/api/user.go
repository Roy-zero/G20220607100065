package api

import (
	"G20220607100065/Go-009/Week02/common"
	"G20220607100065/Go-009/Week02/dao"
	"G20220607100065/Go-009/Week02/service"
	"errors"
	"log"
)

func QueryUser() {
	defer RecoverFunc()
	id := uint64(20220607100065)
	user, err := service.GetUser(id)

	if errors.Is(err, common.ErrNotFound) {
		SendErrMsg(err)
		return
	}
	SendSuccess(user)
}

func RecoverFunc() {
	if err := recover(); err != nil {
		SendErrMsg(err.(error))
	}
}

func SendErrMsg(err error) {
	if errors.Is(err, common.ErrNotFound) {
		log.Printf("stack: %+v", err)
		log.Printf("404 not found, %s \n", err)
		return
	}

	log.Printf("server error: %+v", err)
}

func SendSuccess(user *dao.User) {
	log.Printf("user found %v", user)
	return
}
