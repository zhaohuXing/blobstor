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

func Register(code string, user *model.User) error {
	log.Printf("[info] Register Service process args:"+
		"code = %s, user = %+v", code, user)
	defer log.Println("[info] Register Service done")
	isExist, err := isExistUserWithVerify(user.Phone, code)
	if err != nil {
		return err
	}
	if isExist {
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

func PwdReset(phone, password, code string) error {
	log.Printf("[info] PwdPassword Service process args:"+
		"phone = %s, verify = %s", phone, code)
	defer log.Println("[info] PwdPassword Service done")
	isExist, err := isExistUserWithVerify(phone, code)
	if err != nil {
		return err
	}

	if !isExist {
		log.Printf("[warn] <%s> User is not exist", phone)
		return ErrUserNotExistError
	}

	err = dao.UpdatePassword(phone, password)
	if err != nil {
		log.Println("[error] Update password failed")
		return ErrInternalError
	}
	return nil
}

func isExistUserWithVerify(phone, code string) (bool, error) {
	verify, err := lib.GetMessage(phone)
	if err != nil {
		return false, ErrInternalError
	}

	if verify == "nil" || verify != code {
		log.Println("[error] the verification code does not match")
		return false, ErrNotMatchCodeError
	}

	isExist, err := dao.GetUserByPhone(phone)
	if err != nil {
		log.Println("[error] exec GetUserByPhone failed")
		return false, ErrInternalError
	}
	if isExist != nil {
		log.Printf("[warn] <%s> User is exist", phone)
		return true, nil
	}
	return false, nil
}
