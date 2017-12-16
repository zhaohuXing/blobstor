package ctrls

import (
	"log"
	"net/http"

	. "github.com/zhaohuXing/blobstor/common"
	"github.com/zhaohuXing/blobstor/model"
	"github.com/zhaohuXing/blobstor/service"
)

func Register(w http.ResponseWriter, r *http.Request) {
	log.Printf("[info] %s /register", r.Method)
	defer log.Println("[info] Register controller done")
	var response *Response
	switch r.Method {
	case "POST":
		type Val struct {
			Phone    string `json:"phone"`
			Password string `json:"password"`
			Nickname string `json:"nickname"`
			Verify   string `json:"verify"`
		}
		var val Val
		var user *model.User
		err := LoadBody(r.Body, &val)
		log.Printf("[info] Register controller process args: %+v", val)
		if err != nil || val.Phone == "" || val.Password == "" ||
			val.Nickname == "" || val.Verify == "" {
			log.Println("[error] The Register argument is invalid")
			response = ResponseFactory[ErrInvalidArgumentError]
		} else {
			user = &model.User{
				Phone:    val.Phone,
				Password: val.Password,
				Nickname: val.Nickname,
			}
			err = service.Register(val.Verify, user)
			if err == nil {
				response = ResponseFactory[HTTP_OK_CREATED]
			} else {
				response = ResponseFactory[err]
			}
		}
	default:
		log.Println("[warn] Not supported %s", r.Method)
		response = ResponseFactory[HTTP_METHOD_NOT_ALLOWED]
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response.ToJson()))
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("[info] %s /login", r.Method)
	defer log.Println("[info] Login controller done")
	var response *Response
	switch r.Method {
	case "POST":
		type Val struct {
			Phone    string `json:"phone"`
			Password string `json:"password"`
		}
		var val Val
		err := LoadBody(r.Body, &val)
		if err != nil || val.Phone == "" || val.Password == "" {
			log.Println("[error] The Login argument is invalid")
			response = ResponseFactory[ErrInvalidArgumentError]
			goto JSONEND
		}
		err = service.Login(val.Phone, val.Password, r.RemoteAddr)
		if err == nil {
			response = ResponseFactory[HTTP_OK]
		} else {
			response = ResponseFactory[err]
		}
	default:
		log.Println("[warn] Not supported %s", r.Method)
		response = ResponseFactory[HTTP_METHOD_NOT_ALLOWED]
	}
JSONEND:
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response.ToJson()))
}

func PwdReset(w http.ResponseWriter, r *http.Request) {
	log.Printf("[info] %s /password/reset", r.Method)
	defer log.Println("[info] PwdReset controller done")
	var response *Response
	switch r.Method {
	case "POST":
		type Val struct {
			Phone    string `json:"phone"`
			Password string `json:"password"`
			Verify   string `json:"verify"`
		}
		var val Val
		err := LoadBody(r.Body, &val)
		if err != nil || val.Phone == "" || val.Password == "" || val.Verify == "" {
			log.Println("[error] The PwdReset argument is invalid")
			response = ResponseFactory[ErrInvalidArgumentError]
			goto JSONEND
		}
		err = service.PwdReset(val.Phone, val.Password, val.Verify)
		if err == nil {
			response = ResponseFactory[HTTP_OK]
		} else {
			log.Println("[error] password reset failed")
			response = ResponseFactory[err]
		}
	default:
		log.Println("[warn] Not supported %s", r.Method)
		response = ResponseFactory[HTTP_METHOD_NOT_ALLOWED]
	}
JSONEND:
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response.ToJson()))
}
