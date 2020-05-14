package lang

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RexxIO struct {
}

func RxIO() (rcvr *RexxIO) {
	rcvr = &RexxIO{}
	return
}

func Ask() *Rexx {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil
	}
	line = strings.Replace(line, "\n", "", -1)
	return RxFromString(line)
}

func AskOne() *Rexx {
	one := rune(0)
	reader := bufio.NewReader(os.Stdin)
	one, _, err := reader.ReadRune()
	if err != nil {
		return nil
	}
	return RxFromRune(one)
}

func SayString(str string) bool {
	return Say([]rune(str))
}

func SayRexx(line *Rexx) bool {
	return Say(line.ToRunes())
}

func SayRune(c rune) bool {
	ca := make([]rune, 1)
	ca[0] = c
	return Say(ca)
}

func SayInt64(n int64) bool {
	return Say([]rune(fmt.Sprintf("%v", n)))
}

func SayFloat32(f float32) bool {
	return Say(RxFromFloat32(f).ToRunes())
}

func SayFloat64(d float64) bool {
	return Say(RxFromFloat64(d).ToRunes())
}

func SayBool(b bool) bool {
	return Say(RxFromBool(b).ToRunes())
}

func Say(aline []rune) bool {
	if aline == nil {
		fmt.Println()
	} else {
		if len(aline) == 0 {
			fmt.Println()
		} else if aline[len(aline)-1] != '\000' {
			fmt.Println(string(aline))
		} else {
			bline := make([]rune, len(aline)-1)
			copy(bline, aline)
			fmt.Println(string(bline))
		}
	}
	return false
}
