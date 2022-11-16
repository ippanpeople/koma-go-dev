package main

import (
	"fmt"
	"log"
	"net/http"
)

// URL響應函數, w:通過w給用戶進行響應  r:通過r接收用戶請求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 排除 /hello 以外請求
	if r.URL.Path != "/hello" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	// 只允許GET請求
	if r.Method != "GET" {
		http.Error(w, "方法不支持", http.StatusNotFound)
		return
	}
	// 像用戶端輸出內容
	fmt.Fprintf(w, "Hello Go World URL!")
}

func main() {
	fmt.Println(">>Hello Welcome to Go Server<<")

	// 實例化一個文件服務器, 且讓文件服務器使用本幣文件夾./static
	fileServer := http.FileServer(http.Dir("./static"))
	// 當訪問根（/）時響應這個文件服務器給使用者
	http.Handle("/", fileServer)

	// 響應 /hello 資源請求
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("服務器運行在端口 :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
