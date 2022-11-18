package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 定義標準返回結果類型結果類型的結構體
type ResponseResult struct {
	Result string `json:"result"`
}

// 定義設置共同的Http response Header 訊息的函數
func setSameHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// 定義用來處理跟請求的函數
func index(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	fmt.Fprintf(w, "Hello Go REST API Server")
}

// 定義基礎的CRUD函數, 並調用setSameHeader()函數使CRUD函數能返回JSON類型的資料
func listBooks(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "listBooks",
	})
}
func getBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "getBook",
	})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "createBook",
	})
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "updateBook",
	})
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "deleteBook",
	})
}

func main() {
	fmt.Println(">>Hello Welcome to Go REST API Server<<")

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/books", listBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("服務器運行在端口 :8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
