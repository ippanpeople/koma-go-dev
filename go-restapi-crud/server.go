package main

import "net/http"

// 定義標準返回結果類型結果類型的結構體
type ResponseResult struct {
	Result string `json:"result"`
}

// 設置共同的Http response Header 訊息
func setSameHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
func main() {

}
