# REST API - 建立 CRUD API 服務接口 -> 響應 GET 獲取全部書籍/ 根據id獲取指定書籍

## 步驟
1. 建立微服務工程
2. 編寫Go語言程序,初始化微服務
3. 建立CRUD接口
4. 編寫響應GET /books 獲取全部書籍
5. 編寫透過URL獲取書籍id參數的方法
6. 運行調適

## 實作
- 建立工程目錄
```bash
mkdir go-restapi-get
```
- 進入目錄
```bash
cd go-restapi-get
```
- 初始化工程模塊
```bash
go mod github.com/ippanpeople/go-restapi-get
```
- 安裝依賴模塊(路由模塊)
```bash
go get -u github.com/gorilla/mux
```
- 創建服務器程序
```bash
touch server.go
```
- 編寫CRUD接口
```bash
vim server.go
```
- 編寫GET方法函數
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