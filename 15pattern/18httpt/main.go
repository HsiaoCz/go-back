package main

import (
	"encoding/json"
	"net/http"
)

// http 方式的远程调用
// 错了，这还不是
type Add struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result struct {
	R int    `json:"r"`
	E string `json:"e"`
}

func main() {
	http.HandleFunc("/get/add", HandleAddSomeThing)
}

func HandleAddSomeThing(w http.ResponseWriter, r *http.Request) {
	var add Add
	err := json.NewDecoder(r.Body).Decode(&add)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result := Result{
		R: AddSomeThing(add.A, add.B),
		E: "",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func AddSomeThing(a int, b int) int {
	return a + b
}
