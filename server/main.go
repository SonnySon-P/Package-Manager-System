package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// 定義client請求的資料結構
type RequestData struct {
	API_KEY     string `json:"API_KEY"`
	PackageName string `json:"packageName"`
}

// POST /download
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 檢查Content-Type是否為application/json
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid content type, expected applicçation/json", http.StatusBadRequest)
		return
	}

	// 解析JSON請求體
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Error parsing JSON request", http.StatusBadRequest)
		return
	}

	// 檢查前端傳來的API_KEY是否為null
	if requestData.API_KEY == "" {
		http.Error(w, "API KEY is required", http.StatusBadRequest)
		return
	}

	// 驗證金鑰是否正確
	if requestData.API_KEY != "hGl^X$Lu8&4-s" {
		http.Error(w, "API KEY is error", http.StatusForbidden)
		return
	}

	// 檢查前端傳來的套件名稱是否為null
	packageName := requestData.PackageName
	if packageName == "" {
		http.Error(w, "PackageName is required", http.StatusBadRequest)
		return
	}

	// 檢查套件是否存在
	filePath := fmt.Sprintf("../Package/%s.zip", packageName)
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		http.Error(w, "Package not found", http.StatusNotFound)
		return
	}

	// 設置回應標頭，告訴前端要傳送一個zip檔案
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", packageName))

	// 開啟壓縮檔案並將內容寫入回應
	zipFile, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer zipFile.Close()

	// 使用io.Copy將檔案內容寫入到回應中
	_, err = io.Copy(w, zipFile)
	if err != nil {
		http.Error(w, "Error sending file", http.StatusInternalServerError)
		return
	}
}

func main() {
	// 初始化mux
	r := mux.NewRouter()

	// 創建router和對應函數
	r.HandleFunc("/download", downloadHandler).Methods("POST")

	// 設定server和port
	http.Handle("/", r)
	port := ":8095"
	fmt.Printf("Server starting on port %s...\n", port)

	// 啟動server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
