package lang

import (
	"math"
	"strconv"
)

var zero = RxFromInt32(0)
var one = RxFromInt32(1)
var sixteen = RxFromInt32(16)

type RexxUtil struct {
	*Rexx
}

func RxUtil() (rcvr *RexxUtil) {
	rcvr = &RexxUtil{}
	return
}

func Translate(s []rune, out []rune, in []rune, pad rune) []rune {
	var use rune
	n := 0
	if len(in) == 0 {
		return s
	}
	maxc := -1
	_1 := len(in) - 1
	i := 0
	for ; i <= _1; i++ {
		if int(in[i]) > maxc {
			maxc = int(in[i])
		}
	}
	tran := make([]rune, maxc+1)
	_2 := maxc
	i = 0
	for ; i <= _2; i++ {
		tran[i] = rune(i)
	}
	i = len(in) - 1
	for ; i >= 0; i-- {
		if i < len(out) {
			use = out[i]
		} else {
			use = pad
		}
		tran[int(in[i])] = use
	}
	res := make([]rune, len(s))
	_3 := len(s) - 1
	i = 0
	for ; i <= _3; i++ {
		n = int(s[i])
		if n > maxc {
			res[i] = rune(n)
		} else {
			res[i] = tran[n]
		}
	}
	return res
}

func D2x(d *Rexx, n int32) ([]rune, error) {
	var set *RexxSet
	var work *Rexx
	neg := false
	fill := rune(0)
	var need int32 = 0
	var j int32 = 0
	var newlen int32 = 0
	if d.ind == NotaNum {
		return nil, RxException(9, d.ToString())
	}
	if int32(len(d.mant)) > DefaultDigits {
		set = RxSetWithDigit(int32(len(d.mant)))
	} else {
		set = nil
	}
	if d.ind == isneg {
		if n < 0 {
			return nil, RxException(1, d.ToString())
		}
		rx01, err := d.OpMinus(set)
		if err != nil {
			return nil, err
		}
		rx02, err := rx01.OpSub(set, one)
		if err != nil {
			return nil, err
		}
		work = rx02
		neg = true
		fill = 'F'
	} else {
		work = d
		neg = false
		fill = '0'
	}
	if n < 0 {
		need = int32(len(work.mant))
	} else {
		need = n
	}
	res := make([]rune, need)
	i := int32(len(res)) - 1
i:
	for ; i >= 0; i-- {
		rx03, err := work.OpRem(set, sixteen)
		if err != nil {
			return nil, err
		}
		rem, err := rx03.ToInt32()
		if err != nil {
			return nil, err
		}
		rx04, err := work.OpDivI(set, sixteen)
		if err != nil {
			return nil, err
		}
		work = rx04
		if neg {
			res[i] = Hexes[15-rem]
		} else {
			res[i] = Hexes[rem]
		}
		if work.ind == iszero {
			if n >= 0 {
				j = i - 1
				for ; j >= 0; j-- {
					res[j] = fill
				}
				break i
			}
			if i == 0 {
				return res, nil
			}
			newlen = int32(len(res)) - i
			newres := make([]rune, newlen)
			copy(newres, res[i:])
			return newres, nil
		}
	}
	return res, nil
}

func X2b(x *Rexx) ([]rune, error) {
	var j int32 = 0
	var k int32 = 0
	res := make([]rune, len(x.chars)*4)
	_4 := int32(len(x.chars))
	var i int32 = 0
	for ; _4 > 0; _4-- {
		j = hexint(x.chars[i])
		if j < 0 {
			return nil, RxException(1, ToRxFromRunes([]rune("Bad hexadecimal")).OpCcblank(nil, x).ToString())
		}
		k = i * 4
		if j > 7 {
			res[k] = '1'
		} else {
			res[k] = '0'
		}
		if j&4 != 0 {
			res[k+1] = '1'
		} else {
			res[k+1] = '0'
		}
		if j&2 != 0 {
			res[k+2] = '1'
		} else {
			res[k+2] = '0'
		}
		if j&1 != 0 {
			res[k+3] = '1'
		} else {
			res[k+3] = '0'
		}
		i++
	}
	return res, nil
}

func X2c(x *Rexx) (rune, error) {
	var j int32 = 0
	_5 := int32(len(x.chars)) - 2
	var i int32 = 0
	for ; i <= _5; i++ {
		if x.chars[i] != '0' {
			break
		}
	}
	var acc int32 = 0
	_6 := int32(len(x.chars)) - 1
	for ; i <= _6; i++ {
		j = hexint(x.chars[i])
		if j < 0 {
			return 0, RxException(1, ToRxFromRunes([]rune("Bad hexadecimal")).OpCcblank(nil, x).ToString())
		}
		acc = acc*16 + j
		if acc > 65535 {
			return 0, RxException(1, ToRxFromRunes([]rune("Too big")).OpCcblank(nil, x).ToString())
		}
	}
	return rune(acc), nil
}

