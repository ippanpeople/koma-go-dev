# REST API - 建立微服務工程

## 步驟
1. 建立微服務工程
2. 編寫Go語言程序,初始化微服務
3. 運行調適

## 實作
- 建立工程目錄
```bash
mkdir go-restapi
```
- 進入目錄
```bash
cd go-restapi
```
- 初始化工程模塊
```bash
go mod github.com/ippanpeople/go-restapi
```
- 安裝依賴模塊(路由模塊)
```bash
go get -u github.com/gorilla/mux
```
- 清理工程模塊
```bash
go mod tidy
```
- 編寫服務器程序
```bash
vim server.go
```
- 運行程序
```bash
go run server.go
```
- 編譯程序
```bash
go build
```
-運行程序
```bash
./go-server
```