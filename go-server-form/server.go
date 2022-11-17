package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	// 向用戶端輸出內容
	fmt.Fprintf(w, "Hello Go World URL!")

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 只允許POST請求
	if r.Method != "POST" {
		http.Error(w, "方法不支持", http.StatusNotFound)
	}
	// 解析表單失敗報error
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	name := r.FormValue("name")
	pwd := r.FormValue("pwd")

	fmt.Fprintf(w, "用戶名：%s\n", name)
	fmt.Fprintf(w, "密碼：%s\n", pwd)
}

func main() {
	fmt.Println(">>Welcome to Go Server<<")

	// 實例化一個文件服務器, 且讓文件服務器使用本地文件夾./static
	fileServer := http.FileServer(http.Dir("./static"))
	// ①在使用者請求/目錄時響應調用這個文件服務器
	http.Handle("/", fileServer)

	// ②在使用者請求/hello目錄時響應調用 helloHandler HandleFunc
	http.HandleFunc("/hello", helloHandler)

	// ③在使用者請求/login目錄時響應調用 loginHandler HandleFunc
	http.HandleFunc("/login", loginHandler)

	fmt.Printf("服務器運行在端口 :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
