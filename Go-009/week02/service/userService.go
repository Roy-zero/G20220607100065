package service

import (
	"G20220607100065/Go-009/week02/dao"
	"fmt"
)

func GetUser(id uint64) (*dao.User, error) {
	user, err := dao.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Info: login complete. userID=%d, info=%+v \n", id, user)

	return user, nil
}