func X2d(x *Rexx, n int32) ([]rune, error) {
	var start int32 = 0
	var set *RexxSet
	var nibble int32 = 0
	neg := false
	if (n == 0) || (int32(len(x.chars)) == 0) {
		return []rune("0"), nil
	} else if n < 0 {
		start = 0
	} else if n > int32(len(x.chars)) {
		start = 0
	} else {
		start = int32(len(x.chars)) - n
		if (x.chars[start] > '7') || (x.chars[start] < '0') {
			neg = true
		}
	}
	digs := ((int32(len(x.chars)) * 5) / 4) + 1
	if digs > DefaultDigits {
		set = RxSetWithDigit(digs)
	} else {
		set = nil
	}
	work := RxFromInt32(0)
	_7 := int32(len(x.chars)) - 1
	i := start
	for ; i <= _7; i++ {
		nibble = hexint(x.chars[i])
		if nibble < 0 {
			return nil, RxException(1, ToRxFromRunes([]rune("Bad hexadecimal")).OpCcblank(nil, x).ToString())
		}
		if neg {
			nibble = 15 - nibble
		}
		rx01, err := work.OpMult(set, sixteen)
		if err != nil {
			return nil, err
		}
		rx02, err := rx01.OpAdd(set, RxFromInt32(nibble))
		if err != nil {
			return nil, err
		}
		work = rx02
	}
	if neg {
		rx03, err := work.OpAdd(set, one)
		if err != nil {
			return nil, err
		}
		rx04, err := rx03.OpMinus(set)
		if err != nil {
			return nil, err
		}
		work = rx04
	}
	return work.ToRunes(), nil
}

func Float64Pow(x float64, n int32) (float64, error) {
	neg := false
	if n == 0 {
		return float64(int8(1)), nil
	}
	if n > 0 {
		neg = false
	} else {
		n = int32(-n)
		neg = true
	}
	var lastbit int32 = 31
	acc := float64(1)
	var i int32 = 1
	for ; ; i++ {
		n = n << 1
		if n < 0 {
			acc = acc * x
		}
		if i == lastbit {
			break
		}
		acc = acc * acc
	}
	if math.IsInf(acc, int(n)) {
		return 0, RxException(9, "Overflow")
	}
	if neg {
		return float64(1) / acc, nil
	}
	return acc, nil
}

func Float32Pow(x float64, n int32) (float32, error) {
	value, err := Float64Pow(x, n)
	if err != nil {
		return 0, err
	}
	if math.IsInf(value, int(n)) {
		return 0, RxException(9, "Overflow")
	}
	return float32(value), nil
}

