package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (s *Server) Run(ctx context.Context) error {
	hs := &http.Server{
		Addr:           ":8080",
		Handler:        http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				rw.WriteHeader(400)
				return
			}
			var req struct {
				RPC  string            			`json:"rpc,omitempty"`
				Body map[string] interface{} 	`json:"body,omitempty"`
			} 

			if err := json.Unmarshal(body, &req); err != nil {
				fmt.Println(err)
				rw.WriteHeader(400)
				return
			}
			body, err = json.Marshal(req.Body)
			if err != nil {
				fmt.Println(err)
				rw.WriteHeader(400)
				return
			}
			switch req.RPC {
			case "sendmessage":
				err = s.SendMessage(body)
			default:
				rw.WriteHeader(404)
				return
			}
			if err != nil {
				fmt.Println(err)
				rw.WriteHeader(400)
				return
			}
		}),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return hs.ListenAndServe()
}