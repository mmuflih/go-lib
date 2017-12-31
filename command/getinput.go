package command

import (
	"bufio"
	"os"

	"github.com/astaxie/beego"
)

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return beego.Substr(input, 0, len(input)-1)
}