func ToRxFromFloat64(num float64, digits int32) (*Rexx, error) {
	neg := false
	m := int64(math.Float64bits(num))
	if m >= 0 {
		neg = false
	} else {
		neg = true
		m = m & 9223372036854775807
	}
	if m == 0 {
		return RxFromInt32(0), nil
	}
	fmraw := m % 4503599627370496
	fmant := fmraw | 4503599627370496
	fexp := m / 4503599627370496
	if fexp == 2047 {
		if fmraw == 0 {
			return ToRxFromRunes([]rune("Infinity")), nil
		}
		return ToRxFromRunes([]rune("NaN")), nil
	}
	workset := RxSetWithDigit(digits + 2)
	ftmp := strconv.FormatInt(fmant, 10)
	gtmp := "4503599627370496"
	res, err := RxFromString(ftmp).OpDiv(workset, RxFromString(gtmp))
	if err != nil {
		return nil, err
	}
	if fexp != 0 {
		tmp := strconv.FormatInt(fexp-1023, 10)
		twoexp, err := RxFromInt8(int8(2)).OpPow(workset, RxFromString(tmp))
		if err != nil {
			return nil, err
		}
		value, err := res.OpMult(workset, twoexp)
		if err != nil {
			return nil, err
		}
		res = value
	}
	workset.Digits = digits
	if neg {
		if res.ind == ispos {
			res.ind = isneg
		}
	}
	value, err := res.OpDiv(workset, one)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func Trunc(num *Rexx, after int32) ([]rune, error) {
	var set *RexxSet
	if int32(len(num.mant)) > DefaultDigits {
		set = RxSetWithDigit(int32(len(num.mant)))
	} else {
		set = nil
	}
	num, err := num.OpPlus(set)
	if err != nil {
		return nil, err
	}
	need := int32(len(num.mant)) + after
	if num.exp > 0 {
		need = need + num.exp
	}
	if need > DefaultDigits {
		set = RxSetWithDigit(need)
	} else {
		set = nil
	}
	num.exp = num.exp + after
	rx01, err := num.OpDivI(set, one)
	if err != nil {
		return nil, err
	}
	num = rx01
	if num.ind == iszero {
		num.exp = after
	}
	if num.exp > 0 {
		newmant := make([]rune, int32(len(num.mant))+num.exp)
		copy(newmant, num.mant)
		_8 := len(newmant) - 1
		i := len(num.mant)
		for ; i <= _8; i++ {
			newmant[i] = '0'
		}
		num.mant = newmant
		num.exp = 0
	}
	num.exp = num.exp - after
	return num.layoutnum(), nil
}

func Format(num *Rexx, before int32, after int32, explaces int32, exdigits int32, exform rune) ([]rune, error) {
	var set *RexxSet
	var thisafter int32 = 0
	if int32(len(num.mant)) > DefaultDigits {
		set = RxSetWithDigit(int32(len(num.mant)))
	} else {
		set = nil
	}
	rx01, err := num.OpPlus(set)
	if err != nil {
		return nil, err
	}
	num = rx01
	for {
		if exdigits == -1 {
			exform = 'P'
		} else if num.ind == iszero {
			exform = 'P'
		} else {
			mag := num.exp + int32(len(num.mant))
			if mag > exdigits {
			} else if mag < -5 {
			} else {
				exform = 'P'
			}
		}
		if !(false) {
			break
		}
	}
	if after >= 0 {
	setafter:
		for {
			if exform == 'P' {
				thisafter = int32(-num.exp)
			} else if exform == 'S' {
				thisafter = int32(len(num.mant)) - 1
			} else {
				lead := (num.exp + int32(len(num.mant)) - 1) % 3
				if lead < 0 {
					lead = 3 + lead
				}
				lead++
				if lead >= int32(len(num.mant)) {
					thisafter = 0
				} else {
					thisafter = int32(len(num.mant)) - lead
				}
			}
			if thisafter == after {
				break setafter
			}
			if thisafter < after {
				newmant := make([]rune, (int32(len(num.mant))+after)-thisafter)
				copy(newmant, num.mant)
				_9 := int32(len(newmant)) - 1
				i := int32(len(num.mant))
				for ; i <= _9; i++ {
					newmant[i] = '0'
				}
				num.mant = newmant
				num.exp = num.exp - (after - thisafter)
				if num.exp < MinExp {
					return nil, RxException(5, "")
				}
				break setafter
			}
			need := int32(len(num.mant)) - (thisafter - after)
			if need < 0 {
				num.mant = zero.mant
				num.ind = zero.ind
				num.exp = zero.exp
				continue setafter
			}
			bump := false
			if num.mant[need] >= '5' {
				bump = true
			}
			if need == 0 {
				if !bump {
					num.mant = zero.mant
					num.ind = zero.ind
					num.exp = zero.exp
					continue setafter
				}
			}
			newmant := make([]rune, need)
			copy(newmant, num.mant)
			num.mant = newmant
			num.exp = num.exp + (thisafter - after)
			if num.exp > MaxExp {
				return nil, RxException(5, "")
			}
			if !bump {
				break setafter
			}
			var incr = RxFromInt32(1)
			incr.ind = num.ind
			incr.exp = num.exp
			rx02, err := num.OpAdd(set, incr)
			if err != nil {
				return nil, err
			}
			num = rx02
		}
	}
	if exform == 'S' {
		num.form = SCIENTIFIC
	} else if exform == 'E' {
		num.form = ENGINEERING
	} else {
		num.form = PLAIN
	}
	if exdigits == -1 {
		num.dig = 0
	} else {
		num.dig = exdigits
	}
	a := num.layoutnum()
	if before > 0 {
		p := Pos(ToRunesFromRune('.'), a, 1)
		if p > 0 {
			p--
		} else if exdigits == -1 {
			p = int32(len(a))
		} else {
			p = Pos(ToRunesFromRune('E'), a, 1)
			if p > 0 {
				p--
			} else {
				p = int32(len(a))
			}
		}
		if p > before {
			return nil, RxException(1, strconv.Itoa(int(before)))
		}
		if p < before {
			rx03, err := Rx(a, true).Right(RxFromInt32(int32(len(a))+before-p), ToRxFromRunes([]rune(" ")))
			if err != nil {
				return nil, err
			}
			a = rx03.ToRunes()
		}
	}
	if explaces > 0 {
		epos := Pos(ToRunesFromRune('E'), a, 1)
		if epos == 0 {
			rx04 := Rx(a, true)
			rx05, err := rx04.Left(RxFromInt32(int32(len(a))+explaces+2), ToRxFromRunes([]rune(" ")))
			if err != nil {
				return nil, err
			}
			a = rx05.ToRunes()
		} else {
			places := int32(len(a)) - epos - 1
			if places > explaces {
				return nil, RxException(1, strconv.Itoa(int(explaces)))
			}
			if places < explaces {
				rx06 := Rx(a, true)
				rx07, err := rx06.Insert(ToRxFromRunes([]rune("")), RxFromInt32(epos+1), RxFromInt32(explaces-places), RxFromRune('0'))
				if err != nil {
					return nil, err
				}
				a = rx07.ToRunes()
			}
		}
	}
	return a, nil
}

func hexint(c rune) int32 {
	if c >= '0' {
		if c <= '9' {
			return int32(c) - int32('0')
		}
	}
	if c >= 'A' {
		if c <= 'F' {
			return int32(c) - int32('A') + 10
		}
	}
	if c >= 'a' {
		if c <= 'f' {
			return int32(c) - int32('a') + 10
		}
	}
	return -1
}
