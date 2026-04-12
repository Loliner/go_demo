package main

import (
	"fmt"

	// 引用本模块内的子包：模块名/相对路径
	// 模块名在 go.mod 里定义：module go_demo
	"go_demo/12_modules/mathutil"
	// 引用第三方包（需要先 go get）
	// "github.com/some/package"
)

func main() {
	// ============================================================
	// 1. 使用子包
	// ============================================================
	fmt.Println(mathutil.Add(3, 4))      // 7
	fmt.Println(mathutil.Multiply(3, 4)) // 12
	// mathutil.helper()  // ❌ 编译错误：unexported

	// ============================================================
	// 2. import 别名
	// ============================================================
	// 当包名冲突时可以起别名：
	// import mu "go_demo/12_modules/mathutil"
	// mu.Add(1, 2)

	// 用 _ 导入只执行 init()，不使用包内符号（常见于数据库驱动注册）：
	// import _ "github.com/lib/pq"

	// ============================================================
	// 3. go.mod 核心字段（见项目根目录的 go.mod）
	// ============================================================
	// module go_demo      → 模块名，import 路径的前缀
	// go 1.21             → 最低 Go 版本
	// require (           → 依赖列表，类比 package.json 的 dependencies
	//   github.com/x/y v1.2.3
	// )

	// ============================================================
	// 4. 常用命令（类比 npm）
	// ============================================================
	// go mod init <name>   → npm init
	// go get <pkg>         → npm install <pkg>
	// go get <pkg>@v1.2.3  → npm install <pkg>@1.2.3
	// go mod tidy          → 清理未使用依赖，类比 npm prune
	// go mod vendor        → 把依赖复制到 vendor/，类比 node_modules 提交到 git

	// ============================================================
	// 5. 项目结构惯例
	// ============================================================
	// myapp/
	// ├── go.mod
	// ├── main.go              ← 单入口程序直接放根目录，package main
	// ├── internal/            ← 只有本模块能 import，Go 编译器强制限制
	// │   └── db/              //    类比 JS 里约定的 _private 目录，但这里是硬限制
	// ├── pkg/                 ← 可被外部模块引用的公共包
	// │   └── mathutil/
	// └── cmd/                 ← 有多个可执行文件时用
	//     ├── server/          //    每个子目录是一个 package main
	//     │   └── main.go      //    go build ./cmd/server
	//     └── worker/
	//         └── main.go      //    go build ./cmd/worker
	//
	// internal/ 是 Go 的特殊规则：
	//   go_demo/internal/db 只能被 go_demo/ 下的代码 import
	//   外部模块 import 会直接编译报错
	//   → 适合放不想暴露给外部的实现细节

	fmt.Println("modules demo done")
}
