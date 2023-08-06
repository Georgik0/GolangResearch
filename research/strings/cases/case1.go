package cases

import (
	"fmt"
	"strings"
)

func Case1() {
	s := "footer :text_footer"
	if strings.HasPrefix(s, "footer:") {
		//fmt.Println(strings.TrimPrefix(s, "footer:"))
		fmt.Println(s)
	}
}
