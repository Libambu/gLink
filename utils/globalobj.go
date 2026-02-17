package utils

import (
	"encoding/json"
	"gLink/gIface"
	"os"
)

type GlobalObj struct {
	TcpServer     gIface.IServer //服务器对象
	ServerIp      string         // 服务器IP
	ServerPort    int            //服务器端口
	ServerName    string         //服务器昵称
	ServerVersion string         //服务器版本

	MaxPackageSize int //数据包的最大大小
	MaxConnNum     int //最大链接数
}

var GlobalObject *GlobalObj

// 当你的项目（比如 main 包）导入了 gLink/utils 包时。
// Go 运行时会首先初始化该包内的全局变量（这里是 GlobalObject）。
// 紧接着，就会自动执行 init() 函数。
// 所有包的 init 执行完毕后，才会执行 main 包里的 main() 函数
func init() {
	//初始化GlobalObject对象
	GlobalObject = &GlobalObj{
		ServerIp:       "0.0.0.0",
		ServerPort:     8999,
		ServerName:     "GlinkServerApp",
		ServerVersion:  "v0.4",
		MaxPackageSize: 512, // 512byte
		MaxConnNum:     3,
	}
	//然后加载用户自定义的config/glinc.json
	GlobalObject.reload()
}

func (s *GlobalObj) reload() {
	file, err := os.ReadFile("config/glink.json")
	if err != nil {
		//一旦代码执行了 panic，程序会立即停止当前的逻辑，打印错误堆栈信息，然后退出程序（Crash）。
		panic(err)
	}
	/*
		json.Unmarshal(file, s) (推荐)
		s 是指向结构体的指针。
		效果： 它是填空。它会保留 s 里原本的默认值（比如端口8999），只把 JSON 里有的字段（比如 IP）修改掉。这就是所谓的“合并配置”。

		json.Unmarshal(file, &GlobalObject) (源码写法)
		GlobalObject 是个全局变量，类型是 *GlobalObj。
		&GlobalObject 是指针的指针（指向“全局变量这个指针”的地址）。
		效果： 它是重置。Unmarshal 发现你传的是指针的指针，它往往会先创建一个全新的、空的结构体，把 JSON 读进去，然后把你的 GlobalObject 指针指向这个新结构体。
		后果： 你在 init 里辛苦设置的默认值（比如 MaxConnNum = 3）全丢了，变成了 0（因为新结构体里没初始值）。
	*/
	err = json.Unmarshal(file, s)
	if err != nil {
		panic(err)
	}
}
