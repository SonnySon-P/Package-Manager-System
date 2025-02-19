package utils

import (
	"fmt"
	"os"
	"strings"
)

func ListPackages() error {
	// 確認當前路徑
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current directory: %v", err)
	}

	// 設定管理套件的資料夾
	folderName := "package"
	folderPath := currentDir + "/" + folderName

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("unable to read package folder: %v", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no packages installed")
	}

	// 若有套件則列出
	fmt.Println("Install package:")
	for _, file := range files {
		fmt.Println(removeExtension(file.Name()))
	}
	return nil
}

// 刪除副檔名
func removeExtension(filename string) string {
	extIndex := strings.LastIndex(filename, ".")

	// 如果找不到".", 則返回原始檔名
	if extIndex == -1 {
		return filename
	}

	// 返回副檔名之前的部分
	return filename[:extIndex]
}
