# 簡易Golang套件管理系統

基於Golang程式語言環境，開發的一個簡易套件管理系統框架。

## 壹、基本說明
**一、目標：**
當接觸一門新的程式語言時，除了熟悉基本的輸入輸出、控制結構(如重複執行與邏輯判斷)等語法後，最自然的步驟便是探索該語言所內建提供的函式庫。每種程式語言擁有的函式庫各不相同，且在不同語言之間交錯使用下，往往容易產生混淆。也經常會出現這樣的情況：似乎曾經在某個地方寫過某個常用通用功能，但此時此刻卻無法立即找出其收藏位置，這時候便需要花費大量的時間來搜尋與查找，還不見得能找得到。本專案的核心目標正是為了幫助Golang的程式開發者，在軟體開發過程中能夠縮短開發時間，來快速搭建程式的基礎架構。另一方面，作為一名日常使用各類套件管理軟體的積極用戶，我也藉此機會深入了解這些技術的背景，進一步提升自己的實踐能力。

**二、開發環境：**
以下是開發該平台所採用的環境：
* 虛擬機：Docker
* 映像檔：golang
* 程式語言：Golang
* 程式編輯器：Visual Studio Code

**三、使用相依套件：**
以下是開發該平台所使用到的Golang套件：
* github.com/gorilla/mux（Server的Web應用程式架構）
* github.com/schollz/progressbar/v3（Client的進度bar）

**四、檔案說明：** 
此專案檔案主要分為兩個資料夾：server和client。其中，server資料夾為後端RESTful API平台的主要程式碼，client資料夾則存放使用者套件管理系統。接下來將對各資料夾中的檔案內容與功能進行詳細說明。
```bash
.
├── LICENSE
├── README.md
└── server  # RESTful API程式資料夾
      └── main.py  # 主程式
└── client  # client程式資料夾
      ├── main.py  # 主程式
      └── utils  # 副程式資料夾
            ├── initial.go  # 初始化副程式
            ├── install.go  # 安裝套件副程式
            ├── uninstall.go  # 反安裝套件副程式
            └── list.go  # 列出已安裝套件
```

## 貳、操作說明
**一、安裝程式方式：** 
安裝部分一樣分成兩個部分，Server與Client兩部分，以下將依序進行說明。

1. Server
- 步驟1：將server檔案夾中的程式，部署於一台可架設Server且有安裝Golang的電腦中。
- 步驟2：初始化與下載相依套件。
```bash
go mod init api
go get -u github.com/gorilla/mux
```
- 步驟3：運行server。
```bash
go run main.go
```
- 步驟4：請在同放置server資料夾的位置，創建一個Package資料夾，作為放置套件zip檔的地方。
- 步驟5：以下是此後端平台提供的RESTful API端點，`POST` /download，需在請求中加入以下資訊。
```json
{
  "API_KEY": "hGl^X$Lu8&4-s",
  "packageName": "套件名稱.go"
}
```
<br>

2. Client
- 步驟1：初始化與下載相依套件。
```bash
go mod init api
go get github.com/schollz/progressbar/v3
```
- 步驟2：編譯Golang，生成一個名為`pms`的執行檔。
```bash
go build -o pms main.go
```
- 步驟3：將執行檔放到`bin`資料夾
```bash
mv pms /usr/local/bin/  # 若發生是權限問題，請在最前面加上sudo
```
- 步驟4：設置環境路徑，使得系統可以在任何地方執行`pms`。開啟`.bashrc`或`.zshrc`配置檔(取決於你使用的 shell)，如果使用的是`bash`，需要編輯`~/.bashrc`文件；如果使用的是`zsh`，則是`~/.zshrc`文件。
```bash
nano ~/.bashrc  # 如果是 bash
nano ~/.zshrc   # 如果是 zsh
```
- 步驟5：添加`bin`資料夾到`$PATH`，在配置文件中，加入以下一行：
```bash
export PATH=$PATH:~/bin
```
- 步驟6：重新載入使用者的`.bashrc`或`.zshrc`配置文件。
```bash
source ~/.bashrc  # 或者 source ~/.zshrc
```
- 步驟6：完成上述步驟後，就可以在CLI中直接使用。
```bash
pms initial  # 初始化
pms install <package name>  # 安裝套件
pms uninstall <package name>  # 移除套件
pms list  # 套件列表
```

## 參、建議
**反思：** 若能在此基礎上新增不同用戶的檔案上傳功能，並針對用戶上傳的檔案實現公開與私人模式的區分，將能進一步提升專案的實用性與價值，使其更具意義與深遠影響。
