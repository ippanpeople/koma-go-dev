package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go REST API")
}

func main() {
	fmt.Println(">>Hello Welcome to Go REST API SERVER<<")

	// 從mux模組實例化一個路由
	router := mux.NewRouter()
	// 當使用者請求跟資源時 回應一個index函式並只響應GET方法
	router.HandleFunc("/", index).Methods("GET")

	fmt.Println("無誤器運行在端口： 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
