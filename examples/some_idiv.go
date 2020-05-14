package main

import "agorexx/lang"

func main() {

	rexsult, err := lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi001")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi002")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi003")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddvi007")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddvi008")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddvi009")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi010")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddvi011")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dddvi011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi012")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dddvi012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddvi013")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi014")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi015")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi016")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi017")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi018")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dddvi019")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("dddvi020")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi021")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi021")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dddvi022")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi022")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dddvi023")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi023")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dddvi024")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dddvi024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dddvi025")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dddvi025")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi030")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi030")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dddvi031")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi031")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dddvi032")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dddvi033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dddvi034")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi034")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dddvi035")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dddvi040")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dddvi041")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dddvi042")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dddvi043")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dddvi044")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dddvi045")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi045")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi050")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dddvi051")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dddvi052")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dddvi053")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dddvi054")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dddvi055")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dddvi060")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi060")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dddvi061")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi061")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dddvi062")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi062")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dddvi063")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi063")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dddvi064")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi064")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dddvi065")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi065")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi160")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi160")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddvi161")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi161")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi162")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi162")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dddvi163")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi163")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dddvi164")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi164")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dddvi165")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi165")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dddvi166")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi166")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dddvi167")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi167")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dddvi168")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi168")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddvi170")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dddvi170")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("dddvi171")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi171")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-100"))
	if err != nil {
		lang.SayString("dddvi172")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi172")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1000"))
	if err != nil {
		lang.SayString("dddvi173")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi173")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-10000"))
	if err != nil {
		lang.SayString("dddvi174")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi174")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-100000"))
	if err != nil {
		lang.SayString("dddvi175")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi175")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1000000"))
	if err != nil {
		lang.SayString("dddvi176")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi176")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-10000000"))
	if err != nil {
		lang.SayString("dddvi177")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi177")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-100000000"))
	if err != nil {
		lang.SayString("dddvi178")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi178")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi180")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dddvi180")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddvi181")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi181")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi182")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi182")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dddvi183")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi183")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dddvi184")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi184")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dddvi185")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi185")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dddvi186")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi186")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dddvi187")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi187")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dddvi188")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi188")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddvi190")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi190")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("dddvi191")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi191")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-100"))
	if err != nil {
		lang.SayString("dddvi192")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi192")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1000"))
	if err != nil {
		lang.SayString("dddvi193")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi193")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-10000"))
	if err != nil {
		lang.SayString("dddvi194")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi194")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-100000"))
	if err != nil {
		lang.SayString("dddvi195")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi195")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1000000"))
	if err != nil {
		lang.SayString("dddvi196")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi196")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-10000000"))
	if err != nil {
		lang.SayString("dddvi197")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi197")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-100000000"))
	if err != nil {
		lang.SayString("dddvi198")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi198")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi070")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dddvi070")
		}
	}
	rexsult, err = lang.RxFromString("999999999.4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi071")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dddvi071")
		}
	}
	rexsult, err = lang.RxFromString("999999999.5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi072")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dddvi072")
		}
	}
	rexsult, err = lang.RxFromString("999999999.9").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi073")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dddvi073")
		}
	}
	rexsult, err = lang.RxFromString("999999999.999").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi074")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dddvi074")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi090")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi093")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi094")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi100")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi100")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi101")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi101")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddvi102")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi102")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dddvi103")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi103")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dddvi104")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi104")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("dddvi105")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi105")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dddvi106")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi106")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dddvi107")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi107")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dddvi108")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi108")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddvi109")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi109")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi110")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi110")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi111")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi111")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi112")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dddvi112")
		}
	}
	rexsult, err = lang.RxFromString("4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi113")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("dddvi113")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi114")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("dddvi114")
		}
	}
	rexsult, err = lang.RxFromString("6").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi115")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dddvi115")
		}
	}
	rexsult, err = lang.RxFromString("7").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi116")
	} else {
		if !(rexsult.ToString() == "7") {
			lang.SayString("dddvi116")
		}
	}
	rexsult, err = lang.RxFromString("8").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi117")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("dddvi117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi118")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dddvi118")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi119")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dddvi119")
		}
	}
	rexsult, err = lang.RxFromString("101.3").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi131")
	} else {
		if !(rexsult.ToString() == "101") {
			lang.SayString("dddvi131")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi132")
	} else {
		if !(rexsult.ToString() == "101") {
			lang.SayString("dddvi132")
		}
	}
	rexsult, err = lang.RxFromString("101.3").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddvi133")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("dddvi133")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddvi134")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("dddvi134")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi135")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi135")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi136")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi136")
		}
	}
	rexsult, err = lang.RxFromString("18").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("18"))
	if err != nil {
		lang.SayString("dddvi137")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi137")
		}
	}
	rexsult, err = lang.RxFromString("1120").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dddvi138")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi138")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi139")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi139")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi140")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi140")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dddvi141")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi141")
		}
	}
	rexsult, err = lang.RxFromString("8.005").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dddvi142")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddvi142")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi143")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddvi143")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi144")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi144")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi145")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi145")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("dddvi150")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("dddvi150")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("dddvi151")
	} else {
		if !(rexsult.ToString() == "2473") {
			lang.SayString("dddvi151")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("dddvi152")
	} else {
		if !(rexsult.ToString() == "2519") {
			lang.SayString("dddvi152")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dddvi153")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("dddvi153")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("dddvi154")
	} else {
		if !(rexsult.ToString() == "2420") {
			lang.SayString("dddvi154")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("dddvi155")
	} else {
		if !(rexsult.ToString() == "2464") {
			lang.SayString("dddvi155")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("dddvi156")
	} else {
		if !(rexsult.ToString() == "2468") {
			lang.SayString("dddvi156")
		}
	}
	rexsult, err = lang.RxFromString("101").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("7.6"))
	if err != nil {
		lang.SayString("dddvi157")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("dddvi157")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi301")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi301")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dddvi302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi302")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("dddvi303")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi303")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dddvi304")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi304")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dddvi305")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi305")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi306")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi306")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddvi307")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi307")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dddvi308")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi308")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddvi309")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi309")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dddvi310")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi310")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddvi311")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi311")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dddvi312")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi312")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("dddvi313")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi313")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dddvi314")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi314")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dddvi315")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi315")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddvi330")
	} else {
		if !(rexsult.ToString() == "123456789012345") {
			lang.SayString("dddvi330")
		}
	}
	rexsult, err = lang.RxFromString("1234567890123456").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi331")
	} else {
		if !(rexsult.ToString() == "1234567890123456") {
			lang.SayString("dddvi331")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("dddvi1055")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1055")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("dddvi1056")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1056")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("dddvi1057")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1057")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("dddvi1058")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1058")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+101"))
	if err != nil {
		lang.SayString("dddvi1060")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1060")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+102"))
	if err != nil {
		lang.SayString("dddvi1061")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1061")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+103"))
	if err != nil {
		lang.SayString("dddvi1062")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1062")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+104"))
	if err != nil {
		lang.SayString("dddvi1063")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1063")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+105"))
	if err != nil {
		lang.SayString("dddvi1064")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1064")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+106"))
	if err != nil {
		lang.SayString("dddvi1065")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1065")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+107"))
	if err != nil {
		lang.SayString("dddvi1066")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1066")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+108"))
	if err != nil {
		lang.SayString("dddvi1067")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1067")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+109"))
	if err != nil {
		lang.SayString("dddvi1068")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1068")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+110"))
	if err != nil {
		lang.SayString("dddvi1069")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1069")
		}
	}
	rexsult, err = lang.RxFromString("1.0000E-394").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi1101")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1101")
		}
	}
	rexsult, err = lang.RxFromString("1.000E-394").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+1"))
	if err != nil {
		lang.SayString("dddvi1102")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1102")
		}
	}
	rexsult, err = lang.RxFromString("1.00E-394").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+2"))
	if err != nil {
		lang.SayString("dddvi1103")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1103")
		}
	}
	rexsult, err = lang.RxFromString("1E-394").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddvi1118")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1118")
		}
	}
	rexsult, err = lang.RxFromString("3E-394").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dddvi1119")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1119")
		}
	}
	rexsult, err = lang.RxFromString("5E-394").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dddvi1120")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1120")
		}
	}
	rexsult, err = lang.RxFromString("1E-394").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1e+4"))
	if err != nil {
		lang.SayString("dddvi1124")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1124")
		}
	}
	rexsult, err = lang.RxFromString("3.0E-394").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dddvi1130")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1130")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+200"))
	if err != nil {
		lang.SayString("dddvi1131")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1131")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+199"))
	if err != nil {
		lang.SayString("dddvi1132")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1132")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1e+198"))
	if err != nil {
		lang.SayString("dddvi1133")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1133")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-199").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("2e+198"))
	if err != nil {
		lang.SayString("dddvi1134")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1134")
		}
	}
	rexsult, err = lang.RxFromString("4.0E-199").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("4e+198"))
	if err != nil {
		lang.SayString("dddvi1135")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi1135")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi401")
	} else {
		if !(rexsult.ToString() == "123456780") {
			lang.SayString("dddvi401")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dddvi402")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi402")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddvi403")
	} else {
		if !(rexsult.ToString() == "123456780") {
			lang.SayString("dddvi403")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dddvi404")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi404")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddvi405")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi405")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dddvi406")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi406")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddvi407")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi407")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dddvi408")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi408")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi409")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi409")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dddvi410")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi410")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddvi411")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi411")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dddvi412")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi412")
		}
	}
	rexsult, err = lang.RxFromString("12345678948").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi413")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi413")
		}
	}
	rexsult, err = lang.RxFromString("12345678949").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi414")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi414")
		}
	}
	rexsult, err = lang.RxFromString("12345678950").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi415")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi415")
		}
	}
	rexsult, err = lang.RxFromString("12345678951").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi416")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi416")
		}
	}
	rexsult, err = lang.RxFromString("12345678999").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddvi417")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dddvi417")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi441")
	} else {
		if !(rexsult.ToString() == "12345678000") {
			lang.SayString("dddvi441")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dddvi442")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi442")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi443")
	} else {
		if !(rexsult.ToString() == "1234567800") {
			lang.SayString("dddvi443")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dddvi444")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi444")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi445")
	} else {
		if !(rexsult.ToString() == "1234567890") {
			lang.SayString("dddvi445")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dddvi446")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi446")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi447")
	} else {
		if !(rexsult.ToString() == "1234567891") {
			lang.SayString("dddvi447")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dddvi448")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi448")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi449")
	} else {
		if !(rexsult.ToString() == "12345678901") {
			lang.SayString("dddvi449")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dddvi450")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi450")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi451")
	} else {
		if !(rexsult.ToString() == "1234567896") {
			lang.SayString("dddvi451")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dddvi452")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi452")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("dddvi531")
	} else {
		if !(rexsult.ToString() == "5000") {
			lang.SayString("dddvi531")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddvi541")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi541")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddvi542")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi542")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi543")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi543")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi544")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi544")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddvi551")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi551")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddvi552")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi552")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi553")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi553")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddvi554")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi554")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dddvi561")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi561")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dddvi562")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi562")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddvi563")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi563")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddvi564")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi564")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dddvi571")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi571")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dddvi572")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi572")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddvi573")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi573")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddvi574")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddvi574")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix001")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix002")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix003")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dvix007")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dvix008")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dvix009")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix010")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dvix011")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dvix011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix012")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dvix012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dvix013")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix014")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix015")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix016")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix017")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix018")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dvix019")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("dvix020")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix021")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix021")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dvix022")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix022")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dvix023")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix023")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dvix024")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dvix024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dvix025")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dvix025")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix030")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix030")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dvix031")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix031")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dvix032")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dvix033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dvix034")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix034")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dvix035")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dvix040")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dvix041")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dvix042")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dvix043")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dvix044")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dvix045")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix045")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix050")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dvix051")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dvix052")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dvix053")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dvix054")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dvix055")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dvix060")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix060")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dvix061")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix061")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dvix062")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix062")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dvix063")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix063")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dvix064")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix064")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dvix065")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix065")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix160")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix160")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dvix161")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix161")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix162")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix162")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dvix163")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix163")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dvix164")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix164")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dvix165")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix165")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dvix166")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix166")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dvix167")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix167")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dvix168")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix168")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dvix170")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dvix170")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("dvix171")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix171")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-100"))
	if err != nil {
		lang.SayString("dvix172")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix172")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1000"))
	if err != nil {
		lang.SayString("dvix173")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix173")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-10000"))
	if err != nil {
		lang.SayString("dvix174")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix174")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-100000"))
	if err != nil {
		lang.SayString("dvix175")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix175")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1000000"))
	if err != nil {
		lang.SayString("dvix176")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix176")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-10000000"))
	if err != nil {
		lang.SayString("dvix177")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix177")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-100000000"))
	if err != nil {
		lang.SayString("dvix178")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix178")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix180")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dvix180")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dvix181")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix181")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix182")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix182")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dvix183")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix183")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dvix184")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix184")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dvix185")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix185")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dvix186")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix186")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dvix187")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix187")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dvix188")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix188")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dvix190")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix190")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("dvix191")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix191")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-100"))
	if err != nil {
		lang.SayString("dvix192")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix192")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1000"))
	if err != nil {
		lang.SayString("dvix193")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix193")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-10000"))
	if err != nil {
		lang.SayString("dvix194")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix194")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-100000"))
	if err != nil {
		lang.SayString("dvix195")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix195")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1000000"))
	if err != nil {
		lang.SayString("dvix196")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix196")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-10000000"))
	if err != nil {
		lang.SayString("dvix197")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix197")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-100000000"))
	if err != nil {
		lang.SayString("dvix198")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix198")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix070")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dvix070")
		}
	}
	rexsult, err = lang.RxFromString("999999999.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix071")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dvix071")
		}
	}
	rexsult, err = lang.RxFromString("999999999.5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix072")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dvix072")
		}
	}
	rexsult, err = lang.RxFromString("999999999.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix073")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dvix073")
		}
	}
	rexsult, err = lang.RxFromString("999999999.999").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix074")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dvix074")
		}
	}
	rexsult, err = lang.RxFromString("999999").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix083")
	} else {
		if !(rexsult.ToString() == "999999") {
			lang.SayString("dvix083")
		}
	}
	rexsult, err = lang.RxFromString("99999").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix084")
	} else {
		if !(rexsult.ToString() == "99999") {
			lang.SayString("dvix084")
		}
	}
	rexsult, err = lang.RxFromString("9999").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix085")
	} else {
		if !(rexsult.ToString() == "9999") {
			lang.SayString("dvix085")
		}
	}
	rexsult, err = lang.RxFromString("999").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix086")
	} else {
		if !(rexsult.ToString() == "999") {
			lang.SayString("dvix086")
		}
	}
	rexsult, err = lang.RxFromString("99").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix087")
	} else {
		if !(rexsult.ToString() == "99") {
			lang.SayString("dvix087")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix088")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dvix088")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix090")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix093")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix094")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix100")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix100")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix101")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix101")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dvix102")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix102")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dvix103")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix103")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dvix104")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix104")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("dvix105")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix105")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dvix106")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix106")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dvix107")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix107")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dvix108")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix108")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dvix109")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix109")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix110")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix110")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix111")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix111")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix112")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dvix112")
		}
	}
	rexsult, err = lang.RxFromString("4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix113")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("dvix113")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix114")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("dvix114")
		}
	}
	rexsult, err = lang.RxFromString("6").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix115")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dvix115")
		}
	}
	rexsult, err = lang.RxFromString("7").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix116")
	} else {
		if !(rexsult.ToString() == "7") {
			lang.SayString("dvix116")
		}
	}
	rexsult, err = lang.RxFromString("8").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix117")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("dvix117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix118")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dvix118")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix119")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dvix119")
		}
	}
	rexsult, err = lang.RxFromString("101.3").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix131")
	} else {
		if !(rexsult.ToString() == "101") {
			lang.SayString("dvix131")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix132")
	} else {
		if !(rexsult.ToString() == "101") {
			lang.SayString("dvix132")
		}
	}
	rexsult, err = lang.RxFromString("101.3").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dvix133")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("dvix133")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dvix134")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("dvix134")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix135")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix135")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix136")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix136")
		}
	}
	rexsult, err = lang.RxFromString("18").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("18"))
	if err != nil {
		lang.SayString("dvix137")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix137")
		}
	}
	rexsult, err = lang.RxFromString("1120").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dvix138")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix138")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix139")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix139")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix140")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix140")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dvix141")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix141")
		}
	}
	rexsult, err = lang.RxFromString("8.005").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dvix142")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix142")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix143")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dvix143")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix144")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix144")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix145")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix145")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("dvix150")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("dvix150")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("dvix151")
	} else {
		if !(rexsult.ToString() == "2473") {
			lang.SayString("dvix151")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("dvix152")
	} else {
		if !(rexsult.ToString() == "2519") {
			lang.SayString("dvix152")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dvix153")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("dvix153")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("dvix154")
	} else {
		if !(rexsult.ToString() == "2420") {
			lang.SayString("dvix154")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("dvix155")
	} else {
		if !(rexsult.ToString() == "2464") {
			lang.SayString("dvix155")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("dvix156")
	} else {
		if !(rexsult.ToString() == "2468") {
			lang.SayString("dvix156")
		}
	}
	rexsult, err = lang.RxFromString("101").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7.6"))
	if err != nil {
		lang.SayString("dvix157")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("dvix157")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix301")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix301")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dvix302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix302")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("dvix303")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix303")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dvix304")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix304")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dvix305")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix305")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix306")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix306")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dvix307")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix307")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dvix308")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix308")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dvix309")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix309")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dvix310")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix310")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dvix311")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix311")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dvix312")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix312")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("dvix313")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix313")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dvix314")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix314")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dvix315")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix315")
		}
	}
	rexsult, err = lang.RxFromString("+1.23456789012345E-0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9E+999999999"))
	if err != nil {
		lang.SayString("dvix330")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix330")
		}
	}
	rexsult, err = lang.RxFromString("+0.100").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9E+999999999"))
	if err != nil {
		lang.SayString("dvix332")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix332")
		}
	}
	rexsult, err = lang.RxFromString("9E-999999999").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("+9.100"))
	if err != nil {
		lang.SayString("dvix333")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix333")
		}
	}
	rexsult, err = lang.RxFromString("-1.23456789012345E-0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9E+999999999"))
	if err != nil {
		lang.SayString("dvix335")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix335")
		}
	}
	rexsult, err = lang.RxFromString("-0.100").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9E+999999999"))
	if err != nil {
		lang.SayString("dvix337")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix337")
		}
	}
	rexsult, err = lang.RxFromString("9E-999999999").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-9.100"))
	if err != nil {
		lang.SayString("dvix338")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix338")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix401")
	} else {
		if !(rexsult.ToString() == "123456780") {
			lang.SayString("dvix401")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dvix402")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix402")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dvix403")
	} else {
		if !(rexsult.ToString() == "123456780") {
			lang.SayString("dvix403")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dvix404")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix404")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dvix405")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix405")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dvix406")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix406")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dvix407")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix407")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dvix408")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix408")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix409")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix409")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dvix410")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix410")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dvix411")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix411")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dvix412")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix412")
		}
	}
	rexsult, err = lang.RxFromString("12345678948").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix413")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix413")
		}
	}
	rexsult, err = lang.RxFromString("12345678949").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix414")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix414")
		}
	}
	rexsult, err = lang.RxFromString("12345678950").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix415")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix415")
		}
	}
	rexsult, err = lang.RxFromString("12345678951").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix416")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix416")
		}
	}
	rexsult, err = lang.RxFromString("12345678999").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dvix417")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dvix417")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix441")
	} else {
		if !(rexsult.ToString() == "12345678000") {
			lang.SayString("dvix441")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dvix442")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix442")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix443")
	} else {
		if !(rexsult.ToString() == "1234567800") {
			lang.SayString("dvix443")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dvix444")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix444")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix445")
	} else {
		if !(rexsult.ToString() == "1234567890") {
			lang.SayString("dvix445")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dvix446")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix446")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix447")
	} else {
		if !(rexsult.ToString() == "1234567891") {
			lang.SayString("dvix447")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dvix448")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix448")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix449")
	} else {
		if !(rexsult.ToString() == "12345678901") {
			lang.SayString("dvix449")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dvix450")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix450")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix451")
	} else {
		if !(rexsult.ToString() == "1234567896") {
			lang.SayString("dvix451")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(15), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dvix452")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix452")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("dvix531")
	} else {
		if !(rexsult.ToString() == "5000") {
			lang.SayString("dvix531")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dvix541")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix541")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dvix542")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix542")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix543")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix543")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix544")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix544")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dvix551")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix551")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dvix552")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix552")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix553")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix553")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix554")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix554")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dvix561")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix561")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dvix562")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix562")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dvix563")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix563")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dvix564")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix564")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dvix571")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix571")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dvix572")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix572")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dvix573")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix573")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dvix574")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix574")
		}
	}
	rexsult, err = lang.RxFromString("100000").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix723")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("dvix723")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix724")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("dvix724")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix725")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("dvix725")
		}
	}
	rexsult, err = lang.RxFromString("100").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix726")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("dvix726")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix727")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dvix727")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dvix728")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix728")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(6), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dvix729")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix729")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi001")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi002")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi003")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdvi007")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdvi008")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdvi009")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi010")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdvi011")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dqdvi011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi012")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dqdvi012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdvi013")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi014")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi015")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi016")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi017")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi018")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dqdvi019")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("dqdvi020")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi021")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi021")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqdvi022")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi022")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dqdvi023")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi023")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dqdvi024")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dqdvi024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dqdvi025")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dqdvi025")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi030")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi030")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqdvi031")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi031")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqdvi032")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dqdvi033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dqdvi034")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi034")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dqdvi035")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqdvi040")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dqdvi041")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dqdvi042")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dqdvi043")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dqdvi044")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dqdvi045")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi045")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi050")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqdvi051")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqdvi052")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dqdvi053")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dqdvi054")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dqdvi055")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqdvi060")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi060")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dqdvi061")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi061")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dqdvi062")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi062")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dqdvi063")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi063")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dqdvi064")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi064")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dqdvi065")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi065")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi160")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi160")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdvi161")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi161")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi162")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi162")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqdvi163")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi163")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dqdvi164")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi164")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dqdvi165")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi165")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dqdvi166")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi166")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dqdvi167")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi167")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dqdvi168")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi168")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdvi170")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqdvi170")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("dqdvi171")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi171")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-100"))
	if err != nil {
		lang.SayString("dqdvi172")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi172")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1000"))
	if err != nil {
		lang.SayString("dqdvi173")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi173")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-10000"))
	if err != nil {
		lang.SayString("dqdvi174")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi174")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-100000"))
	if err != nil {
		lang.SayString("dqdvi175")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi175")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1000000"))
	if err != nil {
		lang.SayString("dqdvi176")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi176")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-10000000"))
	if err != nil {
		lang.SayString("dqdvi177")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi177")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-100000000"))
	if err != nil {
		lang.SayString("dqdvi178")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi178")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi180")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqdvi180")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdvi181")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi181")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi182")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi182")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqdvi183")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi183")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dqdvi184")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi184")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dqdvi185")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi185")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dqdvi186")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi186")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dqdvi187")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi187")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dqdvi188")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi188")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdvi190")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi190")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("dqdvi191")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi191")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-100"))
	if err != nil {
		lang.SayString("dqdvi192")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi192")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1000"))
	if err != nil {
		lang.SayString("dqdvi193")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi193")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-10000"))
	if err != nil {
		lang.SayString("dqdvi194")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi194")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-100000"))
	if err != nil {
		lang.SayString("dqdvi195")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi195")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1000000"))
	if err != nil {
		lang.SayString("dqdvi196")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi196")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-10000000"))
	if err != nil {
		lang.SayString("dqdvi197")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi197")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-100000000"))
	if err != nil {
		lang.SayString("dqdvi198")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi198")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi070")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dqdvi070")
		}
	}
	rexsult, err = lang.RxFromString("999999999.4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi071")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dqdvi071")
		}
	}
	rexsult, err = lang.RxFromString("999999999.5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi072")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dqdvi072")
		}
	}
	rexsult, err = lang.RxFromString("999999999.9").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi073")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dqdvi073")
		}
	}
	rexsult, err = lang.RxFromString("999999999.999").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi074")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dqdvi074")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi090")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi093")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi094")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi100")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi100")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi101")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi101")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdvi102")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi102")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqdvi103")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi103")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqdvi104")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi104")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("dqdvi105")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi105")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dqdvi106")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi106")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqdvi107")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi107")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dqdvi108")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi108")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdvi109")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi109")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi110")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi110")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi111")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi111")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi112")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqdvi112")
		}
	}
	rexsult, err = lang.RxFromString("4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi113")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("dqdvi113")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi114")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("dqdvi114")
		}
	}
	rexsult, err = lang.RxFromString("6").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi115")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dqdvi115")
		}
	}
	rexsult, err = lang.RxFromString("7").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi116")
	} else {
		if !(rexsult.ToString() == "7") {
			lang.SayString("dqdvi116")
		}
	}
	rexsult, err = lang.RxFromString("8").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi117")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("dqdvi117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi118")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dqdvi118")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi119")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqdvi119")
		}
	}
	rexsult, err = lang.RxFromString("101.3").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi131")
	} else {
		if !(rexsult.ToString() == "101") {
			lang.SayString("dqdvi131")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi132")
	} else {
		if !(rexsult.ToString() == "101") {
			lang.SayString("dqdvi132")
		}
	}
	rexsult, err = lang.RxFromString("101.3").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdvi133")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("dqdvi133")
		}
	}
	rexsult, err = lang.RxFromString("101.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdvi134")
	} else {
		if !(rexsult.ToString() == "33") {
			lang.SayString("dqdvi134")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi135")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi135")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi136")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi136")
		}
	}
	rexsult, err = lang.RxFromString("18").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("18"))
	if err != nil {
		lang.SayString("dqdvi137")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi137")
		}
	}
	rexsult, err = lang.RxFromString("1120").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqdvi138")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi138")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi139")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi139")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi140")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi140")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dqdvi141")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi141")
		}
	}
	rexsult, err = lang.RxFromString("8.005").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dqdvi142")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdvi142")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi143")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdvi143")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi144")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi144")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi145")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi145")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("dqdvi150")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("dqdvi150")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("dqdvi151")
	} else {
		if !(rexsult.ToString() == "2473") {
			lang.SayString("dqdvi151")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("dqdvi152")
	} else {
		if !(rexsult.ToString() == "2519") {
			lang.SayString("dqdvi152")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqdvi153")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("dqdvi153")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("dqdvi154")
	} else {
		if !(rexsult.ToString() == "2420") {
			lang.SayString("dqdvi154")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("dqdvi155")
	} else {
		if !(rexsult.ToString() == "2464") {
			lang.SayString("dqdvi155")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("dqdvi156")
	} else {
		if !(rexsult.ToString() == "2468") {
			lang.SayString("dqdvi156")
		}
	}
	rexsult, err = lang.RxFromString("101").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("7.6"))
	if err != nil {
		lang.SayString("dqdvi157")
	} else {
		if !(rexsult.ToString() == "13") {
			lang.SayString("dqdvi157")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi301")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi301")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqdvi302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi302")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("dqdvi303")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi303")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dqdvi304")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi304")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dqdvi305")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi305")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi306")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi306")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdvi307")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi307")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dqdvi308")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi308")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdvi309")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi309")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dqdvi310")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi310")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdvi311")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi311")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqdvi312")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi312")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("dqdvi313")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi313")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dqdvi314")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi314")
		}
	}
	rexsult, err = lang.RxFromString("0.12").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("dqdvi315")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi315")
		}
	}
	rexsult, err = lang.RxFromString("1234567987654321987654321890123456").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdvi330")
	} else {
		if !(rexsult.ToString() == "123456798765432198765432189012345") {
			lang.SayString("dqdvi330")
		}
	}
	rexsult, err = lang.RxFromString("1234567987654321987654321890123456").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi331")
	} else {
		if !(rexsult.ToString() == "1234567987654321987654321890123456") {
			lang.SayString("dqdvi331")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("dqdvi1055")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1055")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("dqdvi1056")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1056")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("dqdvi1057")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1057")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("dqdvi1058")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1058")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+101"))
	if err != nil {
		lang.SayString("dqdvi1060")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1060")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+102"))
	if err != nil {
		lang.SayString("dqdvi1061")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1061")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+103"))
	if err != nil {
		lang.SayString("dqdvi1062")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1062")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+104"))
	if err != nil {
		lang.SayString("dqdvi1063")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1063")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+105"))
	if err != nil {
		lang.SayString("dqdvi1064")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1064")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+106"))
	if err != nil {
		lang.SayString("dqdvi1065")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1065")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+107"))
	if err != nil {
		lang.SayString("dqdvi1066")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1066")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+108"))
	if err != nil {
		lang.SayString("dqdvi1067")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1067")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+109"))
	if err != nil {
		lang.SayString("dqdvi1068")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1068")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+110"))
	if err != nil {
		lang.SayString("dqdvi1069")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1069")
		}
	}
	rexsult, err = lang.RxFromString("1.0000E-394").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi1101")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1101")
		}
	}
	rexsult, err = lang.RxFromString("1.000E-394").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+1"))
	if err != nil {
		lang.SayString("dqdvi1102")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1102")
		}
	}
	rexsult, err = lang.RxFromString("1.00E-394").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+2"))
	if err != nil {
		lang.SayString("dqdvi1103")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1103")
		}
	}
	rexsult, err = lang.RxFromString("1E-394").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdvi1118")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1118")
		}
	}
	rexsult, err = lang.RxFromString("3E-394").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dqdvi1119")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1119")
		}
	}
	rexsult, err = lang.RxFromString("5E-394").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dqdvi1120")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1120")
		}
	}
	rexsult, err = lang.RxFromString("1E-394").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1e+4"))
	if err != nil {
		lang.SayString("dqdvi1124")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1124")
		}
	}
	rexsult, err = lang.RxFromString("3.0E-394").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dqdvi1130")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1130")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+200"))
	if err != nil {
		lang.SayString("dqdvi1131")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1131")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+199"))
	if err != nil {
		lang.SayString("dqdvi1132")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1132")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1e+198"))
	if err != nil {
		lang.SayString("dqdvi1133")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1133")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-199").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("2e+198"))
	if err != nil {
		lang.SayString("dqdvi1134")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1134")
		}
	}
	rexsult, err = lang.RxFromString("4.0E-199").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("4e+198"))
	if err != nil {
		lang.SayString("dqdvi1135")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi1135")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi401")
	} else {
		if !(rexsult.ToString() == "123456780") {
			lang.SayString("dqdvi401")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dqdvi402")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi402")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdvi403")
	} else {
		if !(rexsult.ToString() == "123456780") {
			lang.SayString("dqdvi403")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dqdvi404")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi404")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdvi405")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi405")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dqdvi406")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi406")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdvi407")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi407")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dqdvi408")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi408")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi409")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi409")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dqdvi410")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi410")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdvi411")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi411")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dqdvi412")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi412")
		}
	}
	rexsult, err = lang.RxFromString("12345678948").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi413")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi413")
		}
	}
	rexsult, err = lang.RxFromString("12345678949").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi414")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi414")
		}
	}
	rexsult, err = lang.RxFromString("12345678950").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi415")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi415")
		}
	}
	rexsult, err = lang.RxFromString("12345678951").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi416")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi416")
		}
	}
	rexsult, err = lang.RxFromString("12345678999").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdvi417")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("dqdvi417")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi441")
	} else {
		if !(rexsult.ToString() == "12345678000") {
			lang.SayString("dqdvi441")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dqdvi442")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi442")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi443")
	} else {
		if !(rexsult.ToString() == "1234567800") {
			lang.SayString("dqdvi443")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dqdvi444")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi444")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi445")
	} else {
		if !(rexsult.ToString() == "1234567890") {
			lang.SayString("dqdvi445")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dqdvi446")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi446")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi447")
	} else {
		if !(rexsult.ToString() == "1234567891") {
			lang.SayString("dqdvi447")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dqdvi448")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi448")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi449")
	} else {
		if !(rexsult.ToString() == "12345678901") {
			lang.SayString("dqdvi449")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dqdvi450")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi450")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi451")
	} else {
		if !(rexsult.ToString() == "1234567896") {
			lang.SayString("dqdvi451")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dqdvi452")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi452")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("dqdvi531")
	} else {
		if !(rexsult.ToString() == "5000") {
			lang.SayString("dqdvi531")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdvi541")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi541")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdvi542")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi542")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi543")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi543")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi544")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi544")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdvi551")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi551")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdvi552")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi552")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi553")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi553")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdvi554")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi554")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqdvi561")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi561")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqdvi562")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi562")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdvi563")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi563")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdvi564")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi564")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqdvi571")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi571")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqdvi572")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi572")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdvi573")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi573")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdvi574")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi574")
		}
	}
	rexsult, err = lang.RxFromString("8.336804418094040989630006819881709E-6143").OpDivI(lang.RxSetWithDigit(34), lang.RxFromString("8.336804418094040989630006819889000E-6143"))
	if err != nil {
		lang.SayString("dqdvi700")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdvi700")
		}
	}
	rexsult, err = lang.RxFromString("4953734675913.065314738743322579").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("0218.932010396534371704930714860E+797"))
	if err != nil {
		lang.SayString("dvix3001")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3001")
		}
	}
	rexsult, err = lang.RxFromString("9641.684323386955881595490347910E-844").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-78864532047.12287484430980636798E+934"))
	if err != nil {
		lang.SayString("dvix3002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3002")
		}
	}
	rexsult, err = lang.RxFromString("479084.8561808930525417735205519").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("084157571054.2691784660983989931"))
	if err != nil {
		lang.SayString("dvix3004")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3004")
		}
	}
	rexsult, err = lang.RxFromString("-0363750788.573782205664349562931").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-3172.080934464133691909905980096"))
	if err != nil {
		lang.SayString("dvix3005")
	} else {
		if !(rexsult.ToString() == "114672") {
			lang.SayString("dvix3005")
		}
	}
	rexsult, err = lang.RxFromString("1381026551423669919010191878449").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-82.66614775445371254999357800739"))
	if err != nil {
		lang.SayString("dvix3006")
	} else {
		if !(rexsult.ToString() == "-16706071214613552377376639557") {
			lang.SayString("dvix3006")
		}
	}
	rexsult, err = lang.RxFromString("4627.026960423072127953556635585").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-4410583132901.830017479741231131"))
	if err != nil {
		lang.SayString("dvix3007")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3007")
		}
	}
	rexsult, err = lang.RxFromString("75353574493.84484153484918212042").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-8684111695095849922263044191221"))
	if err != nil {
		lang.SayString("dvix3008")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3008")
		}
	}
	rexsult, err = lang.RxFromString("6907058.216435355874729592373011").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("2.857005446917670515662398741545"))
	if err != nil {
		lang.SayString("dvix3009")
	} else {
		if !(rexsult.ToString() == "2417586") {
			lang.SayString("dvix3009")
		}
	}
	rexsult, err = lang.RxFromString("-0708069.025667471996378081482549").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-562842.4701520787831018732202804"))
	if err != nil {
		lang.SayString("dvix3011")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dvix3011")
		}
	}
	rexsult, err = lang.RxFromString("4055087.246994644709729942673976").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-43183146921897.67383476104084575E+211"))
	if err != nil {
		lang.SayString("dvix3012")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3012")
		}
	}
	rexsult, err = lang.RxFromString("4502895892520.396581348110906909E-512").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-815.9047305921862348263521876034"))
	if err != nil {
		lang.SayString("dvix3013")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3013")
		}
	}
	rexsult, err = lang.RxFromString("467.6721295072628100260239179865").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-02.07155073395573569852316073025"))
	if err != nil {
		lang.SayString("dvix3014")
	} else {
		if !(rexsult.ToString() == "-225") {
			lang.SayString("dvix3014")
		}
	}
	rexsult, err = lang.RxFromString("2.156795313311150143949997552501E-571").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-8677147.586389401682712180146855"))
	if err != nil {
		lang.SayString("dvix3015")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3015")
		}
	}
	rexsult, err = lang.RxFromString("-974953.2801637208368002585822457").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-693095793.3667578067802698191246"))
	if err != nil {
		lang.SayString("dvix3016")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3016")
		}
	}
	rexsult, err = lang.RxFromString("262273.0222851186523650889896428E-624").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("74177.21073338090843145838835480"))
	if err != nil {
		lang.SayString("dvix3018")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3018")
		}
	}
	rexsult, err = lang.RxFromString("-8036052748815903177624716581732").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-066677357.4438809548850966167573"))
	if err != nil {
		lang.SayString("dvix3019")
	} else {
		if !(rexsult.ToString() == "120521464210387351732732") {
			lang.SayString("dvix3019")
		}
	}
	rexsult, err = lang.RxFromString("24791301060.37938360567775506973").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-5613327866480.322649080205877564"))
	if err != nil {
		lang.SayString("dvix3021")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3021")
		}
	}
	rexsult, err = lang.RxFromString("-930711443.9474781586162910776139").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-740.3860979292775472622798348030"))
	if err != nil {
		lang.SayString("dvix3022")
	} else {
		if !(rexsult.ToString() == "1257062") {
			lang.SayString("dvix3022")
		}
	}
	rexsult, err = lang.RxFromString("2358276428765.064191082773385539").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("214.3589796082328665878602304469"))
	if err != nil {
		lang.SayString("dvix3023")
	} else {
		if !(rexsult.ToString() == "11001528525") {
			lang.SayString("dvix3023")
		}
	}
	rexsult, err = lang.RxFromString("140422069.5863246490180206814374E-447").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-567195652586.2454217069003186487"))
	if err != nil {
		lang.SayString("dvix3025")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3025")
		}
	}
	rexsult, err = lang.RxFromString("75929096475.63450425339472559646E+153").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-0945260193.503803519572604151290E+459"))
	if err != nil {
		lang.SayString("dvix3026")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3026")
		}
	}
	rexsult, err = lang.RxFromString("6312318309.142044953357460463732").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-5641317823.202274083982487558514E+628"))
	if err != nil {
		lang.SayString("dvix3027")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3027")
		}
	}
	rexsult, err = lang.RxFromString("98471198160.56524417578665886060").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-23994.14313393939743548945165462"))
	if err != nil {
		lang.SayString("dvix3029")
	} else {
		if !(rexsult.ToString() == "-4103968") {
			lang.SayString("dvix3029")
		}
	}
	rexsult, err = lang.RxFromString("329326552.0208398002250836592043").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-02451.10065397010591546041034041"))
	if err != nil {
		lang.SayString("dvix3030")
	} else {
		if !(rexsult.ToString() == "-134358") {
			lang.SayString("dvix3030")
		}
	}
	rexsult, err = lang.RxFromString("-92980.68431371090354435763218439").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-2282178507046019721925800997065"))
	if err != nil {
		lang.SayString("dvix3031")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3031")
		}
	}
	rexsult, err = lang.RxFromString("37.27457578793521166809739140081").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-392550.4790095035979998355569916"))
	if err != nil {
		lang.SayString("dvix3033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3033")
		}
	}
	rexsult, err = lang.RxFromString("-2787.980590304199878755265273703").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("7117631179305319208210387565324"))
	if err != nil {
		lang.SayString("dvix3034")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3034")
		}
	}
	rexsult, err = lang.RxFromString("3944570323.331121750661920475191").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-17360722.28878145641394962484366"))
	if err != nil {
		lang.SayString("dvix3036")
	} else {
		if !(rexsult.ToString() == "-227") {
			lang.SayString("dvix3036")
		}
	}
	rexsult, err = lang.RxFromString("19544.14018503427029002552872707").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("1786697762.885178994182133839546"))
	if err != nil {
		lang.SayString("dvix3037")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3037")
		}
	}
	rexsult, err = lang.RxFromString("-05.75485957937617757983513662981").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("5564476875.989640431173694372083"))
	if err != nil {
		lang.SayString("dvix3038")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3038")
		}
	}
	rexsult, err = lang.RxFromString("-4208820.898718069194008526302746").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("626887.7553774705678201112845462E+206"))
	if err != nil {
		lang.SayString("dvix3039")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3039")
		}
	}
	rexsult, err = lang.RxFromString("-442941.7541811527940918244383454").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-068126768.0563559819156379151016"))
	if err != nil {
		lang.SayString("dvix3041")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3041")
		}
	}
	rexsult, err = lang.RxFromString("-040726778711.8677615616711676159").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("299691.9430345259174614997064916"))
	if err != nil {
		lang.SayString("dvix3042")
	} else {
		if !(rexsult.ToString() == "-135895") {
			lang.SayString("dvix3042")
		}
	}
	rexsult, err = lang.RxFromString("-1934197520.738366912179143085955").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("3.810751422515178400293693371519"))
	if err != nil {
		lang.SayString("dvix3043")
	} else {
		if !(rexsult.ToString() == "-507563287") {
			lang.SayString("dvix3043")
		}
	}
	rexsult, err = lang.RxFromString("813262.7723533833038829559646830").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-303284822716.8282178131118185907"))
	if err != nil {
		lang.SayString("dvix3044")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3044")
		}
	}
	rexsult, err = lang.RxFromString("-075537177538.1814516621962185490").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("26980775255.51542856483122484898"))
	if err != nil {
		lang.SayString("dvix3046")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dvix3046")
		}
	}
	rexsult, err = lang.RxFromString("-6468.903738522951259063099946195").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-7877.324314273694312164407794939E+267"))
	if err != nil {
		lang.SayString("dvix3048")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3048")
		}
	}
	rexsult, err = lang.RxFromString("-9567221.183663236817239254783372E-203").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("1650.198961256061165362319471264"))
	if err != nil {
		lang.SayString("dvix3049")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3049")
		}
	}
	rexsult, err = lang.RxFromString("8812306098770.200752139142033569E-428").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("26790.17380163975186972720427030E+568"))
	if err != nil {
		lang.SayString("dvix3050")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3050")
		}
	}
	rexsult, err = lang.RxFromString("80108033.12724838718736922500904").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-706207255092.7645192310078892869"))
	if err != nil {
		lang.SayString("dvix3051")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3051")
		}
	}
	rexsult, err = lang.RxFromString("-37942846282.76101663789059003505").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-5.649456053942850351313869983197"))
	if err != nil {
		lang.SayString("dvix3052")
	} else {
		if !(rexsult.ToString() == "6716194607") {
			lang.SayString("dvix3052")
		}
	}
	rexsult, err = lang.RxFromString("2838948.589837595494152150647194").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("569547026247.5469563701415715960"))
	if err != nil {
		lang.SayString("dvix3054")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3054")
		}
	}
	rexsult, err = lang.RxFromString("-57131573677452.15449921725097290").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("4669681430736.326858508715643769"))
	if err != nil {
		lang.SayString("dvix3056")
	} else {
		if !(rexsult.ToString() == "-12") {
			lang.SayString("dvix3056")
		}
	}
	rexsult, err = lang.RxFromString("90794826.55528018781830463383411").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-5.471502270351231110027647216128"))
	if err != nil {
		lang.SayString("dvix3057")
	} else {
		if !(rexsult.ToString() == "-16594131") {
			lang.SayString("dvix3057")
		}
	}
	rexsult, err = lang.RxFromString("58508794729.35191160840980489138").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-47060867.24988279680824397447551"))
	if err != nil {
		lang.SayString("dvix3058")
	} else {
		if !(rexsult.ToString() == "-1243") {
			lang.SayString("dvix3058")
		}
	}
	rexsult, err = lang.RxFromString("-746104.0768078474426464219416332E+006").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("9595418.300613754556671852801667E+385"))
	if err != nil {
		lang.SayString("dvix3059")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3059")
		}
	}
	rexsult, err = lang.RxFromString("-41214265628.83801241467317270595").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("1015336323798389903361978271354"))
	if err != nil {
		lang.SayString("dvix3061")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3061")
		}
	}
	rexsult, err = lang.RxFromString("89937.39749201095570357557430822").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("82351554210093.60879476027800331"))
	if err != nil {
		lang.SayString("dvix3062")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3062")
		}
	}
	rexsult, err = lang.RxFromString("-2647593306.528617691373470059213").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-655531558709.4582168930191014461"))
	if err != nil {
		lang.SayString("dvix3064")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3064")
		}
	}
	rexsult, err = lang.RxFromString("2904078690665765116603253099668E-329").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-71.45586619176091599264717047885E+787"))
	if err != nil {
		lang.SayString("dvix3065")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3065")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("dvix3067")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797") {
			lang.SayString("dvix3067")
		}
	}
	rexsult, err = lang.RxFromString("-84172558160661.35863831029352323").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-11271.58916600931155937291904890"))
	if err != nil {
		lang.SayString("dvix3068")
	} else {
		if !(rexsult.ToString() == "7467674426") {
			lang.SayString("dvix3068")
		}
	}
	rexsult, err = lang.RxFromString("-31823131.15691583393820628480997E-440").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("92913.91582947237200286427030028E+771"))
	if err != nil {
		lang.SayString("dvix3071")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3071")
		}
	}
	rexsult, err = lang.RxFromString("55573867888.91575330563698128150").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("599.5231614736232188354393212234"))
	if err != nil {
		lang.SayString("dvix3072")
	} else {
		if !(rexsult.ToString() == "92696782") {
			lang.SayString("dvix3072")
		}
	}
	rexsult, err = lang.RxFromString("-5447727448431680878699555714796E-800").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("5487207.142687001607026665515349E-362"))
	if err != nil {
		lang.SayString("dvix3073")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3073")
		}
	}
	rexsult, err = lang.RxFromString("0418349404834.547110239542290134").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("09819915.92405288066606124554841"))
	if err != nil {
		lang.SayString("dvix3074")
	} else {
		if !(rexsult.ToString() == "42602") {
			lang.SayString("dvix3074")
		}
	}
	rexsult, err = lang.RxFromString("-262021.7565194737396448014286436").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-7983992600094836304387324162042E+390"))
	if err != nil {
		lang.SayString("dvix3075")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3075")
		}
	}
	rexsult, err = lang.RxFromString("48696050631.42565380301204592392E-505").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-33868752339.85057267609967806187E+821"))
	if err != nil {
		lang.SayString("dvix3076")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3076")
		}
	}
	rexsult, err = lang.RxFromString("95316999.19440144356471126680708").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-60791.33805057402845885978390435"))
	if err != nil {
		lang.SayString("dvix3077")
	} else {
		if !(rexsult.ToString() == "-1567") {
			lang.SayString("dvix3077")
		}
	}
	rexsult, err = lang.RxFromString("-5326702296402708234722215224979E-136").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("8032459.450998820205916538543258"))
	if err != nil {
		lang.SayString("dvix3078")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3078")
		}
	}
	rexsult, err = lang.RxFromString("-8739299372114.092482914139281669").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("507610074.7343577029345077385838"))
	if err != nil {
		lang.SayString("dvix3080")
	} else {
		if !(rexsult.ToString() == "-17216") {
			lang.SayString("dvix3080")
		}
	}
	rexsult, err = lang.RxFromString("764578.5204849936912066033177429").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("64603.13571259164812609436832506"))
	if err != nil {
		lang.SayString("dvix3082")
	} else {
		if !(rexsult.ToString() == "11") {
			lang.SayString("dvix3082")
		}
	}
	rexsult, err = lang.RxFromString("079203.7330103777716903518367560").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("846388934347.6324036132959664705"))
	if err != nil {
		lang.SayString("dvix3083")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3083")
		}
	}
	rexsult, err = lang.RxFromString("-4278.581514688669249247007127899E-329").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("5.474973992953902631890208360829"))
	if err != nil {
		lang.SayString("dvix3084")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3084")
		}
	}
	rexsult, err = lang.RxFromString("60867019.81764798845468445196869E+651").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("6.149612565404080501157093851895E+817"))
	if err != nil {
		lang.SayString("dvix3085")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3085")
		}
	}
	rexsult, err = lang.RxFromString("18554417738217.62218590965803605E-382").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-0894505909529.052378474618435782E+527"))
	if err != nil {
		lang.SayString("dvix3086")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3086")
		}
	}
	rexsult, err = lang.RxFromString("69073355517144.36356688642213839").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("997784782535.6104634823627327033E+116"))
	if err != nil {
		lang.SayString("dvix3087")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3087")
		}
	}
	rexsult, err = lang.RxFromString("450282259072.8657099359104277477").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-1791307965314309175477911369824"))
	if err != nil {
		lang.SayString("dvix3088")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3088")
		}
	}
	rexsult, err = lang.RxFromString("954678411.7838149266455177850037").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("142988.7096204254529284334278794"))
	if err != nil {
		lang.SayString("dvix3089")
	} else {
		if !(rexsult.ToString() == "6676") {
			lang.SayString("dvix3089")
		}
	}
	rexsult, err = lang.RxFromString("-75492024.20990197005974241975449").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-14760421311348.35269044633000927"))
	if err != nil {
		lang.SayString("dvix3091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3091")
		}
	}
	rexsult, err = lang.RxFromString("317747.6972215715434186596178036E-452").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("24759763.33144824613591228097330E+092"))
	if err != nil {
		lang.SayString("dvix3092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3092")
		}
	}
	rexsult, err = lang.RxFromString("-8.153334430358647134334545353427").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-9.717872025814596548462854853522"))
	if err != nil {
		lang.SayString("dvix3093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3093")
		}
	}
	rexsult, err = lang.RxFromString("7.267345197492967332320456062961E-478").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("5054015481833.263541189916208065"))
	if err != nil {
		lang.SayString("dvix3094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3094")
		}
	}
	rexsult, err = lang.RxFromString("-1223354029.862567054230912271171").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("8135774223401322785475014855625"))
	if err != nil {
		lang.SayString("dvix3095")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3095")
		}
	}
	rexsult, err = lang.RxFromString("-4673112.663442366293812346653429").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("-3429.998403142546001438238460958"))
	if err != nil {
		lang.SayString("dvix3097")
	} else {
		if !(rexsult.ToString() == "1362") {
			lang.SayString("dvix3097")
		}
	}
	rexsult, err = lang.RxFromString("88.96492479681278079861456051103").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("386939.4621006514751889096510923E+139"))
	if err != nil {
		lang.SayString("dvix3098")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3098")
		}
	}
	rexsult, err = lang.RxFromString("064326846.4286437304788069444326E-942").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("92.23649942010862087149015091350"))
	if err != nil {
		lang.SayString("dvix3099")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3099")
		}
	}
	rexsult, err = lang.RxFromString("504507.0043949324433170405699360").OpDivI(lang.RxSetWithDigit(31), lang.RxFromString("605387.7175522955344659311072099"))
	if err != nil {
		lang.SayString("dvix3100")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3100")
		}
	}
	rexsult, err = lang.RxFromString("1.5283550163839789319142430253644").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-1.6578158484822969520405291379492"))
	if err != nil {
		lang.SayString("dvix3201")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3201")
		}
	}
	rexsult, err = lang.RxFromString("-622903030605.2867503937836507326").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("6519388607.1331855704471328795821"))
	if err != nil {
		lang.SayString("dvix3202")
	} else {
		if !(rexsult.ToString() == "-95") {
			lang.SayString("dvix3202")
		}
	}
	rexsult, err = lang.RxFromString("-5675915.2465457487632250245209054").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("73913909880.381367895205086027416"))
	if err != nil {
		lang.SayString("dvix3203")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3203")
		}
	}
	rexsult, err = lang.RxFromString("97.647321172555144900685605318635").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("4.8620911587547548751209841570885"))
	if err != nil {
		lang.SayString("dvix3204")
	} else {
		if !(rexsult.ToString() == "20") {
			lang.SayString("dvix3204")
		}
	}
	rexsult, err = lang.RxFromString("-9717253267024.5380651435435603552").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-2669.2539695193820424002013488480E+363"))
	if err != nil {
		lang.SayString("dvix3205")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3205")
		}
	}
	rexsult, err = lang.RxFromString("-4.0817391717190128506083001702335E-767").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("12772.807105920712660991033689206"))
	if err != nil {
		lang.SayString("dvix3206")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3206")
		}
	}
	rexsult, err = lang.RxFromString("68625322655934146845003028928647").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-59.634169944840280159782488098700"))
	if err != nil {
		lang.SayString("dvix3207")
	} else {
		if !(rexsult.ToString() == "-1150771826276954946844322988192") {
			lang.SayString("dvix3207")
		}
	}
	rexsult, err = lang.RxFromString("732515.76532049290815665856727641").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-92134479835821.319619827023729829"))
	if err != nil {
		lang.SayString("dvix3208")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3208")
		}
	}
	rexsult, err = lang.RxFromString("-30.458011942978338421676454778733").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-5023372024597665102336430410403E+831"))
	if err != nil {
		lang.SayString("dvix3209")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3209")
		}
	}
	rexsult, err = lang.RxFromString("-89640.094149414644660480286430462").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-58703419758.800889903227509215474"))
	if err != nil {
		lang.SayString("dvix3210")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3210")
		}
	}
	rexsult, err = lang.RxFromString("913391.42744224458216174967853722").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-21051638.816432817393202262710630E+432"))
	if err != nil {
		lang.SayString("dvix3212")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3212")
		}
	}
	rexsult, err = lang.RxFromString("-917591456829.12109027484399536567").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-28892177726858026955513438843371E+708"))
	if err != nil {
		lang.SayString("dvix3213")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3213")
		}
	}
	rexsult, err = lang.RxFromString("34938410840645.913399699219228218").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("30.818220393242402846077755480548"))
	if err != nil {
		lang.SayString("dvix3214")
	} else {
		if !(rexsult.ToString() == "1133693327999") {
			lang.SayString("dvix3214")
		}
	}
	rexsult, err = lang.RxFromString("6034.7374411022598081745006769023E-517").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("29771833428054709077850588904653"))
	if err != nil {
		lang.SayString("dvix3215")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3215")
		}
	}
	rexsult, err = lang.RxFromString("-5565747671734.1686009705574503152").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-490.30899494881071282787487030303"))
	if err != nil {
		lang.SayString("dvix3216")
	} else {
		if !(rexsult.ToString() == "11351510433") {
			lang.SayString("dvix3216")
		}
	}
	rexsult, err = lang.RxFromString("319545511.89203495546689273564728E+036").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-2955943533943321899150310192061"))
	if err != nil {
		lang.SayString("dvix3217")
	} else {
		if !(rexsult.ToString() == "-108102711781422") {
			lang.SayString("dvix3217")
		}
	}
	rexsult, err = lang.RxFromString("-36852134.84194296250843579428931").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-5830629.8347085025808716360357940"))
	if err != nil {
		lang.SayString("dvix3218")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dvix3218")
		}
	}
	rexsult, err = lang.RxFromString("8.6021905001798578659275880018221E-374").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-39505285344943.729681835377530908"))
	if err != nil {
		lang.SayString("dvix3219")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3219")
		}
	}
	rexsult, err = lang.RxFromString("-54863165.152174109720312887805017").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("736.1398476560169141105328256628"))
	if err != nil {
		lang.SayString("dvix3220")
	} else {
		if !(rexsult.ToString() == "-74528") {
			lang.SayString("dvix3220")
		}
	}
	rexsult, err = lang.RxFromString("-3263743464517851012531708810307").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("2457206.2471248382136273643208109"))
	if err != nil {
		lang.SayString("dvix3221")
	} else {
		if !(rexsult.ToString() == "-1328233422952076975055082") {
			lang.SayString("dvix3221")
		}
	}
	rexsult, err = lang.RxFromString("2856586744.0548637797291151154902E-895").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("953545637646.57694835860339582821E+080"))
	if err != nil {
		lang.SayString("dvix3222")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3222")
		}
	}
	rexsult, err = lang.RxFromString("5624157233.3433661009203529937625").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("626098409265.93738586750090160638"))
	if err != nil {
		lang.SayString("dvix3223")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3223")
		}
	}
	rexsult, err = lang.RxFromString("-213499362.91476998701834067092611").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("419272438.02555757699863022643444"))
	if err != nil {
		lang.SayString("dvix3224")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3224")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("dvix3225")
	} else {
		if !(rexsult.ToString() == "6300") {
			lang.SayString("dvix3225")
		}
	}
	rexsult, err = lang.RxFromString("-74396862273800.625679130772935550").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-115616605.52826981284183992013157"))
	if err != nil {
		lang.SayString("dvix3227")
	} else {
		if !(rexsult.ToString() == "643479") {
			lang.SayString("dvix3227")
		}
	}
	rexsult, err = lang.RxFromString("67585560.562561229497720955705979").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("826.96290288608566737177503451613"))
	if err != nil {
		lang.SayString("dvix3228")
	} else {
		if !(rexsult.ToString() == "81727") {
			lang.SayString("dvix3228")
		}
	}
	rexsult, err = lang.RxFromString("6877386868.9498051860742298735156E-232").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("390.3154289860643509393769754551"))
	if err != nil {
		lang.SayString("dvix3229")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3229")
		}
	}
	rexsult, err = lang.RxFromString("-1647335.201144609256134925838937").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-186654823782.50459802235024230856"))
	if err != nil {
		lang.SayString("dvix3230")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3230")
		}
	}
	rexsult, err = lang.RxFromString("-6.6547424012516834662011706165297").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-574454585580.06215974139884746441"))
	if err != nil {
		lang.SayString("dvix3232")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3232")
		}
	}
	rexsult, err = lang.RxFromString("-27627.758745381267780885067447671").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-23385972189441.709586433758111062"))
	if err != nil {
		lang.SayString("dvix3233")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3233")
		}
	}
	rexsult, err = lang.RxFromString("-70454070095869.70717871212601390").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-6200178.370249260009168888392406"))
	if err != nil {
		lang.SayString("dvix3237")
	} else {
		if !(rexsult.ToString() == "11363232") {
			lang.SayString("dvix3237")
		}
	}
	rexsult, err = lang.RxFromString("-5168.2214111091132913776042214525").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-5690274.0971173476527123568627720"))
	if err != nil {
		lang.SayString("dvix3239")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3239")
		}
	}
	rexsult, err = lang.RxFromString("33783.060857197067391462144517964").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-2070.0806959465088198508322815406"))
	if err != nil {
		lang.SayString("dvix3240")
	} else {
		if !(rexsult.ToString() == "-16") {
			lang.SayString("dvix3240")
		}
	}
	rexsult, err = lang.RxFromString("42207435091050.840296353874733169E-905").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("73330633078.828216018536326743325E+976"))
	if err != nil {
		lang.SayString("dvix3241")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3241")
		}
	}
	rexsult, err = lang.RxFromString("-71800.806700868784841045406679641").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-39617456964250697902519150526701"))
	if err != nil {
		lang.SayString("dvix3242")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3242")
		}
	}
	rexsult, err = lang.RxFromString("53627480801.631504892310016062160").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("328259.56947661049313311983109503"))
	if err != nil {
		lang.SayString("dvix3243")
	} else {
		if !(rexsult.ToString() == "163369") {
			lang.SayString("dvix3243")
		}
	}
	rexsult, err = lang.RxFromString("-5052601598.5559371338428368438728").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-97855372.224321664785314782556064"))
	if err != nil {
		lang.SayString("dvix3244")
	} else {
		if !(rexsult.ToString() == "51") {
			lang.SayString("dvix3244")
		}
	}
	rexsult, err = lang.RxFromString("4208134320733.7069742988228068191E+146").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("4270869.1760149477598920960628392E+471"))
	if err != nil {
		lang.SayString("dvix3245")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3245")
		}
	}
	rexsult, err = lang.RxFromString("-440220916.66716743026896931194749").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-102725.01594377871560564824358775"))
	if err != nil {
		lang.SayString("dvix3248")
	} else {
		if !(rexsult.ToString() == "4285") {
			lang.SayString("dvix3248")
		}
	}
	rexsult, err = lang.RxFromString("-46.250735086006350517943464758019").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("14656357555174.263295266074908024"))
	if err != nil {
		lang.SayString("dvix3249")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3249")
		}
	}
	rexsult, err = lang.RxFromString("96668419802749.555738010239087449E-838").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-19498732131365824921639467044927E-542"))
	if err != nil {
		lang.SayString("dvix3251")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3251")
		}
	}
	rexsult, err = lang.RxFromString("-62663404777.352508979582846931050").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("9.2570938837239134052589184917186E+916"))
	if err != nil {
		lang.SayString("dvix3253")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3253")
		}
	}
	rexsult, err = lang.RxFromString("1.744601214474560992754529320172E-827").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-17.353669504253419489494030651507E-631"))
	if err != nil {
		lang.SayString("dvix3254")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3254")
		}
	}
	rexsult, err = lang.RxFromString("0367191549036702224827734853471").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("4410320662415266533763143837742E+721"))
	if err != nil {
		lang.SayString("dvix3255")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3255")
		}
	}
	rexsult, err = lang.RxFromString("097704116.4492566721965710365073").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-96736.400939809433556067504289145"))
	if err != nil {
		lang.SayString("dvix3256")
	} else {
		if !(rexsult.ToString() == "-1010") {
			lang.SayString("dvix3256")
		}
	}
	rexsult, err = lang.RxFromString("-86863.276249466008289214762980838").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("531.50602652732088208397655484476"))
	if err != nil {
		lang.SayString("dvix3260")
	} else {
		if !(rexsult.ToString() == "-163") {
			lang.SayString("dvix3260")
		}
	}
	rexsult, err = lang.RxFromString("-40707.169006771111855573524157083").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-68795521421321853333274411827749"))
	if err != nil {
		lang.SayString("dvix3261")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3261")
		}
	}
	rexsult, err = lang.RxFromString("-4245360967593.9206771555839718158E-682").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-3.119606239042530207103203508057"))
	if err != nil {
		lang.SayString("dvix3263")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3263")
		}
	}
	rexsult, err = lang.RxFromString("-3422145405774.0800213000547612240E-023").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-60810.964656409650839011321706310"))
	if err != nil {
		lang.SayString("dvix3264")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3264")
		}
	}
	rexsult, err = lang.RxFromString("-24521811.07649485796598387627478E-661").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-94860846133404815410816234000694"))
	if err != nil {
		lang.SayString("dvix3265")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3265")
		}
	}
	rexsult, err = lang.RxFromString("-5042529937498.8944492248538951438").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("3891904674.4549170968807145612549"))
	if err != nil {
		lang.SayString("dvix3266")
	} else {
		if !(rexsult.ToString() == "-1295") {
			lang.SayString("dvix3266")
		}
	}
	rexsult, err = lang.RxFromString("-535824163.54531747646293693868651E-665").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("2732988.5891363639325008206099712"))
	if err != nil {
		lang.SayString("dvix3267")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3267")
		}
	}
	rexsult, err = lang.RxFromString("24032.702008553084252925140858134E-509").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("52864854.899420632375589206704068"))
	if err != nil {
		lang.SayString("dvix3268")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3268")
		}
	}
	rexsult, err = lang.RxFromString("71553220259.938950698030519906727E-496").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("754.44220417415325444943566016062"))
	if err != nil {
		lang.SayString("dvix3269")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3269")
		}
	}
	rexsult, err = lang.RxFromString("95.490751127249945886828257312118").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("987.01498316307365714167410690192E+133"))
	if err != nil {
		lang.SayString("dvix3272")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3272")
		}
	}
	rexsult, err = lang.RxFromString("69434850287.460788550936730296153").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-35119136549015044241569827542264"))
	if err != nil {
		lang.SayString("dvix3273")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3273")
		}
	}
	rexsult, err = lang.RxFromString("-392.22739924621965621739098725407").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-65551274.987160998195282109612136"))
	if err != nil {
		lang.SayString("dvix3274")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3274")
		}
	}
	rexsult, err = lang.RxFromString("6413265.4423561191792972085539457").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("24514.222704714139350026165721146"))
	if err != nil {
		lang.SayString("dvix3275")
	} else {
		if !(rexsult.ToString() == "261") {
			lang.SayString("dvix3275")
		}
	}
	rexsult, err = lang.RxFromString("378204716633.40024100602896057615").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-0300218714378.322231269606216516"))
	if err != nil {
		lang.SayString("dvix3277")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dvix3277")
		}
	}
	rexsult, err = lang.RxFromString("-3554.5895974968741465654022772100E-073").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("9752.0428746722497621936998533848E+516"))
	if err != nil {
		lang.SayString("dvix3279")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3279")
		}
	}
	rexsult, err = lang.RxFromString("750187685.63632608923397234391668").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("4633194252863.6730625972669246025"))
	if err != nil {
		lang.SayString("dvix3280")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3280")
		}
	}
	rexsult, err = lang.RxFromString("30190034242853.251165951457589386E-028").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("8038885676.3204238322976087799751E+018"))
	if err != nil {
		lang.SayString("dvix3281")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3281")
		}
	}
	rexsult, err = lang.RxFromString("65.537942676774715953400768803539").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("125946917260.87536506197191782198"))
	if err != nil {
		lang.SayString("dvix3282")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3282")
		}
	}
	rexsult, err = lang.RxFromString("8015272348677.5489394183881813700").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("949.23027111499779258284877920022"))
	if err != nil {
		lang.SayString("dvix3283")
	} else {
		if !(rexsult.ToString() == "8443970438") {
			lang.SayString("dvix3283")
		}
	}
	rexsult, err = lang.RxFromString("-17544189017145.710117633021800426E-537").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("292178000.06450804618299520094843"))
	if err != nil {
		lang.SayString("dvix3285")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3285")
		}
	}
	rexsult, err = lang.RxFromString("-506650.99395649907139204052441630").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("11.018427502031650148962067367158"))
	if err != nil {
		lang.SayString("dvix3286")
	} else {
		if !(rexsult.ToString() == "-45982") {
			lang.SayString("dvix3286")
		}
	}
	rexsult, err = lang.RxFromString("4846835159.5922372307656000769395E-241").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-84.001893040865864590122330800768"))
	if err != nil {
		lang.SayString("dvix3287")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3287")
		}
	}
	rexsult, err = lang.RxFromString("-35.029311013822259358116556164908").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-3994308878.1994645313414534209127"))
	if err != nil {
		lang.SayString("dvix3288")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3288")
		}
	}
	rexsult, err = lang.RxFromString("7606663750.6735265233044420887018E+166").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-5491814639.4484565418284686127552E+365"))
	if err != nil {
		lang.SayString("dvix3289")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3289")
		}
	}
	rexsult, err = lang.RxFromString("97271576094.456406729671729224827").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-1.5412563837540810793697955063295E+554"))
	if err != nil {
		lang.SayString("dvix3291")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3291")
		}
	}
	rexsult, err = lang.RxFromString("41139789894.401826915757391777563").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-1.8729920305671057957156159690823E-020"))
	if err != nil {
		lang.SayString("dvix3292")
	} else {
		if !(rexsult.ToString() == "-2196474369511625389289506461551") {
			lang.SayString("dvix3292")
		}
	}
	rexsult, err = lang.RxFromString("23777.857951114967684767609401626").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("720760.03897144157012301385227528"))
	if err != nil {
		lang.SayString("dvix3295")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3295")
		}
	}
	rexsult, err = lang.RxFromString("-21337853323980858055292469611948").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("6080272840.3071490445256786982100E+532"))
	if err != nil {
		lang.SayString("dvix3296")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3296")
		}
	}
	rexsult, err = lang.RxFromString("-818409238.0423893439849743856947").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("756.39156265972753259267069158760"))
	if err != nil {
		lang.SayString("dvix3297")
	} else {
		if !(rexsult.ToString() == "-1081991") {
			lang.SayString("dvix3297")
		}
	}
	rexsult, err = lang.RxFromString("47567380384943.87013600286155046").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("65.084709374908275826942076480326"))
	if err != nil {
		lang.SayString("dvix3298")
	} else {
		if !(rexsult.ToString() == "730853388480") {
			lang.SayString("dvix3298")
		}
	}
	rexsult, err = lang.RxFromString("-296615544.05897984545127115290251").OpDivI(lang.RxSetWithDigit(32), lang.RxFromString("-5416115.4315053536014016450973264"))
	if err != nil {
		lang.SayString("dvix3299")
	} else {
		if !(rexsult.ToString() == "54") {
			lang.SayString("dvix3299")
		}
	}
	rexsult, err = lang.RxFromString("042.668056830485571428877212944418").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-01364112374639.4941124016455321071"))
	if err != nil {
		lang.SayString("dvix3401")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3401")
		}
	}
	rexsult, err = lang.RxFromString("81721.5803096185422787702538493471").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("900099473.245809628076790757217328"))
	if err != nil {
		lang.SayString("dvix3403")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3403")
		}
	}
	rexsult, err = lang.RxFromString("3991.56734635183403814636354392163E-807").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("72.3239822255871305731314565069132"))
	if err != nil {
		lang.SayString("dvix3404")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3404")
		}
	}
	rexsult, err = lang.RxFromString("-66.3393308595957726456416979163306").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("5.08486573050459213864684589662559"))
	if err != nil {
		lang.SayString("dvix3405")
	} else {
		if !(rexsult.ToString() == "-13") {
			lang.SayString("dvix3405")
		}
	}
	rexsult, err = lang.RxFromString("-019133598.609812524622150421584346").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-858439846.628367734642622922030051"))
	if err != nil {
		lang.SayString("dvix3407")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3407")
		}
	}
	rexsult, err = lang.RxFromString("465.351982159046525762715549761814").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("240444.975944666924657629172844780"))
	if err != nil {
		lang.SayString("dvix3408")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3408")
		}
	}
	rexsult, err = lang.RxFromString("28275236927392.4960902824105246047").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("28212038.4825243127096613158419270E+422"))
	if err != nil {
		lang.SayString("dvix3410")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3410")
		}
	}
	rexsult, err = lang.RxFromString("11791.8644211874630234271801789996").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-8.45457275930363860982261343159741"))
	if err != nil {
		lang.SayString("dvix3411")
	} else {
		if !(rexsult.ToString() == "-1394") {
			lang.SayString("dvix3411")
		}
	}
	rexsult, err = lang.RxFromString("44.7085340739581668975502342787578").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-9337.05408133023920640485556647937"))
	if err != nil {
		lang.SayString("dvix3412")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3412")
		}
	}
	rexsult, err = lang.RxFromString("-367399415798804503177950040443482").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-54845683.9691776397285506712812754"))
	if err != nil {
		lang.SayString("dvix3414")
	} else {
		if !(rexsult.ToString() == "6698784465980529140072174") {
			lang.SayString("dvix3414")
		}
	}
	rexsult, err = lang.RxFromString("-2.87155919781024108503670175443740").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("89529730130.6427881332776797193807"))
	if err != nil {
		lang.SayString("dvix3415")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3415")
		}
	}
	rexsult, err = lang.RxFromString("611655569568.832698912762075889186").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("010182743219.475839030505966016982"))
	if err != nil {
		lang.SayString("dvix3417")
	} else {
		if !(rexsult.ToString() == "60") {
			lang.SayString("dvix3417")
		}
	}
	rexsult, err = lang.RxFromString("3457947.39062863674882672518304442").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-01.9995218868908849056866549811425"))
	if err != nil {
		lang.SayString("dvix3418")
	} else {
		if !(rexsult.ToString() == "-1729387") {
			lang.SayString("dvix3418")
		}
	}
	rexsult, err = lang.RxFromString("-5568057.17870139549478277980540034").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-407906443.141342175740471849723638"))
	if err != nil {
		lang.SayString("dvix3420")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3420")
		}
	}
	rexsult, err = lang.RxFromString("9804385273.49533524416415189990857").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("84.1433929743544659553964804646569"))
	if err != nil {
		lang.SayString("dvix3421")
	} else {
		if !(rexsult.ToString() == "116519965") {
			lang.SayString("dvix3421")
		}
	}
	rexsult, err = lang.RxFromString("-5234910986592.18801727046580014273E-547").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-5874220715892.91440069210512515154"))
	if err != nil {
		lang.SayString("dvix3422")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3422")
		}
	}
	rexsult, err = lang.RxFromString("698416560151955285929747633786867E-495").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("51754681.6784872628933218985216916E-266"))
	if err != nil {
		lang.SayString("dvix3423")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3423")
		}
	}
	rexsult, err = lang.RxFromString("-32174291345686.5371446616670961807").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("79518863759385.5925052747867099091E+408"))
	if err != nil {
		lang.SayString("dvix3425")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3425")
		}
	}
	rexsult, err = lang.RxFromString("-8151730494.53190523620899410544099E+688").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-93173.0631474527142307644239919480E+900"))
	if err != nil {
		lang.SayString("dvix3426")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3426")
		}
	}
	rexsult, err = lang.RxFromString("1.33649801345976199708341799505220").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-56623.0530039528969825480755159562E+459"))
	if err != nil {
		lang.SayString("dvix3427")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3427")
		}
	}
	rexsult, err = lang.RxFromString("4286562.76568866751577306056498271").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("6286.77291578497580015557979349893E+820"))
	if err != nil {
		lang.SayString("dvix3429")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3429")
		}
	}
	rexsult, err = lang.RxFromString("-765782.827432642697305644096365566").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("67.1634368459576834692758114618652"))
	if err != nil {
		lang.SayString("dvix3430")
	} else {
		if !(rexsult.ToString() == "-11401") {
			lang.SayString("dvix3430")
		}
	}
	rexsult, err = lang.RxFromString("46.2835931916106252756465724211276").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("59.2989237834093118332826617957791"))
	if err != nil {
		lang.SayString("dvix3431")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3431")
		}
	}
	rexsult, err = lang.RxFromString("-3029555.82298840234029474459694644").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("857535844655004737373089601128532"))
	if err != nil {
		lang.SayString("dvix3432")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3432")
		}
	}
	rexsult, err = lang.RxFromString("-0138466789523.10694176543700501945E-948").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("481026979918882487383654367924619"))
	if err != nil {
		lang.SayString("dvix3433")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3433")
		}
	}
	rexsult, err = lang.RxFromString("-9593566466.96690575714244442109870").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-87632034347.4845477961976776833770E+769"))
	if err != nil {
		lang.SayString("dvix3434")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3434")
		}
	}
	rexsult, err = lang.RxFromString("-3189.30765477670526823106100241863E-898").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("565688889.355241946154894311253202E-466"))
	if err != nil {
		lang.SayString("dvix3435")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3435")
		}
	}
	rexsult, err = lang.RxFromString("-17084552395.6714834680088150543965").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-631925802672.685034379197328370812E+527"))
	if err != nil {
		lang.SayString("dvix3436")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3436")
		}
	}
	rexsult, err = lang.RxFromString("-763.440067781256632695791981893608").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("19.9263811350611007833220620745413"))
	if err != nil {
		lang.SayString("dvix3438")
	} else {
		if !(rexsult.ToString() == "-38") {
			lang.SayString("dvix3438")
		}
	}
	rexsult, err = lang.RxFromString("-0970725697662.27605454336231195463").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-4541.41897546697187157913886433474"))
	if err != nil {
		lang.SayString("dvix3441")
	} else {
		if !(rexsult.ToString() == "213749425") {
			lang.SayString("dvix3441")
		}
	}
	rexsult, err = lang.RxFromString("-808178238631844268316111259558675").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-598400.265108644514211244980426520"))
	if err != nil {
		lang.SayString("dvix3442")
	} else {
		if !(rexsult.ToString() == "1350564640015847635178945884") {
			lang.SayString("dvix3442")
		}
	}
	rexsult, err = lang.RxFromString("-9.90826595069053564311371766315200").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-031.625916781307847864872329806646"))
	if err != nil {
		lang.SayString("dvix3443")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3443")
		}
	}
	rexsult, err = lang.RxFromString("-196925.469891897719160698483752907").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-41268.9975444533794067723958739778"))
	if err != nil {
		lang.SayString("dvix3444")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("dvix3444")
		}
	}
	rexsult, err = lang.RxFromString("74815000.4716875558358937279052903").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-690425401708167622194241915195001E+891"))
	if err != nil {
		lang.SayString("dvix3447")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3447")
		}
	}
	rexsult, err = lang.RxFromString("-1683993.51210241555668790556759021").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-72394384927344.8402585228267493374"))
	if err != nil {
		lang.SayString("dvix3448")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3448")
		}
	}
	rexsult, err = lang.RxFromString("-763.148530974741766171756970448158").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("517370.808956957601473642272664647"))
	if err != nil {
		lang.SayString("dvix3449")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3449")
		}
	}
	rexsult, err = lang.RxFromString("-77.5841338812312523460591226178754").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-927540422.641025050968830154578151E+524"))
	if err != nil {
		lang.SayString("dvix3450")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3450")
		}
	}
	rexsult, err = lang.RxFromString("5176295309.89943746236102209837813").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-129733.103628797477167908698565465"))
	if err != nil {
		lang.SayString("dvix3451")
	} else {
		if !(rexsult.ToString() == "-39899") {
			lang.SayString("dvix3451")
		}
	}
	rexsult, err = lang.RxFromString("-254346232981353293392888785643245").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-764.416902486152093758287929678445"))
	if err != nil {
		lang.SayString("dvix3454")
	} else {
		if !(rexsult.ToString() == "332732350833857889204406356900") {
			lang.SayString("dvix3454")
		}
	}
	rexsult, err = lang.RxFromString("-2875.36745499544177964804277329726").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-13187.8492045054802205691248267638"))
	if err != nil {
		lang.SayString("dvix3455")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3455")
		}
	}
	rexsult, err = lang.RxFromString("7191753.79974136447257468282073718").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("81.3878426176038487948375777384429"))
	if err != nil {
		lang.SayString("dvix3458")
	} else {
		if !(rexsult.ToString() == "88363") {
			lang.SayString("dvix3458")
		}
	}
	rexsult, err = lang.RxFromString("502975804.069864536104621700404770").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("684.790028432074527960269515227243"))
	if err != nil {
		lang.SayString("dvix3459")
	} else {
		if !(rexsult.ToString() == "734496") {
			lang.SayString("dvix3459")
		}
	}
	rexsult, err = lang.RxFromString("1505368.42063731861590460453659570").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-465242.678439951462767630022819105"))
	if err != nil {
		lang.SayString("dvix3460")
	} else {
		if !(rexsult.ToString() == "-3") {
			lang.SayString("dvix3460")
		}
	}
	rexsult, err = lang.RxFromString("69225023.2850142784063416137144829").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("8584050.06648191798834819995325693"))
	if err != nil {
		lang.SayString("dvix3461")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("dvix3461")
		}
	}
	rexsult, err = lang.RxFromString("-527566.521273992424649346837337602E-408").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-834662.599983953345718523807123972"))
	if err != nil {
		lang.SayString("dvix3463")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3463")
		}
	}
	rexsult, err = lang.RxFromString("4.51117459466491451401835756593793").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("3873385.19981811640063144338087230"))
	if err != nil {
		lang.SayString("dvix3466")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3466")
		}
	}
	rexsult, err = lang.RxFromString("49553763474698.8118661758811091120").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("36.1713861293896216593840817950781E+410"))
	if err != nil {
		lang.SayString("dvix3467")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3467")
		}
	}
	rexsult, err = lang.RxFromString("-20101668541.7472260497594230105456").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-395562148.345003931161532101109964"))
	if err != nil {
		lang.SayString("dvix3469")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("dvix3469")
		}
	}
	rexsult, err = lang.RxFromString("5583903.18072100716072011264668631").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("460868914694.088387067451312500723"))
	if err != nil {
		lang.SayString("dvix3470")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3470")
		}
	}
	rexsult, err = lang.RxFromString("-955638397975240685017992436116257E+020").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-508580.148958769104511751975720470E+662"))
	if err != nil {
		lang.SayString("dvix3471")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3471")
		}
	}
	rexsult, err = lang.RxFromString("-812.266340554281305985524813074211E+396").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-3195.63111559114001594257448970812E+986"))
	if err != nil {
		lang.SayString("dvix3474")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3474")
		}
	}
	rexsult, err = lang.RxFromString("83946.0157801953636255078996185540").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("492670373.235391665758701500314473"))
	if err != nil {
		lang.SayString("dvix3476")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3476")
		}
	}
	rexsult, err = lang.RxFromString("489191747.148674326757767356626016").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("01136942.1182277580093027768490545"))
	if err != nil {
		lang.SayString("dvix3478")
	} else {
		if !(rexsult.ToString() == "430") {
			lang.SayString("dvix3478")
		}
	}
	rexsult, err = lang.RxFromString("-3376883870.85961692148022521960139").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-65247489449.7334589731171980408284"))
	if err != nil {
		lang.SayString("dvix3480")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3480")
		}
	}
	rexsult, err = lang.RxFromString("58.6776780370008364590621772421025").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("01.5925518866529044494309229975160"))
	if err != nil {
		lang.SayString("dvix3481")
	} else {
		if !(rexsult.ToString() == "36") {
			lang.SayString("dvix3481")
		}
	}
	rexsult, err = lang.RxFromString("4099616339.96249499552808575717579").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("290.795187361072489816791525139895"))
	if err != nil {
		lang.SayString("dvix3482")
	} else {
		if !(rexsult.ToString() == "14097951") {
			lang.SayString("dvix3482")
		}
	}
	rexsult, err = lang.RxFromString("85870777.2282833141709970713739108").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-2140392861153.69401346398478113715"))
	if err != nil {
		lang.SayString("dvix3483")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3483")
		}
	}
	rexsult, err = lang.RxFromString("20900.9693761555165742010339929779").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-38.7546147649523793463260940289585"))
	if err != nil {
		lang.SayString("dvix3484")
	} else {
		if !(rexsult.ToString() == "-539") {
			lang.SayString("dvix3484")
		}
	}
	rexsult, err = lang.RxFromString("448.827596155587910947511170319456").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("379130153.382794042652974596286062"))
	if err != nil {
		lang.SayString("dvix3485")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3485")
		}
	}
	rexsult, err = lang.RxFromString("98.4134807921002817357000140482039").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("3404725543.77032945444654351546779"))
	if err != nil {
		lang.SayString("dvix3486")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3486")
		}
	}
	rexsult, err = lang.RxFromString("545746433.649359734136476718176330E-787").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-5149957099709.12830072802043560650E-437"))
	if err != nil {
		lang.SayString("dvix3487")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3487")
		}
	}
	rexsult, err = lang.RxFromString("-706.145005094292315613907254240553").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("4739.82486195739758308735946332234"))
	if err != nil {
		lang.SayString("dvix3489")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3489")
		}
	}
	rexsult, err = lang.RxFromString("-769925786.823099083228795187975893").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-31201.9980469760239870067820594790"))
	if err != nil {
		lang.SayString("dvix3490")
	} else {
		if !(rexsult.ToString() == "24675") {
			lang.SayString("dvix3490")
		}
	}
	rexsult, err = lang.RxFromString("549760.911304725795164589619286514").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("165.160089615604924207754883953484"))
	if err != nil {
		lang.SayString("dvix3492")
	} else {
		if !(rexsult.ToString() == "3328") {
			lang.SayString("dvix3492")
		}
	}
	rexsult, err = lang.RxFromString("3650514.18649737956855828939662794").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("08086721.4036886948248274834735629"))
	if err != nil {
		lang.SayString("dvix3493")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3493")
		}
	}
	rexsult, err = lang.RxFromString("55067723881950.1346958179604099594").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("-8.90481481687182931431054785192083"))
	if err != nil {
		lang.SayString("dvix3494")
	} else {
		if !(rexsult.ToString() == "-6184039198391") {
			lang.SayString("dvix3494")
		}
	}
	rexsult, err = lang.RxFromString("868251123.413992653362860637541060E+019").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("5579665045.37858308541154858567656E+131"))
	if err != nil {
		lang.SayString("dvix3495")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3495")
		}
	}
	rexsult, err = lang.RxFromString("-253280724.939458021588167965038184").OpDivI(lang.RxSetWithDigit(33), lang.RxFromString("616988.426425908872398170896375634E+396"))
	if err != nil {
		lang.SayString("dvix3500")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dvix3500")
		}
	}
	rexsult, err = lang.RxFromString("3915134.7").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-597164907."))
	if err != nil {
		lang.SayString("xdvi002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi002")
		}
	}
	rexsult, err = lang.RxFromString("309759261").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("62663.487"))
	if err != nil {
		lang.SayString("xdvi003")
	} else {
		if !(rexsult.ToString() == "4943") {
			lang.SayString("xdvi003")
		}
	}
	rexsult, err = lang.RxFromString("3.93591888E-236595626").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7242375.00"))
	if err != nil {
		lang.SayString("xdvi004")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi004")
		}
	}
	rexsult, err = lang.RxFromString("5.11970092").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-8807.22036"))
	if err != nil {
		lang.SayString("xdvi006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi006")
		}
	}
	rexsult, err = lang.RxFromString("-7.99874516").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4561.83758"))
	if err != nil {
		lang.SayString("xdvi007")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi007")
		}
	}
	rexsult, err = lang.RxFromString("297802878").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-927206.324"))
	if err != nil {
		lang.SayString("xdvi008")
	} else {
		if !(rexsult.ToString() == "-321") {
			lang.SayString("xdvi008")
		}
	}
	rexsult, err = lang.RxFromString("-766.651824").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("31300.3619"))
	if err != nil {
		lang.SayString("xdvi009")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi009")
		}
	}
	rexsult, err = lang.RxFromString("456417160").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-41346.1024"))
	if err != nil {
		lang.SayString("xdvi011")
	} else {
		if !(rexsult.ToString() == "-11038") {
			lang.SayString("xdvi011")
		}
	}
	rexsult, err = lang.RxFromString("102895.722").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-2.62214826"))
	if err != nil {
		lang.SayString("xdvi012")
	} else {
		if !(rexsult.ToString() == "-39241") {
			lang.SayString("xdvi012")
		}
	}
	rexsult, err = lang.RxFromString("-654645.954").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-9.12535752"))
	if err != nil {
		lang.SayString("xdvi015")
	} else {
		if !(rexsult.ToString() == "71739") {
			lang.SayString("xdvi015")
		}
	}
	rexsult, err = lang.RxFromString("63.1917772E-706014634").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-7.56253257E-138579234"))
	if err != nil {
		lang.SayString("xdvi016")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi016")
		}
	}
	rexsult, err = lang.RxFromString("-39674.7190").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2490607.78"))
	if err != nil {
		lang.SayString("xdvi017")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi017")
		}
	}
	rexsult, err = lang.RxFromString("-3364.59737E-600363681").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("896487.451"))
	if err != nil {
		lang.SayString("xdvi018")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi018")
		}
	}
	rexsult, err = lang.RxFromString("-64138.0578").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("31759011.3E+697488342"))
	if err != nil {
		lang.SayString("xdvi019")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi019")
		}
	}
	rexsult, err = lang.RxFromString("61399.8527").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-64344484.5"))
	if err != nil {
		lang.SayString("xdvi020")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi020")
		}
	}
	rexsult, err = lang.RxFromString("-722960.204").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-26154599.8"))
	if err != nil {
		lang.SayString("xdvi021")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi021")
		}
	}
	rexsult, err = lang.RxFromString("43.7456245").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("547441956."))
	if err != nil {
		lang.SayString("xdvi023")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi023")
		}
	}
	rexsult, err = lang.RxFromString("-73150542E-242017390").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-8.15869954"))
	if err != nil {
		lang.SayString("xdvi024")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi024")
		}
	}
	rexsult, err = lang.RxFromString("29.498114").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-26486451"))
	if err != nil {
		lang.SayString("xdvi026")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi026")
		}
	}
	rexsult, err = lang.RxFromString("-349388.759").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-196215.776"))
	if err != nil {
		lang.SayString("xdvi028")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("xdvi028")
		}
	}
	rexsult, err = lang.RxFromString("-70905112.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-91353968.8"))
	if err != nil {
		lang.SayString("xdvi029")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi029")
		}
	}
	rexsult, err = lang.RxFromString("-225094.28").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-88.7723542"))
	if err != nil {
		lang.SayString("xdvi030")
	} else {
		if !(rexsult.ToString() == "2535") {
			lang.SayString("xdvi030")
		}
	}
	rexsult, err = lang.RxFromString("50.4442340").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("82.7952169E+880120759"))
	if err != nil {
		lang.SayString("xdvi031")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi031")
		}
	}
	rexsult, err = lang.RxFromString("-32311.9037").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8.36379449"))
	if err != nil {
		lang.SayString("xdvi032")
	} else {
		if !(rexsult.ToString() == "-3863") {
			lang.SayString("xdvi032")
		}
	}
	rexsult, err = lang.RxFromString("849.515993E-878446473").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1039.08778"))
	if err != nil {
		lang.SayString("xdvi035")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi035")
		}
	}
	rexsult, err = lang.RxFromString("213361789").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-599.644851"))
	if err != nil {
		lang.SayString("xdvi036")
	} else {
		if !(rexsult.ToString() == "-355813") {
			lang.SayString("xdvi036")
		}
	}
	rexsult, err = lang.RxFromString("-795522555.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-298.037702"))
	if err != nil {
		lang.SayString("xdvi037")
	} else {
		if !(rexsult.ToString() == "2669201") {
			lang.SayString("xdvi037")
		}
	}
	rexsult, err = lang.RxFromString("-1.70781105E-848889023").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("36504769.4"))
	if err != nil {
		lang.SayString("xdvi039")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi039")
		}
	}
	rexsult, err = lang.RxFromString("-5290.54984E-490626676").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("842535254"))
	if err != nil {
		lang.SayString("xdvi040")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi040")
		}
	}
	rexsult, err = lang.RxFromString("-4629035.31").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-167.884398"))
	if err != nil {
		lang.SayString("xdvi042")
	} else {
		if !(rexsult.ToString() == "27572") {
			lang.SayString("xdvi042")
		}
	}
	rexsult, err = lang.RxFromString("-66527378.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-706400268."))
	if err != nil {
		lang.SayString("xdvi043")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi043")
		}
	}
	rexsult, err = lang.RxFromString("-2510497.53").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("372882462."))
	if err != nil {
		lang.SayString("xdvi044")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi044")
		}
	}
	rexsult, err = lang.RxFromString("136.255393E+53329961").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-53494.7201E+720058060"))
	if err != nil {
		lang.SayString("xdvi045")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi045")
		}
	}
	rexsult, err = lang.RxFromString("-876673.100").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-6150.92266"))
	if err != nil {
		lang.SayString("xdvi046")
	} else {
		if !(rexsult.ToString() == "142") {
			lang.SayString("xdvi046")
		}
	}
	rexsult, err = lang.RxFromString("3.61890453E-985606128").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("30664416."))
	if err != nil {
		lang.SayString("xdvi049")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi049")
		}
	}
	rexsult, err = lang.RxFromString("218699.206").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("556944241."))
	if err != nil {
		lang.SayString("xdvi051")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi051")
		}
	}
	rexsult, err = lang.RxFromString("106211716.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-3456793.74"))
	if err != nil {
		lang.SayString("xdvi052")
	} else {
		if !(rexsult.ToString() == "-30") {
			lang.SayString("xdvi052")
		}
	}
	rexsult, err = lang.RxFromString("364.99811").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-46222.0505"))
	if err != nil {
		lang.SayString("xdvi054")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi054")
		}
	}
	rexsult, err = lang.RxFromString("-392217576.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-958364096"))
	if err != nil {
		lang.SayString("xdvi055")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi055")
		}
	}
	rexsult, err = lang.RxFromString("169601285").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("714526.639"))
	if err != nil {
		lang.SayString("xdvi056")
	} else {
		if !(rexsult.ToString() == "237") {
			lang.SayString("xdvi056")
		}
	}
	rexsult, err = lang.RxFromString("-674.094552E+586944319").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("6354.2668E+589657266"))
	if err != nil {
		lang.SayString("xdvi057")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi057")
		}
	}
	rexsult, err = lang.RxFromString("-746.293386").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("927749.647"))
	if err != nil {
		lang.SayString("xdvi059")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi059")
		}
	}
	rexsult, err = lang.RxFromString("6.64377249").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("79161.1070E+619453776"))
	if err != nil {
		lang.SayString("xdvi061")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi061")
		}
	}
	rexsult, err = lang.RxFromString("3146.66571E-313373366").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("88.5282010"))
	if err != nil {
		lang.SayString("xdvi062")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi062")
		}
	}
	rexsult, err = lang.RxFromString("6.44693097").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-87195.8711"))
	if err != nil {
		lang.SayString("xdvi063")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi063")
		}
	}
	rexsult, err = lang.RxFromString("-7701.42814").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("72667.5181"))
	if err != nil {
		lang.SayString("xdvi065")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi065")
		}
	}
	rexsult, err = lang.RxFromString("-851.754789").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-582659.149"))
	if err != nil {
		lang.SayString("xdvi066")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi066")
		}
	}
	rexsult, err = lang.RxFromString("-5.01992943").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7852.16531"))
	if err != nil {
		lang.SayString("xdvi067")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi067")
		}
	}
	rexsult, err = lang.RxFromString("-12393257.2").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("76803689E+949125770"))
	if err != nil {
		lang.SayString("xdvi068")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi068")
		}
	}
	rexsult, err = lang.RxFromString("-296590035").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-481734529"))
	if err != nil {
		lang.SayString("xdvi071")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi071")
		}
	}
	rexsult, err = lang.RxFromString("8.27822605").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9241557.19"))
	if err != nil {
		lang.SayString("xdvi072")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi072")
		}
	}
	rexsult, err = lang.RxFromString("-1.43581098").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7286313.54"))
	if err != nil {
		lang.SayString("xdvi073")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi073")
		}
	}
	rexsult, err = lang.RxFromString("-699036193.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("759263.509E+533543625"))
	if err != nil {
		lang.SayString("xdvi074")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi074")
		}
	}
	rexsult, err = lang.RxFromString("-83.7273615E-305281957").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-287779593.E+458777774"))
	if err != nil {
		lang.SayString("xdvi075")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi075")
		}
	}
	rexsult, err = lang.RxFromString("8.48503224").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("6522.03316"))
	if err != nil {
		lang.SayString("xdvi076")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi076")
		}
	}
	rexsult, err = lang.RxFromString("527916091").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-809.054070"))
	if err != nil {
		lang.SayString("xdvi077")
	} else {
		if !(rexsult.ToString() == "-652510") {
			lang.SayString("xdvi077")
		}
	}
	rexsult, err = lang.RxFromString("3857058.60").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5792997.58E+881077409"))
	if err != nil {
		lang.SayString("xdvi078")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi078")
		}
	}
	rexsult, err = lang.RxFromString("-580.502955").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("38125521.7"))
	if err != nil {
		lang.SayString("xdvi080")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi080")
		}
	}
	rexsult, err = lang.RxFromString("-8378.55499").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("760.131257"))
	if err != nil {
		lang.SayString("xdvi083")
	} else {
		if !(rexsult.ToString() == "-11") {
			lang.SayString("xdvi083")
		}
	}
	rexsult, err = lang.RxFromString("-717.697718").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("984304413"))
	if err != nil {
		lang.SayString("xdvi084")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi084")
		}
	}
	rexsult, err = lang.RxFromString("-76762243.4E-741100094").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-273.706674"))
	if err != nil {
		lang.SayString("xdvi085")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi085")
		}
	}
	rexsult, err = lang.RxFromString("-359866845.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4.57434117"))
	if err != nil {
		lang.SayString("xdvi087")
	} else {
		if !(rexsult.ToString() == "78670748") {
			lang.SayString("xdvi087")
		}
	}
	rexsult, err = lang.RxFromString("779934536.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-76562645.7"))
	if err != nil {
		lang.SayString("xdvi088")
	} else {
		if !(rexsult.ToString() == "-10") {
			lang.SayString("xdvi088")
		}
	}
	rexsult, err = lang.RxFromString("-4820.95451").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3516234.99E+303303176"))
	if err != nil {
		lang.SayString("xdvi089")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi089")
		}
	}
	rexsult, err = lang.RxFromString("69355976.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-9.57838562E+758804984"))
	if err != nil {
		lang.SayString("xdvi090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi090")
		}
	}
	rexsult, err = lang.RxFromString("-671.507198E-908587890").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3057429.32E-555230623"))
	if err != nil {
		lang.SayString("xdvi094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi094")
		}
	}
	rexsult, err = lang.RxFromString("329579114").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("146780548."))
	if err != nil {
		lang.SayString("xdvi096")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("xdvi096")
		}
	}
	rexsult, err = lang.RxFromString("-789904.686E-217225000").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1991.07181E-84080059"))
	if err != nil {
		lang.SayString("xdvi097")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi097")
		}
	}
	rexsult, err = lang.RxFromString("59893.3544").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-408595868"))
	if err != nil {
		lang.SayString("xdvi098")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi098")
		}
	}
	rexsult, err = lang.RxFromString("9866.99208").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("708756501."))
	if err != nil {
		lang.SayString("xdvi100")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi100")
		}
	}
	rexsult, err = lang.RxFromString("-78810.6297").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-399884.68"))
	if err != nil {
		lang.SayString("xdvi101")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi101")
		}
	}
	rexsult, err = lang.RxFromString("409189761").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-771.471460"))
	if err != nil {
		lang.SayString("xdvi102")
	} else {
		if !(rexsult.ToString() == "-530401") {
			lang.SayString("xdvi102")
		}
	}
	rexsult, err = lang.RxFromString("-1.68748838").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("460.46924"))
	if err != nil {
		lang.SayString("xdvi103")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi103")
		}
	}
	rexsult, err = lang.RxFromString("553527296.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-7924.40185"))
	if err != nil {
		lang.SayString("xdvi104")
	} else {
		if !(rexsult.ToString() == "-69850") {
			lang.SayString("xdvi104")
		}
	}
	rexsult, err = lang.RxFromString("-38.7465207").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("64936.2942"))
	if err != nil {
		lang.SayString("xdvi105")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi105")
		}
	}
	rexsult, err = lang.RxFromString("-201075.248").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("845.663928"))
	if err != nil {
		lang.SayString("xdvi106")
	} else {
		if !(rexsult.ToString() == "-237") {
			lang.SayString("xdvi106")
		}
	}
	rexsult, err = lang.RxFromString("91048.4559").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("75953609.3"))
	if err != nil {
		lang.SayString("xdvi107")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi107")
		}
	}
	rexsult, err = lang.RxFromString("6898273.86E-252097460").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("15.3456196"))
	if err != nil {
		lang.SayString("xdvi108")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi108")
		}
	}
	rexsult, err = lang.RxFromString("-17643.39").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.0352568E+304871331"))
	if err != nil {
		lang.SayString("xdvi110")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi110")
		}
	}
	rexsult, err = lang.RxFromString("4589785.16").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7459.04237"))
	if err != nil {
		lang.SayString("xdvi111")
	} else {
		if !(rexsult.ToString() == "615") {
			lang.SayString("xdvi111")
		}
	}
	rexsult, err = lang.RxFromString("-51.1632090E-753968082").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8.96207471E-585797887"))
	if err != nil {
		lang.SayString("xdvi112")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi112")
		}
	}
	rexsult, err = lang.RxFromString("503712056.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-57490703.5E+924890183"))
	if err != nil {
		lang.SayString("xdvi114")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi114")
		}
	}
	rexsult, err = lang.RxFromString("883.849223").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("249259171"))
	if err != nil {
		lang.SayString("xdvi115")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi115")
		}
	}
	rexsult, err = lang.RxFromString("14239029.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-36527.2221"))
	if err != nil {
		lang.SayString("xdvi120")
	} else {
		if !(rexsult.ToString() == "-389") {
			lang.SayString("xdvi120")
		}
	}
	rexsult, err = lang.RxFromString("72333.2654E-544425548").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-690.664836E+662155120"))
	if err != nil {
		lang.SayString("xdvi121")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi121")
		}
	}
	rexsult, err = lang.RxFromString("-37721.1567E-115787341").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-828949864E-76251747"))
	if err != nil {
		lang.SayString("xdvi122")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi122")
		}
	}
	rexsult, err = lang.RxFromString("-2078852.83E-647080089").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-119779858.E+734665461"))
	if err != nil {
		lang.SayString("xdvi123")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi123")
		}
	}
	rexsult, err = lang.RxFromString("-79145.3625").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-7718.57307"))
	if err != nil {
		lang.SayString("xdvi124")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("xdvi124")
		}
	}
	rexsult, err = lang.RxFromString("911249557").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("79810804.9"))
	if err != nil {
		lang.SayString("xdvi126")
	} else {
		if !(rexsult.ToString() == "11") {
			lang.SayString("xdvi126")
		}
	}
	rexsult, err = lang.RxFromString("341134.994").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3.37486292"))
	if err != nil {
		lang.SayString("xdvi127")
	} else {
		if !(rexsult.ToString() == "101081") {
			lang.SayString("xdvi127")
		}
	}
	rexsult, err = lang.RxFromString("699631.893").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-226.423958"))
	if err != nil {
		lang.SayString("xdvi130")
	} else {
		if !(rexsult.ToString() == "-3089") {
			lang.SayString("xdvi130")
		}
	}
	rexsult, err = lang.RxFromString("-249350139.E-571793673").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("775732428."))
	if err != nil {
		lang.SayString("xdvi131")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi131")
		}
	}
	rexsult, err = lang.RxFromString("5.11629020").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-480.53194"))
	if err != nil {
		lang.SayString("xdvi132")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi132")
		}
	}
	rexsult, err = lang.RxFromString("-8.23352673E-446723147").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-530710.866"))
	if err != nil {
		lang.SayString("xdvi133")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi133")
		}
	}
	rexsult, err = lang.RxFromString("7.0598608").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-95908.35"))
	if err != nil {
		lang.SayString("xdvi134")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi134")
		}
	}
	rexsult, err = lang.RxFromString("-7.91189845E+207202706").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1532.71847E+509944335"))
	if err != nil {
		lang.SayString("xdvi135")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi135")
		}
	}
	rexsult, err = lang.RxFromString("208839370.E-215147432").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-75.9420559"))
	if err != nil {
		lang.SayString("xdvi136")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi136")
		}
	}
	rexsult, err = lang.RxFromString("427.754244E-353328369").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4705.0796"))
	if err != nil {
		lang.SayString("xdvi137")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi137")
		}
	}
	rexsult, err = lang.RxFromString("452371821.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4109709.19"))
	if err != nil {
		lang.SayString("xdvi139")
	} else {
		if !(rexsult.ToString() == "-110") {
			lang.SayString("xdvi139")
		}
	}
	rexsult, err = lang.RxFromString("94007.4392").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-9467725.5E+681898234"))
	if err != nil {
		lang.SayString("xdvi140")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi140")
		}
	}
	rexsult, err = lang.RxFromString("99147554.0E-751410586").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("38313.6423"))
	if err != nil {
		lang.SayString("xdvi141")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi141")
		}
	}
	rexsult, err = lang.RxFromString("-7919.30254").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-669.607854"))
	if err != nil {
		lang.SayString("xdvi142")
	} else {
		if !(rexsult.ToString() == "11") {
			lang.SayString("xdvi142")
		}
	}
	rexsult, err = lang.RxFromString("3455755.47E-112465506").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("771.674306"))
	if err != nil {
		lang.SayString("xdvi144")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi144")
		}
	}
	rexsult, err = lang.RxFromString("-477067757.E-961684940").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7.70122608E-741072245"))
	if err != nil {
		lang.SayString("xdvi145")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi145")
		}
	}
	rexsult, err = lang.RxFromString("76482.352").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8237806.8"))
	if err != nil {
		lang.SayString("xdvi146")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi146")
		}
	}
	rexsult, err = lang.RxFromString("1.21505164E-565556504").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9.26146573"))
	if err != nil {
		lang.SayString("xdvi147")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi147")
		}
	}
	rexsult, err = lang.RxFromString("-8303060.25E-169894883").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("901561.985"))
	if err != nil {
		lang.SayString("xdvi148")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi148")
		}
	}
	rexsult, err = lang.RxFromString("-592464.92").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("71445510.7"))
	if err != nil {
		lang.SayString("xdvi149")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi149")
		}
	}
	rexsult, err = lang.RxFromString("-73774.4165").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-39.8243027"))
	if err != nil {
		lang.SayString("xdvi150")
	} else {
		if !(rexsult.ToString() == "1852") {
			lang.SayString("xdvi150")
		}
	}
	rexsult, err = lang.RxFromString("-524724715.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-55763.7937"))
	if err != nil {
		lang.SayString("xdvi151")
	} else {
		if !(rexsult.ToString() == "9409") {
			lang.SayString("xdvi151")
		}
	}
	rexsult, err = lang.RxFromString("37.6027452").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7.22454233"))
	if err != nil {
		lang.SayString("xdvi153")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("xdvi153")
		}
	}
	rexsult, err = lang.RxFromString("2447660.39").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-36981.4253"))
	if err != nil {
		lang.SayString("xdvi154")
	} else {
		if !(rexsult.ToString() == "-66") {
			lang.SayString("xdvi154")
		}
	}
	rexsult, err = lang.RxFromString("2160.36419").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1418.33574E+656265382"))
	if err != nil {
		lang.SayString("xdvi155")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi155")
		}
	}
	rexsult, err = lang.RxFromString("8926.44939").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("54.9430027"))
	if err != nil {
		lang.SayString("xdvi156")
	} else {
		if !(rexsult.ToString() == "162") {
			lang.SayString("xdvi156")
		}
	}
	rexsult, err = lang.RxFromString("861588029").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-41657398E+77955925"))
	if err != nil {
		lang.SayString("xdvi157")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi157")
		}
	}
	rexsult, err = lang.RxFromString("-34.5253062").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("52.6722019"))
	if err != nil {
		lang.SayString("xdvi158")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi158")
		}
	}
	rexsult, err = lang.RxFromString("-18861647.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("99794586.7"))
	if err != nil {
		lang.SayString("xdvi159")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi159")
		}
	}
	rexsult, err = lang.RxFromString("322192.407").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("461.67044"))
	if err != nil {
		lang.SayString("xdvi160")
	} else {
		if !(rexsult.ToString() == "697") {
			lang.SayString("xdvi160")
		}
	}
	rexsult, err = lang.RxFromString("293.773732").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("479899052E+789950177"))
	if err != nil {
		lang.SayString("xdvi162")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi162")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("51897955.3"))
	if err != nil {
		lang.SayString("xdvi163")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("xdvi163")
		}
	}
	rexsult, err = lang.RxFromString("37380.7802").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-277719788."))
	if err != nil {
		lang.SayString("xdvi164")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi164")
		}
	}
	rexsult, err = lang.RxFromString("320133844.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-977517477"))
	if err != nil {
		lang.SayString("xdvi165")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi165")
		}
	}
	rexsult, err = lang.RxFromString("-5409.00482").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-2.16149386"))
	if err != nil {
		lang.SayString("xdvi167")
	} else {
		if !(rexsult.ToString() == "2502") {
			lang.SayString("xdvi167")
		}
	}
	rexsult, err = lang.RxFromString("-957960.367").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("322.858170"))
	if err != nil {
		lang.SayString("xdvi168")
	} else {
		if !(rexsult.ToString() == "-2967") {
			lang.SayString("xdvi168")
		}
	}
	rexsult, err = lang.RxFromString("-11817.8754E+613893442").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-3.84735082E+888333249"))
	if err != nil {
		lang.SayString("xdvi169")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi169")
		}
	}
	rexsult, err = lang.RxFromString("-205842096.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-191342.721"))
	if err != nil {
		lang.SayString("xdvi171")
	} else {
		if !(rexsult.ToString() == "1075") {
			lang.SayString("xdvi171")
		}
	}
	rexsult, err = lang.RxFromString("42501124.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("884.938498E+123341480"))
	if err != nil {
		lang.SayString("xdvi172")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi172")
		}
	}
	rexsult, err = lang.RxFromString("-57809452.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-620380746"))
	if err != nil {
		lang.SayString("xdvi173")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi173")
		}
	}
	rexsult, err = lang.RxFromString("-8022370.31").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9858581.6"))
	if err != nil {
		lang.SayString("xdvi174")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi174")
		}
	}
	rexsult, err = lang.RxFromString("-697273715E-242824870").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-3.81757506"))
	if err != nil {
		lang.SayString("xdvi176")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi176")
		}
	}
	rexsult, err = lang.RxFromString("-7.42204403E-315716280").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-8156111.67E+283261636"))
	if err != nil {
		lang.SayString("xdvi177")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi177")
		}
	}
	rexsult, err = lang.RxFromString("738063892").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("89900467.8"))
	if err != nil {
		lang.SayString("xdvi178")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("xdvi178")
		}
	}
	rexsult, err = lang.RxFromString("613.207774").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-3007.78608"))
	if err != nil {
		lang.SayString("xdvi180")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi180")
		}
	}
	rexsult, err = lang.RxFromString("-93006222.3").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-3.08964619"))
	if err != nil {
		lang.SayString("xdvi181")
	} else {
		if !(rexsult.ToString() == "30102547") {
			lang.SayString("xdvi181")
		}
	}
	rexsult, err = lang.RxFromString("19272386.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-410442379."))
	if err != nil {
		lang.SayString("xdvi183")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi183")
		}
	}
	rexsult, err = lang.RxFromString("571.536725").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("389.899220"))
	if err != nil {
		lang.SayString("xdvi185")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("xdvi185")
		}
	}
	rexsult, err = lang.RxFromString("92427442.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("674334898."))
	if err != nil {
		lang.SayString("xdvi188")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi188")
		}
	}
	rexsult, err = lang.RxFromString("44651895.7").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-910508.438"))
	if err != nil {
		lang.SayString("xdvi189")
	} else {
		if !(rexsult.ToString() == "-49") {
			lang.SayString("xdvi189")
		}
	}
	rexsult, err = lang.RxFromString("25.2592149").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("59.0436981"))
	if err != nil {
		lang.SayString("xdvi191")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi191")
		}
	}
	rexsult, err = lang.RxFromString("-6.850835").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1273.48240"))
	if err != nil {
		lang.SayString("xdvi192")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi192")
		}
	}
	rexsult, err = lang.RxFromString("174.272325").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5638.16229"))
	if err != nil {
		lang.SayString("xdvi193")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi193")
		}
	}
	rexsult, err = lang.RxFromString("3455629.76").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-8.27332322"))
	if err != nil {
		lang.SayString("xdvi194")
	} else {
		if !(rexsult.ToString() == "-417683") {
			lang.SayString("xdvi194")
		}
	}
	rexsult, err = lang.RxFromString("-924337723E-640771235").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("86639377.1"))
	if err != nil {
		lang.SayString("xdvi195")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi195")
		}
	}
	rexsult, err = lang.RxFromString("-18857539.9").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("813013129."))
	if err != nil {
		lang.SayString("xdvi198")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi198")
		}
	}
	rexsult, err = lang.RxFromString("-8.29530327").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3243419.57E+35688332"))
	if err != nil {
		lang.SayString("xdvi199")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi199")
		}
	}
	rexsult, err = lang.RxFromString("-57101683.5").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("763551341E+991491712"))
	if err != nil {
		lang.SayString("xdvi200")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi200")
		}
	}
	rexsult, err = lang.RxFromString("-603326.740").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1710.95183"))
	if err != nil {
		lang.SayString("xdvi201")
	} else {
		if !(rexsult.ToString() == "-352") {
			lang.SayString("xdvi201")
		}
	}
	rexsult, err = lang.RxFromString("-48142763.3").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-943434114"))
	if err != nil {
		lang.SayString("xdvi202")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi202")
		}
	}
	rexsult, err = lang.RxFromString("-204.586767").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-235.531847"))
	if err != nil {
		lang.SayString("xdvi203")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi203")
		}
	}
	rexsult, err = lang.RxFromString("-70.3805581").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("830137.913"))
	if err != nil {
		lang.SayString("xdvi204")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi204")
		}
	}
	rexsult, err = lang.RxFromString("-8818.47606").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-60766.4571"))
	if err != nil {
		lang.SayString("xdvi205")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi205")
		}
	}
	rexsult, err = lang.RxFromString("37060929.3E-168439509").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-79576717.1"))
	if err != nil {
		lang.SayString("xdvi206")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi206")
		}
	}
	rexsult, err = lang.RxFromString("-656285310.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-107221462."))
	if err != nil {
		lang.SayString("xdvi207")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("xdvi207")
		}
	}
	rexsult, err = lang.RxFromString("653397.125").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7195.30990"))
	if err != nil {
		lang.SayString("xdvi208")
	} else {
		if !(rexsult.ToString() == "90") {
			lang.SayString("xdvi208")
		}
	}
	rexsult, err = lang.RxFromString("-62011.4563E-117563240").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-57.1731586E+115657204"))
	if err != nil {
		lang.SayString("xdvi211")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi211")
		}
	}
	rexsult, err = lang.RxFromString("739.944710").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("202949.175"))
	if err != nil {
		lang.SayString("xdvi213")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi213")
		}
	}
	rexsult, err = lang.RxFromString("87686.8016").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4204890.40"))
	if err != nil {
		lang.SayString("xdvi214")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi214")
		}
	}
	rexsult, err = lang.RxFromString("987126721.E-725794834").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4874166.23"))
	if err != nil {
		lang.SayString("xdvi215")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi215")
		}
	}
	rexsult, err = lang.RxFromString("728148726.E-661695938").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("32798.5202"))
	if err != nil {
		lang.SayString("xdvi216")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi216")
		}
	}
	rexsult, err = lang.RxFromString("7428219.97").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("667.326760"))
	if err != nil {
		lang.SayString("xdvi217")
	} else {
		if !(rexsult.ToString() == "11131") {
			lang.SayString("xdvi217")
		}
	}
	rexsult, err = lang.RxFromString("-358.24550").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4447.78675E+601402509"))
	if err != nil {
		lang.SayString("xdvi219")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi219")
		}
	}
	rexsult, err = lang.RxFromString("118.621826").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-2.72010038"))
	if err != nil {
		lang.SayString("xdvi220")
	} else {
		if !(rexsult.ToString() == "-43") {
			lang.SayString("xdvi220")
		}
	}
	rexsult, err = lang.RxFromString("-35544.4029").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-567830.130"))
	if err != nil {
		lang.SayString("xdvi223")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi223")
		}
	}
	rexsult, err = lang.RxFromString("-509.483395").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-147242915."))
	if err != nil {
		lang.SayString("xdvi225")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi225")
		}
	}
	rexsult, err = lang.RxFromString("895612630.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-36.4104040"))
	if err != nil {
		lang.SayString("xdvi227")
	} else {
		if !(rexsult.ToString() == "-24597711") {
			lang.SayString("xdvi227")
		}
	}
	rexsult, err = lang.RxFromString("25455.4973").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2955.00006E+528196218"))
	if err != nil {
		lang.SayString("xdvi228")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi228")
		}
	}
	rexsult, err = lang.RxFromString("62871.2202").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2484.0382E+211662557"))
	if err != nil {
		lang.SayString("xdvi230")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi230")
		}
	}
	rexsult, err = lang.RxFromString("71.9281575").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-9810012.5"))
	if err != nil {
		lang.SayString("xdvi231")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi231")
		}
	}
	rexsult, err = lang.RxFromString("-6388022.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-88.042967"))
	if err != nil {
		lang.SayString("xdvi232")
	} else {
		if !(rexsult.ToString() == "72555") {
			lang.SayString("xdvi232")
		}
	}
	rexsult, err = lang.RxFromString("372567445.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("96.0992141"))
	if err != nil {
		lang.SayString("xdvi233")
	} else {
		if !(rexsult.ToString() == "3876904") {
			lang.SayString("xdvi233")
		}
	}
	rexsult, err = lang.RxFromString("-3.65207541").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("74501982.0"))
	if err != nil {
		lang.SayString("xdvi235")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi235")
		}
	}
	rexsult, err = lang.RxFromString("-5297.76981").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-859.719404"))
	if err != nil {
		lang.SayString("xdvi236")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("xdvi236")
		}
	}
	rexsult, err = lang.RxFromString("-684172.592").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("766.448597E+288361959"))
	if err != nil {
		lang.SayString("xdvi237")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi237")
		}
	}
	rexsult, err = lang.RxFromString("626919.219").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("57469.8727E+13188610"))
	if err != nil {
		lang.SayString("xdvi238")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi238")
		}
	}
	rexsult, err = lang.RxFromString("-77480.5840").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("893265.594E+287982552"))
	if err != nil {
		lang.SayString("xdvi239")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi239")
		}
	}
	rexsult, err = lang.RxFromString("-7177620.29").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7786343.83"))
	if err != nil {
		lang.SayString("xdvi240")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi240")
		}
	}
	rexsult, err = lang.RxFromString("9.6224130").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4.50355112"))
	if err != nil {
		lang.SayString("xdvi241")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("xdvi241")
		}
	}
	rexsult, err = lang.RxFromString("-66.6337347E-597410086").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-818812885"))
	if err != nil {
		lang.SayString("xdvi242")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi242")
		}
	}
	rexsult, err = lang.RxFromString("65587553.7").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("600574.736"))
	if err != nil {
		lang.SayString("xdvi243")
	} else {
		if !(rexsult.ToString() == "109") {
			lang.SayString("xdvi243")
		}
	}
	rexsult, err = lang.RxFromString("-32401.939").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-585200217."))
	if err != nil {
		lang.SayString("xdvi244")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi244")
		}
	}
	rexsult, err = lang.RxFromString("69573.988").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-9.77003465E+740933668"))
	if err != nil {
		lang.SayString("xdvi245")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi245")
		}
	}
	rexsult, err = lang.RxFromString("216741082.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("250290244"))
	if err != nil {
		lang.SayString("xdvi248")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi248")
		}
	}
	rexsult, err = lang.RxFromString("-6364720.49").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5539245.64"))
	if err != nil {
		lang.SayString("xdvi249")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("xdvi249")
		}
	}
	rexsult, err = lang.RxFromString("-814599.475").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-14.5431191"))
	if err != nil {
		lang.SayString("xdvi250")
	} else {
		if !(rexsult.ToString() == "56012") {
			lang.SayString("xdvi250")
		}
	}
	rexsult, err = lang.RxFromString("-162726.257E-597285918").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4391.54799"))
	if err != nil {
		lang.SayString("xdvi253")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi253")
		}
	}
	rexsult, err = lang.RxFromString("700354586.E-99856707").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7198.0493E+436250299"))
	if err != nil {
		lang.SayString("xdvi254")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi254")
		}
	}
	rexsult, err = lang.RxFromString("39617663E-463704664").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-895.290346"))
	if err != nil {
		lang.SayString("xdvi255")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi255")
		}
	}
	rexsult, err = lang.RxFromString("5350882.59").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-36329829"))
	if err != nil {
		lang.SayString("xdvi256")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi256")
		}
	}
	rexsult, err = lang.RxFromString("231899031.E-481759076").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("726.337100"))
	if err != nil {
		lang.SayString("xdvi258")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi258")
		}
	}
	rexsult, err = lang.RxFromString("-9611312.33").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("22109735.9"))
	if err != nil {
		lang.SayString("xdvi259")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi259")
		}
	}
	rexsult, err = lang.RxFromString("-5604938.15E-36812542").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("735937577."))
	if err != nil {
		lang.SayString("xdvi260")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi260")
		}
	}
	rexsult, err = lang.RxFromString("-34865.7378E-368768024").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2297117.88"))
	if err != nil {
		lang.SayString("xdvi262")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi262")
		}
	}
	rexsult, err = lang.RxFromString("1123.32456").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7.86747918E+930888796"))
	if err != nil {
		lang.SayString("xdvi263")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi263")
		}
	}
	rexsult, err = lang.RxFromString("56.6607465E+467812565").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("909552512E+764516200"))
	if err != nil {
		lang.SayString("xdvi264")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi264")
		}
	}
	rexsult, err = lang.RxFromString("34.1935525").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-40767.6450"))
	if err != nil {
		lang.SayString("xdvi266")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi266")
		}
	}
	rexsult, err = lang.RxFromString("51.025101").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4467691.57"))
	if err != nil {
		lang.SayString("xdvi269")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi269")
		}
	}
	rexsult, err = lang.RxFromString("-2214.76582").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("379785372E+223117572"))
	if err != nil {
		lang.SayString("xdvi270")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi270")
		}
	}
	rexsult, err = lang.RxFromString("-2564.75207E-841443929").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-653498187"))
	if err != nil {
		lang.SayString("xdvi271")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi271")
		}
	}
	rexsult, err = lang.RxFromString("513115529.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("27775075.6E+217133352"))
	if err != nil {
		lang.SayString("xdvi272")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi272")
		}
	}
	rexsult, err = lang.RxFromString("-247157.208").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-532990.453"))
	if err != nil {
		lang.SayString("xdvi273")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi273")
		}
	}
	rexsult, err = lang.RxFromString("40.2490764E-339482253").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7626.85442E+594264540"))
	if err != nil {
		lang.SayString("xdvi274")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi274")
		}
	}
	rexsult, err = lang.RxFromString("-1156008.8").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-8870382.36"))
	if err != nil {
		lang.SayString("xdvi275")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi275")
		}
	}
	rexsult, err = lang.RxFromString("880097928.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-52455011.1E+204538218"))
	if err != nil {
		lang.SayString("xdvi276")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi276")
		}
	}
	rexsult, err = lang.RxFromString("5796.2524").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("34458329.7E+832129426"))
	if err != nil {
		lang.SayString("xdvi277")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi277")
		}
	}
	rexsult, err = lang.RxFromString("27.1000923E-218032223").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-45.0198341"))
	if err != nil {
		lang.SayString("xdvi278")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi278")
		}
	}
	rexsult, err = lang.RxFromString("84224841.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("2.62548255E+647087608"))
	if err != nil {
		lang.SayString("xdvi281")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi281")
		}
	}
	rexsult, err = lang.RxFromString("9090950.80").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("436.400932"))
	if err != nil {
		lang.SayString("xdvi284")
	} else {
		if !(rexsult.ToString() == "20831") {
			lang.SayString("xdvi284")
		}
	}
	rexsult, err = lang.RxFromString("-4.18074650E-858746879").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("571035.277E-279409165"))
	if err != nil {
		lang.SayString("xdvi288")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi288")
		}
	}
	rexsult, err = lang.RxFromString("5.15309635").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-695649.219E+451948183"))
	if err != nil {
		lang.SayString("xdvi289")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi289")
		}
	}
	rexsult, err = lang.RxFromString("3336750").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("6.47961126"))
	if err != nil {
		lang.SayString("xdvi292")
	} else {
		if !(rexsult.ToString() == "514961") {
			lang.SayString("xdvi292")
		}
	}
	rexsult, err = lang.RxFromString("904654622.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("692065270.E+329081915"))
	if err != nil {
		lang.SayString("xdvi293")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi293")
		}
	}
	rexsult, err = lang.RxFromString("304804380").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4681.23698"))
	if err != nil {
		lang.SayString("xdvi294")
	} else {
		if !(rexsult.ToString() == "-65111") {
			lang.SayString("xdvi294")
		}
	}
	rexsult, err = lang.RxFromString("674.55569").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-82981.2684E+852890752"))
	if err != nil {
		lang.SayString("xdvi295")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi295")
		}
	}
	rexsult, err = lang.RxFromString("-5111.51025E-108006096").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5448870.4E+279212255"))
	if err != nil {
		lang.SayString("xdvi296")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi296")
		}
	}
	rexsult, err = lang.RxFromString("-2623.45068").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-466463938."))
	if err != nil {
		lang.SayString("xdvi297")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi297")
		}
	}
	rexsult, err = lang.RxFromString("299350.435").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3373.33551"))
	if err != nil {
		lang.SayString("xdvi298")
	} else {
		if !(rexsult.ToString() == "88") {
			lang.SayString("xdvi298")
		}
	}
	rexsult, err = lang.RxFromString("3774.5358E-491090520").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("173.060090"))
	if err != nil {
		lang.SayString("xdvi300")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi300")
		}
	}
	rexsult, err = lang.RxFromString("-13.6783690").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-453.610117"))
	if err != nil {
		lang.SayString("xdvi301")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi301")
		}
	}
	rexsult, err = lang.RxFromString("-990100927.E-615244634").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("223801.421E+247632618"))
	if err != nil {
		lang.SayString("xdvi302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi302")
		}
	}
	rexsult, err = lang.RxFromString("1275.10292").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-667965353"))
	if err != nil {
		lang.SayString("xdvi303")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi303")
		}
	}
	rexsult, err = lang.RxFromString("-8.76375480E-596792197").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("992.077361"))
	if err != nil {
		lang.SayString("xdvi304")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi304")
		}
	}
	rexsult, err = lang.RxFromString("213577152").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-986710073E+31900046"))
	if err != nil {
		lang.SayString("xdvi306")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi306")
		}
	}
	rexsult, err = lang.RxFromString("91393.9398E-323439228").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-135.701000"))
	if err != nil {
		lang.SayString("xdvi307")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi307")
		}
	}
	rexsult, err = lang.RxFromString("59807846.1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.53345254"))
	if err != nil {
		lang.SayString("xdvi309")
	} else {
		if !(rexsult.ToString() == "39002084") {
			lang.SayString("xdvi309")
		}
	}
	rexsult, err = lang.RxFromString("-8046158.45").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8.3635397"))
	if err != nil {
		lang.SayString("xdvi310")
	} else {
		if !(rexsult.ToString() == "-962051") {
			lang.SayString("xdvi310")
		}
	}
	rexsult, err = lang.RxFromString("-948.038054").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3580.84510"))
	if err != nil {
		lang.SayString("xdvi312")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi312")
		}
	}
	rexsult, err = lang.RxFromString("79551.5014").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-538.186229"))
	if err != nil {
		lang.SayString("xdvi314")
	} else {
		if !(rexsult.ToString() == "-147") {
			lang.SayString("xdvi314")
		}
	}
	rexsult, err = lang.RxFromString("-3264204.54").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-42704.501"))
	if err != nil {
		lang.SayString("xdvi317")
	} else {
		if !(rexsult.ToString() == "76") {
			lang.SayString("xdvi317")
		}
	}
	rexsult, err = lang.RxFromString("1.21265492").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("44102.6073"))
	if err != nil {
		lang.SayString("xdvi318")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi318")
		}
	}
	rexsult, err = lang.RxFromString("745.78452").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-1922.00670E+375923302"))
	if err != nil {
		lang.SayString("xdvi320")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi320")
		}
	}
	rexsult, err = lang.RxFromString("-963717836").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-823989308"))
	if err != nil {
		lang.SayString("xdvi321")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("xdvi321")
		}
	}
	rexsult, err = lang.RxFromString("-808328.607E-790810342").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("53075.7082"))
	if err != nil {
		lang.SayString("xdvi323")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi323")
		}
	}
	rexsult, err = lang.RxFromString("700592.720").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-698485.085"))
	if err != nil {
		lang.SayString("xdvi324")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("xdvi324")
		}
	}
	rexsult, err = lang.RxFromString("-80273928.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("661346.239"))
	if err != nil {
		lang.SayString("xdvi325")
	} else {
		if !(rexsult.ToString() == "-121") {
			lang.SayString("xdvi325")
		}
	}
	rexsult, err = lang.RxFromString("-682.796370").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("71131.0224"))
	if err != nil {
		lang.SayString("xdvi328")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi328")
		}
	}
	rexsult, err = lang.RxFromString("89.9997490").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4993.69831"))
	if err != nil {
		lang.SayString("xdvi329")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi329")
		}
	}
	rexsult, err = lang.RxFromString("-8868.72074").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5592399.93"))
	if err != nil {
		lang.SayString("xdvi335")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi335")
		}
	}
	rexsult, err = lang.RxFromString("-74.7852037E-175205809").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("4.14316542"))
	if err != nil {
		lang.SayString("xdvi336")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi336")
		}
	}
	rexsult, err = lang.RxFromString("38660103.1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-6671.73085E+900998477"))
	if err != nil {
		lang.SayString("xdvi338")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi338")
		}
	}
	rexsult, err = lang.RxFromString("-52.2659460").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-296404199E+372050476"))
	if err != nil {
		lang.SayString("xdvi339")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi339")
		}
	}
	rexsult, err = lang.RxFromString("6.06625013").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-276.359186"))
	if err != nil {
		lang.SayString("xdvi340")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi340")
		}
	}
	rexsult, err = lang.RxFromString("-62971617.5E-241444744").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("46266799.3"))
	if err != nil {
		lang.SayString("xdvi341")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi341")
		}
	}
	rexsult, err = lang.RxFromString("2467915.01").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-92.5558322"))
	if err != nil {
		lang.SayString("xdvi343")
	} else {
		if !(rexsult.ToString() == "-26664") {
			lang.SayString("xdvi343")
		}
	}
	rexsult, err = lang.RxFromString("187.232671").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-840.469347"))
	if err != nil {
		lang.SayString("xdvi344")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi344")
		}
	}
	rexsult, err = lang.RxFromString("81233.6823").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-5192.21666E+309315093"))
	if err != nil {
		lang.SayString("xdvi345")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi345")
		}
	}
	rexsult, err = lang.RxFromString("78872665.3").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("172.102119"))
	if err != nil {
		lang.SayString("xdvi347")
	} else {
		if !(rexsult.ToString() == "458289") {
			lang.SayString("xdvi347")
		}
	}
	rexsult, err = lang.RxFromString("328268.1E-436315617").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-204.522245"))
	if err != nil {
		lang.SayString("xdvi348")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi348")
		}
	}
	rexsult, err = lang.RxFromString("-688755561.E-95301699").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("978.275312E+913812609"))
	if err != nil {
		lang.SayString("xdvi350")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi350")
		}
	}
	rexsult, err = lang.RxFromString("-5.47345502").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("59818.7580"))
	if err != nil {
		lang.SayString("xdvi351")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi351")
		}
	}
	rexsult, err = lang.RxFromString("563891620E-361354567").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-845900362."))
	if err != nil {
		lang.SayString("xdvi352")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi352")
		}
	}
	rexsult, err = lang.RxFromString("-69.7231286").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("85773.7504"))
	if err != nil {
		lang.SayString("xdvi353")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi353")
		}
	}
	rexsult, err = lang.RxFromString("-54.6254096").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-332921899."))
	if err != nil {
		lang.SayString("xdvi355")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi355")
		}
	}
	rexsult, err = lang.RxFromString("-9.04778095E-591874079").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8719.40286"))
	if err != nil {
		lang.SayString("xdvi356")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi356")
		}
	}
	rexsult, err = lang.RxFromString("-1546783").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-51935370.4"))
	if err != nil {
		lang.SayString("xdvi358")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi358")
		}
	}
	rexsult, err = lang.RxFromString("61302486.8").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("205.490417"))
	if err != nil {
		lang.SayString("xdvi359")
	} else {
		if !(rexsult.ToString() == "298322") {
			lang.SayString("xdvi359")
		}
	}
	rexsult, err = lang.RxFromString("-546398328.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-27.9149712"))
	if err != nil {
		lang.SayString("xdvi362")
	} else {
		if !(rexsult.ToString() == "19573666") {
			lang.SayString("xdvi362")
		}
	}
	rexsult, err = lang.RxFromString("5402066.1E-284978216").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("622.751128"))
	if err != nil {
		lang.SayString("xdvi363")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi363")
		}
	}
	rexsult, err = lang.RxFromString("18845620").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3129.43753"))
	if err != nil {
		lang.SayString("xdvi364")
	} else {
		if !(rexsult.ToString() == "6022") {
			lang.SayString("xdvi364")
		}
	}
	rexsult, err = lang.RxFromString("13.8003883E-386224921").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-84126481.9E-296378341"))
	if err != nil {
		lang.SayString("xdvi367")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi367")
		}
	}
	rexsult, err = lang.RxFromString("9820.90457").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("46671.5915"))
	if err != nil {
		lang.SayString("xdvi368")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi368")
		}
	}
	rexsult, err = lang.RxFromString("472648900").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-207.784153"))
	if err != nil {
		lang.SayString("xdvi370")
	} else {
		if !(rexsult.ToString() == "-2274711") {
			lang.SayString("xdvi370")
		}
	}
	rexsult, err = lang.RxFromString("-8754.49306").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-818.165153E+631475457"))
	if err != nil {
		lang.SayString("xdvi371")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi371")
		}
	}
	rexsult, err = lang.RxFromString("98750864").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("191380.551"))
	if err != nil {
		lang.SayString("xdvi372")
	} else {
		if !(rexsult.ToString() == "515") {
			lang.SayString("xdvi372")
		}
	}
	rexsult, err = lang.RxFromString("725292561.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-768963606.E+340762986"))
	if err != nil {
		lang.SayString("xdvi373")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi373")
		}
	}
	rexsult, err = lang.RxFromString("1862.80445").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("648254483."))
	if err != nil {
		lang.SayString("xdvi374")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi374")
		}
	}
	rexsult, err = lang.RxFromString("-5549320.1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-93580684.1"))
	if err != nil {
		lang.SayString("xdvi375")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi375")
		}
	}
	rexsult, err = lang.RxFromString("-14677053.1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-25784.7358"))
	if err != nil {
		lang.SayString("xdvi376")
	} else {
		if !(rexsult.ToString() == "569") {
			lang.SayString("xdvi376")
		}
	}
	rexsult, err = lang.RxFromString("-4131738.09").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7579.07566"))
	if err != nil {
		lang.SayString("xdvi378")
	} else {
		if !(rexsult.ToString() == "-545") {
			lang.SayString("xdvi378")
		}
	}
	rexsult, err = lang.RxFromString("829898241").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8912.99114E+929228149"))
	if err != nil {
		lang.SayString("xdvi380")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi380")
		}
	}
	rexsult, err = lang.RxFromString("53.6891691").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-11.2371140"))
	if err != nil {
		lang.SayString("xdvi381")
	} else {
		if !(rexsult.ToString() == "-4") {
			lang.SayString("xdvi381")
		}
	}
	rexsult, err = lang.RxFromString("-93951823.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-25317.8645"))
	if err != nil {
		lang.SayString("xdvi382")
	} else {
		if !(rexsult.ToString() == "3710") {
			lang.SayString("xdvi382")
		}
	}
	rexsult, err = lang.RxFromString("446919.123").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("951338490."))
	if err != nil {
		lang.SayString("xdvi383")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi383")
		}
	}
	rexsult, err = lang.RxFromString("-8.01787748").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-88.3076852"))
	if err != nil {
		lang.SayString("xdvi384")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi384")
		}
	}
	rexsult, err = lang.RxFromString("517458139").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-999731.548"))
	if err != nil {
		lang.SayString("xdvi385")
	} else {
		if !(rexsult.ToString() == "-517") {
			lang.SayString("xdvi385")
		}
	}
	rexsult, err = lang.RxFromString("-405543440").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-4013.18295"))
	if err != nil {
		lang.SayString("xdvi386")
	} else {
		if !(rexsult.ToString() == "101052") {
			lang.SayString("xdvi386")
		}
	}
	rexsult, err = lang.RxFromString("-151144455").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-170371.29"))
	if err != nil {
		lang.SayString("xdvi388")
	} else {
		if !(rexsult.ToString() == "887") {
			lang.SayString("xdvi388")
		}
	}
	rexsult, err = lang.RxFromString("534.394729").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-2369839.37"))
	if err != nil {
		lang.SayString("xdvi390")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi390")
		}
	}
	rexsult, err = lang.RxFromString("89100.1797").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("224.370309"))
	if err != nil {
		lang.SayString("xdvi391")
	} else {
		if !(rexsult.ToString() == "397") {
			lang.SayString("xdvi391")
		}
	}
	rexsult, err = lang.RxFromString("-821377.777").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("38.552821"))
	if err != nil {
		lang.SayString("xdvi392")
	} else {
		if !(rexsult.ToString() == "-21305") {
			lang.SayString("xdvi392")
		}
	}
	rexsult, err = lang.RxFromString("-392640.782").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-2571619.5E+113237865"))
	if err != nil {
		lang.SayString("xdvi393")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi393")
		}
	}
	rexsult, err = lang.RxFromString("-651397.712").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-723.59673"))
	if err != nil {
		lang.SayString("xdvi394")
	} else {
		if !(rexsult.ToString() == "900") {
			lang.SayString("xdvi394")
		}
	}
	rexsult, err = lang.RxFromString("86.6890892").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("940380864"))
	if err != nil {
		lang.SayString("xdvi395")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi395")
		}
	}
	rexsult, err = lang.RxFromString("173398265E-532383158").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3462.51450E+80986915"))
	if err != nil {
		lang.SayString("xdvi397")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi397")
		}
	}
	rexsult, err = lang.RxFromString("-1522176.78").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-6631061.77"))
	if err != nil {
		lang.SayString("xdvi398")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi398")
		}
	}
	rexsult, err = lang.RxFromString("538.10453").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("522934310"))
	if err != nil {
		lang.SayString("xdvi399")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi399")
		}
	}
	rexsult, err = lang.RxFromString("880243.444E-750940977").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-354.601578E-204943740"))
	if err != nil {
		lang.SayString("xdvi400")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi400")
		}
	}
	rexsult, err = lang.RxFromString("968370.780").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("6677268.73"))
	if err != nil {
		lang.SayString("xdvi401")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi401")
		}
	}
	rexsult, err = lang.RxFromString("-97.7474945").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("31248241.5"))
	if err != nil {
		lang.SayString("xdvi402")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi402")
		}
	}
	rexsult, err = lang.RxFromString("-187582786.E+369916952").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("957840435E+744365567"))
	if err != nil {
		lang.SayString("xdvi403")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi403")
		}
	}
	rexsult, err = lang.RxFromString("-328026144.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-125499735."))
	if err != nil {
		lang.SayString("xdvi404")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("xdvi404")
		}
	}
	rexsult, err = lang.RxFromString("-99424150.2E-523662102").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("3712.35030"))
	if err != nil {
		lang.SayString("xdvi405")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi405")
		}
	}
	rexsult, err = lang.RxFromString("14838.0718").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("9489893.28E+830631266"))
	if err != nil {
		lang.SayString("xdvi406")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi406")
		}
	}
	rexsult, err = lang.RxFromString("71207472.8").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-31835.0809"))
	if err != nil {
		lang.SayString("xdvi407")
	} else {
		if !(rexsult.ToString() == "-2236") {
			lang.SayString("xdvi407")
		}
	}
	rexsult, err = lang.RxFromString("-20440.4394").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-44.4064328E+511085806"))
	if err != nil {
		lang.SayString("xdvi408")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi408")
		}
	}
	rexsult, err = lang.RxFromString("-657.186702").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("426844.39"))
	if err != nil {
		lang.SayString("xdvi411")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi411")
		}
	}
	rexsult, err = lang.RxFromString("-41593077.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-688607.564"))
	if err != nil {
		lang.SayString("xdvi412")
	} else {
		if !(rexsult.ToString() == "60") {
			lang.SayString("xdvi412")
		}
	}
	rexsult, err = lang.RxFromString("-5786.38132").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("190556652.E+177045877"))
	if err != nil {
		lang.SayString("xdvi413")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi413")
		}
	}
	rexsult, err = lang.RxFromString("737622.974").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-241560693E+249506565"))
	if err != nil {
		lang.SayString("xdvi414")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi414")
		}
	}
	rexsult, err = lang.RxFromString("644136.179").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-835708.103"))
	if err != nil {
		lang.SayString("xdvi416")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi416")
		}
	}
	rexsult, err = lang.RxFromString("-31068.7549").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-3.41495042E+86001379"))
	if err != nil {
		lang.SayString("xdvi419")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi419")
		}
	}
	rexsult, err = lang.RxFromString("3898.03188").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-82572.615"))
	if err != nil {
		lang.SayString("xdvi422")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi422")
		}
	}
	rexsult, err = lang.RxFromString("-1.7619356").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-2546.64043"))
	if err != nil {
		lang.SayString("xdvi423")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi423")
		}
	}
	rexsult, err = lang.RxFromString("6.88891136E-935467395").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-785049.562E-741671442"))
	if err != nil {
		lang.SayString("xdvi425")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi425")
		}
	}
	rexsult, err = lang.RxFromString("975566251").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-519.858530"))
	if err != nil {
		lang.SayString("xdvi426")
	} else {
		if !(rexsult.ToString() == "-1876599") {
			lang.SayString("xdvi426")
		}
	}
	rexsult, err = lang.RxFromString("307401954").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-231481582."))
	if err != nil {
		lang.SayString("xdvi427")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("xdvi427")
		}
	}
	rexsult, err = lang.RxFromString("6.48674979").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-621732.532E+422575800"))
	if err != nil {
		lang.SayString("xdvi430")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi430")
		}
	}
	rexsult, err = lang.RxFromString("-31401.9418").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("36.3960679"))
	if err != nil {
		lang.SayString("xdvi431")
	} else {
		if !(rexsult.ToString() == "-862") {
			lang.SayString("xdvi431")
		}
	}
	rexsult, err = lang.RxFromString("31345321.1").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("51.5482191"))
	if err != nil {
		lang.SayString("xdvi432")
	} else {
		if !(rexsult.ToString() == "608077") {
			lang.SayString("xdvi432")
		}
	}
	rexsult, err = lang.RxFromString("70437.1551").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-62916.1233"))
	if err != nil {
		lang.SayString("xdvi434")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("xdvi434")
		}
	}
	rexsult, err = lang.RxFromString("916260164").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-58.4017325"))
	if err != nil {
		lang.SayString("xdvi435")
	} else {
		if !(rexsult.ToString() == "-15688920") {
			lang.SayString("xdvi435")
		}
	}
	rexsult, err = lang.RxFromString("19889085.3E-46816480").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1581683.94"))
	if err != nil {
		lang.SayString("xdvi436")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi436")
		}
	}
	rexsult, err = lang.RxFromString("-56312.3383").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("789.466064"))
	if err != nil {
		lang.SayString("xdvi437")
	} else {
		if !(rexsult.ToString() == "-71") {
			lang.SayString("xdvi437")
		}
	}
	rexsult, err = lang.RxFromString("183442.849").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-925876106"))
	if err != nil {
		lang.SayString("xdvi438")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi438")
		}
	}
	rexsult, err = lang.RxFromString("859658551.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("72338.2054"))
	if err != nil {
		lang.SayString("xdvi440")
	} else {
		if !(rexsult.ToString() == "11883") {
			lang.SayString("xdvi440")
		}
	}
	rexsult, err = lang.RxFromString("-969.881818").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("31170.8555"))
	if err != nil {
		lang.SayString("xdvi442")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi442")
		}
	}
	rexsult, err = lang.RxFromString("7980537.27").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("85.4040512"))
	if err != nil {
		lang.SayString("xdvi443")
	} else {
		if !(rexsult.ToString() == "93444") {
			lang.SayString("xdvi443")
		}
	}
	rexsult, err = lang.RxFromString("-114609916.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("7525.14981"))
	if err != nil {
		lang.SayString("xdvi444")
	} else {
		if !(rexsult.ToString() == "-15230") {
			lang.SayString("xdvi444")
		}
	}
	rexsult, err = lang.RxFromString("8.43404682E-500572568").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("474526719"))
	if err != nil {
		lang.SayString("xdvi445")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi445")
		}
	}
	rexsult, err = lang.RxFromString("-9.95836312").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-866466703"))
	if err != nil {
		lang.SayString("xdvi447")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi447")
		}
	}
	rexsult, err = lang.RxFromString("80919339.2E-967231586").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("219.824266"))
	if err != nil {
		lang.SayString("xdvi448")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi448")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-89827.5229"))
	if err != nil {
		lang.SayString("xdvi449")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("xdvi449")
		}
	}
	rexsult, err = lang.RxFromString("-4.54000153").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("6966333.74"))
	if err != nil {
		lang.SayString("xdvi450")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi450")
		}
	}
	rexsult, err = lang.RxFromString("28701538.7E-391015649").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-920999192."))
	if err != nil {
		lang.SayString("xdvi451")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi451")
		}
	}
	rexsult, err = lang.RxFromString("-361382575.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-7976.15286E+898491169"))
	if err != nil {
		lang.SayString("xdvi452")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi452")
		}
	}
	rexsult, err = lang.RxFromString("7021805.61").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1222952.83"))
	if err != nil {
		lang.SayString("xdvi453")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("xdvi453")
		}
	}
	rexsult, err = lang.RxFromString("-40.4811667").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-79655.5635"))
	if err != nil {
		lang.SayString("xdvi454")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi454")
		}
	}
	rexsult, err = lang.RxFromString("-37958476.0").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("584367.935"))
	if err != nil {
		lang.SayString("xdvi457")
	} else {
		if !(rexsult.ToString() == "-64") {
			lang.SayString("xdvi457")
		}
	}
	rexsult, err = lang.RxFromString("495233.553E-414152215").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("62352759.2"))
	if err != nil {
		lang.SayString("xdvi458")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi458")
		}
	}
	rexsult, err = lang.RxFromString("-502343060").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-96828.994"))
	if err != nil {
		lang.SayString("xdvi459")
	} else {
		if !(rexsult.ToString() == "5187") {
			lang.SayString("xdvi459")
		}
	}
	rexsult, err = lang.RxFromString("718180.587E-957473722").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1.66223443"))
	if err != nil {
		lang.SayString("xdvi461")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi461")
		}
	}
	rexsult, err = lang.RxFromString("-51592.2698").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-713885.741"))
	if err != nil {
		lang.SayString("xdvi462")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi462")
		}
	}
	rexsult, err = lang.RxFromString("51.2279848E+80439745").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("207.55925E+865165070"))
	if err != nil {
		lang.SayString("xdvi463")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi463")
		}
	}
	rexsult, err = lang.RxFromString("-5983.23468").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-39.9544513"))
	if err != nil {
		lang.SayString("xdvi464")
	} else {
		if !(rexsult.ToString() == "149") {
			lang.SayString("xdvi464")
		}
	}
	rexsult, err = lang.RxFromString("921639332.E-917542963").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("287325.891"))
	if err != nil {
		lang.SayString("xdvi465")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi465")
		}
	}
	rexsult, err = lang.RxFromString("91095916.8E-787312969").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-58643.418E+58189880"))
	if err != nil {
		lang.SayString("xdvi466")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi466")
		}
	}
	rexsult, err = lang.RxFromString("-6410.5555").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-234964259"))
	if err != nil {
		lang.SayString("xdvi467")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi467")
		}
	}
	rexsult, err = lang.RxFromString("-5.32711606").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-8447286.21"))
	if err != nil {
		lang.SayString("xdvi468")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi468")
		}
	}
	rexsult, err = lang.RxFromString("412411244.E-774339264").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("866452.465"))
	if err != nil {
		lang.SayString("xdvi470")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi470")
		}
	}
	rexsult, err = lang.RxFromString("-31027.8323").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-475378186."))
	if err != nil {
		lang.SayString("xdvi472")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi472")
		}
	}
	rexsult, err = lang.RxFromString("-1199339.72").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-5.73068392E+53774632"))
	if err != nil {
		lang.SayString("xdvi473")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi473")
		}
	}
	rexsult, err = lang.RxFromString("-2376150.83").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-46777583.3"))
	if err != nil {
		lang.SayString("xdvi475")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi475")
		}
	}
	rexsult, err = lang.RxFromString("6.3664211").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-140854908."))
	if err != nil {
		lang.SayString("xdvi476")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi476")
		}
	}
	rexsult, err = lang.RxFromString("-15.791522").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("1902.30210E+90741844"))
	if err != nil {
		lang.SayString("xdvi477")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi477")
		}
	}
	rexsult, err = lang.RxFromString("49436.6528").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("751.919517"))
	if err != nil {
		lang.SayString("xdvi480")
	} else {
		if !(rexsult.ToString() == "65") {
			lang.SayString("xdvi480")
		}
	}
	rexsult, err = lang.RxFromString("552.669453").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("8.3725760E+16223526"))
	if err != nil {
		lang.SayString("xdvi481")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi481")
		}
	}
	rexsult, err = lang.RxFromString("-3266303").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("453741.520"))
	if err != nil {
		lang.SayString("xdvi482")
	} else {
		if !(rexsult.ToString() == "-7") {
			lang.SayString("xdvi482")
		}
	}
	rexsult, err = lang.RxFromString("12302757.4").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("542922.487E+414443353"))
	if err != nil {
		lang.SayString("xdvi483")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi483")
		}
	}
	rexsult, err = lang.RxFromString("-5670757.79E-784754984").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("128144.503"))
	if err != nil {
		lang.SayString("xdvi484")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi484")
		}
	}
	rexsult, err = lang.RxFromString("88.5158199E-980164357").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("325846116"))
	if err != nil {
		lang.SayString("xdvi486")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi486")
		}
	}
	rexsult, err = lang.RxFromString("-22881.0408").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("5.63661562"))
	if err != nil {
		lang.SayString("xdvi487")
	} else {
		if !(rexsult.ToString() == "-4059") {
			lang.SayString("xdvi487")
		}
	}
	rexsult, err = lang.RxFromString("-3066962.41").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-55.3096879"))
	if err != nil {
		lang.SayString("xdvi490")
	} else {
		if !(rexsult.ToString() == "55450") {
			lang.SayString("xdvi490")
		}
	}
	rexsult, err = lang.RxFromString("998890068.").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-92.057879"))
	if err != nil {
		lang.SayString("xdvi492")
	} else {
		if !(rexsult.ToString() == "-10850674") {
			lang.SayString("xdvi492")
		}
	}
	rexsult, err = lang.RxFromString("122.495591").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-407836028."))
	if err != nil {
		lang.SayString("xdvi493")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi493")
		}
	}
	rexsult, err = lang.RxFromString("-286.371320").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("710319152"))
	if err != nil {
		lang.SayString("xdvi497")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi497")
		}
	}
	rexsult, err = lang.RxFromString("-6157.74292").OpDivI(lang.RxSetWithDigit(9), lang.RxFromString("-94075286.2E+92555877"))
	if err != nil {
		lang.SayString("xdvi499")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("xdvi499")
		}
	}

	return
}
