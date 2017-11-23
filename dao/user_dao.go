package dao

import (
	"crypto/sha1"
	"fmt"
	"log"
	"sync"

	"github.com/zhaohuXing/blobstor/model"
)

func InsertUser(user *model.User) (int64, error) {
	log.Printf("[info] InsertUser process args: phone = %s, "+
		"nickname = %s ", user.Phone, user.Nickname)
	defer log.Println("[info] InsertUser done")
	failedErrorf := func(err error) {
		log.Printf("[error] exec InsertUser failed: %s", err)
	}
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	sql := "insert into users (phone, password, nickname, create_time)value(?,?,?,now())"
	stmt, err := Connect().Prepare(sql)
	if err != nil {
		failedErrorf(err)
		return -1, err
	}

	encryptedPwd := fmt.Sprintf("%x", sha1.Sum([]byte(user.Password+user.Phone)))
	res, err := stmt.Exec(user.Phone, encryptedPwd, user.Nickname)
	if err != nil {
		failedErrorf(err)
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		failedErrorf(err)
		return -1, err
	}
	return id, nil
}

func GetUserByAccount(phone, password string) (*model.User, error) {
	log.Printf("[info] GetUserByAccount process args: phone = %s, password = ******", phone)
	defer log.Println("[info] GetUserByAccount done")
	failedErrorf := func(err error) {
		log.Printf("[error] exec GetUserByAccount failed: %s", err)
	}
	sql := "select * from users where phone = ? and password = ?"
	rows, err := Connect().Query(sql, phone, password)
	if err != nil {
		failedErrorf(err)
		return nil, err
	}
	defer rows.Close()

	user := &model.User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Phone, &user.Password,
			&user.Nickname, &user.Role, &user.CreateTime)
		if err != nil {
			failedErrorf(err)
			return nil, err
		}
	}
	return user, nil
}

func GetUserByPhone(phone string) (*model.User, error) {
	log.Printf("[info] GetUserByPhone process args: phone = %s", phone)
	defer log.Println("[info] GetUserByPhone done")
	failedErrorf := func(err error) {
		log.Printf("[error] exec GetUserByPhone failed: %s", err)
	}
	sql := "select * from users where phone = ?"
	rows, err := Connect().Query(sql, phone)
	if err != nil {
		failedErrorf(err)
		return nil, err
	}
	defer rows.Close()

	var user *model.User
	for rows.Next() {
		user = &model.User{}
		err = rows.Scan(&user.Id, &user.Phone, &user.Password,
			&user.Nickname, &user.Role, &user.CreateTime)
		if err != nil {
			failedErrorf(err)
			continue
		}
	}
	return user, nil
}

func UpdateUserByAccount(field, value, phone, password string) (int64, error) {
	log.Printf("[info] UpdateUserByAccount process args: field = %s,"+
		"value = %s, phone = %s, password = ******", field, value, phone)
	defer log.Println("[info] UpdateUserByAccount done")
	failedErrorf := func(err error) {
		log.Printf("[error] exec UpdateUserByAccount failed: %s", err)
	}
	sql := fmt.Sprintf("update users set %s = ? where phone = ? and password = ?", field)
	stmt, err := Connect().Prepare(sql)
	if err != nil {
		failedErrorf(err)
		return -1, err
	}

	res, err := stmt.Exec(value, phone, password)
	if err != nil {
		failedErrorf(err)
		return -1, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		failedErrorf(err)
		return -1, err
	}
	return affect, nil
}

func DeleteUserById(id int64) (int64, error) {
	log.Printf("[info] DeleteUserById process args: id = %d", id)
	defer log.Println("[info] DeleteUserById done")
	failedErrorf := func(err error) {
		log.Printf("[error] exec DeleteUserById failed: %s", err)
	}
	sql := "delete from users where id = ?"
	stmt, err := Connect().Prepare(sql)
	if err != nil {
		failedErrorf(err)
		return -1, err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		failedErrorf(err)
		return -1, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		failedErrorf(err)
		return -1, err
	}
	return affect, nil
}

func GetUsers() ([]*model.User, error) {
	log.Println("[info] GetUsers process")
	defer log.Println(" [info] GetUsers done")
	failedErrorf := func(err error) {
		log.Printf("[error] exec Users failed: %s", err)
	}
	users := make([]*model.User, 0)
	sql := "select id, phone, nickname, create_time from users"
	rows, err := Connect().Query(sql)
	if err != nil {
		failedErrorf(err)
		return nil, err
	}

	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.Id, &user.Phone, &user.Nickname, &user.CreateTime)
		if err != nil {
			failedErrorf(err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
