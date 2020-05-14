package main

import "agorexx/lang"

func main() {

	rexsult, err := lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem001")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem003")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem004")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem007")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem008")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("ddrem008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem009")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem010")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddrem010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem011")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddrem011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem012")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("ddrem012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem013")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("ddrem013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem014")
	} else {
		if !(rexsult.ToString() == "0.40") {
			lang.SayString("ddrem014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem015")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("ddrem015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem016")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddrem016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem017")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("ddrem017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem018")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("ddrem019")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("ddrem020")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem021")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem021")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("ddrem022")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("ddrem022")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("ddrem023")
	} else {
		if !(rexsult.ToString() == "1.000") {
			lang.SayString("ddrem023")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("ddrem024")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("ddrem025")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem025")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem030")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem030")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddrem031")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem031")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("ddrem032")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("ddrem033")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("ddrem034")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem034")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("ddrem035")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("ddrem040")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("ddrem041")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("ddrem042")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("ddrem043")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("ddrem044")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("ddrem045")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem045")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem050")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddrem051")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("ddrem052")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("ddrem053")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("ddrem054")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("ddrem055")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("ddrem060")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem060")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("ddrem061")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem061")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("ddrem062")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem062")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("ddrem063")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem063")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("ddrem064")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem064")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("ddrem065")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddrem065")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem066")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem066")
		}
	}
	rexsult, err = lang.RxFromString("999999999.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem067")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddrem067")
		}
	}
	rexsult, err = lang.RxFromString("999999999.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem068")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem068")
		}
	}
	rexsult, err = lang.RxFromString("999999999.9").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem069")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("ddrem069")
		}
	}
	rexsult, err = lang.RxFromString("999999999.999").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem070")
	} else {
		if !(rexsult.ToString() == "0.999") {
			lang.SayString("ddrem070")
		}
	}
	rexsult, err = lang.RxFromString("999999.999999").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem071")
	} else {
		if !(rexsult.ToString() == "0.999999") {
			lang.SayString("ddrem071")
		}
	}
	rexsult, err = lang.RxFromString("9").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem072")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem072")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem073")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem073")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem074")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem074")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem075")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem075")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddrem076")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem076")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem080")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem080")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem081")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem081")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem082")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem082")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem083")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem083")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+3").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem084")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem084")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+2").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem085")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem085")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem086")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem086")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem087")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem087")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem088")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem088")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem089")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem089")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-2").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem090")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-3").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem091")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem093")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-6").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem094")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem095")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem095")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem130")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem130")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem131")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem131")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem132")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem132")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem133")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem133")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem134")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem134")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem135")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem135")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem136")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem136")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem137")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem137")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem143")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem143")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("ddrem144")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem144")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("ddrem145")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem145")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("ddrem146")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem146")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem147")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("ddrem147")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("ddrem148")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("ddrem148")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("ddrem149")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("ddrem149")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem150")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem150")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem151")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem151")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem152")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem152")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddrem153")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem153")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem154")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem154")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("ddrem155")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem155")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("ddrem156")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem156")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("ddrem157")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem157")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("ddrem158")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem158")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem159")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem159")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem160")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem160")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem161")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem161")
		}
	}
	rexsult, err = lang.RxFromString("3").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem162")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem162")
		}
	}
	rexsult, err = lang.RxFromString("4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem163")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem163")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem164")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem164")
		}
	}
	rexsult, err = lang.RxFromString("6").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem165")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem165")
		}
	}
	rexsult, err = lang.RxFromString("7").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem166")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem166")
		}
	}
	rexsult, err = lang.RxFromString("8").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem167")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem167")
		}
	}
	rexsult, err = lang.RxFromString("9").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem168")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem168")
		}
	}
	rexsult, err = lang.RxFromString("10").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem169")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem169")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem171")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddrem171")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem172")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("ddrem172")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem173")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("ddrem173")
		}
	}
	rexsult, err = lang.RxFromString("0.52").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem174")
	} else {
		if !(rexsult.ToString() == "0.52") {
			lang.SayString("ddrem174")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem175")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("ddrem175")
		}
	}
	rexsult, err = lang.RxFromString("-0.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem231")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("ddrem231")
		}
	}
	rexsult, err = lang.RxFromString("-0.50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem232")
	} else {
		if !(rexsult.ToString() == "-0.50") {
			lang.SayString("ddrem232")
		}
	}
	rexsult, err = lang.RxFromString("-0.51").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem233")
	} else {
		if !(rexsult.ToString() == "-0.51") {
			lang.SayString("ddrem233")
		}
	}
	rexsult, err = lang.RxFromString("-0.52").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem234")
	} else {
		if !(rexsult.ToString() == "-0.52") {
			lang.SayString("ddrem234")
		}
	}
	rexsult, err = lang.RxFromString("-0.6").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("ddrem235")
	} else {
		if !(rexsult.ToString() == "-0.6") {
			lang.SayString("ddrem235")
		}
	}
	rexsult, err = lang.RxFromString("1E+2").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("ddrem240")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem240")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem301")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem301")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem302")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem303")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem303")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("50"))
	if err != nil {
		lang.SayString("ddrem304")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("ddrem304")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddrem305")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("ddrem305")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("ddrem306")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("ddrem306")
		}
	}
	rexsult, err = lang.RxFromString(".13").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem307")
	} else {
		if !(rexsult.ToString() == ".13") {
			lang.SayString("ddrem307")
		}
	}
	rexsult, err = lang.RxFromString("0.133").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem308")
	} else {
		if !(rexsult.ToString() == "0.133") {
			lang.SayString("ddrem308")
		}
	}
	rexsult, err = lang.RxFromString("0.1033").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem309")
	} else {
		if !(rexsult.ToString() == "0.1033") {
			lang.SayString("ddrem309")
		}
	}
	rexsult, err = lang.RxFromString("1.033").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem310")
	} else {
		if !(rexsult.ToString() == "0.033") {
			lang.SayString("ddrem310")
		}
	}
	rexsult, err = lang.RxFromString("10.33").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem311")
	} else {
		if !(rexsult.ToString() == "0.33") {
			lang.SayString("ddrem311")
		}
	}
	rexsult, err = lang.RxFromString("10.33").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem312")
	} else {
		if !(rexsult.ToString() == "0.33") {
			lang.SayString("ddrem312")
		}
	}
	rexsult, err = lang.RxFromString("103.3").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem313")
	} else {
		if !(rexsult.ToString() == "0.3") {
			lang.SayString("ddrem313")
		}
	}
	rexsult, err = lang.RxFromString("133").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem314")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem314")
		}
	}
	rexsult, err = lang.RxFromString("1033").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem315")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem315")
		}
	}
	rexsult, err = lang.RxFromString("1033").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("50"))
	if err != nil {
		lang.SayString("ddrem316")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("ddrem316")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem317")
	} else {
		if !(rexsult.ToString() == "2.0") {
			lang.SayString("ddrem317")
		}
	}
	rexsult, err = lang.RxFromString("102.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem318")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem318")
		}
	}
	rexsult, err = lang.RxFromString("103.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddrem319")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("ddrem319")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem320")
	} else {
		if !(rexsult.ToString() == "0.40") {
			lang.SayString("ddrem320")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem321")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("ddrem321")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem322")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddrem322")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem323")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddrem323")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem324")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("ddrem324")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddrem325")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddrem325")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.30"))
	if err != nil {
		lang.SayString("ddrem326")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("ddrem326")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.300"))
	if err != nil {
		lang.SayString("ddrem327")
	} else {
		if !(rexsult.ToString() == "0.100") {
			lang.SayString("ddrem327")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.3000"))
	if err != nil {
		lang.SayString("ddrem328")
	} else {
		if !(rexsult.ToString() == "0.1000") {
			lang.SayString("ddrem328")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddrem329")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddrem329")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddrem330")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("ddrem330")
		}
	}
	rexsult, err = lang.RxFromString("1.000").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddrem331")
	} else {
		if !(rexsult.ToString() == "0.100") {
			lang.SayString("ddrem331")
		}
	}
	rexsult, err = lang.RxFromString("1.0000").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddrem332")
	} else {
		if !(rexsult.ToString() == "0.1000") {
			lang.SayString("ddrem332")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem333")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem333")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("ddrem334")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem334")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("ddrem335")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem335")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("ddrem336")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem336")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem337")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("ddrem337")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("ddrem338")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("ddrem338")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("ddrem339")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("ddrem339")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.5000001"))
	if err != nil {
		lang.SayString("ddrem340")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem340")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.50000001"))
	if err != nil {
		lang.SayString("ddrem341")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem341")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("ddrem342")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem342")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.5000000001"))
	if err != nil {
		lang.SayString("ddrem343")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem343")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.50000000001"))
	if err != nil {
		lang.SayString("ddrem344")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem344")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.4999999"))
	if err != nil {
		lang.SayString("ddrem345")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("ddrem345")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.49999999"))
	if err != nil {
		lang.SayString("ddrem346")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("ddrem346")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("ddrem347")
	} else {
		if !(rexsult.ToString() == "1E-9") {
			lang.SayString("ddrem347")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.4999999999"))
	if err != nil {
		lang.SayString("ddrem348")
	} else {
		if !(rexsult.ToString() == "1E-10") {
			lang.SayString("ddrem348")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.49999999999"))
	if err != nil {
		lang.SayString("ddrem349")
	} else {
		if !(rexsult.ToString() == "1E-11") {
			lang.SayString("ddrem349")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.499999999999"))
	if err != nil {
		lang.SayString("ddrem350")
	} else {
		if !(rexsult.ToString() == "1E-12") {
			lang.SayString("ddrem350")
		}
	}
	rexsult, err = lang.RxFromString("0.03").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("ddrem351")
	} else {
		if !(rexsult.ToString() == "0.03") {
			lang.SayString("ddrem351")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem352")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem352")
		}
	}
	rexsult, err = lang.RxFromString("4.1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem353")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddrem353")
		}
	}
	rexsult, err = lang.RxFromString("4.01").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem354")
	} else {
		if !(rexsult.ToString() == "0.01") {
			lang.SayString("ddrem354")
		}
	}
	rexsult, err = lang.RxFromString("4.001").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem355")
	} else {
		if !(rexsult.ToString() == "0.001") {
			lang.SayString("ddrem355")
		}
	}
	rexsult, err = lang.RxFromString("4.0001").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem356")
	} else {
		if !(rexsult.ToString() == "0.0001") {
			lang.SayString("ddrem356")
		}
	}
	rexsult, err = lang.RxFromString("4.00001").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem357")
	} else {
		if !(rexsult.ToString() == "0.00001") {
			lang.SayString("ddrem357")
		}
	}
	rexsult, err = lang.RxFromString("4.000001").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem358")
	} else {
		if !(rexsult.ToString() == "0.000001") {
			lang.SayString("ddrem358")
		}
	}
	rexsult, err = lang.RxFromString("4.0000001").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem359")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("ddrem359")
		}
	}
	rexsult, err = lang.RxFromString("1.2").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.7345"))
	if err != nil {
		lang.SayString("ddrem360")
	} else {
		if !(rexsult.ToString() == "0.4655") {
			lang.SayString("ddrem360")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("ddrem361")
	} else {
		if !(rexsult.ToString() == "0.8") {
			lang.SayString("ddrem361")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.2"))
	if err != nil {
		lang.SayString("ddrem362")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem362")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddrem363")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("ddrem363")
		}
	}
	rexsult, err = lang.RxFromString("0.800").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("ddrem364")
	} else {
		if !(rexsult.ToString() == "0.800") {
			lang.SayString("ddrem364")
		}
	}
	rexsult, err = lang.RxFromString("0.800").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.7"))
	if err != nil {
		lang.SayString("ddrem365")
	} else {
		if !(rexsult.ToString() == "0.800") {
			lang.SayString("ddrem365")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem366")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("ddrem366")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddrem371")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("ddrem371")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem381")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem381")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("ddrem382")
	} else {
		if !(rexsult.ToString() == "0.7657") {
			lang.SayString("ddrem382")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("ddrem383")
	} else {
		if !(rexsult.ToString() == "0.668") {
			lang.SayString("ddrem383")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("ddrem384")
	} else {
		if !(rexsult.ToString() == "0.78") {
			lang.SayString("ddrem384")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("ddrem385")
	} else {
		if !(rexsult.ToString() == "0.8") {
			lang.SayString("ddrem385")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddrem386")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem386")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddrem387")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem387")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("ddrem388")
	} else {
		if !(rexsult.ToString() == "2.6912") {
			lang.SayString("ddrem388")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("ddrem389")
	} else {
		if !(rexsult.ToString() == "0.6914") {
			lang.SayString("ddrem389")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("ddrem390")
	} else {
		if !(rexsult.ToString() == "1.9") {
			lang.SayString("ddrem390")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("ddrem391")
	} else {
		if !(rexsult.ToString() == "4.73") {
			lang.SayString("ddrem391")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("ddrem392")
	} else {
		if !(rexsult.ToString() == "2.469") {
			lang.SayString("ddrem392")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("ddrem393")
	} else {
		if !(rexsult.ToString() == "0.2469") {
			lang.SayString("ddrem393")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem394")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem394")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("ddrem395")
	} else {
		if !(rexsult.ToString() == "4.7532") {
			lang.SayString("ddrem395")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("ddrem396")
	} else {
		if !(rexsult.ToString() == "2.532") {
			lang.SayString("ddrem396")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("ddrem397")
	} else {
		if !(rexsult.ToString() == "0.36") {
			lang.SayString("ddrem397")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("ddrem398")
	} else {
		if !(rexsult.ToString() == "3.0") {
			lang.SayString("ddrem398")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem401")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem401")
		}
	}
	rexsult, err = lang.RxFromString("0.55").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem402")
	} else {
		if !(rexsult.ToString() == "0.55") {
			lang.SayString("ddrem402")
		}
	}
	rexsult, err = lang.RxFromString("0.555").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem403")
	} else {
		if !(rexsult.ToString() == "0.555") {
			lang.SayString("ddrem403")
		}
	}
	rexsult, err = lang.RxFromString("0.5555").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem404")
	} else {
		if !(rexsult.ToString() == "0.5555") {
			lang.SayString("ddrem404")
		}
	}
	rexsult, err = lang.RxFromString("0.55555").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem405")
	} else {
		if !(rexsult.ToString() == "0.55555") {
			lang.SayString("ddrem405")
		}
	}
	rexsult, err = lang.RxFromString("0.555555").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem406")
	} else {
		if !(rexsult.ToString() == "0.555555") {
			lang.SayString("ddrem406")
		}
	}
	rexsult, err = lang.RxFromString("0.5555555").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem407")
	} else {
		if !(rexsult.ToString() == "0.5555555") {
			lang.SayString("ddrem407")
		}
	}
	rexsult, err = lang.RxFromString("0.55555555").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem408")
	} else {
		if !(rexsult.ToString() == "0.55555555") {
			lang.SayString("ddrem408")
		}
	}
	rexsult, err = lang.RxFromString("0.555555555").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem409")
	} else {
		if !(rexsult.ToString() == "0.555555555") {
			lang.SayString("ddrem409")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1E+383"))
	if err != nil {
		lang.SayString("ddrem422")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem422")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2E+383"))
	if err != nil {
		lang.SayString("ddrem423")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem423")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3E+383"))
	if err != nil {
		lang.SayString("ddrem424")
	} else {
		if !(rexsult.ToString() == "1E+383") {
			lang.SayString("ddrem424")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4E+383"))
	if err != nil {
		lang.SayString("ddrem425")
	} else {
		if !(rexsult.ToString() == "2E+383") {
			lang.SayString("ddrem425")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5E+383"))
	if err != nil {
		lang.SayString("ddrem426")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem426")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("6E+383"))
	if err != nil {
		lang.SayString("ddrem427")
	} else {
		if !(rexsult.ToString() == "4E+383") {
			lang.SayString("ddrem427")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("7E+383"))
	if err != nil {
		lang.SayString("ddrem428")
	} else {
		if !(rexsult.ToString() == "3E+383") {
			lang.SayString("ddrem428")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("8E+383"))
	if err != nil {
		lang.SayString("ddrem429")
	} else {
		if !(rexsult.ToString() == "2E+383") {
			lang.SayString("ddrem429")
		}
	}
	rexsult, err = lang.RxFromString("1E+384").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("9E+383"))
	if err != nil {
		lang.SayString("ddrem430")
	} else {
		if !(rexsult.ToString() == "1E+383") {
			lang.SayString("ddrem430")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1E-398"))
	if err != nil {
		lang.SayString("ddrem431")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem431")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("2E-398"))
	if err != nil {
		lang.SayString("ddrem432")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem432")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("3E-398"))
	if err != nil {
		lang.SayString("ddrem433")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("ddrem433")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4E-398"))
	if err != nil {
		lang.SayString("ddrem434")
	} else {
		if !(rexsult.ToString() == "2E-398") {
			lang.SayString("ddrem434")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5E-398"))
	if err != nil {
		lang.SayString("ddrem435")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem435")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("6E-398"))
	if err != nil {
		lang.SayString("ddrem436")
	} else {
		if !(rexsult.ToString() == "4E-398") {
			lang.SayString("ddrem436")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("7E-398"))
	if err != nil {
		lang.SayString("ddrem437")
	} else {
		if !(rexsult.ToString() == "3E-398") {
			lang.SayString("ddrem437")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("8E-398"))
	if err != nil {
		lang.SayString("ddrem438")
	} else {
		if !(rexsult.ToString() == "2E-398") {
			lang.SayString("ddrem438")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("9E-398"))
	if err != nil {
		lang.SayString("ddrem439")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("ddrem439")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10E-398"))
	if err != nil {
		lang.SayString("ddrem440")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem440")
		}
	}
	rexsult, err = lang.RxFromString("1E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("11E-398"))
	if err != nil {
		lang.SayString("ddrem441")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("ddrem441")
		}
	}
	rexsult, err = lang.RxFromString("100E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("11E-398"))
	if err != nil {
		lang.SayString("ddrem442")
	} else {
		if !(rexsult.ToString() == "1.0E-397") {
			lang.SayString("ddrem442")
		}
	}
	rexsult, err = lang.RxFromString("100E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("20E-398"))
	if err != nil {
		lang.SayString("ddrem443")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem443")
		}
	}
	rexsult, err = lang.RxFromString("100E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("21E-398"))
	if err != nil {
		lang.SayString("ddrem444")
	} else {
		if !(rexsult.ToString() == "1.3E-397") {
			lang.SayString("ddrem444")
		}
	}
	rexsult, err = lang.RxFromString("100E-397").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("30E-398"))
	if err != nil {
		lang.SayString("ddrem445")
	} else {
		if !(rexsult.ToString() == "1.0E-397") {
			lang.SayString("ddrem445")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem650")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem650")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem651")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem651")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem652")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem652")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem653")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem653")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem654")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem654")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem655")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem655")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem656")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem656")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddrem657")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem657")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem658")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem658")
		}
	}
	rexsult, err = lang.RxFromString("-0.00").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem659")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem659")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem770")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("ddrem770")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem771")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem771")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddrem801")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem801")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("ddrem802")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem802")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem803")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem803")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("ddrem804")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem804")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem805")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem805")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("ddrem806")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem806")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem807")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem807")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("ddrem808")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem808")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddrem809")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem809")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("ddrem810")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem810")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem811")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("ddrem811")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("ddrem812")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem812")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddrem821")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem821")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("ddrem822")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem822")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem823")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem823")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("ddrem824")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem824")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem825")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem825")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("ddrem826")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem826")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem827")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem827")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("ddrem828")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem828")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddrem829")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem829")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("ddrem830")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem830")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddrem831")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("ddrem831")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("ddrem832")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem832")
		}
	}
	rexsult, err = lang.RxFromString("100000000.0").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem840")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddrem840")
		}
	}
	rexsult, err = lang.RxFromString("100000000.4").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem841")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddrem841")
		}
	}
	rexsult, err = lang.RxFromString("100000000.5").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem842")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddrem842")
		}
	}
	rexsult, err = lang.RxFromString("100000000.9").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem843")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("ddrem843")
		}
	}
	rexsult, err = lang.RxFromString("100000000.999").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddrem844")
	} else {
		if !(rexsult.ToString() == "0.999") {
			lang.SayString("ddrem844")
		}
	}
	rexsult, err = lang.RxFromString("100000003").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem850")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem850")
		}
	}
	rexsult, err = lang.RxFromString("10000003").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem851")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem851")
		}
	}
	rexsult, err = lang.RxFromString("1000003").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem852")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem852")
		}
	}
	rexsult, err = lang.RxFromString("100003").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem853")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem853")
		}
	}
	rexsult, err = lang.RxFromString("10003").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem854")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem854")
		}
	}
	rexsult, err = lang.RxFromString("1003").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem855")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem855")
		}
	}
	rexsult, err = lang.RxFromString("103").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem856")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem856")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem857")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddrem857")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("ddrem858")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddrem858")
		}
	}
	rexsult, err = lang.RxFromString("1230").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1000000000000000"))
	if err != nil {
		lang.SayString("ddrem861")
	} else {
		if !(rexsult.ToString() == "1230") {
			lang.SayString("ddrem861")
		}
	}
	rexsult, err = lang.RxFromString("1230").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("ddrem878")
	} else {
		if !(rexsult.ToString() == "1230") {
			lang.SayString("ddrem878")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("ddrem1055")
	} else {
		if !(rexsult.ToString() == "1e-277") {
			lang.SayString("ddrem1055")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("ddrem1056")
	} else {
		if !(rexsult.ToString() == "1e-277") {
			lang.SayString("ddrem1056")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("ddrem1057")
	} else {
		if !(rexsult.ToString() == "-1e-277") {
			lang.SayString("ddrem1057")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("ddrem1058")
	} else {
		if !(rexsult.ToString() == "-1e-277") {
			lang.SayString("ddrem1058")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.000000000000001"))
	if err != nil {
		lang.SayString("ddrem1101")
	} else {
		if !(rexsult.ToString() == "0.765432109876546") {
			lang.SayString("ddrem1101")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.00000000000001"))
	if err != nil {
		lang.SayString("ddrem1102")
	} else {
		if !(rexsult.ToString() == "0.65432109876557") {
			lang.SayString("ddrem1102")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("1.0000000000001"))
	if err != nil {
		lang.SayString("ddrem1103")
	} else {
		if !(rexsult.ToString() == "0.5432109876668") {
			lang.SayString("ddrem1103")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123455").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.000000000000001"))
	if err != nil {
		lang.SayString("ddrem1104")
	} else {
		if !(rexsult.ToString() == "2.691358027469137") {
			lang.SayString("ddrem1104")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.000000000000001"))
	if err != nil {
		lang.SayString("ddrem1105")
	} else {
		if !(rexsult.ToString() == "3.691358027469137") {
			lang.SayString("ddrem1105")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.9999999999999"))
	if err != nil {
		lang.SayString("ddrem1106")
	} else {
		if !(rexsult.ToString() == "0.6913578024696") {
			lang.SayString("ddrem1106")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.99999999999999"))
	if err != nil {
		lang.SayString("ddrem1107")
	} else {
		if !(rexsult.ToString() == "3.46913578024691") {
			lang.SayString("ddrem1107")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("4.999999999999999"))
	if err != nil {
		lang.SayString("ddrem1108")
	} else {
		if !(rexsult.ToString() == "1.246913578024691") {
			lang.SayString("ddrem1108")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5.000000000000001"))
	if err != nil {
		lang.SayString("ddrem1109")
	} else {
		if !(rexsult.ToString() == "0.753086421975309") {
			lang.SayString("ddrem1109")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5.00000000000001"))
	if err != nil {
		lang.SayString("ddrem1110")
	} else {
		if !(rexsult.ToString() == "3.53086421975310") {
			lang.SayString("ddrem1110")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpRem(lang.RxSetWithDigit(16), lang.RxFromString("5.0000000000001"))
	if err != nil {
		lang.SayString("ddrem1111")
	} else {
		if !(rexsult.ToString() == "1.3086421975314") {
			lang.SayString("ddrem1111")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem001")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem003")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem004")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqrem007")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqrem008")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqrem008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqrem009")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem010")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqrem010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem011")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqrem011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem012")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("dqrem012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem013")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("dqrem013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem014")
	} else {
		if !(rexsult.ToString() == "0.40") {
			lang.SayString("dqrem014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem015")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("dqrem015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem016")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqrem016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem017")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("dqrem017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem018")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dqrem019")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("dqrem020")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem021")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem021")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqrem022")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("dqrem022")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dqrem023")
	} else {
		if !(rexsult.ToString() == "1.000") {
			lang.SayString("dqrem023")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dqrem024")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dqrem025")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem025")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem030")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem030")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqrem031")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem031")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqrem032")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dqrem033")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dqrem034")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem034")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dqrem035")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqrem040")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dqrem041")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dqrem042")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dqrem043")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dqrem044")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dqrem045")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem045")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem050")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqrem051")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqrem052")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dqrem053")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dqrem054")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dqrem055")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqrem060")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem060")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dqrem061")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem061")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dqrem062")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem062")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dqrem063")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem063")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dqrem064")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem064")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dqrem065")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqrem065")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem066")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem066")
		}
	}
	rexsult, err = lang.RxFromString("999999999.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem067")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqrem067")
		}
	}
	rexsult, err = lang.RxFromString("999999999.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem068")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem068")
		}
	}
	rexsult, err = lang.RxFromString("999999999.9").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem069")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("dqrem069")
		}
	}
	rexsult, err = lang.RxFromString("999999999.999").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem070")
	} else {
		if !(rexsult.ToString() == "0.999") {
			lang.SayString("dqrem070")
		}
	}
	rexsult, err = lang.RxFromString("999999.999999").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem071")
	} else {
		if !(rexsult.ToString() == "0.999999") {
			lang.SayString("dqrem071")
		}
	}
	rexsult, err = lang.RxFromString("9").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem072")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem072")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem080")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem080")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem081")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem081")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem082")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem082")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem083")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem083")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+3").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem084")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem084")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+2").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem085")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem085")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem086")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem086")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem087")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem087")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem088")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem088")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem089")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem089")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-2").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem090")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-3").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem091")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem093")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-6").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem094")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem095")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem095")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem130")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem130")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem131")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem131")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem132")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem132")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem133")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem133")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem134")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem134")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem135")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem135")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem136")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem136")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem137")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem137")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem143")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem143")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("dqrem144")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem144")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dqrem145")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem145")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("dqrem146")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem146")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem147")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("dqrem147")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dqrem148")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("dqrem148")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("dqrem149")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("dqrem149")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem150")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem150")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem151")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem151")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqrem152")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem152")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqrem153")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem153")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem154")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem154")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("dqrem155")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem155")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dqrem156")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem156")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqrem157")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem157")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dqrem158")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem158")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem159")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem159")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem160")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem160")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem161")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem161")
		}
	}
	rexsult, err = lang.RxFromString("3").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem162")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem162")
		}
	}
	rexsult, err = lang.RxFromString("4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem163")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem163")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem164")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem164")
		}
	}
	rexsult, err = lang.RxFromString("6").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem165")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem165")
		}
	}
	rexsult, err = lang.RxFromString("7").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem166")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem166")
		}
	}
	rexsult, err = lang.RxFromString("8").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem167")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem167")
		}
	}
	rexsult, err = lang.RxFromString("9").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem168")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem168")
		}
	}
	rexsult, err = lang.RxFromString("10").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem169")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem169")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem171")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqrem171")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem172")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("dqrem172")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem173")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("dqrem173")
		}
	}
	rexsult, err = lang.RxFromString("0.52").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem174")
	} else {
		if !(rexsult.ToString() == "0.52") {
			lang.SayString("dqrem174")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem175")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("dqrem175")
		}
	}
	rexsult, err = lang.RxFromString("-0.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem231")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("dqrem231")
		}
	}
	rexsult, err = lang.RxFromString("-0.50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem232")
	} else {
		if !(rexsult.ToString() == "-0.50") {
			lang.SayString("dqrem232")
		}
	}
	rexsult, err = lang.RxFromString("-0.51").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem233")
	} else {
		if !(rexsult.ToString() == "-0.51") {
			lang.SayString("dqrem233")
		}
	}
	rexsult, err = lang.RxFromString("-0.52").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem234")
	} else {
		if !(rexsult.ToString() == "-0.52") {
			lang.SayString("dqrem234")
		}
	}
	rexsult, err = lang.RxFromString("-0.6").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("dqrem235")
	} else {
		if !(rexsult.ToString() == "-0.6") {
			lang.SayString("dqrem235")
		}
	}
	rexsult, err = lang.RxFromString("1E+2").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dqrem240")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem240")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqrem301")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem301")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem302")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem303")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem303")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("50"))
	if err != nil {
		lang.SayString("dqrem304")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("dqrem304")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqrem305")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("dqrem305")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqrem306")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("dqrem306")
		}
	}
	rexsult, err = lang.RxFromString(".13").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem307")
	} else {
		if !(rexsult.ToString() == ".13") {
			lang.SayString("dqrem307")
		}
	}
	rexsult, err = lang.RxFromString("0.133").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem308")
	} else {
		if !(rexsult.ToString() == "0.133") {
			lang.SayString("dqrem308")
		}
	}
	rexsult, err = lang.RxFromString("0.1033").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem309")
	} else {
		if !(rexsult.ToString() == "0.1033") {
			lang.SayString("dqrem309")
		}
	}
	rexsult, err = lang.RxFromString("1.033").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem310")
	} else {
		if !(rexsult.ToString() == "0.033") {
			lang.SayString("dqrem310")
		}
	}
	rexsult, err = lang.RxFromString("10.33").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem311")
	} else {
		if !(rexsult.ToString() == "0.33") {
			lang.SayString("dqrem311")
		}
	}
	rexsult, err = lang.RxFromString("10.33").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem312")
	} else {
		if !(rexsult.ToString() == "0.33") {
			lang.SayString("dqrem312")
		}
	}
	rexsult, err = lang.RxFromString("103.3").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem313")
	} else {
		if !(rexsult.ToString() == "0.3") {
			lang.SayString("dqrem313")
		}
	}
	rexsult, err = lang.RxFromString("133").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem314")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem314")
		}
	}
	rexsult, err = lang.RxFromString("1033").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem315")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem315")
		}
	}
	rexsult, err = lang.RxFromString("1033").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("50"))
	if err != nil {
		lang.SayString("dqrem316")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("dqrem316")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqrem317")
	} else {
		if !(rexsult.ToString() == "2.0") {
			lang.SayString("dqrem317")
		}
	}
	rexsult, err = lang.RxFromString("102.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqrem318")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem318")
		}
	}
	rexsult, err = lang.RxFromString("103.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqrem319")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("dqrem319")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem320")
	} else {
		if !(rexsult.ToString() == "0.40") {
			lang.SayString("dqrem320")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem321")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("dqrem321")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem322")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqrem322")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem323")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqrem323")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem324")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("dqrem324")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqrem325")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqrem325")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.30"))
	if err != nil {
		lang.SayString("dqrem326")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("dqrem326")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.300"))
	if err != nil {
		lang.SayString("dqrem327")
	} else {
		if !(rexsult.ToString() == "0.100") {
			lang.SayString("dqrem327")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.3000"))
	if err != nil {
		lang.SayString("dqrem328")
	} else {
		if !(rexsult.ToString() == "0.1000") {
			lang.SayString("dqrem328")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqrem329")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqrem329")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqrem330")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("dqrem330")
		}
	}
	rexsult, err = lang.RxFromString("1.000").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqrem331")
	} else {
		if !(rexsult.ToString() == "0.100") {
			lang.SayString("dqrem331")
		}
	}
	rexsult, err = lang.RxFromString("1.0000").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqrem332")
	} else {
		if !(rexsult.ToString() == "0.1000") {
			lang.SayString("dqrem332")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem333")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem333")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("dqrem334")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem334")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dqrem335")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem335")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("dqrem336")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem336")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem337")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("dqrem337")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dqrem338")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("dqrem338")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("dqrem339")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("dqrem339")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.5000001"))
	if err != nil {
		lang.SayString("dqrem340")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem340")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.50000001"))
	if err != nil {
		lang.SayString("dqrem341")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem341")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("dqrem342")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem342")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.5000000001"))
	if err != nil {
		lang.SayString("dqrem343")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem343")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.50000000001"))
	if err != nil {
		lang.SayString("dqrem344")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem344")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.4999999"))
	if err != nil {
		lang.SayString("dqrem345")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("dqrem345")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.49999999"))
	if err != nil {
		lang.SayString("dqrem346")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("dqrem346")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("dqrem347")
	} else {
		if !(rexsult.ToString() == "1E-9") {
			lang.SayString("dqrem347")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.4999999999"))
	if err != nil {
		lang.SayString("dqrem348")
	} else {
		if !(rexsult.ToString() == "1E-10") {
			lang.SayString("dqrem348")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.49999999999"))
	if err != nil {
		lang.SayString("dqrem349")
	} else {
		if !(rexsult.ToString() == "1E-11") {
			lang.SayString("dqrem349")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.499999999999"))
	if err != nil {
		lang.SayString("dqrem350")
	} else {
		if !(rexsult.ToString() == "1E-12") {
			lang.SayString("dqrem350")
		}
	}
	rexsult, err = lang.RxFromString("0.03").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dqrem351")
	} else {
		if !(rexsult.ToString() == "0.03") {
			lang.SayString("dqrem351")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem352")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem352")
		}
	}
	rexsult, err = lang.RxFromString("4.1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem353")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqrem353")
		}
	}
	rexsult, err = lang.RxFromString("4.01").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem354")
	} else {
		if !(rexsult.ToString() == "0.01") {
			lang.SayString("dqrem354")
		}
	}
	rexsult, err = lang.RxFromString("4.001").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem355")
	} else {
		if !(rexsult.ToString() == "0.001") {
			lang.SayString("dqrem355")
		}
	}
	rexsult, err = lang.RxFromString("4.0001").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem356")
	} else {
		if !(rexsult.ToString() == "0.0001") {
			lang.SayString("dqrem356")
		}
	}
	rexsult, err = lang.RxFromString("4.00001").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem357")
	} else {
		if !(rexsult.ToString() == "0.00001") {
			lang.SayString("dqrem357")
		}
	}
	rexsult, err = lang.RxFromString("4.000001").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem358")
	} else {
		if !(rexsult.ToString() == "0.000001") {
			lang.SayString("dqrem358")
		}
	}
	rexsult, err = lang.RxFromString("4.0000001").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem359")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("dqrem359")
		}
	}
	rexsult, err = lang.RxFromString("1.2").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.7345"))
	if err != nil {
		lang.SayString("dqrem360")
	} else {
		if !(rexsult.ToString() == "0.4655") {
			lang.SayString("dqrem360")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("dqrem361")
	} else {
		if !(rexsult.ToString() == "0.8") {
			lang.SayString("dqrem361")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.2"))
	if err != nil {
		lang.SayString("dqrem362")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem362")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqrem363")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("dqrem363")
		}
	}
	rexsult, err = lang.RxFromString("0.800").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("dqrem364")
	} else {
		if !(rexsult.ToString() == "0.800") {
			lang.SayString("dqrem364")
		}
	}
	rexsult, err = lang.RxFromString("0.800").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.7"))
	if err != nil {
		lang.SayString("dqrem365")
	} else {
		if !(rexsult.ToString() == "0.800") {
			lang.SayString("dqrem365")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem366")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("dqrem366")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqrem371")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("dqrem371")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem381")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem381")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("dqrem382")
	} else {
		if !(rexsult.ToString() == "0.7657") {
			lang.SayString("dqrem382")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("dqrem383")
	} else {
		if !(rexsult.ToString() == "0.668") {
			lang.SayString("dqrem383")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("dqrem384")
	} else {
		if !(rexsult.ToString() == "0.78") {
			lang.SayString("dqrem384")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("dqrem385")
	} else {
		if !(rexsult.ToString() == "0.8") {
			lang.SayString("dqrem385")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqrem386")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem386")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqrem387")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem387")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("dqrem388")
	} else {
		if !(rexsult.ToString() == "2.6912") {
			lang.SayString("dqrem388")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("dqrem389")
	} else {
		if !(rexsult.ToString() == "0.6914") {
			lang.SayString("dqrem389")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("dqrem390")
	} else {
		if !(rexsult.ToString() == "1.9") {
			lang.SayString("dqrem390")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("dqrem391")
	} else {
		if !(rexsult.ToString() == "4.73") {
			lang.SayString("dqrem391")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("dqrem392")
	} else {
		if !(rexsult.ToString() == "2.469") {
			lang.SayString("dqrem392")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("dqrem393")
	} else {
		if !(rexsult.ToString() == "0.2469") {
			lang.SayString("dqrem393")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem394")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem394")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("dqrem395")
	} else {
		if !(rexsult.ToString() == "4.7532") {
			lang.SayString("dqrem395")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("dqrem396")
	} else {
		if !(rexsult.ToString() == "2.532") {
			lang.SayString("dqrem396")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("dqrem397")
	} else {
		if !(rexsult.ToString() == "0.36") {
			lang.SayString("dqrem397")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("dqrem398")
	} else {
		if !(rexsult.ToString() == "3.0") {
			lang.SayString("dqrem398")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem401")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem401")
		}
	}
	rexsult, err = lang.RxFromString("0.55").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem402")
	} else {
		if !(rexsult.ToString() == "0.55") {
			lang.SayString("dqrem402")
		}
	}
	rexsult, err = lang.RxFromString("0.555").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem403")
	} else {
		if !(rexsult.ToString() == "0.555") {
			lang.SayString("dqrem403")
		}
	}
	rexsult, err = lang.RxFromString("0.5555").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem404")
	} else {
		if !(rexsult.ToString() == "0.5555") {
			lang.SayString("dqrem404")
		}
	}
	rexsult, err = lang.RxFromString("0.55555").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem405")
	} else {
		if !(rexsult.ToString() == "0.55555") {
			lang.SayString("dqrem405")
		}
	}
	rexsult, err = lang.RxFromString("0.555555").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem406")
	} else {
		if !(rexsult.ToString() == "0.555555") {
			lang.SayString("dqrem406")
		}
	}
	rexsult, err = lang.RxFromString("0.5555555").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem407")
	} else {
		if !(rexsult.ToString() == "0.5555555") {
			lang.SayString("dqrem407")
		}
	}
	rexsult, err = lang.RxFromString("0.55555555").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem408")
	} else {
		if !(rexsult.ToString() == "0.55555555") {
			lang.SayString("dqrem408")
		}
	}
	rexsult, err = lang.RxFromString("0.555555555").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem409")
	} else {
		if !(rexsult.ToString() == "0.555555555") {
			lang.SayString("dqrem409")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1E+6143"))
	if err != nil {
		lang.SayString("dqrem422")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem422")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2E+6143"))
	if err != nil {
		lang.SayString("dqrem423")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem423")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3E+6143"))
	if err != nil {
		lang.SayString("dqrem424")
	} else {
		if !(rexsult.ToString() == "1E+6143") {
			lang.SayString("dqrem424")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4E+6143"))
	if err != nil {
		lang.SayString("dqrem425")
	} else {
		if !(rexsult.ToString() == "2E+6143") {
			lang.SayString("dqrem425")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5E+6143"))
	if err != nil {
		lang.SayString("dqrem426")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem426")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("6E+6143"))
	if err != nil {
		lang.SayString("dqrem427")
	} else {
		if !(rexsult.ToString() == "4E+6143") {
			lang.SayString("dqrem427")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("7E+6143"))
	if err != nil {
		lang.SayString("dqrem428")
	} else {
		if !(rexsult.ToString() == "3E+6143") {
			lang.SayString("dqrem428")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("8E+6143"))
	if err != nil {
		lang.SayString("dqrem429")
	} else {
		if !(rexsult.ToString() == "2E+6143") {
			lang.SayString("dqrem429")
		}
	}
	rexsult, err = lang.RxFromString("1E+6144").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("9E+6143"))
	if err != nil {
		lang.SayString("dqrem430")
	} else {
		if !(rexsult.ToString() == "1E+6143") {
			lang.SayString("dqrem430")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1E-6176"))
	if err != nil {
		lang.SayString("dqrem431")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem431")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("2E-6176"))
	if err != nil {
		lang.SayString("dqrem432")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem432")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("3E-6176"))
	if err != nil {
		lang.SayString("dqrem433")
	} else {
		if !(rexsult.ToString() == "1E-6176") {
			lang.SayString("dqrem433")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4E-6176"))
	if err != nil {
		lang.SayString("dqrem434")
	} else {
		if !(rexsult.ToString() == "2E-6176") {
			lang.SayString("dqrem434")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5E-6176"))
	if err != nil {
		lang.SayString("dqrem435")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem435")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("6E-6176"))
	if err != nil {
		lang.SayString("dqrem436")
	} else {
		if !(rexsult.ToString() == "4E-6176") {
			lang.SayString("dqrem436")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("7E-6176"))
	if err != nil {
		lang.SayString("dqrem437")
	} else {
		if !(rexsult.ToString() == "3E-6176") {
			lang.SayString("dqrem437")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("8E-6176"))
	if err != nil {
		lang.SayString("dqrem438")
	} else {
		if !(rexsult.ToString() == "2E-6176") {
			lang.SayString("dqrem438")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("9E-6176"))
	if err != nil {
		lang.SayString("dqrem439")
	} else {
		if !(rexsult.ToString() == "1E-6176") {
			lang.SayString("dqrem439")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10E-6176"))
	if err != nil {
		lang.SayString("dqrem440")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem440")
		}
	}
	rexsult, err = lang.RxFromString("1E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("11E-6176"))
	if err != nil {
		lang.SayString("dqrem441")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqrem441")
		}
	}
	rexsult, err = lang.RxFromString("100E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("11E-6176"))
	if err != nil {
		lang.SayString("dqrem442")
	} else {
		if !(rexsult.ToString() == "1.0E-6175") {
			lang.SayString("dqrem442")
		}
	}
	rexsult, err = lang.RxFromString("100E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("20E-6176"))
	if err != nil {
		lang.SayString("dqrem443")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem443")
		}
	}
	rexsult, err = lang.RxFromString("100E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("21E-6176"))
	if err != nil {
		lang.SayString("dqrem444")
	} else {
		if !(rexsult.ToString() == "1.3E-6175") {
			lang.SayString("dqrem444")
		}
	}
	rexsult, err = lang.RxFromString("100E-6175").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("30E-6176"))
	if err != nil {
		lang.SayString("dqrem445")
	} else {
		if !(rexsult.ToString() == "1.0E-6175") {
			lang.SayString("dqrem445")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem650")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem650")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem651")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem651")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem652")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem652")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem653")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem653")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem654")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem654")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem655")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem655")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem656")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem656")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqrem657")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem657")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem658")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem658")
		}
	}
	rexsult, err = lang.RxFromString("-0.00").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem659")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem659")
		}
	}
	rexsult, err = lang.RxFromString("1234568888888887777777777890123456").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem770")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dqrem770")
		}
	}
	rexsult, err = lang.RxFromString("1234568888888887777777777890123456").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem771")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem771")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqrem801")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem801")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dqrem802")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem802")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem803")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem803")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dqrem804")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem804")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem805")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem805")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dqrem806")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem806")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem807")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem807")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dqrem808")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem808")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqrem809")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem809")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dqrem810")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem810")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem811")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dqrem811")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dqrem812")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem812")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqrem821")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem821")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dqrem822")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem822")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem823")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem823")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dqrem824")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem824")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem825")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem825")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dqrem826")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem826")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem827")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem827")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dqrem828")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem828")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqrem829")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem829")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dqrem830")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem830")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqrem831")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dqrem831")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dqrem832")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem832")
		}
	}
	rexsult, err = lang.RxFromString("100000000.0").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem840")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqrem840")
		}
	}
	rexsult, err = lang.RxFromString("100000000.4").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem841")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqrem841")
		}
	}
	rexsult, err = lang.RxFromString("100000000.5").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem842")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqrem842")
		}
	}
	rexsult, err = lang.RxFromString("100000000.9").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem843")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("dqrem843")
		}
	}
	rexsult, err = lang.RxFromString("100000000.999").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqrem844")
	} else {
		if !(rexsult.ToString() == "0.999") {
			lang.SayString("dqrem844")
		}
	}
	rexsult, err = lang.RxFromString("100000003").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem850")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem850")
		}
	}
	rexsult, err = lang.RxFromString("10000003").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem851")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem851")
		}
	}
	rexsult, err = lang.RxFromString("1000003").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem852")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem852")
		}
	}
	rexsult, err = lang.RxFromString("100003").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem853")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem853")
		}
	}
	rexsult, err = lang.RxFromString("10003").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem854")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem854")
		}
	}
	rexsult, err = lang.RxFromString("1003").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem855")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem855")
		}
	}
	rexsult, err = lang.RxFromString("103").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem856")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem856")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem857")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqrem857")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqrem858")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqrem858")
		}
	}
	rexsult, err = lang.RxFromString("1230").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000"))
	if err != nil {
		lang.SayString("dqrem861")
	} else {
		if !(rexsult.ToString() == "1230") {
			lang.SayString("dqrem861")
		}
	}
	rexsult, err = lang.RxFromString("1230").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dqrem878")
	} else {
		if !(rexsult.ToString() == "1230") {
			lang.SayString("dqrem878")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("dqrem1055")
	} else {
		if !(rexsult.ToString() == "1e-277") {
			lang.SayString("dqrem1055")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("dqrem1056")
	} else {
		if !(rexsult.ToString() == "1e-277") {
			lang.SayString("dqrem1056")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("dqrem1057")
	} else {
		if !(rexsult.ToString() == "-1e-277") {
			lang.SayString("dqrem1057")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("dqrem1058")
	} else {
		if !(rexsult.ToString() == "-1e-277") {
			lang.SayString("dqrem1058")
		}
	}
	rexsult, err = lang.RxFromString("8.336804418094040989630006819881709E-6143").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("8.336804418094040989630006819889000E-6143"))
	if err != nil {
		lang.SayString("dqrem1070")
	} else {
		if !(rexsult.ToString() == "8.336804418094040989630006819881709E-6143") {
			lang.SayString("dqrem1070")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901234").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.000000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqrem1120")
	} else {
		if !(rexsult.ToString() == "0.765432109876543210987654321098768") {
			lang.SayString("dqrem1120")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901234").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.00000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqrem1121")
	} else {
		if !(rexsult.ToString() == "0.65432109876543210987654321098779") {
			lang.SayString("dqrem1121")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901234").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("1.0000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqrem1122")
	} else {
		if !(rexsult.ToString() == "0.5432109876543210987654321098890") {
			lang.SayString("dqrem1122")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901255").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.000000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqrem1123")
	} else {
		if !(rexsult.ToString() == "2.691358027469135802746913580274687") {
			lang.SayString("dqrem1123")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901234").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.000000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqrem1124")
	} else {
		if !(rexsult.ToString() == "1.691358027469135802746913580274692") {
			lang.SayString("dqrem1124")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901234").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.9999999999999999999999999999999"))
	if err != nil {
		lang.SayString("dqrem1125")
	} else {
		if !(rexsult.ToString() == "3.6913578024691357802469135780251") {
			lang.SayString("dqrem1125")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901234").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.99999999999999999999999999999999"))
	if err != nil {
		lang.SayString("dqrem1126")
	} else {
		if !(rexsult.ToString() == "1.46913578024691357802469135780247") {
			lang.SayString("dqrem1126")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901234").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("4.999999999999999999999999999999999"))
	if err != nil {
		lang.SayString("dqrem1127")
	} else {
		if !(rexsult.ToString() == "4.246913578024691357802469135780246") {
			lang.SayString("dqrem1127")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456789012345678901234").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("5.0000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqrem1128")
	} else {
		if !(rexsult.ToString() == "4.3086421975308642197530864219759") {
			lang.SayString("dqrem1128")
		}
	}
	rexsult, err = lang.RxFromString("8.336804418094040989630006819881709E-6143").OpRem(lang.RxSetWithDigit(34), lang.RxFromString("8.336804418094040989630006819889000E-6143"))
	if err != nil {
		lang.SayString("dqrmn1070")
	} else {
		if !(rexsult.ToString() == "8.336804418094040989630006819881709E-6143") {
			lang.SayString("dqrmn1070")
		}
	}
	rexsult, err = lang.RxFromString("4953734675913.065314738743322579").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("0218.932010396534371704930714860E+797"))
	if err != nil {
		lang.SayString("remx3001")
	} else {
		if !(rexsult.ToString() == "4953734675913.065314738743322579") {
			lang.SayString("remx3001")
		}
	}
	rexsult, err = lang.RxFromString("9641.684323386955881595490347910E-844").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-78864532047.12287484430980636798E+934"))
	if err != nil {
		lang.SayString("remx3002")
	} else {
		if !(rexsult.ToString() == "9641.684323386955881595490347910E-844") {
			lang.SayString("remx3002")
		}
	}
	rexsult, err = lang.RxFromString("479084.8561808930525417735205519").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("084157571054.2691784660983989931"))
	if err != nil {
		lang.SayString("remx3004")
	} else {
		if !(rexsult.ToString() == "479084.8561808930525417735205519") {
			lang.SayString("remx3004")
		}
	}
	rexsult, err = lang.RxFromString("-0363750788.573782205664349562931").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-3172.080934464133691909905980096"))
	if err != nil {
		lang.SayString("remx3005")
	} else {
		if !(rexsult.ToString() == "-1923.656911066945656824381431488") {
			lang.SayString("remx3005")
		}
	}
	rexsult, err = lang.RxFromString("1381026551423669919010191878449").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-82.66614775445371254999357800739"))
	if err != nil {
		lang.SayString("remx3006")
	} else {
		if !(rexsult.ToString() == "74.22115953553602036042168767377") {
			lang.SayString("remx3006")
		}
	}
	rexsult, err = lang.RxFromString("4627.026960423072127953556635585").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-4410583132901.830017479741231131"))
	if err != nil {
		lang.SayString("remx3007")
	} else {
		if !(rexsult.ToString() == "4627.026960423072127953556635585") {
			lang.SayString("remx3007")
		}
	}
	rexsult, err = lang.RxFromString("75353574493.84484153484918212042").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-8684111695095849922263044191221"))
	if err != nil {
		lang.SayString("remx3008")
	} else {
		if !(rexsult.ToString() == "75353574493.84484153484918212042") {
			lang.SayString("remx3008")
		}
	}
	rexsult, err = lang.RxFromString("6907058.216435355874729592373011").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("2.857005446917670515662398741545"))
	if err != nil {
		lang.SayString("remx3009")
	} else {
		if !(rexsult.ToString() == "1.846043452483451396449034189630") {
			lang.SayString("remx3009")
		}
	}
	rexsult, err = lang.RxFromString("-0708069.025667471996378081482549").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-562842.4701520787831018732202804"))
	if err != nil {
		lang.SayString("remx3011")
	} else {
		if !(rexsult.ToString() == "-145226.5555153932132762082622686") {
			lang.SayString("remx3011")
		}
	}
	rexsult, err = lang.RxFromString("4055087.246994644709729942673976").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-43183146921897.67383476104084575E+211"))
	if err != nil {
		lang.SayString("remx3012")
	} else {
		if !(rexsult.ToString() == "4055087.246994644709729942673976") {
			lang.SayString("remx3012")
		}
	}
	rexsult, err = lang.RxFromString("4502895892520.396581348110906909E-512").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-815.9047305921862348263521876034"))
	if err != nil {
		lang.SayString("remx3013")
	} else {
		if !(rexsult.ToString() == "4502895892520.396581348110906909E-512") {
			lang.SayString("remx3013")
		}
	}
	rexsult, err = lang.RxFromString("467.6721295072628100260239179865").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-02.07155073395573569852316073025"))
	if err != nil {
		lang.SayString("remx3014")
	} else {
		if !(rexsult.ToString() == "1.57321436722227785831275368025") {
			lang.SayString("remx3014")
		}
	}
	rexsult, err = lang.RxFromString("2.156795313311150143949997552501E-571").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-8677147.586389401682712180146855"))
	if err != nil {
		lang.SayString("remx3015")
	} else {
		if !(rexsult.ToString() == "2.156795313311150143949997552501E-571") {
			lang.SayString("remx3015")
		}
	}
	rexsult, err = lang.RxFromString("-974953.2801637208368002585822457").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-693095793.3667578067802698191246"))
	if err != nil {
		lang.SayString("remx3016")
	} else {
		if !(rexsult.ToString() == "-974953.2801637208368002585822457") {
			lang.SayString("remx3016")
		}
	}
	rexsult, err = lang.RxFromString("262273.0222851186523650889896428E-624").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("74177.21073338090843145838835480"))
	if err != nil {
		lang.SayString("remx3018")
	} else {
		if !(rexsult.ToString() == "262273.0222851186523650889896428E-624") {
			lang.SayString("remx3018")
		}
	}
	rexsult, err = lang.RxFromString("-8036052748815903177624716581732").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-066677357.4438809548850966167573"))
	if err != nil {
		lang.SayString("remx3019")
	} else {
		if !(rexsult.ToString() == "-41816499.5048993028288978900564") {
			lang.SayString("remx3019")
		}
	}
	rexsult, err = lang.RxFromString("24791301060.37938360567775506973").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-5613327866480.322649080205877564"))
	if err != nil {
		lang.SayString("remx3021")
	} else {
		if !(rexsult.ToString() == "24791301060.37938360567775506973") {
			lang.SayString("remx3021")
		}
	}
	rexsult, err = lang.RxFromString("-930711443.9474781586162910776139").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-740.3860979292775472622798348030"))
	if err != nil {
		lang.SayString("remx3022")
	} else {
		if !(rexsult.ToString() == "-214.9123046664996750639167712140") {
			lang.SayString("remx3022")
		}
	}
	rexsult, err = lang.RxFromString("2358276428765.064191082773385539").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("214.3589796082328665878602304469"))
	if err != nil {
		lang.SayString("remx3023")
	} else {
		if !(rexsult.ToString() == "15.1969844739096415643561521775") {
			lang.SayString("remx3023")
		}
	}
	rexsult, err = lang.RxFromString("140422069.5863246490180206814374E-447").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-567195652586.2454217069003186487"))
	if err != nil {
		lang.SayString("remx3025")
	} else {
		if !(rexsult.ToString() == "140422069.5863246490180206814374E-447") {
			lang.SayString("remx3025")
		}
	}
	rexsult, err = lang.RxFromString("75929096475.63450425339472559646E+153").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-0945260193.503803519572604151290E+459"))
	if err != nil {
		lang.SayString("remx3026")
	} else {
		if !(rexsult.ToString() == "75929096475.63450425339472559646E+153") {
			lang.SayString("remx3026")
		}
	}
	rexsult, err = lang.RxFromString("6312318309.142044953357460463732").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-5641317823.202274083982487558514E+628"))
	if err != nil {
		lang.SayString("remx3027")
	} else {
		if !(rexsult.ToString() == "6312318309.142044953357460463732") {
			lang.SayString("remx3027")
		}
	}
	rexsult, err = lang.RxFromString("98471198160.56524417578665886060").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-23994.14313393939743548945165462"))
	if err != nil {
		lang.SayString("remx3029")
	} else {
		if !(rexsult.ToString() == "2551.45824316125588493249246784") {
			lang.SayString("remx3029")
		}
	}
	rexsult, err = lang.RxFromString("329326552.0208398002250836592043").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-02451.10065397010591546041034041"))
	if err != nil {
		lang.SayString("remx3030")
	} else {
		if !(rexsult.ToString() == "1570.35472430963565384668749322") {
			lang.SayString("remx3030")
		}
	}
	rexsult, err = lang.RxFromString("-92980.68431371090354435763218439").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-2282178507046019721925800997065"))
	if err != nil {
		lang.SayString("remx3031")
	} else {
		if !(rexsult.ToString() == "-92980.68431371090354435763218439") {
			lang.SayString("remx3031")
		}
	}
	rexsult, err = lang.RxFromString("37.27457578793521166809739140081").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-392550.4790095035979998355569916"))
	if err != nil {
		lang.SayString("remx3033")
	} else {
		if !(rexsult.ToString() == "37.27457578793521166809739140081") {
			lang.SayString("remx3033")
		}
	}
	rexsult, err = lang.RxFromString("-2787.980590304199878755265273703").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("7117631179305319208210387565324"))
	if err != nil {
		lang.SayString("remx3034")
	} else {
		if !(rexsult.ToString() == "-2787.980590304199878755265273703") {
			lang.SayString("remx3034")
		}
	}
	rexsult, err = lang.RxFromString("3944570323.331121750661920475191").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-17360722.28878145641394962484366"))
	if err != nil {
		lang.SayString("remx3036")
	} else {
		if !(rexsult.ToString() == "3686363.77773114469535563568018") {
			lang.SayString("remx3036")
		}
	}
	rexsult, err = lang.RxFromString("19544.14018503427029002552872707").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("1786697762.885178994182133839546"))
	if err != nil {
		lang.SayString("remx3037")
	} else {
		if !(rexsult.ToString() == "19544.14018503427029002552872707") {
			lang.SayString("remx3037")
		}
	}
	rexsult, err = lang.RxFromString("-05.75485957937617757983513662981").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("5564476875.989640431173694372083"))
	if err != nil {
		lang.SayString("remx3038")
	} else {
		if !(rexsult.ToString() == "-05.75485957937617757983513662981") {
			lang.SayString("remx3038")
		}
	}
	rexsult, err = lang.RxFromString("-4208820.898718069194008526302746").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("626887.7553774705678201112845462E+206"))
	if err != nil {
		lang.SayString("remx3039")
	} else {
		if !(rexsult.ToString() == "-4208820.898718069194008526302746") {
			lang.SayString("remx3039")
		}
	}
	rexsult, err = lang.RxFromString("-442941.7541811527940918244383454").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-068126768.0563559819156379151016"))
	if err != nil {
		lang.SayString("remx3041")
	} else {
		if !(rexsult.ToString() == "-442941.7541811527940918244383454") {
			lang.SayString("remx3041")
		}
	}
	rexsult, err = lang.RxFromString("-040726778711.8677615616711676159").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("299691.9430345259174614997064916"))
	if err != nil {
		lang.SayString("remx3042")
	} else {
		if !(rexsult.ToString() == "-142113.1908620082406650022240180") {
			lang.SayString("remx3042")
		}
	}
	rexsult, err = lang.RxFromString("-1934197520.738366912179143085955").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("3.810751422515178400293693371519"))
	if err != nil {
		lang.SayString("remx3043")
	} else {
		if !(rexsult.ToString() == "-2.786637155934674312936704177047") {
			lang.SayString("remx3043")
		}
	}
	rexsult, err = lang.RxFromString("813262.7723533833038829559646830").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-303284822716.8282178131118185907"))
	if err != nil {
		lang.SayString("remx3044")
	} else {
		if !(rexsult.ToString() == "813262.7723533833038829559646830") {
			lang.SayString("remx3044")
		}
	}
	rexsult, err = lang.RxFromString("-075537177538.1814516621962185490").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("26980775255.51542856483122484898"))
	if err != nil {
		lang.SayString("remx3046")
	} else {
		if !(rexsult.ToString() == "-21575627027.15059453253376885104") {
			lang.SayString("remx3046")
		}
	}
	rexsult, err = lang.RxFromString("-6468.903738522951259063099946195").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-7877.324314273694312164407794939E+267"))
	if err != nil {
		lang.SayString("remx3048")
	} else {
		if !(rexsult.ToString() == "-6468.903738522951259063099946195") {
			lang.SayString("remx3048")
		}
	}
	rexsult, err = lang.RxFromString("-9567221.183663236817239254783372E-203").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("1650.198961256061165362319471264"))
	if err != nil {
		lang.SayString("remx3049")
	} else {
		if !(rexsult.ToString() == "-9567221.183663236817239254783372E-203") {
			lang.SayString("remx3049")
		}
	}
	rexsult, err = lang.RxFromString("8812306098770.200752139142033569E-428").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("26790.17380163975186972720427030E+568"))
	if err != nil {
		lang.SayString("remx3050")
	} else {
		if !(rexsult.ToString() == "8812306098770.200752139142033569E-428") {
			lang.SayString("remx3050")
		}
	}
	rexsult, err = lang.RxFromString("80108033.12724838718736922500904").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-706207255092.7645192310078892869"))
	if err != nil {
		lang.SayString("remx3051")
	} else {
		if !(rexsult.ToString() == "80108033.12724838718736922500904") {
			lang.SayString("remx3051")
		}
	}
	rexsult, err = lang.RxFromString("-37942846282.76101663789059003505").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-5.649456053942850351313869983197"))
	if err != nil {
		lang.SayString("remx3052")
	} else {
		if !(rexsult.ToString() == "-0.786544022188321089603127981421") {
			lang.SayString("remx3052")
		}
	}
	rexsult, err = lang.RxFromString("2838948.589837595494152150647194").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("569547026247.5469563701415715960"))
	if err != nil {
		lang.SayString("remx3054")
	} else {
		if !(rexsult.ToString() == "2838948.589837595494152150647194") {
			lang.SayString("remx3054")
		}
	}
	rexsult, err = lang.RxFromString("-57131573677452.15449921725097290").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("4669681430736.326858508715643769"))
	if err != nil {
		lang.SayString("remx3056")
	} else {
		if !(rexsult.ToString() == "-1095396508616.232197112663247672") {
			lang.SayString("remx3056")
		}
	}
	rexsult, err = lang.RxFromString("90794826.55528018781830463383411").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-5.471502270351231110027647216128"))
	if err != nil {
		lang.SayString("remx3057")
	} else {
		if !(rexsult.ToString() == "1.114274442767230442307896655232") {
			lang.SayString("remx3057")
		}
	}
	rexsult, err = lang.RxFromString("58508794729.35191160840980489138").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-47060867.24988279680824397447551"))
	if err != nil {
		lang.SayString("remx3058")
	} else {
		if !(rexsult.ToString() == "12136737.74759517576254461832107") {
			lang.SayString("remx3058")
		}
	}
	rexsult, err = lang.RxFromString("-746104.0768078474426464219416332E+006").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("9595418.300613754556671852801667E+385"))
	if err != nil {
		lang.SayString("remx3059")
	} else {
		if !(rexsult.ToString() == "-746104.0768078474426464219416332E+006") {
			lang.SayString("remx3059")
		}
	}
	rexsult, err = lang.RxFromString("-41214265628.83801241467317270595").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("1015336323798389903361978271354"))
	if err != nil {
		lang.SayString("remx3061")
	} else {
		if !(rexsult.ToString() == "-41214265628.83801241467317270595") {
			lang.SayString("remx3061")
		}
	}
	rexsult, err = lang.RxFromString("89937.39749201095570357557430822").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("82351554210093.60879476027800331"))
	if err != nil {
		lang.SayString("remx3062")
	} else {
		if !(rexsult.ToString() == "89937.39749201095570357557430822") {
			lang.SayString("remx3062")
		}
	}
	rexsult, err = lang.RxFromString("-2647593306.528617691373470059213").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-655531558709.4582168930191014461"))
	if err != nil {
		lang.SayString("remx3064")
	} else {
		if !(rexsult.ToString() == "-2647593306.528617691373470059213") {
			lang.SayString("remx3064")
		}
	}
	rexsult, err = lang.RxFromString("2904078690665765116603253099668E-329").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-71.45586619176091599264717047885E+787"))
	if err != nil {
		lang.SayString("remx3065")
	} else {
		if !(rexsult.ToString() == "2904078690665765116603253099668E-329") {
			lang.SayString("remx3065")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("remx3067")
	} else {
		if !(rexsult.ToString() == "-66913970168.62046257175566384243") {
			lang.SayString("remx3067")
		}
	}
	rexsult, err = lang.RxFromString("-84172558160661.35863831029352323").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-11271.58916600931155937291904890"))
	if err != nil {
		lang.SayString("remx3068")
	} else {
		if !(rexsult.ToString() == "-5274.95422851496534479122656860") {
			lang.SayString("remx3068")
		}
	}
	rexsult, err = lang.RxFromString("-31823131.15691583393820628480997E-440").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("92913.91582947237200286427030028E+771"))
	if err != nil {
		lang.SayString("remx3071")
	} else {
		if !(rexsult.ToString() == "-31823131.15691583393820628480997E-440") {
			lang.SayString("remx3071")
		}
	}
	rexsult, err = lang.RxFromString("55573867888.91575330563698128150").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("599.5231614736232188354393212234"))
	if err != nil {
		lang.SayString("remx3072")
	} else {
		if !(rexsult.ToString() == "85.8445030391099686478265169012") {
			lang.SayString("remx3072")
		}
	}
	rexsult, err = lang.RxFromString("-5447727448431680878699555714796E-800").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("5487207.142687001607026665515349E-362"))
	if err != nil {
		lang.SayString("remx3073")
	} else {
		if !(rexsult.ToString() == "-5447727448431680878699555714796E-800") {
			lang.SayString("remx3073")
		}
	}
	rexsult, err = lang.RxFromString("0418349404834.547110239542290134").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("09819915.92405288066606124554841"))
	if err != nil {
		lang.SayString("remx3074")
	} else {
		if !(rexsult.ToString() == "1346638.04628810400110728063718") {
			lang.SayString("remx3074")
		}
	}
	rexsult, err = lang.RxFromString("-262021.7565194737396448014286436").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-7983992600094836304387324162042E+390"))
	if err != nil {
		lang.SayString("remx3075")
	} else {
		if !(rexsult.ToString() == "-262021.7565194737396448014286436") {
			lang.SayString("remx3075")
		}
	}
	rexsult, err = lang.RxFromString("48696050631.42565380301204592392E-505").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-33868752339.85057267609967806187E+821"))
	if err != nil {
		lang.SayString("remx3076")
	} else {
		if !(rexsult.ToString() == "48696050631.42565380301204592392E-505") {
			lang.SayString("remx3076")
		}
	}
	rexsult, err = lang.RxFromString("95316999.19440144356471126680708").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-60791.33805057402845885978390435"))
	if err != nil {
		lang.SayString("remx3077")
	} else {
		if !(rexsult.ToString() == "56972.46915194096967798542896355") {
			lang.SayString("remx3077")
		}
	}
	rexsult, err = lang.RxFromString("-5326702296402708234722215224979E-136").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("8032459.450998820205916538543258"))
	if err != nil {
		lang.SayString("remx3078")
	} else {
		if !(rexsult.ToString() == "-5326702296402708234722215224979E-136") {
			lang.SayString("remx3078")
		}
	}
	rexsult, err = lang.RxFromString("-8739299372114.092482914139281669").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("507610074.7343577029345077385838"))
	if err != nil {
		lang.SayString("remx3080")
	} else {
		if !(rexsult.ToString() == "-284325487.3902691936540542102992") {
			lang.SayString("remx3080")
		}
	}
	rexsult, err = lang.RxFromString("764578.5204849936912066033177429").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("64603.13571259164812609436832506"))
	if err != nil {
		lang.SayString("remx3082")
	} else {
		if !(rexsult.ToString() == "53944.02764648556181956526616724") {
			lang.SayString("remx3082")
		}
	}
	rexsult, err = lang.RxFromString("079203.7330103777716903518367560").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("846388934347.6324036132959664705"))
	if err != nil {
		lang.SayString("remx3083")
	} else {
		if !(rexsult.ToString() == "079203.7330103777716903518367560") {
			lang.SayString("remx3083")
		}
	}
	rexsult, err = lang.RxFromString("-4278.581514688669249247007127899E-329").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("5.474973992953902631890208360829"))
	if err != nil {
		lang.SayString("remx3084")
	} else {
		if !(rexsult.ToString() == "-4278.581514688669249247007127899E-329") {
			lang.SayString("remx3084")
		}
	}
	rexsult, err = lang.RxFromString("60867019.81764798845468445196869E+651").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("6.149612565404080501157093851895E+817"))
	if err != nil {
		lang.SayString("remx3085")
	} else {
		if !(rexsult.ToString() == "60867019.81764798845468445196869E+651") {
			lang.SayString("remx3085")
		}
	}
	rexsult, err = lang.RxFromString("18554417738217.62218590965803605E-382").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-0894505909529.052378474618435782E+527"))
	if err != nil {
		lang.SayString("remx3086")
	} else {
		if !(rexsult.ToString() == "18554417738217.62218590965803605E-382") {
			lang.SayString("remx3086")
		}
	}
	rexsult, err = lang.RxFromString("69073355517144.36356688642213839").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("997784782535.6104634823627327033E+116"))
	if err != nil {
		lang.SayString("remx3087")
	} else {
		if !(rexsult.ToString() == "69073355517144.36356688642213839") {
			lang.SayString("remx3087")
		}
	}
	rexsult, err = lang.RxFromString("450282259072.8657099359104277477").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-1791307965314309175477911369824"))
	if err != nil {
		lang.SayString("remx3088")
	} else {
		if !(rexsult.ToString() == "450282259072.8657099359104277477") {
			lang.SayString("remx3088")
		}
	}
	rexsult, err = lang.RxFromString("954678411.7838149266455177850037").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("142988.7096204254529284334278794"))
	if err != nil {
		lang.SayString("remx3089")
	} else {
		if !(rexsult.ToString() == "85786.3578546028952962204808256") {
			lang.SayString("remx3089")
		}
	}
	rexsult, err = lang.RxFromString("-75492024.20990197005974241975449").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-14760421311348.35269044633000927"))
	if err != nil {
		lang.SayString("remx3091")
	} else {
		if !(rexsult.ToString() == "-75492024.20990197005974241975449") {
			lang.SayString("remx3091")
		}
	}
	rexsult, err = lang.RxFromString("317747.6972215715434186596178036E-452").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("24759763.33144824613591228097330E+092"))
	if err != nil {
		lang.SayString("remx3092")
	} else {
		if !(rexsult.ToString() == "317747.6972215715434186596178036E-452") {
			lang.SayString("remx3092")
		}
	}
	rexsult, err = lang.RxFromString("-8.153334430358647134334545353427").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-9.717872025814596548462854853522"))
	if err != nil {
		lang.SayString("remx3093")
	} else {
		if !(rexsult.ToString() == "-8.153334430358647134334545353427") {
			lang.SayString("remx3093")
		}
	}
	rexsult, err = lang.RxFromString("7.267345197492967332320456062961E-478").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("5054015481833.263541189916208065"))
	if err != nil {
		lang.SayString("remx3094")
	} else {
		if !(rexsult.ToString() == "7.267345197492967332320456062961E-478") {
			lang.SayString("remx3094")
		}
	}
	rexsult, err = lang.RxFromString("-1223354029.862567054230912271171").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("8135774223401322785475014855625"))
	if err != nil {
		lang.SayString("remx3095")
	} else {
		if !(rexsult.ToString() == "-1223354029.862567054230912271171") {
			lang.SayString("remx3095")
		}
	}
	rexsult, err = lang.RxFromString("-4673112.663442366293812346653429").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("-3429.998403142546001438238460958"))
	if err != nil {
		lang.SayString("remx3097")
	} else {
		if !(rexsult.ToString() == "-1454.838362218639853465869604204") {
			lang.SayString("remx3097")
		}
	}
	rexsult, err = lang.RxFromString("88.96492479681278079861456051103").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("386939.4621006514751889096510923E+139"))
	if err != nil {
		lang.SayString("remx3098")
	} else {
		if !(rexsult.ToString() == "88.96492479681278079861456051103") {
			lang.SayString("remx3098")
		}
	}
	rexsult, err = lang.RxFromString("064326846.4286437304788069444326E-942").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("92.23649942010862087149015091350"))
	if err != nil {
		lang.SayString("remx3099")
	} else {
		if !(rexsult.ToString() == "064326846.4286437304788069444326E-942") {
			lang.SayString("remx3099")
		}
	}
	rexsult, err = lang.RxFromString("504507.0043949324433170405699360").OpRem(lang.RxSetWithDigit(31), lang.RxFromString("605387.7175522955344659311072099"))
	if err != nil {
		lang.SayString("remx3100")
	} else {
		if !(rexsult.ToString() == "504507.0043949324433170405699360") {
			lang.SayString("remx3100")
		}
	}
	rexsult, err = lang.RxFromString("1.5283550163839789319142430253644").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-1.6578158484822969520405291379492"))
	if err != nil {
		lang.SayString("remx3201")
	} else {
		if !(rexsult.ToString() == "1.5283550163839789319142430253644") {
			lang.SayString("remx3201")
		}
	}
	rexsult, err = lang.RxFromString("-622903030605.2867503937836507326").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("6519388607.1331855704471328795821"))
	if err != nil {
		lang.SayString("remx3202")
	} else {
		if !(rexsult.ToString() == "-3561112927.6341212013060271723005") {
			lang.SayString("remx3202")
		}
	}
	rexsult, err = lang.RxFromString("-5675915.2465457487632250245209054").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("73913909880.381367895205086027416"))
	if err != nil {
		lang.SayString("remx3203")
	} else {
		if !(rexsult.ToString() == "-5675915.2465457487632250245209054") {
			lang.SayString("remx3203")
		}
	}
	rexsult, err = lang.RxFromString("97.647321172555144900685605318635").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("4.8620911587547548751209841570885"))
	if err != nil {
		lang.SayString("remx3204")
	} else {
		if !(rexsult.ToString() == "0.4054979974600473982659221768650") {
			lang.SayString("remx3204")
		}
	}
	rexsult, err = lang.RxFromString("-9717253267024.5380651435435603552").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-2669.2539695193820424002013488480E+363"))
	if err != nil {
		lang.SayString("remx3205")
	} else {
		if !(rexsult.ToString() == "-9717253267024.5380651435435603552") {
			lang.SayString("remx3205")
		}
	}
	rexsult, err = lang.RxFromString("-4.0817391717190128506083001702335E-767").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("12772.807105920712660991033689206"))
	if err != nil {
		lang.SayString("remx3206")
	} else {
		if !(rexsult.ToString() == "-4.0817391717190128506083001702335E-767") {
			lang.SayString("remx3206")
		}
	}
	rexsult, err = lang.RxFromString("68625322655934146845003028928647").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-59.634169944840280159782488098700"))
	if err != nil {
		lang.SayString("remx3207")
	} else {
		if !(rexsult.ToString() == "28.201254004897257552939369449600") {
			lang.SayString("remx3207")
		}
	}
	rexsult, err = lang.RxFromString("732515.76532049290815665856727641").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-92134479835821.319619827023729829"))
	if err != nil {
		lang.SayString("remx3208")
	} else {
		if !(rexsult.ToString() == "732515.76532049290815665856727641") {
			lang.SayString("remx3208")
		}
	}
	rexsult, err = lang.RxFromString("-30.458011942978338421676454778733").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-5023372024597665102336430410403E+831"))
	if err != nil {
		lang.SayString("remx3209")
	} else {
		if !(rexsult.ToString() == "-30.458011942978338421676454778733") {
			lang.SayString("remx3209")
		}
	}
	rexsult, err = lang.RxFromString("-89640.094149414644660480286430462").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-58703419758.800889903227509215474"))
	if err != nil {
		lang.SayString("remx3210")
	} else {
		if !(rexsult.ToString() == "-89640.094149414644660480286430462") {
			lang.SayString("remx3210")
		}
	}
	rexsult, err = lang.RxFromString("913391.42744224458216174967853722").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-21051638.816432817393202262710630E+432"))
	if err != nil {
		lang.SayString("remx3212")
	} else {
		if !(rexsult.ToString() == "913391.42744224458216174967853722") {
			lang.SayString("remx3212")
		}
	}
	rexsult, err = lang.RxFromString("-917591456829.12109027484399536567").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-28892177726858026955513438843371E+708"))
	if err != nil {
		lang.SayString("remx3213")
	} else {
		if !(rexsult.ToString() == "-917591456829.12109027484399536567") {
			lang.SayString("remx3213")
		}
	}
	rexsult, err = lang.RxFromString("34938410840645.913399699219228218").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("30.818220393242402846077755480548"))
	if err != nil {
		lang.SayString("remx3214")
	} else {
		if !(rexsult.ToString() == "24.283226805899273551376371736548") {
			lang.SayString("remx3214")
		}
	}
	rexsult, err = lang.RxFromString("6034.7374411022598081745006769023E-517").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("29771833428054709077850588904653"))
	if err != nil {
		lang.SayString("remx3215")
	} else {
		if !(rexsult.ToString() == "6034.7374411022598081745006769023E-517") {
			lang.SayString("remx3215")
		}
	}
	rexsult, err = lang.RxFromString("-5565747671734.1686009705574503152").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-490.30899494881071282787487030303"))
	if err != nil {
		lang.SayString("remx3216")
	} else {
		if !(rexsult.ToString() == "-178.99949336276892685183308348801") {
			lang.SayString("remx3216")
		}
	}
	rexsult, err = lang.RxFromString("319545511.89203495546689273564728E+036").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-2955943533943321899150310192061"))
	if err != nil {
		lang.SayString("remx3217")
	} else {
		if !(rexsult.ToString() == "2029642017122316721531728309258") {
			lang.SayString("remx3217")
		}
	}
	rexsult, err = lang.RxFromString("-36852134.84194296250843579428931").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-5830629.8347085025808716360357940"))
	if err != nil {
		lang.SayString("remx3218")
	} else {
		if !(rexsult.ToString() == "-1868355.8336919470232059780745460") {
			lang.SayString("remx3218")
		}
	}
	rexsult, err = lang.RxFromString("8.6021905001798578659275880018221E-374").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-39505285344943.729681835377530908"))
	if err != nil {
		lang.SayString("remx3219")
	} else {
		if !(rexsult.ToString() == "8.6021905001798578659275880018221E-374") {
			lang.SayString("remx3219")
		}
	}
	rexsult, err = lang.RxFromString("-54863165.152174109720312887805017").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("736.1398476560169141105328256628"))
	if err != nil {
		lang.SayString("remx3220")
	} else {
		if !(rexsult.ToString() == "-134.5860664811454830973740198416") {
			lang.SayString("remx3220")
		}
	}
	rexsult, err = lang.RxFromString("-3263743464517851012531708810307").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("2457206.2471248382136273643208109"))
	if err != nil {
		lang.SayString("remx3221")
	} else {
		if !(rexsult.ToString() == "-1417336.7573398366062994535940062") {
			lang.SayString("remx3221")
		}
	}
	rexsult, err = lang.RxFromString("2856586744.0548637797291151154902E-895").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("953545637646.57694835860339582821E+080"))
	if err != nil {
		lang.SayString("remx3222")
	} else {
		if !(rexsult.ToString() == "2856586744.0548637797291151154902E-895") {
			lang.SayString("remx3222")
		}
	}
	rexsult, err = lang.RxFromString("5624157233.3433661009203529937625").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("626098409265.93738586750090160638"))
	if err != nil {
		lang.SayString("remx3223")
	} else {
		if !(rexsult.ToString() == "5624157233.3433661009203529937625") {
			lang.SayString("remx3223")
		}
	}
	rexsult, err = lang.RxFromString("-213499362.91476998701834067092611").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("419272438.02555757699863022643444"))
	if err != nil {
		lang.SayString("remx3224")
	} else {
		if !(rexsult.ToString() == "-213499362.91476998701834067092611") {
			lang.SayString("remx3224")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("remx3225")
	} else {
		if !(rexsult.ToString() == "0.6016219662519115373766970119100") {
			lang.SayString("remx3225")
		}
	}
	rexsult, err = lang.RxFromString("-74396862273800.625679130772935550").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-115616605.52826981284183992013157"))
	if err != nil {
		lang.SayString("remx3227")
	} else {
		if !(rexsult.ToString() == "-4565075.09478147646296920746797") {
			lang.SayString("remx3227")
		}
	}
	rexsult, err = lang.RxFromString("67585560.562561229497720955705979").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("826.96290288608566737177503451613"))
	if err != nil {
		lang.SayString("remx3228")
	} else {
		if !(rexsult.ToString() == "363.39839010616042789746007924349") {
			lang.SayString("remx3228")
		}
	}
	rexsult, err = lang.RxFromString("6877386868.9498051860742298735156E-232").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("390.3154289860643509393769754551"))
	if err != nil {
		lang.SayString("remx3229")
	} else {
		if !(rexsult.ToString() == "6877386868.9498051860742298735156E-232") {
			lang.SayString("remx3229")
		}
	}
	rexsult, err = lang.RxFromString("-1647335.201144609256134925838937").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-186654823782.50459802235024230856"))
	if err != nil {
		lang.SayString("remx3230")
	} else {
		if !(rexsult.ToString() == "-1647335.201144609256134925838937") {
			lang.SayString("remx3230")
		}
	}
	rexsult, err = lang.RxFromString("-6.6547424012516834662011706165297").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-574454585580.06215974139884746441"))
	if err != nil {
		lang.SayString("remx3232")
	} else {
		if !(rexsult.ToString() == "-6.6547424012516834662011706165297") {
			lang.SayString("remx3232")
		}
	}
	rexsult, err = lang.RxFromString("-27627.758745381267780885067447671").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-23385972189441.709586433758111062"))
	if err != nil {
		lang.SayString("remx3233")
	} else {
		if !(rexsult.ToString() == "-27627.758745381267780885067447671") {
			lang.SayString("remx3233")
		}
	}
	rexsult, err = lang.RxFromString("-70454070095869.70717871212601390").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-6200178.370249260009168888392406"))
	if err != nil {
		lang.SayString("remx3237")
	} else {
		if !(rexsult.ToString() == "-4833345.467866203920028883583808") {
			lang.SayString("remx3237")
		}
	}
	rexsult, err = lang.RxFromString("-5168.2214111091132913776042214525").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-5690274.0971173476527123568627720"))
	if err != nil {
		lang.SayString("remx3239")
	} else {
		if !(rexsult.ToString() == "-5168.2214111091132913776042214525") {
			lang.SayString("remx3239")
		}
	}
	rexsult, err = lang.RxFromString("33783.060857197067391462144517964").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-2070.0806959465088198508322815406"))
	if err != nil {
		lang.SayString("remx3240")
	} else {
		if !(rexsult.ToString() == "661.7697220529262738488280133144") {
			lang.SayString("remx3240")
		}
	}
	rexsult, err = lang.RxFromString("42207435091050.840296353874733169E-905").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("73330633078.828216018536326743325E+976"))
	if err != nil {
		lang.SayString("remx3241")
	} else {
		if !(rexsult.ToString() == "42207435091050.840296353874733169E-905") {
			lang.SayString("remx3241")
		}
	}
	rexsult, err = lang.RxFromString("-71800.806700868784841045406679641").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-39617456964250697902519150526701"))
	if err != nil {
		lang.SayString("remx3242")
	} else {
		if !(rexsult.ToString() == "-71800.806700868784841045406679641") {
			lang.SayString("remx3242")
		}
	}
	rexsult, err = lang.RxFromString("53627480801.631504892310016062160").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("328259.56947661049313311983109503"))
	if err != nil {
		lang.SayString("remx3243")
	} else {
		if !(rexsult.ToString() == "43195.80712523964536237599604393") {
			lang.SayString("remx3243")
		}
	}
	rexsult, err = lang.RxFromString("-5052601598.5559371338428368438728").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-97855372.224321664785314782556064"))
	if err != nil {
		lang.SayString("remx3244")
	} else {
		if !(rexsult.ToString() == "-61977615.115532229791782933513536") {
			lang.SayString("remx3244")
		}
	}
	rexsult, err = lang.RxFromString("4208134320733.7069742988228068191E+146").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("4270869.1760149477598920960628392E+471"))
	if err != nil {
		lang.SayString("remx3245")
	} else {
		if !(rexsult.ToString() == "4208134320733.7069742988228068191E+146") {
			lang.SayString("remx3245")
		}
	}
	rexsult, err = lang.RxFromString("-440220916.66716743026896931194749").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-102725.01594377871560564824358775"))
	if err != nil {
		lang.SayString("remx3248")
	} else {
		if !(rexsult.ToString() == "-44223.34807563389876658817398125") {
			lang.SayString("remx3248")
		}
	}
	rexsult, err = lang.RxFromString("-46.250735086006350517943464758019").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("14656357555174.263295266074908024"))
	if err != nil {
		lang.SayString("remx3249")
	} else {
		if !(rexsult.ToString() == "-46.250735086006350517943464758019") {
			lang.SayString("remx3249")
		}
	}
	rexsult, err = lang.RxFromString("96668419802749.555738010239087449E-838").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-19498732131365824921639467044927E-542"))
	if err != nil {
		lang.SayString("remx3251")
	} else {
		if !(rexsult.ToString() == "96668419802749.555738010239087449E-838") {
			lang.SayString("remx3251")
		}
	}
	rexsult, err = lang.RxFromString("-62663404777.352508979582846931050").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("9.2570938837239134052589184917186E+916"))
	if err != nil {
		lang.SayString("remx3253")
	} else {
		if !(rexsult.ToString() == "-62663404777.352508979582846931050") {
			lang.SayString("remx3253")
		}
	}
	rexsult, err = lang.RxFromString("1.744601214474560992754529320172E-827").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-17.353669504253419489494030651507E-631"))
	if err != nil {
		lang.SayString("remx3254")
	} else {
		if !(rexsult.ToString() == "1.744601214474560992754529320172E-827") {
			lang.SayString("remx3254")
		}
	}
	rexsult, err = lang.RxFromString("0367191549036702224827734853471").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("4410320662415266533763143837742E+721"))
	if err != nil {
		lang.SayString("remx3255")
	} else {
		if !(rexsult.ToString() == "0367191549036702224827734853471") {
			lang.SayString("remx3255")
		}
	}
	rexsult, err = lang.RxFromString("097704116.4492566721965710365073").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-96736.400939809433556067504289145"))
	if err != nil {
		lang.SayString("remx3256")
	} else {
		if !(rexsult.ToString() == "351.500049144304942857175263550") {
			lang.SayString("remx3256")
		}
	}
	rexsult, err = lang.RxFromString("-86863.276249466008289214762980838").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("531.50602652732088208397655484476"))
	if err != nil {
		lang.SayString("remx3260")
	} else {
		if !(rexsult.ToString() == "-227.79392551270450952658454114212") {
			lang.SayString("remx3260")
		}
	}
	rexsult, err = lang.RxFromString("-40707.169006771111855573524157083").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-68795521421321853333274411827749"))
	if err != nil {
		lang.SayString("remx3261")
	} else {
		if !(rexsult.ToString() == "-40707.169006771111855573524157083") {
			lang.SayString("remx3261")
		}
	}
	rexsult, err = lang.RxFromString("-4245360967593.9206771555839718158E-682").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-3.119606239042530207103203508057"))
	if err != nil {
		lang.SayString("remx3263")
	} else {
		if !(rexsult.ToString() == "-4245360967593.9206771555839718158E-682") {
			lang.SayString("remx3263")
		}
	}
	rexsult, err = lang.RxFromString("-3422145405774.0800213000547612240E-023").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-60810.964656409650839011321706310"))
	if err != nil {
		lang.SayString("remx3264")
	} else {
		if !(rexsult.ToString() == "-3422145405774.0800213000547612240E-023") {
			lang.SayString("remx3264")
		}
	}
	rexsult, err = lang.RxFromString("-24521811.07649485796598387627478E-661").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-94860846133404815410816234000694"))
	if err != nil {
		lang.SayString("remx3265")
	} else {
		if !(rexsult.ToString() == "-24521811.07649485796598387627478E-661") {
			lang.SayString("remx3265")
		}
	}
	rexsult, err = lang.RxFromString("-5042529937498.8944492248538951438").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("3891904674.4549170968807145612549"))
	if err != nil {
		lang.SayString("remx3266")
	} else {
		if !(rexsult.ToString() == "-2513384079.7768087643285383187045") {
			lang.SayString("remx3266")
		}
	}
	rexsult, err = lang.RxFromString("-535824163.54531747646293693868651E-665").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("2732988.5891363639325008206099712"))
	if err != nil {
		lang.SayString("remx3267")
	} else {
		if !(rexsult.ToString() == "-535824163.54531747646293693868651E-665") {
			lang.SayString("remx3267")
		}
	}
	rexsult, err = lang.RxFromString("24032.702008553084252925140858134E-509").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("52864854.899420632375589206704068"))
	if err != nil {
		lang.SayString("remx3268")
	} else {
		if !(rexsult.ToString() == "24032.702008553084252925140858134E-509") {
			lang.SayString("remx3268")
		}
	}
	rexsult, err = lang.RxFromString("71553220259.938950698030519906727E-496").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("754.44220417415325444943566016062"))
	if err != nil {
		lang.SayString("remx3269")
	} else {
		if !(rexsult.ToString() == "71553220259.938950698030519906727E-496") {
			lang.SayString("remx3269")
		}
	}
	rexsult, err = lang.RxFromString("95.490751127249945886828257312118").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("987.01498316307365714167410690192E+133"))
	if err != nil {
		lang.SayString("remx3272")
	} else {
		if !(rexsult.ToString() == "95.490751127249945886828257312118") {
			lang.SayString("remx3272")
		}
	}
	rexsult, err = lang.RxFromString("69434850287.460788550936730296153").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-35119136549015044241569827542264"))
	if err != nil {
		lang.SayString("remx3273")
	} else {
		if !(rexsult.ToString() == "69434850287.460788550936730296153") {
			lang.SayString("remx3273")
		}
	}
	rexsult, err = lang.RxFromString("-392.22739924621965621739098725407").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-65551274.987160998195282109612136"))
	if err != nil {
		lang.SayString("remx3274")
	} else {
		if !(rexsult.ToString() == "-392.22739924621965621739098725407") {
			lang.SayString("remx3274")
		}
	}
	rexsult, err = lang.RxFromString("6413265.4423561191792972085539457").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("24514.222704714139350026165721146"))
	if err != nil {
		lang.SayString("remx3275")
	} else {
		if !(rexsult.ToString() == "15053.316425728808940379300726594") {
			lang.SayString("remx3275")
		}
	}
	rexsult, err = lang.RxFromString("378204716633.40024100602896057615").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-0300218714378.322231269606216516"))
	if err != nil {
		lang.SayString("remx3277")
	} else {
		if !(rexsult.ToString() == "77986002255.07800973642274406015") {
			lang.SayString("remx3277")
		}
	}
	rexsult, err = lang.RxFromString("-3554.5895974968741465654022772100E-073").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("9752.0428746722497621936998533848E+516"))
	if err != nil {
		lang.SayString("remx3279")
	} else {
		if !(rexsult.ToString() == "-3554.5895974968741465654022772100E-073") {
			lang.SayString("remx3279")
		}
	}
	rexsult, err = lang.RxFromString("750187685.63632608923397234391668").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("4633194252863.6730625972669246025"))
	if err != nil {
		lang.SayString("remx3280")
	} else {
		if !(rexsult.ToString() == "750187685.63632608923397234391668") {
			lang.SayString("remx3280")
		}
	}
	rexsult, err = lang.RxFromString("30190034242853.251165951457589386E-028").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("8038885676.3204238322976087799751E+018"))
	if err != nil {
		lang.SayString("remx3281")
	} else {
		if !(rexsult.ToString() == "30190034242853.251165951457589386E-028") {
			lang.SayString("remx3281")
		}
	}
	rexsult, err = lang.RxFromString("65.537942676774715953400768803539").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("125946917260.87536506197191782198"))
	if err != nil {
		lang.SayString("remx3282")
	} else {
		if !(rexsult.ToString() == "65.537942676774715953400768803539") {
			lang.SayString("remx3282")
		}
	}
	rexsult, err = lang.RxFromString("8015272348677.5489394183881813700").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("949.23027111499779258284877920022"))
	if err != nil {
		lang.SayString("remx3283")
	} else {
		if !(rexsult.ToString() == "527.78228041355742397895303690364") {
			lang.SayString("remx3283")
		}
	}
	rexsult, err = lang.RxFromString("-17544189017145.710117633021800426E-537").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("292178000.06450804618299520094843"))
	if err != nil {
		lang.SayString("remx3285")
	} else {
		if !(rexsult.ToString() == "-17544189017145.710117633021800426E-537") {
			lang.SayString("remx3285")
		}
	}
	rexsult, err = lang.RxFromString("-506650.99395649907139204052441630").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("11.018427502031650148962067367158"))
	if err != nil {
		lang.SayString("remx3286")
	} else {
		if !(rexsult.ToString() == "-1.660558079734242466742739640844") {
			lang.SayString("remx3286")
		}
	}
	rexsult, err = lang.RxFromString("4846835159.5922372307656000769395E-241").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-84.001893040865864590122330800768"))
	if err != nil {
		lang.SayString("remx3287")
	} else {
		if !(rexsult.ToString() == "4846835159.5922372307656000769395E-241") {
			lang.SayString("remx3287")
		}
	}
	rexsult, err = lang.RxFromString("-35.029311013822259358116556164908").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-3994308878.1994645313414534209127"))
	if err != nil {
		lang.SayString("remx3288")
	} else {
		if !(rexsult.ToString() == "-35.029311013822259358116556164908") {
			lang.SayString("remx3288")
		}
	}
	rexsult, err = lang.RxFromString("7606663750.6735265233044420887018E+166").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-5491814639.4484565418284686127552E+365"))
	if err != nil {
		lang.SayString("remx3289")
	} else {
		if !(rexsult.ToString() == "7606663750.6735265233044420887018E+166") {
			lang.SayString("remx3289")
		}
	}
	rexsult, err = lang.RxFromString("97271576094.456406729671729224827").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-1.5412563837540810793697955063295E+554"))
	if err != nil {
		lang.SayString("remx3291")
	} else {
		if !(rexsult.ToString() == "97271576094.456406729671729224827") {
			lang.SayString("remx3291")
		}
	}
	rexsult, err = lang.RxFromString("41139789894.401826915757391777563").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-1.8729920305671057957156159690823E-020"))
	if err != nil {
		lang.SayString("remx3292")
	} else {
		if !(rexsult.ToString() == "6.98141022640544018935102953527E-22") {
			lang.SayString("remx3292")
		}
	}
	rexsult, err = lang.RxFromString("23777.857951114967684767609401626").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("720760.03897144157012301385227528"))
	if err != nil {
		lang.SayString("remx3295")
	} else {
		if !(rexsult.ToString() == "23777.857951114967684767609401626") {
			lang.SayString("remx3295")
		}
	}
	rexsult, err = lang.RxFromString("-21337853323980858055292469611948").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("6080272840.3071490445256786982100E+532"))
	if err != nil {
		lang.SayString("remx3296")
	} else {
		if !(rexsult.ToString() == "-21337853323980858055292469611948") {
			lang.SayString("remx3296")
		}
	}
	rexsult, err = lang.RxFromString("-818409238.0423893439849743856947").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("756.39156265972753259267069158760"))
	if err != nil {
		lang.SayString("remx3297")
	} else {
		if !(rexsult.ToString() == "-374.76862809126749803143314108840") {
			lang.SayString("remx3297")
		}
	}
	rexsult, err = lang.RxFromString("47567380384943.87013600286155046").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("65.084709374908275826942076480326"))
	if err != nil {
		lang.SayString("remx3298")
	} else {
		if !(rexsult.ToString() == "56.134058687770878126430844955520") {
			lang.SayString("remx3298")
		}
	}
	rexsult, err = lang.RxFromString("-296615544.05897984545127115290251").OpRem(lang.RxSetWithDigit(32), lang.RxFromString("-5416115.4315053536014016450973264"))
	if err != nil {
		lang.SayString("remx3299")
	} else {
		if !(rexsult.ToString() == "-4145310.7576907509755823176468844") {
			lang.SayString("remx3299")
		}
	}
	rexsult, err = lang.RxFromString("042.668056830485571428877212944418").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-01364112374639.4941124016455321071"))
	if err != nil {
		lang.SayString("remx3401")
	} else {
		if !(rexsult.ToString() == "042.668056830485571428877212944418") {
			lang.SayString("remx3401")
		}
	}
	rexsult, err = lang.RxFromString("81721.5803096185422787702538493471").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("900099473.245809628076790757217328"))
	if err != nil {
		lang.SayString("remx3403")
	} else {
		if !(rexsult.ToString() == "81721.5803096185422787702538493471") {
			lang.SayString("remx3403")
		}
	}
	rexsult, err = lang.RxFromString("3991.56734635183403814636354392163E-807").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("72.3239822255871305731314565069132"))
	if err != nil {
		lang.SayString("remx3404")
	} else {
		if !(rexsult.ToString() == "3991.56734635183403814636354392163E-807") {
			lang.SayString("remx3404")
		}
	}
	rexsult, err = lang.RxFromString("-66.3393308595957726456416979163306").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("5.08486573050459213864684589662559"))
	if err != nil {
		lang.SayString("remx3405")
	} else {
		if !(rexsult.ToString() == "-0.23607636303607484323270126019793") {
			lang.SayString("remx3405")
		}
	}
	rexsult, err = lang.RxFromString("-019133598.609812524622150421584346").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-858439846.628367734642622922030051"))
	if err != nil {
		lang.SayString("remx3407")
	} else {
		if !(rexsult.ToString() == "-019133598.609812524622150421584346") {
			lang.SayString("remx3407")
		}
	}
	rexsult, err = lang.RxFromString("465.351982159046525762715549761814").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("240444.975944666924657629172844780"))
	if err != nil {
		lang.SayString("remx3408")
	} else {
		if !(rexsult.ToString() == "465.351982159046525762715549761814") {
			lang.SayString("remx3408")
		}
	}
	rexsult, err = lang.RxFromString("28275236927392.4960902824105246047").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("28212038.4825243127096613158419270E+422"))
	if err != nil {
		lang.SayString("remx3410")
	} else {
		if !(rexsult.ToString() == "28275236927392.4960902824105246047") {
			lang.SayString("remx3410")
		}
	}
	rexsult, err = lang.RxFromString("11791.8644211874630234271801789996").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-8.45457275930363860982261343159741"))
	if err != nil {
		lang.SayString("remx3411")
	} else {
		if !(rexsult.ToString() == "6.18999471819080133445705535281046") {
			lang.SayString("remx3411")
		}
	}
	rexsult, err = lang.RxFromString("44.7085340739581668975502342787578").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-9337.05408133023920640485556647937"))
	if err != nil {
		lang.SayString("remx3412")
	} else {
		if !(rexsult.ToString() == "44.7085340739581668975502342787578") {
			lang.SayString("remx3412")
		}
	}
	rexsult, err = lang.RxFromString("-367399415798804503177950040443482").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-54845683.9691776397285506712812754"))
	if err != nil {
		lang.SayString("remx3414")
	} else {
		if !(rexsult.ToString() == "-16714095.6549657189177857892292804") {
			lang.SayString("remx3414")
		}
	}
	rexsult, err = lang.RxFromString("-2.87155919781024108503670175443740").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("89529730130.6427881332776797193807"))
	if err != nil {
		lang.SayString("remx3415")
	} else {
		if !(rexsult.ToString() == "-2.87155919781024108503670175443740") {
			lang.SayString("remx3415")
		}
	}
	rexsult, err = lang.RxFromString("611655569568.832698912762075889186").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("010182743219.475839030505966016982"))
	if err != nil {
		lang.SayString("remx3417")
	} else {
		if !(rexsult.ToString() == "690976400.282357082404114870266") {
			lang.SayString("remx3417")
		}
	}
	rexsult, err = lang.RxFromString("3457947.39062863674882672518304442").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-01.9995218868908849056866549811425"))
	if err != nil {
		lang.SayString("remx3418")
	} else {
		if !(rexsult.ToString() == "0.2332240699744359979851713353525") {
			lang.SayString("remx3418")
		}
	}
	rexsult, err = lang.RxFromString("-5568057.17870139549478277980540034").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-407906443.141342175740471849723638"))
	if err != nil {
		lang.SayString("remx3420")
	} else {
		if !(rexsult.ToString() == "-5568057.17870139549478277980540034") {
			lang.SayString("remx3420")
		}
	}
	rexsult, err = lang.RxFromString("9804385273.49533524416415189990857").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("84.1433929743544659553964804646569"))
	if err != nil {
		lang.SayString("remx3421")
	} else {
		if !(rexsult.ToString() == "69.1423069734476624350435642749915") {
			lang.SayString("remx3421")
		}
	}
	rexsult, err = lang.RxFromString("-5234910986592.18801727046580014273E-547").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-5874220715892.91440069210512515154"))
	if err != nil {
		lang.SayString("remx3422")
	} else {
		if !(rexsult.ToString() == "-5234910986592.18801727046580014273E-547") {
			lang.SayString("remx3422")
		}
	}
	rexsult, err = lang.RxFromString("698416560151955285929747633786867E-495").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("51754681.6784872628933218985216916E-266"))
	if err != nil {
		lang.SayString("remx3423")
	} else {
		if !(rexsult.ToString() == "698416560151955285929747633786867E-495") {
			lang.SayString("remx3423")
		}
	}
	rexsult, err = lang.RxFromString("-32174291345686.5371446616670961807").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("79518863759385.5925052747867099091E+408"))
	if err != nil {
		lang.SayString("remx3425")
	} else {
		if !(rexsult.ToString() == "-32174291345686.5371446616670961807") {
			lang.SayString("remx3425")
		}
	}
	rexsult, err = lang.RxFromString("-8151730494.53190523620899410544099E+688").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-93173.0631474527142307644239919480E+900"))
	if err != nil {
		lang.SayString("remx3426")
	} else {
		if !(rexsult.ToString() == "-8151730494.53190523620899410544099E+688") {
			lang.SayString("remx3426")
		}
	}
	rexsult, err = lang.RxFromString("1.33649801345976199708341799505220").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-56623.0530039528969825480755159562E+459"))
	if err != nil {
		lang.SayString("remx3427")
	} else {
		if !(rexsult.ToString() == "1.33649801345976199708341799505220") {
			lang.SayString("remx3427")
		}
	}
	rexsult, err = lang.RxFromString("4286562.76568866751577306056498271").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("6286.77291578497580015557979349893E+820"))
	if err != nil {
		lang.SayString("remx3429")
	} else {
		if !(rexsult.ToString() == "4286562.76568866751577306056498271") {
			lang.SayString("remx3429")
		}
	}
	rexsult, err = lang.RxFromString("-765782.827432642697305644096365566").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("67.1634368459576834692758114618652"))
	if err != nil {
		lang.SayString("remx3430")
	} else {
		if !(rexsult.ToString() == "-52.4839518791480724305698888408548") {
			lang.SayString("remx3430")
		}
	}
	rexsult, err = lang.RxFromString("46.2835931916106252756465724211276").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("59.2989237834093118332826617957791"))
	if err != nil {
		lang.SayString("remx3431")
	} else {
		if !(rexsult.ToString() == "46.2835931916106252756465724211276") {
			lang.SayString("remx3431")
		}
	}
	rexsult, err = lang.RxFromString("-3029555.82298840234029474459694644").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("857535844655004737373089601128532"))
	if err != nil {
		lang.SayString("remx3432")
	} else {
		if !(rexsult.ToString() == "-3029555.82298840234029474459694644") {
			lang.SayString("remx3432")
		}
	}
	rexsult, err = lang.RxFromString("-0138466789523.10694176543700501945E-948").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("481026979918882487383654367924619"))
	if err != nil {
		lang.SayString("remx3433")
	} else {
		if !(rexsult.ToString() == "-0138466789523.10694176543700501945E-948") {
			lang.SayString("remx3433")
		}
	}
	rexsult, err = lang.RxFromString("-9593566466.96690575714244442109870").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-87632034347.4845477961976776833770E+769"))
	if err != nil {
		lang.SayString("remx3434")
	} else {
		if !(rexsult.ToString() == "-9593566466.96690575714244442109870") {
			lang.SayString("remx3434")
		}
	}
	rexsult, err = lang.RxFromString("-3189.30765477670526823106100241863E-898").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("565688889.355241946154894311253202E-466"))
	if err != nil {
		lang.SayString("remx3435")
	} else {
		if !(rexsult.ToString() == "-3189.30765477670526823106100241863E-898") {
			lang.SayString("remx3435")
		}
	}
	rexsult, err = lang.RxFromString("-17084552395.6714834680088150543965").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-631925802672.685034379197328370812E+527"))
	if err != nil {
		lang.SayString("remx3436")
	} else {
		if !(rexsult.ToString() == "-17084552395.6714834680088150543965") {
			lang.SayString("remx3436")
		}
	}
	rexsult, err = lang.RxFromString("-763.440067781256632695791981893608").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("19.9263811350611007833220620745413"))
	if err != nil {
		lang.SayString("remx3438")
	} else {
		if !(rexsult.ToString() == "-6.2375846489348029295536230610386") {
			lang.SayString("remx3438")
		}
	}
	rexsult, err = lang.RxFromString("-0970725697662.27605454336231195463").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-4541.41897546697187157913886433474"))
	if err != nil {
		lang.SayString("remx3441")
	} else {
		if !(rexsult.ToString() == "-2972.12171050214753770792631747550") {
			lang.SayString("remx3441")
		}
	}
	rexsult, err = lang.RxFromString("-808178238631844268316111259558675").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-598400.265108644514211244980426520"))
	if err != nil {
		lang.SayString("remx3442")
	} else {
		if !(rexsult.ToString() == "-585452.097764536570956813681556320") {
			lang.SayString("remx3442")
		}
	}
	rexsult, err = lang.RxFromString("-9.90826595069053564311371766315200").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-031.625916781307847864872329806646"))
	if err != nil {
		lang.SayString("remx3443")
	} else {
		if !(rexsult.ToString() == "-9.90826595069053564311371766315200") {
			lang.SayString("remx3443")
		}
	}
	rexsult, err = lang.RxFromString("-196925.469891897719160698483752907").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-41268.9975444533794067723958739778"))
	if err != nil {
		lang.SayString("remx3444")
	} else {
		if !(rexsult.ToString() == "-31849.4797140842015336089002569958") {
			lang.SayString("remx3444")
		}
	}
	rexsult, err = lang.RxFromString("74815000.4716875558358937279052903").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-690425401708167622194241915195001E+891"))
	if err != nil {
		lang.SayString("remx3447")
	} else {
		if !(rexsult.ToString() == "74815000.4716875558358937279052903") {
			lang.SayString("remx3447")
		}
	}
	rexsult, err = lang.RxFromString("-1683993.51210241555668790556759021").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-72394384927344.8402585228267493374"))
	if err != nil {
		lang.SayString("remx3448")
	} else {
		if !(rexsult.ToString() == "-1683993.51210241555668790556759021") {
			lang.SayString("remx3448")
		}
	}
	rexsult, err = lang.RxFromString("-763.148530974741766171756970448158").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("517370.808956957601473642272664647"))
	if err != nil {
		lang.SayString("remx3449")
	} else {
		if !(rexsult.ToString() == "-763.148530974741766171756970448158") {
			lang.SayString("remx3449")
		}
	}
	rexsult, err = lang.RxFromString("-77.5841338812312523460591226178754").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-927540422.641025050968830154578151E+524"))
	if err != nil {
		lang.SayString("remx3450")
	} else {
		if !(rexsult.ToString() == "-77.5841338812312523460591226178754") {
			lang.SayString("remx3450")
		}
	}
	rexsult, err = lang.RxFromString("5176295309.89943746236102209837813").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-129733.103628797477167908698565465"))
	if err != nil {
		lang.SayString("remx3451")
	} else {
		if !(rexsult.ToString() == "74208.214046920838632934314641965") {
			lang.SayString("remx3451")
		}
	}
	rexsult, err = lang.RxFromString("-254346232981353293392888785643245").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-764.416902486152093758287929678445"))
	if err != nil {
		lang.SayString("remx3454")
	} else {
		if !(rexsult.ToString() == "-508.299323962538610580669092979500") {
			lang.SayString("remx3454")
		}
	}
	rexsult, err = lang.RxFromString("-2875.36745499544177964804277329726").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-13187.8492045054802205691248267638"))
	if err != nil {
		lang.SayString("remx3455")
	} else {
		if !(rexsult.ToString() == "-2875.36745499544177964804277329726") {
			lang.SayString("remx3455")
		}
	}
	rexsult, err = lang.RxFromString("7191753.79974136447257468282073718").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("81.3878426176038487948375777384429"))
	if err != nil {
		lang.SayString("remx3458")
	} else {
		if !(rexsult.ToString() == "79.8625220355815164499390351500273") {
			lang.SayString("remx3458")
		}
	}
	rexsult, err = lang.RxFromString("502975804.069864536104621700404770").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("684.790028432074527960269515227243"))
	if err != nil {
		lang.SayString("remx3459")
	} else {
		if !(rexsult.ToString() == "267.346619523615915582548420925472") {
			lang.SayString("remx3459")
		}
	}
	rexsult, err = lang.RxFromString("1505368.42063731861590460453659570").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-465242.678439951462767630022819105"))
	if err != nil {
		lang.SayString("remx3460")
	} else {
		if !(rexsult.ToString() == "109640.385317464227601714468138385") {
			lang.SayString("remx3460")
		}
	}
	rexsult, err = lang.RxFromString("69225023.2850142784063416137144829").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("8584050.06648191798834819995325693"))
	if err != nil {
		lang.SayString("remx3461")
	} else {
		if !(rexsult.ToString() == "552622.75315893449955601408842746") {
			lang.SayString("remx3461")
		}
	}
	rexsult, err = lang.RxFromString("-527566.521273992424649346837337602E-408").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-834662.599983953345718523807123972"))
	if err != nil {
		lang.SayString("remx3463")
	} else {
		if !(rexsult.ToString() == "-527566.521273992424649346837337602E-408") {
			lang.SayString("remx3463")
		}
	}
	rexsult, err = lang.RxFromString("4.51117459466491451401835756593793").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("3873385.19981811640063144338087230"))
	if err != nil {
		lang.SayString("remx3466")
	} else {
		if !(rexsult.ToString() == "4.51117459466491451401835756593793") {
			lang.SayString("remx3466")
		}
	}
	rexsult, err = lang.RxFromString("49553763474698.8118661758811091120").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("36.1713861293896216593840817950781E+410"))
	if err != nil {
		lang.SayString("remx3467")
	} else {
		if !(rexsult.ToString() == "49553763474698.8118661758811091120") {
			lang.SayString("remx3467")
		}
	}
	rexsult, err = lang.RxFromString("-20101668541.7472260497594230105456").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-395562148.345003931161532101109964"))
	if err != nil {
		lang.SayString("remx3469")
	} else {
		if !(rexsult.ToString() == "-323561124.497029491682817955047400") {
			lang.SayString("remx3469")
		}
	}
	rexsult, err = lang.RxFromString("5583903.18072100716072011264668631").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("460868914694.088387067451312500723"))
	if err != nil {
		lang.SayString("remx3470")
	} else {
		if !(rexsult.ToString() == "5583903.18072100716072011264668631") {
			lang.SayString("remx3470")
		}
	}
	rexsult, err = lang.RxFromString("-955638397975240685017992436116257E+020").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-508580.148958769104511751975720470E+662"))
	if err != nil {
		lang.SayString("remx3471")
	} else {
		if !(rexsult.ToString() == "-955638397975240685017992436116257E+020") {
			lang.SayString("remx3471")
		}
	}
	rexsult, err = lang.RxFromString("-812.266340554281305985524813074211E+396").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-3195.63111559114001594257448970812E+986"))
	if err != nil {
		lang.SayString("remx3474")
	} else {
		if !(rexsult.ToString() == "-812.266340554281305985524813074211E+396") {
			lang.SayString("remx3474")
		}
	}
	rexsult, err = lang.RxFromString("83946.0157801953636255078996185540").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("492670373.235391665758701500314473"))
	if err != nil {
		lang.SayString("remx3476")
	} else {
		if !(rexsult.ToString() == "83946.0157801953636255078996185540") {
			lang.SayString("remx3476")
		}
	}
	rexsult, err = lang.RxFromString("489191747.148674326757767356626016").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("01136942.1182277580093027768490545"))
	if err != nil {
		lang.SayString("remx3478")
	} else {
		if !(rexsult.ToString() == "306636.3107383827575733115325810") {
			lang.SayString("remx3478")
		}
	}
	rexsult, err = lang.RxFromString("-3376883870.85961692148022521960139").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-65247489449.7334589731171980408284"))
	if err != nil {
		lang.SayString("remx3480")
	} else {
		if !(rexsult.ToString() == "-3376883870.85961692148022521960139") {
			lang.SayString("remx3480")
		}
	}
	rexsult, err = lang.RxFromString("58.6776780370008364590621772421025").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("01.5925518866529044494309229975160"))
	if err != nil {
		lang.SayString("remx3481")
	} else {
		if !(rexsult.ToString() == "1.3458101174962762795489493315265") {
			lang.SayString("remx3481")
		}
	}
	rexsult, err = lang.RxFromString("4099616339.96249499552808575717579").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("290.795187361072489816791525139895"))
	if err != nil {
		lang.SayString("remx3482")
	} else {
		if !(rexsult.ToString() == "37.510275726642959858538282144855") {
			lang.SayString("remx3482")
		}
	}
	rexsult, err = lang.RxFromString("85870777.2282833141709970713739108").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-2140392861153.69401346398478113715"))
	if err != nil {
		lang.SayString("remx3483")
	} else {
		if !(rexsult.ToString() == "85870777.2282833141709970713739108") {
			lang.SayString("remx3483")
		}
	}
	rexsult, err = lang.RxFromString("20900.9693761555165742010339929779").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-38.7546147649523793463260940289585"))
	if err != nil {
		lang.SayString("remx3484")
	} else {
		if !(rexsult.ToString() == "12.2320178461841065312693113692685") {
			lang.SayString("remx3484")
		}
	}
	rexsult, err = lang.RxFromString("448.827596155587910947511170319456").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("379130153.382794042652974596286062"))
	if err != nil {
		lang.SayString("remx3485")
	} else {
		if !(rexsult.ToString() == "448.827596155587910947511170319456") {
			lang.SayString("remx3485")
		}
	}
	rexsult, err = lang.RxFromString("98.4134807921002817357000140482039").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("3404725543.77032945444654351546779"))
	if err != nil {
		lang.SayString("remx3486")
	} else {
		if !(rexsult.ToString() == "98.4134807921002817357000140482039") {
			lang.SayString("remx3486")
		}
	}
	rexsult, err = lang.RxFromString("545746433.649359734136476718176330E-787").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-5149957099709.12830072802043560650E-437"))
	if err != nil {
		lang.SayString("remx3487")
	} else {
		if !(rexsult.ToString() == "545746433.649359734136476718176330E-787") {
			lang.SayString("remx3487")
		}
	}
	rexsult, err = lang.RxFromString("-706.145005094292315613907254240553").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("4739.82486195739758308735946332234"))
	if err != nil {
		lang.SayString("remx3489")
	} else {
		if !(rexsult.ToString() == "-706.145005094292315613907254240553") {
			lang.SayString("remx3489")
		}
	}
	rexsult, err = lang.RxFromString("-769925786.823099083228795187975893").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-31201.9980469760239870067820594790"))
	if err != nil {
		lang.SayString("remx3490")
	} else {
		if !(rexsult.ToString() == "-16485.0139656913494028406582486750") {
			lang.SayString("remx3490")
		}
	}
	rexsult, err = lang.RxFromString("549760.911304725795164589619286514").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("165.160089615604924207754883953484"))
	if err != nil {
		lang.SayString("remx3492")
	} else {
		if !(rexsult.ToString() == "108.133063992607401181365489319248") {
			lang.SayString("remx3492")
		}
	}
	rexsult, err = lang.RxFromString("3650514.18649737956855828939662794").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("08086721.4036886948248274834735629"))
	if err != nil {
		lang.SayString("remx3493")
	} else {
		if !(rexsult.ToString() == "3650514.18649737956855828939662794") {
			lang.SayString("remx3493")
		}
	}
	rexsult, err = lang.RxFromString("55067723881950.1346958179604099594").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("-8.90481481687182931431054785192083"))
	if err != nil {
		lang.SayString("remx3494")
	} else {
		if !(rexsult.ToString() == "1.76788075918488693086347720461547") {
			lang.SayString("remx3494")
		}
	}
	rexsult, err = lang.RxFromString("868251123.413992653362860637541060E+019").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("5579665045.37858308541154858567656E+131"))
	if err != nil {
		lang.SayString("remx3495")
	} else {
		if !(rexsult.ToString() == "868251123.413992653362860637541060E+019") {
			lang.SayString("remx3495")
		}
	}
	rexsult, err = lang.RxFromString("-253280724.939458021588167965038184").OpRem(lang.RxSetWithDigit(33), lang.RxFromString("616988.426425908872398170896375634E+396"))
	if err != nil {
		lang.SayString("remx3500")
	} else {
		if !(rexsult.ToString() == "-253280724.939458021588167965038184") {
			lang.SayString("remx3500")
		}
	}
	rexsult, err = lang.RxFromString("3915134.7").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-597164907."))
	if err != nil {
		lang.SayString("xrem002")
	} else {
		if !(rexsult.ToString() == "3915134.7") {
			lang.SayString("xrem002")
		}
	}
	rexsult, err = lang.RxFromString("309759261").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("62663.487"))
	if err != nil {
		lang.SayString("xrem003")
	} else {
		if !(rexsult.ToString() == "13644.759") {
			lang.SayString("xrem003")
		}
	}
	rexsult, err = lang.RxFromString("3.93591888E-236595626").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7242375.00"))
	if err != nil {
		lang.SayString("xrem004")
	} else {
		if !(rexsult.ToString() == "3.93591888E-236595626") {
			lang.SayString("xrem004")
		}
	}
	rexsult, err = lang.RxFromString("5.11970092").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-8807.22036"))
	if err != nil {
		lang.SayString("xrem006")
	} else {
		if !(rexsult.ToString() == "5.11970092") {
			lang.SayString("xrem006")
		}
	}
	rexsult, err = lang.RxFromString("-7.99874516").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4561.83758"))
	if err != nil {
		lang.SayString("xrem007")
	} else {
		if !(rexsult.ToString() == "-7.99874516") {
			lang.SayString("xrem007")
		}
	}
	rexsult, err = lang.RxFromString("297802878").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-927206.324"))
	if err != nil {
		lang.SayString("xrem008")
	} else {
		if !(rexsult.ToString() == "169647.996") {
			lang.SayString("xrem008")
		}
	}
	rexsult, err = lang.RxFromString("-766.651824").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("31300.3619"))
	if err != nil {
		lang.SayString("xrem009")
	} else {
		if !(rexsult.ToString() == "-766.651824") {
			lang.SayString("xrem009")
		}
	}
	rexsult, err = lang.RxFromString("456417160").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-41346.1024"))
	if err != nil {
		lang.SayString("xrem011")
	} else {
		if !(rexsult.ToString() == "38881.7088") {
			lang.SayString("xrem011")
		}
	}
	rexsult, err = lang.RxFromString("102895.722").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-2.62214826"))
	if err != nil {
		lang.SayString("xrem012")
	} else {
		if !(rexsult.ToString() == "0.00212934") {
			lang.SayString("xrem012")
		}
	}
	rexsult, err = lang.RxFromString("-654645.954").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-9.12535752"))
	if err != nil {
		lang.SayString("xrem015")
	} else {
		if !(rexsult.ToString() == "-1.93087272") {
			lang.SayString("xrem015")
		}
	}
	rexsult, err = lang.RxFromString("63.1917772E-706014634").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-7.56253257E-138579234"))
	if err != nil {
		lang.SayString("xrem016")
	} else {
		if !(rexsult.ToString() == "63.1917772E-706014634") {
			lang.SayString("xrem016")
		}
	}
	rexsult, err = lang.RxFromString("-39674.7190").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2490607.78"))
	if err != nil {
		lang.SayString("xrem017")
	} else {
		if !(rexsult.ToString() == "-39674.7190") {
			lang.SayString("xrem017")
		}
	}
	rexsult, err = lang.RxFromString("-3364.59737E-600363681").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("896487.451"))
	if err != nil {
		lang.SayString("xrem018")
	} else {
		if !(rexsult.ToString() == "-3364.59737E-600363681") {
			lang.SayString("xrem018")
		}
	}
	rexsult, err = lang.RxFromString("-64138.0578").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("31759011.3E+697488342"))
	if err != nil {
		lang.SayString("xrem019")
	} else {
		if !(rexsult.ToString() == "-64138.0578") {
			lang.SayString("xrem019")
		}
	}
	rexsult, err = lang.RxFromString("61399.8527").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-64344484.5"))
	if err != nil {
		lang.SayString("xrem020")
	} else {
		if !(rexsult.ToString() == "61399.8527") {
			lang.SayString("xrem020")
		}
	}
	rexsult, err = lang.RxFromString("-722960.204").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-26154599.8"))
	if err != nil {
		lang.SayString("xrem021")
	} else {
		if !(rexsult.ToString() == "-722960.204") {
			lang.SayString("xrem021")
		}
	}
	rexsult, err = lang.RxFromString("43.7456245").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("547441956."))
	if err != nil {
		lang.SayString("xrem023")
	} else {
		if !(rexsult.ToString() == "43.7456245") {
			lang.SayString("xrem023")
		}
	}
	rexsult, err = lang.RxFromString("-73150542E-242017390").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-8.15869954"))
	if err != nil {
		lang.SayString("xrem024")
	} else {
		if !(rexsult.ToString() == "-73150542E-242017390") {
			lang.SayString("xrem024")
		}
	}
	rexsult, err = lang.RxFromString("29.498114").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-26486451"))
	if err != nil {
		lang.SayString("xrem026")
	} else {
		if !(rexsult.ToString() == "29.498114") {
			lang.SayString("xrem026")
		}
	}
	rexsult, err = lang.RxFromString("-349388.759").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-196215.776"))
	if err != nil {
		lang.SayString("xrem028")
	} else {
		if !(rexsult.ToString() == "-153172.983") {
			lang.SayString("xrem028")
		}
	}
	rexsult, err = lang.RxFromString("-70905112.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-91353968.8"))
	if err != nil {
		lang.SayString("xrem029")
	} else {
		if !(rexsult.ToString() == "-70905112.4") {
			lang.SayString("xrem029")
		}
	}
	rexsult, err = lang.RxFromString("-225094.28").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-88.7723542"))
	if err != nil {
		lang.SayString("xrem030")
	} else {
		if !(rexsult.ToString() == "-56.3621030") {
			lang.SayString("xrem030")
		}
	}
	rexsult, err = lang.RxFromString("50.4442340").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("82.7952169E+880120759"))
	if err != nil {
		lang.SayString("xrem031")
	} else {
		if !(rexsult.ToString() == "50.4442340") {
			lang.SayString("xrem031")
		}
	}
	rexsult, err = lang.RxFromString("-32311.9037").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8.36379449"))
	if err != nil {
		lang.SayString("xrem032")
	} else {
		if !(rexsult.ToString() == "-2.56558513") {
			lang.SayString("xrem032")
		}
	}
	rexsult, err = lang.RxFromString("849.515993E-878446473").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1039.08778"))
	if err != nil {
		lang.SayString("xrem035")
	} else {
		if !(rexsult.ToString() == "849.515993E-878446473") {
			lang.SayString("xrem035")
		}
	}
	rexsult, err = lang.RxFromString("213361789").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-599.644851"))
	if err != nil {
		lang.SayString("xrem036")
	} else {
		if !(rexsult.ToString() == "355.631137") {
			lang.SayString("xrem036")
		}
	}
	rexsult, err = lang.RxFromString("-795522555.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-298.037702"))
	if err != nil {
		lang.SayString("xrem037")
	} else {
		if !(rexsult.ToString() == "-22.783898") {
			lang.SayString("xrem037")
		}
	}
	rexsult, err = lang.RxFromString("-1.70781105E-848889023").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("36504769.4"))
	if err != nil {
		lang.SayString("xrem039")
	} else {
		if !(rexsult.ToString() == "-1.70781105E-848889023") {
			lang.SayString("xrem039")
		}
	}
	rexsult, err = lang.RxFromString("-5290.54984E-490626676").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("842535254"))
	if err != nil {
		lang.SayString("xrem040")
	} else {
		if !(rexsult.ToString() == "-5290.54984E-490626676") {
			lang.SayString("xrem040")
		}
	}
	rexsult, err = lang.RxFromString("-4629035.31").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-167.884398"))
	if err != nil {
		lang.SayString("xrem042")
	} else {
		if !(rexsult.ToString() == "-126.688344") {
			lang.SayString("xrem042")
		}
	}
	rexsult, err = lang.RxFromString("-66527378.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-706400268."))
	if err != nil {
		lang.SayString("xrem043")
	} else {
		if !(rexsult.ToString() == "-66527378.") {
			lang.SayString("xrem043")
		}
	}
	rexsult, err = lang.RxFromString("-2510497.53").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("372882462."))
	if err != nil {
		lang.SayString("xrem044")
	} else {
		if !(rexsult.ToString() == "-2510497.53") {
			lang.SayString("xrem044")
		}
	}
	rexsult, err = lang.RxFromString("136.255393E+53329961").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-53494.7201E+720058060"))
	if err != nil {
		lang.SayString("xrem045")
	} else {
		if !(rexsult.ToString() == "136.255393E+53329961") {
			lang.SayString("xrem045")
		}
	}
	rexsult, err = lang.RxFromString("-876673.100").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-6150.92266"))
	if err != nil {
		lang.SayString("xrem046")
	} else {
		if !(rexsult.ToString() == "-3242.08228") {
			lang.SayString("xrem046")
		}
	}
	rexsult, err = lang.RxFromString("3.61890453E-985606128").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("30664416."))
	if err != nil {
		lang.SayString("xrem049")
	} else {
		if !(rexsult.ToString() == "3.61890453E-985606128") {
			lang.SayString("xrem049")
		}
	}
	rexsult, err = lang.RxFromString("218699.206").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("556944241."))
	if err != nil {
		lang.SayString("xrem051")
	} else {
		if !(rexsult.ToString() == "218699.206") {
			lang.SayString("xrem051")
		}
	}
	rexsult, err = lang.RxFromString("106211716.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-3456793.74"))
	if err != nil {
		lang.SayString("xrem052")
	} else {
		if !(rexsult.ToString() == "2507903.80") {
			lang.SayString("xrem052")
		}
	}
	rexsult, err = lang.RxFromString("364.99811").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-46222.0505"))
	if err != nil {
		lang.SayString("xrem054")
	} else {
		if !(rexsult.ToString() == "364.99811") {
			lang.SayString("xrem054")
		}
	}
	rexsult, err = lang.RxFromString("-392217576.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-958364096"))
	if err != nil {
		lang.SayString("xrem055")
	} else {
		if !(rexsult.ToString() == "-392217576") {
			lang.SayString("xrem055")
		}
	}
	rexsult, err = lang.RxFromString("169601285").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("714526.639"))
	if err != nil {
		lang.SayString("xrem056")
	} else {
		if !(rexsult.ToString() == "258471.557") {
			lang.SayString("xrem056")
		}
	}
	rexsult, err = lang.RxFromString("-674.094552E+586944319").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("6354.2668E+589657266"))
	if err != nil {
		lang.SayString("xrem057")
	} else {
		if !(rexsult.ToString() == "-674.094552E+586944319") {
			lang.SayString("xrem057")
		}
	}
	rexsult, err = lang.RxFromString("-746.293386").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("927749.647"))
	if err != nil {
		lang.SayString("xrem059")
	} else {
		if !(rexsult.ToString() == "-746.293386") {
			lang.SayString("xrem059")
		}
	}
	rexsult, err = lang.RxFromString("6.64377249").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("79161.1070E+619453776"))
	if err != nil {
		lang.SayString("xrem061")
	} else {
		if !(rexsult.ToString() == "6.64377249") {
			lang.SayString("xrem061")
		}
	}
	rexsult, err = lang.RxFromString("3146.66571E-313373366").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("88.5282010"))
	if err != nil {
		lang.SayString("xrem062")
	} else {
		if !(rexsult.ToString() == "3146.66571E-313373366") {
			lang.SayString("xrem062")
		}
	}
	rexsult, err = lang.RxFromString("6.44693097").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-87195.8711"))
	if err != nil {
		lang.SayString("xrem063")
	} else {
		if !(rexsult.ToString() == "6.44693097") {
			lang.SayString("xrem063")
		}
	}
	rexsult, err = lang.RxFromString("-7701.42814").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("72667.5181"))
	if err != nil {
		lang.SayString("xrem065")
	} else {
		if !(rexsult.ToString() == "-7701.42814") {
			lang.SayString("xrem065")
		}
	}
	rexsult, err = lang.RxFromString("-851.754789").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-582659.149"))
	if err != nil {
		lang.SayString("xrem066")
	} else {
		if !(rexsult.ToString() == "-851.754789") {
			lang.SayString("xrem066")
		}
	}
	rexsult, err = lang.RxFromString("-5.01992943").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7852.16531"))
	if err != nil {
		lang.SayString("xrem067")
	} else {
		if !(rexsult.ToString() == "-5.01992943") {
			lang.SayString("xrem067")
		}
	}
	rexsult, err = lang.RxFromString("-12393257.2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("76803689E+949125770"))
	if err != nil {
		lang.SayString("xrem068")
	} else {
		if !(rexsult.ToString() == "-12393257.2") {
			lang.SayString("xrem068")
		}
	}
	rexsult, err = lang.RxFromString("-296590035").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-481734529"))
	if err != nil {
		lang.SayString("xrem071")
	} else {
		if !(rexsult.ToString() == "-296590035") {
			lang.SayString("xrem071")
		}
	}
	rexsult, err = lang.RxFromString("8.27822605").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9241557.19"))
	if err != nil {
		lang.SayString("xrem072")
	} else {
		if !(rexsult.ToString() == "8.27822605") {
			lang.SayString("xrem072")
		}
	}
	rexsult, err = lang.RxFromString("-1.43581098").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7286313.54"))
	if err != nil {
		lang.SayString("xrem073")
	} else {
		if !(rexsult.ToString() == "-1.43581098") {
			lang.SayString("xrem073")
		}
	}
	rexsult, err = lang.RxFromString("-699036193.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("759263.509E+533543625"))
	if err != nil {
		lang.SayString("xrem074")
	} else {
		if !(rexsult.ToString() == "-699036193.") {
			lang.SayString("xrem074")
		}
	}
	rexsult, err = lang.RxFromString("-83.7273615E-305281957").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-287779593.E+458777774"))
	if err != nil {
		lang.SayString("xrem075")
	} else {
		if !(rexsult.ToString() == "-83.7273615E-305281957") {
			lang.SayString("xrem075")
		}
	}
	rexsult, err = lang.RxFromString("8.48503224").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("6522.03316"))
	if err != nil {
		lang.SayString("xrem076")
	} else {
		if !(rexsult.ToString() == "8.48503224") {
			lang.SayString("xrem076")
		}
	}
	rexsult, err = lang.RxFromString("527916091").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-809.054070"))
	if err != nil {
		lang.SayString("xrem077")
	} else {
		if !(rexsult.ToString() == "219.784300") {
			lang.SayString("xrem077")
		}
	}
	rexsult, err = lang.RxFromString("3857058.60").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("5792997.58E+881077409"))
	if err != nil {
		lang.SayString("xrem078")
	} else {
		if !(rexsult.ToString() == "3857058.60") {
			lang.SayString("xrem078")
		}
	}
	rexsult, err = lang.RxFromString("-580.502955").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("38125521.7"))
	if err != nil {
		lang.SayString("xrem080")
	} else {
		if !(rexsult.ToString() == "-580.502955") {
			lang.SayString("xrem080")
		}
	}
	rexsult, err = lang.RxFromString("-8378.55499").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("760.131257"))
	if err != nil {
		lang.SayString("xrem083")
	} else {
		if !(rexsult.ToString() == "-17.111163") {
			lang.SayString("xrem083")
		}
	}
	rexsult, err = lang.RxFromString("-717.697718").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("984304413"))
	if err != nil {
		lang.SayString("xrem084")
	} else {
		if !(rexsult.ToString() == "-717.697718") {
			lang.SayString("xrem084")
		}
	}
	rexsult, err = lang.RxFromString("-76762243.4E-741100094").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-273.706674"))
	if err != nil {
		lang.SayString("xrem085")
	} else {
		if !(rexsult.ToString() == "-76762243.4E-741100094") {
			lang.SayString("xrem085")
		}
	}
	rexsult, err = lang.RxFromString("-359866845.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4.57434117"))
	if err != nil {
		lang.SayString("xrem087")
	} else {
		if !(rexsult.ToString() == "-3.54890484") {
			lang.SayString("xrem087")
		}
	}
	rexsult, err = lang.RxFromString("779934536.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-76562645.7"))
	if err != nil {
		lang.SayString("xrem088")
	} else {
		if !(rexsult.ToString() == "14308079.0") {
			lang.SayString("xrem088")
		}
	}
	rexsult, err = lang.RxFromString("-4820.95451").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3516234.99E+303303176"))
	if err != nil {
		lang.SayString("xrem089")
	} else {
		if !(rexsult.ToString() == "-4820.95451") {
			lang.SayString("xrem089")
		}
	}
	rexsult, err = lang.RxFromString("69355976.9").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-9.57838562E+758804984"))
	if err != nil {
		lang.SayString("xrem090")
	} else {
		if !(rexsult.ToString() == "69355976.9") {
			lang.SayString("xrem090")
		}
	}
	rexsult, err = lang.RxFromString("-671.507198E-908587890").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3057429.32E-555230623"))
	if err != nil {
		lang.SayString("xrem094")
	} else {
		if !(rexsult.ToString() == "-671.507198E-908587890") {
			lang.SayString("xrem094")
		}
	}
	rexsult, err = lang.RxFromString("329579114").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("146780548."))
	if err != nil {
		lang.SayString("xrem096")
	} else {
		if !(rexsult.ToString() == "36018018") {
			lang.SayString("xrem096")
		}
	}
	rexsult, err = lang.RxFromString("-789904.686E-217225000").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1991.07181E-84080059"))
	if err != nil {
		lang.SayString("xrem097")
	} else {
		if !(rexsult.ToString() == "-789904.686E-217225000") {
			lang.SayString("xrem097")
		}
	}
	rexsult, err = lang.RxFromString("59893.3544").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-408595868"))
	if err != nil {
		lang.SayString("xrem098")
	} else {
		if !(rexsult.ToString() == "59893.3544") {
			lang.SayString("xrem098")
		}
	}
	rexsult, err = lang.RxFromString("9866.99208").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("708756501."))
	if err != nil {
		lang.SayString("xrem100")
	} else {
		if !(rexsult.ToString() == "9866.99208") {
			lang.SayString("xrem100")
		}
	}
	rexsult, err = lang.RxFromString("-78810.6297").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-399884.68"))
	if err != nil {
		lang.SayString("xrem101")
	} else {
		if !(rexsult.ToString() == "-78810.6297") {
			lang.SayString("xrem101")
		}
	}
	rexsult, err = lang.RxFromString("409189761").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-771.471460"))
	if err != nil {
		lang.SayString("xrem102")
	} else {
		if !(rexsult.ToString() == "527.144540") {
			lang.SayString("xrem102")
		}
	}
	rexsult, err = lang.RxFromString("-1.68748838").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("460.46924"))
	if err != nil {
		lang.SayString("xrem103")
	} else {
		if !(rexsult.ToString() == "-1.68748838") {
			lang.SayString("xrem103")
		}
	}
	rexsult, err = lang.RxFromString("553527296.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-7924.40185"))
	if err != nil {
		lang.SayString("xrem104")
	} else {
		if !(rexsult.ToString() == "7826.77750") {
			lang.SayString("xrem104")
		}
	}
	rexsult, err = lang.RxFromString("-38.7465207").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("64936.2942"))
	if err != nil {
		lang.SayString("xrem105")
	} else {
		if !(rexsult.ToString() == "-38.7465207") {
			lang.SayString("xrem105")
		}
	}
	rexsult, err = lang.RxFromString("-201075.248").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("845.663928"))
	if err != nil {
		lang.SayString("xrem106")
	} else {
		if !(rexsult.ToString() == "-652.897064") {
			lang.SayString("xrem106")
		}
	}
	rexsult, err = lang.RxFromString("91048.4559").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("75953609.3"))
	if err != nil {
		lang.SayString("xrem107")
	} else {
		if !(rexsult.ToString() == "91048.4559") {
			lang.SayString("xrem107")
		}
	}
	rexsult, err = lang.RxFromString("6898273.86E-252097460").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("15.3456196"))
	if err != nil {
		lang.SayString("xrem108")
	} else {
		if !(rexsult.ToString() == "6898273.86E-252097460") {
			lang.SayString("xrem108")
		}
	}
	rexsult, err = lang.RxFromString("-17643.39").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.0352568E+304871331"))
	if err != nil {
		lang.SayString("xrem110")
	} else {
		if !(rexsult.ToString() == "-17643.39") {
			lang.SayString("xrem110")
		}
	}
	rexsult, err = lang.RxFromString("4589785.16").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7459.04237"))
	if err != nil {
		lang.SayString("xrem111")
	} else {
		if !(rexsult.ToString() == "2474.10245") {
			lang.SayString("xrem111")
		}
	}
	rexsult, err = lang.RxFromString("-51.1632090E-753968082").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8.96207471E-585797887"))
	if err != nil {
		lang.SayString("xrem112")
	} else {
		if !(rexsult.ToString() == "-51.1632090E-753968082") {
			lang.SayString("xrem112")
		}
	}
	rexsult, err = lang.RxFromString("503712056.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-57490703.5E+924890183"))
	if err != nil {
		lang.SayString("xrem114")
	} else {
		if !(rexsult.ToString() == "503712056.") {
			lang.SayString("xrem114")
		}
	}
	rexsult, err = lang.RxFromString("883.849223").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("249259171"))
	if err != nil {
		lang.SayString("xrem115")
	} else {
		if !(rexsult.ToString() == "883.849223") {
			lang.SayString("xrem115")
		}
	}
	rexsult, err = lang.RxFromString("14239029.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-36527.2221"))
	if err != nil {
		lang.SayString("xrem120")
	} else {
		if !(rexsult.ToString() == "29939.6031") {
			lang.SayString("xrem120")
		}
	}
	rexsult, err = lang.RxFromString("72333.2654E-544425548").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-690.664836E+662155120"))
	if err != nil {
		lang.SayString("xrem121")
	} else {
		if !(rexsult.ToString() == "72333.2654E-544425548") {
			lang.SayString("xrem121")
		}
	}
	rexsult, err = lang.RxFromString("-37721.1567E-115787341").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-828949864E-76251747"))
	if err != nil {
		lang.SayString("xrem122")
	} else {
		if !(rexsult.ToString() == "-37721.1567E-115787341") {
			lang.SayString("xrem122")
		}
	}
	rexsult, err = lang.RxFromString("-2078852.83E-647080089").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-119779858.E+734665461"))
	if err != nil {
		lang.SayString("xrem123")
	} else {
		if !(rexsult.ToString() == "-2078852.83E-647080089") {
			lang.SayString("xrem123")
		}
	}
	rexsult, err = lang.RxFromString("-79145.3625").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-7718.57307"))
	if err != nil {
		lang.SayString("xrem124")
	} else {
		if !(rexsult.ToString() == "-1959.63180") {
			lang.SayString("xrem124")
		}
	}
	rexsult, err = lang.RxFromString("911249557").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("79810804.9"))
	if err != nil {
		lang.SayString("xrem126")
	} else {
		if !(rexsult.ToString() == "33330703.1") {
			lang.SayString("xrem126")
		}
	}
	rexsult, err = lang.RxFromString("341134.994").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3.37486292"))
	if err != nil {
		lang.SayString("xrem127")
	} else {
		if !(rexsult.ToString() == "0.47518348") {
			lang.SayString("xrem127")
		}
	}
	rexsult, err = lang.RxFromString("699631.893").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-226.423958"))
	if err != nil {
		lang.SayString("xrem130")
	} else {
		if !(rexsult.ToString() == "208.286738") {
			lang.SayString("xrem130")
		}
	}
	rexsult, err = lang.RxFromString("-249350139.E-571793673").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("775732428."))
	if err != nil {
		lang.SayString("xrem131")
	} else {
		if !(rexsult.ToString() == "-249350139.E-571793673") {
			lang.SayString("xrem131")
		}
	}
	rexsult, err = lang.RxFromString("5.11629020").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-480.53194"))
	if err != nil {
		lang.SayString("xrem132")
	} else {
		if !(rexsult.ToString() == "5.11629020") {
			lang.SayString("xrem132")
		}
	}
	rexsult, err = lang.RxFromString("-8.23352673E-446723147").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-530710.866"))
	if err != nil {
		lang.SayString("xrem133")
	} else {
		if !(rexsult.ToString() == "-8.23352673E-446723147") {
			lang.SayString("xrem133")
		}
	}
	rexsult, err = lang.RxFromString("7.0598608").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-95908.35"))
	if err != nil {
		lang.SayString("xrem134")
	} else {
		if !(rexsult.ToString() == "7.0598608") {
			lang.SayString("xrem134")
		}
	}
	rexsult, err = lang.RxFromString("-7.91189845E+207202706").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1532.71847E+509944335"))
	if err != nil {
		lang.SayString("xrem135")
	} else {
		if !(rexsult.ToString() == "-7.91189845E+207202706") {
			lang.SayString("xrem135")
		}
	}
	rexsult, err = lang.RxFromString("208839370.E-215147432").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-75.9420559"))
	if err != nil {
		lang.SayString("xrem136")
	} else {
		if !(rexsult.ToString() == "208839370.E-215147432") {
			lang.SayString("xrem136")
		}
	}
	rexsult, err = lang.RxFromString("427.754244E-353328369").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4705.0796"))
	if err != nil {
		lang.SayString("xrem137")
	} else {
		if !(rexsult.ToString() == "427.754244E-353328369") {
			lang.SayString("xrem137")
		}
	}
	rexsult, err = lang.RxFromString("452371821.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4109709.19"))
	if err != nil {
		lang.SayString("xrem139")
	} else {
		if !(rexsult.ToString() == "303810.10") {
			lang.SayString("xrem139")
		}
	}
	rexsult, err = lang.RxFromString("94007.4392").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-9467725.5E+681898234"))
	if err != nil {
		lang.SayString("xrem140")
	} else {
		if !(rexsult.ToString() == "94007.4392") {
			lang.SayString("xrem140")
		}
	}
	rexsult, err = lang.RxFromString("99147554.0E-751410586").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("38313.6423"))
	if err != nil {
		lang.SayString("xrem141")
	} else {
		if !(rexsult.ToString() == "99147554.0E-751410586") {
			lang.SayString("xrem141")
		}
	}
	rexsult, err = lang.RxFromString("-7919.30254").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-669.607854"))
	if err != nil {
		lang.SayString("xrem142")
	} else {
		if !(rexsult.ToString() == "-553.616146") {
			lang.SayString("xrem142")
		}
	}
	rexsult, err = lang.RxFromString("3455755.47E-112465506").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("771.674306"))
	if err != nil {
		lang.SayString("xrem144")
	} else {
		if !(rexsult.ToString() == "3455755.47E-112465506") {
			lang.SayString("xrem144")
		}
	}
	rexsult, err = lang.RxFromString("-477067757.E-961684940").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7.70122608E-741072245"))
	if err != nil {
		lang.SayString("xrem145")
	} else {
		if !(rexsult.ToString() == "-477067757.E-961684940") {
			lang.SayString("xrem145")
		}
	}
	rexsult, err = lang.RxFromString("76482.352").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8237806.8"))
	if err != nil {
		lang.SayString("xrem146")
	} else {
		if !(rexsult.ToString() == "76482.352") {
			lang.SayString("xrem146")
		}
	}
	rexsult, err = lang.RxFromString("1.21505164E-565556504").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9.26146573"))
	if err != nil {
		lang.SayString("xrem147")
	} else {
		if !(rexsult.ToString() == "1.21505164E-565556504") {
			lang.SayString("xrem147")
		}
	}
	rexsult, err = lang.RxFromString("-8303060.25E-169894883").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("901561.985"))
	if err != nil {
		lang.SayString("xrem148")
	} else {
		if !(rexsult.ToString() == "-8303060.25E-169894883") {
			lang.SayString("xrem148")
		}
	}
	rexsult, err = lang.RxFromString("-592464.92").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("71445510.7"))
	if err != nil {
		lang.SayString("xrem149")
	} else {
		if !(rexsult.ToString() == "-592464.92") {
			lang.SayString("xrem149")
		}
	}
	rexsult, err = lang.RxFromString("-73774.4165").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-39.8243027"))
	if err != nil {
		lang.SayString("xrem150")
	} else {
		if !(rexsult.ToString() == "-19.8078996") {
			lang.SayString("xrem150")
		}
	}
	rexsult, err = lang.RxFromString("-524724715.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-55763.7937"))
	if err != nil {
		lang.SayString("xrem151")
	} else {
		if !(rexsult.ToString() == "-43180.0767") {
			lang.SayString("xrem151")
		}
	}
	rexsult, err = lang.RxFromString("37.6027452").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7.22454233"))
	if err != nil {
		lang.SayString("xrem153")
	} else {
		if !(rexsult.ToString() == "1.48003355") {
			lang.SayString("xrem153")
		}
	}
	rexsult, err = lang.RxFromString("2447660.39").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-36981.4253"))
	if err != nil {
		lang.SayString("xrem154")
	} else {
		if !(rexsult.ToString() == "6886.3202") {
			lang.SayString("xrem154")
		}
	}
	rexsult, err = lang.RxFromString("2160.36419").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1418.33574E+656265382"))
	if err != nil {
		lang.SayString("xrem155")
	} else {
		if !(rexsult.ToString() == "2160.36419") {
			lang.SayString("xrem155")
		}
	}
	rexsult, err = lang.RxFromString("8926.44939").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("54.9430027"))
	if err != nil {
		lang.SayString("xrem156")
	} else {
		if !(rexsult.ToString() == "25.6829526") {
			lang.SayString("xrem156")
		}
	}
	rexsult, err = lang.RxFromString("861588029").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-41657398E+77955925"))
	if err != nil {
		lang.SayString("xrem157")
	} else {
		if !(rexsult.ToString() == "861588029") {
			lang.SayString("xrem157")
		}
	}
	rexsult, err = lang.RxFromString("-34.5253062").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("52.6722019"))
	if err != nil {
		lang.SayString("xrem158")
	} else {
		if !(rexsult.ToString() == "-34.5253062") {
			lang.SayString("xrem158")
		}
	}
	rexsult, err = lang.RxFromString("-18861647.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("99794586.7"))
	if err != nil {
		lang.SayString("xrem159")
	} else {
		if !(rexsult.ToString() == "-18861647") {
			lang.SayString("xrem159")
		}
	}
	rexsult, err = lang.RxFromString("322192.407").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("461.67044"))
	if err != nil {
		lang.SayString("xrem160")
	} else {
		if !(rexsult.ToString() == "408.11032") {
			lang.SayString("xrem160")
		}
	}
	rexsult, err = lang.RxFromString("293.773732").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("479899052E+789950177"))
	if err != nil {
		lang.SayString("xrem162")
	} else {
		if !(rexsult.ToString() == "293.773732") {
			lang.SayString("xrem162")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("51897955.3"))
	if err != nil {
		lang.SayString("xrem163")
	} else {
		if !(rexsult.ToString() == "-51621406.7") {
			lang.SayString("xrem163")
		}
	}
	rexsult, err = lang.RxFromString("37380.7802").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-277719788."))
	if err != nil {
		lang.SayString("xrem164")
	} else {
		if !(rexsult.ToString() == "37380.7802") {
			lang.SayString("xrem164")
		}
	}
	rexsult, err = lang.RxFromString("320133844.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-977517477"))
	if err != nil {
		lang.SayString("xrem165")
	} else {
		if !(rexsult.ToString() == "320133844") {
			lang.SayString("xrem165")
		}
	}
	rexsult, err = lang.RxFromString("-5409.00482").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-2.16149386"))
	if err != nil {
		lang.SayString("xrem167")
	} else {
		if !(rexsult.ToString() == "-0.94718228") {
			lang.SayString("xrem167")
		}
	}
	rexsult, err = lang.RxFromString("-957960.367").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("322.858170"))
	if err != nil {
		lang.SayString("xrem168")
	} else {
		if !(rexsult.ToString() == "-40.176610") {
			lang.SayString("xrem168")
		}
	}
	rexsult, err = lang.RxFromString("-11817.8754E+613893442").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-3.84735082E+888333249"))
	if err != nil {
		lang.SayString("xrem169")
	} else {
		if !(rexsult.ToString() == "-11817.8754E+613893442") {
			lang.SayString("xrem169")
		}
	}
	rexsult, err = lang.RxFromString("-205842096.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-191342.721"))
	if err != nil {
		lang.SayString("xrem171")
	} else {
		if !(rexsult.ToString() == "-148670.925") {
			lang.SayString("xrem171")
		}
	}
	rexsult, err = lang.RxFromString("42501124.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("884.938498E+123341480"))
	if err != nil {
		lang.SayString("xrem172")
	} else {
		if !(rexsult.ToString() == "42501124.") {
			lang.SayString("xrem172")
		}
	}
	rexsult, err = lang.RxFromString("-57809452.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-620380746"))
	if err != nil {
		lang.SayString("xrem173")
	} else {
		if !(rexsult.ToString() == "-57809452.") {
			lang.SayString("xrem173")
		}
	}
	rexsult, err = lang.RxFromString("-8022370.31").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9858581.6"))
	if err != nil {
		lang.SayString("xrem174")
	} else {
		if !(rexsult.ToString() == "-8022370.31") {
			lang.SayString("xrem174")
		}
	}
	rexsult, err = lang.RxFromString("-697273715E-242824870").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-3.81757506"))
	if err != nil {
		lang.SayString("xrem176")
	} else {
		if !(rexsult.ToString() == "-697273715E-242824870") {
			lang.SayString("xrem176")
		}
	}
	rexsult, err = lang.RxFromString("-7.42204403E-315716280").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-8156111.67E+283261636"))
	if err != nil {
		lang.SayString("xrem177")
	} else {
		if !(rexsult.ToString() == "-7.42204403E-315716280") {
			lang.SayString("xrem177")
		}
	}
	rexsult, err = lang.RxFromString("738063892").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("89900467.8"))
	if err != nil {
		lang.SayString("xrem178")
	} else {
		if !(rexsult.ToString() == "18860149.6") {
			lang.SayString("xrem178")
		}
	}
	rexsult, err = lang.RxFromString("613.207774").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-3007.78608"))
	if err != nil {
		lang.SayString("xrem180")
	} else {
		if !(rexsult.ToString() == "613.207774") {
			lang.SayString("xrem180")
		}
	}
	rexsult, err = lang.RxFromString("-93006222.3").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-3.08964619"))
	if err != nil {
		lang.SayString("xrem181")
	} else {
		if !(rexsult.ToString() == "-2.65215407") {
			lang.SayString("xrem181")
		}
	}
	rexsult, err = lang.RxFromString("19272386.9").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-410442379."))
	if err != nil {
		lang.SayString("xrem183")
	} else {
		if !(rexsult.ToString() == "19272386.9") {
			lang.SayString("xrem183")
		}
	}
	rexsult, err = lang.RxFromString("571.536725").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("389.899220"))
	if err != nil {
		lang.SayString("xrem185")
	} else {
		if !(rexsult.ToString() == "181.637505") {
			lang.SayString("xrem185")
		}
	}
	rexsult, err = lang.RxFromString("92427442.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("674334898."))
	if err != nil {
		lang.SayString("xrem188")
	} else {
		if !(rexsult.ToString() == "92427442.4") {
			lang.SayString("xrem188")
		}
	}
	rexsult, err = lang.RxFromString("44651895.7").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-910508.438"))
	if err != nil {
		lang.SayString("xrem189")
	} else {
		if !(rexsult.ToString() == "36982.238") {
			lang.SayString("xrem189")
		}
	}
	rexsult, err = lang.RxFromString("25.2592149").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("59.0436981"))
	if err != nil {
		lang.SayString("xrem191")
	} else {
		if !(rexsult.ToString() == "25.2592149") {
			lang.SayString("xrem191")
		}
	}
	rexsult, err = lang.RxFromString("-6.850835").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1273.48240"))
	if err != nil {
		lang.SayString("xrem192")
	} else {
		if !(rexsult.ToString() == "-6.850835") {
			lang.SayString("xrem192")
		}
	}
	rexsult, err = lang.RxFromString("174.272325").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("5638.16229"))
	if err != nil {
		lang.SayString("xrem193")
	} else {
		if !(rexsult.ToString() == "174.272325") {
			lang.SayString("xrem193")
		}
	}
	rexsult, err = lang.RxFromString("3455629.76").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-8.27332322"))
	if err != nil {
		lang.SayString("xrem194")
	} else {
		if !(rexsult.ToString() == "3.29750074") {
			lang.SayString("xrem194")
		}
	}
	rexsult, err = lang.RxFromString("-924337723E-640771235").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("86639377.1"))
	if err != nil {
		lang.SayString("xrem195")
	} else {
		if !(rexsult.ToString() == "-924337723E-640771235") {
			lang.SayString("xrem195")
		}
	}
	rexsult, err = lang.RxFromString("-18857539.9").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("813013129."))
	if err != nil {
		lang.SayString("xrem198")
	} else {
		if !(rexsult.ToString() == "-18857539.9") {
			lang.SayString("xrem198")
		}
	}
	rexsult, err = lang.RxFromString("-8.29530327").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3243419.57E+35688332"))
	if err != nil {
		lang.SayString("xrem199")
	} else {
		if !(rexsult.ToString() == "-8.29530327") {
			lang.SayString("xrem199")
		}
	}
	rexsult, err = lang.RxFromString("-57101683.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("763551341E+991491712"))
	if err != nil {
		lang.SayString("xrem200")
	} else {
		if !(rexsult.ToString() == "-57101683.5") {
			lang.SayString("xrem200")
		}
	}
	rexsult, err = lang.RxFromString("-603326.740").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1710.95183"))
	if err != nil {
		lang.SayString("xrem201")
	} else {
		if !(rexsult.ToString() == "-1071.69584") {
			lang.SayString("xrem201")
		}
	}
	rexsult, err = lang.RxFromString("-48142763.3").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-943434114"))
	if err != nil {
		lang.SayString("xrem202")
	} else {
		if !(rexsult.ToString() == "-48142763.3") {
			lang.SayString("xrem202")
		}
	}
	rexsult, err = lang.RxFromString("-204.586767").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-235.531847"))
	if err != nil {
		lang.SayString("xrem203")
	} else {
		if !(rexsult.ToString() == "-204.586767") {
			lang.SayString("xrem203")
		}
	}
	rexsult, err = lang.RxFromString("-70.3805581").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("830137.913"))
	if err != nil {
		lang.SayString("xrem204")
	} else {
		if !(rexsult.ToString() == "-70.3805581") {
			lang.SayString("xrem204")
		}
	}
	rexsult, err = lang.RxFromString("-8818.47606").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-60766.4571"))
	if err != nil {
		lang.SayString("xrem205")
	} else {
		if !(rexsult.ToString() == "-8818.47606") {
			lang.SayString("xrem205")
		}
	}
	rexsult, err = lang.RxFromString("37060929.3E-168439509").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-79576717.1"))
	if err != nil {
		lang.SayString("xrem206")
	} else {
		if !(rexsult.ToString() == "37060929.3E-168439509") {
			lang.SayString("xrem206")
		}
	}
	rexsult, err = lang.RxFromString("-656285310.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-107221462."))
	if err != nil {
		lang.SayString("xrem207")
	} else {
		if !(rexsult.ToString() == "-12956538") {
			lang.SayString("xrem207")
		}
	}
	rexsult, err = lang.RxFromString("653397.125").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7195.30990"))
	if err != nil {
		lang.SayString("xrem208")
	} else {
		if !(rexsult.ToString() == "5819.23400") {
			lang.SayString("xrem208")
		}
	}
	rexsult, err = lang.RxFromString("-62011.4563E-117563240").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-57.1731586E+115657204"))
	if err != nil {
		lang.SayString("xrem211")
	} else {
		if !(rexsult.ToString() == "-62011.4563E-117563240") {
			lang.SayString("xrem211")
		}
	}
	rexsult, err = lang.RxFromString("739.944710").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("202949.175"))
	if err != nil {
		lang.SayString("xrem213")
	} else {
		if !(rexsult.ToString() == "739.944710") {
			lang.SayString("xrem213")
		}
	}
	rexsult, err = lang.RxFromString("87686.8016").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4204890.40"))
	if err != nil {
		lang.SayString("xrem214")
	} else {
		if !(rexsult.ToString() == "87686.8016") {
			lang.SayString("xrem214")
		}
	}
	rexsult, err = lang.RxFromString("987126721.E-725794834").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4874166.23"))
	if err != nil {
		lang.SayString("xrem215")
	} else {
		if !(rexsult.ToString() == "987126721.E-725794834") {
			lang.SayString("xrem215")
		}
	}
	rexsult, err = lang.RxFromString("728148726.E-661695938").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("32798.5202"))
	if err != nil {
		lang.SayString("xrem216")
	} else {
		if !(rexsult.ToString() == "728148726.E-661695938") {
			lang.SayString("xrem216")
		}
	}
	rexsult, err = lang.RxFromString("7428219.97").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("667.326760"))
	if err != nil {
		lang.SayString("xrem217")
	} else {
		if !(rexsult.ToString() == "205.804440") {
			lang.SayString("xrem217")
		}
	}
	rexsult, err = lang.RxFromString("-358.24550").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4447.78675E+601402509"))
	if err != nil {
		lang.SayString("xrem219")
	} else {
		if !(rexsult.ToString() == "-358.24550") {
			lang.SayString("xrem219")
		}
	}
	rexsult, err = lang.RxFromString("118.621826").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-2.72010038"))
	if err != nil {
		lang.SayString("xrem220")
	} else {
		if !(rexsult.ToString() == "1.65750966") {
			lang.SayString("xrem220")
		}
	}
	rexsult, err = lang.RxFromString("-35544.4029").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-567830.130"))
	if err != nil {
		lang.SayString("xrem223")
	} else {
		if !(rexsult.ToString() == "-35544.4029") {
			lang.SayString("xrem223")
		}
	}
	rexsult, err = lang.RxFromString("-509.483395").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-147242915."))
	if err != nil {
		lang.SayString("xrem225")
	} else {
		if !(rexsult.ToString() == "-509.483395") {
			lang.SayString("xrem225")
		}
	}
	rexsult, err = lang.RxFromString("895612630.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-36.4104040"))
	if err != nil {
		lang.SayString("xrem227")
	} else {
		if !(rexsult.ToString() == "35.0147560") {
			lang.SayString("xrem227")
		}
	}
	rexsult, err = lang.RxFromString("25455.4973").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2955.00006E+528196218"))
	if err != nil {
		lang.SayString("xrem228")
	} else {
		if !(rexsult.ToString() == "25455.4973") {
			lang.SayString("xrem228")
		}
	}
	rexsult, err = lang.RxFromString("62871.2202").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2484.0382E+211662557"))
	if err != nil {
		lang.SayString("xrem230")
	} else {
		if !(rexsult.ToString() == "62871.2202") {
			lang.SayString("xrem230")
		}
	}
	rexsult, err = lang.RxFromString("71.9281575").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-9810012.5"))
	if err != nil {
		lang.SayString("xrem231")
	} else {
		if !(rexsult.ToString() == "71.9281575") {
			lang.SayString("xrem231")
		}
	}
	rexsult, err = lang.RxFromString("-6388022.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-88.042967"))
	if err != nil {
		lang.SayString("xrem232")
	} else {
		if !(rexsult.ToString() == "-64.529315") {
			lang.SayString("xrem232")
		}
	}
	rexsult, err = lang.RxFromString("372567445.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("96.0992141"))
	if err != nil {
		lang.SayString("xrem233")
	} else {
		if !(rexsult.ToString() == "17.4588536") {
			lang.SayString("xrem233")
		}
	}
	rexsult, err = lang.RxFromString("-3.65207541").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("74501982.0"))
	if err != nil {
		lang.SayString("xrem235")
	} else {
		if !(rexsult.ToString() == "-3.65207541") {
			lang.SayString("xrem235")
		}
	}
	rexsult, err = lang.RxFromString("-5297.76981").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-859.719404"))
	if err != nil {
		lang.SayString("xrem236")
	} else {
		if !(rexsult.ToString() == "-139.453386") {
			lang.SayString("xrem236")
		}
	}
	rexsult, err = lang.RxFromString("-684172.592").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("766.448597E+288361959"))
	if err != nil {
		lang.SayString("xrem237")
	} else {
		if !(rexsult.ToString() == "-684172.592") {
			lang.SayString("xrem237")
		}
	}
	rexsult, err = lang.RxFromString("626919.219").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("57469.8727E+13188610"))
	if err != nil {
		lang.SayString("xrem238")
	} else {
		if !(rexsult.ToString() == "626919.219") {
			lang.SayString("xrem238")
		}
	}
	rexsult, err = lang.RxFromString("-77480.5840").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("893265.594E+287982552"))
	if err != nil {
		lang.SayString("xrem239")
	} else {
		if !(rexsult.ToString() == "-77480.5840") {
			lang.SayString("xrem239")
		}
	}
	rexsult, err = lang.RxFromString("-7177620.29").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7786343.83"))
	if err != nil {
		lang.SayString("xrem240")
	} else {
		if !(rexsult.ToString() == "-7177620.29") {
			lang.SayString("xrem240")
		}
	}
	rexsult, err = lang.RxFromString("9.6224130").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4.50355112"))
	if err != nil {
		lang.SayString("xrem241")
	} else {
		if !(rexsult.ToString() == "0.61531076") {
			lang.SayString("xrem241")
		}
	}
	rexsult, err = lang.RxFromString("-66.6337347E-597410086").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-818812885"))
	if err != nil {
		lang.SayString("xrem242")
	} else {
		if !(rexsult.ToString() == "-66.6337347E-597410086") {
			lang.SayString("xrem242")
		}
	}
	rexsult, err = lang.RxFromString("65587553.7").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("600574.736"))
	if err != nil {
		lang.SayString("xrem243")
	} else {
		if !(rexsult.ToString() == "124907.476") {
			lang.SayString("xrem243")
		}
	}
	rexsult, err = lang.RxFromString("-32401.939").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-585200217."))
	if err != nil {
		lang.SayString("xrem244")
	} else {
		if !(rexsult.ToString() == "-32401.939") {
			lang.SayString("xrem244")
		}
	}
	rexsult, err = lang.RxFromString("69573.988").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-9.77003465E+740933668"))
	if err != nil {
		lang.SayString("xrem245")
	} else {
		if !(rexsult.ToString() == "69573.988") {
			lang.SayString("xrem245")
		}
	}
	rexsult, err = lang.RxFromString("216741082.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("250290244"))
	if err != nil {
		lang.SayString("xrem248")
	} else {
		if !(rexsult.ToString() == "216741082") {
			lang.SayString("xrem248")
		}
	}
	rexsult, err = lang.RxFromString("-6364720.49").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("5539245.64"))
	if err != nil {
		lang.SayString("xrem249")
	} else {
		if !(rexsult.ToString() == "-825474.85") {
			lang.SayString("xrem249")
		}
	}
	rexsult, err = lang.RxFromString("-814599.475").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-14.5431191"))
	if err != nil {
		lang.SayString("xrem250")
	} else {
		if !(rexsult.ToString() == "-10.2879708") {
			lang.SayString("xrem250")
		}
	}
	rexsult, err = lang.RxFromString("-162726.257E-597285918").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4391.54799"))
	if err != nil {
		lang.SayString("xrem253")
	} else {
		if !(rexsult.ToString() == "-162726.257E-597285918") {
			lang.SayString("xrem253")
		}
	}
	rexsult, err = lang.RxFromString("700354586.E-99856707").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7198.0493E+436250299"))
	if err != nil {
		lang.SayString("xrem254")
	} else {
		if !(rexsult.ToString() == "700354586.E-99856707") {
			lang.SayString("xrem254")
		}
	}
	rexsult, err = lang.RxFromString("39617663E-463704664").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-895.290346"))
	if err != nil {
		lang.SayString("xrem255")
	} else {
		if !(rexsult.ToString() == "39617663E-463704664") {
			lang.SayString("xrem255")
		}
	}
	rexsult, err = lang.RxFromString("5350882.59").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-36329829"))
	if err != nil {
		lang.SayString("xrem256")
	} else {
		if !(rexsult.ToString() == "5350882.59") {
			lang.SayString("xrem256")
		}
	}
	rexsult, err = lang.RxFromString("231899031.E-481759076").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("726.337100"))
	if err != nil {
		lang.SayString("xrem258")
	} else {
		if !(rexsult.ToString() == "231899031.E-481759076") {
			lang.SayString("xrem258")
		}
	}
	rexsult, err = lang.RxFromString("-9611312.33").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("22109735.9"))
	if err != nil {
		lang.SayString("xrem259")
	} else {
		if !(rexsult.ToString() == "-9611312.33") {
			lang.SayString("xrem259")
		}
	}
	rexsult, err = lang.RxFromString("-5604938.15E-36812542").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("735937577."))
	if err != nil {
		lang.SayString("xrem260")
	} else {
		if !(rexsult.ToString() == "-5604938.15E-36812542") {
			lang.SayString("xrem260")
		}
	}
	rexsult, err = lang.RxFromString("-34865.7378E-368768024").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2297117.88"))
	if err != nil {
		lang.SayString("xrem262")
	} else {
		if !(rexsult.ToString() == "-34865.7378E-368768024") {
			lang.SayString("xrem262")
		}
	}
	rexsult, err = lang.RxFromString("1123.32456").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7.86747918E+930888796"))
	if err != nil {
		lang.SayString("xrem263")
	} else {
		if !(rexsult.ToString() == "1123.32456") {
			lang.SayString("xrem263")
		}
	}
	rexsult, err = lang.RxFromString("56.6607465E+467812565").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("909552512E+764516200"))
	if err != nil {
		lang.SayString("xrem264")
	} else {
		if !(rexsult.ToString() == "56.6607465E+467812565") {
			lang.SayString("xrem264")
		}
	}
	rexsult, err = lang.RxFromString("34.1935525").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-40767.6450"))
	if err != nil {
		lang.SayString("xrem266")
	} else {
		if !(rexsult.ToString() == "34.1935525") {
			lang.SayString("xrem266")
		}
	}
	rexsult, err = lang.RxFromString("51.025101").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4467691.57"))
	if err != nil {
		lang.SayString("xrem269")
	} else {
		if !(rexsult.ToString() == "51.025101") {
			lang.SayString("xrem269")
		}
	}
	rexsult, err = lang.RxFromString("-2214.76582").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("379785372E+223117572"))
	if err != nil {
		lang.SayString("xrem270")
	} else {
		if !(rexsult.ToString() == "-2214.76582") {
			lang.SayString("xrem270")
		}
	}
	rexsult, err = lang.RxFromString("-2564.75207E-841443929").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-653498187"))
	if err != nil {
		lang.SayString("xrem271")
	} else {
		if !(rexsult.ToString() == "-2564.75207E-841443929") {
			lang.SayString("xrem271")
		}
	}
	rexsult, err = lang.RxFromString("513115529.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("27775075.6E+217133352"))
	if err != nil {
		lang.SayString("xrem272")
	} else {
		if !(rexsult.ToString() == "513115529.") {
			lang.SayString("xrem272")
		}
	}
	rexsult, err = lang.RxFromString("-247157.208").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-532990.453"))
	if err != nil {
		lang.SayString("xrem273")
	} else {
		if !(rexsult.ToString() == "-247157.208") {
			lang.SayString("xrem273")
		}
	}
	rexsult, err = lang.RxFromString("40.2490764E-339482253").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7626.85442E+594264540"))
	if err != nil {
		lang.SayString("xrem274")
	} else {
		if !(rexsult.ToString() == "40.2490764E-339482253") {
			lang.SayString("xrem274")
		}
	}
	rexsult, err = lang.RxFromString("-1156008.8").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-8870382.36"))
	if err != nil {
		lang.SayString("xrem275")
	} else {
		if !(rexsult.ToString() == "-1156008.8") {
			lang.SayString("xrem275")
		}
	}
	rexsult, err = lang.RxFromString("880097928.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-52455011.1E+204538218"))
	if err != nil {
		lang.SayString("xrem276")
	} else {
		if !(rexsult.ToString() == "880097928.") {
			lang.SayString("xrem276")
		}
	}
	rexsult, err = lang.RxFromString("5796.2524").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("34458329.7E+832129426"))
	if err != nil {
		lang.SayString("xrem277")
	} else {
		if !(rexsult.ToString() == "5796.2524") {
			lang.SayString("xrem277")
		}
	}
	rexsult, err = lang.RxFromString("27.1000923E-218032223").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-45.0198341"))
	if err != nil {
		lang.SayString("xrem278")
	} else {
		if !(rexsult.ToString() == "27.1000923E-218032223") {
			lang.SayString("xrem278")
		}
	}
	rexsult, err = lang.RxFromString("84224841.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.62548255E+647087608"))
	if err != nil {
		lang.SayString("xrem281")
	} else {
		if !(rexsult.ToString() == "84224841.0") {
			lang.SayString("xrem281")
		}
	}
	rexsult, err = lang.RxFromString("9090950.80").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("436.400932"))
	if err != nil {
		lang.SayString("xrem284")
	} else {
		if !(rexsult.ToString() == "282.985508") {
			lang.SayString("xrem284")
		}
	}
	rexsult, err = lang.RxFromString("-4.18074650E-858746879").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("571035.277E-279409165"))
	if err != nil {
		lang.SayString("xrem288")
	} else {
		if !(rexsult.ToString() == "-4.18074650E-858746879") {
			lang.SayString("xrem288")
		}
	}
	rexsult, err = lang.RxFromString("5.15309635").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-695649.219E+451948183"))
	if err != nil {
		lang.SayString("xrem289")
	} else {
		if !(rexsult.ToString() == "5.15309635") {
			lang.SayString("xrem289")
		}
	}
	rexsult, err = lang.RxFromString("3336750").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("6.47961126"))
	if err != nil {
		lang.SayString("xrem292")
	} else {
		if !(rexsult.ToString() == "2.90593914") {
			lang.SayString("xrem292")
		}
	}
	rexsult, err = lang.RxFromString("904654622.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("692065270.E+329081915"))
	if err != nil {
		lang.SayString("xrem293")
	} else {
		if !(rexsult.ToString() == "904654622.") {
			lang.SayString("xrem293")
		}
	}
	rexsult, err = lang.RxFromString("304804380").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4681.23698"))
	if err != nil {
		lang.SayString("xrem294")
	} else {
		if !(rexsult.ToString() == "4358.99522") {
			lang.SayString("xrem294")
		}
	}
	rexsult, err = lang.RxFromString("674.55569").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-82981.2684E+852890752"))
	if err != nil {
		lang.SayString("xrem295")
	} else {
		if !(rexsult.ToString() == "674.55569") {
			lang.SayString("xrem295")
		}
	}
	rexsult, err = lang.RxFromString("-5111.51025E-108006096").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("5448870.4E+279212255"))
	if err != nil {
		lang.SayString("xrem296")
	} else {
		if !(rexsult.ToString() == "-5111.51025E-108006096") {
			lang.SayString("xrem296")
		}
	}
	rexsult, err = lang.RxFromString("-2623.45068").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-466463938."))
	if err != nil {
		lang.SayString("xrem297")
	} else {
		if !(rexsult.ToString() == "-2623.45068") {
			lang.SayString("xrem297")
		}
	}
	rexsult, err = lang.RxFromString("299350.435").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3373.33551"))
	if err != nil {
		lang.SayString("xrem298")
	} else {
		if !(rexsult.ToString() == "2496.91012") {
			lang.SayString("xrem298")
		}
	}
	rexsult, err = lang.RxFromString("3774.5358E-491090520").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("173.060090"))
	if err != nil {
		lang.SayString("xrem300")
	} else {
		if !(rexsult.ToString() == "3774.5358E-491090520") {
			lang.SayString("xrem300")
		}
	}
	rexsult, err = lang.RxFromString("-13.6783690").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-453.610117"))
	if err != nil {
		lang.SayString("xrem301")
	} else {
		if !(rexsult.ToString() == "-13.6783690") {
			lang.SayString("xrem301")
		}
	}
	rexsult, err = lang.RxFromString("-990100927.E-615244634").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("223801.421E+247632618"))
	if err != nil {
		lang.SayString("xrem302")
	} else {
		if !(rexsult.ToString() == "-990100927.E-615244634") {
			lang.SayString("xrem302")
		}
	}
	rexsult, err = lang.RxFromString("1275.10292").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-667965353"))
	if err != nil {
		lang.SayString("xrem303")
	} else {
		if !(rexsult.ToString() == "1275.10292") {
			lang.SayString("xrem303")
		}
	}
	rexsult, err = lang.RxFromString("-8.76375480E-596792197").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("992.077361"))
	if err != nil {
		lang.SayString("xrem304")
	} else {
		if !(rexsult.ToString() == "-8.76375480E-596792197") {
			lang.SayString("xrem304")
		}
	}
	rexsult, err = lang.RxFromString("213577152").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-986710073E+31900046"))
	if err != nil {
		lang.SayString("xrem306")
	} else {
		if !(rexsult.ToString() == "213577152") {
			lang.SayString("xrem306")
		}
	}
	rexsult, err = lang.RxFromString("91393.9398E-323439228").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-135.701000"))
	if err != nil {
		lang.SayString("xrem307")
	} else {
		if !(rexsult.ToString() == "91393.9398E-323439228") {
			lang.SayString("xrem307")
		}
	}
	rexsult, err = lang.RxFromString("59807846.1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.53345254"))
	if err != nil {
		lang.SayString("xrem309")
	} else {
		if !(rexsult.ToString() == "1.32490664") {
			lang.SayString("xrem309")
		}
	}
	rexsult, err = lang.RxFromString("-8046158.45").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8.3635397"))
	if err != nil {
		lang.SayString("xrem310")
	} else {
		if !(rexsult.ToString() == "-6.7180753") {
			lang.SayString("xrem310")
		}
	}
	rexsult, err = lang.RxFromString("-948.038054").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3580.84510"))
	if err != nil {
		lang.SayString("xrem312")
	} else {
		if !(rexsult.ToString() == "-948.038054") {
			lang.SayString("xrem312")
		}
	}
	rexsult, err = lang.RxFromString("79551.5014").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-538.186229"))
	if err != nil {
		lang.SayString("xrem314")
	} else {
		if !(rexsult.ToString() == "438.125737") {
			lang.SayString("xrem314")
		}
	}
	rexsult, err = lang.RxFromString("-3264204.54").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-42704.501"))
	if err != nil {
		lang.SayString("xrem317")
	} else {
		if !(rexsult.ToString() == "-18662.464") {
			lang.SayString("xrem317")
		}
	}
	rexsult, err = lang.RxFromString("1.21265492").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("44102.6073"))
	if err != nil {
		lang.SayString("xrem318")
	} else {
		if !(rexsult.ToString() == "1.21265492") {
			lang.SayString("xrem318")
		}
	}
	rexsult, err = lang.RxFromString("745.78452").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1922.00670E+375923302"))
	if err != nil {
		lang.SayString("xrem320")
	} else {
		if !(rexsult.ToString() == "745.78452") {
			lang.SayString("xrem320")
		}
	}
	rexsult, err = lang.RxFromString("-963717836").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-823989308"))
	if err != nil {
		lang.SayString("xrem321")
	} else {
		if !(rexsult.ToString() == "-139728528") {
			lang.SayString("xrem321")
		}
	}
	rexsult, err = lang.RxFromString("-808328.607E-790810342").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("53075.7082"))
	if err != nil {
		lang.SayString("xrem323")
	} else {
		if !(rexsult.ToString() == "-808328.607E-790810342") {
			lang.SayString("xrem323")
		}
	}
	rexsult, err = lang.RxFromString("700592.720").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-698485.085"))
	if err != nil {
		lang.SayString("xrem324")
	} else {
		if !(rexsult.ToString() == "2107.635") {
			lang.SayString("xrem324")
		}
	}
	rexsult, err = lang.RxFromString("-80273928.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("661346.239"))
	if err != nil {
		lang.SayString("xrem325")
	} else {
		if !(rexsult.ToString() == "-251033.081") {
			lang.SayString("xrem325")
		}
	}
	rexsult, err = lang.RxFromString("-682.796370").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("71131.0224"))
	if err != nil {
		lang.SayString("xrem328")
	} else {
		if !(rexsult.ToString() == "-682.796370") {
			lang.SayString("xrem328")
		}
	}
	rexsult, err = lang.RxFromString("89.9997490").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4993.69831"))
	if err != nil {
		lang.SayString("xrem329")
	} else {
		if !(rexsult.ToString() == "89.9997490") {
			lang.SayString("xrem329")
		}
	}
	rexsult, err = lang.RxFromString("-8868.72074").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("5592399.93"))
	if err != nil {
		lang.SayString("xrem335")
	} else {
		if !(rexsult.ToString() == "-8868.72074") {
			lang.SayString("xrem335")
		}
	}
	rexsult, err = lang.RxFromString("-74.7852037E-175205809").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4.14316542"))
	if err != nil {
		lang.SayString("xrem336")
	} else {
		if !(rexsult.ToString() == "-74.7852037E-175205809") {
			lang.SayString("xrem336")
		}
	}
	rexsult, err = lang.RxFromString("38660103.1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-6671.73085E+900998477"))
	if err != nil {
		lang.SayString("xrem338")
	} else {
		if !(rexsult.ToString() == "38660103.1") {
			lang.SayString("xrem338")
		}
	}
	rexsult, err = lang.RxFromString("-52.2659460").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-296404199E+372050476"))
	if err != nil {
		lang.SayString("xrem339")
	} else {
		if !(rexsult.ToString() == "-52.2659460") {
			lang.SayString("xrem339")
		}
	}
	rexsult, err = lang.RxFromString("6.06625013").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-276.359186"))
	if err != nil {
		lang.SayString("xrem340")
	} else {
		if !(rexsult.ToString() == "6.06625013") {
			lang.SayString("xrem340")
		}
	}
	rexsult, err = lang.RxFromString("-62971617.5E-241444744").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("46266799.3"))
	if err != nil {
		lang.SayString("xrem341")
	} else {
		if !(rexsult.ToString() == "-62971617.5E-241444744") {
			lang.SayString("xrem341")
		}
	}
	rexsult, err = lang.RxFromString("2467915.01").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-92.5558322"))
	if err != nil {
		lang.SayString("xrem343")
	} else {
		if !(rexsult.ToString() == "6.3002192") {
			lang.SayString("xrem343")
		}
	}
	rexsult, err = lang.RxFromString("187.232671").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-840.469347"))
	if err != nil {
		lang.SayString("xrem344")
	} else {
		if !(rexsult.ToString() == "187.232671") {
			lang.SayString("xrem344")
		}
	}
	rexsult, err = lang.RxFromString("81233.6823").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-5192.21666E+309315093"))
	if err != nil {
		lang.SayString("xrem345")
	} else {
		if !(rexsult.ToString() == "81233.6823") {
			lang.SayString("xrem345")
		}
	}
	rexsult, err = lang.RxFromString("78872665.3").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("172.102119"))
	if err != nil {
		lang.SayString("xrem347")
	} else {
		if !(rexsult.ToString() == "157.285609") {
			lang.SayString("xrem347")
		}
	}
	rexsult, err = lang.RxFromString("328268.1E-436315617").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-204.522245"))
	if err != nil {
		lang.SayString("xrem348")
	} else {
		if !(rexsult.ToString() == "328268.1E-436315617") {
			lang.SayString("xrem348")
		}
	}
	rexsult, err = lang.RxFromString("-688755561.E-95301699").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("978.275312E+913812609"))
	if err != nil {
		lang.SayString("xrem350")
	} else {
		if !(rexsult.ToString() == "-688755561.E-95301699") {
			lang.SayString("xrem350")
		}
	}
	rexsult, err = lang.RxFromString("-5.47345502").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("59818.7580"))
	if err != nil {
		lang.SayString("xrem351")
	} else {
		if !(rexsult.ToString() == "-5.47345502") {
			lang.SayString("xrem351")
		}
	}
	rexsult, err = lang.RxFromString("563891620E-361354567").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-845900362."))
	if err != nil {
		lang.SayString("xrem352")
	} else {
		if !(rexsult.ToString() == "563891620E-361354567") {
			lang.SayString("xrem352")
		}
	}
	rexsult, err = lang.RxFromString("-69.7231286").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("85773.7504"))
	if err != nil {
		lang.SayString("xrem353")
	} else {
		if !(rexsult.ToString() == "-69.7231286") {
			lang.SayString("xrem353")
		}
	}
	rexsult, err = lang.RxFromString("-54.6254096").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-332921899."))
	if err != nil {
		lang.SayString("xrem355")
	} else {
		if !(rexsult.ToString() == "-54.6254096") {
			lang.SayString("xrem355")
		}
	}
	rexsult, err = lang.RxFromString("-9.04778095E-591874079").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8719.40286"))
	if err != nil {
		lang.SayString("xrem356")
	} else {
		if !(rexsult.ToString() == "-9.04778095E-591874079") {
			lang.SayString("xrem356")
		}
	}
	rexsult, err = lang.RxFromString("-1546783").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-51935370.4"))
	if err != nil {
		lang.SayString("xrem358")
	} else {
		if !(rexsult.ToString() == "-1546783") {
			lang.SayString("xrem358")
		}
	}
	rexsult, err = lang.RxFromString("61302486.8").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("205.490417"))
	if err != nil {
		lang.SayString("xrem359")
	} else {
		if !(rexsult.ToString() == "174.619726") {
			lang.SayString("xrem359")
		}
	}
	rexsult, err = lang.RxFromString("-546398328.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-27.9149712"))
	if err != nil {
		lang.SayString("xrem362")
	} else {
		if !(rexsult.ToString() == "-5.3315808") {
			lang.SayString("xrem362")
		}
	}
	rexsult, err = lang.RxFromString("5402066.1E-284978216").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("622.751128"))
	if err != nil {
		lang.SayString("xrem363")
	} else {
		if !(rexsult.ToString() == "5402066.1E-284978216") {
			lang.SayString("xrem363")
		}
	}
	rexsult, err = lang.RxFromString("18845620").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3129.43753"))
	if err != nil {
		lang.SayString("xrem364")
	} else {
		if !(rexsult.ToString() == "147.19434") {
			lang.SayString("xrem364")
		}
	}
	rexsult, err = lang.RxFromString("13.8003883E-386224921").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-84126481.9E-296378341"))
	if err != nil {
		lang.SayString("xrem367")
	} else {
		if !(rexsult.ToString() == "13.8003883E-386224921") {
			lang.SayString("xrem367")
		}
	}
	rexsult, err = lang.RxFromString("9820.90457").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("46671.5915"))
	if err != nil {
		lang.SayString("xrem368")
	} else {
		if !(rexsult.ToString() == "9820.90457") {
			lang.SayString("xrem368")
		}
	}
	rexsult, err = lang.RxFromString("472648900").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-207.784153"))
	if err != nil {
		lang.SayString("xrem370")
	} else {
		if !(rexsult.ToString() == "1.545217") {
			lang.SayString("xrem370")
		}
	}
	rexsult, err = lang.RxFromString("-8754.49306").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-818.165153E+631475457"))
	if err != nil {
		lang.SayString("xrem371")
	} else {
		if !(rexsult.ToString() == "-8754.49306") {
			lang.SayString("xrem371")
		}
	}
	rexsult, err = lang.RxFromString("98750864").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("191380.551"))
	if err != nil {
		lang.SayString("xrem372")
	} else {
		if !(rexsult.ToString() == "189880.235") {
			lang.SayString("xrem372")
		}
	}
	rexsult, err = lang.RxFromString("725292561.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-768963606.E+340762986"))
	if err != nil {
		lang.SayString("xrem373")
	} else {
		if !(rexsult.ToString() == "725292561.") {
			lang.SayString("xrem373")
		}
	}
	rexsult, err = lang.RxFromString("1862.80445").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("648254483."))
	if err != nil {
		lang.SayString("xrem374")
	} else {
		if !(rexsult.ToString() == "1862.80445") {
			lang.SayString("xrem374")
		}
	}
	rexsult, err = lang.RxFromString("-5549320.1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-93580684.1"))
	if err != nil {
		lang.SayString("xrem375")
	} else {
		if !(rexsult.ToString() == "-5549320.1") {
			lang.SayString("xrem375")
		}
	}
	rexsult, err = lang.RxFromString("-14677053.1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-25784.7358"))
	if err != nil {
		lang.SayString("xrem376")
	} else {
		if !(rexsult.ToString() == "-5538.4298") {
			lang.SayString("xrem376")
		}
	}
	rexsult, err = lang.RxFromString("-4131738.09").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7579.07566"))
	if err != nil {
		lang.SayString("xrem378")
	} else {
		if !(rexsult.ToString() == "-1141.85530") {
			lang.SayString("xrem378")
		}
	}
	rexsult, err = lang.RxFromString("829898241").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8912.99114E+929228149"))
	if err != nil {
		lang.SayString("xrem380")
	} else {
		if !(rexsult.ToString() == "829898241") {
			lang.SayString("xrem380")
		}
	}
	rexsult, err = lang.RxFromString("53.6891691").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-11.2371140"))
	if err != nil {
		lang.SayString("xrem381")
	} else {
		if !(rexsult.ToString() == "8.7407131") {
			lang.SayString("xrem381")
		}
	}
	rexsult, err = lang.RxFromString("-93951823.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-25317.8645"))
	if err != nil {
		lang.SayString("xrem382")
	} else {
		if !(rexsult.ToString() == "-22546.1050") {
			lang.SayString("xrem382")
		}
	}
	rexsult, err = lang.RxFromString("446919.123").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("951338490."))
	if err != nil {
		lang.SayString("xrem383")
	} else {
		if !(rexsult.ToString() == "446919.123") {
			lang.SayString("xrem383")
		}
	}
	rexsult, err = lang.RxFromString("-8.01787748").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-88.3076852"))
	if err != nil {
		lang.SayString("xrem384")
	} else {
		if !(rexsult.ToString() == "-8.01787748") {
			lang.SayString("xrem384")
		}
	}
	rexsult, err = lang.RxFromString("517458139").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-999731.548"))
	if err != nil {
		lang.SayString("xrem385")
	} else {
		if !(rexsult.ToString() == "596928.684") {
			lang.SayString("xrem385")
		}
	}
	rexsult, err = lang.RxFromString("-405543440").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4013.18295"))
	if err != nil {
		lang.SayString("xrem386")
	} else {
		if !(rexsult.ToString() == "-3276.53660") {
			lang.SayString("xrem386")
		}
	}
	rexsult, err = lang.RxFromString("-151144455").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-170371.29"))
	if err != nil {
		lang.SayString("xrem388")
	} else {
		if !(rexsult.ToString() == "-25120.77") {
			lang.SayString("xrem388")
		}
	}
	rexsult, err = lang.RxFromString("534.394729").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-2369839.37"))
	if err != nil {
		lang.SayString("xrem390")
	} else {
		if !(rexsult.ToString() == "534.394729") {
			lang.SayString("xrem390")
		}
	}
	rexsult, err = lang.RxFromString("89100.1797").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("224.370309"))
	if err != nil {
		lang.SayString("xrem391")
	} else {
		if !(rexsult.ToString() == "25.167027") {
			lang.SayString("xrem391")
		}
	}
	rexsult, err = lang.RxFromString("-821377.777").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("38.552821"))
	if err != nil {
		lang.SayString("xrem392")
	} else {
		if !(rexsult.ToString() == "-9.925595") {
			lang.SayString("xrem392")
		}
	}
	rexsult, err = lang.RxFromString("-392640.782").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-2571619.5E+113237865"))
	if err != nil {
		lang.SayString("xrem393")
	} else {
		if !(rexsult.ToString() == "-392640.782") {
			lang.SayString("xrem393")
		}
	}
	rexsult, err = lang.RxFromString("-651397.712").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-723.59673"))
	if err != nil {
		lang.SayString("xrem394")
	} else {
		if !(rexsult.ToString() == "-160.65500") {
			lang.SayString("xrem394")
		}
	}
	rexsult, err = lang.RxFromString("86.6890892").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("940380864"))
	if err != nil {
		lang.SayString("xrem395")
	} else {
		if !(rexsult.ToString() == "86.6890892") {
			lang.SayString("xrem395")
		}
	}
	rexsult, err = lang.RxFromString("173398265E-532383158").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3462.51450E+80986915"))
	if err != nil {
		lang.SayString("xrem397")
	} else {
		if !(rexsult.ToString() == "173398265E-532383158") {
			lang.SayString("xrem397")
		}
	}
	rexsult, err = lang.RxFromString("-1522176.78").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-6631061.77"))
	if err != nil {
		lang.SayString("xrem398")
	} else {
		if !(rexsult.ToString() == "-1522176.78") {
			lang.SayString("xrem398")
		}
	}
	rexsult, err = lang.RxFromString("538.10453").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("522934310"))
	if err != nil {
		lang.SayString("xrem399")
	} else {
		if !(rexsult.ToString() == "538.10453") {
			lang.SayString("xrem399")
		}
	}
	rexsult, err = lang.RxFromString("880243.444E-750940977").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-354.601578E-204943740"))
	if err != nil {
		lang.SayString("xrem400")
	} else {
		if !(rexsult.ToString() == "880243.444E-750940977") {
			lang.SayString("xrem400")
		}
	}
	rexsult, err = lang.RxFromString("968370.780").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("6677268.73"))
	if err != nil {
		lang.SayString("xrem401")
	} else {
		if !(rexsult.ToString() == "968370.780") {
			lang.SayString("xrem401")
		}
	}
	rexsult, err = lang.RxFromString("-97.7474945").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("31248241.5"))
	if err != nil {
		lang.SayString("xrem402")
	} else {
		if !(rexsult.ToString() == "-97.7474945") {
			lang.SayString("xrem402")
		}
	}
	rexsult, err = lang.RxFromString("-187582786.E+369916952").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("957840435E+744365567"))
	if err != nil {
		lang.SayString("xrem403")
	} else {
		if !(rexsult.ToString() == "-187582786.E+369916952") {
			lang.SayString("xrem403")
		}
	}
	rexsult, err = lang.RxFromString("-328026144.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-125499735."))
	if err != nil {
		lang.SayString("xrem404")
	} else {
		if !(rexsult.ToString() == "-77026674") {
			lang.SayString("xrem404")
		}
	}
	rexsult, err = lang.RxFromString("-99424150.2E-523662102").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3712.35030"))
	if err != nil {
		lang.SayString("xrem405")
	} else {
		if !(rexsult.ToString() == "-99424150.2E-523662102") {
			lang.SayString("xrem405")
		}
	}
	rexsult, err = lang.RxFromString("14838.0718").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9489893.28E+830631266"))
	if err != nil {
		lang.SayString("xrem406")
	} else {
		if !(rexsult.ToString() == "14838.0718") {
			lang.SayString("xrem406")
		}
	}
	rexsult, err = lang.RxFromString("71207472.8").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-31835.0809"))
	if err != nil {
		lang.SayString("xrem407")
	} else {
		if !(rexsult.ToString() == "24231.9076") {
			lang.SayString("xrem407")
		}
	}
	rexsult, err = lang.RxFromString("-20440.4394").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-44.4064328E+511085806"))
	if err != nil {
		lang.SayString("xrem408")
	} else {
		if !(rexsult.ToString() == "-20440.4394") {
			lang.SayString("xrem408")
		}
	}
	rexsult, err = lang.RxFromString("-657.186702").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("426844.39"))
	if err != nil {
		lang.SayString("xrem411")
	} else {
		if !(rexsult.ToString() == "-657.186702") {
			lang.SayString("xrem411")
		}
	}
	rexsult, err = lang.RxFromString("-41593077.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-688607.564"))
	if err != nil {
		lang.SayString("xrem412")
	} else {
		if !(rexsult.ToString() == "-276623.160") {
			lang.SayString("xrem412")
		}
	}
	rexsult, err = lang.RxFromString("-5786.38132").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("190556652.E+177045877"))
	if err != nil {
		lang.SayString("xrem413")
	} else {
		if !(rexsult.ToString() == "-5786.38132") {
			lang.SayString("xrem413")
		}
	}
	rexsult, err = lang.RxFromString("737622.974").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-241560693E+249506565"))
	if err != nil {
		lang.SayString("xrem414")
	} else {
		if !(rexsult.ToString() == "737622.974") {
			lang.SayString("xrem414")
		}
	}
	rexsult, err = lang.RxFromString("644136.179").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-835708.103"))
	if err != nil {
		lang.SayString("xrem416")
	} else {
		if !(rexsult.ToString() == "644136.179") {
			lang.SayString("xrem416")
		}
	}
	rexsult, err = lang.RxFromString("-31068.7549").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-3.41495042E+86001379"))
	if err != nil {
		lang.SayString("xrem419")
	} else {
		if !(rexsult.ToString() == "-31068.7549") {
			lang.SayString("xrem419")
		}
	}
	rexsult, err = lang.RxFromString("3898.03188").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-82572.615"))
	if err != nil {
		lang.SayString("xrem422")
	} else {
		if !(rexsult.ToString() == "3898.03188") {
			lang.SayString("xrem422")
		}
	}
	rexsult, err = lang.RxFromString("-1.7619356").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-2546.64043"))
	if err != nil {
		lang.SayString("xrem423")
	} else {
		if !(rexsult.ToString() == "-1.7619356") {
			lang.SayString("xrem423")
		}
	}
	rexsult, err = lang.RxFromString("6.88891136E-935467395").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-785049.562E-741671442"))
	if err != nil {
		lang.SayString("xrem425")
	} else {
		if !(rexsult.ToString() == "6.88891136E-935467395") {
			lang.SayString("xrem425")
		}
	}
	rexsult, err = lang.RxFromString("975566251").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-519.858530"))
	if err != nil {
		lang.SayString("xrem426")
	} else {
		if !(rexsult.ToString() == "253.460530") {
			lang.SayString("xrem426")
		}
	}
	rexsult, err = lang.RxFromString("307401954").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-231481582."))
	if err != nil {
		lang.SayString("xrem427")
	} else {
		if !(rexsult.ToString() == "75920372") {
			lang.SayString("xrem427")
		}
	}
	rexsult, err = lang.RxFromString("6.48674979").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-621732.532E+422575800"))
	if err != nil {
		lang.SayString("xrem430")
	} else {
		if !(rexsult.ToString() == "6.48674979") {
			lang.SayString("xrem430")
		}
	}
	rexsult, err = lang.RxFromString("-31401.9418").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("36.3960679"))
	if err != nil {
		lang.SayString("xrem431")
	} else {
		if !(rexsult.ToString() == "-28.5312702") {
			lang.SayString("xrem431")
		}
	}
	rexsult, err = lang.RxFromString("31345321.1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("51.5482191"))
	if err != nil {
		lang.SayString("xrem432")
	} else {
		if !(rexsult.ToString() == "34.6743293") {
			lang.SayString("xrem432")
		}
	}
	rexsult, err = lang.RxFromString("70437.1551").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-62916.1233"))
	if err != nil {
		lang.SayString("xrem434")
	} else {
		if !(rexsult.ToString() == "7521.0318") {
			lang.SayString("xrem434")
		}
	}
	rexsult, err = lang.RxFromString("916260164").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-58.4017325"))
	if err != nil {
		lang.SayString("xrem435")
	} else {
		if !(rexsult.ToString() == "54.9461000") {
			lang.SayString("xrem435")
		}
	}
	rexsult, err = lang.RxFromString("19889085.3E-46816480").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1581683.94"))
	if err != nil {
		lang.SayString("xrem436")
	} else {
		if !(rexsult.ToString() == "19889085.3E-46816480") {
			lang.SayString("xrem436")
		}
	}
	rexsult, err = lang.RxFromString("-56312.3383").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("789.466064"))
	if err != nil {
		lang.SayString("xrem437")
	} else {
		if !(rexsult.ToString() == "-260.247756") {
			lang.SayString("xrem437")
		}
	}
	rexsult, err = lang.RxFromString("183442.849").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-925876106"))
	if err != nil {
		lang.SayString("xrem438")
	} else {
		if !(rexsult.ToString() == "183442.849") {
			lang.SayString("xrem438")
		}
	}
	rexsult, err = lang.RxFromString("859658551.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("72338.2054"))
	if err != nil {
		lang.SayString("xrem440")
	} else {
		if !(rexsult.ToString() == "63656.2318") {
			lang.SayString("xrem440")
		}
	}
	rexsult, err = lang.RxFromString("-969.881818").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("31170.8555"))
	if err != nil {
		lang.SayString("xrem442")
	} else {
		if !(rexsult.ToString() == "-969.881818") {
			lang.SayString("xrem442")
		}
	}
	rexsult, err = lang.RxFromString("7980537.27").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("85.4040512"))
	if err != nil {
		lang.SayString("xrem443")
	} else {
		if !(rexsult.ToString() == "41.1096672") {
			lang.SayString("xrem443")
		}
	}
	rexsult, err = lang.RxFromString("-114609916.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7525.14981"))
	if err != nil {
		lang.SayString("xrem444")
	} else {
		if !(rexsult.ToString() == "-1884.39370") {
			lang.SayString("xrem444")
		}
	}
	rexsult, err = lang.RxFromString("8.43404682E-500572568").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("474526719"))
	if err != nil {
		lang.SayString("xrem445")
	} else {
		if !(rexsult.ToString() == "8.43404682E-500572568") {
			lang.SayString("xrem445")
		}
	}
	rexsult, err = lang.RxFromString("-9.95836312").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-866466703"))
	if err != nil {
		lang.SayString("xrem447")
	} else {
		if !(rexsult.ToString() == "-9.95836312") {
			lang.SayString("xrem447")
		}
	}
	rexsult, err = lang.RxFromString("80919339.2E-967231586").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("219.824266"))
	if err != nil {
		lang.SayString("xrem448")
	} else {
		if !(rexsult.ToString() == "80919339.2E-967231586") {
			lang.SayString("xrem448")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-89827.5229"))
	if err != nil {
		lang.SayString("xrem449")
	} else {
		if !(rexsult.ToString() == "69751.9211") {
			lang.SayString("xrem449")
		}
	}
	rexsult, err = lang.RxFromString("-4.54000153").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("6966333.74"))
	if err != nil {
		lang.SayString("xrem450")
	} else {
		if !(rexsult.ToString() == "-4.54000153") {
			lang.SayString("xrem450")
		}
	}
	rexsult, err = lang.RxFromString("28701538.7E-391015649").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-920999192."))
	if err != nil {
		lang.SayString("xrem451")
	} else {
		if !(rexsult.ToString() == "28701538.7E-391015649") {
			lang.SayString("xrem451")
		}
	}
	rexsult, err = lang.RxFromString("-361382575.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-7976.15286E+898491169"))
	if err != nil {
		lang.SayString("xrem452")
	} else {
		if !(rexsult.ToString() == "-361382575.") {
			lang.SayString("xrem452")
		}
	}
	rexsult, err = lang.RxFromString("7021805.61").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1222952.83"))
	if err != nil {
		lang.SayString("xrem453")
	} else {
		if !(rexsult.ToString() == "907041.46") {
			lang.SayString("xrem453")
		}
	}
	rexsult, err = lang.RxFromString("-40.4811667").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-79655.5635"))
	if err != nil {
		lang.SayString("xrem454")
	} else {
		if !(rexsult.ToString() == "-40.4811667") {
			lang.SayString("xrem454")
		}
	}
	rexsult, err = lang.RxFromString("-37958476.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("584367.935"))
	if err != nil {
		lang.SayString("xrem457")
	} else {
		if !(rexsult.ToString() == "-558928.160") {
			lang.SayString("xrem457")
		}
	}
	rexsult, err = lang.RxFromString("495233.553E-414152215").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("62352759.2"))
	if err != nil {
		lang.SayString("xrem458")
	} else {
		if !(rexsult.ToString() == "495233.553E-414152215") {
			lang.SayString("xrem458")
		}
	}
	rexsult, err = lang.RxFromString("-502343060").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-96828.994"))
	if err != nil {
		lang.SayString("xrem459")
	} else {
		if !(rexsult.ToString() == "-91068.122") {
			lang.SayString("xrem459")
		}
	}
	rexsult, err = lang.RxFromString("718180.587E-957473722").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.66223443"))
	if err != nil {
		lang.SayString("xrem461")
	} else {
		if !(rexsult.ToString() == "718180.587E-957473722") {
			lang.SayString("xrem461")
		}
	}
	rexsult, err = lang.RxFromString("-51592.2698").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-713885.741"))
	if err != nil {
		lang.SayString("xrem462")
	} else {
		if !(rexsult.ToString() == "-51592.2698") {
			lang.SayString("xrem462")
		}
	}
	rexsult, err = lang.RxFromString("51.2279848E+80439745").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("207.55925E+865165070"))
	if err != nil {
		lang.SayString("xrem463")
	} else {
		if !(rexsult.ToString() == "51.2279848E+80439745") {
			lang.SayString("xrem463")
		}
	}
	rexsult, err = lang.RxFromString("-5983.23468").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-39.9544513"))
	if err != nil {
		lang.SayString("xrem464")
	} else {
		if !(rexsult.ToString() == "-30.0214363") {
			lang.SayString("xrem464")
		}
	}
	rexsult, err = lang.RxFromString("921639332.E-917542963").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("287325.891"))
	if err != nil {
		lang.SayString("xrem465")
	} else {
		if !(rexsult.ToString() == "921639332.E-917542963") {
			lang.SayString("xrem465")
		}
	}
	rexsult, err = lang.RxFromString("91095916.8E-787312969").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-58643.418E+58189880"))
	if err != nil {
		lang.SayString("xrem466")
	} else {
		if !(rexsult.ToString() == "91095916.8E-787312969") {
			lang.SayString("xrem466")
		}
	}
	rexsult, err = lang.RxFromString("-6410.5555").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-234964259"))
	if err != nil {
		lang.SayString("xrem467")
	} else {
		if !(rexsult.ToString() == "-6410.5555") {
			lang.SayString("xrem467")
		}
	}
	rexsult, err = lang.RxFromString("-5.32711606").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-8447286.21"))
	if err != nil {
		lang.SayString("xrem468")
	} else {
		if !(rexsult.ToString() == "-5.32711606") {
			lang.SayString("xrem468")
		}
	}
	rexsult, err = lang.RxFromString("412411244.E-774339264").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("866452.465"))
	if err != nil {
		lang.SayString("xrem470")
	} else {
		if !(rexsult.ToString() == "412411244.E-774339264") {
			lang.SayString("xrem470")
		}
	}
	rexsult, err = lang.RxFromString("-31027.8323").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-475378186."))
	if err != nil {
		lang.SayString("xrem472")
	} else {
		if !(rexsult.ToString() == "-31027.8323") {
			lang.SayString("xrem472")
		}
	}
	rexsult, err = lang.RxFromString("-1199339.72").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-5.73068392E+53774632"))
	if err != nil {
		lang.SayString("xrem473")
	} else {
		if !(rexsult.ToString() == "-1199339.72") {
			lang.SayString("xrem473")
		}
	}
	rexsult, err = lang.RxFromString("-2376150.83").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-46777583.3"))
	if err != nil {
		lang.SayString("xrem475")
	} else {
		if !(rexsult.ToString() == "-2376150.83") {
			lang.SayString("xrem475")
		}
	}
	rexsult, err = lang.RxFromString("6.3664211").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-140854908."))
	if err != nil {
		lang.SayString("xrem476")
	} else {
		if !(rexsult.ToString() == "6.3664211") {
			lang.SayString("xrem476")
		}
	}
	rexsult, err = lang.RxFromString("-15.791522").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1902.30210E+90741844"))
	if err != nil {
		lang.SayString("xrem477")
	} else {
		if !(rexsult.ToString() == "-15.791522") {
			lang.SayString("xrem477")
		}
	}
	rexsult, err = lang.RxFromString("49436.6528").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("751.919517"))
	if err != nil {
		lang.SayString("xrem480")
	} else {
		if !(rexsult.ToString() == "561.884195") {
			lang.SayString("xrem480")
		}
	}
	rexsult, err = lang.RxFromString("552.669453").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8.3725760E+16223526"))
	if err != nil {
		lang.SayString("xrem481")
	} else {
		if !(rexsult.ToString() == "552.669453") {
			lang.SayString("xrem481")
		}
	}
	rexsult, err = lang.RxFromString("-3266303").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("453741.520"))
	if err != nil {
		lang.SayString("xrem482")
	} else {
		if !(rexsult.ToString() == "-90112.360") {
			lang.SayString("xrem482")
		}
	}
	rexsult, err = lang.RxFromString("12302757.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("542922.487E+414443353"))
	if err != nil {
		lang.SayString("xrem483")
	} else {
		if !(rexsult.ToString() == "12302757.4") {
			lang.SayString("xrem483")
		}
	}
	rexsult, err = lang.RxFromString("-5670757.79E-784754984").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("128144.503"))
	if err != nil {
		lang.SayString("xrem484")
	} else {
		if !(rexsult.ToString() == "-5670757.79E-784754984") {
			lang.SayString("xrem484")
		}
	}
	rexsult, err = lang.RxFromString("88.5158199E-980164357").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("325846116"))
	if err != nil {
		lang.SayString("xrem486")
	} else {
		if !(rexsult.ToString() == "88.5158199E-980164357") {
			lang.SayString("xrem486")
		}
	}
	rexsult, err = lang.RxFromString("-22881.0408").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("5.63661562"))
	if err != nil {
		lang.SayString("xrem487")
	} else {
		if !(rexsult.ToString() == "-2.01799842") {
			lang.SayString("xrem487")
		}
	}
	rexsult, err = lang.RxFromString("-3066962.41").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-55.3096879"))
	if err != nil {
		lang.SayString("xrem490")
	} else {
		if !(rexsult.ToString() == "-40.2159450") {
			lang.SayString("xrem490")
		}
	}
	rexsult, err = lang.RxFromString("998890068.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-92.057879"))
	if err != nil {
		lang.SayString("xrem492")
	} else {
		if !(rexsult.ToString() == "33.839554") {
			lang.SayString("xrem492")
		}
	}
	rexsult, err = lang.RxFromString("122.495591").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-407836028."))
	if err != nil {
		lang.SayString("xrem493")
	} else {
		if !(rexsult.ToString() == "122.495591") {
			lang.SayString("xrem493")
		}
	}
	rexsult, err = lang.RxFromString("-286.371320").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("710319152"))
	if err != nil {
		lang.SayString("xrem497")
	} else {
		if !(rexsult.ToString() == "-286.371320") {
			lang.SayString("xrem497")
		}
	}
	rexsult, err = lang.RxFromString("-6157.74292").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-94075286.2E+92555877"))
	if err != nil {
		lang.SayString("xrem499")
	} else {
		if !(rexsult.ToString() == "-6157.74292") {
			lang.SayString("xrem499")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx001")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx003")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx004")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("remx007")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("remx008")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("remx008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("remx009")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx010")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("remx010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx011")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("remx011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx012")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("remx012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx013")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("remx013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx014")
	} else {
		if !(rexsult.ToString() == "0.40") {
			lang.SayString("remx014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx015")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("remx015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx016")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("remx016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx017")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("remx017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx018")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("remx019")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("remx020")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx021")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx021")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("remx022")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("remx022")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("remx023")
	} else {
		if !(rexsult.ToString() == "1.000") {
			lang.SayString("remx023")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("remx024")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("remx025")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx025")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx030")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx030")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("remx031")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx031")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("remx032")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("remx033")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("remx034")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx034")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("remx035")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("remx040")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("remx041")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("remx042")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("remx043")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("remx044")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("remx045")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx045")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx050")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("remx051")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("remx052")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("remx053")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("remx054")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("remx055")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("remx060")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx060")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("remx061")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx061")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("remx062")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx062")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("remx063")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx063")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("remx064")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx064")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("remx065")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("remx065")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx066")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx066")
		}
	}
	rexsult, err = lang.RxFromString("999999999.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx067")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("remx067")
		}
	}
	rexsult, err = lang.RxFromString("999999999.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx068")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx068")
		}
	}
	rexsult, err = lang.RxFromString("999999999.9").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx069")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("remx069")
		}
	}
	rexsult, err = lang.RxFromString("999999999.999").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx070")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("remx070")
		}
	}
	rexsult, err = lang.RxFromString("999999").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx074")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx074")
		}
	}
	rexsult, err = lang.RxFromString("99999").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx075")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx075")
		}
	}
	rexsult, err = lang.RxFromString("9999").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx076")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx076")
		}
	}
	rexsult, err = lang.RxFromString("999").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx077")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx077")
		}
	}
	rexsult, err = lang.RxFromString("99").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx078")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx078")
		}
	}
	rexsult, err = lang.RxFromString("9").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx079")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx079")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx080")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx080")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx081")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx081")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx082")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx082")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx083")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx083")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+3").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx084")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx084")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx085")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx085")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx086")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx086")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx087")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx087")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx088")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx088")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx089")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx089")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx090")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-3").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx091")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx093")
		}
	}
	rexsult, err = lang.RxFromString("0.00E-6").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx094")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx095")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx095")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx130")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx130")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx131")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx131")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx132")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx132")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx133")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx133")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx134")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx134")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx135")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx135")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx136")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx136")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx137")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx137")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx143")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx143")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("remx144")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx144")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("remx145")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx145")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("remx146")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx146")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx147")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("remx147")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("remx148")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("remx148")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("remx149")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("remx149")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx150")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx150")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx151")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx151")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("remx152")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx152")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("remx153")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx153")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx154")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx154")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("remx155")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx155")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("remx156")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx156")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("remx157")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx157")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("remx158")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx158")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx159")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx159")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx160")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx160")
		}
	}
	rexsult, err = lang.RxFromString("2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx161")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx161")
		}
	}
	rexsult, err = lang.RxFromString("3").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx162")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx162")
		}
	}
	rexsult, err = lang.RxFromString("4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx163")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx163")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx164")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx164")
		}
	}
	rexsult, err = lang.RxFromString("6").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx165")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx165")
		}
	}
	rexsult, err = lang.RxFromString("7").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx166")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx166")
		}
	}
	rexsult, err = lang.RxFromString("8").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx167")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx167")
		}
	}
	rexsult, err = lang.RxFromString("9").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx168")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx168")
		}
	}
	rexsult, err = lang.RxFromString("10").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx169")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx169")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx171")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("remx171")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx172")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("remx172")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx173")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("remx173")
		}
	}
	rexsult, err = lang.RxFromString("0.52").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx174")
	} else {
		if !(rexsult.ToString() == "0.52") {
			lang.SayString("remx174")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx175")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("remx175")
		}
	}
	rexsult, err = lang.RxFromString("-0.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx231")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("remx231")
		}
	}
	rexsult, err = lang.RxFromString("-0.50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx232")
	} else {
		if !(rexsult.ToString() == "-0.50") {
			lang.SayString("remx232")
		}
	}
	rexsult, err = lang.RxFromString("-0.51").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx233")
	} else {
		if !(rexsult.ToString() == "-0.51") {
			lang.SayString("remx233")
		}
	}
	rexsult, err = lang.RxFromString("-0.52").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx234")
	} else {
		if !(rexsult.ToString() == "-0.52") {
			lang.SayString("remx234")
		}
	}
	rexsult, err = lang.RxFromString("-0.6").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.020"))
	if err != nil {
		lang.SayString("remx235")
	} else {
		if !(rexsult.ToString() == "-0.6") {
			lang.SayString("remx235")
		}
	}
	rexsult, err = lang.RxFromString("1E+2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("remx240")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx240")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("remx301")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx301")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx302")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx303")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx303")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("50"))
	if err != nil {
		lang.SayString("remx304")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("remx304")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("remx305")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("remx305")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("remx306")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("remx306")
		}
	}
	rexsult, err = lang.RxFromString(".13").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx307")
	} else {
		if !(rexsult.ToString() == ".13") {
			lang.SayString("remx307")
		}
	}
	rexsult, err = lang.RxFromString("0.133").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx308")
	} else {
		if !(rexsult.ToString() == "0.133") {
			lang.SayString("remx308")
		}
	}
	rexsult, err = lang.RxFromString("0.1033").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx309")
	} else {
		if !(rexsult.ToString() == "0.1033") {
			lang.SayString("remx309")
		}
	}
	rexsult, err = lang.RxFromString("1.033").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx310")
	} else {
		if !(rexsult.ToString() == "0.033") {
			lang.SayString("remx310")
		}
	}
	rexsult, err = lang.RxFromString("10.33").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx311")
	} else {
		if !(rexsult.ToString() == "0.33") {
			lang.SayString("remx311")
		}
	}
	rexsult, err = lang.RxFromString("10.33").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx312")
	} else {
		if !(rexsult.ToString() == "0.33") {
			lang.SayString("remx312")
		}
	}
	rexsult, err = lang.RxFromString("103.3").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx313")
	} else {
		if !(rexsult.ToString() == "0.3") {
			lang.SayString("remx313")
		}
	}
	rexsult, err = lang.RxFromString("133").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx314")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx314")
		}
	}
	rexsult, err = lang.RxFromString("1033").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx315")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx315")
		}
	}
	rexsult, err = lang.RxFromString("1033").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("50"))
	if err != nil {
		lang.SayString("remx316")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("remx316")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("remx317")
	} else {
		if !(rexsult.ToString() == "2.0") {
			lang.SayString("remx317")
		}
	}
	rexsult, err = lang.RxFromString("102.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("remx318")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx318")
		}
	}
	rexsult, err = lang.RxFromString("103.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("remx319")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("remx319")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx320")
	} else {
		if !(rexsult.ToString() == "0.40") {
			lang.SayString("remx320")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx321")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("remx321")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx322")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("remx322")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx323")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("remx323")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx324")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("remx324")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("remx325")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("remx325")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.30"))
	if err != nil {
		lang.SayString("remx326")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("remx326")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.300"))
	if err != nil {
		lang.SayString("remx327")
	} else {
		if !(rexsult.ToString() == "0.100") {
			lang.SayString("remx327")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.3000"))
	if err != nil {
		lang.SayString("remx328")
	} else {
		if !(rexsult.ToString() == "0.1000") {
			lang.SayString("remx328")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("remx329")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("remx329")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("remx330")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("remx330")
		}
	}
	rexsult, err = lang.RxFromString("1.000").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("remx331")
	} else {
		if !(rexsult.ToString() == "0.100") {
			lang.SayString("remx331")
		}
	}
	rexsult, err = lang.RxFromString("1.0000").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("remx332")
	} else {
		if !(rexsult.ToString() == "0.1000") {
			lang.SayString("remx332")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx333")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx333")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("remx334")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx334")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("remx335")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx335")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("remx336")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx336")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx337")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("remx337")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("remx338")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("remx338")
		}
	}
	rexsult, err = lang.RxFromString("0.50").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("remx339")
	} else {
		if !(rexsult.ToString() == "0.50") {
			lang.SayString("remx339")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.5000001"))
	if err != nil {
		lang.SayString("remx340")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx340")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.50000001"))
	if err != nil {
		lang.SayString("remx341")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx341")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("remx342")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx342")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.5000000001"))
	if err != nil {
		lang.SayString("remx343")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx343")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.50000000001"))
	if err != nil {
		lang.SayString("remx344")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx344")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.4999999"))
	if err != nil {
		lang.SayString("remx345")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("remx345")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.49999999"))
	if err != nil {
		lang.SayString("remx346")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("remx346")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("remx347")
	} else {
		if !(rexsult.ToString() == "1E-9") {
			lang.SayString("remx347")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.4999999999"))
	if err != nil {
		lang.SayString("remx348")
	} else {
		if !(rexsult.ToString() == "1E-10") {
			lang.SayString("remx348")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.49999999999"))
	if err != nil {
		lang.SayString("remx349")
	} else {
		if !(rexsult.ToString() == "1E-10") {
			lang.SayString("remx349")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.499999999999"))
	if err != nil {
		lang.SayString("remx350")
	} else {
		if !(rexsult.ToString() == "1E-10") {
			lang.SayString("remx350")
		}
	}
	rexsult, err = lang.RxFromString("0.03").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("remx351")
	} else {
		if !(rexsult.ToString() == "0.03") {
			lang.SayString("remx351")
		}
	}
	rexsult, err = lang.RxFromString("5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx352")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx352")
		}
	}
	rexsult, err = lang.RxFromString("4.1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx353")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("remx353")
		}
	}
	rexsult, err = lang.RxFromString("4.01").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx354")
	} else {
		if !(rexsult.ToString() == "0.01") {
			lang.SayString("remx354")
		}
	}
	rexsult, err = lang.RxFromString("4.001").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx355")
	} else {
		if !(rexsult.ToString() == "0.001") {
			lang.SayString("remx355")
		}
	}
	rexsult, err = lang.RxFromString("4.0001").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx356")
	} else {
		if !(rexsult.ToString() == "0.0001") {
			lang.SayString("remx356")
		}
	}
	rexsult, err = lang.RxFromString("4.00001").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx357")
	} else {
		if !(rexsult.ToString() == "0.00001") {
			lang.SayString("remx357")
		}
	}
	rexsult, err = lang.RxFromString("4.000001").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx358")
	} else {
		if !(rexsult.ToString() == "0.000001") {
			lang.SayString("remx358")
		}
	}
	rexsult, err = lang.RxFromString("4.0000001").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx359")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("remx359")
		}
	}
	rexsult, err = lang.RxFromString("1.2").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.7345"))
	if err != nil {
		lang.SayString("remx360")
	} else {
		if !(rexsult.ToString() == "0.4655") {
			lang.SayString("remx360")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("remx361")
	} else {
		if !(rexsult.ToString() == "0.8") {
			lang.SayString("remx361")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.2"))
	if err != nil {
		lang.SayString("remx362")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx362")
		}
	}
	rexsult, err = lang.RxFromString("0.8").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("remx363")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("remx363")
		}
	}
	rexsult, err = lang.RxFromString("0.800").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("remx364")
	} else {
		if !(rexsult.ToString() == "0.800") {
			lang.SayString("remx364")
		}
	}
	rexsult, err = lang.RxFromString("0.800").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1.7"))
	if err != nil {
		lang.SayString("remx365")
	} else {
		if !(rexsult.ToString() == "0.800") {
			lang.SayString("remx365")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx366")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("remx366")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("remx371")
	} else {
		if !(rexsult.ToString() == "0.400") {
			lang.SayString("remx371")
		}
	}
	rexsult, err = lang.RxFromString("12345678900000").OpRem(lang.RxSetWithDigit(3), lang.RxFromString("12e+12"))
	if err != nil {
		lang.SayString("remx372")
	} else {
		if !(rexsult.ToString() == "3.4E+11") {
			lang.SayString("remx372")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx381")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx381")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("remx382")
	} else {
		if !(rexsult.ToString() == "0.7657") {
			lang.SayString("remx382")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("remx383")
	} else {
		if !(rexsult.ToString() == "0.668") {
			lang.SayString("remx383")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("remx384")
	} else {
		if !(rexsult.ToString() == "0.78") {
			lang.SayString("remx384")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("remx385")
	} else {
		if !(rexsult.ToString() == "0.8") {
			lang.SayString("remx385")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("remx386")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx386")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("remx387")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx387")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("remx388")
	} else {
		if !(rexsult.ToString() == "2.6912") {
			lang.SayString("remx388")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("remx389")
	} else {
		if !(rexsult.ToString() == "0.6914") {
			lang.SayString("remx389")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("remx390")
	} else {
		if !(rexsult.ToString() == "1.9") {
			lang.SayString("remx390")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("remx391")
	} else {
		if !(rexsult.ToString() == "4.73") {
			lang.SayString("remx391")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("remx392")
	} else {
		if !(rexsult.ToString() == "2.469") {
			lang.SayString("remx392")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("remx393")
	} else {
		if !(rexsult.ToString() == "0.2469") {
			lang.SayString("remx393")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx394")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx394")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("remx395")
	} else {
		if !(rexsult.ToString() == "4.7532") {
			lang.SayString("remx395")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("remx396")
	} else {
		if !(rexsult.ToString() == "2.532") {
			lang.SayString("remx396")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("remx397")
	} else {
		if !(rexsult.ToString() == "0.36") {
			lang.SayString("remx397")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpRem(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("remx398")
	} else {
		if !(rexsult.ToString() == "3.0") {
			lang.SayString("remx398")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx401")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("remx401")
		}
	}
	rexsult, err = lang.RxFromString("0.55").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx402")
	} else {
		if !(rexsult.ToString() == "0.55") {
			lang.SayString("remx402")
		}
	}
	rexsult, err = lang.RxFromString("0.555").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx403")
	} else {
		if !(rexsult.ToString() == "0.555") {
			lang.SayString("remx403")
		}
	}
	rexsult, err = lang.RxFromString("0.5555").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx404")
	} else {
		if !(rexsult.ToString() == "0.5555") {
			lang.SayString("remx404")
		}
	}
	rexsult, err = lang.RxFromString("0.55555").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx405")
	} else {
		if !(rexsult.ToString() == "0.55555") {
			lang.SayString("remx405")
		}
	}
	rexsult, err = lang.RxFromString("0.555555").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx406")
	} else {
		if !(rexsult.ToString() == "0.555555") {
			lang.SayString("remx406")
		}
	}
	rexsult, err = lang.RxFromString("0.5555555").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx407")
	} else {
		if !(rexsult.ToString() == "0.5555555") {
			lang.SayString("remx407")
		}
	}
	rexsult, err = lang.RxFromString("0.55555555").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx408")
	} else {
		if !(rexsult.ToString() == "0.55555555") {
			lang.SayString("remx408")
		}
	}
	rexsult, err = lang.RxFromString("0.555555555").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx409")
	} else {
		if !(rexsult.ToString() == "0.555555555") {
			lang.SayString("remx409")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx650")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx650")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx651")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx651")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx652")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx652")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx653")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx653")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx654")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx654")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx655")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx655")
		}
	}
	rexsult, err = lang.RxFromString("0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx656")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx656")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("remx657")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx657")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx658")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx658")
		}
	}
	rexsult, err = lang.RxFromString("-0.00").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("remx659")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx659")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("remx801")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx801")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("remx802")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx802")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx803")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx803")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("remx804")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx804")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx805")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx805")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("remx806")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx806")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx807")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx807")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("remx808")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx808")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("remx809")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx809")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("remx810")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx810")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx811")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("remx811")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("remx812")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx812")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("remx821")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx821")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("remx822")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx822")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx823")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx823")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("remx824")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx824")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx825")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("remx825")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("remx826")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx826")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx827")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx827")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("remx828")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx828")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("remx829")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx829")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("remx830")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx830")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("remx831")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("remx831")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(15), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("remx832")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx832")
		}
	}
	rexsult, err = lang.RxFromString("1000003").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx852")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx852")
		}
	}
	rexsult, err = lang.RxFromString("100003").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx853")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx853")
		}
	}
	rexsult, err = lang.RxFromString("10003").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx854")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx854")
		}
	}
	rexsult, err = lang.RxFromString("1003").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx855")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx855")
		}
	}
	rexsult, err = lang.RxFromString("103").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx856")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx856")
		}
	}
	rexsult, err = lang.RxFromString("13").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx857")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("remx857")
		}
	}
	rexsult, err = lang.RxFromString("1").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("remx858")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("remx858")
		}
	}
	rexsult, err = lang.RxFromString("1230").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("10000000000000000"))
	if err != nil {
		lang.SayString("remx861")
	} else {
		if !(rexsult.ToString() == "1230") {
			lang.SayString("remx861")
		}
	}
	rexsult, err = lang.RxFromString("1230").OpRem(lang.RxSetWithDigit(6), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("remx878")
	} else {
		if !(rexsult.ToString() == "1230") {
			lang.SayString("remx878")
		}
	}
	rexsult, err = lang.RxFromString("+1.23456789012345E-0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9E+999999999"))
	if err != nil {
		lang.SayString("remx990")
	} else {
		if !(rexsult.ToString() == "1.23456789") {
			lang.SayString("remx990")
		}
	}
	rexsult, err = lang.RxFromString("+0.100").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9E+999999999"))
	if err != nil {
		lang.SayString("remx992")
	} else {
		if !(rexsult.ToString() == "+0.100") {
			lang.SayString("remx992")
		}
	}
	rexsult, err = lang.RxFromString("9E-999999999").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("+9.100"))
	if err != nil {
		lang.SayString("remx993")
	} else {
		if !(rexsult.ToString() == "9E-999999999") {
			lang.SayString("remx993")
		}
	}
	rexsult, err = lang.RxFromString("-1.23456789012345E-0").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9E+999999999"))
	if err != nil {
		lang.SayString("remx995")
	} else {
		if !(rexsult.ToString() == "-1.23456789") {
			lang.SayString("remx995")
		}
	}
	rexsult, err = lang.RxFromString("-0.100").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("9E+999999999"))
	if err != nil {
		lang.SayString("remx997")
	} else {
		if !(rexsult.ToString() == "-0.100") {
			lang.SayString("remx997")
		}
	}
	rexsult, err = lang.RxFromString("9E-999999999").OpRem(lang.RxSetWithDigit(9), lang.RxFromString("-9.100"))
	if err != nil {
		lang.SayString("remx998")
	} else {
		if !(rexsult.ToString() == "9E-999999999") {
			lang.SayString("remx998")
		}
	}

	return
}
