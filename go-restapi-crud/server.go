package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
func updatBook(w http.ResponseWriter, r *http.Request) {
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

}
