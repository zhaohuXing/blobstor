package service

import (
	"crypto/sha1"
	"fmt"
	"log"

	. "github.com/zhaohuXing/blobstor/common"
	"github.com/zhaohuXing/blobstor/dao"
	"github.com/zhaohuXing/blobstor/lib"
	"github.com/zhaohuXing/blobstor/model"
)

func Register(verifyCode string, user *model.User) error {
	log.Printf("[info] Register Service process args:"+
		"code = %s, user = %+v", verifyCode, user)
	defer log.Println("[info] Register Service done")
	verify, err := lib.GetMessage(user.Phone)
	if err != nil {
		return ErrInternalError
	}

	if verify == "nil" || verify != verifyCode {
		log.Println("[error] the verification code does not match")
		return ErrNotMatchCodeError
	}

	isExist, err := dao.GetUserByPhone(user.Phone)
	if err != nil {
		log.Println("[error] exec GetUserByPhone failed")
		return ErrInternalError
	}
	if isExist != nil {
		log.Printf("[warn] <%s> User is exist", user.Phone)
		return ErrUserExistError
	}

	_, err = dao.InsertUser(user)
	if err != nil {
		log.Println("[error] exec InsertUser failed")
		return ErrInternalError
	}
	return nil
}

func Login(phone, password, addr string) error {
	log.Printf("[info] Login Service process args:"+
		"phone = %s, password = ******", phone)
	defer log.Println("[info] Login Service done")
	encryptedPwd := fmt.Sprintf("%x", sha1.Sum([]byte(password+phone)))
	_, err := dao.GetUserByAccount(phone, encryptedPwd)
	if err != nil {
		log.Println("[error] exec GetUserByAccount failed")
		return ErrInternalError
	}

	// Set User to Redis Session
	lib.SetSession(phone, addr)

	return nil
}
