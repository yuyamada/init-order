package b

import (
	"fmt"

	_ "github.com/yuyamada/init-order/c"
)

func init() {
	fmt.Println("b.init")
}
