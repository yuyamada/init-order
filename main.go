package main

import (
	"fmt"

	_ "github.com/yuyamada/init-order/a"
	_ "github.com/yuyamada/init-order/b"
	_ "github.com/yuyamada/init-order/c"
	_ "github.com/yuyamada/init-order/d"
)

func init() {
	fmt.Println("main.init")
}

func main() {
	fmt.Println("main.main")
}
