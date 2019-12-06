package osArgs

import (
	"fmt"
	"os"
)

func main() {

	//获取命令行参数
	fmt.Println("命令行的参数有", len(os.Args))

	//遍历os.Args切片，得到每一个参数
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v\n", i, v)
	}
}
