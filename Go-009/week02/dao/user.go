package dao

import (
	"database/sql"
	"github.com/pkg/errors"

	"G20220607100065/Go-009/week02/common"
)

type userInfo struct {
	name string
}

// User info
type User struct {
	ID   uint64
	Info userInfo
}

func GetUserByID(id uint64) (*User, error) {
	// access DB...
	sqlQuery := "select id, name from user where id = ?"
	data, err := queryData(sqlQuery)
	if err == sql.ErrNoRows {
		return nil, errors.Wrapf(common.ErrNotFound, "sql: %s error: %v", sqlQuery, err)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "sql: %s error: %v", sqlQuery, err)
	}

	return data, nil
}

func queryData(sqlQuery string) (*User, error) {
	//user := User{123456, userInfo{"test"}}
	//return &user, nil
	return nil, sql.ErrNoRows
}
