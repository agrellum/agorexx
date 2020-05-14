package main

import "agorexx/lang"

func main() {

	rexsult, err := lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-384"))
	if err != nil {
		lang.SayString("addx574")
	} else {
		if !(rexsult.ToString() == "9E-384") {
			lang.SayString("addx574")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-398"))
	if err != nil {
		lang.SayString("addx575")
	} else {
		if !(rexsult.ToString() == "9.99999999999999E-384") {
			lang.SayString("addx575")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-398"))
	if err != nil {
		lang.SayString("addx576")
	} else {
		if !(rexsult.ToString() == "9.99999999999999E-384") {
			lang.SayString("addx576")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-399"))
	if err != nil {
		lang.SayString("addx577")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E-383") {
			lang.SayString("addx577")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-400"))
	if err != nil {
		lang.SayString("addx578")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E-383") {
			lang.SayString("addx578")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-401"))
	if err != nil {
		lang.SayString("addx579")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E-383") {
			lang.SayString("addx579")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-402"))
	if err != nil {
		lang.SayString("addx580")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E-383") {
			lang.SayString("addx580")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("120E-2"))
	if err != nil {
		lang.SayString("addx1705")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("addx1705")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("12E-1"))
	if err != nil {
		lang.SayString("addx1706")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("addx1706")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-384"))
	if err != nil {
		lang.SayString("addx6574")
	} else {
		if !(rexsult.ToString() == "9E-384") {
			lang.SayString("addx6574")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-398"))
	if err != nil {
		lang.SayString("addx6575")
	} else {
		if !(rexsult.ToString() == "9.99999999999999E-384") {
			lang.SayString("addx6575")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-398"))
	if err != nil {
		lang.SayString("addx6576")
	} else {
		if !(rexsult.ToString() == "9.99999999999999E-384") {
			lang.SayString("addx6576")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-399"))
	if err != nil {
		lang.SayString("addx6577")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E-383") {
			lang.SayString("addx6577")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-400"))
	if err != nil {
		lang.SayString("addx6578")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E-383") {
			lang.SayString("addx6578")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-401"))
	if err != nil {
		lang.SayString("addx6579")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E-383") {
			lang.SayString("addx6579")
		}
	}
	rexsult, err = lang.RxFromString("1E-383").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1E-402"))
	if err != nil {
		lang.SayString("addx6580")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E-383") {
			lang.SayString("addx6580")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("120E-2"))
	if err != nil {
		lang.SayString("addx61705")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("addx61705")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("12E-1"))
	if err != nil {
		lang.SayString("addx61706")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("addx61706")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub001")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub001")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddsub003")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddsub003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddsub004")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddsub005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub005")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddsub006")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddsub006")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddsub007")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddsub007")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub011")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub011")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub012")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("ddsub012")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddsub013")
	} else {
		if !(rexsult.ToString() == "-3") {
			lang.SayString("ddsub013")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub014")
	} else {
		if !(rexsult.ToString() == "-3") {
			lang.SayString("ddsub014")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddsub015")
	} else {
		if !(rexsult.ToString() == "-4") {
			lang.SayString("ddsub015")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("ddsub016")
	} else {
		if !(rexsult.ToString() == "-5") {
			lang.SayString("ddsub016")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddsub017")
	} else {
		if !(rexsult.ToString() == "-5") {
			lang.SayString("ddsub017")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("ddsub021")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub021")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddsub022")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("ddsub022")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("ddsub023")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddsub023")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("ddsub024")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("ddsub024")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("ddsub025")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("ddsub025")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("ddsub026")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("ddsub026")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("ddsub027")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("ddsub027")
		}
	}
	rexsult, err = lang.RxFromString("11").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub030")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("ddsub030")
		}
	}
	rexsult, err = lang.RxFromString("10").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub031")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("ddsub031")
		}
	}
	rexsult, err = lang.RxFromString("9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub032")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("ddsub032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub033")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub034")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddsub034")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub035")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("ddsub035")
		}
	}
	rexsult, err = lang.RxFromString("-9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub036")
	} else {
		if !(rexsult.ToString() == "-10") {
			lang.SayString("ddsub036")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub037")
	} else {
		if !(rexsult.ToString() == "-11") {
			lang.SayString("ddsub037")
		}
	}
	rexsult, err = lang.RxFromString("-11").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub038")
	} else {
		if !(rexsult.ToString() == "-12") {
			lang.SayString("ddsub038")
		}
	}
	rexsult, err = lang.RxFromString("5.75").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("3.3"))
	if err != nil {
		lang.SayString("ddsub040")
	} else {
		if !(rexsult.ToString() == "2.45") {
			lang.SayString("ddsub040")
		}
	}
	rexsult, err = lang.RxFromString("5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("ddsub041")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("ddsub041")
		}
	}
	rexsult, err = lang.RxFromString("-5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("ddsub042")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("ddsub042")
		}
	}
	rexsult, err = lang.RxFromString("-7").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("2.5"))
	if err != nil {
		lang.SayString("ddsub043")
	} else {
		if !(rexsult.ToString() == "-9.5") {
			lang.SayString("ddsub043")
		}
	}
	rexsult, err = lang.RxFromString("0.7").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddsub044")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddsub044")
		}
	}
	rexsult, err = lang.RxFromString("1.3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddsub045")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("ddsub045")
		}
	}
	rexsult, err = lang.RxFromString("1.25").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.25"))
	if err != nil {
		lang.SayString("ddsub046")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub046")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00000000"))
	if err != nil {
		lang.SayString("ddsub050")
	} else {
		if !(rexsult.ToString() == "0.23456789") {
			lang.SayString("ddsub050")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00000089"))
	if err != nil {
		lang.SayString("ddsub051")
	} else {
		if !(rexsult.ToString() == "0.23456700") {
			lang.SayString("ddsub051")
		}
	}
	rexsult, err = lang.RxFromString("70").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+16"))
	if err != nil {
		lang.SayString("ddsub060")
	} else {
		if !(rexsult.ToString() == "-1.000000000000000E+20") {
			lang.SayString("ddsub060")
		}
	}
	rexsult, err = lang.RxFromString("700").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+16"))
	if err != nil {
		lang.SayString("ddsub061")
	} else {
		if !(rexsult.ToString() == "-1.000000000000000E+20") {
			lang.SayString("ddsub061")
		}
	}
	rexsult, err = lang.RxFromString("7000").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+16"))
	if err != nil {
		lang.SayString("ddsub062")
	} else {
		if !(rexsult.ToString() == "-1.000000000000000E+20") {
			lang.SayString("ddsub062")
		}
	}
	rexsult, err = lang.RxFromString("70000").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+16"))
	if err != nil {
		lang.SayString("ddsub063")
	} else {
		if !(rexsult.ToString() == "-9.99999999999999E+19") {
			lang.SayString("ddsub063")
		}
	}
	rexsult, err = lang.RxFromString("700000").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+16"))
	if err != nil {
		lang.SayString("ddsub064")
	} else {
		if !(rexsult.ToString() == "-9.99999999999993E+19") {
			lang.SayString("ddsub064")
		}
	}
	rexsult, err = lang.RxFromString("10000e+16").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("70"))
	if err != nil {
		lang.SayString("ddsub065")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E+20") {
			lang.SayString("ddsub065")
		}
	}
	rexsult, err = lang.RxFromString("10000e+16").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("700"))
	if err != nil {
		lang.SayString("ddsub066")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E+20") {
			lang.SayString("ddsub066")
		}
	}
	rexsult, err = lang.RxFromString("10000e+16").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("7000"))
	if err != nil {
		lang.SayString("ddsub067")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E+20") {
			lang.SayString("ddsub067")
		}
	}
	rexsult, err = lang.RxFromString("10000e+16").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("ddsub068")
	} else {
		if !(rexsult.ToString() == "9.99999999999999E+19") {
			lang.SayString("ddsub068")
		}
	}
	rexsult, err = lang.RxFromString("10000e+16").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("700000"))
	if err != nil {
		lang.SayString("ddsub069")
	} else {
		if !(rexsult.ToString() == "9.99999999999993E+19") {
			lang.SayString("ddsub069")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("ddsub090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub090")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.00"))
	if err != nil {
		lang.SayString("ddsub091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("00.0"))
	if err != nil {
		lang.SayString("ddsub092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub092")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.00"))
	if err != nil {
		lang.SayString("ddsub093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub093")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("00.0"))
	if err != nil {
		lang.SayString("ddsub094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub094")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("ddsub095")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("ddsub095")
		}
	}
	rexsult, err = lang.RxFromString("3.").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("ddsub096")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("ddsub096")
		}
	}
	rexsult, err = lang.RxFromString("3.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("ddsub097")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("ddsub097")
		}
	}
	rexsult, err = lang.RxFromString("3.00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("ddsub098")
	} else {
		if !(rexsult.ToString() == "2.70") {
			lang.SayString("ddsub098")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("ddsub099")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub099")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("+3"))
	if err != nil {
		lang.SayString("ddsub100")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub100")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("ddsub101")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("ddsub101")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddsub102")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("ddsub102")
		}
	}
	rexsult, err = lang.RxFromString("3.").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddsub103")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("ddsub103")
		}
	}
	rexsult, err = lang.RxFromString("3.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddsub104")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("ddsub104")
		}
	}
	rexsult, err = lang.RxFromString("3.00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("ddsub105")
	} else {
		if !(rexsult.ToString() == "2.70") {
			lang.SayString("ddsub105")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("3.0"))
	if err != nil {
		lang.SayString("ddsub106")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub106")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("+3.0"))
	if err != nil {
		lang.SayString("ddsub107")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub107")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-3.0"))
	if err != nil {
		lang.SayString("ddsub108")
	} else {
		if !(rexsult.ToString() == "6.0") {
			lang.SayString("ddsub108")
		}
	}
	rexsult, err = lang.RxFromString("10.23456784").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub120")
	} else {
		if !(rexsult.ToString() == "-5E-8") {
			lang.SayString("ddsub120")
		}
	}
	rexsult, err = lang.RxFromString("10.23456785").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub121")
	} else {
		if !(rexsult.ToString() == "-4E-8") {
			lang.SayString("ddsub121")
		}
	}
	rexsult, err = lang.RxFromString("10.23456786").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub122")
	} else {
		if !(rexsult.ToString() == "-3E-8") {
			lang.SayString("ddsub122")
		}
	}
	rexsult, err = lang.RxFromString("10.23456787").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub123")
	} else {
		if !(rexsult.ToString() == "-2E-8") {
			lang.SayString("ddsub123")
		}
	}
	rexsult, err = lang.RxFromString("10.23456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub124")
	} else {
		if !(rexsult.ToString() == "-1E-8") {
			lang.SayString("ddsub124")
		}
	}
	rexsult, err = lang.RxFromString("10.23456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub125")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub125")
		}
	}
	rexsult, err = lang.RxFromString("10.23456790").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub126")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("ddsub126")
		}
	}
	rexsult, err = lang.RxFromString("10.23456791").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub127")
	} else {
		if !(rexsult.ToString() == "2E-8") {
			lang.SayString("ddsub127")
		}
	}
	rexsult, err = lang.RxFromString("10.23456792").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub128")
	} else {
		if !(rexsult.ToString() == "3E-8") {
			lang.SayString("ddsub128")
		}
	}
	rexsult, err = lang.RxFromString("10.23456793").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub129")
	} else {
		if !(rexsult.ToString() == "4E-8") {
			lang.SayString("ddsub129")
		}
	}
	rexsult, err = lang.RxFromString("10.23456794").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("ddsub130")
	} else {
		if !(rexsult.ToString() == "5E-8") {
			lang.SayString("ddsub130")
		}
	}
	rexsult, err = lang.RxFromString("10.23456781").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub131")
	} else {
		if !(rexsult.ToString() == "-5E-8") {
			lang.SayString("ddsub131")
		}
	}
	rexsult, err = lang.RxFromString("10.23456782").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub132")
	} else {
		if !(rexsult.ToString() == "-4E-8") {
			lang.SayString("ddsub132")
		}
	}
	rexsult, err = lang.RxFromString("10.23456783").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub133")
	} else {
		if !(rexsult.ToString() == "-3E-8") {
			lang.SayString("ddsub133")
		}
	}
	rexsult, err = lang.RxFromString("10.23456784").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub134")
	} else {
		if !(rexsult.ToString() == "-2E-8") {
			lang.SayString("ddsub134")
		}
	}
	rexsult, err = lang.RxFromString("10.23456785").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub135")
	} else {
		if !(rexsult.ToString() == "-1E-8") {
			lang.SayString("ddsub135")
		}
	}
	rexsult, err = lang.RxFromString("10.23456786").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub136")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("ddsub136")
		}
	}
	rexsult, err = lang.RxFromString("10.23456787").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub137")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("ddsub137")
		}
	}
	rexsult, err = lang.RxFromString("10.23456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub138")
	} else {
		if !(rexsult.ToString() == "2E-8") {
			lang.SayString("ddsub138")
		}
	}
	rexsult, err = lang.RxFromString("10.23456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub139")
	} else {
		if !(rexsult.ToString() == "3E-8") {
			lang.SayString("ddsub139")
		}
	}
	rexsult, err = lang.RxFromString("10.23456790").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub140")
	} else {
		if !(rexsult.ToString() == "4E-8") {
			lang.SayString("ddsub140")
		}
	}
	rexsult, err = lang.RxFromString("10.23456791").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("ddsub141")
	} else {
		if !(rexsult.ToString() == "5E-8") {
			lang.SayString("ddsub141")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("ddsub142")
	} else {
		if !(rexsult.ToString() == "1E-9") {
			lang.SayString("ddsub142")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub143")
	} else {
		if !(rexsult.ToString() == "-1E-9") {
			lang.SayString("ddsub143")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456780").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("ddsub144")
	} else {
		if !(rexsult.ToString() == "6E-8") {
			lang.SayString("ddsub144")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456790").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("ddsub145")
	} else {
		if !(rexsult.ToString() == "-4E-8") {
			lang.SayString("ddsub145")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456791").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("ddsub146")
	} else {
		if !(rexsult.ToString() == "-5E-8") {
			lang.SayString("ddsub146")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".1"))
	if err != nil {
		lang.SayString("ddsub160")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("ddsub160")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".97983"))
	if err != nil {
		lang.SayString("ddsub161")
	} else {
		if !(rexsult.ToString() == "-0.97983") {
			lang.SayString("ddsub161")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".9"))
	if err != nil {
		lang.SayString("ddsub162")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("ddsub162")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.102"))
	if err != nil {
		lang.SayString("ddsub163")
	} else {
		if !(rexsult.ToString() == "-0.102") {
			lang.SayString("ddsub163")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".4"))
	if err != nil {
		lang.SayString("ddsub164")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("ddsub164")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".307"))
	if err != nil {
		lang.SayString("ddsub165")
	} else {
		if !(rexsult.ToString() == "-0.307") {
			lang.SayString("ddsub165")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".43822"))
	if err != nil {
		lang.SayString("ddsub166")
	} else {
		if !(rexsult.ToString() == "-0.43822") {
			lang.SayString("ddsub166")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".911"))
	if err != nil {
		lang.SayString("ddsub167")
	} else {
		if !(rexsult.ToString() == "-0.911") {
			lang.SayString("ddsub167")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".02"))
	if err != nil {
		lang.SayString("ddsub168")
	} else {
		if !(rexsult.ToString() == "-0.02") {
			lang.SayString("ddsub168")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".392"))
	if err != nil {
		lang.SayString("ddsub169")
	} else {
		if !(rexsult.ToString() == "-0.392") {
			lang.SayString("ddsub169")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".26"))
	if err != nil {
		lang.SayString("ddsub170")
	} else {
		if !(rexsult.ToString() == "-0.26") {
			lang.SayString("ddsub170")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("ddsub171")
	} else {
		if !(rexsult.ToString() == "-0.51") {
			lang.SayString("ddsub171")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".2234"))
	if err != nil {
		lang.SayString("ddsub172")
	} else {
		if !(rexsult.ToString() == "-0.2234") {
			lang.SayString("ddsub172")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".2"))
	if err != nil {
		lang.SayString("ddsub173")
	} else {
		if !(rexsult.ToString() == "-0.2") {
			lang.SayString("ddsub173")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".0008"))
	if err != nil {
		lang.SayString("ddsub174")
	} else {
		if !(rexsult.ToString() == "-0.0008") {
			lang.SayString("ddsub174")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.1"))
	if err != nil {
		lang.SayString("ddsub180")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddsub180")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.97983"))
	if err != nil {
		lang.SayString("ddsub181")
	} else {
		if !(rexsult.ToString() == "0.97983") {
			lang.SayString("ddsub181")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.9"))
	if err != nil {
		lang.SayString("ddsub182")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("ddsub182")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.102"))
	if err != nil {
		lang.SayString("ddsub183")
	} else {
		if !(rexsult.ToString() == "0.102") {
			lang.SayString("ddsub183")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.4"))
	if err != nil {
		lang.SayString("ddsub184")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddsub184")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.307"))
	if err != nil {
		lang.SayString("ddsub185")
	} else {
		if !(rexsult.ToString() == "0.307") {
			lang.SayString("ddsub185")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.43822"))
	if err != nil {
		lang.SayString("ddsub186")
	} else {
		if !(rexsult.ToString() == "0.43822") {
			lang.SayString("ddsub186")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.911"))
	if err != nil {
		lang.SayString("ddsub187")
	} else {
		if !(rexsult.ToString() == "0.911") {
			lang.SayString("ddsub187")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.02"))
	if err != nil {
		lang.SayString("ddsub188")
	} else {
		if !(rexsult.ToString() == "0.02") {
			lang.SayString("ddsub188")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.392"))
	if err != nil {
		lang.SayString("ddsub189")
	} else {
		if !(rexsult.ToString() == "0.392") {
			lang.SayString("ddsub189")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.26"))
	if err != nil {
		lang.SayString("ddsub190")
	} else {
		if !(rexsult.ToString() == "0.26") {
			lang.SayString("ddsub190")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.51"))
	if err != nil {
		lang.SayString("ddsub191")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("ddsub191")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.2234"))
	if err != nil {
		lang.SayString("ddsub192")
	} else {
		if !(rexsult.ToString() == "0.2234") {
			lang.SayString("ddsub192")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.2"))
	if err != nil {
		lang.SayString("ddsub193")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("ddsub193")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.0008"))
	if err != nil {
		lang.SayString("ddsub194")
	} else {
		if !(rexsult.ToString() == "0.0008") {
			lang.SayString("ddsub194")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.1"))
	if err != nil {
		lang.SayString("ddsub200")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddsub200")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.97983"))
	if err != nil {
		lang.SayString("ddsub201")
	} else {
		if !(rexsult.ToString() == "0.97983") {
			lang.SayString("ddsub201")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.9"))
	if err != nil {
		lang.SayString("ddsub202")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("ddsub202")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.102"))
	if err != nil {
		lang.SayString("ddsub203")
	} else {
		if !(rexsult.ToString() == "0.102") {
			lang.SayString("ddsub203")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.4"))
	if err != nil {
		lang.SayString("ddsub204")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("ddsub204")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.307"))
	if err != nil {
		lang.SayString("ddsub205")
	} else {
		if !(rexsult.ToString() == "0.307") {
			lang.SayString("ddsub205")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.43822"))
	if err != nil {
		lang.SayString("ddsub206")
	} else {
		if !(rexsult.ToString() == "0.43822") {
			lang.SayString("ddsub206")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.911"))
	if err != nil {
		lang.SayString("ddsub207")
	} else {
		if !(rexsult.ToString() == "0.911") {
			lang.SayString("ddsub207")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.02"))
	if err != nil {
		lang.SayString("ddsub208")
	} else {
		if !(rexsult.ToString() == "0.02") {
			lang.SayString("ddsub208")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.392"))
	if err != nil {
		lang.SayString("ddsub209")
	} else {
		if !(rexsult.ToString() == "0.392") {
			lang.SayString("ddsub209")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.26"))
	if err != nil {
		lang.SayString("ddsub210")
	} else {
		if !(rexsult.ToString() == "0.26") {
			lang.SayString("ddsub210")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.51"))
	if err != nil {
		lang.SayString("ddsub211")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("ddsub211")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.2234"))
	if err != nil {
		lang.SayString("ddsub212")
	} else {
		if !(rexsult.ToString() == "0.2234") {
			lang.SayString("ddsub212")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.2"))
	if err != nil {
		lang.SayString("ddsub213")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("ddsub213")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-.0008"))
	if err != nil {
		lang.SayString("ddsub214")
	} else {
		if !(rexsult.ToString() == "0.0008") {
			lang.SayString("ddsub214")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-12").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub220")
	} else {
		if !(rexsult.ToString() == "-5.6267E-8") {
			lang.SayString("ddsub220")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-11").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub221")
	} else {
		if !(rexsult.ToString() == "-5.6267E-7") {
			lang.SayString("ddsub221")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-10").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub222")
	} else {
		if !(rexsult.ToString() == "-0.0000056267") {
			lang.SayString("ddsub222")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub223")
	} else {
		if !(rexsult.ToString() == "-0.000056267") {
			lang.SayString("ddsub223")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-8").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub224")
	} else {
		if !(rexsult.ToString() == "-0.00056267") {
			lang.SayString("ddsub224")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-7").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub225")
	} else {
		if !(rexsult.ToString() == "-0.0056267") {
			lang.SayString("ddsub225")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-6").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub226")
	} else {
		if !(rexsult.ToString() == "-0.056267") {
			lang.SayString("ddsub226")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub227")
	} else {
		if !(rexsult.ToString() == "-0.56267") {
			lang.SayString("ddsub227")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub228")
	} else {
		if !(rexsult.ToString() == "-562.67") {
			lang.SayString("ddsub228")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub229")
	} else {
		if !(rexsult.ToString() == "-5626.7") {
			lang.SayString("ddsub229")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub230")
	} else {
		if !(rexsult.ToString() == "-56267") {
			lang.SayString("ddsub230")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-12"))
	if err != nil {
		lang.SayString("ddsub240")
	} else {
		if !(rexsult.ToString() == "5.6267E-8") {
			lang.SayString("ddsub240")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-11"))
	if err != nil {
		lang.SayString("ddsub241")
	} else {
		if !(rexsult.ToString() == "5.6267E-7") {
			lang.SayString("ddsub241")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-10"))
	if err != nil {
		lang.SayString("ddsub242")
	} else {
		if !(rexsult.ToString() == "0.0000056267") {
			lang.SayString("ddsub242")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-9"))
	if err != nil {
		lang.SayString("ddsub243")
	} else {
		if !(rexsult.ToString() == "0.000056267") {
			lang.SayString("ddsub243")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-8"))
	if err != nil {
		lang.SayString("ddsub244")
	} else {
		if !(rexsult.ToString() == "0.00056267") {
			lang.SayString("ddsub244")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-7"))
	if err != nil {
		lang.SayString("ddsub245")
	} else {
		if !(rexsult.ToString() == "0.0056267") {
			lang.SayString("ddsub245")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-6"))
	if err != nil {
		lang.SayString("ddsub246")
	} else {
		if !(rexsult.ToString() == "0.056267") {
			lang.SayString("ddsub246")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-5"))
	if err != nil {
		lang.SayString("ddsub247")
	} else {
		if !(rexsult.ToString() == "0.56267") {
			lang.SayString("ddsub247")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-2"))
	if err != nil {
		lang.SayString("ddsub248")
	} else {
		if !(rexsult.ToString() == "562.67") {
			lang.SayString("ddsub248")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-1"))
	if err != nil {
		lang.SayString("ddsub249")
	} else {
		if !(rexsult.ToString() == "5626.7") {
			lang.SayString("ddsub249")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-0"))
	if err != nil {
		lang.SayString("ddsub250")
	} else {
		if !(rexsult.ToString() == "56267") {
			lang.SayString("ddsub250")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00000000"))
	if err != nil {
		lang.SayString("ddsub301")
	} else {
		if !(rexsult.ToString() == "0.23456789") {
			lang.SayString("ddsub301")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00000011"))
	if err != nil {
		lang.SayString("ddsub302")
	} else {
		if !(rexsult.ToString() == "0.23456778") {
			lang.SayString("ddsub302")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.0000"))
	if err != nil {
		lang.SayString("ddsub321")
	} else {
		if !(rexsult.ToString() == "0.9998") {
			lang.SayString("ddsub321")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.0001"))
	if err != nil {
		lang.SayString("ddsub322")
	} else {
		if !(rexsult.ToString() == "0.9997") {
			lang.SayString("ddsub322")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.0002"))
	if err != nil {
		lang.SayString("ddsub323")
	} else {
		if !(rexsult.ToString() == "0.9996") {
			lang.SayString("ddsub323")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.0003"))
	if err != nil {
		lang.SayString("ddsub324")
	} else {
		if !(rexsult.ToString() == "0.9995") {
			lang.SayString("ddsub324")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.0000"))
	if err != nil {
		lang.SayString("ddsub325")
	} else {
		if !(rexsult.ToString() == "0.9998") {
			lang.SayString("ddsub325")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.0001"))
	if err != nil {
		lang.SayString("ddsub326")
	} else {
		if !(rexsult.ToString() == "0.9999") {
			lang.SayString("ddsub326")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.0002"))
	if err != nil {
		lang.SayString("ddsub327")
	} else {
		if !(rexsult.ToString() == "1.0000") {
			lang.SayString("ddsub327")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.0003"))
	if err != nil {
		lang.SayString("ddsub328")
	} else {
		if !(rexsult.ToString() == "1.0001") {
			lang.SayString("ddsub328")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("ddsub346")
	} else {
		if !(rexsult.ToString() == "9999999999993") {
			lang.SayString("ddsub346")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("70"))
	if err != nil {
		lang.SayString("ddsub347")
	} else {
		if !(rexsult.ToString() == "9999999999930") {
			lang.SayString("ddsub347")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("700"))
	if err != nil {
		lang.SayString("ddsub348")
	} else {
		if !(rexsult.ToString() == "9999999999300") {
			lang.SayString("ddsub348")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("7000"))
	if err != nil {
		lang.SayString("ddsub349")
	} else {
		if !(rexsult.ToString() == "9999999993000") {
			lang.SayString("ddsub349")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("ddsub350")
	} else {
		if !(rexsult.ToString() == "9999999930000") {
			lang.SayString("ddsub350")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("700000"))
	if err != nil {
		lang.SayString("ddsub351")
	} else {
		if !(rexsult.ToString() == "9999999300000") {
			lang.SayString("ddsub351")
		}
	}
	rexsult, err = lang.RxFromString("7").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("ddsub352")
	} else {
		if !(rexsult.ToString() == "-9999999999993") {
			lang.SayString("ddsub352")
		}
	}
	rexsult, err = lang.RxFromString("70").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("ddsub353")
	} else {
		if !(rexsult.ToString() == "-9999999999930") {
			lang.SayString("ddsub353")
		}
	}
	rexsult, err = lang.RxFromString("700").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("ddsub354")
	} else {
		if !(rexsult.ToString() == "-9999999999300") {
			lang.SayString("ddsub354")
		}
	}
	rexsult, err = lang.RxFromString("7000").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("ddsub355")
	} else {
		if !(rexsult.ToString() == "-9999999993000") {
			lang.SayString("ddsub355")
		}
	}
	rexsult, err = lang.RxFromString("70000").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("ddsub356")
	} else {
		if !(rexsult.ToString() == "-9999999930000") {
			lang.SayString("ddsub356")
		}
	}
	rexsult, err = lang.RxFromString("700000").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("ddsub357")
	} else {
		if !(rexsult.ToString() == "-9999999300000") {
			lang.SayString("ddsub357")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.0001"))
	if err != nil {
		lang.SayString("ddsub361")
	} else {
		if !(rexsult.ToString() == "0.9999") {
			lang.SayString("ddsub361")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.00001"))
	if err != nil {
		lang.SayString("ddsub362")
	} else {
		if !(rexsult.ToString() == "0.99999") {
			lang.SayString("ddsub362")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("ddsub363")
	} else {
		if !(rexsult.ToString() == "0.999999") {
			lang.SayString("ddsub363")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.0000000000000001"))
	if err != nil {
		lang.SayString("ddsub364")
	} else {
		if !(rexsult.ToString() == "1.000000000000000") {
			lang.SayString("ddsub364")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.00000000000000001"))
	if err != nil {
		lang.SayString("ddsub365")
	} else {
		if !(rexsult.ToString() == "1.000000000000000") {
			lang.SayString("ddsub365")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.000000000000000001"))
	if err != nil {
		lang.SayString("ddsub366")
	} else {
		if !(rexsult.ToString() == "1.000000000000000") {
			lang.SayString("ddsub366")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub370")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddsub370")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0."))
	if err != nil {
		lang.SayString("ddsub371")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddsub371")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString(".0"))
	if err != nil {
		lang.SayString("ddsub372")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddsub372")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("ddsub373")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("ddsub373")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub374")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddsub374")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub375")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddsub375")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub376")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddsub376")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub377")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("ddsub377")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-51897955.3"))
	if err != nil {
		lang.SayString("ddsub910")
	} else {
		if !(rexsult.ToString() == "-51621406.7") {
			lang.SayString("ddsub910")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("89827.5229"))
	if err != nil {
		lang.SayString("ddsub911")
	} else {
		if !(rexsult.ToString() == "69751.9211") {
			lang.SayString("ddsub911")
		}
	}
	rexsult, err = lang.RxFromString("333.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("33.00000001234566"))
	if err != nil {
		lang.SayString("ddsub920")
	} else {
		if !(rexsult.ToString() == "299.9999999999999") {
			lang.SayString("ddsub920")
		}
	}
	rexsult, err = lang.RxFromString("333.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("33.00000001234565"))
	if err != nil {
		lang.SayString("ddsub921")
	} else {
		if !(rexsult.ToString() == "300.0000000000000") {
			lang.SayString("ddsub921")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("33.00000001234565"))
	if err != nil {
		lang.SayString("ddsub922")
	} else {
		if !(rexsult.ToString() == "100.0000000000000") {
			lang.SayString("ddsub922")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("33.00000001234564"))
	if err != nil {
		lang.SayString("ddsub923")
	} else {
		if !(rexsult.ToString() == "100.0000000000000") {
			lang.SayString("ddsub923")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("33.00000001234540"))
	if err != nil {
		lang.SayString("ddsub924")
	} else {
		if !(rexsult.ToString() == "100.0000000000002") {
			lang.SayString("ddsub924")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("43.00000001234560"))
	if err != nil {
		lang.SayString("ddsub925")
	} else {
		if !(rexsult.ToString() == "90.0000000000000") {
			lang.SayString("ddsub925")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("43.00000001234561"))
	if err != nil {
		lang.SayString("ddsub926")
	} else {
		if !(rexsult.ToString() == "90.0000000000000") {
			lang.SayString("ddsub926")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("43.00000001234566"))
	if err != nil {
		lang.SayString("ddsub927")
	} else {
		if !(rexsult.ToString() == "89.9999999999999") {
			lang.SayString("ddsub927")
		}
	}
	rexsult, err = lang.RxFromString("101.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("91.00000001234566"))
	if err != nil {
		lang.SayString("ddsub928")
	} else {
		if !(rexsult.ToString() == "9.9999999999999") {
			lang.SayString("ddsub928")
		}
	}
	rexsult, err = lang.RxFromString("101.0000000123456").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("99.00000001234566"))
	if err != nil {
		lang.SayString("ddsub929")
	} else {
		if !(rexsult.ToString() == "1.9999999999999") {
			lang.SayString("ddsub929")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-10").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub390")
	} else {
		if !(rexsult.ToString() == "-0.0000056267") {
			lang.SayString("ddsub390")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-6").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub391")
	} else {
		if !(rexsult.ToString() == "-0.056267") {
			lang.SayString("ddsub391")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub392")
	} else {
		if !(rexsult.ToString() == "-0.56267") {
			lang.SayString("ddsub392")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-4").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub393")
	} else {
		if !(rexsult.ToString() == "-5.6267") {
			lang.SayString("ddsub393")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub394")
	} else {
		if !(rexsult.ToString() == "-56.267") {
			lang.SayString("ddsub394")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-2").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub395")
	} else {
		if !(rexsult.ToString() == "-562.67") {
			lang.SayString("ddsub395")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub396")
	} else {
		if !(rexsult.ToString() == "-5626.7") {
			lang.SayString("ddsub396")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub397")
	} else {
		if !(rexsult.ToString() == "-56267") {
			lang.SayString("ddsub397")
		}
	}
	rexsult, err = lang.RxFromString("-5E-10").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub398")
	} else {
		if !(rexsult.ToString() == "-5E-10") {
			lang.SayString("ddsub398")
		}
	}
	rexsult, err = lang.RxFromString("-5E-7").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub399")
	} else {
		if !(rexsult.ToString() == "-5E-7") {
			lang.SayString("ddsub399")
		}
	}
	rexsult, err = lang.RxFromString("-5E-6").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub400")
	} else {
		if !(rexsult.ToString() == "-0.000005") {
			lang.SayString("ddsub400")
		}
	}
	rexsult, err = lang.RxFromString("-5E-5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub401")
	} else {
		if !(rexsult.ToString() == "-0.00005") {
			lang.SayString("ddsub401")
		}
	}
	rexsult, err = lang.RxFromString("-5E-4").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub402")
	} else {
		if !(rexsult.ToString() == "-0.0005") {
			lang.SayString("ddsub402")
		}
	}
	rexsult, err = lang.RxFromString("-5E-1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub403")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("ddsub403")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-10"))
	if err != nil {
		lang.SayString("ddsub420")
	} else {
		if !(rexsult.ToString() == "0.0000056267") {
			lang.SayString("ddsub420")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-6"))
	if err != nil {
		lang.SayString("ddsub421")
	} else {
		if !(rexsult.ToString() == "0.056267") {
			lang.SayString("ddsub421")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-5"))
	if err != nil {
		lang.SayString("ddsub422")
	} else {
		if !(rexsult.ToString() == "0.56267") {
			lang.SayString("ddsub422")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-4"))
	if err != nil {
		lang.SayString("ddsub423")
	} else {
		if !(rexsult.ToString() == "5.6267") {
			lang.SayString("ddsub423")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-3"))
	if err != nil {
		lang.SayString("ddsub424")
	} else {
		if !(rexsult.ToString() == "56.267") {
			lang.SayString("ddsub424")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-2"))
	if err != nil {
		lang.SayString("ddsub425")
	} else {
		if !(rexsult.ToString() == "562.67") {
			lang.SayString("ddsub425")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-1"))
	if err != nil {
		lang.SayString("ddsub426")
	} else {
		if !(rexsult.ToString() == "5626.7") {
			lang.SayString("ddsub426")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-56267E-0"))
	if err != nil {
		lang.SayString("ddsub427")
	} else {
		if !(rexsult.ToString() == "56267") {
			lang.SayString("ddsub427")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-5E-10"))
	if err != nil {
		lang.SayString("ddsub428")
	} else {
		if !(rexsult.ToString() == "5E-10") {
			lang.SayString("ddsub428")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-5E-7"))
	if err != nil {
		lang.SayString("ddsub429")
	} else {
		if !(rexsult.ToString() == "5E-7") {
			lang.SayString("ddsub429")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-5E-6"))
	if err != nil {
		lang.SayString("ddsub430")
	} else {
		if !(rexsult.ToString() == "0.000005") {
			lang.SayString("ddsub430")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-5E-5"))
	if err != nil {
		lang.SayString("ddsub431")
	} else {
		if !(rexsult.ToString() == "0.00005") {
			lang.SayString("ddsub431")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-5E-4"))
	if err != nil {
		lang.SayString("ddsub432")
	} else {
		if !(rexsult.ToString() == "0.0005") {
			lang.SayString("ddsub432")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-5E-1"))
	if err != nil {
		lang.SayString("ddsub433")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("ddsub433")
		}
	}
	rexsult, err = lang.RxFromString("1E+16").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub461")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E+16") {
			lang.SayString("ddsub461")
		}
	}
	rexsult, err = lang.RxFromString("1E+12").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-1.111"))
	if err != nil {
		lang.SayString("ddsub462")
	} else {
		if !(rexsult.ToString() == "1000000000001.111") {
			lang.SayString("ddsub462")
		}
	}
	rexsult, err = lang.RxFromString("1.111").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-1E+12"))
	if err != nil {
		lang.SayString("ddsub463")
	} else {
		if !(rexsult.ToString() == "1000000000001.111") {
			lang.SayString("ddsub463")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-1E+16"))
	if err != nil {
		lang.SayString("ddsub464")
	} else {
		if !(rexsult.ToString() == "1.000000000000000E+16") {
			lang.SayString("ddsub464")
		}
	}
	rexsult, err = lang.RxFromString("7E+15").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub465")
	} else {
		if !(rexsult.ToString() == "6999999999999999") {
			lang.SayString("ddsub465")
		}
	}
	rexsult, err = lang.RxFromString("7E+12").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-1.111"))
	if err != nil {
		lang.SayString("ddsub466")
	} else {
		if !(rexsult.ToString() == "7000000000001.111") {
			lang.SayString("ddsub466")
		}
	}
	rexsult, err = lang.RxFromString("1.111").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-7E+12"))
	if err != nil {
		lang.SayString("ddsub467")
	} else {
		if !(rexsult.ToString() == "7000000000001.111") {
			lang.SayString("ddsub467")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-7E+15"))
	if err != nil {
		lang.SayString("ddsub468")
	} else {
		if !(rexsult.ToString() == "6999999999999999") {
			lang.SayString("ddsub468")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555563"))
	if err != nil {
		lang.SayString("ddsub470")
	} else {
		if !(rexsult.ToString() == "1.000000000000001") {
			lang.SayString("ddsub470")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555562"))
	if err != nil {
		lang.SayString("ddsub471")
	} else {
		if !(rexsult.ToString() == "1.000000000000001") {
			lang.SayString("ddsub471")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555561"))
	if err != nil {
		lang.SayString("ddsub472")
	} else {
		if !(rexsult.ToString() == "1.000000000000001") {
			lang.SayString("ddsub472")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555560"))
	if err != nil {
		lang.SayString("ddsub473")
	} else {
		if !(rexsult.ToString() == "1.000000000000000") {
			lang.SayString("ddsub473")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555559"))
	if err != nil {
		lang.SayString("ddsub474")
	} else {
		if !(rexsult.ToString() == "1.000000000000000") {
			lang.SayString("ddsub474")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555558"))
	if err != nil {
		lang.SayString("ddsub475")
	} else {
		if !(rexsult.ToString() == "1.000000000000000") {
			lang.SayString("ddsub475")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555557"))
	if err != nil {
		lang.SayString("ddsub476")
	} else {
		if !(rexsult.ToString() == "1.000000000000000") {
			lang.SayString("ddsub476")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555556"))
	if err != nil {
		lang.SayString("ddsub477")
	} else {
		if !(rexsult.ToString() == "1.000000000000000") {
			lang.SayString("ddsub477")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555555"))
	if err != nil {
		lang.SayString("ddsub478")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999") {
			lang.SayString("ddsub478")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555554"))
	if err != nil {
		lang.SayString("ddsub479")
	} else {
		if !(rexsult.ToString() == "0.9999999999999998") {
			lang.SayString("ddsub479")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555553"))
	if err != nil {
		lang.SayString("ddsub480")
	} else {
		if !(rexsult.ToString() == "0.9999999999999997") {
			lang.SayString("ddsub480")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555552"))
	if err != nil {
		lang.SayString("ddsub481")
	} else {
		if !(rexsult.ToString() == "0.9999999999999996") {
			lang.SayString("ddsub481")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555551"))
	if err != nil {
		lang.SayString("ddsub482")
	} else {
		if !(rexsult.ToString() == "0.9999999999999995") {
			lang.SayString("ddsub482")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("-0.5555555555555550"))
	if err != nil {
		lang.SayString("ddsub483")
	} else {
		if !(rexsult.ToString() == "0.9999999999999994") {
			lang.SayString("ddsub483")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub500")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub500")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("ddsub501")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub501")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("ddsub502")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub502")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("ddsub503")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub503")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("ddsub504")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub504")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("ddsub505")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub505")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("ddsub506")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub506")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("ddsub507")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub507")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("ddsub508")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub508")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("ddsub509")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub509")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("ddsub510")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub510")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("ddsub511")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub511")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("ddsub512")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub512")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("ddsub513")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub513")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("ddsub514")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub514")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("ddsub515")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub515")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub516")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub516")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.000000001"))
	if err != nil {
		lang.SayString("ddsub517")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub517")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("ddsub518")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub518")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("ddsub519")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub519")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub520")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub520")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("ddsub521")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub521")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("ddsub522")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub522")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("ddsub523")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub523")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("ddsub524")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub524")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("ddsub525")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub525")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("ddsub526")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub526")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("ddsub527")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub527")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("ddsub528")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub528")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("ddsub529")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub529")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("ddsub530")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub530")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("ddsub531")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub531")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("ddsub532")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub532")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("ddsub533")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub533")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("ddsub534")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub534")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("ddsub535")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub535")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub536")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub536")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00000001"))
	if err != nil {
		lang.SayString("ddsub537")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub537")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("ddsub538")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub538")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("ddsub539")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub539")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("ddsub540")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub540")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("ddsub541")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub541")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("ddsub542")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub542")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("ddsub550")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub550")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("ddsub551")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub551")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("ddsub552")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub552")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("ddsub553")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub553")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("ddsub554")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub554")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("ddsub555")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub555")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("ddsub556")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub556")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("ddsub557")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub557")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("ddsub558")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub558")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("ddsub559")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub559")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("ddsub560")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub560")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("ddsub561")
	} else {
		if !(rexsult.ToString() == "1231234567456789") {
			lang.SayString("ddsub561")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("ddsub562")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub562")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("ddsub563")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub563")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("ddsub564")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub564")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("ddsub565")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub565")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("ddsub566")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub566")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00000001"))
	if err != nil {
		lang.SayString("ddsub567")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub567")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("ddsub568")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub568")
		}
	}
	rexsult, err = lang.RxFromString("1231234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("ddsub569")
	} else {
		if !(rexsult.ToString() == "1231234567456788") {
			lang.SayString("ddsub569")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub600")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub600")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub601")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub601")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub602")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub602")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub603")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub603")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub604")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub604")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub605")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub605")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub606")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub606")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub607")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub607")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub608")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub608")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub609")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub609")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub610")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub610")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub611")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub611")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub612")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub612")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub613")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub613")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub614")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub614")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub615")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub615")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub616")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub616")
		}
	}
	rexsult, err = lang.RxFromString("1.000000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub617")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub617")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub618")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub618")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub619")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub619")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub620")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub620")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub621")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub621")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub622")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub622")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub623")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub623")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub624")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub624")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub625")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub625")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub626")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub626")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub627")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub627")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub628")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub628")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub629")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub629")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub630")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub630")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub631")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub631")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub632")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub632")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub633")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub633")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub634")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub634")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub635")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub635")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub636")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub636")
		}
	}
	rexsult, err = lang.RxFromString("1.00000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub637")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub637")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub638")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub638")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub639")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub639")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456788"))
	if err != nil {
		lang.SayString("ddsub640")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub640")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456788"))
	if err != nil {
		lang.SayString("ddsub641")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub641")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456788"))
	if err != nil {
		lang.SayString("ddsub642")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub642")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub650")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub650")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub651")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub651")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub652")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub652")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub653")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub653")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub654")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub654")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub655")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub655")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub656")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub656")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub657")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub657")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub658")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub658")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub659")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub659")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub660")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub660")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub661")
	} else {
		if !(rexsult.ToString() == "-1231234567456789") {
			lang.SayString("ddsub661")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub662")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub662")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub663")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub663")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub664")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub664")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub665")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub665")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub666")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub666")
		}
	}
	rexsult, err = lang.RxFromString("1.00000001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub667")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub667")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub668")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub668")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1231234567456789"))
	if err != nil {
		lang.SayString("ddsub669")
	} else {
		if !(rexsult.ToString() == "-1231234567456788") {
			lang.SayString("ddsub669")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.1"))
	if err != nil {
		lang.SayString("ddsub670")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("ddsub670")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.9"))
	if err != nil {
		lang.SayString("ddsub671")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddsub671")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.1"))
	if err != nil {
		lang.SayString("ddsub672")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("ddsub672")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.5"))
	if err != nil {
		lang.SayString("ddsub673")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("ddsub673")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.9"))
	if err != nil {
		lang.SayString("ddsub674")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("ddsub674")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.1"))
	if err != nil {
		lang.SayString("ddsub680")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("ddsub680")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.9"))
	if err != nil {
		lang.SayString("ddsub681")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddsub681")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.1"))
	if err != nil {
		lang.SayString("ddsub682")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("ddsub682")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.5"))
	if err != nil {
		lang.SayString("ddsub683")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("ddsub683")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.9"))
	if err != nil {
		lang.SayString("ddsub684")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("ddsub684")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456787.1"))
	if err != nil {
		lang.SayString("ddsub685")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("ddsub685")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456787.9"))
	if err != nil {
		lang.SayString("ddsub686")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddsub686")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.1"))
	if err != nil {
		lang.SayString("ddsub687")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("ddsub687")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.5"))
	if err != nil {
		lang.SayString("ddsub688")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("ddsub688")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.9"))
	if err != nil {
		lang.SayString("ddsub689")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("ddsub689")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.1"))
	if err != nil {
		lang.SayString("ddsub690")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("ddsub690")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456788.9"))
	if err != nil {
		lang.SayString("ddsub691")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("ddsub691")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.1"))
	if err != nil {
		lang.SayString("ddsub692")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("ddsub692")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.5"))
	if err != nil {
		lang.SayString("ddsub693")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("ddsub693")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1234567456789.9"))
	if err != nil {
		lang.SayString("ddsub694")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("ddsub694")
		}
	}
	rexsult, err = lang.RxFromString("2.E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.002"))
	if err != nil {
		lang.SayString("ddsub901")
	} else {
		if !(rexsult.ToString() == "-1.000") {
			lang.SayString("ddsub901")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.002"))
	if err != nil {
		lang.SayString("ddsub902")
	} else {
		if !(rexsult.ToString() == "-1.0000") {
			lang.SayString("ddsub902")
		}
	}
	rexsult, err = lang.RxFromString("2.00E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.0020"))
	if err != nil {
		lang.SayString("ddsub903")
	} else {
		if !(rexsult.ToString() == "-1.00000") {
			lang.SayString("ddsub903")
		}
	}
	rexsult, err = lang.RxFromString("2.000E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00200"))
	if err != nil {
		lang.SayString("ddsub904")
	} else {
		if !(rexsult.ToString() == "-1.000000") {
			lang.SayString("ddsub904")
		}
	}
	rexsult, err = lang.RxFromString("2.0000E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.002000"))
	if err != nil {
		lang.SayString("ddsub905")
	} else {
		if !(rexsult.ToString() == "-1.0000000") {
			lang.SayString("ddsub905")
		}
	}
	rexsult, err = lang.RxFromString("2.00000E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.0020000"))
	if err != nil {
		lang.SayString("ddsub906")
	} else {
		if !(rexsult.ToString() == "-1.00000000") {
			lang.SayString("ddsub906")
		}
	}
	rexsult, err = lang.RxFromString("2.000000E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.00200000"))
	if err != nil {
		lang.SayString("ddsub907")
	} else {
		if !(rexsult.ToString() == "-1.000000000") {
			lang.SayString("ddsub907")
		}
	}
	rexsult, err = lang.RxFromString("2.0000000E-3").OpSub(lang.RxSetWithDigit(16), lang.RxFromString("1.002000000"))
	if err != nil {
		lang.SayString("ddsub908")
	} else {
		if !(rexsult.ToString() == "-1.0000000000") {
			lang.SayString("ddsub908")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub001")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub001")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqsub003")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqsub003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqsub004")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqsub005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub005")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqsub006")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqsub006")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqsub007")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqsub007")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub011")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub011")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub012")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dqsub012")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqsub013")
	} else {
		if !(rexsult.ToString() == "-3") {
			lang.SayString("dqsub013")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub014")
	} else {
		if !(rexsult.ToString() == "-3") {
			lang.SayString("dqsub014")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqsub015")
	} else {
		if !(rexsult.ToString() == "-4") {
			lang.SayString("dqsub015")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqsub016")
	} else {
		if !(rexsult.ToString() == "-5") {
			lang.SayString("dqsub016")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqsub017")
	} else {
		if !(rexsult.ToString() == "-5") {
			lang.SayString("dqsub017")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("dqsub021")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub021")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqsub022")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqsub022")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqsub023")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqsub023")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqsub024")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqsub024")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqsub025")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("dqsub025")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqsub026")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("dqsub026")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("dqsub027")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("dqsub027")
		}
	}
	rexsult, err = lang.RxFromString("11").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub030")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqsub030")
		}
	}
	rexsult, err = lang.RxFromString("10").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub031")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dqsub031")
		}
	}
	rexsult, err = lang.RxFromString("9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub032")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("dqsub032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub033")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub034")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqsub034")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub035")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dqsub035")
		}
	}
	rexsult, err = lang.RxFromString("-9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub036")
	} else {
		if !(rexsult.ToString() == "-10") {
			lang.SayString("dqsub036")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub037")
	} else {
		if !(rexsult.ToString() == "-11") {
			lang.SayString("dqsub037")
		}
	}
	rexsult, err = lang.RxFromString("-11").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub038")
	} else {
		if !(rexsult.ToString() == "-12") {
			lang.SayString("dqsub038")
		}
	}
	rexsult, err = lang.RxFromString("5.75").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("3.3"))
	if err != nil {
		lang.SayString("dqsub040")
	} else {
		if !(rexsult.ToString() == "2.45") {
			lang.SayString("dqsub040")
		}
	}
	rexsult, err = lang.RxFromString("5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("dqsub041")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("dqsub041")
		}
	}
	rexsult, err = lang.RxFromString("-5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("dqsub042")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("dqsub042")
		}
	}
	rexsult, err = lang.RxFromString("-7").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("2.5"))
	if err != nil {
		lang.SayString("dqsub043")
	} else {
		if !(rexsult.ToString() == "-9.5") {
			lang.SayString("dqsub043")
		}
	}
	rexsult, err = lang.RxFromString("0.7").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqsub044")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqsub044")
		}
	}
	rexsult, err = lang.RxFromString("1.3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqsub045")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("dqsub045")
		}
	}
	rexsult, err = lang.RxFromString("1.25").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.25"))
	if err != nil {
		lang.SayString("dqsub046")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub046")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00000000"))
	if err != nil {
		lang.SayString("dqsub050")
	} else {
		if !(rexsult.ToString() == "0.23456789") {
			lang.SayString("dqsub050")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00000089"))
	if err != nil {
		lang.SayString("dqsub051")
	} else {
		if !(rexsult.ToString() == "0.23456700") {
			lang.SayString("dqsub051")
		}
	}
	rexsult, err = lang.RxFromString("70").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+34"))
	if err != nil {
		lang.SayString("dqsub060")
	} else {
		if !(rexsult.ToString() == "-1.000000000000000000000000000000000E+38") {
			lang.SayString("dqsub060")
		}
	}
	rexsult, err = lang.RxFromString("700").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+34"))
	if err != nil {
		lang.SayString("dqsub061")
	} else {
		if !(rexsult.ToString() == "-1.000000000000000000000000000000000E+38") {
			lang.SayString("dqsub061")
		}
	}
	rexsult, err = lang.RxFromString("7000").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+34"))
	if err != nil {
		lang.SayString("dqsub062")
	} else {
		if !(rexsult.ToString() == "-1.000000000000000000000000000000000E+38") {
			lang.SayString("dqsub062")
		}
	}
	rexsult, err = lang.RxFromString("70000").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+34"))
	if err != nil {
		lang.SayString("dqsub063")
	} else {
		if !(rexsult.ToString() == "-9.99999999999999999999999999999999E+37") {
			lang.SayString("dqsub063")
		}
	}
	rexsult, err = lang.RxFromString("700000").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+34"))
	if err != nil {
		lang.SayString("dqsub064")
	} else {
		if !(rexsult.ToString() == "-9.99999999999999999999999999999993E+37") {
			lang.SayString("dqsub064")
		}
	}
	rexsult, err = lang.RxFromString("10000e+34").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("70"))
	if err != nil {
		lang.SayString("dqsub065")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000E+38") {
			lang.SayString("dqsub065")
		}
	}
	rexsult, err = lang.RxFromString("10000e+34").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("700"))
	if err != nil {
		lang.SayString("dqsub066")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000E+38") {
			lang.SayString("dqsub066")
		}
	}
	rexsult, err = lang.RxFromString("10000e+34").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("7000"))
	if err != nil {
		lang.SayString("dqsub067")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000E+38") {
			lang.SayString("dqsub067")
		}
	}
	rexsult, err = lang.RxFromString("10000e+34").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("dqsub068")
	} else {
		if !(rexsult.ToString() == "9.99999999999999999999999999999999E+37") {
			lang.SayString("dqsub068")
		}
	}
	rexsult, err = lang.RxFromString("10000e+34").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("700000"))
	if err != nil {
		lang.SayString("dqsub069")
	} else {
		if !(rexsult.ToString() == "9.99999999999999999999999999999993E+37") {
			lang.SayString("dqsub069")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("dqsub090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub090")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.00"))
	if err != nil {
		lang.SayString("dqsub091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("00.0"))
	if err != nil {
		lang.SayString("dqsub092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub092")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.00"))
	if err != nil {
		lang.SayString("dqsub093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub093")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("00.0"))
	if err != nil {
		lang.SayString("dqsub094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub094")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("dqsub095")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("dqsub095")
		}
	}
	rexsult, err = lang.RxFromString("3.").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("dqsub096")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("dqsub096")
		}
	}
	rexsult, err = lang.RxFromString("3.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("dqsub097")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("dqsub097")
		}
	}
	rexsult, err = lang.RxFromString("3.00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("dqsub098")
	} else {
		if !(rexsult.ToString() == "2.70") {
			lang.SayString("dqsub098")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqsub099")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub099")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("+3"))
	if err != nil {
		lang.SayString("dqsub100")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub100")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("dqsub101")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dqsub101")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqsub102")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("dqsub102")
		}
	}
	rexsult, err = lang.RxFromString("3.").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqsub103")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("dqsub103")
		}
	}
	rexsult, err = lang.RxFromString("3.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqsub104")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("dqsub104")
		}
	}
	rexsult, err = lang.RxFromString("3.00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("dqsub105")
	} else {
		if !(rexsult.ToString() == "2.70") {
			lang.SayString("dqsub105")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("3.0"))
	if err != nil {
		lang.SayString("dqsub106")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub106")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("+3.0"))
	if err != nil {
		lang.SayString("dqsub107")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub107")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-3.0"))
	if err != nil {
		lang.SayString("dqsub108")
	} else {
		if !(rexsult.ToString() == "6.0") {
			lang.SayString("dqsub108")
		}
	}
	rexsult, err = lang.RxFromString("10.23456784").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub120")
	} else {
		if !(rexsult.ToString() == "-5E-8") {
			lang.SayString("dqsub120")
		}
	}
	rexsult, err = lang.RxFromString("10.23456785").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub121")
	} else {
		if !(rexsult.ToString() == "-4E-8") {
			lang.SayString("dqsub121")
		}
	}
	rexsult, err = lang.RxFromString("10.23456786").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub122")
	} else {
		if !(rexsult.ToString() == "-3E-8") {
			lang.SayString("dqsub122")
		}
	}
	rexsult, err = lang.RxFromString("10.23456787").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub123")
	} else {
		if !(rexsult.ToString() == "-2E-8") {
			lang.SayString("dqsub123")
		}
	}
	rexsult, err = lang.RxFromString("10.23456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub124")
	} else {
		if !(rexsult.ToString() == "-1E-8") {
			lang.SayString("dqsub124")
		}
	}
	rexsult, err = lang.RxFromString("10.23456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub125")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub125")
		}
	}
	rexsult, err = lang.RxFromString("10.23456790").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub126")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("dqsub126")
		}
	}
	rexsult, err = lang.RxFromString("10.23456791").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub127")
	} else {
		if !(rexsult.ToString() == "2E-8") {
			lang.SayString("dqsub127")
		}
	}
	rexsult, err = lang.RxFromString("10.23456792").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub128")
	} else {
		if !(rexsult.ToString() == "3E-8") {
			lang.SayString("dqsub128")
		}
	}
	rexsult, err = lang.RxFromString("10.23456793").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub129")
	} else {
		if !(rexsult.ToString() == "4E-8") {
			lang.SayString("dqsub129")
		}
	}
	rexsult, err = lang.RxFromString("10.23456794").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("dqsub130")
	} else {
		if !(rexsult.ToString() == "5E-8") {
			lang.SayString("dqsub130")
		}
	}
	rexsult, err = lang.RxFromString("10.23456781").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub131")
	} else {
		if !(rexsult.ToString() == "-5E-8") {
			lang.SayString("dqsub131")
		}
	}
	rexsult, err = lang.RxFromString("10.23456782").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub132")
	} else {
		if !(rexsult.ToString() == "-4E-8") {
			lang.SayString("dqsub132")
		}
	}
	rexsult, err = lang.RxFromString("10.23456783").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub133")
	} else {
		if !(rexsult.ToString() == "-3E-8") {
			lang.SayString("dqsub133")
		}
	}
	rexsult, err = lang.RxFromString("10.23456784").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub134")
	} else {
		if !(rexsult.ToString() == "-2E-8") {
			lang.SayString("dqsub134")
		}
	}
	rexsult, err = lang.RxFromString("10.23456785").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub135")
	} else {
		if !(rexsult.ToString() == "-1E-8") {
			lang.SayString("dqsub135")
		}
	}
	rexsult, err = lang.RxFromString("10.23456786").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub136")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqsub136")
		}
	}
	rexsult, err = lang.RxFromString("10.23456787").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub137")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("dqsub137")
		}
	}
	rexsult, err = lang.RxFromString("10.23456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub138")
	} else {
		if !(rexsult.ToString() == "2E-8") {
			lang.SayString("dqsub138")
		}
	}
	rexsult, err = lang.RxFromString("10.23456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub139")
	} else {
		if !(rexsult.ToString() == "3E-8") {
			lang.SayString("dqsub139")
		}
	}
	rexsult, err = lang.RxFromString("10.23456790").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub140")
	} else {
		if !(rexsult.ToString() == "4E-8") {
			lang.SayString("dqsub140")
		}
	}
	rexsult, err = lang.RxFromString("10.23456791").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("dqsub141")
	} else {
		if !(rexsult.ToString() == "5E-8") {
			lang.SayString("dqsub141")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("dqsub142")
	} else {
		if !(rexsult.ToString() == "1E-9") {
			lang.SayString("dqsub142")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub143")
	} else {
		if !(rexsult.ToString() == "-1E-9") {
			lang.SayString("dqsub143")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456780").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("dqsub144")
	} else {
		if !(rexsult.ToString() == "6E-8") {
			lang.SayString("dqsub144")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456790").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("dqsub145")
	} else {
		if !(rexsult.ToString() == "-4E-8") {
			lang.SayString("dqsub145")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456791").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("dqsub146")
	} else {
		if !(rexsult.ToString() == "-5E-8") {
			lang.SayString("dqsub146")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".1"))
	if err != nil {
		lang.SayString("dqsub160")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("dqsub160")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".97983"))
	if err != nil {
		lang.SayString("dqsub161")
	} else {
		if !(rexsult.ToString() == "-0.97983") {
			lang.SayString("dqsub161")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".9"))
	if err != nil {
		lang.SayString("dqsub162")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("dqsub162")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.102"))
	if err != nil {
		lang.SayString("dqsub163")
	} else {
		if !(rexsult.ToString() == "-0.102") {
			lang.SayString("dqsub163")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".4"))
	if err != nil {
		lang.SayString("dqsub164")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("dqsub164")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".307"))
	if err != nil {
		lang.SayString("dqsub165")
	} else {
		if !(rexsult.ToString() == "-0.307") {
			lang.SayString("dqsub165")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".43822"))
	if err != nil {
		lang.SayString("dqsub166")
	} else {
		if !(rexsult.ToString() == "-0.43822") {
			lang.SayString("dqsub166")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".911"))
	if err != nil {
		lang.SayString("dqsub167")
	} else {
		if !(rexsult.ToString() == "-0.911") {
			lang.SayString("dqsub167")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".02"))
	if err != nil {
		lang.SayString("dqsub168")
	} else {
		if !(rexsult.ToString() == "-0.02") {
			lang.SayString("dqsub168")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".392"))
	if err != nil {
		lang.SayString("dqsub169")
	} else {
		if !(rexsult.ToString() == "-0.392") {
			lang.SayString("dqsub169")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".26"))
	if err != nil {
		lang.SayString("dqsub170")
	} else {
		if !(rexsult.ToString() == "-0.26") {
			lang.SayString("dqsub170")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("dqsub171")
	} else {
		if !(rexsult.ToString() == "-0.51") {
			lang.SayString("dqsub171")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".2234"))
	if err != nil {
		lang.SayString("dqsub172")
	} else {
		if !(rexsult.ToString() == "-0.2234") {
			lang.SayString("dqsub172")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".2"))
	if err != nil {
		lang.SayString("dqsub173")
	} else {
		if !(rexsult.ToString() == "-0.2") {
			lang.SayString("dqsub173")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".0008"))
	if err != nil {
		lang.SayString("dqsub174")
	} else {
		if !(rexsult.ToString() == "-0.0008") {
			lang.SayString("dqsub174")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.1"))
	if err != nil {
		lang.SayString("dqsub180")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqsub180")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.97983"))
	if err != nil {
		lang.SayString("dqsub181")
	} else {
		if !(rexsult.ToString() == "0.97983") {
			lang.SayString("dqsub181")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.9"))
	if err != nil {
		lang.SayString("dqsub182")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("dqsub182")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.102"))
	if err != nil {
		lang.SayString("dqsub183")
	} else {
		if !(rexsult.ToString() == "0.102") {
			lang.SayString("dqsub183")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.4"))
	if err != nil {
		lang.SayString("dqsub184")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqsub184")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.307"))
	if err != nil {
		lang.SayString("dqsub185")
	} else {
		if !(rexsult.ToString() == "0.307") {
			lang.SayString("dqsub185")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.43822"))
	if err != nil {
		lang.SayString("dqsub186")
	} else {
		if !(rexsult.ToString() == "0.43822") {
			lang.SayString("dqsub186")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.911"))
	if err != nil {
		lang.SayString("dqsub187")
	} else {
		if !(rexsult.ToString() == "0.911") {
			lang.SayString("dqsub187")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.02"))
	if err != nil {
		lang.SayString("dqsub188")
	} else {
		if !(rexsult.ToString() == "0.02") {
			lang.SayString("dqsub188")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.392"))
	if err != nil {
		lang.SayString("dqsub189")
	} else {
		if !(rexsult.ToString() == "0.392") {
			lang.SayString("dqsub189")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.26"))
	if err != nil {
		lang.SayString("dqsub190")
	} else {
		if !(rexsult.ToString() == "0.26") {
			lang.SayString("dqsub190")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.51"))
	if err != nil {
		lang.SayString("dqsub191")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("dqsub191")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.2234"))
	if err != nil {
		lang.SayString("dqsub192")
	} else {
		if !(rexsult.ToString() == "0.2234") {
			lang.SayString("dqsub192")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.2"))
	if err != nil {
		lang.SayString("dqsub193")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("dqsub193")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.0008"))
	if err != nil {
		lang.SayString("dqsub194")
	} else {
		if !(rexsult.ToString() == "0.0008") {
			lang.SayString("dqsub194")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.1"))
	if err != nil {
		lang.SayString("dqsub200")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqsub200")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.97983"))
	if err != nil {
		lang.SayString("dqsub201")
	} else {
		if !(rexsult.ToString() == "0.97983") {
			lang.SayString("dqsub201")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.9"))
	if err != nil {
		lang.SayString("dqsub202")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("dqsub202")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.102"))
	if err != nil {
		lang.SayString("dqsub203")
	} else {
		if !(rexsult.ToString() == "0.102") {
			lang.SayString("dqsub203")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.4"))
	if err != nil {
		lang.SayString("dqsub204")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("dqsub204")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.307"))
	if err != nil {
		lang.SayString("dqsub205")
	} else {
		if !(rexsult.ToString() == "0.307") {
			lang.SayString("dqsub205")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.43822"))
	if err != nil {
		lang.SayString("dqsub206")
	} else {
		if !(rexsult.ToString() == "0.43822") {
			lang.SayString("dqsub206")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.911"))
	if err != nil {
		lang.SayString("dqsub207")
	} else {
		if !(rexsult.ToString() == "0.911") {
			lang.SayString("dqsub207")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.02"))
	if err != nil {
		lang.SayString("dqsub208")
	} else {
		if !(rexsult.ToString() == "0.02") {
			lang.SayString("dqsub208")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.392"))
	if err != nil {
		lang.SayString("dqsub209")
	} else {
		if !(rexsult.ToString() == "0.392") {
			lang.SayString("dqsub209")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.26"))
	if err != nil {
		lang.SayString("dqsub210")
	} else {
		if !(rexsult.ToString() == "0.26") {
			lang.SayString("dqsub210")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.51"))
	if err != nil {
		lang.SayString("dqsub211")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("dqsub211")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.2234"))
	if err != nil {
		lang.SayString("dqsub212")
	} else {
		if !(rexsult.ToString() == "0.2234") {
			lang.SayString("dqsub212")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.2"))
	if err != nil {
		lang.SayString("dqsub213")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("dqsub213")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-.0008"))
	if err != nil {
		lang.SayString("dqsub214")
	} else {
		if !(rexsult.ToString() == "0.0008") {
			lang.SayString("dqsub214")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-12").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub220")
	} else {
		if !(rexsult.ToString() == "-5.6267E-8") {
			lang.SayString("dqsub220")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-11").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub221")
	} else {
		if !(rexsult.ToString() == "-5.6267E-7") {
			lang.SayString("dqsub221")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-10").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub222")
	} else {
		if !(rexsult.ToString() == "-0.0000056267") {
			lang.SayString("dqsub222")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub223")
	} else {
		if !(rexsult.ToString() == "-0.000056267") {
			lang.SayString("dqsub223")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-8").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub224")
	} else {
		if !(rexsult.ToString() == "-0.00056267") {
			lang.SayString("dqsub224")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-7").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub225")
	} else {
		if !(rexsult.ToString() == "-0.0056267") {
			lang.SayString("dqsub225")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-6").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub226")
	} else {
		if !(rexsult.ToString() == "-0.056267") {
			lang.SayString("dqsub226")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub227")
	} else {
		if !(rexsult.ToString() == "-0.56267") {
			lang.SayString("dqsub227")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub228")
	} else {
		if !(rexsult.ToString() == "-562.67") {
			lang.SayString("dqsub228")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub229")
	} else {
		if !(rexsult.ToString() == "-5626.7") {
			lang.SayString("dqsub229")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub230")
	} else {
		if !(rexsult.ToString() == "-56267") {
			lang.SayString("dqsub230")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-12"))
	if err != nil {
		lang.SayString("dqsub240")
	} else {
		if !(rexsult.ToString() == "5.6267E-8") {
			lang.SayString("dqsub240")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-11"))
	if err != nil {
		lang.SayString("dqsub241")
	} else {
		if !(rexsult.ToString() == "5.6267E-7") {
			lang.SayString("dqsub241")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-10"))
	if err != nil {
		lang.SayString("dqsub242")
	} else {
		if !(rexsult.ToString() == "0.0000056267") {
			lang.SayString("dqsub242")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-9"))
	if err != nil {
		lang.SayString("dqsub243")
	} else {
		if !(rexsult.ToString() == "0.000056267") {
			lang.SayString("dqsub243")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-8"))
	if err != nil {
		lang.SayString("dqsub244")
	} else {
		if !(rexsult.ToString() == "0.00056267") {
			lang.SayString("dqsub244")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-7"))
	if err != nil {
		lang.SayString("dqsub245")
	} else {
		if !(rexsult.ToString() == "0.0056267") {
			lang.SayString("dqsub245")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-6"))
	if err != nil {
		lang.SayString("dqsub246")
	} else {
		if !(rexsult.ToString() == "0.056267") {
			lang.SayString("dqsub246")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-5"))
	if err != nil {
		lang.SayString("dqsub247")
	} else {
		if !(rexsult.ToString() == "0.56267") {
			lang.SayString("dqsub247")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-2"))
	if err != nil {
		lang.SayString("dqsub248")
	} else {
		if !(rexsult.ToString() == "562.67") {
			lang.SayString("dqsub248")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-1"))
	if err != nil {
		lang.SayString("dqsub249")
	} else {
		if !(rexsult.ToString() == "5626.7") {
			lang.SayString("dqsub249")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-0"))
	if err != nil {
		lang.SayString("dqsub250")
	} else {
		if !(rexsult.ToString() == "56267") {
			lang.SayString("dqsub250")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00000000"))
	if err != nil {
		lang.SayString("dqsub301")
	} else {
		if !(rexsult.ToString() == "0.23456789") {
			lang.SayString("dqsub301")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00000011"))
	if err != nil {
		lang.SayString("dqsub302")
	} else {
		if !(rexsult.ToString() == "0.23456778") {
			lang.SayString("dqsub302")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.0000"))
	if err != nil {
		lang.SayString("dqsub321")
	} else {
		if !(rexsult.ToString() == "0.9998") {
			lang.SayString("dqsub321")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.0001"))
	if err != nil {
		lang.SayString("dqsub322")
	} else {
		if !(rexsult.ToString() == "0.9997") {
			lang.SayString("dqsub322")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.0002"))
	if err != nil {
		lang.SayString("dqsub323")
	} else {
		if !(rexsult.ToString() == "0.9996") {
			lang.SayString("dqsub323")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.0003"))
	if err != nil {
		lang.SayString("dqsub324")
	} else {
		if !(rexsult.ToString() == "0.9995") {
			lang.SayString("dqsub324")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.0000"))
	if err != nil {
		lang.SayString("dqsub325")
	} else {
		if !(rexsult.ToString() == "0.9998") {
			lang.SayString("dqsub325")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.0001"))
	if err != nil {
		lang.SayString("dqsub326")
	} else {
		if !(rexsult.ToString() == "0.9999") {
			lang.SayString("dqsub326")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.0002"))
	if err != nil {
		lang.SayString("dqsub327")
	} else {
		if !(rexsult.ToString() == "1.0000") {
			lang.SayString("dqsub327")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.0003"))
	if err != nil {
		lang.SayString("dqsub328")
	} else {
		if !(rexsult.ToString() == "1.0001") {
			lang.SayString("dqsub328")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dqsub346")
	} else {
		if !(rexsult.ToString() == "9999999999993") {
			lang.SayString("dqsub346")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("70"))
	if err != nil {
		lang.SayString("dqsub347")
	} else {
		if !(rexsult.ToString() == "9999999999930") {
			lang.SayString("dqsub347")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("700"))
	if err != nil {
		lang.SayString("dqsub348")
	} else {
		if !(rexsult.ToString() == "9999999999300") {
			lang.SayString("dqsub348")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("7000"))
	if err != nil {
		lang.SayString("dqsub349")
	} else {
		if !(rexsult.ToString() == "9999999993000") {
			lang.SayString("dqsub349")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("dqsub350")
	} else {
		if !(rexsult.ToString() == "9999999930000") {
			lang.SayString("dqsub350")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("700000"))
	if err != nil {
		lang.SayString("dqsub351")
	} else {
		if !(rexsult.ToString() == "9999999300000") {
			lang.SayString("dqsub351")
		}
	}
	rexsult, err = lang.RxFromString("7").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("dqsub352")
	} else {
		if !(rexsult.ToString() == "-9999999999993") {
			lang.SayString("dqsub352")
		}
	}
	rexsult, err = lang.RxFromString("70").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("dqsub353")
	} else {
		if !(rexsult.ToString() == "-9999999999930") {
			lang.SayString("dqsub353")
		}
	}
	rexsult, err = lang.RxFromString("700").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("dqsub354")
	} else {
		if !(rexsult.ToString() == "-9999999999300") {
			lang.SayString("dqsub354")
		}
	}
	rexsult, err = lang.RxFromString("7000").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("dqsub355")
	} else {
		if !(rexsult.ToString() == "-9999999993000") {
			lang.SayString("dqsub355")
		}
	}
	rexsult, err = lang.RxFromString("70000").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("dqsub356")
	} else {
		if !(rexsult.ToString() == "-9999999930000") {
			lang.SayString("dqsub356")
		}
	}
	rexsult, err = lang.RxFromString("700000").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("dqsub357")
	} else {
		if !(rexsult.ToString() == "-9999999300000") {
			lang.SayString("dqsub357")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.0001"))
	if err != nil {
		lang.SayString("dqsub361")
	} else {
		if !(rexsult.ToString() == "0.9999") {
			lang.SayString("dqsub361")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.00001"))
	if err != nil {
		lang.SayString("dqsub362")
	} else {
		if !(rexsult.ToString() == "0.99999") {
			lang.SayString("dqsub362")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("dqsub363")
	} else {
		if !(rexsult.ToString() == "0.999999") {
			lang.SayString("dqsub363")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.0000000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqsub364")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000") {
			lang.SayString("dqsub364")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.00000000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqsub365")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000") {
			lang.SayString("dqsub365")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.000000000000000000000000000000000001"))
	if err != nil {
		lang.SayString("dqsub366")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000") {
			lang.SayString("dqsub366")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub370")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqsub370")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0."))
	if err != nil {
		lang.SayString("dqsub371")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqsub371")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString(".0"))
	if err != nil {
		lang.SayString("dqsub372")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqsub372")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("dqsub373")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqsub373")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub374")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqsub374")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub375")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqsub375")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub376")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqsub376")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub377")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("dqsub377")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-51897955.3"))
	if err != nil {
		lang.SayString("dqsub910")
	} else {
		if !(rexsult.ToString() == "-51621406.7") {
			lang.SayString("dqsub910")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("89827.5229"))
	if err != nil {
		lang.SayString("dqsub911")
	} else {
		if !(rexsult.ToString() == "69751.9211") {
			lang.SayString("dqsub911")
		}
	}
	rexsult, err = lang.RxFromString("333.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("33.00000000000000000000000001234566"))
	if err != nil {
		lang.SayString("dqsub920")
	} else {
		if !(rexsult.ToString() == "299.9999999999999999999999999999999") {
			lang.SayString("dqsub920")
		}
	}
	rexsult, err = lang.RxFromString("333.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("33.00000000000000000000000001234565"))
	if err != nil {
		lang.SayString("dqsub921")
	} else {
		if !(rexsult.ToString() == "300.0000000000000000000000000000000") {
			lang.SayString("dqsub921")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("33.00000000000000000000000001234565"))
	if err != nil {
		lang.SayString("dqsub922")
	} else {
		if !(rexsult.ToString() == "100.0000000000000000000000000000000") {
			lang.SayString("dqsub922")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("33.00000000000000000000000001234564"))
	if err != nil {
		lang.SayString("dqsub923")
	} else {
		if !(rexsult.ToString() == "100.0000000000000000000000000000000") {
			lang.SayString("dqsub923")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("33.00000000000000000000000001234540"))
	if err != nil {
		lang.SayString("dqsub924")
	} else {
		if !(rexsult.ToString() == "100.0000000000000000000000000000002") {
			lang.SayString("dqsub924")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("43.00000000000000000000000001234560"))
	if err != nil {
		lang.SayString("dqsub925")
	} else {
		if !(rexsult.ToString() == "90.0000000000000000000000000000000") {
			lang.SayString("dqsub925")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("43.00000000000000000000000001234561"))
	if err != nil {
		lang.SayString("dqsub926")
	} else {
		if !(rexsult.ToString() == "90.0000000000000000000000000000000") {
			lang.SayString("dqsub926")
		}
	}
	rexsult, err = lang.RxFromString("133.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("43.00000000000000000000000001234566"))
	if err != nil {
		lang.SayString("dqsub927")
	} else {
		if !(rexsult.ToString() == "89.9999999999999999999999999999999") {
			lang.SayString("dqsub927")
		}
	}
	rexsult, err = lang.RxFromString("101.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("91.00000000000000000000000001234566"))
	if err != nil {
		lang.SayString("dqsub928")
	} else {
		if !(rexsult.ToString() == "9.9999999999999999999999999999999") {
			lang.SayString("dqsub928")
		}
	}
	rexsult, err = lang.RxFromString("101.0000000000000000000000000123456").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("99.00000000000000000000000001234566"))
	if err != nil {
		lang.SayString("dqsub929")
	} else {
		if !(rexsult.ToString() == "1.9999999999999999999999999999999") {
			lang.SayString("dqsub929")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-10").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub390")
	} else {
		if !(rexsult.ToString() == "-0.0000056267") {
			lang.SayString("dqsub390")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-6").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub391")
	} else {
		if !(rexsult.ToString() == "-0.056267") {
			lang.SayString("dqsub391")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub392")
	} else {
		if !(rexsult.ToString() == "-0.56267") {
			lang.SayString("dqsub392")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-4").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub393")
	} else {
		if !(rexsult.ToString() == "-5.6267") {
			lang.SayString("dqsub393")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub394")
	} else {
		if !(rexsult.ToString() == "-56.267") {
			lang.SayString("dqsub394")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub395")
	} else {
		if !(rexsult.ToString() == "-562.67") {
			lang.SayString("dqsub395")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub396")
	} else {
		if !(rexsult.ToString() == "-5626.7") {
			lang.SayString("dqsub396")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub397")
	} else {
		if !(rexsult.ToString() == "-56267") {
			lang.SayString("dqsub397")
		}
	}
	rexsult, err = lang.RxFromString("-5E-10").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub398")
	} else {
		if !(rexsult.ToString() == "-5E-10") {
			lang.SayString("dqsub398")
		}
	}
	rexsult, err = lang.RxFromString("-5E-7").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub399")
	} else {
		if !(rexsult.ToString() == "-5E-7") {
			lang.SayString("dqsub399")
		}
	}
	rexsult, err = lang.RxFromString("-5E-6").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub400")
	} else {
		if !(rexsult.ToString() == "-0.000005") {
			lang.SayString("dqsub400")
		}
	}
	rexsult, err = lang.RxFromString("-5E-5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub401")
	} else {
		if !(rexsult.ToString() == "-0.00005") {
			lang.SayString("dqsub401")
		}
	}
	rexsult, err = lang.RxFromString("-5E-4").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub402")
	} else {
		if !(rexsult.ToString() == "-0.0005") {
			lang.SayString("dqsub402")
		}
	}
	rexsult, err = lang.RxFromString("-5E-1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub403")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dqsub403")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-10"))
	if err != nil {
		lang.SayString("dqsub420")
	} else {
		if !(rexsult.ToString() == "0.0000056267") {
			lang.SayString("dqsub420")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-6"))
	if err != nil {
		lang.SayString("dqsub421")
	} else {
		if !(rexsult.ToString() == "0.056267") {
			lang.SayString("dqsub421")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-5"))
	if err != nil {
		lang.SayString("dqsub422")
	} else {
		if !(rexsult.ToString() == "0.56267") {
			lang.SayString("dqsub422")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-4"))
	if err != nil {
		lang.SayString("dqsub423")
	} else {
		if !(rexsult.ToString() == "5.6267") {
			lang.SayString("dqsub423")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-3"))
	if err != nil {
		lang.SayString("dqsub424")
	} else {
		if !(rexsult.ToString() == "56.267") {
			lang.SayString("dqsub424")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-2"))
	if err != nil {
		lang.SayString("dqsub425")
	} else {
		if !(rexsult.ToString() == "562.67") {
			lang.SayString("dqsub425")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-1"))
	if err != nil {
		lang.SayString("dqsub426")
	} else {
		if !(rexsult.ToString() == "5626.7") {
			lang.SayString("dqsub426")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-56267E-0"))
	if err != nil {
		lang.SayString("dqsub427")
	} else {
		if !(rexsult.ToString() == "56267") {
			lang.SayString("dqsub427")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-5E-10"))
	if err != nil {
		lang.SayString("dqsub428")
	} else {
		if !(rexsult.ToString() == "5E-10") {
			lang.SayString("dqsub428")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-5E-7"))
	if err != nil {
		lang.SayString("dqsub429")
	} else {
		if !(rexsult.ToString() == "5E-7") {
			lang.SayString("dqsub429")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-5E-6"))
	if err != nil {
		lang.SayString("dqsub430")
	} else {
		if !(rexsult.ToString() == "0.000005") {
			lang.SayString("dqsub430")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-5E-5"))
	if err != nil {
		lang.SayString("dqsub431")
	} else {
		if !(rexsult.ToString() == "0.00005") {
			lang.SayString("dqsub431")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-5E-4"))
	if err != nil {
		lang.SayString("dqsub432")
	} else {
		if !(rexsult.ToString() == "0.0005") {
			lang.SayString("dqsub432")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-5E-1"))
	if err != nil {
		lang.SayString("dqsub433")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqsub433")
		}
	}
	rexsult, err = lang.RxFromString("1E+16").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub461")
	} else {
		if !(rexsult.ToString() == "9999999999999999") {
			lang.SayString("dqsub461")
		}
	}
	rexsult, err = lang.RxFromString("1E+12").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-1.111"))
	if err != nil {
		lang.SayString("dqsub462")
	} else {
		if !(rexsult.ToString() == "1000000000001.111") {
			lang.SayString("dqsub462")
		}
	}
	rexsult, err = lang.RxFromString("1.111").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-1E+12"))
	if err != nil {
		lang.SayString("dqsub463")
	} else {
		if !(rexsult.ToString() == "1000000000001.111") {
			lang.SayString("dqsub463")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-1E+16"))
	if err != nil {
		lang.SayString("dqsub464")
	} else {
		if !(rexsult.ToString() == "9999999999999999") {
			lang.SayString("dqsub464")
		}
	}
	rexsult, err = lang.RxFromString("7E+15").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub465")
	} else {
		if !(rexsult.ToString() == "6999999999999999") {
			lang.SayString("dqsub465")
		}
	}
	rexsult, err = lang.RxFromString("7E+12").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-1.111"))
	if err != nil {
		lang.SayString("dqsub466")
	} else {
		if !(rexsult.ToString() == "7000000000001.111") {
			lang.SayString("dqsub466")
		}
	}
	rexsult, err = lang.RxFromString("1.111").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-7E+12"))
	if err != nil {
		lang.SayString("dqsub467")
	} else {
		if !(rexsult.ToString() == "7000000000001.111") {
			lang.SayString("dqsub467")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-7E+15"))
	if err != nil {
		lang.SayString("dqsub468")
	} else {
		if !(rexsult.ToString() == "6999999999999999") {
			lang.SayString("dqsub468")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555563"))
	if err != nil {
		lang.SayString("dqsub470")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000001") {
			lang.SayString("dqsub470")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555562"))
	if err != nil {
		lang.SayString("dqsub471")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000001") {
			lang.SayString("dqsub471")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555561"))
	if err != nil {
		lang.SayString("dqsub472")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000001") {
			lang.SayString("dqsub472")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555560"))
	if err != nil {
		lang.SayString("dqsub473")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000") {
			lang.SayString("dqsub473")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555559"))
	if err != nil {
		lang.SayString("dqsub474")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000") {
			lang.SayString("dqsub474")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555558"))
	if err != nil {
		lang.SayString("dqsub475")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000") {
			lang.SayString("dqsub475")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555557"))
	if err != nil {
		lang.SayString("dqsub476")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000") {
			lang.SayString("dqsub476")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555556"))
	if err != nil {
		lang.SayString("dqsub477")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000000") {
			lang.SayString("dqsub477")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555555"))
	if err != nil {
		lang.SayString("dqsub478")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999999999") {
			lang.SayString("dqsub478")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555554"))
	if err != nil {
		lang.SayString("dqsub479")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999999998") {
			lang.SayString("dqsub479")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555553"))
	if err != nil {
		lang.SayString("dqsub480")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999999997") {
			lang.SayString("dqsub480")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555552"))
	if err != nil {
		lang.SayString("dqsub481")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999999996") {
			lang.SayString("dqsub481")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555551"))
	if err != nil {
		lang.SayString("dqsub482")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999999995") {
			lang.SayString("dqsub482")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444444444444444444444444444").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("-0.5555555555555555555555555555555550"))
	if err != nil {
		lang.SayString("dqsub483")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999999994") {
			lang.SayString("dqsub483")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub500")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub500")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("dqsub501")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub501")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("dqsub502")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub502")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqsub503")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub503")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("dqsub504")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub504")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("dqsub505")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub505")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("dqsub506")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub506")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("dqsub507")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub507")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("dqsub508")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub508")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("dqsub509")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub509")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("dqsub510")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub510")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("dqsub511")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub511")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("dqsub512")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub512")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("dqsub513")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub513")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("dqsub514")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub514")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("dqsub515")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub515")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub516")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub516")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.000000001"))
	if err != nil {
		lang.SayString("dqsub517")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub517")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("dqsub518")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub518")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("dqsub519")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub519")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub520")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub520")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("dqsub521")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub521")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("dqsub522")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub522")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqsub523")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub523")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("dqsub524")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub524")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("dqsub525")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub525")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("dqsub526")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub526")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("dqsub527")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub527")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("dqsub528")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub528")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("dqsub529")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub529")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("dqsub530")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub530")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("dqsub531")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub531")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("dqsub532")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub532")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("dqsub533")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub533")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("dqsub534")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub534")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("dqsub535")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub535")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub536")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub536")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00000001"))
	if err != nil {
		lang.SayString("dqsub537")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub537")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("dqsub538")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub538")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("dqsub539")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub539")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("dqsub540")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub540")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("dqsub541")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub541")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("dqsub542")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub542")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("dqsub550")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub550")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("dqsub551")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub551")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("dqsub552")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub552")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqsub553")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub553")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("dqsub554")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub554")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("dqsub555")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub555")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("dqsub556")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub556")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("dqsub557")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub557")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("dqsub558")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub558")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("dqsub559")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub559")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("dqsub560")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub560")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("dqsub561")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456789") {
			lang.SayString("dqsub561")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("dqsub562")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub562")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("dqsub563")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub563")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("dqsub564")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub564")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("dqsub565")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub565")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqsub566")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub566")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00000001"))
	if err != nil {
		lang.SayString("dqsub567")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub567")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("dqsub568")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub568")
		}
	}
	rexsult, err = lang.RxFromString("1231234555555555555555555567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("dqsub569")
	} else {
		if !(rexsult.ToString() == "1231234555555555555555555567456788") {
			lang.SayString("dqsub569")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub600")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub600")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub601")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub601")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub602")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub602")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub603")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub603")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub604")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub604")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub605")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub605")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub606")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub606")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub607")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub607")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub608")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub608")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub609")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub609")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub610")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub610")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub611")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub611")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub612")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub612")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub613")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub613")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub614")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub614")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub615")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub615")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub616")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub616")
		}
	}
	rexsult, err = lang.RxFromString("1.000000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub617")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub617")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub618")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub618")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub619")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub619")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub620")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub620")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub621")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub621")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub622")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub622")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub623")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub623")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub624")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub624")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub625")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub625")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub626")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub626")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub627")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub627")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub628")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub628")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub629")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub629")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub630")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub630")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub631")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub631")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub632")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub632")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub633")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub633")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub634")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub634")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub635")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub635")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub636")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub636")
		}
	}
	rexsult, err = lang.RxFromString("1.00000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub637")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub637")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub638")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub638")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub639")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub639")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456788"))
	if err != nil {
		lang.SayString("dqsub640")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub640")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456788"))
	if err != nil {
		lang.SayString("dqsub641")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub641")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456788"))
	if err != nil {
		lang.SayString("dqsub642")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub642")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub650")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub650")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub651")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub651")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub652")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub652")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub653")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub653")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub654")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub654")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub655")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub655")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub656")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub656")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub657")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub657")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub658")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub658")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub659")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub659")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub660")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub660")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub661")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456789") {
			lang.SayString("dqsub661")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub662")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub662")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub663")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub663")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub664")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub664")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub665")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub665")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub666")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub666")
		}
	}
	rexsult, err = lang.RxFromString("1.00000001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub667")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub667")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub668")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub668")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1231234555555555555555555567456789"))
	if err != nil {
		lang.SayString("dqsub669")
	} else {
		if !(rexsult.ToString() == "-1231234555555555555555555567456788") {
			lang.SayString("dqsub669")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.1"))
	if err != nil {
		lang.SayString("dqsub670")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("dqsub670")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.9"))
	if err != nil {
		lang.SayString("dqsub671")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqsub671")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.1"))
	if err != nil {
		lang.SayString("dqsub672")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("dqsub672")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.5"))
	if err != nil {
		lang.SayString("dqsub673")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dqsub673")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.9"))
	if err != nil {
		lang.SayString("dqsub674")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("dqsub674")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.1"))
	if err != nil {
		lang.SayString("dqsub680")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("dqsub680")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.9"))
	if err != nil {
		lang.SayString("dqsub681")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqsub681")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.1"))
	if err != nil {
		lang.SayString("dqsub682")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("dqsub682")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.5"))
	if err != nil {
		lang.SayString("dqsub683")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dqsub683")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.9"))
	if err != nil {
		lang.SayString("dqsub684")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("dqsub684")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456787.1"))
	if err != nil {
		lang.SayString("dqsub685")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("dqsub685")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456787.9"))
	if err != nil {
		lang.SayString("dqsub686")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqsub686")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.1"))
	if err != nil {
		lang.SayString("dqsub687")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("dqsub687")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.5"))
	if err != nil {
		lang.SayString("dqsub688")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dqsub688")
		}
	}
	rexsult, err = lang.RxFromString("1234567456788").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.9"))
	if err != nil {
		lang.SayString("dqsub689")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("dqsub689")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.1"))
	if err != nil {
		lang.SayString("dqsub690")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("dqsub690")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456788.9"))
	if err != nil {
		lang.SayString("dqsub691")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqsub691")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.1"))
	if err != nil {
		lang.SayString("dqsub692")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("dqsub692")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.5"))
	if err != nil {
		lang.SayString("dqsub693")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dqsub693")
		}
	}
	rexsult, err = lang.RxFromString("1234567456789").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1234567456789.9"))
	if err != nil {
		lang.SayString("dqsub694")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("dqsub694")
		}
	}
	rexsult, err = lang.RxFromString("2.E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.002"))
	if err != nil {
		lang.SayString("dqsub901")
	} else {
		if !(rexsult.ToString() == "-1.000") {
			lang.SayString("dqsub901")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.002"))
	if err != nil {
		lang.SayString("dqsub902")
	} else {
		if !(rexsult.ToString() == "-1.0000") {
			lang.SayString("dqsub902")
		}
	}
	rexsult, err = lang.RxFromString("2.00E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.0020"))
	if err != nil {
		lang.SayString("dqsub903")
	} else {
		if !(rexsult.ToString() == "-1.00000") {
			lang.SayString("dqsub903")
		}
	}
	rexsult, err = lang.RxFromString("2.000E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00200"))
	if err != nil {
		lang.SayString("dqsub904")
	} else {
		if !(rexsult.ToString() == "-1.000000") {
			lang.SayString("dqsub904")
		}
	}
	rexsult, err = lang.RxFromString("2.0000E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.002000"))
	if err != nil {
		lang.SayString("dqsub905")
	} else {
		if !(rexsult.ToString() == "-1.0000000") {
			lang.SayString("dqsub905")
		}
	}
	rexsult, err = lang.RxFromString("2.00000E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.0020000"))
	if err != nil {
		lang.SayString("dqsub906")
	} else {
		if !(rexsult.ToString() == "-1.00000000") {
			lang.SayString("dqsub906")
		}
	}
	rexsult, err = lang.RxFromString("2.000000E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.00200000"))
	if err != nil {
		lang.SayString("dqsub907")
	} else {
		if !(rexsult.ToString() == "-1.000000000") {
			lang.SayString("dqsub907")
		}
	}
	rexsult, err = lang.RxFromString("2.0000000E-3").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("1.002000000"))
	if err != nil {
		lang.SayString("dqsub908")
	} else {
		if !(rexsult.ToString() == "-1.0000000000") {
			lang.SayString("dqsub908")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("120E-2"))
	if err != nil {
		lang.SayString("dqsub1125")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("dqsub1125")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("12E-1"))
	if err != nil {
		lang.SayString("dqsub1126")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("dqsub1126")
		}
	}
	rexsult, err = lang.RxFromString("4953734675913.065314738743322579").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("0218.932010396534371704930714860E+797"))
	if err != nil {
		lang.SayString("subx3001")
	} else {
		if !(rexsult.ToString() == "-2.189320103965343717049307148600E+799") {
			lang.SayString("subx3001")
		}
	}
	rexsult, err = lang.RxFromString("9641.684323386955881595490347910E-844").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-78864532047.12287484430980636798E+934"))
	if err != nil {
		lang.SayString("subx3002")
	} else {
		if !(rexsult.ToString() == "7.886453204712287484430980636798E+944") {
			lang.SayString("subx3002")
		}
	}
	rexsult, err = lang.RxFromString("-1.028048571628326871054964307774E+529").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("49200008645699.35577937582714739"))
	if err != nil {
		lang.SayString("subx3003")
	} else {
		if !(rexsult.ToString() == "-1.028048571628326871054964307774E+529") {
			lang.SayString("subx3003")
		}
	}
	rexsult, err = lang.RxFromString("479084.8561808930525417735205519").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("084157571054.2691784660983989931"))
	if err != nil {
		lang.SayString("subx3004")
	} else {
		if !(rexsult.ToString() == "-84157091969.41299757304585721958") {
			lang.SayString("subx3004")
		}
	}
	rexsult, err = lang.RxFromString("-0363750788.573782205664349562931").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-3172.080934464133691909905980096"))
	if err != nil {
		lang.SayString("subx3005")
	} else {
		if !(rexsult.ToString() == "-363747616.4928477415306576530250") {
			lang.SayString("subx3005")
		}
	}
	rexsult, err = lang.RxFromString("1381026551423669919010191878449").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-82.66614775445371254999357800739"))
	if err != nil {
		lang.SayString("subx3006")
	} else {
		if !(rexsult.ToString() == "1381026551423669919010191878532") {
			lang.SayString("subx3006")
		}
	}
	rexsult, err = lang.RxFromString("4627.026960423072127953556635585").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-4410583132901.830017479741231131"))
	if err != nil {
		lang.SayString("subx3007")
	} else {
		if !(rexsult.ToString() == "4410583137528.856977902813359085") {
			lang.SayString("subx3007")
		}
	}
	rexsult, err = lang.RxFromString("75353574493.84484153484918212042").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-8684111695095849922263044191221"))
	if err != nil {
		lang.SayString("subx3008")
	} else {
		if !(rexsult.ToString() == "8684111695095849922338397765715") {
			lang.SayString("subx3008")
		}
	}
	rexsult, err = lang.RxFromString("6907058.216435355874729592373011").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("2.857005446917670515662398741545"))
	if err != nil {
		lang.SayString("subx3009")
	} else {
		if !(rexsult.ToString() == "6907055.359429908957059076710612") {
			lang.SayString("subx3009")
		}
	}
	rexsult, err = lang.RxFromString("-38949530427253.24030680468677190").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("712168021.1265384466442576619064E-992"))
	if err != nil {
		lang.SayString("subx3010")
	} else {
		if !(rexsult.ToString() == "-38949530427253.24030680468677190") {
			lang.SayString("subx3010")
		}
	}
	rexsult, err = lang.RxFromString("-0708069.025667471996378081482549").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-562842.4701520787831018732202804"))
	if err != nil {
		lang.SayString("subx3011")
	} else {
		if !(rexsult.ToString() == "-145226.5555153932132762082622686") {
			lang.SayString("subx3011")
		}
	}
	rexsult, err = lang.RxFromString("4055087.246994644709729942673976").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-43183146921897.67383476104084575E+211"))
	if err != nil {
		lang.SayString("subx3012")
	} else {
		if !(rexsult.ToString() == "4.318314692189767383476104084575E+224") {
			lang.SayString("subx3012")
		}
	}
	rexsult, err = lang.RxFromString("4502895892520.396581348110906909E-512").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-815.9047305921862348263521876034"))
	if err != nil {
		lang.SayString("subx3013")
	} else {
		if !(rexsult.ToString() == "815.9047305921862348263521876034") {
			lang.SayString("subx3013")
		}
	}
	rexsult, err = lang.RxFromString("467.6721295072628100260239179865").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-02.07155073395573569852316073025"))
	if err != nil {
		lang.SayString("subx3014")
	} else {
		if !(rexsult.ToString() == "469.7436802412185457245470787168") {
			lang.SayString("subx3014")
		}
	}
	rexsult, err = lang.RxFromString("2.156795313311150143949997552501E-571").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-8677147.586389401682712180146855"))
	if err != nil {
		lang.SayString("subx3015")
	} else {
		if !(rexsult.ToString() == "8677147.586389401682712180146855") {
			lang.SayString("subx3015")
		}
	}
	rexsult, err = lang.RxFromString("-974953.2801637208368002585822457").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-693095793.3667578067802698191246"))
	if err != nil {
		lang.SayString("subx3016")
	} else {
		if !(rexsult.ToString() == "692120840.0865940859434695605424") {
			lang.SayString("subx3016")
		}
	}
	rexsult, err = lang.RxFromString("-7634680140009571846155654339781").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("3009630949502.035852433434214413E-490"))
	if err != nil {
		lang.SayString("subx3017")
	} else {
		if !(rexsult.ToString() == "-7634680140009571846155654339781") {
			lang.SayString("subx3017")
		}
	}
	rexsult, err = lang.RxFromString("262273.0222851186523650889896428E-624").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("74177.21073338090843145838835480"))
	if err != nil {
		lang.SayString("subx3018")
	} else {
		if !(rexsult.ToString() == "-74177.21073338090843145838835480") {
			lang.SayString("subx3018")
		}
	}
	rexsult, err = lang.RxFromString("-8036052748815903177624716581732").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-066677357.4438809548850966167573"))
	if err != nil {
		lang.SayString("subx3019")
	} else {
		if !(rexsult.ToString() == "-8036052748815903177624649904375") {
			lang.SayString("subx3019")
		}
	}
	rexsult, err = lang.RxFromString("883429.5928031498103637713570166E+765").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-43978.97283712939198111043032726"))
	if err != nil {
		lang.SayString("subx3020")
	} else {
		if !(rexsult.ToString() == "8.834295928031498103637713570166E+770") {
			lang.SayString("subx3020")
		}
	}
	rexsult, err = lang.RxFromString("24791301060.37938360567775506973").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-5613327866480.322649080205877564"))
	if err != nil {
		lang.SayString("subx3021")
	} else {
		if !(rexsult.ToString() == "5638119167540.702032685883632634") {
			lang.SayString("subx3021")
		}
	}
	rexsult, err = lang.RxFromString("-930711443.9474781586162910776139").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-740.3860979292775472622798348030"))
	if err != nil {
		lang.SayString("subx3022")
	} else {
		if !(rexsult.ToString() == "-930710703.5613802293387438153341") {
			lang.SayString("subx3022")
		}
	}
	rexsult, err = lang.RxFromString("2358276428765.064191082773385539").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("214.3589796082328665878602304469"))
	if err != nil {
		lang.SayString("subx3023")
	} else {
		if !(rexsult.ToString() == "2358276428550.705211474540518951") {
			lang.SayString("subx3023")
		}
	}
	rexsult, err = lang.RxFromString("-3.868744449795653651638308926987E+750").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("8270.472492965559872384018329418"))
	if err != nil {
		lang.SayString("subx3024")
	} else {
		if !(rexsult.ToString() == "-3.868744449795653651638308926987E+750") {
			lang.SayString("subx3024")
		}
	}
	rexsult, err = lang.RxFromString("140422069.5863246490180206814374E-447").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-567195652586.2454217069003186487"))
	if err != nil {
		lang.SayString("subx3025")
	} else {
		if !(rexsult.ToString() == "567195652586.2454217069003186487") {
			lang.SayString("subx3025")
		}
	}
	rexsult, err = lang.RxFromString("75929096475.63450425339472559646E+153").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-0945260193.503803519572604151290E+459"))
	if err != nil {
		lang.SayString("subx3026")
	} else {
		if !(rexsult.ToString() == "9.452601935038035195726041512900E+467") {
			lang.SayString("subx3026")
		}
	}
	rexsult, err = lang.RxFromString("6312318309.142044953357460463732").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-5641317823.202274083982487558514E+628"))
	if err != nil {
		lang.SayString("subx3027")
	} else {
		if !(rexsult.ToString() == "5.641317823202274083982487558514E+637") {
			lang.SayString("subx3027")
		}
	}
	rexsult, err = lang.RxFromString("93793652428100.52105928239469937").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("917.2571313109730433369594936416E-712"))
	if err != nil {
		lang.SayString("subx3028")
	} else {
		if !(rexsult.ToString() == "93793652428100.52105928239469937") {
			lang.SayString("subx3028")
		}
	}
	rexsult, err = lang.RxFromString("98471198160.56524417578665886060").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-23994.14313393939743548945165462"))
	if err != nil {
		lang.SayString("subx3029")
	} else {
		if !(rexsult.ToString() == "98471222154.70837811518409435005") {
			lang.SayString("subx3029")
		}
	}
	rexsult, err = lang.RxFromString("329326552.0208398002250836592043").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-02451.10065397010591546041034041"))
	if err != nil {
		lang.SayString("subx3030")
	} else {
		if !(rexsult.ToString() == "329329003.1214937703309991196146") {
			lang.SayString("subx3030")
		}
	}
	rexsult, err = lang.RxFromString("-92980.68431371090354435763218439").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-2282178507046019721925800997065"))
	if err != nil {
		lang.SayString("subx3031")
	} else {
		if !(rexsult.ToString() == "2282178507046019721925800904084") {
			lang.SayString("subx3031")
		}
	}
	rexsult, err = lang.RxFromString("12135817762.27858606259822256987E+738").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("98.35649167872356132249561021910E-902"))
	if err != nil {
		lang.SayString("subx3032")
	} else {
		if !(rexsult.ToString() == "1.213581776227858606259822256987E+748") {
			lang.SayString("subx3032")
		}
	}
	rexsult, err = lang.RxFromString("37.27457578793521166809739140081").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-392550.4790095035979998355569916"))
	if err != nil {
		lang.SayString("subx3033")
	} else {
		if !(rexsult.ToString() == "392587.7535852915332115036543830") {
			lang.SayString("subx3033")
		}
	}
	rexsult, err = lang.RxFromString("-2787.980590304199878755265273703").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("7117631179305319208210387565324"))
	if err != nil {
		lang.SayString("subx3034")
	} else {
		if !(rexsult.ToString() == "-7117631179305319208210387568112") {
			lang.SayString("subx3034")
		}
	}
	rexsult, err = lang.RxFromString("-9890633.854609434943559831911276E+971").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-1939985729.436827777055699361237"))
	if err != nil {
		lang.SayString("subx3035")
	} else {
		if !(rexsult.ToString() == "-9.890633854609434943559831911276E+977") {
			lang.SayString("subx3035")
		}
	}
	rexsult, err = lang.RxFromString("3944570323.331121750661920475191").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-17360722.28878145641394962484366"))
	if err != nil {
		lang.SayString("subx3036")
	} else {
		if !(rexsult.ToString() == "3961931045.619903207075870100035") {
			lang.SayString("subx3036")
		}
	}
	rexsult, err = lang.RxFromString("19544.14018503427029002552872707").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("1786697762.885178994182133839546"))
	if err != nil {
		lang.SayString("subx3037")
	} else {
		if !(rexsult.ToString() == "-1786678218.744993959911843814017") {
			lang.SayString("subx3037")
		}
	}
	rexsult, err = lang.RxFromString("-05.75485957937617757983513662981").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("5564476875.989640431173694372083"))
	if err != nil {
		lang.SayString("subx3038")
	} else {
		if !(rexsult.ToString() == "-5564476881.744500010549871951918") {
			lang.SayString("subx3038")
		}
	}
	rexsult, err = lang.RxFromString("-4208820.898718069194008526302746").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("626887.7553774705678201112845462E+206"))
	if err != nil {
		lang.SayString("subx3039")
	} else {
		if !(rexsult.ToString() == "-6.268877553774705678201112845462E+211") {
			lang.SayString("subx3039")
		}
	}
	rexsult, err = lang.RxFromString("-70077195478066.30896979085821269E+549").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("4607.163248554155483681430013073"))
	if err != nil {
		lang.SayString("subx3040")
	} else {
		if !(rexsult.ToString() == "-7.007719547806630896979085821269E+562") {
			lang.SayString("subx3040")
		}
	}
	rexsult, err = lang.RxFromString("-442941.7541811527940918244383454").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-068126768.0563559819156379151016"))
	if err != nil {
		lang.SayString("subx3041")
	} else {
		if !(rexsult.ToString() == "67683826.30217482912154609066326") {
			lang.SayString("subx3041")
		}
	}
	rexsult, err = lang.RxFromString("-040726778711.8677615616711676159").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("299691.9430345259174614997064916"))
	if err != nil {
		lang.SayString("subx3042")
	} else {
		if !(rexsult.ToString() == "-40727078403.81079608758862911561") {
			lang.SayString("subx3042")
		}
	}
	rexsult, err = lang.RxFromString("-1934197520.738366912179143085955").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("3.810751422515178400293693371519"))
	if err != nil {
		lang.SayString("subx3043")
	} else {
		if !(rexsult.ToString() == "-1934197524.549118334694321486249") {
			lang.SayString("subx3043")
		}
	}
	rexsult, err = lang.RxFromString("813262.7723533833038829559646830").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-303284822716.8282178131118185907"))
	if err != nil {
		lang.SayString("subx3044")
	} else {
		if !(rexsult.ToString() == "303285635979.6005711964157015467") {
			lang.SayString("subx3044")
		}
	}
	rexsult, err = lang.RxFromString("36105954884.94621434979365589311").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("745558205.7692397481313005659523E-952"))
	if err != nil {
		lang.SayString("subx3045")
	} else {
		if !(rexsult.ToString() == "36105954884.94621434979365589311") {
			lang.SayString("subx3045")
		}
	}
	rexsult, err = lang.RxFromString("-075537177538.1814516621962185490").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("26980775255.51542856483122484898"))
	if err != nil {
		lang.SayString("subx3046")
	} else {
		if !(rexsult.ToString() == "-102517952793.6968802270274433980") {
			lang.SayString("subx3046")
		}
	}
	rexsult, err = lang.RxFromString("-4223765.415319564898840040697647").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-2590590305497454185455459149918E-215"))
	if err != nil {
		lang.SayString("subx3047")
	} else {
		if !(rexsult.ToString() == "-4223765.415319564898840040697647") {
			lang.SayString("subx3047")
		}
	}
	rexsult, err = lang.RxFromString("-6468.903738522951259063099946195").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-7877.324314273694312164407794939E+267"))
	if err != nil {
		lang.SayString("subx3048")
	} else {
		if !(rexsult.ToString() == "7.877324314273694312164407794939E+270") {
			lang.SayString("subx3048")
		}
	}
	rexsult, err = lang.RxFromString("-9567221.183663236817239254783372E-203").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("1650.198961256061165362319471264"))
	if err != nil {
		lang.SayString("subx3049")
	} else {
		if !(rexsult.ToString() == "-1650.198961256061165362319471264") {
			lang.SayString("subx3049")
		}
	}
	rexsult, err = lang.RxFromString("8812306098770.200752139142033569E-428").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("26790.17380163975186972720427030E+568"))
	if err != nil {
		lang.SayString("subx3050")
	} else {
		if !(rexsult.ToString() == "-2.679017380163975186972720427030E+572") {
			lang.SayString("subx3050")
		}
	}
	rexsult, err = lang.RxFromString("80108033.12724838718736922500904").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-706207255092.7645192310078892869"))
	if err != nil {
		lang.SayString("subx3051")
	} else {
		if !(rexsult.ToString() == "706287363125.8917676181952585119") {
			lang.SayString("subx3051")
		}
	}
	rexsult, err = lang.RxFromString("-37942846282.76101663789059003505").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-5.649456053942850351313869983197"))
	if err != nil {
		lang.SayString("subx3052")
	} else {
		if !(rexsult.ToString() == "-37942846277.11156058394773968374") {
			lang.SayString("subx3052")
		}
	}
	rexsult, err = lang.RxFromString("92659632115305.13735437728445541").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("6483438.317862851676468094261410E-139"))
	if err != nil {
		lang.SayString("subx3053")
	} else {
		if !(rexsult.ToString() == "92659632115305.13735437728445541") {
			lang.SayString("subx3053")
		}
	}
	rexsult, err = lang.RxFromString("2838948.589837595494152150647194").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("569547026247.5469563701415715960"))
	if err != nil {
		lang.SayString("subx3054")
	} else {
		if !(rexsult.ToString() == "-569544187298.9571187746474194454") {
			lang.SayString("subx3054")
		}
	}
	rexsult, err = lang.RxFromString("524995204523.6053307941775794287E+694").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("1589600879689517100527293028553"))
	if err != nil {
		lang.SayString("subx3055")
	} else {
		if !(rexsult.ToString() == "5.249952045236053307941775794287E+705") {
			lang.SayString("subx3055")
		}
	}
	rexsult, err = lang.RxFromString("-57131573677452.15449921725097290").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("4669681430736.326858508715643769"))
	if err != nil {
		lang.SayString("subx3056")
	} else {
		if !(rexsult.ToString() == "-61801255108188.48135772596661667") {
			lang.SayString("subx3056")
		}
	}
	rexsult, err = lang.RxFromString("90794826.55528018781830463383411").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-5.471502270351231110027647216128"))
	if err != nil {
		lang.SayString("subx3057")
	} else {
		if !(rexsult.ToString() == "90794832.02678245816953574386176") {
			lang.SayString("subx3057")
		}
	}
	rexsult, err = lang.RxFromString("58508794729.35191160840980489138").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-47060867.24988279680824397447551"))
	if err != nil {
		lang.SayString("subx3058")
	} else {
		if !(rexsult.ToString() == "58555855596.60179440521804886586") {
			lang.SayString("subx3058")
		}
	}
	rexsult, err = lang.RxFromString("-746104.0768078474426464219416332E+006").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("9595418.300613754556671852801667E+385"))
	if err != nil {
		lang.SayString("subx3059")
	} else {
		if !(rexsult.ToString() == "-9.595418300613754556671852801667E+391") {
			lang.SayString("subx3059")
		}
	}
	rexsult, err = lang.RxFromString("55.99427632688387400403789659459E+119").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-9.170530450881612853998489340127"))
	if err != nil {
		lang.SayString("subx3060")
	} else {
		if !(rexsult.ToString() == "5.599427632688387400403789659459E+120") {
			lang.SayString("subx3060")
		}
	}
	rexsult, err = lang.RxFromString("-41214265628.83801241467317270595").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("1015336323798389903361978271354"))
	if err != nil {
		lang.SayString("subx3061")
	} else {
		if !(rexsult.ToString() == "-1015336323798389903403192536983") {
			lang.SayString("subx3061")
		}
	}
	rexsult, err = lang.RxFromString("89937.39749201095570357557430822").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("82351554210093.60879476027800331"))
	if err != nil {
		lang.SayString("subx3062")
	} else {
		if !(rexsult.ToString() == "-82351554120156.21130274932229974") {
			lang.SayString("subx3062")
		}
	}
	rexsult, err = lang.RxFromString("01712661.64677082156284125486943E+359").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("57932.78435529483241552042115837E-037"))
	if err != nil {
		lang.SayString("subx3063")
	} else {
		if !(rexsult.ToString() == "1.712661646770821562841254869430E+365") {
			lang.SayString("subx3063")
		}
	}
	rexsult, err = lang.RxFromString("-2647593306.528617691373470059213").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-655531558709.4582168930191014461"))
	if err != nil {
		lang.SayString("subx3064")
	} else {
		if !(rexsult.ToString() == "652883965402.9295992016456313869") {
			lang.SayString("subx3064")
		}
	}
	rexsult, err = lang.RxFromString("2904078690665765116603253099668E-329").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-71.45586619176091599264717047885E+787"))
	if err != nil {
		lang.SayString("subx3065")
	} else {
		if !(rexsult.ToString() == "7.145586619176091599264717047885E+788") {
			lang.SayString("subx3065")
		}
	}
	rexsult, err = lang.RxFromString("22094338972.39109726522477999515").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-409846549371.3900805039668417203E-499"))
	if err != nil {
		lang.SayString("subx3066")
	} else {
		if !(rexsult.ToString() == "22094338972.39109726522477999515") {
			lang.SayString("subx3066")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("subx3067")
	} else {
		if !(rexsult.ToString() == "-3374988581607586061337836096173") {
			lang.SayString("subx3067")
		}
	}
	rexsult, err = lang.RxFromString("-84172558160661.35863831029352323").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-11271.58916600931155937291904890"))
	if err != nil {
		lang.SayString("subx3068")
	} else {
		if !(rexsult.ToString() == "-84172558149389.76947230098196386") {
			lang.SayString("subx3068")
		}
	}
	rexsult, err = lang.RxFromString("-70046932324614.90596396237508541E-568").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("33.63163964004608865836577297698E-918"))
	if err != nil {
		lang.SayString("subx3069")
	} else {
		if !(rexsult.ToString() == "-7.004693232461490596396237508541E-555") {
			lang.SayString("subx3069")
		}
	}
	rexsult, err = lang.RxFromString("0004125384407.053782660115680886").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-391429084.5847321402514385603223E-648"))
	if err != nil {
		lang.SayString("subx3070")
	} else {
		if !(rexsult.ToString() == "4125384407.053782660115680886000") {
			lang.SayString("subx3070")
		}
	}
	rexsult, err = lang.RxFromString("-31823131.15691583393820628480997E-440").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("92913.91582947237200286427030028E+771"))
	if err != nil {
		lang.SayString("subx3071")
	} else {
		if !(rexsult.ToString() == "-9.291391582947237200286427030028E+775") {
			lang.SayString("subx3071")
		}
	}
	rexsult, err = lang.RxFromString("55573867888.91575330563698128150").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("599.5231614736232188354393212234"))
	if err != nil {
		lang.SayString("subx3072")
	} else {
		if !(rexsult.ToString() == "55573867289.39259183201376244606") {
			lang.SayString("subx3072")
		}
	}
	rexsult, err = lang.RxFromString("-5447727448431680878699555714796E-800").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("5487207.142687001607026665515349E-362"))
	if err != nil {
		lang.SayString("subx3073")
	} else {
		if !(rexsult.ToString() == "-5.487207142687001607026665515349E-356") {
			lang.SayString("subx3073")
		}
	}
	rexsult, err = lang.RxFromString("0418349404834.547110239542290134").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("09819915.92405288066606124554841"))
	if err != nil {
		lang.SayString("subx3074")
	} else {
		if !(rexsult.ToString() == "418339584918.6230573588762288885") {
			lang.SayString("subx3074")
		}
	}
	rexsult, err = lang.RxFromString("-262021.7565194737396448014286436").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-7983992600094836304387324162042E+390"))
	if err != nil {
		lang.SayString("subx3075")
	} else {
		if !(rexsult.ToString() == "7.983992600094836304387324162042E+420") {
			lang.SayString("subx3075")
		}
	}
	rexsult, err = lang.RxFromString("48696050631.42565380301204592392E-505").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-33868752339.85057267609967806187E+821"))
	if err != nil {
		lang.SayString("subx3076")
	} else {
		if !(rexsult.ToString() == "3.386875233985057267609967806187E+831") {
			lang.SayString("subx3076")
		}
	}
	rexsult, err = lang.RxFromString("95316999.19440144356471126680708").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-60791.33805057402845885978390435"))
	if err != nil {
		lang.SayString("subx3077")
	} else {
		if !(rexsult.ToString() == "95377790.53245201759317012659098") {
			lang.SayString("subx3077")
		}
	}
	rexsult, err = lang.RxFromString("-5326702296402708234722215224979E-136").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("8032459.450998820205916538543258"))
	if err != nil {
		lang.SayString("subx3078")
	} else {
		if !(rexsult.ToString() == "-8032459.450998820205916538543258") {
			lang.SayString("subx3078")
		}
	}
	rexsult, err = lang.RxFromString("67.18750684079501575335482615780E-281").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("734.1168841683438410314843011541E-854"))
	if err != nil {
		lang.SayString("subx3079")
	} else {
		if !(rexsult.ToString() == "6.718750684079501575335482615780E-280") {
			lang.SayString("subx3079")
		}
	}
	rexsult, err = lang.RxFromString("-8739299372114.092482914139281669").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("507610074.7343577029345077385838"))
	if err != nil {
		lang.SayString("subx3080")
	} else {
		if !(rexsult.ToString() == "-8739806982188.826840617073789408") {
			lang.SayString("subx3080")
		}
	}
	rexsult, err = lang.RxFromString("2454.002078468928665008217727731").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("583546039.6233842869119950982009E-147"))
	if err != nil {
		lang.SayString("subx3081")
	} else {
		if !(rexsult.ToString() == "2454.002078468928665008217727731") {
			lang.SayString("subx3081")
		}
	}
	rexsult, err = lang.RxFromString("764578.5204849936912066033177429").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("64603.13571259164812609436832506"))
	if err != nil {
		lang.SayString("subx3082")
	} else {
		if !(rexsult.ToString() == "699975.3847724020430805089494178") {
			lang.SayString("subx3082")
		}
	}
	rexsult, err = lang.RxFromString("079203.7330103777716903518367560").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("846388934347.6324036132959664705"))
	if err != nil {
		lang.SayString("subx3083")
	} else {
		if !(rexsult.ToString() == "-846388855143.8993932355242761187") {
			lang.SayString("subx3083")
		}
	}
	rexsult, err = lang.RxFromString("-4278.581514688669249247007127899E-329").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("5.474973992953902631890208360829"))
	if err != nil {
		lang.SayString("subx3084")
	} else {
		if !(rexsult.ToString() == "-5.474973992953902631890208360829") {
			lang.SayString("subx3084")
		}
	}
	rexsult, err = lang.RxFromString("60867019.81764798845468445196869E+651").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("6.149612565404080501157093851895E+817"))
	if err != nil {
		lang.SayString("subx3085")
	} else {
		if !(rexsult.ToString() == "-6.149612565404080501157093851895E+817") {
			lang.SayString("subx3085")
		}
	}
	rexsult, err = lang.RxFromString("18554417738217.62218590965803605E-382").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-0894505909529.052378474618435782E+527"))
	if err != nil {
		lang.SayString("subx3086")
	} else {
		if !(rexsult.ToString() == "8.945059095290523784746184357820E+538") {
			lang.SayString("subx3086")
		}
	}
	rexsult, err = lang.RxFromString("69073355517144.36356688642213839").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("997784782535.6104634823627327033E+116"))
	if err != nil {
		lang.SayString("subx3087")
	} else {
		if !(rexsult.ToString() == "-9.977847825356104634823627327033E+127") {
			lang.SayString("subx3087")
		}
	}
	rexsult, err = lang.RxFromString("450282259072.8657099359104277477").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-1791307965314309175477911369824"))
	if err != nil {
		lang.SayString("subx3088")
	} else {
		if !(rexsult.ToString() == "1791307965314309175928193628897") {
			lang.SayString("subx3088")
		}
	}
	rexsult, err = lang.RxFromString("954678411.7838149266455177850037").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("142988.7096204254529284334278794"))
	if err != nil {
		lang.SayString("subx3089")
	} else {
		if !(rexsult.ToString() == "954535423.0741945011925893515758") {
			lang.SayString("subx3089")
		}
	}
	rexsult, err = lang.RxFromString("-9244530976.220812127155852389807E+557").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("541089.4715446858896619078627941"))
	if err != nil {
		lang.SayString("subx3090")
	} else {
		if !(rexsult.ToString() == "-9.244530976220812127155852389807E+566") {
			lang.SayString("subx3090")
		}
	}
	rexsult, err = lang.RxFromString("-75492024.20990197005974241975449").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-14760421311348.35269044633000927"))
	if err != nil {
		lang.SayString("subx3091")
	} else {
		if !(rexsult.ToString() == "14760345819324.14278847627026685") {
			lang.SayString("subx3091")
		}
	}
	rexsult, err = lang.RxFromString("317747.6972215715434186596178036E-452").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("24759763.33144824613591228097330E+092"))
	if err != nil {
		lang.SayString("subx3092")
	} else {
		if !(rexsult.ToString() == "-2.475976333144824613591228097330E+99") {
			lang.SayString("subx3092")
		}
	}
	rexsult, err = lang.RxFromString("-8.153334430358647134334545353427").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-9.717872025814596548462854853522"))
	if err != nil {
		lang.SayString("subx3093")
	} else {
		if !(rexsult.ToString() == "1.564537595455949414128309500095") {
			lang.SayString("subx3093")
		}
	}
	rexsult, err = lang.RxFromString("7.267345197492967332320456062961E-478").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("5054015481833.263541189916208065"))
	if err != nil {
		lang.SayString("subx3094")
	} else {
		if !(rexsult.ToString() == "-5054015481833.263541189916208065") {
			lang.SayString("subx3094")
		}
	}
	rexsult, err = lang.RxFromString("-1223354029.862567054230912271171").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("8135774223401322785475014855625"))
	if err != nil {
		lang.SayString("subx3095")
	} else {
		if !(rexsult.ToString() == "-8135774223401322785476238209655") {
			lang.SayString("subx3095")
		}
	}
	rexsult, err = lang.RxFromString("285397644111.5655679961211349982E+645").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-2479499427613157519359627280704"))
	if err != nil {
		lang.SayString("subx3096")
	} else {
		if !(rexsult.ToString() == "2.853976441115655679961211349982E+656") {
			lang.SayString("subx3096")
		}
	}
	rexsult, err = lang.RxFromString("-4673112.663442366293812346653429").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("-3429.998403142546001438238460958"))
	if err != nil {
		lang.SayString("subx3097")
	} else {
		if !(rexsult.ToString() == "-4669682.665039223747810908414968") {
			lang.SayString("subx3097")
		}
	}
	rexsult, err = lang.RxFromString("88.96492479681278079861456051103").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("386939.4621006514751889096510923E+139"))
	if err != nil {
		lang.SayString("subx3098")
	} else {
		if !(rexsult.ToString() == "-3.869394621006514751889096510923E+144") {
			lang.SayString("subx3098")
		}
	}
	rexsult, err = lang.RxFromString("064326846.4286437304788069444326E-942").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("92.23649942010862087149015091350"))
	if err != nil {
		lang.SayString("subx3099")
	} else {
		if !(rexsult.ToString() == "-92.23649942010862087149015091350") {
			lang.SayString("subx3099")
		}
	}
	rexsult, err = lang.RxFromString("504507.0043949324433170405699360").OpSub(lang.RxSetWithDigit(31), lang.RxFromString("605387.7175522955344659311072099"))
	if err != nil {
		lang.SayString("subx3100")
	} else {
		if !(rexsult.ToString() == "-100880.7131573630911488905372739") {
			lang.SayString("subx3100")
		}
	}
	rexsult, err = lang.RxFromString("1.5283550163839789319142430253644").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-1.6578158484822969520405291379492"))
	if err != nil {
		lang.SayString("subx3201")
	} else {
		if !(rexsult.ToString() == "3.1861708648662758839547721633136") {
			lang.SayString("subx3201")
		}
	}
	rexsult, err = lang.RxFromString("-622903030605.2867503937836507326").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("6519388607.1331855704471328795821"))
	if err != nil {
		lang.SayString("subx3202")
	} else {
		if !(rexsult.ToString() == "-629422419212.41993596423078361218") {
			lang.SayString("subx3202")
		}
	}
	rexsult, err = lang.RxFromString("-5675915.2465457487632250245209054").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("73913909880.381367895205086027416"))
	if err != nil {
		lang.SayString("subx3203")
	} else {
		if !(rexsult.ToString() == "-73919585795.627913643968311051937") {
			lang.SayString("subx3203")
		}
	}
	rexsult, err = lang.RxFromString("97.647321172555144900685605318635").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("4.8620911587547548751209841570885"))
	if err != nil {
		lang.SayString("subx3204")
	} else {
		if !(rexsult.ToString() == "92.785230013800390025564621161547") {
			lang.SayString("subx3204")
		}
	}
	rexsult, err = lang.RxFromString("-9717253267024.5380651435435603552").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-2669.2539695193820424002013488480E+363"))
	if err != nil {
		lang.SayString("subx3205")
	} else {
		if !(rexsult.ToString() == "2.6692539695193820424002013488480E+366") {
			lang.SayString("subx3205")
		}
	}
	rexsult, err = lang.RxFromString("-4.0817391717190128506083001702335E-767").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("12772.807105920712660991033689206"))
	if err != nil {
		lang.SayString("subx3206")
	} else {
		if !(rexsult.ToString() == "-12772.807105920712660991033689206") {
			lang.SayString("subx3206")
		}
	}
	rexsult, err = lang.RxFromString("68625322655934146845003028928647").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-59.634169944840280159782488098700"))
	if err != nil {
		lang.SayString("subx3207")
	} else {
		if !(rexsult.ToString() == "68625322655934146845003028928707") {
			lang.SayString("subx3207")
		}
	}
	rexsult, err = lang.RxFromString("732515.76532049290815665856727641").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-92134479835821.319619827023729829"))
	if err != nil {
		lang.SayString("subx3208")
	} else {
		if !(rexsult.ToString() == "92134480568337.084940319931886488") {
			lang.SayString("subx3208")
		}
	}
	rexsult, err = lang.RxFromString("-30.458011942978338421676454778733").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-5023372024597665102336430410403E+831"))
	if err != nil {
		lang.SayString("subx3209")
	} else {
		if !(rexsult.ToString() == "5.0233720245976651023364304104030E+861") {
			lang.SayString("subx3209")
		}
	}
	rexsult, err = lang.RxFromString("-89640.094149414644660480286430462").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-58703419758.800889903227509215474"))
	if err != nil {
		lang.SayString("subx3210")
	} else {
		if !(rexsult.ToString() == "58703330118.706740488582848735188") {
			lang.SayString("subx3210")
		}
	}
	rexsult, err = lang.RxFromString("458653.1567870081810052917714259").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("18353106238.516235116080449814053E-038"))
	if err != nil {
		lang.SayString("subx3211")
	} else {
		if !(rexsult.ToString() == "458653.15678700818100529177142590") {
			lang.SayString("subx3211")
		}
	}
	rexsult, err = lang.RxFromString("913391.42744224458216174967853722").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-21051638.816432817393202262710630E+432"))
	if err != nil {
		lang.SayString("subx3212")
	} else {
		if !(rexsult.ToString() == "2.1051638816432817393202262710630E+439") {
			lang.SayString("subx3212")
		}
	}
	rexsult, err = lang.RxFromString("-917591456829.12109027484399536567").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-28892177726858026955513438843371E+708"))
	if err != nil {
		lang.SayString("subx3213")
	} else {
		if !(rexsult.ToString() == "2.8892177726858026955513438843371E+739") {
			lang.SayString("subx3213")
		}
	}
	rexsult, err = lang.RxFromString("34938410840645.913399699219228218").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("30.818220393242402846077755480548"))
	if err != nil {
		lang.SayString("subx3214")
	} else {
		if !(rexsult.ToString() == "34938410840615.095179305976825372") {
			lang.SayString("subx3214")
		}
	}
	rexsult, err = lang.RxFromString("6034.7374411022598081745006769023E-517").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("29771833428054709077850588904653"))
	if err != nil {
		lang.SayString("subx3215")
	} else {
		if !(rexsult.ToString() == "-29771833428054709077850588904653") {
			lang.SayString("subx3215")
		}
	}
	rexsult, err = lang.RxFromString("-5565747671734.1686009705574503152").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-490.30899494881071282787487030303"))
	if err != nil {
		lang.SayString("subx3216")
	} else {
		if !(rexsult.ToString() == "-5565747671243.8596060217467374873") {
			lang.SayString("subx3216")
		}
	}
	rexsult, err = lang.RxFromString("319545511.89203495546689273564728E+036").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-2955943533943321899150310192061"))
	if err != nil {
		lang.SayString("subx3217")
	} else {
		if !(rexsult.ToString() == "3.1954551189203791141042667896918E+44") {
			lang.SayString("subx3217")
		}
	}
	rexsult, err = lang.RxFromString("-36852134.84194296250843579428931").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-5830629.8347085025808716360357940"))
	if err != nil {
		lang.SayString("subx3218")
	} else {
		if !(rexsult.ToString() == "-31021505.007234459927564158253516") {
			lang.SayString("subx3218")
		}
	}
	rexsult, err = lang.RxFromString("8.6021905001798578659275880018221E-374").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-39505285344943.729681835377530908"))
	if err != nil {
		lang.SayString("subx3219")
	} else {
		if !(rexsult.ToString() == "39505285344943.729681835377530908") {
			lang.SayString("subx3219")
		}
	}
	rexsult, err = lang.RxFromString("-54863165.152174109720312887805017").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("736.1398476560169141105328256628"))
	if err != nil {
		lang.SayString("subx3220")
	} else {
		if !(rexsult.ToString() == "-54863901.292021765737226998337843") {
			lang.SayString("subx3220")
		}
	}
	rexsult, err = lang.RxFromString("-3263743464517851012531708810307").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("2457206.2471248382136273643208109"))
	if err != nil {
		lang.SayString("subx3221")
	} else {
		if !(rexsult.ToString() == "-3263743464517851012531711267513.2") {
			lang.SayString("subx3221")
		}
	}
	rexsult, err = lang.RxFromString("2856586744.0548637797291151154902E-895").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("953545637646.57694835860339582821E+080"))
	if err != nil {
		lang.SayString("subx3222")
	} else {
		if !(rexsult.ToString() == "-9.5354563764657694835860339582821E+91") {
			lang.SayString("subx3222")
		}
	}
	rexsult, err = lang.RxFromString("5624157233.3433661009203529937625").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("626098409265.93738586750090160638"))
	if err != nil {
		lang.SayString("subx3223")
	} else {
		if !(rexsult.ToString() == "-620474252032.59401976658054861262") {
			lang.SayString("subx3223")
		}
	}
	rexsult, err = lang.RxFromString("-213499362.91476998701834067092611").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("419272438.02555757699863022643444"))
	if err != nil {
		lang.SayString("subx3224")
	} else {
		if !(rexsult.ToString() == "-632771800.94032756401697089736055") {
			lang.SayString("subx3224")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("subx3225")
	} else {
		if !(rexsult.ToString() == "30264.783154666903063637658139590") {
			lang.SayString("subx3225")
		}
	}
	rexsult, err = lang.RxFromString("47.525676459351505682005359699680E+704").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-58631943508436657383369760970210"))
	if err != nil {
		lang.SayString("subx3226")
	} else {
		if !(rexsult.ToString() == "4.7525676459351505682005359699680E+705") {
			lang.SayString("subx3226")
		}
	}
	rexsult, err = lang.RxFromString("-74396862273800.625679130772935550").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-115616605.52826981284183992013157"))
	if err != nil {
		lang.SayString("subx3227")
	} else {
		if !(rexsult.ToString() == "-74396746657195.097409317931095630") {
			lang.SayString("subx3227")
		}
	}
	rexsult, err = lang.RxFromString("67585560.562561229497720955705979").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("826.96290288608566737177503451613"))
	if err != nil {
		lang.SayString("subx3228")
	} else {
		if !(rexsult.ToString() == "67584733.599658343412053583930945") {
			lang.SayString("subx3228")
		}
	}
	rexsult, err = lang.RxFromString("6877386868.9498051860742298735156E-232").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("390.3154289860643509393769754551"))
	if err != nil {
		lang.SayString("subx3229")
	} else {
		if !(rexsult.ToString() == "-390.31542898606435093937697545510") {
			lang.SayString("subx3229")
		}
	}
	rexsult, err = lang.RxFromString("-1647335.201144609256134925838937").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-186654823782.50459802235024230856"))
	if err != nil {
		lang.SayString("subx3230")
	} else {
		if !(rexsult.ToString() == "186653176447.30345341309410738272") {
			lang.SayString("subx3230")
		}
	}
	rexsult, err = lang.RxFromString("41407818140948.866630923934138155").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-5156.7624534000311342310106671627E-963"))
	if err != nil {
		lang.SayString("subx3231")
	} else {
		if !(rexsult.ToString() == "41407818140948.866630923934138155") {
			lang.SayString("subx3231")
		}
	}
	rexsult, err = lang.RxFromString("-6.6547424012516834662011706165297").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-574454585580.06215974139884746441"))
	if err != nil {
		lang.SayString("subx3232")
	} else {
		if !(rexsult.ToString() == "574454585573.40741734014716399821") {
			lang.SayString("subx3232")
		}
	}
	rexsult, err = lang.RxFromString("-27627.758745381267780885067447671").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-23385972189441.709586433758111062"))
	if err != nil {
		lang.SayString("subx3233")
	} else {
		if !(rexsult.ToString() == "23385972161813.950841052490330177") {
			lang.SayString("subx3233")
		}
	}
	rexsult, err = lang.RxFromString("209819.74379099914752963711944307E-228").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-440230.6700989532467831370320266E-960"))
	if err != nil {
		lang.SayString("subx3234")
	} else {
		if !(rexsult.ToString() == "2.0981974379099914752963711944307E-223") {
			lang.SayString("subx3234")
		}
	}
	rexsult, err = lang.RxFromString("2.3488457600415474270259330865184").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("9182434.6660212482500376220508605E-612"))
	if err != nil {
		lang.SayString("subx3235")
	} else {
		if !(rexsult.ToString() == "2.3488457600415474270259330865184") {
			lang.SayString("subx3235")
		}
	}
	rexsult, err = lang.RxFromString("-5107586300197.9703941034404557409").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("56609.05486055057838678039496686E-768"))
	if err != nil {
		lang.SayString("subx3236")
	} else {
		if !(rexsult.ToString() == "-5107586300197.9703941034404557409") {
			lang.SayString("subx3236")
		}
	}
	rexsult, err = lang.RxFromString("-70454070095869.70717871212601390").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-6200178.370249260009168888392406"))
	if err != nil {
		lang.SayString("subx3237")
	} else {
		if !(rexsult.ToString() == "-70454063895691.336929452116845012") {
			lang.SayString("subx3237")
		}
	}
	rexsult, err = lang.RxFromString("29119.220621511046558757900645228").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("3517612.8810761470018676311863010E-222"))
	if err != nil {
		lang.SayString("subx3238")
	} else {
		if !(rexsult.ToString() == "29119.220621511046558757900645228") {
			lang.SayString("subx3238")
		}
	}
	rexsult, err = lang.RxFromString("-5168.2214111091132913776042214525").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-5690274.0971173476527123568627720"))
	if err != nil {
		lang.SayString("subx3239")
	} else {
		if !(rexsult.ToString() == "5685105.8757062385394209792585506") {
			lang.SayString("subx3239")
		}
	}
	rexsult, err = lang.RxFromString("33783.060857197067391462144517964").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-2070.0806959465088198508322815406"))
	if err != nil {
		lang.SayString("subx3240")
	} else {
		if !(rexsult.ToString() == "35853.141553143576211312976799505") {
			lang.SayString("subx3240")
		}
	}
	rexsult, err = lang.RxFromString("42207435091050.840296353874733169E-905").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("73330633078.828216018536326743325E+976"))
	if err != nil {
		lang.SayString("subx3241")
	} else {
		if !(rexsult.ToString() == "-7.3330633078828216018536326743325E+986") {
			lang.SayString("subx3241")
		}
	}
	rexsult, err = lang.RxFromString("-71800.806700868784841045406679641").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-39617456964250697902519150526701"))
	if err != nil {
		lang.SayString("subx3242")
	} else {
		if !(rexsult.ToString() == "39617456964250697902519150454900") {
			lang.SayString("subx3242")
		}
	}
	rexsult, err = lang.RxFromString("53627480801.631504892310016062160").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("328259.56947661049313311983109503"))
	if err != nil {
		lang.SayString("subx3243")
	} else {
		if !(rexsult.ToString() == "53627152542.062028281816882942329") {
			lang.SayString("subx3243")
		}
	}
	rexsult, err = lang.RxFromString("-5052601598.5559371338428368438728").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-97855372.224321664785314782556064"))
	if err != nil {
		lang.SayString("subx3244")
	} else {
		if !(rexsult.ToString() == "-4954746226.3316154690575220613167") {
			lang.SayString("subx3244")
		}
	}
	rexsult, err = lang.RxFromString("4208134320733.7069742988228068191E+146").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("4270869.1760149477598920960628392E+471"))
	if err != nil {
		lang.SayString("subx3245")
	} else {
		if !(rexsult.ToString() == "-4.2708691760149477598920960628392E+477") {
			lang.SayString("subx3245")
		}
	}
	rexsult, err = lang.RxFromString("-8.5077009657942581515590471189084E+308").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("9652145155.374217047842114042376E-250"))
	if err != nil {
		lang.SayString("subx3246")
	} else {
		if !(rexsult.ToString() == "-8.5077009657942581515590471189084E+308") {
			lang.SayString("subx3246")
		}
	}
	rexsult, err = lang.RxFromString("-9504.9703032286960790904181078063E+619").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-86.245956949049186533469206485003"))
	if err != nil {
		lang.SayString("subx3247")
	} else {
		if !(rexsult.ToString() == "-9.5049703032286960790904181078063E+622") {
			lang.SayString("subx3247")
		}
	}
	rexsult, err = lang.RxFromString("-440220916.66716743026896931194749").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-102725.01594377871560564824358775"))
	if err != nil {
		lang.SayString("subx3248")
	} else {
		if !(rexsult.ToString() == "-440118191.65122365155336366370390") {
			lang.SayString("subx3248")
		}
	}
	rexsult, err = lang.RxFromString("-46.250735086006350517943464758019").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("14656357555174.263295266074908024"))
	if err != nil {
		lang.SayString("subx3249")
	} else {
		if !(rexsult.ToString() == "-14656357555220.514030352081258542") {
			lang.SayString("subx3249")
		}
	}
	rexsult, err = lang.RxFromString("-61641121299391.316420647102699627E+763").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-91896469863.461006903590004188187E+474"))
	if err != nil {
		lang.SayString("subx3250")
	} else {
		if !(rexsult.ToString() == "-6.1641121299391316420647102699627E+776") {
			lang.SayString("subx3250")
		}
	}
	rexsult, err = lang.RxFromString("96668419802749.555738010239087449E-838").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-19498732131365824921639467044927E-542"))
	if err != nil {
		lang.SayString("subx3251")
	} else {
		if !(rexsult.ToString() == "1.9498732131365824921639467044927E-511") {
			lang.SayString("subx3251")
		}
	}
	rexsult, err = lang.RxFromString("-8534543911197995906031245719519E+124").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("16487117050031.594886541650897974"))
	if err != nil {
		lang.SayString("subx3252")
	} else {
		if !(rexsult.ToString() == "-8.5345439111979959060312457195190E+154") {
			lang.SayString("subx3252")
		}
	}
	rexsult, err = lang.RxFromString("-62663404777.352508979582846931050").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("9.2570938837239134052589184917186E+916"))
	if err != nil {
		lang.SayString("subx3253")
	} else {
		if !(rexsult.ToString() == "-9.2570938837239134052589184917186E+916") {
			lang.SayString("subx3253")
		}
	}
	rexsult, err = lang.RxFromString("1.744601214474560992754529320172E-827").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-17.353669504253419489494030651507E-631"))
	if err != nil {
		lang.SayString("subx3254")
	} else {
		if !(rexsult.ToString() == "1.7353669504253419489494030651507E-630") {
			lang.SayString("subx3254")
		}
	}
	rexsult, err = lang.RxFromString("0367191549036702224827734853471").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("4410320662415266533763143837742E+721"))
	if err != nil {
		lang.SayString("subx3255")
	} else {
		if !(rexsult.ToString() == "-4.4103206624152665337631438377420E+751") {
			lang.SayString("subx3255")
		}
	}
	rexsult, err = lang.RxFromString("097704116.4492566721965710365073").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-96736.400939809433556067504289145"))
	if err != nil {
		lang.SayString("subx3256")
	} else {
		if !(rexsult.ToString() == "97800852.850196481630127104011589") {
			lang.SayString("subx3256")
		}
	}
	rexsult, err = lang.RxFromString("19533298.147150158931958733807878").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("80.141668338350708476637377647025E-641"))
	if err != nil {
		lang.SayString("subx3257")
	} else {
		if !(rexsult.ToString() == "19533298.147150158931958733807878") {
			lang.SayString("subx3257")
		}
	}
	rexsult, err = lang.RxFromString("-23765003221220177430797028997378").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-15203369569.373411506379096871224E-945"))
	if err != nil {
		lang.SayString("subx3258")
	} else {
		if !(rexsult.ToString() == "-23765003221220177430797028997378") {
			lang.SayString("subx3258")
		}
	}
	rexsult, err = lang.RxFromString("129255.41937932433359193338910552E+932").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-281253953.38990382799508873560320"))
	if err != nil {
		lang.SayString("subx3259")
	} else {
		if !(rexsult.ToString() == "1.2925541937932433359193338910552E+937") {
			lang.SayString("subx3259")
		}
	}
	rexsult, err = lang.RxFromString("-86863.276249466008289214762980838").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("531.50602652732088208397655484476"))
	if err != nil {
		lang.SayString("subx3260")
	} else {
		if !(rexsult.ToString() == "-87394.782275993329171298739535683") {
			lang.SayString("subx3260")
		}
	}
	rexsult, err = lang.RxFromString("-40707.169006771111855573524157083").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-68795521421321853333274411827749"))
	if err != nil {
		lang.SayString("subx3261")
	} else {
		if !(rexsult.ToString() == "68795521421321853333274411787042") {
			lang.SayString("subx3261")
		}
	}
	rexsult, err = lang.RxFromString("-90838752568673.728630494658778003E+095").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-738.01370301217606577533107981431"))
	if err != nil {
		lang.SayString("subx3262")
	} else {
		if !(rexsult.ToString() == "-9.0838752568673728630494658778003E+108") {
			lang.SayString("subx3262")
		}
	}
	rexsult, err = lang.RxFromString("-4245360967593.9206771555839718158E-682").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-3.119606239042530207103203508057"))
	if err != nil {
		lang.SayString("subx3263")
	} else {
		if !(rexsult.ToString() == "3.1196062390425302071032035080570") {
			lang.SayString("subx3263")
		}
	}
	rexsult, err = lang.RxFromString("-3422145405774.0800213000547612240E-023").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-60810.964656409650839011321706310"))
	if err != nil {
		lang.SayString("subx3264")
	} else {
		if !(rexsult.ToString() == "60810.964656409616617557263965510") {
			lang.SayString("subx3264")
		}
	}
	rexsult, err = lang.RxFromString("-24521811.07649485796598387627478E-661").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-94860846133404815410816234000694"))
	if err != nil {
		lang.SayString("subx3265")
	} else {
		if !(rexsult.ToString() == "94860846133404815410816234000694") {
			lang.SayString("subx3265")
		}
	}
	rexsult, err = lang.RxFromString("-5042529937498.8944492248538951438").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("3891904674.4549170968807145612549"))
	if err != nil {
		lang.SayString("subx3266")
	} else {
		if !(rexsult.ToString() == "-5046421842173.3493663217346097051") {
			lang.SayString("subx3266")
		}
	}
	rexsult, err = lang.RxFromString("-535824163.54531747646293693868651E-665").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("2732988.5891363639325008206099712"))
	if err != nil {
		lang.SayString("subx3267")
	} else {
		if !(rexsult.ToString() == "-2732988.5891363639325008206099712") {
			lang.SayString("subx3267")
		}
	}
	rexsult, err = lang.RxFromString("24032.702008553084252925140858134E-509").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("52864854.899420632375589206704068"))
	if err != nil {
		lang.SayString("subx3268")
	} else {
		if !(rexsult.ToString() == "-52864854.899420632375589206704068") {
			lang.SayString("subx3268")
		}
	}
	rexsult, err = lang.RxFromString("71553220259.938950698030519906727E-496").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("754.44220417415325444943566016062"))
	if err != nil {
		lang.SayString("subx3269")
	} else {
		if !(rexsult.ToString() == "-754.44220417415325444943566016062") {
			lang.SayString("subx3269")
		}
	}
	rexsult, err = lang.RxFromString("35572.960284795962697740953932508").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("520.39506364687594082725754878910E-731"))
	if err != nil {
		lang.SayString("subx3270")
	} else {
		if !(rexsult.ToString() == "35572.960284795962697740953932508") {
			lang.SayString("subx3270")
		}
	}
	rexsult, err = lang.RxFromString("53035405018123760598334895413057E+818").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-9558464247240.4476790042911379151"))
	if err != nil {
		lang.SayString("subx3271")
	} else {
		if !(rexsult.ToString() == "5.3035405018123760598334895413057E+849") {
			lang.SayString("subx3271")
		}
	}
	rexsult, err = lang.RxFromString("95.490751127249945886828257312118").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("987.01498316307365714167410690192E+133"))
	if err != nil {
		lang.SayString("subx3272")
	} else {
		if !(rexsult.ToString() == "-9.8701498316307365714167410690192E+135") {
			lang.SayString("subx3272")
		}
	}
	rexsult, err = lang.RxFromString("69434850287.460788550936730296153").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-35119136549015044241569827542264"))
	if err != nil {
		lang.SayString("subx3273")
	} else {
		if !(rexsult.ToString() == "35119136549015044241639262392551") {
			lang.SayString("subx3273")
		}
	}
	rexsult, err = lang.RxFromString("-392.22739924621965621739098725407").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-65551274.987160998195282109612136"))
	if err != nil {
		lang.SayString("subx3274")
	} else {
		if !(rexsult.ToString() == "65550882.759761751975625892221149") {
			lang.SayString("subx3274")
		}
	}
	rexsult, err = lang.RxFromString("6413265.4423561191792972085539457").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("24514.222704714139350026165721146"))
	if err != nil {
		lang.SayString("subx3275")
	} else {
		if !(rexsult.ToString() == "6388751.2196514050399471823882246") {
			lang.SayString("subx3275")
		}
	}
	rexsult, err = lang.RxFromString("-6.9667706389122107760046184064057E+487").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("32.405810703971538278419625527234"))
	if err != nil {
		lang.SayString("subx3276")
	} else {
		if !(rexsult.ToString() == "-6.9667706389122107760046184064057E+487") {
			lang.SayString("subx3276")
		}
	}
	rexsult, err = lang.RxFromString("378204716633.40024100602896057615").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-0300218714378.322231269606216516"))
	if err != nil {
		lang.SayString("subx3277")
	} else {
		if !(rexsult.ToString() == "678423431011.72247227563517709215") {
			lang.SayString("subx3277")
		}
	}
	rexsult, err = lang.RxFromString("-44234.512012457148027685282969235E+501").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("2132572.4571987908375002100894927"))
	if err != nil {
		lang.SayString("subx3278")
	} else {
		if !(rexsult.ToString() == "-4.4234512012457148027685282969235E+505") {
			lang.SayString("subx3278")
		}
	}
	rexsult, err = lang.RxFromString("-3554.5895974968741465654022772100E-073").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("9752.0428746722497621936998533848E+516"))
	if err != nil {
		lang.SayString("subx3279")
	} else {
		if !(rexsult.ToString() == "-9.7520428746722497621936998533848E+519") {
			lang.SayString("subx3279")
		}
	}
	rexsult, err = lang.RxFromString("750187685.63632608923397234391668").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("4633194252863.6730625972669246025"))
	if err != nil {
		lang.SayString("subx3280")
	} else {
		if !(rexsult.ToString() == "-4632444065178.0367365080329522586") {
			lang.SayString("subx3280")
		}
	}
	rexsult, err = lang.RxFromString("30190034242853.251165951457589386E-028").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("8038885676.3204238322976087799751E+018"))
	if err != nil {
		lang.SayString("subx3281")
	} else {
		if !(rexsult.ToString() == "-8038885676320423832297608779.9751") {
			lang.SayString("subx3281")
		}
	}
	rexsult, err = lang.RxFromString("65.537942676774715953400768803539").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("125946917260.87536506197191782198"))
	if err != nil {
		lang.SayString("subx3282")
	} else {
		if !(rexsult.ToString() == "-125946917195.33742238519720186858") {
			lang.SayString("subx3282")
		}
	}
	rexsult, err = lang.RxFromString("8015272348677.5489394183881813700").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("949.23027111499779258284877920022"))
	if err != nil {
		lang.SayString("subx3283")
	} else {
		if !(rexsult.ToString() == "8015272347728.3186683033903887872") {
			lang.SayString("subx3283")
		}
	}
	rexsult, err = lang.RxFromString("-32595333982.67068622120451804225").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("69130.052233649808750113141502465E-862"))
	if err != nil {
		lang.SayString("subx3284")
	} else {
		if !(rexsult.ToString() == "-32595333982.670686221204518042250") {
			lang.SayString("subx3284")
		}
	}
	rexsult, err = lang.RxFromString("-17544189017145.710117633021800426E-537").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("292178000.06450804618299520094843"))
	if err != nil {
		lang.SayString("subx3285")
	} else {
		if !(rexsult.ToString() == "-292178000.06450804618299520094843") {
			lang.SayString("subx3285")
		}
	}
	rexsult, err = lang.RxFromString("-506650.99395649907139204052441630").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("11.018427502031650148962067367158"))
	if err != nil {
		lang.SayString("subx3286")
	} else {
		if !(rexsult.ToString() == "-506662.01238400110304218948648367") {
			lang.SayString("subx3286")
		}
	}
	rexsult, err = lang.RxFromString("4846835159.5922372307656000769395E-241").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-84.001893040865864590122330800768"))
	if err != nil {
		lang.SayString("subx3287")
	} else {
		if !(rexsult.ToString() == "84.001893040865864590122330800768") {
			lang.SayString("subx3287")
		}
	}
	rexsult, err = lang.RxFromString("-35.029311013822259358116556164908").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-3994308878.1994645313414534209127"))
	if err != nil {
		lang.SayString("subx3288")
	} else {
		if !(rexsult.ToString() == "3994308843.1701535175191940627962") {
			lang.SayString("subx3288")
		}
	}
	rexsult, err = lang.RxFromString("7606663750.6735265233044420887018E+166").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-5491814639.4484565418284686127552E+365"))
	if err != nil {
		lang.SayString("subx3289")
	} else {
		if !(rexsult.ToString() == "5.4918146394484565418284686127552E+374") {
			lang.SayString("subx3289")
		}
	}
	rexsult, err = lang.RxFromString("-25677.829660831741274207326697052E-163").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-94135395124193048560172705082029E-862"))
	if err != nil {
		lang.SayString("subx3290")
	} else {
		if !(rexsult.ToString() == "-2.5677829660831741274207326697052E-159") {
			lang.SayString("subx3290")
		}
	}
	rexsult, err = lang.RxFromString("97271576094.456406729671729224827").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-1.5412563837540810793697955063295E+554"))
	if err != nil {
		lang.SayString("subx3291")
	} else {
		if !(rexsult.ToString() == "1.5412563837540810793697955063295E+554") {
			lang.SayString("subx3291")
		}
	}
	rexsult, err = lang.RxFromString("41139789894.401826915757391777563").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-1.8729920305671057957156159690823E-020"))
	if err != nil {
		lang.SayString("subx3292")
	} else {
		if !(rexsult.ToString() == "41139789894.401826915757391777582") {
			lang.SayString("subx3292")
		}
	}
	rexsult, err = lang.RxFromString("-83310831287241.777598696853498149").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-54799825033.797100418162985103519E-330"))
	if err != nil {
		lang.SayString("subx3293")
	} else {
		if !(rexsult.ToString() == "-83310831287241.777598696853498149") {
			lang.SayString("subx3293")
		}
	}
	rexsult, err = lang.RxFromString("4506653461.4414974233678331771169").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-74955470.977653038010031457181957E-887"))
	if err != nil {
		lang.SayString("subx3294")
	} else {
		if !(rexsult.ToString() == "4506653461.4414974233678331771169") {
			lang.SayString("subx3294")
		}
	}
	rexsult, err = lang.RxFromString("23777.857951114967684767609401626").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("720760.03897144157012301385227528"))
	if err != nil {
		lang.SayString("subx3295")
	} else {
		if !(rexsult.ToString() == "-696982.18102032660243824624287365") {
			lang.SayString("subx3295")
		}
	}
	rexsult, err = lang.RxFromString("-21337853323980858055292469611948").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("6080272840.3071490445256786982100E+532"))
	if err != nil {
		lang.SayString("subx3296")
	} else {
		if !(rexsult.ToString() == "-6.0802728403071490445256786982100E+541") {
			lang.SayString("subx3296")
		}
	}
	rexsult, err = lang.RxFromString("-818409238.0423893439849743856947").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("756.39156265972753259267069158760"))
	if err != nil {
		lang.SayString("subx3297")
	} else {
		if !(rexsult.ToString() == "-818409994.43395200371250697836539") {
			lang.SayString("subx3297")
		}
	}
	rexsult, err = lang.RxFromString("47567380384943.87013600286155046").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("65.084709374908275826942076480326"))
	if err != nil {
		lang.SayString("subx3298")
	} else {
		if !(rexsult.ToString() == "47567380384878.785426627953274633") {
			lang.SayString("subx3298")
		}
	}
	rexsult, err = lang.RxFromString("-296615544.05897984545127115290251").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-5416115.4315053536014016450973264"))
	if err != nil {
		lang.SayString("subx3299")
	} else {
		if !(rexsult.ToString() == "-291199428.62747449184986950780518") {
			lang.SayString("subx3299")
		}
	}
	rexsult, err = lang.RxFromString("61391705914.046707180652185247584E+739").OpSub(lang.RxSetWithDigit(32), lang.RxFromString("-675982087721.91856456125242297346"))
	if err != nil {
		lang.SayString("subx3300")
	} else {
		if !(rexsult.ToString() == "6.1391705914046707180652185247584E+749") {
			lang.SayString("subx3300")
		}
	}
	rexsult, err = lang.RxFromString("042.668056830485571428877212944418").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-01364112374639.4941124016455321071"))
	if err != nil {
		lang.SayString("subx3401")
	} else {
		if !(rexsult.ToString() == "1364112374682.16216923213110353598") {
			lang.SayString("subx3401")
		}
	}
	rexsult, err = lang.RxFromString("-327.179426341653256363531809227344E+453").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("759067457.107518663444899356760793"))
	if err != nil {
		lang.SayString("subx3402")
	} else {
		if !(rexsult.ToString() == "-3.27179426341653256363531809227344E+455") {
			lang.SayString("subx3402")
		}
	}
	rexsult, err = lang.RxFromString("81721.5803096185422787702538493471").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("900099473.245809628076790757217328"))
	if err != nil {
		lang.SayString("subx3403")
	} else {
		if !(rexsult.ToString() == "-900017751.665500009534511986963479") {
			lang.SayString("subx3403")
		}
	}
	rexsult, err = lang.RxFromString("3991.56734635183403814636354392163E-807").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("72.3239822255871305731314565069132"))
	if err != nil {
		lang.SayString("subx3404")
	} else {
		if !(rexsult.ToString() == "-72.3239822255871305731314565069132") {
			lang.SayString("subx3404")
		}
	}
	rexsult, err = lang.RxFromString("-66.3393308595957726456416979163306").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("5.08486573050459213864684589662559"))
	if err != nil {
		lang.SayString("subx3405")
	} else {
		if !(rexsult.ToString() == "-71.4241965901003647842885438129562") {
			lang.SayString("subx3405")
		}
	}
	rexsult, err = lang.RxFromString("-393606.873703567753255097095208112E+111").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-2124339550.86891093200758095660557"))
	if err != nil {
		lang.SayString("subx3406")
	} else {
		if !(rexsult.ToString() == "-3.93606873703567753255097095208112E+116") {
			lang.SayString("subx3406")
		}
	}
	rexsult, err = lang.RxFromString("-019133598.609812524622150421584346").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-858439846.628367734642622922030051"))
	if err != nil {
		lang.SayString("subx3407")
	} else {
		if !(rexsult.ToString() == "839306248.018555210020472500445705") {
			lang.SayString("subx3407")
		}
	}
	rexsult, err = lang.RxFromString("465.351982159046525762715549761814").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("240444.975944666924657629172844780"))
	if err != nil {
		lang.SayString("subx3408")
	} else {
		if !(rexsult.ToString() == "-239979.623962507878131866457295018") {
			lang.SayString("subx3408")
		}
	}
	rexsult, err = lang.RxFromString("28066955004783.1076824222873384828").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("571699969.220753535758504907561016E-718"))
	if err != nil {
		lang.SayString("subx3409")
	} else {
		if !(rexsult.ToString() == "28066955004783.1076824222873384828") {
			lang.SayString("subx3409")
		}
	}
	rexsult, err = lang.RxFromString("28275236927392.4960902824105246047").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("28212038.4825243127096613158419270E+422"))
	if err != nil {
		lang.SayString("subx3410")
	} else {
		if !(rexsult.ToString() == "-2.82120384825243127096613158419270E+429") {
			lang.SayString("subx3410")
		}
	}
	rexsult, err = lang.RxFromString("11791.8644211874630234271801789996").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-8.45457275930363860982261343159741"))
	if err != nil {
		lang.SayString("subx3411")
	} else {
		if !(rexsult.ToString() == "11800.3189939467666620370027924312") {
			lang.SayString("subx3411")
		}
	}
	rexsult, err = lang.RxFromString("44.7085340739581668975502342787578").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-9337.05408133023920640485556647937"))
	if err != nil {
		lang.SayString("subx3412")
	} else {
		if !(rexsult.ToString() == "9381.76261540419737330240580075813") {
			lang.SayString("subx3412")
		}
	}
	rexsult, err = lang.RxFromString("93354527428804.5458053295581965867E+576").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-856525909852.318790321300941615314"))
	if err != nil {
		lang.SayString("subx3413")
	} else {
		if !(rexsult.ToString() == "9.33545274288045458053295581965867E+589") {
			lang.SayString("subx3413")
		}
	}
	rexsult, err = lang.RxFromString("-367399415798804503177950040443482").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-54845683.9691776397285506712812754"))
	if err != nil {
		lang.SayString("subx3414")
	} else {
		if !(rexsult.ToString() == "-367399415798804503177949985597798") {
			lang.SayString("subx3414")
		}
	}
	rexsult, err = lang.RxFromString("-2.87155919781024108503670175443740").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("89529730130.6427881332776797193807"))
	if err != nil {
		lang.SayString("subx3415")
	} else {
		if !(rexsult.ToString() == "-89529730133.5143473310879208044174") {
			lang.SayString("subx3415")
		}
	}
	rexsult, err = lang.RxFromString("-010.693934338179479652178057279204E+188").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("26484.8887731973153745666514260684"))
	if err != nil {
		lang.SayString("subx3416")
	} else {
		if !(rexsult.ToString() == "-1.06939343381794796521780572792040E+189") {
			lang.SayString("subx3416")
		}
	}
	rexsult, err = lang.RxFromString("611655569568.832698912762075889186").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("010182743219.475839030505966016982"))
	if err != nil {
		lang.SayString("subx3417")
	} else {
		if !(rexsult.ToString() == "601472826349.356859882256109872204") {
			lang.SayString("subx3417")
		}
	}
	rexsult, err = lang.RxFromString("3457947.39062863674882672518304442").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-01.9995218868908849056866549811425"))
	if err != nil {
		lang.SayString("subx3418")
	} else {
		if !(rexsult.ToString() == "3457949.39015052363971163086969940") {
			lang.SayString("subx3418")
		}
	}
	rexsult, err = lang.RxFromString("-53308666960535.7393391289364591513").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-6527.00547629475578694521436764596E-442"))
	if err != nil {
		lang.SayString("subx3419")
	} else {
		if !(rexsult.ToString() == "-53308666960535.7393391289364591513") {
			lang.SayString("subx3419")
		}
	}
	rexsult, err = lang.RxFromString("-5568057.17870139549478277980540034").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-407906443.141342175740471849723638"))
	if err != nil {
		lang.SayString("subx3420")
	} else {
		if !(rexsult.ToString() == "402338385.962640780245689069918238") {
			lang.SayString("subx3420")
		}
	}
	rexsult, err = lang.RxFromString("9804385273.49533524416415189990857").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("84.1433929743544659553964804646569"))
	if err != nil {
		lang.SayString("subx3421")
	} else {
		if !(rexsult.ToString() == "9804385189.35194226980968594451209") {
			lang.SayString("subx3421")
		}
	}
	rexsult, err = lang.RxFromString("-5234910986592.18801727046580014273E-547").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-5874220715892.91440069210512515154"))
	if err != nil {
		lang.SayString("subx3422")
	} else {
		if !(rexsult.ToString() == "5874220715892.91440069210512515154") {
			lang.SayString("subx3422")
		}
	}
	rexsult, err = lang.RxFromString("698416560151955285929747633786867E-495").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("51754681.6784872628933218985216916E-266"))
	if err != nil {
		lang.SayString("subx3423")
	} else {
		if !(rexsult.ToString() == "-5.17546816784872628933218985216916E-259") {
			lang.SayString("subx3423")
		}
	}
	rexsult, err = lang.RxFromString("107635.497735316515080720330536027").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-3972075.83989512668362609609006425E-605"))
	if err != nil {
		lang.SayString("subx3424")
	} else {
		if !(rexsult.ToString() == "107635.497735316515080720330536027") {
			lang.SayString("subx3424")
		}
	}
	rexsult, err = lang.RxFromString("-32174291345686.5371446616670961807").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("79518863759385.5925052747867099091E+408"))
	if err != nil {
		lang.SayString("subx3425")
	} else {
		if !(rexsult.ToString() == "-7.95188637593855925052747867099091E+421") {
			lang.SayString("subx3425")
		}
	}
	rexsult, err = lang.RxFromString("-8151730494.53190523620899410544099E+688").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-93173.0631474527142307644239919480E+900"))
	if err != nil {
		lang.SayString("subx3426")
	} else {
		if !(rexsult.ToString() == "9.31730631474527142307644239919480E+904") {
			lang.SayString("subx3426")
		}
	}
	rexsult, err = lang.RxFromString("1.33649801345976199708341799505220").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-56623.0530039528969825480755159562E+459"))
	if err != nil {
		lang.SayString("subx3427")
	} else {
		if !(rexsult.ToString() == "5.66230530039528969825480755159562E+463") {
			lang.SayString("subx3427")
		}
	}
	rexsult, err = lang.RxFromString("67762238162788.6551061476018185196").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-6140.75837959248100352788853809376E-822"))
	if err != nil {
		lang.SayString("subx3428")
	} else {
		if !(rexsult.ToString() == "67762238162788.6551061476018185196") {
			lang.SayString("subx3428")
		}
	}
	rexsult, err = lang.RxFromString("4286562.76568866751577306056498271").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("6286.77291578497580015557979349893E+820"))
	if err != nil {
		lang.SayString("subx3429")
	} else {
		if !(rexsult.ToString() == "-6.28677291578497580015557979349893E+823") {
			lang.SayString("subx3429")
		}
	}
	rexsult, err = lang.RxFromString("-765782.827432642697305644096365566").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("67.1634368459576834692758114618652"))
	if err != nil {
		lang.SayString("subx3430")
	} else {
		if !(rexsult.ToString() == "-765849.990869488654989113372177028") {
			lang.SayString("subx3430")
		}
	}
	rexsult, err = lang.RxFromString("46.2835931916106252756465724211276").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("59.2989237834093118332826617957791"))
	if err != nil {
		lang.SayString("subx3431")
	} else {
		if !(rexsult.ToString() == "-13.0153305917986865576360893746515") {
			lang.SayString("subx3431")
		}
	}
	rexsult, err = lang.RxFromString("-3029555.82298840234029474459694644").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("857535844655004737373089601128532"))
	if err != nil {
		lang.SayString("subx3432")
	} else {
		if !(rexsult.ToString() == "-857535844655004737373089604158088") {
			lang.SayString("subx3432")
		}
	}
	rexsult, err = lang.RxFromString("-0138466789523.10694176543700501945E-948").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("481026979918882487383654367924619"))
	if err != nil {
		lang.SayString("subx3433")
	} else {
		if !(rexsult.ToString() == "-481026979918882487383654367924619") {
			lang.SayString("subx3433")
		}
	}
	rexsult, err = lang.RxFromString("-9593566466.96690575714244442109870").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-87632034347.4845477961976776833770E+769"))
	if err != nil {
		lang.SayString("subx3434")
	} else {
		if !(rexsult.ToString() == "8.76320343474845477961976776833770E+779") {
			lang.SayString("subx3434")
		}
	}
	rexsult, err = lang.RxFromString("-3189.30765477670526823106100241863E-898").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("565688889.355241946154894311253202E-466"))
	if err != nil {
		lang.SayString("subx3435")
	} else {
		if !(rexsult.ToString() == "-5.65688889355241946154894311253202E-458") {
			lang.SayString("subx3435")
		}
	}
	rexsult, err = lang.RxFromString("-17084552395.6714834680088150543965").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-631925802672.685034379197328370812E+527"))
	if err != nil {
		lang.SayString("subx3436")
	} else {
		if !(rexsult.ToString() == "6.31925802672685034379197328370812E+538") {
			lang.SayString("subx3436")
		}
	}
	rexsult, err = lang.RxFromString("034956830.349823306815911887469760").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-61600816.0672274126966042956781665E-667"))
	if err != nil {
		lang.SayString("subx3437")
	} else {
		if !(rexsult.ToString() == "34956830.3498233068159118874697600") {
			lang.SayString("subx3437")
		}
	}
	rexsult, err = lang.RxFromString("-763.440067781256632695791981893608").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("19.9263811350611007833220620745413"))
	if err != nil {
		lang.SayString("subx3438")
	} else {
		if !(rexsult.ToString() == "-783.366448916317733479114043968149") {
			lang.SayString("subx3438")
		}
	}
	rexsult, err = lang.RxFromString("-510472027868440667684575147556654E+789").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("834872378550801889983927148587909"))
	if err != nil {
		lang.SayString("subx3439")
	} else {
		if !(rexsult.ToString() == "-5.10472027868440667684575147556654E+821") {
			lang.SayString("subx3439")
		}
	}
	rexsult, err = lang.RxFromString("070304761.560517086676993503034828E-094").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-17773.7446959771077104057845273992E-761"))
	if err != nil {
		lang.SayString("subx3440")
	} else {
		if !(rexsult.ToString() == "7.03047615605170866769935030348280E-87") {
			lang.SayString("subx3440")
		}
	}
	rexsult, err = lang.RxFromString("-0970725697662.27605454336231195463").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-4541.41897546697187157913886433474"))
	if err != nil {
		lang.SayString("subx3441")
	} else {
		if !(rexsult.ToString() == "-970725693120.857079076390440375491") {
			lang.SayString("subx3441")
		}
	}
	rexsult, err = lang.RxFromString("-808178238631844268316111259558675").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-598400.265108644514211244980426520"))
	if err != nil {
		lang.SayString("subx3442")
	} else {
		if !(rexsult.ToString() == "-808178238631844268316111258960275") {
			lang.SayString("subx3442")
		}
	}
	rexsult, err = lang.RxFromString("-9.90826595069053564311371766315200").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-031.625916781307847864872329806646"))
	if err != nil {
		lang.SayString("subx3443")
	} else {
		if !(rexsult.ToString() == "21.7176508306173122217586121434940") {
			lang.SayString("subx3443")
		}
	}
	rexsult, err = lang.RxFromString("-196925.469891897719160698483752907").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-41268.9975444533794067723958739778"))
	if err != nil {
		lang.SayString("subx3444")
	} else {
		if !(rexsult.ToString() == "-155656.472347444339753926087878929") {
			lang.SayString("subx3444")
		}
	}
	rexsult, err = lang.RxFromString("421071135212152225162086005824310").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("1335320330.08964354845796510145246E-604"))
	if err != nil {
		lang.SayString("subx3445")
	} else {
		if !(rexsult.ToString() == "421071135212152225162086005824310") {
			lang.SayString("subx3445")
		}
	}
	rexsult, err = lang.RxFromString("1249441.46421514282301182772247227").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-0289848.71208912281976374705180836E-676"))
	if err != nil {
		lang.SayString("subx3446")
	} else {
		if !(rexsult.ToString() == "1249441.46421514282301182772247227") {
			lang.SayString("subx3446")
		}
	}
	rexsult, err = lang.RxFromString("74815000.4716875558358937279052903").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-690425401708167622194241915195001E+891"))
	if err != nil {
		lang.SayString("subx3447")
	} else {
		if !(rexsult.ToString() == "6.90425401708167622194241915195001E+923") {
			lang.SayString("subx3447")
		}
	}
	rexsult, err = lang.RxFromString("-1683993.51210241555668790556759021").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-72394384927344.8402585228267493374"))
	if err != nil {
		lang.SayString("subx3448")
	} else {
		if !(rexsult.ToString() == "72394383243351.3281561072700614318") {
			lang.SayString("subx3448")
		}
	}
	rexsult, err = lang.RxFromString("-763.148530974741766171756970448158").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("517370.808956957601473642272664647"))
	if err != nil {
		lang.SayString("subx3449")
	} else {
		if !(rexsult.ToString() == "-518133.957487932343239814029635095") {
			lang.SayString("subx3449")
		}
	}
	rexsult, err = lang.RxFromString("-77.5841338812312523460591226178754").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-927540422.641025050968830154578151E+524"))
	if err != nil {
		lang.SayString("subx3450")
	} else {
		if !(rexsult.ToString() == "9.27540422641025050968830154578151E+532") {
			lang.SayString("subx3450")
		}
	}
	rexsult, err = lang.RxFromString("5176295309.89943746236102209837813").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-129733.103628797477167908698565465"))
	if err != nil {
		lang.SayString("subx3451")
	} else {
		if !(rexsult.ToString() == "5176425043.00306625983819000707670") {
			lang.SayString("subx3451")
		}
	}
	rexsult, err = lang.RxFromString("4471634841166.90197229286530307326E+172").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("31511104397.8671727003201890825879E-955"))
	if err != nil {
		lang.SayString("subx3452")
	} else {
		if !(rexsult.ToString() == "4.47163484116690197229286530307326E+184") {
			lang.SayString("subx3452")
		}
	}
	rexsult, err = lang.RxFromString("-8189130.15945231049670285726774317").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("2.57912402871404325831670923864936E-366"))
	if err != nil {
		lang.SayString("subx3453")
	} else {
		if !(rexsult.ToString() == "-8189130.15945231049670285726774317") {
			lang.SayString("subx3453")
		}
	}
	rexsult, err = lang.RxFromString("-254346232981353293392888785643245").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-764.416902486152093758287929678445"))
	if err != nil {
		lang.SayString("subx3454")
	} else {
		if !(rexsult.ToString() == "-254346232981353293392888785642481") {
			lang.SayString("subx3454")
		}
	}
	rexsult, err = lang.RxFromString("-2875.36745499544177964804277329726").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-13187.8492045054802205691248267638"))
	if err != nil {
		lang.SayString("subx3455")
	} else {
		if !(rexsult.ToString() == "10312.4817495100384409210820534665") {
			lang.SayString("subx3455")
		}
	}
	rexsult, err = lang.RxFromString("-7.26927570984219944693602140223103").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("0160883021541.32275286769110003971E-438"))
	if err != nil {
		lang.SayString("subx3456")
	} else {
		if !(rexsult.ToString() == "-7.26927570984219944693602140223103") {
			lang.SayString("subx3456")
		}
	}
	rexsult, err = lang.RxFromString("-28567151.6868762752241056540028515E+498").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-4470.15455137546427645290210989675"))
	if err != nil {
		lang.SayString("subx3457")
	} else {
		if !(rexsult.ToString() == "-2.85671516868762752241056540028515E+505") {
			lang.SayString("subx3457")
		}
	}
	rexsult, err = lang.RxFromString("7191753.79974136447257468282073718").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("81.3878426176038487948375777384429"))
	if err != nil {
		lang.SayString("subx3458")
	} else {
		if !(rexsult.ToString() == "7191672.41189874686872588798315944") {
			lang.SayString("subx3458")
		}
	}
	rexsult, err = lang.RxFromString("502975804.069864536104621700404770").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("684.790028432074527960269515227243"))
	if err != nil {
		lang.SayString("subx3459")
	} else {
		if !(rexsult.ToString() == "502975119.279836104030093740135255") {
			lang.SayString("subx3459")
		}
	}
	rexsult, err = lang.RxFromString("1505368.42063731861590460453659570").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-465242.678439951462767630022819105"))
	if err != nil {
		lang.SayString("subx3460")
	} else {
		if !(rexsult.ToString() == "1970611.09907727007867223455941481") {
			lang.SayString("subx3460")
		}
	}
	rexsult, err = lang.RxFromString("69225023.2850142784063416137144829").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("8584050.06648191798834819995325693"))
	if err != nil {
		lang.SayString("subx3461")
	} else {
		if !(rexsult.ToString() == "60640973.2185323604179934137612260") {
			lang.SayString("subx3461")
		}
	}
	rexsult, err = lang.RxFromString("-55669501853.7751006841940471339310E+614").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("061400538.186044693233816566977189"))
	if err != nil {
		lang.SayString("subx3462")
	} else {
		if !(rexsult.ToString() == "-5.56695018537751006841940471339310E+624") {
			lang.SayString("subx3462")
		}
	}
	rexsult, err = lang.RxFromString("-527566.521273992424649346837337602E-408").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-834662.599983953345718523807123972"))
	if err != nil {
		lang.SayString("subx3463")
	} else {
		if !(rexsult.ToString() == "834662.599983953345718523807123972") {
			lang.SayString("subx3463")
		}
	}
	rexsult, err = lang.RxFromString("69065510.0459653699418083455335366").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("694848643848212520086960886818157E-853"))
	if err != nil {
		lang.SayString("subx3464")
	} else {
		if !(rexsult.ToString() == "69065510.0459653699418083455335366") {
			lang.SayString("subx3464")
		}
	}
	rexsult, err = lang.RxFromString("-2921982.82411285505229122890041841").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-72994.6523840298017471962569778803E-763"))
	if err != nil {
		lang.SayString("subx3465")
	} else {
		if !(rexsult.ToString() == "-2921982.82411285505229122890041841") {
			lang.SayString("subx3465")
		}
	}
	rexsult, err = lang.RxFromString("4.51117459466491451401835756593793").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("3873385.19981811640063144338087230"))
	if err != nil {
		lang.SayString("subx3466")
	} else {
		if !(rexsult.ToString() == "-3873380.68864352173571692936251474") {
			lang.SayString("subx3466")
		}
	}
	rexsult, err = lang.RxFromString("49553763474698.8118661758811091120").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("36.1713861293896216593840817950781E+410"))
	if err != nil {
		lang.SayString("subx3467")
	} else {
		if !(rexsult.ToString() == "-3.61713861293896216593840817950781E+411") {
			lang.SayString("subx3467")
		}
	}
	rexsult, err = lang.RxFromString("755985583344.379951506071499170749E+956").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("746921095569971477373643487285780"))
	if err != nil {
		lang.SayString("subx3468")
	} else {
		if !(rexsult.ToString() == "7.55985583344379951506071499170749E+967") {
			lang.SayString("subx3468")
		}
	}
	rexsult, err = lang.RxFromString("-20101668541.7472260497594230105456").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-395562148.345003931161532101109964"))
	if err != nil {
		lang.SayString("subx3469")
	} else {
		if !(rexsult.ToString() == "-19706106393.4022221185978909094356") {
			lang.SayString("subx3469")
		}
	}
	rexsult, err = lang.RxFromString("5583903.18072100716072011264668631").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("460868914694.088387067451312500723"))
	if err != nil {
		lang.SayString("subx3470")
	} else {
		if !(rexsult.ToString() == "-460863330790.907666060290592388076") {
			lang.SayString("subx3470")
		}
	}
	rexsult, err = lang.RxFromString("-955638397975240685017992436116257E+020").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-508580.148958769104511751975720470E+662"))
	if err != nil {
		lang.SayString("subx3471")
	} else {
		if !(rexsult.ToString() == "5.08580148958769104511751975720470E+667") {
			lang.SayString("subx3471")
		}
	}
	rexsult, err = lang.RxFromString("-136243796098020983814115584402407E+796").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("6589108083.99750311651581032447390"))
	if err != nil {
		lang.SayString("subx3472")
	} else {
		if !(rexsult.ToString() == "-1.36243796098020983814115584402407E+828") {
			lang.SayString("subx3472")
		}
	}
	rexsult, err = lang.RxFromString("-808498.482718304598213092937543934E+521").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("48005.1465097914355096301483788905"))
	if err != nil {
		lang.SayString("subx3473")
	} else {
		if !(rexsult.ToString() == "-8.08498482718304598213092937543934E+526") {
			lang.SayString("subx3473")
		}
	}
	rexsult, err = lang.RxFromString("-812.266340554281305985524813074211E+396").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-3195.63111559114001594257448970812E+986"))
	if err != nil {
		lang.SayString("subx3474")
	} else {
		if !(rexsult.ToString() == "3.19563111559114001594257448970812E+989") {
			lang.SayString("subx3474")
		}
	}
	rexsult, err = lang.RxFromString("-929889.720905183397678866648217793E+134").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-280300190774.057595671079264841349"))
	if err != nil {
		lang.SayString("subx3475")
	} else {
		if !(rexsult.ToString() == "-9.29889720905183397678866648217793E+139") {
			lang.SayString("subx3475")
		}
	}
	rexsult, err = lang.RxFromString("83946.0157801953636255078996185540").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("492670373.235391665758701500314473"))
	if err != nil {
		lang.SayString("subx3476")
	} else {
		if !(rexsult.ToString() == "-492586427.219611470395075992414855") {
			lang.SayString("subx3476")
		}
	}
	rexsult, err = lang.RxFromString("7812758113817.99135851825223122508").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("3037492.36716301443309571918002123E-157"))
	if err != nil {
		lang.SayString("subx3477")
	} else {
		if !(rexsult.ToString() == "7812758113817.99135851825223122508") {
			lang.SayString("subx3477")
		}
	}
	rexsult, err = lang.RxFromString("489191747.148674326757767356626016").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("01136942.1182277580093027768490545"))
	if err != nil {
		lang.SayString("subx3478")
	} else {
		if !(rexsult.ToString() == "488054805.030446568748464579776962") {
			lang.SayString("subx3478")
		}
	}
	rexsult, err = lang.RxFromString("-599369540.373174482335865567937853E+289").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-38288383205.6749439588707955585209"))
	if err != nil {
		lang.SayString("subx3479")
	} else {
		if !(rexsult.ToString() == "-5.99369540373174482335865567937853E+297") {
			lang.SayString("subx3479")
		}
	}
	rexsult, err = lang.RxFromString("-3376883870.85961692148022521960139").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-65247489449.7334589731171980408284"))
	if err != nil {
		lang.SayString("subx3480")
	} else {
		if !(rexsult.ToString() == "61870605578.8738420516369728212270") {
			lang.SayString("subx3480")
		}
	}
	rexsult, err = lang.RxFromString("58.6776780370008364590621772421025").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("01.5925518866529044494309229975160"))
	if err != nil {
		lang.SayString("subx3481")
	} else {
		if !(rexsult.ToString() == "57.0851261503479320096312542445865") {
			lang.SayString("subx3481")
		}
	}
	rexsult, err = lang.RxFromString("4099616339.96249499552808575717579").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("290.795187361072489816791525139895"))
	if err != nil {
		lang.SayString("subx3482")
	} else {
		if !(rexsult.ToString() == "4099616049.16730763445559594038427") {
			lang.SayString("subx3482")
		}
	}
	rexsult, err = lang.RxFromString("85870777.2282833141709970713739108").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-2140392861153.69401346398478113715"))
	if err != nil {
		lang.SayString("subx3483")
	} else {
		if !(rexsult.ToString() == "2140478731930.92229677815577820852") {
			lang.SayString("subx3483")
		}
	}
	rexsult, err = lang.RxFromString("20900.9693761555165742010339929779").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-38.7546147649523793463260940289585"))
	if err != nil {
		lang.SayString("subx3484")
	} else {
		if !(rexsult.ToString() == "20939.7239909204689535473600870069") {
			lang.SayString("subx3484")
		}
	}
	rexsult, err = lang.RxFromString("448.827596155587910947511170319456").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("379130153.382794042652974596286062"))
	if err != nil {
		lang.SayString("subx3485")
	} else {
		if !(rexsult.ToString() == "-379129704.555197887065063648774892") {
			lang.SayString("subx3485")
		}
	}
	rexsult, err = lang.RxFromString("98.4134807921002817357000140482039").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("3404725543.77032945444654351546779"))
	if err != nil {
		lang.SayString("subx3486")
	} else {
		if !(rexsult.ToString() == "-3404725445.35684866234626177976778") {
			lang.SayString("subx3486")
		}
	}
	rexsult, err = lang.RxFromString("545746433.649359734136476718176330E-787").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-5149957099709.12830072802043560650E-437"))
	if err != nil {
		lang.SayString("subx3487")
	} else {
		if !(rexsult.ToString() == "5.14995709970912830072802043560650E-425") {
			lang.SayString("subx3487")
		}
	}
	rexsult, err = lang.RxFromString("741304513547.273820525801608231737").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("0396.22823128272584928019323186355E-830"))
	if err != nil {
		lang.SayString("subx3488")
	} else {
		if !(rexsult.ToString() == "741304513547.273820525801608231737") {
			lang.SayString("subx3488")
		}
	}
	rexsult, err = lang.RxFromString("-706.145005094292315613907254240553").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("4739.82486195739758308735946332234"))
	if err != nil {
		lang.SayString("subx3489")
	} else {
		if !(rexsult.ToString() == "-5445.96986705168989870126671756289") {
			lang.SayString("subx3489")
		}
	}
	rexsult, err = lang.RxFromString("-769925786.823099083228795187975893").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-31201.9980469760239870067820594790"))
	if err != nil {
		lang.SayString("subx3490")
	} else {
		if !(rexsult.ToString() == "-769894584.825052107204808181193834") {
			lang.SayString("subx3490")
		}
	}
	rexsult, err = lang.RxFromString("84438610546049.7256507419289692857E+906").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("052604240766736461898844111790311"))
	if err != nil {
		lang.SayString("subx3491")
	} else {
		if !(rexsult.ToString() == "8.44386105460497256507419289692857E+919") {
			lang.SayString("subx3491")
		}
	}
	rexsult, err = lang.RxFromString("549760.911304725795164589619286514").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("165.160089615604924207754883953484"))
	if err != nil {
		lang.SayString("subx3492")
	} else {
		if !(rexsult.ToString() == "549595.751215110190240381864402561") {
			lang.SayString("subx3492")
		}
	}
	rexsult, err = lang.RxFromString("3650514.18649737956855828939662794").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("08086721.4036886948248274834735629"))
	if err != nil {
		lang.SayString("subx3493")
	} else {
		if !(rexsult.ToString() == "-4436207.21719131525626919407693496") {
			lang.SayString("subx3493")
		}
	}
	rexsult, err = lang.RxFromString("55067723881950.1346958179604099594").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-8.90481481687182931431054785192083"))
	if err != nil {
		lang.SayString("subx3494")
	} else {
		if !(rexsult.ToString() == "55067723881959.0395106348322392737") {
			lang.SayString("subx3494")
		}
	}
	rexsult, err = lang.RxFromString("868251123.413992653362860637541060E+019").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("5579665045.37858308541154858567656E+131"))
	if err != nil {
		lang.SayString("subx3495")
	} else {
		if !(rexsult.ToString() == "-5.57966504537858308541154858567656E+140") {
			lang.SayString("subx3495")
		}
	}
	rexsult, err = lang.RxFromString("-646.464431574014407536004990059069").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-798.679560020414523841321724649594E-037"))
	if err != nil {
		lang.SayString("subx3496")
	} else {
		if !(rexsult.ToString() == "-646.464431574014407536004990059069") {
			lang.SayString("subx3496")
		}
	}
	rexsult, err = lang.RxFromString("354.546679975219753598558273421556").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-7039.46386812239015144581761752927E-448"))
	if err != nil {
		lang.SayString("subx3497")
	} else {
		if !(rexsult.ToString() == "354.546679975219753598558273421556") {
			lang.SayString("subx3497")
		}
	}
	rexsult, err = lang.RxFromString("91936087917435.5974889495278215874").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("-67080823344.8903392584327136082486E-757"))
	if err != nil {
		lang.SayString("subx3498")
	} else {
		if !(rexsult.ToString() == "91936087917435.5974889495278215874") {
			lang.SayString("subx3498")
		}
	}
	rexsult, err = lang.RxFromString("-07345.6422518528556136521417259811E-600").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("41188325.7041362608934957584583381E-763"))
	if err != nil {
		lang.SayString("subx3499")
	} else {
		if !(rexsult.ToString() == "-7.34564225185285561365214172598110E-597") {
			lang.SayString("subx3499")
		}
	}
	rexsult, err = lang.RxFromString("-253280724.939458021588167965038184").OpSub(lang.RxSetWithDigit(33), lang.RxFromString("616988.426425908872398170896375634E+396"))
	if err != nil {
		lang.SayString("subx3500")
	} else {
		if !(rexsult.ToString() == "-6.16988426425908872398170896375634E+401") {
			lang.SayString("subx3500")
		}
	}
	rexsult, err = lang.RxFromString("905.67402").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-202896611.E-780472620"))
	if err != nil {
		lang.SayString("xsub001")
	} else {
		if !(rexsult.ToString() == "905.674020") {
			lang.SayString("xsub001")
		}
	}
	rexsult, err = lang.RxFromString("3915134.7").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-597164907."))
	if err != nil {
		lang.SayString("xsub002")
	} else {
		if !(rexsult.ToString() == "601080042") {
			lang.SayString("xsub002")
		}
	}
	rexsult, err = lang.RxFromString("309759261").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("62663.487"))
	if err != nil {
		lang.SayString("xsub003")
	} else {
		if !(rexsult.ToString() == "309696598") {
			lang.SayString("xsub003")
		}
	}
	rexsult, err = lang.RxFromString("3.93591888E-236595626").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7242375.00"))
	if err != nil {
		lang.SayString("xsub004")
	} else {
		if !(rexsult.ToString() == "-7242375.00") {
			lang.SayString("xsub004")
		}
	}
	rexsult, err = lang.RxFromString("323902.714").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-608669.607E-657060568"))
	if err != nil {
		lang.SayString("xsub005")
	} else {
		if !(rexsult.ToString() == "323902.714") {
			lang.SayString("xsub005")
		}
	}
	rexsult, err = lang.RxFromString("5.11970092").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-8807.22036"))
	if err != nil {
		lang.SayString("xsub006")
	} else {
		if !(rexsult.ToString() == "8812.34006") {
			lang.SayString("xsub006")
		}
	}
	rexsult, err = lang.RxFromString("-7.99874516").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("4561.83758"))
	if err != nil {
		lang.SayString("xsub007")
	} else {
		if !(rexsult.ToString() == "-4569.83633") {
			lang.SayString("xsub007")
		}
	}
	rexsult, err = lang.RxFromString("297802878").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-927206.324"))
	if err != nil {
		lang.SayString("xsub008")
	} else {
		if !(rexsult.ToString() == "298730084") {
			lang.SayString("xsub008")
		}
	}
	rexsult, err = lang.RxFromString("-766.651824").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("31300.3619"))
	if err != nil {
		lang.SayString("xsub009")
	} else {
		if !(rexsult.ToString() == "-32067.0137") {
			lang.SayString("xsub009")
		}
	}
	rexsult, err = lang.RxFromString("-56746.8689E+934981942").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("471002521."))
	if err != nil {
		lang.SayString("xsub010")
	} else {
		if !(rexsult.ToString() == "-5.67468689E+934981946") {
			lang.SayString("xsub010")
		}
	}
	rexsult, err = lang.RxFromString("456417160").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-41346.1024"))
	if err != nil {
		lang.SayString("xsub011")
	} else {
		if !(rexsult.ToString() == "456458506") {
			lang.SayString("xsub011")
		}
	}
	rexsult, err = lang.RxFromString("102895.722").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2.62214826"))
	if err != nil {
		lang.SayString("xsub012")
	} else {
		if !(rexsult.ToString() == "102898.344") {
			lang.SayString("xsub012")
		}
	}
	rexsult, err = lang.RxFromString("61.3033331E+157644141").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-567740.918E-893439456"))
	if err != nil {
		lang.SayString("xsub013")
	} else {
		if !(rexsult.ToString() == "6.13033331E+157644142") {
			lang.SayString("xsub013")
		}
	}
	rexsult, err = lang.RxFromString("80223.3897").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("73921.0383E-467772675"))
	if err != nil {
		lang.SayString("xsub014")
	} else {
		if !(rexsult.ToString() == "80223.3897") {
			lang.SayString("xsub014")
		}
	}
	rexsult, err = lang.RxFromString("-654645.954").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-9.12535752"))
	if err != nil {
		lang.SayString("xsub015")
	} else {
		if !(rexsult.ToString() == "-654636.829") {
			lang.SayString("xsub015")
		}
	}
	rexsult, err = lang.RxFromString("63.1917772E-706014634").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7.56253257E-138579234"))
	if err != nil {
		lang.SayString("xsub016")
	} else {
		if !(rexsult.ToString() == "7.56253257E-138579234") {
			lang.SayString("xsub016")
		}
	}
	rexsult, err = lang.RxFromString("-39674.7190").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2490607.78"))
	if err != nil {
		lang.SayString("xsub017")
	} else {
		if !(rexsult.ToString() == "-2530282.50") {
			lang.SayString("xsub017")
		}
	}
	rexsult, err = lang.RxFromString("-3364.59737E-600363681").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("896487.451"))
	if err != nil {
		lang.SayString("xsub018")
	} else {
		if !(rexsult.ToString() == "-896487.451") {
			lang.SayString("xsub018")
		}
	}
	rexsult, err = lang.RxFromString("-64138.0578").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("31759011.3E+697488342"))
	if err != nil {
		lang.SayString("xsub019")
	} else {
		if !(rexsult.ToString() == "-3.17590113E+697488349") {
			lang.SayString("xsub019")
		}
	}
	rexsult, err = lang.RxFromString("61399.8527").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-64344484.5"))
	if err != nil {
		lang.SayString("xsub020")
	} else {
		if !(rexsult.ToString() == "64405884.4") {
			lang.SayString("xsub020")
		}
	}
	rexsult, err = lang.RxFromString("-722960.204").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-26154599.8"))
	if err != nil {
		lang.SayString("xsub021")
	} else {
		if !(rexsult.ToString() == "25431639.6") {
			lang.SayString("xsub021")
		}
	}
	rexsult, err = lang.RxFromString("9.47109959E+230565093").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("73354723.2"))
	if err != nil {
		lang.SayString("xsub022")
	} else {
		if !(rexsult.ToString() == "9.47109959E+230565093") {
			lang.SayString("xsub022")
		}
	}
	rexsult, err = lang.RxFromString("43.7456245").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("547441956."))
	if err != nil {
		lang.SayString("xsub023")
	} else {
		if !(rexsult.ToString() == "-547441912") {
			lang.SayString("xsub023")
		}
	}
	rexsult, err = lang.RxFromString("-73150542E-242017390").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-8.15869954"))
	if err != nil {
		lang.SayString("xsub024")
	} else {
		if !(rexsult.ToString() == "8.15869954") {
			lang.SayString("xsub024")
		}
	}
	rexsult, err = lang.RxFromString("2015.62109E+299897596").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-11788916.1"))
	if err != nil {
		lang.SayString("xsub025")
	} else {
		if !(rexsult.ToString() == "2.01562109E+299897599") {
			lang.SayString("xsub025")
		}
	}
	rexsult, err = lang.RxFromString("29.498114").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-26486451"))
	if err != nil {
		lang.SayString("xsub026")
	} else {
		if !(rexsult.ToString() == "26486480.5") {
			lang.SayString("xsub026")
		}
	}
	rexsult, err = lang.RxFromString("244375043.E+130840878").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-9.44522029"))
	if err != nil {
		lang.SayString("xsub027")
	} else {
		if !(rexsult.ToString() == "2.44375043E+130840886") {
			lang.SayString("xsub027")
		}
	}
	rexsult, err = lang.RxFromString("-349388.759").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-196215.776"))
	if err != nil {
		lang.SayString("xsub028")
	} else {
		if !(rexsult.ToString() == "-153172.983") {
			lang.SayString("xsub028")
		}
	}
	rexsult, err = lang.RxFromString("-70905112.4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-91353968.8"))
	if err != nil {
		lang.SayString("xsub029")
	} else {
		if !(rexsult.ToString() == "20448856.4") {
			lang.SayString("xsub029")
		}
	}
	rexsult, err = lang.RxFromString("-225094.28").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-88.7723542"))
	if err != nil {
		lang.SayString("xsub030")
	} else {
		if !(rexsult.ToString() == "-225005.508") {
			lang.SayString("xsub030")
		}
	}
	rexsult, err = lang.RxFromString("50.4442340").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("82.7952169E+880120759"))
	if err != nil {
		lang.SayString("xsub031")
	} else {
		if !(rexsult.ToString() == "-8.27952169E+880120760") {
			lang.SayString("xsub031")
		}
	}
	rexsult, err = lang.RxFromString("-32311.9037").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8.36379449"))
	if err != nil {
		lang.SayString("xsub032")
	} else {
		if !(rexsult.ToString() == "-32320.2675") {
			lang.SayString("xsub032")
		}
	}
	rexsult, err = lang.RxFromString("615396156.E+549895291").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-29530247.4"))
	if err != nil {
		lang.SayString("xsub033")
	} else {
		if !(rexsult.ToString() == "6.15396156E+549895299") {
			lang.SayString("xsub033")
		}
	}
	rexsult, err = lang.RxFromString("592.142173E-419941416").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3.46079109E-844011845"))
	if err != nil {
		lang.SayString("xsub034")
	} else {
		if !(rexsult.ToString() == "5.92142173E-419941414") {
			lang.SayString("xsub034")
		}
	}
	rexsult, err = lang.RxFromString("849.515993E-878446473").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1039.08778"))
	if err != nil {
		lang.SayString("xsub035")
	} else {
		if !(rexsult.ToString() == "1039.08778") {
			lang.SayString("xsub035")
		}
	}
	rexsult, err = lang.RxFromString("213361789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-599.644851"))
	if err != nil {
		lang.SayString("xsub036")
	} else {
		if !(rexsult.ToString() == "213362389") {
			lang.SayString("xsub036")
		}
	}
	rexsult, err = lang.RxFromString("-795522555.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-298.037702"))
	if err != nil {
		lang.SayString("xsub037")
	} else {
		if !(rexsult.ToString() == "-795522257") {
			lang.SayString("xsub037")
		}
	}
	rexsult, err = lang.RxFromString("-501260651.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-8761893.0E-689281479"))
	if err != nil {
		lang.SayString("xsub038")
	} else {
		if !(rexsult.ToString() == "-501260651") {
			lang.SayString("xsub038")
		}
	}
	rexsult, err = lang.RxFromString("-1.70781105E-848889023").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("36504769.4"))
	if err != nil {
		lang.SayString("xsub039")
	} else {
		if !(rexsult.ToString() == "-36504769.4") {
			lang.SayString("xsub039")
		}
	}
	rexsult, err = lang.RxFromString("-5290.54984E-490626676").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("842535254"))
	if err != nil {
		lang.SayString("xsub040")
	} else {
		if !(rexsult.ToString() == "-842535254") {
			lang.SayString("xsub040")
		}
	}
	rexsult, err = lang.RxFromString("608.31825E+535268120").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-59609.0993"))
	if err != nil {
		lang.SayString("xsub041")
	} else {
		if !(rexsult.ToString() == "6.08318250E+535268122") {
			lang.SayString("xsub041")
		}
	}
	rexsult, err = lang.RxFromString("-4629035.31").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-167.884398"))
	if err != nil {
		lang.SayString("xsub042")
	} else {
		if !(rexsult.ToString() == "-4628867.43") {
			lang.SayString("xsub042")
		}
	}
	rexsult, err = lang.RxFromString("-66527378.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-706400268."))
	if err != nil {
		lang.SayString("xsub043")
	} else {
		if !(rexsult.ToString() == "639872890") {
			lang.SayString("xsub043")
		}
	}
	rexsult, err = lang.RxFromString("-2510497.53").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("372882462."))
	if err != nil {
		lang.SayString("xsub044")
	} else {
		if !(rexsult.ToString() == "-375392960") {
			lang.SayString("xsub044")
		}
	}
	rexsult, err = lang.RxFromString("136.255393E+53329961").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-53494.7201E+720058060"))
	if err != nil {
		lang.SayString("xsub045")
	} else {
		if !(rexsult.ToString() == "5.34947201E+720058064") {
			lang.SayString("xsub045")
		}
	}
	rexsult, err = lang.RxFromString("-876673.100").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-6150.92266"))
	if err != nil {
		lang.SayString("xsub046")
	} else {
		if !(rexsult.ToString() == "-870522.177") {
			lang.SayString("xsub046")
		}
	}
	rexsult, err = lang.RxFromString("-2.45151797E+911306117").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("27235771"))
	if err != nil {
		lang.SayString("xsub047")
	} else {
		if !(rexsult.ToString() == "-2.45151797E+911306117") {
			lang.SayString("xsub047")
		}
	}
	rexsult, err = lang.RxFromString("-9.15117551").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4.95100733E-314511326"))
	if err != nil {
		lang.SayString("xsub048")
	} else {
		if !(rexsult.ToString() == "-9.15117551") {
			lang.SayString("xsub048")
		}
	}
	rexsult, err = lang.RxFromString("3.61890453E-985606128").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("30664416."))
	if err != nil {
		lang.SayString("xsub049")
	} else {
		if !(rexsult.ToString() == "-30664416.0") {
			lang.SayString("xsub049")
		}
	}
	rexsult, err = lang.RxFromString("-257674602E+216723382").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-70820959.4"))
	if err != nil {
		lang.SayString("xsub050")
	} else {
		if !(rexsult.ToString() == "-2.57674602E+216723390") {
			lang.SayString("xsub050")
		}
	}
	rexsult, err = lang.RxFromString("218699.206").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("556944241."))
	if err != nil {
		lang.SayString("xsub051")
	} else {
		if !(rexsult.ToString() == "-556725542") {
			lang.SayString("xsub051")
		}
	}
	rexsult, err = lang.RxFromString("106211716.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3456793.74"))
	if err != nil {
		lang.SayString("xsub052")
	} else {
		if !(rexsult.ToString() == "109668510") {
			lang.SayString("xsub052")
		}
	}
	rexsult, err = lang.RxFromString("1.25018078").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("399856.763E-726816740"))
	if err != nil {
		lang.SayString("xsub053")
	} else {
		if !(rexsult.ToString() == "1.25018078") {
			lang.SayString("xsub053")
		}
	}
	rexsult, err = lang.RxFromString("364.99811").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-46222.0505"))
	if err != nil {
		lang.SayString("xsub054")
	} else {
		if !(rexsult.ToString() == "46587.0486") {
			lang.SayString("xsub054")
		}
	}
	rexsult, err = lang.RxFromString("-392217576.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-958364096"))
	if err != nil {
		lang.SayString("xsub055")
	} else {
		if !(rexsult.ToString() == "566146520") {
			lang.SayString("xsub055")
		}
	}
	rexsult, err = lang.RxFromString("169601285").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("714526.639"))
	if err != nil {
		lang.SayString("xsub056")
	} else {
		if !(rexsult.ToString() == "168886758") {
			lang.SayString("xsub056")
		}
	}
	rexsult, err = lang.RxFromString("-674.094552E+586944319").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("6354.2668E+589657266"))
	if err != nil {
		lang.SayString("xsub057")
	} else {
		if !(rexsult.ToString() == "-6.35426680E+589657269") {
			lang.SayString("xsub057")
		}
	}
	rexsult, err = lang.RxFromString("151795163E-371727182").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-488.09788E-738852245"))
	if err != nil {
		lang.SayString("xsub058")
	} else {
		if !(rexsult.ToString() == "1.51795163E-371727174") {
			lang.SayString("xsub058")
		}
	}
	rexsult, err = lang.RxFromString("-746.293386").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("927749.647"))
	if err != nil {
		lang.SayString("xsub059")
	} else {
		if !(rexsult.ToString() == "-928495.940") {
			lang.SayString("xsub059")
		}
	}
	rexsult, err = lang.RxFromString("888946471E+241331592").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-235739.595"))
	if err != nil {
		lang.SayString("xsub060")
	} else {
		if !(rexsult.ToString() == "8.88946471E+241331600") {
			lang.SayString("xsub060")
		}
	}
	rexsult, err = lang.RxFromString("6.64377249").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("79161.1070E+619453776"))
	if err != nil {
		lang.SayString("xsub061")
	} else {
		if !(rexsult.ToString() == "-7.91611070E+619453780") {
			lang.SayString("xsub061")
		}
	}
	rexsult, err = lang.RxFromString("3146.66571E-313373366").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("88.5282010"))
	if err != nil {
		lang.SayString("xsub062")
	} else {
		if !(rexsult.ToString() == "-88.5282010") {
			lang.SayString("xsub062")
		}
	}
	rexsult, err = lang.RxFromString("6.44693097").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-87195.8711"))
	if err != nil {
		lang.SayString("xsub063")
	} else {
		if !(rexsult.ToString() == "87202.3180") {
			lang.SayString("xsub063")
		}
	}
	rexsult, err = lang.RxFromString("-2113132.56E+577957840").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("981125821"))
	if err != nil {
		lang.SayString("xsub064")
	} else {
		if !(rexsult.ToString() == "-2.11313256E+577957846") {
			lang.SayString("xsub064")
		}
	}
	rexsult, err = lang.RxFromString("-7701.42814").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("72667.5181"))
	if err != nil {
		lang.SayString("xsub065")
	} else {
		if !(rexsult.ToString() == "-80368.9462") {
			lang.SayString("xsub065")
		}
	}
	rexsult, err = lang.RxFromString("-851.754789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-582659.149"))
	if err != nil {
		lang.SayString("xsub066")
	} else {
		if !(rexsult.ToString() == "581807.394") {
			lang.SayString("xsub066")
		}
	}
	rexsult, err = lang.RxFromString("-5.01992943").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7852.16531"))
	if err != nil {
		lang.SayString("xsub067")
	} else {
		if !(rexsult.ToString() == "-7857.18524") {
			lang.SayString("xsub067")
		}
	}
	rexsult, err = lang.RxFromString("-12393257.2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("76803689E+949125770"))
	if err != nil {
		lang.SayString("xsub068")
	} else {
		if !(rexsult.ToString() == "-7.68036890E+949125777") {
			lang.SayString("xsub068")
		}
	}
	rexsult, err = lang.RxFromString("-754771634.E+716555026").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-292336.311"))
	if err != nil {
		lang.SayString("xsub069")
	} else {
		if !(rexsult.ToString() == "-7.54771634E+716555034") {
			lang.SayString("xsub069")
		}
	}
	rexsult, err = lang.RxFromString("-915006.171E+614548652").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-314086965."))
	if err != nil {
		lang.SayString("xsub070")
	} else {
		if !(rexsult.ToString() == "-9.15006171E+614548657") {
			lang.SayString("xsub070")
		}
	}
	rexsult, err = lang.RxFromString("-296590035").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-481734529"))
	if err != nil {
		lang.SayString("xsub071")
	} else {
		if !(rexsult.ToString() == "185144494") {
			lang.SayString("xsub071")
		}
	}
	rexsult, err = lang.RxFromString("8.27822605").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("9241557.19"))
	if err != nil {
		lang.SayString("xsub072")
	} else {
		if !(rexsult.ToString() == "-9241548.91") {
			lang.SayString("xsub072")
		}
	}
	rexsult, err = lang.RxFromString("-1.43581098").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7286313.54"))
	if err != nil {
		lang.SayString("xsub073")
	} else {
		if !(rexsult.ToString() == "-7286314.98") {
			lang.SayString("xsub073")
		}
	}
	rexsult, err = lang.RxFromString("-699036193.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("759263.509E+533543625"))
	if err != nil {
		lang.SayString("xsub074")
	} else {
		if !(rexsult.ToString() == "-7.59263509E+533543630") {
			lang.SayString("xsub074")
		}
	}
	rexsult, err = lang.RxFromString("-83.7273615E-305281957").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-287779593.E+458777774"))
	if err != nil {
		lang.SayString("xsub075")
	} else {
		if !(rexsult.ToString() == "2.87779593E+458777782") {
			lang.SayString("xsub075")
		}
	}
	rexsult, err = lang.RxFromString("8.48503224").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("6522.03316"))
	if err != nil {
		lang.SayString("xsub076")
	} else {
		if !(rexsult.ToString() == "-6513.54813") {
			lang.SayString("xsub076")
		}
	}
	rexsult, err = lang.RxFromString("527916091").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-809.054070"))
	if err != nil {
		lang.SayString("xsub077")
	} else {
		if !(rexsult.ToString() == "527916900") {
			lang.SayString("xsub077")
		}
	}
	rexsult, err = lang.RxFromString("3857058.60").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("5792997.58E+881077409"))
	if err != nil {
		lang.SayString("xsub078")
	} else {
		if !(rexsult.ToString() == "-5.79299758E+881077415") {
			lang.SayString("xsub078")
		}
	}
	rexsult, err = lang.RxFromString("-66587363.E+556538173").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-551902402E+357309146"))
	if err != nil {
		lang.SayString("xsub079")
	} else {
		if !(rexsult.ToString() == "-6.65873630E+556538180") {
			lang.SayString("xsub079")
		}
	}
	rexsult, err = lang.RxFromString("-580.502955").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("38125521.7"))
	if err != nil {
		lang.SayString("xsub080")
	} else {
		if !(rexsult.ToString() == "-38126102.2") {
			lang.SayString("xsub080")
		}
	}
	rexsult, err = lang.RxFromString("-9627363.00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-80616885E-749891394"))
	if err != nil {
		lang.SayString("xsub081")
	} else {
		if !(rexsult.ToString() == "-9627363.00") {
			lang.SayString("xsub081")
		}
	}
	rexsult, err = lang.RxFromString("-526.594855E+803110107").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-64.5451639"))
	if err != nil {
		lang.SayString("xsub082")
	} else {
		if !(rexsult.ToString() == "-5.26594855E+803110109") {
			lang.SayString("xsub082")
		}
	}
	rexsult, err = lang.RxFromString("-8378.55499").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("760.131257"))
	if err != nil {
		lang.SayString("xsub083")
	} else {
		if !(rexsult.ToString() == "-9138.68625") {
			lang.SayString("xsub083")
		}
	}
	rexsult, err = lang.RxFromString("-717.697718").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("984304413"))
	if err != nil {
		lang.SayString("xsub084")
	} else {
		if !(rexsult.ToString() == "-984305131") {
			lang.SayString("xsub084")
		}
	}
	rexsult, err = lang.RxFromString("-76762243.4E-741100094").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-273.706674"))
	if err != nil {
		lang.SayString("xsub085")
	} else {
		if !(rexsult.ToString() == "273.706674") {
			lang.SayString("xsub085")
		}
	}
	rexsult, err = lang.RxFromString("-701.518354E+786274918").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8822750.68E+243052107"))
	if err != nil {
		lang.SayString("xsub086")
	} else {
		if !(rexsult.ToString() == "-7.01518354E+786274920") {
			lang.SayString("xsub086")
		}
	}
	rexsult, err = lang.RxFromString("-359866845.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4.57434117"))
	if err != nil {
		lang.SayString("xsub087")
	} else {
		if !(rexsult.ToString() == "-359866841") {
			lang.SayString("xsub087")
		}
	}
	rexsult, err = lang.RxFromString("779934536.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-76562645.7"))
	if err != nil {
		lang.SayString("xsub088")
	} else {
		if !(rexsult.ToString() == "856497182") {
			lang.SayString("xsub088")
		}
	}
	rexsult, err = lang.RxFromString("-4820.95451").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3516234.99E+303303176"))
	if err != nil {
		lang.SayString("xsub089")
	} else {
		if !(rexsult.ToString() == "-3.51623499E+303303182") {
			lang.SayString("xsub089")
		}
	}
	rexsult, err = lang.RxFromString("69355976.9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-9.57838562E+758804984"))
	if err != nil {
		lang.SayString("xsub090")
	} else {
		if !(rexsult.ToString() == "9.57838562E+758804984") {
			lang.SayString("xsub090")
		}
	}
	rexsult, err = lang.RxFromString("-12672093.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8569.78255E-382866025"))
	if err != nil {
		lang.SayString("xsub091")
	} else {
		if !(rexsult.ToString() == "-12672093.1") {
			lang.SayString("xsub091")
		}
	}
	rexsult, err = lang.RxFromString("-5910750.2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("66150383E-662459241"))
	if err != nil {
		lang.SayString("xsub092")
	} else {
		if !(rexsult.ToString() == "-5910750.20") {
			lang.SayString("xsub092")
		}
	}
	rexsult, err = lang.RxFromString("-532577268.E-163806629").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-240650398E-650110558"))
	if err != nil {
		lang.SayString("xsub093")
	} else {
		if !(rexsult.ToString() == "-5.32577268E-163806621") {
			lang.SayString("xsub093")
		}
	}
	rexsult, err = lang.RxFromString("-671.507198E-908587890").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3057429.32E-555230623"))
	if err != nil {
		lang.SayString("xsub094")
	} else {
		if !(rexsult.ToString() == "-3.05742932E-555230617") {
			lang.SayString("xsub094")
		}
	}
	rexsult, err = lang.RxFromString("-294.994352E+346452027").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-6061853.0"))
	if err != nil {
		lang.SayString("xsub095")
	} else {
		if !(rexsult.ToString() == "-2.94994352E+346452029") {
			lang.SayString("xsub095")
		}
	}
	rexsult, err = lang.RxFromString("329579114").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("146780548."))
	if err != nil {
		lang.SayString("xsub096")
	} else {
		if !(rexsult.ToString() == "182798566") {
			lang.SayString("xsub096")
		}
	}
	rexsult, err = lang.RxFromString("-789904.686E-217225000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1991.07181E-84080059"))
	if err != nil {
		lang.SayString("xsub097")
	} else {
		if !(rexsult.ToString() == "1.99107181E-84080056") {
			lang.SayString("xsub097")
		}
	}
	rexsult, err = lang.RxFromString("59893.3544").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-408595868"))
	if err != nil {
		lang.SayString("xsub098")
	} else {
		if !(rexsult.ToString() == "408655761") {
			lang.SayString("xsub098")
		}
	}
	rexsult, err = lang.RxFromString("129.878613").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-54652.7288E-963564940"))
	if err != nil {
		lang.SayString("xsub099")
	} else {
		if !(rexsult.ToString() == "129.878613") {
			lang.SayString("xsub099")
		}
	}
	rexsult, err = lang.RxFromString("9866.99208").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("708756501."))
	if err != nil {
		lang.SayString("xsub100")
	} else {
		if !(rexsult.ToString() == "-708746634") {
			lang.SayString("xsub100")
		}
	}
	rexsult, err = lang.RxFromString("-78810.6297").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-399884.68"))
	if err != nil {
		lang.SayString("xsub101")
	} else {
		if !(rexsult.ToString() == "321074.050") {
			lang.SayString("xsub101")
		}
	}
	rexsult, err = lang.RxFromString("409189761").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-771.471460"))
	if err != nil {
		lang.SayString("xsub102")
	} else {
		if !(rexsult.ToString() == "409190532") {
			lang.SayString("xsub102")
		}
	}
	rexsult, err = lang.RxFromString("-1.68748838").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("460.46924"))
	if err != nil {
		lang.SayString("xsub103")
	} else {
		if !(rexsult.ToString() == "-462.156728") {
			lang.SayString("xsub103")
		}
	}
	rexsult, err = lang.RxFromString("553527296.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7924.40185"))
	if err != nil {
		lang.SayString("xsub104")
	} else {
		if !(rexsult.ToString() == "553535220") {
			lang.SayString("xsub104")
		}
	}
	rexsult, err = lang.RxFromString("-38.7465207").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("64936.2942"))
	if err != nil {
		lang.SayString("xsub105")
	} else {
		if !(rexsult.ToString() == "-64975.0407") {
			lang.SayString("xsub105")
		}
	}
	rexsult, err = lang.RxFromString("-201075.248").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("845.663928"))
	if err != nil {
		lang.SayString("xsub106")
	} else {
		if !(rexsult.ToString() == "-201920.912") {
			lang.SayString("xsub106")
		}
	}
	rexsult, err = lang.RxFromString("91048.4559").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("75953609.3"))
	if err != nil {
		lang.SayString("xsub107")
	} else {
		if !(rexsult.ToString() == "-75862560.9") {
			lang.SayString("xsub107")
		}
	}
	rexsult, err = lang.RxFromString("6898273.86E-252097460").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("15.3456196"))
	if err != nil {
		lang.SayString("xsub108")
	} else {
		if !(rexsult.ToString() == "-15.3456196") {
			lang.SayString("xsub108")
		}
	}
	rexsult, err = lang.RxFromString("88.4370343").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-980709105E-869899289"))
	if err != nil {
		lang.SayString("xsub109")
	} else {
		if !(rexsult.ToString() == "88.4370343") {
			lang.SayString("xsub109")
		}
	}
	rexsult, err = lang.RxFromString("-17643.39").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2.0352568E+304871331"))
	if err != nil {
		lang.SayString("xsub110")
	} else {
		if !(rexsult.ToString() == "-2.03525680E+304871331") {
			lang.SayString("xsub110")
		}
	}
	rexsult, err = lang.RxFromString("4589785.16").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7459.04237"))
	if err != nil {
		lang.SayString("xsub111")
	} else {
		if !(rexsult.ToString() == "4582326.12") {
			lang.SayString("xsub111")
		}
	}
	rexsult, err = lang.RxFromString("-51.1632090E-753968082").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8.96207471E-585797887"))
	if err != nil {
		lang.SayString("xsub112")
	} else {
		if !(rexsult.ToString() == "-8.96207471E-585797887") {
			lang.SayString("xsub112")
		}
	}
	rexsult, err = lang.RxFromString("982.217817").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7224909.4E-45243816"))
	if err != nil {
		lang.SayString("xsub113")
	} else {
		if !(rexsult.ToString() == "982.217817") {
			lang.SayString("xsub113")
		}
	}
	rexsult, err = lang.RxFromString("503712056.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-57490703.5E+924890183"))
	if err != nil {
		lang.SayString("xsub114")
	} else {
		if !(rexsult.ToString() == "5.74907035E+924890190") {
			lang.SayString("xsub114")
		}
	}
	rexsult, err = lang.RxFromString("883.849223").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("249259171"))
	if err != nil {
		lang.SayString("xsub115")
	} else {
		if !(rexsult.ToString() == "-249258287") {
			lang.SayString("xsub115")
		}
	}
	rexsult, err = lang.RxFromString("245.092853E+872642874").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("828195.152E+419771532"))
	if err != nil {
		lang.SayString("xsub116")
	} else {
		if !(rexsult.ToString() == "2.45092853E+872642876") {
			lang.SayString("xsub116")
		}
	}
	rexsult, err = lang.RxFromString("-83658638.6E+728551928").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2952478.42"))
	if err != nil {
		lang.SayString("xsub117")
	} else {
		if !(rexsult.ToString() == "-8.36586386E+728551935") {
			lang.SayString("xsub117")
		}
	}
	rexsult, err = lang.RxFromString("-6291780.97").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("269967.394E-22000817"))
	if err != nil {
		lang.SayString("xsub118")
	} else {
		if !(rexsult.ToString() == "-6291780.97") {
			lang.SayString("xsub118")
		}
	}
	rexsult, err = lang.RxFromString("978571348.E+222382547").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("6006.19370"))
	if err != nil {
		lang.SayString("xsub119")
	} else {
		if !(rexsult.ToString() == "9.78571348E+222382555") {
			lang.SayString("xsub119")
		}
	}
	rexsult, err = lang.RxFromString("14239029.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-36527.2221"))
	if err != nil {
		lang.SayString("xsub120")
	} else {
		if !(rexsult.ToString() == "14275556.2") {
			lang.SayString("xsub120")
		}
	}
	rexsult, err = lang.RxFromString("72333.2654E-544425548").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-690.664836E+662155120"))
	if err != nil {
		lang.SayString("xsub121")
	} else {
		if !(rexsult.ToString() == "6.90664836E+662155122") {
			lang.SayString("xsub121")
		}
	}
	rexsult, err = lang.RxFromString("-37721.1567E-115787341").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-828949864E-76251747"))
	if err != nil {
		lang.SayString("xsub122")
	} else {
		if !(rexsult.ToString() == "8.28949864E-76251739") {
			lang.SayString("xsub122")
		}
	}
	rexsult, err = lang.RxFromString("-2078852.83E-647080089").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-119779858.E+734665461"))
	if err != nil {
		lang.SayString("xsub123")
	} else {
		if !(rexsult.ToString() == "1.19779858E+734665469") {
			lang.SayString("xsub123")
		}
	}
	rexsult, err = lang.RxFromString("-79145.3625").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7718.57307"))
	if err != nil {
		lang.SayString("xsub124")
	} else {
		if !(rexsult.ToString() == "-71426.7894") {
			lang.SayString("xsub124")
		}
	}
	rexsult, err = lang.RxFromString("2103890.49E+959247237").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("20024.3017"))
	if err != nil {
		lang.SayString("xsub125")
	} else {
		if !(rexsult.ToString() == "2.10389049E+959247243") {
			lang.SayString("xsub125")
		}
	}
	rexsult, err = lang.RxFromString("911249557").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("79810804.9"))
	if err != nil {
		lang.SayString("xsub126")
	} else {
		if !(rexsult.ToString() == "831438752") {
			lang.SayString("xsub126")
		}
	}
	rexsult, err = lang.RxFromString("341134.994").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3.37486292"))
	if err != nil {
		lang.SayString("xsub127")
	} else {
		if !(rexsult.ToString() == "341131.619") {
			lang.SayString("xsub127")
		}
	}
	rexsult, err = lang.RxFromString("244.23634").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("512706190E-341459836"))
	if err != nil {
		lang.SayString("xsub128")
	} else {
		if !(rexsult.ToString() == "244.236340") {
			lang.SayString("xsub128")
		}
	}
	rexsult, err = lang.RxFromString("-9.22783849E+171585954").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-99.0946800"))
	if err != nil {
		lang.SayString("xsub129")
	} else {
		if !(rexsult.ToString() == "-9.22783849E+171585954") {
			lang.SayString("xsub129")
		}
	}
	rexsult, err = lang.RxFromString("699631.893").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-226.423958"))
	if err != nil {
		lang.SayString("xsub130")
	} else {
		if !(rexsult.ToString() == "699858.317") {
			lang.SayString("xsub130")
		}
	}
	rexsult, err = lang.RxFromString("-249350139.E-571793673").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("775732428."))
	if err != nil {
		lang.SayString("xsub131")
	} else {
		if !(rexsult.ToString() == "-775732428") {
			lang.SayString("xsub131")
		}
	}
	rexsult, err = lang.RxFromString("5.11629020").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-480.53194"))
	if err != nil {
		lang.SayString("xsub132")
	} else {
		if !(rexsult.ToString() == "485.648230") {
			lang.SayString("xsub132")
		}
	}
	rexsult, err = lang.RxFromString("-8.23352673E-446723147").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-530710.866"))
	if err != nil {
		lang.SayString("xsub133")
	} else {
		if !(rexsult.ToString() == "530710.866") {
			lang.SayString("xsub133")
		}
	}
	rexsult, err = lang.RxFromString("7.0598608").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-95908.35"))
	if err != nil {
		lang.SayString("xsub134")
	} else {
		if !(rexsult.ToString() == "95915.4099") {
			lang.SayString("xsub134")
		}
	}
	rexsult, err = lang.RxFromString("-7.91189845E+207202706").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1532.71847E+509944335"))
	if err != nil {
		lang.SayString("xsub135")
	} else {
		if !(rexsult.ToString() == "-1.53271847E+509944338") {
			lang.SayString("xsub135")
		}
	}
	rexsult, err = lang.RxFromString("208839370.E-215147432").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-75.9420559"))
	if err != nil {
		lang.SayString("xsub136")
	} else {
		if !(rexsult.ToString() == "75.9420559") {
			lang.SayString("xsub136")
		}
	}
	rexsult, err = lang.RxFromString("427.754244E-353328369").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("4705.0796"))
	if err != nil {
		lang.SayString("xsub137")
	} else {
		if !(rexsult.ToString() == "-4705.07960") {
			lang.SayString("xsub137")
		}
	}
	rexsult, err = lang.RxFromString("44911.089").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-95.1733605E-313081848"))
	if err != nil {
		lang.SayString("xsub138")
	} else {
		if !(rexsult.ToString() == "44911.0890") {
			lang.SayString("xsub138")
		}
	}
	rexsult, err = lang.RxFromString("452371821.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4109709.19"))
	if err != nil {
		lang.SayString("xsub139")
	} else {
		if !(rexsult.ToString() == "456481530") {
			lang.SayString("xsub139")
		}
	}
	rexsult, err = lang.RxFromString("94007.4392").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-9467725.5E+681898234"))
	if err != nil {
		lang.SayString("xsub140")
	} else {
		if !(rexsult.ToString() == "9.46772550E+681898240") {
			lang.SayString("xsub140")
		}
	}
	rexsult, err = lang.RxFromString("99147554.0E-751410586").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("38313.6423"))
	if err != nil {
		lang.SayString("xsub141")
	} else {
		if !(rexsult.ToString() == "-38313.6423") {
			lang.SayString("xsub141")
		}
	}
	rexsult, err = lang.RxFromString("-7919.30254").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-669.607854"))
	if err != nil {
		lang.SayString("xsub142")
	} else {
		if !(rexsult.ToString() == "-7249.69469") {
			lang.SayString("xsub142")
		}
	}
	rexsult, err = lang.RxFromString("461.58280E+136110821").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("710666052.E-383754231"))
	if err != nil {
		lang.SayString("xsub143")
	} else {
		if !(rexsult.ToString() == "4.61582800E+136110823") {
			lang.SayString("xsub143")
		}
	}
	rexsult, err = lang.RxFromString("3455755.47E-112465506").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("771.674306"))
	if err != nil {
		lang.SayString("xsub144")
	} else {
		if !(rexsult.ToString() == "-771.674306") {
			lang.SayString("xsub144")
		}
	}
	rexsult, err = lang.RxFromString("-477067757.E-961684940").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7.70122608E-741072245"))
	if err != nil {
		lang.SayString("xsub145")
	} else {
		if !(rexsult.ToString() == "-7.70122608E-741072245") {
			lang.SayString("xsub145")
		}
	}
	rexsult, err = lang.RxFromString("76482.352").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8237806.8"))
	if err != nil {
		lang.SayString("xsub146")
	} else {
		if !(rexsult.ToString() == "-8161324.45") {
			lang.SayString("xsub146")
		}
	}
	rexsult, err = lang.RxFromString("1.21505164E-565556504").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("9.26146573"))
	if err != nil {
		lang.SayString("xsub147")
	} else {
		if !(rexsult.ToString() == "-9.26146573") {
			lang.SayString("xsub147")
		}
	}
	rexsult, err = lang.RxFromString("-8303060.25E-169894883").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("901561.985"))
	if err != nil {
		lang.SayString("xsub148")
	} else {
		if !(rexsult.ToString() == "-901561.985") {
			lang.SayString("xsub148")
		}
	}
	rexsult, err = lang.RxFromString("-592464.92").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("71445510.7"))
	if err != nil {
		lang.SayString("xsub149")
	} else {
		if !(rexsult.ToString() == "-72037975.6") {
			lang.SayString("xsub149")
		}
	}
	rexsult, err = lang.RxFromString("-73774.4165").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-39.8243027"))
	if err != nil {
		lang.SayString("xsub150")
	} else {
		if !(rexsult.ToString() == "-73734.5922") {
			lang.SayString("xsub150")
		}
	}
	rexsult, err = lang.RxFromString("-524724715.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-55763.7937"))
	if err != nil {
		lang.SayString("xsub151")
	} else {
		if !(rexsult.ToString() == "-524668951") {
			lang.SayString("xsub151")
		}
	}
	rexsult, err = lang.RxFromString("7.53800427").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("784873768E-9981146"))
	if err != nil {
		lang.SayString("xsub152")
	} else {
		if !(rexsult.ToString() == "7.53800427") {
			lang.SayString("xsub152")
		}
	}
	rexsult, err = lang.RxFromString("37.6027452").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7.22454233"))
	if err != nil {
		lang.SayString("xsub153")
	} else {
		if !(rexsult.ToString() == "30.3782029") {
			lang.SayString("xsub153")
		}
	}
	rexsult, err = lang.RxFromString("2447660.39").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-36981.4253"))
	if err != nil {
		lang.SayString("xsub154")
	} else {
		if !(rexsult.ToString() == "2484641.82") {
			lang.SayString("xsub154")
		}
	}
	rexsult, err = lang.RxFromString("2160.36419").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1418.33574E+656265382"))
	if err != nil {
		lang.SayString("xsub155")
	} else {
		if !(rexsult.ToString() == "-1.41833574E+656265385") {
			lang.SayString("xsub155")
		}
	}
	rexsult, err = lang.RxFromString("8926.44939").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("54.9430027"))
	if err != nil {
		lang.SayString("xsub156")
	} else {
		if !(rexsult.ToString() == "8871.50639") {
			lang.SayString("xsub156")
		}
	}
	rexsult, err = lang.RxFromString("861588029").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-41657398E+77955925"))
	if err != nil {
		lang.SayString("xsub157")
	} else {
		if !(rexsult.ToString() == "4.16573980E+77955932") {
			lang.SayString("xsub157")
		}
	}
	rexsult, err = lang.RxFromString("-34.5253062").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("52.6722019"))
	if err != nil {
		lang.SayString("xsub158")
	} else {
		if !(rexsult.ToString() == "-87.1975081") {
			lang.SayString("xsub158")
		}
	}
	rexsult, err = lang.RxFromString("-18861647.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("99794586.7"))
	if err != nil {
		lang.SayString("xsub159")
	} else {
		if !(rexsult.ToString() == "-118656234") {
			lang.SayString("xsub159")
		}
	}
	rexsult, err = lang.RxFromString("322192.407").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("461.67044"))
	if err != nil {
		lang.SayString("xsub160")
	} else {
		if !(rexsult.ToString() == "321730.737") {
			lang.SayString("xsub160")
		}
	}
	rexsult, err = lang.RxFromString("-896298518E+61412314").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("78873.8049"))
	if err != nil {
		lang.SayString("xsub161")
	} else {
		if !(rexsult.ToString() == "-8.96298518E+61412322") {
			lang.SayString("xsub161")
		}
	}
	rexsult, err = lang.RxFromString("293.773732").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("479899052E+789950177"))
	if err != nil {
		lang.SayString("xsub162")
	} else {
		if !(rexsult.ToString() == "-4.79899052E+789950185") {
			lang.SayString("xsub162")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("51897955.3"))
	if err != nil {
		lang.SayString("xsub163")
	} else {
		if !(rexsult.ToString() == "-155417317") {
			lang.SayString("xsub163")
		}
	}
	rexsult, err = lang.RxFromString("37380.7802").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-277719788."))
	if err != nil {
		lang.SayString("xsub164")
	} else {
		if !(rexsult.ToString() == "277757169") {
			lang.SayString("xsub164")
		}
	}
	rexsult, err = lang.RxFromString("320133844.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-977517477"))
	if err != nil {
		lang.SayString("xsub165")
	} else {
		if !(rexsult.ToString() == "1.29765132E+9") {
			lang.SayString("xsub165")
		}
	}
	rexsult, err = lang.RxFromString("721776701E+933646161").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5689279.64E+669903645"))
	if err != nil {
		lang.SayString("xsub166")
	} else {
		if !(rexsult.ToString() == "7.21776701E+933646169") {
			lang.SayString("xsub166")
		}
	}
	rexsult, err = lang.RxFromString("-5409.00482").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2.16149386"))
	if err != nil {
		lang.SayString("xsub167")
	} else {
		if !(rexsult.ToString() == "-5406.84333") {
			lang.SayString("xsub167")
		}
	}
	rexsult, err = lang.RxFromString("-957960.367").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("322.858170"))
	if err != nil {
		lang.SayString("xsub168")
	} else {
		if !(rexsult.ToString() == "-958283.225") {
			lang.SayString("xsub168")
		}
	}
	rexsult, err = lang.RxFromString("-11817.8754E+613893442").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3.84735082E+888333249"))
	if err != nil {
		lang.SayString("xsub169")
	} else {
		if !(rexsult.ToString() == "3.84735082E+888333249") {
			lang.SayString("xsub169")
		}
	}
	rexsult, err = lang.RxFromString("840258203").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("58363.980E-906584723"))
	if err != nil {
		lang.SayString("xsub170")
	} else {
		if !(rexsult.ToString() == "840258203") {
			lang.SayString("xsub170")
		}
	}
	rexsult, err = lang.RxFromString("-205842096.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-191342.721"))
	if err != nil {
		lang.SayString("xsub171")
	} else {
		if !(rexsult.ToString() == "-205650753") {
			lang.SayString("xsub171")
		}
	}
	rexsult, err = lang.RxFromString("42501124.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("884.938498E+123341480"))
	if err != nil {
		lang.SayString("xsub172")
	} else {
		if !(rexsult.ToString() == "-8.84938498E+123341482") {
			lang.SayString("xsub172")
		}
	}
	rexsult, err = lang.RxFromString("-57809452.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-620380746"))
	if err != nil {
		lang.SayString("xsub173")
	} else {
		if !(rexsult.ToString() == "562571294") {
			lang.SayString("xsub173")
		}
	}
	rexsult, err = lang.RxFromString("-8022370.31").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("9858581.6"))
	if err != nil {
		lang.SayString("xsub174")
	} else {
		if !(rexsult.ToString() == "-17880951.9") {
			lang.SayString("xsub174")
		}
	}
	rexsult, err = lang.RxFromString("2.49065060E+592139141").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5432.72014E-419965357"))
	if err != nil {
		lang.SayString("xsub175")
	} else {
		if !(rexsult.ToString() == "2.49065060E+592139141") {
			lang.SayString("xsub175")
		}
	}
	rexsult, err = lang.RxFromString("-697273715E-242824870").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3.81757506"))
	if err != nil {
		lang.SayString("xsub176")
	} else {
		if !(rexsult.ToString() == "3.81757506") {
			lang.SayString("xsub176")
		}
	}
	rexsult, err = lang.RxFromString("-7.42204403E-315716280").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-8156111.67E+283261636"))
	if err != nil {
		lang.SayString("xsub177")
	} else {
		if !(rexsult.ToString() == "8.15611167E+283261642") {
			lang.SayString("xsub177")
		}
	}
	rexsult, err = lang.RxFromString("738063892").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("89900467.8"))
	if err != nil {
		lang.SayString("xsub178")
	} else {
		if !(rexsult.ToString() == "648163424") {
			lang.SayString("xsub178")
		}
	}
	rexsult, err = lang.RxFromString("-630309366").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-884783.338E-21595410"))
	if err != nil {
		lang.SayString("xsub179")
	} else {
		if !(rexsult.ToString() == "-630309366") {
			lang.SayString("xsub179")
		}
	}
	rexsult, err = lang.RxFromString("613.207774").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3007.78608"))
	if err != nil {
		lang.SayString("xsub180")
	} else {
		if !(rexsult.ToString() == "3620.99385") {
			lang.SayString("xsub180")
		}
	}
	rexsult, err = lang.RxFromString("-93006222.3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3.08964619"))
	if err != nil {
		lang.SayString("xsub181")
	} else {
		if !(rexsult.ToString() == "-93006219.2") {
			lang.SayString("xsub181")
		}
	}
	rexsult, err = lang.RxFromString("-18116.0621").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("34096.306E-270347092"))
	if err != nil {
		lang.SayString("xsub182")
	} else {
		if !(rexsult.ToString() == "-18116.0621") {
			lang.SayString("xsub182")
		}
	}
	rexsult, err = lang.RxFromString("19272386.9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-410442379."))
	if err != nil {
		lang.SayString("xsub183")
	} else {
		if !(rexsult.ToString() == "429714766") {
			lang.SayString("xsub183")
		}
	}
	rexsult, err = lang.RxFromString("4180.30821").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1.6439543E-624759104"))
	if err != nil {
		lang.SayString("xsub184")
	} else {
		if !(rexsult.ToString() == "4180.30821") {
			lang.SayString("xsub184")
		}
	}
	rexsult, err = lang.RxFromString("571.536725").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("389.899220"))
	if err != nil {
		lang.SayString("xsub185")
	} else {
		if !(rexsult.ToString() == "181.637505") {
			lang.SayString("xsub185")
		}
	}
	rexsult, err = lang.RxFromString("-622007306.E+159924886").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-126.971745"))
	if err != nil {
		lang.SayString("xsub186")
	} else {
		if !(rexsult.ToString() == "-6.22007306E+159924894") {
			lang.SayString("xsub186")
		}
	}
	rexsult, err = lang.RxFromString("-29.356551E-282816139").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("37141748E-903397821"))
	if err != nil {
		lang.SayString("xsub187")
	} else {
		if !(rexsult.ToString() == "-2.93565510E-282816138") {
			lang.SayString("xsub187")
		}
	}
	rexsult, err = lang.RxFromString("92427442.4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("674334898."))
	if err != nil {
		lang.SayString("xsub188")
	} else {
		if !(rexsult.ToString() == "-581907456") {
			lang.SayString("xsub188")
		}
	}
	rexsult, err = lang.RxFromString("44651895.7").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-910508.438"))
	if err != nil {
		lang.SayString("xsub189")
	} else {
		if !(rexsult.ToString() == "45562404.1") {
			lang.SayString("xsub189")
		}
	}
	rexsult, err = lang.RxFromString("647897872.E+374021790").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-467.423029"))
	if err != nil {
		lang.SayString("xsub190")
	} else {
		if !(rexsult.ToString() == "6.47897872E+374021798") {
			lang.SayString("xsub190")
		}
	}
	rexsult, err = lang.RxFromString("25.2592149").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("59.0436981"))
	if err != nil {
		lang.SayString("xsub191")
	} else {
		if !(rexsult.ToString() == "-33.7844832") {
			lang.SayString("xsub191")
		}
	}
	rexsult, err = lang.RxFromString("-6.850835").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1273.48240"))
	if err != nil {
		lang.SayString("xsub192")
	} else {
		if !(rexsult.ToString() == "1266.63157") {
			lang.SayString("xsub192")
		}
	}
	rexsult, err = lang.RxFromString("174.272325").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("5638.16229"))
	if err != nil {
		lang.SayString("xsub193")
	} else {
		if !(rexsult.ToString() == "-5463.88997") {
			lang.SayString("xsub193")
		}
	}
	rexsult, err = lang.RxFromString("3455629.76").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-8.27332322"))
	if err != nil {
		lang.SayString("xsub194")
	} else {
		if !(rexsult.ToString() == "3455638.03") {
			lang.SayString("xsub194")
		}
	}
	rexsult, err = lang.RxFromString("-924337723E-640771235").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("86639377.1"))
	if err != nil {
		lang.SayString("xsub195")
	} else {
		if !(rexsult.ToString() == "-86639377.1") {
			lang.SayString("xsub195")
		}
	}
	rexsult, err = lang.RxFromString("-620236932.E+656823969").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3364722.73"))
	if err != nil {
		lang.SayString("xsub196")
	} else {
		if !(rexsult.ToString() == "-6.20236932E+656823977") {
			lang.SayString("xsub196")
		}
	}
	rexsult, err = lang.RxFromString("9.10025079").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("702777882E-8192234"))
	if err != nil {
		lang.SayString("xsub197")
	} else {
		if !(rexsult.ToString() == "9.10025079") {
			lang.SayString("xsub197")
		}
	}
	rexsult, err = lang.RxFromString("-18857539.9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("813013129."))
	if err != nil {
		lang.SayString("xsub198")
	} else {
		if !(rexsult.ToString() == "-831870669") {
			lang.SayString("xsub198")
		}
	}
	rexsult, err = lang.RxFromString("-8.29530327").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3243419.57E+35688332"))
	if err != nil {
		lang.SayString("xsub199")
	} else {
		if !(rexsult.ToString() == "-3.24341957E+35688338") {
			lang.SayString("xsub199")
		}
	}
	rexsult, err = lang.RxFromString("-57101683.5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("763551341E+991491712"))
	if err != nil {
		lang.SayString("xsub200")
	} else {
		if !(rexsult.ToString() == "-7.63551341E+991491720") {
			lang.SayString("xsub200")
		}
	}
	rexsult, err = lang.RxFromString("-603326.740").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1710.95183"))
	if err != nil {
		lang.SayString("xsub201")
	} else {
		if !(rexsult.ToString() == "-605037.692") {
			lang.SayString("xsub201")
		}
	}
	rexsult, err = lang.RxFromString("-48142763.3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-943434114"))
	if err != nil {
		lang.SayString("xsub202")
	} else {
		if !(rexsult.ToString() == "895291351") {
			lang.SayString("xsub202")
		}
	}
	rexsult, err = lang.RxFromString("-204.586767").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-235.531847"))
	if err != nil {
		lang.SayString("xsub203")
	} else {
		if !(rexsult.ToString() == "30.945080") {
			lang.SayString("xsub203")
		}
	}
	rexsult, err = lang.RxFromString("-70.3805581").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("830137.913"))
	if err != nil {
		lang.SayString("xsub204")
	} else {
		if !(rexsult.ToString() == "-830208.294") {
			lang.SayString("xsub204")
		}
	}
	rexsult, err = lang.RxFromString("-8818.47606").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-60766.4571"))
	if err != nil {
		lang.SayString("xsub205")
	} else {
		if !(rexsult.ToString() == "51947.9810") {
			lang.SayString("xsub205")
		}
	}
	rexsult, err = lang.RxFromString("37060929.3E-168439509").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-79576717.1"))
	if err != nil {
		lang.SayString("xsub206")
	} else {
		if !(rexsult.ToString() == "79576717.1") {
			lang.SayString("xsub206")
		}
	}
	rexsult, err = lang.RxFromString("-656285310.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-107221462."))
	if err != nil {
		lang.SayString("xsub207")
	} else {
		if !(rexsult.ToString() == "-549063848") {
			lang.SayString("xsub207")
		}
	}
	rexsult, err = lang.RxFromString("653397.125").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7195.30990"))
	if err != nil {
		lang.SayString("xsub208")
	} else {
		if !(rexsult.ToString() == "646201.815") {
			lang.SayString("xsub208")
		}
	}
	rexsult, err = lang.RxFromString("56221910.0E+857909374").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-58.7247929"))
	if err != nil {
		lang.SayString("xsub209")
	} else {
		if !(rexsult.ToString() == "5.62219100E+857909381") {
			lang.SayString("xsub209")
		}
	}
	rexsult, err = lang.RxFromString("809862859E+643769974").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5.06784016"))
	if err != nil {
		lang.SayString("xsub210")
	} else {
		if !(rexsult.ToString() == "8.09862859E+643769982") {
			lang.SayString("xsub210")
		}
	}
	rexsult, err = lang.RxFromString("-62011.4563E-117563240").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-57.1731586E+115657204"))
	if err != nil {
		lang.SayString("xsub211")
	} else {
		if !(rexsult.ToString() == "5.71731586E+115657205") {
			lang.SayString("xsub211")
		}
	}
	rexsult, err = lang.RxFromString("315.33351").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("91588.837E-536020149"))
	if err != nil {
		lang.SayString("xsub212")
	} else {
		if !(rexsult.ToString() == "315.333510") {
			lang.SayString("xsub212")
		}
	}
	rexsult, err = lang.RxFromString("739.944710").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("202949.175"))
	if err != nil {
		lang.SayString("xsub213")
	} else {
		if !(rexsult.ToString() == "-202209.230") {
			lang.SayString("xsub213")
		}
	}
	rexsult, err = lang.RxFromString("87686.8016").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("4204890.40"))
	if err != nil {
		lang.SayString("xsub214")
	} else {
		if !(rexsult.ToString() == "-4117203.60") {
			lang.SayString("xsub214")
		}
	}
	rexsult, err = lang.RxFromString("987126721.E-725794834").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("4874166.23"))
	if err != nil {
		lang.SayString("xsub215")
	} else {
		if !(rexsult.ToString() == "-4874166.23") {
			lang.SayString("xsub215")
		}
	}
	rexsult, err = lang.RxFromString("728148726.E-661695938").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("32798.5202"))
	if err != nil {
		lang.SayString("xsub216")
	} else {
		if !(rexsult.ToString() == "-32798.5202") {
			lang.SayString("xsub216")
		}
	}
	rexsult, err = lang.RxFromString("7428219.97").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("667.326760"))
	if err != nil {
		lang.SayString("xsub217")
	} else {
		if !(rexsult.ToString() == "7427552.64") {
			lang.SayString("xsub217")
		}
	}
	rexsult, err = lang.RxFromString("-7291.19212").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("209.64966E-588526476"))
	if err != nil {
		lang.SayString("xsub218")
	} else {
		if !(rexsult.ToString() == "-7291.19212") {
			lang.SayString("xsub218")
		}
	}
	rexsult, err = lang.RxFromString("-358.24550").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4447.78675E+601402509"))
	if err != nil {
		lang.SayString("xsub219")
	} else {
		if !(rexsult.ToString() == "4.44778675E+601402512") {
			lang.SayString("xsub219")
		}
	}
	rexsult, err = lang.RxFromString("118.621826").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2.72010038"))
	if err != nil {
		lang.SayString("xsub220")
	} else {
		if !(rexsult.ToString() == "121.341926") {
			lang.SayString("xsub220")
		}
	}
	rexsult, err = lang.RxFromString("8071961.94").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-135533740.E-102451543"))
	if err != nil {
		lang.SayString("xsub221")
	} else {
		if !(rexsult.ToString() == "8071961.94") {
			lang.SayString("xsub221")
		}
	}
	rexsult, err = lang.RxFromString("64262528.5E+812118682").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-8692.94447E-732186947"))
	if err != nil {
		lang.SayString("xsub222")
	} else {
		if !(rexsult.ToString() == "6.42625285E+812118689") {
			lang.SayString("xsub222")
		}
	}
	rexsult, err = lang.RxFromString("-35544.4029").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-567830.130"))
	if err != nil {
		lang.SayString("xsub223")
	} else {
		if !(rexsult.ToString() == "532285.727") {
			lang.SayString("xsub223")
		}
	}
	rexsult, err = lang.RxFromString("-7.16513047E+59297103").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("87767.8211"))
	if err != nil {
		lang.SayString("xsub224")
	} else {
		if !(rexsult.ToString() == "-7.16513047E+59297103") {
			lang.SayString("xsub224")
		}
	}
	rexsult, err = lang.RxFromString("-509.483395").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-147242915."))
	if err != nil {
		lang.SayString("xsub225")
	} else {
		if !(rexsult.ToString() == "147242406") {
			lang.SayString("xsub225")
		}
	}
	rexsult, err = lang.RxFromString("-7919047.28E+956041629").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-367667329"))
	if err != nil {
		lang.SayString("xsub226")
	} else {
		if !(rexsult.ToString() == "-7.91904728E+956041635") {
			lang.SayString("xsub226")
		}
	}
	rexsult, err = lang.RxFromString("895612630.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-36.4104040"))
	if err != nil {
		lang.SayString("xsub227")
	} else {
		if !(rexsult.ToString() == "895612666") {
			lang.SayString("xsub227")
		}
	}
	rexsult, err = lang.RxFromString("25455.4973").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2955.00006E+528196218"))
	if err != nil {
		lang.SayString("xsub228")
	} else {
		if !(rexsult.ToString() == "-2.95500006E+528196221") {
			lang.SayString("xsub228")
		}
	}
	rexsult, err = lang.RxFromString("-112.294144E+273414172").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-71448007.7"))
	if err != nil {
		lang.SayString("xsub229")
	} else {
		if !(rexsult.ToString() == "-1.12294144E+273414174") {
			lang.SayString("xsub229")
		}
	}
	rexsult, err = lang.RxFromString("62871.2202").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2484.0382E+211662557"))
	if err != nil {
		lang.SayString("xsub230")
	} else {
		if !(rexsult.ToString() == "-2.48403820E+211662560") {
			lang.SayString("xsub230")
		}
	}
	rexsult, err = lang.RxFromString("71.9281575").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-9810012.5"))
	if err != nil {
		lang.SayString("xsub231")
	} else {
		if !(rexsult.ToString() == "9810084.43") {
			lang.SayString("xsub231")
		}
	}
	rexsult, err = lang.RxFromString("-6388022.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-88.042967"))
	if err != nil {
		lang.SayString("xsub232")
	} else {
		if !(rexsult.ToString() == "-6387933.96") {
			lang.SayString("xsub232")
		}
	}
	rexsult, err = lang.RxFromString("372567445.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("96.0992141"))
	if err != nil {
		lang.SayString("xsub233")
	} else {
		if !(rexsult.ToString() == "372567349") {
			lang.SayString("xsub233")
		}
	}
	rexsult, err = lang.RxFromString("802.156517").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-174409310.E-255338020"))
	if err != nil {
		lang.SayString("xsub234")
	} else {
		if !(rexsult.ToString() == "802.156517") {
			lang.SayString("xsub234")
		}
	}
	rexsult, err = lang.RxFromString("-3.65207541").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("74501982.0"))
	if err != nil {
		lang.SayString("xsub235")
	} else {
		if !(rexsult.ToString() == "-74501985.7") {
			lang.SayString("xsub235")
		}
	}
	rexsult, err = lang.RxFromString("-5297.76981").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-859.719404"))
	if err != nil {
		lang.SayString("xsub236")
	} else {
		if !(rexsult.ToString() == "-4438.05041") {
			lang.SayString("xsub236")
		}
	}
	rexsult, err = lang.RxFromString("-684172.592").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("766.448597E+288361959"))
	if err != nil {
		lang.SayString("xsub237")
	} else {
		if !(rexsult.ToString() == "-7.66448597E+288361961") {
			lang.SayString("xsub237")
		}
	}
	rexsult, err = lang.RxFromString("626919.219").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("57469.8727E+13188610"))
	if err != nil {
		lang.SayString("xsub238")
	} else {
		if !(rexsult.ToString() == "-5.74698727E+13188614") {
			lang.SayString("xsub238")
		}
	}
	rexsult, err = lang.RxFromString("-77480.5840").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("893265.594E+287982552"))
	if err != nil {
		lang.SayString("xsub239")
	} else {
		if !(rexsult.ToString() == "-8.93265594E+287982557") {
			lang.SayString("xsub239")
		}
	}
	rexsult, err = lang.RxFromString("-7177620.29").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7786343.83"))
	if err != nil {
		lang.SayString("xsub240")
	} else {
		if !(rexsult.ToString() == "-14963964.1") {
			lang.SayString("xsub240")
		}
	}
	rexsult, err = lang.RxFromString("9.6224130").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("4.50355112"))
	if err != nil {
		lang.SayString("xsub241")
	} else {
		if !(rexsult.ToString() == "5.11886188") {
			lang.SayString("xsub241")
		}
	}
	rexsult, err = lang.RxFromString("-66.6337347E-597410086").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-818812885"))
	if err != nil {
		lang.SayString("xsub242")
	} else {
		if !(rexsult.ToString() == "818812885") {
			lang.SayString("xsub242")
		}
	}
	rexsult, err = lang.RxFromString("65587553.7").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("600574.736"))
	if err != nil {
		lang.SayString("xsub243")
	} else {
		if !(rexsult.ToString() == "64986979.0") {
			lang.SayString("xsub243")
		}
	}
	rexsult, err = lang.RxFromString("-32401.939").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-585200217."))
	if err != nil {
		lang.SayString("xsub244")
	} else {
		if !(rexsult.ToString() == "585167815") {
			lang.SayString("xsub244")
		}
	}
	rexsult, err = lang.RxFromString("69573.988").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-9.77003465E+740933668"))
	if err != nil {
		lang.SayString("xsub245")
	} else {
		if !(rexsult.ToString() == "9.77003465E+740933668") {
			lang.SayString("xsub245")
		}
	}
	rexsult, err = lang.RxFromString("2362.06251").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-433149546.E-152643629"))
	if err != nil {
		lang.SayString("xsub246")
	} else {
		if !(rexsult.ToString() == "2362.06251") {
			lang.SayString("xsub246")
		}
	}
	rexsult, err = lang.RxFromString("-615.23488E+249953452").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-21437483.7"))
	if err != nil {
		lang.SayString("xsub247")
	} else {
		if !(rexsult.ToString() == "-6.15234880E+249953454") {
			lang.SayString("xsub247")
		}
	}
	rexsult, err = lang.RxFromString("216741082.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("250290244"))
	if err != nil {
		lang.SayString("xsub248")
	} else {
		if !(rexsult.ToString() == "-33549162") {
			lang.SayString("xsub248")
		}
	}
	rexsult, err = lang.RxFromString("-6364720.49").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("5539245.64"))
	if err != nil {
		lang.SayString("xsub249")
	} else {
		if !(rexsult.ToString() == "-11903966.1") {
			lang.SayString("xsub249")
		}
	}
	rexsult, err = lang.RxFromString("-814599.475").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-14.5431191"))
	if err != nil {
		lang.SayString("xsub250")
	} else {
		if !(rexsult.ToString() == "-814584.932") {
			lang.SayString("xsub250")
		}
	}
	rexsult, err = lang.RxFromString("-877498.755").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("507408724E-168628106"))
	if err != nil {
		lang.SayString("xsub251")
	} else {
		if !(rexsult.ToString() == "-877498.755") {
			lang.SayString("xsub251")
		}
	}
	rexsult, err = lang.RxFromString("10634446.5E+475783861").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("50.7213056E+17807809"))
	if err != nil {
		lang.SayString("xsub252")
	} else {
		if !(rexsult.ToString() == "1.06344465E+475783868") {
			lang.SayString("xsub252")
		}
	}
	rexsult, err = lang.RxFromString("-162726.257E-597285918").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4391.54799"))
	if err != nil {
		lang.SayString("xsub253")
	} else {
		if !(rexsult.ToString() == "4391.54799") {
			lang.SayString("xsub253")
		}
	}
	rexsult, err = lang.RxFromString("700354586.E-99856707").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7198.0493E+436250299"))
	if err != nil {
		lang.SayString("xsub254")
	} else {
		if !(rexsult.ToString() == "-7.19804930E+436250302") {
			lang.SayString("xsub254")
		}
	}
	rexsult, err = lang.RxFromString("39617663E-463704664").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-895.290346"))
	if err != nil {
		lang.SayString("xsub255")
	} else {
		if !(rexsult.ToString() == "895.290346") {
			lang.SayString("xsub255")
		}
	}
	rexsult, err = lang.RxFromString("5350882.59").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-36329829"))
	if err != nil {
		lang.SayString("xsub256")
	} else {
		if !(rexsult.ToString() == "41680711.6") {
			lang.SayString("xsub256")
		}
	}
	rexsult, err = lang.RxFromString("91966.4084E+210382952").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("166740.46E-42001390"))
	if err != nil {
		lang.SayString("xsub257")
	} else {
		if !(rexsult.ToString() == "9.19664084E+210382956") {
			lang.SayString("xsub257")
		}
	}
	rexsult, err = lang.RxFromString("231899031.E-481759076").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("726.337100"))
	if err != nil {
		lang.SayString("xsub258")
	} else {
		if !(rexsult.ToString() == "-726.337100") {
			lang.SayString("xsub258")
		}
	}
	rexsult, err = lang.RxFromString("-9611312.33").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("22109735.9"))
	if err != nil {
		lang.SayString("xsub259")
	} else {
		if !(rexsult.ToString() == "-31721048.2") {
			lang.SayString("xsub259")
		}
	}
	rexsult, err = lang.RxFromString("-5604938.15E-36812542").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("735937577."))
	if err != nil {
		lang.SayString("xsub260")
	} else {
		if !(rexsult.ToString() == "-735937577") {
			lang.SayString("xsub260")
		}
	}
	rexsult, err = lang.RxFromString("693881413.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("260547224E-480281418"))
	if err != nil {
		lang.SayString("xsub261")
	} else {
		if !(rexsult.ToString() == "693881413") {
			lang.SayString("xsub261")
		}
	}
	rexsult, err = lang.RxFromString("-34865.7378E-368768024").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2297117.88"))
	if err != nil {
		lang.SayString("xsub262")
	} else {
		if !(rexsult.ToString() == "-2297117.88") {
			lang.SayString("xsub262")
		}
	}
	rexsult, err = lang.RxFromString("1123.32456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7.86747918E+930888796"))
	if err != nil {
		lang.SayString("xsub263")
	} else {
		if !(rexsult.ToString() == "-7.86747918E+930888796") {
			lang.SayString("xsub263")
		}
	}
	rexsult, err = lang.RxFromString("56.6607465E+467812565").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("909552512E+764516200"))
	if err != nil {
		lang.SayString("xsub264")
	} else {
		if !(rexsult.ToString() == "-9.09552512E+764516208") {
			lang.SayString("xsub264")
		}
	}
	rexsult, err = lang.RxFromString("-1.85771840E+365552540").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-73028339.7"))
	if err != nil {
		lang.SayString("xsub265")
	} else {
		if !(rexsult.ToString() == "-1.85771840E+365552540") {
			lang.SayString("xsub265")
		}
	}
	rexsult, err = lang.RxFromString("34.1935525").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-40767.6450"))
	if err != nil {
		lang.SayString("xsub266")
	} else {
		if !(rexsult.ToString() == "40801.8386") {
			lang.SayString("xsub266")
		}
	}
	rexsult, err = lang.RxFromString("26.0009168E+751618294").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-304019.929"))
	if err != nil {
		lang.SayString("xsub267")
	} else {
		if !(rexsult.ToString() == "2.60009168E+751618295") {
			lang.SayString("xsub267")
		}
	}
	rexsult, err = lang.RxFromString("-58.4853072E+588540055").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4647.3205"))
	if err != nil {
		lang.SayString("xsub268")
	} else {
		if !(rexsult.ToString() == "-5.84853072E+588540056") {
			lang.SayString("xsub268")
		}
	}
	rexsult, err = lang.RxFromString("51.025101").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4467691.57"))
	if err != nil {
		lang.SayString("xsub269")
	} else {
		if !(rexsult.ToString() == "4467742.60") {
			lang.SayString("xsub269")
		}
	}
	rexsult, err = lang.RxFromString("-2214.76582").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("379785372E+223117572"))
	if err != nil {
		lang.SayString("xsub270")
	} else {
		if !(rexsult.ToString() == "-3.79785372E+223117580") {
			lang.SayString("xsub270")
		}
	}
	rexsult, err = lang.RxFromString("-2564.75207E-841443929").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-653498187"))
	if err != nil {
		lang.SayString("xsub271")
	} else {
		if !(rexsult.ToString() == "653498187") {
			lang.SayString("xsub271")
		}
	}
	rexsult, err = lang.RxFromString("513115529.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("27775075.6E+217133352"))
	if err != nil {
		lang.SayString("xsub272")
	} else {
		if !(rexsult.ToString() == "-2.77750756E+217133359") {
			lang.SayString("xsub272")
		}
	}
	rexsult, err = lang.RxFromString("-247157.208").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-532990.453"))
	if err != nil {
		lang.SayString("xsub273")
	} else {
		if !(rexsult.ToString() == "285833.245") {
			lang.SayString("xsub273")
		}
	}
	rexsult, err = lang.RxFromString("40.2490764E-339482253").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7626.85442E+594264540"))
	if err != nil {
		lang.SayString("xsub274")
	} else {
		if !(rexsult.ToString() == "-7.62685442E+594264543") {
			lang.SayString("xsub274")
		}
	}
	rexsult, err = lang.RxFromString("-1156008.8").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-8870382.36"))
	if err != nil {
		lang.SayString("xsub275")
	} else {
		if !(rexsult.ToString() == "7714373.56") {
			lang.SayString("xsub275")
		}
	}
	rexsult, err = lang.RxFromString("880097928.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-52455011.1E+204538218"))
	if err != nil {
		lang.SayString("xsub276")
	} else {
		if !(rexsult.ToString() == "5.24550111E+204538225") {
			lang.SayString("xsub276")
		}
	}
	rexsult, err = lang.RxFromString("5796.2524").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("34458329.7E+832129426"))
	if err != nil {
		lang.SayString("xsub277")
	} else {
		if !(rexsult.ToString() == "-3.44583297E+832129433") {
			lang.SayString("xsub277")
		}
	}
	rexsult, err = lang.RxFromString("27.1000923E-218032223").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-45.0198341"))
	if err != nil {
		lang.SayString("xsub278")
	} else {
		if !(rexsult.ToString() == "45.0198341") {
			lang.SayString("xsub278")
		}
	}
	rexsult, err = lang.RxFromString("42643477.8").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("26118465E-730390549"))
	if err != nil {
		lang.SayString("xsub279")
	} else {
		if !(rexsult.ToString() == "42643477.8") {
			lang.SayString("xsub279")
		}
	}
	rexsult, err = lang.RxFromString("-31918.9176E-163031657").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-21.5422824E-807317258"))
	if err != nil {
		lang.SayString("xsub280")
	} else {
		if !(rexsult.ToString() == "-3.19189176E-163031653") {
			lang.SayString("xsub280")
		}
	}
	rexsult, err = lang.RxFromString("84224841.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2.62548255E+647087608"))
	if err != nil {
		lang.SayString("xsub281")
	} else {
		if !(rexsult.ToString() == "-2.62548255E+647087608") {
			lang.SayString("xsub281")
		}
	}
	rexsult, err = lang.RxFromString("-64413698.9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-6674.1055E-701047852"))
	if err != nil {
		lang.SayString("xsub282")
	} else {
		if !(rexsult.ToString() == "-64413698.9") {
			lang.SayString("xsub282")
		}
	}
	rexsult, err = lang.RxFromString("-62.5059208").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("9.5795779E-898350012"))
	if err != nil {
		lang.SayString("xsub283")
	} else {
		if !(rexsult.ToString() == "-62.5059208") {
			lang.SayString("xsub283")
		}
	}
	rexsult, err = lang.RxFromString("9090950.80").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("436.400932"))
	if err != nil {
		lang.SayString("xsub284")
	} else {
		if !(rexsult.ToString() == "9090514.40") {
			lang.SayString("xsub284")
		}
	}
	rexsult, err = lang.RxFromString("-89833825.7E+329205393").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-779430.194"))
	if err != nil {
		lang.SayString("xsub285")
	} else {
		if !(rexsult.ToString() == "-8.98338257E+329205400") {
			lang.SayString("xsub285")
		}
	}
	rexsult, err = lang.RxFromString("-714562.019E+750205688").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("704079764"))
	if err != nil {
		lang.SayString("xsub286")
	} else {
		if !(rexsult.ToString() == "-7.14562019E+750205693") {
			lang.SayString("xsub286")
		}
	}
	rexsult, err = lang.RxFromString("-584537670.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("31139.7737E-146687560"))
	if err != nil {
		lang.SayString("xsub287")
	} else {
		if !(rexsult.ToString() == "-584537670") {
			lang.SayString("xsub287")
		}
	}
	rexsult, err = lang.RxFromString("-4.18074650E-858746879").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("571035.277E-279409165"))
	if err != nil {
		lang.SayString("xsub288")
	} else {
		if !(rexsult.ToString() == "-5.71035277E-279409160") {
			lang.SayString("xsub288")
		}
	}
	rexsult, err = lang.RxFromString("5.15309635").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-695649.219E+451948183"))
	if err != nil {
		lang.SayString("xsub289")
	} else {
		if !(rexsult.ToString() == "6.95649219E+451948188") {
			lang.SayString("xsub289")
		}
	}
	rexsult, err = lang.RxFromString("-940030153.E+83797657").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4.11510193"))
	if err != nil {
		lang.SayString("xsub290")
	} else {
		if !(rexsult.ToString() == "-9.40030153E+83797665") {
			lang.SayString("xsub290")
		}
	}
	rexsult, err = lang.RxFromString("89088.9683E+587739290").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.31932110"))
	if err != nil {
		lang.SayString("xsub291")
	} else {
		if !(rexsult.ToString() == "8.90889683E+587739294") {
			lang.SayString("xsub291")
		}
	}
	rexsult, err = lang.RxFromString("3336750").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("6.47961126"))
	if err != nil {
		lang.SayString("xsub292")
	} else {
		if !(rexsult.ToString() == "3336743.52") {
			lang.SayString("xsub292")
		}
	}
	rexsult, err = lang.RxFromString("904654622.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("692065270.E+329081915"))
	if err != nil {
		lang.SayString("xsub293")
	} else {
		if !(rexsult.ToString() == "-6.92065270E+329081923") {
			lang.SayString("xsub293")
		}
	}
	rexsult, err = lang.RxFromString("304804380").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4681.23698"))
	if err != nil {
		lang.SayString("xsub294")
	} else {
		if !(rexsult.ToString() == "304809061") {
			lang.SayString("xsub294")
		}
	}
	rexsult, err = lang.RxFromString("674.55569").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-82981.2684E+852890752"))
	if err != nil {
		lang.SayString("xsub295")
	} else {
		if !(rexsult.ToString() == "8.29812684E+852890756") {
			lang.SayString("xsub295")
		}
	}
	rexsult, err = lang.RxFromString("-5111.51025E-108006096").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("5448870.4E+279212255"))
	if err != nil {
		lang.SayString("xsub296")
	} else {
		if !(rexsult.ToString() == "-5.44887040E+279212261") {
			lang.SayString("xsub296")
		}
	}
	rexsult, err = lang.RxFromString("-2623.45068").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-466463938."))
	if err != nil {
		lang.SayString("xsub297")
	} else {
		if !(rexsult.ToString() == "466461315") {
			lang.SayString("xsub297")
		}
	}
	rexsult, err = lang.RxFromString("299350.435").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3373.33551"))
	if err != nil {
		lang.SayString("xsub298")
	} else {
		if !(rexsult.ToString() == "295977.100") {
			lang.SayString("xsub298")
		}
	}
	rexsult, err = lang.RxFromString("-6589947.80").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2448.75933E-591549734"))
	if err != nil {
		lang.SayString("xsub299")
	} else {
		if !(rexsult.ToString() == "-6589947.80") {
			lang.SayString("xsub299")
		}
	}
	rexsult, err = lang.RxFromString("3774.5358E-491090520").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("173.060090"))
	if err != nil {
		lang.SayString("xsub300")
	} else {
		if !(rexsult.ToString() == "-173.060090") {
			lang.SayString("xsub300")
		}
	}
	rexsult, err = lang.RxFromString("-13.6783690").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-453.610117"))
	if err != nil {
		lang.SayString("xsub301")
	} else {
		if !(rexsult.ToString() == "439.931748") {
			lang.SayString("xsub301")
		}
	}
	rexsult, err = lang.RxFromString("-990100927.E-615244634").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("223801.421E+247632618"))
	if err != nil {
		lang.SayString("xsub302")
	} else {
		if !(rexsult.ToString() == "-2.23801421E+247632623") {
			lang.SayString("xsub302")
		}
	}
	rexsult, err = lang.RxFromString("1275.10292").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-667965353"))
	if err != nil {
		lang.SayString("xsub303")
	} else {
		if !(rexsult.ToString() == "667966628") {
			lang.SayString("xsub303")
		}
	}
	rexsult, err = lang.RxFromString("-8.76375480E-596792197").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("992.077361"))
	if err != nil {
		lang.SayString("xsub304")
	} else {
		if !(rexsult.ToString() == "-992.077361") {
			lang.SayString("xsub304")
		}
	}
	rexsult, err = lang.RxFromString("953.976935E+385444720").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("96503.3378"))
	if err != nil {
		lang.SayString("xsub305")
	} else {
		if !(rexsult.ToString() == "9.53976935E+385444722") {
			lang.SayString("xsub305")
		}
	}
	rexsult, err = lang.RxFromString("213577152").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-986710073E+31900046"))
	if err != nil {
		lang.SayString("xsub306")
	} else {
		if !(rexsult.ToString() == "9.86710073E+31900054") {
			lang.SayString("xsub306")
		}
	}
	rexsult, err = lang.RxFromString("91393.9398E-323439228").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-135.701000"))
	if err != nil {
		lang.SayString("xsub307")
	} else {
		if !(rexsult.ToString() == "135.701000") {
			lang.SayString("xsub307")
		}
	}
	rexsult, err = lang.RxFromString("-396.503557").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("45757264.E-254363788"))
	if err != nil {
		lang.SayString("xsub308")
	} else {
		if !(rexsult.ToString() == "-396.503557") {
			lang.SayString("xsub308")
		}
	}
	rexsult, err = lang.RxFromString("59807846.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.53345254"))
	if err != nil {
		lang.SayString("xsub309")
	} else {
		if !(rexsult.ToString() == "59807844.6") {
			lang.SayString("xsub309")
		}
	}
	rexsult, err = lang.RxFromString("-8046158.45").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8.3635397"))
	if err != nil {
		lang.SayString("xsub310")
	} else {
		if !(rexsult.ToString() == "-8046166.81") {
			lang.SayString("xsub310")
		}
	}
	rexsult, err = lang.RxFromString("55.1123381E+50627250").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-94.0355047E-162540316"))
	if err != nil {
		lang.SayString("xsub311")
	} else {
		if !(rexsult.ToString() == "5.51123381E+50627251") {
			lang.SayString("xsub311")
		}
	}
	rexsult, err = lang.RxFromString("-948.038054").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3580.84510"))
	if err != nil {
		lang.SayString("xsub312")
	} else {
		if !(rexsult.ToString() == "-4528.88315") {
			lang.SayString("xsub312")
		}
	}
	rexsult, err = lang.RxFromString("-6026.42752").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-14.2286406E-334921364"))
	if err != nil {
		lang.SayString("xsub313")
	} else {
		if !(rexsult.ToString() == "-6026.42752") {
			lang.SayString("xsub313")
		}
	}
	rexsult, err = lang.RxFromString("79551.5014").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-538.186229"))
	if err != nil {
		lang.SayString("xsub314")
	} else {
		if !(rexsult.ToString() == "80089.6876") {
			lang.SayString("xsub314")
		}
	}
	rexsult, err = lang.RxFromString("42706056.E+623578292").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-690.327745"))
	if err != nil {
		lang.SayString("xsub315")
	} else {
		if !(rexsult.ToString() == "4.27060560E+623578299") {
			lang.SayString("xsub315")
		}
	}
	rexsult, err = lang.RxFromString("2454136.08E+502374077").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("856268.795E-356664934"))
	if err != nil {
		lang.SayString("xsub316")
	} else {
		if !(rexsult.ToString() == "2.45413608E+502374083") {
			lang.SayString("xsub316")
		}
	}
	rexsult, err = lang.RxFromString("-3264204.54").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-42704.501"))
	if err != nil {
		lang.SayString("xsub317")
	} else {
		if !(rexsult.ToString() == "-3221500.04") {
			lang.SayString("xsub317")
		}
	}
	rexsult, err = lang.RxFromString("1.21265492").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("44102.6073"))
	if err != nil {
		lang.SayString("xsub318")
	} else {
		if !(rexsult.ToString() == "-44101.3947") {
			lang.SayString("xsub318")
		}
	}
	rexsult, err = lang.RxFromString("-19.054711E+975514652").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-22144.0822"))
	if err != nil {
		lang.SayString("xsub319")
	} else {
		if !(rexsult.ToString() == "-1.90547110E+975514653") {
			lang.SayString("xsub319")
		}
	}
	rexsult, err = lang.RxFromString("745.78452").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1922.00670E+375923302"))
	if err != nil {
		lang.SayString("xsub320")
	} else {
		if !(rexsult.ToString() == "1.92200670E+375923305") {
			lang.SayString("xsub320")
		}
	}
	rexsult, err = lang.RxFromString("-963717836").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-823989308"))
	if err != nil {
		lang.SayString("xsub321")
	} else {
		if !(rexsult.ToString() == "-139728528") {
			lang.SayString("xsub321")
		}
	}
	rexsult, err = lang.RxFromString("82.4185291E-321919303").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-215747737.E-995147400"))
	if err != nil {
		lang.SayString("xsub322")
	} else {
		if !(rexsult.ToString() == "8.24185291E-321919302") {
			lang.SayString("xsub322")
		}
	}
	rexsult, err = lang.RxFromString("-808328.607E-790810342").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("53075.7082"))
	if err != nil {
		lang.SayString("xsub323")
	} else {
		if !(rexsult.ToString() == "-53075.7082") {
			lang.SayString("xsub323")
		}
	}
	rexsult, err = lang.RxFromString("700592.720").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-698485.085"))
	if err != nil {
		lang.SayString("xsub324")
	} else {
		if !(rexsult.ToString() == "1399077.81") {
			lang.SayString("xsub324")
		}
	}
	rexsult, err = lang.RxFromString("-80273928.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("661346.239"))
	if err != nil {
		lang.SayString("xsub325")
	} else {
		if !(rexsult.ToString() == "-80935274.2") {
			lang.SayString("xsub325")
		}
	}
	rexsult, err = lang.RxFromString("-24018251.0E+819786764").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("59141.9600E-167165065"))
	if err != nil {
		lang.SayString("xsub326")
	} else {
		if !(rexsult.ToString() == "-2.40182510E+819786771") {
			lang.SayString("xsub326")
		}
	}
	rexsult, err = lang.RxFromString("2512953.3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3769170.35E-993621645"))
	if err != nil {
		lang.SayString("xsub327")
	} else {
		if !(rexsult.ToString() == "2512953.30") {
			lang.SayString("xsub327")
		}
	}
	rexsult, err = lang.RxFromString("-682.796370").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("71131.0224"))
	if err != nil {
		lang.SayString("xsub328")
	} else {
		if !(rexsult.ToString() == "-71813.8188") {
			lang.SayString("xsub328")
		}
	}
	rexsult, err = lang.RxFromString("89.9997490").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4993.69831"))
	if err != nil {
		lang.SayString("xsub329")
	} else {
		if !(rexsult.ToString() == "5083.69806") {
			lang.SayString("xsub329")
		}
	}
	rexsult, err = lang.RxFromString("76563354.6E-112338836").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("278271.585E-511481095"))
	if err != nil {
		lang.SayString("xsub330")
	} else {
		if !(rexsult.ToString() == "7.65633546E-112338829") {
			lang.SayString("xsub330")
		}
	}
	rexsult, err = lang.RxFromString("-932499.010").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("873.377701E-502190452"))
	if err != nil {
		lang.SayString("xsub331")
	} else {
		if !(rexsult.ToString() == "-932499.010") {
			lang.SayString("xsub331")
		}
	}
	rexsult, err = lang.RxFromString("-7735918.21E+799514797").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7748.78023"))
	if err != nil {
		lang.SayString("xsub332")
	} else {
		if !(rexsult.ToString() == "-7.73591821E+799514803") {
			lang.SayString("xsub332")
		}
	}
	rexsult, err = lang.RxFromString("-3708780.75E+445232787").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("980.006567E-780728623"))
	if err != nil {
		lang.SayString("xsub333")
	} else {
		if !(rexsult.ToString() == "-3.70878075E+445232793") {
			lang.SayString("xsub333")
		}
	}
	rexsult, err = lang.RxFromString("-5205124.44E-140588661").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-495394029.E-620856313"))
	if err != nil {
		lang.SayString("xsub334")
	} else {
		if !(rexsult.ToString() == "-5.20512444E-140588655") {
			lang.SayString("xsub334")
		}
	}
	rexsult, err = lang.RxFromString("-8868.72074").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("5592399.93"))
	if err != nil {
		lang.SayString("xsub335")
	} else {
		if !(rexsult.ToString() == "-5601268.65") {
			lang.SayString("xsub335")
		}
	}
	rexsult, err = lang.RxFromString("-74.7852037E-175205809").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("4.14316542"))
	if err != nil {
		lang.SayString("xsub336")
	} else {
		if !(rexsult.ToString() == "-4.14316542") {
			lang.SayString("xsub336")
		}
	}
	rexsult, err = lang.RxFromString("84196.1091E+242628748").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8.07523036E-288231467"))
	if err != nil {
		lang.SayString("xsub337")
	} else {
		if !(rexsult.ToString() == "8.41961091E+242628752") {
			lang.SayString("xsub337")
		}
	}
	rexsult, err = lang.RxFromString("38660103.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-6671.73085E+900998477"))
	if err != nil {
		lang.SayString("xsub338")
	} else {
		if !(rexsult.ToString() == "6.67173085E+900998480") {
			lang.SayString("xsub338")
		}
	}
	rexsult, err = lang.RxFromString("-52.2659460").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-296404199E+372050476"))
	if err != nil {
		lang.SayString("xsub339")
	} else {
		if !(rexsult.ToString() == "2.96404199E+372050484") {
			lang.SayString("xsub339")
		}
	}
	rexsult, err = lang.RxFromString("6.06625013").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-276.359186"))
	if err != nil {
		lang.SayString("xsub340")
	} else {
		if !(rexsult.ToString() == "282.425436") {
			lang.SayString("xsub340")
		}
	}
	rexsult, err = lang.RxFromString("-62971617.5E-241444744").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("46266799.3"))
	if err != nil {
		lang.SayString("xsub341")
	} else {
		if !(rexsult.ToString() == "-46266799.3") {
			lang.SayString("xsub341")
		}
	}
	rexsult, err = lang.RxFromString("-5.36917800").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-311124593.E-976066491"))
	if err != nil {
		lang.SayString("xsub342")
	} else {
		if !(rexsult.ToString() == "-5.36917800") {
			lang.SayString("xsub342")
		}
	}
	rexsult, err = lang.RxFromString("2467915.01").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-92.5558322"))
	if err != nil {
		lang.SayString("xsub343")
	} else {
		if !(rexsult.ToString() == "2468007.57") {
			lang.SayString("xsub343")
		}
	}
	rexsult, err = lang.RxFromString("187.232671").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-840.469347"))
	if err != nil {
		lang.SayString("xsub344")
	} else {
		if !(rexsult.ToString() == "1027.70202") {
			lang.SayString("xsub344")
		}
	}
	rexsult, err = lang.RxFromString("81233.6823").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5192.21666E+309315093"))
	if err != nil {
		lang.SayString("xsub345")
	} else {
		if !(rexsult.ToString() == "5.19221666E+309315096") {
			lang.SayString("xsub345")
		}
	}
	rexsult, err = lang.RxFromString("-854.586113").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-79.8715762E-853065103"))
	if err != nil {
		lang.SayString("xsub346")
	} else {
		if !(rexsult.ToString() == "-854.586113") {
			lang.SayString("xsub346")
		}
	}
	rexsult, err = lang.RxFromString("78872665.3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("172.102119"))
	if err != nil {
		lang.SayString("xsub347")
	} else {
		if !(rexsult.ToString() == "78872493.2") {
			lang.SayString("xsub347")
		}
	}
	rexsult, err = lang.RxFromString("328268.1E-436315617").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-204.522245"))
	if err != nil {
		lang.SayString("xsub348")
	} else {
		if !(rexsult.ToString() == "204.522245") {
			lang.SayString("xsub348")
		}
	}
	rexsult, err = lang.RxFromString("-4037911.02E+641367645").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("29.5713010"))
	if err != nil {
		lang.SayString("xsub349")
	} else {
		if !(rexsult.ToString() == "-4.03791102E+641367651") {
			lang.SayString("xsub349")
		}
	}
	rexsult, err = lang.RxFromString("-688755561.E-95301699").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("978.275312E+913812609"))
	if err != nil {
		lang.SayString("xsub350")
	} else {
		if !(rexsult.ToString() == "-9.78275312E+913812611") {
			lang.SayString("xsub350")
		}
	}
	rexsult, err = lang.RxFromString("-5.47345502").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("59818.7580"))
	if err != nil {
		lang.SayString("xsub351")
	} else {
		if !(rexsult.ToString() == "-59824.2315") {
			lang.SayString("xsub351")
		}
	}
	rexsult, err = lang.RxFromString("563891620E-361354567").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-845900362."))
	if err != nil {
		lang.SayString("xsub352")
	} else {
		if !(rexsult.ToString() == "845900362") {
			lang.SayString("xsub352")
		}
	}
	rexsult, err = lang.RxFromString("-69.7231286").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("85773.7504"))
	if err != nil {
		lang.SayString("xsub353")
	} else {
		if !(rexsult.ToString() == "-85843.4735") {
			lang.SayString("xsub353")
		}
	}
	rexsult, err = lang.RxFromString("5125.51188").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("73814638.4E-500934741"))
	if err != nil {
		lang.SayString("xsub354")
	} else {
		if !(rexsult.ToString() == "5125.51188") {
			lang.SayString("xsub354")
		}
	}
	rexsult, err = lang.RxFromString("-54.6254096").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-332921899."))
	if err != nil {
		lang.SayString("xsub355")
	} else {
		if !(rexsult.ToString() == "332921844") {
			lang.SayString("xsub355")
		}
	}
	rexsult, err = lang.RxFromString("-9.04778095E-591874079").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8719.40286"))
	if err != nil {
		lang.SayString("xsub356")
	} else {
		if !(rexsult.ToString() == "-8719.40286") {
			lang.SayString("xsub356")
		}
	}
	rexsult, err = lang.RxFromString("-21006.1733E+884684431").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-48872.9175"))
	if err != nil {
		lang.SayString("xsub357")
	} else {
		if !(rexsult.ToString() == "-2.10061733E+884684435") {
			lang.SayString("xsub357")
		}
	}
	rexsult, err = lang.RxFromString("-1546783").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-51935370.4"))
	if err != nil {
		lang.SayString("xsub358")
	} else {
		if !(rexsult.ToString() == "50388587.4") {
			lang.SayString("xsub358")
		}
	}
	rexsult, err = lang.RxFromString("61302486.8").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("205.490417"))
	if err != nil {
		lang.SayString("xsub359")
	} else {
		if !(rexsult.ToString() == "61302281.3") {
			lang.SayString("xsub359")
		}
	}
	rexsult, err = lang.RxFromString("-318180109.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-54008744.6E-170931002"))
	if err != nil {
		lang.SayString("xsub360")
	} else {
		if !(rexsult.ToString() == "-318180109") {
			lang.SayString("xsub360")
		}
	}
	rexsult, err = lang.RxFromString("-28486137.1E+901441714").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-42454.940"))
	if err != nil {
		lang.SayString("xsub361")
	} else {
		if !(rexsult.ToString() == "-2.84861371E+901441721") {
			lang.SayString("xsub361")
		}
	}
	rexsult, err = lang.RxFromString("-546398328.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-27.9149712"))
	if err != nil {
		lang.SayString("xsub362")
	} else {
		if !(rexsult.ToString() == "-546398300") {
			lang.SayString("xsub362")
		}
	}
	rexsult, err = lang.RxFromString("5402066.1E-284978216").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("622.751128"))
	if err != nil {
		lang.SayString("xsub363")
	} else {
		if !(rexsult.ToString() == "-622.751128") {
			lang.SayString("xsub363")
		}
	}
	rexsult, err = lang.RxFromString("18845620").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3129.43753"))
	if err != nil {
		lang.SayString("xsub364")
	} else {
		if !(rexsult.ToString() == "18842490.6") {
			lang.SayString("xsub364")
		}
	}
	rexsult, err = lang.RxFromString("50707.1412E+912475670").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-198098.186E+701407524"))
	if err != nil {
		lang.SayString("xsub365")
	} else {
		if !(rexsult.ToString() == "5.07071412E+912475674") {
			lang.SayString("xsub365")
		}
	}
	rexsult, err = lang.RxFromString("55.8245006E+928885991").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("99170843.9E-47402167"))
	if err != nil {
		lang.SayString("xsub366")
	} else {
		if !(rexsult.ToString() == "5.58245006E+928885992") {
			lang.SayString("xsub366")
		}
	}
	rexsult, err = lang.RxFromString("13.8003883E-386224921").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-84126481.9E-296378341"))
	if err != nil {
		lang.SayString("xsub367")
	} else {
		if !(rexsult.ToString() == "8.41264819E-296378334") {
			lang.SayString("xsub367")
		}
	}
	rexsult, err = lang.RxFromString("9820.90457").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("46671.5915"))
	if err != nil {
		lang.SayString("xsub368")
	} else {
		if !(rexsult.ToString() == "-36850.6869") {
			lang.SayString("xsub368")
		}
	}
	rexsult, err = lang.RxFromString("7.22436006E+831949153").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-11168830E+322331045"))
	if err != nil {
		lang.SayString("xsub369")
	} else {
		if !(rexsult.ToString() == "7.22436006E+831949153") {
			lang.SayString("xsub369")
		}
	}
	rexsult, err = lang.RxFromString("472648900").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-207.784153"))
	if err != nil {
		lang.SayString("xsub370")
	} else {
		if !(rexsult.ToString() == "472649108") {
			lang.SayString("xsub370")
		}
	}
	rexsult, err = lang.RxFromString("-8754.49306").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-818.165153E+631475457"))
	if err != nil {
		lang.SayString("xsub371")
	} else {
		if !(rexsult.ToString() == "8.18165153E+631475459") {
			lang.SayString("xsub371")
		}
	}
	rexsult, err = lang.RxFromString("98750864").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("191380.551"))
	if err != nil {
		lang.SayString("xsub372")
	} else {
		if !(rexsult.ToString() == "98559483.5") {
			lang.SayString("xsub372")
		}
	}
	rexsult, err = lang.RxFromString("725292561.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-768963606.E+340762986"))
	if err != nil {
		lang.SayString("xsub373")
	} else {
		if !(rexsult.ToString() == "7.68963606E+340762994") {
			lang.SayString("xsub373")
		}
	}
	rexsult, err = lang.RxFromString("1862.80445").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("648254483."))
	if err != nil {
		lang.SayString("xsub374")
	} else {
		if !(rexsult.ToString() == "-648252620") {
			lang.SayString("xsub374")
		}
	}
	rexsult, err = lang.RxFromString("-5549320.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-93580684.1"))
	if err != nil {
		lang.SayString("xsub375")
	} else {
		if !(rexsult.ToString() == "88031364.0") {
			lang.SayString("xsub375")
		}
	}
	rexsult, err = lang.RxFromString("-14677053.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-25784.7358"))
	if err != nil {
		lang.SayString("xsub376")
	} else {
		if !(rexsult.ToString() == "-14651268.4") {
			lang.SayString("xsub376")
		}
	}
	rexsult, err = lang.RxFromString("547402.308E+571687617").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7835797.01E+500067364"))
	if err != nil {
		lang.SayString("xsub377")
	} else {
		if !(rexsult.ToString() == "5.47402308E+571687622") {
			lang.SayString("xsub377")
		}
	}
	rexsult, err = lang.RxFromString("-4131738.09").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7579.07566"))
	if err != nil {
		lang.SayString("xsub378")
	} else {
		if !(rexsult.ToString() == "-4139317.17") {
			lang.SayString("xsub378")
		}
	}
	rexsult, err = lang.RxFromString("504544.648").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7678.96133E-662143268"))
	if err != nil {
		lang.SayString("xsub379")
	} else {
		if !(rexsult.ToString() == "504544.648") {
			lang.SayString("xsub379")
		}
	}
	rexsult, err = lang.RxFromString("829898241").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8912.99114E+929228149"))
	if err != nil {
		lang.SayString("xsub380")
	} else {
		if !(rexsult.ToString() == "-8.91299114E+929228152") {
			lang.SayString("xsub380")
		}
	}
	rexsult, err = lang.RxFromString("53.6891691").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-11.2371140"))
	if err != nil {
		lang.SayString("xsub381")
	} else {
		if !(rexsult.ToString() == "64.9262831") {
			lang.SayString("xsub381")
		}
	}
	rexsult, err = lang.RxFromString("-93951823.4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-25317.8645"))
	if err != nil {
		lang.SayString("xsub382")
	} else {
		if !(rexsult.ToString() == "-93926505.5") {
			lang.SayString("xsub382")
		}
	}
	rexsult, err = lang.RxFromString("446919.123").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("951338490."))
	if err != nil {
		lang.SayString("xsub383")
	} else {
		if !(rexsult.ToString() == "-950891571") {
			lang.SayString("xsub383")
		}
	}
	rexsult, err = lang.RxFromString("-8.01787748").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-88.3076852"))
	if err != nil {
		lang.SayString("xsub384")
	} else {
		if !(rexsult.ToString() == "80.2898077") {
			lang.SayString("xsub384")
		}
	}
	rexsult, err = lang.RxFromString("517458139").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-999731.548"))
	if err != nil {
		lang.SayString("xsub385")
	} else {
		if !(rexsult.ToString() == "518457871") {
			lang.SayString("xsub385")
		}
	}
	rexsult, err = lang.RxFromString("-405543440").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4013.18295"))
	if err != nil {
		lang.SayString("xsub386")
	} else {
		if !(rexsult.ToString() == "-405539427") {
			lang.SayString("xsub386")
		}
	}
	rexsult, err = lang.RxFromString("-49245250.1E+682760825").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-848776.637"))
	if err != nil {
		lang.SayString("xsub387")
	} else {
		if !(rexsult.ToString() == "-4.92452501E+682760832") {
			lang.SayString("xsub387")
		}
	}
	rexsult, err = lang.RxFromString("-151144455").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-170371.29"))
	if err != nil {
		lang.SayString("xsub388")
	} else {
		if !(rexsult.ToString() == "-150974084") {
			lang.SayString("xsub388")
		}
	}
	rexsult, err = lang.RxFromString("-729236746.E+662737067").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("9.10823602"))
	if err != nil {
		lang.SayString("xsub389")
	} else {
		if !(rexsult.ToString() == "-7.29236746E+662737075") {
			lang.SayString("xsub389")
		}
	}
	rexsult, err = lang.RxFromString("534.394729").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2369839.37"))
	if err != nil {
		lang.SayString("xsub390")
	} else {
		if !(rexsult.ToString() == "2370373.76") {
			lang.SayString("xsub390")
		}
	}
	rexsult, err = lang.RxFromString("89100.1797").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("224.370309"))
	if err != nil {
		lang.SayString("xsub391")
	} else {
		if !(rexsult.ToString() == "88875.8094") {
			lang.SayString("xsub391")
		}
	}
	rexsult, err = lang.RxFromString("-821377.777").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("38.552821"))
	if err != nil {
		lang.SayString("xsub392")
	} else {
		if !(rexsult.ToString() == "-821416.330") {
			lang.SayString("xsub392")
		}
	}
	rexsult, err = lang.RxFromString("-392640.782").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2571619.5E+113237865"))
	if err != nil {
		lang.SayString("xsub393")
	} else {
		if !(rexsult.ToString() == "2.57161950E+113237871") {
			lang.SayString("xsub393")
		}
	}
	rexsult, err = lang.RxFromString("-651397.712").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-723.59673"))
	if err != nil {
		lang.SayString("xsub394")
	} else {
		if !(rexsult.ToString() == "-650674.115") {
			lang.SayString("xsub394")
		}
	}
	rexsult, err = lang.RxFromString("86.6890892").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("940380864"))
	if err != nil {
		lang.SayString("xsub395")
	} else {
		if !(rexsult.ToString() == "-940380777") {
			lang.SayString("xsub395")
		}
	}
	rexsult, err = lang.RxFromString("4880.06442E-382222621").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-115627239E-912834031"))
	if err != nil {
		lang.SayString("xsub396")
	} else {
		if !(rexsult.ToString() == "4.88006442E-382222618") {
			lang.SayString("xsub396")
		}
	}
	rexsult, err = lang.RxFromString("173398265E-532383158").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3462.51450E+80986915"))
	if err != nil {
		lang.SayString("xsub397")
	} else {
		if !(rexsult.ToString() == "-3.46251450E+80986918") {
			lang.SayString("xsub397")
		}
	}
	rexsult, err = lang.RxFromString("-1522176.78").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-6631061.77"))
	if err != nil {
		lang.SayString("xsub398")
	} else {
		if !(rexsult.ToString() == "5108884.99") {
			lang.SayString("xsub398")
		}
	}
	rexsult, err = lang.RxFromString("538.10453").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("522934310"))
	if err != nil {
		lang.SayString("xsub399")
	} else {
		if !(rexsult.ToString() == "-522933772") {
			lang.SayString("xsub399")
		}
	}
	rexsult, err = lang.RxFromString("880243.444E-750940977").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-354.601578E-204943740"))
	if err != nil {
		lang.SayString("xsub400")
	} else {
		if !(rexsult.ToString() == "3.54601578E-204943738") {
			lang.SayString("xsub400")
		}
	}
	rexsult, err = lang.RxFromString("968370.780").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("6677268.73"))
	if err != nil {
		lang.SayString("xsub401")
	} else {
		if !(rexsult.ToString() == "-5708897.95") {
			lang.SayString("xsub401")
		}
	}
	rexsult, err = lang.RxFromString("-97.7474945").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("31248241.5"))
	if err != nil {
		lang.SayString("xsub402")
	} else {
		if !(rexsult.ToString() == "-31248339.2") {
			lang.SayString("xsub402")
		}
	}
	rexsult, err = lang.RxFromString("-187582786.E+369916952").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("957840435E+744365567"))
	if err != nil {
		lang.SayString("xsub403")
	} else {
		if !(rexsult.ToString() == "-9.57840435E+744365575") {
			lang.SayString("xsub403")
		}
	}
	rexsult, err = lang.RxFromString("-328026144.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-125499735."))
	if err != nil {
		lang.SayString("xsub404")
	} else {
		if !(rexsult.ToString() == "-202526409") {
			lang.SayString("xsub404")
		}
	}
	rexsult, err = lang.RxFromString("-99424150.2E-523662102").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3712.35030"))
	if err != nil {
		lang.SayString("xsub405")
	} else {
		if !(rexsult.ToString() == "-3712.35030") {
			lang.SayString("xsub405")
		}
	}
	rexsult, err = lang.RxFromString("14838.0718").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("9489893.28E+830631266"))
	if err != nil {
		lang.SayString("xsub406")
	} else {
		if !(rexsult.ToString() == "-9.48989328E+830631272") {
			lang.SayString("xsub406")
		}
	}
	rexsult, err = lang.RxFromString("71207472.8").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-31835.0809"))
	if err != nil {
		lang.SayString("xsub407")
	} else {
		if !(rexsult.ToString() == "71239307.9") {
			lang.SayString("xsub407")
		}
	}
	rexsult, err = lang.RxFromString("-20440.4394").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-44.4064328E+511085806"))
	if err != nil {
		lang.SayString("xsub408")
	} else {
		if !(rexsult.ToString() == "4.44064328E+511085807") {
			lang.SayString("xsub408")
		}
	}
	rexsult, err = lang.RxFromString("-54.3684171E-807210192").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.04592973E-984041807"))
	if err != nil {
		lang.SayString("xsub409")
	} else {
		if !(rexsult.ToString() == "-5.43684171E-807210191") {
			lang.SayString("xsub409")
		}
	}
	rexsult, err = lang.RxFromString("54310060.5E+948159739").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("274320701.E+205880484"))
	if err != nil {
		lang.SayString("xsub410")
	} else {
		if !(rexsult.ToString() == "5.43100605E+948159746") {
			lang.SayString("xsub410")
		}
	}
	rexsult, err = lang.RxFromString("-657.186702").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("426844.39"))
	if err != nil {
		lang.SayString("xsub411")
	} else {
		if !(rexsult.ToString() == "-427501.577") {
			lang.SayString("xsub411")
		}
	}
	rexsult, err = lang.RxFromString("-41593077.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-688607.564"))
	if err != nil {
		lang.SayString("xsub412")
	} else {
		if !(rexsult.ToString() == "-40904469.4") {
			lang.SayString("xsub412")
		}
	}
	rexsult, err = lang.RxFromString("-5786.38132").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("190556652.E+177045877"))
	if err != nil {
		lang.SayString("xsub413")
	} else {
		if !(rexsult.ToString() == "-1.90556652E+177045885") {
			lang.SayString("xsub413")
		}
	}
	rexsult, err = lang.RxFromString("737622.974").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-241560693E+249506565"))
	if err != nil {
		lang.SayString("xsub414")
	} else {
		if !(rexsult.ToString() == "2.41560693E+249506573") {
			lang.SayString("xsub414")
		}
	}
	rexsult, err = lang.RxFromString("5615373.52").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7.27583808E-949781048"))
	if err != nil {
		lang.SayString("xsub415")
	} else {
		if !(rexsult.ToString() == "5615373.52") {
			lang.SayString("xsub415")
		}
	}
	rexsult, err = lang.RxFromString("644136.179").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-835708.103"))
	if err != nil {
		lang.SayString("xsub416")
	} else {
		if !(rexsult.ToString() == "1479844.28") {
			lang.SayString("xsub416")
		}
	}
	rexsult, err = lang.RxFromString("-307.419521E+466861843").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-738689976.E-199032711"))
	if err != nil {
		lang.SayString("xsub417")
	} else {
		if !(rexsult.ToString() == "-3.07419521E+466861845") {
			lang.SayString("xsub417")
		}
	}
	rexsult, err = lang.RxFromString("-619642.130").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-226740537.E-902590153"))
	if err != nil {
		lang.SayString("xsub418")
	} else {
		if !(rexsult.ToString() == "-619642.130") {
			lang.SayString("xsub418")
		}
	}
	rexsult, err = lang.RxFromString("-31068.7549").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3.41495042E+86001379"))
	if err != nil {
		lang.SayString("xsub419")
	} else {
		if !(rexsult.ToString() == "3.41495042E+86001379") {
			lang.SayString("xsub419")
		}
	}
	rexsult, err = lang.RxFromString("-68951173.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-211804977.E-97318126"))
	if err != nil {
		lang.SayString("xsub420")
	} else {
		if !(rexsult.ToString() == "-68951173.0") {
			lang.SayString("xsub420")
		}
	}
	rexsult, err = lang.RxFromString("-4.09492571E-301749490").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("434.20199E-749390952"))
	if err != nil {
		lang.SayString("xsub421")
	} else {
		if !(rexsult.ToString() == "-4.09492571E-301749490") {
			lang.SayString("xsub421")
		}
	}
	rexsult, err = lang.RxFromString("3898.03188").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-82572.615"))
	if err != nil {
		lang.SayString("xsub422")
	} else {
		if !(rexsult.ToString() == "86470.6469") {
			lang.SayString("xsub422")
		}
	}
	rexsult, err = lang.RxFromString("-1.7619356").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2546.64043"))
	if err != nil {
		lang.SayString("xsub423")
	} else {
		if !(rexsult.ToString() == "2544.87850") {
			lang.SayString("xsub423")
		}
	}
	rexsult, err = lang.RxFromString("59714.1968").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("29734388.6E-564525525"))
	if err != nil {
		lang.SayString("xsub424")
	} else {
		if !(rexsult.ToString() == "59714.1968") {
			lang.SayString("xsub424")
		}
	}
	rexsult, err = lang.RxFromString("6.88891136E-935467395").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-785049.562E-741671442"))
	if err != nil {
		lang.SayString("xsub425")
	} else {
		if !(rexsult.ToString() == "7.85049562E-741671437") {
			lang.SayString("xsub425")
		}
	}
	rexsult, err = lang.RxFromString("975566251").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-519.858530"))
	if err != nil {
		lang.SayString("xsub426")
	} else {
		if !(rexsult.ToString() == "975566771") {
			lang.SayString("xsub426")
		}
	}
	rexsult, err = lang.RxFromString("307401954").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-231481582."))
	if err != nil {
		lang.SayString("xsub427")
	} else {
		if !(rexsult.ToString() == "538883536") {
			lang.SayString("xsub427")
		}
	}
	rexsult, err = lang.RxFromString("2237645.48E+992947388").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-60618055.3E-857316706"))
	if err != nil {
		lang.SayString("xsub428")
	} else {
		if !(rexsult.ToString() == "2.23764548E+992947394") {
			lang.SayString("xsub428")
		}
	}
	rexsult, err = lang.RxFromString("-403903.851").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("35.5049687E-72095155"))
	if err != nil {
		lang.SayString("xsub429")
	} else {
		if !(rexsult.ToString() == "-403903.851") {
			lang.SayString("xsub429")
		}
	}
	rexsult, err = lang.RxFromString("6.48674979").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-621732.532E+422575800"))
	if err != nil {
		lang.SayString("xsub430")
	} else {
		if !(rexsult.ToString() == "6.21732532E+422575805") {
			lang.SayString("xsub430")
		}
	}
	rexsult, err = lang.RxFromString("-31401.9418").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("36.3960679"))
	if err != nil {
		lang.SayString("xsub431")
	} else {
		if !(rexsult.ToString() == "-31438.3379") {
			lang.SayString("xsub431")
		}
	}
	rexsult, err = lang.RxFromString("31345321.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("51.5482191"))
	if err != nil {
		lang.SayString("xsub432")
	} else {
		if !(rexsult.ToString() == "31345269.6") {
			lang.SayString("xsub432")
		}
	}
	rexsult, err = lang.RxFromString("-64.172844").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-28506227.2E-767965800"))
	if err != nil {
		lang.SayString("xsub433")
	} else {
		if !(rexsult.ToString() == "-64.1728440") {
			lang.SayString("xsub433")
		}
	}
	rexsult, err = lang.RxFromString("70437.1551").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-62916.1233"))
	if err != nil {
		lang.SayString("xsub434")
	} else {
		if !(rexsult.ToString() == "133353.278") {
			lang.SayString("xsub434")
		}
	}
	rexsult, err = lang.RxFromString("916260164").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-58.4017325"))
	if err != nil {
		lang.SayString("xsub435")
	} else {
		if !(rexsult.ToString() == "916260222") {
			lang.SayString("xsub435")
		}
	}
	rexsult, err = lang.RxFromString("19889085.3E-46816480").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1581683.94"))
	if err != nil {
		lang.SayString("xsub436")
	} else {
		if !(rexsult.ToString() == "-1581683.94") {
			lang.SayString("xsub436")
		}
	}
	rexsult, err = lang.RxFromString("-56312.3383").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("789.466064"))
	if err != nil {
		lang.SayString("xsub437")
	} else {
		if !(rexsult.ToString() == "-57101.8044") {
			lang.SayString("xsub437")
		}
	}
	rexsult, err = lang.RxFromString("183442.849").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-925876106"))
	if err != nil {
		lang.SayString("xsub438")
	} else {
		if !(rexsult.ToString() == "926059549") {
			lang.SayString("xsub438")
		}
	}
	rexsult, err = lang.RxFromString("971113.655E-695540249").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-419351120E-977743823"))
	if err != nil {
		lang.SayString("xsub439")
	} else {
		if !(rexsult.ToString() == "9.71113655E-695540244") {
			lang.SayString("xsub439")
		}
	}
	rexsult, err = lang.RxFromString("859658551.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("72338.2054"))
	if err != nil {
		lang.SayString("xsub440")
	} else {
		if !(rexsult.ToString() == "859586213") {
			lang.SayString("xsub440")
		}
	}
	rexsult, err = lang.RxFromString("-3.86446630E+426816068").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-664.534737"))
	if err != nil {
		lang.SayString("xsub441")
	} else {
		if !(rexsult.ToString() == "-3.86446630E+426816068") {
			lang.SayString("xsub441")
		}
	}
	rexsult, err = lang.RxFromString("-969.881818").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("31170.8555"))
	if err != nil {
		lang.SayString("xsub442")
	} else {
		if !(rexsult.ToString() == "-32140.7373") {
			lang.SayString("xsub442")
		}
	}
	rexsult, err = lang.RxFromString("7980537.27").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("85.4040512"))
	if err != nil {
		lang.SayString("xsub443")
	} else {
		if !(rexsult.ToString() == "7980451.87") {
			lang.SayString("xsub443")
		}
	}
	rexsult, err = lang.RxFromString("-114609916.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7525.14981"))
	if err != nil {
		lang.SayString("xsub444")
	} else {
		if !(rexsult.ToString() == "-114617441") {
			lang.SayString("xsub444")
		}
	}
	rexsult, err = lang.RxFromString("8.43404682E-500572568").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("474526719"))
	if err != nil {
		lang.SayString("xsub445")
	} else {
		if !(rexsult.ToString() == "-474526719") {
			lang.SayString("xsub445")
		}
	}
	rexsult, err = lang.RxFromString("188006433").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2260.17037E-978192525"))
	if err != nil {
		lang.SayString("xsub446")
	} else {
		if !(rexsult.ToString() == "188006433") {
			lang.SayString("xsub446")
		}
	}
	rexsult, err = lang.RxFromString("-9.95836312").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-866466703"))
	if err != nil {
		lang.SayString("xsub447")
	} else {
		if !(rexsult.ToString() == "866466693") {
			lang.SayString("xsub447")
		}
	}
	rexsult, err = lang.RxFromString("80919339.2E-967231586").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("219.824266"))
	if err != nil {
		lang.SayString("xsub448")
	} else {
		if !(rexsult.ToString() == "-219.824266") {
			lang.SayString("xsub448")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-89827.5229"))
	if err != nil {
		lang.SayString("xsub449")
	} else {
		if !(rexsult.ToString() == "249406.967") {
			lang.SayString("xsub449")
		}
	}
	rexsult, err = lang.RxFromString("-4.54000153").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("6966333.74"))
	if err != nil {
		lang.SayString("xsub450")
	} else {
		if !(rexsult.ToString() == "-6966338.28") {
			lang.SayString("xsub450")
		}
	}
	rexsult, err = lang.RxFromString("28701538.7E-391015649").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-920999192."))
	if err != nil {
		lang.SayString("xsub451")
	} else {
		if !(rexsult.ToString() == "920999192") {
			lang.SayString("xsub451")
		}
	}
	rexsult, err = lang.RxFromString("-361382575.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7976.15286E+898491169"))
	if err != nil {
		lang.SayString("xsub452")
	} else {
		if !(rexsult.ToString() == "7.97615286E+898491172") {
			lang.SayString("xsub452")
		}
	}
	rexsult, err = lang.RxFromString("7021805.61").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1222952.83"))
	if err != nil {
		lang.SayString("xsub453")
	} else {
		if !(rexsult.ToString() == "5798852.78") {
			lang.SayString("xsub453")
		}
	}
	rexsult, err = lang.RxFromString("-40.4811667").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-79655.5635"))
	if err != nil {
		lang.SayString("xsub454")
	} else {
		if !(rexsult.ToString() == "79615.0823") {
			lang.SayString("xsub454")
		}
	}
	rexsult, err = lang.RxFromString("-8755674.38E+117168177").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("148.903404"))
	if err != nil {
		lang.SayString("xsub455")
	} else {
		if !(rexsult.ToString() == "-8.75567438E+117168183") {
			lang.SayString("xsub455")
		}
	}
	rexsult, err = lang.RxFromString("34.5329781E+382829392").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-45.2177309"))
	if err != nil {
		lang.SayString("xsub456")
	} else {
		if !(rexsult.ToString() == "3.45329781E+382829393") {
			lang.SayString("xsub456")
		}
	}
	rexsult, err = lang.RxFromString("-37958476.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("584367.935"))
	if err != nil {
		lang.SayString("xsub457")
	} else {
		if !(rexsult.ToString() == "-38542843.9") {
			lang.SayString("xsub457")
		}
	}
	rexsult, err = lang.RxFromString("495233.553E-414152215").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("62352759.2"))
	if err != nil {
		lang.SayString("xsub458")
	} else {
		if !(rexsult.ToString() == "-62352759.2") {
			lang.SayString("xsub458")
		}
	}
	rexsult, err = lang.RxFromString("-502343060").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-96828.994"))
	if err != nil {
		lang.SayString("xsub459")
	} else {
		if !(rexsult.ToString() == "-502246231") {
			lang.SayString("xsub459")
		}
	}
	rexsult, err = lang.RxFromString("-22.439639E+916362878").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-39.4037681"))
	if err != nil {
		lang.SayString("xsub460")
	} else {
		if !(rexsult.ToString() == "-2.24396390E+916362879") {
			lang.SayString("xsub460")
		}
	}
	rexsult, err = lang.RxFromString("718180.587E-957473722").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.66223443"))
	if err != nil {
		lang.SayString("xsub461")
	} else {
		if !(rexsult.ToString() == "-1.66223443") {
			lang.SayString("xsub461")
		}
	}
	rexsult, err = lang.RxFromString("-51592.2698").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-713885.741"))
	if err != nil {
		lang.SayString("xsub462")
	} else {
		if !(rexsult.ToString() == "662293.471") {
			lang.SayString("xsub462")
		}
	}
	rexsult, err = lang.RxFromString("51.2279848E+80439745").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("207.55925E+865165070"))
	if err != nil {
		lang.SayString("xsub463")
	} else {
		if !(rexsult.ToString() == "-2.07559250E+865165072") {
			lang.SayString("xsub463")
		}
	}
	rexsult, err = lang.RxFromString("-5983.23468").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-39.9544513"))
	if err != nil {
		lang.SayString("xsub464")
	} else {
		if !(rexsult.ToString() == "-5943.28023") {
			lang.SayString("xsub464")
		}
	}
	rexsult, err = lang.RxFromString("921639332.E-917542963").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("287325.891"))
	if err != nil {
		lang.SayString("xsub465")
	} else {
		if !(rexsult.ToString() == "-287325.891") {
			lang.SayString("xsub465")
		}
	}
	rexsult, err = lang.RxFromString("91095916.8E-787312969").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-58643.418E+58189880"))
	if err != nil {
		lang.SayString("xsub466")
	} else {
		if !(rexsult.ToString() == "5.86434180E+58189884") {
			lang.SayString("xsub466")
		}
	}
	rexsult, err = lang.RxFromString("-6410.5555").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-234964259"))
	if err != nil {
		lang.SayString("xsub467")
	} else {
		if !(rexsult.ToString() == "234957849") {
			lang.SayString("xsub467")
		}
	}
	rexsult, err = lang.RxFromString("-5.32711606").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-8447286.21"))
	if err != nil {
		lang.SayString("xsub468")
	} else {
		if !(rexsult.ToString() == "8447280.88") {
			lang.SayString("xsub468")
		}
	}
	rexsult, err = lang.RxFromString("-82272171.8").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-776.238587E-372690416"))
	if err != nil {
		lang.SayString("xsub469")
	} else {
		if !(rexsult.ToString() == "-82272171.8") {
			lang.SayString("xsub469")
		}
	}
	rexsult, err = lang.RxFromString("412411244.E-774339264").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("866452.465"))
	if err != nil {
		lang.SayString("xsub470")
	} else {
		if !(rexsult.ToString() == "-866452.465") {
			lang.SayString("xsub470")
		}
	}
	rexsult, err = lang.RxFromString("-103.474598").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3.01660661E-446661257"))
	if err != nil {
		lang.SayString("xsub471")
	} else {
		if !(rexsult.ToString() == "-103.474598") {
			lang.SayString("xsub471")
		}
	}
	rexsult, err = lang.RxFromString("-31027.8323").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-475378186."))
	if err != nil {
		lang.SayString("xsub472")
	} else {
		if !(rexsult.ToString() == "475347158") {
			lang.SayString("xsub472")
		}
	}
	rexsult, err = lang.RxFromString("-1199339.72").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5.73068392E+53774632"))
	if err != nil {
		lang.SayString("xsub473")
	} else {
		if !(rexsult.ToString() == "5.73068392E+53774632") {
			lang.SayString("xsub473")
		}
	}
	rexsult, err = lang.RxFromString("-732908.930E+364345433").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3486146.26"))
	if err != nil {
		lang.SayString("xsub474")
	} else {
		if !(rexsult.ToString() == "-7.32908930E+364345438") {
			lang.SayString("xsub474")
		}
	}
	rexsult, err = lang.RxFromString("-2376150.83").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-46777583.3"))
	if err != nil {
		lang.SayString("xsub475")
	} else {
		if !(rexsult.ToString() == "44401432.5") {
			lang.SayString("xsub475")
		}
	}
	rexsult, err = lang.RxFromString("6.3664211").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-140854908."))
	if err != nil {
		lang.SayString("xsub476")
	} else {
		if !(rexsult.ToString() == "140854914") {
			lang.SayString("xsub476")
		}
	}
	rexsult, err = lang.RxFromString("-15.791522").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1902.30210E+90741844"))
	if err != nil {
		lang.SayString("xsub477")
	} else {
		if !(rexsult.ToString() == "-1.90230210E+90741847") {
			lang.SayString("xsub477")
		}
	}
	rexsult, err = lang.RxFromString("15356.1505E+373950429").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2.88020400"))
	if err != nil {
		lang.SayString("xsub478")
	} else {
		if !(rexsult.ToString() == "1.53561505E+373950433") {
			lang.SayString("xsub478")
		}
	}
	rexsult, err = lang.RxFromString("-3.12001326E+318884762").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("9567.21595"))
	if err != nil {
		lang.SayString("xsub479")
	} else {
		if !(rexsult.ToString() == "-3.12001326E+318884762") {
			lang.SayString("xsub479")
		}
	}
	rexsult, err = lang.RxFromString("49436.6528").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("751.919517"))
	if err != nil {
		lang.SayString("xsub480")
	} else {
		if !(rexsult.ToString() == "48684.7333") {
			lang.SayString("xsub480")
		}
	}
	rexsult, err = lang.RxFromString("552.669453").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("8.3725760E+16223526"))
	if err != nil {
		lang.SayString("xsub481")
	} else {
		if !(rexsult.ToString() == "-8.37257600E+16223526") {
			lang.SayString("xsub481")
		}
	}
	rexsult, err = lang.RxFromString("-3266303").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("453741.520"))
	if err != nil {
		lang.SayString("xsub482")
	} else {
		if !(rexsult.ToString() == "-3720044.52") {
			lang.SayString("xsub482")
		}
	}
	rexsult, err = lang.RxFromString("12302757.4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("542922.487E+414443353"))
	if err != nil {
		lang.SayString("xsub483")
	} else {
		if !(rexsult.ToString() == "-5.42922487E+414443358") {
			lang.SayString("xsub483")
		}
	}
	rexsult, err = lang.RxFromString("-5670757.79E-784754984").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("128144.503"))
	if err != nil {
		lang.SayString("xsub484")
	} else {
		if !(rexsult.ToString() == "-128144.503") {
			lang.SayString("xsub484")
		}
	}
	rexsult, err = lang.RxFromString("22.7721968E+842530698").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("5223.70462"))
	if err != nil {
		lang.SayString("xsub485")
	} else {
		if !(rexsult.ToString() == "2.27721968E+842530699") {
			lang.SayString("xsub485")
		}
	}
	rexsult, err = lang.RxFromString("88.5158199E-980164357").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("325846116"))
	if err != nil {
		lang.SayString("xsub486")
	} else {
		if !(rexsult.ToString() == "-325846116") {
			lang.SayString("xsub486")
		}
	}
	rexsult, err = lang.RxFromString("-22881.0408").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("5.63661562"))
	if err != nil {
		lang.SayString("xsub487")
	} else {
		if !(rexsult.ToString() == "-22886.6774") {
			lang.SayString("xsub487")
		}
	}
	rexsult, err = lang.RxFromString("-7157.57449").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-76.4455519E-85647047"))
	if err != nil {
		lang.SayString("xsub488")
	} else {
		if !(rexsult.ToString() == "-7157.57449") {
			lang.SayString("xsub488")
		}
	}
	rexsult, err = lang.RxFromString("-503113.801").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-9715149.82E-612184422"))
	if err != nil {
		lang.SayString("xsub489")
	} else {
		if !(rexsult.ToString() == "-503113.801") {
			lang.SayString("xsub489")
		}
	}
	rexsult, err = lang.RxFromString("-3066962.41").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-55.3096879"))
	if err != nil {
		lang.SayString("xsub490")
	} else {
		if !(rexsult.ToString() == "-3066907.10") {
			lang.SayString("xsub490")
		}
	}
	rexsult, err = lang.RxFromString("-53311.5738E+156608936").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-7.45890666"))
	if err != nil {
		lang.SayString("xsub491")
	} else {
		if !(rexsult.ToString() == "-5.33115738E+156608940") {
			lang.SayString("xsub491")
		}
	}
	rexsult, err = lang.RxFromString("998890068.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-92.057879"))
	if err != nil {
		lang.SayString("xsub492")
	} else {
		if !(rexsult.ToString() == "998890160") {
			lang.SayString("xsub492")
		}
	}
	rexsult, err = lang.RxFromString("122.495591").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-407836028."))
	if err != nil {
		lang.SayString("xsub493")
	} else {
		if !(rexsult.ToString() == "407836150") {
			lang.SayString("xsub493")
		}
	}
	rexsult, err = lang.RxFromString("187098.488").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("6220.05584E-236541249"))
	if err != nil {
		lang.SayString("xsub494")
	} else {
		if !(rexsult.ToString() == "187098.488") {
			lang.SayString("xsub494")
		}
	}
	rexsult, err = lang.RxFromString("4819899.21E+432982550").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-727441917"))
	if err != nil {
		lang.SayString("xsub495")
	} else {
		if !(rexsult.ToString() == "4.81989921E+432982556") {
			lang.SayString("xsub495")
		}
	}
	rexsult, err = lang.RxFromString("5770.01020E+507459752").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-4208339.33E-129766680"))
	if err != nil {
		lang.SayString("xsub496")
	} else {
		if !(rexsult.ToString() == "5.77001020E+507459755") {
			lang.SayString("xsub496")
		}
	}
	rexsult, err = lang.RxFromString("-286.371320").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("710319152"))
	if err != nil {
		lang.SayString("xsub497")
	} else {
		if !(rexsult.ToString() == "-710319438") {
			lang.SayString("xsub497")
		}
	}
	rexsult, err = lang.RxFromString("-7.27403536").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-481469656E-835183700"))
	if err != nil {
		lang.SayString("xsub498")
	} else {
		if !(rexsult.ToString() == "-7.27403536") {
			lang.SayString("xsub498")
		}
	}
	rexsult, err = lang.RxFromString("-6157.74292").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-94075286.2E+92555877"))
	if err != nil {
		lang.SayString("xsub499")
	} else {
		if !(rexsult.ToString() == "9.40752862E+92555884") {
			lang.SayString("xsub499")
		}
	}
	rexsult, err = lang.RxFromString("-525445087.E+231529167").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("188227460"))
	if err != nil {
		lang.SayString("xsub500")
	} else {
		if !(rexsult.ToString() == "-5.25445087E+231529175") {
			lang.SayString("xsub500")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx001")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx001")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("subx003")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx004")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("subx005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx005")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("subx006")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx006")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("subx007")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx007")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx011")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx011")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx012")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("subx012")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("subx013")
	} else {
		if !(rexsult.ToString() == "-3") {
			lang.SayString("subx013")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx014")
	} else {
		if !(rexsult.ToString() == "-3") {
			lang.SayString("subx014")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("subx015")
	} else {
		if !(rexsult.ToString() == "-4") {
			lang.SayString("subx015")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("subx016")
	} else {
		if !(rexsult.ToString() == "-5") {
			lang.SayString("subx016")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("subx017")
	} else {
		if !(rexsult.ToString() == "-5") {
			lang.SayString("subx017")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("subx021")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx021")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("subx022")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("subx022")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("subx023")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("subx023")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("subx024")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("subx024")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("subx025")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("subx025")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("subx026")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("subx026")
		}
	}
	rexsult, err = lang.RxFromString("2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("subx027")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("subx027")
		}
	}
	rexsult, err = lang.RxFromString("11").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx030")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("subx030")
		}
	}
	rexsult, err = lang.RxFromString("10").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx031")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("subx031")
		}
	}
	rexsult, err = lang.RxFromString("9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx032")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("subx032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx033")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx033")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx034")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx034")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx035")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("subx035")
		}
	}
	rexsult, err = lang.RxFromString("-9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx036")
	} else {
		if !(rexsult.ToString() == "-10") {
			lang.SayString("subx036")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx037")
	} else {
		if !(rexsult.ToString() == "-11") {
			lang.SayString("subx037")
		}
	}
	rexsult, err = lang.RxFromString("-11").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx038")
	} else {
		if !(rexsult.ToString() == "-12") {
			lang.SayString("subx038")
		}
	}
	rexsult, err = lang.RxFromString("5.75").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3.3"))
	if err != nil {
		lang.SayString("subx040")
	} else {
		if !(rexsult.ToString() == "2.45") {
			lang.SayString("subx040")
		}
	}
	rexsult, err = lang.RxFromString("5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("subx041")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("subx041")
		}
	}
	rexsult, err = lang.RxFromString("-5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("subx042")
	} else {
		if !(rexsult.ToString() == "-2") {
			lang.SayString("subx042")
		}
	}
	rexsult, err = lang.RxFromString("-7").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2.5"))
	if err != nil {
		lang.SayString("subx043")
	} else {
		if !(rexsult.ToString() == "-9.5") {
			lang.SayString("subx043")
		}
	}
	rexsult, err = lang.RxFromString("0.7").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("subx044")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("subx044")
		}
	}
	rexsult, err = lang.RxFromString("1.3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("subx045")
	} else {
		if !(rexsult.ToString() == "1.0") {
			lang.SayString("subx045")
		}
	}
	rexsult, err = lang.RxFromString("1.25").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.25"))
	if err != nil {
		lang.SayString("subx046")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx046")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00000000"))
	if err != nil {
		lang.SayString("subx050")
	} else {
		if !(rexsult.ToString() == "0.23456789") {
			lang.SayString("subx050")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00000089"))
	if err != nil {
		lang.SayString("subx051")
	} else {
		if !(rexsult.ToString() == "0.23456700") {
			lang.SayString("subx051")
		}
	}
	rexsult, err = lang.RxFromString("0.5555555559").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.0000000001"))
	if err != nil {
		lang.SayString("subx052")
	} else {
		if !(rexsult.ToString() == "0.555555556") {
			lang.SayString("subx052")
		}
	}
	rexsult, err = lang.RxFromString("0.5555555559").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.0000000005"))
	if err != nil {
		lang.SayString("subx053")
	} else {
		if !(rexsult.ToString() == "0.555555555") {
			lang.SayString("subx053")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.1111111111"))
	if err != nil {
		lang.SayString("subx054")
	} else {
		if !(rexsult.ToString() == "0.333333333") {
			lang.SayString("subx054")
		}
	}
	rexsult, err = lang.RxFromString("1.0000000000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.00000001"))
	if err != nil {
		lang.SayString("subx055")
	} else {
		if !(rexsult.ToString() == "0.99999999") {
			lang.SayString("subx055")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx056")
	} else {
		if !(rexsult.ToString() == "0.444444444") {
			lang.SayString("subx056")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444445000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx057")
	} else {
		if !(rexsult.ToString() == "0.444444445") {
			lang.SayString("subx057")
		}
	}
	rexsult, err = lang.RxFromString("70").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx060")
	} else {
		if !(rexsult.ToString() == "-1.00000000E+13") {
			lang.SayString("subx060")
		}
	}
	rexsult, err = lang.RxFromString("700").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx061")
	} else {
		if !(rexsult.ToString() == "-1.00000000E+13") {
			lang.SayString("subx061")
		}
	}
	rexsult, err = lang.RxFromString("7000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx062")
	} else {
		if !(rexsult.ToString() == "-1.00000000E+13") {
			lang.SayString("subx062")
		}
	}
	rexsult, err = lang.RxFromString("70000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx063")
	} else {
		if !(rexsult.ToString() == "-9.9999999E+12") {
			lang.SayString("subx063")
		}
	}
	rexsult, err = lang.RxFromString("700000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx064")
	} else {
		if !(rexsult.ToString() == "-9.9999993E+12") {
			lang.SayString("subx064")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("70"))
	if err != nil {
		lang.SayString("subx065")
	} else {
		if !(rexsult.ToString() == "1.00000000E+13") {
			lang.SayString("subx065")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("700"))
	if err != nil {
		lang.SayString("subx066")
	} else {
		if !(rexsult.ToString() == "1.00000000E+13") {
			lang.SayString("subx066")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7000"))
	if err != nil {
		lang.SayString("subx067")
	} else {
		if !(rexsult.ToString() == "1.00000000E+13") {
			lang.SayString("subx067")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("subx068")
	} else {
		if !(rexsult.ToString() == "9.9999999E+12") {
			lang.SayString("subx068")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("700000"))
	if err != nil {
		lang.SayString("subx069")
	} else {
		if !(rexsult.ToString() == "9.9999993E+12") {
			lang.SayString("subx069")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("subx080")
	} else {
		if !(rexsult.ToString() == "9.9999999E+12") {
			lang.SayString("subx080")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("subx081")
	} else {
		if !(rexsult.ToString() == "1.00000E+13") {
			lang.SayString("subx081")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("subx090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx090")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.00"))
	if err != nil {
		lang.SayString("subx091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("00.0"))
	if err != nil {
		lang.SayString("subx092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx092")
		}
	}
	rexsult, err = lang.RxFromString("00.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.00"))
	if err != nil {
		lang.SayString("subx093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx093")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("00.0"))
	if err != nil {
		lang.SayString("subx094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx094")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("subx095")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("subx095")
		}
	}
	rexsult, err = lang.RxFromString("3.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("subx096")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("subx096")
		}
	}
	rexsult, err = lang.RxFromString("3.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("subx097")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("subx097")
		}
	}
	rexsult, err = lang.RxFromString("3.00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".3"))
	if err != nil {
		lang.SayString("subx098")
	} else {
		if !(rexsult.ToString() == "2.70") {
			lang.SayString("subx098")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("subx099")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx099")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("+3"))
	if err != nil {
		lang.SayString("subx100")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx100")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("subx101")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("subx101")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("subx102")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("subx102")
		}
	}
	rexsult, err = lang.RxFromString("3.").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("subx103")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("subx103")
		}
	}
	rexsult, err = lang.RxFromString("3.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("subx104")
	} else {
		if !(rexsult.ToString() == "2.7") {
			lang.SayString("subx104")
		}
	}
	rexsult, err = lang.RxFromString("3.00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.3"))
	if err != nil {
		lang.SayString("subx105")
	} else {
		if !(rexsult.ToString() == "2.70") {
			lang.SayString("subx105")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("3.0"))
	if err != nil {
		lang.SayString("subx106")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx106")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("+3.0"))
	if err != nil {
		lang.SayString("subx107")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx107")
		}
	}
	rexsult, err = lang.RxFromString("3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-3.0"))
	if err != nil {
		lang.SayString("subx108")
	} else {
		if !(rexsult.ToString() == "6.0") {
			lang.SayString("subx108")
		}
	}
	rexsult, err = lang.RxFromString("10.23456784").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx120")
	} else {
		if !(rexsult.ToString() == "-1E-7") {
			lang.SayString("subx120")
		}
	}
	rexsult, err = lang.RxFromString("10.23456785").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx121")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx121")
		}
	}
	rexsult, err = lang.RxFromString("10.23456786").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx122")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx122")
		}
	}
	rexsult, err = lang.RxFromString("10.23456787").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx123")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx123")
		}
	}
	rexsult, err = lang.RxFromString("10.23456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx124")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx124")
		}
	}
	rexsult, err = lang.RxFromString("10.23456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx125")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx125")
		}
	}
	rexsult, err = lang.RxFromString("10.23456790").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx126")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx126")
		}
	}
	rexsult, err = lang.RxFromString("10.23456791").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx127")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx127")
		}
	}
	rexsult, err = lang.RxFromString("10.23456792").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx128")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx128")
		}
	}
	rexsult, err = lang.RxFromString("10.23456793").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx129")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx129")
		}
	}
	rexsult, err = lang.RxFromString("10.23456794").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if err != nil {
		lang.SayString("subx130")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("subx130")
		}
	}
	rexsult, err = lang.RxFromString("10.23456781").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx131")
	} else {
		if !(rexsult.ToString() == "-1E-7") {
			lang.SayString("subx131")
		}
	}
	rexsult, err = lang.RxFromString("10.23456782").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx132")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx132")
		}
	}
	rexsult, err = lang.RxFromString("10.23456783").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx133")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx133")
		}
	}
	rexsult, err = lang.RxFromString("10.23456784").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx134")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx134")
		}
	}
	rexsult, err = lang.RxFromString("10.23456785").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx135")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx135")
		}
	}
	rexsult, err = lang.RxFromString("10.23456786").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx136")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx136")
		}
	}
	rexsult, err = lang.RxFromString("10.23456787").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx137")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx137")
		}
	}
	rexsult, err = lang.RxFromString("10.23456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx138")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx138")
		}
	}
	rexsult, err = lang.RxFromString("10.23456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx139")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx139")
		}
	}
	rexsult, err = lang.RxFromString("10.23456790").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx140")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx140")
		}
	}
	rexsult, err = lang.RxFromString("10.23456791").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456786"))
	if err != nil {
		lang.SayString("subx141")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("subx141")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("subx142")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx142")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx143")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx143")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456780").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("subx144")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("subx144")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456790").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("subx145")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx145")
		}
	}
	rexsult, err = lang.RxFromString("-10.23456791").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-10.23456786"))
	if err != nil {
		lang.SayString("subx146")
	} else {
		if !(rexsult.ToString() == "-1E-7") {
			lang.SayString("subx146")
		}
	}
	rexsult, err = lang.RxFromString("12345678900000").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("9999999999999"))
	if err != nil {
		lang.SayString("subx150")
	} else {
		if !(rexsult.ToString() == "2.4E+12") {
			lang.SayString("subx150")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("12345678900000"))
	if err != nil {
		lang.SayString("subx151")
	} else {
		if !(rexsult.ToString() == "-2.4E+12") {
			lang.SayString("subx151")
		}
	}
	rexsult, err = lang.RxFromString("12345678900000").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("9999999999999"))
	if err != nil {
		lang.SayString("subx152")
	} else {
		if !(rexsult.ToString() == "2.3457E+12") {
			lang.SayString("subx152")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("12345678900000"))
	if err != nil {
		lang.SayString("subx153")
	} else {
		if !(rexsult.ToString() == "-2.3457E+12") {
			lang.SayString("subx153")
		}
	}
	rexsult, err = lang.RxFromString("12345678900000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("9999999999999"))
	if err != nil {
		lang.SayString("subx154")
	} else {
		if !(rexsult.ToString() == "2.3456789E+12") {
			lang.SayString("subx154")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("12345678900000"))
	if err != nil {
		lang.SayString("subx155")
	} else {
		if !(rexsult.ToString() == "-2.3456789E+12") {
			lang.SayString("subx155")
		}
	}
	rexsult, err = lang.RxFromString("12345678900000").OpSub(lang.RxSetWithDigit(12), lang.RxFromString("9999999999999"))
	if err != nil {
		lang.SayString("subx156")
	} else {
		if !(rexsult.ToString() == "2.3456789000E+12") {
			lang.SayString("subx156")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999").OpSub(lang.RxSetWithDigit(12), lang.RxFromString("12345678900000"))
	if err != nil {
		lang.SayString("subx157")
	} else {
		if !(rexsult.ToString() == "-2.3456789000E+12") {
			lang.SayString("subx157")
		}
	}
	rexsult, err = lang.RxFromString("12345678900000").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("9999999999999"))
	if err != nil {
		lang.SayString("subx158")
	} else {
		if !(rexsult.ToString() == "2345678900001") {
			lang.SayString("subx158")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("12345678900000"))
	if err != nil {
		lang.SayString("subx159")
	} else {
		if !(rexsult.ToString() == "-2345678900001") {
			lang.SayString("subx159")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".1"))
	if err != nil {
		lang.SayString("subx160")
	} else {
		if !(rexsult.ToString() == "-0.1") {
			lang.SayString("subx160")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".97983"))
	if err != nil {
		lang.SayString("subx161")
	} else {
		if !(rexsult.ToString() == "-0.97983") {
			lang.SayString("subx161")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".9"))
	if err != nil {
		lang.SayString("subx162")
	} else {
		if !(rexsult.ToString() == "-0.9") {
			lang.SayString("subx162")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.102"))
	if err != nil {
		lang.SayString("subx163")
	} else {
		if !(rexsult.ToString() == "-0.102") {
			lang.SayString("subx163")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".4"))
	if err != nil {
		lang.SayString("subx164")
	} else {
		if !(rexsult.ToString() == "-0.4") {
			lang.SayString("subx164")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".307"))
	if err != nil {
		lang.SayString("subx165")
	} else {
		if !(rexsult.ToString() == "-0.307") {
			lang.SayString("subx165")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".43822"))
	if err != nil {
		lang.SayString("subx166")
	} else {
		if !(rexsult.ToString() == "-0.43822") {
			lang.SayString("subx166")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".911"))
	if err != nil {
		lang.SayString("subx167")
	} else {
		if !(rexsult.ToString() == "-0.911") {
			lang.SayString("subx167")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".02"))
	if err != nil {
		lang.SayString("subx168")
	} else {
		if !(rexsult.ToString() == "-0.02") {
			lang.SayString("subx168")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".392"))
	if err != nil {
		lang.SayString("subx169")
	} else {
		if !(rexsult.ToString() == "-0.392") {
			lang.SayString("subx169")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".26"))
	if err != nil {
		lang.SayString("subx170")
	} else {
		if !(rexsult.ToString() == "-0.26") {
			lang.SayString("subx170")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("subx171")
	} else {
		if !(rexsult.ToString() == "-0.51") {
			lang.SayString("subx171")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".2234"))
	if err != nil {
		lang.SayString("subx172")
	} else {
		if !(rexsult.ToString() == "-0.2234") {
			lang.SayString("subx172")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".2"))
	if err != nil {
		lang.SayString("subx173")
	} else {
		if !(rexsult.ToString() == "-0.2") {
			lang.SayString("subx173")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString(".0008"))
	if err != nil {
		lang.SayString("subx174")
	} else {
		if !(rexsult.ToString() == "-0.0008") {
			lang.SayString("subx174")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.1"))
	if err != nil {
		lang.SayString("subx180")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("subx180")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.97983"))
	if err != nil {
		lang.SayString("subx181")
	} else {
		if !(rexsult.ToString() == "0.97983") {
			lang.SayString("subx181")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.9"))
	if err != nil {
		lang.SayString("subx182")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("subx182")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0.102"))
	if err != nil {
		lang.SayString("subx183")
	} else {
		if !(rexsult.ToString() == "0.102") {
			lang.SayString("subx183")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.4"))
	if err != nil {
		lang.SayString("subx184")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("subx184")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.307"))
	if err != nil {
		lang.SayString("subx185")
	} else {
		if !(rexsult.ToString() == "0.307") {
			lang.SayString("subx185")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.43822"))
	if err != nil {
		lang.SayString("subx186")
	} else {
		if !(rexsult.ToString() == "0.43822") {
			lang.SayString("subx186")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.911"))
	if err != nil {
		lang.SayString("subx187")
	} else {
		if !(rexsult.ToString() == "0.911") {
			lang.SayString("subx187")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.02"))
	if err != nil {
		lang.SayString("subx188")
	} else {
		if !(rexsult.ToString() == "0.02") {
			lang.SayString("subx188")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.392"))
	if err != nil {
		lang.SayString("subx189")
	} else {
		if !(rexsult.ToString() == "0.392") {
			lang.SayString("subx189")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.26"))
	if err != nil {
		lang.SayString("subx190")
	} else {
		if !(rexsult.ToString() == "0.26") {
			lang.SayString("subx190")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0.51"))
	if err != nil {
		lang.SayString("subx191")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("subx191")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.2234"))
	if err != nil {
		lang.SayString("subx192")
	} else {
		if !(rexsult.ToString() == "0.2234") {
			lang.SayString("subx192")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.2"))
	if err != nil {
		lang.SayString("subx193")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("subx193")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.0008"))
	if err != nil {
		lang.SayString("subx194")
	} else {
		if !(rexsult.ToString() == "0.0008") {
			lang.SayString("subx194")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.1"))
	if err != nil {
		lang.SayString("subx200")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("subx200")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.97983"))
	if err != nil {
		lang.SayString("subx201")
	} else {
		if !(rexsult.ToString() == "0.97983") {
			lang.SayString("subx201")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.9"))
	if err != nil {
		lang.SayString("subx202")
	} else {
		if !(rexsult.ToString() == "0.9") {
			lang.SayString("subx202")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0.102"))
	if err != nil {
		lang.SayString("subx203")
	} else {
		if !(rexsult.ToString() == "0.102") {
			lang.SayString("subx203")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.4"))
	if err != nil {
		lang.SayString("subx204")
	} else {
		if !(rexsult.ToString() == "0.4") {
			lang.SayString("subx204")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.307"))
	if err != nil {
		lang.SayString("subx205")
	} else {
		if !(rexsult.ToString() == "0.307") {
			lang.SayString("subx205")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.43822"))
	if err != nil {
		lang.SayString("subx206")
	} else {
		if !(rexsult.ToString() == "0.43822") {
			lang.SayString("subx206")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.911"))
	if err != nil {
		lang.SayString("subx207")
	} else {
		if !(rexsult.ToString() == "0.911") {
			lang.SayString("subx207")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.02"))
	if err != nil {
		lang.SayString("subx208")
	} else {
		if !(rexsult.ToString() == "0.02") {
			lang.SayString("subx208")
		}
	}
	rexsult, err = lang.RxFromString("00").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.392"))
	if err != nil {
		lang.SayString("subx209")
	} else {
		if !(rexsult.ToString() == "0.392") {
			lang.SayString("subx209")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.26"))
	if err != nil {
		lang.SayString("subx210")
	} else {
		if !(rexsult.ToString() == "0.26") {
			lang.SayString("subx210")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0.51"))
	if err != nil {
		lang.SayString("subx211")
	} else {
		if !(rexsult.ToString() == "0.51") {
			lang.SayString("subx211")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.2234"))
	if err != nil {
		lang.SayString("subx212")
	} else {
		if !(rexsult.ToString() == "0.2234") {
			lang.SayString("subx212")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.2"))
	if err != nil {
		lang.SayString("subx213")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("subx213")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-.0008"))
	if err != nil {
		lang.SayString("subx214")
	} else {
		if !(rexsult.ToString() == "0.0008") {
			lang.SayString("subx214")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-12").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx220")
	} else {
		if !(rexsult.ToString() == "-5.6267E-8") {
			lang.SayString("subx220")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-11").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx221")
	} else {
		if !(rexsult.ToString() == "-5.6267E-7") {
			lang.SayString("subx221")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-10").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx222")
	} else {
		if !(rexsult.ToString() == "-0.0000056267") {
			lang.SayString("subx222")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx223")
	} else {
		if !(rexsult.ToString() == "-0.000056267") {
			lang.SayString("subx223")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-8").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx224")
	} else {
		if !(rexsult.ToString() == "-0.00056267") {
			lang.SayString("subx224")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-7").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx225")
	} else {
		if !(rexsult.ToString() == "-0.0056267") {
			lang.SayString("subx225")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-6").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx226")
	} else {
		if !(rexsult.ToString() == "-0.056267") {
			lang.SayString("subx226")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx227")
	} else {
		if !(rexsult.ToString() == "-0.56267") {
			lang.SayString("subx227")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx228")
	} else {
		if !(rexsult.ToString() == "-562.67") {
			lang.SayString("subx228")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx229")
	} else {
		if !(rexsult.ToString() == "-5626.7") {
			lang.SayString("subx229")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx230")
	} else {
		if !(rexsult.ToString() == "-56267") {
			lang.SayString("subx230")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-12"))
	if err != nil {
		lang.SayString("subx240")
	} else {
		if !(rexsult.ToString() == "5.6267E-8") {
			lang.SayString("subx240")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-11"))
	if err != nil {
		lang.SayString("subx241")
	} else {
		if !(rexsult.ToString() == "5.6267E-7") {
			lang.SayString("subx241")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-10"))
	if err != nil {
		lang.SayString("subx242")
	} else {
		if !(rexsult.ToString() == "0.0000056267") {
			lang.SayString("subx242")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-9"))
	if err != nil {
		lang.SayString("subx243")
	} else {
		if !(rexsult.ToString() == "0.000056267") {
			lang.SayString("subx243")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-8"))
	if err != nil {
		lang.SayString("subx244")
	} else {
		if !(rexsult.ToString() == "0.00056267") {
			lang.SayString("subx244")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-7"))
	if err != nil {
		lang.SayString("subx245")
	} else {
		if !(rexsult.ToString() == "0.0056267") {
			lang.SayString("subx245")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-6"))
	if err != nil {
		lang.SayString("subx246")
	} else {
		if !(rexsult.ToString() == "0.056267") {
			lang.SayString("subx246")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-5"))
	if err != nil {
		lang.SayString("subx247")
	} else {
		if !(rexsult.ToString() == "0.56267") {
			lang.SayString("subx247")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-2"))
	if err != nil {
		lang.SayString("subx248")
	} else {
		if !(rexsult.ToString() == "562.67") {
			lang.SayString("subx248")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-1"))
	if err != nil {
		lang.SayString("subx249")
	} else {
		if !(rexsult.ToString() == "5626.7") {
			lang.SayString("subx249")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-0"))
	if err != nil {
		lang.SayString("subx250")
	} else {
		if !(rexsult.ToString() == "56267") {
			lang.SayString("subx250")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00000000"))
	if err != nil {
		lang.SayString("subx301")
	} else {
		if !(rexsult.ToString() == "0.23456789") {
			lang.SayString("subx301")
		}
	}
	rexsult, err = lang.RxFromString("1.23456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00000011"))
	if err != nil {
		lang.SayString("subx302")
	} else {
		if !(rexsult.ToString() == "0.23456778") {
			lang.SayString("subx302")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.5555555555"))
	if err != nil {
		lang.SayString("subx311")
	} else {
		if !(rexsult.ToString() == "-0.111111111") {
			lang.SayString("subx311")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444440").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.5555555555"))
	if err != nil {
		lang.SayString("subx312")
	} else {
		if !(rexsult.ToString() == "-0.111111112") {
			lang.SayString("subx312")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.5555555550"))
	if err != nil {
		lang.SayString("subx313")
	} else {
		if !(rexsult.ToString() == "-0.111111111") {
			lang.SayString("subx313")
		}
	}
	rexsult, err = lang.RxFromString("0.44444444449").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx314")
	} else {
		if !(rexsult.ToString() == "0.444444444") {
			lang.SayString("subx314")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444499").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx315")
	} else {
		if !(rexsult.ToString() == "0.444444444") {
			lang.SayString("subx315")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444444999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx316")
	} else {
		if !(rexsult.ToString() == "0.444444444") {
			lang.SayString("subx316")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444445000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx317")
	} else {
		if !(rexsult.ToString() == "0.444444445") {
			lang.SayString("subx317")
		}
	}
	rexsult, err = lang.RxFromString("0.4444444445001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx318")
	} else {
		if !(rexsult.ToString() == "0.444444445") {
			lang.SayString("subx318")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444501").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx319")
	} else {
		if !(rexsult.ToString() == "0.444444445") {
			lang.SayString("subx319")
		}
	}
	rexsult, err = lang.RxFromString("0.44444444451").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx320")
	} else {
		if !(rexsult.ToString() == "0.444444445") {
			lang.SayString("subx320")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.0000"))
	if err != nil {
		lang.SayString("subx321")
	} else {
		if !(rexsult.ToString() == "0.9998") {
			lang.SayString("subx321")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.0001"))
	if err != nil {
		lang.SayString("subx322")
	} else {
		if !(rexsult.ToString() == "0.9997") {
			lang.SayString("subx322")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.0002"))
	if err != nil {
		lang.SayString("subx323")
	} else {
		if !(rexsult.ToString() == "0.9996") {
			lang.SayString("subx323")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.0003"))
	if err != nil {
		lang.SayString("subx324")
	} else {
		if !(rexsult.ToString() == "0.9995") {
			lang.SayString("subx324")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0.0000"))
	if err != nil {
		lang.SayString("subx325")
	} else {
		if !(rexsult.ToString() == "0.9998") {
			lang.SayString("subx325")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0.0001"))
	if err != nil {
		lang.SayString("subx326")
	} else {
		if !(rexsult.ToString() == "0.9999") {
			lang.SayString("subx326")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0.0002"))
	if err != nil {
		lang.SayString("subx327")
	} else {
		if !(rexsult.ToString() == "1.0000") {
			lang.SayString("subx327")
		}
	}
	rexsult, err = lang.RxFromString("0.9998").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-0.0003"))
	if err != nil {
		lang.SayString("subx328")
	} else {
		if !(rexsult.ToString() == "1.0001") {
			lang.SayString("subx328")
		}
	}
	rexsult, err = lang.RxFromString("70").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx330")
	} else {
		if !(rexsult.ToString() == "-1.00000000E+13") {
			lang.SayString("subx330")
		}
	}
	rexsult, err = lang.RxFromString("700").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx331")
	} else {
		if !(rexsult.ToString() == "-1.00000000E+13") {
			lang.SayString("subx331")
		}
	}
	rexsult, err = lang.RxFromString("7000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx332")
	} else {
		if !(rexsult.ToString() == "-1.00000000E+13") {
			lang.SayString("subx332")
		}
	}
	rexsult, err = lang.RxFromString("70000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx333")
	} else {
		if !(rexsult.ToString() == "-9.9999999E+12") {
			lang.SayString("subx333")
		}
	}
	rexsult, err = lang.RxFromString("700000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx334")
	} else {
		if !(rexsult.ToString() == "-9.9999993E+12") {
			lang.SayString("subx334")
		}
	}
	rexsult, err = lang.RxFromString("7000000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx335")
	} else {
		if !(rexsult.ToString() == "-9.9999930E+12") {
			lang.SayString("subx335")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("70"))
	if err != nil {
		lang.SayString("subx340")
	} else {
		if !(rexsult.ToString() == "1.00000000E+13") {
			lang.SayString("subx340")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("700"))
	if err != nil {
		lang.SayString("subx341")
	} else {
		if !(rexsult.ToString() == "1.00000000E+13") {
			lang.SayString("subx341")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7000"))
	if err != nil {
		lang.SayString("subx342")
	} else {
		if !(rexsult.ToString() == "1.00000000E+13") {
			lang.SayString("subx342")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("subx343")
	} else {
		if !(rexsult.ToString() == "9.9999999E+12") {
			lang.SayString("subx343")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("700000"))
	if err != nil {
		lang.SayString("subx344")
	} else {
		if !(rexsult.ToString() == "9.9999993E+12") {
			lang.SayString("subx344")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("7000000"))
	if err != nil {
		lang.SayString("subx345")
	} else {
		if !(rexsult.ToString() == "9.9999930E+12") {
			lang.SayString("subx345")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("subx346")
	} else {
		if !(rexsult.ToString() == "9999999999993") {
			lang.SayString("subx346")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("70"))
	if err != nil {
		lang.SayString("subx347")
	} else {
		if !(rexsult.ToString() == "9999999999930") {
			lang.SayString("subx347")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("700"))
	if err != nil {
		lang.SayString("subx348")
	} else {
		if !(rexsult.ToString() == "9999999999300") {
			lang.SayString("subx348")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("7000"))
	if err != nil {
		lang.SayString("subx349")
	} else {
		if !(rexsult.ToString() == "9999999993000") {
			lang.SayString("subx349")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("subx350")
	} else {
		if !(rexsult.ToString() == "9999999930000") {
			lang.SayString("subx350")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("700000"))
	if err != nil {
		lang.SayString("subx351")
	} else {
		if !(rexsult.ToString() == "9999999300000") {
			lang.SayString("subx351")
		}
	}
	rexsult, err = lang.RxFromString("7").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx352")
	} else {
		if !(rexsult.ToString() == "-9999999999993") {
			lang.SayString("subx352")
		}
	}
	rexsult, err = lang.RxFromString("70").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx353")
	} else {
		if !(rexsult.ToString() == "-9999999999930") {
			lang.SayString("subx353")
		}
	}
	rexsult, err = lang.RxFromString("700").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx354")
	} else {
		if !(rexsult.ToString() == "-9999999999300") {
			lang.SayString("subx354")
		}
	}
	rexsult, err = lang.RxFromString("7000").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx355")
	} else {
		if !(rexsult.ToString() == "-9999999993000") {
			lang.SayString("subx355")
		}
	}
	rexsult, err = lang.RxFromString("70000").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx356")
	} else {
		if !(rexsult.ToString() == "-9999999930000") {
			lang.SayString("subx356")
		}
	}
	rexsult, err = lang.RxFromString("700000").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("10000e+9"))
	if err != nil {
		lang.SayString("subx357")
	} else {
		if !(rexsult.ToString() == "-9999999300000") {
			lang.SayString("subx357")
		}
	}
	rexsult, err = lang.RxFromString("10000e+9").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("70000"))
	if err != nil {
		lang.SayString("subx360")
	} else {
		if !(rexsult.ToString() == "1.00000E+13") {
			lang.SayString("subx360")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("0.0001"))
	if err != nil {
		lang.SayString("subx361")
	} else {
		if !(rexsult.ToString() == "0.9999") {
			lang.SayString("subx361")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("0.00001"))
	if err != nil {
		lang.SayString("subx362")
	} else {
		if !(rexsult.ToString() == "0.99999") {
			lang.SayString("subx362")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("subx363")
	} else {
		if !(rexsult.ToString() == "1.00000") {
			lang.SayString("subx363")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("0.0000001"))
	if err != nil {
		lang.SayString("subx364")
	} else {
		if !(rexsult.ToString() == "1.00000") {
			lang.SayString("subx364")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("0.00000001"))
	if err != nil {
		lang.SayString("subx365")
	} else {
		if !(rexsult.ToString() == "1.00000") {
			lang.SayString("subx365")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx370")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx370")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("0."))
	if err != nil {
		lang.SayString("subx371")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx371")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString(".0"))
	if err != nil {
		lang.SayString("subx372")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx372")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("0.0"))
	if err != nil {
		lang.SayString("subx373")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx373")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx374")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx374")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx375")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx375")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx376")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx376")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx377")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx377")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-51897955.3"))
	if err != nil {
		lang.SayString("subx910")
	} else {
		if !(rexsult.ToString() == "-51621407") {
			lang.SayString("subx910")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("89827.5229"))
	if err != nil {
		lang.SayString("subx911")
	} else {
		if !(rexsult.ToString() == "69751.921") {
			lang.SayString("subx911")
		}
	}
	rexsult, err = lang.RxFromString("333.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("33.1234566"))
	if err != nil {
		lang.SayString("subx920")
	} else {
		if !(rexsult.ToString() == "299.999999") {
			lang.SayString("subx920")
		}
	}
	rexsult, err = lang.RxFromString("333.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("33.1234565"))
	if err != nil {
		lang.SayString("subx921")
	} else {
		if !(rexsult.ToString() == "300.000000") {
			lang.SayString("subx921")
		}
	}
	rexsult, err = lang.RxFromString("133.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("33.1234565"))
	if err != nil {
		lang.SayString("subx922")
	} else {
		if !(rexsult.ToString() == "100.000000") {
			lang.SayString("subx922")
		}
	}
	rexsult, err = lang.RxFromString("133.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("33.1234564"))
	if err != nil {
		lang.SayString("subx923")
	} else {
		if !(rexsult.ToString() == "100.000000") {
			lang.SayString("subx923")
		}
	}
	rexsult, err = lang.RxFromString("133.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("33.1234540"))
	if err != nil {
		lang.SayString("subx924")
	} else {
		if !(rexsult.ToString() == "100.000002") {
			lang.SayString("subx924")
		}
	}
	rexsult, err = lang.RxFromString("133.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("43.1234560"))
	if err != nil {
		lang.SayString("subx925")
	} else {
		if !(rexsult.ToString() == "90.000000") {
			lang.SayString("subx925")
		}
	}
	rexsult, err = lang.RxFromString("133.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("43.1234561"))
	if err != nil {
		lang.SayString("subx926")
	} else {
		if !(rexsult.ToString() == "90.000000") {
			lang.SayString("subx926")
		}
	}
	rexsult, err = lang.RxFromString("133.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("43.1234566"))
	if err != nil {
		lang.SayString("subx927")
	} else {
		if !(rexsult.ToString() == "89.999999") {
			lang.SayString("subx927")
		}
	}
	rexsult, err = lang.RxFromString("101.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("91.1234566"))
	if err != nil {
		lang.SayString("subx928")
	} else {
		if !(rexsult.ToString() == "9.999999") {
			lang.SayString("subx928")
		}
	}
	rexsult, err = lang.RxFromString("101.123456").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("99.1234566"))
	if err != nil {
		lang.SayString("subx929")
	} else {
		if !(rexsult.ToString() == "1.999999") {
			lang.SayString("subx929")
		}
	}
	rexsult, err = lang.RxFromString("11").OpSub(lang.RxSetWithDigit(1), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("subx930")
	} else {
		if !(rexsult.ToString() == "1E+1") {
			lang.SayString("subx930")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(2), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("subx932")
	} else {
		if !(rexsult.ToString() == "1.0E+2") {
			lang.SayString("subx932")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("2.1"))
	if err != nil {
		lang.SayString("subx934")
	} else {
		if !(rexsult.ToString() == "99") {
			lang.SayString("subx934")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("92.01"))
	if err != nil {
		lang.SayString("subx935")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("subx935")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(4), lang.RxFromString("2.01"))
	if err != nil {
		lang.SayString("subx936")
	} else {
		if !(rexsult.ToString() == "99.0") {
			lang.SayString("subx936")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(4), lang.RxFromString("92.01"))
	if err != nil {
		lang.SayString("subx937")
	} else {
		if !(rexsult.ToString() == "9.0") {
			lang.SayString("subx937")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(4), lang.RxFromString("92.006"))
	if err != nil {
		lang.SayString("subx938")
	} else {
		if !(rexsult.ToString() == "9.0") {
			lang.SayString("subx938")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("2.001"))
	if err != nil {
		lang.SayString("subx939")
	} else {
		if !(rexsult.ToString() == "99.00") {
			lang.SayString("subx939")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("92.001"))
	if err != nil {
		lang.SayString("subx940")
	} else {
		if !(rexsult.ToString() == "9.00") {
			lang.SayString("subx940")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("92.0006"))
	if err != nil {
		lang.SayString("subx941")
	} else {
		if !(rexsult.ToString() == "9.00") {
			lang.SayString("subx941")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("2.0001"))
	if err != nil {
		lang.SayString("subx942")
	} else {
		if !(rexsult.ToString() == "99.000") {
			lang.SayString("subx942")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("92.0001"))
	if err != nil {
		lang.SayString("subx943")
	} else {
		if !(rexsult.ToString() == "9.000") {
			lang.SayString("subx943")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(6), lang.RxFromString("92.00006"))
	if err != nil {
		lang.SayString("subx944")
	} else {
		if !(rexsult.ToString() == "9.000") {
			lang.SayString("subx944")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(7), lang.RxFromString("2.00001"))
	if err != nil {
		lang.SayString("subx945")
	} else {
		if !(rexsult.ToString() == "99.0000") {
			lang.SayString("subx945")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(7), lang.RxFromString("92.00001"))
	if err != nil {
		lang.SayString("subx946")
	} else {
		if !(rexsult.ToString() == "9.0000") {
			lang.SayString("subx946")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(7), lang.RxFromString("92.000006"))
	if err != nil {
		lang.SayString("subx947")
	} else {
		if !(rexsult.ToString() == "9.0000") {
			lang.SayString("subx947")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(8), lang.RxFromString("2.000001"))
	if err != nil {
		lang.SayString("subx948")
	} else {
		if !(rexsult.ToString() == "99.00000") {
			lang.SayString("subx948")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(8), lang.RxFromString("92.000001"))
	if err != nil {
		lang.SayString("subx949")
	} else {
		if !(rexsult.ToString() == "9.00000") {
			lang.SayString("subx949")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(8), lang.RxFromString("92.0000006"))
	if err != nil {
		lang.SayString("subx950")
	} else {
		if !(rexsult.ToString() == "9.00000") {
			lang.SayString("subx950")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("2.0000001"))
	if err != nil {
		lang.SayString("subx951")
	} else {
		if !(rexsult.ToString() == "99.000000") {
			lang.SayString("subx951")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("92.0000001"))
	if err != nil {
		lang.SayString("subx952")
	} else {
		if !(rexsult.ToString() == "9.000000") {
			lang.SayString("subx952")
		}
	}
	rexsult, err = lang.RxFromString("101").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("92.00000006"))
	if err != nil {
		lang.SayString("subx953")
	} else {
		if !(rexsult.ToString() == "9.000000") {
			lang.SayString("subx953")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-10").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx390")
	} else {
		if !(rexsult.ToString() == "-0.0000056267") {
			lang.SayString("subx390")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-6").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx391")
	} else {
		if !(rexsult.ToString() == "-0.056267") {
			lang.SayString("subx391")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx392")
	} else {
		if !(rexsult.ToString() == "-0.56267") {
			lang.SayString("subx392")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx393")
	} else {
		if !(rexsult.ToString() == "-5.6267") {
			lang.SayString("subx393")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-3").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx394")
	} else {
		if !(rexsult.ToString() == "-56.267") {
			lang.SayString("subx394")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-2").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx395")
	} else {
		if !(rexsult.ToString() == "-562.67") {
			lang.SayString("subx395")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx396")
	} else {
		if !(rexsult.ToString() == "-5626.7") {
			lang.SayString("subx396")
		}
	}
	rexsult, err = lang.RxFromString("-56267E-0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx397")
	} else {
		if !(rexsult.ToString() == "-56267") {
			lang.SayString("subx397")
		}
	}
	rexsult, err = lang.RxFromString("-5E-10").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx398")
	} else {
		if !(rexsult.ToString() == "-5E-10") {
			lang.SayString("subx398")
		}
	}
	rexsult, err = lang.RxFromString("-5E-7").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx399")
	} else {
		if !(rexsult.ToString() == "-5E-7") {
			lang.SayString("subx399")
		}
	}
	rexsult, err = lang.RxFromString("-5E-6").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx400")
	} else {
		if !(rexsult.ToString() == "-0.000005") {
			lang.SayString("subx400")
		}
	}
	rexsult, err = lang.RxFromString("-5E-5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx401")
	} else {
		if !(rexsult.ToString() == "-0.00005") {
			lang.SayString("subx401")
		}
	}
	rexsult, err = lang.RxFromString("-5E-4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx402")
	} else {
		if !(rexsult.ToString() == "-0.0005") {
			lang.SayString("subx402")
		}
	}
	rexsult, err = lang.RxFromString("-5E-1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx403")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("subx403")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-10"))
	if err != nil {
		lang.SayString("subx420")
	} else {
		if !(rexsult.ToString() == "0.0000056267") {
			lang.SayString("subx420")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-6"))
	if err != nil {
		lang.SayString("subx421")
	} else {
		if !(rexsult.ToString() == "0.056267") {
			lang.SayString("subx421")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-5"))
	if err != nil {
		lang.SayString("subx422")
	} else {
		if !(rexsult.ToString() == "0.56267") {
			lang.SayString("subx422")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-4"))
	if err != nil {
		lang.SayString("subx423")
	} else {
		if !(rexsult.ToString() == "5.6267") {
			lang.SayString("subx423")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-3"))
	if err != nil {
		lang.SayString("subx424")
	} else {
		if !(rexsult.ToString() == "56.267") {
			lang.SayString("subx424")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-2"))
	if err != nil {
		lang.SayString("subx425")
	} else {
		if !(rexsult.ToString() == "562.67") {
			lang.SayString("subx425")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-1"))
	if err != nil {
		lang.SayString("subx426")
	} else {
		if !(rexsult.ToString() == "5626.7") {
			lang.SayString("subx426")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-56267E-0"))
	if err != nil {
		lang.SayString("subx427")
	} else {
		if !(rexsult.ToString() == "56267") {
			lang.SayString("subx427")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5E-10"))
	if err != nil {
		lang.SayString("subx428")
	} else {
		if !(rexsult.ToString() == "5E-10") {
			lang.SayString("subx428")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5E-7"))
	if err != nil {
		lang.SayString("subx429")
	} else {
		if !(rexsult.ToString() == "5E-7") {
			lang.SayString("subx429")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5E-6"))
	if err != nil {
		lang.SayString("subx430")
	} else {
		if !(rexsult.ToString() == "0.000005") {
			lang.SayString("subx430")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5E-5"))
	if err != nil {
		lang.SayString("subx431")
	} else {
		if !(rexsult.ToString() == "0.00005") {
			lang.SayString("subx431")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5E-4"))
	if err != nil {
		lang.SayString("subx432")
	} else {
		if !(rexsult.ToString() == "0.0005") {
			lang.SayString("subx432")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-5E-1"))
	if err != nil {
		lang.SayString("subx433")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("subx433")
		}
	}
	rexsult, err = lang.RxFromString("1E+12").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx461")
	} else {
		if !(rexsult.ToString() == "999999999999") {
			lang.SayString("subx461")
		}
	}
	rexsult, err = lang.RxFromString("1E+12").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-1.11"))
	if err != nil {
		lang.SayString("subx462")
	} else {
		if !(rexsult.ToString() == "1000000000001.11") {
			lang.SayString("subx462")
		}
	}
	rexsult, err = lang.RxFromString("1.11").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-1E+12"))
	if err != nil {
		lang.SayString("subx463")
	} else {
		if !(rexsult.ToString() == "1000000000001.11") {
			lang.SayString("subx463")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-1E+12"))
	if err != nil {
		lang.SayString("subx464")
	} else {
		if !(rexsult.ToString() == "999999999999") {
			lang.SayString("subx464")
		}
	}
	rexsult, err = lang.RxFromString("7E+12").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx465")
	} else {
		if !(rexsult.ToString() == "6999999999999") {
			lang.SayString("subx465")
		}
	}
	rexsult, err = lang.RxFromString("7E+12").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-1.11"))
	if err != nil {
		lang.SayString("subx466")
	} else {
		if !(rexsult.ToString() == "7000000000001.11") {
			lang.SayString("subx466")
		}
	}
	rexsult, err = lang.RxFromString("1.11").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-7E+12"))
	if err != nil {
		lang.SayString("subx467")
	} else {
		if !(rexsult.ToString() == "7000000000001.11") {
			lang.SayString("subx467")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-7E+12"))
	if err != nil {
		lang.SayString("subx468")
	} else {
		if !(rexsult.ToString() == "6999999999999") {
			lang.SayString("subx468")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555563"))
	if err != nil {
		lang.SayString("subx470")
	} else {
		if !(rexsult.ToString() == "1.00000000000001") {
			lang.SayString("subx470")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555562"))
	if err != nil {
		lang.SayString("subx471")
	} else {
		if !(rexsult.ToString() == "1.00000000000001") {
			lang.SayString("subx471")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555561"))
	if err != nil {
		lang.SayString("subx472")
	} else {
		if !(rexsult.ToString() == "1.00000000000001") {
			lang.SayString("subx472")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555560"))
	if err != nil {
		lang.SayString("subx473")
	} else {
		if !(rexsult.ToString() == "1.00000000000000") {
			lang.SayString("subx473")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555559"))
	if err != nil {
		lang.SayString("subx474")
	} else {
		if !(rexsult.ToString() == "1.00000000000000") {
			lang.SayString("subx474")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555558"))
	if err != nil {
		lang.SayString("subx475")
	} else {
		if !(rexsult.ToString() == "1.00000000000000") {
			lang.SayString("subx475")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555557"))
	if err != nil {
		lang.SayString("subx476")
	} else {
		if !(rexsult.ToString() == "1.00000000000000") {
			lang.SayString("subx476")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555556"))
	if err != nil {
		lang.SayString("subx477")
	} else {
		if !(rexsult.ToString() == "1.00000000000000") {
			lang.SayString("subx477")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555555"))
	if err != nil {
		lang.SayString("subx478")
	} else {
		if !(rexsult.ToString() == "0.999999999999999") {
			lang.SayString("subx478")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555554"))
	if err != nil {
		lang.SayString("subx479")
	} else {
		if !(rexsult.ToString() == "0.999999999999998") {
			lang.SayString("subx479")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555553"))
	if err != nil {
		lang.SayString("subx480")
	} else {
		if !(rexsult.ToString() == "0.999999999999997") {
			lang.SayString("subx480")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555552"))
	if err != nil {
		lang.SayString("subx481")
	} else {
		if !(rexsult.ToString() == "0.999999999999996") {
			lang.SayString("subx481")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555551"))
	if err != nil {
		lang.SayString("subx482")
	} else {
		if !(rexsult.ToString() == "0.999999999999995") {
			lang.SayString("subx482")
		}
	}
	rexsult, err = lang.RxFromString("0.444444444444444").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("-0.555555555555550"))
	if err != nil {
		lang.SayString("subx483")
	} else {
		if !(rexsult.ToString() == "0.999999999999994") {
			lang.SayString("subx483")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx500")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx500")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("subx501")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx501")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("subx502")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx502")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("subx503")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx503")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("subx504")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx504")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("subx505")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx505")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("subx506")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx506")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("subx507")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx507")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("subx508")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx508")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("subx509")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx509")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("subx510")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx510")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("subx511")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx511")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("subx512")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx512")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("subx513")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx513")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("subx514")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx514")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("subx515")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx515")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx516")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx516")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.000000001"))
	if err != nil {
		lang.SayString("subx517")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx517")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("subx518")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx518")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("subx519")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx519")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx520")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx520")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("subx521")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx521")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("subx522")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx522")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("subx523")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx523")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("subx524")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx524")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("subx525")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx525")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("subx526")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx526")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("subx527")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx527")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("subx528")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx528")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("subx529")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx529")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("subx530")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx530")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("subx531")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx531")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("subx532")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx532")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("subx533")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx533")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("subx534")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx534")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("subx535")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx535")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx536")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx536")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00000001"))
	if err != nil {
		lang.SayString("subx537")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx537")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("subx538")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx538")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("subx539")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx539")
		}
	}
	rexsult, err = lang.RxFromString("123456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("subx540")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx540")
		}
	}
	rexsult, err = lang.RxFromString("123456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("subx541")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx541")
		}
	}
	rexsult, err = lang.RxFromString("123456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("subx542")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx542")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx550")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx550")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.000000001"))
	if err != nil {
		lang.SayString("subx551")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx551")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.000001"))
	if err != nil {
		lang.SayString("subx552")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx552")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("subx553")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx553")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.4"))
	if err != nil {
		lang.SayString("subx554")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx554")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.49"))
	if err != nil {
		lang.SayString("subx555")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx555")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.499999"))
	if err != nil {
		lang.SayString("subx556")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx556")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.499999999"))
	if err != nil {
		lang.SayString("subx557")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx557")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.5"))
	if err != nil {
		lang.SayString("subx558")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx558")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.500000001"))
	if err != nil {
		lang.SayString("subx559")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx559")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.500001"))
	if err != nil {
		lang.SayString("subx560")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx560")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.51"))
	if err != nil {
		lang.SayString("subx561")
	} else {
		if !(rexsult.ToString() == "123456789") {
			lang.SayString("subx561")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.6"))
	if err != nil {
		lang.SayString("subx562")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx562")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.9"))
	if err != nil {
		lang.SayString("subx563")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx563")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.99999"))
	if err != nil {
		lang.SayString("subx564")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx564")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0.999999999"))
	if err != nil {
		lang.SayString("subx565")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx565")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("subx566")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx566")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00000001"))
	if err != nil {
		lang.SayString("subx567")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx567")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.00001"))
	if err != nil {
		lang.SayString("subx568")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx568")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("subx569")
	} else {
		if !(rexsult.ToString() == "123456788") {
			lang.SayString("subx569")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx600")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx600")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx601")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx601")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx602")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx602")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx603")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx603")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx604")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx604")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx605")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx605")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx606")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx606")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx607")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx607")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx608")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx608")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx609")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx609")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx610")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx610")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx611")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx611")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx612")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx612")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx613")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx613")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx614")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx614")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx615")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx615")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx616")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx616")
		}
	}
	rexsult, err = lang.RxFromString("1.000000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx617")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx617")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx618")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx618")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx619")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx619")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx620")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx620")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx621")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx621")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx622")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx622")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx623")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx623")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx624")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx624")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx625")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx625")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx626")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx626")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx627")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx627")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx628")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx628")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx629")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx629")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx630")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx630")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx631")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx631")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx632")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx632")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx633")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx633")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx634")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx634")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx635")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx635")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx636")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx636")
		}
	}
	rexsult, err = lang.RxFromString("1.00000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx637")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx637")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx638")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx638")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx639")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx639")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788"))
	if err != nil {
		lang.SayString("subx640")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx640")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788"))
	if err != nil {
		lang.SayString("subx641")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx641")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788"))
	if err != nil {
		lang.SayString("subx642")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx642")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx650")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx650")
		}
	}
	rexsult, err = lang.RxFromString("0.000000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx651")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx651")
		}
	}
	rexsult, err = lang.RxFromString("0.000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx652")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx652")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx653")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx653")
		}
	}
	rexsult, err = lang.RxFromString("0.4").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx654")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx654")
		}
	}
	rexsult, err = lang.RxFromString("0.49").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx655")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx655")
		}
	}
	rexsult, err = lang.RxFromString("0.499999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx656")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx656")
		}
	}
	rexsult, err = lang.RxFromString("0.499999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx657")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx657")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx658")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx658")
		}
	}
	rexsult, err = lang.RxFromString("0.500000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx659")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx659")
		}
	}
	rexsult, err = lang.RxFromString("0.500001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx660")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx660")
		}
	}
	rexsult, err = lang.RxFromString("0.51").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx661")
	} else {
		if !(rexsult.ToString() == "-123456789") {
			lang.SayString("subx661")
		}
	}
	rexsult, err = lang.RxFromString("0.6").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx662")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx662")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx663")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx663")
		}
	}
	rexsult, err = lang.RxFromString("0.99999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx664")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx664")
		}
	}
	rexsult, err = lang.RxFromString("0.999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx665")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx665")
		}
	}
	rexsult, err = lang.RxFromString("1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx666")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx666")
		}
	}
	rexsult, err = lang.RxFromString("1.00000001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx667")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx667")
		}
	}
	rexsult, err = lang.RxFromString("1.00001").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx668")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx668")
		}
	}
	rexsult, err = lang.RxFromString("1.1").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789"))
	if err != nil {
		lang.SayString("subx669")
	} else {
		if !(rexsult.ToString() == "-123456788") {
			lang.SayString("subx669")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.1"))
	if err != nil {
		lang.SayString("subx670")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx670")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.9"))
	if err != nil {
		lang.SayString("subx671")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx671")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.1"))
	if err != nil {
		lang.SayString("subx672")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx672")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.5"))
	if err != nil {
		lang.SayString("subx673")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx673")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.9"))
	if err != nil {
		lang.SayString("subx674")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx674")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.1"))
	if err != nil {
		lang.SayString("subx680")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx680")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.9"))
	if err != nil {
		lang.SayString("subx681")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx681")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.1"))
	if err != nil {
		lang.SayString("subx682")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx682")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.5"))
	if err != nil {
		lang.SayString("subx683")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx683")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.9"))
	if err != nil {
		lang.SayString("subx684")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx684")
		}
	}
	rexsult, err = lang.RxFromString("123456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456787.1"))
	if err != nil {
		lang.SayString("subx685")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx685")
		}
	}
	rexsult, err = lang.RxFromString("123456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456787.9"))
	if err != nil {
		lang.SayString("subx686")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx686")
		}
	}
	rexsult, err = lang.RxFromString("123456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.1"))
	if err != nil {
		lang.SayString("subx687")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx687")
		}
	}
	rexsult, err = lang.RxFromString("123456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.5"))
	if err != nil {
		lang.SayString("subx688")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx688")
		}
	}
	rexsult, err = lang.RxFromString("123456788").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.9"))
	if err != nil {
		lang.SayString("subx689")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx689")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.1"))
	if err != nil {
		lang.SayString("subx690")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("subx690")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456788.9"))
	if err != nil {
		lang.SayString("subx691")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx691")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.1"))
	if err != nil {
		lang.SayString("subx692")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx692")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.5"))
	if err != nil {
		lang.SayString("subx693")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx693")
		}
	}
	rexsult, err = lang.RxFromString("123456789").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("123456789.9"))
	if err != nil {
		lang.SayString("subx694")
	} else {
		if !(rexsult.ToString() == "-1") {
			lang.SayString("subx694")
		}
	}
	rexsult, err = lang.RxFromString("12345678900000").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-9999999999999"))
	if err != nil {
		lang.SayString("subx700")
	} else {
		if !(rexsult.ToString() == "2.23E+13") {
			lang.SayString("subx700")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-12345678900000"))
	if err != nil {
		lang.SayString("subx701")
	} else {
		if !(rexsult.ToString() == "2.23E+13") {
			lang.SayString("subx701")
		}
	}
	rexsult, err = lang.RxFromString("12E+3").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-3456"))
	if err != nil {
		lang.SayString("subx702")
	} else {
		if !(rexsult.ToString() == "1.55E+4") {
			lang.SayString("subx702")
		}
	}
	rexsult, err = lang.RxFromString("12E+3").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-3446"))
	if err != nil {
		lang.SayString("subx703")
	} else {
		if !(rexsult.ToString() == "1.54E+4") {
			lang.SayString("subx703")
		}
	}
	rexsult, err = lang.RxFromString("12E+3").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-3454"))
	if err != nil {
		lang.SayString("subx704")
	} else {
		if !(rexsult.ToString() == "1.55E+4") {
			lang.SayString("subx704")
		}
	}
	rexsult, err = lang.RxFromString("12E+3").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-3444"))
	if err != nil {
		lang.SayString("subx705")
	} else {
		if !(rexsult.ToString() == "1.54E+4") {
			lang.SayString("subx705")
		}
	}
	rexsult, err = lang.RxFromString("3456").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-12E+3"))
	if err != nil {
		lang.SayString("subx706")
	} else {
		if !(rexsult.ToString() == "1.55E+4") {
			lang.SayString("subx706")
		}
	}
	rexsult, err = lang.RxFromString("3446").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-12E+3"))
	if err != nil {
		lang.SayString("subx707")
	} else {
		if !(rexsult.ToString() == "1.54E+4") {
			lang.SayString("subx707")
		}
	}
	rexsult, err = lang.RxFromString("3454").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-12E+3"))
	if err != nil {
		lang.SayString("subx708")
	} else {
		if !(rexsult.ToString() == "1.55E+4") {
			lang.SayString("subx708")
		}
	}
	rexsult, err = lang.RxFromString("3444").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-12E+3"))
	if err != nil {
		lang.SayString("subx709")
	} else {
		if !(rexsult.ToString() == "1.54E+4") {
			lang.SayString("subx709")
		}
	}
	rexsult, err = lang.RxFromString("-1.1E-999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1E-999999999"))
	if err != nil {
		lang.SayString("subx714")
	} else {
		if !(rexsult.ToString() == "-1E-1000000000") {
			lang.SayString("subx714")
		}
	}
	rexsult, err = lang.RxFromString("1E-999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("+1.1e-999999999"))
	if err != nil {
		lang.SayString("subx715")
	} else {
		if !(rexsult.ToString() == "-1E-1000000000") {
			lang.SayString("subx715")
		}
	}
	rexsult, err = lang.RxFromString("+1.1E-999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("+1E-999999999"))
	if err != nil {
		lang.SayString("subx718")
	} else {
		if !(rexsult.ToString() == "1E-1000000000") {
			lang.SayString("subx718")
		}
	}
	rexsult, err = lang.RxFromString("-1E-999999999").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("-1.1e-999999999"))
	if err != nil {
		lang.SayString("subx719")
	} else {
		if !(rexsult.ToString() == "1E-1000000000") {
			lang.SayString("subx719")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub731")
	} else {
		if !(rexsult.ToString() == "1.23456780E+10") {
			lang.SayString("sub731")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("sub732")
	} else {
		if !(rexsult.ToString() == "-1.23456780E+10") {
			lang.SayString("sub732")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub733")
	} else {
		if !(rexsult.ToString() == "1.23456780E+9") {
			lang.SayString("sub733")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("sub734")
	} else {
		if !(rexsult.ToString() == "-1.23456780E+9") {
			lang.SayString("sub734")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub735")
	} else {
		if !(rexsult.ToString() == "1.23456789E+9") {
			lang.SayString("sub735")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("sub736")
	} else {
		if !(rexsult.ToString() == "-1.23456789E+9") {
			lang.SayString("sub736")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub737")
	} else {
		if !(rexsult.ToString() == "1.23456789E+9") {
			lang.SayString("sub737")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("sub738")
	} else {
		if !(rexsult.ToString() == "-1.23456789E+9") {
			lang.SayString("sub738")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub739")
	} else {
		if !(rexsult.ToString() == "1.23456789E+10") {
			lang.SayString("sub739")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("sub740")
	} else {
		if !(rexsult.ToString() == "-1.23456789E+10") {
			lang.SayString("sub740")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub741")
	} else {
		if !(rexsult.ToString() == "1.23456790E+9") {
			lang.SayString("sub741")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("sub742")
	} else {
		if !(rexsult.ToString() == "-1.23456790E+9") {
			lang.SayString("sub742")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub751")
	} else {
		if !(rexsult.ToString() == "12345678000") {
			lang.SayString("sub751")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("sub752")
	} else {
		if !(rexsult.ToString() == "-12345678000") {
			lang.SayString("sub752")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub753")
	} else {
		if !(rexsult.ToString() == "1234567800") {
			lang.SayString("sub753")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("sub754")
	} else {
		if !(rexsult.ToString() == "-1234567800") {
			lang.SayString("sub754")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub755")
	} else {
		if !(rexsult.ToString() == "1234567890") {
			lang.SayString("sub755")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("sub756")
	} else {
		if !(rexsult.ToString() == "-1234567890") {
			lang.SayString("sub756")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub757")
	} else {
		if !(rexsult.ToString() == "1234567891") {
			lang.SayString("sub757")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("sub758")
	} else {
		if !(rexsult.ToString() == "-1234567891") {
			lang.SayString("sub758")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub759")
	} else {
		if !(rexsult.ToString() == "12345678901") {
			lang.SayString("sub759")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("sub760")
	} else {
		if !(rexsult.ToString() == "-12345678901") {
			lang.SayString("sub760")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("sub761")
	} else {
		if !(rexsult.ToString() == "1234567896") {
			lang.SayString("sub761")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("sub762")
	} else {
		if !(rexsult.ToString() == "-1234567896") {
			lang.SayString("sub762")
		}
	}
	rexsult, err = lang.RxFromString("2.E-3").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1.002"))
	if err != nil {
		lang.SayString("subx901")
	} else {
		if !(rexsult.ToString() == "-1.000") {
			lang.SayString("subx901")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-3").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1.002"))
	if err != nil {
		lang.SayString("subx902")
	} else {
		if !(rexsult.ToString() == "-1.0000") {
			lang.SayString("subx902")
		}
	}
	rexsult, err = lang.RxFromString("2.00E-3").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1.0020"))
	if err != nil {
		lang.SayString("subx903")
	} else {
		if !(rexsult.ToString() == "-1.00000") {
			lang.SayString("subx903")
		}
	}
	rexsult, err = lang.RxFromString("2.000E-3").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1.00200"))
	if err != nil {
		lang.SayString("subx904")
	} else {
		if !(rexsult.ToString() == "-1.000000") {
			lang.SayString("subx904")
		}
	}
	rexsult, err = lang.RxFromString("2.0000E-3").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1.002000"))
	if err != nil {
		lang.SayString("subx905")
	} else {
		if !(rexsult.ToString() == "-1.0000000") {
			lang.SayString("subx905")
		}
	}
	rexsult, err = lang.RxFromString("2.00000E-3").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1.0020000"))
	if err != nil {
		lang.SayString("subx906")
	} else {
		if !(rexsult.ToString() == "-1.00000000") {
			lang.SayString("subx906")
		}
	}
	rexsult, err = lang.RxFromString("2.000000E-3").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1.00200000"))
	if err != nil {
		lang.SayString("subx907")
	} else {
		if !(rexsult.ToString() == "-1.000000000") {
			lang.SayString("subx907")
		}
	}
	rexsult, err = lang.RxFromString("2.0000000E-3").OpSub(lang.RxSetWithDigit(15), lang.RxFromString("1.002000000"))
	if err != nil {
		lang.SayString("subx908")
	} else {
		if !(rexsult.ToString() == "-1.0000000000") {
			lang.SayString("subx908")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("1.00E-999"))
	if err != nil {
		lang.SayString("subx1010")
	} else {
		if !(rexsult.ToString() == "-1.00E-999") {
			lang.SayString("subx1010")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1011")
	} else {
		if !(rexsult.ToString() == "-1E-1000") {
			lang.SayString("subx1011")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.10E-999"))
	if err != nil {
		lang.SayString("subx1012")
	} else {
		if !(rexsult.ToString() == "-1.0E-1000") {
			lang.SayString("subx1012")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.100E-999"))
	if err != nil {
		lang.SayString("subx1013")
	} else {
		if !(rexsult.ToString() == "-1.00E-1000") {
			lang.SayString("subx1013")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.01E-999"))
	if err != nil {
		lang.SayString("subx1014")
	} else {
		if !(rexsult.ToString() == "-1E-1001") {
			lang.SayString("subx1014")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.999E-999"))
	if err != nil {
		lang.SayString("subx1015")
	} else {
		if !(rexsult.ToString() == "-9.99E-1000") {
			lang.SayString("subx1015")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.099E-999"))
	if err != nil {
		lang.SayString("subx1016")
	} else {
		if !(rexsult.ToString() == "-9.9E-1001") {
			lang.SayString("subx1016")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.009E-999"))
	if err != nil {
		lang.SayString("subx1017")
	} else {
		if !(rexsult.ToString() == "-9E-1002") {
			lang.SayString("subx1017")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.001E-999"))
	if err != nil {
		lang.SayString("subx1018")
	} else {
		if !(rexsult.ToString() == "-1E-1002") {
			lang.SayString("subx1018")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.0009E-999"))
	if err != nil {
		lang.SayString("subx1019")
	} else {
		if !(rexsult.ToString() == "-9E-1003") {
			lang.SayString("subx1019")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.0001E-999"))
	if err != nil {
		lang.SayString("subx1020")
	} else {
		if !(rexsult.ToString() == "-1E-1003") {
			lang.SayString("subx1020")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-1.00E-999"))
	if err != nil {
		lang.SayString("subx1030")
	} else {
		if !(rexsult.ToString() == "1.00E-999") {
			lang.SayString("subx1030")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.1E-999"))
	if err != nil {
		lang.SayString("subx1031")
	} else {
		if !(rexsult.ToString() == "1E-1000") {
			lang.SayString("subx1031")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.10E-999"))
	if err != nil {
		lang.SayString("subx1032")
	} else {
		if !(rexsult.ToString() == "1.0E-1000") {
			lang.SayString("subx1032")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.100E-999"))
	if err != nil {
		lang.SayString("subx1033")
	} else {
		if !(rexsult.ToString() == "1.00E-1000") {
			lang.SayString("subx1033")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.01E-999"))
	if err != nil {
		lang.SayString("subx1034")
	} else {
		if !(rexsult.ToString() == "1E-1001") {
			lang.SayString("subx1034")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.999E-999"))
	if err != nil {
		lang.SayString("subx1035")
	} else {
		if !(rexsult.ToString() == "9.99E-1000") {
			lang.SayString("subx1035")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.099E-999"))
	if err != nil {
		lang.SayString("subx1036")
	} else {
		if !(rexsult.ToString() == "9.9E-1001") {
			lang.SayString("subx1036")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.009E-999"))
	if err != nil {
		lang.SayString("subx1037")
	} else {
		if !(rexsult.ToString() == "9E-1002") {
			lang.SayString("subx1037")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.001E-999"))
	if err != nil {
		lang.SayString("subx1038")
	} else {
		if !(rexsult.ToString() == "1E-1002") {
			lang.SayString("subx1038")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.0009E-999"))
	if err != nil {
		lang.SayString("subx1039")
	} else {
		if !(rexsult.ToString() == "9E-1003") {
			lang.SayString("subx1039")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("-0.0001E-999"))
	if err != nil {
		lang.SayString("subx1040")
	} else {
		if !(rexsult.ToString() == "1E-1003") {
			lang.SayString("subx1040")
		}
	}
	rexsult, err = lang.RxFromString("1.00E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1050")
	} else {
		if !(rexsult.ToString() == "9.0E-1000") {
			lang.SayString("subx1050")
		}
	}
	rexsult, err = lang.RxFromString("0.1E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1051")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx1051")
		}
	}
	rexsult, err = lang.RxFromString("0.10E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1052")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx1052")
		}
	}
	rexsult, err = lang.RxFromString("0.100E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1053")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx1053")
		}
	}
	rexsult, err = lang.RxFromString("0.01E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1054")
	} else {
		if !(rexsult.ToString() == "-9E-1001") {
			lang.SayString("subx1054")
		}
	}
	rexsult, err = lang.RxFromString("0.999E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1055")
	} else {
		if !(rexsult.ToString() == "8.99E-1000") {
			lang.SayString("subx1055")
		}
	}
	rexsult, err = lang.RxFromString("0.099E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1056")
	} else {
		if !(rexsult.ToString() == "-1E-1002") {
			lang.SayString("subx1056")
		}
	}
	rexsult, err = lang.RxFromString("0.009E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1057")
	} else {
		if !(rexsult.ToString() == "-9.1E-1001") {
			lang.SayString("subx1057")
		}
	}
	rexsult, err = lang.RxFromString("0.001E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1058")
	} else {
		if !(rexsult.ToString() == "-9.9E-1001") {
			lang.SayString("subx1058")
		}
	}
	rexsult, err = lang.RxFromString("0.0009E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1059")
	} else {
		if !(rexsult.ToString() == "-9.9E-1001") {
			lang.SayString("subx1059")
		}
	}
	rexsult, err = lang.RxFromString("0.0001E-999").OpSub(lang.RxSetWithDigit(3), lang.RxFromString("0.1E-999"))
	if err != nil {
		lang.SayString("subx1060")
	} else {
		if !(rexsult.ToString() == "-1.00E-1000") {
			lang.SayString("subx1060")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.52444E-80"))
	if err != nil {
		lang.SayString("subx1101")
	} else {
		if !(rexsult.ToString() == "-1.5244E-80") {
			lang.SayString("subx1101")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.52445E-80"))
	if err != nil {
		lang.SayString("subx1102")
	} else {
		if !(rexsult.ToString() == "-1.5245E-80") {
			lang.SayString("subx1102")
		}
	}
	rexsult, err = lang.RxFromString("0").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.52446E-80"))
	if err != nil {
		lang.SayString("subx1103")
	} else {
		if !(rexsult.ToString() == "-1.5245E-80") {
			lang.SayString("subx1103")
		}
	}
	rexsult, err = lang.RxFromString("1.52444E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx1104")
	} else {
		if !(rexsult.ToString() == "1.5244E-80") {
			lang.SayString("subx1104")
		}
	}
	rexsult, err = lang.RxFromString("1.52445E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx1105")
	} else {
		if !(rexsult.ToString() == "1.5245E-80") {
			lang.SayString("subx1105")
		}
	}
	rexsult, err = lang.RxFromString("1.52446E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("subx1106")
	} else {
		if !(rexsult.ToString() == "1.5245E-80") {
			lang.SayString("subx1106")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.2345671E-80"))
	if err != nil {
		lang.SayString("subx1111")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx1111")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.2345618E-80"))
	if err != nil {
		lang.SayString("subx1112")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("subx1112")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.2345178E-80"))
	if err != nil {
		lang.SayString("subx1113")
	} else {
		if !(rexsult.ToString() == "1E-84") {
			lang.SayString("subx1113")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.2341678E-80"))
	if err != nil {
		lang.SayString("subx1114")
	} else {
		if !(rexsult.ToString() == "4E-84") {
			lang.SayString("subx1114")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.2315678E-80"))
	if err != nil {
		lang.SayString("subx1115")
	} else {
		if !(rexsult.ToString() == "3.0E-83") {
			lang.SayString("subx1115")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.2145678E-80"))
	if err != nil {
		lang.SayString("subx1116")
	} else {
		if !(rexsult.ToString() == "2.00E-82") {
			lang.SayString("subx1116")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("1.1345678E-80"))
	if err != nil {
		lang.SayString("subx1117")
	} else {
		if !(rexsult.ToString() == "1.000E-81") {
			lang.SayString("subx1117")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678E-80").OpSub(lang.RxSetWithDigit(5), lang.RxFromString("0.2345678E-80"))
	if err != nil {
		lang.SayString("subx1118")
	} else {
		if !(rexsult.ToString() == "1.0000E-80") {
			lang.SayString("subx1118")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("120E-2"))
	if err != nil {
		lang.SayString("subx1125")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("subx1125")
		}
	}
	rexsult, err = lang.RxFromString("130E-2").OpSub(lang.RxSetWithDigit(34), lang.RxFromString("12E-1"))
	if err != nil {
		lang.SayString("subx1126")
	} else {
		if !(rexsult.ToString() == "0.10") {
			lang.SayString("subx1126")
		}
	}

	return
}
