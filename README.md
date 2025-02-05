# 自創Package Manager程式

**安裝方式**
步驟一：將程式編譯成一個名為 mycli 的可執行檔案。
```shell
go build -o mycli main.go
```
步驟二：請將名為 mycli 的可執行檔案，移動可執行檔到 bin 目錄（通常是 /usr/local/bin/ 或 /usr/bin/，具體位置依你系統配置而定）
```shell
sudo mv mycli /usr/local/bin/
```
步驟三：確保 PATH 變數包含 /usr/local/bin/。你可以在終端機中輸入以下命令檢查：
```shell
echo $PATH
```
步驟四：確保 PATH 變數包含 /usr/local/bin/。你可以在終端機中輸入以下命令檢查：

在 CLI 上執行程式
完成上述步驟後，你就可以在命令行中直接使用 mycli 指令來執行你的程式了。例如：
