package error

import (
	"fmt"
	"os"
)

func Exit(info string, code int) {
	fmt.Println("!-------------!")
	fmt.Println(info)
	fmt.Println("按任意键退出程序...")
	fmt.Scanln()
	os.Exit(code)
}
