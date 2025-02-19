# 簡易Golang套件管理系統

基於Golang程式語言環境開發，所開發的簡易套件管理系統框架。

## 壹、基本說明
**一、目標：**
當接觸一門新的程式語言時，除了熟悉基本的輸入輸出、控制結構（如重複執行與邏輯判斷）等語法後，最自然的步驟便是探索該語言所內建提供的函式庫。每種程式語言擁有的函式庫各不相同，且在不同語言之間交錯使用時，往往容易產生混淆。經常會出現這樣的情況：似乎曾經在某個地方寫過某個功能，但此刻卻無法立即找出其位置，這時候便需要花費大量的時間來搜尋與查找。本專案的核心目標正是為了幫助Golang的使用者，在軟體開發過程中能夠縮短開發時間，來快速搭建程式的基礎架構。另一方面，作為一名日常使用各類套件管理軟體的用戶，我也藉此機會深入了解這些技術的背景，進一步提升自己的實踐能力。

**二、開發環境：**
以下是開發該平台所採用的環境：
* 虛擬機：Docker
* 映像檔：golang
* 程式語言：Golang
* 程式編輯器：Visual Studio Code

**三、檔案說明：** 
此專案檔案（指coding這個資料夾）主要分為兩個資料夾：nodejs和tests。其中，nodejs資料夾為後端平台的主要程式碼，tests資料夾則存放使用jest框架進行的單元測試。接下來將對各資料夾中的檔案內容進行詳細說明。
```bash
.
├── LICENSE
├── README.md
└──  code  # 開發程式資料夾
      ├── main.py  # 主程式
      ├── readFile.py  # 讀取組語模組
      ├── RV32IMemory.py  # 模擬memory模組
      ├── RV32IRegisters.py  # 模擬register模組
      ├── cpuCore.py  # 模擬CPU模組
      ├── instructionTyple.py  # instruction與Typle的對應模組
      ├── rType.py  # 模擬R-Type instruction運行模組
      ├── iType.py  # 模擬I-Type instruction運行模組
      ├── sType.py  # 模擬S-Type instruction運行模組
      ├── bType.py  # 模擬B-Type instruction運行模組
      ├── uType.py  # 模擬U-Type instruction運行模組
      ├── jType.py  # 模擬J-Type instruction運行模組
      ├── otherType.py  # 模擬ecall instruction運行模組
      └── try.asm  # 測試檔案
```

## 貳、操作說明
**一、安裝程式方式：** 
將一個編譯好的執行檔放置到`bin`資料夾並設置好環境路徑，步驟如下：
***步驟1: 編譯Golang，生成一個名為`vm`的執行檔。
```bash
go build -o vm main.go
```

***步驟 2: 將執行檔放到`bin`資料夾
請在UNIX類系統(如Linux或macOS)中，將執行檔放到`/usr/local/bin`或`~/bin` 通常是用來存放可執行檔的目錄。
```bash
mkdir -p ~/bin
mv vm ~/bin/
```

***步驟 3: 設置環境路徑
接下來，您需要設置您的環境變數，使得系統可以找到您放置的`bin`資料夾。使得可以在任何地方執行`vm`。

1. 開啟`.bashrc`或`.zshrc`配置檔(取決於您使用的 shell)，如果使用的是`bash`，需要編輯`~/.bashrc`文件；如果使用的是`zsh`，則是`~/.zshrc`文件。
```bash
nano ~/.bashrc  # 如果是 bash
# 或者
nano ~/.zshrc   # 如果是 zsh
```

2. 添加`bin`資料夾到`$PATH`，在配置文件中，加入以下一行：
```bash
export PATH=$PATH:~/bin
```

**二、運行程式方式：**
```bash
vm -load <path>  #  載入映像檔，將指定的映像檔解壓到容器目錄中
vm -save <path>  #  將當前容器保存為新映像檔
vm -create <path>  #  從映像檔創建容器
vm -start <containerID>  #  啟動容器
vm -stop <containerID>  #  停止容器
vm -delete <containerID>  #  刪除容器
```





# 自創Golang Package Manager程式

在本專案主要想達成的目標為，讓客服端的Golang的使用者，在開發軟體過程中，若希望能簡化開發時間，採倚靠已寫好的各種不同功能的副程式套件，作為撰寫程式的基礎。

客服端的部分
使用者可以在已有安裝Golang編譯環境的電腦進行使用，當安裝完本套件後，可以進行註冊 

RESTful API的部分
關於這一部分所提供的功能主要有以下三部分，第一為提供Package的下載功能。第二為查詢套件、第三為上傳套件。每個人都有自己私人的套件資料庫，當然若有人願意分享，也可以將檔案放置公共區之中，供所有用戶可以進行查詢與下載。在後端伺服器管理上，主要是透過一個MongoDB的資料庫來達成，每一個Package在資料庫中都記載著該名稱、上傳者、權限及使用說明或概要。

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
