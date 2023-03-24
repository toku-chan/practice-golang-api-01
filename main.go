package main

import (
	"fmt"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func main() {
	// https://pkg.go.dev/net/http

	// HandleFunc
	// 第1引数にpath
	// 第2引数にhttpのres, reqを扱ったFunctionを持つっぽい
	// pathに対して、どういう処理を登録するかって感じかな
	http.HandleFunc("/", helloHandler)

	// 第1引数にポート番号を登録
	// 第2引数はハンドラーらしいが、nilを公式でも指定しているので、とりあえず真似てみる
	http.ListenAndServe(":8080", nil)
}
