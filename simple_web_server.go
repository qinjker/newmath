package newmath

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello")
}

func ListenReq() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
