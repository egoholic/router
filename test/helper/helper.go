package helper

import (
	"fmt"
	"net/http"

	"github.com/egoholic/router/params"
)

func TestHandlerFunc(w http.ResponseWriter, r *http.Request, p *params.Params) {
	w.WriteHeader(200)
	w.Header().Add("TEST-HEADER", p.Param("header")[0])
	_, err := w.Write([]byte("hello!"))
	if err != nil {
		fmt.Printf("\n\t\tERROR: %s\n\n", err.Error())
	}
	var body = []byte{}
	r.Body.Read(body)
	w.Write(body)
}
