package adapter

import (
	"fmt"
	"net/http"
)

type HttpAdapter struct {
	addr string
}

func NewHttpAdapter(addr string) *HttpAdapter {
	return &HttpAdapter{
		addr: addr,
	}
}

func (a *HttpAdapter) Listen(ch chan NotifyData) {
	http.HandleFunc("/reg", func(w http.ResponseWriter, req *http.Request) {
		len := req.ContentLength
		body := make([]byte, len)
		req.Body.Read(body)

		obj, err := ParseRegPayload(string(body))
		if err == nil {
			ch <- obj
		}

		fmt.Fprintln(w, "OK")
	})

	httpsvr := http.Server{
		Addr: a.addr,
	}

	httpsvr.ListenAndServe()
}
