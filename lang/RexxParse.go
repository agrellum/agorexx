package lang

const EOP = '\000'
const EOS = '\001'
const STRING = '\002'
const ABS = '\003'
const PLUS = '\004'
const MINUS = '\005'
const VAR = '\006'
const VARABS = '\007'
const VARPLUS = '\010'
const VARMINUS = '\011'
const VARLIST = '\012'
const MinCol int32 = 0
const MaxCol int32 = 999999999

type RexxParse struct {
}

func RxParse() (rcvr *RexxParse) {
	rcvr = &RexxParse{}
	return
}

func Parse(obj *Rexx, list []rune, vars []*Rexx) error {
	var needle []rune
	var slen int32 = 0
	var i int32 = 0
	var vbeg int32 = 0
	var vend int32 = 0
	var leng int32 = 0
	chars := obj.ToRunes()
	var onext int32 = 0
	var omatch int32 = 0
	var oend int32 = 0
	var ip int32 = 0
	ins := list[0]
instruction:
	for {
		obeg := onext
		for {
			switch ins {
			case STRING:
				fallthrough
			case VAR:
				if ins == VAR {
					needle = vars[int(list[ip+1])].ToRunes()
					slen = int32(len(needle))
					ip = ip + 2
				} else {
					slen = int32(list[ip+1])
					needle = make([]rune, slen)
					ip = ip + 2
					_1 := slen - 1
					i = 0
					for ; i <= _1; i++ {
						needle[i] = list[ip]
						ip++
					}
				}
				omatch = Pos(needle, chars, obeg+1) - 1
				if omatch < 0 {
					oend = int32(len(chars))
					omatch = oend
					onext = oend
				} else {
					oend = omatch
					onext = omatch + slen
				}
			case EOP:
				break instruction
			case EOS:
				oend = int32(len(chars))
				omatch = oend
				onext = oend
				ip++
			case VARABS:
				fallthrough
			case VARPLUS:
				fallthrough
			case VARMINUS:
				nextcol, err := vars[int(list[ip+1])].ToInt32()
				if err != nil {
					return err
				}
				if (nextcol < MinCol) || (nextcol > MaxCol) {
					return RxException(2, (vars[int(list[ip+1])]).ToString())
				}
				ip = ip + 2
				if ins == VARABS {
					omatch = nextcol - 1
				} else if ins == VARPLUS {
					obeg = omatch
					omatch = omatch + nextcol
				} else if ins == VARMINUS {
					obeg = omatch
					omatch = omatch - nextcol
				} else {
					return RxException(6, "")
				}
				if omatch < 0 {
					omatch = 0
				}
				oend = omatch
				onext = oend
				if oend <= obeg {
					oend = int32(len(chars))
				}
			default:
				var nextcol int32 = 0
				ip++
				_2 := int32(list[ip])
				i = 1
				for ; i <= _2; i++ {
					ip++
					nextcol = nextcol*256 + int32(list[ip])
				}
				ip++
				if ins == ABS {
					omatch = nextcol - 1
				} else if ins == PLUS {
					obeg = omatch
					omatch = omatch + nextcol
				} else if ins == MINUS {
					obeg = omatch
					omatch = omatch - nextcol
				} else {
					return RxException(6, "")
				}
				if omatch < 0 {
					omatch = 0
				}
				oend = omatch
				onext = oend
				if oend <= obeg {
					oend = int32(len(chars))
				}
			}
			if !(false) {
				break
			}
		}
		if obeg > int32(len(chars)) {
			obeg = int32(len(chars))
		}
		if oend > int32(len(chars)) {
			oend = int32(len(chars))
		}
		ins = list[ip]
		if ins != VARLIST {
			continue instruction
		}
		varcount := int32(list[ip+1])
		ip = ip + 2
		_3 := varcount
		var v int32 = 1
		for ; v <= _3; v++ {
			if v == varcount {
				vbeg = obeg
				vend = oend
				obeg = oend
			} else {
				vbeg = obeg
			vbeg:
				for ; ; vbeg++ {
					if !(vbeg < oend) {
						break
					}
					if chars[vbeg] != ' ' {
						break vbeg
					}
				}
				vend = vbeg
			vend:
				for ; ; vend++ {
					if !(vend < oend) {
						break
					}
					if chars[vend] == ' ' {
						break vend
					}
				}
				if vend < oend {
					obeg = vend + 1
				} else {
					obeg = vend
				}
			}
			leng = vend - vbeg
			value := make([]rune, leng)
			_4 := leng - 1
			i = 0
			for ; i <= _4; i++ {
				value[i] = chars[vbeg+i]
			}
			vars[int(list[ip])] = RxFromRunes(value)
			ip++
		}
		ins = list[ip]
	}
	return nil
}
