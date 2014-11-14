package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleStuff(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("TestPage.txt")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, string(body))
	fmt.Println(r.URL)
}
