package main

import (
	"fmt"
	"os"
	"pms/utils"
)

func main() {
	// 檢查是否有action參數
	if len(os.Args) < 2 {
		fmt.Println("Error: action is required (initial, install, uninstall, list)")
		os.Exit(1)
	}

	// 根據action執行不同的任務
	switch os.Args[1] {
	case "initial":
		err := utils.InitialProject()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "install":
		if len(os.Args) < 3 {
			fmt.Println("Usage: install <package name>")
			os.Exit(1)
		}
		err := utils.InstallPackage(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "uninstall":
		if len(os.Args) < 3 {
			fmt.Println("Usage: uninstall <package name>")
			os.Exit(1)
		}
		err := utils.UninstallPackage(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "list":
		err := utils.ListPackages()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Error: invalid action. Choices are (initial, install, uninstall, list)")
		os.Exit(1)
	}
}
