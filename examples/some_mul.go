package main

import "agorexx/lang"

func main() {

	rexsult, err := lang.RxFromString("2").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddmul000")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("ddmul000")
		}
	}
	rexsult, err = lang.RxFromString("2").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddmul001")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("ddmul001")
		}
	}
	rexsult, err = lang.RxFromString("5").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul002")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("ddmul002")
		}
	}
	rexsult, err = lang.RxFromString("5").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddmul003")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("ddmul003")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddmul004")
	} else {
		if !(rexsult.ToString() == "2.40") {
			lang.SayString("ddmul004")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddmul005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul005")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("ddmul006")
	} else {
		if !(rexsult.ToString() == "-2.40") {
			lang.SayString("ddmul006")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddmul007")
	} else {
		if !(rexsult.ToString() == "-2.40") {
			lang.SayString("ddmul007")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddmul008")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul008")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("ddmul009")
	} else {
		if !(rexsult.ToString() == "2.40") {
			lang.SayString("ddmul009")
		}
	}
	rexsult, err = lang.RxFromString("5.09").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("7.1"))
	if err != nil {
		lang.SayString("ddmul010")
	} else {
		if !(rexsult.ToString() == "36.139") {
			lang.SayString("ddmul010")
		}
	}
	rexsult, err = lang.RxFromString("2.5").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddmul011")
	} else {
		if !(rexsult.ToString() == "10.0") {
			lang.SayString("ddmul011")
		}
	}
	rexsult, err = lang.RxFromString("2.50").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddmul012")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("ddmul012")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.00000000"))
	if err != nil {
		lang.SayString("ddmul013")
	} else {
		if !(rexsult.ToString() == "1.234567890000000") {
			lang.SayString("ddmul013")
		}
	}
	rexsult, err = lang.RxFromString("2.50").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("ddmul015")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("ddmul015")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9.999999999"))
	if err != nil {
		lang.SayString("ddmul016")
	} else {
		if !(rexsult.ToString() == "99.99999998000000") {
			lang.SayString("ddmul016")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-9.999999999"))
	if err != nil {
		lang.SayString("ddmul017")
	} else {
		if !(rexsult.ToString() == "-99.99999998000000") {
			lang.SayString("ddmul017")
		}
	}
	rexsult, err = lang.RxFromString("-9.999999999").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9.999999999"))
	if err != nil {
		lang.SayString("ddmul018")
	} else {
		if !(rexsult.ToString() == "-99.99999998000000") {
			lang.SayString("ddmul018")
		}
	}
	rexsult, err = lang.RxFromString("-9.999999999").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-9.999999999"))
	if err != nil {
		lang.SayString("ddmul019")
	} else {
		if !(rexsult.ToString() == "99.99999998000000") {
			lang.SayString("ddmul019")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddmul021")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul021")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("ddmul022")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul022")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddmul023")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul023")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("ddmul024")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul024")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("ddmul025")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul025")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("ddmul026")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul026")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("ddmul027")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul027")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("ddmul028")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul028")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("ddmul030")
	} else {
		if !(rexsult.ToString() == "0.00500") {
			lang.SayString("ddmul030")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.000"))
	if err != nil {
		lang.SayString("ddmul031")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul031")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0E-3"))
	if err != nil {
		lang.SayString("ddmul032")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul032")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("00.00"))
	if err != nil {
		lang.SayString("ddmul033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul033")
		}
	}
	rexsult, err = lang.RxFromString("-5.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("ddmul034")
	} else {
		if !(rexsult.ToString() == "-0.00500") {
			lang.SayString("ddmul034")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.000"))
	if err != nil {
		lang.SayString("ddmul035")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul035")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0E-3"))
	if err != nil {
		lang.SayString("ddmul036")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul036")
		}
	}
	rexsult, err = lang.RxFromString("-0E-3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("00.00"))
	if err != nil {
		lang.SayString("ddmul037")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul037")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1E-3"))
	if err != nil {
		lang.SayString("ddmul038")
	} else {
		if !(rexsult.ToString() == "-0.00500") {
			lang.SayString("ddmul038")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.000"))
	if err != nil {
		lang.SayString("ddmul039")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul039")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0E-3"))
	if err != nil {
		lang.SayString("ddmul040")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul040")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-00.00"))
	if err != nil {
		lang.SayString("ddmul041")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul041")
		}
	}
	rexsult, err = lang.RxFromString("-5.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1E-3"))
	if err != nil {
		lang.SayString("ddmul042")
	} else {
		if !(rexsult.ToString() == "0.00500") {
			lang.SayString("ddmul042")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.000"))
	if err != nil {
		lang.SayString("ddmul043")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul043")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0E-3"))
	if err != nil {
		lang.SayString("ddmul044")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul044")
		}
	}
	rexsult, err = lang.RxFromString("-0E-3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-00.00"))
	if err != nil {
		lang.SayString("ddmul045")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul045")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddmul050")
	} else {
		if !(rexsult.ToString() == "3.60") {
			lang.SayString("ddmul050")
		}
	}
	rexsult, err = lang.RxFromString("7").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddmul051")
	} else {
		if !(rexsult.ToString() == "21") {
			lang.SayString("ddmul051")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.8"))
	if err != nil {
		lang.SayString("ddmul052")
	} else {
		if !(rexsult.ToString() == "0.72") {
			lang.SayString("ddmul052")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("ddmul053")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul053")
		}
	}
	rexsult, err = lang.RxFromString("654321").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("654321"))
	if err != nil {
		lang.SayString("ddmul054")
	} else {
		if !(rexsult.ToString() == "428135971041") {
			lang.SayString("ddmul054")
		}
	}
	rexsult, err = lang.RxFromString("123.45").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+9"))
	if err != nil {
		lang.SayString("ddmul062")
	} else {
		if !(rexsult.ToString() == "123450000000") {
			lang.SayString("ddmul062")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1230123456456789"))
	if err != nil {
		lang.SayString("ddmul080")
	} else {
		if !(rexsult.ToString() == "123012345645678.9") {
			lang.SayString("ddmul080")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1230123456456789"))
	if err != nil {
		lang.SayString("ddmul084")
	} else {
		if !(rexsult.ToString() == "123012345645678.9") {
			lang.SayString("ddmul084")
		}
	}
	rexsult, err = lang.RxFromString("1230123456456789").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("ddmul090")
	} else {
		if !(rexsult.ToString() == "123012345645678.9") {
			lang.SayString("ddmul090")
		}
	}
	rexsult, err = lang.RxFromString("1230123456456789").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("ddmul094")
	} else {
		if !(rexsult.ToString() == "123012345645678.9") {
			lang.SayString("ddmul094")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("ddmul101")
	} else {
		if !(rexsult.ToString() == "81") {
			lang.SayString("ddmul101")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90"))
	if err != nil {
		lang.SayString("ddmul102")
	} else {
		if !(rexsult.ToString() == "810") {
			lang.SayString("ddmul102")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900"))
	if err != nil {
		lang.SayString("ddmul103")
	} else {
		if !(rexsult.ToString() == "8100") {
			lang.SayString("ddmul103")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000"))
	if err != nil {
		lang.SayString("ddmul104")
	} else {
		if !(rexsult.ToString() == "81000") {
			lang.SayString("ddmul104")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000"))
	if err != nil {
		lang.SayString("ddmul105")
	} else {
		if !(rexsult.ToString() == "810000") {
			lang.SayString("ddmul105")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000"))
	if err != nil {
		lang.SayString("ddmul106")
	} else {
		if !(rexsult.ToString() == "8100000") {
			lang.SayString("ddmul106")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000"))
	if err != nil {
		lang.SayString("ddmul107")
	} else {
		if !(rexsult.ToString() == "81000000") {
			lang.SayString("ddmul107")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000"))
	if err != nil {
		lang.SayString("ddmul108")
	} else {
		if !(rexsult.ToString() == "810000000") {
			lang.SayString("ddmul108")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000000"))
	if err != nil {
		lang.SayString("ddmul109")
	} else {
		if !(rexsult.ToString() == "8100000000") {
			lang.SayString("ddmul109")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000"))
	if err != nil {
		lang.SayString("ddmul110")
	} else {
		if !(rexsult.ToString() == "81000000000") {
			lang.SayString("ddmul110")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000000"))
	if err != nil {
		lang.SayString("ddmul111")
	} else {
		if !(rexsult.ToString() == "810000000000") {
			lang.SayString("ddmul111")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000000000"))
	if err != nil {
		lang.SayString("ddmul112")
	} else {
		if !(rexsult.ToString() == "8100000000000") {
			lang.SayString("ddmul112")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000"))
	if err != nil {
		lang.SayString("ddmul113")
	} else {
		if !(rexsult.ToString() == "81000000000000") {
			lang.SayString("ddmul113")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000000000"))
	if err != nil {
		lang.SayString("ddmul114")
	} else {
		if !(rexsult.ToString() == "810000000000000") {
			lang.SayString("ddmul114")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000000000000"))
	if err != nil {
		lang.SayString("ddmul115")
	} else {
		if !(rexsult.ToString() == "8100000000000000") {
			lang.SayString("ddmul115")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("--ddmul116")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+16") {
			lang.SayString("--ddmul116")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000000000000"))
	if err != nil {
		lang.SayString("--ddmul117")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+17") {
			lang.SayString("--ddmul117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000000000000000"))
	if err != nil {
		lang.SayString("--ddmul118")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+18") {
			lang.SayString("--ddmul118")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000000"))
	if err != nil {
		lang.SayString("--ddmul119")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+19") {
			lang.SayString("--ddmul119")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000000000000000"))
	if err != nil {
		lang.SayString("--ddmul120")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+20") {
			lang.SayString("--ddmul120")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000000000000000000"))
	if err != nil {
		lang.SayString("--ddmul121")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+21") {
			lang.SayString("--ddmul121")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000000000"))
	if err != nil {
		lang.SayString("--ddmul122")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+22") {
			lang.SayString("--ddmul122")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000000000000000000"))
	if err != nil {
		lang.SayString("--ddmul123")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+23") {
			lang.SayString("--ddmul123")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddmul131")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("ddmul131")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("30"))
	if err != nil {
		lang.SayString("ddmul132")
	} else {
		if !(rexsult.ToString() == "90") {
			lang.SayString("ddmul132")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("300"))
	if err != nil {
		lang.SayString("ddmul133")
	} else {
		if !(rexsult.ToString() == "900") {
			lang.SayString("ddmul133")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("3000"))
	if err != nil {
		lang.SayString("ddmul134")
	} else {
		if !(rexsult.ToString() == "9000") {
			lang.SayString("ddmul134")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("30000"))
	if err != nil {
		lang.SayString("ddmul135")
	} else {
		if !(rexsult.ToString() == "90000") {
			lang.SayString("ddmul135")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("300000"))
	if err != nil {
		lang.SayString("ddmul136")
	} else {
		if !(rexsult.ToString() == "900000") {
			lang.SayString("ddmul136")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("3000000"))
	if err != nil {
		lang.SayString("ddmul137")
	} else {
		if !(rexsult.ToString() == "9000000") {
			lang.SayString("ddmul137")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("30000000"))
	if err != nil {
		lang.SayString("ddmul138")
	} else {
		if !(rexsult.ToString() == "90000000") {
			lang.SayString("ddmul138")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("300000000"))
	if err != nil {
		lang.SayString("ddmul139")
	} else {
		if !(rexsult.ToString() == "900000000") {
			lang.SayString("ddmul139")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("3000000000"))
	if err != nil {
		lang.SayString("ddmul140")
	} else {
		if !(rexsult.ToString() == "9000000000") {
			lang.SayString("ddmul140")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("30000000000"))
	if err != nil {
		lang.SayString("ddmul141")
	} else {
		if !(rexsult.ToString() == "90000000000") {
			lang.SayString("ddmul141")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("300000000000"))
	if err != nil {
		lang.SayString("ddmul142")
	} else {
		if !(rexsult.ToString() == "900000000000") {
			lang.SayString("ddmul142")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("3000000000000"))
	if err != nil {
		lang.SayString("ddmul143")
	} else {
		if !(rexsult.ToString() == "9000000000000") {
			lang.SayString("ddmul143")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("30000000000000"))
	if err != nil {
		lang.SayString("ddmul144")
	} else {
		if !(rexsult.ToString() == "90000000000000") {
			lang.SayString("ddmul144")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("300000000000000"))
	if err != nil {
		lang.SayString("ddmul145")
	} else {
		if !(rexsult.ToString() == "900000000000000") {
			lang.SayString("ddmul145")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("ddmul301")
	} else {
		if !(rexsult.ToString() == "81") {
			lang.SayString("ddmul301")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90"))
	if err != nil {
		lang.SayString("ddmul302")
	} else {
		if !(rexsult.ToString() == "810") {
			lang.SayString("ddmul302")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900"))
	if err != nil {
		lang.SayString("ddmul303")
	} else {
		if !(rexsult.ToString() == "8100") {
			lang.SayString("ddmul303")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000"))
	if err != nil {
		lang.SayString("ddmul304")
	} else {
		if !(rexsult.ToString() == "81000") {
			lang.SayString("ddmul304")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000"))
	if err != nil {
		lang.SayString("ddmul305")
	} else {
		if !(rexsult.ToString() == "810000") {
			lang.SayString("ddmul305")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000"))
	if err != nil {
		lang.SayString("ddmul306")
	} else {
		if !(rexsult.ToString() == "8100000") {
			lang.SayString("ddmul306")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000"))
	if err != nil {
		lang.SayString("ddmul307")
	} else {
		if !(rexsult.ToString() == "81000000") {
			lang.SayString("ddmul307")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000"))
	if err != nil {
		lang.SayString("ddmul308")
	} else {
		if !(rexsult.ToString() == "810000000") {
			lang.SayString("ddmul308")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000000"))
	if err != nil {
		lang.SayString("ddmul309")
	} else {
		if !(rexsult.ToString() == "8100000000") {
			lang.SayString("ddmul309")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000"))
	if err != nil {
		lang.SayString("ddmul310")
	} else {
		if !(rexsult.ToString() == "81000000000") {
			lang.SayString("ddmul310")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000000"))
	if err != nil {
		lang.SayString("ddmul311")
	} else {
		if !(rexsult.ToString() == "810000000000") {
			lang.SayString("ddmul311")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000000000"))
	if err != nil {
		lang.SayString("ddmul312")
	} else {
		if !(rexsult.ToString() == "8100000000000") {
			lang.SayString("ddmul312")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000"))
	if err != nil {
		lang.SayString("ddmul313")
	} else {
		if !(rexsult.ToString() == "81000000000000") {
			lang.SayString("ddmul313")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("90000000000000"))
	if err != nil {
		lang.SayString("ddmul314")
	} else {
		if !(rexsult.ToString() == "810000000000000") {
			lang.SayString("ddmul314")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("900000000000000"))
	if err != nil {
		lang.SayString("ddmul315")
	} else {
		if !(rexsult.ToString() == "8100000000000000") {
			lang.SayString("ddmul315")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("ddmul316")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+16") {
			lang.SayString("ddmul316")
		}
	}
	rexsult, err = lang.RxFromString("90").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("ddmul317")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+17") {
			lang.SayString("ddmul317")
		}
	}
	rexsult, err = lang.RxFromString("900").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("ddmul318")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+18") {
			lang.SayString("ddmul318")
		}
	}
	rexsult, err = lang.RxFromString("9000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("ddmul319")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+19") {
			lang.SayString("ddmul319")
		}
	}
	rexsult, err = lang.RxFromString("90000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("ddmul320")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+20") {
			lang.SayString("ddmul320")
		}
	}
	rexsult, err = lang.RxFromString("900000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("ddmul321")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+21") {
			lang.SayString("ddmul321")
		}
	}
	rexsult, err = lang.RxFromString("9000000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("ddmul322")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+22") {
			lang.SayString("ddmul322")
		}
	}
	rexsult, err = lang.RxFromString("90000000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("ddmul323")
	} else {
		if !(rexsult.ToString() == "8.100000000000000E+23") {
			lang.SayString("ddmul323")
		}
	}
	rexsult, err = lang.RxFromString("0E-260").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000E-260"))
	if err != nil {
		lang.SayString("ddmul504")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul504")
		}
	}
	rexsult, err = lang.RxFromString("100E+260").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0E+260"))
	if err != nil {
		lang.SayString("ddmul505")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul505")
		}
	}
	rexsult, err = lang.RxFromString("77.1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("850"))
	if err != nil {
		lang.SayString("ddmul506")
	} else {
		if !(rexsult.ToString() == "65535.0") {
			lang.SayString("ddmul506")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddmul541")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul541")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddmul542")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul542")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul543")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul543")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul544")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul544")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddmul545")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul545")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("ddmul546")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul546")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddmul547")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul547")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("ddmul548")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul548")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddmul551")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul551")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddmul552")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul552")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul553")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul553")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul554")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul554")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddmul555")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul555")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("ddmul556")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul556")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddmul557")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul557")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("ddmul558")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul558")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("ddmul561")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul561")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("ddmul562")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul562")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("ddmul563")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul563")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("ddmul564")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul564")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("ddmul565")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul565")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("ddmul566")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul566")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("ddmul567")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul567")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("ddmul568")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul568")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("ddmul571")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul571")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("ddmul572")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul572")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("ddmul573")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul573")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("ddmul574")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul574")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("ddmul575")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul575")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("ddmul576")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul576")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("ddmul577")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul577")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("ddmul578")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddmul578")
		}
	}
	rexsult, err = lang.RxFromString("1e+277").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("ddmul751")
	} else {
		if !(rexsult.ToString() == "1E+588") {
			lang.SayString("ddmul751")
		}
	}
	rexsult, err = lang.RxFromString("1e+277").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("ddmul752")
	} else {
		if !(rexsult.ToString() == "-1E+588") {
			lang.SayString("ddmul752")
		}
	}
	rexsult, err = lang.RxFromString("-1e+277").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("ddmul753")
	} else {
		if !(rexsult.ToString() == "-1E+588") {
			lang.SayString("ddmul753")
		}
	}
	rexsult, err = lang.RxFromString("-1e+277").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("ddmul754")
	} else {
		if !(rexsult.ToString() == "1E+588") {
			lang.SayString("ddmul754")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-311"))
	if err != nil {
		lang.SayString("ddmul755")
	} else {
		if !(rexsult.ToString() == "1E-588") {
			lang.SayString("ddmul755")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e-311"))
	if err != nil {
		lang.SayString("ddmul756")
	} else {
		if !(rexsult.ToString() == "-1E-588") {
			lang.SayString("ddmul756")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-311"))
	if err != nil {
		lang.SayString("ddmul757")
	} else {
		if !(rexsult.ToString() == "-1E-588") {
			lang.SayString("ddmul757")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e-311"))
	if err != nil {
		lang.SayString("ddmul758")
	} else {
		if !(rexsult.ToString() == "1E-588") {
			lang.SayString("ddmul758")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-101"))
	if err != nil {
		lang.SayString("ddmul760")
	} else {
		if !(rexsult.ToString() == "1E-392") {
			lang.SayString("ddmul760")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-102"))
	if err != nil {
		lang.SayString("ddmul761")
	} else {
		if !(rexsult.ToString() == "1E-393") {
			lang.SayString("ddmul761")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-103"))
	if err != nil {
		lang.SayString("ddmul762")
	} else {
		if !(rexsult.ToString() == "1E-394") {
			lang.SayString("ddmul762")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-104"))
	if err != nil {
		lang.SayString("ddmul763")
	} else {
		if !(rexsult.ToString() == "1E-395") {
			lang.SayString("ddmul763")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-105"))
	if err != nil {
		lang.SayString("ddmul764")
	} else {
		if !(rexsult.ToString() == "1E-396") {
			lang.SayString("ddmul764")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-106"))
	if err != nil {
		lang.SayString("ddmul765")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("ddmul765")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-107"))
	if err != nil {
		lang.SayString("ddmul766")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("ddmul766")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-108"))
	if err != nil {
		lang.SayString("ddmul767")
	} else {
		if !(rexsult.ToString() == "1E-399") {
			lang.SayString("ddmul767")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-109"))
	if err != nil {
		lang.SayString("ddmul768")
	} else {
		if !(rexsult.ToString() == "1E-400") {
			lang.SayString("ddmul768")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-110"))
	if err != nil {
		lang.SayString("ddmul769")
	} else {
		if !(rexsult.ToString() == "1E-401") {
			lang.SayString("ddmul769")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+321"))
	if err != nil {
		lang.SayString("ddmul770")
	} else {
		if !(rexsult.ToString() == "1E+381") {
			lang.SayString("ddmul770")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+322"))
	if err != nil {
		lang.SayString("ddmul771")
	} else {
		if !(rexsult.ToString() == "1E+382") {
			lang.SayString("ddmul771")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+323"))
	if err != nil {
		lang.SayString("ddmul772")
	} else {
		if !(rexsult.ToString() == "1E+383") {
			lang.SayString("ddmul772")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+324"))
	if err != nil {
		lang.SayString("ddmul773")
	} else {
		if !(rexsult.ToString() == "1E+384") {
			lang.SayString("ddmul773")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+325"))
	if err != nil {
		lang.SayString("ddmul774")
	} else {
		if !(rexsult.ToString() == "1E+385") {
			lang.SayString("ddmul774")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+326"))
	if err != nil {
		lang.SayString("ddmul775")
	} else {
		if !(rexsult.ToString() == "1E+386") {
			lang.SayString("ddmul775")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+327"))
	if err != nil {
		lang.SayString("ddmul776")
	} else {
		if !(rexsult.ToString() == "1E+387") {
			lang.SayString("ddmul776")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+328"))
	if err != nil {
		lang.SayString("ddmul777")
	} else {
		if !(rexsult.ToString() == "1E+388") {
			lang.SayString("ddmul777")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+329"))
	if err != nil {
		lang.SayString("ddmul778")
	} else {
		if !(rexsult.ToString() == "1E+389") {
			lang.SayString("ddmul778")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e+330"))
	if err != nil {
		lang.SayString("ddmul779")
	} else {
		if !(rexsult.ToString() == "1E+390") {
			lang.SayString("ddmul779")
		}
	}
	rexsult, err = lang.RxFromString("1.0000E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul801")
	} else {
		if !(rexsult.ToString() == "1.0000E-394") {
			lang.SayString("ddmul801")
		}
	}
	rexsult, err = lang.RxFromString("1.000E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-1"))
	if err != nil {
		lang.SayString("ddmul802")
	} else {
		if !(rexsult.ToString() == "1.000E-395") {
			lang.SayString("ddmul802")
		}
	}
	rexsult, err = lang.RxFromString("1.00E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-2"))
	if err != nil {
		lang.SayString("ddmul803")
	} else {
		if !(rexsult.ToString() == "1.00E-396") {
			lang.SayString("ddmul803")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-3"))
	if err != nil {
		lang.SayString("ddmul804")
	} else {
		if !(rexsult.ToString() == "1.0E-397") {
			lang.SayString("ddmul804")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul805")
	} else {
		if !(rexsult.ToString() == "1.0E-398") {
			lang.SayString("ddmul805")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul806")
	} else {
		if !(rexsult.ToString() == "1.3E-398") {
			lang.SayString("ddmul806")
		}
	}
	rexsult, err = lang.RxFromString("1.5E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul807")
	} else {
		if !(rexsult.ToString() == "1.5E-398") {
			lang.SayString("ddmul807")
		}
	}
	rexsult, err = lang.RxFromString("1.7E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul808")
	} else {
		if !(rexsult.ToString() == "1.7E-398") {
			lang.SayString("ddmul808")
		}
	}
	rexsult, err = lang.RxFromString("2.3E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul809")
	} else {
		if !(rexsult.ToString() == "2.3E-398") {
			lang.SayString("ddmul809")
		}
	}
	rexsult, err = lang.RxFromString("2.5E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul810")
	} else {
		if !(rexsult.ToString() == "2.5E-398") {
			lang.SayString("ddmul810")
		}
	}
	rexsult, err = lang.RxFromString("2.7E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul811")
	} else {
		if !(rexsult.ToString() == "2.7E-398") {
			lang.SayString("ddmul811")
		}
	}
	rexsult, err = lang.RxFromString("1.49E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul812")
	} else {
		if !(rexsult.ToString() == "1.49E-398") {
			lang.SayString("ddmul812")
		}
	}
	rexsult, err = lang.RxFromString("1.50E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul813")
	} else {
		if !(rexsult.ToString() == "1.50E-398") {
			lang.SayString("ddmul813")
		}
	}
	rexsult, err = lang.RxFromString("1.51E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul814")
	} else {
		if !(rexsult.ToString() == "1.51E-398") {
			lang.SayString("ddmul814")
		}
	}
	rexsult, err = lang.RxFromString("2.49E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul815")
	} else {
		if !(rexsult.ToString() == "2.49E-398") {
			lang.SayString("ddmul815")
		}
	}
	rexsult, err = lang.RxFromString("2.50E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul816")
	} else {
		if !(rexsult.ToString() == "2.50E-398") {
			lang.SayString("ddmul816")
		}
	}
	rexsult, err = lang.RxFromString("2.51E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul817")
	} else {
		if !(rexsult.ToString() == "2.51E-398") {
			lang.SayString("ddmul817")
		}
	}
	rexsult, err = lang.RxFromString("1E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("ddmul818")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("ddmul818")
		}
	}
	rexsult, err = lang.RxFromString("3E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("ddmul819")
	} else {
		if !(rexsult.ToString() == "3E-399") {
			lang.SayString("ddmul819")
		}
	}
	rexsult, err = lang.RxFromString("5E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("ddmul820")
	} else {
		if !(rexsult.ToString() == "5E-399") {
			lang.SayString("ddmul820")
		}
	}
	rexsult, err = lang.RxFromString("7E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("ddmul821")
	} else {
		if !(rexsult.ToString() == "7E-399") {
			lang.SayString("ddmul821")
		}
	}
	rexsult, err = lang.RxFromString("9E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("ddmul822")
	} else {
		if !(rexsult.ToString() == "9E-399") {
			lang.SayString("ddmul822")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("ddmul823")
	} else {
		if !(rexsult.ToString() == "9.9E-399") {
			lang.SayString("ddmul823")
		}
	}
	rexsult, err = lang.RxFromString("1E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e-4"))
	if err != nil {
		lang.SayString("ddmul824")
	} else {
		if !(rexsult.ToString() == "-1E-398") {
			lang.SayString("ddmul824")
		}
	}
	rexsult, err = lang.RxFromString("3E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("ddmul825")
	} else {
		if !(rexsult.ToString() == "-3E-399") {
			lang.SayString("ddmul825")
		}
	}
	rexsult, err = lang.RxFromString("-5E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("ddmul826")
	} else {
		if !(rexsult.ToString() == "-5E-399") {
			lang.SayString("ddmul826")
		}
	}
	rexsult, err = lang.RxFromString("7E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("ddmul827")
	} else {
		if !(rexsult.ToString() == "-7E-399") {
			lang.SayString("ddmul827")
		}
	}
	rexsult, err = lang.RxFromString("-9E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("ddmul828")
	} else {
		if !(rexsult.ToString() == "-9E-399") {
			lang.SayString("ddmul828")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("ddmul829")
	} else {
		if !(rexsult.ToString() == "-9.9E-399") {
			lang.SayString("ddmul829")
		}
	}
	rexsult, err = lang.RxFromString("3.0E-394").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("ddmul830")
	} else {
		if !(rexsult.ToString() == "-3.0E-399") {
			lang.SayString("ddmul830")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-200"))
	if err != nil {
		lang.SayString("ddmul831")
	} else {
		if !(rexsult.ToString() == "1.0E-399") {
			lang.SayString("ddmul831")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-199"))
	if err != nil {
		lang.SayString("ddmul832")
	} else {
		if !(rexsult.ToString() == "1.0E-398") {
			lang.SayString("ddmul832")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1e-198"))
	if err != nil {
		lang.SayString("ddmul833")
	} else {
		if !(rexsult.ToString() == "1.0E-397") {
			lang.SayString("ddmul833")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("2e-198"))
	if err != nil {
		lang.SayString("ddmul834")
	} else {
		if !(rexsult.ToString() == "4.0E-397") {
			lang.SayString("ddmul834")
		}
	}
	rexsult, err = lang.RxFromString("4.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("4e-198"))
	if err != nil {
		lang.SayString("ddmul835")
	} else {
		if !(rexsult.ToString() == "1.60E-396") {
			lang.SayString("ddmul835")
		}
	}
	rexsult, err = lang.RxFromString("10.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10e-198"))
	if err != nil {
		lang.SayString("ddmul836")
	} else {
		if !(rexsult.ToString() == "1.000E-395") {
			lang.SayString("ddmul836")
		}
	}
	rexsult, err = lang.RxFromString("30.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("30e-198"))
	if err != nil {
		lang.SayString("ddmul837")
	} else {
		if !(rexsult.ToString() == "9.000E-395") {
			lang.SayString("ddmul837")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("40e-188"))
	if err != nil {
		lang.SayString("ddmul838")
	} else {
		if !(rexsult.ToString() == "1.6000E-384") {
			lang.SayString("ddmul838")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("40e-187"))
	if err != nil {
		lang.SayString("ddmul839")
	} else {
		if !(rexsult.ToString() == "1.6000E-383") {
			lang.SayString("ddmul839")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-199").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("40e-186"))
	if err != nil {
		lang.SayString("ddmul840")
	} else {
		if !(rexsult.ToString() == "1.6000E-382") {
			lang.SayString("ddmul840")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9.999E+383"))
	if err != nil {
		lang.SayString("ddmul870")
	} else {
		if !(rexsult.ToString() == "9.99900E+385") {
			lang.SayString("ddmul870")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("-9.999E+383"))
	if err != nil {
		lang.SayString("ddmul871")
	} else {
		if !(rexsult.ToString() == "-9.99900E+385") {
			lang.SayString("ddmul871")
		}
	}
	rexsult, err = lang.RxFromString("9.999E+383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddmul872")
	} else {
		if !(rexsult.ToString() == "9.99900E+385") {
			lang.SayString("ddmul872")
		}
	}
	rexsult, err = lang.RxFromString("-9.999E+383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddmul873")
	} else {
		if !(rexsult.ToString() == "-9.99900E+385") {
			lang.SayString("ddmul873")
		}
	}
	rexsult, err = lang.RxFromString("1.2347E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.2347E-40"))
	if err != nil {
		lang.SayString("ddmul881")
	} else {
		if !(rexsult.ToString() == "1.52448409E-395") {
			lang.SayString("ddmul881")
		}
	}
	rexsult, err = lang.RxFromString("1.234E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("ddmul882")
	} else {
		if !(rexsult.ToString() == "1.522756E-395") {
			lang.SayString("ddmul882")
		}
	}
	rexsult, err = lang.RxFromString("1.23E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.23E-40"))
	if err != nil {
		lang.SayString("ddmul883")
	} else {
		if !(rexsult.ToString() == "1.5129E-395") {
			lang.SayString("ddmul883")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.2E-40"))
	if err != nil {
		lang.SayString("ddmul884")
	} else {
		if !(rexsult.ToString() == "1.44E-395") {
			lang.SayString("ddmul884")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.2E-41"))
	if err != nil {
		lang.SayString("ddmul885")
	} else {
		if !(rexsult.ToString() == "1.44E-396") {
			lang.SayString("ddmul885")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.2E-42"))
	if err != nil {
		lang.SayString("ddmul886")
	} else {
		if !(rexsult.ToString() == "1.44E-397") {
			lang.SayString("ddmul886")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.3E-42"))
	if err != nil {
		lang.SayString("ddmul887")
	} else {
		if !(rexsult.ToString() == "1.56E-397") {
			lang.SayString("ddmul887")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.3E-42"))
	if err != nil {
		lang.SayString("ddmul888")
	} else {
		if !(rexsult.ToString() == "1.69E-397") {
			lang.SayString("ddmul888")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-355").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.3E-43"))
	if err != nil {
		lang.SayString("ddmul889")
	} else {
		if !(rexsult.ToString() == "1.69E-398") {
			lang.SayString("ddmul889")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-356").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.3E-43"))
	if err != nil {
		lang.SayString("ddmul890")
	} else {
		if !(rexsult.ToString() == "1.69E-399") {
			lang.SayString("ddmul890")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-39").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.234E-355"))
	if err != nil {
		lang.SayString("ddmul891")
	} else {
		if !(rexsult.ToString() == "1.5233730E-394") {
			lang.SayString("ddmul891")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-39").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.234E-355"))
	if err != nil {
		lang.SayString("ddmul892")
	} else {
		if !(rexsult.ToString() == "1.52344704E-394") {
			lang.SayString("ddmul892")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-40").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.234E-355"))
	if err != nil {
		lang.SayString("ddmul893")
	} else {
		if !(rexsult.ToString() == "1.5233730E-395") {
			lang.SayString("ddmul893")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-40").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.234E-355"))
	if err != nil {
		lang.SayString("ddmul894")
	} else {
		if !(rexsult.ToString() == "1.52344704E-395") {
			lang.SayString("ddmul894")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-41").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.234E-355"))
	if err != nil {
		lang.SayString("ddmul895")
	} else {
		if !(rexsult.ToString() == "1.5233730E-396") {
			lang.SayString("ddmul895")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-41").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1.234E-355"))
	if err != nil {
		lang.SayString("ddmul896")
	} else {
		if !(rexsult.ToString() == "1.52344704E-396") {
			lang.SayString("ddmul896")
		}
	}
	rexsult, err = lang.RxFromString("0.3000000000E-191").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.3000000000E-191"))
	if err != nil {
		lang.SayString("ddmul900")
	} else {
		if !(rexsult.ToString() == "9.000000000000000E-384") {
			lang.SayString("ddmul900")
		}
	}
	rexsult, err = lang.RxFromString("0.3000000001E-191").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.3000000001E-191"))
	if err != nil {
		lang.SayString("ddmul901")
	} else {
		if !(rexsult.ToString() == "9.000000006000000E-384") {
			lang.SayString("ddmul901")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0999999999999"))
	if err != nil {
		lang.SayString("ddmul902")
	} else {
		if !(rexsult.ToString() == "9.999999999989999E-384") {
			lang.SayString("ddmul902")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.09999999999999"))
	if err != nil {
		lang.SayString("ddmul903")
	} else {
		if !(rexsult.ToString() == "9.999999999998999E-384") {
			lang.SayString("ddmul903")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.099999999999999"))
	if err != nil {
		lang.SayString("ddmul904")
	} else {
		if !(rexsult.ToString() == "9.999999999999899E-384") {
			lang.SayString("ddmul904")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0999999999999999"))
	if err != nil {
		lang.SayString("ddmul905")
	} else {
		if !(rexsult.ToString() == "9.999999999999989E-384") {
			lang.SayString("ddmul905")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.09999999999999999"))
	if err != nil {
		lang.SayString("ddmul906")
	} else {
		if !(rexsult.ToString() == "9.999999999999998E-384") {
			lang.SayString("ddmul906")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul907")
	} else {
		if !(rexsult.ToString() == "9.999999999999999E-383") {
			lang.SayString("ddmul907")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.09999999999999999"))
	if err != nil {
		lang.SayString("ddmul908")
	} else {
		if !(rexsult.ToString() == "0.09999999999999999") {
			lang.SayString("ddmul908")
		}
	}
	rexsult, err = lang.RxFromString("1e-398").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.99"))
	if err != nil {
		lang.SayString("ddmul910")
	} else {
		if !(rexsult.ToString() == "9.9E-399") {
			lang.SayString("ddmul910")
		}
	}
	rexsult, err = lang.RxFromString("1e-398").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.75"))
	if err != nil {
		lang.SayString("ddmul911")
	} else {
		if !(rexsult.ToString() == "7.5E-399") {
			lang.SayString("ddmul911")
		}
	}
	rexsult, err = lang.RxFromString("1e-398").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("ddmul912")
	} else {
		if !(rexsult.ToString() == "5E-399") {
			lang.SayString("ddmul912")
		}
	}
	rexsult, err = lang.RxFromString("1e-398").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.25"))
	if err != nil {
		lang.SayString("ddmul913")
	} else {
		if !(rexsult.ToString() == "2.5E-399") {
			lang.SayString("ddmul913")
		}
	}
	rexsult, err = lang.RxFromString("1e-398").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.01"))
	if err != nil {
		lang.SayString("ddmul914")
	} else {
		if !(rexsult.ToString() == "1E-400") {
			lang.SayString("ddmul914")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("9999999999999999"))
	if err != nil {
		lang.SayString("ddmul920")
	} else {
		if !(rexsult.ToString() == "9.999999999999998E+31") {
			lang.SayString("ddmul920")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddmul1001")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("ddmul1001")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddmul1002")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("ddmul1002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("ddmul1003")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("ddmul1003")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("ddmul1004")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("ddmul1004")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("ddmul1005")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("ddmul1005")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("ddmul1006")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("ddmul1006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("ddmul1007")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("ddmul1007")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("ddmul1008")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("ddmul1008")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("ddmul1009")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("ddmul1009")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("ddmul1010")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("ddmul1010")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("ddmul1011")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("ddmul1011")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000000"))
	if err != nil {
		lang.SayString("ddmul1012")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("ddmul1012")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000000000"))
	if err != nil {
		lang.SayString("ddmul1013")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("ddmul1013")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000000000"))
	if err != nil {
		lang.SayString("ddmul1014")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("ddmul1014")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000000000"))
	if err != nil {
		lang.SayString("ddmul1015")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("ddmul1015")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul1021")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("ddmul1021")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddmul1022")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("ddmul1022")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddmul1023")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("ddmul1023")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("ddmul1024")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("ddmul1024")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("ddmul1025")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("ddmul1025")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("ddmul1026")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("ddmul1026")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("ddmul1027")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("ddmul1027")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("ddmul1028")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("ddmul1028")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("ddmul1029")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("ddmul1029")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("ddmul1030")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("ddmul1030")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("ddmul1031")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("ddmul1031")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("ddmul1032")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("ddmul1032")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000000"))
	if err != nil {
		lang.SayString("ddmul1033")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("ddmul1033")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000000000"))
	if err != nil {
		lang.SayString("ddmul1034")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("ddmul1034")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000000000"))
	if err != nil {
		lang.SayString("ddmul1035")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("ddmul1035")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("ddmul1041")
	} else {
		if !(rexsult.ToString() == "10.0") {
			lang.SayString("ddmul1041")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul1042")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("ddmul1042")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddmul1043")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("ddmul1043")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddmul1044")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("ddmul1044")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("ddmul1045")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("ddmul1045")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("ddmul1046")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("ddmul1046")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("ddmul1047")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("ddmul1047")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("ddmul1048")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("ddmul1048")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("ddmul1049")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("ddmul1049")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("ddmul1050")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("ddmul1050")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("ddmul1051")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("ddmul1051")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("ddmul1052")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("ddmul1052")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("ddmul1053")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("ddmul1053")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000000"))
	if err != nil {
		lang.SayString("ddmul1054")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("ddmul1054")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000000000"))
	if err != nil {
		lang.SayString("ddmul1055")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("ddmul1055")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.01"))
	if err != nil {
		lang.SayString("ddmul1061")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("ddmul1061")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("ddmul1062")
	} else {
		if !(rexsult.ToString() == "100.0") {
			lang.SayString("ddmul1062")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul1063")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("ddmul1063")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddmul1064")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("ddmul1064")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddmul1065")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("ddmul1065")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("ddmul1066")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("ddmul1066")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("ddmul1067")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("ddmul1067")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("ddmul1068")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("ddmul1068")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("ddmul1069")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("ddmul1069")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("ddmul1070")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("ddmul1070")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("ddmul1071")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("ddmul1071")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("ddmul1072")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("ddmul1072")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("ddmul1073")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("ddmul1073")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("ddmul1074")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("ddmul1074")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000000"))
	if err != nil {
		lang.SayString("ddmul1075")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("ddmul1075")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.001"))
	if err != nil {
		lang.SayString("ddmul1081")
	} else {
		if !(rexsult.ToString() == "10.000") {
			lang.SayString("ddmul1081")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.01"))
	if err != nil {
		lang.SayString("ddmul1082")
	} else {
		if !(rexsult.ToString() == "100.00") {
			lang.SayString("ddmul1082")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("ddmul1083")
	} else {
		if !(rexsult.ToString() == "1000.0") {
			lang.SayString("ddmul1083")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddmul1084")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("ddmul1084")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("ddmul1085")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("ddmul1085")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("ddmul1086")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("ddmul1086")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("ddmul1087")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("ddmul1087")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("ddmul1088")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("ddmul1088")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("ddmul1089")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("ddmul1089")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("ddmul1090")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("ddmul1090")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("ddmul1091")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("ddmul1091")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("ddmul1092")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("ddmul1092")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("ddmul1093")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("ddmul1093")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("ddmul1094")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("ddmul1094")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("ddmul1095")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("ddmul1095")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("99999999999"))
	if err != nil {
		lang.SayString("ddmul1097")
	} else {
		if !(rexsult.ToString() == "999999999990000") {
			lang.SayString("ddmul1097")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("99999999999"))
	if err != nil {
		lang.SayString("ddmul1098")
	} else {
		if !(rexsult.ToString() == "999999999990000") {
			lang.SayString("ddmul1098")
		}
	}
	rexsult, err = lang.RxFromString("999999999999999").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("999999999999999"))
	if err != nil {
		lang.SayString("ddm2")
	} else {
		if !(rexsult.ToString() == "9.999999999999980E+29") {
			lang.SayString("ddm2")
		}
	}
	rexsult, err = lang.RxFromString("2").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqmul000")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("dqmul000")
		}
	}
	rexsult, err = lang.RxFromString("2").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqmul001")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dqmul001")
		}
	}
	rexsult, err = lang.RxFromString("5").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul002")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("dqmul002")
		}
	}
	rexsult, err = lang.RxFromString("5").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqmul003")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqmul003")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqmul004")
	} else {
		if !(rexsult.ToString() == "2.40") {
			lang.SayString("dqmul004")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqmul005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul005")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqmul006")
	} else {
		if !(rexsult.ToString() == "-2.40") {
			lang.SayString("dqmul006")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqmul007")
	} else {
		if !(rexsult.ToString() == "-2.40") {
			lang.SayString("dqmul007")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqmul008")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul008")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqmul009")
	} else {
		if !(rexsult.ToString() == "2.40") {
			lang.SayString("dqmul009")
		}
	}
	rexsult, err = lang.RxFromString("5.09").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("7.1"))
	if err != nil {
		lang.SayString("dqmul010")
	} else {
		if !(rexsult.ToString() == "36.139") {
			lang.SayString("dqmul010")
		}
	}
	rexsult, err = lang.RxFromString("2.5").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqmul011")
	} else {
		if !(rexsult.ToString() == "10.0") {
			lang.SayString("dqmul011")
		}
	}
	rexsult, err = lang.RxFromString("2.50").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqmul012")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("dqmul012")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.0000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul013")
	} else {
		if !(rexsult.ToString() == "1.234567890000000000000000000000000") {
			lang.SayString("dqmul013")
		}
	}
	rexsult, err = lang.RxFromString("2.50").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqmul015")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("dqmul015")
		}
	}
	rexsult, err = lang.RxFromString("9.99999999999999999").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9.99999999999999999"))
	if err != nil {
		lang.SayString("dqmul016")
	} else {
		if !(rexsult.ToString() == "99.99999999999999980000000000000000") {
			lang.SayString("dqmul016")
		}
	}
	rexsult, err = lang.RxFromString("9.99999999999999999").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-9.99999999999999999"))
	if err != nil {
		lang.SayString("dqmul017")
	} else {
		if !(rexsult.ToString() == "-99.99999999999999980000000000000000") {
			lang.SayString("dqmul017")
		}
	}
	rexsult, err = lang.RxFromString("-9.99999999999999999").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9.99999999999999999"))
	if err != nil {
		lang.SayString("dqmul018")
	} else {
		if !(rexsult.ToString() == "-99.99999999999999980000000000000000") {
			lang.SayString("dqmul018")
		}
	}
	rexsult, err = lang.RxFromString("-9.99999999999999999").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-9.99999999999999999"))
	if err != nil {
		lang.SayString("dqmul019")
	} else {
		if !(rexsult.ToString() == "99.99999999999999980000000000000000") {
			lang.SayString("dqmul019")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqmul021")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul021")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("dqmul022")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul022")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqmul023")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul023")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("dqmul024")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul024")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("dqmul025")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul025")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("dqmul026")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul026")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("dqmul027")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul027")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("dqmul028")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul028")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("dqmul030")
	} else {
		if !(rexsult.ToString() == "0.00500") {
			lang.SayString("dqmul030")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.000"))
	if err != nil {
		lang.SayString("dqmul031")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul031")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0E-3"))
	if err != nil {
		lang.SayString("dqmul032")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul032")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("00.00"))
	if err != nil {
		lang.SayString("dqmul033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul033")
		}
	}
	rexsult, err = lang.RxFromString("-5.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("dqmul034")
	} else {
		if !(rexsult.ToString() == "-0.00500") {
			lang.SayString("dqmul034")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.000"))
	if err != nil {
		lang.SayString("dqmul035")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul035")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0E-3"))
	if err != nil {
		lang.SayString("dqmul036")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul036")
		}
	}
	rexsult, err = lang.RxFromString("-0E-3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("00.00"))
	if err != nil {
		lang.SayString("dqmul037")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul037")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1E-3"))
	if err != nil {
		lang.SayString("dqmul038")
	} else {
		if !(rexsult.ToString() == "-0.00500") {
			lang.SayString("dqmul038")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.000"))
	if err != nil {
		lang.SayString("dqmul039")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul039")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0E-3"))
	if err != nil {
		lang.SayString("dqmul040")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul040")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-00.00"))
	if err != nil {
		lang.SayString("dqmul041")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul041")
		}
	}
	rexsult, err = lang.RxFromString("-5.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1E-3"))
	if err != nil {
		lang.SayString("dqmul042")
	} else {
		if !(rexsult.ToString() == "0.00500") {
			lang.SayString("dqmul042")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.000"))
	if err != nil {
		lang.SayString("dqmul043")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul043")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0E-3"))
	if err != nil {
		lang.SayString("dqmul044")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul044")
		}
	}
	rexsult, err = lang.RxFromString("-0E-3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-00.00"))
	if err != nil {
		lang.SayString("dqmul045")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul045")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqmul050")
	} else {
		if !(rexsult.ToString() == "3.60") {
			lang.SayString("dqmul050")
		}
	}
	rexsult, err = lang.RxFromString("7").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqmul051")
	} else {
		if !(rexsult.ToString() == "21") {
			lang.SayString("dqmul051")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.8"))
	if err != nil {
		lang.SayString("dqmul052")
	} else {
		if !(rexsult.ToString() == "0.72") {
			lang.SayString("dqmul052")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("dqmul053")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul053")
		}
	}
	rexsult, err = lang.RxFromString("654321").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("654321"))
	if err != nil {
		lang.SayString("dqmul054")
	} else {
		if !(rexsult.ToString() == "428135971041") {
			lang.SayString("dqmul054")
		}
	}
	rexsult, err = lang.RxFromString("123.45").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+9"))
	if err != nil {
		lang.SayString("dqmul062")
	} else {
		if !(rexsult.ToString() == "123450000000") {
			lang.SayString("dqmul062")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1230123456456789"))
	if err != nil {
		lang.SayString("dqmul080")
	} else {
		if !(rexsult.ToString() == "123012345645678.9") {
			lang.SayString("dqmul080")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1230123456456789"))
	if err != nil {
		lang.SayString("dqmul084")
	} else {
		if !(rexsult.ToString() == "123012345645678.9") {
			lang.SayString("dqmul084")
		}
	}
	rexsult, err = lang.RxFromString("1230123456456789").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqmul090")
	} else {
		if !(rexsult.ToString() == "123012345645678.9") {
			lang.SayString("dqmul090")
		}
	}
	rexsult, err = lang.RxFromString("1230123456456789").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqmul094")
	} else {
		if !(rexsult.ToString() == "123012345645678.9") {
			lang.SayString("dqmul094")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dqmul101")
	} else {
		if !(rexsult.ToString() == "81") {
			lang.SayString("dqmul101")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90"))
	if err != nil {
		lang.SayString("dqmul102")
	} else {
		if !(rexsult.ToString() == "810") {
			lang.SayString("dqmul102")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900"))
	if err != nil {
		lang.SayString("dqmul103")
	} else {
		if !(rexsult.ToString() == "8100") {
			lang.SayString("dqmul103")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000"))
	if err != nil {
		lang.SayString("dqmul104")
	} else {
		if !(rexsult.ToString() == "81000") {
			lang.SayString("dqmul104")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000"))
	if err != nil {
		lang.SayString("dqmul105")
	} else {
		if !(rexsult.ToString() == "810000") {
			lang.SayString("dqmul105")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000"))
	if err != nil {
		lang.SayString("dqmul106")
	} else {
		if !(rexsult.ToString() == "8100000") {
			lang.SayString("dqmul106")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000"))
	if err != nil {
		lang.SayString("dqmul107")
	} else {
		if !(rexsult.ToString() == "81000000") {
			lang.SayString("dqmul107")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000"))
	if err != nil {
		lang.SayString("dqmul108")
	} else {
		if !(rexsult.ToString() == "810000000") {
			lang.SayString("dqmul108")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000000"))
	if err != nil {
		lang.SayString("dqmul109")
	} else {
		if !(rexsult.ToString() == "8100000000") {
			lang.SayString("dqmul109")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000"))
	if err != nil {
		lang.SayString("dqmul110")
	} else {
		if !(rexsult.ToString() == "81000000000") {
			lang.SayString("dqmul110")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000000"))
	if err != nil {
		lang.SayString("dqmul111")
	} else {
		if !(rexsult.ToString() == "810000000000") {
			lang.SayString("dqmul111")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000000000"))
	if err != nil {
		lang.SayString("dqmul112")
	} else {
		if !(rexsult.ToString() == "8100000000000") {
			lang.SayString("dqmul112")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000"))
	if err != nil {
		lang.SayString("dqmul113")
	} else {
		if !(rexsult.ToString() == "81000000000000") {
			lang.SayString("dqmul113")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000000000"))
	if err != nil {
		lang.SayString("dqmul114")
	} else {
		if !(rexsult.ToString() == "810000000000000") {
			lang.SayString("dqmul114")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000000000000"))
	if err != nil {
		lang.SayString("dqmul115")
	} else {
		if !(rexsult.ToString() == "8100000000000000") {
			lang.SayString("dqmul115")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("--dqmul116")
	} else {
		if !(rexsult.ToString() == "81000000000000000") {
			lang.SayString("--dqmul116")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000000000000"))
	if err != nil {
		lang.SayString("--dqmul117")
	} else {
		if !(rexsult.ToString() == "810000000000000000") {
			lang.SayString("--dqmul117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000000000000000"))
	if err != nil {
		lang.SayString("--dqmul118")
	} else {
		if !(rexsult.ToString() == "8100000000000000000") {
			lang.SayString("--dqmul118")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000000"))
	if err != nil {
		lang.SayString("--dqmul119")
	} else {
		if !(rexsult.ToString() == "81000000000000000000") {
			lang.SayString("--dqmul119")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000000000000000"))
	if err != nil {
		lang.SayString("--dqmul120")
	} else {
		if !(rexsult.ToString() == "810000000000000000000") {
			lang.SayString("--dqmul120")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000000000000000000"))
	if err != nil {
		lang.SayString("--dqmul121")
	} else {
		if !(rexsult.ToString() == "8100000000000000000000") {
			lang.SayString("--dqmul121")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000000000"))
	if err != nil {
		lang.SayString("--dqmul122")
	} else {
		if !(rexsult.ToString() == "81000000000000000000000") {
			lang.SayString("--dqmul122")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000000000000000000"))
	if err != nil {
		lang.SayString("--dqmul123")
	} else {
		if !(rexsult.ToString() == "810000000000000000000000") {
			lang.SayString("--dqmul123")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqmul131")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dqmul131")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30"))
	if err != nil {
		lang.SayString("dqmul132")
	} else {
		if !(rexsult.ToString() == "90") {
			lang.SayString("dqmul132")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("300"))
	if err != nil {
		lang.SayString("dqmul133")
	} else {
		if !(rexsult.ToString() == "900") {
			lang.SayString("dqmul133")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3000"))
	if err != nil {
		lang.SayString("dqmul134")
	} else {
		if !(rexsult.ToString() == "9000") {
			lang.SayString("dqmul134")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30000"))
	if err != nil {
		lang.SayString("dqmul135")
	} else {
		if !(rexsult.ToString() == "90000") {
			lang.SayString("dqmul135")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("300000"))
	if err != nil {
		lang.SayString("dqmul136")
	} else {
		if !(rexsult.ToString() == "900000") {
			lang.SayString("dqmul136")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3000000"))
	if err != nil {
		lang.SayString("dqmul137")
	} else {
		if !(rexsult.ToString() == "9000000") {
			lang.SayString("dqmul137")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30000000"))
	if err != nil {
		lang.SayString("dqmul138")
	} else {
		if !(rexsult.ToString() == "90000000") {
			lang.SayString("dqmul138")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("300000000"))
	if err != nil {
		lang.SayString("dqmul139")
	} else {
		if !(rexsult.ToString() == "900000000") {
			lang.SayString("dqmul139")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3000000000"))
	if err != nil {
		lang.SayString("dqmul140")
	} else {
		if !(rexsult.ToString() == "9000000000") {
			lang.SayString("dqmul140")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30000000000"))
	if err != nil {
		lang.SayString("dqmul141")
	} else {
		if !(rexsult.ToString() == "90000000000") {
			lang.SayString("dqmul141")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("300000000000"))
	if err != nil {
		lang.SayString("dqmul142")
	} else {
		if !(rexsult.ToString() == "900000000000") {
			lang.SayString("dqmul142")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3000000000000"))
	if err != nil {
		lang.SayString("dqmul143")
	} else {
		if !(rexsult.ToString() == "9000000000000") {
			lang.SayString("dqmul143")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30000000000000"))
	if err != nil {
		lang.SayString("dqmul144")
	} else {
		if !(rexsult.ToString() == "90000000000000") {
			lang.SayString("dqmul144")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("300000000000000"))
	if err != nil {
		lang.SayString("dqmul145")
	} else {
		if !(rexsult.ToString() == "900000000000000") {
			lang.SayString("dqmul145")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3000000000000000"))
	if err != nil {
		lang.SayString("dqmul146")
	} else {
		if !(rexsult.ToString() == "9000000000000000") {
			lang.SayString("dqmul146")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30000000000000000"))
	if err != nil {
		lang.SayString("dqmul147")
	} else {
		if !(rexsult.ToString() == "90000000000000000") {
			lang.SayString("dqmul147")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("300000000000000000"))
	if err != nil {
		lang.SayString("dqmul148")
	} else {
		if !(rexsult.ToString() == "900000000000000000") {
			lang.SayString("dqmul148")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3000000000000000000"))
	if err != nil {
		lang.SayString("dqmul149")
	} else {
		if !(rexsult.ToString() == "9000000000000000000") {
			lang.SayString("dqmul149")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30000000000000000000"))
	if err != nil {
		lang.SayString("dqmul150")
	} else {
		if !(rexsult.ToString() == "90000000000000000000") {
			lang.SayString("dqmul150")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("300000000000000000000"))
	if err != nil {
		lang.SayString("dqmul151")
	} else {
		if !(rexsult.ToString() == "900000000000000000000") {
			lang.SayString("dqmul151")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("3000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul152")
	} else {
		if !(rexsult.ToString() == "9000000000000000000000") {
			lang.SayString("dqmul152")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul153")
	} else {
		if !(rexsult.ToString() == "90000000000000000000000") {
			lang.SayString("dqmul153")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("dqmul263")
	} else {
		if !(rexsult.ToString() == "145433.2908011933696719165119928296") {
			lang.SayString("dqmul263")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dqmul301")
	} else {
		if !(rexsult.ToString() == "8100000000000000000") {
			lang.SayString("dqmul301")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90"))
	if err != nil {
		lang.SayString("dqmul302")
	} else {
		if !(rexsult.ToString() == "81000000000000000000") {
			lang.SayString("dqmul302")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900"))
	if err != nil {
		lang.SayString("dqmul303")
	} else {
		if !(rexsult.ToString() == "810000000000000000000") {
			lang.SayString("dqmul303")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000"))
	if err != nil {
		lang.SayString("dqmul304")
	} else {
		if !(rexsult.ToString() == "8100000000000000000000") {
			lang.SayString("dqmul304")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000"))
	if err != nil {
		lang.SayString("dqmul305")
	} else {
		if !(rexsult.ToString() == "81000000000000000000000") {
			lang.SayString("dqmul305")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000"))
	if err != nil {
		lang.SayString("dqmul306")
	} else {
		if !(rexsult.ToString() == "810000000000000000000000") {
			lang.SayString("dqmul306")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000"))
	if err != nil {
		lang.SayString("dqmul307")
	} else {
		if !(rexsult.ToString() == "8100000000000000000000000") {
			lang.SayString("dqmul307")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000"))
	if err != nil {
		lang.SayString("dqmul308")
	} else {
		if !(rexsult.ToString() == "81000000000000000000000000") {
			lang.SayString("dqmul308")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000000"))
	if err != nil {
		lang.SayString("dqmul309")
	} else {
		if !(rexsult.ToString() == "810000000000000000000000000") {
			lang.SayString("dqmul309")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000"))
	if err != nil {
		lang.SayString("dqmul310")
	} else {
		if !(rexsult.ToString() == "8100000000000000000000000000") {
			lang.SayString("dqmul310")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000000"))
	if err != nil {
		lang.SayString("dqmul311")
	} else {
		if !(rexsult.ToString() == "81000000000000000000000000000") {
			lang.SayString("dqmul311")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000000000"))
	if err != nil {
		lang.SayString("dqmul312")
	} else {
		if !(rexsult.ToString() == "810000000000000000000000000000") {
			lang.SayString("dqmul312")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000"))
	if err != nil {
		lang.SayString("dqmul313")
	} else {
		if !(rexsult.ToString() == "8100000000000000000000000000000") {
			lang.SayString("dqmul313")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("90000000000000"))
	if err != nil {
		lang.SayString("dqmul314")
	} else {
		if !(rexsult.ToString() == "81000000000000000000000000000000") {
			lang.SayString("dqmul314")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("900000000000000"))
	if err != nil {
		lang.SayString("dqmul315")
	} else {
		if !(rexsult.ToString() == "810000000000000000000000000000000") {
			lang.SayString("dqmul315")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("dqmul316")
	} else {
		if !(rexsult.ToString() == "8100000000000000000000000000000000") {
			lang.SayString("dqmul316")
		}
	}
	rexsult, err = lang.RxFromString("9000000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("dqmul317")
	} else {
		if !(rexsult.ToString() == "8.100000000000000000000000000000000E+34") {
			lang.SayString("dqmul317")
		}
	}
	rexsult, err = lang.RxFromString("90000000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("dqmul318")
	} else {
		if !(rexsult.ToString() == "8.100000000000000000000000000000000E+35") {
			lang.SayString("dqmul318")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("dqmul319")
	} else {
		if !(rexsult.ToString() == "8.100000000000000000000000000000000E+36") {
			lang.SayString("dqmul319")
		}
	}
	rexsult, err = lang.RxFromString("9000000000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("dqmul320")
	} else {
		if !(rexsult.ToString() == "8.100000000000000000000000000000000E+37") {
			lang.SayString("dqmul320")
		}
	}
	rexsult, err = lang.RxFromString("90000000000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("dqmul321")
	} else {
		if !(rexsult.ToString() == "8.100000000000000000000000000000000E+38") {
			lang.SayString("dqmul321")
		}
	}
	rexsult, err = lang.RxFromString("900000000000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("dqmul322")
	} else {
		if !(rexsult.ToString() == "8.100000000000000000000000000000000E+39") {
			lang.SayString("dqmul322")
		}
	}
	rexsult, err = lang.RxFromString("9000000000000000000000000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("dqmul323")
	} else {
		if !(rexsult.ToString() == "8.100000000000000000000000000000000E+40") {
			lang.SayString("dqmul323")
		}
	}
	rexsult, err = lang.RxFromString("0E-4260").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000E-4260"))
	if err != nil {
		lang.SayString("dqmul504")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul504")
		}
	}
	rexsult, err = lang.RxFromString("100E+4260").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0E+4260"))
	if err != nil {
		lang.SayString("dqmul505")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul505")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqmul541")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul541")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqmul542")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul542")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul543")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul543")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul544")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul544")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqmul545")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul545")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("dqmul546")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul546")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqmul547")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul547")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("dqmul548")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul548")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqmul551")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul551")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqmul552")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul552")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul553")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul553")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul554")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul554")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqmul555")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul555")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("dqmul556")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul556")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqmul557")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul557")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("dqmul558")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul558")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqmul561")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul561")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqmul562")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul562")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqmul563")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul563")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqmul564")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul564")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("dqmul565")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul565")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("dqmul566")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul566")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("dqmul567")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul567")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("dqmul568")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul568")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqmul571")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul571")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqmul572")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul572")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqmul573")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul573")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqmul574")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul574")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("dqmul575")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul575")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("dqmul576")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul576")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("dqmul577")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul577")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("dqmul578")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqmul578")
		}
	}
	rexsult, err = lang.RxFromString("1e+4277").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+3311"))
	if err != nil {
		lang.SayString("dqmul751")
	} else {
		if !(rexsult.ToString() == "1E+7588") {
			lang.SayString("dqmul751")
		}
	}
	rexsult, err = lang.RxFromString("1e+4277").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e+3311"))
	if err != nil {
		lang.SayString("dqmul752")
	} else {
		if !(rexsult.ToString() == "-1E+7588") {
			lang.SayString("dqmul752")
		}
	}
	rexsult, err = lang.RxFromString("-1e+4277").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+3311"))
	if err != nil {
		lang.SayString("dqmul753")
	} else {
		if !(rexsult.ToString() == "-1E+7588") {
			lang.SayString("dqmul753")
		}
	}
	rexsult, err = lang.RxFromString("-1e+4277").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e+3311"))
	if err != nil {
		lang.SayString("dqmul754")
	} else {
		if !(rexsult.ToString() == "1E+7588") {
			lang.SayString("dqmul754")
		}
	}
	rexsult, err = lang.RxFromString("1e-4277").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-3311"))
	if err != nil {
		lang.SayString("dqmul755")
	} else {
		if !(rexsult.ToString() == "1E-7588") {
			lang.SayString("dqmul755")
		}
	}
	rexsult, err = lang.RxFromString("1e-4277").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e-3311"))
	if err != nil {
		lang.SayString("dqmul756")
	} else {
		if !(rexsult.ToString() == "-1E-7588") {
			lang.SayString("dqmul756")
		}
	}
	rexsult, err = lang.RxFromString("-1e-4277").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-3311"))
	if err != nil {
		lang.SayString("dqmul757")
	} else {
		if !(rexsult.ToString() == "-1E-7588") {
			lang.SayString("dqmul757")
		}
	}
	rexsult, err = lang.RxFromString("-1e-4277").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e-3311"))
	if err != nil {
		lang.SayString("dqmul758")
	} else {
		if !(rexsult.ToString() == "1E-7588") {
			lang.SayString("dqmul758")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-101"))
	if err != nil {
		lang.SayString("dqmul760")
	} else {
		if !(rexsult.ToString() == "1E-6170") {
			lang.SayString("dqmul760")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-102"))
	if err != nil {
		lang.SayString("dqmul761")
	} else {
		if !(rexsult.ToString() == "1E-6171") {
			lang.SayString("dqmul761")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-103"))
	if err != nil {
		lang.SayString("dqmul762")
	} else {
		if !(rexsult.ToString() == "1E-6172") {
			lang.SayString("dqmul762")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-104"))
	if err != nil {
		lang.SayString("dqmul763")
	} else {
		if !(rexsult.ToString() == "1E-6173") {
			lang.SayString("dqmul763")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-105"))
	if err != nil {
		lang.SayString("dqmul764")
	} else {
		if !(rexsult.ToString() == "1E-6174") {
			lang.SayString("dqmul764")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-106"))
	if err != nil {
		lang.SayString("dqmul765")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqmul765")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-107"))
	if err != nil {
		lang.SayString("dqmul766")
	} else {
		if !(rexsult.ToString() == "1E-6176") {
			lang.SayString("dqmul766")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-108"))
	if err != nil {
		lang.SayString("dqmul767")
	} else {
		if !(rexsult.ToString() == "1E-6177") {
			lang.SayString("dqmul767")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-109"))
	if err != nil {
		lang.SayString("dqmul768")
	} else {
		if !(rexsult.ToString() == "1E-6178") {
			lang.SayString("dqmul768")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-110"))
	if err != nil {
		lang.SayString("dqmul769")
	} else {
		if !(rexsult.ToString() == "1E-6179") {
			lang.SayString("dqmul769")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6101"))
	if err != nil {
		lang.SayString("dqmul770")
	} else {
		if !(rexsult.ToString() == "1E+6141") {
			lang.SayString("dqmul770")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6102"))
	if err != nil {
		lang.SayString("dqmul771")
	} else {
		if !(rexsult.ToString() == "1E+6142") {
			lang.SayString("dqmul771")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6103"))
	if err != nil {
		lang.SayString("dqmul772")
	} else {
		if !(rexsult.ToString() == "1E+6143") {
			lang.SayString("dqmul772")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6104"))
	if err != nil {
		lang.SayString("dqmul773")
	} else {
		if !(rexsult.ToString() == "1E+6144") {
			lang.SayString("dqmul773")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6105"))
	if err != nil {
		lang.SayString("dqmul774")
	} else {
		if !(rexsult.ToString() == "1E+6145") {
			lang.SayString("dqmul774")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6106"))
	if err != nil {
		lang.SayString("dqmul775")
	} else {
		if !(rexsult.ToString() == "1E+6146") {
			lang.SayString("dqmul775")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6107"))
	if err != nil {
		lang.SayString("dqmul776")
	} else {
		if !(rexsult.ToString() == "1E+6147") {
			lang.SayString("dqmul776")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6108"))
	if err != nil {
		lang.SayString("dqmul777")
	} else {
		if !(rexsult.ToString() == "1E+6148") {
			lang.SayString("dqmul777")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6109"))
	if err != nil {
		lang.SayString("dqmul778")
	} else {
		if !(rexsult.ToString() == "1E+6149") {
			lang.SayString("dqmul778")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e+6110"))
	if err != nil {
		lang.SayString("dqmul779")
	} else {
		if !(rexsult.ToString() == "1E+6150") {
			lang.SayString("dqmul779")
		}
	}
	rexsult, err = lang.RxFromString("1.0000E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul801")
	} else {
		if !(rexsult.ToString() == "1.0000E-6172") {
			lang.SayString("dqmul801")
		}
	}
	rexsult, err = lang.RxFromString("1.000E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-1"))
	if err != nil {
		lang.SayString("dqmul802")
	} else {
		if !(rexsult.ToString() == "1.000E-6173") {
			lang.SayString("dqmul802")
		}
	}
	rexsult, err = lang.RxFromString("1.00E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-2"))
	if err != nil {
		lang.SayString("dqmul803")
	} else {
		if !(rexsult.ToString() == "1.00E-6174") {
			lang.SayString("dqmul803")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-3"))
	if err != nil {
		lang.SayString("dqmul804")
	} else {
		if !(rexsult.ToString() == "1.0E-6175") {
			lang.SayString("dqmul804")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul805")
	} else {
		if !(rexsult.ToString() == "1.0E-6176") {
			lang.SayString("dqmul805")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul806")
	} else {
		if !(rexsult.ToString() == "1.3E-6176") {
			lang.SayString("dqmul806")
		}
	}
	rexsult, err = lang.RxFromString("1.5E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul807")
	} else {
		if !(rexsult.ToString() == "1.5E-6176") {
			lang.SayString("dqmul807")
		}
	}
	rexsult, err = lang.RxFromString("1.7E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul808")
	} else {
		if !(rexsult.ToString() == "1.7E-6176") {
			lang.SayString("dqmul808")
		}
	}
	rexsult, err = lang.RxFromString("2.3E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul809")
	} else {
		if !(rexsult.ToString() == "2.3E-6176") {
			lang.SayString("dqmul809")
		}
	}
	rexsult, err = lang.RxFromString("2.5E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul810")
	} else {
		if !(rexsult.ToString() == "2.5E-6176") {
			lang.SayString("dqmul810")
		}
	}
	rexsult, err = lang.RxFromString("2.7E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul811")
	} else {
		if !(rexsult.ToString() == "2.7E-6176") {
			lang.SayString("dqmul811")
		}
	}
	rexsult, err = lang.RxFromString("1.49E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul812")
	} else {
		if !(rexsult.ToString() == "1.49E-6176") {
			lang.SayString("dqmul812")
		}
	}
	rexsult, err = lang.RxFromString("1.50E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul813")
	} else {
		if !(rexsult.ToString() == "1.50E-6176") {
			lang.SayString("dqmul813")
		}
	}
	rexsult, err = lang.RxFromString("1.51E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul814")
	} else {
		if !(rexsult.ToString() == "1.51E-6176") {
			lang.SayString("dqmul814")
		}
	}
	rexsult, err = lang.RxFromString("2.49E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul815")
	} else {
		if !(rexsult.ToString() == "2.49E-6176") {
			lang.SayString("dqmul815")
		}
	}
	rexsult, err = lang.RxFromString("2.50E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul816")
	} else {
		if !(rexsult.ToString() == "2.50E-6176") {
			lang.SayString("dqmul816")
		}
	}
	rexsult, err = lang.RxFromString("2.51E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul817")
	} else {
		if !(rexsult.ToString() == "2.51E-6176") {
			lang.SayString("dqmul817")
		}
	}
	rexsult, err = lang.RxFromString("1E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("dqmul818")
	} else {
		if !(rexsult.ToString() == "1E-6176") {
			lang.SayString("dqmul818")
		}
	}
	rexsult, err = lang.RxFromString("3E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("dqmul819")
	} else {
		if !(rexsult.ToString() == "3E-6177") {
			lang.SayString("dqmul819")
		}
	}
	rexsult, err = lang.RxFromString("5E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("dqmul820")
	} else {
		if !(rexsult.ToString() == "5E-6177") {
			lang.SayString("dqmul820")
		}
	}
	rexsult, err = lang.RxFromString("7E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("dqmul821")
	} else {
		if !(rexsult.ToString() == "7E-6177") {
			lang.SayString("dqmul821")
		}
	}
	rexsult, err = lang.RxFromString("9E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("dqmul822")
	} else {
		if !(rexsult.ToString() == "9E-6177") {
			lang.SayString("dqmul822")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("dqmul823")
	} else {
		if !(rexsult.ToString() == "9.9E-6177") {
			lang.SayString("dqmul823")
		}
	}
	rexsult, err = lang.RxFromString("1E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e-4"))
	if err != nil {
		lang.SayString("dqmul824")
	} else {
		if !(rexsult.ToString() == "-1E-6176") {
			lang.SayString("dqmul824")
		}
	}
	rexsult, err = lang.RxFromString("3E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("dqmul825")
	} else {
		if !(rexsult.ToString() == "-3E-6177") {
			lang.SayString("dqmul825")
		}
	}
	rexsult, err = lang.RxFromString("-5E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("dqmul826")
	} else {
		if !(rexsult.ToString() == "-5E-6177") {
			lang.SayString("dqmul826")
		}
	}
	rexsult, err = lang.RxFromString("7E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("dqmul827")
	} else {
		if !(rexsult.ToString() == "-7E-6177") {
			lang.SayString("dqmul827")
		}
	}
	rexsult, err = lang.RxFromString("-9E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("dqmul828")
	} else {
		if !(rexsult.ToString() == "-9E-6177") {
			lang.SayString("dqmul828")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("dqmul829")
	} else {
		if !(rexsult.ToString() == "-9.9E-6177") {
			lang.SayString("dqmul829")
		}
	}
	rexsult, err = lang.RxFromString("3.0E-6172").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("dqmul830")
	} else {
		if !(rexsult.ToString() == "-3.0E-6177") {
			lang.SayString("dqmul830")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-5977").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-200"))
	if err != nil {
		lang.SayString("dqmul831")
	} else {
		if !(rexsult.ToString() == "1.0E-6177") {
			lang.SayString("dqmul831")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-5977").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-199"))
	if err != nil {
		lang.SayString("dqmul832")
	} else {
		if !(rexsult.ToString() == "1.0E-6176") {
			lang.SayString("dqmul832")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-5977").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1e-198"))
	if err != nil {
		lang.SayString("dqmul833")
	} else {
		if !(rexsult.ToString() == "1.0E-6175") {
			lang.SayString("dqmul833")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-5977").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("2e-198"))
	if err != nil {
		lang.SayString("dqmul834")
	} else {
		if !(rexsult.ToString() == "4.0E-6175") {
			lang.SayString("dqmul834")
		}
	}
	rexsult, err = lang.RxFromString("4.0E-5977").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("4e-198"))
	if err != nil {
		lang.SayString("dqmul835")
	} else {
		if !(rexsult.ToString() == "1.60E-6174") {
			lang.SayString("dqmul835")
		}
	}
	rexsult, err = lang.RxFromString("10.0E-5977").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10e-198"))
	if err != nil {
		lang.SayString("dqmul836")
	} else {
		if !(rexsult.ToString() == "1.000E-6173") {
			lang.SayString("dqmul836")
		}
	}
	rexsult, err = lang.RxFromString("30.0E-5977").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("30e-198"))
	if err != nil {
		lang.SayString("dqmul837")
	} else {
		if !(rexsult.ToString() == "9.000E-6173") {
			lang.SayString("dqmul837")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-5982").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("40e-166"))
	if err != nil {
		lang.SayString("dqmul838")
	} else {
		if !(rexsult.ToString() == "1.6000E-6145") {
			lang.SayString("dqmul838")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-5982").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("40e-165"))
	if err != nil {
		lang.SayString("dqmul839")
	} else {
		if !(rexsult.ToString() == "1.6000E-6144") {
			lang.SayString("dqmul839")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-5982").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("40e-164"))
	if err != nil {
		lang.SayString("dqmul840")
	} else {
		if !(rexsult.ToString() == "1.6000E-6143") {
			lang.SayString("dqmul840")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9.999E+6143"))
	if err != nil {
		lang.SayString("dqmul870")
	} else {
		if !(rexsult.ToString() == "9.99900E+6145") {
			lang.SayString("dqmul870")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("-9.999E+6143"))
	if err != nil {
		lang.SayString("dqmul871")
	} else {
		if !(rexsult.ToString() == "-9.99900E+6145") {
			lang.SayString("dqmul871")
		}
	}
	rexsult, err = lang.RxFromString("9.999E+6143").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqmul872")
	} else {
		if !(rexsult.ToString() == "9.99900E+6145") {
			lang.SayString("dqmul872")
		}
	}
	rexsult, err = lang.RxFromString("-9.999E+6143").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqmul873")
	} else {
		if !(rexsult.ToString() == "-9.99900E+6145") {
			lang.SayString("dqmul873")
		}
	}
	rexsult, err = lang.RxFromString("1.2347E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.2347E-40"))
	if err != nil {
		lang.SayString("dqmul881")
	} else {
		if !(rexsult.ToString() == "1.52448409E-6173") {
			lang.SayString("dqmul881")
		}
	}
	rexsult, err = lang.RxFromString("1.234E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("dqmul882")
	} else {
		if !(rexsult.ToString() == "1.522756E-6173") {
			lang.SayString("dqmul882")
		}
	}
	rexsult, err = lang.RxFromString("1.23E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.23E-40"))
	if err != nil {
		lang.SayString("dqmul883")
	} else {
		if !(rexsult.ToString() == "1.5129E-6173") {
			lang.SayString("dqmul883")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.2E-40"))
	if err != nil {
		lang.SayString("dqmul884")
	} else {
		if !(rexsult.ToString() == "1.44E-6173") {
			lang.SayString("dqmul884")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.2E-41"))
	if err != nil {
		lang.SayString("dqmul885")
	} else {
		if !(rexsult.ToString() == "1.44E-6174") {
			lang.SayString("dqmul885")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.2E-42"))
	if err != nil {
		lang.SayString("dqmul886")
	} else {
		if !(rexsult.ToString() == "1.44E-6175") {
			lang.SayString("dqmul886")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.3E-42"))
	if err != nil {
		lang.SayString("dqmul887")
	} else {
		if !(rexsult.ToString() == "1.56E-6175") {
			lang.SayString("dqmul887")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.3E-42"))
	if err != nil {
		lang.SayString("dqmul888")
	} else {
		if !(rexsult.ToString() == "1.69E-6175") {
			lang.SayString("dqmul888")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-6133").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.3E-43"))
	if err != nil {
		lang.SayString("dqmul889")
	} else {
		if !(rexsult.ToString() == "1.69E-6176") {
			lang.SayString("dqmul889")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-6134").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.3E-43"))
	if err != nil {
		lang.SayString("dqmul890")
	} else {
		if !(rexsult.ToString() == "1.69E-6177") {
			lang.SayString("dqmul890")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-39").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.234E-6133"))
	if err != nil {
		lang.SayString("dqmul891")
	} else {
		if !(rexsult.ToString() == "1.5233730E-6172") {
			lang.SayString("dqmul891")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-39").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.234E-6133"))
	if err != nil {
		lang.SayString("dqmul892")
	} else {
		if !(rexsult.ToString() == "1.52344704E-6172") {
			lang.SayString("dqmul892")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.234E-6133"))
	if err != nil {
		lang.SayString("dqmul893")
	} else {
		if !(rexsult.ToString() == "1.5233730E-6173") {
			lang.SayString("dqmul893")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-40").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.234E-6133"))
	if err != nil {
		lang.SayString("dqmul894")
	} else {
		if !(rexsult.ToString() == "1.52344704E-6173") {
			lang.SayString("dqmul894")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-41").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.234E-6133"))
	if err != nil {
		lang.SayString("dqmul895")
	} else {
		if !(rexsult.ToString() == "1.5233730E-6174") {
			lang.SayString("dqmul895")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-41").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1.234E-6133"))
	if err != nil {
		lang.SayString("dqmul896")
	} else {
		if !(rexsult.ToString() == "1.52344704E-6174") {
			lang.SayString("dqmul896")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999999999999999999999E-6143").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul906")
	} else {
		if !(rexsult.ToString() == "9.999999999999999999999999999999999E-6143") {
			lang.SayString("dqmul906")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.09999999999999999999999999999999999"))
	if err != nil {
		lang.SayString("dqmul907")
	} else {
		if !(rexsult.ToString() == "0.09999999999999999999999999999999999") {
			lang.SayString("dqmul907")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999999999999999999999E-6143").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.09999999999999999999999999999999999"))
	if err != nil {
		lang.SayString("dqmul908")
	} else {
		if !(rexsult.ToString() == "9.999999999999999999999999999999998E-6144") {
			lang.SayString("dqmul908")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999999999999999999999").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("9999999999999999999999999999999999"))
	if err != nil {
		lang.SayString("dqmul909")
	} else {
		if !(rexsult.ToString() == "9.999999999999999999999999999999998E+67") {
			lang.SayString("dqmul909")
		}
	}
	rexsult, err = lang.RxFromString("8.81125000000001349436E-1548").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("8.000000000000000000E-1550"))
	if err != nil {
		lang.SayString("dqmul910")
	} else {
		if !(rexsult.ToString() == "7.049000000000010795488000000000000E-3097") {
			lang.SayString("dqmul910")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("120E-2"))
	if err != nil {
		lang.SayString("dqmul911")
	} else {
		if !(rexsult.ToString() == "1.5600") {
			lang.SayString("dqmul911")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("12E-1"))
	if err != nil {
		lang.SayString("dqmul912")
	} else {
		if !(rexsult.ToString() == "1.560") {
			lang.SayString("dqmul912")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqmul1001")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqmul1001")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqmul1002")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("dqmul1002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqmul1003")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("dqmul1003")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dqmul1004")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("dqmul1004")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dqmul1005")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("dqmul1005")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dqmul1006")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("dqmul1006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dqmul1007")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("dqmul1007")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dqmul1008")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("dqmul1008")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("dqmul1009")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("dqmul1009")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("dqmul1010")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("dqmul1010")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("dqmul1011")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("dqmul1011")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000"))
	if err != nil {
		lang.SayString("dqmul1012")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("dqmul1012")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000"))
	if err != nil {
		lang.SayString("dqmul1013")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("dqmul1013")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000"))
	if err != nil {
		lang.SayString("dqmul1014")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("dqmul1014")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000"))
	if err != nil {
		lang.SayString("dqmul1015")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("dqmul1015")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1016")
	} else {
		if !(rexsult.ToString() == "1000000000000000000") {
			lang.SayString("dqmul1016")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1017")
	} else {
		if !(rexsult.ToString() == "100000000000000000000000000") {
			lang.SayString("dqmul1017")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1018")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000") {
			lang.SayString("dqmul1018")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1019")
	} else {
		if !(rexsult.ToString() == "10000000000000000000000000000") {
			lang.SayString("dqmul1019")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1020")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000000000") {
			lang.SayString("dqmul1020")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul1021")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqmul1021")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqmul1022")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("dqmul1022")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqmul1023")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("dqmul1023")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqmul1024")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("dqmul1024")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dqmul1025")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("dqmul1025")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dqmul1026")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("dqmul1026")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dqmul1027")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("dqmul1027")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dqmul1028")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("dqmul1028")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dqmul1029")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("dqmul1029")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("dqmul1030")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("dqmul1030")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("dqmul1031")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("dqmul1031")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("dqmul1032")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("dqmul1032")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000"))
	if err != nil {
		lang.SayString("dqmul1033")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("dqmul1033")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000"))
	if err != nil {
		lang.SayString("dqmul1034")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("dqmul1034")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000"))
	if err != nil {
		lang.SayString("dqmul1035")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("dqmul1035")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000000"))
	if err != nil {
		lang.SayString("dqmul1036")
	} else {
		if !(rexsult.ToString() == "1000000000000000000") {
			lang.SayString("dqmul1036")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1037")
	} else {
		if !(rexsult.ToString() == "100000000000000000000000000") {
			lang.SayString("dqmul1037")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1038")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000") {
			lang.SayString("dqmul1038")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1039")
	} else {
		if !(rexsult.ToString() == "10000000000000000000000000000") {
			lang.SayString("dqmul1039")
		}
	}
	rexsult, err = lang.RxFromString("10").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1040")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000000000") {
			lang.SayString("dqmul1040")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqmul1041")
	} else {
		if !(rexsult.ToString() == "10.0") {
			lang.SayString("dqmul1041")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul1042")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("dqmul1042")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqmul1043")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("dqmul1043")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqmul1044")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("dqmul1044")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqmul1045")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("dqmul1045")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dqmul1046")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("dqmul1046")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dqmul1047")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("dqmul1047")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dqmul1048")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("dqmul1048")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dqmul1049")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("dqmul1049")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dqmul1050")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("dqmul1050")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("dqmul1051")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("dqmul1051")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("dqmul1052")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("dqmul1052")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("dqmul1053")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("dqmul1053")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000"))
	if err != nil {
		lang.SayString("dqmul1054")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("dqmul1054")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000"))
	if err != nil {
		lang.SayString("dqmul1055")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("dqmul1055")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000000"))
	if err != nil {
		lang.SayString("dqmul1056")
	} else {
		if !(rexsult.ToString() == "1000000000000000000") {
			lang.SayString("dqmul1056")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1057")
	} else {
		if !(rexsult.ToString() == "100000000000000000000000000") {
			lang.SayString("dqmul1057")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1058")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000") {
			lang.SayString("dqmul1058")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1059")
	} else {
		if !(rexsult.ToString() == "10000000000000000000000000000") {
			lang.SayString("dqmul1059")
		}
	}
	rexsult, err = lang.RxFromString("100").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1060")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000000000") {
			lang.SayString("dqmul1060")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.01"))
	if err != nil {
		lang.SayString("dqmul1061")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("dqmul1061")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqmul1062")
	} else {
		if !(rexsult.ToString() == "100.0") {
			lang.SayString("dqmul1062")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul1063")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("dqmul1063")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqmul1064")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("dqmul1064")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqmul1065")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("dqmul1065")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqmul1066")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("dqmul1066")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dqmul1067")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("dqmul1067")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dqmul1068")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("dqmul1068")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dqmul1069")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("dqmul1069")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dqmul1070")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("dqmul1070")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dqmul1071")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("dqmul1071")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("dqmul1072")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("dqmul1072")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("dqmul1073")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("dqmul1073")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("dqmul1074")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("dqmul1074")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000"))
	if err != nil {
		lang.SayString("dqmul1075")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("dqmul1075")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000"))
	if err != nil {
		lang.SayString("dqmul1076")
	} else {
		if !(rexsult.ToString() == "1000000000000000000") {
			lang.SayString("dqmul1076")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1077")
	} else {
		if !(rexsult.ToString() == "100000000000000000000000000") {
			lang.SayString("dqmul1077")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1078")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000") {
			lang.SayString("dqmul1078")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1079")
	} else {
		if !(rexsult.ToString() == "10000000000000000000000000000") {
			lang.SayString("dqmul1079")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1080")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000000000") {
			lang.SayString("dqmul1080")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.001"))
	if err != nil {
		lang.SayString("dqmul1081")
	} else {
		if !(rexsult.ToString() == "10.000") {
			lang.SayString("dqmul1081")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.01"))
	if err != nil {
		lang.SayString("dqmul1082")
	} else {
		if !(rexsult.ToString() == "100.00") {
			lang.SayString("dqmul1082")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqmul1083")
	} else {
		if !(rexsult.ToString() == "1000.0") {
			lang.SayString("dqmul1083")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqmul1084")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("dqmul1084")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqmul1085")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("dqmul1085")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqmul1086")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("dqmul1086")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqmul1087")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("dqmul1087")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("dqmul1088")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("dqmul1088")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000"))
	if err != nil {
		lang.SayString("dqmul1089")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("dqmul1089")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("dqmul1090")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("dqmul1090")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000"))
	if err != nil {
		lang.SayString("dqmul1091")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("dqmul1091")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("dqmul1092")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("dqmul1092")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000"))
	if err != nil {
		lang.SayString("dqmul1093")
	} else {
		if !(rexsult.ToString() == "10000000000000") {
			lang.SayString("dqmul1093")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000"))
	if err != nil {
		lang.SayString("dqmul1094")
	} else {
		if !(rexsult.ToString() == "100000000000000") {
			lang.SayString("dqmul1094")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000"))
	if err != nil {
		lang.SayString("dqmul1095")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("dqmul1095")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000"))
	if err != nil {
		lang.SayString("dqmul1096")
	} else {
		if !(rexsult.ToString() == "1000000000000000000") {
			lang.SayString("dqmul1096")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("10000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1097")
	} else {
		if !(rexsult.ToString() == "100000000000000000000000000") {
			lang.SayString("dqmul1097")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1098")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000") {
			lang.SayString("dqmul1098")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1099")
	} else {
		if !(rexsult.ToString() == "10000000000000000000000000000") {
			lang.SayString("dqmul1099")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("100000000000000000000000000000"))
	if err != nil {
		lang.SayString("dqmul1100")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000000000") {
			lang.SayString("dqmul1100")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("99999999999"))
	if err != nil {
		lang.SayString("dqmul1107")
	} else {
		if !(rexsult.ToString() == "999999999990000") {
			lang.SayString("dqmul1107")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("99999999999"))
	if err != nil {
		lang.SayString("dqmul1108")
	} else {
		if !(rexsult.ToString() == "999999999990000") {
			lang.SayString("dqmul1108")
		}
	}
	rexsult, err = lang.RxFromString("0.0003333333333333333333333333333333333").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqmul1110")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333333333333333333333") {
			lang.SayString("dqmul1110")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("inx101")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("inx101")
		}
	}
	rexsult, err = lang.RxFromString("9000").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("9000"))
	if err != nil {
		lang.SayString("inx102")
	} else {
		if !(rexsult.ToString() == "81000000") {
			lang.SayString("inx102")
		}
	}
	rexsult, err = lang.RxFromString("9999").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("9999"))
	if err != nil {
		lang.SayString("inx103")
	} else {
		if !(rexsult.ToString() == "99980001") {
			lang.SayString("inx103")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("inx104")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("inx104")
		}
	}
	rexsult, err = lang.RxFromString("10000").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("inx105")
	} else {
		if !(rexsult.ToString() == "1.0000000E+8") {
			lang.SayString("inx105")
		}
	}
	rexsult, err = lang.RxFromString("10001").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("10000"))
	if err != nil {
		lang.SayString("inx106")
	} else {
		if !(rexsult.ToString() == "1.0001000E+8") {
			lang.SayString("inx106")
		}
	}
	rexsult, err = lang.RxFromString("10001").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("10001"))
	if err != nil {
		lang.SayString("inx107")
	} else {
		if !(rexsult.ToString() == "1.0002000E+8") {
			lang.SayString("inx107")
		}
	}
	rexsult, err = lang.RxFromString("10101").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("10001"))
	if err != nil {
		lang.SayString("inx108")
	} else {
		if !(rexsult.ToString() == "1.0102010E+8") {
			lang.SayString("inx108")
		}
	}
	rexsult, err = lang.RxFromString("10001").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("10101"))
	if err != nil {
		lang.SayString("inx109")
	} else {
		if !(rexsult.ToString() == "1.0102010E+8") {
			lang.SayString("inx109")
		}
	}
	rexsult, err = lang.RxFromString("2").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("mulx000")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("mulx000")
		}
	}
	rexsult, err = lang.RxFromString("2").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("mulx001")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("mulx001")
		}
	}
	rexsult, err = lang.RxFromString("5").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx002")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("mulx002")
		}
	}
	rexsult, err = lang.RxFromString("5").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("mulx003")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("mulx003")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("mulx004")
	} else {
		if !(rexsult.ToString() == "2.40") {
			lang.SayString("mulx004")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("mulx005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx005")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("mulx006")
	} else {
		if !(rexsult.ToString() == "-2.40") {
			lang.SayString("mulx006")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("mulx007")
	} else {
		if !(rexsult.ToString() == "-2.40") {
			lang.SayString("mulx007")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("mulx008")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx008")
		}
	}
	rexsult, err = lang.RxFromString("-1.20").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("mulx009")
	} else {
		if !(rexsult.ToString() == "2.40") {
			lang.SayString("mulx009")
		}
	}
	rexsult, err = lang.RxFromString("5.09").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7.1"))
	if err != nil {
		lang.SayString("mulx010")
	} else {
		if !(rexsult.ToString() == "36.139") {
			lang.SayString("mulx010")
		}
	}
	rexsult, err = lang.RxFromString("2.5").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("mulx011")
	} else {
		if !(rexsult.ToString() == "10.0") {
			lang.SayString("mulx011")
		}
	}
	rexsult, err = lang.RxFromString("2.50").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("mulx012")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("mulx012")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.00000000"))
	if err != nil {
		lang.SayString("mulx013")
	} else {
		if !(rexsult.ToString() == "1.23456789") {
			lang.SayString("mulx013")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9.999999999"))
	if err != nil {
		lang.SayString("mulx014")
	} else {
		if !(rexsult.ToString() == "100.000000") {
			lang.SayString("mulx014")
		}
	}
	rexsult, err = lang.RxFromString("2.50").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("mulx015")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("mulx015")
		}
	}
	rexsult, err = lang.RxFromString("2.50").OpMult(lang.RxSetWithDigit(6), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("mulx016")
	} else {
		if !(rexsult.ToString() == "10.00") {
			lang.SayString("mulx016")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999").OpMult(lang.RxSetWithDigit(6), lang.RxFromString("9.999999999"))
	if err != nil {
		lang.SayString("mulx017")
	} else {
		if !(rexsult.ToString() == "100.000") {
			lang.SayString("mulx017")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999").OpMult(lang.RxSetWithDigit(6), lang.RxFromString("-9.999999999"))
	if err != nil {
		lang.SayString("mulx018")
	} else {
		if !(rexsult.ToString() == "-100.000") {
			lang.SayString("mulx018")
		}
	}
	rexsult, err = lang.RxFromString("-9.999999999").OpMult(lang.RxSetWithDigit(6), lang.RxFromString("9.999999999"))
	if err != nil {
		lang.SayString("mulx019")
	} else {
		if !(rexsult.ToString() == "-100.000") {
			lang.SayString("mulx019")
		}
	}
	rexsult, err = lang.RxFromString("-9.999999999").OpMult(lang.RxSetWithDigit(6), lang.RxFromString("-9.999999999"))
	if err != nil {
		lang.SayString("mulx020")
	} else {
		if !(rexsult.ToString() == "100.000") {
			lang.SayString("mulx020")
		}
	}
	rexsult, err = lang.RxFromString("999999999999").OpMult(lang.RxSetWithDigit(15), lang.RxFromString("9765625"))
	if err != nil {
		lang.SayString("mulx059")
	} else {
		if !(rexsult.ToString() == "9.76562499999023E+18") {
			lang.SayString("mulx059")
		}
	}
	rexsult, err = lang.RxFromString("999999999999").OpMult(lang.RxSetWithDigit(30), lang.RxFromString("9765625"))
	if err != nil {
		lang.SayString("mulx160")
	} else {
		if !(rexsult.ToString() == "9765624999990234375") {
			lang.SayString("mulx160")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("mulx021")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx021")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("mulx022")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx022")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("mulx023")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx023")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("mulx024")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx024")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("mulx025")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx025")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("mulx026")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx026")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("mulx027")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx027")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("mulx028")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx028")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("mulx030")
	} else {
		if !(rexsult.ToString() == "0.00500") {
			lang.SayString("mulx030")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.000"))
	if err != nil {
		lang.SayString("mulx031")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx031")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0E-3"))
	if err != nil {
		lang.SayString("mulx032")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx032")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("00.00"))
	if err != nil {
		lang.SayString("mulx033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx033")
		}
	}
	rexsult, err = lang.RxFromString("-5.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("mulx034")
	} else {
		if !(rexsult.ToString() == "-0.00500") {
			lang.SayString("mulx034")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.000"))
	if err != nil {
		lang.SayString("mulx035")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx035")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0E-3"))
	if err != nil {
		lang.SayString("mulx036")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx036")
		}
	}
	rexsult, err = lang.RxFromString("-0E-3").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("00.00"))
	if err != nil {
		lang.SayString("mulx037")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx037")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1E-3"))
	if err != nil {
		lang.SayString("mulx038")
	} else {
		if !(rexsult.ToString() == "-0.00500") {
			lang.SayString("mulx038")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.000"))
	if err != nil {
		lang.SayString("mulx039")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx039")
		}
	}
	rexsult, err = lang.RxFromString("00.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0E-3"))
	if err != nil {
		lang.SayString("mulx040")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx040")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-00.00"))
	if err != nil {
		lang.SayString("mulx041")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx041")
		}
	}
	rexsult, err = lang.RxFromString("-5.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1E-3"))
	if err != nil {
		lang.SayString("mulx042")
	} else {
		if !(rexsult.ToString() == "0.00500") {
			lang.SayString("mulx042")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.000"))
	if err != nil {
		lang.SayString("mulx043")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx043")
		}
	}
	rexsult, err = lang.RxFromString("-00.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0E-3"))
	if err != nil {
		lang.SayString("mulx044")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx044")
		}
	}
	rexsult, err = lang.RxFromString("-0E-3").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-00.00"))
	if err != nil {
		lang.SayString("mulx045")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx045")
		}
	}
	rexsult, err = lang.RxFromString("1.20").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("mulx050")
	} else {
		if !(rexsult.ToString() == "3.60") {
			lang.SayString("mulx050")
		}
	}
	rexsult, err = lang.RxFromString("7").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("mulx051")
	} else {
		if !(rexsult.ToString() == "21") {
			lang.SayString("mulx051")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.8"))
	if err != nil {
		lang.SayString("mulx052")
	} else {
		if !(rexsult.ToString() == "0.72") {
			lang.SayString("mulx052")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("mulx053")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx053")
		}
	}
	rexsult, err = lang.RxFromString("654321").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("654321"))
	if err != nil {
		lang.SayString("mulx054")
	} else {
		if !(rexsult.ToString() == "4.28135971E+11") {
			lang.SayString("mulx054")
		}
	}
	rexsult, err = lang.RxFromString("123.45").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1e+9"))
	if err != nil {
		lang.SayString("mulx062")
	} else {
		if !(rexsult.ToString() == "1.2345E+11") {
			lang.SayString("mulx062")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("mulx080")
	} else {
		if !(rexsult.ToString() == "12345678.9") {
			lang.SayString("mulx080")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("mulx081")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("mulx081")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("12345678912"))
	if err != nil {
		lang.SayString("mulx082")
	} else {
		if !(rexsult.ToString() == "1.23456789E+9") {
			lang.SayString("mulx082")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("12345678912345"))
	if err != nil {
		lang.SayString("mulx083")
	} else {
		if !(rexsult.ToString() == "1.23456789E+12") {
			lang.SayString("mulx083")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("mulx084")
	} else {
		if !(rexsult.ToString() == "12345678.9") {
			lang.SayString("mulx084")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("12345678912"))
	if err != nil {
		lang.SayString("mulx085")
	} else {
		if !(rexsult.ToString() == "1.2345679E+9") {
			lang.SayString("mulx085")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("12345678912345"))
	if err != nil {
		lang.SayString("mulx086")
	} else {
		if !(rexsult.ToString() == "1.2345679E+12") {
			lang.SayString("mulx086")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(7), lang.RxFromString("12345678912"))
	if err != nil {
		lang.SayString("mulx087")
	} else {
		if !(rexsult.ToString() == "1.234568E+9") {
			lang.SayString("mulx087")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(7), lang.RxFromString("12345678912345"))
	if err != nil {
		lang.SayString("mulx088")
	} else {
		if !(rexsult.ToString() == "1.234568E+12") {
			lang.SayString("mulx088")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx090")
	} else {
		if !(rexsult.ToString() == "12345678.9") {
			lang.SayString("mulx090")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx091")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("mulx091")
		}
	}
	rexsult, err = lang.RxFromString("12345678912").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx092")
	} else {
		if !(rexsult.ToString() == "1.23456789E+9") {
			lang.SayString("mulx092")
		}
	}
	rexsult, err = lang.RxFromString("12345678912345").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx093")
	} else {
		if !(rexsult.ToString() == "1.23456789E+12") {
			lang.SayString("mulx093")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx094")
	} else {
		if !(rexsult.ToString() == "12345678.9") {
			lang.SayString("mulx094")
		}
	}
	rexsult, err = lang.RxFromString("12345678912").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx095")
	} else {
		if !(rexsult.ToString() == "1.2345679E+9") {
			lang.SayString("mulx095")
		}
	}
	rexsult, err = lang.RxFromString("12345678912345").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx096")
	} else {
		if !(rexsult.ToString() == "1.2345679E+12") {
			lang.SayString("mulx096")
		}
	}
	rexsult, err = lang.RxFromString("12345678912").OpMult(lang.RxSetWithDigit(7), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx097")
	} else {
		if !(rexsult.ToString() == "1.234568E+9") {
			lang.SayString("mulx097")
		}
	}
	rexsult, err = lang.RxFromString("12345678912345").OpMult(lang.RxSetWithDigit(7), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx098")
	} else {
		if !(rexsult.ToString() == "1.234568E+12") {
			lang.SayString("mulx098")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("mulx101")
	} else {
		if !(rexsult.ToString() == "81") {
			lang.SayString("mulx101")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("90"))
	if err != nil {
		lang.SayString("mulx102")
	} else {
		if !(rexsult.ToString() == "810") {
			lang.SayString("mulx102")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("900"))
	if err != nil {
		lang.SayString("mulx103")
	} else {
		if !(rexsult.ToString() == "8100") {
			lang.SayString("mulx103")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("9000"))
	if err != nil {
		lang.SayString("mulx104")
	} else {
		if !(rexsult.ToString() == "81000") {
			lang.SayString("mulx104")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("90000"))
	if err != nil {
		lang.SayString("mulx105")
	} else {
		if !(rexsult.ToString() == "810000") {
			lang.SayString("mulx105")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("900000"))
	if err != nil {
		lang.SayString("mulx106")
	} else {
		if !(rexsult.ToString() == "8100000") {
			lang.SayString("mulx106")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("9000000"))
	if err != nil {
		lang.SayString("mulx107")
	} else {
		if !(rexsult.ToString() == "81000000") {
			lang.SayString("mulx107")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("90000000"))
	if err != nil {
		lang.SayString("mulx108")
	} else {
		if !(rexsult.ToString() == "810000000") {
			lang.SayString("mulx108")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("900000000"))
	if err != nil {
		lang.SayString("mulx109")
	} else {
		if !(rexsult.ToString() == "8100000000") {
			lang.SayString("mulx109")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("9000000000"))
	if err != nil {
		lang.SayString("mulx110")
	} else {
		if !(rexsult.ToString() == "81000000000") {
			lang.SayString("mulx110")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("90000000000"))
	if err != nil {
		lang.SayString("mulx111")
	} else {
		if !(rexsult.ToString() == "810000000000") {
			lang.SayString("mulx111")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("900000000000"))
	if err != nil {
		lang.SayString("mulx112")
	} else {
		if !(rexsult.ToString() == "8100000000000") {
			lang.SayString("mulx112")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("9000000000000"))
	if err != nil {
		lang.SayString("mulx113")
	} else {
		if !(rexsult.ToString() == "81000000000000") {
			lang.SayString("mulx113")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("90000000000000"))
	if err != nil {
		lang.SayString("mulx114")
	} else {
		if !(rexsult.ToString() == "810000000000000") {
			lang.SayString("mulx114")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("900000000000000"))
	if err != nil {
		lang.SayString("mulx115")
	} else {
		if !(rexsult.ToString() == "8100000000000000") {
			lang.SayString("mulx115")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("mulx116")
	} else {
		if !(rexsult.ToString() == "81000000000000000") {
			lang.SayString("mulx116")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("90000000000000000"))
	if err != nil {
		lang.SayString("mulx117")
	} else {
		if !(rexsult.ToString() == "810000000000000000") {
			lang.SayString("mulx117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("900000000000000000"))
	if err != nil {
		lang.SayString("mulx118")
	} else {
		if !(rexsult.ToString() == "8100000000000000000") {
			lang.SayString("mulx118")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("9000000000000000000"))
	if err != nil {
		lang.SayString("mulx119")
	} else {
		if !(rexsult.ToString() == "81000000000000000000") {
			lang.SayString("mulx119")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("90000000000000000000"))
	if err != nil {
		lang.SayString("mulx120")
	} else {
		if !(rexsult.ToString() == "810000000000000000000") {
			lang.SayString("mulx120")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("900000000000000000000"))
	if err != nil {
		lang.SayString("mulx121")
	} else {
		if !(rexsult.ToString() == "8100000000000000000000") {
			lang.SayString("mulx121")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("9000000000000000000000"))
	if err != nil {
		lang.SayString("mulx122")
	} else {
		if !(rexsult.ToString() == "81000000000000000000000") {
			lang.SayString("mulx122")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("90000000000000000000000"))
	if err != nil {
		lang.SayString("mulx123")
	} else {
		if !(rexsult.ToString() == "810000000000000000000000") {
			lang.SayString("mulx123")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("mulx131")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("mulx131")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("30"))
	if err != nil {
		lang.SayString("mulx132")
	} else {
		if !(rexsult.ToString() == "90") {
			lang.SayString("mulx132")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("300"))
	if err != nil {
		lang.SayString("mulx133")
	} else {
		if !(rexsult.ToString() == "900") {
			lang.SayString("mulx133")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3000"))
	if err != nil {
		lang.SayString("mulx134")
	} else {
		if !(rexsult.ToString() == "9000") {
			lang.SayString("mulx134")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("30000"))
	if err != nil {
		lang.SayString("mulx135")
	} else {
		if !(rexsult.ToString() == "90000") {
			lang.SayString("mulx135")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("300000"))
	if err != nil {
		lang.SayString("mulx136")
	} else {
		if !(rexsult.ToString() == "900000") {
			lang.SayString("mulx136")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3000000"))
	if err != nil {
		lang.SayString("mulx137")
	} else {
		if !(rexsult.ToString() == "9000000") {
			lang.SayString("mulx137")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("30000000"))
	if err != nil {
		lang.SayString("mulx138")
	} else {
		if !(rexsult.ToString() == "90000000") {
			lang.SayString("mulx138")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("300000000"))
	if err != nil {
		lang.SayString("mulx139")
	} else {
		if !(rexsult.ToString() == "900000000") {
			lang.SayString("mulx139")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3000000000"))
	if err != nil {
		lang.SayString("mulx140")
	} else {
		if !(rexsult.ToString() == "9000000000") {
			lang.SayString("mulx140")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("30000000000"))
	if err != nil {
		lang.SayString("mulx141")
	} else {
		if !(rexsult.ToString() == "90000000000") {
			lang.SayString("mulx141")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("300000000000"))
	if err != nil {
		lang.SayString("mulx142")
	} else {
		if !(rexsult.ToString() == "900000000000") {
			lang.SayString("mulx142")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3000000000000"))
	if err != nil {
		lang.SayString("mulx143")
	} else {
		if !(rexsult.ToString() == "9000000000000") {
			lang.SayString("mulx143")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("30000000000000"))
	if err != nil {
		lang.SayString("mulx144")
	} else {
		if !(rexsult.ToString() == "90000000000000") {
			lang.SayString("mulx144")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("300000000000000"))
	if err != nil {
		lang.SayString("mulx145")
	} else {
		if !(rexsult.ToString() == "900000000000000") {
			lang.SayString("mulx145")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3000000000000000"))
	if err != nil {
		lang.SayString("mulx146")
	} else {
		if !(rexsult.ToString() == "9000000000000000") {
			lang.SayString("mulx146")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("30000000000000000"))
	if err != nil {
		lang.SayString("mulx147")
	} else {
		if !(rexsult.ToString() == "90000000000000000") {
			lang.SayString("mulx147")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("300000000000000000"))
	if err != nil {
		lang.SayString("mulx148")
	} else {
		if !(rexsult.ToString() == "900000000000000000") {
			lang.SayString("mulx148")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3000000000000000000"))
	if err != nil {
		lang.SayString("mulx149")
	} else {
		if !(rexsult.ToString() == "9000000000000000000") {
			lang.SayString("mulx149")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("30000000000000000000"))
	if err != nil {
		lang.SayString("mulx150")
	} else {
		if !(rexsult.ToString() == "90000000000000000000") {
			lang.SayString("mulx150")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("300000000000000000000"))
	if err != nil {
		lang.SayString("mulx151")
	} else {
		if !(rexsult.ToString() == "900000000000000000000") {
			lang.SayString("mulx151")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3000000000000000000000"))
	if err != nil {
		lang.SayString("mulx152")
	} else {
		if !(rexsult.ToString() == "9000000000000000000000") {
			lang.SayString("mulx152")
		}
	}
	rexsult, err = lang.RxFromString("3").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("30000000000000000000000"))
	if err != nil {
		lang.SayString("mulx153")
	} else {
		if !(rexsult.ToString() == "90000000000000000000000") {
			lang.SayString("mulx153")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9e-999999998"))
	if err != nil {
		lang.SayString("mulx180")
	} else {
		if !(rexsult.ToString() == "9E-999999999") {
			lang.SayString("mulx180")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("99e-999999998"))
	if err != nil {
		lang.SayString("mulx181")
	} else {
		if !(rexsult.ToString() == "9.9E-999999998") {
			lang.SayString("mulx181")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("999e-999999998"))
	if err != nil {
		lang.SayString("mulx182")
	} else {
		if !(rexsult.ToString() == "9.99E-999999997") {
			lang.SayString("mulx182")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9e-999999998"))
	if err != nil {
		lang.SayString("mulx183")
	} else {
		if !(rexsult.ToString() == "9E-999999999") {
			lang.SayString("mulx183")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("99e-999999998"))
	if err != nil {
		lang.SayString("mulx184")
	} else {
		if !(rexsult.ToString() == "9.9E-999999998") {
			lang.SayString("mulx184")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("999e-999999998"))
	if err != nil {
		lang.SayString("mulx185")
	} else {
		if !(rexsult.ToString() == "9.99E-999999997") {
			lang.SayString("mulx185")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("999e-999999997"))
	if err != nil {
		lang.SayString("mulx186")
	} else {
		if !(rexsult.ToString() == "9.99E-999999996") {
			lang.SayString("mulx186")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9999e-999999997"))
	if err != nil {
		lang.SayString("mulx187")
	} else {
		if !(rexsult.ToString() == "9.999E-999999995") {
			lang.SayString("mulx187")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("99999e-999999997"))
	if err != nil {
		lang.SayString("mulx188")
	} else {
		if !(rexsult.ToString() == "9.9999E-999999994") {
			lang.SayString("mulx188")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9e-999999998"))
	if err != nil {
		lang.SayString("mulx190")
	} else {
		if !(rexsult.ToString() == "9E-999999998") {
			lang.SayString("mulx190")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("99e-999999998"))
	if err != nil {
		lang.SayString("mulx191")
	} else {
		if !(rexsult.ToString() == "9.9E-999999997") {
			lang.SayString("mulx191")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("999e-999999998"))
	if err != nil {
		lang.SayString("mulx192")
	} else {
		if !(rexsult.ToString() == "9.99E-999999996") {
			lang.SayString("mulx192")
		}
	}
	rexsult, err = lang.RxFromString("9e-999999998").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx193")
	} else {
		if !(rexsult.ToString() == "9E-999999998") {
			lang.SayString("mulx193")
		}
	}
	rexsult, err = lang.RxFromString("99e-999999998").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx194")
	} else {
		if !(rexsult.ToString() == "9.9E-999999997") {
			lang.SayString("mulx194")
		}
	}
	rexsult, err = lang.RxFromString("999e-999999998").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx195")
	} else {
		if !(rexsult.ToString() == "9.99E-999999996") {
			lang.SayString("mulx195")
		}
	}
	rexsult, err = lang.RxFromString("1e-599999999").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1e-400000000"))
	if err != nil {
		lang.SayString("mulx196")
	} else {
		if !(rexsult.ToString() == "1E-999999999") {
			lang.SayString("mulx196")
		}
	}
	rexsult, err = lang.RxFromString("1e-600000000").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1e-399999999"))
	if err != nil {
		lang.SayString("mulx197")
	} else {
		if !(rexsult.ToString() == "1E-999999999") {
			lang.SayString("mulx197")
		}
	}
	rexsult, err = lang.RxFromString("1.2e-599999999").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.2e-400000000"))
	if err != nil {
		lang.SayString("mulx198")
	} else {
		if !(rexsult.ToString() == "1.44E-999999999") {
			lang.SayString("mulx198")
		}
	}
	rexsult, err = lang.RxFromString("1.2e-600000000").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.2e-399999999"))
	if err != nil {
		lang.SayString("mulx199")
	} else {
		if !(rexsult.ToString() == "1.44E-999999999") {
			lang.SayString("mulx199")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx246")
	} else {
		if !(rexsult.ToString() == "145433.290801193369671916511992830") {
			lang.SayString("mulx246")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx247")
	} else {
		if !(rexsult.ToString() == "145433.29080119336967191651199283") {
			lang.SayString("mulx247")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx248")
	} else {
		if !(rexsult.ToString() == "145433.2908011933696719165119928") {
			lang.SayString("mulx248")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(30), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx249")
	} else {
		if !(rexsult.ToString() == "145433.290801193369671916511993") {
			lang.SayString("mulx249")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(29), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx250")
	} else {
		if !(rexsult.ToString() == "145433.29080119336967191651199") {
			lang.SayString("mulx250")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(28), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx251")
	} else {
		if !(rexsult.ToString() == "145433.2908011933696719165120") {
			lang.SayString("mulx251")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(27), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx252")
	} else {
		if !(rexsult.ToString() == "145433.290801193369671916512") {
			lang.SayString("mulx252")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(26), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx253")
	} else {
		if !(rexsult.ToString() == "145433.29080119336967191651") {
			lang.SayString("mulx253")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(25), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx254")
	} else {
		if !(rexsult.ToString() == "145433.2908011933696719165") {
			lang.SayString("mulx254")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(24), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx255")
	} else {
		if !(rexsult.ToString() == "145433.290801193369671916") {
			lang.SayString("mulx255")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(23), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx256")
	} else {
		if !(rexsult.ToString() == "145433.29080119336967192") {
			lang.SayString("mulx256")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(22), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx257")
	} else {
		if !(rexsult.ToString() == "145433.2908011933696719") {
			lang.SayString("mulx257")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(21), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx258")
	} else {
		if !(rexsult.ToString() == "145433.290801193369672") {
			lang.SayString("mulx258")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(20), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx259")
	} else {
		if !(rexsult.ToString() == "145433.29080119336967") {
			lang.SayString("mulx259")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(19), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx260")
	} else {
		if !(rexsult.ToString() == "145433.2908011933697") {
			lang.SayString("mulx260")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(18), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx261")
	} else {
		if !(rexsult.ToString() == "145433.290801193370") {
			lang.SayString("mulx261")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(17), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx262")
	} else {
		if !(rexsult.ToString() == "145433.29080119337") {
			lang.SayString("mulx262")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx263")
	} else {
		if !(rexsult.ToString() == "145433.2908011934") {
			lang.SayString("mulx263")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(15), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx264")
	} else {
		if !(rexsult.ToString() == "145433.290801193") {
			lang.SayString("mulx264")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(14), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx265")
	} else {
		if !(rexsult.ToString() == "145433.29080119") {
			lang.SayString("mulx265")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(13), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx266")
	} else {
		if !(rexsult.ToString() == "145433.2908012") {
			lang.SayString("mulx266")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(12), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx267")
	} else {
		if !(rexsult.ToString() == "145433.290801") {
			lang.SayString("mulx267")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(11), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx268")
	} else {
		if !(rexsult.ToString() == "145433.29080") {
			lang.SayString("mulx268")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(10), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx269")
	} else {
		if !(rexsult.ToString() == "145433.2908") {
			lang.SayString("mulx269")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx270")
	} else {
		if !(rexsult.ToString() == "145433.291") {
			lang.SayString("mulx270")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(8), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx271")
	} else {
		if !(rexsult.ToString() == "145433.29") {
			lang.SayString("mulx271")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(7), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx272")
	} else {
		if !(rexsult.ToString() == "145433.3") {
			lang.SayString("mulx272")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(6), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx273")
	} else {
		if !(rexsult.ToString() == "145433") {
			lang.SayString("mulx273")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx274")
	} else {
		if !(rexsult.ToString() == "1.4543E+5") {
			lang.SayString("mulx274")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(4), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx275")
	} else {
		if !(rexsult.ToString() == "1.454E+5") {
			lang.SayString("mulx275")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(3), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx276")
	} else {
		if !(rexsult.ToString() == "1.45E+5") {
			lang.SayString("mulx276")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(2), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx277")
	} else {
		if !(rexsult.ToString() == "1.4E+5") {
			lang.SayString("mulx277")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(1), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx278")
	} else {
		if !(rexsult.ToString() == "1E+5") {
			lang.SayString("mulx278")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("mulx301")
	} else {
		if !(rexsult.ToString() == "81") {
			lang.SayString("mulx301")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("90"))
	if err != nil {
		lang.SayString("mulx302")
	} else {
		if !(rexsult.ToString() == "810") {
			lang.SayString("mulx302")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("900"))
	if err != nil {
		lang.SayString("mulx303")
	} else {
		if !(rexsult.ToString() == "8100") {
			lang.SayString("mulx303")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9000"))
	if err != nil {
		lang.SayString("mulx304")
	} else {
		if !(rexsult.ToString() == "81000") {
			lang.SayString("mulx304")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("90000"))
	if err != nil {
		lang.SayString("mulx305")
	} else {
		if !(rexsult.ToString() == "810000") {
			lang.SayString("mulx305")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("900000"))
	if err != nil {
		lang.SayString("mulx306")
	} else {
		if !(rexsult.ToString() == "8100000") {
			lang.SayString("mulx306")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9000000"))
	if err != nil {
		lang.SayString("mulx307")
	} else {
		if !(rexsult.ToString() == "81000000") {
			lang.SayString("mulx307")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("90000000"))
	if err != nil {
		lang.SayString("mulx308")
	} else {
		if !(rexsult.ToString() == "810000000") {
			lang.SayString("mulx308")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("900000000"))
	if err != nil {
		lang.SayString("mulx309")
	} else {
		if !(rexsult.ToString() == "8.10000000E+9") {
			lang.SayString("mulx309")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9000000000"))
	if err != nil {
		lang.SayString("mulx310")
	} else {
		if !(rexsult.ToString() == "8.10000000E+10") {
			lang.SayString("mulx310")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("90000000000"))
	if err != nil {
		lang.SayString("mulx311")
	} else {
		if !(rexsult.ToString() == "8.10000000E+11") {
			lang.SayString("mulx311")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("900000000000"))
	if err != nil {
		lang.SayString("mulx312")
	} else {
		if !(rexsult.ToString() == "8.10000000E+12") {
			lang.SayString("mulx312")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9000000000000"))
	if err != nil {
		lang.SayString("mulx313")
	} else {
		if !(rexsult.ToString() == "8.10000000E+13") {
			lang.SayString("mulx313")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("90000000000000"))
	if err != nil {
		lang.SayString("mulx314")
	} else {
		if !(rexsult.ToString() == "8.10000000E+14") {
			lang.SayString("mulx314")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("900000000000000"))
	if err != nil {
		lang.SayString("mulx315")
	} else {
		if !(rexsult.ToString() == "8.10000000E+15") {
			lang.SayString("mulx315")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9000000000000000"))
	if err != nil {
		lang.SayString("mulx316")
	} else {
		if !(rexsult.ToString() == "8.10000000E+16") {
			lang.SayString("mulx316")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("90000000000000000"))
	if err != nil {
		lang.SayString("mulx317")
	} else {
		if !(rexsult.ToString() == "8.10000000E+17") {
			lang.SayString("mulx317")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("900000000000000000"))
	if err != nil {
		lang.SayString("mulx318")
	} else {
		if !(rexsult.ToString() == "8.10000000E+18") {
			lang.SayString("mulx318")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9000000000000000000"))
	if err != nil {
		lang.SayString("mulx319")
	} else {
		if !(rexsult.ToString() == "8.10000000E+19") {
			lang.SayString("mulx319")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("90000000000000000000"))
	if err != nil {
		lang.SayString("mulx320")
	} else {
		if !(rexsult.ToString() == "8.10000000E+20") {
			lang.SayString("mulx320")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("900000000000000000000"))
	if err != nil {
		lang.SayString("mulx321")
	} else {
		if !(rexsult.ToString() == "8.10000000E+21") {
			lang.SayString("mulx321")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9000000000000000000000"))
	if err != nil {
		lang.SayString("mulx322")
	} else {
		if !(rexsult.ToString() == "8.10000000E+22") {
			lang.SayString("mulx322")
		}
	}
	rexsult, err = lang.RxFromString("9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("90000000000000000000000"))
	if err != nil {
		lang.SayString("mulx323")
	} else {
		if !(rexsult.ToString() == "8.10000000E+23") {
			lang.SayString("mulx323")
		}
	}
	rexsult, err = lang.RxFromString("1.491824697641270317824852952837224").OpMult(lang.RxSetWithDigit(29), lang.RxFromString("1.105170918075647624811707826490246514675628614562883537345747603"))
	if err != nil {
		lang.SayString("mulx330")
	} else {
		if !(rexsult.ToString() == "1.6487212707001281468486507878") {
			lang.SayString("mulx330")
		}
	}
	rexsult, err = lang.RxFromString("0.8958341352965282506768545828765117803873717284891040428").OpMult(lang.RxSetWithDigit(55), lang.RxFromString("0.8958341352965282506768545828765117803873717284891040428"))
	if err != nil {
		lang.SayString("mulx331")
	} else {
		if !(rexsult.ToString() == "0.8025187979624784829842553829934069955890983696752228299") {
			lang.SayString("mulx331")
		}
	}
	rexsult, err = lang.RxFromString("0E-60").OpMult(lang.RxSetWithDigit(7), lang.RxFromString("1000E-60"))
	if err != nil {
		lang.SayString("mulx504")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx504")
		}
	}
	rexsult, err = lang.RxFromString("100E+60").OpMult(lang.RxSetWithDigit(7), lang.RxFromString("0E+60"))
	if err != nil {
		lang.SayString("mulx505")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx505")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("mulx541")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx541")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("mulx542")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx542")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx543")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx543")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx544")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx544")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("mulx545")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx545")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("mulx546")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx546")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("mulx547")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx547")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("mulx548")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx548")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("mulx551")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx551")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("mulx552")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx552")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx553")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx553")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx554")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx554")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("mulx555")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx555")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("mulx556")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx556")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("mulx557")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx557")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("mulx558")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx558")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("mulx561")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx561")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("mulx562")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx562")
		}
	}
	rexsult, err = lang.RxFromString("0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("mulx563")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx563")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("mulx564")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx564")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("mulx565")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx565")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("mulx566")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx566")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("mulx567")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx567")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("mulx568")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx568")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("mulx571")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx571")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("mulx572")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx572")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("mulx573")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx573")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("mulx574")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx574")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("mulx575")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx575")
		}
	}
	rexsult, err = lang.RxFromString("-1.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("mulx576")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx576")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("mulx577")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx577")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-0.0"))
	if err != nil {
		lang.SayString("mulx578")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("mulx578")
		}
	}
	rexsult, err = lang.RxFromString("1e-599999999").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1e-400000000"))
	if err != nil {
		lang.SayString("mulx740")
	} else {
		if !(rexsult.ToString() == "1E-999999999") {
			lang.SayString("mulx740")
		}
	}
	rexsult, err = lang.RxFromString("9e-999999998").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("mulx743")
	} else {
		if !(rexsult.ToString() == "9E-999999999") {
			lang.SayString("mulx743")
		}
	}
	rexsult, err = lang.RxFromString("1.0000E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx801")
	} else {
		if !(rexsult.ToString() == "1.0000E-999") {
			lang.SayString("mulx801")
		}
	}
	rexsult, err = lang.RxFromString("1.000E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-1"))
	if err != nil {
		lang.SayString("mulx802")
	} else {
		if !(rexsult.ToString() == "1.000E-1000") {
			lang.SayString("mulx802")
		}
	}
	rexsult, err = lang.RxFromString("1.00E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-2"))
	if err != nil {
		lang.SayString("mulx803")
	} else {
		if !(rexsult.ToString() == "1.00E-1001") {
			lang.SayString("mulx803")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-3"))
	if err != nil {
		lang.SayString("mulx804")
	} else {
		if !(rexsult.ToString() == "1.0E-1002") {
			lang.SayString("mulx804")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx805")
	} else {
		if !(rexsult.ToString() == "1.0E-1003") {
			lang.SayString("mulx805")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx806")
	} else {
		if !(rexsult.ToString() == "1.3E-1003") {
			lang.SayString("mulx806")
		}
	}
	rexsult, err = lang.RxFromString("1.5E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx807")
	} else {
		if !(rexsult.ToString() == "1.5E-1003") {
			lang.SayString("mulx807")
		}
	}
	rexsult, err = lang.RxFromString("1.7E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx808")
	} else {
		if !(rexsult.ToString() == "1.7E-1003") {
			lang.SayString("mulx808")
		}
	}
	rexsult, err = lang.RxFromString("2.3E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx809")
	} else {
		if !(rexsult.ToString() == "2.3E-1003") {
			lang.SayString("mulx809")
		}
	}
	rexsult, err = lang.RxFromString("2.5E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx810")
	} else {
		if !(rexsult.ToString() == "2.5E-1003") {
			lang.SayString("mulx810")
		}
	}
	rexsult, err = lang.RxFromString("2.7E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx811")
	} else {
		if !(rexsult.ToString() == "2.7E-1003") {
			lang.SayString("mulx811")
		}
	}
	rexsult, err = lang.RxFromString("1.49E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx812")
	} else {
		if !(rexsult.ToString() == "1.49E-1003") {
			lang.SayString("mulx812")
		}
	}
	rexsult, err = lang.RxFromString("1.50E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx813")
	} else {
		if !(rexsult.ToString() == "1.50E-1003") {
			lang.SayString("mulx813")
		}
	}
	rexsult, err = lang.RxFromString("1.51E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx814")
	} else {
		if !(rexsult.ToString() == "1.51E-1003") {
			lang.SayString("mulx814")
		}
	}
	rexsult, err = lang.RxFromString("2.49E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx815")
	} else {
		if !(rexsult.ToString() == "2.49E-1003") {
			lang.SayString("mulx815")
		}
	}
	rexsult, err = lang.RxFromString("2.50E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx816")
	} else {
		if !(rexsult.ToString() == "2.50E-1003") {
			lang.SayString("mulx816")
		}
	}
	rexsult, err = lang.RxFromString("2.51E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx817")
	} else {
		if !(rexsult.ToString() == "2.51E-1003") {
			lang.SayString("mulx817")
		}
	}
	rexsult, err = lang.RxFromString("1E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-4"))
	if err != nil {
		lang.SayString("mulx818")
	} else {
		if !(rexsult.ToString() == "1E-1003") {
			lang.SayString("mulx818")
		}
	}
	rexsult, err = lang.RxFromString("3E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("mulx819")
	} else {
		if !(rexsult.ToString() == "3E-1004") {
			lang.SayString("mulx819")
		}
	}
	rexsult, err = lang.RxFromString("5E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("mulx820")
	} else {
		if !(rexsult.ToString() == "5E-1004") {
			lang.SayString("mulx820")
		}
	}
	rexsult, err = lang.RxFromString("7E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("mulx821")
	} else {
		if !(rexsult.ToString() == "7E-1004") {
			lang.SayString("mulx821")
		}
	}
	rexsult, err = lang.RxFromString("9E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("mulx822")
	} else {
		if !(rexsult.ToString() == "9E-1004") {
			lang.SayString("mulx822")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("mulx823")
	} else {
		if !(rexsult.ToString() == "9.9E-1004") {
			lang.SayString("mulx823")
		}
	}
	rexsult, err = lang.RxFromString("1E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("-1e-4"))
	if err != nil {
		lang.SayString("mulx824")
	} else {
		if !(rexsult.ToString() == "-1E-1003") {
			lang.SayString("mulx824")
		}
	}
	rexsult, err = lang.RxFromString("3E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("mulx825")
	} else {
		if !(rexsult.ToString() == "-3E-1004") {
			lang.SayString("mulx825")
		}
	}
	rexsult, err = lang.RxFromString("-5E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("mulx826")
	} else {
		if !(rexsult.ToString() == "-5E-1004") {
			lang.SayString("mulx826")
		}
	}
	rexsult, err = lang.RxFromString("7E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("mulx827")
	} else {
		if !(rexsult.ToString() == "-7E-1004") {
			lang.SayString("mulx827")
		}
	}
	rexsult, err = lang.RxFromString("-9E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-5"))
	if err != nil {
		lang.SayString("mulx828")
	} else {
		if !(rexsult.ToString() == "-9E-1004") {
			lang.SayString("mulx828")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("mulx829")
	} else {
		if !(rexsult.ToString() == "-9.9E-1004") {
			lang.SayString("mulx829")
		}
	}
	rexsult, err = lang.RxFromString("3.0E-999").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("-1e-5"))
	if err != nil {
		lang.SayString("mulx830")
	} else {
		if !(rexsult.ToString() == "-3.0E-1004") {
			lang.SayString("mulx830")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-501"))
	if err != nil {
		lang.SayString("mulx831")
	} else {
		if !(rexsult.ToString() == "1.0E-1002") {
			lang.SayString("mulx831")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("2e-501"))
	if err != nil {
		lang.SayString("mulx832")
	} else {
		if !(rexsult.ToString() == "4.0E-1002") {
			lang.SayString("mulx832")
		}
	}
	rexsult, err = lang.RxFromString("4.0E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4e-501"))
	if err != nil {
		lang.SayString("mulx833")
	} else {
		if !(rexsult.ToString() == "1.60E-1001") {
			lang.SayString("mulx833")
		}
	}
	rexsult, err = lang.RxFromString("10.0E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("10e-501"))
	if err != nil {
		lang.SayString("mulx834")
	} else {
		if !(rexsult.ToString() == "1.000E-1000") {
			lang.SayString("mulx834")
		}
	}
	rexsult, err = lang.RxFromString("30.0E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("30e-501"))
	if err != nil {
		lang.SayString("mulx835")
	} else {
		if !(rexsult.ToString() == "9.000E-1000") {
			lang.SayString("mulx835")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("40e-501"))
	if err != nil {
		lang.SayString("mulx836")
	} else {
		if !(rexsult.ToString() == "1.6000E-999") {
			lang.SayString("mulx836")
		}
	}
	rexsult, err = lang.RxFromString("1E-502").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-502"))
	if err != nil {
		lang.SayString("mulx840")
	} else {
		if !(rexsult.ToString() == "1E-1004") {
			lang.SayString("mulx840")
		}
	}
	rexsult, err = lang.RxFromString("1E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-501"))
	if err != nil {
		lang.SayString("mulx841")
	} else {
		if !(rexsult.ToString() == "1E-1002") {
			lang.SayString("mulx841")
		}
	}
	rexsult, err = lang.RxFromString("2E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("2e-501"))
	if err != nil {
		lang.SayString("mulx842")
	} else {
		if !(rexsult.ToString() == "4E-1002") {
			lang.SayString("mulx842")
		}
	}
	rexsult, err = lang.RxFromString("4E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4e-501"))
	if err != nil {
		lang.SayString("mulx843")
	} else {
		if !(rexsult.ToString() == "1.6E-1001") {
			lang.SayString("mulx843")
		}
	}
	rexsult, err = lang.RxFromString("10E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("10e-501"))
	if err != nil {
		lang.SayString("mulx844")
	} else {
		if !(rexsult.ToString() == "1.00E-1000") {
			lang.SayString("mulx844")
		}
	}
	rexsult, err = lang.RxFromString("30E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("30e-501"))
	if err != nil {
		lang.SayString("mulx845")
	} else {
		if !(rexsult.ToString() == "9.00E-1000") {
			lang.SayString("mulx845")
		}
	}
	rexsult, err = lang.RxFromString("40E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("40e-501"))
	if err != nil {
		lang.SayString("mulx846")
	} else {
		if !(rexsult.ToString() == "1.600E-999") {
			lang.SayString("mulx846")
		}
	}
	rexsult, err = lang.RxFromString("1E-670").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-335"))
	if err != nil {
		lang.SayString("mulx850")
	} else {
		if !(rexsult.ToString() == "1E-1005") {
			lang.SayString("mulx850")
		}
	}
	rexsult, err = lang.RxFromString("1E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-334"))
	if err != nil {
		lang.SayString("mulx851")
	} else {
		if !(rexsult.ToString() == "1E-1002") {
			lang.SayString("mulx851")
		}
	}
	rexsult, err = lang.RxFromString("4E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("2e-334"))
	if err != nil {
		lang.SayString("mulx852")
	} else {
		if !(rexsult.ToString() == "8E-1002") {
			lang.SayString("mulx852")
		}
	}
	rexsult, err = lang.RxFromString("9E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("3e-334"))
	if err != nil {
		lang.SayString("mulx853")
	} else {
		if !(rexsult.ToString() == "2.7E-1001") {
			lang.SayString("mulx853")
		}
	}
	rexsult, err = lang.RxFromString("16E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4e-334"))
	if err != nil {
		lang.SayString("mulx854")
	} else {
		if !(rexsult.ToString() == "6.4E-1001") {
			lang.SayString("mulx854")
		}
	}
	rexsult, err = lang.RxFromString("25E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5e-334"))
	if err != nil {
		lang.SayString("mulx855")
	} else {
		if !(rexsult.ToString() == "1.25E-1000") {
			lang.SayString("mulx855")
		}
	}
	rexsult, err = lang.RxFromString("10E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("100e-334"))
	if err != nil {
		lang.SayString("mulx856")
	} else {
		if !(rexsult.ToString() == "1.000E-999") {
			lang.SayString("mulx856")
		}
	}
	rexsult, err = lang.RxFromString("6636851557994578716E-520").OpMult(lang.RxSetWithDigit(19), lang.RxFromString("6636851557994578716E-520"))
	if err != nil {
		lang.SayString("mulx860")
	} else {
		if !(rexsult.ToString() == "4.404779860285506685E-1003") {
			lang.SayString("mulx860")
		}
	}
	rexsult, err = lang.RxFromString("1.2347E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.2347E-40"))
	if err != nil {
		lang.SayString("mulx881")
	} else {
		if !(rexsult.ToString() == "1.5245E-80") {
			lang.SayString("mulx881")
		}
	}
	rexsult, err = lang.RxFromString("1.234E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("mulx882")
	} else {
		if !(rexsult.ToString() == "1.5228E-80") {
			lang.SayString("mulx882")
		}
	}
	rexsult, err = lang.RxFromString("1.23E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.23E-40"))
	if err != nil {
		lang.SayString("mulx883")
	} else {
		if !(rexsult.ToString() == "1.5129E-80") {
			lang.SayString("mulx883")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.2E-40"))
	if err != nil {
		lang.SayString("mulx884")
	} else {
		if !(rexsult.ToString() == "1.44E-80") {
			lang.SayString("mulx884")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.2E-41"))
	if err != nil {
		lang.SayString("mulx885")
	} else {
		if !(rexsult.ToString() == "1.44E-81") {
			lang.SayString("mulx885")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.2E-42"))
	if err != nil {
		lang.SayString("mulx886")
	} else {
		if !(rexsult.ToString() == "1.44E-82") {
			lang.SayString("mulx886")
		}
	}
	rexsult, err = lang.RxFromString("1.2E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.3E-42"))
	if err != nil {
		lang.SayString("mulx887")
	} else {
		if !(rexsult.ToString() == "1.56E-82") {
			lang.SayString("mulx887")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.3E-42"))
	if err != nil {
		lang.SayString("mulx888")
	} else {
		if !(rexsult.ToString() == "1.69E-82") {
			lang.SayString("mulx888")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.3E-43"))
	if err != nil {
		lang.SayString("mulx889")
	} else {
		if !(rexsult.ToString() == "1.69E-83") {
			lang.SayString("mulx889")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-41").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.3E-43"))
	if err != nil {
		lang.SayString("mulx890")
	} else {
		if !(rexsult.ToString() == "1.69E-84") {
			lang.SayString("mulx890")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-39").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("mulx891")
	} else {
		if !(rexsult.ToString() == "1.5234E-79") {
			lang.SayString("mulx891")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-39").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("mulx892")
	} else {
		if !(rexsult.ToString() == "1.5234E-79") {
			lang.SayString("mulx892")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("mulx893")
	} else {
		if !(rexsult.ToString() == "1.5234E-80") {
			lang.SayString("mulx893")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-40").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("mulx894")
	} else {
		if !(rexsult.ToString() == "1.5234E-80") {
			lang.SayString("mulx894")
		}
	}
	rexsult, err = lang.RxFromString("1.2345E-41").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("mulx895")
	} else {
		if !(rexsult.ToString() == "1.5234E-81") {
			lang.SayString("mulx895")
		}
	}
	rexsult, err = lang.RxFromString("1.23456E-41").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.234E-40"))
	if err != nil {
		lang.SayString("mulx896")
	} else {
		if !(rexsult.ToString() == "1.5234E-81") {
			lang.SayString("mulx896")
		}
	}
	rexsult, err = lang.RxFromString("0.3000000000E-191").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.3000000000E-191"))
	if err != nil {
		lang.SayString("mulx900")
	} else {
		if !(rexsult.ToString() == "9.000000000000000E-384") {
			lang.SayString("mulx900")
		}
	}
	rexsult, err = lang.RxFromString("0.3000000001E-191").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.3000000001E-191"))
	if err != nil {
		lang.SayString("mulx901")
	} else {
		if !(rexsult.ToString() == "9.000000006000000E-384") {
			lang.SayString("mulx901")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0999999999999"))
	if err != nil {
		lang.SayString("mulx902")
	} else {
		if !(rexsult.ToString() == "9.999999999989999E-384") {
			lang.SayString("mulx902")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.09999999999999"))
	if err != nil {
		lang.SayString("mulx903")
	} else {
		if !(rexsult.ToString() == "9.999999999998999E-384") {
			lang.SayString("mulx903")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.099999999999999"))
	if err != nil {
		lang.SayString("mulx904")
	} else {
		if !(rexsult.ToString() == "9.999999999999899E-384") {
			lang.SayString("mulx904")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0999999999999999"))
	if err != nil {
		lang.SayString("mulx905")
	} else {
		if !(rexsult.ToString() == "9.999999999999989E-384") {
			lang.SayString("mulx905")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("mulx906")
	} else {
		if !(rexsult.ToString() == "9.999999999999999E-383") {
			lang.SayString("mulx906")
		}
	}
	rexsult, err = lang.RxFromString("1").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.09999999999999999"))
	if err != nil {
		lang.SayString("mulx907")
	} else {
		if !(rexsult.ToString() == "0.09999999999999999") {
			lang.SayString("mulx907")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.09999999999999999"))
	if err != nil {
		lang.SayString("mulx908")
	} else {
		if !(rexsult.ToString() == "9.999999999999998E-384") {
			lang.SayString("mulx908")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.099999999999999999"))
	if err != nil {
		lang.SayString("mulx909")
	} else {
		if !(rexsult.ToString() == "9.999999999999999E-384") {
			lang.SayString("mulx909")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.0999999999999999999"))
	if err != nil {
		lang.SayString("mulx910")
	} else {
		if !(rexsult.ToString() == "9.999999999999999E-384") {
			lang.SayString("mulx910")
		}
	}
	rexsult, err = lang.RxFromString("9.999999999999999E-383").OpMult(lang.RxSetWithDigit(16), lang.RxFromString("0.09999999999999999999"))
	if err != nil {
		lang.SayString("mulx911")
	} else {
		if !(rexsult.ToString() == "9.999999999999999E-384") {
			lang.SayString("mulx911")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("120E-2"))
	if err != nil {
		lang.SayString("mulx1001")
	} else {
		if !(rexsult.ToString() == "1.5600") {
			lang.SayString("mulx1001")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpMult(lang.RxSetWithDigit(34), lang.RxFromString("12E-1"))
	if err != nil {
		lang.SayString("mulx1002")
	} else {
		if !(rexsult.ToString() == "1.560") {
			lang.SayString("mulx1002")
		}
	}
	rexsult, err = lang.RxFromString("1E-502").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-502"))
	if err != nil {
		lang.SayString("mulx400")
	} else {
		if !(rexsult.ToString() == "1E-1004") {
			lang.SayString("mulx400")
		}
	}
	rexsult, err = lang.RxFromString("1E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-501"))
	if err != nil {
		lang.SayString("mulx401")
	} else {
		if !(rexsult.ToString() == "1E-1002") {
			lang.SayString("mulx401")
		}
	}
	rexsult, err = lang.RxFromString("2E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("2e-501"))
	if err != nil {
		lang.SayString("mulx402")
	} else {
		if !(rexsult.ToString() == "4E-1002") {
			lang.SayString("mulx402")
		}
	}
	rexsult, err = lang.RxFromString("4E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4e-501"))
	if err != nil {
		lang.SayString("mulx403")
	} else {
		if !(rexsult.ToString() == "1.6E-1001") {
			lang.SayString("mulx403")
		}
	}
	rexsult, err = lang.RxFromString("10E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("10e-501"))
	if err != nil {
		lang.SayString("mulx404")
	} else {
		if !(rexsult.ToString() == "1.00E-1000") {
			lang.SayString("mulx404")
		}
	}
	rexsult, err = lang.RxFromString("30E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("30e-501"))
	if err != nil {
		lang.SayString("mulx405")
	} else {
		if !(rexsult.ToString() == "9.00E-1000") {
			lang.SayString("mulx405")
		}
	}
	rexsult, err = lang.RxFromString("40E-501").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("40e-501"))
	if err != nil {
		lang.SayString("mulx406")
	} else {
		if !(rexsult.ToString() == "1.600E-999") {
			lang.SayString("mulx406")
		}
	}
	rexsult, err = lang.RxFromString("1E-670").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-335"))
	if err != nil {
		lang.SayString("mulx410")
	} else {
		if !(rexsult.ToString() == "1E-1005") {
			lang.SayString("mulx410")
		}
	}
	rexsult, err = lang.RxFromString("1E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1e-334"))
	if err != nil {
		lang.SayString("mulx411")
	} else {
		if !(rexsult.ToString() == "1E-1002") {
			lang.SayString("mulx411")
		}
	}
	rexsult, err = lang.RxFromString("4E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("2e-334"))
	if err != nil {
		lang.SayString("mulx412")
	} else {
		if !(rexsult.ToString() == "8E-1002") {
			lang.SayString("mulx412")
		}
	}
	rexsult, err = lang.RxFromString("9E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("3e-334"))
	if err != nil {
		lang.SayString("mulx413")
	} else {
		if !(rexsult.ToString() == "2.7E-1001") {
			lang.SayString("mulx413")
		}
	}
	rexsult, err = lang.RxFromString("16E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4e-334"))
	if err != nil {
		lang.SayString("mulx414")
	} else {
		if !(rexsult.ToString() == "6.4E-1001") {
			lang.SayString("mulx414")
		}
	}
	rexsult, err = lang.RxFromString("25E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5e-334"))
	if err != nil {
		lang.SayString("mulx415")
	} else {
		if !(rexsult.ToString() == "1.25E-1000") {
			lang.SayString("mulx415")
		}
	}
	rexsult, err = lang.RxFromString("10E-668").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("100e-334"))
	if err != nil {
		lang.SayString("mulx416")
	} else {
		if !(rexsult.ToString() == "1.000E-999") {
			lang.SayString("mulx416")
		}
	}
	rexsult, err = lang.RxFromString("4953734675913.065314738743322579").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("0218.932010396534371704930714860E+797"))
	if err != nil {
		lang.SayString("mulx3001")
	} else {
		if !(rexsult.ToString() == "1.084531091568672041923151632066E+812") {
			lang.SayString("mulx3001")
		}
	}
	rexsult, err = lang.RxFromString("9641.684323386955881595490347910E-844").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-78864532047.12287484430980636798E+934"))
	if err != nil {
		lang.SayString("mulx3002")
	} else {
		if !(rexsult.ToString() == "-7.603869223099928141659831589905E+104") {
			lang.SayString("mulx3002")
		}
	}
	rexsult, err = lang.RxFromString("-1.028048571628326871054964307774E+529").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("49200008645699.35577937582714739"))
	if err != nil {
		lang.SayString("mulx3003")
	} else {
		if !(rexsult.ToString() == "-5.057999861231255549283737861207E+542") {
			lang.SayString("mulx3003")
		}
	}
	rexsult, err = lang.RxFromString("479084.8561808930525417735205519").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("084157571054.2691784660983989931"))
	if err != nil {
		lang.SayString("mulx3004")
	} else {
		if !(rexsult.ToString() == "40318617825067837.47317700523687") {
			lang.SayString("mulx3004")
		}
	}
	rexsult, err = lang.RxFromString("-0363750788.573782205664349562931").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-3172.080934464133691909905980096"))
	if err != nil {
		lang.SayString("mulx3005")
	} else {
		if !(rexsult.ToString() == "1153846941331.188583292239230818") {
			lang.SayString("mulx3005")
		}
	}
	rexsult, err = lang.RxFromString("1381026551423669919010191878449").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-82.66614775445371254999357800739"))
	if err != nil {
		lang.SayString("mulx3006")
	} else {
		if !(rexsult.ToString() == "-1.141641449528127656560770057228E+32") {
			lang.SayString("mulx3006")
		}
	}
	rexsult, err = lang.RxFromString("4627.026960423072127953556635585").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-4410583132901.830017479741231131"))
	if err != nil {
		lang.SayString("mulx3007")
	} else {
		if !(rexsult.ToString() == "-20407887067124025.31576887565113") {
			lang.SayString("mulx3007")
		}
	}
	rexsult, err = lang.RxFromString("75353574493.84484153484918212042").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-8684111695095849922263044191221"))
	if err != nil {
		lang.SayString("mulx3008")
	} else {
		if !(rexsult.ToString() == "-6.543788575292743281456830701127E+41") {
			lang.SayString("mulx3008")
		}
	}
	rexsult, err = lang.RxFromString("6907058.216435355874729592373011").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("2.857005446917670515662398741545"))
	if err != nil {
		lang.SayString("mulx3009")
	} else {
		if !(rexsult.ToString() == "19733502.94653326211623698034717") {
			lang.SayString("mulx3009")
		}
	}
	rexsult, err = lang.RxFromString("-38949530427253.24030680468677190").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("712168021.1265384466442576619064E-992"))
	if err != nil {
		lang.SayString("mulx3010")
	} else {
		if !(rexsult.ToString() == "-2.773861000818483769292240109417E-970") {
			lang.SayString("mulx3010")
		}
	}
	rexsult, err = lang.RxFromString("-0708069.025667471996378081482549").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-562842.4701520787831018732202804"))
	if err != nil {
		lang.SayString("mulx3011")
	} else {
		if !(rexsult.ToString() == "398531319444.8556128729086112205") {
			lang.SayString("mulx3011")
		}
	}
	rexsult, err = lang.RxFromString("4055087.246994644709729942673976").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-43183146921897.67383476104084575E+211"))
	if err != nil {
		lang.SayString("mulx3012")
	} else {
		if !(rexsult.ToString() == "-1.751114283680833039197637874453E+231") {
			lang.SayString("mulx3012")
		}
	}
	rexsult, err = lang.RxFromString("4502895892520.396581348110906909E-512").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-815.9047305921862348263521876034"))
	if err != nil {
		lang.SayString("mulx3013")
	} else {
		if !(rexsult.ToString() == "-3.673934060071516156604453756541E-497") {
			lang.SayString("mulx3013")
		}
	}
	rexsult, err = lang.RxFromString("467.6721295072628100260239179865").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-02.07155073395573569852316073025"))
	if err != nil {
		lang.SayString("mulx3014")
	} else {
		if !(rexsult.ToString() == "-968.8065431314121523074875069807") {
			lang.SayString("mulx3014")
		}
	}
	rexsult, err = lang.RxFromString("2.156795313311150143949997552501E-571").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-8677147.586389401682712180146855"))
	if err != nil {
		lang.SayString("mulx3015")
	} else {
		if !(rexsult.ToString() == "-1.871483124723381986272837942577E-564") {
			lang.SayString("mulx3015")
		}
	}
	rexsult, err = lang.RxFromString("-974953.2801637208368002585822457").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-693095793.3667578067802698191246"))
	if err != nil {
		lang.SayString("mulx3016")
	} else {
		if !(rexsult.ToString() == "675736017210596.9899587749991363") {
			lang.SayString("mulx3016")
		}
	}
	rexsult, err = lang.RxFromString("-7634680140009571846155654339781").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("3009630949502.035852433434214413E-490"))
	if err != nil {
		lang.SayString("mulx3017")
	} else {
		if !(rexsult.ToString() == "-2.297756963892134373657544025107E-447") {
			lang.SayString("mulx3017")
		}
	}
	rexsult, err = lang.RxFromString("262273.0222851186523650889896428E-624").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("74177.21073338090843145838835480"))
	if err != nil {
		lang.SayString("mulx3018")
	} else {
		if !(rexsult.ToString() == "1.945468124372395349192665031675E-614") {
			lang.SayString("mulx3018")
		}
	}
	rexsult, err = lang.RxFromString("-8036052748815903177624716581732").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-066677357.4438809548850966167573"))
	if err != nil {
		lang.SayString("mulx3019")
	} else {
		if !(rexsult.ToString() == "5.358227615706800711033262124598E+38") {
			lang.SayString("mulx3019")
		}
	}
	rexsult, err = lang.RxFromString("883429.5928031498103637713570166E+765").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-43978.97283712939198111043032726"))
	if err != nil {
		lang.SayString("mulx3020")
	} else {
		if !(rexsult.ToString() == "-3.885232606540600490321438191516E+775") {
			lang.SayString("mulx3020")
		}
	}
	rexsult, err = lang.RxFromString("24791301060.37938360567775506973").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-5613327866480.322649080205877564"))
	if err != nil {
		lang.SayString("mulx3021")
	} else {
		if !(rexsult.ToString() == "-139161701088530765925120.8408852") {
			lang.SayString("mulx3021")
		}
	}
	rexsult, err = lang.RxFromString("-930711443.9474781586162910776139").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-740.3860979292775472622798348030"))
	if err != nil {
		lang.SayString("mulx3022")
	} else {
		if !(rexsult.ToString() == "689085814282.3968746911100154133") {
			lang.SayString("mulx3022")
		}
	}
	rexsult, err = lang.RxFromString("2358276428765.064191082773385539").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("214.3589796082328665878602304469"))
	if err != nil {
		lang.SayString("mulx3023")
	} else {
		if !(rexsult.ToString() == "505517728904226.6233443209659001") {
			lang.SayString("mulx3023")
		}
	}
	rexsult, err = lang.RxFromString("-3.868744449795653651638308926987E+750").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("8270.472492965559872384018329418"))
	if err != nil {
		lang.SayString("mulx3024")
	} else {
		if !(rexsult.ToString() == "-3.199634455434813294426505526063E+754") {
			lang.SayString("mulx3024")
		}
	}
	rexsult, err = lang.RxFromString("140422069.5863246490180206814374E-447").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-567195652586.2454217069003186487"))
	if err != nil {
		lang.SayString("mulx3025")
	} else {
		if !(rexsult.ToString() == "-7.964678739652657498503799559950E-428") {
			lang.SayString("mulx3025")
		}
	}
	rexsult, err = lang.RxFromString("75929096475.63450425339472559646E+153").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-0945260193.503803519572604151290E+459"))
	if err != nil {
		lang.SayString("mulx3026")
	} else {
		if !(rexsult.ToString() == "-7.177275242712723733041569606882E+631") {
			lang.SayString("mulx3026")
		}
	}
	rexsult, err = lang.RxFromString("6312318309.142044953357460463732").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-5641317823.202274083982487558514E+628"))
	if err != nil {
		lang.SayString("mulx3027")
	} else {
		if !(rexsult.ToString() == "-3.560979378308906043783023726787E+647") {
			lang.SayString("mulx3027")
		}
	}
	rexsult, err = lang.RxFromString("93793652428100.52105928239469937").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("917.2571313109730433369594936416E-712"))
	if err != nil {
		lang.SayString("mulx3028")
	} else {
		if !(rexsult.ToString() == "8.603289656137796526769786965341E-696") {
			lang.SayString("mulx3028")
		}
	}
	rexsult, err = lang.RxFromString("98471198160.56524417578665886060").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-23994.14313393939743548945165462"))
	if err != nil {
		lang.SayString("mulx3029")
	} else {
		if !(rexsult.ToString() == "-2362732023235112.375960528304974") {
			lang.SayString("mulx3029")
		}
	}
	rexsult, err = lang.RxFromString("329326552.0208398002250836592043").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-02451.10065397010591546041034041"))
	if err != nil {
		lang.SayString("mulx3030")
	} else {
		if !(rexsult.ToString() == "-807212527028.0005401736893474430") {
			lang.SayString("mulx3030")
		}
	}
	rexsult, err = lang.RxFromString("-92980.68431371090354435763218439").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-2282178507046019721925800997065"))
	if err != nil {
		lang.SayString("mulx3031")
	} else {
		if !(rexsult.ToString() == "2.121985193111820147170707717938E+35") {
			lang.SayString("mulx3031")
		}
	}
	rexsult, err = lang.RxFromString("12135817762.27858606259822256987E+738").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("98.35649167872356132249561021910E-902"))
	if err != nil {
		lang.SayString("mulx3032")
	} else {
		if !(rexsult.ToString() == "1.193636458750059340733188876015E-152") {
			lang.SayString("mulx3032")
		}
	}
	rexsult, err = lang.RxFromString("37.27457578793521166809739140081").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-392550.4790095035979998355569916"))
	if err != nil {
		lang.SayString("mulx3033")
	} else {
		if !(rexsult.ToString() == "-14632152.58043001234518095997140") {
			lang.SayString("mulx3033")
		}
	}
	rexsult, err = lang.RxFromString("-2787.980590304199878755265273703").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("7117631179305319208210387565324"))
	if err != nil {
		lang.SayString("mulx3034")
	} else {
		if !(rexsult.ToString() == "-1.984381757684722217801410305714E+34") {
			lang.SayString("mulx3034")
		}
	}
	rexsult, err = lang.RxFromString("-9890633.854609434943559831911276E+971").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-1939985729.436827777055699361237"))
	if err != nil {
		lang.SayString("mulx3035")
	} else {
		if !(rexsult.ToString() == "1.918768853302706825964087702307E+987") {
			lang.SayString("mulx3035")
		}
	}
	rexsult, err = lang.RxFromString("3944570323.331121750661920475191").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-17360722.28878145641394962484366"))
	if err != nil {
		lang.SayString("mulx3036")
	} else {
		if !(rexsult.ToString() == "-68480589931920481.56020043213767") {
			lang.SayString("mulx3036")
		}
	}
	rexsult, err = lang.RxFromString("19544.14018503427029002552872707").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("1786697762.885178994182133839546"))
	if err != nil {
		lang.SayString("mulx3037")
	} else {
		if !(rexsult.ToString() == "34919471546115.05897163496162290") {
			lang.SayString("mulx3037")
		}
	}
	rexsult, err = lang.RxFromString("-05.75485957937617757983513662981").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("5564476875.989640431173694372083"))
	if err != nil {
		lang.SayString("mulx3038")
	} else {
		if !(rexsult.ToString() == "-32022783054.00620878436398990135") {
			lang.SayString("mulx3038")
		}
	}
	rexsult, err = lang.RxFromString("-4208820.898718069194008526302746").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("626887.7553774705678201112845462E+206"))
	if err != nil {
		lang.SayString("mulx3039")
	} else {
		if !(rexsult.ToString() == "-2.638458285983158789458925170267E+218") {
			lang.SayString("mulx3039")
		}
	}
	rexsult, err = lang.RxFromString("-70077195478066.30896979085821269E+549").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("4607.163248554155483681430013073"))
	if err != nil {
		lang.SayString("mulx3040")
	} else {
		if !(rexsult.ToString() == "-3.228570795682925509478191397878E+566") {
			lang.SayString("mulx3040")
		}
	}
	rexsult, err = lang.RxFromString("-442941.7541811527940918244383454").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-068126768.0563559819156379151016"))
	if err != nil {
		lang.SayString("mulx3041")
	} else {
		if !(rexsult.ToString() == "30176190149574.84386395947593970") {
			lang.SayString("mulx3041")
		}
	}
	rexsult, err = lang.RxFromString("-040726778711.8677615616711676159").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("299691.9430345259174614997064916"))
	if err != nil {
		lang.SayString("mulx3042")
	} else {
		if !(rexsult.ToString() == "-12205487445696816.02175665622242") {
			lang.SayString("mulx3042")
		}
	}
	rexsult, err = lang.RxFromString("-1934197520.738366912179143085955").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("3.810751422515178400293693371519"))
	if err != nil {
		lang.SayString("mulx3043")
	} else {
		if !(rexsult.ToString() == "-7370745953.579062985130438309023") {
			lang.SayString("mulx3043")
		}
	}
	rexsult, err = lang.RxFromString("813262.7723533833038829559646830").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-303284822716.8282178131118185907"))
	if err != nil {
		lang.SayString("mulx3044")
	} else {
		if !(rexsult.ToString() == "-246650255735392080.1357404280431") {
			lang.SayString("mulx3044")
		}
	}
	rexsult, err = lang.RxFromString("36105954884.94621434979365589311").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("745558205.7692397481313005659523E-952"))
	if err != nil {
		lang.SayString("mulx3045")
	} else {
		if !(rexsult.ToString() == "2.691909094160561673391352743869E-933") {
			lang.SayString("mulx3045")
		}
	}
	rexsult, err = lang.RxFromString("-075537177538.1814516621962185490").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("26980775255.51542856483122484898"))
	if err != nil {
		lang.SayString("mulx3046")
	} else {
		if !(rexsult.ToString() == "-2038051610593641947717.268652175") {
			lang.SayString("mulx3046")
		}
	}
	rexsult, err = lang.RxFromString("-4223765.415319564898840040697647").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-2590590305497454185455459149918E-215"))
	if err != nil {
		lang.SayString("mulx3047")
	} else {
		if !(rexsult.ToString() == "1.094204573762229308798604845395E-178") {
			lang.SayString("mulx3047")
		}
	}
	rexsult, err = lang.RxFromString("-6468.903738522951259063099946195").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-7877.324314273694312164407794939E+267"))
	if err != nil {
		lang.SayString("mulx3048")
	} else {
		if !(rexsult.ToString() == "5.095765270616284455922747530676E+274") {
			lang.SayString("mulx3048")
		}
	}
	rexsult, err = lang.RxFromString("-9567221.183663236817239254783372E-203").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("1650.198961256061165362319471264"))
	if err != nil {
		lang.SayString("mulx3049")
	} else {
		if !(rexsult.ToString() == "-1.578781845938805737527304303976E-193") {
			lang.SayString("mulx3049")
		}
	}
	rexsult, err = lang.RxFromString("8812306098770.200752139142033569E-428").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("26790.17380163975186972720427030E+568"))
	if err != nil {
		lang.SayString("mulx3050")
	} else {
		if !(rexsult.ToString() == "2.360832119793036398127652187732E+157") {
			lang.SayString("mulx3050")
		}
	}
	rexsult, err = lang.RxFromString("80108033.12724838718736922500904").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-706207255092.7645192310078892869"))
	if err != nil {
		lang.SayString("mulx3051")
	} else {
		if !(rexsult.ToString() == "-56572874185674332398.36004114372") {
			lang.SayString("mulx3051")
		}
	}
	rexsult, err = lang.RxFromString("-37942846282.76101663789059003505").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-5.649456053942850351313869983197"))
	if err != nil {
		lang.SayString("mulx3052")
	} else {
		if !(rexsult.ToString() == "214356442635.9672009449140933366") {
			lang.SayString("mulx3052")
		}
	}
	rexsult, err = lang.RxFromString("92659632115305.13735437728445541").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("6483438.317862851676468094261410E-139"))
	if err != nil {
		lang.SayString("mulx3053")
	} else {
		if !(rexsult.ToString() == "6.007530093754446085819255987878E-119") {
			lang.SayString("mulx3053")
		}
	}
	rexsult, err = lang.RxFromString("2838948.589837595494152150647194").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("569547026247.5469563701415715960"))
	if err != nil {
		lang.SayString("mulx3054")
	} else {
		if !(rexsult.ToString() == "1616914727011669419.390959984273") {
			lang.SayString("mulx3054")
		}
	}
	rexsult, err = lang.RxFromString("524995204523.6053307941775794287E+694").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("1589600879689517100527293028553"))
	if err != nil {
		lang.SayString("mulx3055")
	} else {
		if !(rexsult.ToString() == "8.345328389435009812933599889447E+735") {
			lang.SayString("mulx3055")
		}
	}
	rexsult, err = lang.RxFromString("-57131573677452.15449921725097290").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("4669681430736.326858508715643769"))
	if err != nil {
		lang.SayString("mulx3056")
	} else {
		if !(rexsult.ToString() == "-266786248710342647746063322.0544") {
			lang.SayString("mulx3056")
		}
	}
	rexsult, err = lang.RxFromString("90794826.55528018781830463383411").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-5.471502270351231110027647216128"))
	if err != nil {
		lang.SayString("mulx3057")
	} else {
		if !(rexsult.ToString() == "-496784099.6333617958496589124964") {
			lang.SayString("mulx3057")
		}
	}
	rexsult, err = lang.RxFromString("58508794729.35191160840980489138").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-47060867.24988279680824397447551"))
	if err != nil {
		lang.SayString("mulx3058")
	} else {
		if !(rexsult.ToString() == "-2753474621708672573.249029643967") {
			lang.SayString("mulx3058")
		}
	}
	rexsult, err = lang.RxFromString("-746104.0768078474426464219416332E+006").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("9595418.300613754556671852801667E+385"))
	if err != nil {
		lang.SayString("mulx3059")
	} else {
		if !(rexsult.ToString() == "-7.159180712764549711669939947084E+403") {
			lang.SayString("mulx3059")
		}
	}
	rexsult, err = lang.RxFromString("55.99427632688387400403789659459E+119").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-9.170530450881612853998489340127"))
	if err != nil {
		lang.SayString("mulx3060")
	} else {
		if !(rexsult.ToString() == "-5.134972161307679939281170944556E+121") {
			lang.SayString("mulx3060")
		}
	}
	rexsult, err = lang.RxFromString("-41214265628.83801241467317270595").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("1015336323798389903361978271354"))
	if err != nil {
		lang.SayString("mulx3061")
	} else {
		if !(rexsult.ToString() == "-4.184634095163472384028549378392E+40") {
			lang.SayString("mulx3061")
		}
	}
	rexsult, err = lang.RxFromString("89937.39749201095570357557430822").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("82351554210093.60879476027800331"))
	if err != nil {
		lang.SayString("mulx3062")
	} else {
		if !(rexsult.ToString() == "7406484465078077191.920015793662") {
			lang.SayString("mulx3062")
		}
	}
	rexsult, err = lang.RxFromString("01712661.64677082156284125486943E+359").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("57932.78435529483241552042115837E-037"))
	if err != nil {
		lang.SayString("mulx3063")
	} else {
		if !(rexsult.ToString() == "9.921925785595813587655312307930E+332") {
			lang.SayString("mulx3063")
		}
	}
	rexsult, err = lang.RxFromString("-2647593306.528617691373470059213").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-655531558709.4582168930191014461"))
	if err != nil {
		lang.SayString("mulx3064")
	} else {
		if !(rexsult.ToString() == "1735580967057433153120.099643641") {
			lang.SayString("mulx3064")
		}
	}
	rexsult, err = lang.RxFromString("2904078690665765116603253099668E-329").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-71.45586619176091599264717047885E+787"))
	if err != nil {
		lang.SayString("mulx3065")
	} else {
		if !(rexsult.ToString() == "-2.075134583305571527962710017262E+490") {
			lang.SayString("mulx3065")
		}
	}
	rexsult, err = lang.RxFromString("22094338972.39109726522477999515").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-409846549371.3900805039668417203E-499"))
	if err != nil {
		lang.SayString("mulx3066")
	} else {
		if !(rexsult.ToString() == "-9.055288588476315822113975426730E-478") {
			lang.SayString("mulx3066")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("mulx3067")
	} else {
		if !(rexsult.ToString() == "-2.777409563825512202793336132310E+41") {
			lang.SayString("mulx3067")
		}
	}
	rexsult, err = lang.RxFromString("-84172558160661.35863831029352323").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-11271.58916600931155937291904890"))
	if err != nil {
		lang.SayString("mulx3068")
	} else {
		if !(rexsult.ToString() == "948758494638999235.1953022970755") {
			lang.SayString("mulx3068")
		}
	}
	rexsult, err = lang.RxFromString("-70046932324614.90596396237508541E-568").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("33.63163964004608865836577297698E-918"))
	if err != nil {
		lang.SayString("mulx3069")
	} else {
		if !(rexsult.ToString() == "-2.355793185832144388285949021738E-1471") {
			lang.SayString("mulx3069")
		}
	}
	rexsult, err = lang.RxFromString("0004125384407.053782660115680886").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-391429084.5847321402514385603223E-648"))
	if err != nil {
		lang.SayString("mulx3070")
	} else {
		if !(rexsult.ToString() == "-1.614795442013190139080634449273E-630") {
			lang.SayString("mulx3070")
		}
	}
	rexsult, err = lang.RxFromString("-31823131.15691583393820628480997E-440").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("92913.91582947237200286427030028E+771"))
	if err != nil {
		lang.SayString("mulx3071")
	} else {
		if !(rexsult.ToString() == "-2.956811729743937541973845029816E+343") {
			lang.SayString("mulx3071")
		}
	}
	rexsult, err = lang.RxFromString("55573867888.91575330563698128150").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("599.5231614736232188354393212234"))
	if err != nil {
		lang.SayString("mulx3072")
	} else {
		if !(rexsult.ToString() == "33317820972080.24347717542221477") {
			lang.SayString("mulx3072")
		}
	}
	rexsult, err = lang.RxFromString("-5447727448431680878699555714796E-800").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("5487207.142687001607026665515349E-362"))
	if err != nil {
		lang.SayString("mulx3073")
	} else {
		if !(rexsult.ToString() == "-2.989280896644635352838087864373E-1125") {
			lang.SayString("mulx3073")
		}
	}
	rexsult, err = lang.RxFromString("0418349404834.547110239542290134").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("09819915.92405288066606124554841"))
	if err != nil {
		lang.SayString("mulx3074")
	} else {
		if !(rexsult.ToString() == "4108155982352814348.343441299082") {
			lang.SayString("mulx3074")
		}
	}
	rexsult, err = lang.RxFromString("-262021.7565194737396448014286436").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-7983992600094836304387324162042E+390"))
	if err != nil {
		lang.SayString("mulx3075")
	} else {
		if !(rexsult.ToString() == "2.091979765115329268275803385534E+426") {
			lang.SayString("mulx3075")
		}
	}
	rexsult, err = lang.RxFromString("48696050631.42565380301204592392E-505").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-33868752339.85057267609967806187E+821"))
	if err != nil {
		lang.SayString("mulx3076")
	} else {
		if !(rexsult.ToString() == "-1.649274478764579569246425611629E+337") {
			lang.SayString("mulx3076")
		}
	}
	rexsult, err = lang.RxFromString("95316999.19440144356471126680708").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-60791.33805057402845885978390435"))
	if err != nil {
		lang.SayString("mulx3077")
	} else {
		if !(rexsult.ToString() == "-5794447919993.150493301061195714") {
			lang.SayString("mulx3077")
		}
	}
	rexsult, err = lang.RxFromString("-5326702296402708234722215224979E-136").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("8032459.450998820205916538543258"))
	if err != nil {
		lang.SayString("mulx3078")
	} else {
		if !(rexsult.ToString() == "-4.278652020339705265013632757349E-99") {
			lang.SayString("mulx3078")
		}
	}
	rexsult, err = lang.RxFromString("67.18750684079501575335482615780E-281").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("734.1168841683438410314843011541E-854"))
	if err != nil {
		lang.SayString("mulx3079")
	} else {
		if !(rexsult.ToString() == "4.932348317700372401849231767007E-1131") {
			lang.SayString("mulx3079")
		}
	}
	rexsult, err = lang.RxFromString("-8739299372114.092482914139281669").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("507610074.7343577029345077385838"))
	if err != nil {
		lang.SayString("mulx3080")
	} else {
		if !(rexsult.ToString() == "-4436156407404759833857.580707024") {
			lang.SayString("mulx3080")
		}
	}
	rexsult, err = lang.RxFromString("2454.002078468928665008217727731").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("583546039.6233842869119950982009E-147"))
	if err != nil {
		lang.SayString("mulx3081")
	} else {
		if !(rexsult.ToString() == "1.432023194118096842806010293027E-135") {
			lang.SayString("mulx3081")
		}
	}
	rexsult, err = lang.RxFromString("764578.5204849936912066033177429").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("64603.13571259164812609436832506"))
	if err != nil {
		lang.SayString("mulx3082")
	} else {
		if !(rexsult.ToString() == "49394169921.82458094138096628957") {
			lang.SayString("mulx3082")
		}
	}
	rexsult, err = lang.RxFromString("079203.7330103777716903518367560").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("846388934347.6324036132959664705"))
	if err != nil {
		lang.SayString("mulx3083")
	} else {
		if !(rexsult.ToString() == "67037163179008037.19983564789203") {
			lang.SayString("mulx3083")
		}
	}
	rexsult, err = lang.RxFromString("-4278.581514688669249247007127899E-329").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("5.474973992953902631890208360829"))
	if err != nil {
		lang.SayString("mulx3084")
	} else {
		if !(rexsult.ToString() == "-2.342512251965378028433584538870E-325") {
			lang.SayString("mulx3084")
		}
	}
	rexsult, err = lang.RxFromString("60867019.81764798845468445196869E+651").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("6.149612565404080501157093851895E+817"))
	if err != nil {
		lang.SayString("mulx3085")
	} else {
		if !(rexsult.ToString() == "3.743085898893072544197564013497E+1476") {
			lang.SayString("mulx3085")
		}
	}
	rexsult, err = lang.RxFromString("18554417738217.62218590965803605E-382").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-0894505909529.052378474618435782E+527"))
	if err != nil {
		lang.SayString("mulx3086")
	} else {
		if !(rexsult.ToString() == "-1.659703631470633700884136887614E+170") {
			lang.SayString("mulx3086")
		}
	}
	rexsult, err = lang.RxFromString("69073355517144.36356688642213839").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("997784782535.6104634823627327033E+116"))
	if err != nil {
		lang.SayString("mulx3087")
	} else {
		if !(rexsult.ToString() == "6.892034301367879802693422066425E+141") {
			lang.SayString("mulx3087")
		}
	}
	rexsult, err = lang.RxFromString("450282259072.8657099359104277477").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-1791307965314309175477911369824"))
	if err != nil {
		lang.SayString("mulx3088")
	} else {
		if !(rexsult.ToString() == "-8.065941973169457071650996861677E+41") {
			lang.SayString("mulx3088")
		}
	}
	rexsult, err = lang.RxFromString("954678411.7838149266455177850037").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("142988.7096204254529284334278794"))
	if err != nil {
		lang.SayString("mulx3089")
	} else {
		if !(rexsult.ToString() == "136508234203444.8694879431412375") {
			lang.SayString("mulx3089")
		}
	}
	rexsult, err = lang.RxFromString("-9244530976.220812127155852389807E+557").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("541089.4715446858896619078627941"))
	if err != nil {
		lang.SayString("mulx3090")
	} else {
		if !(rexsult.ToString() == "-5.002118380601798392363043558941E+572") {
			lang.SayString("mulx3090")
		}
	}
	rexsult, err = lang.RxFromString("-75492024.20990197005974241975449").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-14760421311348.35269044633000927"))
	if err != nil {
		lang.SayString("mulx3091")
	} else {
		if !(rexsult.ToString() == "1114294082984662825831.464787487") {
			lang.SayString("mulx3091")
		}
	}
	rexsult, err = lang.RxFromString("317747.6972215715434186596178036E-452").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("24759763.33144824613591228097330E+092"))
	if err != nil {
		lang.SayString("mulx3092")
	} else {
		if !(rexsult.ToString() == "7.867357782318786860404997647513E-348") {
			lang.SayString("mulx3092")
		}
	}
	rexsult, err = lang.RxFromString("-8.153334430358647134334545353427").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-9.717872025814596548462854853522"))
	if err != nil {
		lang.SayString("mulx3093")
	} else {
		if !(rexsult.ToString() == "79.23306057789328578902960605222") {
			lang.SayString("mulx3093")
		}
	}
	rexsult, err = lang.RxFromString("7.267345197492967332320456062961E-478").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("5054015481833.263541189916208065"))
	if err != nil {
		lang.SayString("mulx3094")
	} else {
		if !(rexsult.ToString() == "3.672927513995607308048737751972E-465") {
			lang.SayString("mulx3094")
		}
	}
	rexsult, err = lang.RxFromString("-1223354029.862567054230912271171").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("8135774223401322785475014855625"))
	if err != nil {
		lang.SayString("mulx3095")
	} else {
		if !(rexsult.ToString() == "-9.952932182250005119307429060894E+39") {
			lang.SayString("mulx3095")
		}
	}
	rexsult, err = lang.RxFromString("285397644111.5655679961211349982E+645").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-2479499427613157519359627280704"))
	if err != nil {
		lang.SayString("mulx3096")
	} else {
		if !(rexsult.ToString() == "-7.076432952167704614138411740001E+686") {
			lang.SayString("mulx3096")
		}
	}
	rexsult, err = lang.RxFromString("-4673112.663442366293812346653429").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("-3429.998403142546001438238460958"))
	if err != nil {
		lang.SayString("mulx3097")
	} else {
		if !(rexsult.ToString() == "16028768973.31252639476148371361") {
			lang.SayString("mulx3097")
		}
	}
	rexsult, err = lang.RxFromString("88.96492479681278079861456051103").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("386939.4621006514751889096510923E+139"))
	if err != nil {
		lang.SayString("mulx3098")
	} else {
		if !(rexsult.ToString() == "3.442404014670364763780946297856E+146") {
			lang.SayString("mulx3098")
		}
	}
	rexsult, err = lang.RxFromString("064326846.4286437304788069444326E-942").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("92.23649942010862087149015091350"))
	if err != nil {
		lang.SayString("mulx3099")
	} else {
		if !(rexsult.ToString() == "5.933283133313013755814405436342E-933") {
			lang.SayString("mulx3099")
		}
	}
	rexsult, err = lang.RxFromString("504507.0043949324433170405699360").OpMult(lang.RxSetWithDigit(31), lang.RxFromString("605387.7175522955344659311072099"))
	if err != nil {
		lang.SayString("mulx3100")
	} else {
		if !(rexsult.ToString() == "305422343879.7940838630401656585") {
			lang.SayString("mulx3100")
		}
	}
	rexsult, err = lang.RxFromString("1.5283550163839789319142430253644").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-1.6578158484822969520405291379492"))
	if err != nil {
		lang.SayString("mulx3201")
	} else {
		if !(rexsult.ToString() == "-2.5337311682687808926633910761614") {
			lang.SayString("mulx3201")
		}
	}
	rexsult, err = lang.RxFromString("-622903030605.2867503937836507326").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("6519388607.1331855704471328795821"))
	if err != nil {
		lang.SayString("mulx3202")
	} else {
		if !(rexsult.ToString() == "-4060946921076840449949.6988828486") {
			lang.SayString("mulx3202")
		}
	}
	rexsult, err = lang.RxFromString("-5675915.2465457487632250245209054").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("73913909880.381367895205086027416"))
	if err != nil {
		lang.SayString("mulx3203")
	} else {
		if !(rexsult.ToString() == "-419529088021865067.23307352973589") {
			lang.SayString("mulx3203")
		}
	}
	rexsult, err = lang.RxFromString("97.647321172555144900685605318635").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("4.8620911587547548751209841570885"))
	if err != nil {
		lang.SayString("mulx3204")
	} else {
		if !(rexsult.ToString() == "474.77017694916635398652276042175") {
			lang.SayString("mulx3204")
		}
	}
	rexsult, err = lang.RxFromString("-9717253267024.5380651435435603552").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-2669.2539695193820424002013488480E+363"))
	if err != nil {
		lang.SayString("mulx3205")
	} else {
		if !(rexsult.ToString() == "2.5937816855830431899123217912144E+379") {
			lang.SayString("mulx3205")
		}
	}
	rexsult, err = lang.RxFromString("-4.0817391717190128506083001702335E-767").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("12772.807105920712660991033689206"))
	if err != nil {
		lang.SayString("mulx3206")
	} else {
		if !(rexsult.ToString() == "-5.2135267097047531336100750110314E-763") {
			lang.SayString("mulx3206")
		}
	}
	rexsult, err = lang.RxFromString("68625322655934146845003028928647").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-59.634169944840280159782488098700"))
	if err != nil {
		lang.SayString("mulx3207")
	} else {
		if !(rexsult.ToString() == "-4.0924141537834748501140151997778E+33") {
			lang.SayString("mulx3207")
		}
	}
	rexsult, err = lang.RxFromString("732515.76532049290815665856727641").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-92134479835821.319619827023729829"))
	if err != nil {
		lang.SayString("mulx3208")
	} else {
		if !(rexsult.ToString() == "-67489959009342175728.710494356322") {
			lang.SayString("mulx3208")
		}
	}
	rexsult, err = lang.RxFromString("-30.458011942978338421676454778733").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-5023372024597665102336430410403E+831"))
	if err != nil {
		lang.SayString("mulx3209")
	} else {
		if !(rexsult.ToString() == "1.5300192511921895929031818638961E+863") {
			lang.SayString("mulx3209")
		}
	}
	rexsult, err = lang.RxFromString("-89640.094149414644660480286430462").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-58703419758.800889903227509215474"))
	if err != nil {
		lang.SayString("mulx3210")
	} else {
		if !(rexsult.ToString() == "5262180074071519.7018252171579753") {
			lang.SayString("mulx3210")
		}
	}
	rexsult, err = lang.RxFromString("458653.1567870081810052917714259").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("18353106238.516235116080449814053E-038"))
	if err != nil {
		lang.SayString("mulx3211")
	} else {
		if !(rexsult.ToString() == "8.4177101131428047497998594379593E-23") {
			lang.SayString("mulx3211")
		}
	}
	rexsult, err = lang.RxFromString("913391.42744224458216174967853722").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-21051638.816432817393202262710630E+432"))
	if err != nil {
		lang.SayString("mulx3212")
	} else {
		if !(rexsult.ToString() == "-1.9228386428540135340600836707270E+445") {
			lang.SayString("mulx3212")
		}
	}
	rexsult, err = lang.RxFromString("-917591456829.12109027484399536567").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-28892177726858026955513438843371E+708"))
	if err != nil {
		lang.SayString("mulx3213")
	} else {
		if !(rexsult.ToString() == "2.6511215451353541156703914721725E+751") {
			lang.SayString("mulx3213")
		}
	}
	rexsult, err = lang.RxFromString("34938410840645.913399699219228218").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("30.818220393242402846077755480548"))
	if err != nil {
		lang.SayString("mulx3214")
	} else {
		if !(rexsult.ToString() == "1076739645476675.3318519289128961") {
			lang.SayString("mulx3214")
		}
	}
	rexsult, err = lang.RxFromString("6034.7374411022598081745006769023E-517").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("29771833428054709077850588904653"))
	if err != nil {
		lang.SayString("mulx3215")
	} else {
		if !(rexsult.ToString() == "1.7966519787854159464382359411642E-482") {
			lang.SayString("mulx3215")
		}
	}
	rexsult, err = lang.RxFromString("-5565747671734.1686009705574503152").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-490.30899494881071282787487030303"))
	if err != nil {
		lang.SayString("mulx3216")
	} else {
		if !(rexsult.ToString() == "2728936147066663.4580064428639745") {
			lang.SayString("mulx3216")
		}
	}
	rexsult, err = lang.RxFromString("319545511.89203495546689273564728E+036").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-2955943533943321899150310192061"))
	if err != nil {
		lang.SayString("mulx3217")
	} else {
		if !(rexsult.ToString() == "-9.4455848967786959996525702197139E+74") {
			lang.SayString("mulx3217")
		}
	}
	rexsult, err = lang.RxFromString("-36852134.84194296250843579428931").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-5830629.8347085025808716360357940"))
	if err != nil {
		lang.SayString("mulx3218")
	} else {
		if !(rexsult.ToString() == "214871156882133.34437417534873098") {
			lang.SayString("mulx3218")
		}
	}
	rexsult, err = lang.RxFromString("8.6021905001798578659275880018221E-374").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-39505285344943.729681835377530908"))
	if err != nil {
		lang.SayString("mulx3219")
	} else {
		if !(rexsult.ToString() == "-3.3983199030116951081865430362053E-360") {
			lang.SayString("mulx3219")
		}
	}
	rexsult, err = lang.RxFromString("-54863165.152174109720312887805017").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("736.1398476560169141105328256628"))
	if err != nil {
		lang.SayString("mulx3220")
	} else {
		if !(rexsult.ToString() == "-40386962037.048345148338122539405") {
			lang.SayString("mulx3220")
		}
	}
	rexsult, err = lang.RxFromString("-3263743464517851012531708810307").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("2457206.2471248382136273643208109"))
	if err != nil {
		lang.SayString("mulx3221")
	} else {
		if !(rexsult.ToString() == "-8.0196908300261262548565838031943E+36") {
			lang.SayString("mulx3221")
		}
	}
	rexsult, err = lang.RxFromString("2856586744.0548637797291151154902E-895").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("953545637646.57694835860339582821E+080"))
	if err != nil {
		lang.SayString("mulx3222")
	} else {
		if !(rexsult.ToString() == "2.7238858283525541854826594343954E-794") {
			lang.SayString("mulx3222")
		}
	}
	rexsult, err = lang.RxFromString("5624157233.3433661009203529937625").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("626098409265.93738586750090160638"))
	if err != nil {
		lang.SayString("mulx3223")
	} else {
		if !(rexsult.ToString() == "3521275897257796938833.8975037909") {
			lang.SayString("mulx3223")
		}
	}
	rexsult, err = lang.RxFromString("-213499362.91476998701834067092611").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("419272438.02555757699863022643444"))
	if err != nil {
		lang.SayString("mulx3224")
	} else {
		if !(rexsult.ToString() == "-89514398406178925.073260776410672") {
			lang.SayString("mulx3224")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("mulx3225")
	} else {
		if !(rexsult.ToString() == "145433.29080119336967191651199283") {
			lang.SayString("mulx3225")
		}
	}
	rexsult, err = lang.RxFromString("47.525676459351505682005359699680E+704").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-58631943508436657383369760970210"))
	if err != nil {
		lang.SayString("mulx3226")
	} else {
		if !(rexsult.ToString() == "-2.7865227773649353769876975366506E+737") {
			lang.SayString("mulx3226")
		}
	}
	rexsult, err = lang.RxFromString("-74396862273800.625679130772935550").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-115616605.52826981284183992013157"))
	if err != nil {
		lang.SayString("mulx3227")
	} else {
		if !(rexsult.ToString() == "8601512678051025297297.7169654467") {
			lang.SayString("mulx3227")
		}
	}
	rexsult, err = lang.RxFromString("67585560.562561229497720955705979").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("826.96290288608566737177503451613"))
	if err != nil {
		lang.SayString("mulx3228")
	} else {
		if !(rexsult.ToString() == "55890751355.998983433895910295596") {
			lang.SayString("mulx3228")
		}
	}
	rexsult, err = lang.RxFromString("6877386868.9498051860742298735156E-232").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("390.3154289860643509393769754551"))
	if err != nil {
		lang.SayString("mulx3229")
	} else {
		if !(rexsult.ToString() == "2.6843502060572691408091663822732E-220") {
			lang.SayString("mulx3229")
		}
	}
	rexsult, err = lang.RxFromString("-1647335.201144609256134925838937").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-186654823782.50459802235024230856"))
	if err != nil {
		lang.SayString("mulx3230")
	} else {
		if !(rexsult.ToString() == "307483061680363807.48775619333285") {
			lang.SayString("mulx3230")
		}
	}
	rexsult, err = lang.RxFromString("41407818140948.866630923934138155").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-5156.7624534000311342310106671627E-963"))
	if err != nil {
		lang.SayString("mulx3231")
	} else {
		if !(rexsult.ToString() == "-2.1353028186646179369220834691156E-946") {
			lang.SayString("mulx3231")
		}
	}
	rexsult, err = lang.RxFromString("-6.6547424012516834662011706165297").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-574454585580.06215974139884746441"))
	if err != nil {
		lang.SayString("mulx3232")
	} else {
		if !(rexsult.ToString() == "3822847288253.1035559206691532826") {
			lang.SayString("mulx3232")
		}
	}
	rexsult, err = lang.RxFromString("-27627.758745381267780885067447671").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-23385972189441.709586433758111062"))
	if err != nil {
		lang.SayString("mulx3233")
	} else {
		if !(rexsult.ToString() == "646101997676091306.41485393678655") {
			lang.SayString("mulx3233")
		}
	}
	rexsult, err = lang.RxFromString("209819.74379099914752963711944307E-228").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-440230.6700989532467831370320266E-960"))
	if err != nil {
		lang.SayString("mulx3234")
	} else {
		if !(rexsult.ToString() == "-9.2369086409102239573726316593648E-1178") {
			lang.SayString("mulx3234")
		}
	}
	rexsult, err = lang.RxFromString("2.3488457600415474270259330865184").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("9182434.6660212482500376220508605E-612"))
	if err != nil {
		lang.SayString("mulx3235")
	} else {
		if !(rexsult.ToString() == "2.1568122732142531556215204459407E-605") {
			lang.SayString("mulx3235")
		}
	}
	rexsult, err = lang.RxFromString("-5107586300197.9703941034404557409").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("56609.05486055057838678039496686E-768"))
	if err != nil {
		lang.SayString("mulx3236")
	} else {
		if !(rexsult.ToString() == "-2.8913563307290346152596212593532E-751") {
			lang.SayString("mulx3236")
		}
	}
	rexsult, err = lang.RxFromString("-70454070095869.70717871212601390").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-6200178.370249260009168888392406"))
	if err != nil {
		lang.SayString("mulx3237")
	} else {
		if !(rexsult.ToString() == "436827801504436566945.76663687924") {
			lang.SayString("mulx3237")
		}
	}
	rexsult, err = lang.RxFromString("29119.220621511046558757900645228").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("3517612.8810761470018676311863010E-222"))
	if err != nil {
		lang.SayString("mulx3238")
	} else {
		if !(rexsult.ToString() == "1.0243014554512542440592768088600E-211") {
			lang.SayString("mulx3238")
		}
	}
	rexsult, err = lang.RxFromString("-5168.2214111091132913776042214525").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-5690274.0971173476527123568627720"))
	if err != nil {
		lang.SayString("mulx3239")
	} else {
		if !(rexsult.ToString() == "29408596423.801454053855793898323") {
			lang.SayString("mulx3239")
		}
	}
	rexsult, err = lang.RxFromString("33783.060857197067391462144517964").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-2070.0806959465088198508322815406"))
	if err != nil {
		lang.SayString("mulx3240")
	} else {
		if !(rexsult.ToString() == "-69933662.130469766080574235843448") {
			lang.SayString("mulx3240")
		}
	}
	rexsult, err = lang.RxFromString("42207435091050.840296353874733169E-905").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("73330633078.828216018536326743325E+976"))
	if err != nil {
		lang.SayString("mulx3241")
	} else {
		if !(rexsult.ToString() == "3.0950979358603075650592433398939E+95") {
			lang.SayString("mulx3241")
		}
	}
	rexsult, err = lang.RxFromString("-71800.806700868784841045406679641").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-39617456964250697902519150526701"))
	if err != nil {
		lang.SayString("mulx3242")
	} else {
		if !(rexsult.ToString() == "2.8445653694701522164901827524538E+36") {
			lang.SayString("mulx3242")
		}
	}
	rexsult, err = lang.RxFromString("53627480801.631504892310016062160").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("328259.56947661049313311983109503"))
	if err != nil {
		lang.SayString("mulx3243")
	} else {
		if !(rexsult.ToString() == "17603733760058752.363123585224369") {
			lang.SayString("mulx3243")
		}
	}
	rexsult, err = lang.RxFromString("-5052601598.5559371338428368438728").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-97855372.224321664785314782556064"))
	if err != nil {
		lang.SayString("mulx3244")
	} else {
		if !(rexsult.ToString() == "494424210127893893.12581512954787") {
			lang.SayString("mulx3244")
		}
	}
	rexsult, err = lang.RxFromString("4208134320733.7069742988228068191E+146").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("4270869.1760149477598920960628392E+471"))
	if err != nil {
		lang.SayString("mulx3245")
	} else {
		if !(rexsult.ToString() == "1.7972391158952189002169082753183E+636") {
			lang.SayString("mulx3245")
		}
	}
	rexsult, err = lang.RxFromString("-8.5077009657942581515590471189084E+308").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("9652145155.374217047842114042376E-250"))
	if err != nil {
		lang.SayString("mulx3246")
	} else {
		if !(rexsult.ToString() == "-8.2117564660363596283732942091852E+68") {
			lang.SayString("mulx3246")
		}
	}
	rexsult, err = lang.RxFromString("-9504.9703032286960790904181078063E+619").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-86.245956949049186533469206485003"))
	if err != nil {
		lang.SayString("mulx3247")
	} else {
		if !(rexsult.ToString() == "8.1976525957425311427858087117655E+624") {
			lang.SayString("mulx3247")
		}
	}
	rexsult, err = lang.RxFromString("-440220916.66716743026896931194749").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-102725.01594377871560564824358775"))
	if err != nil {
		lang.SayString("mulx3248")
	} else {
		if !(rexsult.ToString() == "45221700683419.655596771711603505") {
			lang.SayString("mulx3248")
		}
	}
	rexsult, err = lang.RxFromString("-46.250735086006350517943464758019").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("14656357555174.263295266074908024"))
	if err != nil {
		lang.SayString("mulx3249")
	} else {
		if !(rexsult.ToString() == "-677867310610152.55569620459788530") {
			lang.SayString("mulx3249")
		}
	}
	rexsult, err = lang.RxFromString("-61641121299391.316420647102699627E+763").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-91896469863.461006903590004188187E+474"))
	if err != nil {
		lang.SayString("mulx3250")
	} else {
		if !(rexsult.ToString() == "5.6646014458394584921579417504939E+1261") {
			lang.SayString("mulx3250")
		}
	}
	rexsult, err = lang.RxFromString("96668419802749.555738010239087449E-838").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-19498732131365824921639467044927E-542"))
	if err != nil {
		lang.SayString("mulx3251")
	} else {
		if !(rexsult.ToString() == "-1.8849116232962331617140676274611E-1335") {
			lang.SayString("mulx3251")
		}
	}
	rexsult, err = lang.RxFromString("-8534543911197995906031245719519E+124").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("16487117050031.594886541650897974"))
	if err != nil {
		lang.SayString("mulx3252")
	} else {
		if !(rexsult.ToString() == "-1.4071002443255581217471698731240E+168") {
			lang.SayString("mulx3252")
		}
	}
	rexsult, err = lang.RxFromString("-62663404777.352508979582846931050").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("9.2570938837239134052589184917186E+916"))
	if err != nil {
		lang.SayString("mulx3253")
	} else {
		if !(rexsult.ToString() == "-5.8008102109774576654709018012876E+927") {
			lang.SayString("mulx3253")
		}
	}
	rexsult, err = lang.RxFromString("1.744601214474560992754529320172E-827").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-17.353669504253419489494030651507E-631"))
	if err != nil {
		lang.SayString("mulx3254")
	} else {
		if !(rexsult.ToString() == "-3.0275232892710668432895049546233E-1457") {
			lang.SayString("mulx3254")
		}
	}
	rexsult, err = lang.RxFromString("0367191549036702224827734853471").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("4410320662415266533763143837742E+721"))
	if err != nil {
		lang.SayString("mulx3255")
	} else {
		if !(rexsult.ToString() == "1.6194324757808363802947192054966E+781") {
			lang.SayString("mulx3255")
		}
	}
	rexsult, err = lang.RxFromString("097704116.4492566721965710365073").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-96736.400939809433556067504289145"))
	if err != nil {
		lang.SayString("mulx3256")
	} else {
		if !(rexsult.ToString() == "-9451544582305.1234805483449772252") {
			lang.SayString("mulx3256")
		}
	}
	rexsult, err = lang.RxFromString("19533298.147150158931958733807878").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("80.141668338350708476637377647025E-641"))
	if err != nil {
		lang.SayString("mulx3257")
	} else {
		if !(rexsult.ToString() == "1.5654311016630284502459158971272E-632") {
			lang.SayString("mulx3257")
		}
	}
	rexsult, err = lang.RxFromString("-23765003221220177430797028997378").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-15203369569.373411506379096871224E-945"))
	if err != nil {
		lang.SayString("mulx3258")
	} else {
		if !(rexsult.ToString() == "3.6130812678955994625210007005216E-904") {
			lang.SayString("mulx3258")
		}
	}
	rexsult, err = lang.RxFromString("129255.41937932433359193338910552E+932").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-281253953.38990382799508873560320"))
	if err != nil {
		lang.SayString("mulx3259")
	} else {
		if !(rexsult.ToString() == "-3.6353597697504958096931088780367E+945") {
			lang.SayString("mulx3259")
		}
	}
	rexsult, err = lang.RxFromString("-86863.276249466008289214762980838").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("531.50602652732088208397655484476"))
	if err != nil {
		lang.SayString("mulx3260")
	} else {
		if !(rexsult.ToString() == "-46168354.810498682140456143534524") {
			lang.SayString("mulx3260")
		}
	}
	rexsult, err = lang.RxFromString("-40707.169006771111855573524157083").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-68795521421321853333274411827749"))
	if err != nil {
		lang.SayString("mulx3261")
	} else {
		if !(rexsult.ToString() == "2.8004709174066910577370895499575E+36") {
			lang.SayString("mulx3261")
		}
	}
	rexsult, err = lang.RxFromString("-90838752568673.728630494658778003E+095").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-738.01370301217606577533107981431"))
	if err != nil {
		lang.SayString("mulx3262")
	} else {
		if !(rexsult.ToString() == "6.7040244160213718891633678248127E+111") {
			lang.SayString("mulx3262")
		}
	}
	rexsult, err = lang.RxFromString("-4245360967593.9206771555839718158E-682").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-3.119606239042530207103203508057"))
	if err != nil {
		lang.SayString("mulx3263")
	} else {
		if !(rexsult.ToString() == "1.3243854561493627844105290415330E-669") {
			lang.SayString("mulx3263")
		}
	}
	rexsult, err = lang.RxFromString("-3422145405774.0800213000547612240E-023").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-60810.964656409650839011321706310"))
	if err != nil {
		lang.SayString("mulx3264")
	} else {
		if !(rexsult.ToString() == "0.0000020810396331962224323288744910607") {
			lang.SayString("mulx3264")
		}
	}
	rexsult, err = lang.RxFromString("-24521811.07649485796598387627478E-661").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-94860846133404815410816234000694"))
	if err != nil {
		lang.SayString("mulx3265")
	} else {
		if !(rexsult.ToString() == "2.3261597474398006215017751785104E-622") {
			lang.SayString("mulx3265")
		}
	}
	rexsult, err = lang.RxFromString("-5042529937498.8944492248538951438").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("3891904674.4549170968807145612549"))
	if err != nil {
		lang.SayString("mulx3266")
	} else {
		if !(rexsult.ToString() == "-19625045834830808256871.952659048") {
			lang.SayString("mulx3266")
		}
	}
	rexsult, err = lang.RxFromString("-535824163.54531747646293693868651E-665").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("2732988.5891363639325008206099712"))
	if err != nil {
		lang.SayString("mulx3267")
	} else {
		if !(rexsult.ToString() == "-1.4644013247528895376254850705597E-650") {
			lang.SayString("mulx3267")
		}
	}
	rexsult, err = lang.RxFromString("24032.702008553084252925140858134E-509").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("52864854.899420632375589206704068"))
	if err != nil {
		lang.SayString("mulx3268")
	} else {
		if !(rexsult.ToString() == "1.2704853045231735885074945710576E-497") {
			lang.SayString("mulx3268")
		}
	}
	rexsult, err = lang.RxFromString("71553220259.938950698030519906727E-496").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("754.44220417415325444943566016062"))
	if err != nil {
		lang.SayString("mulx3269")
	} else {
		if !(rexsult.ToString() == "5.3982769208667021044675146787248E-483") {
			lang.SayString("mulx3269")
		}
	}
	rexsult, err = lang.RxFromString("35572.960284795962697740953932508").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("520.39506364687594082725754878910E-731"))
	if err != nil {
		lang.SayString("mulx3270")
	} else {
		if !(rexsult.ToString() == "1.8511992931514185102474609686066E-724") {
			lang.SayString("mulx3270")
		}
	}
	rexsult, err = lang.RxFromString("53035405018123760598334895413057E+818").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-9558464247240.4476790042911379151"))
	if err != nil {
		lang.SayString("mulx3271")
	} else {
		if !(rexsult.ToString() == "-5.0693702270365259274203181894613E+862") {
			lang.SayString("mulx3271")
		}
	}
	rexsult, err = lang.RxFromString("95.490751127249945886828257312118").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("987.01498316307365714167410690192E+133"))
	if err != nil {
		lang.SayString("mulx3272")
	} else {
		if !(rexsult.ToString() == "9.4250802116091862185764800227004E+137") {
			lang.SayString("mulx3272")
		}
	}
	rexsult, err = lang.RxFromString("69434850287.460788550936730296153").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-35119136549015044241569827542264"))
	if err != nil {
		lang.SayString("mulx3273")
	} else {
		if !(rexsult.ToString() == "-2.4384919885057519302646522425980E+42") {
			lang.SayString("mulx3273")
		}
	}
	rexsult, err = lang.RxFromString("-392.22739924621965621739098725407").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-65551274.987160998195282109612136"))
	if err != nil {
		lang.SayString("mulx3274")
	} else {
		if !(rexsult.ToString() == "25711006105.487929108329637701882") {
			lang.SayString("mulx3274")
		}
	}
	rexsult, err = lang.RxFromString("6413265.4423561191792972085539457").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("24514.222704714139350026165721146"))
	if err != nil {
		lang.SayString("mulx3275")
	} else {
		if !(rexsult.ToString() == "157216217318.36494525300694583138") {
			lang.SayString("mulx3275")
		}
	}
	rexsult, err = lang.RxFromString("-6.9667706389122107760046184064057E+487").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("32.405810703971538278419625527234"))
	if err != nil {
		lang.SayString("mulx3276")
	} else {
		if !(rexsult.ToString() == "-2.2576385054257595259511556258470E+489") {
			lang.SayString("mulx3276")
		}
	}
	rexsult, err = lang.RxFromString("378204716633.40024100602896057615").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-0300218714378.322231269606216516"))
	if err != nil {
		lang.SayString("mulx3277")
	} else {
		if !(rexsult.ToString() == "-113544133799497082075557.21180430") {
			lang.SayString("mulx3277")
		}
	}
	rexsult, err = lang.RxFromString("-44234.512012457148027685282969235E+501").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("2132572.4571987908375002100894927"))
	if err != nil {
		lang.SayString("mulx3278")
	} else {
		if !(rexsult.ToString() == "-9.4333301975395170465982968019915E+511") {
			lang.SayString("mulx3278")
		}
	}
	rexsult, err = lang.RxFromString("-3554.5895974968741465654022772100E-073").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("9752.0428746722497621936998533848E+516"))
	if err != nil {
		lang.SayString("mulx3279")
	} else {
		if !(rexsult.ToString() == "-3.4664510156653491769901435777060E+450") {
			lang.SayString("mulx3279")
		}
	}
	rexsult, err = lang.RxFromString("750187685.63632608923397234391668").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("4633194252863.6730625972669246025"))
	if err != nil {
		lang.SayString("mulx3280")
	} else {
		if !(rexsult.ToString() == "3475765273659325895012.7612107556") {
			lang.SayString("mulx3280")
		}
	}
	rexsult, err = lang.RxFromString("30190034242853.251165951457589386E-028").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("8038885676.3204238322976087799751E+018"))
	if err != nil {
		lang.SayString("mulx3281")
	} else {
		if !(rexsult.ToString() == "24269423384249.611263728854793731") {
			lang.SayString("mulx3281")
		}
	}
	rexsult, err = lang.RxFromString("65.537942676774715953400768803539").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("125946917260.87536506197191782198"))
	if err != nil {
		lang.SayString("mulx3282")
	} else {
		if !(rexsult.ToString() == "8254301843759.7376990957355411370") {
			lang.SayString("mulx3282")
		}
	}
	rexsult, err = lang.RxFromString("8015272348677.5489394183881813700").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("949.23027111499779258284877920022"))
	if err != nil {
		lang.SayString("mulx3283")
	} else {
		if !(rexsult.ToString() == "7608339144595734.8984281431471741") {
			lang.SayString("mulx3283")
		}
	}
	rexsult, err = lang.RxFromString("-32595333982.67068622120451804225").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("69130.052233649808750113141502465E-862"))
	if err != nil {
		lang.SayString("mulx3284")
	} else {
		if !(rexsult.ToString() == "-2.2533171407952851885446213697715E-847") {
			lang.SayString("mulx3284")
		}
	}
	rexsult, err = lang.RxFromString("-17544189017145.710117633021800426E-537").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("292178000.06450804618299520094843"))
	if err != nil {
		lang.SayString("mulx3285")
	} else {
		if !(rexsult.ToString() == "-5.1260260597833406461110136952456E-516") {
			lang.SayString("mulx3285")
		}
	}
	rexsult, err = lang.RxFromString("-506650.99395649907139204052441630").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("11.018427502031650148962067367158"))
	if err != nil {
		lang.SayString("mulx3286")
	} else {
		if !(rexsult.ToString() == "-5582497.2457419607392940234271222") {
			lang.SayString("mulx3286")
		}
	}
	rexsult, err = lang.RxFromString("4846835159.5922372307656000769395E-241").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-84.001893040865864590122330800768"))
	if err != nil {
		lang.SayString("mulx3287")
	} else {
		if !(rexsult.ToString() == "-4.0714332866277514481192856925775E-230") {
			lang.SayString("mulx3287")
		}
	}
	rexsult, err = lang.RxFromString("-35.029311013822259358116556164908").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-3994308878.1994645313414534209127"))
	if err != nil {
		lang.SayString("mulx3288")
	} else {
		if !(rexsult.ToString() == "139917887979.72053637272961120639") {
			lang.SayString("mulx3288")
		}
	}
	rexsult, err = lang.RxFromString("7606663750.6735265233044420887018E+166").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-5491814639.4484565418284686127552E+365"))
	if err != nil {
		lang.SayString("mulx3289")
	} else {
		if !(rexsult.ToString() == "-4.1774387343310777190917128006589E+550") {
			lang.SayString("mulx3289")
		}
	}
	rexsult, err = lang.RxFromString("-25677.829660831741274207326697052E-163").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-94135395124193048560172705082029E-862"))
	if err != nil {
		lang.SayString("mulx3290")
	} else {
		if !(rexsult.ToString() == "2.4171926410541199393728294762559E-989") {
			lang.SayString("mulx3290")
		}
	}
	rexsult, err = lang.RxFromString("97271576094.456406729671729224827").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-1.5412563837540810793697955063295E+554"))
	if err != nil {
		lang.SayString("mulx3291")
	} else {
		if !(rexsult.ToString() == "-1.4992043761340180288065959300090E+565") {
			lang.SayString("mulx3291")
		}
	}
	rexsult, err = lang.RxFromString("41139789894.401826915757391777563").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-1.8729920305671057957156159690823E-020"))
	if err != nil {
		lang.SayString("mulx3292")
	} else {
		if !(rexsult.ToString() == "-7.7054498611419776714291080928601E-10") {
			lang.SayString("mulx3292")
		}
	}
	rexsult, err = lang.RxFromString("-83310831287241.777598696853498149").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-54799825033.797100418162985103519E-330"))
	if err != nil {
		lang.SayString("mulx3293")
	} else {
		if !(rexsult.ToString() == "4.5654189779610386760330527839588E-306") {
			lang.SayString("mulx3293")
		}
	}
	rexsult, err = lang.RxFromString("4506653461.4414974233678331771169").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-74955470.977653038010031457181957E-887"))
	if err != nil {
		lang.SayString("mulx3294")
	} else {
		if !(rexsult.ToString() == "-3.3779833273541776470902903512949E-870") {
			lang.SayString("mulx3294")
		}
	}
	rexsult, err = lang.RxFromString("23777.857951114967684767609401626").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("720760.03897144157012301385227528"))
	if err != nil {
		lang.SayString("mulx3295")
	} else {
		if !(rexsult.ToString() == "17138129823.503025913034987537096") {
			lang.SayString("mulx3295")
		}
	}
	rexsult, err = lang.RxFromString("-21337853323980858055292469611948").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("6080272840.3071490445256786982100E+532"))
	if err != nil {
		lang.SayString("mulx3296")
	} else {
		if !(rexsult.ToString() == "-1.2973997003625843317417981902198E+573") {
			lang.SayString("mulx3296")
		}
	}
	rexsult, err = lang.RxFromString("-818409238.0423893439849743856947").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("756.39156265972753259267069158760"))
	if err != nil {
		lang.SayString("mulx3297")
	} else {
		if !(rexsult.ToString() == "-619037842458.03980537370328252817") {
			lang.SayString("mulx3297")
		}
	}
	rexsult, err = lang.RxFromString("47567380384943.87013600286155046").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("65.084709374908275826942076480326"))
	if err != nil {
		lang.SayString("mulx3298")
	} else {
		if !(rexsult.ToString() == "3095909128079784.3348591472999468") {
			lang.SayString("mulx3298")
		}
	}
	rexsult, err = lang.RxFromString("-296615544.05897984545127115290251").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-5416115.4315053536014016450973264"))
	if err != nil {
		lang.SayString("mulx3299")
	} else {
		if !(rexsult.ToString() == "1606504025402196.8484885386501478") {
			lang.SayString("mulx3299")
		}
	}
	rexsult, err = lang.RxFromString("61391705914.046707180652185247584E+739").OpMult(lang.RxSetWithDigit(32), lang.RxFromString("-675982087721.91856456125242297346"))
	if err != nil {
		lang.SayString("mulx3300")
	} else {
		if !(rexsult.ToString() == "-4.1499693532587347944890300176290E+761") {
			lang.SayString("mulx3300")
		}
	}
	rexsult, err = lang.RxFromString("042.668056830485571428877212944418").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-01364112374639.4941124016455321071"))
	if err != nil {
		lang.SayString("mulx3401")
	} else {
		if !(rexsult.ToString() == "-58204024324286.5595453066065234923") {
			lang.SayString("mulx3401")
		}
	}
	rexsult, err = lang.RxFromString("-327.179426341653256363531809227344E+453").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("759067457.107518663444899356760793"))
	if err != nil {
		lang.SayString("mulx3402")
	} else {
		if !(rexsult.ToString() == "-2.48351255171055445110558613627379E+464") {
			lang.SayString("mulx3402")
		}
	}
	rexsult, err = lang.RxFromString("81721.5803096185422787702538493471").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("900099473.245809628076790757217328"))
	if err != nil {
		lang.SayString("mulx3403")
	} else {
		if !(rexsult.ToString() == "73557551389502.7779979042453102926") {
			lang.SayString("mulx3403")
		}
	}
	rexsult, err = lang.RxFromString("3991.56734635183403814636354392163E-807").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("72.3239822255871305731314565069132"))
	if err != nil {
		lang.SayString("mulx3404")
	} else {
		if !(rexsult.ToString() == "2.88686045809784034794803928177854E-802") {
			lang.SayString("mulx3404")
		}
	}
	rexsult, err = lang.RxFromString("-66.3393308595957726456416979163306").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("5.08486573050459213864684589662559"))
	if err != nil {
		lang.SayString("mulx3405")
	} else {
		if !(rexsult.ToString() == "-337.326590072564290813539036280205") {
			lang.SayString("mulx3405")
		}
	}
	rexsult, err = lang.RxFromString("-393606.873703567753255097095208112E+111").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-2124339550.86891093200758095660557"))
	if err != nil {
		lang.SayString("mulx3406")
	} else {
		if !(rexsult.ToString() == "8.36154649302353269818801263275941E+125") {
			lang.SayString("mulx3406")
		}
	}
	rexsult, err = lang.RxFromString("-019133598.609812524622150421584346").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-858439846.628367734642622922030051"))
	if err != nil {
		lang.SayString("mulx3407")
	} else {
		if !(rexsult.ToString() == "16425043456056213.7395191514029288") {
			lang.SayString("mulx3407")
		}
	}
	rexsult, err = lang.RxFromString("465.351982159046525762715549761814").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("240444.975944666924657629172844780"))
	if err != nil {
		lang.SayString("mulx3408")
	} else {
		if !(rexsult.ToString() == "111891546.156035013780371395668674") {
			lang.SayString("mulx3408")
		}
	}
	rexsult, err = lang.RxFromString("28066955004783.1076824222873384828").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("571699969.220753535758504907561016E-718"))
	if err != nil {
		lang.SayString("mulx3409")
	} else {
		if !(rexsult.ToString() == "1.60458773123547770690452195569223E-696") {
			lang.SayString("mulx3409")
		}
	}
	rexsult, err = lang.RxFromString("28275236927392.4960902824105246047").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("28212038.4825243127096613158419270E+422"))
	if err != nil {
		lang.SayString("mulx3410")
	} else {
		if !(rexsult.ToString() == "7.97702072298089605706798770013561E+442") {
			lang.SayString("mulx3410")
		}
	}
	rexsult, err = lang.RxFromString("11791.8644211874630234271801789996").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-8.45457275930363860982261343159741"))
	if err != nil {
		lang.SayString("mulx3411")
	} else {
		if !(rexsult.ToString() == "-99695.1757167732926302533138186716") {
			lang.SayString("mulx3411")
		}
	}
	rexsult, err = lang.RxFromString("44.7085340739581668975502342787578").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-9337.05408133023920640485556647937"))
	if err != nil {
		lang.SayString("mulx3412")
	} else {
		if !(rexsult.ToString() == "-417446.000545543168866158913077419") {
			lang.SayString("mulx3412")
		}
	}
	rexsult, err = lang.RxFromString("93354527428804.5458053295581965867E+576").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-856525909852.318790321300941615314"))
	if err != nil {
		lang.SayString("mulx3413")
	} else {
		if !(rexsult.ToString() == "-7.99605715447900642683774360731254E+601") {
			lang.SayString("mulx3413")
		}
	}
	rexsult, err = lang.RxFromString("-367399415798804503177950040443482").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-54845683.9691776397285506712812754"))
	if err != nil {
		lang.SayString("mulx3414")
	} else {
		if !(rexsult.ToString() == "2.01502722493617222018040789291414E+40") {
			lang.SayString("mulx3414")
		}
	}
	rexsult, err = lang.RxFromString("-2.87155919781024108503670175443740").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("89529730130.6427881332776797193807"))
	if err != nil {
		lang.SayString("mulx3415")
	} else {
		if !(rexsult.ToString() == "-257089920034.115975469931085527642") {
			lang.SayString("mulx3415")
		}
	}
	rexsult, err = lang.RxFromString("-010.693934338179479652178057279204E+188").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("26484.8887731973153745666514260684"))
	if err != nil {
		lang.SayString("mulx3416")
	} else {
		if !(rexsult.ToString() == "-2.83227661494558963558481633880647E+193") {
			lang.SayString("mulx3416")
		}
	}
	rexsult, err = lang.RxFromString("611655569568.832698912762075889186").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("010182743219.475839030505966016982"))
	if err != nil {
		lang.SayString("mulx3417")
	} else {
		if !(rexsult.ToString() == "6228331603681663511826.60450280350") {
			lang.SayString("mulx3417")
		}
	}
	rexsult, err = lang.RxFromString("3457947.39062863674882672518304442").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-01.9995218868908849056866549811425"))
	if err != nil {
		lang.SayString("mulx3418")
	} else {
		if !(rexsult.ToString() == "-6914241.49127918361259252956576654") {
			lang.SayString("mulx3418")
		}
	}
	rexsult, err = lang.RxFromString("-53308666960535.7393391289364591513").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-6527.00547629475578694521436764596E-442"))
	if err != nil {
		lang.SayString("mulx3419")
	} else {
		if !(rexsult.ToString() == "3.47945961185390084641156250100085E-425") {
			lang.SayString("mulx3419")
		}
	}
	rexsult, err = lang.RxFromString("-5568057.17870139549478277980540034").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-407906443.141342175740471849723638"))
	if err != nil {
		lang.SayString("mulx3420")
	} else {
		if !(rexsult.ToString() == "2271246398971702.91169807728132089") {
			lang.SayString("mulx3420")
		}
	}
	rexsult, err = lang.RxFromString("9804385273.49533524416415189990857").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("84.1433929743544659553964804646569"))
	if err != nil {
		lang.SayString("mulx3421")
	} else {
		if !(rexsult.ToString() == "824974242939.691780798621180901714") {
			lang.SayString("mulx3421")
		}
	}
	rexsult, err = lang.RxFromString("-5234910986592.18801727046580014273E-547").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-5874220715892.91440069210512515154"))
	if err != nil {
		lang.SayString("mulx3422")
	} else {
		if !(rexsult.ToString() == "3.07510225632952455144944282925583E-522") {
			lang.SayString("mulx3422")
		}
	}
	rexsult, err = lang.RxFromString("698416560151955285929747633786867E-495").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("51754681.6784872628933218985216916E-266"))
	if err != nil {
		lang.SayString("mulx3423")
	} else {
		if !(rexsult.ToString() == "3.61463267496484976064271305679796E-721") {
			lang.SayString("mulx3423")
		}
	}
	rexsult, err = lang.RxFromString("107635.497735316515080720330536027").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-3972075.83989512668362609609006425E-605"))
	if err != nil {
		lang.SayString("mulx3424")
	} else {
		if !(rexsult.ToString() == "-4.27536360069537352698066408021773E-594") {
			lang.SayString("mulx3424")
		}
	}
	rexsult, err = lang.RxFromString("-32174291345686.5371446616670961807").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("79518863759385.5925052747867099091E+408"))
	if err != nil {
		lang.SayString("mulx3425")
	} else {
		if !(rexsult.ToString() == "-2.55846309007242668513226814043593E+435") {
			lang.SayString("mulx3425")
		}
	}
	rexsult, err = lang.RxFromString("-8151730494.53190523620899410544099E+688").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-93173.0631474527142307644239919480E+900"))
	if err != nil {
		lang.SayString("mulx3426")
	} else {
		if !(rexsult.ToString() == "7.59521700128037149179751467730962E+1602") {
			lang.SayString("mulx3426")
		}
	}
	rexsult, err = lang.RxFromString("1.33649801345976199708341799505220").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-56623.0530039528969825480755159562E+459"))
	if err != nil {
		lang.SayString("mulx3427")
	} else {
		if !(rexsult.ToString() == "-7.56765978558098558928268129700052E+463") {
			lang.SayString("mulx3427")
		}
	}
	rexsult, err = lang.RxFromString("67762238162788.6551061476018185196").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-6140.75837959248100352788853809376E-822"))
	if err != nil {
		lang.SayString("mulx3428")
	} else {
		if !(rexsult.ToString() == "-4.16111531818085838717201828773857E-805") {
			lang.SayString("mulx3428")
		}
	}
	rexsult, err = lang.RxFromString("4286562.76568866751577306056498271").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("6286.77291578497580015557979349893E+820"))
	if err != nil {
		lang.SayString("mulx3429")
	} else {
		if !(rexsult.ToString() == "2.69486466971438542975159893306219E+830") {
			lang.SayString("mulx3429")
		}
	}
	rexsult, err = lang.RxFromString("-765782.827432642697305644096365566").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("67.1634368459576834692758114618652"))
	if err != nil {
		lang.SayString("mulx3430")
	} else {
		if !(rexsult.ToString() == "-51432606.5679912088468256122315944") {
			lang.SayString("mulx3430")
		}
	}
	rexsult, err = lang.RxFromString("46.2835931916106252756465724211276").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("59.2989237834093118332826617957791"))
	if err != nil {
		lang.SayString("mulx3431")
	} else {
		if !(rexsult.ToString() == "2744.56726509164060561370653286614") {
			lang.SayString("mulx3431")
		}
	}
	rexsult, err = lang.RxFromString("-3029555.82298840234029474459694644").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("857535844655004737373089601128532"))
	if err != nil {
		lang.SayString("mulx3432")
	} else {
		if !(rexsult.ToString() == "-2.59795271159584761928986181925721E+39") {
			lang.SayString("mulx3432")
		}
	}
	rexsult, err = lang.RxFromString("-0138466789523.10694176543700501945E-948").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("481026979918882487383654367924619"))
	if err != nil {
		lang.SayString("mulx3433")
	} else {
		if !(rexsult.ToString() == "-6.66062615833636908683785283687416E-905") {
			lang.SayString("mulx3433")
		}
	}
	rexsult, err = lang.RxFromString("-9593566466.96690575714244442109870").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-87632034347.4845477961976776833770E+769"))
	if err != nil {
		lang.SayString("mulx3434")
	} else {
		if !(rexsult.ToString() == "8.40703746148119867711463485065336E+789") {
			lang.SayString("mulx3434")
		}
	}
	rexsult, err = lang.RxFromString("-3189.30765477670526823106100241863E-898").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("565688889.355241946154894311253202E-466"))
	if err != nil {
		lang.SayString("mulx3435")
	} else {
		if !(rexsult.ToString() == "-1.80415590504280580443565448126548E-1352") {
			lang.SayString("mulx3435")
		}
	}
	rexsult, err = lang.RxFromString("-17084552395.6714834680088150543965").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-631925802672.685034379197328370812E+527"))
	if err != nil {
		lang.SayString("mulx3436")
	} else {
		if !(rexsult.ToString() == "1.07961694859382462346866817306769E+549") {
			lang.SayString("mulx3436")
		}
	}
	rexsult, err = lang.RxFromString("034956830.349823306815911887469760").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-61600816.0672274126966042956781665E-667"))
	if err != nil {
		lang.SayString("mulx3437")
	} else {
		if !(rexsult.ToString() == "-2.15336927667273841617128781173293E-652") {
			lang.SayString("mulx3437")
		}
	}
	rexsult, err = lang.RxFromString("-763.440067781256632695791981893608").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("19.9263811350611007833220620745413"))
	if err != nil {
		lang.SayString("mulx3438")
	} else {
		if !(rexsult.ToString() == "-15212.5977643862002585039364868883") {
			lang.SayString("mulx3438")
		}
	}
	rexsult, err = lang.RxFromString("-510472027868440667684575147556654E+789").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("834872378550801889983927148587909"))
	if err != nil {
		lang.SayString("mulx3439")
	} else {
		if !(rexsult.ToString() == "-4.26178996090176289115594057419892E+854") {
			lang.SayString("mulx3439")
		}
	}
	rexsult, err = lang.RxFromString("070304761.560517086676993503034828E-094").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-17773.7446959771077104057845273992E-761"))
	if err != nil {
		lang.SayString("mulx3440")
	} else {
		if !(rexsult.ToString() == "-1.24957888288817581538108991453732E-843") {
			lang.SayString("mulx3440")
		}
	}
	rexsult, err = lang.RxFromString("-0970725697662.27605454336231195463").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-4541.41897546697187157913886433474"))
	if err != nil {
		lang.SayString("mulx3441")
	} else {
		if !(rexsult.ToString() == "4408472103336875.21161867891724392") {
			lang.SayString("mulx3441")
		}
	}
	rexsult, err = lang.RxFromString("-808178238631844268316111259558675").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-598400.265108644514211244980426520"))
	if err != nil {
		lang.SayString("mulx3442")
	} else {
		if !(rexsult.ToString() == "4.83614072252332979731348423145208E+38") {
			lang.SayString("mulx3442")
		}
	}
	rexsult, err = lang.RxFromString("-9.90826595069053564311371766315200").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-031.625916781307847864872329806646"))
	if err != nil {
		lang.SayString("mulx3443")
	} else {
		if !(rexsult.ToString() == "313.357994403604968250936036978086") {
			lang.SayString("mulx3443")
		}
	}
	rexsult, err = lang.RxFromString("-196925.469891897719160698483752907").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-41268.9975444533794067723958739778"))
	if err != nil {
		lang.SayString("mulx3444")
	} else {
		if !(rexsult.ToString() == "8126916733.40905487026003135987472") {
			lang.SayString("mulx3444")
		}
	}
	rexsult, err = lang.RxFromString("421071135212152225162086005824310").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("1335320330.08964354845796510145246E-604"))
	if err != nil {
		lang.SayString("mulx3445")
	} else {
		if !(rexsult.ToString() == "5.62264847262712040027311932121460E-563") {
			lang.SayString("mulx3445")
		}
	}
	rexsult, err = lang.RxFromString("1249441.46421514282301182772247227").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-0289848.71208912281976374705180836E-676"))
	if err != nil {
		lang.SayString("mulx3446")
	} else {
		if !(rexsult.ToString() == "-3.62148999233506984566620611700349E-665") {
			lang.SayString("mulx3446")
		}
	}
	rexsult, err = lang.RxFromString("74815000.4716875558358937279052903").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-690425401708167622194241915195001E+891"))
	if err != nil {
		lang.SayString("mulx3447")
	} else {
		if !(rexsult.ToString() == "-5.16541767544616308732028810026275E+931") {
			lang.SayString("mulx3447")
		}
	}
	rexsult, err = lang.RxFromString("-1683993.51210241555668790556759021").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-72394384927344.8402585228267493374"))
	if err != nil {
		lang.SayString("mulx3448")
	} else {
		if !(rexsult.ToString() == "121911674530293613615.441384822381") {
			lang.SayString("mulx3448")
		}
	}
	rexsult, err = lang.RxFromString("-763.148530974741766171756970448158").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("517370.808956957601473642272664647"))
	if err != nil {
		lang.SayString("mulx3449")
	} else {
		if !(rexsult.ToString() == "-394830772.824715962925351447322187") {
			lang.SayString("mulx3449")
		}
	}
	rexsult, err = lang.RxFromString("-77.5841338812312523460591226178754").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-927540422.641025050968830154578151E+524"))
	if err != nil {
		lang.SayString("mulx3450")
	} else {
		if !(rexsult.ToString() == "7.19624203304351070562409746475943E+534") {
			lang.SayString("mulx3450")
		}
	}
	rexsult, err = lang.RxFromString("5176295309.89943746236102209837813").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-129733.103628797477167908698565465"))
	if err != nil {
		lang.SayString("mulx3451")
	} else {
		if !(rexsult.ToString() == "-671536855852442.071887385512001794") {
			lang.SayString("mulx3451")
		}
	}
	rexsult, err = lang.RxFromString("4471634841166.90197229286530307326E+172").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("31511104397.8671727003201890825879E-955"))
	if err != nil {
		lang.SayString("mulx3452")
	} else {
		if !(rexsult.ToString() == "1.40906152309150441010046222214810E-760") {
			lang.SayString("mulx3452")
		}
	}
	rexsult, err = lang.RxFromString("-8189130.15945231049670285726774317").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("2.57912402871404325831670923864936E-366"))
	if err != nil {
		lang.SayString("mulx3453")
	} else {
		if !(rexsult.ToString() == "-2.11207823685103185039979144161848E-359") {
			lang.SayString("mulx3453")
		}
	}
	rexsult, err = lang.RxFromString("-254346232981353293392888785643245").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-764.416902486152093758287929678445"))
	if err != nil {
		lang.SayString("mulx3454")
	} else {
		if !(rexsult.ToString() == "1.94426559574627262006307326336289E+35") {
			lang.SayString("mulx3454")
		}
	}
	rexsult, err = lang.RxFromString("-2875.36745499544177964804277329726").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-13187.8492045054802205691248267638"))
	if err != nil {
		lang.SayString("mulx3455")
	} else {
		if !(rexsult.ToString() == "37919912.4040225840727281633024665") {
			lang.SayString("mulx3455")
		}
	}
	rexsult, err = lang.RxFromString("-7.26927570984219944693602140223103").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("0160883021541.32275286769110003971E-438"))
	if err != nil {
		lang.SayString("mulx3456")
	} else {
		if !(rexsult.ToString() == "-1.16950304061635681891361504442479E-426") {
			lang.SayString("mulx3456")
		}
	}
	rexsult, err = lang.RxFromString("-28567151.6868762752241056540028515E+498").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-4470.15455137546427645290210989675"))
	if err != nil {
		lang.SayString("mulx3457")
	} else {
		if !(rexsult.ToString() == "1.27699583132923253605397736797000E+509") {
			lang.SayString("mulx3457")
		}
	}
	rexsult, err = lang.RxFromString("7191753.79974136447257468282073718").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("81.3878426176038487948375777384429"))
	if err != nil {
		lang.SayString("mulx3458")
	} else {
		if !(rexsult.ToString() == "585321326.397904638863485891524555") {
			lang.SayString("mulx3458")
		}
	}
	rexsult, err = lang.RxFromString("502975804.069864536104621700404770").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("684.790028432074527960269515227243"))
	if err != nil {
		lang.SayString("mulx3459")
	} else {
		if !(rexsult.ToString() == "344432815169.648082754214631086270") {
			lang.SayString("mulx3459")
		}
	}
	rexsult, err = lang.RxFromString("1505368.42063731861590460453659570").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-465242.678439951462767630022819105"))
	if err != nil {
		lang.SayString("mulx3460")
	} else {
		if !(rexsult.ToString() == "-700361636056.225618266296899048765") {
			lang.SayString("mulx3460")
		}
	}
	rexsult, err = lang.RxFromString("69225023.2850142784063416137144829").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("8584050.06648191798834819995325693"))
	if err != nil {
		lang.SayString("mulx3461")
	} else {
		if !(rexsult.ToString() == "594231065731939.137329770485497261") {
			lang.SayString("mulx3461")
		}
	}
	rexsult, err = lang.RxFromString("-55669501853.7751006841940471339310E+614").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("061400538.186044693233816566977189"))
	if err != nil {
		lang.SayString("mulx3462")
	} else {
		if !(rexsult.ToString() == "-3.41813737437080390787865389703565E+632") {
			lang.SayString("mulx3462")
		}
	}
	rexsult, err = lang.RxFromString("-527566.521273992424649346837337602E-408").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-834662.599983953345718523807123972"))
	if err != nil {
		lang.SayString("mulx3463")
	} else {
		if !(rexsult.ToString() == "4.40340044311040151960763108019957E-397") {
			lang.SayString("mulx3463")
		}
	}
	rexsult, err = lang.RxFromString("69065510.0459653699418083455335366").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("694848643848212520086960886818157E-853"))
	if err != nil {
		lang.SayString("mulx3464")
	} else {
		if !(rexsult.ToString() == "4.79900759921241352562381181332720E-813") {
			lang.SayString("mulx3464")
		}
	}
	rexsult, err = lang.RxFromString("-2921982.82411285505229122890041841").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-72994.6523840298017471962569778803E-763"))
	if err != nil {
		lang.SayString("mulx3465")
	} else {
		if !(rexsult.ToString() == "2.13289120518223547921212412642411E-752") {
			lang.SayString("mulx3465")
		}
	}
	rexsult, err = lang.RxFromString("4.51117459466491451401835756593793").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3873385.19981811640063144338087230"))
	if err != nil {
		lang.SayString("mulx3466")
	} else {
		if !(rexsult.ToString() == "17473516.9087705701652062546164705") {
			lang.SayString("mulx3466")
		}
	}
	rexsult, err = lang.RxFromString("49553763474698.8118661758811091120").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("36.1713861293896216593840817950781E+410"))
	if err != nil {
		lang.SayString("mulx3467")
	} else {
		if !(rexsult.ToString() == "1.79242831280777466554271332425735E+425") {
			lang.SayString("mulx3467")
		}
	}
	rexsult, err = lang.RxFromString("755985583344.379951506071499170749E+956").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("746921095569971477373643487285780"))
	if err != nil {
		lang.SayString("mulx3468")
	} else {
		if !(rexsult.ToString() == "5.64661580146688255286933753616580E+1000") {
			lang.SayString("mulx3468")
		}
	}
	rexsult, err = lang.RxFromString("-20101668541.7472260497594230105456").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-395562148.345003931161532101109964"))
	if err != nil {
		lang.SayString("mulx3469")
	} else {
		if !(rexsult.ToString() == "7951459193692715079.09328760016914") {
			lang.SayString("mulx3469")
		}
	}
	rexsult, err = lang.RxFromString("5583903.18072100716072011264668631").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("460868914694.088387067451312500723"))
	if err != nil {
		lang.SayString("mulx3470")
	} else {
		if !(rexsult.ToString() == "2573447398655758659.39475672905225") {
			lang.SayString("mulx3470")
		}
	}
	rexsult, err = lang.RxFromString("-955638397975240685017992436116257E+020").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-508580.148958769104511751975720470E+662"))
	if err != nil {
		lang.SayString("mulx3471")
	} else {
		if !(rexsult.ToString() == "4.86018718792967378985838739820108E+720") {
			lang.SayString("mulx3471")
		}
	}
	rexsult, err = lang.RxFromString("-136243796098020983814115584402407E+796").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("6589108083.99750311651581032447390"))
	if err != nil {
		lang.SayString("mulx3472")
	} else {
		if !(rexsult.ToString() == "-8.97725098263977535966921696143011E+837") {
			lang.SayString("mulx3472")
		}
	}
	rexsult, err = lang.RxFromString("-808498.482718304598213092937543934E+521").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("48005.1465097914355096301483788905"))
	if err != nil {
		lang.SayString("mulx3473")
	} else {
		if !(rexsult.ToString() == "-3.88120881158362912220132691803539E+531") {
			lang.SayString("mulx3473")
		}
	}
	rexsult, err = lang.RxFromString("-812.266340554281305985524813074211E+396").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-3195.63111559114001594257448970812E+986"))
	if err != nil {
		lang.SayString("mulx3474")
	} else {
		if !(rexsult.ToString() == "2.59570359202261082537505332325404E+1388") {
			lang.SayString("mulx3474")
		}
	}
	rexsult, err = lang.RxFromString("-929889.720905183397678866648217793E+134").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-280300190774.057595671079264841349"))
	if err != nil {
		lang.SayString("mulx3475")
	} else {
		if !(rexsult.ToString() == "2.60648266168558079957349074609920E+151") {
			lang.SayString("mulx3475")
		}
	}
	rexsult, err = lang.RxFromString("83946.0157801953636255078996185540").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("492670373.235391665758701500314473"))
	if err != nil {
		lang.SayString("mulx3476")
	} else {
		if !(rexsult.ToString() == "41357714926052.9282985560380064649") {
			lang.SayString("mulx3476")
		}
	}
	rexsult, err = lang.RxFromString("7812758113817.99135851825223122508").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3037492.36716301443309571918002123E-157"))
	if err != nil {
		lang.SayString("mulx3477")
	} else {
		if !(rexsult.ToString() == "2.37311931372130583136091717093935E-138") {
			lang.SayString("mulx3477")
		}
	}
	rexsult, err = lang.RxFromString("489191747.148674326757767356626016").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("01136942.1182277580093027768490545"))
	if err != nil {
		lang.SayString("mulx3478")
	} else {
		if !(rexsult.ToString() == "556182701222751.588454129518830550") {
			lang.SayString("mulx3478")
		}
	}
	rexsult, err = lang.RxFromString("-599369540.373174482335865567937853E+289").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-38288383205.6749439588707955585209"))
	if err != nil {
		lang.SayString("mulx3479")
	} else {
		if !(rexsult.ToString() == "2.29488906436173641324091638963715E+308") {
			lang.SayString("mulx3479")
		}
	}
	rexsult, err = lang.RxFromString("-3376883870.85961692148022521960139").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-65247489449.7334589731171980408284"))
	if err != nil {
		lang.SayString("mulx3480")
	} else {
		if !(rexsult.ToString() == "220333194736887939420.719579906257") {
			lang.SayString("mulx3480")
		}
	}
	rexsult, err = lang.RxFromString("58.6776780370008364590621772421025").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("01.5925518866529044494309229975160"))
	if err != nil {
		lang.SayString("mulx3481")
	} else {
		if !(rexsult.ToString() == "93.4472468622373769590900258060029") {
			lang.SayString("mulx3481")
		}
	}
	rexsult, err = lang.RxFromString("4099616339.96249499552808575717579").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("290.795187361072489816791525139895"))
	if err != nil {
		lang.SayString("mulx3482")
	} else {
		if !(rexsult.ToString() == "1192148701687.90798437501397900174") {
			lang.SayString("mulx3482")
		}
	}
	rexsult, err = lang.RxFromString("85870777.2282833141709970713739108").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-2140392861153.69401346398478113715"))
	if err != nil {
		lang.SayString("mulx3483")
	} else {
		if !(rexsult.ToString() == "-183797198561136797328.508878254848") {
			lang.SayString("mulx3483")
		}
	}
	rexsult, err = lang.RxFromString("20900.9693761555165742010339929779").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-38.7546147649523793463260940289585"))
	if err != nil {
		lang.SayString("mulx3484")
	} else {
		if !(rexsult.ToString() == "-810009.016386974103738622793670566") {
			lang.SayString("mulx3484")
		}
	}
	rexsult, err = lang.RxFromString("448.827596155587910947511170319456").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("379130153.382794042652974596286062"))
	if err != nil {
		lang.SayString("mulx3485")
	} else {
		if !(rexsult.ToString() == "170164075372.898786469094460692097") {
			lang.SayString("mulx3485")
		}
	}
	rexsult, err = lang.RxFromString("98.4134807921002817357000140482039").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("3404725543.77032945444654351546779"))
	if err != nil {
		lang.SayString("mulx3486")
	} else {
		if !(rexsult.ToString() == "335070891904.214504811798212040413") {
			lang.SayString("mulx3486")
		}
	}
	rexsult, err = lang.RxFromString("545746433.649359734136476718176330E-787").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-5149957099709.12830072802043560650E-437"))
	if err != nil {
		lang.SayString("mulx3487")
	} else {
		if !(rexsult.ToString() == "-2.81057072061345688074304873033317E-1203") {
			lang.SayString("mulx3487")
		}
	}
	rexsult, err = lang.RxFromString("741304513547.273820525801608231737").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("0396.22823128272584928019323186355E-830"))
	if err != nil {
		lang.SayString("mulx3488")
	} else {
		if !(rexsult.ToString() == "2.93725776244737788947443361076095E-816") {
			lang.SayString("mulx3488")
		}
	}
	rexsult, err = lang.RxFromString("-706.145005094292315613907254240553").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("4739.82486195739758308735946332234"))
	if err != nil {
		lang.SayString("mulx3489")
	} else {
		if !(rexsult.ToString() == "-3347003.65129295988793454267973464") {
			lang.SayString("mulx3489")
		}
	}
	rexsult, err = lang.RxFromString("-769925786.823099083228795187975893").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-31201.9980469760239870067820594790"))
	if err != nil {
		lang.SayString("mulx3490")
	} else {
		if !(rexsult.ToString() == "24023222896770.8161787236737395477") {
			lang.SayString("mulx3490")
		}
	}
	rexsult, err = lang.RxFromString("84438610546049.7256507419289692857E+906").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("052604240766736461898844111790311"))
	if err != nil {
		lang.SayString("mulx3491")
	} else {
		if !(rexsult.ToString() == "4.44182899917309231779837668210610E+951") {
			lang.SayString("mulx3491")
		}
	}
	rexsult, err = lang.RxFromString("549760.911304725795164589619286514").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("165.160089615604924207754883953484"))
	if err != nil {
		lang.SayString("mulx3492")
	} else {
		if !(rexsult.ToString() == "90798561.3782451425861113694732484") {
			lang.SayString("mulx3492")
		}
	}
	rexsult, err = lang.RxFromString("3650514.18649737956855828939662794").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("08086721.4036886948248274834735629"))
	if err != nil {
		lang.SayString("mulx3493")
	} else {
		if !(rexsult.ToString() == "29520691206417.5831886752808745421") {
			lang.SayString("mulx3493")
		}
	}
	rexsult, err = lang.RxFromString("55067723881950.1346958179604099594").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-8.90481481687182931431054785192083"))
	if err != nil {
		lang.SayString("mulx3494")
	} else {
		if !(rexsult.ToString() == "-490367883555396.250365158593373279") {
			lang.SayString("mulx3494")
		}
	}
	rexsult, err = lang.RxFromString("868251123.413992653362860637541060E+019").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("5579665045.37858308541154858567656E+131"))
	if err != nil {
		lang.SayString("mulx3495")
	} else {
		if !(rexsult.ToString() == "4.84455044392374106106966779322483E+168") {
			lang.SayString("mulx3495")
		}
	}
	rexsult, err = lang.RxFromString("-646.464431574014407536004990059069").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-798.679560020414523841321724649594E-037"))
	if err != nil {
		lang.SayString("mulx3496")
	} else {
		if !(rexsult.ToString() == "5.16317927778381197995451363439626E-32") {
			lang.SayString("mulx3496")
		}
	}
	rexsult, err = lang.RxFromString("354.546679975219753598558273421556").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-7039.46386812239015144581761752927E-448"))
	if err != nil {
		lang.SayString("mulx3497")
	} else {
		if !(rexsult.ToString() == "-2.49581854324831161267369292071408E-442") {
			lang.SayString("mulx3497")
		}
	}
	rexsult, err = lang.RxFromString("91936087917435.5974889495278215874").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("-67080823344.8903392584327136082486E-757"))
	if err != nil {
		lang.SayString("mulx3498")
	} else {
		if !(rexsult.ToString() == "-6.16714847260980448099292763939423E-733") {
			lang.SayString("mulx3498")
		}
	}
	rexsult, err = lang.RxFromString("-07345.6422518528556136521417259811E-600").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("41188325.7041362608934957584583381E-763"))
	if err != nil {
		lang.SayString("mulx3499")
	} else {
		if !(rexsult.ToString() == "-3.02554705575380338274126867655676E-1352") {
			lang.SayString("mulx3499")
		}
	}
	rexsult, err = lang.RxFromString("-253280724.939458021588167965038184").OpMult(lang.RxSetWithDigit(33), lang.RxFromString("616988.426425908872398170896375634E+396"))
	if err != nil {
		lang.SayString("mulx3500")
	} else {
		if !(rexsult.ToString() == "-1.56271275924409657991913620522315E+410") {
			lang.SayString("mulx3500")
		}
	}
	rexsult, err = lang.RxFromString("905.67402").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-202896611.E-780472620"))
	if err != nil {
		lang.SayString("xmul001")
	} else {
		if !(rexsult.ToString() == "-1.83758189E-780472609") {
			lang.SayString("xmul001")
		}
	}
	rexsult, err = lang.RxFromString("3915134.7").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-597164907."))
	if err != nil {
		lang.SayString("xmul002")
	} else {
		if !(rexsult.ToString() == "-2.33798105E+15") {
			lang.SayString("xmul002")
		}
	}
	rexsult, err = lang.RxFromString("309759261").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("62663.487"))
	if err != nil {
		lang.SayString("xmul003")
	} else {
		if !(rexsult.ToString() == "1.94105954E+13") {
			lang.SayString("xmul003")
		}
	}
	rexsult, err = lang.RxFromString("3.93591888E-236595626").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7242375.00"))
	if err != nil {
		lang.SayString("xmul004")
	} else {
		if !(rexsult.ToString() == "2.85054005E-236595619") {
			lang.SayString("xmul004")
		}
	}
	rexsult, err = lang.RxFromString("323902.714").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-608669.607E-657060568"))
	if err != nil {
		lang.SayString("xmul005")
	} else {
		if !(rexsult.ToString() == "-1.97149738E-657060557") {
			lang.SayString("xmul005")
		}
	}
	rexsult, err = lang.RxFromString("5.11970092").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-8807.22036"))
	if err != nil {
		lang.SayString("xmul006")
	} else {
		if !(rexsult.ToString() == "-45090.3342") {
			lang.SayString("xmul006")
		}
	}
	rexsult, err = lang.RxFromString("-7.99874516").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4561.83758"))
	if err != nil {
		lang.SayString("xmul007")
	} else {
		if !(rexsult.ToString() == "-36488.9763") {
			lang.SayString("xmul007")
		}
	}
	rexsult, err = lang.RxFromString("297802878").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-927206.324"))
	if err != nil {
		lang.SayString("xmul008")
	} else {
		if !(rexsult.ToString() == "-2.76124712E+14") {
			lang.SayString("xmul008")
		}
	}
	rexsult, err = lang.RxFromString("-766.651824").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("31300.3619"))
	if err != nil {
		lang.SayString("xmul009")
	} else {
		if !(rexsult.ToString() == "-23996479.5") {
			lang.SayString("xmul009")
		}
	}
	rexsult, err = lang.RxFromString("-56746.8689E+934981942").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("471002521."))
	if err != nil {
		lang.SayString("xmul010")
	} else {
		if !(rexsult.ToString() == "-2.67279183E+934981955") {
			lang.SayString("xmul010")
		}
	}
	rexsult, err = lang.RxFromString("456417160").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-41346.1024"))
	if err != nil {
		lang.SayString("xmul011")
	} else {
		if !(rexsult.ToString() == "-1.88710706E+13") {
			lang.SayString("xmul011")
		}
	}
	rexsult, err = lang.RxFromString("102895.722").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2.62214826"))
	if err != nil {
		lang.SayString("xmul012")
	} else {
		if !(rexsult.ToString() == "-269807.838") {
			lang.SayString("xmul012")
		}
	}
	rexsult, err = lang.RxFromString("61.3033331E+157644141").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-567740.918E-893439456"))
	if err != nil {
		lang.SayString("xmul013")
	} else {
		if !(rexsult.ToString() == "-3.48044106E-735795308") {
			lang.SayString("xmul013")
		}
	}
	rexsult, err = lang.RxFromString("80223.3897").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("73921.0383E-467772675"))
	if err != nil {
		lang.SayString("xmul014")
	} else {
		if !(rexsult.ToString() == "5.93019626E-467772666") {
			lang.SayString("xmul014")
		}
	}
	rexsult, err = lang.RxFromString("-654645.954").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-9.12535752"))
	if err != nil {
		lang.SayString("xmul015")
	} else {
		if !(rexsult.ToString() == "5973878.38") {
			lang.SayString("xmul015")
		}
	}
	rexsult, err = lang.RxFromString("63.1917772E-706014634").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-7.56253257E-138579234"))
	if err != nil {
		lang.SayString("xmul016")
	} else {
		if !(rexsult.ToString() == "-4.77889873E-844593866") {
			lang.SayString("xmul016")
		}
	}
	rexsult, err = lang.RxFromString("-39674.7190").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2490607.78"))
	if err != nil {
		lang.SayString("xmul017")
	} else {
		if !(rexsult.ToString() == "-9.88141638E+10") {
			lang.SayString("xmul017")
		}
	}
	rexsult, err = lang.RxFromString("-3364.59737E-600363681").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("896487.451"))
	if err != nil {
		lang.SayString("xmul018")
	} else {
		if !(rexsult.ToString() == "-3.01631932E-600363672") {
			lang.SayString("xmul018")
		}
	}
	rexsult, err = lang.RxFromString("-64138.0578").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("31759011.3E+697488342"))
	if err != nil {
		lang.SayString("xmul019")
	} else {
		if !(rexsult.ToString() == "-2.03696130E+697488354") {
			lang.SayString("xmul019")
		}
	}
	rexsult, err = lang.RxFromString("61399.8527").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-64344484.5"))
	if err != nil {
		lang.SayString("xmul020")
	} else {
		if !(rexsult.ToString() == "-3.95074187E+12") {
			lang.SayString("xmul020")
		}
	}
	rexsult, err = lang.RxFromString("-722960.204").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-26154599.8"))
	if err != nil {
		lang.SayString("xmul021")
	} else {
		if !(rexsult.ToString() == "1.89087348E+13") {
			lang.SayString("xmul021")
		}
	}
	rexsult, err = lang.RxFromString("9.47109959E+230565093").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("73354723.2"))
	if err != nil {
		lang.SayString("xmul022")
	} else {
		if !(rexsult.ToString() == "6.94749889E+230565101") {
			lang.SayString("xmul022")
		}
	}
	rexsult, err = lang.RxFromString("43.7456245").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("547441956."))
	if err != nil {
		lang.SayString("xmul023")
	} else {
		if !(rexsult.ToString() == "2.39481902E+10") {
			lang.SayString("xmul023")
		}
	}
	rexsult, err = lang.RxFromString("-73150542E-242017390").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-8.15869954"))
	if err != nil {
		lang.SayString("xmul024")
	} else {
		if !(rexsult.ToString() == "5.96813293E-242017382") {
			lang.SayString("xmul024")
		}
	}
	rexsult, err = lang.RxFromString("2015.62109E+299897596").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-11788916.1"))
	if err != nil {
		lang.SayString("xmul025")
	} else {
		if !(rexsult.ToString() == "-2.37619879E+299897606") {
			lang.SayString("xmul025")
		}
	}
	rexsult, err = lang.RxFromString("29.498114").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-26486451"))
	if err != nil {
		lang.SayString("xmul026")
	} else {
		if !(rexsult.ToString() == "-781300351") {
			lang.SayString("xmul026")
		}
	}
	rexsult, err = lang.RxFromString("244375043.E+130840878").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-9.44522029"))
	if err != nil {
		lang.SayString("xmul027")
	} else {
		if !(rexsult.ToString() == "-2.30817611E+130840887") {
			lang.SayString("xmul027")
		}
	}
	rexsult, err = lang.RxFromString("-349388.759").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-196215.776"))
	if err != nil {
		lang.SayString("xmul028")
	} else {
		if !(rexsult.ToString() == "6.85555865E+10") {
			lang.SayString("xmul028")
		}
	}
	rexsult, err = lang.RxFromString("-70905112.4").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-91353968.8"))
	if err != nil {
		lang.SayString("xmul029")
	} else {
		if !(rexsult.ToString() == "6.47746343E+15") {
			lang.SayString("xmul029")
		}
	}
	rexsult, err = lang.RxFromString("-225094.28").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-88.7723542"))
	if err != nil {
		lang.SayString("xmul030")
	} else {
		if !(rexsult.ToString() == "19982149.2") {
			lang.SayString("xmul030")
		}
	}
	rexsult, err = lang.RxFromString("50.4442340").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("82.7952169E+880120759"))
	if err != nil {
		lang.SayString("xmul031")
	} else {
		if !(rexsult.ToString() == "4.17654130E+880120762") {
			lang.SayString("xmul031")
		}
	}
	rexsult, err = lang.RxFromString("-32311.9037").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("8.36379449"))
	if err != nil {
		lang.SayString("xmul032")
	} else {
		if !(rexsult.ToString() == "-270250.122") {
			lang.SayString("xmul032")
		}
	}
	rexsult, err = lang.RxFromString("615396156.E+549895291").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-29530247.4"))
	if err != nil {
		lang.SayString("xmul033")
	} else {
		if !(rexsult.ToString() == "-1.81728007E+549895307") {
			lang.SayString("xmul033")
		}
	}
	rexsult, err = lang.RxFromString("849.515993E-878446473").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1039.08778"))
	if err != nil {
		lang.SayString("xmul035")
	} else {
		if !(rexsult.ToString() == "-8.82721687E-878446468") {
			lang.SayString("xmul035")
		}
	}
	rexsult, err = lang.RxFromString("213361789").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-599.644851"))
	if err != nil {
		lang.SayString("xmul036")
	} else {
		if !(rexsult.ToString() == "-1.27941298E+11") {
			lang.SayString("xmul036")
		}
	}
	rexsult, err = lang.RxFromString("-795522555.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-298.037702"))
	if err != nil {
		lang.SayString("xmul037")
	} else {
		if !(rexsult.ToString() == "2.37095714E+11") {
			lang.SayString("xmul037")
		}
	}
	rexsult, err = lang.RxFromString("-501260651.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-8761893.0E-689281479"))
	if err != nil {
		lang.SayString("xmul038")
	} else {
		if !(rexsult.ToString() == "4.39199219E-689281464") {
			lang.SayString("xmul038")
		}
	}
	rexsult, err = lang.RxFromString("-1.70781105E-848889023").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("36504769.4"))
	if err != nil {
		lang.SayString("xmul039")
	} else {
		if !(rexsult.ToString() == "-6.23432486E-848889016") {
			lang.SayString("xmul039")
		}
	}
	rexsult, err = lang.RxFromString("-5290.54984E-490626676").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("842535254"))
	if err != nil {
		lang.SayString("xmul040")
	} else {
		if !(rexsult.ToString() == "-4.45747475E-490626664") {
			lang.SayString("xmul040")
		}
	}
	rexsult, err = lang.RxFromString("608.31825E+535268120").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-59609.0993"))
	if err != nil {
		lang.SayString("xmul041")
	} else {
		if !(rexsult.ToString() == "-3.62613030E+535268127") {
			lang.SayString("xmul041")
		}
	}
	rexsult, err = lang.RxFromString("-4629035.31").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-167.884398"))
	if err != nil {
		lang.SayString("xmul042")
	} else {
		if !(rexsult.ToString() == "777142806") {
			lang.SayString("xmul042")
		}
	}
	rexsult, err = lang.RxFromString("-66527378.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-706400268."))
	if err != nil {
		lang.SayString("xmul043")
	} else {
		if !(rexsult.ToString() == "4.69949576E+16") {
			lang.SayString("xmul043")
		}
	}
	rexsult, err = lang.RxFromString("-2510497.53").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("372882462."))
	if err != nil {
		lang.SayString("xmul044")
	} else {
		if !(rexsult.ToString() == "-9.36120500E+14") {
			lang.SayString("xmul044")
		}
	}
	rexsult, err = lang.RxFromString("136.255393E+53329961").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-53494.7201E+720058060"))
	if err != nil {
		lang.SayString("xmul045")
	} else {
		if !(rexsult.ToString() == "-7.28894411E+773388027") {
			lang.SayString("xmul045")
		}
	}
	rexsult, err = lang.RxFromString("-876673.100").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-6150.92266"))
	if err != nil {
		lang.SayString("xmul046")
	} else {
		if !(rexsult.ToString() == "5.39234844E+9") {
			lang.SayString("xmul046")
		}
	}
	rexsult, err = lang.RxFromString("-2.45151797E+911306117").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("27235771"))
	if err != nil {
		lang.SayString("xmul047")
	} else {
		if !(rexsult.ToString() == "-6.67689820E+911306124") {
			lang.SayString("xmul047")
		}
	}
	rexsult, err = lang.RxFromString("-9.15117551").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4.95100733E-314511326"))
	if err != nil {
		lang.SayString("xmul048")
	} else {
		if !(rexsult.ToString() == "4.53075370E-314511325") {
			lang.SayString("xmul048")
		}
	}
	rexsult, err = lang.RxFromString("3.61890453E-985606128").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("30664416."))
	if err != nil {
		lang.SayString("xmul049")
	} else {
		if !(rexsult.ToString() == "1.10971594E-985606120") {
			lang.SayString("xmul049")
		}
	}
	rexsult, err = lang.RxFromString("-257674602E+216723382").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-70820959.4"))
	if err != nil {
		lang.SayString("xmul050")
	} else {
		if !(rexsult.ToString() == "1.82487625E+216723398") {
			lang.SayString("xmul050")
		}
	}
	rexsult, err = lang.RxFromString("218699.206").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("556944241."))
	if err != nil {
		lang.SayString("xmul051")
	} else {
		if !(rexsult.ToString() == "1.21803263E+14") {
			lang.SayString("xmul051")
		}
	}
	rexsult, err = lang.RxFromString("106211716.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-3456793.74"))
	if err != nil {
		lang.SayString("xmul052")
	} else {
		if !(rexsult.ToString() == "-3.67151995E+14") {
			lang.SayString("xmul052")
		}
	}
	rexsult, err = lang.RxFromString("1.25018078").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("399856.763E-726816740"))
	if err != nil {
		lang.SayString("xmul053")
	} else {
		if !(rexsult.ToString() == "4.99893240E-726816735") {
			lang.SayString("xmul053")
		}
	}
	rexsult, err = lang.RxFromString("364.99811").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-46222.0505"))
	if err != nil {
		lang.SayString("xmul054")
	} else {
		if !(rexsult.ToString() == "-16870961.1") {
			lang.SayString("xmul054")
		}
	}
	rexsult, err = lang.RxFromString("-392217576.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-958364096"))
	if err != nil {
		lang.SayString("xmul055")
	} else {
		if !(rexsult.ToString() == "3.75887243E+17") {
			lang.SayString("xmul055")
		}
	}
	rexsult, err = lang.RxFromString("169601285").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("714526.639"))
	if err != nil {
		lang.SayString("xmul056")
	} else {
		if !(rexsult.ToString() == "1.21184636E+14") {
			lang.SayString("xmul056")
		}
	}
	rexsult, err = lang.RxFromString("-746.293386").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("927749.647"))
	if err != nil {
		lang.SayString("xmul059")
	} else {
		if !(rexsult.ToString() == "-692373425") {
			lang.SayString("xmul059")
		}
	}
	rexsult, err = lang.RxFromString("888946471E+241331592").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-235739.595"))
	if err != nil {
		lang.SayString("xmul060")
	} else {
		if !(rexsult.ToString() == "-2.09559881E+241331606") {
			lang.SayString("xmul060")
		}
	}
	rexsult, err = lang.RxFromString("6.64377249").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("79161.1070E+619453776"))
	if err != nil {
		lang.SayString("xmul061")
	} else {
		if !(rexsult.ToString() == "5.25928385E+619453781") {
			lang.SayString("xmul061")
		}
	}
	rexsult, err = lang.RxFromString("3146.66571E-313373366").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("88.5282010"))
	if err != nil {
		lang.SayString("xmul062")
	} else {
		if !(rexsult.ToString() == "2.78568654E-313373361") {
			lang.SayString("xmul062")
		}
	}
	rexsult, err = lang.RxFromString("6.44693097").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-87195.8711"))
	if err != nil {
		lang.SayString("xmul063")
	} else {
		if !(rexsult.ToString() == "-562145.762") {
			lang.SayString("xmul063")
		}
	}
	rexsult, err = lang.RxFromString("-2113132.56E+577957840").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("981125821"))
	if err != nil {
		lang.SayString("xmul064")
	} else {
		if !(rexsult.ToString() == "-2.07324892E+577957855") {
			lang.SayString("xmul064")
		}
	}
	rexsult, err = lang.RxFromString("-7701.42814").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("72667.5181"))
	if err != nil {
		lang.SayString("xmul065")
	} else {
		if !(rexsult.ToString() == "-559643669") {
			lang.SayString("xmul065")
		}
	}
	rexsult, err = lang.RxFromString("-851.754789").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-582659.149"))
	if err != nil {
		lang.SayString("xmul066")
	} else {
		if !(rexsult.ToString() == "496282721") {
			lang.SayString("xmul066")
		}
	}
	rexsult, err = lang.RxFromString("-5.01992943").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7852.16531"))
	if err != nil {
		lang.SayString("xmul067")
	} else {
		if !(rexsult.ToString() == "-39417.3157") {
			lang.SayString("xmul067")
		}
	}
	rexsult, err = lang.RxFromString("-12393257.2").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("76803689E+949125770"))
	if err != nil {
		lang.SayString("xmul068")
	} else {
		if !(rexsult.ToString() == "-9.51847872E+949125784") {
			lang.SayString("xmul068")
		}
	}
	rexsult, err = lang.RxFromString("-754771634.E+716555026").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-292336.311"))
	if err != nil {
		lang.SayString("xmul069")
	} else {
		if !(rexsult.ToString() == "2.20647155E+716555040") {
			lang.SayString("xmul069")
		}
	}
	rexsult, err = lang.RxFromString("-915006.171E+614548652").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-314086965."))
	if err != nil {
		lang.SayString("xmul070")
	} else {
		if !(rexsult.ToString() == "2.87391511E+614548666") {
			lang.SayString("xmul070")
		}
	}
	rexsult, err = lang.RxFromString("-296590035").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-481734529"))
	if err != nil {
		lang.SayString("xmul071")
	} else {
		if !(rexsult.ToString() == "1.42877661E+17") {
			lang.SayString("xmul071")
		}
	}
	rexsult, err = lang.RxFromString("8.27822605").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9241557.19"))
	if err != nil {
		lang.SayString("xmul072")
	} else {
		if !(rexsult.ToString() == "76503699.5") {
			lang.SayString("xmul072")
		}
	}
	rexsult, err = lang.RxFromString("-1.43581098").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7286313.54"))
	if err != nil {
		lang.SayString("xmul073")
	} else {
		if !(rexsult.ToString() == "-10461769.0") {
			lang.SayString("xmul073")
		}
	}
	rexsult, err = lang.RxFromString("-699036193.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("759263.509E+533543625"))
	if err != nil {
		lang.SayString("xmul074")
	} else {
		if !(rexsult.ToString() == "-5.30752673E+533543639") {
			lang.SayString("xmul074")
		}
	}
	rexsult, err = lang.RxFromString("-83.7273615E-305281957").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-287779593.E+458777774"))
	if err != nil {
		lang.SayString("xmul075")
	} else {
		if !(rexsult.ToString() == "2.40950260E+153495827") {
			lang.SayString("xmul075")
		}
	}
	rexsult, err = lang.RxFromString("8.48503224").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("6522.03316"))
	if err != nil {
		lang.SayString("xmul076")
	} else {
		if !(rexsult.ToString() == "55339.6616") {
			lang.SayString("xmul076")
		}
	}
	rexsult, err = lang.RxFromString("527916091").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-809.054070"))
	if err != nil {
		lang.SayString("xmul077")
	} else {
		if !(rexsult.ToString() == "-4.27112662E+11") {
			lang.SayString("xmul077")
		}
	}
	rexsult, err = lang.RxFromString("3857058.60").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("5792997.58E+881077409"))
	if err != nil {
		lang.SayString("xmul078")
	} else {
		if !(rexsult.ToString() == "2.23439311E+881077422") {
			lang.SayString("xmul078")
		}
	}
	rexsult, err = lang.RxFromString("-66587363.E+556538173").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-551902402E+357309146"))
	if err != nil {
		lang.SayString("xmul079")
	} else {
		if !(rexsult.ToString() == "3.67497256E+913847335") {
			lang.SayString("xmul079")
		}
	}
	rexsult, err = lang.RxFromString("-580.502955").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("38125521.7"))
	if err != nil {
		lang.SayString("xmul080")
	} else {
		if !(rexsult.ToString() == "-2.21319780E+10") {
			lang.SayString("xmul080")
		}
	}
	rexsult, err = lang.RxFromString("-9627363.00").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-80616885E-749891394"))
	if err != nil {
		lang.SayString("xmul081")
	} else {
		if !(rexsult.ToString() == "7.76128016E-749891380") {
			lang.SayString("xmul081")
		}
	}
	rexsult, err = lang.RxFromString("-526.594855E+803110107").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-64.5451639"))
	if err != nil {
		lang.SayString("xmul082")
	} else {
		if !(rexsult.ToString() == "3.39891512E+803110111") {
			lang.SayString("xmul082")
		}
	}
	rexsult, err = lang.RxFromString("-8378.55499").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("760.131257"))
	if err != nil {
		lang.SayString("xmul083")
	} else {
		if !(rexsult.ToString() == "-6368801.54") {
			lang.SayString("xmul083")
		}
	}
	rexsult, err = lang.RxFromString("-717.697718").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("984304413"))
	if err != nil {
		lang.SayString("xmul084")
	} else {
		if !(rexsult.ToString() == "-7.06433031E+11") {
			lang.SayString("xmul084")
		}
	}
	rexsult, err = lang.RxFromString("-76762243.4E-741100094").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-273.706674"))
	if err != nil {
		lang.SayString("xmul085")
	} else {
		if !(rexsult.ToString() == "2.10103383E-741100084") {
			lang.SayString("xmul085")
		}
	}
	rexsult, err = lang.RxFromString("-359866845.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4.57434117"))
	if err != nil {
		lang.SayString("xmul087")
	} else {
		if !(rexsult.ToString() == "1.64615372E+9") {
			lang.SayString("xmul087")
		}
	}
	rexsult, err = lang.RxFromString("779934536.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-76562645.7"))
	if err != nil {
		lang.SayString("xmul088")
	} else {
		if !(rexsult.ToString() == "-5.97138515E+16") {
			lang.SayString("xmul088")
		}
	}
	rexsult, err = lang.RxFromString("-4820.95451").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3516234.99E+303303176"))
	if err != nil {
		lang.SayString("xmul089")
	} else {
		if !(rexsult.ToString() == "-1.69516089E+303303186") {
			lang.SayString("xmul089")
		}
	}
	rexsult, err = lang.RxFromString("69355976.9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-9.57838562E+758804984"))
	if err != nil {
		lang.SayString("xmul090")
	} else {
		if !(rexsult.ToString() == "-6.64318292E+758804992") {
			lang.SayString("xmul090")
		}
	}
	rexsult, err = lang.RxFromString("-12672093.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("8569.78255E-382866025"))
	if err != nil {
		lang.SayString("xmul091")
	} else {
		if !(rexsult.ToString() == "-1.08597082E-382866014") {
			lang.SayString("xmul091")
		}
	}
	rexsult, err = lang.RxFromString("-5910750.2").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("66150383E-662459241"))
	if err != nil {
		lang.SayString("xmul092")
	} else {
		if !(rexsult.ToString() == "-3.90998390E-662459227") {
			lang.SayString("xmul092")
		}
	}
	rexsult, err = lang.RxFromString("-532577268.E-163806629").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-240650398E-650110558"))
	if err != nil {
		lang.SayString("xmul093")
	} else {
		if !(rexsult.ToString() == "1.28164932E-813917170") {
			lang.SayString("xmul093")
		}
	}
	rexsult, err = lang.RxFromString("-294.994352E+346452027").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-6061853.0"))
	if err != nil {
		lang.SayString("xmul095")
	} else {
		if !(rexsult.ToString() == "1.78821240E+346452036") {
			lang.SayString("xmul095")
		}
	}
	rexsult, err = lang.RxFromString("329579114").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("146780548."))
	if err != nil {
		lang.SayString("xmul096")
	} else {
		if !(rexsult.ToString() == "4.83758030E+16") {
			lang.SayString("xmul096")
		}
	}
	rexsult, err = lang.RxFromString("-789904.686E-217225000").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1991.07181E-84080059"))
	if err != nil {
		lang.SayString("xmul097")
	} else {
		if !(rexsult.ToString() == "1.57275695E-301305050") {
			lang.SayString("xmul097")
		}
	}
	rexsult, err = lang.RxFromString("59893.3544").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-408595868"))
	if err != nil {
		lang.SayString("xmul098")
	} else {
		if !(rexsult.ToString() == "-2.44721771E+13") {
			lang.SayString("xmul098")
		}
	}
	rexsult, err = lang.RxFromString("129.878613").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-54652.7288E-963564940"))
	if err != nil {
		lang.SayString("xmul099")
	} else {
		if !(rexsult.ToString() == "-7.09822061E-963564934") {
			lang.SayString("xmul099")
		}
	}
	rexsult, err = lang.RxFromString("9866.99208").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("708756501."))
	if err != nil {
		lang.SayString("xmul100")
	} else {
		if !(rexsult.ToString() == "6.99329478E+12") {
			lang.SayString("xmul100")
		}
	}
	rexsult, err = lang.RxFromString("-78810.6297").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-399884.68"))
	if err != nil {
		lang.SayString("xmul101")
	} else {
		if !(rexsult.ToString() == "3.15151634E+10") {
			lang.SayString("xmul101")
		}
	}
	rexsult, err = lang.RxFromString("409189761").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-771.471460"))
	if err != nil {
		lang.SayString("xmul102")
	} else {
		if !(rexsult.ToString() == "-3.15678222E+11") {
			lang.SayString("xmul102")
		}
	}
	rexsult, err = lang.RxFromString("-1.68748838").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("460.46924"))
	if err != nil {
		lang.SayString("xmul103")
	} else {
		if !(rexsult.ToString() == "-777.036492") {
			lang.SayString("xmul103")
		}
	}
	rexsult, err = lang.RxFromString("553527296.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-7924.40185"))
	if err != nil {
		lang.SayString("xmul104")
	} else {
		if !(rexsult.ToString() == "-4.38637273E+12") {
			lang.SayString("xmul104")
		}
	}
	rexsult, err = lang.RxFromString("-38.7465207").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("64936.2942"))
	if err != nil {
		lang.SayString("xmul105")
	} else {
		if !(rexsult.ToString() == "-2516055.47") {
			lang.SayString("xmul105")
		}
	}
	rexsult, err = lang.RxFromString("-201075.248").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("845.663928"))
	if err != nil {
		lang.SayString("xmul106")
	} else {
		if !(rexsult.ToString() == "-170042084") {
			lang.SayString("xmul106")
		}
	}
	rexsult, err = lang.RxFromString("91048.4559").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("75953609.3"))
	if err != nil {
		lang.SayString("xmul107")
	} else {
		if !(rexsult.ToString() == "6.91545885E+12") {
			lang.SayString("xmul107")
		}
	}
	rexsult, err = lang.RxFromString("6898273.86E-252097460").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("15.3456196"))
	if err != nil {
		lang.SayString("xmul108")
	} else {
		if !(rexsult.ToString() == "1.05858287E-252097452") {
			lang.SayString("xmul108")
		}
	}
	rexsult, err = lang.RxFromString("88.4370343").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-980709105E-869899289"))
	if err != nil {
		lang.SayString("xmul109")
	} else {
		if !(rexsult.ToString() == "-8.67310048E-869899279") {
			lang.SayString("xmul109")
		}
	}
	rexsult, err = lang.RxFromString("-17643.39").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2.0352568E+304871331"))
	if err != nil {
		lang.SayString("xmul110")
	} else {
		if !(rexsult.ToString() == "-3.59088295E+304871335") {
			lang.SayString("xmul110")
		}
	}
	rexsult, err = lang.RxFromString("4589785.16").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7459.04237"))
	if err != nil {
		lang.SayString("xmul111")
	} else {
		if !(rexsult.ToString() == "3.42354020E+10") {
			lang.SayString("xmul111")
		}
	}
	rexsult, err = lang.RxFromString("982.217817").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7224909.4E-45243816"))
	if err != nil {
		lang.SayString("xmul113")
	} else {
		if !(rexsult.ToString() == "7.09643474E-45243807") {
			lang.SayString("xmul113")
		}
	}
	rexsult, err = lang.RxFromString("503712056.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-57490703.5E+924890183"))
	if err != nil {
		lang.SayString("xmul114")
	} else {
		if !(rexsult.ToString() == "-2.89587605E+924890199") {
			lang.SayString("xmul114")
		}
	}
	rexsult, err = lang.RxFromString("883.849223").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("249259171"))
	if err != nil {
		lang.SayString("xmul115")
	} else {
		if !(rexsult.ToString() == "2.20307525E+11") {
			lang.SayString("xmul115")
		}
	}
	rexsult, err = lang.RxFromString("-83658638.6E+728551928").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2952478.42"))
	if err != nil {
		lang.SayString("xmul117")
	} else {
		if !(rexsult.ToString() == "-2.47000325E+728551942") {
			lang.SayString("xmul117")
		}
	}
	rexsult, err = lang.RxFromString("-6291780.97").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("269967.394E-22000817"))
	if err != nil {
		lang.SayString("xmul118")
	} else {
		if !(rexsult.ToString() == "-1.69857571E-22000805") {
			lang.SayString("xmul118")
		}
	}
	rexsult, err = lang.RxFromString("978571348.E+222382547").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("6006.19370"))
	if err != nil {
		lang.SayString("xmul119")
	} else {
		if !(rexsult.ToString() == "5.87748907E+222382559") {
			lang.SayString("xmul119")
		}
	}
	rexsult, err = lang.RxFromString("14239029.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-36527.2221"))
	if err != nil {
		lang.SayString("xmul120")
	} else {
		if !(rexsult.ToString() == "-5.20112175E+11") {
			lang.SayString("xmul120")
		}
	}
	rexsult, err = lang.RxFromString("72333.2654E-544425548").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-690.664836E+662155120"))
	if err != nil {
		lang.SayString("xmul121")
	} else {
		if !(rexsult.ToString() == "-4.99580429E+117729579") {
			lang.SayString("xmul121")
		}
	}
	rexsult, err = lang.RxFromString("-37721.1567E-115787341").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-828949864E-76251747"))
	if err != nil {
		lang.SayString("xmul122")
	} else {
		if !(rexsult.ToString() == "3.12689477E-192039075") {
			lang.SayString("xmul122")
		}
	}
	rexsult, err = lang.RxFromString("-2078852.83E-647080089").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-119779858.E+734665461"))
	if err != nil {
		lang.SayString("xmul123")
	} else {
		if !(rexsult.ToString() == "2.49004697E+87585386") {
			lang.SayString("xmul123")
		}
	}
	rexsult, err = lang.RxFromString("-79145.3625").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-7718.57307"))
	if err != nil {
		lang.SayString("xmul124")
	} else {
		if !(rexsult.ToString() == "610889264") {
			lang.SayString("xmul124")
		}
	}
	rexsult, err = lang.RxFromString("2103890.49E+959247237").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("20024.3017"))
	if err != nil {
		lang.SayString("xmul125")
	} else {
		if !(rexsult.ToString() == "4.21289379E+959247247") {
			lang.SayString("xmul125")
		}
	}
	rexsult, err = lang.RxFromString("911249557").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("79810804.9"))
	if err != nil {
		lang.SayString("xmul126")
	} else {
		if !(rexsult.ToString() == "7.27275606E+16") {
			lang.SayString("xmul126")
		}
	}
	rexsult, err = lang.RxFromString("341134.994").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3.37486292"))
	if err != nil {
		lang.SayString("xmul127")
	} else {
		if !(rexsult.ToString() == "1151283.84") {
			lang.SayString("xmul127")
		}
	}
	rexsult, err = lang.RxFromString("244.23634").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("512706190E-341459836"))
	if err != nil {
		lang.SayString("xmul128")
	} else {
		if !(rexsult.ToString() == "1.25221483E-341459825") {
			lang.SayString("xmul128")
		}
	}
	rexsult, err = lang.RxFromString("-9.22783849E+171585954").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-99.0946800"))
	if err != nil {
		lang.SayString("xmul129")
	} else {
		if !(rexsult.ToString() == "9.14429702E+171585956") {
			lang.SayString("xmul129")
		}
	}
	rexsult, err = lang.RxFromString("699631.893").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-226.423958"))
	if err != nil {
		lang.SayString("xmul130")
	} else {
		if !(rexsult.ToString() == "-158413422") {
			lang.SayString("xmul130")
		}
	}
	rexsult, err = lang.RxFromString("-249350139.E-571793673").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("775732428."))
	if err != nil {
		lang.SayString("xmul131")
	} else {
		if !(rexsult.ToString() == "-1.93428989E-571793656") {
			lang.SayString("xmul131")
		}
	}
	rexsult, err = lang.RxFromString("5.11629020").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-480.53194"))
	if err != nil {
		lang.SayString("xmul132")
	} else {
		if !(rexsult.ToString() == "-2458.54086") {
			lang.SayString("xmul132")
		}
	}
	rexsult, err = lang.RxFromString("-8.23352673E-446723147").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-530710.866"))
	if err != nil {
		lang.SayString("xmul133")
	} else {
		if !(rexsult.ToString() == "4.36962210E-446723141") {
			lang.SayString("xmul133")
		}
	}
	rexsult, err = lang.RxFromString("7.0598608").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-95908.35"))
	if err != nil {
		lang.SayString("xmul134")
	} else {
		if !(rexsult.ToString() == "-677099.601") {
			lang.SayString("xmul134")
		}
	}
	rexsult, err = lang.RxFromString("-7.91189845E+207202706").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1532.71847E+509944335"))
	if err != nil {
		lang.SayString("xmul135")
	} else {
		if !(rexsult.ToString() == "-1.21267129E+717147045") {
			lang.SayString("xmul135")
		}
	}
	rexsult, err = lang.RxFromString("208839370.E-215147432").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-75.9420559"))
	if err != nil {
		lang.SayString("xmul136")
	} else {
		if !(rexsult.ToString() == "-1.58596911E-215147422") {
			lang.SayString("xmul136")
		}
	}
	rexsult, err = lang.RxFromString("427.754244E-353328369").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4705.0796"))
	if err != nil {
		lang.SayString("xmul137")
	} else {
		if !(rexsult.ToString() == "2.01261777E-353328363") {
			lang.SayString("xmul137")
		}
	}
	rexsult, err = lang.RxFromString("44911.089").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-95.1733605E-313081848"))
	if err != nil {
		lang.SayString("xmul138")
	} else {
		if !(rexsult.ToString() == "-4.27433926E-313081842") {
			lang.SayString("xmul138")
		}
	}
	rexsult, err = lang.RxFromString("452371821.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4109709.19"))
	if err != nil {
		lang.SayString("xmul139")
	} else {
		if !(rexsult.ToString() == "-1.85911663E+15") {
			lang.SayString("xmul139")
		}
	}
	rexsult, err = lang.RxFromString("94007.4392").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-9467725.5E+681898234"))
	if err != nil {
		lang.SayString("xmul140")
	} else {
		if !(rexsult.ToString() == "-8.90036629E+681898245") {
			lang.SayString("xmul140")
		}
	}
	rexsult, err = lang.RxFromString("99147554.0E-751410586").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("38313.6423"))
	if err != nil {
		lang.SayString("xmul141")
	} else {
		if !(rexsult.ToString() == "3.79870392E-751410574") {
			lang.SayString("xmul141")
		}
	}
	rexsult, err = lang.RxFromString("-7919.30254").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-669.607854"))
	if err != nil {
		lang.SayString("xmul142")
	} else {
		if !(rexsult.ToString() == "5302827.18") {
			lang.SayString("xmul142")
		}
	}
	rexsult, err = lang.RxFromString("461.58280E+136110821").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("710666052.E-383754231"))
	if err != nil {
		lang.SayString("xmul143")
	} else {
		if !(rexsult.ToString() == "3.28031226E-247643399") {
			lang.SayString("xmul143")
		}
	}
	rexsult, err = lang.RxFromString("3455755.47E-112465506").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("771.674306"))
	if err != nil {
		lang.SayString("xmul144")
	} else {
		if !(rexsult.ToString() == "2.66671770E-112465497") {
			lang.SayString("xmul144")
		}
	}
	rexsult, err = lang.RxFromString("76482.352").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("8237806.8"))
	if err != nil {
		lang.SayString("xmul146")
	} else {
		if !(rexsult.ToString() == "6.30046839E+11") {
			lang.SayString("xmul146")
		}
	}
	rexsult, err = lang.RxFromString("1.21505164E-565556504").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9.26146573"))
	if err != nil {
		lang.SayString("xmul147")
	} else {
		if !(rexsult.ToString() == "1.12531591E-565556503") {
			lang.SayString("xmul147")
		}
	}
	rexsult, err = lang.RxFromString("-8303060.25E-169894883").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("901561.985"))
	if err != nil {
		lang.SayString("xmul148")
	} else {
		if !(rexsult.ToString() == "-7.48572348E-169894871") {
			lang.SayString("xmul148")
		}
	}
	rexsult, err = lang.RxFromString("-592464.92").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("71445510.7"))
	if err != nil {
		lang.SayString("xmul149")
	} else {
		if !(rexsult.ToString() == "-4.23289588E+13") {
			lang.SayString("xmul149")
		}
	}
	rexsult, err = lang.RxFromString("-73774.4165").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-39.8243027"))
	if err != nil {
		lang.SayString("xmul150")
	} else {
		if !(rexsult.ToString() == "2938014.69") {
			lang.SayString("xmul150")
		}
	}
	rexsult, err = lang.RxFromString("-524724715.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-55763.7937"))
	if err != nil {
		lang.SayString("xmul151")
	} else {
		if !(rexsult.ToString() == "2.92606408E+13") {
			lang.SayString("xmul151")
		}
	}
	rexsult, err = lang.RxFromString("7.53800427").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("784873768E-9981146"))
	if err != nil {
		lang.SayString("xmul152")
	} else {
		if !(rexsult.ToString() == "5.91638181E-9981137") {
			lang.SayString("xmul152")
		}
	}
	rexsult, err = lang.RxFromString("37.6027452").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7.22454233"))
	if err != nil {
		lang.SayString("xmul153")
	} else {
		if !(rexsult.ToString() == "271.662624") {
			lang.SayString("xmul153")
		}
	}
	rexsult, err = lang.RxFromString("2447660.39").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-36981.4253"))
	if err != nil {
		lang.SayString("xmul154")
	} else {
		if !(rexsult.ToString() == "-9.05179699E+10") {
			lang.SayString("xmul154")
		}
	}
	rexsult, err = lang.RxFromString("2160.36419").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1418.33574E+656265382"))
	if err != nil {
		lang.SayString("xmul155")
	} else {
		if !(rexsult.ToString() == "3.06412174E+656265388") {
			lang.SayString("xmul155")
		}
	}
	rexsult, err = lang.RxFromString("8926.44939").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("54.9430027"))
	if err != nil {
		lang.SayString("xmul156")
	} else {
		if !(rexsult.ToString() == "490445.933") {
			lang.SayString("xmul156")
		}
	}
	rexsult, err = lang.RxFromString("861588029").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-41657398E+77955925"))
	if err != nil {
		lang.SayString("xmul157")
	} else {
		if !(rexsult.ToString() == "-3.58915154E+77955941") {
			lang.SayString("xmul157")
		}
	}
	rexsult, err = lang.RxFromString("-34.5253062").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("52.6722019"))
	if err != nil {
		lang.SayString("xmul158")
	} else {
		if !(rexsult.ToString() == "-1818.52390") {
			lang.SayString("xmul158")
		}
	}
	rexsult, err = lang.RxFromString("-18861647.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("99794586.7"))
	if err != nil {
		lang.SayString("xmul159")
	} else {
		if !(rexsult.ToString() == "-1.88229027E+15") {
			lang.SayString("xmul159")
		}
	}
	rexsult, err = lang.RxFromString("322192.407").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("461.67044"))
	if err != nil {
		lang.SayString("xmul160")
	} else {
		if !(rexsult.ToString() == "148746710") {
			lang.SayString("xmul160")
		}
	}
	rexsult, err = lang.RxFromString("-896298518E+61412314").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("78873.8049"))
	if err != nil {
		lang.SayString("xmul161")
	} else {
		if !(rexsult.ToString() == "-7.06944744E+61412327") {
			lang.SayString("xmul161")
		}
	}
	rexsult, err = lang.RxFromString("293.773732").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("479899052E+789950177"))
	if err != nil {
		lang.SayString("xmul162")
	} else {
		if !(rexsult.ToString() == "1.40981735E+789950188") {
			lang.SayString("xmul162")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("51897955.3"))
	if err != nil {
		lang.SayString("xmul163")
	} else {
		if !(rexsult.ToString() == "-5.37244322E+15") {
			lang.SayString("xmul163")
		}
	}
	rexsult, err = lang.RxFromString("37380.7802").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-277719788."))
	if err != nil {
		lang.SayString("xmul164")
	} else {
		if !(rexsult.ToString() == "-1.03813824E+13") {
			lang.SayString("xmul164")
		}
	}
	rexsult, err = lang.RxFromString("320133844.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-977517477"))
	if err != nil {
		lang.SayString("xmul165")
	} else {
		if !(rexsult.ToString() == "-3.12936427E+17") {
			lang.SayString("xmul165")
		}
	}
	rexsult, err = lang.RxFromString("-5409.00482").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2.16149386"))
	if err != nil {
		lang.SayString("xmul167")
	} else {
		if !(rexsult.ToString() == "11691.5307") {
			lang.SayString("xmul167")
		}
	}
	rexsult, err = lang.RxFromString("-957960.367").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("322.858170"))
	if err != nil {
		lang.SayString("xmul168")
	} else {
		if !(rexsult.ToString() == "-309285331") {
			lang.SayString("xmul168")
		}
	}
	rexsult, err = lang.RxFromString("840258203").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("58363.980E-906584723"))
	if err != nil {
		lang.SayString("xmul170")
	} else {
		if !(rexsult.ToString() == "4.90408130E-906584710") {
			lang.SayString("xmul170")
		}
	}
	rexsult, err = lang.RxFromString("-205842096.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-191342.721"))
	if err != nil {
		lang.SayString("xmul171")
	} else {
		if !(rexsult.ToString() == "3.93863867E+13") {
			lang.SayString("xmul171")
		}
	}
	rexsult, err = lang.RxFromString("42501124.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("884.938498E+123341480"))
	if err != nil {
		lang.SayString("xmul172")
	} else {
		if !(rexsult.ToString() == "3.76108808E+123341490") {
			lang.SayString("xmul172")
		}
	}
	rexsult, err = lang.RxFromString("-57809452.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-620380746"))
	if err != nil {
		lang.SayString("xmul173")
	} else {
		if !(rexsult.ToString() == "3.58638710E+16") {
			lang.SayString("xmul173")
		}
	}
	rexsult, err = lang.RxFromString("-8022370.31").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9858581.6"))
	if err != nil {
		lang.SayString("xmul174")
	} else {
		if !(rexsult.ToString() == "-7.90891923E+13") {
			lang.SayString("xmul174")
		}
	}
	rexsult, err = lang.RxFromString("2.49065060E+592139141").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-5432.72014E-419965357"))
	if err != nil {
		lang.SayString("xmul175")
	} else {
		if !(rexsult.ToString() == "-1.35310077E+172173788") {
			lang.SayString("xmul175")
		}
	}
	rexsult, err = lang.RxFromString("-697273715E-242824870").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-3.81757506"))
	if err != nil {
		lang.SayString("xmul176")
	} else {
		if !(rexsult.ToString() == "2.66189474E-242824861") {
			lang.SayString("xmul176")
		}
	}
	rexsult, err = lang.RxFromString("-7.42204403E-315716280").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-8156111.67E+283261636"))
	if err != nil {
		lang.SayString("xmul177")
	} else {
		if !(rexsult.ToString() == "6.05350199E-32454637") {
			lang.SayString("xmul177")
		}
	}
	rexsult, err = lang.RxFromString("738063892").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("89900467.8"))
	if err != nil {
		lang.SayString("xmul178")
	} else {
		if !(rexsult.ToString() == "6.63522892E+16") {
			lang.SayString("xmul178")
		}
	}
	rexsult, err = lang.RxFromString("-630309366").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-884783.338E-21595410"))
	if err != nil {
		lang.SayString("xmul179")
	} else {
		if !(rexsult.ToString() == "5.57687225E-21595396") {
			lang.SayString("xmul179")
		}
	}
	rexsult, err = lang.RxFromString("613.207774").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-3007.78608"))
	if err != nil {
		lang.SayString("xmul180")
	} else {
		if !(rexsult.ToString() == "-1844397.81") {
			lang.SayString("xmul180")
		}
	}
	rexsult, err = lang.RxFromString("-93006222.3").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-3.08964619"))
	if err != nil {
		lang.SayString("xmul181")
	} else {
		if !(rexsult.ToString() == "287356320") {
			lang.SayString("xmul181")
		}
	}
	rexsult, err = lang.RxFromString("-18116.0621").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("34096.306E-270347092"))
	if err != nil {
		lang.SayString("xmul182")
	} else {
		if !(rexsult.ToString() == "-6.17690797E-270347084") {
			lang.SayString("xmul182")
		}
	}
	rexsult, err = lang.RxFromString("19272386.9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-410442379."))
	if err != nil {
		lang.SayString("xmul183")
	} else {
		if !(rexsult.ToString() == "-7.91020433E+15") {
			lang.SayString("xmul183")
		}
	}
	rexsult, err = lang.RxFromString("4180.30821").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1.6439543E-624759104"))
	if err != nil {
		lang.SayString("xmul184")
	} else {
		if !(rexsult.ToString() == "-6.87223566E-624759101") {
			lang.SayString("xmul184")
		}
	}
	rexsult, err = lang.RxFromString("571.536725").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("389.899220"))
	if err != nil {
		lang.SayString("xmul185")
	} else {
		if !(rexsult.ToString() == "222841.723") {
			lang.SayString("xmul185")
		}
	}
	rexsult, err = lang.RxFromString("-622007306.E+159924886").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-126.971745"))
	if err != nil {
		lang.SayString("xmul186")
	} else {
		if !(rexsult.ToString() == "7.89773530E+159924896") {
			lang.SayString("xmul186")
		}
	}
	rexsult, err = lang.RxFromString("92427442.4").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("674334898."))
	if err != nil {
		lang.SayString("xmul188")
	} else {
		if !(rexsult.ToString() == "6.23270499E+16") {
			lang.SayString("xmul188")
		}
	}
	rexsult, err = lang.RxFromString("44651895.7").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-910508.438"))
	if err != nil {
		lang.SayString("xmul189")
	} else {
		if !(rexsult.ToString() == "-4.06559278E+13") {
			lang.SayString("xmul189")
		}
	}
	rexsult, err = lang.RxFromString("647897872.E+374021790").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-467.423029"))
	if err != nil {
		lang.SayString("xmul190")
	} else {
		if !(rexsult.ToString() == "-3.02842386E+374021801") {
			lang.SayString("xmul190")
		}
	}
	rexsult, err = lang.RxFromString("25.2592149").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("59.0436981"))
	if err != nil {
		lang.SayString("xmul191")
	} else {
		if !(rexsult.ToString() == "1491.39746") {
			lang.SayString("xmul191")
		}
	}
	rexsult, err = lang.RxFromString("-6.850835").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1273.48240"))
	if err != nil {
		lang.SayString("xmul192")
	} else {
		if !(rexsult.ToString() == "8724.41780") {
			lang.SayString("xmul192")
		}
	}
	rexsult, err = lang.RxFromString("174.272325").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("5638.16229"))
	if err != nil {
		lang.SayString("xmul193")
	} else {
		if !(rexsult.ToString() == "982575.651") {
			lang.SayString("xmul193")
		}
	}
	rexsult, err = lang.RxFromString("3455629.76").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-8.27332322"))
	if err != nil {
		lang.SayString("xmul194")
	} else {
		if !(rexsult.ToString() == "-28589541.9") {
			lang.SayString("xmul194")
		}
	}
	rexsult, err = lang.RxFromString("-924337723E-640771235").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("86639377.1"))
	if err != nil {
		lang.SayString("xmul195")
	} else {
		if !(rexsult.ToString() == "-8.00840446E-640771219") {
			lang.SayString("xmul195")
		}
	}
	rexsult, err = lang.RxFromString("-620236932.E+656823969").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3364722.73"))
	if err != nil {
		lang.SayString("xmul196")
	} else {
		if !(rexsult.ToString() == "-2.08692530E+656823984") {
			lang.SayString("xmul196")
		}
	}
	rexsult, err = lang.RxFromString("9.10025079").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("702777882E-8192234"))
	if err != nil {
		lang.SayString("xmul197")
	} else {
		if !(rexsult.ToString() == "6.39545498E-8192225") {
			lang.SayString("xmul197")
		}
	}
	rexsult, err = lang.RxFromString("-18857539.9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("813013129."))
	if err != nil {
		lang.SayString("xmul198")
	} else {
		if !(rexsult.ToString() == "-1.53314275E+16") {
			lang.SayString("xmul198")
		}
	}
	rexsult, err = lang.RxFromString("-8.29530327").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3243419.57E+35688332"))
	if err != nil {
		lang.SayString("xmul199")
	} else {
		if !(rexsult.ToString() == "-2.69051490E+35688339") {
			lang.SayString("xmul199")
		}
	}
	rexsult, err = lang.RxFromString("-57101683.5").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("763551341E+991491712"))
	if err != nil {
		lang.SayString("xmul200")
	} else {
		if !(rexsult.ToString() == "-4.36000670E+991491728") {
			lang.SayString("xmul200")
		}
	}
	rexsult, err = lang.RxFromString("-603326.740").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1710.95183"))
	if err != nil {
		lang.SayString("xmul201")
	} else {
		if !(rexsult.ToString() == "-1.03226299E+9") {
			lang.SayString("xmul201")
		}
	}
	rexsult, err = lang.RxFromString("-48142763.3").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-943434114"))
	if err != nil {
		lang.SayString("xmul202")
	} else {
		if !(rexsult.ToString() == "4.54195252E+16") {
			lang.SayString("xmul202")
		}
	}
	rexsult, err = lang.RxFromString("-204.586767").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-235.531847"))
	if err != nil {
		lang.SayString("xmul203")
	} else {
		if !(rexsult.ToString() == "48186.6991") {
			lang.SayString("xmul203")
		}
	}
	rexsult, err = lang.RxFromString("-70.3805581").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("830137.913"))
	if err != nil {
		lang.SayString("xmul204")
	} else {
		if !(rexsult.ToString() == "-58425569.6") {
			lang.SayString("xmul204")
		}
	}
	rexsult, err = lang.RxFromString("-8818.47606").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-60766.4571"))
	if err != nil {
		lang.SayString("xmul205")
	} else {
		if !(rexsult.ToString() == "535867547") {
			lang.SayString("xmul205")
		}
	}
	rexsult, err = lang.RxFromString("37060929.3E-168439509").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-79576717.1"))
	if err != nil {
		lang.SayString("xmul206")
	} else {
		if !(rexsult.ToString() == "-2.94918709E-168439494") {
			lang.SayString("xmul206")
		}
	}
	rexsult, err = lang.RxFromString("-656285310.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-107221462."))
	if err != nil {
		lang.SayString("xmul207")
	} else {
		if !(rexsult.ToString() == "7.03678704E+16") {
			lang.SayString("xmul207")
		}
	}
	rexsult, err = lang.RxFromString("653397.125").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7195.30990"))
	if err != nil {
		lang.SayString("xmul208")
	} else {
		if !(rexsult.ToString() == "4.70139480E+9") {
			lang.SayString("xmul208")
		}
	}
	rexsult, err = lang.RxFromString("56221910.0E+857909374").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-58.7247929"))
	if err != nil {
		lang.SayString("xmul209")
	} else {
		if !(rexsult.ToString() == "-3.30162002E+857909383") {
			lang.SayString("xmul209")
		}
	}
	rexsult, err = lang.RxFromString("809862859E+643769974").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-5.06784016"))
	if err != nil {
		lang.SayString("xmul210")
	} else {
		if !(rexsult.ToString() == "-4.10425552E+643769983") {
			lang.SayString("xmul210")
		}
	}
	rexsult, err = lang.RxFromString("-62011.4563E-117563240").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-57.1731586E+115657204"))
	if err != nil {
		lang.SayString("xmul211")
	} else {
		if !(rexsult.ToString() == "3.54539083E-1906030") {
			lang.SayString("xmul211")
		}
	}
	rexsult, err = lang.RxFromString("315.33351").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("91588.837E-536020149"))
	if err != nil {
		lang.SayString("xmul212")
	} else {
		if !(rexsult.ToString() == "2.88810294E-536020142") {
			lang.SayString("xmul212")
		}
	}
	rexsult, err = lang.RxFromString("739.944710").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("202949.175"))
	if err != nil {
		lang.SayString("xmul213")
	} else {
		if !(rexsult.ToString() == "150171168") {
			lang.SayString("xmul213")
		}
	}
	rexsult, err = lang.RxFromString("87686.8016").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4204890.40"))
	if err != nil {
		lang.SayString("xmul214")
	} else {
		if !(rexsult.ToString() == "3.68713390E+11") {
			lang.SayString("xmul214")
		}
	}
	rexsult, err = lang.RxFromString("987126721.E-725794834").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4874166.23"))
	if err != nil {
		lang.SayString("xmul215")
	} else {
		if !(rexsult.ToString() == "4.81141973E-725794819") {
			lang.SayString("xmul215")
		}
	}
	rexsult, err = lang.RxFromString("728148726.E-661695938").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("32798.5202"))
	if err != nil {
		lang.SayString("xmul216")
	} else {
		if !(rexsult.ToString() == "2.38822007E-661695925") {
			lang.SayString("xmul216")
		}
	}
	rexsult, err = lang.RxFromString("7428219.97").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("667.326760"))
	if err != nil {
		lang.SayString("xmul217")
	} else {
		if !(rexsult.ToString() == "4.95704997E+9") {
			lang.SayString("xmul217")
		}
	}
	rexsult, err = lang.RxFromString("-7291.19212").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("209.64966E-588526476"))
	if err != nil {
		lang.SayString("xmul218")
	} else {
		if !(rexsult.ToString() == "-1.52859595E-588526470") {
			lang.SayString("xmul218")
		}
	}
	rexsult, err = lang.RxFromString("-358.24550").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4447.78675E+601402509"))
	if err != nil {
		lang.SayString("xmul219")
	} else {
		if !(rexsult.ToString() == "1.59339959E+601402515") {
			lang.SayString("xmul219")
		}
	}
	rexsult, err = lang.RxFromString("118.621826").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2.72010038"))
	if err != nil {
		lang.SayString("xmul220")
	} else {
		if !(rexsult.ToString() == "-322.663274") {
			lang.SayString("xmul220")
		}
	}
	rexsult, err = lang.RxFromString("8071961.94").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-135533740.E-102451543"))
	if err != nil {
		lang.SayString("xmul221")
	} else {
		if !(rexsult.ToString() == "-1.09402319E-102451528") {
			lang.SayString("xmul221")
		}
	}
	rexsult, err = lang.RxFromString("64262528.5E+812118682").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-8692.94447E-732186947"))
	if err != nil {
		lang.SayString("xmul222")
	} else {
		if !(rexsult.ToString() == "-5.58630592E+79931746") {
			lang.SayString("xmul222")
		}
	}
	rexsult, err = lang.RxFromString("-35544.4029").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-567830.130"))
	if err != nil {
		lang.SayString("xmul223")
	} else {
		if !(rexsult.ToString() == "2.01831829E+10") {
			lang.SayString("xmul223")
		}
	}
	rexsult, err = lang.RxFromString("-7.16513047E+59297103").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("87767.8211"))
	if err != nil {
		lang.SayString("xmul224")
	} else {
		if !(rexsult.ToString() == "-6.28867889E+59297108") {
			lang.SayString("xmul224")
		}
	}
	rexsult, err = lang.RxFromString("-509.483395").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-147242915."))
	if err != nil {
		lang.SayString("xmul225")
	} else {
		if !(rexsult.ToString() == "7.50178202E+10") {
			lang.SayString("xmul225")
		}
	}
	rexsult, err = lang.RxFromString("-7919047.28E+956041629").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-367667329"))
	if err != nil {
		lang.SayString("xmul226")
	} else {
		if !(rexsult.ToString() == "2.91157496E+956041644") {
			lang.SayString("xmul226")
		}
	}
	rexsult, err = lang.RxFromString("895612630.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-36.4104040"))
	if err != nil {
		lang.SayString("xmul227")
	} else {
		if !(rexsult.ToString() == "-3.26096177E+10") {
			lang.SayString("xmul227")
		}
	}
	rexsult, err = lang.RxFromString("25455.4973").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2955.00006E+528196218"))
	if err != nil {
		lang.SayString("xmul228")
	} else {
		if !(rexsult.ToString() == "7.52209960E+528196225") {
			lang.SayString("xmul228")
		}
	}
	rexsult, err = lang.RxFromString("-112.294144E+273414172").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-71448007.7"))
	if err != nil {
		lang.SayString("xmul229")
	} else {
		if !(rexsult.ToString() == "8.02319287E+273414181") {
			lang.SayString("xmul229")
		}
	}
	rexsult, err = lang.RxFromString("62871.2202").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2484.0382E+211662557"))
	if err != nil {
		lang.SayString("xmul230")
	} else {
		if !(rexsult.ToString() == "1.56174513E+211662565") {
			lang.SayString("xmul230")
		}
	}
	rexsult, err = lang.RxFromString("71.9281575").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-9810012.5"))
	if err != nil {
		lang.SayString("xmul231")
	} else {
		if !(rexsult.ToString() == "-705616124") {
			lang.SayString("xmul231")
		}
	}
	rexsult, err = lang.RxFromString("-6388022.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-88.042967"))
	if err != nil {
		lang.SayString("xmul232")
	} else {
		if !(rexsult.ToString() == "562420410") {
			lang.SayString("xmul232")
		}
	}
	rexsult, err = lang.RxFromString("372567445.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("96.0992141"))
	if err != nil {
		lang.SayString("xmul233")
	} else {
		if !(rexsult.ToString() == "3.58034387E+10") {
			lang.SayString("xmul233")
		}
	}
	rexsult, err = lang.RxFromString("802.156517").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-174409310.E-255338020"))
	if err != nil {
		lang.SayString("xmul234")
	} else {
		if !(rexsult.ToString() == "-1.39903565E-255338009") {
			lang.SayString("xmul234")
		}
	}
	rexsult, err = lang.RxFromString("-3.65207541").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("74501982.0"))
	if err != nil {
		lang.SayString("xmul235")
	} else {
		if !(rexsult.ToString() == "-272086856") {
			lang.SayString("xmul235")
		}
	}
	rexsult, err = lang.RxFromString("-5297.76981").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-859.719404"))
	if err != nil {
		lang.SayString("xmul236")
	} else {
		if !(rexsult.ToString() == "4554595.50") {
			lang.SayString("xmul236")
		}
	}
	rexsult, err = lang.RxFromString("-684172.592").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("766.448597E+288361959"))
	if err != nil {
		lang.SayString("xmul237")
	} else {
		if !(rexsult.ToString() == "-5.24383123E+288361967") {
			lang.SayString("xmul237")
		}
	}
	rexsult, err = lang.RxFromString("626919.219").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("57469.8727E+13188610"))
	if err != nil {
		lang.SayString("xmul238")
	} else {
		if !(rexsult.ToString() == "3.60289677E+13188620") {
			lang.SayString("xmul238")
		}
	}
	rexsult, err = lang.RxFromString("-77480.5840").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("893265.594E+287982552"))
	if err != nil {
		lang.SayString("xmul239")
	} else {
		if !(rexsult.ToString() == "-6.92107399E+287982562") {
			lang.SayString("xmul239")
		}
	}
	rexsult, err = lang.RxFromString("-7177620.29").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7786343.83"))
	if err != nil {
		lang.SayString("xmul240")
	} else {
		if !(rexsult.ToString() == "-5.58874195E+13") {
			lang.SayString("xmul240")
		}
	}
	rexsult, err = lang.RxFromString("9.6224130").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4.50355112"))
	if err != nil {
		lang.SayString("xmul241")
	} else {
		if !(rexsult.ToString() == "43.3350288") {
			lang.SayString("xmul241")
		}
	}
	rexsult, err = lang.RxFromString("-66.6337347E-597410086").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-818812885"))
	if err != nil {
		lang.SayString("xmul242")
	} else {
		if !(rexsult.ToString() == "5.45605605E-597410076") {
			lang.SayString("xmul242")
		}
	}
	rexsult, err = lang.RxFromString("65587553.7").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("600574.736"))
	if err != nil {
		lang.SayString("xmul243")
	} else {
		if !(rexsult.ToString() == "3.93902277E+13") {
			lang.SayString("xmul243")
		}
	}
	rexsult, err = lang.RxFromString("-32401.939").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-585200217."))
	if err != nil {
		lang.SayString("xmul244")
	} else {
		if !(rexsult.ToString() == "1.89616217E+13") {
			lang.SayString("xmul244")
		}
	}
	rexsult, err = lang.RxFromString("69573.988").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-9.77003465E+740933668"))
	if err != nil {
		lang.SayString("xmul245")
	} else {
		if !(rexsult.ToString() == "-6.79740273E+740933673") {
			lang.SayString("xmul245")
		}
	}
	rexsult, err = lang.RxFromString("2362.06251").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-433149546.E-152643629"))
	if err != nil {
		lang.SayString("xmul246")
	} else {
		if !(rexsult.ToString() == "-1.02312630E-152643617") {
			lang.SayString("xmul246")
		}
	}
	rexsult, err = lang.RxFromString("-615.23488E+249953452").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-21437483.7"))
	if err != nil {
		lang.SayString("xmul247")
	} else {
		if !(rexsult.ToString() == "1.31890877E+249953462") {
			lang.SayString("xmul247")
		}
	}
	rexsult, err = lang.RxFromString("216741082.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("250290244"))
	if err != nil {
		lang.SayString("xmul248")
	} else {
		if !(rexsult.ToString() == "5.42481783E+16") {
			lang.SayString("xmul248")
		}
	}
	rexsult, err = lang.RxFromString("-6364720.49").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("5539245.64"))
	if err != nil {
		lang.SayString("xmul249")
	} else {
		if !(rexsult.ToString() == "-3.52557502E+13") {
			lang.SayString("xmul249")
		}
	}
	rexsult, err = lang.RxFromString("-814599.475").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-14.5431191"))
	if err != nil {
		lang.SayString("xmul250")
	} else {
		if !(rexsult.ToString() == "11846817.2") {
			lang.SayString("xmul250")
		}
	}
	rexsult, err = lang.RxFromString("-877498.755").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("507408724E-168628106"))
	if err != nil {
		lang.SayString("xmul251")
	} else {
		if !(rexsult.ToString() == "-4.45250524E-168628092") {
			lang.SayString("xmul251")
		}
	}
	rexsult, err = lang.RxFromString("10634446.5E+475783861").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("50.7213056E+17807809"))
	if err != nil {
		lang.SayString("xmul252")
	} else {
		if !(rexsult.ToString() == "5.39393011E+493591678") {
			lang.SayString("xmul252")
		}
	}
	rexsult, err = lang.RxFromString("-162726.257E-597285918").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4391.54799"))
	if err != nil {
		lang.SayString("xmul253")
	} else {
		if !(rexsult.ToString() == "7.14620167E-597285910") {
			lang.SayString("xmul253")
		}
	}
	rexsult, err = lang.RxFromString("700354586.E-99856707").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7198.0493E+436250299"))
	if err != nil {
		lang.SayString("xmul254")
	} else {
		if !(rexsult.ToString() == "5.04118684E+336393604") {
			lang.SayString("xmul254")
		}
	}
	rexsult, err = lang.RxFromString("39617663E-463704664").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-895.290346"))
	if err != nil {
		lang.SayString("xmul255")
	} else {
		if !(rexsult.ToString() == "-3.54693112E-463704654") {
			lang.SayString("xmul255")
		}
	}
	rexsult, err = lang.RxFromString("5350882.59").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-36329829"))
	if err != nil {
		lang.SayString("xmul256")
	} else {
		if !(rexsult.ToString() == "-1.94396649E+14") {
			lang.SayString("xmul256")
		}
	}
	rexsult, err = lang.RxFromString("91966.4084E+210382952").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("166740.46E-42001390"))
	if err != nil {
		lang.SayString("xmul257")
	} else {
		if !(rexsult.ToString() == "1.53345212E+168381572") {
			lang.SayString("xmul257")
		}
	}
	rexsult, err = lang.RxFromString("231899031.E-481759076").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("726.337100"))
	if err != nil {
		lang.SayString("xmul258")
	} else {
		if !(rexsult.ToString() == "1.68436870E-481759065") {
			lang.SayString("xmul258")
		}
	}
	rexsult, err = lang.RxFromString("-9611312.33").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("22109735.9"))
	if err != nil {
		lang.SayString("xmul259")
	} else {
		if !(rexsult.ToString() == "-2.12503577E+14") {
			lang.SayString("xmul259")
		}
	}
	rexsult, err = lang.RxFromString("-5604938.15E-36812542").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("735937577."))
	if err != nil {
		lang.SayString("xmul260")
	} else {
		if !(rexsult.ToString() == "-4.12488460E-36812527") {
			lang.SayString("xmul260")
		}
	}
	rexsult, err = lang.RxFromString("693881413.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("260547224E-480281418"))
	if err != nil {
		lang.SayString("xmul261")
	} else {
		if !(rexsult.ToString() == "1.80788876E-480281401") {
			lang.SayString("xmul261")
		}
	}
	rexsult, err = lang.RxFromString("-34865.7378E-368768024").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2297117.88"))
	if err != nil {
		lang.SayString("xmul262")
	} else {
		if !(rexsult.ToString() == "-8.00907097E-368768014") {
			lang.SayString("xmul262")
		}
	}
	rexsult, err = lang.RxFromString("1123.32456").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7.86747918E+930888796"))
	if err != nil {
		lang.SayString("xmul263")
	} else {
		if !(rexsult.ToString() == "8.83773259E+930888799") {
			lang.SayString("xmul263")
		}
	}
	rexsult, err = lang.RxFromString("-1.85771840E+365552540").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-73028339.7"))
	if err != nil {
		lang.SayString("xmul265")
	} else {
		if !(rexsult.ToString() == "1.35666090E+365552548") {
			lang.SayString("xmul265")
		}
	}
	rexsult, err = lang.RxFromString("34.1935525").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-40767.6450"))
	if err != nil {
		lang.SayString("xmul266")
	} else {
		if !(rexsult.ToString() == "-1393990.61") {
			lang.SayString("xmul266")
		}
	}
	rexsult, err = lang.RxFromString("26.0009168E+751618294").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-304019.929"))
	if err != nil {
		lang.SayString("xmul267")
	} else {
		if !(rexsult.ToString() == "-7.90479688E+751618300") {
			lang.SayString("xmul267")
		}
	}
	rexsult, err = lang.RxFromString("-58.4853072E+588540055").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4647.3205"))
	if err != nil {
		lang.SayString("xmul268")
	} else {
		if !(rexsult.ToString() == "2.71799967E+588540060") {
			lang.SayString("xmul268")
		}
	}
	rexsult, err = lang.RxFromString("51.025101").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4467691.57"))
	if err != nil {
		lang.SayString("xmul269")
	} else {
		if !(rexsult.ToString() == "-227964414") {
			lang.SayString("xmul269")
		}
	}
	rexsult, err = lang.RxFromString("-2214.76582").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("379785372E+223117572"))
	if err != nil {
		lang.SayString("xmul270")
	} else {
		if !(rexsult.ToString() == "-8.41135661E+223117583") {
			lang.SayString("xmul270")
		}
	}
	rexsult, err = lang.RxFromString("-2564.75207E-841443929").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-653498187"))
	if err != nil {
		lang.SayString("xmul271")
	} else {
		if !(rexsult.ToString() == "1.67606083E-841443917") {
			lang.SayString("xmul271")
		}
	}
	rexsult, err = lang.RxFromString("513115529.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("27775075.6E+217133352"))
	if err != nil {
		lang.SayString("xmul272")
	} else {
		if !(rexsult.ToString() == "1.42518226E+217133368") {
			lang.SayString("xmul272")
		}
	}
	rexsult, err = lang.RxFromString("-247157.208").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-532990.453"))
	if err != nil {
		lang.SayString("xmul273")
	} else {
		if !(rexsult.ToString() == "1.31732432E+11") {
			lang.SayString("xmul273")
		}
	}
	rexsult, err = lang.RxFromString("40.2490764E-339482253").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7626.85442E+594264540"))
	if err != nil {
		lang.SayString("xmul274")
	} else {
		if !(rexsult.ToString() == "3.06973846E+254782292") {
			lang.SayString("xmul274")
		}
	}
	rexsult, err = lang.RxFromString("-1156008.8").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-8870382.36"))
	if err != nil {
		lang.SayString("xmul275")
	} else {
		if !(rexsult.ToString() == "1.02542401E+13") {
			lang.SayString("xmul275")
		}
	}
	rexsult, err = lang.RxFromString("880097928.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-52455011.1E+204538218"))
	if err != nil {
		lang.SayString("xmul276")
	} else {
		if !(rexsult.ToString() == "-4.61655466E+204538234") {
			lang.SayString("xmul276")
		}
	}
	rexsult, err = lang.RxFromString("5796.2524").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("34458329.7E+832129426"))
	if err != nil {
		lang.SayString("xmul277")
	} else {
		if !(rexsult.ToString() == "1.99729176E+832129437") {
			lang.SayString("xmul277")
		}
	}
	rexsult, err = lang.RxFromString("27.1000923E-218032223").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-45.0198341"))
	if err != nil {
		lang.SayString("xmul278")
	} else {
		if !(rexsult.ToString() == "-1.22004166E-218032220") {
			lang.SayString("xmul278")
		}
	}
	rexsult, err = lang.RxFromString("42643477.8").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("26118465E-730390549"))
	if err != nil {
		lang.SayString("xmul279")
	} else {
		if !(rexsult.ToString() == "1.11378218E-730390534") {
			lang.SayString("xmul279")
		}
	}
	rexsult, err = lang.RxFromString("-31918.9176E-163031657").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-21.5422824E-807317258"))
	if err != nil {
		lang.SayString("xmul280")
	} else {
		if !(rexsult.ToString() == "6.87606337E-970348910") {
			lang.SayString("xmul280")
		}
	}
	rexsult, err = lang.RxFromString("84224841.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2.62548255E+647087608"))
	if err != nil {
		lang.SayString("xmul281")
	} else {
		if !(rexsult.ToString() == "2.21130850E+647087616") {
			lang.SayString("xmul281")
		}
	}
	rexsult, err = lang.RxFromString("-64413698.9").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-6674.1055E-701047852"))
	if err != nil {
		lang.SayString("xmul282")
	} else {
		if !(rexsult.ToString() == "4.29903822E-701047841") {
			lang.SayString("xmul282")
		}
	}
	rexsult, err = lang.RxFromString("-62.5059208").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9.5795779E-898350012"))
	if err != nil {
		lang.SayString("xmul283")
	} else {
		if !(rexsult.ToString() == "-5.98780338E-898350010") {
			lang.SayString("xmul283")
		}
	}
	rexsult, err = lang.RxFromString("9090950.80").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("436.400932"))
	if err != nil {
		lang.SayString("xmul284")
	} else {
		if !(rexsult.ToString() == "3.96729940E+9") {
			lang.SayString("xmul284")
		}
	}
	rexsult, err = lang.RxFromString("-89833825.7E+329205393").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-779430.194"))
	if err != nil {
		lang.SayString("xmul285")
	} else {
		if !(rexsult.ToString() == "7.00191962E+329205406") {
			lang.SayString("xmul285")
		}
	}
	rexsult, err = lang.RxFromString("-714562.019E+750205688").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("704079764"))
	if err != nil {
		lang.SayString("xmul286")
	} else {
		if !(rexsult.ToString() == "-5.03108658E+750205702") {
			lang.SayString("xmul286")
		}
	}
	rexsult, err = lang.RxFromString("-584537670.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("31139.7737E-146687560"))
	if err != nil {
		lang.SayString("xmul287")
	} else {
		if !(rexsult.ToString() == "-1.82023708E-146687547") {
			lang.SayString("xmul287")
		}
	}
	rexsult, err = lang.RxFromString("5.15309635").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-695649.219E+451948183"))
	if err != nil {
		lang.SayString("xmul289")
	} else {
		if !(rexsult.ToString() == "-3.58474745E+451948189") {
			lang.SayString("xmul289")
		}
	}
	rexsult, err = lang.RxFromString("-940030153.E+83797657").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4.11510193"))
	if err != nil {
		lang.SayString("xmul290")
	} else {
		if !(rexsult.ToString() == "3.86831990E+83797666") {
			lang.SayString("xmul290")
		}
	}
	rexsult, err = lang.RxFromString("89088.9683E+587739290").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.31932110"))
	if err != nil {
		lang.SayString("xmul291")
	} else {
		if !(rexsult.ToString() == "1.17536956E+587739295") {
			lang.SayString("xmul291")
		}
	}
	rexsult, err = lang.RxFromString("3336750").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("6.47961126"))
	if err != nil {
		lang.SayString("xmul292")
	} else {
		if !(rexsult.ToString() == "21620842.9") {
			lang.SayString("xmul292")
		}
	}
	rexsult, err = lang.RxFromString("904654622.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("692065270.E+329081915"))
	if err != nil {
		lang.SayString("xmul293")
	} else {
		if !(rexsult.ToString() == "6.26080045E+329081932") {
			lang.SayString("xmul293")
		}
	}
	rexsult, err = lang.RxFromString("304804380").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4681.23698"))
	if err != nil {
		lang.SayString("xmul294")
	} else {
		if !(rexsult.ToString() == "-1.42686154E+12") {
			lang.SayString("xmul294")
		}
	}
	rexsult, err = lang.RxFromString("674.55569").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-82981.2684E+852890752"))
	if err != nil {
		lang.SayString("xmul295")
	} else {
		if !(rexsult.ToString() == "-5.59754868E+852890759") {
			lang.SayString("xmul295")
		}
	}
	rexsult, err = lang.RxFromString("-5111.51025E-108006096").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("5448870.4E+279212255"))
	if err != nil {
		lang.SayString("xmul296")
	} else {
		if !(rexsult.ToString() == "-2.78519569E+171206169") {
			lang.SayString("xmul296")
		}
	}
	rexsult, err = lang.RxFromString("-2623.45068").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-466463938."))
	if err != nil {
		lang.SayString("xmul297")
	} else {
		if !(rexsult.ToString() == "1.22374514E+12") {
			lang.SayString("xmul297")
		}
	}
	rexsult, err = lang.RxFromString("299350.435").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3373.33551"))
	if err != nil {
		lang.SayString("xmul298")
	} else {
		if !(rexsult.ToString() == "1.00980945E+9") {
			lang.SayString("xmul298")
		}
	}
	rexsult, err = lang.RxFromString("-6589947.80").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2448.75933E-591549734"))
	if err != nil {
		lang.SayString("xmul299")
	} else {
		if !(rexsult.ToString() == "1.61371962E-591549724") {
			lang.SayString("xmul299")
		}
	}
	rexsult, err = lang.RxFromString("3774.5358E-491090520").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("173.060090"))
	if err != nil {
		lang.SayString("xmul300")
	} else {
		if !(rexsult.ToString() == "6.53221505E-491090515") {
			lang.SayString("xmul300")
		}
	}
	rexsult, err = lang.RxFromString("-13.6783690").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-453.610117"))
	if err != nil {
		lang.SayString("xmul301")
	} else {
		if !(rexsult.ToString() == "6204.64656") {
			lang.SayString("xmul301")
		}
	}
	rexsult, err = lang.RxFromString("-990100927.E-615244634").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("223801.421E+247632618"))
	if err != nil {
		lang.SayString("xmul302")
	} else {
		if !(rexsult.ToString() == "-2.21585994E-367612002") {
			lang.SayString("xmul302")
		}
	}
	rexsult, err = lang.RxFromString("1275.10292").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-667965353"))
	if err != nil {
		lang.SayString("xmul303")
	} else {
		if !(rexsult.ToString() == "-8.51724572E+11") {
			lang.SayString("xmul303")
		}
	}
	rexsult, err = lang.RxFromString("-8.76375480E-596792197").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("992.077361"))
	if err != nil {
		lang.SayString("xmul304")
	} else {
		if !(rexsult.ToString() == "-8.69432273E-596792194") {
			lang.SayString("xmul304")
		}
	}
	rexsult, err = lang.RxFromString("953.976935E+385444720").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("96503.3378"))
	if err != nil {
		lang.SayString("xmul305")
	} else {
		if !(rexsult.ToString() == "9.20619584E+385444727") {
			lang.SayString("xmul305")
		}
	}
	rexsult, err = lang.RxFromString("213577152").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-986710073E+31900046"))
	if err != nil {
		lang.SayString("xmul306")
	} else {
		if !(rexsult.ToString() == "-2.10738727E+31900063") {
			lang.SayString("xmul306")
		}
	}
	rexsult, err = lang.RxFromString("91393.9398E-323439228").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-135.701000"))
	if err != nil {
		lang.SayString("xmul307")
	} else {
		if !(rexsult.ToString() == "-1.24022490E-323439221") {
			lang.SayString("xmul307")
		}
	}
	rexsult, err = lang.RxFromString("-396.503557").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("45757264.E-254363788"))
	if err != nil {
		lang.SayString("xmul308")
	} else {
		if !(rexsult.ToString() == "-1.81429179E-254363778") {
			lang.SayString("xmul308")
		}
	}
	rexsult, err = lang.RxFromString("59807846.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.53345254"))
	if err != nil {
		lang.SayString("xmul309")
	} else {
		if !(rexsult.ToString() == "91712493.5") {
			lang.SayString("xmul309")
		}
	}
	rexsult, err = lang.RxFromString("-8046158.45").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("8.3635397"))
	if err != nil {
		lang.SayString("xmul310")
	} else {
		if !(rexsult.ToString() == "-67294365.6") {
			lang.SayString("xmul310")
		}
	}
	rexsult, err = lang.RxFromString("55.1123381E+50627250").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-94.0355047E-162540316"))
	if err != nil {
		lang.SayString("xmul311")
	} else {
		if !(rexsult.ToString() == "-5.18251653E-111913063") {
			lang.SayString("xmul311")
		}
	}
	rexsult, err = lang.RxFromString("-948.038054").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3580.84510"))
	if err != nil {
		lang.SayString("xmul312")
	} else {
		if !(rexsult.ToString() == "-3394777.42") {
			lang.SayString("xmul312")
		}
	}
	rexsult, err = lang.RxFromString("-6026.42752").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-14.2286406E-334921364"))
	if err != nil {
		lang.SayString("xmul313")
	} else {
		if !(rexsult.ToString() == "8.57478713E-334921360") {
			lang.SayString("xmul313")
		}
	}
	rexsult, err = lang.RxFromString("79551.5014").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-538.186229"))
	if err != nil {
		lang.SayString("xmul314")
	} else {
		if !(rexsult.ToString() == "-42813522.5") {
			lang.SayString("xmul314")
		}
	}
	rexsult, err = lang.RxFromString("42706056.E+623578292").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-690.327745"))
	if err != nil {
		lang.SayString("xmul315")
	} else {
		if !(rexsult.ToString() == "-2.94811753E+623578302") {
			lang.SayString("xmul315")
		}
	}
	rexsult, err = lang.RxFromString("2454136.08E+502374077").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("856268.795E-356664934"))
	if err != nil {
		lang.SayString("xmul316")
	} else {
		if !(rexsult.ToString() == "2.10140014E+145709155") {
			lang.SayString("xmul316")
		}
	}
	rexsult, err = lang.RxFromString("-3264204.54").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-42704.501"))
	if err != nil {
		lang.SayString("xmul317")
	} else {
		if !(rexsult.ToString() == "1.39396226E+11") {
			lang.SayString("xmul317")
		}
	}
	rexsult, err = lang.RxFromString("1.21265492").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("44102.6073"))
	if err != nil {
		lang.SayString("xmul318")
	} else {
		if !(rexsult.ToString() == "53481.2437") {
			lang.SayString("xmul318")
		}
	}
	rexsult, err = lang.RxFromString("-19.054711E+975514652").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-22144.0822"))
	if err != nil {
		lang.SayString("xmul319")
	} else {
		if !(rexsult.ToString() == "4.21949087E+975514657") {
			lang.SayString("xmul319")
		}
	}
	rexsult, err = lang.RxFromString("745.78452").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-1922.00670E+375923302"))
	if err != nil {
		lang.SayString("xmul320")
	} else {
		if !(rexsult.ToString() == "-1.43340284E+375923308") {
			lang.SayString("xmul320")
		}
	}
	rexsult, err = lang.RxFromString("-963717836").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-823989308"))
	if err != nil {
		lang.SayString("xmul321")
	} else {
		if !(rexsult.ToString() == "7.94093193E+17") {
			lang.SayString("xmul321")
		}
	}
	rexsult, err = lang.RxFromString("-808328.607E-790810342").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("53075.7082"))
	if err != nil {
		lang.SayString("xmul323")
	} else {
		if !(rexsult.ToString() == "-4.29026133E-790810332") {
			lang.SayString("xmul323")
		}
	}
	rexsult, err = lang.RxFromString("700592.720").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-698485.085"))
	if err != nil {
		lang.SayString("xmul324")
	} else {
		if !(rexsult.ToString() == "-4.89353566E+11") {
			lang.SayString("xmul324")
		}
	}
	rexsult, err = lang.RxFromString("-80273928.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("661346.239"))
	if err != nil {
		lang.SayString("xmul325")
	} else {
		if !(rexsult.ToString() == "-5.30888604E+13") {
			lang.SayString("xmul325")
		}
	}
	rexsult, err = lang.RxFromString("-24018251.0E+819786764").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("59141.9600E-167165065"))
	if err != nil {
		lang.SayString("xmul326")
	} else {
		if !(rexsult.ToString() == "-1.42048644E+652621711") {
			lang.SayString("xmul326")
		}
	}
	rexsult, err = lang.RxFromString("2512953.3").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-3769170.35E-993621645"))
	if err != nil {
		lang.SayString("xmul327")
	} else {
		if !(rexsult.ToString() == "-9.47174907E-993621633") {
			lang.SayString("xmul327")
		}
	}
	rexsult, err = lang.RxFromString("-682.796370").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("71131.0224"))
	if err != nil {
		lang.SayString("xmul328")
	} else {
		if !(rexsult.ToString() == "-48568003.9") {
			lang.SayString("xmul328")
		}
	}
	rexsult, err = lang.RxFromString("89.9997490").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4993.69831"))
	if err != nil {
		lang.SayString("xmul329")
	} else {
		if !(rexsult.ToString() == "-449431.594") {
			lang.SayString("xmul329")
		}
	}
	rexsult, err = lang.RxFromString("76563354.6E-112338836").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("278271.585E-511481095"))
	if err != nil {
		lang.SayString("xmul330")
	} else {
		if !(rexsult.ToString() == "2.13054060E-623819918") {
			lang.SayString("xmul330")
		}
	}
	rexsult, err = lang.RxFromString("-932499.010").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("873.377701E-502190452"))
	if err != nil {
		lang.SayString("xmul331")
	} else {
		if !(rexsult.ToString() == "-8.14423842E-502190444") {
			lang.SayString("xmul331")
		}
	}
	rexsult, err = lang.RxFromString("-7735918.21E+799514797").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-7748.78023"))
	if err != nil {
		lang.SayString("xmul332")
	} else {
		if !(rexsult.ToString() == "5.99439301E+799514807") {
			lang.SayString("xmul332")
		}
	}
	rexsult, err = lang.RxFromString("-3708780.75E+445232787").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("980.006567E-780728623"))
	if err != nil {
		lang.SayString("xmul333")
	} else {
		if !(rexsult.ToString() == "-3.63462949E-335495827") {
			lang.SayString("xmul333")
		}
	}
	rexsult, err = lang.RxFromString("-5205124.44E-140588661").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-495394029.E-620856313"))
	if err != nil {
		lang.SayString("xmul334")
	} else {
		if !(rexsult.ToString() == "2.57858757E-761444959") {
			lang.SayString("xmul334")
		}
	}
	rexsult, err = lang.RxFromString("-8868.72074").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("5592399.93"))
	if err != nil {
		lang.SayString("xmul335")
	} else {
		if !(rexsult.ToString() == "-4.95974332E+10") {
			lang.SayString("xmul335")
		}
	}
	rexsult, err = lang.RxFromString("-74.7852037E-175205809").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("4.14316542"))
	if err != nil {
		lang.SayString("xmul336")
	} else {
		if !(rexsult.ToString() == "-3.09847470E-175205807") {
			lang.SayString("xmul336")
		}
	}
	rexsult, err = lang.RxFromString("84196.1091E+242628748").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("8.07523036E-288231467"))
	if err != nil {
		lang.SayString("xmul337")
	} else {
		if !(rexsult.ToString() == "6.79902976E-45602714") {
			lang.SayString("xmul337")
		}
	}
	rexsult, err = lang.RxFromString("38660103.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-6671.73085E+900998477"))
	if err != nil {
		lang.SayString("xmul338")
	} else {
		if !(rexsult.ToString() == "-2.57929803E+900998488") {
			lang.SayString("xmul338")
		}
	}
	rexsult, err = lang.RxFromString("-52.2659460").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-296404199E+372050476"))
	if err != nil {
		lang.SayString("xmul339")
	} else {
		if !(rexsult.ToString() == "1.54918459E+372050486") {
			lang.SayString("xmul339")
		}
	}
	rexsult, err = lang.RxFromString("6.06625013").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-276.359186"))
	if err != nil {
		lang.SayString("xmul340")
	} else {
		if !(rexsult.ToString() == "-1676.46395") {
			lang.SayString("xmul340")
		}
	}
	rexsult, err = lang.RxFromString("-62971617.5E-241444744").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("46266799.3"))
	if err != nil {
		lang.SayString("xmul341")
	} else {
		if !(rexsult.ToString() == "-2.91349519E-241444729") {
			lang.SayString("xmul341")
		}
	}
	rexsult, err = lang.RxFromString("-5.36917800").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-311124593.E-976066491"))
	if err != nil {
		lang.SayString("xmul342")
	} else {
		if !(rexsult.ToString() == "1.67048332E-976066482") {
			lang.SayString("xmul342")
		}
	}
	rexsult, err = lang.RxFromString("2467915.01").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-92.5558322"))
	if err != nil {
		lang.SayString("xmul343")
	} else {
		if !(rexsult.ToString() == "-228419928") {
			lang.SayString("xmul343")
		}
	}
	rexsult, err = lang.RxFromString("187.232671").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-840.469347"))
	if err != nil {
		lang.SayString("xmul344")
	} else {
		if !(rexsult.ToString() == "-157363.321") {
			lang.SayString("xmul344")
		}
	}
	rexsult, err = lang.RxFromString("81233.6823").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-5192.21666E+309315093"))
	if err != nil {
		lang.SayString("xmul345")
	} else {
		if !(rexsult.ToString() == "-4.21782879E+309315101") {
			lang.SayString("xmul345")
		}
	}
	rexsult, err = lang.RxFromString("-854.586113").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-79.8715762E-853065103"))
	if err != nil {
		lang.SayString("xmul346")
	} else {
		if !(rexsult.ToString() == "6.82571398E-853065099") {
			lang.SayString("xmul346")
		}
	}
	rexsult, err = lang.RxFromString("78872665.3").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("172.102119"))
	if err != nil {
		lang.SayString("xmul347")
	} else {
		if !(rexsult.ToString() == "1.35741528E+10") {
			lang.SayString("xmul347")
		}
	}
	rexsult, err = lang.RxFromString("328268.1E-436315617").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-204.522245"))
	if err != nil {
		lang.SayString("xmul348")
	} else {
		if !(rexsult.ToString() == "-6.71381288E-436315610") {
			lang.SayString("xmul348")
		}
	}
	rexsult, err = lang.RxFromString("-4037911.02E+641367645").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("29.5713010"))
	if err != nil {
		lang.SayString("xmul349")
	} else {
		if !(rexsult.ToString() == "-1.19406282E+641367653") {
			lang.SayString("xmul349")
		}
	}
	rexsult, err = lang.RxFromString("-688755561.E-95301699").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("978.275312E+913812609"))
	if err != nil {
		lang.SayString("xmul350")
	} else {
		if !(rexsult.ToString() == "-6.73792561E+818510921") {
			lang.SayString("xmul350")
		}
	}
	rexsult, err = lang.RxFromString("-5.47345502").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("59818.7580"))
	if err != nil {
		lang.SayString("xmul351")
	} else {
		if !(rexsult.ToString() == "-327415.281") {
			lang.SayString("xmul351")
		}
	}
	rexsult, err = lang.RxFromString("563891620E-361354567").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-845900362."))
	if err != nil {
		lang.SayString("xmul352")
	} else {
		if !(rexsult.ToString() == "-4.76996125E-361354550") {
			lang.SayString("xmul352")
		}
	}
	rexsult, err = lang.RxFromString("-69.7231286").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("85773.7504"))
	if err != nil {
		lang.SayString("xmul353")
	} else {
		if !(rexsult.ToString() == "-5980414.23") {
			lang.SayString("xmul353")
		}
	}
	rexsult, err = lang.RxFromString("5125.51188").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("73814638.4E-500934741"))
	if err != nil {
		lang.SayString("xmul354")
	} else {
		if !(rexsult.ToString() == "3.78337806E-500934730") {
			lang.SayString("xmul354")
		}
	}
	rexsult, err = lang.RxFromString("-54.6254096").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-332921899."))
	if err != nil {
		lang.SayString("xmul355")
	} else {
		if !(rexsult.ToString() == "1.81859951E+10") {
			lang.SayString("xmul355")
		}
	}
	rexsult, err = lang.RxFromString("-9.04778095E-591874079").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("8719.40286"))
	if err != nil {
		lang.SayString("xmul356")
	} else {
		if !(rexsult.ToString() == "-7.88912471E-591874075") {
			lang.SayString("xmul356")
		}
	}
	rexsult, err = lang.RxFromString("-21006.1733E+884684431").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-48872.9175"))
	if err != nil {
		lang.SayString("xmul357")
	} else {
		if !(rexsult.ToString() == "1.02663297E+884684440") {
			lang.SayString("xmul357")
		}
	}
	rexsult, err = lang.RxFromString("-1546783").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-51935370.4"))
	if err != nil {
		lang.SayString("xmul358")
	} else {
		if !(rexsult.ToString() == "8.03327480E+13") {
			lang.SayString("xmul358")
		}
	}
	rexsult, err = lang.RxFromString("61302486.8").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("205.490417"))
	if err != nil {
		lang.SayString("xmul359")
	} else {
		if !(rexsult.ToString() == "1.25970736E+10") {
			lang.SayString("xmul359")
		}
	}
	rexsult, err = lang.RxFromString("-318180109.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-54008744.6E-170931002"))
	if err != nil {
		lang.SayString("xmul360")
	} else {
		if !(rexsult.ToString() == "1.71845082E-170930986") {
			lang.SayString("xmul360")
		}
	}
	rexsult, err = lang.RxFromString("-28486137.1E+901441714").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-42454.940"))
	if err != nil {
		lang.SayString("xmul361")
	} else {
		if !(rexsult.ToString() == "1.20937724E+901441726") {
			lang.SayString("xmul361")
		}
	}
	rexsult, err = lang.RxFromString("-546398328.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-27.9149712"))
	if err != nil {
		lang.SayString("xmul362")
	} else {
		if !(rexsult.ToString() == "1.52526936E+10") {
			lang.SayString("xmul362")
		}
	}
	rexsult, err = lang.RxFromString("5402066.1E-284978216").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("622.751128"))
	if err != nil {
		lang.SayString("xmul363")
	} else {
		if !(rexsult.ToString() == "3.36414276E-284978207") {
			lang.SayString("xmul363")
		}
	}
	rexsult, err = lang.RxFromString("18845620").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3129.43753"))
	if err != nil {
		lang.SayString("xmul364")
	} else {
		if !(rexsult.ToString() == "5.89761905E+10") {
			lang.SayString("xmul364")
		}
	}
	rexsult, err = lang.RxFromString("55.8245006E+928885991").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("99170843.9E-47402167"))
	if err != nil {
		lang.SayString("xmul366")
	} else {
		if !(rexsult.ToString() == "5.53616283E+881483833") {
			lang.SayString("xmul366")
		}
	}
	rexsult, err = lang.RxFromString("13.8003883E-386224921").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-84126481.9E-296378341"))
	if err != nil {
		lang.SayString("xmul367")
	} else {
		if !(rexsult.ToString() == "-1.16097812E-682603253") {
			lang.SayString("xmul367")
		}
	}
	rexsult, err = lang.RxFromString("9820.90457").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("46671.5915"))
	if err != nil {
		lang.SayString("xmul368")
	} else {
		if !(rexsult.ToString() == "458357246") {
			lang.SayString("xmul368")
		}
	}
	rexsult, err = lang.RxFromString("472648900").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-207.784153"))
	if err != nil {
		lang.SayString("xmul370")
	} else {
		if !(rexsult.ToString() == "-9.82089514E+10") {
			lang.SayString("xmul370")
		}
	}
	rexsult, err = lang.RxFromString("-8754.49306").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-818.165153E+631475457"))
	if err != nil {
		lang.SayString("xmul371")
	} else {
		if !(rexsult.ToString() == "7.16262115E+631475463") {
			lang.SayString("xmul371")
		}
	}
	rexsult, err = lang.RxFromString("98750864").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("191380.551"))
	if err != nil {
		lang.SayString("xmul372")
	} else {
		if !(rexsult.ToString() == "1.88989948E+13") {
			lang.SayString("xmul372")
		}
	}
	rexsult, err = lang.RxFromString("725292561.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-768963606.E+340762986"))
	if err != nil {
		lang.SayString("xmul373")
	} else {
		if !(rexsult.ToString() == "-5.57723583E+340763003") {
			lang.SayString("xmul373")
		}
	}
	rexsult, err = lang.RxFromString("1862.80445").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("648254483."))
	if err != nil {
		lang.SayString("xmul374")
	} else {
		if !(rexsult.ToString() == "1.20757134E+12") {
			lang.SayString("xmul374")
		}
	}
	rexsult, err = lang.RxFromString("-5549320.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-93580684.1"))
	if err != nil {
		lang.SayString("xmul375")
	} else {
		if !(rexsult.ToString() == "5.19309171E+14") {
			lang.SayString("xmul375")
		}
	}
	rexsult, err = lang.RxFromString("-14677053.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-25784.7358"))
	if err != nil {
		lang.SayString("xmul376")
	} else {
		if !(rexsult.ToString() == "3.78443937E+11") {
			lang.SayString("xmul376")
		}
	}
	rexsult, err = lang.RxFromString("-4131738.09").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7579.07566"))
	if err != nil {
		lang.SayString("xmul378")
	} else {
		if !(rexsult.ToString() == "-3.13147556E+10") {
			lang.SayString("xmul378")
		}
	}
	rexsult, err = lang.RxFromString("504544.648").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-7678.96133E-662143268"))
	if err != nil {
		lang.SayString("xmul379")
	} else {
		if !(rexsult.ToString() == "-3.87437884E-662143259") {
			lang.SayString("xmul379")
		}
	}
	rexsult, err = lang.RxFromString("829898241").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("8912.99114E+929228149"))
	if err != nil {
		lang.SayString("xmul380")
	} else {
		if !(rexsult.ToString() == "7.39687567E+929228161") {
			lang.SayString("xmul380")
		}
	}
	rexsult, err = lang.RxFromString("53.6891691").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-11.2371140"))
	if err != nil {
		lang.SayString("xmul381")
	} else {
		if !(rexsult.ToString() == "-603.311314") {
			lang.SayString("xmul381")
		}
	}
	rexsult, err = lang.RxFromString("-93951823.4").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-25317.8645"))
	if err != nil {
		lang.SayString("xmul382")
	} else {
		if !(rexsult.ToString() == "2.37865953E+12") {
			lang.SayString("xmul382")
		}
	}
	rexsult, err = lang.RxFromString("446919.123").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("951338490."))
	if err != nil {
		lang.SayString("xmul383")
	} else {
		if !(rexsult.ToString() == "4.25171364E+14") {
			lang.SayString("xmul383")
		}
	}
	rexsult, err = lang.RxFromString("-8.01787748").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-88.3076852"))
	if err != nil {
		lang.SayString("xmul384")
	} else {
		if !(rexsult.ToString() == "708.040200") {
			lang.SayString("xmul384")
		}
	}
	rexsult, err = lang.RxFromString("517458139").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-999731.548"))
	if err != nil {
		lang.SayString("xmul385")
	} else {
		if !(rexsult.ToString() == "-5.17319226E+14") {
			lang.SayString("xmul385")
		}
	}
	rexsult, err = lang.RxFromString("-405543440").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4013.18295"))
	if err != nil {
		lang.SayString("xmul386")
	} else {
		if !(rexsult.ToString() == "1.62752002E+12") {
			lang.SayString("xmul386")
		}
	}
	rexsult, err = lang.RxFromString("-49245250.1E+682760825").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-848776.637"))
	if err != nil {
		lang.SayString("xmul387")
	} else {
		if !(rexsult.ToString() == "4.17982178E+682760838") {
			lang.SayString("xmul387")
		}
	}
	rexsult, err = lang.RxFromString("-151144455").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-170371.29"))
	if err != nil {
		lang.SayString("xmul388")
	} else {
		if !(rexsult.ToString() == "2.57506758E+13") {
			lang.SayString("xmul388")
		}
	}
	rexsult, err = lang.RxFromString("-729236746.E+662737067").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9.10823602"))
	if err != nil {
		lang.SayString("xmul389")
	} else {
		if !(rexsult.ToString() == "-6.64206040E+662737076") {
			lang.SayString("xmul389")
		}
	}
	rexsult, err = lang.RxFromString("534.394729").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2369839.37"))
	if err != nil {
		lang.SayString("xmul390")
	} else {
		if !(rexsult.ToString() == "-1.26642967E+9") {
			lang.SayString("xmul390")
		}
	}
	rexsult, err = lang.RxFromString("89100.1797").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("224.370309"))
	if err != nil {
		lang.SayString("xmul391")
	} else {
		if !(rexsult.ToString() == "19991434.9") {
			lang.SayString("xmul391")
		}
	}
	rexsult, err = lang.RxFromString("-821377.777").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("38.552821"))
	if err != nil {
		lang.SayString("xmul392")
	} else {
		if !(rexsult.ToString() == "-31666430.4") {
			lang.SayString("xmul392")
		}
	}
	rexsult, err = lang.RxFromString("-392640.782").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2571619.5E+113237865"))
	if err != nil {
		lang.SayString("xmul393")
	} else {
		if !(rexsult.ToString() == "1.00972269E+113237877") {
			lang.SayString("xmul393")
		}
	}
	rexsult, err = lang.RxFromString("-651397.712").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-723.59673"))
	if err != nil {
		lang.SayString("xmul394")
	} else {
		if !(rexsult.ToString() == "471349254") {
			lang.SayString("xmul394")
		}
	}
	rexsult, err = lang.RxFromString("86.6890892").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("940380864"))
	if err != nil {
		lang.SayString("xmul395")
	} else {
		if !(rexsult.ToString() == "8.15207606E+10") {
			lang.SayString("xmul395")
		}
	}
	rexsult, err = lang.RxFromString("173398265E-532383158").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3462.51450E+80986915"))
	if err != nil {
		lang.SayString("xmul397")
	} else {
		if !(rexsult.ToString() == "6.00394007E-451396232") {
			lang.SayString("xmul397")
		}
	}
	rexsult, err = lang.RxFromString("-1522176.78").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-6631061.77"))
	if err != nil {
		lang.SayString("xmul398")
	} else {
		if !(rexsult.ToString() == "1.00936483E+13") {
			lang.SayString("xmul398")
		}
	}
	rexsult, err = lang.RxFromString("538.10453").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("522934310"))
	if err != nil {
		lang.SayString("xmul399")
	} else {
		if !(rexsult.ToString() == "2.81393321E+11") {
			lang.SayString("xmul399")
		}
	}
	rexsult, err = lang.RxFromString("880243.444E-750940977").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-354.601578E-204943740"))
	if err != nil {
		lang.SayString("xmul400")
	} else {
		if !(rexsult.ToString() == "-3.12135714E-955884709") {
			lang.SayString("xmul400")
		}
	}
	rexsult, err = lang.RxFromString("968370.780").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("6677268.73"))
	if err != nil {
		lang.SayString("xmul401")
	} else {
		if !(rexsult.ToString() == "6.46607193E+12") {
			lang.SayString("xmul401")
		}
	}
	rexsult, err = lang.RxFromString("-97.7474945").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("31248241.5"))
	if err != nil {
		lang.SayString("xmul402")
	} else {
		if !(rexsult.ToString() == "-3.05443731E+9") {
			lang.SayString("xmul402")
		}
	}
	rexsult, err = lang.RxFromString("-328026144.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-125499735."))
	if err != nil {
		lang.SayString("xmul404")
	} else {
		if !(rexsult.ToString() == "4.11671941E+16") {
			lang.SayString("xmul404")
		}
	}
	rexsult, err = lang.RxFromString("-99424150.2E-523662102").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("3712.35030"))
	if err != nil {
		lang.SayString("xmul405")
	} else {
		if !(rexsult.ToString() == "-3.69097274E-523662091") {
			lang.SayString("xmul405")
		}
	}
	rexsult, err = lang.RxFromString("14838.0718").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9489893.28E+830631266"))
	if err != nil {
		lang.SayString("xmul406")
	} else {
		if !(rexsult.ToString() == "1.40811718E+830631277") {
			lang.SayString("xmul406")
		}
	}
	rexsult, err = lang.RxFromString("71207472.8").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-31835.0809"))
	if err != nil {
		lang.SayString("xmul407")
	} else {
		if !(rexsult.ToString() == "-2.26689566E+12") {
			lang.SayString("xmul407")
		}
	}
	rexsult, err = lang.RxFromString("-20440.4394").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-44.4064328E+511085806"))
	if err != nil {
		lang.SayString("xmul408")
	} else {
		if !(rexsult.ToString() == "9.07686999E+511085811") {
			lang.SayString("xmul408")
		}
	}
	rexsult, err = lang.RxFromString("-657.186702").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("426844.39"))
	if err != nil {
		lang.SayString("xmul411")
	} else {
		if !(rexsult.ToString() == "-280516457") {
			lang.SayString("xmul411")
		}
	}
	rexsult, err = lang.RxFromString("-41593077.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-688607.564"))
	if err != nil {
		lang.SayString("xmul412")
	} else {
		if !(rexsult.ToString() == "2.86413074E+13") {
			lang.SayString("xmul412")
		}
	}
	rexsult, err = lang.RxFromString("-5786.38132").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("190556652.E+177045877"))
	if err != nil {
		lang.SayString("xmul413")
	} else {
		if !(rexsult.ToString() == "-1.10263345E+177045889") {
			lang.SayString("xmul413")
		}
	}
	rexsult, err = lang.RxFromString("737622.974").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-241560693E+249506565"))
	if err != nil {
		lang.SayString("xmul414")
	} else {
		if !(rexsult.ToString() == "-1.78180717E+249506579") {
			lang.SayString("xmul414")
		}
	}
	rexsult, err = lang.RxFromString("5615373.52").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-7.27583808E-949781048"))
	if err != nil {
		lang.SayString("xmul415")
	} else {
		if !(rexsult.ToString() == "-4.08565485E-949781041") {
			lang.SayString("xmul415")
		}
	}
	rexsult, err = lang.RxFromString("644136.179").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-835708.103"))
	if err != nil {
		lang.SayString("xmul416")
	} else {
		if !(rexsult.ToString() == "-5.38309824E+11") {
			lang.SayString("xmul416")
		}
	}
	rexsult, err = lang.RxFromString("-307.419521E+466861843").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-738689976.E-199032711"))
	if err != nil {
		lang.SayString("xmul417")
	} else {
		if !(rexsult.ToString() == "2.27087719E+267829143") {
			lang.SayString("xmul417")
		}
	}
	rexsult, err = lang.RxFromString("-619642.130").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-226740537.E-902590153"))
	if err != nil {
		lang.SayString("xmul418")
	} else {
		if !(rexsult.ToString() == "1.40497989E-902590139") {
			lang.SayString("xmul418")
		}
	}
	rexsult, err = lang.RxFromString("-31068.7549").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-3.41495042E+86001379"))
	if err != nil {
		lang.SayString("xmul419")
	} else {
		if !(rexsult.ToString() == "1.06098258E+86001384") {
			lang.SayString("xmul419")
		}
	}
	rexsult, err = lang.RxFromString("-68951173.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-211804977.E-97318126"))
	if err != nil {
		lang.SayString("xmul420")
	} else {
		if !(rexsult.ToString() == "1.46042016E-97318110") {
			lang.SayString("xmul420")
		}
	}
	rexsult, err = lang.RxFromString("3898.03188").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-82572.615"))
	if err != nil {
		lang.SayString("xmul422")
	} else {
		if !(rexsult.ToString() == "-321870686") {
			lang.SayString("xmul422")
		}
	}
	rexsult, err = lang.RxFromString("-1.7619356").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-2546.64043"))
	if err != nil {
		lang.SayString("xmul423")
	} else {
		if !(rexsult.ToString() == "4487.01643") {
			lang.SayString("xmul423")
		}
	}
	rexsult, err = lang.RxFromString("59714.1968").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("29734388.6E-564525525"))
	if err != nil {
		lang.SayString("xmul424")
	} else {
		if !(rexsult.ToString() == "1.77556513E-564525513") {
			lang.SayString("xmul424")
		}
	}
	rexsult, err = lang.RxFromString("975566251").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-519.858530"))
	if err != nil {
		lang.SayString("xmul426")
	} else {
		if !(rexsult.ToString() == "-5.07156437E+11") {
			lang.SayString("xmul426")
		}
	}
	rexsult, err = lang.RxFromString("307401954").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-231481582."))
	if err != nil {
		lang.SayString("xmul427")
	} else {
		if !(rexsult.ToString() == "-7.11578906E+16") {
			lang.SayString("xmul427")
		}
	}
	rexsult, err = lang.RxFromString("2237645.48E+992947388").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-60618055.3E-857316706"))
	if err != nil {
		lang.SayString("xmul428")
	} else {
		if !(rexsult.ToString() == "-1.35641717E+135630696") {
			lang.SayString("xmul428")
		}
	}
	rexsult, err = lang.RxFromString("-403903.851").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("35.5049687E-72095155"))
	if err != nil {
		lang.SayString("xmul429")
	} else {
		if !(rexsult.ToString() == "-1.43405936E-72095148") {
			lang.SayString("xmul429")
		}
	}
	rexsult, err = lang.RxFromString("6.48674979").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-621732.532E+422575800"))
	if err != nil {
		lang.SayString("xmul430")
	} else {
		if !(rexsult.ToString() == "-4.03302337E+422575806") {
			lang.SayString("xmul430")
		}
	}
	rexsult, err = lang.RxFromString("-31401.9418").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("36.3960679"))
	if err != nil {
		lang.SayString("xmul431")
	} else {
		if !(rexsult.ToString() == "-1142907.21") {
			lang.SayString("xmul431")
		}
	}
	rexsult, err = lang.RxFromString("31345321.1").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("51.5482191"))
	if err != nil {
		lang.SayString("xmul432")
	} else {
		if !(rexsult.ToString() == "1.61579548E+9") {
			lang.SayString("xmul432")
		}
	}
	rexsult, err = lang.RxFromString("-64.172844").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-28506227.2E-767965800"))
	if err != nil {
		lang.SayString("xmul433")
	} else {
		if !(rexsult.ToString() == "1.82932567E-767965791") {
			lang.SayString("xmul433")
		}
	}
	rexsult, err = lang.RxFromString("70437.1551").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-62916.1233"))
	if err != nil {
		lang.SayString("xmul434")
	} else {
		if !(rexsult.ToString() == "-4.43163274E+9") {
			lang.SayString("xmul434")
		}
	}
	rexsult, err = lang.RxFromString("916260164").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-58.4017325"))
	if err != nil {
		lang.SayString("xmul435")
	} else {
		if !(rexsult.ToString() == "-5.35111810E+10") {
			lang.SayString("xmul435")
		}
	}
	rexsult, err = lang.RxFromString("19889085.3E-46816480").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1581683.94"))
	if err != nil {
		lang.SayString("xmul436")
	} else {
		if !(rexsult.ToString() == "3.14582468E-46816467") {
			lang.SayString("xmul436")
		}
	}
	rexsult, err = lang.RxFromString("-56312.3383").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("789.466064"))
	if err != nil {
		lang.SayString("xmul437")
	} else {
		if !(rexsult.ToString() == "-44456680.1") {
			lang.SayString("xmul437")
		}
	}
	rexsult, err = lang.RxFromString("183442.849").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-925876106"))
	if err != nil {
		lang.SayString("xmul438")
	} else {
		if !(rexsult.ToString() == "-1.69845351E+14") {
			lang.SayString("xmul438")
		}
	}
	rexsult, err = lang.RxFromString("859658551.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("72338.2054"))
	if err != nil {
		lang.SayString("xmul440")
	} else {
		if !(rexsult.ToString() == "6.21861568E+13") {
			lang.SayString("xmul440")
		}
	}
	rexsult, err = lang.RxFromString("-3.86446630E+426816068").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-664.534737"))
	if err != nil {
		lang.SayString("xmul441")
	} else {
		if !(rexsult.ToString() == "2.56807210E+426816071") {
			lang.SayString("xmul441")
		}
	}
	rexsult, err = lang.RxFromString("-969.881818").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("31170.8555"))
	if err != nil {
		lang.SayString("xmul442")
	} else {
		if !(rexsult.ToString() == "-30232046.0") {
			lang.SayString("xmul442")
		}
	}
	rexsult, err = lang.RxFromString("7980537.27").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("85.4040512"))
	if err != nil {
		lang.SayString("xmul443")
	} else {
		if !(rexsult.ToString() == "681570214") {
			lang.SayString("xmul443")
		}
	}
	rexsult, err = lang.RxFromString("-114609916.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("7525.14981"))
	if err != nil {
		lang.SayString("xmul444")
	} else {
		if !(rexsult.ToString() == "-8.62456788E+11") {
			lang.SayString("xmul444")
		}
	}
	rexsult, err = lang.RxFromString("8.43404682E-500572568").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("474526719"))
	if err != nil {
		lang.SayString("xmul445")
	} else {
		if !(rexsult.ToString() == "4.00218057E-500572559") {
			lang.SayString("xmul445")
		}
	}
	rexsult, err = lang.RxFromString("188006433").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2260.17037E-978192525"))
	if err != nil {
		lang.SayString("xmul446")
	} else {
		if !(rexsult.ToString() == "4.24926569E-978192514") {
			lang.SayString("xmul446")
		}
	}
	rexsult, err = lang.RxFromString("-9.95836312").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-866466703"))
	if err != nil {
		lang.SayString("xmul447")
	} else {
		if !(rexsult.ToString() == "8.62859006E+9") {
			lang.SayString("xmul447")
		}
	}
	rexsult, err = lang.RxFromString("80919339.2E-967231586").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("219.824266"))
	if err != nil {
		lang.SayString("xmul448")
	} else {
		if !(rexsult.ToString() == "1.77880343E-967231576") {
			lang.SayString("xmul448")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-89827.5229"))
	if err != nil {
		lang.SayString("xmul449")
	} else {
		if !(rexsult.ToString() == "-1.43346262E+10") {
			lang.SayString("xmul449")
		}
	}
	rexsult, err = lang.RxFromString("-4.54000153").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("6966333.74"))
	if err != nil {
		lang.SayString("xmul450")
	} else {
		if !(rexsult.ToString() == "-31627165.8") {
			lang.SayString("xmul450")
		}
	}
	rexsult, err = lang.RxFromString("28701538.7E-391015649").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-920999192."))
	if err != nil {
		lang.SayString("xmul451")
	} else {
		if !(rexsult.ToString() == "-2.64340940E-391015633") {
			lang.SayString("xmul451")
		}
	}
	rexsult, err = lang.RxFromString("-361382575.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-7976.15286E+898491169"))
	if err != nil {
		lang.SayString("xmul452")
	} else {
		if !(rexsult.ToString() == "2.88244266E+898491181") {
			lang.SayString("xmul452")
		}
	}
	rexsult, err = lang.RxFromString("7021805.61").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1222952.83"))
	if err != nil {
		lang.SayString("xmul453")
	} else {
		if !(rexsult.ToString() == "8.58733704E+12") {
			lang.SayString("xmul453")
		}
	}
	rexsult, err = lang.RxFromString("-40.4811667").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-79655.5635"))
	if err != nil {
		lang.SayString("xmul454")
	} else {
		if !(rexsult.ToString() == "3224550.14") {
			lang.SayString("xmul454")
		}
	}
	rexsult, err = lang.RxFromString("-8755674.38E+117168177").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("148.903404"))
	if err != nil {
		lang.SayString("xmul455")
	} else {
		if !(rexsult.ToString() == "-1.30374972E+117168186") {
			lang.SayString("xmul455")
		}
	}
	rexsult, err = lang.RxFromString("34.5329781E+382829392").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-45.2177309"))
	if err != nil {
		lang.SayString("xmul456")
	} else {
		if !(rexsult.ToString() == "-1.56150291E+382829395") {
			lang.SayString("xmul456")
		}
	}
	rexsult, err = lang.RxFromString("-37958476.0").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("584367.935"))
	if err != nil {
		lang.SayString("xmul457")
	} else {
		if !(rexsult.ToString() == "-2.21817162E+13") {
			lang.SayString("xmul457")
		}
	}
	rexsult, err = lang.RxFromString("495233.553E-414152215").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("62352759.2"))
	if err != nil {
		lang.SayString("xmul458")
	} else {
		if !(rexsult.ToString() == "3.08791785E-414152202") {
			lang.SayString("xmul458")
		}
	}
	rexsult, err = lang.RxFromString("-502343060").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-96828.994"))
	if err != nil {
		lang.SayString("xmul459")
	} else {
		if !(rexsult.ToString() == "4.86413731E+13") {
			lang.SayString("xmul459")
		}
	}
	rexsult, err = lang.RxFromString("-22.439639E+916362878").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-39.4037681"))
	if err != nil {
		lang.SayString("xmul460")
	} else {
		if !(rexsult.ToString() == "8.84206331E+916362880") {
			lang.SayString("xmul460")
		}
	}
	rexsult, err = lang.RxFromString("718180.587E-957473722").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1.66223443"))
	if err != nil {
		lang.SayString("xmul461")
	} else {
		if !(rexsult.ToString() == "1.19378450E-957473716") {
			lang.SayString("xmul461")
		}
	}
	rexsult, err = lang.RxFromString("-51592.2698").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-713885.741"))
	if err != nil {
		lang.SayString("xmul462")
	} else {
		if !(rexsult.ToString() == "3.68309858E+10") {
			lang.SayString("xmul462")
		}
	}
	rexsult, err = lang.RxFromString("51.2279848E+80439745").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("207.55925E+865165070"))
	if err != nil {
		lang.SayString("xmul463")
	} else {
		if !(rexsult.ToString() == "1.06328421E+945604819") {
			lang.SayString("xmul463")
		}
	}
	rexsult, err = lang.RxFromString("-5983.23468").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-39.9544513"))
	if err != nil {
		lang.SayString("xmul464")
	} else {
		if !(rexsult.ToString() == "239056.859") {
			lang.SayString("xmul464")
		}
	}
	rexsult, err = lang.RxFromString("921639332.E-917542963").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("287325.891"))
	if err != nil {
		lang.SayString("xmul465")
	} else {
		if !(rexsult.ToString() == "2.64810842E-917542949") {
			lang.SayString("xmul465")
		}
	}
	rexsult, err = lang.RxFromString("91095916.8E-787312969").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-58643.418E+58189880"))
	if err != nil {
		lang.SayString("xmul466")
	} else {
		if !(rexsult.ToString() == "-5.34217593E-729123077") {
			lang.SayString("xmul466")
		}
	}
	rexsult, err = lang.RxFromString("-6410.5555").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-234964259"))
	if err != nil {
		lang.SayString("xmul467")
	} else {
		if !(rexsult.ToString() == "1.50625142E+12") {
			lang.SayString("xmul467")
		}
	}
	rexsult, err = lang.RxFromString("-5.32711606").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-8447286.21"))
	if err != nil {
		lang.SayString("xmul468")
	} else {
		if !(rexsult.ToString() == "44999674.0") {
			lang.SayString("xmul468")
		}
	}
	rexsult, err = lang.RxFromString("-82272171.8").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-776.238587E-372690416"))
	if err != nil {
		lang.SayString("xmul469")
	} else {
		if !(rexsult.ToString() == "6.38628344E-372690406") {
			lang.SayString("xmul469")
		}
	}
	rexsult, err = lang.RxFromString("412411244.E-774339264").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("866452.465"))
	if err != nil {
		lang.SayString("xmul470")
	} else {
		if !(rexsult.ToString() == "3.57334739E-774339250") {
			lang.SayString("xmul470")
		}
	}
	rexsult, err = lang.RxFromString("-103.474598").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-3.01660661E-446661257"))
	if err != nil {
		lang.SayString("xmul471")
	} else {
		if !(rexsult.ToString() == "3.12142156E-446661255") {
			lang.SayString("xmul471")
		}
	}
	rexsult, err = lang.RxFromString("-31027.8323").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-475378186."))
	if err != nil {
		lang.SayString("xmul472")
	} else {
		if !(rexsult.ToString() == "1.47499546E+13") {
			lang.SayString("xmul472")
		}
	}
	rexsult, err = lang.RxFromString("-1199339.72").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-5.73068392E+53774632"))
	if err != nil {
		lang.SayString("xmul473")
	} else {
		if !(rexsult.ToString() == "6.87303685E+53774638") {
			lang.SayString("xmul473")
		}
	}
	rexsult, err = lang.RxFromString("-732908.930E+364345433").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-3486146.26"))
	if err != nil {
		lang.SayString("xmul474")
	} else {
		if !(rexsult.ToString() == "2.55502773E+364345445") {
			lang.SayString("xmul474")
		}
	}
	rexsult, err = lang.RxFromString("-2376150.83").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-46777583.3"))
	if err != nil {
		lang.SayString("xmul475")
	} else {
		if !(rexsult.ToString() == "1.11150593E+14") {
			lang.SayString("xmul475")
		}
	}
	rexsult, err = lang.RxFromString("6.3664211").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-140854908."))
	if err != nil {
		lang.SayString("xmul476")
	} else {
		if !(rexsult.ToString() == "-896741658") {
			lang.SayString("xmul476")
		}
	}
	rexsult, err = lang.RxFromString("-15.791522").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("1902.30210E+90741844"))
	if err != nil {
		lang.SayString("xmul477")
	} else {
		if !(rexsult.ToString() == "-3.00402455E+90741848") {
			lang.SayString("xmul477")
		}
	}
	rexsult, err = lang.RxFromString("15356.1505E+373950429").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("2.88020400"))
	if err != nil {
		lang.SayString("xmul478")
	} else {
		if !(rexsult.ToString() == "4.42288461E+373950433") {
			lang.SayString("xmul478")
		}
	}
	rexsult, err = lang.RxFromString("-3.12001326E+318884762").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("9567.21595"))
	if err != nil {
		lang.SayString("xmul479")
	} else {
		if !(rexsult.ToString() == "-2.98498406E+318884766") {
			lang.SayString("xmul479")
		}
	}
	rexsult, err = lang.RxFromString("49436.6528").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("751.919517"))
	if err != nil {
		lang.SayString("xmul480")
	} else {
		if !(rexsult.ToString() == "37172384.1") {
			lang.SayString("xmul480")
		}
	}
	rexsult, err = lang.RxFromString("552.669453").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("8.3725760E+16223526"))
	if err != nil {
		lang.SayString("xmul481")
	} else {
		if !(rexsult.ToString() == "4.62726700E+16223529") {
			lang.SayString("xmul481")
		}
	}
	rexsult, err = lang.RxFromString("-3266303").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("453741.520"))
	if err != nil {
		lang.SayString("xmul482")
	} else {
		if !(rexsult.ToString() == "-1.48205729E+12") {
			lang.SayString("xmul482")
		}
	}
	rexsult, err = lang.RxFromString("12302757.4").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("542922.487E+414443353"))
	if err != nil {
		lang.SayString("xmul483")
	} else {
		if !(rexsult.ToString() == "6.67944364E+414443365") {
			lang.SayString("xmul483")
		}
	}
	rexsult, err = lang.RxFromString("-5670757.79E-784754984").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("128144.503"))
	if err != nil {
		lang.SayString("xmul484")
	} else {
		if !(rexsult.ToString() == "-7.26676439E-784754973") {
			lang.SayString("xmul484")
		}
	}
	rexsult, err = lang.RxFromString("22.7721968E+842530698").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("5223.70462"))
	if err != nil {
		lang.SayString("xmul485")
	} else {
		if !(rexsult.ToString() == "1.18955230E+842530703") {
			lang.SayString("xmul485")
		}
	}
	rexsult, err = lang.RxFromString("88.5158199E-980164357").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("325846116"))
	if err != nil {
		lang.SayString("xmul486")
	} else {
		if !(rexsult.ToString() == "2.88425361E-980164347") {
			lang.SayString("xmul486")
		}
	}
	rexsult, err = lang.RxFromString("-22881.0408").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("5.63661562"))
	if err != nil {
		lang.SayString("xmul487")
	} else {
		if !(rexsult.ToString() == "-128971.632") {
			lang.SayString("xmul487")
		}
	}
	rexsult, err = lang.RxFromString("-7157.57449").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-76.4455519E-85647047"))
	if err != nil {
		lang.SayString("xmul488")
	} else {
		if !(rexsult.ToString() == "5.47164732E-85647042") {
			lang.SayString("xmul488")
		}
	}
	rexsult, err = lang.RxFromString("-503113.801").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-9715149.82E-612184422"))
	if err != nil {
		lang.SayString("xmul489")
	} else {
		if !(rexsult.ToString() == "4.88782595E-612184410") {
			lang.SayString("xmul489")
		}
	}
	rexsult, err = lang.RxFromString("-3066962.41").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-55.3096879"))
	if err != nil {
		lang.SayString("xmul490")
	} else {
		if !(rexsult.ToString() == "169632734") {
			lang.SayString("xmul490")
		}
	}
	rexsult, err = lang.RxFromString("-53311.5738E+156608936").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-7.45890666"))
	if err != nil {
		lang.SayString("xmul491")
	} else {
		if !(rexsult.ToString() == "3.97646053E+156608941") {
			lang.SayString("xmul491")
		}
	}
	rexsult, err = lang.RxFromString("998890068.").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-92.057879"))
	if err != nil {
		lang.SayString("xmul492")
	} else {
		if !(rexsult.ToString() == "-9.19557010E+10") {
			lang.SayString("xmul492")
		}
	}
	rexsult, err = lang.RxFromString("122.495591").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-407836028."))
	if err != nil {
		lang.SayString("xmul493")
	} else {
		if !(rexsult.ToString() == "-4.99581153E+10") {
			lang.SayString("xmul493")
		}
	}
	rexsult, err = lang.RxFromString("187098.488").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("6220.05584E-236541249"))
	if err != nil {
		lang.SayString("xmul494")
	} else {
		if !(rexsult.ToString() == "1.16376304E-236541240") {
			lang.SayString("xmul494")
		}
	}
	rexsult, err = lang.RxFromString("4819899.21E+432982550").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-727441917"))
	if err != nil {
		lang.SayString("xmul495")
	} else {
		if !(rexsult.ToString() == "-3.50619672E+432982565") {
			lang.SayString("xmul495")
		}
	}
	rexsult, err = lang.RxFromString("5770.01020E+507459752").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-4208339.33E-129766680"))
	if err != nil {
		lang.SayString("xmul496")
	} else {
		if !(rexsult.ToString() == "-2.42821609E+377693082") {
			lang.SayString("xmul496")
		}
	}
	rexsult, err = lang.RxFromString("-286.371320").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("710319152"))
	if err != nil {
		lang.SayString("xmul497")
	} else {
		if !(rexsult.ToString() == "-2.03415033E+11") {
			lang.SayString("xmul497")
		}
	}
	rexsult, err = lang.RxFromString("-7.27403536").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-481469656E-835183700"))
	if err != nil {
		lang.SayString("xmul498")
	} else {
		if !(rexsult.ToString() == "3.50222730E-835183691") {
			lang.SayString("xmul498")
		}
	}
	rexsult, err = lang.RxFromString("-6157.74292").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("-94075286.2E+92555877"))
	if err != nil {
		lang.SayString("xmul499")
	} else {
		if !(rexsult.ToString() == "5.79291428E+92555888") {
			lang.SayString("xmul499")
		}
	}
	rexsult, err = lang.RxFromString("-525445087.E+231529167").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("188227460"))
	if err != nil {
		lang.SayString("xmul500")
	} else {
		if !(rexsult.ToString() == "-9.89031941E+231529183") {
			lang.SayString("xmul500")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rmux101")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rmux101")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rmux102")
	} else {
		if !(rexsult.ToString() == "12346") {
			lang.SayString("rmux102")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rmux103")
	} else {
		if !(rexsult.ToString() == "12357") {
			lang.SayString("rmux103")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rmux104")
	} else {
		if !(rexsult.ToString() == "12468") {
			lang.SayString("rmux104")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rmux105")
	} else {
		if !(rexsult.ToString() == "13580") {
			lang.SayString("rmux105")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rmux106")
	} else {
		if !(rexsult.ToString() == "49380") {
			lang.SayString("rmux106")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rmux107")
	} else {
		if !(rexsult.ToString() == "49381") {
			lang.SayString("rmux107")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rmux108")
	} else {
		if !(rexsult.ToString() == "60491") {
			lang.SayString("rmux108")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rmux109")
	} else {
		if !(rexsult.ToString() == "61602") {
			lang.SayString("rmux109")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rmux110")
	} else {
		if !(rexsult.ToString() == "61713") {
			lang.SayString("rmux110")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rmux111")
	} else {
		if !(rexsult.ToString() == "61724") {
			lang.SayString("rmux111")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rmux112")
	} else {
		if !(rexsult.ToString() == "61725") {
			lang.SayString("rmux112")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rmux113")
	} else {
		if !(rexsult.ToString() == "61726") {
			lang.SayString("rmux113")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rmux114")
	} else {
		if !(rexsult.ToString() == "61737") {
			lang.SayString("rmux114")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rmux115")
	} else {
		if !(rexsult.ToString() == "61848") {
			lang.SayString("rmux115")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux116")
	} else {
		if !(rexsult.ToString() == "1.4814E+5") {
			lang.SayString("rmux116")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux117")
	} else {
		if !(rexsult.ToString() == "1.6049E+5") {
			lang.SayString("rmux117")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux118")
	} else {
		if !(rexsult.ToString() == "1.4826E+5") {
			lang.SayString("rmux118")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux119")
	} else {
		if !(rexsult.ToString() == "1.6062E+5") {
			lang.SayString("rmux119")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rmux201")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rmux201")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rmux202")
	} else {
		if !(rexsult.ToString() == "12346") {
			lang.SayString("rmux202")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rmux203")
	} else {
		if !(rexsult.ToString() == "12357") {
			lang.SayString("rmux203")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rmux204")
	} else {
		if !(rexsult.ToString() == "12468") {
			lang.SayString("rmux204")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rmux205")
	} else {
		if !(rexsult.ToString() == "13580") {
			lang.SayString("rmux205")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rmux206")
	} else {
		if !(rexsult.ToString() == "49380") {
			lang.SayString("rmux206")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rmux207")
	} else {
		if !(rexsult.ToString() == "49381") {
			lang.SayString("rmux207")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rmux208")
	} else {
		if !(rexsult.ToString() == "60491") {
			lang.SayString("rmux208")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rmux209")
	} else {
		if !(rexsult.ToString() == "61602") {
			lang.SayString("rmux209")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rmux210")
	} else {
		if !(rexsult.ToString() == "61713") {
			lang.SayString("rmux210")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rmux211")
	} else {
		if !(rexsult.ToString() == "61724") {
			lang.SayString("rmux211")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rmux212")
	} else {
		if !(rexsult.ToString() == "61725") {
			lang.SayString("rmux212")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rmux213")
	} else {
		if !(rexsult.ToString() == "61726") {
			lang.SayString("rmux213")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rmux214")
	} else {
		if !(rexsult.ToString() == "61737") {
			lang.SayString("rmux214")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rmux215")
	} else {
		if !(rexsult.ToString() == "61848") {
			lang.SayString("rmux215")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux216")
	} else {
		if !(rexsult.ToString() == "1.4814E+5") {
			lang.SayString("rmux216")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux217")
	} else {
		if !(rexsult.ToString() == "1.6049E+5") {
			lang.SayString("rmux217")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux218")
	} else {
		if !(rexsult.ToString() == "1.4826E+5") {
			lang.SayString("rmux218")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux219")
	} else {
		if !(rexsult.ToString() == "1.6062E+5") {
			lang.SayString("rmux219")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rmux301")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rmux301")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rmux302")
	} else {
		if !(rexsult.ToString() == "12346") {
			lang.SayString("rmux302")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rmux303")
	} else {
		if !(rexsult.ToString() == "12357") {
			lang.SayString("rmux303")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rmux304")
	} else {
		if !(rexsult.ToString() == "12468") {
			lang.SayString("rmux304")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rmux305")
	} else {
		if !(rexsult.ToString() == "13580") {
			lang.SayString("rmux305")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rmux306")
	} else {
		if !(rexsult.ToString() == "49380") {
			lang.SayString("rmux306")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rmux307")
	} else {
		if !(rexsult.ToString() == "49381") {
			lang.SayString("rmux307")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rmux308")
	} else {
		if !(rexsult.ToString() == "60491") {
			lang.SayString("rmux308")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rmux309")
	} else {
		if !(rexsult.ToString() == "61602") {
			lang.SayString("rmux309")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rmux310")
	} else {
		if !(rexsult.ToString() == "61713") {
			lang.SayString("rmux310")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rmux311")
	} else {
		if !(rexsult.ToString() == "61724") {
			lang.SayString("rmux311")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rmux312")
	} else {
		if !(rexsult.ToString() == "61725") {
			lang.SayString("rmux312")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rmux313")
	} else {
		if !(rexsult.ToString() == "61726") {
			lang.SayString("rmux313")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rmux314")
	} else {
		if !(rexsult.ToString() == "61737") {
			lang.SayString("rmux314")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rmux315")
	} else {
		if !(rexsult.ToString() == "61848") {
			lang.SayString("rmux315")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux316")
	} else {
		if !(rexsult.ToString() == "1.4814E+5") {
			lang.SayString("rmux316")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux317")
	} else {
		if !(rexsult.ToString() == "1.6049E+5") {
			lang.SayString("rmux317")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux318")
	} else {
		if !(rexsult.ToString() == "1.4826E+5") {
			lang.SayString("rmux318")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux319")
	} else {
		if !(rexsult.ToString() == "1.6062E+5") {
			lang.SayString("rmux319")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rmux401")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rmux401")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rmux402")
	} else {
		if !(rexsult.ToString() == "12346") {
			lang.SayString("rmux402")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rmux403")
	} else {
		if !(rexsult.ToString() == "12357") {
			lang.SayString("rmux403")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rmux404")
	} else {
		if !(rexsult.ToString() == "12468") {
			lang.SayString("rmux404")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rmux405")
	} else {
		if !(rexsult.ToString() == "13580") {
			lang.SayString("rmux405")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rmux406")
	} else {
		if !(rexsult.ToString() == "49380") {
			lang.SayString("rmux406")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rmux407")
	} else {
		if !(rexsult.ToString() == "49381") {
			lang.SayString("rmux407")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rmux408")
	} else {
		if !(rexsult.ToString() == "60491") {
			lang.SayString("rmux408")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rmux409")
	} else {
		if !(rexsult.ToString() == "61602") {
			lang.SayString("rmux409")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rmux410")
	} else {
		if !(rexsult.ToString() == "61713") {
			lang.SayString("rmux410")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rmux411")
	} else {
		if !(rexsult.ToString() == "61724") {
			lang.SayString("rmux411")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rmux412")
	} else {
		if !(rexsult.ToString() == "61725") {
			lang.SayString("rmux412")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rmux413")
	} else {
		if !(rexsult.ToString() == "61726") {
			lang.SayString("rmux413")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rmux414")
	} else {
		if !(rexsult.ToString() == "61737") {
			lang.SayString("rmux414")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rmux415")
	} else {
		if !(rexsult.ToString() == "61848") {
			lang.SayString("rmux415")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux416")
	} else {
		if !(rexsult.ToString() == "1.4814E+5") {
			lang.SayString("rmux416")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux417")
	} else {
		if !(rexsult.ToString() == "1.6049E+5") {
			lang.SayString("rmux417")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux418")
	} else {
		if !(rexsult.ToString() == "1.4826E+5") {
			lang.SayString("rmux418")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux419")
	} else {
		if !(rexsult.ToString() == "1.6062E+5") {
			lang.SayString("rmux419")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rmux501")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rmux501")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rmux502")
	} else {
		if !(rexsult.ToString() == "12346") {
			lang.SayString("rmux502")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rmux503")
	} else {
		if !(rexsult.ToString() == "12357") {
			lang.SayString("rmux503")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rmux504")
	} else {
		if !(rexsult.ToString() == "12468") {
			lang.SayString("rmux504")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rmux505")
	} else {
		if !(rexsult.ToString() == "13580") {
			lang.SayString("rmux505")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rmux506")
	} else {
		if !(rexsult.ToString() == "49380") {
			lang.SayString("rmux506")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rmux507")
	} else {
		if !(rexsult.ToString() == "49381") {
			lang.SayString("rmux507")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rmux508")
	} else {
		if !(rexsult.ToString() == "60491") {
			lang.SayString("rmux508")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rmux509")
	} else {
		if !(rexsult.ToString() == "61602") {
			lang.SayString("rmux509")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rmux510")
	} else {
		if !(rexsult.ToString() == "61713") {
			lang.SayString("rmux510")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rmux511")
	} else {
		if !(rexsult.ToString() == "61724") {
			lang.SayString("rmux511")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rmux512")
	} else {
		if !(rexsult.ToString() == "61725") {
			lang.SayString("rmux512")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rmux513")
	} else {
		if !(rexsult.ToString() == "61726") {
			lang.SayString("rmux513")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rmux514")
	} else {
		if !(rexsult.ToString() == "61737") {
			lang.SayString("rmux514")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rmux515")
	} else {
		if !(rexsult.ToString() == "61848") {
			lang.SayString("rmux515")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux516")
	} else {
		if !(rexsult.ToString() == "1.4814E+5") {
			lang.SayString("rmux516")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux517")
	} else {
		if !(rexsult.ToString() == "1.6049E+5") {
			lang.SayString("rmux517")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux518")
	} else {
		if !(rexsult.ToString() == "1.4826E+5") {
			lang.SayString("rmux518")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux519")
	} else {
		if !(rexsult.ToString() == "1.6062E+5") {
			lang.SayString("rmux519")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rmux601")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rmux601")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rmux602")
	} else {
		if !(rexsult.ToString() == "12346") {
			lang.SayString("rmux602")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rmux603")
	} else {
		if !(rexsult.ToString() == "12357") {
			lang.SayString("rmux603")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rmux604")
	} else {
		if !(rexsult.ToString() == "12468") {
			lang.SayString("rmux604")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rmux605")
	} else {
		if !(rexsult.ToString() == "13580") {
			lang.SayString("rmux605")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rmux606")
	} else {
		if !(rexsult.ToString() == "49380") {
			lang.SayString("rmux606")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rmux607")
	} else {
		if !(rexsult.ToString() == "49381") {
			lang.SayString("rmux607")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rmux608")
	} else {
		if !(rexsult.ToString() == "60491") {
			lang.SayString("rmux608")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rmux609")
	} else {
		if !(rexsult.ToString() == "61602") {
			lang.SayString("rmux609")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rmux610")
	} else {
		if !(rexsult.ToString() == "61713") {
			lang.SayString("rmux610")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rmux611")
	} else {
		if !(rexsult.ToString() == "61724") {
			lang.SayString("rmux611")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rmux612")
	} else {
		if !(rexsult.ToString() == "61725") {
			lang.SayString("rmux612")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rmux613")
	} else {
		if !(rexsult.ToString() == "61726") {
			lang.SayString("rmux613")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rmux614")
	} else {
		if !(rexsult.ToString() == "61737") {
			lang.SayString("rmux614")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rmux615")
	} else {
		if !(rexsult.ToString() == "61848") {
			lang.SayString("rmux615")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux616")
	} else {
		if !(rexsult.ToString() == "1.4814E+5") {
			lang.SayString("rmux616")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux617")
	} else {
		if !(rexsult.ToString() == "1.6049E+5") {
			lang.SayString("rmux617")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux618")
	} else {
		if !(rexsult.ToString() == "1.4826E+5") {
			lang.SayString("rmux618")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux619")
	} else {
		if !(rexsult.ToString() == "1.6062E+5") {
			lang.SayString("rmux619")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rmux701")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rmux701")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rmux702")
	} else {
		if !(rexsult.ToString() == "12346") {
			lang.SayString("rmux702")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rmux703")
	} else {
		if !(rexsult.ToString() == "12357") {
			lang.SayString("rmux703")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rmux704")
	} else {
		if !(rexsult.ToString() == "12468") {
			lang.SayString("rmux704")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rmux705")
	} else {
		if !(rexsult.ToString() == "13580") {
			lang.SayString("rmux705")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rmux706")
	} else {
		if !(rexsult.ToString() == "49380") {
			lang.SayString("rmux706")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rmux707")
	} else {
		if !(rexsult.ToString() == "49381") {
			lang.SayString("rmux707")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rmux708")
	} else {
		if !(rexsult.ToString() == "60491") {
			lang.SayString("rmux708")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rmux709")
	} else {
		if !(rexsult.ToString() == "61602") {
			lang.SayString("rmux709")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rmux710")
	} else {
		if !(rexsult.ToString() == "61713") {
			lang.SayString("rmux710")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rmux711")
	} else {
		if !(rexsult.ToString() == "61724") {
			lang.SayString("rmux711")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rmux712")
	} else {
		if !(rexsult.ToString() == "61725") {
			lang.SayString("rmux712")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rmux713")
	} else {
		if !(rexsult.ToString() == "61726") {
			lang.SayString("rmux713")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rmux714")
	} else {
		if !(rexsult.ToString() == "61737") {
			lang.SayString("rmux714")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rmux715")
	} else {
		if !(rexsult.ToString() == "61848") {
			lang.SayString("rmux715")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux716")
	} else {
		if !(rexsult.ToString() == "1.4814E+5") {
			lang.SayString("rmux716")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux717")
	} else {
		if !(rexsult.ToString() == "1.6049E+5") {
			lang.SayString("rmux717")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("rmux718")
	} else {
		if !(rexsult.ToString() == "1.4826E+5") {
			lang.SayString("rmux718")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("rmux719")
	} else {
		if !(rexsult.ToString() == "1.6062E+5") {
			lang.SayString("rmux719")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("r0mux101")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("r0mux101")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("r0mux102")
	} else {
		if !(rexsult.ToString() == "12346") {
			lang.SayString("r0mux102")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("r0mux103")
	} else {
		if !(rexsult.ToString() == "12357") {
			lang.SayString("r0mux103")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("r0mux104")
	} else {
		if !(rexsult.ToString() == "12468") {
			lang.SayString("r0mux104")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("r0mux105")
	} else {
		if !(rexsult.ToString() == "13580") {
			lang.SayString("r0mux105")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("r0mux106")
	} else {
		if !(rexsult.ToString() == "49380") {
			lang.SayString("r0mux106")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("r0mux107")
	} else {
		if !(rexsult.ToString() == "49381") {
			lang.SayString("r0mux107")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("r0mux108")
	} else {
		if !(rexsult.ToString() == "60491") {
			lang.SayString("r0mux108")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("r0mux109")
	} else {
		if !(rexsult.ToString() == "61602") {
			lang.SayString("r0mux109")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("r0mux110")
	} else {
		if !(rexsult.ToString() == "61713") {
			lang.SayString("r0mux110")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("r0mux111")
	} else {
		if !(rexsult.ToString() == "61724") {
			lang.SayString("r0mux111")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("r0mux112")
	} else {
		if !(rexsult.ToString() == "61725") {
			lang.SayString("r0mux112")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("r0mux113")
	} else {
		if !(rexsult.ToString() == "61726") {
			lang.SayString("r0mux113")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("r0mux114")
	} else {
		if !(rexsult.ToString() == "61737") {
			lang.SayString("r0mux114")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("r0mux115")
	} else {
		if !(rexsult.ToString() == "61848") {
			lang.SayString("r0mux115")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("r0mux116")
	} else {
		if !(rexsult.ToString() == "1.4814E+5") {
			lang.SayString("r0mux116")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("r0mux117")
	} else {
		if !(rexsult.ToString() == "1.6049E+5") {
			lang.SayString("r0mux117")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("r0mux118")
	} else {
		if !(rexsult.ToString() == "1.4826E+5") {
			lang.SayString("r0mux118")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpMult(lang.RxSetWithDigit(5), lang.RxFromString("13"))
	if err != nil {
		lang.SayString("r0mux119")
	} else {
		if !(rexsult.ToString() == "1.6062E+5") {
			lang.SayString("r0mux119")
		}
	}

	return
}
