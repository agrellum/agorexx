package lang

type RexxWords struct {
}

func RxWords() (rcvr *RexxWords) {
	rcvr = &RexxWords{}
	return
}

func Abbrev(a []rune, b []rune, leng int) bool {
	if (len(a) < leng) || (len(b) < leng) {
		return false
	}
	if len(b) > len(a) {
		return false
	}
	if (len(b) == 0) && (leng == 0) {
		return true
	}
	_1 := len(b) - 1
	i := 0
	for ; i <= _1; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Centre(s []rune, wid int, pad rune) []rune {
	chop := len(s) - wid
	if chop == 0 {
		return s
	}
	ret := make([]rune, wid)
	if chop > 0 {
		copy(ret, s[chop/2:])
	} else {
		gap := int(-chop) / 2
		_2 := gap - 1
		i := 0
		for ; i <= _2; i++ {
			ret[i] = pad
		}
		copy(ret[gap:], s)
		_3 := len(ret) - 1
		i = gap + len(s)
		for ; i <= _3; i++ {
			ret[i] = pad
		}
	}
	return ret
}

func Changestr(needle []rune, haystack []rune, _new []rune) []rune {
	p := 0
	i := 0
	reps := CountStr(needle, haystack)
	newlen := len(haystack) + reps*(len(_new)-len(needle))
	res := make([]rune, newlen)
	oin := 0
	out := 0
	_4 := reps
	for ; _4 > 0; _4-- {
		p = Pos(needle, haystack, oin+1)
		_5 := p - 2
		i = oin
		for ; i <= _5; i++ {
			res[out] = haystack[i]
			out++
		}
		_6 := len(_new) - 1
		i = 0
		for ; i <= _6; i++ {
			res[out] = _new[i]
			out++
		}
		oin = p + len(needle) - 1
	}
	_7 := len(haystack) - 1
	i = oin
	for ; i <= _7; i++ {
		res[out] = haystack[i]
		out++
	}
	return res
}

func Compare(a []rune, b []rune, pad rune) int {
	maxlen := len(a)
	if len(b) > maxlen {
		maxlen = len(b)
	}
	_8 := maxlen
	i := 1
	for ; i <= _8; i++ {
		if i > len(a) {
			if b[i-1] != pad {
				return i
			}
		} else if i > len(b) {
			if a[i-1] != pad {
				return i
			}
		} else {
			if a[i-1] != b[i-1] {
				return i
			}
		}
	}
	return 0
}

func CountStr(needle []rune, haystack []rune) int {
	count := 0
	p := Pos(needle, haystack, 1)
	for {
		if !(p > 0) {
			break
		}
		count++
		p = Pos(needle, haystack, p+len(needle))
	}
	return count
}

func DelStr(s []rune, start int, leng int) []rune {
	if leng == 0 {
		return s
	}
	if start > len(s) {
		return s
	}
	fin := start + leng
	if fin > len(s) {
		fin = len(s) + 1
	}
	need := start + len(s) - fin
	res := make([]rune, need)
	if start > 1 {
		copy(res, s)
	}
	if fin <= len(s) {
		copy(res[start-1:], s[fin-1:])
	}
	return res
}

func DelWord(s []rune, num int, leng int) []rune {
	if leng == 0 {
		return s
	}
	start := WordIndex(s, num)
	if start == 0 {
		return s
	}
	fin := WordIndex(s, num+leng)
	if fin == 0 {
		fin = len(s) + 1
	}
	need := start + len(s) - fin
	res := make([]rune, need)
	if start > 1 {
		copy(res, s)
	}
	if fin <= len(s) {
		copy(res[start-1:], s[fin-1:])
	}
	return res
}

func Insert(chars []rune, newchars []rune, num int, leng int, padchar rune) []rune {
	i := 0
	reslen := num + leng
	if num < len(chars) {
		reslen = reslen + len(chars) - num
	}
	res := make([]rune, reslen)
	if num > 0 {
		if num <= len(chars) {
			copy(res, chars[0:num])
		} else {
			copy(res, chars)
			_9 := num - 1
			i = len(chars)
			for ; i <= _9; i++ {
				res[i] = padchar
			}
		}
	}
	if leng > 0 {
		if leng <= len(newchars) {
			copy(res[num:], newchars[0:leng])
		} else {
			copy(res[num:], newchars)
			_10 := leng - 1
			i = len(newchars)
			for ; i <= _10; i++ {
				res[num+i] = padchar
			}
		}
	}
	if num < len(chars) {
		copy(res[num+leng:], chars[num:])
	}
	return res
}

func Overlay(chars []rune, newchars []rune, num int, leng int, padchar rune) []rune {
	i := 0
	reslen := num + leng - 1
	if reslen < len(chars) {
		reslen = len(chars)
	}
	res := make([]rune, reslen)
	if num > 1 {
		if num-1 <= len(chars) {
			copy(res, chars[0:num-1]) //FIXME
		} else {
			copy(res, chars)
			_11 := num - 2
			i = len(chars)
			for ; i <= _11; i++ {
				res[i] = padchar
			}
		}
	}
	if leng > 0 {
		if leng <= len(newchars) {
			copy(res[num-1:], newchars[0:leng]) //FIXME
		} else {
			copy(res[num-1:], newchars)
			_12 := leng - 1
			i = len(newchars)
			for ; i <= _12; i++ {
				res[num-1+i] = padchar
			}
		}
	}
	if num+leng-1 < len(chars) {
		copy(res[(num+leng)-1:], chars[(num+leng)-1:]) //FIXME
	}
	return res
}

func Nextblank(s []rune, start int) int {
	_13 := len(s)
	i := start
	for ; i <= _13; i++ {
		if s[i-1] == ' ' {
			return i
		}
	}
	return 0
}

func Nextnonblank(s []rune, start int) int {
	_14 := len(s)
	i := start
	for ; i <= _14; i++ {
		if s[i-1] != ' ' {
			return i
		}
	}
	return 0
}

func Pos(needle []rune, haystack []rune, start int) int {
	if len(needle) == 0 {
		return 0
	}
	_16 := len(haystack) - len(needle) + 1
	i := start
i:
	for ; i <= _16; i++ {
		_17 := len(needle) - 1
		j := 0
		for ; j <= _17; j++ {
			if needle[j] != haystack[i+j-1] {
				continue i
			}
		}
		return i
	}
	return 0
}

func Space(s []rune, gap int, pad rune) []rune {
	start := 1
	count := 0
	nonspaces := 0
	for {
		start = Nextnonblank(s, start)
		if start == 0 {
			break
		}
		count++
		nextblank := Nextblank(s, start+1)
		if nextblank == 0 {
			nonspaces = nonspaces + len(s) + 1 - start
			break
		}
		nonspaces = nonspaces + nextblank - start
		start = nextblank
	}
	if count == 0 {
		return make([]rune, 0)
	}
	leng := (count-1)*gap + nonspaces
	res := make([]rune, leng)
	start = 1
	out := 0
	_19 := count
	c := 1
c:
	for ; c <= _19; c++ {
		start = Nextnonblank(s, start)
		if start == 0 {
			break c
		}
	_20:
		for {
			res[out] = s[start-1]
			start++
			if start > len(s) {
				break c
			}
			out++
			if s[start-1] == ' ' {
				break _20
			}
		}
		if c == count {
			break c
		}
		_21 := gap
		for ; _21 > 0; _21-- {
			res[out] = pad
			out++
		}
	}
	return res
}

func SubWord(s []rune, num int, leng int) []rune {
	if leng == 0 {
		return make([]rune, 0)
	}
	start := WordIndex(s, num)
	if start == 0 {
		return make([]rune, 0)
	}
	fin := WordIndex(s, num+leng)
	if fin == 0 {
		fin = len(s) + 1
	}
	_22 := start
	fin = fin - 1
	for ; fin >= _22; fin-- {
		if s[fin-1] != ' ' {
			break
		}
	}
	need := fin - start + 1
	res := make([]rune, need)
	copy(res, s[start-1:])
	return res
}

func Verify(s []rune, v []rune, opt rune, start int) int {
	if opt == 'N' {
		return Verifyn(s, v, start)
	}
	return Verifym(s, v, start)
}

func Verifym(s []rune, v []rune, start int) int {
	_try := rune(0)
	t := 0
	last := len(s)
	if start > last {
		return 0
	}
	if len(v) == 0 {
		return 0
	}
	_23 := last
	i := start
	for ; i <= _23; i++ {
		_try = s[i-1]
		if _try == v[0] {
			return i
		}
		t = Pos(ToRunesFromRune(_try), v, 2)
		if t > 0 {
			return i
		}
	}
	return 0
}

func Verifyn(s []rune, v []rune, start int) int {
	t := 0
	last := len(s)
	if start > last {
		return 0
	}
	if len(v) == 0 {
		return start
	}
	_24 := last
	i := start
	for ; i <= _24; i++ {
		t = Pos(ToRunesFromRune(s[i-1]), v, 1)
		if t == 0 {
			return i
		}
	}
	return 0
}

func Word(s []rune, num int) []rune {
	return SubWord(s, num, 1)
}

func WordIndex(s []rune, num int) int {
	start := 1
	for count := 1; ; count++ {
		start = Nextnonblank(s, start)
		if start == 0 {
			return 0
		}
		if count == num {
			break
		}
		start = Nextblank(s, start+1)
		if start == 0 {
			return 0
		}
	}
	return start
}

func WordLength(s []rune, num int) int {
	start := WordIndex(s, num)
	if start == 0 {
		return 0
	}
	fin := Nextblank(s, start+1)
	if fin == 0 {
		fin = len(s) + 1
	}
	return fin - start
}

func WordPos(needle []rune, haystack []rune, wpos int) int {
	nlen := len(needle)
	if nlen == 0 {
		return 0
	}
	nbeg := Nextnonblank(needle, 1)
	if nbeg == 0 {
		return 0
	}
	hbeg := WordIndex(haystack, wpos)
	if hbeg == 0 {
		return 0
	}
	hlen := len(haystack)
_26:
	for {
		nb := nbeg
		hb := hbeg
	compare:
		for {
			nend := Nextblank(needle, nb+1)
			if nend == 0 {
				nend = nlen + 1
			}
			hend := Nextblank(haystack, hb+1)
			if hend == 0 {
				hend = hlen + 1
			}
			if hend-hb != nend-nb {
				break compare
			}
			h := hb - 1
			_25 := nend - 2
			n := nb - 1
			for ; n <= _25; n++ {
				if needle[n] != haystack[h] {
					break compare
				}
				h++
			}
			if nend > nlen {
				return wpos
			}
			nb = Nextnonblank(needle, nend)
			if nb == 0 {
				return wpos
			}
			if hend > hlen {
				break compare
			}
			hb = Nextnonblank(haystack, hend)
			if hb == 0 {
				break compare
			}
		}
		wpos++
		hbeg = Nextblank(haystack, hbeg+1)
		if hbeg == 0 {
			break _26
		}
		hbeg = Nextnonblank(haystack, hbeg+1)
		if hbeg == 0 {
			break _26
		}
	}
	return 0
}

func Words(s []rune) int {
	start := 1
	count := 0
	for {
		start = Nextnonblank(s, start)
		if start == 0 {
			break
		}
		count++
		start = Nextblank(s, start+1)
		if start == 0 {
			break
		}
	}
	return count
}
