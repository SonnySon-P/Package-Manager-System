package utils

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/schollz/progressbar/v3"
)

// 定義POST請求的資料結構
type requestData struct {
	API_KEY     string `json:"API_KEY"`
	PackageName string `json:"packageName"`
}

// 安裝套件
func InstallPackage(packageName string) error {
	// 確認當前路徑
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current directory: %v", err)
	}

	// 設定管理套件的資料夾和檔案名稱，及套件名稱
	folderPath := currentDir + "/package"
	packageName = packageName + ".go"
	packagePath := folderPath + "/" + packageName
	zipPackagPath := folderPath + "/" + packageName + ".zip"

	// 檢查套件是否已存在
	if _, err := os.Stat(packagePath); err == nil {
		return fmt.Errorf("package already exists: %v", packagePath)
	}

	// 設定下載的URL
	zipPackagName := packageName + ".zip"
	url := "http://PMS-RESTfulAPI:8095/download"

	// 設定請求的資料
	requestData := requestData{
		API_KEY:     "hGl^X$Lu8&4-s",
		PackageName: packageName,
	}

	// 開始下載檔案
	fmt.Println("Starting download...")
	err = downloadFile(url, folderPath, zipPackagName, requestData)
	if err != nil {
		return err
	}

	// 解壓縮檔案
	err = unzipFile(zipPackagPath, packagePath)
	if err != nil {
		return err
	}

	// 成功下載
	fmt.Println("Package downloaded successfully")
	return nil
}

func downloadFile(url string, folderPath string, fileName string, requestData requestData) error {
	// 序列化requestData為JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("error marshaling request data: %v", err)
	}

	// 發送POST請求
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// 確認server回應狀態是200
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", response.Status)
	}

	// 獲取回應內容的大小
	contentLength := response.ContentLength
	if contentLength <= 0 {
		return fmt.Errorf("invalid content length: %d", contentLength)
	}

	// 創建進度條
	bar := progressbar.NewOptions(int(contentLength),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetDescription("Downloading"),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionSetRenderBlankState(true),
	)

	// 使用io.TeeReader讀取並更新進度條
	teeReader := io.TeeReader(response.Body, bar)

	// 創建本地端一個空檔案
	outFile, err := os.Create(filepath.Join(folderPath, fileName))
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 將server回應寫入檔案
	_, err = io.Copy(outFile, teeReader)
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}

// 解壓縮檔案
func unzipFile(zipPackagName, packagePath string) error {
	// 打開zip檔案
	zipReader, err := zip.OpenReader(zipPackagName)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	// 取得檔案
	file := zipReader.File[0]

	// 打開zip中的檔案
	fileInZip, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file inside zip: %v", err)
	}
	defer fileInZip.Close()

	// 創建目標檔案
	outFile, err := os.Create(packagePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	// 解壓縮檔案內容到目標檔案
	_, err = io.Copy(outFile, fileInZip)
	if err != nil {
		return fmt.Errorf("failed to extract file: %v", err)
	}

	// 刪除壓縮檔
	err = os.Remove(zipPackagName)
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %v", zipPackagName, err)
	}

	return nil
}
