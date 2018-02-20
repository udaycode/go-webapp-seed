package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"encoding/json"
)

type Resp struct {
    Error bool
    Msg string
}

func HomeH(w http.ResponseWriter, r *http.Request) {
	resp := Resp{false, ""}
    b, _ := json.Marshal(resp)
    w.Write(b)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeH)
    http.Handle("/", r)
    fmt.Println("Listening")
    http.ListenAndServe(":8000", r)
}