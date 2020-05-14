package main

import "fmt"
import "agorexx/lang"

func main() {

	data, err := lang.RxFromString("22").OpDiv(lang.RxSetWithDigit(1000), lang.RxFromString("7"))
	if err != nil {
		fmt.Println("BUG")
	} else {
		lang.SayRexx(data)
	}

	return
}
