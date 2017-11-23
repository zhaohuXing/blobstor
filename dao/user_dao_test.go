package dao

import (
	"fmt"
	"testing"

	"github.com/zhaohuXing/blobstor/model"
)

var users = []*model.User{
	&model.User{
		Phone:    "17853551517",
		Password: "xxx*sfhyit",
		Nickname: "testname1",
	},
	&model.User{
		Phone:    "17853518518",
		Password: "xxx*sfhyit",
		Nickname: "testname2",
	},
}

func TestInsertUsers(t *testing.T) {
	for _, user := range users {
		InsertUser(user)
	}
}

func TestGetUsers(t *testing.T) {
	ret, err := GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v, %+v", ret[0], ret[1])
}

func TestInsertUser(t *testing.T) {
	user := &model.User{
		Phone:    "17862834237",
		Password: "xxx*sdf8",
		Nickname: "Joe",
	}
	InsertUser(user)
}

func TestGetUserByAccount(t *testing.T) {
	user, err := GetUserByAccount("17862812345", "xxx*sdf8")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)
}

func TestUpdateUserByAccount(t *testing.T) {
	phone := "17862812345"
	password := "xxx*sdf8"
	newPassword := "root123456"
	UpdateUserByAccount("password", newPassword, phone, password)
}

func TestDeleteUserById(t *testing.T) {
	num, err := DeleteUserById(int64(1000))
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}
