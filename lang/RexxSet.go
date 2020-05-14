package lang

const SCIENTIFIC int8 = 0
const ENGINEERING int8 = 1
const PLAIN int8 = 2
const DEFAULT_DIGITS int32 = 9

const DEFAULT_FORM = SCIENTIFIC

type RexxSet struct {
	Digits int32
	Form   int8
}

func RxSet() (rcvr *RexxSet) {
	rcvr = &RexxSet{}
	rcvr.Digits = DEFAULT_DIGITS
	rcvr.Form = DEFAULT_FORM
	return
}

func RxSetWithDigit(newdigits int32) (rcvr *RexxSet) {
	rcvr = RxSet()
	rcvr.Digits = newdigits
	return
}

func RxSetWithDigitandForm(newdigits int32, newform int8) (rcvr *RexxSet) {
	rcvr = RxSet()
	rcvr.Digits = newdigits
	rcvr.Form = newform
	return
}

func RxSetFromOld(old *RexxSet) (rcvr *RexxSet) {
	rcvr = RxSet()
	rcvr.Digits = old.Digits
	rcvr.Form = old.Form
	return
}

func (rcvr *RexxSet) Formword() *Rexx {
	if rcvr.Form == SCIENTIFIC {
		return ToRxFromRunes([]rune("scientific"))
	}
	return ToRxFromRunes([]rune("engineering"))
}

func (rcvr *RexxSet) SetDigits(d *Rexx) error {
	r, err := d.OpPlus(rcvr)
	if err != nil {
		return err
	}
	if r.ind == ispos {
		rx01, err := r.DataType(RxFromRune('w'))
		if err != nil {
			return err
		}
		test := rx01.OpEqS(nil, RxFromRune('1'))
		if test {
			if int32(len(r.mant))+r.exp <= 9 {
				num, err := r.ToInt32()
				if err != nil {
					return err
				}
				rcvr.Digits = num
				return nil
			}
		}
	}
	return RxException(3, d.ToString())
}

func (rcvr *RexxSet) SetForm(f *Rexx) error {
	eng, err := f.OpEq(nil, ToRxFromRunes([]rune("engineering")))
	if err != nil {
		return err
	}
	sci, err := f.OpEq(nil, ToRxFromRunes([]rune("scientific")))
	if err != nil {
		return err
	}
	if eng {
		rcvr.Form = ENGINEERING
	} else if sci {
		rcvr.Form = SCIENTIFIC
	} else {
		return RxException(3, f.ToString())
	}
	return nil
}
