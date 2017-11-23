package router

import (
	"net/http"

	"github.com/zhaohuXing/blobstor/ctrls"
)

func init() {
	http.HandleFunc("/verify", ctrls.SendMessage)
	http.HandleFunc("/register", ctrls.Register)
	http.HandleFunc("/login", ctrls.Login)
}
