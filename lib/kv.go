package kv

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	. "github.com/zhaohuXing/blobstor/common"

	"github.com/go-redis/redis"
	"github.com/yunpian/yunpian-go-sdk/sdk"
)

var op = &redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
}

func SendMessage(phone string) error {
	log.Printf("[info] SendMessage process args: phone = %s", phone)
	defer log.Println("[info] SendMessage done")
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	err := setValue(phone, vcode, 60*10*time.Second)
	log.Printf("[info] Send code <%s> to %s", vcode, phone)
	if err != nil {
		log.Printf("[error] The verify message store failed: %s", err)
		return ErrInternalError
	}

	producer := sdk.New("ab27823f4e2de4a556619e21d3a00594")
	param := sdk.NewParam(2)
	param[sdk.MOBILE] = phone
	param[sdk.TEXT] = fmt.Sprintf("【BStor云存储】(%s)"+
		" 注册/登录验证码10分钟内有效，请尽快完成验证。", vcode)
	producer.Sms().SingleSend(param)
	return nil
}

func GetMessage(phone string) (string, error) {
	log.Printf("[info] GetMessage process args: phone = %s", phone)
	defer log.Println("[info] GetMessage done")
	val, err := getValue(phone)
	if err != nil {
		log.Println("[error] the verify message get failed from redis")
		return "", err
	}
	return val, nil
}

func SetSession(phone, addr string) error {
	log.Printf("[info] SetSession process args: phone = %s, addr = %s", phone, addr)
	defer log.Println("[info] SetSession done")

	err := setValue(phone+addr, "true", 5*60*60*time.Second)
	if err != nil {
		log.Printf("[error] The session set failed: %s", err)
		return ErrInternalError
	}

	return nil
}

func GetSession(phone, addr string) error {
	log.Printf("[info] GetSession process args: phone = %s, addr = %s", phone, addr)
	defer log.Println("[info] GetSession done")
	val, err := getValue(phone + addr)
	if err != nil {
		log.Println("[error] The session get failed from redis")
		return "", err
	}
	return val, nil
}

func getValue(key string) (string, error) {
	client := redis.NewClient(op)
	defer client.Close()
	val, err := client.Get(key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return "nil", nil
		}
		return "", err
	}
	return val, nil
}

func setValue(key, value string, expired time.Duration) error {
	client := redis.NewClient(op)
	defer client.Close()
	err := client.Set(key, value, expired).Err()
	if err != nil {
		return err
	}
	return nil
}
