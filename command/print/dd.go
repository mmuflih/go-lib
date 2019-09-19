package print

import (
	"encoding/json"
	"fmt"
)

func DD(data ...interface{}) {
	o, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(o))
}
