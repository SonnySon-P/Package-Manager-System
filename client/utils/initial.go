package utils

import (
	"fmt"
	"os"
)

func InitialProject() error {
	// 確認當前路徑
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current directory: %v", err)
	}

	// 設定管理套件的資料夾
	folderPath := currentDir + "/package"

	// 檢查管理套件的資料夾是否存在
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			return fmt.Errorf("unable to create folder: %v", err)
		}
	} else {
		fmt.Println("folder already exists:", folderPath)
	}

	fmt.Println("Project setup successful")
	return nil
}
