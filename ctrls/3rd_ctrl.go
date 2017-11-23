package ctrls

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/zhaohuXing/blobstor/common"
	"github.com/zhaohuXing/blobstor/lib"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	log.Printf("[info] %s /verify", r.Method)
	defer log.Println("[info] 3rd controller done")
	var response *Response
	var err error
	if r.Method == "POST" {
		type Val struct {
			Phone string `json:"phone"`
		}
		var val Val
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&val)
		if err != nil || val.Phone == "" {
			log.Println("[error] The verify argument is invalid")
			response = ResponseFactory[ErrInvalidArgumentError]
			goto END
		}
		err = lib.SendMessage(val.Phone)
		if err == nil {
			response = ResponseFactory[HTTP_OK]
		} else {
			response = ResponseFactory[err]
		}
	END:
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response.ToJson()))
	}
}
