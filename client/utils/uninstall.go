package utils

import (
	"fmt"
	"os"
)

func UninstallPackage(packageName string) error {
	// 確認當前路徑
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current directory: %v", err)
	}

	// 設定管理套件的資料夾
	folderPath := currentDir + "/package"
	packagePath := folderPath + "/" + packageName + ".go"

	// 檢查管理套件的資料夾是否存在
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return fmt.Errorf("unable to read package folder: %v", err)
	} else {
		// 刪除套件
		err := os.Remove(packagePath)
		if err != nil {
			return fmt.Errorf("failed to delete file %s: %v", packagePath, err)
		}
	}

	fmt.Println("Delete packages successful")
	return nil
}
