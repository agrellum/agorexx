package main

import "agorexx/lang"

func main() {

	rexsult, err := lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv001")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv002")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv003")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv007")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv008")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv009")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv010")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dddiv010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddiv011")
	} else {
		if !(rexsult.ToString() == "-2.4") {
			lang.SayString("dddiv011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv012")
	} else {
		if !(rexsult.ToString() == "-2.4") {
			lang.SayString("dddiv012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddiv013")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dddiv013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv014")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dddiv014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv015")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dddiv015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv016")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dddiv016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv017")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dddiv017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv018")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dddiv019")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("dddiv020")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv021")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv021")
		}
	}
	rexsult, err = lang.RxFromString("50").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dddiv022")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv022")
		}
	}
	rexsult, err = lang.RxFromString("500").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("200"))
	if err != nil {
		lang.SayString("dddiv023")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv023")
		}
	}
	rexsult, err = lang.RxFromString("50.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("20.0"))
	if err != nil {
		lang.SayString("dddiv024")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv024")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dddiv025")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv025")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dddiv026")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv026")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dddiv027")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv027")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("0.20"))
	if err != nil {
		lang.SayString("dddiv028")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dddiv028")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dddiv029")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dddiv029")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv030")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dddiv030")
		}
	}
	rexsult, err = lang.RxFromString("100").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv031")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("dddiv031")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv032")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("dddiv032")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddiv033")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dddiv033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv035")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dddiv036")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dddiv036")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dddiv037")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("dddiv037")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dddiv038")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("dddiv038")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dddiv039")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("dddiv039")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dddiv040")
	} else {
		if !(rexsult.ToString() == "0.015625") {
			lang.SayString("dddiv040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dddiv041")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dddiv041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dddiv042")
	} else {
		if !(rexsult.ToString() == "-0.25") {
			lang.SayString("dddiv042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dddiv043")
	} else {
		if !(rexsult.ToString() == "-0.125") {
			lang.SayString("dddiv043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dddiv044")
	} else {
		if !(rexsult.ToString() == "-0.0625") {
			lang.SayString("dddiv044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dddiv045")
	} else {
		if !(rexsult.ToString() == "-0.03125") {
			lang.SayString("dddiv045")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dddiv046")
	} else {
		if !(rexsult.ToString() == "-0.015625") {
			lang.SayString("dddiv046")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv047")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dddiv047")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dddiv048")
	} else {
		if !(rexsult.ToString() == "-0.25") {
			lang.SayString("dddiv048")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dddiv049")
	} else {
		if !(rexsult.ToString() == "-0.125") {
			lang.SayString("dddiv049")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dddiv050")
	} else {
		if !(rexsult.ToString() == "-0.0625") {
			lang.SayString("dddiv050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dddiv051")
	} else {
		if !(rexsult.ToString() == "-0.03125") {
			lang.SayString("dddiv051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dddiv052")
	} else {
		if !(rexsult.ToString() == "-0.015625") {
			lang.SayString("dddiv052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dddiv053")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dddiv054")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dddiv054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dddiv055")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("dddiv055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dddiv056")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("dddiv056")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dddiv057")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("dddiv057")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dddiv058")
	} else {
		if !(rexsult.ToString() == "0.015625") {
			lang.SayString("dddiv058")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dddiv060")
	} else {
		if !(rexsult.ToString() == "0.1428571428571429") {
			lang.SayString("dddiv060")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.9876543"))
	if err != nil {
		lang.SayString("dddiv061")
	} else {
		if !(rexsult.ToString() == "0.6211179680490717") {
			lang.SayString("dddiv061")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv071")
	} else {
		if !(rexsult.ToString() == "9999999999999999") {
			lang.SayString("dddiv071")
		}
	}
	rexsult, err = lang.RxFromString("999999999999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv072")
	} else {
		if !(rexsult.ToString() == "999999999999999") {
			lang.SayString("dddiv072")
		}
	}
	rexsult, err = lang.RxFromString("99999999999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv073")
	} else {
		if !(rexsult.ToString() == "99999999999999") {
			lang.SayString("dddiv073")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv074")
	} else {
		if !(rexsult.ToString() == "9999999999999") {
			lang.SayString("dddiv074")
		}
	}
	rexsult, err = lang.RxFromString("999999999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv075")
	} else {
		if !(rexsult.ToString() == "999999999999") {
			lang.SayString("dddiv075")
		}
	}
	rexsult, err = lang.RxFromString("99999999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv076")
	} else {
		if !(rexsult.ToString() == "99999999999") {
			lang.SayString("dddiv076")
		}
	}
	rexsult, err = lang.RxFromString("9999999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv077")
	} else {
		if !(rexsult.ToString() == "9999999999") {
			lang.SayString("dddiv077")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv078")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dddiv078")
		}
	}
	rexsult, err = lang.RxFromString("99999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv079")
	} else {
		if !(rexsult.ToString() == "99999999") {
			lang.SayString("dddiv079")
		}
	}
	rexsult, err = lang.RxFromString("9999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv080")
	} else {
		if !(rexsult.ToString() == "9999999") {
			lang.SayString("dddiv080")
		}
	}
	rexsult, err = lang.RxFromString("999999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv081")
	} else {
		if !(rexsult.ToString() == "999999") {
			lang.SayString("dddiv081")
		}
	}
	rexsult, err = lang.RxFromString("99999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv082")
	} else {
		if !(rexsult.ToString() == "99999") {
			lang.SayString("dddiv082")
		}
	}
	rexsult, err = lang.RxFromString("9999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv083")
	} else {
		if !(rexsult.ToString() == "9999") {
			lang.SayString("dddiv083")
		}
	}
	rexsult, err = lang.RxFromString("999").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv084")
	} else {
		if !(rexsult.ToString() == "999") {
			lang.SayString("dddiv084")
		}
	}
	rexsult, err = lang.RxFromString("99").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv085")
	} else {
		if !(rexsult.ToString() == "99") {
			lang.SayString("dddiv085")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv086")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dddiv086")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv090")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv093")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv094")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1E-8"))
	if err != nil {
		lang.SayString("dddiv095")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("dddiv095")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1E-9"))
	if err != nil {
		lang.SayString("dddiv096")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("dddiv096")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1E-10"))
	if err != nil {
		lang.SayString("dddiv097")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("dddiv097")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1E-11"))
	if err != nil {
		lang.SayString("dddiv098")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("dddiv098")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1E-12"))
	if err != nil {
		lang.SayString("dddiv099")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("dddiv099")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv100")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv100")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv101")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv101")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv102")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv102")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dddiv103")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dddiv103")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dddiv104")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("dddiv104")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("dddiv105")
	} else {
		if !(rexsult.ToString() == "0.1666666666666667") {
			lang.SayString("dddiv105")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dddiv106")
	} else {
		if !(rexsult.ToString() == "0.1428571428571429") {
			lang.SayString("dddiv106")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dddiv107")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("dddiv107")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dddiv108")
	} else {
		if !(rexsult.ToString() == "0.1111111111111111") {
			lang.SayString("dddiv108")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddiv109")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dddiv109")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv110")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv110")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv111")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv111")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv112")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dddiv112")
		}
	}
	rexsult, err = lang.RxFromString("4").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv113")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("dddiv113")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv114")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("dddiv114")
		}
	}
	rexsult, err = lang.RxFromString("6").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv115")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dddiv115")
		}
	}
	rexsult, err = lang.RxFromString("7").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv116")
	} else {
		if !(rexsult.ToString() == "7") {
			lang.SayString("dddiv116")
		}
	}
	rexsult, err = lang.RxFromString("8").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv117")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("dddiv117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv118")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dddiv118")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv119")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dddiv119")
		}
	}
	rexsult, err = lang.RxFromString("3E+1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("0.001"))
	if err != nil {
		lang.SayString("dddiv120")
	} else {
		if !(rexsult.ToString() == "30000") {
			lang.SayString("dddiv120")
		}
	}
	rexsult, err = lang.RxFromString("2.200").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv121")
	} else {
		if !(rexsult.ToString() == "1.1") {
			lang.SayString("dddiv121")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("dddiv130")
	} else {
		if !(rexsult.ToString() == "2469.493898779756") {
			lang.SayString("dddiv130")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("dddiv131")
	} else {
		if !(rexsult.ToString() == "2473.947895791583") {
			lang.SayString("dddiv131")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("dddiv132")
	} else {
		if !(rexsult.ToString() == "2519.387755102041") {
			lang.SayString("dddiv132")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dddiv133")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("dddiv133")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("dddiv134")
	} else {
		if !(rexsult.ToString() == "2420.588235294118") {
			lang.SayString("dddiv134")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("dddiv135")
	} else {
		if !(rexsult.ToString() == "2464.071856287425") {
			lang.SayString("dddiv135")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("dddiv136")
	} else {
		if !(rexsult.ToString() == "2468.506298740252") {
			lang.SayString("dddiv136")
		}
	}
	rexsult, err = lang.RxFromString("391").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("597"))
	if err != nil {
		lang.SayString("dddiv220")
	} else {
		if !(rexsult.ToString() == "0.6549413735343384") {
			lang.SayString("dddiv220")
		}
	}
	rexsult, err = lang.RxFromString("391").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-597"))
	if err != nil {
		lang.SayString("dddiv221")
	} else {
		if !(rexsult.ToString() == "-0.6549413735343384") {
			lang.SayString("dddiv221")
		}
	}
	rexsult, err = lang.RxFromString("-391").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("597"))
	if err != nil {
		lang.SayString("dddiv222")
	} else {
		if !(rexsult.ToString() == "-0.6549413735343384") {
			lang.SayString("dddiv222")
		}
	}
	rexsult, err = lang.RxFromString("-391").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-597"))
	if err != nil {
		lang.SayString("dddiv223")
	} else {
		if !(rexsult.ToString() == "0.6549413735343384") {
			lang.SayString("dddiv223")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dddiv301")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv301")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dddiv302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv302")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dddiv303")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv303")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dddiv304")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv304")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dddiv305")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv305")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+6"))
	if err != nil {
		lang.SayString("dddiv306")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv306")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+7"))
	if err != nil {
		lang.SayString("dddiv307")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv307")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("70E-5"))
	if err != nil {
		lang.SayString("dddiv308")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv308")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("70E-1"))
	if err != nil {
		lang.SayString("dddiv309")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv309")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("70E+0"))
	if err != nil {
		lang.SayString("dddiv310")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv310")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("70E+1"))
	if err != nil {
		lang.SayString("dddiv311")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv311")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("70E+5"))
	if err != nil {
		lang.SayString("dddiv312")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv312")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("70E+6"))
	if err != nil {
		lang.SayString("dddiv313")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv313")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("70E+7"))
	if err != nil {
		lang.SayString("dddiv314")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv314")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("700E-5"))
	if err != nil {
		lang.SayString("dddiv315")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv315")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("700E-1"))
	if err != nil {
		lang.SayString("dddiv316")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv316")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("700E+0"))
	if err != nil {
		lang.SayString("dddiv317")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv317")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("700E+1"))
	if err != nil {
		lang.SayString("dddiv318")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv318")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("700E+5"))
	if err != nil {
		lang.SayString("dddiv319")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv319")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("700E+6"))
	if err != nil {
		lang.SayString("dddiv320")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv320")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("700E+7"))
	if err != nil {
		lang.SayString("dddiv321")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv321")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("700E+77"))
	if err != nil {
		lang.SayString("dddiv322")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv322")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dddiv331")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv331")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dddiv332")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv332")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dddiv333")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv333")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dddiv334")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv334")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dddiv335")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv335")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dddiv336")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv336")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dddiv337")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv337")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dddiv338")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv338")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dddiv339")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv339")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dddiv340")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv340")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dddiv341")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv341")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dddiv342")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv342")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dddiv343")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv343")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dddiv344")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv344")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dddiv345")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv345")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dddiv346")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv346")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv441")
	} else {
		if !(rexsult.ToString() == "12345678000") {
			lang.SayString("dddiv441")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dddiv442")
	} else {
		if !(rexsult.ToString() == "8.100000664200054E-11") {
			lang.SayString("dddiv442")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv443")
	} else {
		if !(rexsult.ToString() == "1234567800") {
			lang.SayString("dddiv443")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dddiv444")
	} else {
		if !(rexsult.ToString() == "8.100000664200054E-10") {
			lang.SayString("dddiv444")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv445")
	} else {
		if !(rexsult.ToString() == "1234567890") {
			lang.SayString("dddiv445")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dddiv446")
	} else {
		if !(rexsult.ToString() == "8.100000073710001E-10") {
			lang.SayString("dddiv446")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv447")
	} else {
		if !(rexsult.ToString() == "1234567891") {
			lang.SayString("dddiv447")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dddiv448")
	} else {
		if !(rexsult.ToString() == "8.100000067149001E-10") {
			lang.SayString("dddiv448")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv449")
	} else {
		if !(rexsult.ToString() == "12345678901") {
			lang.SayString("dddiv449")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dddiv450")
	} else {
		if !(rexsult.ToString() == "8.100000073053901E-11") {
			lang.SayString("dddiv450")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv451")
	} else {
		if !(rexsult.ToString() == "1234567896") {
			lang.SayString("dddiv451")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dddiv452")
	} else {
		if !(rexsult.ToString() == "8.100000034344E-10") {
			lang.SayString("dddiv452")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv453")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dddiv453")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddiv454")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dddiv454")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dddiv455")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dddiv455")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv456")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("dddiv456")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dddiv457")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("dddiv457")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dddiv458")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("dddiv458")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("dddiv466")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv466")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("dddiv467")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv467")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("dddiv468")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv468")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("dddiv470")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv470")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("dddiv471")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv471")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("dddiv472")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv472")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("dddiv474")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv474")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("dddiv475")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv475")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("dddiv476")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dddiv476")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0E+33"))
	if err != nil {
		lang.SayString("dddiv480")
	} else {
		if !(rexsult.ToString() == "1E-33") {
			lang.SayString("dddiv480")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10E+33"))
	if err != nil {
		lang.SayString("dddiv481")
	} else {
		if !(rexsult.ToString() == "1E-34") {
			lang.SayString("dddiv481")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0E-33"))
	if err != nil {
		lang.SayString("dddiv482")
	} else {
		if !(rexsult.ToString() == "1E+33") {
			lang.SayString("dddiv482")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10E-33"))
	if err != nil {
		lang.SayString("dddiv483")
	} else {
		if !(rexsult.ToString() == "1E+32") {
			lang.SayString("dddiv483")
		}
	}
	rexsult, err = lang.RxFromString("0E+380").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1000E-13"))
	if err != nil {
		lang.SayString("dddiv497")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv497")
		}
	}
	rexsult, err = lang.RxFromString("0E-390").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1000E+13"))
	if err != nil {
		lang.SayString("dddiv498")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv498")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9.9"))
	if err != nil {
		lang.SayString("dddiv500")
	} else {
		if !(rexsult.ToString() == "0.101010101010101") {
			lang.SayString("dddiv500")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9.09"))
	if err != nil {
		lang.SayString("dddiv501")
	} else {
		if !(rexsult.ToString() == "0.11001100110011") {
			lang.SayString("dddiv501")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9.009"))
	if err != nil {
		lang.SayString("dddiv502")
	} else {
		if !(rexsult.ToString() == "0.111000111000111") {
			lang.SayString("dddiv502")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv511")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv511")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv512")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv512")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv513")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv513")
		}
	}
	rexsult, err = lang.RxFromString("1.000").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv514")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv514")
		}
	}
	rexsult, err = lang.RxFromString("1.0000").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv515")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv515")
		}
	}
	rexsult, err = lang.RxFromString("1.00000").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv516")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv516")
		}
	}
	rexsult, err = lang.RxFromString("1.000000").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv517")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv517")
		}
	}
	rexsult, err = lang.RxFromString("1.0000000").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv518")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv518")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dddiv519")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dddiv519")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv521")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv521")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddiv522")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv522")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dddiv523")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv523")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.000"))
	if err != nil {
		lang.SayString("dddiv524")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv524")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0000"))
	if err != nil {
		lang.SayString("dddiv525")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv525")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.00000"))
	if err != nil {
		lang.SayString("dddiv526")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv526")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.000000"))
	if err != nil {
		lang.SayString("dddiv527")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv527")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0000000"))
	if err != nil {
		lang.SayString("dddiv528")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv528")
		}
	}
	rexsult, err = lang.RxFromString("2.00").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dddiv529")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dddiv529")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv530")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dddiv530")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dddiv531")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("dddiv531")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddiv532")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("dddiv532")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dddiv533")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dddiv533")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4.0"))
	if err != nil {
		lang.SayString("dddiv534")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("dddiv534")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10.0"))
	if err != nil {
		lang.SayString("dddiv535")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("dddiv535")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dddiv536")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dddiv536")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4.00"))
	if err != nil {
		lang.SayString("dddiv537")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("dddiv537")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10.00"))
	if err != nil {
		lang.SayString("dddiv538")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("dddiv538")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dddiv539")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dddiv539")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("0.01"))
	if err != nil {
		lang.SayString("dddiv540")
	} else {
		if !(rexsult.ToString() == "90") {
			lang.SayString("dddiv540")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("0.001"))
	if err != nil {
		lang.SayString("dddiv541")
	} else {
		if !(rexsult.ToString() == "900") {
			lang.SayString("dddiv541")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv542")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv542")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dddiv543")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv543")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dddiv544")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dddiv544")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dddiv545")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dddiv545")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("20.0"))
	if err != nil {
		lang.SayString("dddiv546")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dddiv546")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv547")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dddiv547")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dddiv548")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dddiv548")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.400"))
	if err != nil {
		lang.SayString("dddiv549")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv549")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv550")
	} else {
		if !(rexsult.ToString() == "240") {
			lang.SayString("dddiv550")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddiv551")
	} else {
		if !(rexsult.ToString() == "24") {
			lang.SayString("dddiv551")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddiv552")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dddiv552")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dddiv553")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("dddiv553")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv554")
	} else {
		if !(rexsult.ToString() == "2400") {
			lang.SayString("dddiv554")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dddiv555")
	} else {
		if !(rexsult.ToString() == "240") {
			lang.SayString("dddiv555")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dddiv556")
	} else {
		if !(rexsult.ToString() == "24") {
			lang.SayString("dddiv556")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dddiv557")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dddiv557")
		}
	}
	rexsult, err = lang.RxFromString("2.4E+9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv600")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dddiv600")
		}
	}
	rexsult, err = lang.RxFromString("2.40E+9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv601")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dddiv601")
		}
	}
	rexsult, err = lang.RxFromString("2.400E+9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv602")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dddiv602")
		}
	}
	rexsult, err = lang.RxFromString("2.4000E+9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv603")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dddiv603")
		}
	}
	rexsult, err = lang.RxFromString("24E+8").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv604")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dddiv604")
		}
	}
	rexsult, err = lang.RxFromString("240E+7").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv605")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dddiv605")
		}
	}
	rexsult, err = lang.RxFromString("2400E+6").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv606")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dddiv606")
		}
	}
	rexsult, err = lang.RxFromString("24000E+5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dddiv607")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dddiv607")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("dddiv731")
	} else {
		if !(rexsult.ToString() == "5000") {
			lang.SayString("dddiv731")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddiv741")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv741")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddiv742")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv742")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv743")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv743")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv744")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv744")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddiv751")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv751")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dddiv752")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv752")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv753")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv753")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv754")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv754")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dddiv761")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv761")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dddiv762")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv762")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddiv763")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv763")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddiv764")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv764")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dddiv771")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv771")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dddiv772")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv772")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddiv773")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv773")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dddiv774")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dddiv774")
		}
	}
	rexsult, err = lang.RxFromString("100E-2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1000E-3"))
	if err != nil {
		lang.SayString("dddiv1024")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dddiv1024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dddiv1040")
	} else {
		if !(rexsult.ToString() == "0.5555555555555556") {
			lang.SayString("dddiv1040")
		}
	}
	rexsult, err = lang.RxFromString("6").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("11"))
	if err != nil {
		lang.SayString("dddiv1041")
	} else {
		if !(rexsult.ToString() == "0.5454545454545455") {
			lang.SayString("dddiv1041")
		}
	}
	rexsult, err = lang.RxFromString("1e+277").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-311"))
	if err != nil {
		lang.SayString("dddiv1051")
	} else {
		if !(rexsult.ToString() == "1E+588") {
			lang.SayString("dddiv1051")
		}
	}
	rexsult, err = lang.RxFromString("1e+277").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e-311"))
	if err != nil {
		lang.SayString("dddiv1052")
	} else {
		if !(rexsult.ToString() == "-1E+588") {
			lang.SayString("dddiv1052")
		}
	}
	rexsult, err = lang.RxFromString("-1e+277").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-311"))
	if err != nil {
		lang.SayString("dddiv1053")
	} else {
		if !(rexsult.ToString() == "-1E+588") {
			lang.SayString("dddiv1053")
		}
	}
	rexsult, err = lang.RxFromString("-1e+277").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e-311"))
	if err != nil {
		lang.SayString("dddiv1054")
	} else {
		if !(rexsult.ToString() == "1E+588") {
			lang.SayString("dddiv1054")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("dddiv1055")
	} else {
		if !(rexsult.ToString() == "1E-588") {
			lang.SayString("dddiv1055")
		}
	}
	rexsult, err = lang.RxFromString("1e-277").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("dddiv1056")
	} else {
		if !(rexsult.ToString() == "-1E-588") {
			lang.SayString("dddiv1056")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+311"))
	if err != nil {
		lang.SayString("dddiv1057")
	} else {
		if !(rexsult.ToString() == "-1E-588") {
			lang.SayString("dddiv1057")
		}
	}
	rexsult, err = lang.RxFromString("-1e-277").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e+311"))
	if err != nil {
		lang.SayString("dddiv1058")
	} else {
		if !(rexsult.ToString() == "1E-588") {
			lang.SayString("dddiv1058")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+101"))
	if err != nil {
		lang.SayString("dddiv1060")
	} else {
		if !(rexsult.ToString() == "1E-392") {
			lang.SayString("dddiv1060")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+102"))
	if err != nil {
		lang.SayString("dddiv1061")
	} else {
		if !(rexsult.ToString() == "1E-393") {
			lang.SayString("dddiv1061")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+103"))
	if err != nil {
		lang.SayString("dddiv1062")
	} else {
		if !(rexsult.ToString() == "1E-394") {
			lang.SayString("dddiv1062")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+104"))
	if err != nil {
		lang.SayString("dddiv1063")
	} else {
		if !(rexsult.ToString() == "1E-395") {
			lang.SayString("dddiv1063")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+105"))
	if err != nil {
		lang.SayString("dddiv1064")
	} else {
		if !(rexsult.ToString() == "1E-396") {
			lang.SayString("dddiv1064")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+106"))
	if err != nil {
		lang.SayString("dddiv1065")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("dddiv1065")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+107"))
	if err != nil {
		lang.SayString("dddiv1066")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("dddiv1066")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+108"))
	if err != nil {
		lang.SayString("dddiv1067")
	} else {
		if !(rexsult.ToString() == "1E-399") {
			lang.SayString("dddiv1067")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+109"))
	if err != nil {
		lang.SayString("dddiv1068")
	} else {
		if !(rexsult.ToString() == "1E-400") {
			lang.SayString("dddiv1068")
		}
	}
	rexsult, err = lang.RxFromString("1e-291").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+110"))
	if err != nil {
		lang.SayString("dddiv1069")
	} else {
		if !(rexsult.ToString() == "1E-401") {
			lang.SayString("dddiv1069")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-321"))
	if err != nil {
		lang.SayString("dddiv1070")
	} else {
		if !(rexsult.ToString() == "1E+381") {
			lang.SayString("dddiv1070")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-322"))
	if err != nil {
		lang.SayString("dddiv1071")
	} else {
		if !(rexsult.ToString() == "1E+382") {
			lang.SayString("dddiv1071")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-323"))
	if err != nil {
		lang.SayString("dddiv1072")
	} else {
		if !(rexsult.ToString() == "1E+383") {
			lang.SayString("dddiv1072")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-324"))
	if err != nil {
		lang.SayString("dddiv1073")
	} else {
		if !(rexsult.ToString() == "1E+384") {
			lang.SayString("dddiv1073")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-325"))
	if err != nil {
		lang.SayString("dddiv1074")
	} else {
		if !(rexsult.ToString() == "1E+385") {
			lang.SayString("dddiv1074")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-326"))
	if err != nil {
		lang.SayString("dddiv1075")
	} else {
		if !(rexsult.ToString() == "1E+386") {
			lang.SayString("dddiv1075")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-327"))
	if err != nil {
		lang.SayString("dddiv1076")
	} else {
		if !(rexsult.ToString() == "1E+387") {
			lang.SayString("dddiv1076")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-328"))
	if err != nil {
		lang.SayString("dddiv1077")
	} else {
		if !(rexsult.ToString() == "1E+388") {
			lang.SayString("dddiv1077")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-329"))
	if err != nil {
		lang.SayString("dddiv1078")
	} else {
		if !(rexsult.ToString() == "1E+389") {
			lang.SayString("dddiv1078")
		}
	}
	rexsult, err = lang.RxFromString("1e+60").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e-330"))
	if err != nil {
		lang.SayString("dddiv1079")
	} else {
		if !(rexsult.ToString() == "1E+390") {
			lang.SayString("dddiv1079")
		}
	}
	rexsult, err = lang.RxFromString("1.0000E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dddiv1101")
	} else {
		if !(rexsult.ToString() == "1E-394") {
			lang.SayString("dddiv1101")
		}
	}
	rexsult, err = lang.RxFromString("1.000E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+1"))
	if err != nil {
		lang.SayString("dddiv1102")
	} else {
		if !(rexsult.ToString() == "1E-395") {
			lang.SayString("dddiv1102")
		}
	}
	rexsult, err = lang.RxFromString("1.00E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+2"))
	if err != nil {
		lang.SayString("dddiv1103")
	} else {
		if !(rexsult.ToString() == "1E-396") {
			lang.SayString("dddiv1103")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+3"))
	if err != nil {
		lang.SayString("dddiv1104")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("dddiv1104")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1105")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("dddiv1105")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1106")
	} else {
		if !(rexsult.ToString() == "1.3E-398") {
			lang.SayString("dddiv1106")
		}
	}
	rexsult, err = lang.RxFromString("1.5E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1107")
	} else {
		if !(rexsult.ToString() == "1.5E-398") {
			lang.SayString("dddiv1107")
		}
	}
	rexsult, err = lang.RxFromString("1.7E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1108")
	} else {
		if !(rexsult.ToString() == "1.7E-398") {
			lang.SayString("dddiv1108")
		}
	}
	rexsult, err = lang.RxFromString("2.3E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1109")
	} else {
		if !(rexsult.ToString() == "2.3E-398") {
			lang.SayString("dddiv1109")
		}
	}
	rexsult, err = lang.RxFromString("2.5E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1110")
	} else {
		if !(rexsult.ToString() == "2.5E-398") {
			lang.SayString("dddiv1110")
		}
	}
	rexsult, err = lang.RxFromString("2.7E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1111")
	} else {
		if !(rexsult.ToString() == "2.7E-398") {
			lang.SayString("dddiv1111")
		}
	}
	rexsult, err = lang.RxFromString("1.49E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1112")
	} else {
		if !(rexsult.ToString() == "1.49E-398") {
			lang.SayString("dddiv1112")
		}
	}
	rexsult, err = lang.RxFromString("1.50E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1113")
	} else {
		if !(rexsult.ToString() == "1.5E-398") {
			lang.SayString("dddiv1113")
		}
	}
	rexsult, err = lang.RxFromString("1.51E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1114")
	} else {
		if !(rexsult.ToString() == "1.51E-398") {
			lang.SayString("dddiv1114")
		}
	}
	rexsult, err = lang.RxFromString("2.49E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1115")
	} else {
		if !(rexsult.ToString() == "2.49E-398") {
			lang.SayString("dddiv1115")
		}
	}
	rexsult, err = lang.RxFromString("2.50E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1116")
	} else {
		if !(rexsult.ToString() == "2.5E-398") {
			lang.SayString("dddiv1116")
		}
	}
	rexsult, err = lang.RxFromString("2.51E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1117")
	} else {
		if !(rexsult.ToString() == "2.51E-398") {
			lang.SayString("dddiv1117")
		}
	}
	rexsult, err = lang.RxFromString("1E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dddiv1118")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("dddiv1118")
		}
	}
	rexsult, err = lang.RxFromString("3E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dddiv1119")
	} else {
		if !(rexsult.ToString() == "3E-399") {
			lang.SayString("dddiv1119")
		}
	}
	rexsult, err = lang.RxFromString("5E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dddiv1120")
	} else {
		if !(rexsult.ToString() == "5E-399") {
			lang.SayString("dddiv1120")
		}
	}
	rexsult, err = lang.RxFromString("7E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dddiv1121")
	} else {
		if !(rexsult.ToString() == "7E-399") {
			lang.SayString("dddiv1121")
		}
	}
	rexsult, err = lang.RxFromString("9E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dddiv1122")
	} else {
		if !(rexsult.ToString() == "9E-399") {
			lang.SayString("dddiv1122")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dddiv1123")
	} else {
		if !(rexsult.ToString() == "9.9E-399") {
			lang.SayString("dddiv1123")
		}
	}
	rexsult, err = lang.RxFromString("1E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e+4"))
	if err != nil {
		lang.SayString("dddiv1124")
	} else {
		if !(rexsult.ToString() == "-1E-398") {
			lang.SayString("dddiv1124")
		}
	}
	rexsult, err = lang.RxFromString("3E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dddiv1125")
	} else {
		if !(rexsult.ToString() == "-3E-399") {
			lang.SayString("dddiv1125")
		}
	}
	rexsult, err = lang.RxFromString("-5E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dddiv1126")
	} else {
		if !(rexsult.ToString() == "-5E-399") {
			lang.SayString("dddiv1126")
		}
	}
	rexsult, err = lang.RxFromString("7E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dddiv1127")
	} else {
		if !(rexsult.ToString() == "-7E-399") {
			lang.SayString("dddiv1127")
		}
	}
	rexsult, err = lang.RxFromString("-9E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dddiv1128")
	} else {
		if !(rexsult.ToString() == "-9E-399") {
			lang.SayString("dddiv1128")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dddiv1129")
	} else {
		if !(rexsult.ToString() == "-9.9E-399") {
			lang.SayString("dddiv1129")
		}
	}
	rexsult, err = lang.RxFromString("3.0E-394").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dddiv1130")
	} else {
		if !(rexsult.ToString() == "-3E-399") {
			lang.SayString("dddiv1130")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+200"))
	if err != nil {
		lang.SayString("dddiv1131")
	} else {
		if !(rexsult.ToString() == "1E-399") {
			lang.SayString("dddiv1131")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+199"))
	if err != nil {
		lang.SayString("dddiv1132")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("dddiv1132")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-199").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1e+198"))
	if err != nil {
		lang.SayString("dddiv1133")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("dddiv1133")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-199").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2e+198"))
	if err != nil {
		lang.SayString("dddiv1134")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("dddiv1134")
		}
	}
	rexsult, err = lang.RxFromString("4.0E-199").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4e+198"))
	if err != nil {
		lang.SayString("dddiv1135")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("dddiv1135")
		}
	}
	rexsult, err = lang.RxFromString("10.0E-199").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("10e+198"))
	if err != nil {
		lang.SayString("dddiv1136")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("dddiv1136")
		}
	}
	rexsult, err = lang.RxFromString("30.0E-199").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("30e+198"))
	if err != nil {
		lang.SayString("dddiv1137")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("dddiv1137")
		}
	}
	rexsult, err = lang.RxFromString("-3.303226714900711E-35").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8.796578842713183E+73"))
	if err != nil {
		lang.SayString("dddiv2010")
	} else {
		if !(rexsult.ToString() == "-3.755126594058783E-109") {
			lang.SayString("dddiv2010")
		}
	}
	rexsult, err = lang.RxFromString("933153327821073.6").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("68782181090246.25"))
	if err != nil {
		lang.SayString("dddiv2011")
	} else {
		if !(rexsult.ToString() == "13.56678885475763") {
			lang.SayString("dddiv2011")
		}
	}
	rexsult, err = lang.RxFromString("5.04752436057906E-72").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-8.179481771238642E+64"))
	if err != nil {
		lang.SayString("dddiv2012")
	} else {
		if !(rexsult.ToString() == "-6.170958627632835E-137") {
			lang.SayString("dddiv2012")
		}
	}
	rexsult, err = lang.RxFromString("-3707613309582318").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3394911196503.048"))
	if err != nil {
		lang.SayString("dddiv2013")
	} else {
		if !(rexsult.ToString() == "-1092.109070010836") {
			lang.SayString("dddiv2013")
		}
	}
	rexsult, err = lang.RxFromString("99689.0555190461").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-4.735208553891464"))
	if err != nil {
		lang.SayString("dddiv2014")
	} else {
		if !(rexsult.ToString() == "-21052.72753765411") {
			lang.SayString("dddiv2014")
		}
	}
	rexsult, err = lang.RxFromString("-1447915775613329").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("269750797.8184875"))
	if err != nil {
		lang.SayString("dddiv2015")
	} else {
		if !(rexsult.ToString() == "-5367605.164925653") {
			lang.SayString("dddiv2015")
		}
	}
	rexsult, err = lang.RxFromString("-9.394881304225258E-19").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-830585.0252671636"))
	if err != nil {
		lang.SayString("dddiv2016")
	} else {
		if !(rexsult.ToString() == "1.131116143251358E-24") {
			lang.SayString("dddiv2016")
		}
	}
	rexsult, err = lang.RxFromString("-1.056283432738934").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("88.58754555124013"))
	if err != nil {
		lang.SayString("dddiv2017")
	} else {
		if !(rexsult.ToString() == "-0.01192361100159352") {
			lang.SayString("dddiv2017")
		}
	}
	rexsult, err = lang.RxFromString("5763220933343.081").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("689089567025052.1"))
	if err != nil {
		lang.SayString("dddiv2018")
	} else {
		if !(rexsult.ToString() == "0.008363529516524456") {
			lang.SayString("dddiv2018")
		}
	}
	rexsult, err = lang.RxFromString("873819.122103216").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9.740612494523300E-49"))
	if err != nil {
		lang.SayString("dddiv2019")
	} else {
		if !(rexsult.ToString() == "8.970884763093948E+53") {
			lang.SayString("dddiv2019")
		}
	}
	rexsult, err = lang.RxFromString("8022914.838533576").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6178.566801742713"))
	if err != nil {
		lang.SayString("dddiv2020")
	} else {
		if !(rexsult.ToString() == "1298.507420243583") {
			lang.SayString("dddiv2020")
		}
	}
	rexsult, err = lang.RxFromString("203982.7605650363").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-2158.283639053435"))
	if err != nil {
		lang.SayString("dddiv2021")
	} else {
		if !(rexsult.ToString() == "-94.51156320422168") {
			lang.SayString("dddiv2021")
		}
	}
	rexsult, err = lang.RxFromString("803.6310547013030").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7101143795399.238"))
	if err != nil {
		lang.SayString("dddiv2022")
	} else {
		if !(rexsult.ToString() == "1.131692411611166E-10") {
			lang.SayString("dddiv2022")
		}
	}
	rexsult, err = lang.RxFromString("9.251697842123399E-82").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.342350220606119E-7"))
	if err != nil {
		lang.SayString("dddiv2023")
	} else {
		if !(rexsult.ToString() == "-6.892163982321936E-75") {
			lang.SayString("dddiv2023")
		}
	}
	rexsult, err = lang.RxFromString("-1.980600645637992E-53").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-5.474262753214457E+77"))
	if err != nil {
		lang.SayString("dddiv2024")
	} else {
		if !(rexsult.ToString() == "3.618022617703168E-131") {
			lang.SayString("dddiv2024")
		}
	}
	rexsult, err = lang.RxFromString("-210.0322996351690").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-8.580951835872843E+80"))
	if err != nil {
		lang.SayString("dddiv2025")
	} else {
		if !(rexsult.ToString() == "2.447657365434971E-79") {
			lang.SayString("dddiv2025")
		}
	}
	rexsult, err = lang.RxFromString("-1.821980314020370E+85").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-3.018915267138165"))
	if err != nil {
		lang.SayString("dddiv2026")
	} else {
		if !(rexsult.ToString() == "6.035215144503042E+84") {
			lang.SayString("dddiv2026")
		}
	}
	rexsult, err = lang.RxFromString("-772264503601.1047").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5.158258271408988E-86"))
	if err != nil {
		lang.SayString("dddiv2027")
	} else {
		if !(rexsult.ToString() == "-1.497141986630614E+97") {
			lang.SayString("dddiv2027")
		}
	}
	rexsult, err = lang.RxFromString("-767.0532415847106").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.700027228028939E-59"))
	if err != nil {
		lang.SayString("dddiv2028")
	} else {
		if !(rexsult.ToString() == "-2.840909282772941E+61") {
			lang.SayString("dddiv2028")
		}
	}
	rexsult, err = lang.RxFromString("496724.8548250093").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7.32700588163100E+66"))
	if err != nil {
		lang.SayString("dddiv2029")
	} else {
		if !(rexsult.ToString() == "6.779370220929013E-62") {
			lang.SayString("dddiv2029")
		}
	}
	rexsult, err = lang.RxFromString("-304232651447703.9").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-108.9730808657440"))
	if err != nil {
		lang.SayString("dddiv2030")
	} else {
		if !(rexsult.ToString() == "2791814721862.565") {
			lang.SayString("dddiv2030")
		}
	}
	rexsult, err = lang.RxFromString("-7.233817192699405E+42").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-5711302004.149411"))
	if err != nil {
		lang.SayString("dddiv2031")
	} else {
		if !(rexsult.ToString() == "1.26657935221143E+33") {
			lang.SayString("dddiv2031")
		}
	}
	rexsult, err = lang.RxFromString("-9.999221444912745E+96").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4010569406446197"))
	if err != nil {
		lang.SayString("dddiv2032")
	} else {
		if !(rexsult.ToString() == "-2.49321740420225E+81") {
			lang.SayString("dddiv2032")
		}
	}
	rexsult, err = lang.RxFromString("-1837272.061937622").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8.356322838066762"))
	if err != nil {
		lang.SayString("dddiv2033")
	} else {
		if !(rexsult.ToString() == "-219866.0939196882") {
			lang.SayString("dddiv2033")
		}
	}
	rexsult, err = lang.RxFromString("2168.517555606529").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("209.1910258615061"))
	if err != nil {
		lang.SayString("dddiv2034")
	} else {
		if !(rexsult.ToString() == "10.36620737756784") {
			lang.SayString("dddiv2034")
		}
	}
	rexsult, err = lang.RxFromString("-1.884389790576371E+88").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2.95181953870583E+20"))
	if err != nil {
		lang.SayString("dddiv2035")
	} else {
		if !(rexsult.ToString() == "-6.383824505079828E+67") {
			lang.SayString("dddiv2035")
		}
	}
	rexsult, err = lang.RxFromString("732263.6037438196").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("961222.3634446889"))
	if err != nil {
		lang.SayString("dddiv2036")
	} else {
		if !(rexsult.ToString() == "0.7618045850698269") {
			lang.SayString("dddiv2036")
		}
	}
	rexsult, err = lang.RxFromString("-813461419.0348336").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5.376293753809143E+84"))
	if err != nil {
		lang.SayString("dddiv2037")
	} else {
		if !(rexsult.ToString() == "-1.513052404285927E-76") {
			lang.SayString("dddiv2037")
		}
	}
	rexsult, err = lang.RxFromString("-45562133508108.50").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-9.776843494690107E+51"))
	if err != nil {
		lang.SayString("dddiv2038")
	} else {
		if !(rexsult.ToString() == "4.660208945029519E-39") {
			lang.SayString("dddiv2038")
		}
	}
	rexsult, err = lang.RxFromString("-6.489393172441016E+80").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-9101965.097852113"))
	if err != nil {
		lang.SayString("dddiv2039")
	} else {
		if !(rexsult.ToString() == "7.129661674897421E+73") {
			lang.SayString("dddiv2039")
		}
	}
	rexsult, err = lang.RxFromString("3.694576237117349E+93").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6683512.012622003"))
	if err != nil {
		lang.SayString("dddiv2040")
	} else {
		if !(rexsult.ToString() == "5.527896456443912E+86") {
			lang.SayString("dddiv2040")
		}
	}
	rexsult, err = lang.RxFromString("-2.252877726403272E+19").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-7451913256.181367"))
	if err != nil {
		lang.SayString("dddiv2041")
	} else {
		if !(rexsult.ToString() == "3023220546.125531") {
			lang.SayString("dddiv2041")
		}
	}
	rexsult, err = lang.RxFromString("518303.1989111842").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("50.01587020474133"))
	if err != nil {
		lang.SayString("dddiv2042")
	} else {
		if !(rexsult.ToString() == "10362.77479107123") {
			lang.SayString("dddiv2042")
		}
	}
	rexsult, err = lang.RxFromString("2.902087881880103E+24").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("33.32400992305702"))
	if err != nil {
		lang.SayString("dddiv2043")
	} else {
		if !(rexsult.ToString() == "8.708699488989578E+22") {
			lang.SayString("dddiv2043")
		}
	}
	rexsult, err = lang.RxFromString("549619.4559510557").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1660824845196338"))
	if err != nil {
		lang.SayString("dddiv2044")
	} else {
		if !(rexsult.ToString() == "3.309316196351104E-10") {
			lang.SayString("dddiv2044")
		}
	}
	rexsult, err = lang.RxFromString("-6775670774684043").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8292152023.077262"))
	if err != nil {
		lang.SayString("dddiv2045")
	} else {
		if !(rexsult.ToString() == "-817118.4941891062") {
			lang.SayString("dddiv2045")
		}
	}
	rexsult, err = lang.RxFromString("-77.50923921524079").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-5.636882655425815E+74"))
	if err != nil {
		lang.SayString("dddiv2046")
	} else {
		if !(rexsult.ToString() == "1.375037302588405E-73") {
			lang.SayString("dddiv2046")
		}
	}
	rexsult, err = lang.RxFromString("-2.984889459605149E-10").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-88106156784122.99"))
	if err != nil {
		lang.SayString("dddiv2047")
	} else {
		if !(rexsult.ToString() == "3.387833005721384E-24") {
			lang.SayString("dddiv2047")
		}
	}
	rexsult, err = lang.RxFromString("0.949517293997085").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("44767115.96450998"))
	if err != nil {
		lang.SayString("dddiv2048")
	} else {
		if !(rexsult.ToString() == "2.121015110175589E-8") {
			lang.SayString("dddiv2048")
		}
	}
	rexsult, err = lang.RxFromString("-2760937211.084521").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1087015876975408"))
	if err != nil {
		lang.SayString("dddiv2049")
	} else {
		if !(rexsult.ToString() == "0.000002539923537057024") {
			lang.SayString("dddiv2049")
		}
	}
	rexsult, err = lang.RxFromString("28438351.85030536").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-4.209397904088624E-47"))
	if err != nil {
		lang.SayString("dddiv2050")
	} else {
		if !(rexsult.ToString() == "-6.755919135770688E+53") {
			lang.SayString("dddiv2050")
		}
	}
	rexsult, err = lang.RxFromString("-85562731.6820956").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-7.166045442530185E+45"))
	if err != nil {
		lang.SayString("dddiv2051")
	} else {
		if !(rexsult.ToString() == "1.194002080621542E-38") {
			lang.SayString("dddiv2051")
		}
	}
	rexsult, err = lang.RxFromString("2533802852165.25").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7154.119606235955"))
	if err != nil {
		lang.SayString("dddiv2052")
	} else {
		if !(rexsult.ToString() == "354173957.3317501") {
			lang.SayString("dddiv2052")
		}
	}
	rexsult, err = lang.RxFromString("-8858831346851.474").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("97.59734208801716"))
	if err != nil {
		lang.SayString("dddiv2053")
	} else {
		if !(rexsult.ToString() == "-90769186509.83577") {
			lang.SayString("dddiv2053")
		}
	}
	rexsult, err = lang.RxFromString("176783629801387.5").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("840073263.3109817"))
	if err != nil {
		lang.SayString("dddiv2054")
	} else {
		if !(rexsult.ToString() == "210438.3480848206") {
			lang.SayString("dddiv2054")
		}
	}
	rexsult, err = lang.RxFromString("-493506471796175.6").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("79733894790822.03"))
	if err != nil {
		lang.SayString("dddiv2055")
	} else {
		if !(rexsult.ToString() == "-6.189418854940746") {
			lang.SayString("dddiv2055")
		}
	}
	rexsult, err = lang.RxFromString("790.1682542103445").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("829.9449370367435"))
	if err != nil {
		lang.SayString("dddiv2056")
	} else {
		if !(rexsult.ToString() == "0.9520731062371214") {
			lang.SayString("dddiv2056")
		}
	}
	rexsult, err = lang.RxFromString("-8920459838.583164").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-4767.889187899214"))
	if err != nil {
		lang.SayString("dddiv2057")
	} else {
		if !(rexsult.ToString() == "1870945.294035581") {
			lang.SayString("dddiv2057")
		}
	}
	rexsult, err = lang.RxFromString("53536687164422.1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("53137.5007032689"))
	if err != nil {
		lang.SayString("dddiv2058")
	} else {
		if !(rexsult.ToString() == "1007512330.385698") {
			lang.SayString("dddiv2058")
		}
	}
	rexsult, err = lang.RxFromString("4.051532311146561E-74").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-2.343089768972261E+94"))
	if err != nil {
		lang.SayString("dddiv2059")
	} else {
		if !(rexsult.ToString() == "-1.729140882606332E-168") {
			lang.SayString("dddiv2059")
		}
	}
	rexsult, err = lang.RxFromString("-14847758778636.88").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3.062543516383807E-43"))
	if err != nil {
		lang.SayString("dddiv2060")
	} else {
		if !(rexsult.ToString() == "-4.848178874587497E+55") {
			lang.SayString("dddiv2060")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3001")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv3001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3002")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv3002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("99999"))
	if err != nil {
		lang.SayString("dddiv3003")
	} else {
		if !(rexsult.ToString() == "0.00001000010000100001") {
			lang.SayString("dddiv3003")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("dddiv3004")
	} else {
		if !(rexsult.ToString() == "0.000001000001000001") {
			lang.SayString("dddiv3004")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3011")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv3011")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3012")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv3012")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("99999"))
	if err != nil {
		lang.SayString("dddiv3013")
	} else {
		if !(rexsult.ToString() == "0.00001000010000100001") {
			lang.SayString("dddiv3013")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("dddiv3014")
	} else {
		if !(rexsult.ToString() == "0.000001000001000001") {
			lang.SayString("dddiv3014")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3021")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv3021")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3022")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv3022")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("99999"))
	if err != nil {
		lang.SayString("dddiv3023")
	} else {
		if !(rexsult.ToString() == "0.00001000010000100001") {
			lang.SayString("dddiv3023")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("dddiv3024")
	} else {
		if !(rexsult.ToString() == "0.000001000001000001") {
			lang.SayString("dddiv3024")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3031")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv3031")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3032")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv3032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("99999"))
	if err != nil {
		lang.SayString("dddiv3033")
	} else {
		if !(rexsult.ToString() == "0.00001000010000100001") {
			lang.SayString("dddiv3033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("dddiv3034")
	} else {
		if !(rexsult.ToString() == "0.000001000001000001") {
			lang.SayString("dddiv3034")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3041")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv3041")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3042")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv3042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("99999"))
	if err != nil {
		lang.SayString("dddiv3043")
	} else {
		if !(rexsult.ToString() == "0.00001000010000100001") {
			lang.SayString("dddiv3043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("dddiv3044")
	} else {
		if !(rexsult.ToString() == "0.000001000001000001") {
			lang.SayString("dddiv3044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3051")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv3051")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3052")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv3052")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("99999"))
	if err != nil {
		lang.SayString("dddiv3053")
	} else {
		if !(rexsult.ToString() == "0.00001000010000100001") {
			lang.SayString("dddiv3053")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("dddiv3054")
	} else {
		if !(rexsult.ToString() == "0.000001000001000001") {
			lang.SayString("dddiv3054")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3061")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv3061")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3062")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv3062")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("99999"))
	if err != nil {
		lang.SayString("dddiv3063")
	} else {
		if !(rexsult.ToString() == "0.00001000010000100001") {
			lang.SayString("dddiv3063")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("dddiv3064")
	} else {
		if !(rexsult.ToString() == "0.000001000001000001") {
			lang.SayString("dddiv3064")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3071")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333") {
			lang.SayString("dddiv3071")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dddiv3072")
	} else {
		if !(rexsult.ToString() == "0.6666666666666667") {
			lang.SayString("dddiv3072")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("99999"))
	if err != nil {
		lang.SayString("dddiv3073")
	} else {
		if !(rexsult.ToString() == "0.00001000010000100001") {
			lang.SayString("dddiv3073")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("dddiv3074")
	} else {
		if !(rexsult.ToString() == "0.000001000001000001") {
			lang.SayString("dddiv3074")
		}
	}
	rexsult, err = lang.RxFromString("3195385192916917").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3195385192946695"))
	if err != nil {
		lang.SayString("dddiv4001")
	} else {
		if !(rexsult.ToString() == "0.9999999999906809") {
			lang.SayString("dddiv4001")
		}
	}
	rexsult, err = lang.RxFromString("1393723067526993").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1393723067519475"))
	if err != nil {
		lang.SayString("dddiv4002")
	} else {
		if !(rexsult.ToString() == "1.000000000005394") {
			lang.SayString("dddiv4002")
		}
	}
	rexsult, err = lang.RxFromString("759985543702302").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("759985543674015"))
	if err != nil {
		lang.SayString("dddiv4003")
	} else {
		if !(rexsult.ToString() == "1.00000000003722") {
			lang.SayString("dddiv4003")
		}
	}
	rexsult, err = lang.RxFromString("9579158456027302").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9579158456036864"))
	if err != nil {
		lang.SayString("dddiv4004")
	} else {
		if !(rexsult.ToString() == "0.9999999999990018") {
			lang.SayString("dddiv4004")
		}
	}
	rexsult, err = lang.RxFromString("7079398299143569").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7079398299156904"))
	if err != nil {
		lang.SayString("dddiv4005")
	} else {
		if !(rexsult.ToString() == "0.9999999999981164") {
			lang.SayString("dddiv4005")
		}
	}
	rexsult, err = lang.RxFromString("6636169255366598").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6636169255336386"))
	if err != nil {
		lang.SayString("dddiv4006")
	} else {
		if !(rexsult.ToString() == "1.000000000004553") {
			lang.SayString("dddiv4006")
		}
	}
	rexsult, err = lang.RxFromString("6964813971340090").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6964813971321554"))
	if err != nil {
		lang.SayString("dddiv4007")
	} else {
		if !(rexsult.ToString() == "1.000000000002661") {
			lang.SayString("dddiv4007")
		}
	}
	rexsult, err = lang.RxFromString("4182275225480784").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4182275225454009"))
	if err != nil {
		lang.SayString("dddiv4008")
	} else {
		if !(rexsult.ToString() == "1.000000000006402") {
			lang.SayString("dddiv4008")
		}
	}
	rexsult, err = lang.RxFromString("9228325124938029").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9228325124918730"))
	if err != nil {
		lang.SayString("dddiv4009")
	} else {
		if !(rexsult.ToString() == "1.000000000002091") {
			lang.SayString("dddiv4009")
		}
	}
	rexsult, err = lang.RxFromString("3428346338630192").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3428346338609843"))
	if err != nil {
		lang.SayString("dddiv4010")
	} else {
		if !(rexsult.ToString() == "1.000000000005936") {
			lang.SayString("dddiv4010")
		}
	}
	rexsult, err = lang.RxFromString("2143511550722893").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2143511550751754"))
	if err != nil {
		lang.SayString("dddiv4011")
	} else {
		if !(rexsult.ToString() == "0.9999999999865356") {
			lang.SayString("dddiv4011")
		}
	}
	rexsult, err = lang.RxFromString("1672732924396785").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1672732924401811"))
	if err != nil {
		lang.SayString("dddiv4012")
	} else {
		if !(rexsult.ToString() == "0.9999999999969953") {
			lang.SayString("dddiv4012")
		}
	}
	rexsult, err = lang.RxFromString("4190714611948216").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4190714611948664"))
	if err != nil {
		lang.SayString("dddiv4013")
	} else {
		if !(rexsult.ToString() == "0.9999999999998931") {
			lang.SayString("dddiv4013")
		}
	}
	rexsult, err = lang.RxFromString("3942254800848877").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3942254800814556"))
	if err != nil {
		lang.SayString("dddiv4014")
	} else {
		if !(rexsult.ToString() == "1.000000000008706") {
			lang.SayString("dddiv4014")
		}
	}
	rexsult, err = lang.RxFromString("2854459826952334").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2854459826960762"))
	if err != nil {
		lang.SayString("dddiv4015")
	} else {
		if !(rexsult.ToString() == "0.9999999999970474") {
			lang.SayString("dddiv4015")
		}
	}
	rexsult, err = lang.RxFromString("2853258953664731").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2853258953684471"))
	if err != nil {
		lang.SayString("dddiv4016")
	} else {
		if !(rexsult.ToString() == "0.9999999999930816") {
			lang.SayString("dddiv4016")
		}
	}
	rexsult, err = lang.RxFromString("9453512638125978").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9453512638146425"))
	if err != nil {
		lang.SayString("dddiv4017")
	} else {
		if !(rexsult.ToString() == "0.9999999999978371") {
			lang.SayString("dddiv4017")
		}
	}
	rexsult, err = lang.RxFromString("339476633940369").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("339476633912887"))
	if err != nil {
		lang.SayString("dddiv4018")
	} else {
		if !(rexsult.ToString() == "1.000000000080954") {
			lang.SayString("dddiv4018")
		}
	}
	rexsult, err = lang.RxFromString("4542181492688467").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4542181492697735"))
	if err != nil {
		lang.SayString("dddiv4019")
	} else {
		if !(rexsult.ToString() == "0.9999999999979596") {
			lang.SayString("dddiv4019")
		}
	}
	rexsult, err = lang.RxFromString("7312600192399197").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7312600192395424"))
	if err != nil {
		lang.SayString("dddiv4020")
	} else {
		if !(rexsult.ToString() == "1.000000000000516") {
			lang.SayString("dddiv4020")
		}
	}
	rexsult, err = lang.RxFromString("1811674985570111").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1811674985603935"))
	if err != nil {
		lang.SayString("dddiv4021")
	} else {
		if !(rexsult.ToString() == "0.99999999998133") {
			lang.SayString("dddiv4021")
		}
	}
	rexsult, err = lang.RxFromString("1706462639003481").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1706462639017740"))
	if err != nil {
		lang.SayString("dddiv4022")
	} else {
		if !(rexsult.ToString() == "0.9999999999916441") {
			lang.SayString("dddiv4022")
		}
	}
	rexsult, err = lang.RxFromString("6697052654940368").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6697052654934110"))
	if err != nil {
		lang.SayString("dddiv4023")
	} else {
		if !(rexsult.ToString() == "1.000000000000934") {
			lang.SayString("dddiv4023")
		}
	}
	rexsult, err = lang.RxFromString("5015283664277539").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5015283664310719"))
	if err != nil {
		lang.SayString("dddiv4024")
	} else {
		if !(rexsult.ToString() == "0.9999999999933842") {
			lang.SayString("dddiv4024")
		}
	}
	rexsult, err = lang.RxFromString("2359501561537464").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2359501561502464"))
	if err != nil {
		lang.SayString("dddiv4025")
	} else {
		if !(rexsult.ToString() == "1.000000000014834") {
			lang.SayString("dddiv4025")
		}
	}
	rexsult, err = lang.RxFromString("2669850227909157").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2669850227901548"))
	if err != nil {
		lang.SayString("dddiv4026")
	} else {
		if !(rexsult.ToString() == "1.00000000000285") {
			lang.SayString("dddiv4026")
		}
	}
	rexsult, err = lang.RxFromString("9329725546974648").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9329725547002445"))
	if err != nil {
		lang.SayString("dddiv4027")
	} else {
		if !(rexsult.ToString() == "0.9999999999970206") {
			lang.SayString("dddiv4027")
		}
	}
	rexsult, err = lang.RxFromString("3228562867071248").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3228562867106206"))
	if err != nil {
		lang.SayString("dddiv4028")
	} else {
		if !(rexsult.ToString() == "0.9999999999891723") {
			lang.SayString("dddiv4028")
		}
	}
	rexsult, err = lang.RxFromString("4862226644921175").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4862226644909380"))
	if err != nil {
		lang.SayString("dddiv4029")
	} else {
		if !(rexsult.ToString() == "1.000000000002426") {
			lang.SayString("dddiv4029")
		}
	}
	rexsult, err = lang.RxFromString("1022267997054529").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1022267997071329"))
	if err != nil {
		lang.SayString("dddiv4030")
	} else {
		if !(rexsult.ToString() == "0.999999999983566") {
			lang.SayString("dddiv4030")
		}
	}
	rexsult, err = lang.RxFromString("1048777482023719").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1048777482000948"))
	if err != nil {
		lang.SayString("dddiv4031")
	} else {
		if !(rexsult.ToString() == "1.000000000021712") {
			lang.SayString("dddiv4031")
		}
	}
	rexsult, err = lang.RxFromString("9980113777337098").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9980113777330539"))
	if err != nil {
		lang.SayString("dddiv4032")
	} else {
		if !(rexsult.ToString() == "1.000000000000657") {
			lang.SayString("dddiv4032")
		}
	}
	rexsult, err = lang.RxFromString("7506839167963908").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7506839167942901"))
	if err != nil {
		lang.SayString("dddiv4033")
	} else {
		if !(rexsult.ToString() == "1.000000000002798") {
			lang.SayString("dddiv4033")
		}
	}
	rexsult, err = lang.RxFromString("231119751977860").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("231119751962453"))
	if err != nil {
		lang.SayString("dddiv4034")
	} else {
		if !(rexsult.ToString() == "1.000000000066662") {
			lang.SayString("dddiv4034")
		}
	}
	rexsult, err = lang.RxFromString("4034903664762962").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4034903664795526"))
	if err != nil {
		lang.SayString("dddiv4035")
	} else {
		if !(rexsult.ToString() == "0.9999999999919294") {
			lang.SayString("dddiv4035")
		}
	}
	rexsult, err = lang.RxFromString("5700122152274696").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5700122152251386"))
	if err != nil {
		lang.SayString("dddiv4036")
	} else {
		if !(rexsult.ToString() == "1.000000000004089") {
			lang.SayString("dddiv4036")
		}
	}
	rexsult, err = lang.RxFromString("6869599590293110").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6869599590293495"))
	if err != nil {
		lang.SayString("dddiv4037")
	} else {
		if !(rexsult.ToString() == "0.999999999999944") {
			lang.SayString("dddiv4037")
		}
	}
	rexsult, err = lang.RxFromString("5576281960092797").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5576281960105579"))
	if err != nil {
		lang.SayString("dddiv4038")
	} else {
		if !(rexsult.ToString() == "0.9999999999977078") {
			lang.SayString("dddiv4038")
		}
	}
	rexsult, err = lang.RxFromString("2304844888381318").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2304844888353073"))
	if err != nil {
		lang.SayString("dddiv4039")
	} else {
		if !(rexsult.ToString() == "1.000000000012255") {
			lang.SayString("dddiv4039")
		}
	}
	rexsult, err = lang.RxFromString("3265933651656452").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3265933651682779"))
	if err != nil {
		lang.SayString("dddiv4040")
	} else {
		if !(rexsult.ToString() == "0.9999999999919389") {
			lang.SayString("dddiv4040")
		}
	}
	rexsult, err = lang.RxFromString("5235714985079914").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5235714985066131"))
	if err != nil {
		lang.SayString("dddiv4041")
	} else {
		if !(rexsult.ToString() == "1.000000000002632") {
			lang.SayString("dddiv4041")
		}
	}
	rexsult, err = lang.RxFromString("5578481572827551").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5578481572822945"))
	if err != nil {
		lang.SayString("dddiv4042")
	} else {
		if !(rexsult.ToString() == "1.000000000000826") {
			lang.SayString("dddiv4042")
		}
	}
	rexsult, err = lang.RxFromString("4909616081396134").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4909616081373076"))
	if err != nil {
		lang.SayString("dddiv4043")
	} else {
		if !(rexsult.ToString() == "1.000000000004696") {
			lang.SayString("dddiv4043")
		}
	}
	rexsult, err = lang.RxFromString("636447224349537").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("636447224338757"))
	if err != nil {
		lang.SayString("dddiv4044")
	} else {
		if !(rexsult.ToString() == "1.000000000016938") {
			lang.SayString("dddiv4044")
		}
	}
	rexsult, err = lang.RxFromString("1539373428396640").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1539373428364727"))
	if err != nil {
		lang.SayString("dddiv4045")
	} else {
		if !(rexsult.ToString() == "1.000000000020731") {
			lang.SayString("dddiv4045")
		}
	}
	rexsult, err = lang.RxFromString("2028786707377893").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2028786707378866"))
	if err != nil {
		lang.SayString("dddiv4046")
	} else {
		if !(rexsult.ToString() == "0.9999999999995204") {
			lang.SayString("dddiv4046")
		}
	}
	rexsult, err = lang.RxFromString("137643260486222").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("137643260487419"))
	if err != nil {
		lang.SayString("dddiv4047")
	} else {
		if !(rexsult.ToString() == "0.9999999999913036") {
			lang.SayString("dddiv4047")
		}
	}
	rexsult, err = lang.RxFromString("247451519746765").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("247451519752267"))
	if err != nil {
		lang.SayString("dddiv4048")
	} else {
		if !(rexsult.ToString() == "0.9999999999777653") {
			lang.SayString("dddiv4048")
		}
	}
	rexsult, err = lang.RxFromString("7877858475022054").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7877858474999794"))
	if err != nil {
		lang.SayString("dddiv4049")
	} else {
		if !(rexsult.ToString() == "1.000000000002826") {
			lang.SayString("dddiv4049")
		}
	}
	rexsult, err = lang.RxFromString("7333242694766258").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7333242694744628"))
	if err != nil {
		lang.SayString("dddiv4050")
	} else {
		if !(rexsult.ToString() == "1.00000000000295") {
			lang.SayString("dddiv4050")
		}
	}
	rexsult, err = lang.RxFromString("124051503698592").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("124051503699397"))
	if err != nil {
		lang.SayString("dddiv4051")
	} else {
		if !(rexsult.ToString() == "0.9999999999935108") {
			lang.SayString("dddiv4051")
		}
	}
	rexsult, err = lang.RxFromString("8944737432385188").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8944737432406860"))
	if err != nil {
		lang.SayString("dddiv4052")
	} else {
		if !(rexsult.ToString() == "0.9999999999975771") {
			lang.SayString("dddiv4052")
		}
	}
	rexsult, err = lang.RxFromString("9883948923406874").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9883948923424843"))
	if err != nil {
		lang.SayString("dddiv4053")
	} else {
		if !(rexsult.ToString() == "0.999999999998182") {
			lang.SayString("dddiv4053")
		}
	}
	rexsult, err = lang.RxFromString("6829178741654284").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6829178741671973"))
	if err != nil {
		lang.SayString("dddiv4054")
	} else {
		if !(rexsult.ToString() == "0.9999999999974098") {
			lang.SayString("dddiv4054")
		}
	}
	rexsult, err = lang.RxFromString("7342752479768122").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7342752479793385"))
	if err != nil {
		lang.SayString("dddiv4055")
	} else {
		if !(rexsult.ToString() == "0.9999999999965595") {
			lang.SayString("dddiv4055")
		}
	}
	rexsult, err = lang.RxFromString("8066426579008783").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8066426578977563"))
	if err != nil {
		lang.SayString("dddiv4056")
	} else {
		if !(rexsult.ToString() == "1.00000000000387") {
			lang.SayString("dddiv4056")
		}
	}
	rexsult, err = lang.RxFromString("8992775071383295").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8992775071352712"))
	if err != nil {
		lang.SayString("dddiv4057")
	} else {
		if !(rexsult.ToString() == "1.000000000003401") {
			lang.SayString("dddiv4057")
		}
	}
	rexsult, err = lang.RxFromString("5485011755545641").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5485011755543611"))
	if err != nil {
		lang.SayString("dddiv4058")
	} else {
		if !(rexsult.ToString() == "1.00000000000037") {
			lang.SayString("dddiv4058")
		}
	}
	rexsult, err = lang.RxFromString("5779983054353918").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5779983054365300"))
	if err != nil {
		lang.SayString("dddiv4059")
	} else {
		if !(rexsult.ToString() == "0.9999999999980308") {
			lang.SayString("dddiv4059")
		}
	}
	rexsult, err = lang.RxFromString("9502265102713774").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9502265102735208"))
	if err != nil {
		lang.SayString("dddiv4060")
	} else {
		if !(rexsult.ToString() == "0.9999999999977443") {
			lang.SayString("dddiv4060")
		}
	}
	rexsult, err = lang.RxFromString("2109558399130981").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2109558399116281"))
	if err != nil {
		lang.SayString("dddiv4061")
	} else {
		if !(rexsult.ToString() == "1.000000000006968") {
			lang.SayString("dddiv4061")
		}
	}
	rexsult, err = lang.RxFromString("5296182636350471").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5296182636351521"))
	if err != nil {
		lang.SayString("dddiv4062")
	} else {
		if !(rexsult.ToString() == "0.9999999999998017") {
			lang.SayString("dddiv4062")
		}
	}
	rexsult, err = lang.RxFromString("1440019225591883").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1440019225601844"))
	if err != nil {
		lang.SayString("dddiv4063")
	} else {
		if !(rexsult.ToString() == "0.9999999999930827") {
			lang.SayString("dddiv4063")
		}
	}
	rexsult, err = lang.RxFromString("8182110791881341").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8182110791847174"))
	if err != nil {
		lang.SayString("dddiv4064")
	} else {
		if !(rexsult.ToString() == "1.000000000004176") {
			lang.SayString("dddiv4064")
		}
	}
	rexsult, err = lang.RxFromString("489098235512060").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("489098235534516"))
	if err != nil {
		lang.SayString("dddiv4065")
	} else {
		if !(rexsult.ToString() == "0.9999999999540869") {
			lang.SayString("dddiv4065")
		}
	}
	rexsult, err = lang.RxFromString("6475687084782038").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6475687084756089"))
	if err != nil {
		lang.SayString("dddiv4066")
	} else {
		if !(rexsult.ToString() == "1.000000000004007") {
			lang.SayString("dddiv4066")
		}
	}
	rexsult, err = lang.RxFromString("8094348555736948").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8094348555759236"))
	if err != nil {
		lang.SayString("dddiv4067")
	} else {
		if !(rexsult.ToString() == "0.9999999999972465") {
			lang.SayString("dddiv4067")
		}
	}
	rexsult, err = lang.RxFromString("1982766816291543").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1982766816309463"))
	if err != nil {
		lang.SayString("dddiv4068")
	} else {
		if !(rexsult.ToString() == "0.9999999999909621") {
			lang.SayString("dddiv4068")
		}
	}
	rexsult, err = lang.RxFromString("9277314300113251").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9277314300084467"))
	if err != nil {
		lang.SayString("dddiv4069")
	} else {
		if !(rexsult.ToString() == "1.000000000003103") {
			lang.SayString("dddiv4069")
		}
	}
	rexsult, err = lang.RxFromString("4335532959318934").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4335532959293167"))
	if err != nil {
		lang.SayString("dddiv4070")
	} else {
		if !(rexsult.ToString() == "1.000000000005943") {
			lang.SayString("dddiv4070")
		}
	}
	rexsult, err = lang.RxFromString("7767113032981348").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7767113032968132"))
	if err != nil {
		lang.SayString("dddiv4071")
	} else {
		if !(rexsult.ToString() == "1.000000000001702") {
			lang.SayString("dddiv4071")
		}
	}
	rexsult, err = lang.RxFromString("1578548053342868").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1578548053370448"))
	if err != nil {
		lang.SayString("dddiv4072")
	} else {
		if !(rexsult.ToString() == "0.9999999999825282") {
			lang.SayString("dddiv4072")
		}
	}
	rexsult, err = lang.RxFromString("3790420686666898").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3790420686636315"))
	if err != nil {
		lang.SayString("dddiv4073")
	} else {
		if !(rexsult.ToString() == "1.000000000008068") {
			lang.SayString("dddiv4073")
		}
	}
	rexsult, err = lang.RxFromString("871682421955147").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("871682421976441"))
	if err != nil {
		lang.SayString("dddiv4074")
	} else {
		if !(rexsult.ToString() == "0.9999999999755714") {
			lang.SayString("dddiv4074")
		}
	}
	rexsult, err = lang.RxFromString("744141054479940").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("744141054512329"))
	if err != nil {
		lang.SayString("dddiv4075")
	} else {
		if !(rexsult.ToString() == "0.9999999999564746") {
			lang.SayString("dddiv4075")
		}
	}
	rexsult, err = lang.RxFromString("8956824183670735").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8956824183641741"))
	if err != nil {
		lang.SayString("dddiv4076")
	} else {
		if !(rexsult.ToString() == "1.000000000003237") {
			lang.SayString("dddiv4076")
		}
	}
	rexsult, err = lang.RxFromString("8337291694485682").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8337291694451193"))
	if err != nil {
		lang.SayString("dddiv4077")
	} else {
		if !(rexsult.ToString() == "1.000000000004137") {
			lang.SayString("dddiv4077")
		}
	}
	rexsult, err = lang.RxFromString("4107775944683669").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4107775944657097"))
	if err != nil {
		lang.SayString("dddiv4078")
	} else {
		if !(rexsult.ToString() == "1.000000000006469") {
			lang.SayString("dddiv4078")
		}
	}
	rexsult, err = lang.RxFromString("8691900057964648").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("8691900057997555"))
	if err != nil {
		lang.SayString("dddiv4079")
	} else {
		if !(rexsult.ToString() == "0.9999999999962141") {
			lang.SayString("dddiv4079")
		}
	}
	rexsult, err = lang.RxFromString("2229528520536462").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2229528520502337"))
	if err != nil {
		lang.SayString("dddiv4080")
	} else {
		if !(rexsult.ToString() == "1.000000000015306") {
			lang.SayString("dddiv4080")
		}
	}
	rexsult, err = lang.RxFromString("398442083774322").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("398442083746273"))
	if err != nil {
		lang.SayString("dddiv4081")
	} else {
		if !(rexsult.ToString() == "1.000000000070397") {
			lang.SayString("dddiv4081")
		}
	}
	rexsult, err = lang.RxFromString("5319819776808759").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5319819776838313"))
	if err != nil {
		lang.SayString("dddiv4082")
	} else {
		if !(rexsult.ToString() == "0.9999999999944445") {
			lang.SayString("dddiv4082")
		}
	}
	rexsult, err = lang.RxFromString("7710491299066855").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7710491299041858"))
	if err != nil {
		lang.SayString("dddiv4083")
	} else {
		if !(rexsult.ToString() == "1.000000000003242") {
			lang.SayString("dddiv4083")
		}
	}
	rexsult, err = lang.RxFromString("9083231296087266").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9083231296058160"))
	if err != nil {
		lang.SayString("dddiv4084")
	} else {
		if !(rexsult.ToString() == "1.000000000003204") {
			lang.SayString("dddiv4084")
		}
	}
	rexsult, err = lang.RxFromString("3566873574904559").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3566873574890328"))
	if err != nil {
		lang.SayString("dddiv4085")
	} else {
		if !(rexsult.ToString() == "1.00000000000399") {
			lang.SayString("dddiv4085")
		}
	}
	rexsult, err = lang.RxFromString("596343290550525").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("596343290555614"))
	if err != nil {
		lang.SayString("dddiv4086")
	} else {
		if !(rexsult.ToString() == "0.9999999999914663") {
			lang.SayString("dddiv4086")
		}
	}
	rexsult, err = lang.RxFromString("278227925093192").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("278227925068104"))
	if err != nil {
		lang.SayString("dddiv4087")
	} else {
		if !(rexsult.ToString() == "1.000000000090171") {
			lang.SayString("dddiv4087")
		}
	}
	rexsult, err = lang.RxFromString("3292902958490649").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3292902958519881"))
	if err != nil {
		lang.SayString("dddiv4088")
	} else {
		if !(rexsult.ToString() == "0.9999999999911227") {
			lang.SayString("dddiv4088")
		}
	}
	rexsult, err = lang.RxFromString("5521871364245881").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5521871364229536"))
	if err != nil {
		lang.SayString("dddiv4089")
	} else {
		if !(rexsult.ToString() == "1.00000000000296") {
			lang.SayString("dddiv4089")
		}
	}
	rexsult, err = lang.RxFromString("2406505602883617").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("2406505602857997"))
	if err != nil {
		lang.SayString("dddiv4090")
	} else {
		if !(rexsult.ToString() == "1.000000000010646") {
			lang.SayString("dddiv4090")
		}
	}
	rexsult, err = lang.RxFromString("7741146984869208").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("7741146984867255"))
	if err != nil {
		lang.SayString("dddiv4091")
	} else {
		if !(rexsult.ToString() == "1.000000000000252") {
			lang.SayString("dddiv4091")
		}
	}
	rexsult, err = lang.RxFromString("4576041832414909").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("4576041832405102"))
	if err != nil {
		lang.SayString("dddiv4092")
	} else {
		if !(rexsult.ToString() == "1.000000000002143") {
			lang.SayString("dddiv4092")
		}
	}
	rexsult, err = lang.RxFromString("9183756982878057").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("9183756982901934"))
	if err != nil {
		lang.SayString("dddiv4093")
	} else {
		if !(rexsult.ToString() == "0.9999999999974001") {
			lang.SayString("dddiv4093")
		}
	}
	rexsult, err = lang.RxFromString("6215736513855159").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6215736513870342"))
	if err != nil {
		lang.SayString("dddiv4094")
	} else {
		if !(rexsult.ToString() == "0.9999999999975573") {
			lang.SayString("dddiv4094")
		}
	}
	rexsult, err = lang.RxFromString("248554968534533").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("248554968551417"))
	if err != nil {
		lang.SayString("dddiv4095")
	} else {
		if !(rexsult.ToString() == "0.9999999999320714") {
			lang.SayString("dddiv4095")
		}
	}
	rexsult, err = lang.RxFromString("376314165668645").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("376314165659755"))
	if err != nil {
		lang.SayString("dddiv4096")
	} else {
		if !(rexsult.ToString() == "1.000000000023624") {
			lang.SayString("dddiv4096")
		}
	}
	rexsult, err = lang.RxFromString("5513569249809718").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("5513569249808906"))
	if err != nil {
		lang.SayString("dddiv4097")
	} else {
		if !(rexsult.ToString() == "1.000000000000147") {
			lang.SayString("dddiv4097")
		}
	}
	rexsult, err = lang.RxFromString("3367992242167904").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("3367992242156228"))
	if err != nil {
		lang.SayString("dddiv4098")
	} else {
		if !(rexsult.ToString() == "1.000000000003467") {
			lang.SayString("dddiv4098")
		}
	}
	rexsult, err = lang.RxFromString("6134869538966967").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("6134869538985986"))
	if err != nil {
		lang.SayString("dddiv4099")
	} else {
		if !(rexsult.ToString() == "0.9999999999968999") {
			lang.SayString("dddiv4099")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx001")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx002")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx003")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("divx007")
	} else {
		if !(rexsult.ToString() == "0.333333333") {
			lang.SayString("divx007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("divx008")
	} else {
		if !(rexsult.ToString() == "0.666666667") {
			lang.SayString("divx008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("divx009")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx010")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("divx010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("divx011")
	} else {
		if !(rexsult.ToString() == "-2.4") {
			lang.SayString("divx011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx012")
	} else {
		if !(rexsult.ToString() == "-2.4") {
			lang.SayString("divx012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("divx013")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("divx013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx014")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("divx014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx015")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("divx015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx016")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("divx016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx017")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("divx017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx018")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("divx019")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("divx020")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx021")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx021")
		}
	}
	rexsult, err = lang.RxFromString("50").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("divx022")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx022")
		}
	}
	rexsult, err = lang.RxFromString("500").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("200"))
	if err != nil {
		lang.SayString("divx023")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx023")
		}
	}
	rexsult, err = lang.RxFromString("50.0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("20.0"))
	if err != nil {
		lang.SayString("divx024")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx024")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("divx025")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx025")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("divx026")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx026")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("divx027")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx027")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("0.20"))
	if err != nil {
		lang.SayString("divx028")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("divx028")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("divx029")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("divx029")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx030")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("divx030")
		}
	}
	rexsult, err = lang.RxFromString("100").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx031")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("divx031")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx032")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("divx032")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("divx033")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("divx033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx035")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("divx036")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("divx036")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("divx037")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("divx037")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("divx038")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("divx038")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("divx039")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("divx039")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("divx040")
	} else {
		if !(rexsult.ToString() == "0.015625") {
			lang.SayString("divx040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("divx041")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("divx041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("divx042")
	} else {
		if !(rexsult.ToString() == "-0.25") {
			lang.SayString("divx042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("divx043")
	} else {
		if !(rexsult.ToString() == "-0.125") {
			lang.SayString("divx043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("divx044")
	} else {
		if !(rexsult.ToString() == "-0.0625") {
			lang.SayString("divx044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("divx045")
	} else {
		if !(rexsult.ToString() == "-0.03125") {
			lang.SayString("divx045")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("divx046")
	} else {
		if !(rexsult.ToString() == "-0.015625") {
			lang.SayString("divx046")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx047")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("divx047")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("divx048")
	} else {
		if !(rexsult.ToString() == "-0.25") {
			lang.SayString("divx048")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("divx049")
	} else {
		if !(rexsult.ToString() == "-0.125") {
			lang.SayString("divx049")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("divx050")
	} else {
		if !(rexsult.ToString() == "-0.0625") {
			lang.SayString("divx050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("divx051")
	} else {
		if !(rexsult.ToString() == "-0.03125") {
			lang.SayString("divx051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("divx052")
	} else {
		if !(rexsult.ToString() == "-0.015625") {
			lang.SayString("divx052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("divx053")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("divx054")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("divx054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("divx055")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("divx055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("divx056")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("divx056")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("divx057")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("divx057")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("divx058")
	} else {
		if !(rexsult.ToString() == "0.015625") {
			lang.SayString("divx058")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx070")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("divx070")
		}
	}
	rexsult, err = lang.RxFromString("999999999.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx071")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("divx071")
		}
	}
	rexsult, err = lang.RxFromString("999999999.5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx072")
	} else {
		if !(rexsult.ToString() == "1E+9") {
			lang.SayString("divx072")
		}
	}
	rexsult, err = lang.RxFromString("999999999.9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx073")
	} else {
		if !(rexsult.ToString() == "1E+9") {
			lang.SayString("divx073")
		}
	}
	rexsult, err = lang.RxFromString("999999999.999").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx074")
	} else {
		if !(rexsult.ToString() == "1E+9") {
			lang.SayString("divx074")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx080")
	} else {
		if !(rexsult.ToString() == "1E+9") {
			lang.SayString("divx080")
		}
	}
	rexsult, err = lang.RxFromString("99999999").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx081")
	} else {
		if !(rexsult.ToString() == "1E+8") {
			lang.SayString("divx081")
		}
	}
	rexsult, err = lang.RxFromString("9999999").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx082")
	} else {
		if !(rexsult.ToString() == "1E+7") {
			lang.SayString("divx082")
		}
	}
	rexsult, err = lang.RxFromString("999999").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx083")
	} else {
		if !(rexsult.ToString() == "999999") {
			lang.SayString("divx083")
		}
	}
	rexsult, err = lang.RxFromString("99999").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx084")
	} else {
		if !(rexsult.ToString() == "99999") {
			lang.SayString("divx084")
		}
	}
	rexsult, err = lang.RxFromString("9999").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx085")
	} else {
		if !(rexsult.ToString() == "9999") {
			lang.SayString("divx085")
		}
	}
	rexsult, err = lang.RxFromString("999").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx086")
	} else {
		if !(rexsult.ToString() == "999") {
			lang.SayString("divx086")
		}
	}
	rexsult, err = lang.RxFromString("99").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx087")
	} else {
		if !(rexsult.ToString() == "99") {
			lang.SayString("divx087")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx088")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("divx088")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx090")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx093")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx094")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1E-8"))
	if err != nil {
		lang.SayString("divx095")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("divx095")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1E-9"))
	if err != nil {
		lang.SayString("divx096")
	} else {
		if !(rexsult.ToString() == "1E+9") {
			lang.SayString("divx096")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1E-10"))
	if err != nil {
		lang.SayString("divx097")
	} else {
		if !(rexsult.ToString() == "1E+10") {
			lang.SayString("divx097")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1E-11"))
	if err != nil {
		lang.SayString("divx098")
	} else {
		if !(rexsult.ToString() == "1E+11") {
			lang.SayString("divx098")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1E-12"))
	if err != nil {
		lang.SayString("divx099")
	} else {
		if !(rexsult.ToString() == "1E+12") {
			lang.SayString("divx099")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx100")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx100")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx101")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx101")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("divx102")
	} else {
		if !(rexsult.ToString() == "0.333333333") {
			lang.SayString("divx102")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("divx103")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("divx103")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("divx104")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("divx104")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("divx105")
	} else {
		if !(rexsult.ToString() == "0.166666667") {
			lang.SayString("divx105")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("divx106")
	} else {
		if !(rexsult.ToString() == "0.142857143") {
			lang.SayString("divx106")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("divx107")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("divx107")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("divx108")
	} else {
		if !(rexsult.ToString() == "0.111111111") {
			lang.SayString("divx108")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("divx109")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("divx109")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx110")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx110")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx111")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx111")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx112")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("divx112")
		}
	}
	rexsult, err = lang.RxFromString("4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx113")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("divx113")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx114")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("divx114")
		}
	}
	rexsult, err = lang.RxFromString("6").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx115")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("divx115")
		}
	}
	rexsult, err = lang.RxFromString("7").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx116")
	} else {
		if !(rexsult.ToString() == "7") {
			lang.SayString("divx116")
		}
	}
	rexsult, err = lang.RxFromString("8").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx117")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("divx117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx118")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("divx118")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx119")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("divx119")
		}
	}
	rexsult, err = lang.RxFromString("3E+1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("0.001"))
	if err != nil {
		lang.SayString("divx120")
	} else {
		if !(rexsult.ToString() == "30000") {
			lang.SayString("divx120")
		}
	}
	rexsult, err = lang.RxFromString("2.200").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx121")
	} else {
		if !(rexsult.ToString() == "1.1") {
			lang.SayString("divx121")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("divx130")
	} else {
		if !(rexsult.ToString() == "2469.4939") {
			lang.SayString("divx130")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("divx131")
	} else {
		if !(rexsult.ToString() == "2473.9479") {
			lang.SayString("divx131")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("divx132")
	} else {
		if !(rexsult.ToString() == "2519.38776") {
			lang.SayString("divx132")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("divx133")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("divx133")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("divx134")
	} else {
		if !(rexsult.ToString() == "2420.58824") {
			lang.SayString("divx134")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("divx135")
	} else {
		if !(rexsult.ToString() == "2464.07186") {
			lang.SayString("divx135")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("divx136")
	} else {
		if !(rexsult.ToString() == "2468.5063") {
			lang.SayString("divx136")
		}
	}
	rexsult, err = lang.RxFromString("391").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("597"))
	if err != nil {
		lang.SayString("divx220")
	} else {
		if !(rexsult.ToString() == "0.654941374") {
			lang.SayString("divx220")
		}
	}
	rexsult, err = lang.RxFromString("391").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-597"))
	if err != nil {
		lang.SayString("divx221")
	} else {
		if !(rexsult.ToString() == "-0.654941374") {
			lang.SayString("divx221")
		}
	}
	rexsult, err = lang.RxFromString("-391").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("597"))
	if err != nil {
		lang.SayString("divx222")
	} else {
		if !(rexsult.ToString() == "-0.654941374") {
			lang.SayString("divx222")
		}
	}
	rexsult, err = lang.RxFromString("-391").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-597"))
	if err != nil {
		lang.SayString("divx223")
	} else {
		if !(rexsult.ToString() == "0.654941374") {
			lang.SayString("divx223")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9e-999999999"))
	if err != nil {
		lang.SayString("divx280")
	} else {
		if !(rexsult.ToString() == "1.11111111E+999999997") {
			lang.SayString("divx280")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("99e-999999999"))
	if err != nil {
		lang.SayString("divx281")
	} else {
		if !(rexsult.ToString() == "1.01010101E+999999996") {
			lang.SayString("divx281")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("999e-999999999"))
	if err != nil {
		lang.SayString("divx282")
	} else {
		if !(rexsult.ToString() == "1.001001E+999999995") {
			lang.SayString("divx282")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9e-999999998"))
	if err != nil {
		lang.SayString("divx283")
	} else {
		if !(rexsult.ToString() == "1.11111111E+999999996") {
			lang.SayString("divx283")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("99e-999999998"))
	if err != nil {
		lang.SayString("divx284")
	} else {
		if !(rexsult.ToString() == "1.01010101E+999999995") {
			lang.SayString("divx284")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("999e-999999998"))
	if err != nil {
		lang.SayString("divx285")
	} else {
		if !(rexsult.ToString() == "1.001001E+999999994") {
			lang.SayString("divx285")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("999e-999999997"))
	if err != nil {
		lang.SayString("divx286")
	} else {
		if !(rexsult.ToString() == "1.001001E+999999993") {
			lang.SayString("divx286")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9999e-999999997"))
	if err != nil {
		lang.SayString("divx287")
	} else {
		if !(rexsult.ToString() == "1.00010001E+999999992") {
			lang.SayString("divx287")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("99999e-999999997"))
	if err != nil {
		lang.SayString("divx288")
	} else {
		if !(rexsult.ToString() == "1.00001E+999999991") {
			lang.SayString("divx288")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("divx301")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx301")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("divx302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx302")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("divx303")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx303")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("divx304")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx304")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("divx305")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx305")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+6"))
	if err != nil {
		lang.SayString("divx306")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx306")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+7"))
	if err != nil {
		lang.SayString("divx307")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx307")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("70E-5"))
	if err != nil {
		lang.SayString("divx308")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx308")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("70E-1"))
	if err != nil {
		lang.SayString("divx309")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx309")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("70E+0"))
	if err != nil {
		lang.SayString("divx310")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx310")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("70E+1"))
	if err != nil {
		lang.SayString("divx311")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx311")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("70E+5"))
	if err != nil {
		lang.SayString("divx312")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx312")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("70E+6"))
	if err != nil {
		lang.SayString("divx313")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx313")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("70E+7"))
	if err != nil {
		lang.SayString("divx314")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx314")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("700E-5"))
	if err != nil {
		lang.SayString("divx315")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx315")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("700E-1"))
	if err != nil {
		lang.SayString("divx316")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx316")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("700E+0"))
	if err != nil {
		lang.SayString("divx317")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx317")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("700E+1"))
	if err != nil {
		lang.SayString("divx318")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx318")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("700E+5"))
	if err != nil {
		lang.SayString("divx319")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx319")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("700E+6"))
	if err != nil {
		lang.SayString("divx320")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx320")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("700E+7"))
	if err != nil {
		lang.SayString("divx321")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx321")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("700E+77"))
	if err != nil {
		lang.SayString("divx322")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx322")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("divx331")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx331")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("divx332")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx332")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("divx333")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx333")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("divx334")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx334")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("divx335")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx335")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("divx336")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx336")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("divx337")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx337")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("divx338")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx338")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("divx339")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx339")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("divx340")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx340")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("divx341")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx341")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("divx342")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx342")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("divx343")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx343")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("divx344")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx344")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("divx345")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx345")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("divx346")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx346")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("divx351")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx351")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("divx352")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx352")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("divx353")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx353")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E+6"))
	if err != nil {
		lang.SayString("divx354")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx354")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E+7"))
	if err != nil {
		lang.SayString("divx355")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx355")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E-1"))
	if err != nil {
		lang.SayString("divx356")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx356")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E+1"))
	if err != nil {
		lang.SayString("divx357")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx357")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E+3"))
	if err != nil {
		lang.SayString("divx358")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx358")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E+4"))
	if err != nil {
		lang.SayString("divx359")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx359")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E+5"))
	if err != nil {
		lang.SayString("divx360")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx360")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E+6"))
	if err != nil {
		lang.SayString("divx361")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx361")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E+7"))
	if err != nil {
		lang.SayString("divx362")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx362")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E+92"))
	if err != nil {
		lang.SayString("divx363")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx363")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E-1"))
	if err != nil {
		lang.SayString("divx371")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx371")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E+1"))
	if err != nil {
		lang.SayString("divx372")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx372")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E+3"))
	if err != nil {
		lang.SayString("divx373")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx373")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E+4"))
	if err != nil {
		lang.SayString("divx374")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx374")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E+5"))
	if err != nil {
		lang.SayString("divx375")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx375")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E+6"))
	if err != nil {
		lang.SayString("divx376")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx376")
		}
	}
	rexsult, err = lang.RxFromString("0E-92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E+7"))
	if err != nil {
		lang.SayString("divx377")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx377")
		}
	}
	rexsult, err = lang.RxFromString("0E+92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("divx381")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx381")
		}
	}
	rexsult, err = lang.RxFromString("0E+92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E+0"))
	if err != nil {
		lang.SayString("divx382")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx382")
		}
	}
	rexsult, err = lang.RxFromString("0E+92").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("divx383")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx383")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E+1"))
	if err != nil {
		lang.SayString("divx384")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx384")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E-1"))
	if err != nil {
		lang.SayString("divx385")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx385")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E-2"))
	if err != nil {
		lang.SayString("divx386")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx386")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E-3"))
	if err != nil {
		lang.SayString("divx387")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx387")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("777E-4"))
	if err != nil {
		lang.SayString("divx388")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx388")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E+1"))
	if err != nil {
		lang.SayString("divx391")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx391")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E-1"))
	if err != nil {
		lang.SayString("divx392")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx392")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E-2"))
	if err != nil {
		lang.SayString("divx393")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx393")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E-3"))
	if err != nil {
		lang.SayString("divx394")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx394")
		}
	}
	rexsult, err = lang.RxFromString("0E+90").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("700E-4"))
	if err != nil {
		lang.SayString("divx395")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx395")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx401")
	} else {
		if !(rexsult.ToString() == "1.2345678E+10") {
			lang.SayString("divx401")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("divx402")
	} else {
		if !(rexsult.ToString() == "8.10000066E-11") {
			lang.SayString("divx402")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx403")
	} else {
		if !(rexsult.ToString() == "1.2345678E+9") {
			lang.SayString("divx403")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("divx404")
	} else {
		if !(rexsult.ToString() == "8.10000066E-10") {
			lang.SayString("divx404")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx405")
	} else {
		if !(rexsult.ToString() == "1.23456789E+9") {
			lang.SayString("divx405")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("divx406")
	} else {
		if !(rexsult.ToString() == "8.10000007E-10") {
			lang.SayString("divx406")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx407")
	} else {
		if !(rexsult.ToString() == "1.23456789E+9") {
			lang.SayString("divx407")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("divx408")
	} else {
		if !(rexsult.ToString() == "8.10000007E-10") {
			lang.SayString("divx408")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx409")
	} else {
		if !(rexsult.ToString() == "1.23456789E+10") {
			lang.SayString("divx409")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("divx410")
	} else {
		if !(rexsult.ToString() == "8.10000007E-11") {
			lang.SayString("divx410")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx411")
	} else {
		if !(rexsult.ToString() == "1.2345679E+9") {
			lang.SayString("divx411")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("divx412")
	} else {
		if !(rexsult.ToString() == "8.10000003E-10") {
			lang.SayString("divx412")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567897"))
	if err != nil {
		lang.SayString("divx413")
	} else {
		if !(rexsult.ToString() == "8.10000003E-10") {
			lang.SayString("divx413")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567898"))
	if err != nil {
		lang.SayString("divx414")
	} else {
		if !(rexsult.ToString() == "8.10000002E-10") {
			lang.SayString("divx414")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567899"))
	if err != nil {
		lang.SayString("divx415")
	} else {
		if !(rexsult.ToString() == "8.10000001E-10") {
			lang.SayString("divx415")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567900"))
	if err != nil {
		lang.SayString("divx416")
	} else {
		if !(rexsult.ToString() == "8.10000001E-10") {
			lang.SayString("divx416")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567901"))
	if err != nil {
		lang.SayString("divx417")
	} else {
		if !(rexsult.ToString() == "8.1E-10") {
			lang.SayString("divx417")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567902"))
	if err != nil {
		lang.SayString("divx418")
	} else {
		if !(rexsult.ToString() == "8.09999999E-10") {
			lang.SayString("divx418")
		}
	}
	rexsult, err = lang.RxFromString("1234567896.000000000000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx421")
	} else {
		if !(rexsult.ToString() == "1.2345679E+9") {
			lang.SayString("divx421")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567896.000000000000"))
	if err != nil {
		lang.SayString("divx422")
	} else {
		if !(rexsult.ToString() == "8.10000003E-10") {
			lang.SayString("divx422")
		}
	}
	rexsult, err = lang.RxFromString("1234567896.000000000001").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx423")
	} else {
		if !(rexsult.ToString() == "1.2345679E+9") {
			lang.SayString("divx423")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567896.000000000001"))
	if err != nil {
		lang.SayString("divx424")
	} else {
		if !(rexsult.ToString() == "8.10000003E-10") {
			lang.SayString("divx424")
		}
	}
	rexsult, err = lang.RxFromString("1234567896.000000000000000000000000000000000000000009").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx425")
	} else {
		if !(rexsult.ToString() == "1.2345679E+9") {
			lang.SayString("divx425")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567896.000000000000000000000000000000000000000009"))
	if err != nil {
		lang.SayString("divx426")
	} else {
		if !(rexsult.ToString() == "8.10000003E-10") {
			lang.SayString("divx426")
		}
	}
	rexsult, err = lang.RxFromString("1234567897.900010000000000000000000000000000000000009").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx427")
	} else {
		if !(rexsult.ToString() == "1.2345679E+9") {
			lang.SayString("divx427")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1234567897.900010000000000000000000000000000000000009"))
	if err != nil {
		lang.SayString("divx428")
	} else {
		if !(rexsult.ToString() == "8.10000003E-10") {
			lang.SayString("divx428")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx441")
	} else {
		if !(rexsult.ToString() == "12345678000") {
			lang.SayString("divx441")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("divx442")
	} else {
		if !(rexsult.ToString() == "8.10000066420005E-11") {
			lang.SayString("divx442")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx443")
	} else {
		if !(rexsult.ToString() == "1234567800") {
			lang.SayString("divx443")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("divx444")
	} else {
		if !(rexsult.ToString() == "8.10000066420005E-10") {
			lang.SayString("divx444")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx445")
	} else {
		if !(rexsult.ToString() == "1234567890") {
			lang.SayString("divx445")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("divx446")
	} else {
		if !(rexsult.ToString() == "8.10000007371E-10") {
			lang.SayString("divx446")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx447")
	} else {
		if !(rexsult.ToString() == "1234567891") {
			lang.SayString("divx447")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("divx448")
	} else {
		if !(rexsult.ToString() == "8.100000067149E-10") {
			lang.SayString("divx448")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx449")
	} else {
		if !(rexsult.ToString() == "12345678901") {
			lang.SayString("divx449")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("divx450")
	} else {
		if !(rexsult.ToString() == "8.1000000730539E-11") {
			lang.SayString("divx450")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx451")
	} else {
		if !(rexsult.ToString() == "1234567896") {
			lang.SayString("divx451")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("divx452")
	} else {
		if !(rexsult.ToString() == "8.100000034344E-10") {
			lang.SayString("divx452")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx453")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("divx453")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("divx454")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("divx454")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("divx455")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("divx455")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx456")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("divx456")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("divx457")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("divx457")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("divx458")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("divx458")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("divx466")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx466")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("divx467")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx467")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("divx468")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx468")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("divx470")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx470")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("divx471")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx471")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("divx472")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx472")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("divx474")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx474")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("divx475")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx475")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("divx476")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("divx476")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1.0E+33"))
	if err != nil {
		lang.SayString("divx480")
	} else {
		if !(rexsult.ToString() == "1E-33") {
			lang.SayString("divx480")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("10E+33"))
	if err != nil {
		lang.SayString("divx481")
	} else {
		if !(rexsult.ToString() == "1E-34") {
			lang.SayString("divx481")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("1.0E-33"))
	if err != nil {
		lang.SayString("divx482")
	} else {
		if !(rexsult.ToString() == "1E+33") {
			lang.SayString("divx482")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("10E-33"))
	if err != nil {
		lang.SayString("divx483")
	} else {
		if !(rexsult.ToString() == "1E+32") {
			lang.SayString("divx483")
		}
	}
	rexsult, err = lang.RxFromString("0E+86").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("1000E-13"))
	if err != nil {
		lang.SayString("divx497")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx497")
		}
	}
	rexsult, err = lang.RxFromString("0E-98").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("1000E+13"))
	if err != nil {
		lang.SayString("divx498")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx498")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9.9"))
	if err != nil {
		lang.SayString("divx500")
	} else {
		if !(rexsult.ToString() == "0.101010101") {
			lang.SayString("divx500")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(8), lang.RxFromString("9.9"))
	if err != nil {
		lang.SayString("divx501")
	} else {
		if !(rexsult.ToString() == "0.1010101") {
			lang.SayString("divx501")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("9.9"))
	if err != nil {
		lang.SayString("divx502")
	} else {
		if !(rexsult.ToString() == "0.1010101") {
			lang.SayString("divx502")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("9.9"))
	if err != nil {
		lang.SayString("divx503")
	} else {
		if !(rexsult.ToString() == "0.10101") {
			lang.SayString("divx503")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx511")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx511")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx512")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx512")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx513")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx513")
		}
	}
	rexsult, err = lang.RxFromString("1.000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx514")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx514")
		}
	}
	rexsult, err = lang.RxFromString("1.0000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx515")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx515")
		}
	}
	rexsult, err = lang.RxFromString("1.00000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx516")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx516")
		}
	}
	rexsult, err = lang.RxFromString("1.000000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx517")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx517")
		}
	}
	rexsult, err = lang.RxFromString("1.0000000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx518")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx518")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("divx519")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("divx519")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx521")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx521")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("divx522")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx522")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("divx523")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx523")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.000"))
	if err != nil {
		lang.SayString("divx524")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx524")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.0000"))
	if err != nil {
		lang.SayString("divx525")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx525")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.00000"))
	if err != nil {
		lang.SayString("divx526")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx526")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.000000"))
	if err != nil {
		lang.SayString("divx527")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx527")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.0000000"))
	if err != nil {
		lang.SayString("divx528")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx528")
		}
	}
	rexsult, err = lang.RxFromString("2.00").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("divx529")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("divx529")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx530")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("divx530")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("divx531")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("divx531")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("divx532")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("divx532")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("divx533")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("divx533")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4.0"))
	if err != nil {
		lang.SayString("divx534")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("divx534")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("10.0"))
	if err != nil {
		lang.SayString("divx535")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("divx535")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("divx536")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("divx536")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4.00"))
	if err != nil {
		lang.SayString("divx537")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("divx537")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("10.00"))
	if err != nil {
		lang.SayString("divx538")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("divx538")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("divx539")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("divx539")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("0.01"))
	if err != nil {
		lang.SayString("divx540")
	} else {
		if !(rexsult.ToString() == "90") {
			lang.SayString("divx540")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("0.001"))
	if err != nil {
		lang.SayString("divx541")
	} else {
		if !(rexsult.ToString() == "900") {
			lang.SayString("divx541")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx542")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx542")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("divx543")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx543")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("divx544")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("divx544")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("divx545")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("divx545")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("20.0"))
	if err != nil {
		lang.SayString("divx546")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("divx546")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx547")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("divx547")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("divx548")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("divx548")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.400"))
	if err != nil {
		lang.SayString("divx549")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx549")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx550")
	} else {
		if !(rexsult.ToString() == "240") {
			lang.SayString("divx550")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("divx551")
	} else {
		if !(rexsult.ToString() == "24") {
			lang.SayString("divx551")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("divx552")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("divx552")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("divx553")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("divx553")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx554")
	} else {
		if !(rexsult.ToString() == "2400") {
			lang.SayString("divx554")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("divx555")
	} else {
		if !(rexsult.ToString() == "240") {
			lang.SayString("divx555")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("divx556")
	} else {
		if !(rexsult.ToString() == "24") {
			lang.SayString("divx556")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("divx557")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("divx557")
		}
	}
	rexsult, err = lang.RxFromString("2.4E+6").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx570")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx570")
		}
	}
	rexsult, err = lang.RxFromString("2.40E+6").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx571")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx571")
		}
	}
	rexsult, err = lang.RxFromString("2.400E+6").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx572")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx572")
		}
	}
	rexsult, err = lang.RxFromString("2.4000E+6").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx573")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx573")
		}
	}
	rexsult, err = lang.RxFromString("24E+5").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx574")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx574")
		}
	}
	rexsult, err = lang.RxFromString("240E+4").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx575")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx575")
		}
	}
	rexsult, err = lang.RxFromString("2400E+3").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx576")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx576")
		}
	}
	rexsult, err = lang.RxFromString("24000E+2").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx577")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx577")
		}
	}
	rexsult, err = lang.RxFromString("2.4E+6").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx580")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx580")
		}
	}
	rexsult, err = lang.RxFromString("2.40E+6").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx581")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx581")
		}
	}
	rexsult, err = lang.RxFromString("2.400E+6").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx582")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx582")
		}
	}
	rexsult, err = lang.RxFromString("2.4000E+6").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx583")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx583")
		}
	}
	rexsult, err = lang.RxFromString("24E+5").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx584")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx584")
		}
	}
	rexsult, err = lang.RxFromString("240E+4").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx585")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx585")
		}
	}
	rexsult, err = lang.RxFromString("2400E+3").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx586")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx586")
		}
	}
	rexsult, err = lang.RxFromString("24000E+2").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx587")
	} else {
		if !(rexsult.ToString() == "1.2E+6") {
			lang.SayString("divx587")
		}
	}
	rexsult, err = lang.RxFromString("2.4E+6").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx590")
	} else {
		if !(rexsult.ToString() == "1200000") {
			lang.SayString("divx590")
		}
	}
	rexsult, err = lang.RxFromString("2.40E+6").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx591")
	} else {
		if !(rexsult.ToString() == "1200000") {
			lang.SayString("divx591")
		}
	}
	rexsult, err = lang.RxFromString("2.400E+6").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx592")
	} else {
		if !(rexsult.ToString() == "1200000") {
			lang.SayString("divx592")
		}
	}
	rexsult, err = lang.RxFromString("2.4000E+6").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx593")
	} else {
		if !(rexsult.ToString() == "1200000") {
			lang.SayString("divx593")
		}
	}
	rexsult, err = lang.RxFromString("24E+5").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx594")
	} else {
		if !(rexsult.ToString() == "1200000") {
			lang.SayString("divx594")
		}
	}
	rexsult, err = lang.RxFromString("240E+4").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx595")
	} else {
		if !(rexsult.ToString() == "1200000") {
			lang.SayString("divx595")
		}
	}
	rexsult, err = lang.RxFromString("2400E+3").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx596")
	} else {
		if !(rexsult.ToString() == "1200000") {
			lang.SayString("divx596")
		}
	}
	rexsult, err = lang.RxFromString("24000E+2").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx597")
	} else {
		if !(rexsult.ToString() == "1200000") {
			lang.SayString("divx597")
		}
	}
	rexsult, err = lang.RxFromString("2.4E+9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx600")
	} else {
		if !(rexsult.ToString() == "1.2E+9") {
			lang.SayString("divx600")
		}
	}
	rexsult, err = lang.RxFromString("2.40E+9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx601")
	} else {
		if !(rexsult.ToString() == "1.2E+9") {
			lang.SayString("divx601")
		}
	}
	rexsult, err = lang.RxFromString("2.400E+9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx602")
	} else {
		if !(rexsult.ToString() == "1.2E+9") {
			lang.SayString("divx602")
		}
	}
	rexsult, err = lang.RxFromString("2.4000E+9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx603")
	} else {
		if !(rexsult.ToString() == "1.2E+9") {
			lang.SayString("divx603")
		}
	}
	rexsult, err = lang.RxFromString("24E+8").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx604")
	} else {
		if !(rexsult.ToString() == "1.2E+9") {
			lang.SayString("divx604")
		}
	}
	rexsult, err = lang.RxFromString("240E+7").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx605")
	} else {
		if !(rexsult.ToString() == "1.2E+9") {
			lang.SayString("divx605")
		}
	}
	rexsult, err = lang.RxFromString("2400E+6").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx606")
	} else {
		if !(rexsult.ToString() == "1.2E+9") {
			lang.SayString("divx606")
		}
	}
	rexsult, err = lang.RxFromString("24000E+5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("divx607")
	} else {
		if !(rexsult.ToString() == "1.2E+9") {
			lang.SayString("divx607")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx610")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.8131097703792") {
			lang.SayString("divx610")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx611")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.813109770379") {
			lang.SayString("divx611")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx612")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.81310977038") {
			lang.SayString("divx612")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(30), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx613")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.8131097704") {
			lang.SayString("divx613")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(29), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx614")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.81310977") {
			lang.SayString("divx614")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(28), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx615")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.81310977") {
			lang.SayString("divx615")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(27), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx616")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.8131098") {
			lang.SayString("divx616")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(26), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx617")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.81311") {
			lang.SayString("divx617")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(25), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx618")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.81311") {
			lang.SayString("divx618")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(24), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx619")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.8131") {
			lang.SayString("divx619")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(23), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx620")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.813") {
			lang.SayString("divx620")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(22), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx621")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.81") {
			lang.SayString("divx621")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(21), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx622")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.8") {
			lang.SayString("divx622")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(20), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx623")
	} else {
		if !(rexsult.ToString() == "-41011408883796817798") {
			lang.SayString("divx623")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(19), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx624")
	} else {
		if !(rexsult.ToString() == "-4.10114088837968178E+19") {
			lang.SayString("divx624")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(18), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx625")
	} else {
		if !(rexsult.ToString() == "-4.10114088837968178E+19") {
			lang.SayString("divx625")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(17), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx626")
	} else {
		if !(rexsult.ToString() == "-4.1011408883796818E+19") {
			lang.SayString("divx626")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx627")
	} else {
		if !(rexsult.ToString() == "-4.101140888379682E+19") {
			lang.SayString("divx627")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(15), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx628")
	} else {
		if !(rexsult.ToString() == "-4.10114088837968E+19") {
			lang.SayString("divx628")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(14), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx629")
	} else {
		if !(rexsult.ToString() == "-4.1011408883797E+19") {
			lang.SayString("divx629")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(13), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx630")
	} else {
		if !(rexsult.ToString() == "-4.10114088838E+19") {
			lang.SayString("divx630")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(12), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx631")
	} else {
		if !(rexsult.ToString() == "-4.10114088838E+19") {
			lang.SayString("divx631")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(11), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx632")
	} else {
		if !(rexsult.ToString() == "-4.1011408884E+19") {
			lang.SayString("divx632")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(10), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx633")
	} else {
		if !(rexsult.ToString() == "-4.101140888E+19") {
			lang.SayString("divx633")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx634")
	} else {
		if !(rexsult.ToString() == "-4.10114089E+19") {
			lang.SayString("divx634")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(8), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx635")
	} else {
		if !(rexsult.ToString() == "-4.1011409E+19") {
			lang.SayString("divx635")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx636")
	} else {
		if !(rexsult.ToString() == "-4.101141E+19") {
			lang.SayString("divx636")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(6), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx637")
	} else {
		if !(rexsult.ToString() == "-4.10114E+19") {
			lang.SayString("divx637")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx638")
	} else {
		if !(rexsult.ToString() == "-4.1011E+19") {
			lang.SayString("divx638")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx639")
	} else {
		if !(rexsult.ToString() == "-4.101E+19") {
			lang.SayString("divx639")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx640")
	} else {
		if !(rexsult.ToString() == "-4.1E+19") {
			lang.SayString("divx640")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx641")
	} else {
		if !(rexsult.ToString() == "-4.1E+19") {
			lang.SayString("divx641")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx642")
	} else {
		if !(rexsult.ToString() == "-4E+19") {
			lang.SayString("divx642")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("divx731")
	} else {
		if !(rexsult.ToString() == "5000") {
			lang.SayString("divx731")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("divx741")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx741")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("divx742")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx742")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx743")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx743")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx744")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx744")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("divx751")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx751")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("divx752")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx752")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx753")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx753")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx754")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx754")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("divx761")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx761")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("divx762")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx762")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("divx763")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx763")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("divx764")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx764")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("divx771")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx771")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("divx772")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx772")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("divx773")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx773")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(16), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("divx774")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("divx774")
		}
	}
	rexsult, err = lang.RxFromString("1.52444E-80").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx1001")
	} else {
		if !(rexsult.ToString() == "1.5244E-80") {
			lang.SayString("divx1001")
		}
	}
	rexsult, err = lang.RxFromString("1.52445E-80").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx1002")
	} else {
		if !(rexsult.ToString() == "1.5245E-80") {
			lang.SayString("divx1002")
		}
	}
	rexsult, err = lang.RxFromString("1.52446E-80").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("divx1003")
	} else {
		if !(rexsult.ToString() == "1.5245E-80") {
			lang.SayString("divx1003")
		}
	}
	rexsult, err = lang.RxFromString("100E-2").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("1000E-3"))
	if err != nil {
		lang.SayString("divx1024")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("divx1024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("divx1050")
	} else {
		if !(rexsult.ToString() == "0.5555556") {
			lang.SayString("divx1050")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(7), lang.RxFromString("11"))
	if err != nil {
		lang.SayString("divx1051")
	} else {
		if !(rexsult.ToString() == "0.4545455") {
			lang.SayString("divx1051")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv001")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv001")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv002")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv002")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv003")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv003")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv004")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv005")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv005")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv006")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv006")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdiv007")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333333333333333333333") {
			lang.SayString("dqdiv007")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdiv008")
	} else {
		if !(rexsult.ToString() == "0.6666666666666666666666666666666667") {
			lang.SayString("dqdiv008")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdiv009")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv009")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv010")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dqdiv010")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdiv011")
	} else {
		if !(rexsult.ToString() == "-2.4") {
			lang.SayString("dqdiv011")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv012")
	} else {
		if !(rexsult.ToString() == "-2.4") {
			lang.SayString("dqdiv012")
		}
	}
	rexsult, err = lang.RxFromString("-2.4").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdiv013")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dqdiv013")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv014")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dqdiv014")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv015")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dqdiv015")
		}
	}
	rexsult, err = lang.RxFromString("2.4").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv016")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dqdiv016")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv017")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dqdiv017")
		}
	}
	rexsult, err = lang.RxFromString("2.").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv018")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv018")
		}
	}
	rexsult, err = lang.RxFromString("20").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dqdiv019")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv019")
		}
	}
	rexsult, err = lang.RxFromString("187").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("187"))
	if err != nil {
		lang.SayString("dqdiv020")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv020")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv021")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv021")
		}
	}
	rexsult, err = lang.RxFromString("50").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dqdiv022")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv022")
		}
	}
	rexsult, err = lang.RxFromString("500").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("200"))
	if err != nil {
		lang.SayString("dqdiv023")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv023")
		}
	}
	rexsult, err = lang.RxFromString("50.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("20.0"))
	if err != nil {
		lang.SayString("dqdiv024")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv024")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dqdiv025")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv025")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqdiv026")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv026")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.000"))
	if err != nil {
		lang.SayString("dqdiv027")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv027")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("0.20"))
	if err != nil {
		lang.SayString("dqdiv028")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dqdiv028")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("0.200"))
	if err != nil {
		lang.SayString("dqdiv029")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("dqdiv029")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv030")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqdiv030")
		}
	}
	rexsult, err = lang.RxFromString("100").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv031")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("dqdiv031")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv032")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("dqdiv032")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdiv033")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqdiv033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv035")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqdiv036")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dqdiv036")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqdiv037")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("dqdiv037")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dqdiv038")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("dqdiv038")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dqdiv039")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("dqdiv039")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dqdiv040")
	} else {
		if !(rexsult.ToString() == "0.015625") {
			lang.SayString("dqdiv040")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqdiv041")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dqdiv041")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dqdiv042")
	} else {
		if !(rexsult.ToString() == "-0.25") {
			lang.SayString("dqdiv042")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dqdiv043")
	} else {
		if !(rexsult.ToString() == "-0.125") {
			lang.SayString("dqdiv043")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dqdiv044")
	} else {
		if !(rexsult.ToString() == "-0.0625") {
			lang.SayString("dqdiv044")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dqdiv045")
	} else {
		if !(rexsult.ToString() == "-0.03125") {
			lang.SayString("dqdiv045")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dqdiv046")
	} else {
		if !(rexsult.ToString() == "-0.015625") {
			lang.SayString("dqdiv046")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv047")
	} else {
		if !(rexsult.ToString() == "-0.5") {
			lang.SayString("dqdiv047")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqdiv048")
	} else {
		if !(rexsult.ToString() == "-0.25") {
			lang.SayString("dqdiv048")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqdiv049")
	} else {
		if !(rexsult.ToString() == "-0.125") {
			lang.SayString("dqdiv049")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("dqdiv050")
	} else {
		if !(rexsult.ToString() == "-0.0625") {
			lang.SayString("dqdiv050")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("dqdiv051")
	} else {
		if !(rexsult.ToString() == "-0.03125") {
			lang.SayString("dqdiv051")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("dqdiv052")
	} else {
		if !(rexsult.ToString() == "-0.015625") {
			lang.SayString("dqdiv052")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("dqdiv053")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv053")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("dqdiv054")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dqdiv054")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("dqdiv055")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("dqdiv055")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("dqdiv056")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("dqdiv056")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("dqdiv057")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("dqdiv057")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("dqdiv058")
	} else {
		if !(rexsult.ToString() == "0.015625") {
			lang.SayString("dqdiv058")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dqdiv060")
	} else {
		if !(rexsult.ToString() == "0.1428571428571428571428571428571429") {
			lang.SayString("dqdiv060")
		}
	}
	rexsult, err = lang.RxFromString("1.2345678").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.9876543"))
	if err != nil {
		lang.SayString("dqdiv061")
	} else {
		if !(rexsult.ToString() == "0.6211179680490717123193907511985359") {
			lang.SayString("dqdiv061")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999999999999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv067")
	} else {
		if !(rexsult.ToString() == "9999999999999999999999999999999999") {
			lang.SayString("dqdiv067")
		}
	}
	rexsult, err = lang.RxFromString("999999999999999999999999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv068")
	} else {
		if !(rexsult.ToString() == "999999999999999999999999999999999") {
			lang.SayString("dqdiv068")
		}
	}
	rexsult, err = lang.RxFromString("99999999999999999999999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv069")
	} else {
		if !(rexsult.ToString() == "99999999999999999999999999999999") {
			lang.SayString("dqdiv069")
		}
	}
	rexsult, err = lang.RxFromString("99999999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv070")
	} else {
		if !(rexsult.ToString() == "99999999999999999") {
			lang.SayString("dqdiv070")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv071")
	} else {
		if !(rexsult.ToString() == "9999999999999999") {
			lang.SayString("dqdiv071")
		}
	}
	rexsult, err = lang.RxFromString("999999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv072")
	} else {
		if !(rexsult.ToString() == "999999999999999") {
			lang.SayString("dqdiv072")
		}
	}
	rexsult, err = lang.RxFromString("99999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv073")
	} else {
		if !(rexsult.ToString() == "99999999999999") {
			lang.SayString("dqdiv073")
		}
	}
	rexsult, err = lang.RxFromString("9999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv074")
	} else {
		if !(rexsult.ToString() == "9999999999999") {
			lang.SayString("dqdiv074")
		}
	}
	rexsult, err = lang.RxFromString("999999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv075")
	} else {
		if !(rexsult.ToString() == "999999999999") {
			lang.SayString("dqdiv075")
		}
	}
	rexsult, err = lang.RxFromString("99999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv076")
	} else {
		if !(rexsult.ToString() == "99999999999") {
			lang.SayString("dqdiv076")
		}
	}
	rexsult, err = lang.RxFromString("9999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv077")
	} else {
		if !(rexsult.ToString() == "9999999999") {
			lang.SayString("dqdiv077")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv078")
	} else {
		if !(rexsult.ToString() == "999999999") {
			lang.SayString("dqdiv078")
		}
	}
	rexsult, err = lang.RxFromString("99999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv079")
	} else {
		if !(rexsult.ToString() == "99999999") {
			lang.SayString("dqdiv079")
		}
	}
	rexsult, err = lang.RxFromString("9999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv080")
	} else {
		if !(rexsult.ToString() == "9999999") {
			lang.SayString("dqdiv080")
		}
	}
	rexsult, err = lang.RxFromString("999999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv081")
	} else {
		if !(rexsult.ToString() == "999999") {
			lang.SayString("dqdiv081")
		}
	}
	rexsult, err = lang.RxFromString("99999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv082")
	} else {
		if !(rexsult.ToString() == "99999") {
			lang.SayString("dqdiv082")
		}
	}
	rexsult, err = lang.RxFromString("9999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv083")
	} else {
		if !(rexsult.ToString() == "9999") {
			lang.SayString("dqdiv083")
		}
	}
	rexsult, err = lang.RxFromString("999").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv084")
	} else {
		if !(rexsult.ToString() == "999") {
			lang.SayString("dqdiv084")
		}
	}
	rexsult, err = lang.RxFromString("99").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv085")
	} else {
		if !(rexsult.ToString() == "99") {
			lang.SayString("dqdiv085")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv086")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dqdiv086")
		}
	}
	rexsult, err = lang.RxFromString("0.").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv090")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv090")
		}
	}
	rexsult, err = lang.RxFromString(".0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv091")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv091")
		}
	}
	rexsult, err = lang.RxFromString("0.00").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv092")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv092")
		}
	}
	rexsult, err = lang.RxFromString("0.00E+9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv093")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv093")
		}
	}
	rexsult, err = lang.RxFromString("0.0000E-50").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv094")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv094")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1E-8"))
	if err != nil {
		lang.SayString("dqdiv095")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("dqdiv095")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1E-9"))
	if err != nil {
		lang.SayString("dqdiv096")
	} else {
		if !(rexsult.ToString() == "1000000000") {
			lang.SayString("dqdiv096")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1E-10"))
	if err != nil {
		lang.SayString("dqdiv097")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("dqdiv097")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1E-11"))
	if err != nil {
		lang.SayString("dqdiv098")
	} else {
		if !(rexsult.ToString() == "100000000000") {
			lang.SayString("dqdiv098")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1E-12"))
	if err != nil {
		lang.SayString("dqdiv099")
	} else {
		if !(rexsult.ToString() == "1000000000000") {
			lang.SayString("dqdiv099")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv100")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv100")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv101")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv101")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("dqdiv102")
	} else {
		if !(rexsult.ToString() == "0.3333333333333333333333333333333333") {
			lang.SayString("dqdiv102")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqdiv103")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dqdiv103")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqdiv104")
	} else {
		if !(rexsult.ToString() == "0.2") {
			lang.SayString("dqdiv104")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("dqdiv105")
	} else {
		if !(rexsult.ToString() == "0.1666666666666666666666666666666667") {
			lang.SayString("dqdiv105")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dqdiv106")
	} else {
		if !(rexsult.ToString() == "0.1428571428571428571428571428571429") {
			lang.SayString("dqdiv106")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("dqdiv107")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("dqdiv107")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dqdiv108")
	} else {
		if !(rexsult.ToString() == "0.1111111111111111111111111111111111") {
			lang.SayString("dqdiv108")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdiv109")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("dqdiv109")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv110")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv110")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv111")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv111")
		}
	}
	rexsult, err = lang.RxFromString("3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv112")
	} else {
		if !(rexsult.ToString() == "3") {
			lang.SayString("dqdiv112")
		}
	}
	rexsult, err = lang.RxFromString("4").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv113")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("dqdiv113")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv114")
	} else {
		if !(rexsult.ToString() == "5") {
			lang.SayString("dqdiv114")
		}
	}
	rexsult, err = lang.RxFromString("6").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv115")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("dqdiv115")
		}
	}
	rexsult, err = lang.RxFromString("7").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv116")
	} else {
		if !(rexsult.ToString() == "7") {
			lang.SayString("dqdiv116")
		}
	}
	rexsult, err = lang.RxFromString("8").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv117")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("dqdiv117")
		}
	}
	rexsult, err = lang.RxFromString("9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv118")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dqdiv118")
		}
	}
	rexsult, err = lang.RxFromString("10").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv119")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqdiv119")
		}
	}
	rexsult, err = lang.RxFromString("3E+1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("0.001"))
	if err != nil {
		lang.SayString("dqdiv120")
	} else {
		if !(rexsult.ToString() == "30000") {
			lang.SayString("dqdiv120")
		}
	}
	rexsult, err = lang.RxFromString("2.200").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv121")
	} else {
		if !(rexsult.ToString() == "1.1") {
			lang.SayString("dqdiv121")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("dqdiv130")
	} else {
		if !(rexsult.ToString() == "2469.493898779755951190238047609522") {
			lang.SayString("dqdiv130")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("dqdiv131")
	} else {
		if !(rexsult.ToString() == "2473.947895791583166332665330661323") {
			lang.SayString("dqdiv131")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("dqdiv132")
	} else {
		if !(rexsult.ToString() == "2519.387755102040816326530612244898") {
			lang.SayString("dqdiv132")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("dqdiv133")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("dqdiv133")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("dqdiv134")
	} else {
		if !(rexsult.ToString() == "2420.588235294117647058823529411765") {
			lang.SayString("dqdiv134")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("dqdiv135")
	} else {
		if !(rexsult.ToString() == "2464.07185628742514970059880239521") {
			lang.SayString("dqdiv135")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("dqdiv136")
	} else {
		if !(rexsult.ToString() == "2468.506298740251949610077984403119") {
			lang.SayString("dqdiv136")
		}
	}
	rexsult, err = lang.RxFromString("391").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("597"))
	if err != nil {
		lang.SayString("dqdiv220")
	} else {
		if !(rexsult.ToString() == "0.6549413735343383584589614740368509") {
			lang.SayString("dqdiv220")
		}
	}
	rexsult, err = lang.RxFromString("391").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-597"))
	if err != nil {
		lang.SayString("dqdiv221")
	} else {
		if !(rexsult.ToString() == "-0.6549413735343383584589614740368509") {
			lang.SayString("dqdiv221")
		}
	}
	rexsult, err = lang.RxFromString("-391").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("597"))
	if err != nil {
		lang.SayString("dqdiv222")
	} else {
		if !(rexsult.ToString() == "-0.6549413735343383584589614740368509") {
			lang.SayString("dqdiv222")
		}
	}
	rexsult, err = lang.RxFromString("-391").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-597"))
	if err != nil {
		lang.SayString("dqdiv223")
	} else {
		if !(rexsult.ToString() == "0.6549413735343383584589614740368509") {
			lang.SayString("dqdiv223")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("dqdiv301")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv301")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dqdiv302")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv302")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dqdiv303")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv303")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dqdiv304")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv304")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dqdiv305")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv305")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+6"))
	if err != nil {
		lang.SayString("dqdiv306")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv306")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+7"))
	if err != nil {
		lang.SayString("dqdiv307")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv307")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("70E-5"))
	if err != nil {
		lang.SayString("dqdiv308")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv308")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("70E-1"))
	if err != nil {
		lang.SayString("dqdiv309")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv309")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("70E+0"))
	if err != nil {
		lang.SayString("dqdiv310")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv310")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("70E+1"))
	if err != nil {
		lang.SayString("dqdiv311")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv311")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("70E+5"))
	if err != nil {
		lang.SayString("dqdiv312")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv312")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("70E+6"))
	if err != nil {
		lang.SayString("dqdiv313")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv313")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("70E+7"))
	if err != nil {
		lang.SayString("dqdiv314")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv314")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("700E-5"))
	if err != nil {
		lang.SayString("dqdiv315")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv315")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("700E-1"))
	if err != nil {
		lang.SayString("dqdiv316")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv316")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("700E+0"))
	if err != nil {
		lang.SayString("dqdiv317")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv317")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("700E+1"))
	if err != nil {
		lang.SayString("dqdiv318")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv318")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("700E+5"))
	if err != nil {
		lang.SayString("dqdiv319")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv319")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("700E+6"))
	if err != nil {
		lang.SayString("dqdiv320")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv320")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("700E+7"))
	if err != nil {
		lang.SayString("dqdiv321")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv321")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("700E+77"))
	if err != nil {
		lang.SayString("dqdiv322")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv322")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dqdiv331")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv331")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dqdiv332")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv332")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dqdiv333")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv333")
		}
	}
	rexsult, err = lang.RxFromString("0E-3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dqdiv334")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv334")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dqdiv335")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv335")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dqdiv336")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv336")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dqdiv337")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv337")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dqdiv338")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv338")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dqdiv339")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv339")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dqdiv340")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv340")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dqdiv341")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv341")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dqdiv342")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv342")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-5"))
	if err != nil {
		lang.SayString("dqdiv343")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv343")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E-1"))
	if err != nil {
		lang.SayString("dqdiv344")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv344")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+1"))
	if err != nil {
		lang.SayString("dqdiv345")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv345")
		}
	}
	rexsult, err = lang.RxFromString("0E+3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7E+5"))
	if err != nil {
		lang.SayString("dqdiv346")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv346")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv441")
	} else {
		if !(rexsult.ToString() == "12345678000") {
			lang.SayString("dqdiv441")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("12345678000"))
	if err != nil {
		lang.SayString("dqdiv442")
	} else {
		if !(rexsult.ToString() == "8.100000664200054464404466081166219E-11") {
			lang.SayString("dqdiv442")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv443")
	} else {
		if !(rexsult.ToString() == "1234567800") {
			lang.SayString("dqdiv443")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1234567800"))
	if err != nil {
		lang.SayString("dqdiv444")
	} else {
		if !(rexsult.ToString() == "8.100000664200054464404466081166219E-10") {
			lang.SayString("dqdiv444")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv445")
	} else {
		if !(rexsult.ToString() == "1234567890") {
			lang.SayString("dqdiv445")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1234567890"))
	if err != nil {
		lang.SayString("dqdiv446")
	} else {
		if !(rexsult.ToString() == "8.100000073710000670761006103925156E-10") {
			lang.SayString("dqdiv446")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv447")
	} else {
		if !(rexsult.ToString() == "1234567891") {
			lang.SayString("dqdiv447")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1234567891"))
	if err != nil {
		lang.SayString("dqdiv448")
	} else {
		if !(rexsult.ToString() == "8.100000067149000556665214614754629E-10") {
			lang.SayString("dqdiv448")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv449")
	} else {
		if !(rexsult.ToString() == "12345678901") {
			lang.SayString("dqdiv449")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("12345678901"))
	if err != nil {
		lang.SayString("dqdiv450")
	} else {
		if !(rexsult.ToString() == "8.10000007305390065887313004237676E-11") {
			lang.SayString("dqdiv450")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv451")
	} else {
		if !(rexsult.ToString() == "1234567896") {
			lang.SayString("dqdiv451")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1234567896"))
	if err != nil {
		lang.SayString("dqdiv452")
	} else {
		if !(rexsult.ToString() == "8.100000034344000145618560617422697E-10") {
			lang.SayString("dqdiv452")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv453")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqdiv453")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdiv454")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqdiv454")
		}
	}
	rexsult, err = lang.RxFromString("1e+1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dqdiv455")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("dqdiv455")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv456")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("dqdiv456")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqdiv457")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("dqdiv457")
		}
	}
	rexsult, err = lang.RxFromString("1e+2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dqdiv458")
	} else {
		if !(rexsult.ToString() == "50") {
			lang.SayString("dqdiv458")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("dqdiv466")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv466")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("dqdiv467")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv467")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("20e-1"))
	if err != nil {
		lang.SayString("dqdiv468")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv468")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("dqdiv470")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv470")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("dqdiv471")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv471")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("200e-2"))
	if err != nil {
		lang.SayString("dqdiv472")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv472")
		}
	}
	rexsult, err = lang.RxFromString("30e-1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("dqdiv474")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv474")
		}
	}
	rexsult, err = lang.RxFromString("300e-2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("dqdiv475")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv475")
		}
	}
	rexsult, err = lang.RxFromString("3000e-3").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2000e-3"))
	if err != nil {
		lang.SayString("dqdiv476")
	} else {
		if !(rexsult.ToString() == "1.5") {
			lang.SayString("dqdiv476")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0E+33"))
	if err != nil {
		lang.SayString("dqdiv480")
	} else {
		if !(rexsult.ToString() == "1E-33") {
			lang.SayString("dqdiv480")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10E+33"))
	if err != nil {
		lang.SayString("dqdiv481")
	} else {
		if !(rexsult.ToString() == "1E-34") {
			lang.SayString("dqdiv481")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0E-33"))
	if err != nil {
		lang.SayString("dqdiv482")
	} else {
		if !(rexsult.ToString() == "1000000000000000000000000000000000") {
			lang.SayString("dqdiv482")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10E-33"))
	if err != nil {
		lang.SayString("dqdiv483")
	} else {
		if !(rexsult.ToString() == "100000000000000000000000000000000") {
			lang.SayString("dqdiv483")
		}
	}
	rexsult, err = lang.RxFromString("0E+6108").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1000E-33"))
	if err != nil {
		lang.SayString("dqdiv497")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv497")
		}
	}
	rexsult, err = lang.RxFromString("0E-6170").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1000E+33"))
	if err != nil {
		lang.SayString("dqdiv498")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv498")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9.9"))
	if err != nil {
		lang.SayString("dqdiv500")
	} else {
		if !(rexsult.ToString() == "0.101010101010101010101010101010101") {
			lang.SayString("dqdiv500")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9.09"))
	if err != nil {
		lang.SayString("dqdiv501")
	} else {
		if !(rexsult.ToString() == "0.1100110011001100110011001100110011") {
			lang.SayString("dqdiv501")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9.009"))
	if err != nil {
		lang.SayString("dqdiv502")
	} else {
		if !(rexsult.ToString() == "0.111000111000111000111000111000111") {
			lang.SayString("dqdiv502")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv511")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv511")
		}
	}
	rexsult, err = lang.RxFromString("1.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv512")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv512")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv513")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv513")
		}
	}
	rexsult, err = lang.RxFromString("1.000").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv514")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv514")
		}
	}
	rexsult, err = lang.RxFromString("1.0000").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv515")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv515")
		}
	}
	rexsult, err = lang.RxFromString("1.00000").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv516")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv516")
		}
	}
	rexsult, err = lang.RxFromString("1.000000").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv517")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv517")
		}
	}
	rexsult, err = lang.RxFromString("1.0000000").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv518")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv518")
		}
	}
	rexsult, err = lang.RxFromString("1.00").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dqdiv519")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("dqdiv519")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv521")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv521")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdiv522")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv522")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dqdiv523")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv523")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.000"))
	if err != nil {
		lang.SayString("dqdiv524")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv524")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0000"))
	if err != nil {
		lang.SayString("dqdiv525")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv525")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.00000"))
	if err != nil {
		lang.SayString("dqdiv526")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv526")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.000000"))
	if err != nil {
		lang.SayString("dqdiv527")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv527")
		}
	}
	rexsult, err = lang.RxFromString("2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0000000"))
	if err != nil {
		lang.SayString("dqdiv528")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv528")
		}
	}
	rexsult, err = lang.RxFromString("2.00").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("dqdiv529")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("dqdiv529")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv530")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dqdiv530")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("dqdiv531")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("dqdiv531")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdiv532")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("dqdiv532")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqdiv533")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dqdiv533")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4.0"))
	if err != nil {
		lang.SayString("dqdiv534")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("dqdiv534")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10.0"))
	if err != nil {
		lang.SayString("dqdiv535")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("dqdiv535")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dqdiv536")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dqdiv536")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4.00"))
	if err != nil {
		lang.SayString("dqdiv537")
	} else {
		if !(rexsult.ToString() == "0.6") {
			lang.SayString("dqdiv537")
		}
	}
	rexsult, err = lang.RxFromString("2.40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10.00"))
	if err != nil {
		lang.SayString("dqdiv538")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("dqdiv538")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("0.1"))
	if err != nil {
		lang.SayString("dqdiv539")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("dqdiv539")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("0.01"))
	if err != nil {
		lang.SayString("dqdiv540")
	} else {
		if !(rexsult.ToString() == "90") {
			lang.SayString("dqdiv540")
		}
	}
	rexsult, err = lang.RxFromString("0.9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("0.001"))
	if err != nil {
		lang.SayString("dqdiv541")
	} else {
		if !(rexsult.ToString() == "900") {
			lang.SayString("dqdiv541")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv542")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv542")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqdiv543")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv543")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("dqdiv544")
	} else {
		if !(rexsult.ToString() == "2.5") {
			lang.SayString("dqdiv544")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("dqdiv545")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dqdiv545")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("20.0"))
	if err != nil {
		lang.SayString("dqdiv546")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("dqdiv546")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv547")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dqdiv547")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.0"))
	if err != nil {
		lang.SayString("dqdiv548")
	} else {
		if !(rexsult.ToString() == "1.2") {
			lang.SayString("dqdiv548")
		}
	}
	rexsult, err = lang.RxFromString("2.400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2.400"))
	if err != nil {
		lang.SayString("dqdiv549")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv549")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv550")
	} else {
		if !(rexsult.ToString() == "240") {
			lang.SayString("dqdiv550")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdiv551")
	} else {
		if !(rexsult.ToString() == "24") {
			lang.SayString("dqdiv551")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdiv552")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dqdiv552")
		}
	}
	rexsult, err = lang.RxFromString("240").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqdiv553")
	} else {
		if !(rexsult.ToString() == "0.24") {
			lang.SayString("dqdiv553")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv554")
	} else {
		if !(rexsult.ToString() == "2400") {
			lang.SayString("dqdiv554")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("dqdiv555")
	} else {
		if !(rexsult.ToString() == "240") {
			lang.SayString("dqdiv555")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("100"))
	if err != nil {
		lang.SayString("dqdiv556")
	} else {
		if !(rexsult.ToString() == "24") {
			lang.SayString("dqdiv556")
		}
	}
	rexsult, err = lang.RxFromString("2400").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("dqdiv557")
	} else {
		if !(rexsult.ToString() == "2.4") {
			lang.SayString("dqdiv557")
		}
	}
	rexsult, err = lang.RxFromString("2.4E+9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv600")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dqdiv600")
		}
	}
	rexsult, err = lang.RxFromString("2.40E+9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv601")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dqdiv601")
		}
	}
	rexsult, err = lang.RxFromString("2.400E+9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv602")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dqdiv602")
		}
	}
	rexsult, err = lang.RxFromString("2.4000E+9").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv603")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dqdiv603")
		}
	}
	rexsult, err = lang.RxFromString("24E+8").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv604")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dqdiv604")
		}
	}
	rexsult, err = lang.RxFromString("240E+7").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv605")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dqdiv605")
		}
	}
	rexsult, err = lang.RxFromString("2400E+6").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv606")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dqdiv606")
		}
	}
	rexsult, err = lang.RxFromString("24000E+5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("dqdiv607")
	} else {
		if !(rexsult.ToString() == "1200000000") {
			lang.SayString("dqdiv607")
		}
	}
	rexsult, err = lang.RxFromString("5.00").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1E-3"))
	if err != nil {
		lang.SayString("dqdiv731")
	} else {
		if !(rexsult.ToString() == "5000") {
			lang.SayString("dqdiv731")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdiv741")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv741")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdiv742")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv742")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv743")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv743")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv744")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv744")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdiv751")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv751")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("dqdiv752")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv752")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv753")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv753")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv754")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv754")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqdiv761")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv761")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqdiv762")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv762")
		}
	}
	rexsult, err = lang.RxFromString("0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdiv763")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv763")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdiv764")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv764")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqdiv771")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv771")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1.0"))
	if err != nil {
		lang.SayString("dqdiv772")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv772")
		}
	}
	rexsult, err = lang.RxFromString("0.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdiv773")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv773")
		}
	}
	rexsult, err = lang.RxFromString("-0.0").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.0"))
	if err != nil {
		lang.SayString("dqdiv774")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("dqdiv774")
		}
	}
	rexsult, err = lang.RxFromString("100E-2").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1000E-3"))
	if err != nil {
		lang.SayString("dqdiv1024")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("dqdiv1024")
		}
	}
	rexsult, err = lang.RxFromString("5").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("dqdiv1040")
	} else {
		if !(rexsult.ToString() == "0.5555555555555555555555555555555556") {
			lang.SayString("dqdiv1040")
		}
	}
	rexsult, err = lang.RxFromString("6").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("11"))
	if err != nil {
		lang.SayString("dqdiv1041")
	} else {
		if !(rexsult.ToString() == "0.5454545454545454545454545454545455") {
			lang.SayString("dqdiv1041")
		}
	}
	rexsult, err = lang.RxFromString("8.336804418094040989630006819881709E-6143").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8.336804418094040989630006819889000E-6143"))
	if err != nil {
		lang.SayString("dqdiv1050")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999991254") {
			lang.SayString("dqdiv1050")
		}
	}
	rexsult, err = lang.RxFromString("1e+4277").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-3311"))
	if err != nil {
		lang.SayString("dqdiv1751")
	} else {
		if !(rexsult.ToString() == "1E+7588") {
			lang.SayString("dqdiv1751")
		}
	}
	rexsult, err = lang.RxFromString("1e+4277").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e-3311"))
	if err != nil {
		lang.SayString("dqdiv1752")
	} else {
		if !(rexsult.ToString() == "-1E+7588") {
			lang.SayString("dqdiv1752")
		}
	}
	rexsult, err = lang.RxFromString("-1e+4277").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-3311"))
	if err != nil {
		lang.SayString("dqdiv1753")
	} else {
		if !(rexsult.ToString() == "-1E+7588") {
			lang.SayString("dqdiv1753")
		}
	}
	rexsult, err = lang.RxFromString("-1e+4277").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e-3311"))
	if err != nil {
		lang.SayString("dqdiv1754")
	} else {
		if !(rexsult.ToString() == "1E+7588") {
			lang.SayString("dqdiv1754")
		}
	}
	rexsult, err = lang.RxFromString("1e-4277").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+3311"))
	if err != nil {
		lang.SayString("dqdiv1755")
	} else {
		if !(rexsult.ToString() == "1E-7588") {
			lang.SayString("dqdiv1755")
		}
	}
	rexsult, err = lang.RxFromString("1e-4277").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e+3311"))
	if err != nil {
		lang.SayString("dqdiv1756")
	} else {
		if !(rexsult.ToString() == "-1E-7588") {
			lang.SayString("dqdiv1756")
		}
	}
	rexsult, err = lang.RxFromString("-1e-4277").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+3311"))
	if err != nil {
		lang.SayString("dqdiv1757")
	} else {
		if !(rexsult.ToString() == "-1E-7588") {
			lang.SayString("dqdiv1757")
		}
	}
	rexsult, err = lang.RxFromString("-1e-4277").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e+3311"))
	if err != nil {
		lang.SayString("dqdiv1758")
	} else {
		if !(rexsult.ToString() == "1E-7588") {
			lang.SayString("dqdiv1758")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+101"))
	if err != nil {
		lang.SayString("dqdiv1760")
	} else {
		if !(rexsult.ToString() == "1E-6170") {
			lang.SayString("dqdiv1760")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+102"))
	if err != nil {
		lang.SayString("dqdiv1761")
	} else {
		if !(rexsult.ToString() == "1E-6171") {
			lang.SayString("dqdiv1761")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+103"))
	if err != nil {
		lang.SayString("dqdiv1762")
	} else {
		if !(rexsult.ToString() == "1E-6172") {
			lang.SayString("dqdiv1762")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+104"))
	if err != nil {
		lang.SayString("dqdiv1763")
	} else {
		if !(rexsult.ToString() == "1E-6173") {
			lang.SayString("dqdiv1763")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+105"))
	if err != nil {
		lang.SayString("dqdiv1764")
	} else {
		if !(rexsult.ToString() == "1E-6174") {
			lang.SayString("dqdiv1764")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+106"))
	if err != nil {
		lang.SayString("dqdiv1765")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqdiv1765")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+107"))
	if err != nil {
		lang.SayString("dqdiv1766")
	} else {
		if !(rexsult.ToString() == "1E-6176") {
			lang.SayString("dqdiv1766")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+108"))
	if err != nil {
		lang.SayString("dqdiv1767")
	} else {
		if !(rexsult.ToString() == "1E-6177") {
			lang.SayString("dqdiv1767")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+109"))
	if err != nil {
		lang.SayString("dqdiv1768")
	} else {
		if !(rexsult.ToString() == "1E-6178") {
			lang.SayString("dqdiv1768")
		}
	}
	rexsult, err = lang.RxFromString("1e-6069").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+110"))
	if err != nil {
		lang.SayString("dqdiv1769")
	} else {
		if !(rexsult.ToString() == "1E-6179") {
			lang.SayString("dqdiv1769")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6101"))
	if err != nil {
		lang.SayString("dqdiv1770")
	} else {
		if !(rexsult.ToString() == "1E+6141") {
			lang.SayString("dqdiv1770")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6102"))
	if err != nil {
		lang.SayString("dqdiv1771")
	} else {
		if !(rexsult.ToString() == "1E+6142") {
			lang.SayString("dqdiv1771")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6103"))
	if err != nil {
		lang.SayString("dqdiv1772")
	} else {
		if !(rexsult.ToString() == "1E+6143") {
			lang.SayString("dqdiv1772")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6104"))
	if err != nil {
		lang.SayString("dqdiv1773")
	} else {
		if !(rexsult.ToString() == "1E+6144") {
			lang.SayString("dqdiv1773")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6105"))
	if err != nil {
		lang.SayString("dqdiv1774")
	} else {
		if !(rexsult.ToString() == "1E+6145") {
			lang.SayString("dqdiv1774")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6106"))
	if err != nil {
		lang.SayString("dqdiv1775")
	} else {
		if !(rexsult.ToString() == "1E+6146") {
			lang.SayString("dqdiv1775")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6107"))
	if err != nil {
		lang.SayString("dqdiv1776")
	} else {
		if !(rexsult.ToString() == "1E+6147") {
			lang.SayString("dqdiv1776")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6108"))
	if err != nil {
		lang.SayString("dqdiv1777")
	} else {
		if !(rexsult.ToString() == "1E+6148") {
			lang.SayString("dqdiv1777")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6109"))
	if err != nil {
		lang.SayString("dqdiv1778")
	} else {
		if !(rexsult.ToString() == "1E+6149") {
			lang.SayString("dqdiv1778")
		}
	}
	rexsult, err = lang.RxFromString("1e+40").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e-6110"))
	if err != nil {
		lang.SayString("dqdiv1779")
	} else {
		if !(rexsult.ToString() == "1E+6150") {
			lang.SayString("dqdiv1779")
		}
	}
	rexsult, err = lang.RxFromString("1.0000E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("dqdiv1801")
	} else {
		if !(rexsult.ToString() == "1E-6172") {
			lang.SayString("dqdiv1801")
		}
	}
	rexsult, err = lang.RxFromString("1.000E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+1"))
	if err != nil {
		lang.SayString("dqdiv1802")
	} else {
		if !(rexsult.ToString() == "1E-6173") {
			lang.SayString("dqdiv1802")
		}
	}
	rexsult, err = lang.RxFromString("1.00E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+2"))
	if err != nil {
		lang.SayString("dqdiv1803")
	} else {
		if !(rexsult.ToString() == "1E-6174") {
			lang.SayString("dqdiv1803")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+3"))
	if err != nil {
		lang.SayString("dqdiv1804")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqdiv1804")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1805")
	} else {
		if !(rexsult.ToString() == "1E-6176") {
			lang.SayString("dqdiv1805")
		}
	}
	rexsult, err = lang.RxFromString("1.3E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1806")
	} else {
		if !(rexsult.ToString() == "1.3E-6176") {
			lang.SayString("dqdiv1806")
		}
	}
	rexsult, err = lang.RxFromString("1.5E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1807")
	} else {
		if !(rexsult.ToString() == "1.5E-6176") {
			lang.SayString("dqdiv1807")
		}
	}
	rexsult, err = lang.RxFromString("1.7E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1808")
	} else {
		if !(rexsult.ToString() == "1.7E-6176") {
			lang.SayString("dqdiv1808")
		}
	}
	rexsult, err = lang.RxFromString("2.3E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1809")
	} else {
		if !(rexsult.ToString() == "2.3E-6176") {
			lang.SayString("dqdiv1809")
		}
	}
	rexsult, err = lang.RxFromString("2.5E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1810")
	} else {
		if !(rexsult.ToString() == "2.5E-6176") {
			lang.SayString("dqdiv1810")
		}
	}
	rexsult, err = lang.RxFromString("2.7E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1811")
	} else {
		if !(rexsult.ToString() == "2.7E-6176") {
			lang.SayString("dqdiv1811")
		}
	}
	rexsult, err = lang.RxFromString("1.49E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1812")
	} else {
		if !(rexsult.ToString() == "1.49E-6176") {
			lang.SayString("dqdiv1812")
		}
	}
	rexsult, err = lang.RxFromString("1.50E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1813")
	} else {
		if !(rexsult.ToString() == "1.5E-6176") {
			lang.SayString("dqdiv1813")
		}
	}
	rexsult, err = lang.RxFromString("1.51E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1814")
	} else {
		if !(rexsult.ToString() == "1.51E-6176") {
			lang.SayString("dqdiv1814")
		}
	}
	rexsult, err = lang.RxFromString("2.49E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1815")
	} else {
		if !(rexsult.ToString() == "2.49E-6176") {
			lang.SayString("dqdiv1815")
		}
	}
	rexsult, err = lang.RxFromString("2.50E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1816")
	} else {
		if !(rexsult.ToString() == "2.5E-6176") {
			lang.SayString("dqdiv1816")
		}
	}
	rexsult, err = lang.RxFromString("2.51E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1817")
	} else {
		if !(rexsult.ToString() == "2.51E-6176") {
			lang.SayString("dqdiv1817")
		}
	}
	rexsult, err = lang.RxFromString("1E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+4"))
	if err != nil {
		lang.SayString("dqdiv1818")
	} else {
		if !(rexsult.ToString() == "1E-6176") {
			lang.SayString("dqdiv1818")
		}
	}
	rexsult, err = lang.RxFromString("3E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dqdiv1819")
	} else {
		if !(rexsult.ToString() == "3E-6177") {
			lang.SayString("dqdiv1819")
		}
	}
	rexsult, err = lang.RxFromString("5E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dqdiv1820")
	} else {
		if !(rexsult.ToString() == "5E-6177") {
			lang.SayString("dqdiv1820")
		}
	}
	rexsult, err = lang.RxFromString("7E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dqdiv1821")
	} else {
		if !(rexsult.ToString() == "7E-6177") {
			lang.SayString("dqdiv1821")
		}
	}
	rexsult, err = lang.RxFromString("9E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dqdiv1822")
	} else {
		if !(rexsult.ToString() == "9E-6177") {
			lang.SayString("dqdiv1822")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dqdiv1823")
	} else {
		if !(rexsult.ToString() == "9.9E-6177") {
			lang.SayString("dqdiv1823")
		}
	}
	rexsult, err = lang.RxFromString("1E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e+4"))
	if err != nil {
		lang.SayString("dqdiv1824")
	} else {
		if !(rexsult.ToString() == "-1E-6176") {
			lang.SayString("dqdiv1824")
		}
	}
	rexsult, err = lang.RxFromString("3E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dqdiv1825")
	} else {
		if !(rexsult.ToString() == "-3E-6177") {
			lang.SayString("dqdiv1825")
		}
	}
	rexsult, err = lang.RxFromString("-5E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dqdiv1826")
	} else {
		if !(rexsult.ToString() == "-5E-6177") {
			lang.SayString("dqdiv1826")
		}
	}
	rexsult, err = lang.RxFromString("7E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dqdiv1827")
	} else {
		if !(rexsult.ToString() == "-7E-6177") {
			lang.SayString("dqdiv1827")
		}
	}
	rexsult, err = lang.RxFromString("-9E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+5"))
	if err != nil {
		lang.SayString("dqdiv1828")
	} else {
		if !(rexsult.ToString() == "-9E-6177") {
			lang.SayString("dqdiv1828")
		}
	}
	rexsult, err = lang.RxFromString("9.9E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dqdiv1829")
	} else {
		if !(rexsult.ToString() == "-9.9E-6177") {
			lang.SayString("dqdiv1829")
		}
	}
	rexsult, err = lang.RxFromString("3.0E-6172").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-1e+5"))
	if err != nil {
		lang.SayString("dqdiv1830")
	} else {
		if !(rexsult.ToString() == "-3E-6177") {
			lang.SayString("dqdiv1830")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-5977").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+200"))
	if err != nil {
		lang.SayString("dqdiv1831")
	} else {
		if !(rexsult.ToString() == "1E-6177") {
			lang.SayString("dqdiv1831")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-5977").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+199"))
	if err != nil {
		lang.SayString("dqdiv1832")
	} else {
		if !(rexsult.ToString() == "1E-6176") {
			lang.SayString("dqdiv1832")
		}
	}
	rexsult, err = lang.RxFromString("1.0E-5977").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1e+198"))
	if err != nil {
		lang.SayString("dqdiv1833")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqdiv1833")
		}
	}
	rexsult, err = lang.RxFromString("2.0E-5977").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2e+198"))
	if err != nil {
		lang.SayString("dqdiv1834")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqdiv1834")
		}
	}
	rexsult, err = lang.RxFromString("4.0E-5977").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4e+198"))
	if err != nil {
		lang.SayString("dqdiv1835")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqdiv1835")
		}
	}
	rexsult, err = lang.RxFromString("10.0E-5977").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("10e+198"))
	if err != nil {
		lang.SayString("dqdiv1836")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqdiv1836")
		}
	}
	rexsult, err = lang.RxFromString("30.0E-5977").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("30e+198"))
	if err != nil {
		lang.SayString("dqdiv1837")
	} else {
		if !(rexsult.ToString() == "1E-6175") {
			lang.SayString("dqdiv1837")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-5982").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("40e+166"))
	if err != nil {
		lang.SayString("dqdiv1838")
	} else {
		if !(rexsult.ToString() == "1E-6148") {
			lang.SayString("dqdiv1838")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-5982").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("40e+165"))
	if err != nil {
		lang.SayString("dqdiv1839")
	} else {
		if !(rexsult.ToString() == "1E-6147") {
			lang.SayString("dqdiv1839")
		}
	}
	rexsult, err = lang.RxFromString("40.0E-5982").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("40e+164"))
	if err != nil {
		lang.SayString("dqdiv1840")
	} else {
		if !(rexsult.ToString() == "1E-6146") {
			lang.SayString("dqdiv1840")
		}
	}
	rexsult, err = lang.RxFromString("-5231195652931651968034356117118850").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-7243718664422548573203260970.34995"))
	if err != nil {
		lang.SayString("dqdiv2010")
	} else {
		if !(rexsult.ToString() == "722169.909583128462473605146055068") {
			lang.SayString("dqdiv2010")
		}
	}
	rexsult, err = lang.RxFromString("-89584669773927.82711237350022515352").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-42077943728529635884.21142627532985"))
	if err != nil {
		lang.SayString("dqdiv2011")
	} else {
		if !(rexsult.ToString() == "0.000002129017291146471565928125887527266") {
			lang.SayString("dqdiv2011")
		}
	}
	rexsult, err = lang.RxFromString("-2.828201693360723203806974891946180E-232").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("812596541221823960386384403089240.9"))
	if err != nil {
		lang.SayString("dqdiv2012")
	} else {
		if !(rexsult.ToString() == "-3.48045007564052132004005575912512E-265") {
			lang.SayString("dqdiv2012")
		}
	}
	rexsult, err = lang.RxFromString("-6442775372761069267502937539408720").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("24904085056.69185465145182606089196"))
	if err != nil {
		lang.SayString("dqdiv2013")
	} else {
		if !(rexsult.ToString() == "-258703556388226463687701.4884719589") {
			lang.SayString("dqdiv2013")
		}
	}
	rexsult, err = lang.RxFromString("5.535520011272625629610079879714705").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-44343664650.57203052003068113531208"))
	if err != nil {
		lang.SayString("dqdiv2014")
	} else {
		if !(rexsult.ToString() == "-1.248322630728089308975940533493562E-10") {
			lang.SayString("dqdiv2014")
		}
	}
	rexsult, err = lang.RxFromString("65919273712517865964325.99419625010").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-314733354141381737378622515.7789054"))
	if err != nil {
		lang.SayString("dqdiv2015")
	} else {
		if !(rexsult.ToString() == "-0.0002094448295521490616379784758911632") {
			lang.SayString("dqdiv2015")
		}
	}
	rexsult, err = lang.RxFromString("-7.779172568193197107115275140431129E+759").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-140453015639.3988987652895178782143"))
	if err != nil {
		lang.SayString("dqdiv2016")
	} else {
		if !(rexsult.ToString() == "5.538629792161641534962774244238115E+748") {
			lang.SayString("dqdiv2016")
		}
	}
	rexsult, err = lang.RxFromString("644314832597569.0181226067518178797").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-115024585257425.1635759521565201075"))
	if err != nil {
		lang.SayString("dqdiv2017")
	} else {
		if !(rexsult.ToString() == "-5.601540150356479257367687450922795") {
			lang.SayString("dqdiv2017")
		}
	}
	rexsult, err = lang.RxFromString("6.898640941579611450676592553286870E-47").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-11272429881407851485163914999.25943"))
	if err != nil {
		lang.SayString("dqdiv2018")
	} else {
		if !(rexsult.ToString() == "-6.11992357828533868937113764831928E-75") {
			lang.SayString("dqdiv2018")
		}
	}
	rexsult, err = lang.RxFromString("-3591344544888727133.30819750163254").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5329395.423792795661446561090331037"))
	if err != nil {
		lang.SayString("dqdiv2019")
	} else {
		if !(rexsult.ToString() == "-673874662941.196852558946053372529") {
			lang.SayString("dqdiv2019")
		}
	}
	rexsult, err = lang.RxFromString("-7.682356781384631313156462724425838E+747").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-6.60375855512219057281922141809940E+703"))
	if err != nil {
		lang.SayString("dqdiv2020")
	} else {
		if !(rexsult.ToString() == "1.163330960279556016678379128875149E+44") {
			lang.SayString("dqdiv2020")
		}
	}
	rexsult, err = lang.RxFromString("-4511495596596941820863224.274679699").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3365395017.263329795449661616090724"))
	if err != nil {
		lang.SayString("dqdiv2021")
	} else {
		if !(rexsult.ToString() == "-1340554548115304.904166888018346299") {
			lang.SayString("dqdiv2021")
		}
	}
	rexsult, err = lang.RxFromString("5.211164127840931517263639608151299").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("164.5566381356276567012533847006453"))
	if err != nil {
		lang.SayString("dqdiv2022")
	} else {
		if !(rexsult.ToString() == "0.0316679058765522886447826015715651") {
			lang.SayString("dqdiv2022")
		}
	}
	rexsult, err = lang.RxFromString("-49891.2243893458830384077684620383").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-47179.9312961860747554053371171530"))
	if err != nil {
		lang.SayString("dqdiv2023")
	} else {
		if !(rexsult.ToString() == "1.057467084386767291602189656430268") {
			lang.SayString("dqdiv2023")
		}
	}
	rexsult, err = lang.RxFromString("15065477.47214268488077415462413353").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4366211.120892953261309529740552596"))
	if err != nil {
		lang.SayString("dqdiv2024")
	} else {
		if !(rexsult.ToString() == "3.450469309661227984244545513441359") {
			lang.SayString("dqdiv2024")
		}
	}
	rexsult, err = lang.RxFromString("1.575670269440761846109602429612644E+370").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("653199649324740300.006185482643439"))
	if err != nil {
		lang.SayString("dqdiv2025")
	} else {
		if !(rexsult.ToString() == "2.412233795700359170904588548041481E+352") {
			lang.SayString("dqdiv2025")
		}
	}
	rexsult, err = lang.RxFromString("-2112422311733448924573432192.620145").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-80067206.03590693153848215848613406"))
	if err != nil {
		lang.SayString("dqdiv2026")
	} else {
		if !(rexsult.ToString() == "26383115089417660175.20102646756574") {
			lang.SayString("dqdiv2026")
		}
	}
	rexsult, err = lang.RxFromString("-67096536051279809.32218611548721839").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-869685412881941081664251990181.1049"))
	if err != nil {
		lang.SayString("dqdiv2027")
	} else {
		if !(rexsult.ToString() == "7.715035236584805921278566365231168E-14") {
			lang.SayString("dqdiv2027")
		}
	}
	rexsult, err = lang.RxFromString("-58612908548962047.21866913425488972").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-978449597531.3873665583475633831644"))
	if err != nil {
		lang.SayString("dqdiv2028")
	} else {
		if !(rexsult.ToString() == "59903.86085991703091236507859837023") {
			lang.SayString("dqdiv2028")
		}
	}
	rexsult, err = lang.RxFromString("-133032412010942.1476864138213319796").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-7.882059293498670705446528648201359E-428"))
	if err != nil {
		lang.SayString("dqdiv2029")
	} else {
		if !(rexsult.ToString() == "1.687787506504433064549515681693715E+441") {
			lang.SayString("dqdiv2029")
		}
	}
	rexsult, err = lang.RxFromString("1.83746698338966029492299716360513E+977").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-9.897926608979649951672839879128603E+154"))
	if err != nil {
		lang.SayString("dqdiv2030")
	} else {
		if !(rexsult.ToString() == "-1.856416051542212552042390218062458E+822") {
			lang.SayString("dqdiv2030")
		}
	}
	rexsult, err = lang.RxFromString("-113742475841399236307128962.1507063").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8298602.203049834732657567965262989"))
	if err != nil {
		lang.SayString("dqdiv2031")
	} else {
		if !(rexsult.ToString() == "-13706221006665137826.16557393919929") {
			lang.SayString("dqdiv2031")
		}
	}
	rexsult, err = lang.RxFromString("196.4787574650754152995941808331862").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("929.6553388472318094427422117172394"))
	if err != nil {
		lang.SayString("dqdiv2032")
	} else {
		if !(rexsult.ToString() == "0.2113458066176526651006917922814018") {
			lang.SayString("dqdiv2032")
		}
	}
	rexsult, err = lang.RxFromString("71931221465.43867996282803628130350").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3838685934206426257090718.402248853"))
	if err != nil {
		lang.SayString("dqdiv2033")
	} else {
		if !(rexsult.ToString() == "1.873850132527423413607199513324021E-14") {
			lang.SayString("dqdiv2033")
		}
	}
	rexsult, err = lang.RxFromString("488.4282502289651653783596246312885").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-80.68940956806634280078706577953188"))
	if err != nil {
		lang.SayString("dqdiv2034")
	} else {
		if !(rexsult.ToString() == "-6.053189047280693318844801899473272") {
			lang.SayString("dqdiv2034")
		}
	}
	rexsult, err = lang.RxFromString("9.001764344963921754981762913247394E-162").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-8.585540973667205753734967645386919E-729"))
	if err != nil {
		lang.SayString("dqdiv2035")
	} else {
		if !(rexsult.ToString() == "-1.048479574271827326396012573232934E+567") {
			lang.SayString("dqdiv2035")
		}
	}
	rexsult, err = lang.RxFromString("-7.404133959409894743706402857145471E-828").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-51.38159929460289711134684843086265"))
	if err != nil {
		lang.SayString("dqdiv2036")
	} else {
		if !(rexsult.ToString() == "1.441008855516029461032061785219773E-829") {
			lang.SayString("dqdiv2036")
		}
	}
	rexsult, err = lang.RxFromString("2.967520235574419794048994436040717E-613").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-6252513855.91394894949879262731889"))
	if err != nil {
		lang.SayString("dqdiv2037")
	} else {
		if !(rexsult.ToString() == "-4.746123405656409127572998751885338E-623") {
			lang.SayString("dqdiv2037")
		}
	}
	rexsult, err = lang.RxFromString("-18826852654824040505.83920366765051").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-6336924877942437992590557460147340"))
	if err != nil {
		lang.SayString("dqdiv2038")
	} else {
		if !(rexsult.ToString() == "2.970976146546494669807886278519194E-15") {
			lang.SayString("dqdiv2038")
		}
	}
	rexsult, err = lang.RxFromString("-8.101406784809197604949584001735949E+561").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4.823300306948942821076681658771635E+361"))
	if err != nil {
		lang.SayString("dqdiv2039")
	} else {
		if !(rexsult.ToString() == "-1.679639721610839204738445747238987E+200") {
			lang.SayString("dqdiv2039")
		}
	}
	rexsult, err = lang.RxFromString("-6.11981977773094052331062585191723E+295").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1.507610253755339328302779005586534E+238"))
	if err != nil {
		lang.SayString("dqdiv2040")
	} else {
		if !(rexsult.ToString() == "-4.059285058911577244044418416044763E+57") {
			lang.SayString("dqdiv2040")
		}
	}
	rexsult, err = lang.RxFromString("6.472638850046815880599220534274055E-596").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-4.475233712083047516933911786159972"))
	if err != nil {
		lang.SayString("dqdiv2041")
	} else {
		if !(rexsult.ToString() == "-1.446324207062261745520496475778879E-596") {
			lang.SayString("dqdiv2041")
		}
	}
	rexsult, err = lang.RxFromString("-84438593330.71277839631144509397112").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-586684596204401664208947.4054879633"))
	if err != nil {
		lang.SayString("dqdiv2042")
	} else {
		if !(rexsult.ToString() == "1.439250218550041228759983937772504E-13") {
			lang.SayString("dqdiv2042")
		}
	}
	rexsult, err = lang.RxFromString("9.354533233294022616695815656704369E-24").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("405.500390626135304252144163591746"))
	if err != nil {
		lang.SayString("dqdiv2043")
	} else {
		if !(rexsult.ToString() == "2.306911028827774549740571229736198E-26") {
			lang.SayString("dqdiv2043")
		}
	}
	rexsult, err = lang.RxFromString("985606423350210.7374876650149957881").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-36811563697.41925681866694859828794"))
	if err != nil {
		lang.SayString("dqdiv2044")
	} else {
		if !(rexsult.ToString() == "-26774.36990864119445335813354717711") {
			lang.SayString("dqdiv2044")
		}
	}
	rexsult, err = lang.RxFromString("-8.187280774177715706278002247766311E-123").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-38784124393.91212870828430001300068"))
	if err != nil {
		lang.SayString("dqdiv2045")
	} else {
		if !(rexsult.ToString() == "2.110987653356139147357240727794365E-133") {
			lang.SayString("dqdiv2045")
		}
	}
	rexsult, err = lang.RxFromString("-4.612203126350070903459245798371657E+912").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7.971562182727956290901984736800519E+64"))
	if err != nil {
		lang.SayString("dqdiv2046")
	} else {
		if !(rexsult.ToString() == "-5.785820922708683237098826662769748E+847") {
			lang.SayString("dqdiv2046")
		}
	}
	rexsult, err = lang.RxFromString("4.661015909421485298247928967977089E+888").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-6.360911253323922338737311563845581E+388"))
	if err != nil {
		lang.SayString("dqdiv2047")
	} else {
		if !(rexsult.ToString() == "-7.327591478321365980156654539638836E+499") {
			lang.SayString("dqdiv2047")
		}
	}
	rexsult, err = lang.RxFromString("9156078172903.257500003260710833030").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7.189796653262147139071634237964074E-90"))
	if err != nil {
		lang.SayString("dqdiv2048")
	} else {
		if !(rexsult.ToString() == "1.273482215766000994365201545096026E+102") {
			lang.SayString("dqdiv2048")
		}
	}
	rexsult, err = lang.RxFromString("-1.710722303327476586373477781276586E-311").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-3167561628260156837329323.729380695"))
	if err != nil {
		lang.SayString("dqdiv2049")
	} else {
		if !(rexsult.ToString() == "5.400754599578613984875752958645655E-336") {
			lang.SayString("dqdiv2049")
		}
	}
	rexsult, err = lang.RxFromString("-4.647935210881806238321616345413021E-878").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("209388.5431867744648177308460639582"))
	if err != nil {
		lang.SayString("dqdiv2050")
	} else {
		if !(rexsult.ToString() == "-2.21976577139459373314049429738814E-883") {
			lang.SayString("dqdiv2050")
		}
	}
	rexsult, err = lang.RxFromString("5958.694728395760992719084781582700").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4.541510156564315632536353171846096E-746"))
	if err != nil {
		lang.SayString("dqdiv2051")
	} else {
		if !(rexsult.ToString() == "1.31205139325363866494785269300548E+749") {
			lang.SayString("dqdiv2051")
		}
	}
	rexsult, err = lang.RxFromString("-7.935732544649702175256699886872093E-489").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-7.433329073664793138998765647467971E+360"))
	if err != nil {
		lang.SayString("dqdiv2052")
	} else {
		if !(rexsult.ToString() == "1.067587949626076917672271619664656E-849") {
			lang.SayString("dqdiv2052")
		}
	}
	rexsult, err = lang.RxFromString("-2746650864601157.863589959939901350").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7.016684945507647528907184694359598E+548"))
	if err != nil {
		lang.SayString("dqdiv2053")
	} else {
		if !(rexsult.ToString() == "-3.914456593009309529351254950429932E-534") {
			lang.SayString("dqdiv2053")
		}
	}
	rexsult, err = lang.RxFromString("3605149408631197365447953.994569178").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-75614025825649082.78264864428237833"))
	if err != nil {
		lang.SayString("dqdiv2054")
	} else {
		if !(rexsult.ToString() == "-47678315.88472693507060063188020532") {
			lang.SayString("dqdiv2054")
		}
	}
	rexsult, err = lang.RxFromString("788194320921798404906375214.196349").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-6.222718148433247384932573401976337E-418"))
	if err != nil {
		lang.SayString("dqdiv2055")
	} else {
		if !(rexsult.ToString() == "-1.266639918634671803982222244977287E+444") {
			lang.SayString("dqdiv2055")
		}
	}
	rexsult, err = lang.RxFromString("5620722730534752.758208943447603211").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6.843552841168538319123000917657759E-139"))
	if err != nil {
		lang.SayString("dqdiv2056")
	} else {
		if !(rexsult.ToString() == "8.213164800485434666629970443739554E+153") {
			lang.SayString("dqdiv2056")
		}
	}
	rexsult, err = lang.RxFromString("7304534676713703938102.403949019402").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-576169.3685010935108153023803590835"))
	if err != nil {
		lang.SayString("dqdiv2057")
	} else {
		if !(rexsult.ToString() == "-12677756014201995.31969237144394772") {
			lang.SayString("dqdiv2057")
		}
	}
	rexsult, err = lang.RxFromString("8067918762.134621639254916786945547").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-8.774771480055536009105596163864758E+954"))
	if err != nil {
		lang.SayString("dqdiv2058")
	} else {
		if !(rexsult.ToString() == "-9.194448858836332156766764605125245E-946") {
			lang.SayString("dqdiv2058")
		}
	}
	rexsult, err = lang.RxFromString("8.702093454123046507578256899537563E-324").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("-5.875399733016018404580201176576293E-401"))
	if err != nil {
		lang.SayString("dqdiv2059")
	} else {
		if !(rexsult.ToString() == "-1.481106622452052581470443526957335E+77") {
			lang.SayString("dqdiv2059")
		}
	}
	rexsult, err = lang.RxFromString("-41426.01662518451861386352415092356").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("90.00146621684478300510769802013464"))
	if err != nil {
		lang.SayString("dqdiv2060")
	} else {
		if !(rexsult.ToString() == "-460.28157502873186927320677091762") {
			lang.SayString("dqdiv2060")
		}
	}
	rexsult, err = lang.RxFromString("2003100352770753969878925664524900").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2003100352770753969878925664497824"))
	if err != nil {
		lang.SayString("dqdiv4001")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000013517") {
			lang.SayString("dqdiv4001")
		}
	}
	rexsult, err = lang.RxFromString("4817785793916490652579552318371645").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4817785793916490652579552318362097"))
	if err != nil {
		lang.SayString("dqdiv4002")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000001982") {
			lang.SayString("dqdiv4002")
		}
	}
	rexsult, err = lang.RxFromString("8299187410920067325648068439560282").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8299187410920067325648068439591159"))
	if err != nil {
		lang.SayString("dqdiv4003")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999962795") {
			lang.SayString("dqdiv4003")
		}
	}
	rexsult, err = lang.RxFromString("5641088455897407044544461785365899").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5641088455897407044544461785389965"))
	if err != nil {
		lang.SayString("dqdiv4004")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999957338") {
			lang.SayString("dqdiv4004")
		}
	}
	rexsult, err = lang.RxFromString("5752274694706545359326361313490424").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5752274694706545359326361313502723"))
	if err != nil {
		lang.SayString("dqdiv4005")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999978619") {
			lang.SayString("dqdiv4005")
		}
	}
	rexsult, err = lang.RxFromString("6762079477373670594829319346099665").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6762079477373670594829319346132579"))
	if err != nil {
		lang.SayString("dqdiv4006")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999951326") {
			lang.SayString("dqdiv4006")
		}
	}
	rexsult, err = lang.RxFromString("7286425153691890341633023222602916").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7286425153691890341633023222606556"))
	if err != nil {
		lang.SayString("dqdiv4007")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999995004") {
			lang.SayString("dqdiv4007")
		}
	}
	rexsult, err = lang.RxFromString("9481233991901305727648306421946655").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9481233991901305727648306421919124"))
	if err != nil {
		lang.SayString("dqdiv4008")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002904") {
			lang.SayString("dqdiv4008")
		}
	}
	rexsult, err = lang.RxFromString("4282053941893951742029444065614311").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4282053941893951742029444065583077"))
	if err != nil {
		lang.SayString("dqdiv4009")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000007294") {
			lang.SayString("dqdiv4009")
		}
	}
	rexsult, err = lang.RxFromString("626888225441250639741781850338695").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("626888225441250639741781850327299"))
	if err != nil {
		lang.SayString("dqdiv4010")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000018179") {
			lang.SayString("dqdiv4010")
		}
	}
	rexsult, err = lang.RxFromString("3860973649222028009456598604468547").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3860973649222028009456598604476849"))
	if err != nil {
		lang.SayString("dqdiv4011")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999978498") {
			lang.SayString("dqdiv4011")
		}
	}
	rexsult, err = lang.RxFromString("4753157080127468127908060607821839").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4753157080127468127908060607788379"))
	if err != nil {
		lang.SayString("dqdiv4012")
	} else {
		if !(rexsult.ToString() == "1.00000000000000000000000000000704") {
			lang.SayString("dqdiv4012")
		}
	}
	rexsult, err = lang.RxFromString("552448546203754062805706277880419").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("552448546203754062805706277881903"))
	if err != nil {
		lang.SayString("dqdiv4013")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999973138") {
			lang.SayString("dqdiv4013")
		}
	}
	rexsult, err = lang.RxFromString("8405954527952158455323713728917395").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8405954527952158455323713728933866"))
	if err != nil {
		lang.SayString("dqdiv4014")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999980406") {
			lang.SayString("dqdiv4014")
		}
	}
	rexsult, err = lang.RxFromString("7554096502235321142555802238016116").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7554096502235321142555802238026546"))
	if err != nil {
		lang.SayString("dqdiv4015")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999986193") {
			lang.SayString("dqdiv4015")
		}
	}
	rexsult, err = lang.RxFromString("4053257674127518606871054934746782").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4053257674127518606871054934767355"))
	if err != nil {
		lang.SayString("dqdiv4016")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999949243") {
			lang.SayString("dqdiv4016")
		}
	}
	rexsult, err = lang.RxFromString("7112419420755090454716888844011582").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7112419420755090454716888844038105"))
	if err != nil {
		lang.SayString("dqdiv4017")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999962709") {
			lang.SayString("dqdiv4017")
		}
	}
	rexsult, err = lang.RxFromString("3132302137520072728164549730911846").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3132302137520072728164549730908416"))
	if err != nil {
		lang.SayString("dqdiv4018")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000001095") {
			lang.SayString("dqdiv4018")
		}
	}
	rexsult, err = lang.RxFromString("4788374045841416355706715048161013").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4788374045841416355706715048190077"))
	if err != nil {
		lang.SayString("dqdiv4019")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999939303") {
			lang.SayString("dqdiv4019")
		}
	}
	rexsult, err = lang.RxFromString("9466021636047630218238075099510597").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9466021636047630218238075099484053"))
	if err != nil {
		lang.SayString("dqdiv4020")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002804") {
			lang.SayString("dqdiv4020")
		}
	}
	rexsult, err = lang.RxFromString("912742745646765625597399692138650").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("912742745646765625597399692139042"))
	if err != nil {
		lang.SayString("dqdiv4021")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999995705") {
			lang.SayString("dqdiv4021")
		}
	}
	rexsult, err = lang.RxFromString("9508402742933643208806264897188504").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9508402742933643208806264897195973"))
	if err != nil {
		lang.SayString("dqdiv4022")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999992145") {
			lang.SayString("dqdiv4022")
		}
	}
	rexsult, err = lang.RxFromString("1186956795727233704962361914360895").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1186956795727233704962361914329577"))
	if err != nil {
		lang.SayString("dqdiv4023")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000026385") {
			lang.SayString("dqdiv4023")
		}
	}
	rexsult, err = lang.RxFromString("5972210268839014812696916170967938").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5972210268839014812696916170954974"))
	if err != nil {
		lang.SayString("dqdiv4024")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002171") {
			lang.SayString("dqdiv4024")
		}
	}
	rexsult, err = lang.RxFromString("2303801625521619930894460139793140").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2303801625521619930894460139799643"))
	if err != nil {
		lang.SayString("dqdiv4025")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999971773") {
			lang.SayString("dqdiv4025")
		}
	}
	rexsult, err = lang.RxFromString("6022231560002898264777393473966595").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6022231560002898264777393473947198"))
	if err != nil {
		lang.SayString("dqdiv4026")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000003221") {
			lang.SayString("dqdiv4026")
		}
	}
	rexsult, err = lang.RxFromString("8426148335801396199969346032210893").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8426148335801396199969346032203179"))
	if err != nil {
		lang.SayString("dqdiv4027")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000915") {
			lang.SayString("dqdiv4027")
		}
	}
	rexsult, err = lang.RxFromString("8812278947028784637382847098411749").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8812278947028784637382847098385317"))
	if err != nil {
		lang.SayString("dqdiv4028")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002999") {
			lang.SayString("dqdiv4028")
		}
	}
	rexsult, err = lang.RxFromString("8145282002348367383264197170116146").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8145282002348367383264197170083988"))
	if err != nil {
		lang.SayString("dqdiv4029")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000003948") {
			lang.SayString("dqdiv4029")
		}
	}
	rexsult, err = lang.RxFromString("6821577571876840153123510107387026").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6821577571876840153123510107418008"))
	if err != nil {
		lang.SayString("dqdiv4030")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999954582") {
			lang.SayString("dqdiv4030")
		}
	}
	rexsult, err = lang.RxFromString("9018555319518966970480565482023720").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9018555319518966970480565482013346"))
	if err != nil {
		lang.SayString("dqdiv4031")
	} else {
		if !(rexsult.ToString() == "1.00000000000000000000000000000115") {
			lang.SayString("dqdiv4031")
		}
	}
	rexsult, err = lang.RxFromString("4602155712998228449640717252788864").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4602155712998228449640717252818502"))
	if err != nil {
		lang.SayString("dqdiv4032")
	} else {
		if !(rexsult.ToString() == "0.99999999999999999999999999999356") {
			lang.SayString("dqdiv4032")
		}
	}
	rexsult, err = lang.RxFromString("6675607481522785614506828292264472").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6675607481522785614506828292277100"))
	if err != nil {
		lang.SayString("dqdiv4033")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999981083") {
			lang.SayString("dqdiv4033")
		}
	}
	rexsult, err = lang.RxFromString("4015881516871833897766945836264472").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4015881516871833897766945836262645"))
	if err != nil {
		lang.SayString("dqdiv4034")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000455") {
			lang.SayString("dqdiv4034")
		}
	}
	rexsult, err = lang.RxFromString("1415580205933411837595459716910365").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1415580205933411837595459716880139"))
	if err != nil {
		lang.SayString("dqdiv4035")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000021352") {
			lang.SayString("dqdiv4035")
		}
	}
	rexsult, err = lang.RxFromString("9432968297069542816752035276361552").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9432968297069542816752035276353054"))
	if err != nil {
		lang.SayString("dqdiv4036")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000901") {
			lang.SayString("dqdiv4036")
		}
	}
	rexsult, err = lang.RxFromString("4799319591303848500532766682140658").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4799319591303848500532766682172655"))
	if err != nil {
		lang.SayString("dqdiv4037")
	} else {
		if !(rexsult.ToString() == "0.999999999999999999999999999993333") {
			lang.SayString("dqdiv4037")
		}
	}
	rexsult, err = lang.RxFromString("316854270732839529790584284987472").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("316854270732839529790584285004832"))
	if err != nil {
		lang.SayString("dqdiv4038")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999452114") {
			lang.SayString("dqdiv4038")
		}
	}
	rexsult, err = lang.RxFromString("3598981300592490427826027975697415").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3598981300592490427826027975686712"))
	if err != nil {
		lang.SayString("dqdiv4039")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002974") {
			lang.SayString("dqdiv4039")
		}
	}
	rexsult, err = lang.RxFromString("1664315435694461371155800682196520").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1664315435694461371155800682195617"))
	if err != nil {
		lang.SayString("dqdiv4040")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000543") {
			lang.SayString("dqdiv4040")
		}
	}
	rexsult, err = lang.RxFromString("1680872316531128890102855316510581").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1680872316531128890102855316495545"))
	if err != nil {
		lang.SayString("dqdiv4041")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000008945") {
			lang.SayString("dqdiv4041")
		}
	}
	rexsult, err = lang.RxFromString("9881274879566405475755499281644730").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9881274879566405475755499281615743"))
	if err != nil {
		lang.SayString("dqdiv4042")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002934") {
			lang.SayString("dqdiv4042")
		}
	}
	rexsult, err = lang.RxFromString("4737225957717466960447204232279216").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4737225957717466960447204232277452"))
	if err != nil {
		lang.SayString("dqdiv4043")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000372") {
			lang.SayString("dqdiv4043")
		}
	}
	rexsult, err = lang.RxFromString("2482097379414867061213319346418288").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2482097379414867061213319346387936"))
	if err != nil {
		lang.SayString("dqdiv4044")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000012228") {
			lang.SayString("dqdiv4044")
		}
	}
	rexsult, err = lang.RxFromString("7406977595233762723576434122161868").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7406977595233762723576434122189042"))
	if err != nil {
		lang.SayString("dqdiv4045")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999963313") {
			lang.SayString("dqdiv4045")
		}
	}
	rexsult, err = lang.RxFromString("228782057757566047086593281773577").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("228782057757566047086593281769727"))
	if err != nil {
		lang.SayString("dqdiv4046")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000016828") {
			lang.SayString("dqdiv4046")
		}
	}
	rexsult, err = lang.RxFromString("2956594270240579648823270540367653").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2956594270240579648823270540368556"))
	if err != nil {
		lang.SayString("dqdiv4047")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999996946") {
			lang.SayString("dqdiv4047")
		}
	}
	rexsult, err = lang.RxFromString("6326964098897620620534136767634340").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6326964098897620620534136767619339"))
	if err != nil {
		lang.SayString("dqdiv4048")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002371") {
			lang.SayString("dqdiv4048")
		}
	}
	rexsult, err = lang.RxFromString("414586440456590215247002678327800").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("414586440456590215247002678316922"))
	if err != nil {
		lang.SayString("dqdiv4049")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000026238") {
			lang.SayString("dqdiv4049")
		}
	}
	rexsult, err = lang.RxFromString("7364552208570039386220505636779125").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7364552208570039386220505636803548"))
	if err != nil {
		lang.SayString("dqdiv4050")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999966837") {
			lang.SayString("dqdiv4050")
		}
	}
	rexsult, err = lang.RxFromString("5626266749902369710022824950590056").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5626266749902369710022824950591008"))
	if err != nil {
		lang.SayString("dqdiv4051")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999998308") {
			lang.SayString("dqdiv4051")
		}
	}
	rexsult, err = lang.RxFromString("4863278293916197454987481343460484").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4863278293916197454987481343442522"))
	if err != nil {
		lang.SayString("dqdiv4052")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000003693") {
			lang.SayString("dqdiv4052")
		}
	}
	rexsult, err = lang.RxFromString("1170713582030637359713249796835483").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1170713582030637359713249796823345"))
	if err != nil {
		lang.SayString("dqdiv4053")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000010368") {
			lang.SayString("dqdiv4053")
		}
	}
	rexsult, err = lang.RxFromString("9838062494725965667776326556052931").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9838062494725965667776326556061002"))
	if err != nil {
		lang.SayString("dqdiv4054")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999991796") {
			lang.SayString("dqdiv4054")
		}
	}
	rexsult, err = lang.RxFromString("4071388731298861093005687091498922").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4071388731298861093005687091498278"))
	if err != nil {
		lang.SayString("dqdiv4055")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000158") {
			lang.SayString("dqdiv4055")
		}
	}
	rexsult, err = lang.RxFromString("8753155722324706795855038590272526").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8753155722324706795855038590276656"))
	if err != nil {
		lang.SayString("dqdiv4056")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999995282") {
			lang.SayString("dqdiv4056")
		}
	}
	rexsult, err = lang.RxFromString("4399941911533273418844742658240485").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4399941911533273418844742658219891"))
	if err != nil {
		lang.SayString("dqdiv4057")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000004681") {
			lang.SayString("dqdiv4057")
		}
	}
	rexsult, err = lang.RxFromString("4127884159949503677776430620050269").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4127884159949503677776430620026091"))
	if err != nil {
		lang.SayString("dqdiv4058")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000005857") {
			lang.SayString("dqdiv4058")
		}
	}
	rexsult, err = lang.RxFromString("5536160822360800067042528317438808").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5536160822360800067042528317450687"))
	if err != nil {
		lang.SayString("dqdiv4059")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999978543") {
			lang.SayString("dqdiv4059")
		}
	}
	rexsult, err = lang.RxFromString("3973234998468664936671088237710246").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3973234998468664936671088237741886"))
	if err != nil {
		lang.SayString("dqdiv4060")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999920367") {
			lang.SayString("dqdiv4060")
		}
	}
	rexsult, err = lang.RxFromString("9824855935638263593410444142327358").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9824855935638263593410444142328576"))
	if err != nil {
		lang.SayString("dqdiv4061")
	} else {
		if !(rexsult.ToString() == "0.999999999999999999999999999999876") {
			lang.SayString("dqdiv4061")
		}
	}
	rexsult, err = lang.RxFromString("5917078517340218131867327300814867").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5917078517340218131867327300788701"))
	if err != nil {
		lang.SayString("dqdiv4062")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000004422") {
			lang.SayString("dqdiv4062")
		}
	}
	rexsult, err = lang.RxFromString("4354236601830544882286139612521362").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4354236601830544882286139612543223"))
	if err != nil {
		lang.SayString("dqdiv4063")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999949794") {
			lang.SayString("dqdiv4063")
		}
	}
	rexsult, err = lang.RxFromString("8058474772375259017342110013891294").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8058474772375259017342110013906792"))
	if err != nil {
		lang.SayString("dqdiv4064")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999980768") {
			lang.SayString("dqdiv4064")
		}
	}
	rexsult, err = lang.RxFromString("5519604020981748170517093746166328").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5519604020981748170517093746181763"))
	if err != nil {
		lang.SayString("dqdiv4065")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999972036") {
			lang.SayString("dqdiv4065")
		}
	}
	rexsult, err = lang.RxFromString("1502130966879805458831323782443139").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1502130966879805458831323782412213"))
	if err != nil {
		lang.SayString("dqdiv4066")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000020588") {
			lang.SayString("dqdiv4066")
		}
	}
	rexsult, err = lang.RxFromString("562795633719481212915159787980270").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("562795633719481212915159788007066"))
	if err != nil {
		lang.SayString("dqdiv4067")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999523877") {
			lang.SayString("dqdiv4067")
		}
	}
	rexsult, err = lang.RxFromString("6584743324494664273941281557268878").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6584743324494664273941281557258945"))
	if err != nil {
		lang.SayString("dqdiv4068")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000001508") {
			lang.SayString("dqdiv4068")
		}
	}
	rexsult, err = lang.RxFromString("3632000327285743997976431109416500").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3632000327285743997976431109408107"))
	if err != nil {
		lang.SayString("dqdiv4069")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002311") {
			lang.SayString("dqdiv4069")
		}
	}
	rexsult, err = lang.RxFromString("1145827237315430089388953838561450").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("1145827237315430089388953838527332"))
	if err != nil {
		lang.SayString("dqdiv4070")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000029776") {
			lang.SayString("dqdiv4070")
		}
	}
	rexsult, err = lang.RxFromString("8874431010357691869725372317350380").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8874431010357691869725372317316472"))
	if err != nil {
		lang.SayString("dqdiv4071")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000003821") {
			lang.SayString("dqdiv4071")
		}
	}
	rexsult, err = lang.RxFromString("992948718902804648119753141202196").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("992948718902804648119753141235222"))
	if err != nil {
		lang.SayString("dqdiv4072")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999667395") {
			lang.SayString("dqdiv4072")
		}
	}
	rexsult, err = lang.RxFromString("2522735183374218505142417265439989").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2522735183374218505142417265453779"))
	if err != nil {
		lang.SayString("dqdiv4073")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999945337") {
			lang.SayString("dqdiv4073")
		}
	}
	rexsult, err = lang.RxFromString("2668419161912936508006872303501052").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2668419161912936508006872303471036"))
	if err != nil {
		lang.SayString("dqdiv4074")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000011249") {
			lang.SayString("dqdiv4074")
		}
	}
	rexsult, err = lang.RxFromString("3036169085665186712590941111775092").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3036169085665186712590941111808846"))
	if err != nil {
		lang.SayString("dqdiv4075")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999888827") {
			lang.SayString("dqdiv4075")
		}
	}
	rexsult, err = lang.RxFromString("9441634604917231638508898934006147").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("9441634604917231638508898934000288"))
	if err != nil {
		lang.SayString("dqdiv4076")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000000621") {
			lang.SayString("dqdiv4076")
		}
	}
	rexsult, err = lang.RxFromString("2677301353164377091111458811839190").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2677301353164377091111458811867722"))
	if err != nil {
		lang.SayString("dqdiv4077")
	} else {
		if !(rexsult.ToString() == "0.999999999999999999999999999989343") {
			lang.SayString("dqdiv4077")
		}
	}
	rexsult, err = lang.RxFromString("6844979203112066166583765857171426").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6844979203112066166583765857189682"))
	if err != nil {
		lang.SayString("dqdiv4078")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999973329") {
			lang.SayString("dqdiv4078")
		}
	}
	rexsult, err = lang.RxFromString("2220337435141796724323783960231661").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2220337435141796724323783960208778"))
	if err != nil {
		lang.SayString("dqdiv4079")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000010306") {
			lang.SayString("dqdiv4079")
		}
	}
	rexsult, err = lang.RxFromString("6447424700019783931569996989561380").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6447424700019783931569996989572454"))
	if err != nil {
		lang.SayString("dqdiv4080")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999982824") {
			lang.SayString("dqdiv4080")
		}
	}
	rexsult, err = lang.RxFromString("7512856762696607119847092195587180").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7512856762696607119847092195557346"))
	if err != nil {
		lang.SayString("dqdiv4081")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000003971") {
			lang.SayString("dqdiv4081")
		}
	}
	rexsult, err = lang.RxFromString("7395261981193960399087819077237482").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7395261981193960399087819077242487"))
	if err != nil {
		lang.SayString("dqdiv4082")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999993232") {
			lang.SayString("dqdiv4082")
		}
	}
	rexsult, err = lang.RxFromString("2253442467682584035792724884376735").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2253442467682584035792724884407178"))
	if err != nil {
		lang.SayString("dqdiv4083")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999864904") {
			lang.SayString("dqdiv4083")
		}
	}
	rexsult, err = lang.RxFromString("8153138680300213135577336466190997").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("8153138680300213135577336466220607"))
	if err != nil {
		lang.SayString("dqdiv4084")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999963683") {
			lang.SayString("dqdiv4084")
		}
	}
	rexsult, err = lang.RxFromString("4668731252254148074041022681801390").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4668731252254148074041022681778101"))
	if err != nil {
		lang.SayString("dqdiv4085")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000004988") {
			lang.SayString("dqdiv4085")
		}
	}
	rexsult, err = lang.RxFromString("6078404557993669696040425501815056").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6078404557993669696040425501797612"))
	if err != nil {
		lang.SayString("dqdiv4086")
	} else {
		if !(rexsult.ToString() == "1.00000000000000000000000000000287") {
			lang.SayString("dqdiv4086")
		}
	}
	rexsult, err = lang.RxFromString("2306352359874261623223356878316278").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2306352359874261623223356878335612"))
	if err != nil {
		lang.SayString("dqdiv4087")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999916171") {
			lang.SayString("dqdiv4087")
		}
	}
	rexsult, err = lang.RxFromString("3264842186668480362900909564091908").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3264842186668480362900909564058658"))
	if err != nil {
		lang.SayString("dqdiv4088")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000010184") {
			lang.SayString("dqdiv4088")
		}
	}
	rexsult, err = lang.RxFromString("6971985047279636878957959608612204").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("6971985047279636878957959608615088"))
	if err != nil {
		lang.SayString("dqdiv4089")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999995863") {
			lang.SayString("dqdiv4089")
		}
	}
	rexsult, err = lang.RxFromString("5262810889952721235466445973816257").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5262810889952721235466445973783077"))
	if err != nil {
		lang.SayString("dqdiv4090")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000006305") {
			lang.SayString("dqdiv4090")
		}
	}
	rexsult, err = lang.RxFromString("7947944731035267178548357070080288").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("7947944731035267178548357070061339"))
	if err != nil {
		lang.SayString("dqdiv4091")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000002384") {
			lang.SayString("dqdiv4091")
		}
	}
	rexsult, err = lang.RxFromString("5071808908395375108383035800443229").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5071808908395375108383035800412429"))
	if err != nil {
		lang.SayString("dqdiv4092")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000006073") {
			lang.SayString("dqdiv4092")
		}
	}
	rexsult, err = lang.RxFromString("2043146542084503655511507209262969").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2043146542084503655511507209249263"))
	if err != nil {
		lang.SayString("dqdiv4093")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000006708") {
			lang.SayString("dqdiv4093")
		}
	}
	rexsult, err = lang.RxFromString("4097632735384534181661959731264802").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4097632735384534181661959731234499"))
	if err != nil {
		lang.SayString("dqdiv4094")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000007395") {
			lang.SayString("dqdiv4094")
		}
	}
	rexsult, err = lang.RxFromString("3061477642831387489729464587044430").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3061477642831387489729464587059452"))
	if err != nil {
		lang.SayString("dqdiv4095")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999950932") {
			lang.SayString("dqdiv4095")
		}
	}
	rexsult, err = lang.RxFromString("3429854941039776159498802936252638").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("3429854941039776159498802936246415"))
	if err != nil {
		lang.SayString("dqdiv4096")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000001814") {
			lang.SayString("dqdiv4096")
		}
	}
	rexsult, err = lang.RxFromString("4874324979578599700024133278284545").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("4874324979578599700024133278262131"))
	if err != nil {
		lang.SayString("dqdiv4097")
	} else {
		if !(rexsult.ToString() == "1.000000000000000000000000000004598") {
			lang.SayString("dqdiv4097")
		}
	}
	rexsult, err = lang.RxFromString("5701652369691833541455978515820882").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("5701652369691833541455978515834854"))
	if err != nil {
		lang.SayString("dqdiv4098")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999975495") {
			lang.SayString("dqdiv4098")
		}
	}
	rexsult, err = lang.RxFromString("2928205728402945266953255632343113").OpDiv(lang.RxSetWithDigit(34), lang.RxFromString("2928205728402945266953255632373794"))
	if err != nil {
		lang.SayString("dqdiv4099")
	} else {
		if !(rexsult.ToString() == "0.9999999999999999999999999999895223") {
			lang.SayString("dqdiv4099")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("1000"))
	if err != nil {
		lang.SayString("inx201")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("inx201")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("inx202")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("inx202")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("inx203")
	} else {
		if !(rexsult.ToString() == "500") {
			lang.SayString("inx203")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("inx204")
	} else {
		if !(rexsult.ToString() == "333.3") {
			lang.SayString("inx204")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("inx205")
	} else {
		if !(rexsult.ToString() == "250") {
			lang.SayString("inx205")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("inx206")
	} else {
		if !(rexsult.ToString() == "200") {
			lang.SayString("inx206")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("inx207")
	} else {
		if !(rexsult.ToString() == "166.7") {
			lang.SayString("inx207")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("inx208")
	} else {
		if !(rexsult.ToString() == "142.9") {
			lang.SayString("inx208")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("inx209")
	} else {
		if !(rexsult.ToString() == "125") {
			lang.SayString("inx209")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("inx210")
	} else {
		if !(rexsult.ToString() == "111.1") {
			lang.SayString("inx210")
		}
	}
	rexsult, err = lang.RxFromString("1000").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("inx211")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("inx211")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("inx220")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("inx220")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("inx221")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("inx221")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("inx222")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("inx222")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("inx223")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("inx223")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("inx224")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("inx224")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("inx225")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("inx225")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("inx226")
	} else {
		if !(rexsult.ToString() == "0.01563") {
			lang.SayString("inx226")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(4), lang.RxFromString("128"))
	if err != nil {
		lang.SayString("inx227")
	} else {
		if !(rexsult.ToString() == "0.007813") {
			lang.SayString("inx227")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("inx230")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("inx230")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("inx231")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("inx231")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("inx232")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("inx232")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("inx233")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("inx233")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("inx234")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("inx234")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("inx235")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("inx235")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("inx236")
	} else {
		if !(rexsult.ToString() == "0.015625") {
			lang.SayString("inx236")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("128"))
	if err != nil {
		lang.SayString("inx237")
	} else {
		if !(rexsult.ToString() == "0.0078125") {
			lang.SayString("inx237")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("inx240")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("inx240")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("inx241")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("inx241")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("inx242")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("inx242")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("inx243")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("inx243")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("inx244")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("inx244")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("inx245")
	} else {
		if !(rexsult.ToString() == "0.0313") {
			lang.SayString("inx245")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("inx246")
	} else {
		if !(rexsult.ToString() == "0.0156") {
			lang.SayString("inx246")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(3), lang.RxFromString("128"))
	if err != nil {
		lang.SayString("inx247")
	} else {
		if !(rexsult.ToString() == "0.00781") {
			lang.SayString("inx247")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("inx250")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("inx250")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("inx251")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("inx251")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("inx252")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("inx252")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("inx253")
	} else {
		if !(rexsult.ToString() == "0.13") {
			lang.SayString("inx253")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("inx254")
	} else {
		if !(rexsult.ToString() == "0.063") {
			lang.SayString("inx254")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("inx255")
	} else {
		if !(rexsult.ToString() == "0.031") {
			lang.SayString("inx255")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("inx256")
	} else {
		if !(rexsult.ToString() == "0.016") {
			lang.SayString("inx256")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(2), lang.RxFromString("128"))
	if err != nil {
		lang.SayString("inx257")
	} else {
		if !(rexsult.ToString() == "0.0078") {
			lang.SayString("inx257")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("inx260")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("inx260")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("inx261")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("inx261")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("inx262")
	} else {
		if !(rexsult.ToString() == "0.3") {
			lang.SayString("inx262")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("inx263")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("inx263")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("inx264")
	} else {
		if !(rexsult.ToString() == "0.06") {
			lang.SayString("inx264")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("inx265")
	} else {
		if !(rexsult.ToString() == "0.03") {
			lang.SayString("inx265")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("64"))
	if err != nil {
		lang.SayString("inx266")
	} else {
		if !(rexsult.ToString() == "0.02") {
			lang.SayString("inx266")
		}
	}
	rexsult, err = lang.RxFromString("1").OpDiv(lang.RxSetWithDigit(1), lang.RxFromString("128"))
	if err != nil {
		lang.SayString("inx267")
	} else {
		if !(rexsult.ToString() == "0.008") {
			lang.SayString("inx267")
		}
	}
	rexsult, err = lang.RxFromString("4953734675913.065314738743322579").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("0218.932010396534371704930714860E+797"))
	if err != nil {
		lang.SayString("divx3001")
	} else {
		if !(rexsult.ToString() == "2.262681764507965005284080800438E-787") {
			lang.SayString("divx3001")
		}
	}
	rexsult, err = lang.RxFromString("9641.684323386955881595490347910E-844").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-78864532047.12287484430980636798E+934"))
	if err != nil {
		lang.SayString("divx3002")
	} else {
		if !(rexsult.ToString() == "-1.222562801441069667849402782716E-1785") {
			lang.SayString("divx3002")
		}
	}
	rexsult, err = lang.RxFromString("-1.028048571628326871054964307774E+529").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("49200008645699.35577937582714739"))
	if err != nil {
		lang.SayString("divx3003")
	} else {
		if !(rexsult.ToString() == "-2.089529249946971482861843692465E+515") {
			lang.SayString("divx3003")
		}
	}
	rexsult, err = lang.RxFromString("479084.8561808930525417735205519").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("084157571054.2691784660983989931"))
	if err != nil {
		lang.SayString("divx3004")
	} else {
		if !(rexsult.ToString() == "0.000005692712493709617905493710207969") {
			lang.SayString("divx3004")
		}
	}
	rexsult, err = lang.RxFromString("-0363750788.573782205664349562931").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-3172.080934464133691909905980096"))
	if err != nil {
		lang.SayString("divx3005")
	} else {
		if !(rexsult.ToString() == "114672.606433742016709629529089") {
			lang.SayString("divx3005")
		}
	}
	rexsult, err = lang.RxFromString("1381026551423669919010191878449").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-82.66614775445371254999357800739"))
	if err != nil {
		lang.SayString("divx3006")
	} else {
		if !(rexsult.ToString() == "-16706071214613552377376639557.9") {
			lang.SayString("divx3006")
		}
	}
	rexsult, err = lang.RxFromString("4627.026960423072127953556635585").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-4410583132901.830017479741231131"))
	if err != nil {
		lang.SayString("divx3007")
	} else {
		if !(rexsult.ToString() == "-1.049073743992404570569003129346E-9") {
			lang.SayString("divx3007")
		}
	}
	rexsult, err = lang.RxFromString("75353574493.84484153484918212042").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-8684111695095849922263044191221"))
	if err != nil {
		lang.SayString("divx3008")
	} else {
		if !(rexsult.ToString() == "-8.677177026223536475531592432118E-21") {
			lang.SayString("divx3008")
		}
	}
	rexsult, err = lang.RxFromString("6907058.216435355874729592373011").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("2.857005446917670515662398741545"))
	if err != nil {
		lang.SayString("divx3009")
	} else {
		if !(rexsult.ToString() == "2417586.646146283856436864121104") {
			lang.SayString("divx3009")
		}
	}
	rexsult, err = lang.RxFromString("-38949530427253.24030680468677190").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("712168021.1265384466442576619064E-992"))
	if err != nil {
		lang.SayString("divx3010")
	} else {
		if !(rexsult.ToString() == "-5.469149031100999700489221122509E+996") {
			lang.SayString("divx3010")
		}
	}
	rexsult, err = lang.RxFromString("-0708069.025667471996378081482549").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-562842.4701520787831018732202804"))
	if err != nil {
		lang.SayString("divx3011")
	} else {
		if !(rexsult.ToString() == "1.258023449218665608349145394069") {
			lang.SayString("divx3011")
		}
	}
	rexsult, err = lang.RxFromString("4055087.246994644709729942673976").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-43183146921897.67383476104084575E+211"))
	if err != nil {
		lang.SayString("divx3012")
	} else {
		if !(rexsult.ToString() == "-9.390439409913307906923909630247E-219") {
			lang.SayString("divx3012")
		}
	}
	rexsult, err = lang.RxFromString("4502895892520.396581348110906909E-512").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-815.9047305921862348263521876034"))
	if err != nil {
		lang.SayString("divx3013")
	} else {
		if !(rexsult.ToString() == "-5.518899111238367862234798433551E-503") {
			lang.SayString("divx3013")
		}
	}
	rexsult, err = lang.RxFromString("467.6721295072628100260239179865").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-02.07155073395573569852316073025"))
	if err != nil {
		lang.SayString("divx3014")
	} else {
		if !(rexsult.ToString() == "-225.7594380101027705997496045999") {
			lang.SayString("divx3014")
		}
	}
	rexsult, err = lang.RxFromString("2.156795313311150143949997552501E-571").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-8677147.586389401682712180146855"))
	if err != nil {
		lang.SayString("divx3015")
	} else {
		if !(rexsult.ToString() == "-2.485604044230163799604243529005E-578") {
			lang.SayString("divx3015")
		}
	}
	rexsult, err = lang.RxFromString("-974953.2801637208368002585822457").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-693095793.3667578067802698191246"))
	if err != nil {
		lang.SayString("divx3016")
	} else {
		if !(rexsult.ToString() == "0.001406664546942092941961075608769") {
			lang.SayString("divx3016")
		}
	}
	rexsult, err = lang.RxFromString("-7634680140009571846155654339781").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("3009630949502.035852433434214413E-490"))
	if err != nil {
		lang.SayString("divx3017")
	} else {
		if !(rexsult.ToString() == "-2.536749610869326753741024659948E+508") {
			lang.SayString("divx3017")
		}
	}
	rexsult, err = lang.RxFromString("262273.0222851186523650889896428E-624").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("74177.21073338090843145838835480"))
	if err != nil {
		lang.SayString("divx3018")
	} else {
		if !(rexsult.ToString() == "3.535762799545274329358292065343E-624") {
			lang.SayString("divx3018")
		}
	}
	rexsult, err = lang.RxFromString("-8036052748815903177624716581732").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-066677357.4438809548850966167573"))
	if err != nil {
		lang.SayString("divx3019")
	} else {
		if !(rexsult.ToString() == "120521464210387351732732.6271469") {
			lang.SayString("divx3019")
		}
	}
	rexsult, err = lang.RxFromString("883429.5928031498103637713570166E+765").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-43978.97283712939198111043032726"))
	if err != nil {
		lang.SayString("divx3020")
	} else {
		if !(rexsult.ToString() == "-2.008754492913739633208672455025E+766") {
			lang.SayString("divx3020")
		}
	}
	rexsult, err = lang.RxFromString("24791301060.37938360567775506973").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-5613327866480.322649080205877564"))
	if err != nil {
		lang.SayString("divx3021")
	} else {
		if !(rexsult.ToString() == "-0.004416506865458415275182120038399") {
			lang.SayString("divx3021")
		}
	}
	rexsult, err = lang.RxFromString("-930711443.9474781586162910776139").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-740.3860979292775472622798348030"))
	if err != nil {
		lang.SayString("divx3022")
	} else {
		if !(rexsult.ToString() == "1257062.290270583507131602958799") {
			lang.SayString("divx3022")
		}
	}
	rexsult, err = lang.RxFromString("2358276428765.064191082773385539").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("214.3589796082328665878602304469"))
	if err != nil {
		lang.SayString("divx3023")
	} else {
		if !(rexsult.ToString() == "11001528525.07089502152736489473") {
			lang.SayString("divx3023")
		}
	}
	rexsult, err = lang.RxFromString("-3.868744449795653651638308926987E+750").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("8270.472492965559872384018329418"))
	if err != nil {
		lang.SayString("divx3024")
	} else {
		if !(rexsult.ToString() == "-4.677779235812959233092739433453E+746") {
			lang.SayString("divx3024")
		}
	}
	rexsult, err = lang.RxFromString("140422069.5863246490180206814374E-447").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-567195652586.2454217069003186487"))
	if err != nil {
		lang.SayString("divx3025")
	} else {
		if !(rexsult.ToString() == "-2.475725421131866851190640203633E-451") {
			lang.SayString("divx3025")
		}
	}
	rexsult, err = lang.RxFromString("75929096475.63450425339472559646E+153").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-0945260193.503803519572604151290E+459"))
	if err != nil {
		lang.SayString("divx3026")
	} else {
		if !(rexsult.ToString() == "-8.032613347885465805613265604973E-305") {
			lang.SayString("divx3026")
		}
	}
	rexsult, err = lang.RxFromString("6312318309.142044953357460463732").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-5641317823.202274083982487558514E+628"))
	if err != nil {
		lang.SayString("divx3027")
	} else {
		if !(rexsult.ToString() == "-1.118943925332481944765809682502E-628") {
			lang.SayString("divx3027")
		}
	}
	rexsult, err = lang.RxFromString("93793652428100.52105928239469937").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("917.2571313109730433369594936416E-712"))
	if err != nil {
		lang.SayString("divx3028")
	} else {
		if !(rexsult.ToString() == "1.022544815694674972559924997256E+723") {
			lang.SayString("divx3028")
		}
	}
	rexsult, err = lang.RxFromString("98471198160.56524417578665886060").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-23994.14313393939743548945165462"))
	if err != nil {
		lang.SayString("divx3029")
	} else {
		if !(rexsult.ToString() == "-4103968.106336710126241266685434") {
			lang.SayString("divx3029")
		}
	}
	rexsult, err = lang.RxFromString("329326552.0208398002250836592043").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-02451.10065397010591546041034041"))
	if err != nil {
		lang.SayString("divx3030")
	} else {
		if !(rexsult.ToString() == "-134358.6406732917173739187421978") {
			lang.SayString("divx3030")
		}
	}
	rexsult, err = lang.RxFromString("-92980.68431371090354435763218439").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-2282178507046019721925800997065"))
	if err != nil {
		lang.SayString("divx3031")
	} else {
		if !(rexsult.ToString() == "4.074207342968196863070496994457E-26") {
			lang.SayString("divx3031")
		}
	}
	rexsult, err = lang.RxFromString("12135817762.27858606259822256987E+738").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("98.35649167872356132249561021910E-902"))
	if err != nil {
		lang.SayString("divx3032")
	} else {
		if !(rexsult.ToString() == "1.233860374149945561886955398724E+1648") {
			lang.SayString("divx3032")
		}
	}
	rexsult, err = lang.RxFromString("37.27457578793521166809739140081").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-392550.4790095035979998355569916"))
	if err != nil {
		lang.SayString("divx3033")
	} else {
		if !(rexsult.ToString() == "-0.00009495486002714264641177211062199") {
			lang.SayString("divx3033")
		}
	}
	rexsult, err = lang.RxFromString("-2787.980590304199878755265273703").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("7117631179305319208210387565324"))
	if err != nil {
		lang.SayString("divx3034")
	} else {
		if !(rexsult.ToString() == "-3.91700626243506309347514025087E-28") {
			lang.SayString("divx3034")
		}
	}
	rexsult, err = lang.RxFromString("-9890633.854609434943559831911276E+971").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-1939985729.436827777055699361237"))
	if err != nil {
		lang.SayString("divx3035")
	} else {
		if !(rexsult.ToString() == "5.098302376420396260404821158158E+968") {
			lang.SayString("divx3035")
		}
	}
	rexsult, err = lang.RxFromString("3944570323.331121750661920475191").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-17360722.28878145641394962484366"))
	if err != nil {
		lang.SayString("divx3036")
	} else {
		if !(rexsult.ToString() == "-227.2123393091837706827708196101") {
			lang.SayString("divx3036")
		}
	}
	rexsult, err = lang.RxFromString("19544.14018503427029002552872707").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("1786697762.885178994182133839546"))
	if err != nil {
		lang.SayString("divx3037")
	} else {
		if !(rexsult.ToString() == "0.00001093869404832867759234359871991") {
			lang.SayString("divx3037")
		}
	}
	rexsult, err = lang.RxFromString("-05.75485957937617757983513662981").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("5564476875.989640431173694372083"))
	if err != nil {
		lang.SayString("divx3038")
	} else {
		if !(rexsult.ToString() == "-1.034213944568271324841608825136E-9") {
			lang.SayString("divx3038")
		}
	}
	rexsult, err = lang.RxFromString("-4208820.898718069194008526302746").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("626887.7553774705678201112845462E+206"))
	if err != nil {
		lang.SayString("divx3039")
	} else {
		if !(rexsult.ToString() == "-6.713834913211527184907421856434E-206") {
			lang.SayString("divx3039")
		}
	}
	rexsult, err = lang.RxFromString("-70077195478066.30896979085821269E+549").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("4607.163248554155483681430013073"))
	if err != nil {
		lang.SayString("divx3040")
	} else {
		if !(rexsult.ToString() == "-1.521048673498997627360230078306E+559") {
			lang.SayString("divx3040")
		}
	}
	rexsult, err = lang.RxFromString("-442941.7541811527940918244383454").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-068126768.0563559819156379151016"))
	if err != nil {
		lang.SayString("divx3041")
	} else {
		if !(rexsult.ToString() == "0.006501728568934042143913111768557") {
			lang.SayString("divx3041")
		}
	}
	rexsult, err = lang.RxFromString("-040726778711.8677615616711676159").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("299691.9430345259174614997064916"))
	if err != nil {
		lang.SayString("divx3042")
	} else {
		if !(rexsult.ToString() == "-135895.4741975690872548233111888") {
			lang.SayString("divx3042")
		}
	}
	rexsult, err = lang.RxFromString("-1934197520.738366912179143085955").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("3.810751422515178400293693371519"))
	if err != nil {
		lang.SayString("divx3043")
	} else {
		if !(rexsult.ToString() == "-507563287.7312566071537233697473") {
			lang.SayString("divx3043")
		}
	}
	rexsult, err = lang.RxFromString("813262.7723533833038829559646830").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-303284822716.8282178131118185907"))
	if err != nil {
		lang.SayString("divx3044")
	} else {
		if !(rexsult.ToString() == "-0.000002681514904267770294213381485108") {
			lang.SayString("divx3044")
		}
	}
	rexsult, err = lang.RxFromString("36105954884.94621434979365589311").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("745558205.7692397481313005659523E-952"))
	if err != nil {
		lang.SayString("divx3045")
	} else {
		if !(rexsult.ToString() == "4.842808328786805821411674302686E+953") {
			lang.SayString("divx3045")
		}
	}
	rexsult, err = lang.RxFromString("-075537177538.1814516621962185490").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("26980775255.51542856483122484898"))
	if err != nil {
		lang.SayString("divx3046")
	} else {
		if !(rexsult.ToString() == "-2.799666682029089956269018541649") {
			lang.SayString("divx3046")
		}
	}
	rexsult, err = lang.RxFromString("-4223765.415319564898840040697647").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-2590590305497454185455459149918E-215"))
	if err != nil {
		lang.SayString("divx3047")
	} else {
		if !(rexsult.ToString() == "1.630425855588347356570076909053E+191") {
			lang.SayString("divx3047")
		}
	}
	rexsult, err = lang.RxFromString("-6468.903738522951259063099946195").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-7877.324314273694312164407794939E+267"))
	if err != nil {
		lang.SayString("divx3048")
	} else {
		if !(rexsult.ToString() == "8.212057140774706874666307246628E-268") {
			lang.SayString("divx3048")
		}
	}
	rexsult, err = lang.RxFromString("-9567221.183663236817239254783372E-203").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("1650.198961256061165362319471264"))
	if err != nil {
		lang.SayString("divx3049")
	} else {
		if !(rexsult.ToString() == "-5.797616777301250711985729776957E-200") {
			lang.SayString("divx3049")
		}
	}
	rexsult, err = lang.RxFromString("8812306098770.200752139142033569E-428").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("26790.17380163975186972720427030E+568"))
	if err != nil {
		lang.SayString("divx3050")
	} else {
		if !(rexsult.ToString() == "3.289379965960065573444140749635E-988") {
			lang.SayString("divx3050")
		}
	}
	rexsult, err = lang.RxFromString("80108033.12724838718736922500904").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-706207255092.7645192310078892869"))
	if err != nil {
		lang.SayString("divx3051")
	} else {
		if !(rexsult.ToString() == "-0.0001134341690057060105325397863996") {
			lang.SayString("divx3051")
		}
	}
	rexsult, err = lang.RxFromString("-37942846282.76101663789059003505").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-5.649456053942850351313869983197"))
	if err != nil {
		lang.SayString("divx3052")
	} else {
		if !(rexsult.ToString() == "6716194607.13922473503256632896") {
			lang.SayString("divx3052")
		}
	}
	rexsult, err = lang.RxFromString("92659632115305.13735437728445541").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("6483438.317862851676468094261410E-139"))
	if err != nil {
		lang.SayString("divx3053")
	} else {
		if !(rexsult.ToString() == "1.429174267919135710410529211791E+146") {
			lang.SayString("divx3053")
		}
	}
	rexsult, err = lang.RxFromString("2838948.589837595494152150647194").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("569547026247.5469563701415715960"))
	if err != nil {
		lang.SayString("divx3054")
	} else {
		if !(rexsult.ToString() == "0.000004984572755198057481907281080406") {
			lang.SayString("divx3054")
		}
	}
	rexsult, err = lang.RxFromString("524995204523.6053307941775794287E+694").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("1589600879689517100527293028553"))
	if err != nil {
		lang.SayString("divx3055")
	} else {
		if !(rexsult.ToString() == "3.302685669286670708554753139233E+675") {
			lang.SayString("divx3055")
		}
	}
	rexsult, err = lang.RxFromString("-57131573677452.15449921725097290").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("4669681430736.326858508715643769"))
	if err != nil {
		lang.SayString("divx3056")
	} else {
		if !(rexsult.ToString() == "-12.23457628210057733643575143694") {
			lang.SayString("divx3056")
		}
	}
	rexsult, err = lang.RxFromString("90794826.55528018781830463383411").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-5.471502270351231110027647216128"))
	if err != nil {
		lang.SayString("divx3057")
	} else {
		if !(rexsult.ToString() == "-16594131.20365054928428313232246") {
			lang.SayString("divx3057")
		}
	}
	rexsult, err = lang.RxFromString("58508794729.35191160840980489138").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-47060867.24988279680824397447551"))
	if err != nil {
		lang.SayString("divx3058")
	} else {
		if !(rexsult.ToString() == "-1243.257894477021678809337875304") {
			lang.SayString("divx3058")
		}
	}
	rexsult, err = lang.RxFromString("-746104.0768078474426464219416332E+006").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("9595418.300613754556671852801667E+385"))
	if err != nil {
		lang.SayString("divx3059")
	} else {
		if !(rexsult.ToString() == "-7.775628465932789700547872511745E-381") {
			lang.SayString("divx3059")
		}
	}
	rexsult, err = lang.RxFromString("55.99427632688387400403789659459E+119").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-9.170530450881612853998489340127"))
	if err != nil {
		lang.SayString("divx3060")
	} else {
		if !(rexsult.ToString() == "-6.105892851759828176655685111491E+119") {
			lang.SayString("divx3060")
		}
	}
	rexsult, err = lang.RxFromString("-41214265628.83801241467317270595").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("1015336323798389903361978271354"))
	if err != nil {
		lang.SayString("divx3061")
	} else {
		if !(rexsult.ToString() == "-4.059173759750342247620706384027E-20") {
			lang.SayString("divx3061")
		}
	}
	rexsult, err = lang.RxFromString("89937.39749201095570357557430822").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("82351554210093.60879476027800331"))
	if err != nil {
		lang.SayString("divx3062")
	} else {
		if !(rexsult.ToString() == "1.092115362662913415592930982129E-9") {
			lang.SayString("divx3062")
		}
	}
	rexsult, err = lang.RxFromString("01712661.64677082156284125486943E+359").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("57932.78435529483241552042115837E-037"))
	if err != nil {
		lang.SayString("divx3063")
	} else {
		if !(rexsult.ToString() == "2.956290925475414185960999788848E+397") {
			lang.SayString("divx3063")
		}
	}
	rexsult, err = lang.RxFromString("-2647593306.528617691373470059213").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-655531558709.4582168930191014461"))
	if err != nil {
		lang.SayString("divx3064")
	} else {
		if !(rexsult.ToString() == "0.004038849497560303158639192895544") {
			lang.SayString("divx3064")
		}
	}
	rexsult, err = lang.RxFromString("2904078690665765116603253099668E-329").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-71.45586619176091599264717047885E+787"))
	if err != nil {
		lang.SayString("divx3065")
	} else {
		if !(rexsult.ToString() == "-4.064157144036712325084472022316E-1088") {
			lang.SayString("divx3065")
		}
	}
	rexsult, err = lang.RxFromString("22094338972.39109726522477999515").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-409846549371.3900805039668417203E-499"))
	if err != nil {
		lang.SayString("divx3066")
	} else {
		if !(rexsult.ToString() == "-5.390880808019174194010224736965E+497") {
			lang.SayString("divx3066")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("82293895124.90045271504836568681"))
	if err != nil {
		lang.SayString("divx3067")
	} else {
		if !(rexsult.ToString() == "-41011408883796817797.81310977038") {
			lang.SayString("divx3067")
		}
	}
	rexsult, err = lang.RxFromString("-84172558160661.35863831029352323").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-11271.58916600931155937291904890"))
	if err != nil {
		lang.SayString("divx3068")
	} else {
		if !(rexsult.ToString() == "7467674426.467986736459678347587") {
			lang.SayString("divx3068")
		}
	}
	rexsult, err = lang.RxFromString("-70046932324614.90596396237508541E-568").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("33.63163964004608865836577297698E-918"))
	if err != nil {
		lang.SayString("divx3069")
	} else {
		if !(rexsult.ToString() == "-2.082768876995463487926920072359E+362") {
			lang.SayString("divx3069")
		}
	}
	rexsult, err = lang.RxFromString("0004125384407.053782660115680886").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-391429084.5847321402514385603223E-648"))
	if err != nil {
		lang.SayString("divx3070")
	} else {
		if !(rexsult.ToString() == "-1.053928941287132717250540955457E+649") {
			lang.SayString("divx3070")
		}
	}
	rexsult, err = lang.RxFromString("-31823131.15691583393820628480997E-440").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("92913.91582947237200286427030028E+771"))
	if err != nil {
		lang.SayString("divx3071")
	} else {
		if !(rexsult.ToString() == "-3.425012375468251447194400841658E-1209") {
			lang.SayString("divx3071")
		}
	}
	rexsult, err = lang.RxFromString("55573867888.91575330563698128150").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("599.5231614736232188354393212234"))
	if err != nil {
		lang.SayString("divx3072")
	} else {
		if !(rexsult.ToString() == "92696782.14318796763098335498657") {
			lang.SayString("divx3072")
		}
	}
	rexsult, err = lang.RxFromString("-5447727448431680878699555714796E-800").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("5487207.142687001607026665515349E-362"))
	if err != nil {
		lang.SayString("divx3073")
	} else {
		if !(rexsult.ToString() == "-9.928051387110587327889009363069E-415") {
			lang.SayString("divx3073")
		}
	}
	rexsult, err = lang.RxFromString("0418349404834.547110239542290134").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("09819915.92405288066606124554841"))
	if err != nil {
		lang.SayString("divx3074")
	} else {
		if !(rexsult.ToString() == "42602.13713335803513874339309132") {
			lang.SayString("divx3074")
		}
	}
	rexsult, err = lang.RxFromString("-262021.7565194737396448014286436").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-7983992600094836304387324162042E+390"))
	if err != nil {
		lang.SayString("divx3075")
	} else {
		if !(rexsult.ToString() == "3.281838669494274896180376328433E-416") {
			lang.SayString("divx3075")
		}
	}
	rexsult, err = lang.RxFromString("48696050631.42565380301204592392E-505").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-33868752339.85057267609967806187E+821"))
	if err != nil {
		lang.SayString("divx3076")
	} else {
		if !(rexsult.ToString() == "-1.43778696489297658200995217242E-1326") {
			lang.SayString("divx3076")
		}
	}
	rexsult, err = lang.RxFromString("95316999.19440144356471126680708").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-60791.33805057402845885978390435"))
	if err != nil {
		lang.SayString("divx3077")
	} else {
		if !(rexsult.ToString() == "-1567.937180706641856870286122623") {
			lang.SayString("divx3077")
		}
	}
	rexsult, err = lang.RxFromString("-5326702296402708234722215224979E-136").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("8032459.450998820205916538543258"))
	if err != nil {
		lang.SayString("divx3078")
	} else {
		if !(rexsult.ToString() == "-6.631471131473117487839243582873E-113") {
			lang.SayString("divx3078")
		}
	}
	rexsult, err = lang.RxFromString("67.18750684079501575335482615780E-281").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("734.1168841683438410314843011541E-854"))
	if err != nil {
		lang.SayString("divx3079")
	} else {
		if !(rexsult.ToString() == "9.152153872187460598958616592442E+571") {
			lang.SayString("divx3079")
		}
	}
	rexsult, err = lang.RxFromString("-8739299372114.092482914139281669").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("507610074.7343577029345077385838"))
	if err != nil {
		lang.SayString("divx3080")
	} else {
		if !(rexsult.ToString() == "-17216.5601257767373161213006813") {
			lang.SayString("divx3080")
		}
	}
	rexsult, err = lang.RxFromString("2454.002078468928665008217727731").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("583546039.6233842869119950982009E-147"))
	if err != nil {
		lang.SayString("divx3081")
	} else {
		if !(rexsult.ToString() == "4.205327278123112611006652533618E+141") {
			lang.SayString("divx3081")
		}
	}
	rexsult, err = lang.RxFromString("764578.5204849936912066033177429").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("64603.13571259164812609436832506"))
	if err != nil {
		lang.SayString("divx3082")
	} else {
		if !(rexsult.ToString() == "11.83500633601553578851124281417") {
			lang.SayString("divx3082")
		}
	}
	rexsult, err = lang.RxFromString("079203.7330103777716903518367560").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("846388934347.6324036132959664705"))
	if err != nil {
		lang.SayString("divx3083")
	} else {
		if !(rexsult.ToString() == "9.357841270860339858146471876044E-8") {
			lang.SayString("divx3083")
		}
	}
	rexsult, err = lang.RxFromString("-4278.581514688669249247007127899E-329").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("5.474973992953902631890208360829"))
	if err != nil {
		lang.SayString("divx3084")
	} else {
		if !(rexsult.ToString() == "-7.814797878848469282033896969532E-327") {
			lang.SayString("divx3084")
		}
	}
	rexsult, err = lang.RxFromString("60867019.81764798845468445196869E+651").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("6.149612565404080501157093851895E+817"))
	if err != nil {
		lang.SayString("divx3085")
	} else {
		if !(rexsult.ToString() == "9.897699923417617920996187420968E-160") {
			lang.SayString("divx3085")
		}
	}
	rexsult, err = lang.RxFromString("18554417738217.62218590965803605E-382").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-0894505909529.052378474618435782E+527"))
	if err != nil {
		lang.SayString("divx3086")
	} else {
		if !(rexsult.ToString() == "-2.074264411286709228674841672954E-908") {
			lang.SayString("divx3086")
		}
	}
	rexsult, err = lang.RxFromString("69073355517144.36356688642213839").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("997784782535.6104634823627327033E+116"))
	if err != nil {
		lang.SayString("divx3087")
	} else {
		if !(rexsult.ToString() == "6.922670772910807388395384866884E-115") {
			lang.SayString("divx3087")
		}
	}
	rexsult, err = lang.RxFromString("450282259072.8657099359104277477").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-1791307965314309175477911369824"))
	if err != nil {
		lang.SayString("divx3088")
	} else {
		if !(rexsult.ToString() == "-2.513706564096350714213771006483E-19") {
			lang.SayString("divx3088")
		}
	}
	rexsult, err = lang.RxFromString("954678411.7838149266455177850037").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("142988.7096204254529284334278794"))
	if err != nil {
		lang.SayString("divx3089")
	} else {
		if !(rexsult.ToString() == "6676.599951968811589335427770046") {
			lang.SayString("divx3089")
		}
	}
	rexsult, err = lang.RxFromString("-9244530976.220812127155852389807E+557").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("541089.4715446858896619078627941"))
	if err != nil {
		lang.SayString("divx3090")
	} else {
		if !(rexsult.ToString() == "-1.708503207395591002370649848757E+561") {
			lang.SayString("divx3090")
		}
	}
	rexsult, err = lang.RxFromString("-75492024.20990197005974241975449").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-14760421311348.35269044633000927"))
	if err != nil {
		lang.SayString("divx3091")
	} else {
		if !(rexsult.ToString() == "0.000005114489797920668836278344635108") {
			lang.SayString("divx3091")
		}
	}
	rexsult, err = lang.RxFromString("317747.6972215715434186596178036E-452").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("24759763.33144824613591228097330E+092"))
	if err != nil {
		lang.SayString("divx3092")
	} else {
		if !(rexsult.ToString() == "1.283322837007852247594216151634E-546") {
			lang.SayString("divx3092")
		}
	}
	rexsult, err = lang.RxFromString("-8.153334430358647134334545353427").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-9.717872025814596548462854853522"))
	if err != nil {
		lang.SayString("divx3093")
	} else {
		if !(rexsult.ToString() == "0.8390040956188859972044344532019") {
			lang.SayString("divx3093")
		}
	}
	rexsult, err = lang.RxFromString("7.267345197492967332320456062961E-478").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("5054015481833.263541189916208065"))
	if err != nil {
		lang.SayString("divx3094")
	} else {
		if !(rexsult.ToString() == "1.437934890309606594895299558654E-490") {
			lang.SayString("divx3094")
		}
	}
	rexsult, err = lang.RxFromString("-1223354029.862567054230912271171").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("8135774223401322785475014855625"))
	if err != nil {
		lang.SayString("divx3095")
	} else {
		if !(rexsult.ToString() == "-1.503672540892020337688277553692E-22") {
			lang.SayString("divx3095")
		}
	}
	rexsult, err = lang.RxFromString("285397644111.5655679961211349982E+645").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-2479499427613157519359627280704"))
	if err != nil {
		lang.SayString("divx3096")
	} else {
		if !(rexsult.ToString() == "-1.151029280076495626421134733122E+626") {
			lang.SayString("divx3096")
		}
	}
	rexsult, err = lang.RxFromString("-4673112.663442366293812346653429").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("-3429.998403142546001438238460958"))
	if err != nil {
		lang.SayString("divx3097")
	} else {
		if !(rexsult.ToString() == "1362.424151323477505064686589716") {
			lang.SayString("divx3097")
		}
	}
	rexsult, err = lang.RxFromString("88.96492479681278079861456051103").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("386939.4621006514751889096510923E+139"))
	if err != nil {
		lang.SayString("divx3098")
	} else {
		if !(rexsult.ToString() == "2.299194926095985647821385937618E-143") {
			lang.SayString("divx3098")
		}
	}
	rexsult, err = lang.RxFromString("064326846.4286437304788069444326E-942").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("92.23649942010862087149015091350"))
	if err != nil {
		lang.SayString("divx3099")
	} else {
		if !(rexsult.ToString() == "6.974120530708230229344349531719E-937") {
			lang.SayString("divx3099")
		}
	}
	rexsult, err = lang.RxFromString("504507.0043949324433170405699360").OpDiv(lang.RxSetWithDigit(31), lang.RxFromString("605387.7175522955344659311072099"))
	if err != nil {
		lang.SayString("divx3100")
	} else {
		if !(rexsult.ToString() == "0.8333618105678718895216067463832") {
			lang.SayString("divx3100")
		}
	}
	rexsult, err = lang.RxFromString("1.5283550163839789319142430253644").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-1.6578158484822969520405291379492"))
	if err != nil {
		lang.SayString("divx3201")
	} else {
		if !(rexsult.ToString() == "-0.9219087981232431363028298011028") {
			lang.SayString("divx3201")
		}
	}
	rexsult, err = lang.RxFromString("-622903030605.2867503937836507326").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("6519388607.1331855704471328795821"))
	if err != nil {
		lang.SayString("divx3202")
	} else {
		if !(rexsult.ToString() == "-95.54623418578511049167689415351") {
			lang.SayString("divx3202")
		}
	}
	rexsult, err = lang.RxFromString("-5675915.2465457487632250245209054").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("73913909880.381367895205086027416"))
	if err != nil {
		lang.SayString("divx3203")
	} else {
		if !(rexsult.ToString() == "-0.000076790894376056827552388054657082") {
			lang.SayString("divx3203")
		}
	}
	rexsult, err = lang.RxFromString("97.647321172555144900685605318635").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("4.8620911587547548751209841570885"))
	if err != nil {
		lang.SayString("divx3204")
	} else {
		if !(rexsult.ToString() == "20.083399916665466374741708949621") {
			lang.SayString("divx3204")
		}
	}
	rexsult, err = lang.RxFromString("-9717253267024.5380651435435603552").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-2669.2539695193820424002013488480E+363"))
	if err != nil {
		lang.SayString("divx3205")
	} else {
		if !(rexsult.ToString() == "3.6404378818903462695633337631098E-354") {
			lang.SayString("divx3205")
		}
	}
	rexsult, err = lang.RxFromString("-4.0817391717190128506083001702335E-767").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("12772.807105920712660991033689206"))
	if err != nil {
		lang.SayString("divx3206")
	} else {
		if !(rexsult.ToString() == "-3.1956477052150593175206769891434E-771") {
			lang.SayString("divx3206")
		}
	}
	rexsult, err = lang.RxFromString("68625322655934146845003028928647").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-59.634169944840280159782488098700"))
	if err != nil {
		lang.SayString("divx3207")
	} else {
		if !(rexsult.ToString() == "-1150771826276954946844322988192.5") {
			lang.SayString("divx3207")
		}
	}
	rexsult, err = lang.RxFromString("732515.76532049290815665856727641").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-92134479835821.319619827023729829"))
	if err != nil {
		lang.SayString("divx3208")
	} else {
		if !(rexsult.ToString() == "-7.9505063318943846655593887991914E-9") {
			lang.SayString("divx3208")
		}
	}
	rexsult, err = lang.RxFromString("-30.458011942978338421676454778733").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-5023372024597665102336430410403E+831"))
	if err != nil {
		lang.SayString("divx3209")
	} else {
		if !(rexsult.ToString() == "6.063260255031141082148300130501E-861") {
			lang.SayString("divx3209")
		}
	}
	rexsult, err = lang.RxFromString("-89640.094149414644660480286430462").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-58703419758.800889903227509215474"))
	if err != nil {
		lang.SayString("divx3210")
	} else {
		if !(rexsult.ToString() == "0.0000015269995260536025237167199970238") {
			lang.SayString("divx3210")
		}
	}
	rexsult, err = lang.RxFromString("458653.1567870081810052917714259").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("18353106238.516235116080449814053E-038"))
	if err != nil {
		lang.SayString("divx3211")
	} else {
		if !(rexsult.ToString() == "2.4990492117594160215641311760498E+33") {
			lang.SayString("divx3211")
		}
	}
	rexsult, err = lang.RxFromString("913391.42744224458216174967853722").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-21051638.816432817393202262710630E+432"))
	if err != nil {
		lang.SayString("divx3212")
	} else {
		if !(rexsult.ToString() == "-4.3388138824102151127273259092613E-434") {
			lang.SayString("divx3212")
		}
	}
	rexsult, err = lang.RxFromString("-917591456829.12109027484399536567").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-28892177726858026955513438843371E+708"))
	if err != nil {
		lang.SayString("divx3213")
	} else {
		if !(rexsult.ToString() == "3.1759165595057674196644927106447E-728") {
			lang.SayString("divx3213")
		}
	}
	rexsult, err = lang.RxFromString("34938410840645.913399699219228218").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("30.818220393242402846077755480548"))
	if err != nil {
		lang.SayString("divx3214")
	} else {
		if !(rexsult.ToString() == "1133693327999.7879503260098666966") {
			lang.SayString("divx3214")
		}
	}
	rexsult, err = lang.RxFromString("6034.7374411022598081745006769023E-517").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("29771833428054709077850588904653"))
	if err != nil {
		lang.SayString("divx3215")
	} else {
		if !(rexsult.ToString() == "2.0269955680376683526099444523691E-545") {
			lang.SayString("divx3215")
		}
	}
	rexsult, err = lang.RxFromString("-5565747671734.1686009705574503152").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-490.30899494881071282787487030303"))
	if err != nil {
		lang.SayString("divx3216")
	} else {
		if !(rexsult.ToString() == "11351510433.365074871574519756245") {
			lang.SayString("divx3216")
		}
	}
	rexsult, err = lang.RxFromString("319545511.89203495546689273564728E+036").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-2955943533943321899150310192061"))
	if err != nil {
		lang.SayString("divx3217")
	} else {
		if !(rexsult.ToString() == "-108102711781422.68663084859902931") {
			lang.SayString("divx3217")
		}
	}
	rexsult, err = lang.RxFromString("-36852134.84194296250843579428931").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-5830629.8347085025808716360357940"))
	if err != nil {
		lang.SayString("divx3218")
	} else {
		if !(rexsult.ToString() == "6.320438080731865547545904741016") {
			lang.SayString("divx3218")
		}
	}
	rexsult, err = lang.RxFromString("8.6021905001798578659275880018221E-374").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-39505285344943.729681835377530908"))
	if err != nil {
		lang.SayString("divx3219")
	} else {
		if !(rexsult.ToString() == "-2.1774783867700502002511486885272E-387") {
			lang.SayString("divx3219")
		}
	}
	rexsult, err = lang.RxFromString("-54863165.152174109720312887805017").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("736.1398476560169141105328256628"))
	if err != nil {
		lang.SayString("divx3220")
	} else {
		if !(rexsult.ToString() == "-74528.182826764384088602813142847") {
			lang.SayString("divx3220")
		}
	}
	rexsult, err = lang.RxFromString("-3263743464517851012531708810307").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("2457206.2471248382136273643208109"))
	if err != nil {
		lang.SayString("divx3221")
	} else {
		if !(rexsult.ToString() == "-1328233422952076975055082.5768082") {
			lang.SayString("divx3221")
		}
	}
	rexsult, err = lang.RxFromString("2856586744.0548637797291151154902E-895").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("953545637646.57694835860339582821E+080"))
	if err != nil {
		lang.SayString("divx3222")
	} else {
		if !(rexsult.ToString() == "2.99575251700079800087128289683E-978") {
			lang.SayString("divx3222")
		}
	}
	rexsult, err = lang.RxFromString("5624157233.3433661009203529937625").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("626098409265.93738586750090160638"))
	if err != nil {
		lang.SayString("divx3223")
	} else {
		if !(rexsult.ToString() == "0.0089828645946207580492752544218316") {
			lang.SayString("divx3223")
		}
	}
	rexsult, err = lang.RxFromString("-213499362.91476998701834067092611").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("419272438.02555757699863022643444"))
	if err != nil {
		lang.SayString("divx3224")
	} else {
		if !(rexsult.ToString() == "-0.50921392286166855779828061147786") {
			lang.SayString("divx3224")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("4.8046009735990873395936309640543"))
	if err != nil {
		lang.SayString("divx3225")
	} else {
		if !(rexsult.ToString() == "6300.1252178837655359831527173832") {
			lang.SayString("divx3225")
		}
	}
	rexsult, err = lang.RxFromString("47.525676459351505682005359699680E+704").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-58631943508436657383369760970210"))
	if err != nil {
		lang.SayString("divx3226")
	} else {
		if !(rexsult.ToString() == "-8.1057651538555854520994438038537E+673") {
			lang.SayString("divx3226")
		}
	}
	rexsult, err = lang.RxFromString("-74396862273800.625679130772935550").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-115616605.52826981284183992013157"))
	if err != nil {
		lang.SayString("divx3227")
	} else {
		if !(rexsult.ToString() == "643479.03948459716424778005613112") {
			lang.SayString("divx3227")
		}
	}
	rexsult, err = lang.RxFromString("67585560.562561229497720955705979").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("826.96290288608566737177503451613"))
	if err != nil {
		lang.SayString("divx3228")
	} else {
		if !(rexsult.ToString() == "81727.43943735424878985271558651") {
			lang.SayString("divx3228")
		}
	}
	rexsult, err = lang.RxFromString("6877386868.9498051860742298735156E-232").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("390.3154289860643509393769754551"))
	if err != nil {
		lang.SayString("divx3229")
	} else {
		if !(rexsult.ToString() == "1.7620074325054038174571564409871E-225") {
			lang.SayString("divx3229")
		}
	}
	rexsult, err = lang.RxFromString("-1647335.201144609256134925838937").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-186654823782.50459802235024230856"))
	if err != nil {
		lang.SayString("divx3230")
	} else {
		if !(rexsult.ToString() == "0.0000088255699357876233458660331146583") {
			lang.SayString("divx3230")
		}
	}
	rexsult, err = lang.RxFromString("41407818140948.866630923934138155").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-5156.7624534000311342310106671627E-963"))
	if err != nil {
		lang.SayString("divx3231")
	} else {
		if !(rexsult.ToString() == "-8.0298091128179204076796507697517E+972") {
			lang.SayString("divx3231")
		}
	}
	rexsult, err = lang.RxFromString("-6.6547424012516834662011706165297").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-574454585580.06215974139884746441"))
	if err != nil {
		lang.SayString("divx3232")
	} else {
		if !(rexsult.ToString() == "1.1584453442097568745411568037078E-11") {
			lang.SayString("divx3232")
		}
	}
	rexsult, err = lang.RxFromString("-27627.758745381267780885067447671").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-23385972189441.709586433758111062"))
	if err != nil {
		lang.SayString("divx3233")
	} else {
		if !(rexsult.ToString() == "1.1813816642548920194709898111624E-9") {
			lang.SayString("divx3233")
		}
	}
	rexsult, err = lang.RxFromString("209819.74379099914752963711944307E-228").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-440230.6700989532467831370320266E-960"))
	if err != nil {
		lang.SayString("divx3234")
	} else {
		if !(rexsult.ToString() == "-4.7661318949867060595545765053187E+731") {
			lang.SayString("divx3234")
		}
	}
	rexsult, err = lang.RxFromString("2.3488457600415474270259330865184").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("9182434.6660212482500376220508605E-612"))
	if err != nil {
		lang.SayString("divx3235")
	} else {
		if !(rexsult.ToString() == "2.5579771002708402753412266574941E+605") {
			lang.SayString("divx3235")
		}
	}
	rexsult, err = lang.RxFromString("-5107586300197.9703941034404557409").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("56609.05486055057838678039496686E-768"))
	if err != nil {
		lang.SayString("divx3236")
	} else {
		if !(rexsult.ToString() == "-9.0225606358909877855326357402242E+775") {
			lang.SayString("divx3236")
		}
	}
	rexsult, err = lang.RxFromString("-70454070095869.70717871212601390").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-6200178.370249260009168888392406"))
	if err != nil {
		lang.SayString("divx3237")
	} else {
		if !(rexsult.ToString() == "11363232.779549422490548997517194") {
			lang.SayString("divx3237")
		}
	}
	rexsult, err = lang.RxFromString("29119.220621511046558757900645228").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("3517612.8810761470018676311863010E-222"))
	if err != nil {
		lang.SayString("divx3238")
	} else {
		if !(rexsult.ToString() == "8.2781197380089684063239752337467E+219") {
			lang.SayString("divx3238")
		}
	}
	rexsult, err = lang.RxFromString("-5168.2214111091132913776042214525").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-5690274.0971173476527123568627720"))
	if err != nil {
		lang.SayString("divx3239")
	} else {
		if !(rexsult.ToString() == "0.00090825526554639915580539316714451") {
			lang.SayString("divx3239")
		}
	}
	rexsult, err = lang.RxFromString("33783.060857197067391462144517964").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-2070.0806959465088198508322815406"))
	if err != nil {
		lang.SayString("divx3240")
	} else {
		if !(rexsult.ToString() == "-16.31968305551989288139435844922") {
			lang.SayString("divx3240")
		}
	}
	rexsult, err = lang.RxFromString("42207435091050.840296353874733169E-905").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("73330633078.828216018536326743325E+976"))
	if err != nil {
		lang.SayString("divx3241")
	} else {
		if !(rexsult.ToString() == "5.7557712676064206636178247554056E-1879") {
			lang.SayString("divx3241")
		}
	}
	rexsult, err = lang.RxFromString("-71800.806700868784841045406679641").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-39617456964250697902519150526701"))
	if err != nil {
		lang.SayString("divx3242")
	} else {
		if !(rexsult.ToString() == "1.8123527405017220178579049964126E-27") {
			lang.SayString("divx3242")
		}
	}
	rexsult, err = lang.RxFromString("53627480801.631504892310016062160").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("328259.56947661049313311983109503"))
	if err != nil {
		lang.SayString("divx3243")
	} else {
		if !(rexsult.ToString() == "163369.13159039717901500465109839") {
			lang.SayString("divx3243")
		}
	}
	rexsult, err = lang.RxFromString("-5052601598.5559371338428368438728").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-97855372.224321664785314782556064"))
	if err != nil {
		lang.SayString("divx3244")
	} else {
		if !(rexsult.ToString() == "51.633359351732432283879274192947") {
			lang.SayString("divx3244")
		}
	}
	rexsult, err = lang.RxFromString("4208134320733.7069742988228068191E+146").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("4270869.1760149477598920960628392E+471"))
	if err != nil {
		lang.SayString("divx3245")
	} else {
		if !(rexsult.ToString() == "9.8531098643021951048744078027283E-320") {
			lang.SayString("divx3245")
		}
	}
	rexsult, err = lang.RxFromString("-8.5077009657942581515590471189084E+308").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("9652145155.374217047842114042376E-250"))
	if err != nil {
		lang.SayString("divx3246")
	} else {
		if !(rexsult.ToString() == "-8.814311045723608997807041904797E+548") {
			lang.SayString("divx3246")
		}
	}
	rexsult, err = lang.RxFromString("-9504.9703032286960790904181078063E+619").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-86.245956949049186533469206485003"))
	if err != nil {
		lang.SayString("divx3247")
	} else {
		if !(rexsult.ToString() == "1.1020772033225707125391212519421E+621") {
			lang.SayString("divx3247")
		}
	}
	rexsult, err = lang.RxFromString("-440220916.66716743026896931194749").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-102725.01594377871560564824358775"))
	if err != nil {
		lang.SayString("divx3248")
	} else {
		if !(rexsult.ToString() == "4285.4305022264473269770246126234") {
			lang.SayString("divx3248")
		}
	}
	rexsult, err = lang.RxFromString("-46.250735086006350517943464758019").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("14656357555174.263295266074908024"))
	if err != nil {
		lang.SayString("divx3249")
	} else {
		if !(rexsult.ToString() == "-3.1556773169523313932207725522866E-12") {
			lang.SayString("divx3249")
		}
	}
	rexsult, err = lang.RxFromString("-61641121299391.316420647102699627E+763").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-91896469863.461006903590004188187E+474"))
	if err != nil {
		lang.SayString("divx3250")
	} else {
		if !(rexsult.ToString() == "6.7076702065897819604716946852581E+291") {
			lang.SayString("divx3250")
		}
	}
	rexsult, err = lang.RxFromString("96668419802749.555738010239087449E-838").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-19498732131365824921639467044927E-542"))
	if err != nil {
		lang.SayString("divx3251")
	} else {
		if !(rexsult.ToString() == "-4.9576772044192514715453215933704E-314") {
			lang.SayString("divx3251")
		}
	}
	rexsult, err = lang.RxFromString("-8534543911197995906031245719519E+124").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("16487117050031.594886541650897974"))
	if err != nil {
		lang.SayString("divx3252")
	} else {
		if !(rexsult.ToString() == "-5.1764925822381062287959523371316E+141") {
			lang.SayString("divx3252")
		}
	}
	rexsult, err = lang.RxFromString("-62663404777.352508979582846931050").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("9.2570938837239134052589184917186E+916"))
	if err != nil {
		lang.SayString("divx3253")
	} else {
		if !(rexsult.ToString() == "-6.7692307720384142592597124956951E-907") {
			lang.SayString("divx3253")
		}
	}
	rexsult, err = lang.RxFromString("1.744601214474560992754529320172E-827").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-17.353669504253419489494030651507E-631"))
	if err != nil {
		lang.SayString("divx3254")
	} else {
		if !(rexsult.ToString() == "-1.0053212169604565230497117966004E-197") {
			lang.SayString("divx3254")
		}
	}
	rexsult, err = lang.RxFromString("0367191549036702224827734853471").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("4410320662415266533763143837742E+721"))
	if err != nil {
		lang.SayString("divx3255")
	} else {
		if !(rexsult.ToString() == "8.3257335949720619093963917942525E-723") {
			lang.SayString("divx3255")
		}
	}
	rexsult, err = lang.RxFromString("097704116.4492566721965710365073").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-96736.400939809433556067504289145"))
	if err != nil {
		lang.SayString("divx3256")
	} else {
		if !(rexsult.ToString() == "-1010.0036335861757252324592571874") {
			lang.SayString("divx3256")
		}
	}
	rexsult, err = lang.RxFromString("19533298.147150158931958733807878").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("80.141668338350708476637377647025E-641"))
	if err != nil {
		lang.SayString("divx3257")
	} else {
		if !(rexsult.ToString() == "2.4373460837728485395672882395171E+646") {
			lang.SayString("divx3257")
		}
	}
	rexsult, err = lang.RxFromString("-23765003221220177430797028997378").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-15203369569.373411506379096871224E-945"))
	if err != nil {
		lang.SayString("divx3258")
	} else {
		if !(rexsult.ToString() == "1.5631405336020930064824135669541E+966") {
			lang.SayString("divx3258")
		}
	}
	rexsult, err = lang.RxFromString("129255.41937932433359193338910552E+932").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-281253953.38990382799508873560320"))
	if err != nil {
		lang.SayString("divx3259")
	} else {
		if !(rexsult.ToString() == "-4.5956836453828213050223260551064E+928") {
			lang.SayString("divx3259")
		}
	}
	rexsult, err = lang.RxFromString("-86863.276249466008289214762980838").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("531.50602652732088208397655484476"))
	if err != nil {
		lang.SayString("divx3260")
	} else {
		if !(rexsult.ToString() == "-163.42858201815891408475902229649") {
			lang.SayString("divx3260")
		}
	}
	rexsult, err = lang.RxFromString("-40707.169006771111855573524157083").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-68795521421321853333274411827749"))
	if err != nil {
		lang.SayString("divx3261")
	} else {
		if !(rexsult.ToString() == "5.9171248601300236694386185513139E-28") {
			lang.SayString("divx3261")
		}
	}
	rexsult, err = lang.RxFromString("-90838752568673.728630494658778003E+095").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-738.01370301217606577533107981431"))
	if err != nil {
		lang.SayString("divx3262")
	} else {
		if !(rexsult.ToString() == "1.2308545518588430875268553851424E+106") {
			lang.SayString("divx3262")
		}
	}
	rexsult, err = lang.RxFromString("-4245360967593.9206771555839718158E-682").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-3.119606239042530207103203508057"))
	if err != nil {
		lang.SayString("divx3263")
	} else {
		if !(rexsult.ToString() == "1.3608643662980066356437236081969E-670") {
			lang.SayString("divx3263")
		}
	}
	rexsult, err = lang.RxFromString("-3422145405774.0800213000547612240E-023").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-60810.964656409650839011321706310"))
	if err != nil {
		lang.SayString("divx3264")
	} else {
		if !(rexsult.ToString() == "5.627513763528788287591412474265E-16") {
			lang.SayString("divx3264")
		}
	}
	rexsult, err = lang.RxFromString("-24521811.07649485796598387627478E-661").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-94860846133404815410816234000694"))
	if err != nil {
		lang.SayString("divx3265")
	} else {
		if !(rexsult.ToString() == "2.5850297647576657819483988845904E-686") {
			lang.SayString("divx3265")
		}
	}
	rexsult, err = lang.RxFromString("-5042529937498.8944492248538951438").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("3891904674.4549170968807145612549"))
	if err != nil {
		lang.SayString("divx3266")
	} else {
		if !(rexsult.ToString() == "-1295.6457979549894351378127413283") {
			lang.SayString("divx3266")
		}
	}
	rexsult, err = lang.RxFromString("-535824163.54531747646293693868651E-665").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("2732988.5891363639325008206099712"))
	if err != nil {
		lang.SayString("divx3267")
	} else {
		if !(rexsult.ToString() == "-1.9605795855687791246611683328346E-663") {
			lang.SayString("divx3267")
		}
	}
	rexsult, err = lang.RxFromString("24032.702008553084252925140858134E-509").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("52864854.899420632375589206704068"))
	if err != nil {
		lang.SayString("divx3268")
	} else {
		if !(rexsult.ToString() == "4.5460641203455697917573431961511E-513") {
			lang.SayString("divx3268")
		}
	}
	rexsult, err = lang.RxFromString("71553220259.938950698030519906727E-496").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("754.44220417415325444943566016062"))
	if err != nil {
		lang.SayString("divx3269")
	} else {
		if !(rexsult.ToString() == "9.4842547068617879794218050008353E-489") {
			lang.SayString("divx3269")
		}
	}
	rexsult, err = lang.RxFromString("35572.960284795962697740953932508").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("520.39506364687594082725754878910E-731"))
	if err != nil {
		lang.SayString("divx3270")
	} else {
		if !(rexsult.ToString() == "6.8357605153869556504869061469852E+732") {
			lang.SayString("divx3270")
		}
	}
	rexsult, err = lang.RxFromString("53035405018123760598334895413057E+818").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-9558464247240.4476790042911379151"))
	if err != nil {
		lang.SayString("divx3271")
	} else {
		if !(rexsult.ToString() == "-5.5485278436266802470202487233285E+836") {
			lang.SayString("divx3271")
		}
	}
	rexsult, err = lang.RxFromString("95.490751127249945886828257312118").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("987.01498316307365714167410690192E+133"))
	if err != nil {
		lang.SayString("divx3272")
	} else {
		if !(rexsult.ToString() == "9.6747012716293341927566515915016E-135") {
			lang.SayString("divx3272")
		}
	}
	rexsult, err = lang.RxFromString("69434850287.460788550936730296153").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-35119136549015044241569827542264"))
	if err != nil {
		lang.SayString("divx3273")
	} else {
		if !(rexsult.ToString() == "-1.9771229338327273644129394734299E-21") {
			lang.SayString("divx3273")
		}
	}
	rexsult, err = lang.RxFromString("-392.22739924621965621739098725407").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-65551274.987160998195282109612136"))
	if err != nil {
		lang.SayString("divx3274")
	} else {
		if !(rexsult.ToString() == "0.0000059835205237890809449684317245033") {
			lang.SayString("divx3274")
		}
	}
	rexsult, err = lang.RxFromString("6413265.4423561191792972085539457").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("24514.222704714139350026165721146"))
	if err != nil {
		lang.SayString("divx3275")
	} else {
		if !(rexsult.ToString() == "261.61406460270241498757868681883") {
			lang.SayString("divx3275")
		}
	}
	rexsult, err = lang.RxFromString("-6.9667706389122107760046184064057E+487").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("32.405810703971538278419625527234"))
	if err != nil {
		lang.SayString("divx3276")
	} else {
		if !(rexsult.ToString() == "-2.1498522911689997341347293419761E+486") {
			lang.SayString("divx3276")
		}
	}
	rexsult, err = lang.RxFromString("378204716633.40024100602896057615").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-0300218714378.322231269606216516"))
	if err != nil {
		lang.SayString("divx3277")
	} else {
		if !(rexsult.ToString() == "-1.2597639604731753284599748820876") {
			lang.SayString("divx3277")
		}
	}
	rexsult, err = lang.RxFromString("-44234.512012457148027685282969235E+501").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("2132572.4571987908375002100894927"))
	if err != nil {
		lang.SayString("divx3278")
	} else {
		if !(rexsult.ToString() == "-2.0742325477916347193181603963257E+499") {
			lang.SayString("divx3278")
		}
	}
	rexsult, err = lang.RxFromString("-3554.5895974968741465654022772100E-073").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("9752.0428746722497621936998533848E+516"))
	if err != nil {
		lang.SayString("divx3279")
	} else {
		if !(rexsult.ToString() == "-3.6449692061227100572768330912162E-590") {
			lang.SayString("divx3279")
		}
	}
	rexsult, err = lang.RxFromString("750187685.63632608923397234391668").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("4633194252863.6730625972669246025"))
	if err != nil {
		lang.SayString("divx3280")
	} else {
		if !(rexsult.ToString() == "0.00016191587157664541463871807382759") {
			lang.SayString("divx3280")
		}
	}
	rexsult, err = lang.RxFromString("30190034242853.251165951457589386E-028").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("8038885676.3204238322976087799751E+018"))
	if err != nil {
		lang.SayString("divx3281")
	} else {
		if !(rexsult.ToString() == "3.755499886231980729590334896028E-43") {
			lang.SayString("divx3281")
		}
	}
	rexsult, err = lang.RxFromString("65.537942676774715953400768803539").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("125946917260.87536506197191782198"))
	if err != nil {
		lang.SayString("divx3282")
	} else {
		if !(rexsult.ToString() == "5.2036162616846894920389414735878E-10") {
			lang.SayString("divx3282")
		}
	}
	rexsult, err = lang.RxFromString("8015272348677.5489394183881813700").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("949.23027111499779258284877920022"))
	if err != nil {
		lang.SayString("divx3283")
	} else {
		if !(rexsult.ToString() == "8443970438.556010797879008443011") {
			lang.SayString("divx3283")
		}
	}
	rexsult, err = lang.RxFromString("-32595333982.67068622120451804225").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("69130.052233649808750113141502465E-862"))
	if err != nil {
		lang.SayString("divx3284")
	} else {
		if !(rexsult.ToString() == "-4.7150744038935574574681609457727E+867") {
			lang.SayString("divx3284")
		}
	}
	rexsult, err = lang.RxFromString("-17544189017145.710117633021800426E-537").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("292178000.06450804618299520094843"))
	if err != nil {
		lang.SayString("divx3285")
	} else {
		if !(rexsult.ToString() == "-6.0046235559392715334668277026896E-533") {
			lang.SayString("divx3285")
		}
	}
	rexsult, err = lang.RxFromString("-506650.99395649907139204052441630").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("11.018427502031650148962067367158"))
	if err != nil {
		lang.SayString("divx3286")
	} else {
		if !(rexsult.ToString() == "-45982.150707356329027698717189104") {
			lang.SayString("divx3286")
		}
	}
	rexsult, err = lang.RxFromString("4846835159.5922372307656000769395E-241").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-84.001893040865864590122330800768"))
	if err != nil {
		lang.SayString("divx3287")
	} else {
		if !(rexsult.ToString() == "-5.7699118247660357814641813260524E-234") {
			lang.SayString("divx3287")
		}
	}
	rexsult, err = lang.RxFromString("-35.029311013822259358116556164908").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-3994308878.1994645313414534209127"))
	if err != nil {
		lang.SayString("divx3288")
	} else {
		if !(rexsult.ToString() == "8.7698052609323004543538163061774E-9") {
			lang.SayString("divx3288")
		}
	}
	rexsult, err = lang.RxFromString("7606663750.6735265233044420887018E+166").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-5491814639.4484565418284686127552E+365"))
	if err != nil {
		lang.SayString("divx3289")
	} else {
		if !(rexsult.ToString() == "-1.3850911310869487895947733090204E-199") {
			lang.SayString("divx3289")
		}
	}
	rexsult, err = lang.RxFromString("-25677.829660831741274207326697052E-163").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-94135395124193048560172705082029E-862"))
	if err != nil {
		lang.SayString("divx3290")
	} else {
		if !(rexsult.ToString() == "2.7277550199853009708493167299838E+671") {
			lang.SayString("divx3290")
		}
	}
	rexsult, err = lang.RxFromString("97271576094.456406729671729224827").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-1.5412563837540810793697955063295E+554"))
	if err != nil {
		lang.SayString("divx3291")
	} else {
		if !(rexsult.ToString() == "-6.311187231389064614447365264503E-544") {
			lang.SayString("divx3291")
		}
	}
	rexsult, err = lang.RxFromString("41139789894.401826915757391777563").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-1.8729920305671057957156159690823E-020"))
	if err != nil {
		lang.SayString("divx3292")
	} else {
		if !(rexsult.ToString() == "-2196474369511625389289506461551") {
			lang.SayString("divx3292")
		}
	}
	rexsult, err = lang.RxFromString("-83310831287241.777598696853498149").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-54799825033.797100418162985103519E-330"))
	if err != nil {
		lang.SayString("divx3293")
	} else {
		if !(rexsult.ToString() == "1.5202754978845438636605170857478E+333") {
			lang.SayString("divx3293")
		}
	}
	rexsult, err = lang.RxFromString("4506653461.4414974233678331771169").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-74955470.977653038010031457181957E-887"))
	if err != nil {
		lang.SayString("divx3294")
	} else {
		if !(rexsult.ToString() == "-6.0124409901781490054438220048629E+888") {
			lang.SayString("divx3294")
		}
	}
	rexsult, err = lang.RxFromString("23777.857951114967684767609401626").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("720760.03897144157012301385227528"))
	if err != nil {
		lang.SayString("divx3295")
	} else {
		if !(rexsult.ToString() == "0.032989978169498808275308039034795") {
			lang.SayString("divx3295")
		}
	}
	rexsult, err = lang.RxFromString("-21337853323980858055292469611948").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("6080272840.3071490445256786982100E+532"))
	if err != nil {
		lang.SayString("divx3296")
	} else {
		if !(rexsult.ToString() == "-3.5093578667274020123788514069885E-511") {
			lang.SayString("divx3296")
		}
	}
	rexsult, err = lang.RxFromString("-818409238.0423893439849743856947").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("756.39156265972753259267069158760"))
	if err != nil {
		lang.SayString("divx3297")
	} else {
		if !(rexsult.ToString() == "-1081991.4954690752676494129493403") {
			lang.SayString("divx3297")
		}
	}
	rexsult, err = lang.RxFromString("47567380384943.87013600286155046").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("65.084709374908275826942076480326"))
	if err != nil {
		lang.SayString("divx3298")
	} else {
		if !(rexsult.ToString() == "730853388480.86247690474303493972") {
			lang.SayString("divx3298")
		}
	}
	rexsult, err = lang.RxFromString("-296615544.05897984545127115290251").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-5416115.4315053536014016450973264"))
	if err != nil {
		lang.SayString("divx3299")
	} else {
		if !(rexsult.ToString() == "54.765366028496664530688259272591") {
			lang.SayString("divx3299")
		}
	}
	rexsult, err = lang.RxFromString("61391705914.046707180652185247584E+739").OpDiv(lang.RxSetWithDigit(32), lang.RxFromString("-675982087721.91856456125242297346"))
	if err != nil {
		lang.SayString("divx3300")
	} else {
		if !(rexsult.ToString() == "-9.0818539468906248593699700040068E+737") {
			lang.SayString("divx3300")
		}
	}
	rexsult, err = lang.RxFromString("042.668056830485571428877212944418").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-01364112374639.4941124016455321071"))
	if err != nil {
		lang.SayString("divx3401")
	} else {
		if !(rexsult.ToString() == "-3.12789896373176963160811150593867E-11") {
			lang.SayString("divx3401")
		}
	}
	rexsult, err = lang.RxFromString("-327.179426341653256363531809227344E+453").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("759067457.107518663444899356760793"))
	if err != nil {
		lang.SayString("divx3402")
	} else {
		if !(rexsult.ToString() == "-4.31028129684803083255704680611589E+446") {
			lang.SayString("divx3402")
		}
	}
	rexsult, err = lang.RxFromString("81721.5803096185422787702538493471").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("900099473.245809628076790757217328"))
	if err != nil {
		lang.SayString("divx3403")
	} else {
		if !(rexsult.ToString() == "0.0000907917210693679220610511319812826") {
			lang.SayString("divx3403")
		}
	}
	rexsult, err = lang.RxFromString("3991.56734635183403814636354392163E-807").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("72.3239822255871305731314565069132"))
	if err != nil {
		lang.SayString("divx3404")
	} else {
		if !(rexsult.ToString() == "5.5190093569539066498459824811529E-806") {
			lang.SayString("divx3404")
		}
	}
	rexsult, err = lang.RxFromString("-66.3393308595957726456416979163306").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("5.08486573050459213864684589662559"))
	if err != nil {
		lang.SayString("divx3405")
	} else {
		if !(rexsult.ToString() == "-13.046427256007927669474992491585") {
			lang.SayString("divx3405")
		}
	}
	rexsult, err = lang.RxFromString("-393606.873703567753255097095208112E+111").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-2124339550.86891093200758095660557"))
	if err != nil {
		lang.SayString("divx3406")
	} else {
		if !(rexsult.ToString() == "1.85284350396137075010428736564737E+107") {
			lang.SayString("divx3406")
		}
	}
	rexsult, err = lang.RxFromString("-019133598.609812524622150421584346").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-858439846.628367734642622922030051"))
	if err != nil {
		lang.SayString("divx3407")
	} else {
		if !(rexsult.ToString() == "0.022288805307631256579746065031107") {
			lang.SayString("divx3407")
		}
	}
	rexsult, err = lang.RxFromString("465.351982159046525762715549761814").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("240444.975944666924657629172844780"))
	if err != nil {
		lang.SayString("divx3408")
	} else {
		if !(rexsult.ToString() == "0.00193537827243326122782974132829095") {
			lang.SayString("divx3408")
		}
	}
	rexsult, err = lang.RxFromString("28066955004783.1076824222873384828").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("571699969.220753535758504907561016E-718"))
	if err != nil {
		lang.SayString("divx3409")
	} else {
		if !(rexsult.ToString() == "4.90938543219432390013656968123815E+722") {
			lang.SayString("divx3409")
		}
	}
	rexsult, err = lang.RxFromString("28275236927392.4960902824105246047").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("28212038.4825243127096613158419270E+422"))
	if err != nil {
		lang.SayString("divx3410")
	} else {
		if !(rexsult.ToString() == "1.00224012330435927467559203688861E-416") {
			lang.SayString("divx3410")
		}
	}
	rexsult, err = lang.RxFromString("11791.8644211874630234271801789996").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-8.45457275930363860982261343159741"))
	if err != nil {
		lang.SayString("divx3411")
	} else {
		if !(rexsult.ToString() == "-1394.73214754836418731335761858151") {
			lang.SayString("divx3411")
		}
	}
	rexsult, err = lang.RxFromString("44.7085340739581668975502342787578").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-9337.05408133023920640485556647937"))
	if err != nil {
		lang.SayString("divx3412")
	} else {
		if !(rexsult.ToString() == "-0.00478829121953512281527242631775613") {
			lang.SayString("divx3412")
		}
	}
	rexsult, err = lang.RxFromString("93354527428804.5458053295581965867E+576").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-856525909852.318790321300941615314"))
	if err != nil {
		lang.SayString("divx3413")
	} else {
		if !(rexsult.ToString() == "-1.08992064752484400353231056271614E+578") {
			lang.SayString("divx3413")
		}
	}
	rexsult, err = lang.RxFromString("-367399415798804503177950040443482").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-54845683.9691776397285506712812754"))
	if err != nil {
		lang.SayString("divx3414")
	} else {
		if !(rexsult.ToString() == "6698784465980529140072174.30474769") {
			lang.SayString("divx3414")
		}
	}
	rexsult, err = lang.RxFromString("-2.87155919781024108503670175443740").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("89529730130.6427881332776797193807"))
	if err != nil {
		lang.SayString("divx3415")
	} else {
		if !(rexsult.ToString() == "-3.2073806026445401317483592875443E-11") {
			lang.SayString("divx3415")
		}
	}
	rexsult, err = lang.RxFromString("-010.693934338179479652178057279204E+188").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("26484.8887731973153745666514260684"))
	if err != nil {
		lang.SayString("divx3416")
	} else {
		if !(rexsult.ToString() == "-4.03774938598259547575707503087638E+184") {
			lang.SayString("divx3416")
		}
	}
	rexsult, err = lang.RxFromString("611655569568.832698912762075889186").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("010182743219.475839030505966016982"))
	if err != nil {
		lang.SayString("divx3417")
	} else {
		if !(rexsult.ToString() == "60.0678575886074367081836436812959") {
			lang.SayString("divx3417")
		}
	}
	rexsult, err = lang.RxFromString("3457947.39062863674882672518304442").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-01.9995218868908849056866549811425"))
	if err != nil {
		lang.SayString("divx3418")
	} else {
		if !(rexsult.ToString() == "-1729387.11663991852426428263230475") {
			lang.SayString("divx3418")
		}
	}
	rexsult, err = lang.RxFromString("-53308666960535.7393391289364591513").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-6527.00547629475578694521436764596E-442"))
	if err != nil {
		lang.SayString("divx3419")
	} else {
		if !(rexsult.ToString() == "8.16740037282731870883136714441204E+451") {
			lang.SayString("divx3419")
		}
	}
	rexsult, err = lang.RxFromString("-5568057.17870139549478277980540034").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-407906443.141342175740471849723638"))
	if err != nil {
		lang.SayString("divx3420")
	} else {
		if !(rexsult.ToString() == "0.0136503290701197094953429018013146") {
			lang.SayString("divx3420")
		}
	}
	rexsult, err = lang.RxFromString("9804385273.49533524416415189990857").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("84.1433929743544659553964804646569"))
	if err != nil {
		lang.SayString("divx3421")
	} else {
		if !(rexsult.ToString() == "116519965.821719977402398190558439") {
			lang.SayString("divx3421")
		}
	}
	rexsult, err = lang.RxFromString("-5234910986592.18801727046580014273E-547").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-5874220715892.91440069210512515154"))
	if err != nil {
		lang.SayString("divx3422")
	} else {
		if !(rexsult.ToString() == "8.91166886601477021757439826903776E-548") {
			lang.SayString("divx3422")
		}
	}
	rexsult, err = lang.RxFromString("698416560151955285929747633786867E-495").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("51754681.6784872628933218985216916E-266"))
	if err != nil {
		lang.SayString("divx3423")
	} else {
		if !(rexsult.ToString() == "1.34947513442491971488363250398908E-204") {
			lang.SayString("divx3423")
		}
	}
	rexsult, err = lang.RxFromString("107635.497735316515080720330536027").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-3972075.83989512668362609609006425E-605"))
	if err != nil {
		lang.SayString("divx3424")
	} else {
		if !(rexsult.ToString() == "-2.70980469844599888443309571235597E+603") {
			lang.SayString("divx3424")
		}
	}
	rexsult, err = lang.RxFromString("-32174291345686.5371446616670961807").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("79518863759385.5925052747867099091E+408"))
	if err != nil {
		lang.SayString("divx3425")
	} else {
		if !(rexsult.ToString() == "-4.04612060894658007715621807881076E-409") {
			lang.SayString("divx3425")
		}
	}
	rexsult, err = lang.RxFromString("-8151730494.53190523620899410544099E+688").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-93173.0631474527142307644239919480E+900"))
	if err != nil {
		lang.SayString("divx3426")
	} else {
		if !(rexsult.ToString() == "8.74902060655796717043678441884283E-208") {
			lang.SayString("divx3426")
		}
	}
	rexsult, err = lang.RxFromString("1.33649801345976199708341799505220").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-56623.0530039528969825480755159562E+459"))
	if err != nil {
		lang.SayString("divx3427")
	} else {
		if !(rexsult.ToString() == "-2.36034255052700900395787131334608E-464") {
			lang.SayString("divx3427")
		}
	}
	rexsult, err = lang.RxFromString("67762238162788.6551061476018185196").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-6140.75837959248100352788853809376E-822"))
	if err != nil {
		lang.SayString("divx3428")
	} else {
		if !(rexsult.ToString() == "-1.10348321777294157014941951870409E+832") {
			lang.SayString("divx3428")
		}
	}
	rexsult, err = lang.RxFromString("4286562.76568866751577306056498271").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("6286.77291578497580015557979349893E+820"))
	if err != nil {
		lang.SayString("divx3429")
	} else {
		if !(rexsult.ToString() == "6.81838333133660025740681459349372E-818") {
			lang.SayString("divx3429")
		}
	}
	rexsult, err = lang.RxFromString("-765782.827432642697305644096365566").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("67.1634368459576834692758114618652"))
	if err != nil {
		lang.SayString("divx3430")
	} else {
		if !(rexsult.ToString() == "-11401.7814363639478774761697650867") {
			lang.SayString("divx3430")
		}
	}
	rexsult, err = lang.RxFromString("46.2835931916106252756465724211276").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("59.2989237834093118332826617957791"))
	if err != nil {
		lang.SayString("divx3431")
	} else {
		if !(rexsult.ToString() == "0.780513207299722975882416995140701") {
			lang.SayString("divx3431")
		}
	}
	rexsult, err = lang.RxFromString("-3029555.82298840234029474459694644").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("857535844655004737373089601128532"))
	if err != nil {
		lang.SayString("divx3432")
	} else {
		if !(rexsult.ToString() == "-3.53286202771759704502126811323937E-27") {
			lang.SayString("divx3432")
		}
	}
	rexsult, err = lang.RxFromString("-0138466789523.10694176543700501945E-948").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("481026979918882487383654367924619"))
	if err != nil {
		lang.SayString("divx3433")
	} else {
		if !(rexsult.ToString() == "-2.87856597038397207797777811199804E-970") {
			lang.SayString("divx3433")
		}
	}
	rexsult, err = lang.RxFromString("-9593566466.96690575714244442109870").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-87632034347.4845477961976776833770E+769"))
	if err != nil {
		lang.SayString("divx3434")
	} else {
		if !(rexsult.ToString() == "1.09475564939253134070730299863765E-770") {
			lang.SayString("divx3434")
		}
	}
	rexsult, err = lang.RxFromString("-3189.30765477670526823106100241863E-898").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("565688889.355241946154894311253202E-466"))
	if err != nil {
		lang.SayString("divx3435")
	} else {
		if !(rexsult.ToString() == "-5.63791814686655486612569970629128E-438") {
			lang.SayString("divx3435")
		}
	}
	rexsult, err = lang.RxFromString("-17084552395.6714834680088150543965").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-631925802672.685034379197328370812E+527"))
	if err != nil {
		lang.SayString("divx3436")
	} else {
		if !(rexsult.ToString() == "2.70356936263934622050341328519534E-529") {
			lang.SayString("divx3436")
		}
	}
	rexsult, err = lang.RxFromString("034956830.349823306815911887469760").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-61600816.0672274126966042956781665E-667"))
	if err != nil {
		lang.SayString("divx3437")
	} else {
		if !(rexsult.ToString() == "-5.67473494371787737607169979602343E+666") {
			lang.SayString("divx3437")
		}
	}
	rexsult, err = lang.RxFromString("-763.440067781256632695791981893608").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("19.9263811350611007833220620745413"))
	if err != nil {
		lang.SayString("divx3438")
	} else {
		if !(rexsult.ToString() == "-38.3130314835722766807703585435688") {
			lang.SayString("divx3438")
		}
	}
	rexsult, err = lang.RxFromString("-510472027868440667684575147556654E+789").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("834872378550801889983927148587909"))
	if err != nil {
		lang.SayString("divx3439")
	} else {
		if !(rexsult.ToString() == "-6.11437198047603754107526874071737E+788") {
			lang.SayString("divx3439")
		}
	}
	rexsult, err = lang.RxFromString("070304761.560517086676993503034828E-094").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-17773.7446959771077104057845273992E-761"))
	if err != nil {
		lang.SayString("divx3440")
	} else {
		if !(rexsult.ToString() == "-3.95554019499502537743883483402608E+670") {
			lang.SayString("divx3440")
		}
	}
	rexsult, err = lang.RxFromString("-0970725697662.27605454336231195463").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-4541.41897546697187157913886433474"))
	if err != nil {
		lang.SayString("divx3441")
	} else {
		if !(rexsult.ToString() == "213749425.654447811698884007553614") {
			lang.SayString("divx3441")
		}
	}
	rexsult, err = lang.RxFromString("-808178238631844268316111259558675").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-598400.265108644514211244980426520"))
	if err != nil {
		lang.SayString("divx3442")
	} else {
		if !(rexsult.ToString() == "1350564640015847635178945884.97836") {
			lang.SayString("divx3442")
		}
	}
	rexsult, err = lang.RxFromString("-9.90826595069053564311371766315200").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-031.625916781307847864872329806646"))
	if err != nil {
		lang.SayString("divx3443")
	} else {
		if !(rexsult.ToString() == "0.313295770023233218639213140599856") {
			lang.SayString("divx3443")
		}
	}
	rexsult, err = lang.RxFromString("-196925.469891897719160698483752907").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-41268.9975444533794067723958739778"))
	if err != nil {
		lang.SayString("divx3444")
	} else {
		if !(rexsult.ToString() == "4.77175317088274715226553516820589") {
			lang.SayString("divx3444")
		}
	}
	rexsult, err = lang.RxFromString("421071135212152225162086005824310").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("1335320330.08964354845796510145246E-604"))
	if err != nil {
		lang.SayString("divx3445")
	} else {
		if !(rexsult.ToString() == "3.15333426537349744281860005497304E+627") {
			lang.SayString("divx3445")
		}
	}
	rexsult, err = lang.RxFromString("1249441.46421514282301182772247227").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-0289848.71208912281976374705180836E-676"))
	if err != nil {
		lang.SayString("divx3446")
	} else {
		if !(rexsult.ToString() == "-4.31066764178328992440635387255816E+676") {
			lang.SayString("divx3446")
		}
	}
	rexsult, err = lang.RxFromString("74815000.4716875558358937279052903").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-690425401708167622194241915195001E+891"))
	if err != nil {
		lang.SayString("divx3447")
	} else {
		if !(rexsult.ToString() == "-1.08360729901578455109968388309079E-916") {
			lang.SayString("divx3447")
		}
	}
	rexsult, err = lang.RxFromString("-1683993.51210241555668790556759021").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-72394384927344.8402585228267493374"))
	if err != nil {
		lang.SayString("divx3448")
	} else {
		if !(rexsult.ToString() == "2.32613829621244113284301004158794E-8") {
			lang.SayString("divx3448")
		}
	}
	rexsult, err = lang.RxFromString("-763.148530974741766171756970448158").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("517370.808956957601473642272664647"))
	if err != nil {
		lang.SayString("divx3449")
	} else {
		if !(rexsult.ToString() == "-0.00147505139014951946381155525173867") {
			lang.SayString("divx3449")
		}
	}
	rexsult, err = lang.RxFromString("-77.5841338812312523460591226178754").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-927540422.641025050968830154578151E+524"))
	if err != nil {
		lang.SayString("divx3450")
	} else {
		if !(rexsult.ToString() == "8.36450164191471769978415758342237E-532") {
			lang.SayString("divx3450")
		}
	}
	rexsult, err = lang.RxFromString("5176295309.89943746236102209837813").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-129733.103628797477167908698565465"))
	if err != nil {
		lang.SayString("divx3451")
	} else {
		if !(rexsult.ToString() == "-39899.5720067736855444089432524094") {
			lang.SayString("divx3451")
		}
	}
	rexsult, err = lang.RxFromString("4471634841166.90197229286530307326E+172").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("31511104397.8671727003201890825879E-955"))
	if err != nil {
		lang.SayString("divx3452")
	} else {
		if !(rexsult.ToString() == "1.41906636616314987705536737025932E+1129") {
			lang.SayString("divx3452")
		}
	}
	rexsult, err = lang.RxFromString("-8189130.15945231049670285726774317").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("2.57912402871404325831670923864936E-366"))
	if err != nil {
		lang.SayString("divx3453")
	} else {
		if !(rexsult.ToString() == "-3.17515949922556778343526099830093E+372") {
			lang.SayString("divx3453")
		}
	}
	rexsult, err = lang.RxFromString("-254346232981353293392888785643245").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-764.416902486152093758287929678445"))
	if err != nil {
		lang.SayString("divx3454")
	} else {
		if !(rexsult.ToString() == "332732350833857889204406356900.665") {
			lang.SayString("divx3454")
		}
	}
	rexsult, err = lang.RxFromString("-2875.36745499544177964804277329726").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-13187.8492045054802205691248267638"))
	if err != nil {
		lang.SayString("divx3455")
	} else {
		if !(rexsult.ToString() == "0.218031569091122520391599541575615") {
			lang.SayString("divx3455")
		}
	}
	rexsult, err = lang.RxFromString("-7.26927570984219944693602140223103").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("0160883021541.32275286769110003971E-438"))
	if err != nil {
		lang.SayString("divx3456")
	} else {
		if !(rexsult.ToString() == "-4.51836100553039917574557235275173E+427") {
			lang.SayString("divx3456")
		}
	}
	rexsult, err = lang.RxFromString("-28567151.6868762752241056540028515E+498").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-4470.15455137546427645290210989675"))
	if err != nil {
		lang.SayString("divx3457")
	} else {
		if !(rexsult.ToString() == "6.39064071690455919792707589054106E+501") {
			lang.SayString("divx3457")
		}
	}
	rexsult, err = lang.RxFromString("7191753.79974136447257468282073718").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("81.3878426176038487948375777384429"))
	if err != nil {
		lang.SayString("divx3458")
	} else {
		if !(rexsult.ToString() == "88363.98125861881867339355698741") {
			lang.SayString("divx3458")
		}
	}
	rexsult, err = lang.RxFromString("502975804.069864536104621700404770").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("684.790028432074527960269515227243"))
	if err != nil {
		lang.SayString("divx3459")
	} else {
		if !(rexsult.ToString() == "734496.390406706323899801641278933") {
			lang.SayString("divx3459")
		}
	}
	rexsult, err = lang.RxFromString("1505368.42063731861590460453659570").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-465242.678439951462767630022819105"))
	if err != nil {
		lang.SayString("divx3460")
	} else {
		if !(rexsult.ToString() == "-3.23566278503319947059213001405065") {
			lang.SayString("divx3460")
		}
	}
	rexsult, err = lang.RxFromString("69225023.2850142784063416137144829").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("8584050.06648191798834819995325693"))
	if err != nil {
		lang.SayString("divx3461")
	} else {
		if !(rexsult.ToString() == "8.06437785764050431295652411163382") {
			lang.SayString("divx3461")
		}
	}
	rexsult, err = lang.RxFromString("-55669501853.7751006841940471339310E+614").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("061400538.186044693233816566977189"))
	if err != nil {
		lang.SayString("divx3462")
	} else {
		if !(rexsult.ToString() == "-9.06661464189378059067792554300676E+616") {
			lang.SayString("divx3462")
		}
	}
	rexsult, err = lang.RxFromString("-527566.521273992424649346837337602E-408").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-834662.599983953345718523807123972"))
	if err != nil {
		lang.SayString("divx3463")
	} else {
		if !(rexsult.ToString() == "6.32071595497552015656909892339226E-409") {
			lang.SayString("divx3463")
		}
	}
	rexsult, err = lang.RxFromString("69065510.0459653699418083455335366").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("694848643848212520086960886818157E-853"))
	if err != nil {
		lang.SayString("divx3464")
	} else {
		if !(rexsult.ToString() == "9.93964810285396646889830919492683E+827") {
			lang.SayString("divx3464")
		}
	}
	rexsult, err = lang.RxFromString("-2921982.82411285505229122890041841").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-72994.6523840298017471962569778803E-763"))
	if err != nil {
		lang.SayString("divx3465")
	} else {
		if !(rexsult.ToString() == "4.00300943792444663467732029501716E+764") {
			lang.SayString("divx3465")
		}
	}
	rexsult, err = lang.RxFromString("4.51117459466491451401835756593793").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("3873385.19981811640063144338087230"))
	if err != nil {
		lang.SayString("divx3466")
	} else {
		if !(rexsult.ToString() == "0.00000116465942888322776753062580106351") {
			lang.SayString("divx3466")
		}
	}
	rexsult, err = lang.RxFromString("49553763474698.8118661758811091120").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("36.1713861293896216593840817950781E+410"))
	if err != nil {
		lang.SayString("divx3467")
	} else {
		if !(rexsult.ToString() == "1.36997137177543416190811827685231E-398") {
			lang.SayString("divx3467")
		}
	}
	rexsult, err = lang.RxFromString("755985583344.379951506071499170749E+956").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("746921095569971477373643487285780"))
	if err != nil {
		lang.SayString("divx3468")
	} else {
		if !(rexsult.ToString() == "1.01213580367212873025671916758669E+935") {
			lang.SayString("divx3468")
		}
	}
	rexsult, err = lang.RxFromString("-20101668541.7472260497594230105456").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-395562148.345003931161532101109964"))
	if err != nil {
		lang.SayString("divx3469")
	} else {
		if !(rexsult.ToString() == "50.8179779735012053661447873666816") {
			lang.SayString("divx3469")
		}
	}
	rexsult, err = lang.RxFromString("5583903.18072100716072011264668631").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("460868914694.088387067451312500723"))
	if err != nil {
		lang.SayString("divx3470")
	} else {
		if !(rexsult.ToString() == "0.0000121160334374633240168068405467307") {
			lang.SayString("divx3470")
		}
	}
	rexsult, err = lang.RxFromString("-955638397975240685017992436116257E+020").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-508580.148958769104511751975720470E+662"))
	if err != nil {
		lang.SayString("divx3471")
	} else {
		if !(rexsult.ToString() == "1.87903204624039476408191264564568E-615") {
			lang.SayString("divx3471")
		}
	}
	rexsult, err = lang.RxFromString("-136243796098020983814115584402407E+796").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("6589108083.99750311651581032447390"))
	if err != nil {
		lang.SayString("divx3472")
	} else {
		if !(rexsult.ToString() == "-2.0677122663825560063493937136592E+818") {
			lang.SayString("divx3472")
		}
	}
	rexsult, err = lang.RxFromString("-808498.482718304598213092937543934E+521").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("48005.1465097914355096301483788905"))
	if err != nil {
		lang.SayString("divx3473")
	} else {
		if !(rexsult.ToString() == "-1.68419126177106468565397017107736E+522") {
			lang.SayString("divx3473")
		}
	}
	rexsult, err = lang.RxFromString("-812.266340554281305985524813074211E+396").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-3195.63111559114001594257448970812E+986"))
	if err != nil {
		lang.SayString("divx3474")
	} else {
		if !(rexsult.ToString() == "2.5418025772477972144848478105604E-591") {
			lang.SayString("divx3474")
		}
	}
	rexsult, err = lang.RxFromString("-929889.720905183397678866648217793E+134").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-280300190774.057595671079264841349"))
	if err != nil {
		lang.SayString("divx3475")
	} else {
		if !(rexsult.ToString() == "3.31747801646964399331556971055197E+128") {
			lang.SayString("divx3475")
		}
	}
	rexsult, err = lang.RxFromString("83946.0157801953636255078996185540").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("492670373.235391665758701500314473"))
	if err != nil {
		lang.SayString("divx3476")
	} else {
		if !(rexsult.ToString() == "0.000170389819117633485695890041185782") {
			lang.SayString("divx3476")
		}
	}
	rexsult, err = lang.RxFromString("7812758113817.99135851825223122508").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("3037492.36716301443309571918002123E-157"))
	if err != nil {
		lang.SayString("divx3477")
	} else {
		if !(rexsult.ToString() == "2.57210790001590171809512871857381E+163") {
			lang.SayString("divx3477")
		}
	}
	rexsult, err = lang.RxFromString("489191747.148674326757767356626016").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("01136942.1182277580093027768490545"))
	if err != nil {
		lang.SayString("divx3478")
	} else {
		if !(rexsult.ToString() == "430.269702657525223124148258641358") {
			lang.SayString("divx3478")
		}
	}
	rexsult, err = lang.RxFromString("-599369540.373174482335865567937853E+289").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-38288383205.6749439588707955585209"))
	if err != nil {
		lang.SayString("divx3479")
	} else {
		if !(rexsult.ToString() == "1.56540833065089684132688143737586E+287") {
			lang.SayString("divx3479")
		}
	}
	rexsult, err = lang.RxFromString("-3376883870.85961692148022521960139").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-65247489449.7334589731171980408284"))
	if err != nil {
		lang.SayString("divx3480")
	} else {
		if !(rexsult.ToString() == "0.0517550008335747813596332404664731") {
			lang.SayString("divx3480")
		}
	}
	rexsult, err = lang.RxFromString("58.6776780370008364590621772421025").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("01.5925518866529044494309229975160"))
	if err != nil {
		lang.SayString("divx3481")
	} else {
		if !(rexsult.ToString() == "36.8450651616286048437498576613622") {
			lang.SayString("divx3481")
		}
	}
	rexsult, err = lang.RxFromString("4099616339.96249499552808575717579").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("290.795187361072489816791525139895"))
	if err != nil {
		lang.SayString("divx3482")
	} else {
		if !(rexsult.ToString() == "14097951.1289920788134209002390834") {
			lang.SayString("divx3482")
		}
	}
	rexsult, err = lang.RxFromString("85870777.2282833141709970713739108").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-2140392861153.69401346398478113715"))
	if err != nil {
		lang.SayString("divx3483")
	} else {
		if !(rexsult.ToString() == "-0.0000401191663393971853092748263233128") {
			lang.SayString("divx3483")
		}
	}
	rexsult, err = lang.RxFromString("20900.9693761555165742010339929779").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-38.7546147649523793463260940289585"))
	if err != nil {
		lang.SayString("divx3484")
	} else {
		if !(rexsult.ToString() == "-539.315627388386430188627412639767") {
			lang.SayString("divx3484")
		}
	}
	rexsult, err = lang.RxFromString("448.827596155587910947511170319456").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("379130153.382794042652974596286062"))
	if err != nil {
		lang.SayString("divx3485")
	} else {
		if !(rexsult.ToString() == "0.00000118383513458615061394140895596979") {
			lang.SayString("divx3485")
		}
	}
	rexsult, err = lang.RxFromString("98.4134807921002817357000140482039").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("3404725543.77032945444654351546779"))
	if err != nil {
		lang.SayString("divx3486")
	} else {
		if !(rexsult.ToString() == "2.89049673833970863420201979291523E-8") {
			lang.SayString("divx3486")
		}
	}
	rexsult, err = lang.RxFromString("545746433.649359734136476718176330E-787").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-5149957099709.12830072802043560650E-437"))
	if err != nil {
		lang.SayString("divx3487")
	} else {
		if !(rexsult.ToString() == "-1.05971064046375011086850722752614E-354") {
			lang.SayString("divx3487")
		}
	}
	rexsult, err = lang.RxFromString("741304513547.273820525801608231737").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("0396.22823128272584928019323186355E-830"))
	if err != nil {
		lang.SayString("divx3488")
	} else {
		if !(rexsult.ToString() == "1.87090281565101612623398174727653E+839") {
			lang.SayString("divx3488")
		}
	}
	rexsult, err = lang.RxFromString("-706.145005094292315613907254240553").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("4739.82486195739758308735946332234"))
	if err != nil {
		lang.SayString("divx3489")
	} else {
		if !(rexsult.ToString() == "-0.148981244172527671907534117771626") {
			lang.SayString("divx3489")
		}
	}
	rexsult, err = lang.RxFromString("-769925786.823099083228795187975893").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-31201.9980469760239870067820594790"))
	if err != nil {
		lang.SayString("divx3490")
	} else {
		if !(rexsult.ToString() == "24675.5283319978698932292028650803") {
			lang.SayString("divx3490")
		}
	}
	rexsult, err = lang.RxFromString("84438610546049.7256507419289692857E+906").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("052604240766736461898844111790311"))
	if err != nil {
		lang.SayString("divx3491")
	} else {
		if !(rexsult.ToString() == "1.60516736512701978695559003341922E+888") {
			lang.SayString("divx3491")
		}
	}
	rexsult, err = lang.RxFromString("549760.911304725795164589619286514").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("165.160089615604924207754883953484"))
	if err != nil {
		lang.SayString("divx3492")
	} else {
		if !(rexsult.ToString() == "3328.65471667062107598395714348089") {
			lang.SayString("divx3492")
		}
	}
	rexsult, err = lang.RxFromString("3650514.18649737956855828939662794").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("08086721.4036886948248274834735629"))
	if err != nil {
		lang.SayString("divx3493")
	} else {
		if !(rexsult.ToString() == "0.451420792712387250865423208234291") {
			lang.SayString("divx3493")
		}
	}
	rexsult, err = lang.RxFromString("55067723881950.1346958179604099594").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-8.90481481687182931431054785192083"))
	if err != nil {
		lang.SayString("divx3494")
	} else {
		if !(rexsult.ToString() == "-6184039198391.19853088419484117054") {
			lang.SayString("divx3494")
		}
	}
	rexsult, err = lang.RxFromString("868251123.413992653362860637541060E+019").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("5579665045.37858308541154858567656E+131"))
	if err != nil {
		lang.SayString("divx3495")
	} else {
		if !(rexsult.ToString() == "1.55609900657590706155251902725027E-113") {
			lang.SayString("divx3495")
		}
	}
	rexsult, err = lang.RxFromString("-646.464431574014407536004990059069").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-798.679560020414523841321724649594E-037"))
	if err != nil {
		lang.SayString("divx3496")
	} else {
		if !(rexsult.ToString() == "8.09416521887063886613527228353543E+36") {
			lang.SayString("divx3496")
		}
	}
	rexsult, err = lang.RxFromString("354.546679975219753598558273421556").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-7039.46386812239015144581761752927E-448"))
	if err != nil {
		lang.SayString("divx3497")
	} else {
		if !(rexsult.ToString() == "-5.03655799102477192579414523352028E+446") {
			lang.SayString("divx3497")
		}
	}
	rexsult, err = lang.RxFromString("91936087917435.5974889495278215874").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("-67080823344.8903392584327136082486E-757"))
	if err != nil {
		lang.SayString("divx3498")
	} else {
		if !(rexsult.ToString() == "-1.37052712434303366569304688993783E+760") {
			lang.SayString("divx3498")
		}
	}
	rexsult, err = lang.RxFromString("-07345.6422518528556136521417259811E-600").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("41188325.7041362608934957584583381E-763"))
	if err != nil {
		lang.SayString("divx3499")
	} else {
		if !(rexsult.ToString() == "-1.78342822299163842247184303878022E+159") {
			lang.SayString("divx3499")
		}
	}
	rexsult, err = lang.RxFromString("-253280724.939458021588167965038184").OpDiv(lang.RxSetWithDigit(33), lang.RxFromString("616988.426425908872398170896375634E+396"))
	if err != nil {
		lang.SayString("divx3500")
	} else {
		if !(rexsult.ToString() == "-4.1051130635733775335165551186617E-394") {
			lang.SayString("divx3500")
		}
	}
	rexsult, err = lang.RxFromString("905.67402").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-202896611.E-780472620"))
	if err != nil {
		lang.SayString("xdiv001")
	} else {
		if !(rexsult.ToString() == "-4.46372177E+780472614") {
			lang.SayString("xdiv001")
		}
	}
	rexsult, err = lang.RxFromString("3915134.7").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-597164907."))
	if err != nil {
		lang.SayString("xdiv002")
	} else {
		if !(rexsult.ToString() == "-0.00655620358") {
			lang.SayString("xdiv002")
		}
	}
	rexsult, err = lang.RxFromString("309759261").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("62663.487"))
	if err != nil {
		lang.SayString("xdiv003")
	} else {
		if !(rexsult.ToString() == "4943.21775") {
			lang.SayString("xdiv003")
		}
	}
	rexsult, err = lang.RxFromString("3.93591888E-236595626").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7242375.00"))
	if err != nil {
		lang.SayString("xdiv004")
	} else {
		if !(rexsult.ToString() == "5.4345693E-236595633") {
			lang.SayString("xdiv004")
		}
	}
	rexsult, err = lang.RxFromString("323902.714").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-608669.607E-657060568"))
	if err != nil {
		lang.SayString("xdiv005")
	} else {
		if !(rexsult.ToString() == "-5.32148657E+657060567") {
			lang.SayString("xdiv005")
		}
	}
	rexsult, err = lang.RxFromString("5.11970092").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8807.22036"))
	if err != nil {
		lang.SayString("xdiv006")
	} else {
		if !(rexsult.ToString() == "-0.000581307236") {
			lang.SayString("xdiv006")
		}
	}
	rexsult, err = lang.RxFromString("-7.99874516").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4561.83758"))
	if err != nil {
		lang.SayString("xdiv007")
	} else {
		if !(rexsult.ToString() == "-0.0017534042") {
			lang.SayString("xdiv007")
		}
	}
	rexsult, err = lang.RxFromString("297802878").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-927206.324"))
	if err != nil {
		lang.SayString("xdiv008")
	} else {
		if !(rexsult.ToString() == "-321.182967") {
			lang.SayString("xdiv008")
		}
	}
	rexsult, err = lang.RxFromString("-766.651824").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("31300.3619"))
	if err != nil {
		lang.SayString("xdiv009")
	} else {
		if !(rexsult.ToString() == "-0.0244933853") {
			lang.SayString("xdiv009")
		}
	}
	rexsult, err = lang.RxFromString("-56746.8689E+934981942").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("471002521."))
	if err != nil {
		lang.SayString("xdiv010")
	} else {
		if !(rexsult.ToString() == "-1.2048103E+934981938") {
			lang.SayString("xdiv010")
		}
	}
	rexsult, err = lang.RxFromString("456417160").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-41346.1024"))
	if err != nil {
		lang.SayString("xdiv011")
	} else {
		if !(rexsult.ToString() == "-11038.9404") {
			lang.SayString("xdiv011")
		}
	}
	rexsult, err = lang.RxFromString("102895.722").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2.62214826"))
	if err != nil {
		lang.SayString("xdiv012")
	} else {
		if !(rexsult.ToString() == "-39241.0008") {
			lang.SayString("xdiv012")
		}
	}
	rexsult, err = lang.RxFromString("80223.3897").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("73921.0383E-467772675"))
	if err != nil {
		lang.SayString("xdiv014")
	} else {
		if !(rexsult.ToString() == "1.08525789E+467772675") {
			lang.SayString("xdiv014")
		}
	}
	rexsult, err = lang.RxFromString("-654645.954").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-9.12535752"))
	if err != nil {
		lang.SayString("xdiv015")
	} else {
		if !(rexsult.ToString() == "71739.2116") {
			lang.SayString("xdiv015")
		}
	}
	rexsult, err = lang.RxFromString("63.1917772E-706014634").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7.56253257E-138579234"))
	if err != nil {
		lang.SayString("xdiv016")
	} else {
		if !(rexsult.ToString() == "-8.35590149E-567435400") {
			lang.SayString("xdiv016")
		}
	}
	rexsult, err = lang.RxFromString("-39674.7190").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2490607.78"))
	if err != nil {
		lang.SayString("xdiv017")
	} else {
		if !(rexsult.ToString() == "-0.0159297338") {
			lang.SayString("xdiv017")
		}
	}
	rexsult, err = lang.RxFromString("-3364.59737E-600363681").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("896487.451"))
	if err != nil {
		lang.SayString("xdiv018")
	} else {
		if !(rexsult.ToString() == "-3.7530892E-600363684") {
			lang.SayString("xdiv018")
		}
	}
	rexsult, err = lang.RxFromString("-64138.0578").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("31759011.3E+697488342"))
	if err != nil {
		lang.SayString("xdiv019")
	} else {
		if !(rexsult.ToString() == "-2.01952313E-697488345") {
			lang.SayString("xdiv019")
		}
	}
	rexsult, err = lang.RxFromString("61399.8527").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-64344484.5"))
	if err != nil {
		lang.SayString("xdiv020")
	} else {
		if !(rexsult.ToString() == "-0.000954236454") {
			lang.SayString("xdiv020")
		}
	}
	rexsult, err = lang.RxFromString("-722960.204").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-26154599.8"))
	if err != nil {
		lang.SayString("xdiv021")
	} else {
		if !(rexsult.ToString() == "0.0276417995") {
			lang.SayString("xdiv021")
		}
	}
	rexsult, err = lang.RxFromString("9.47109959E+230565093").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("73354723.2"))
	if err != nil {
		lang.SayString("xdiv022")
	} else {
		if !(rexsult.ToString() == "1.29113698E+230565086") {
			lang.SayString("xdiv022")
		}
	}
	rexsult, err = lang.RxFromString("43.7456245").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("547441956."))
	if err != nil {
		lang.SayString("xdiv023")
	} else {
		if !(rexsult.ToString() == "7.99091557E-8") {
			lang.SayString("xdiv023")
		}
	}
	rexsult, err = lang.RxFromString("-73150542E-242017390").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8.15869954"))
	if err != nil {
		lang.SayString("xdiv024")
	} else {
		if !(rexsult.ToString() == "8.96595611E-242017384") {
			lang.SayString("xdiv024")
		}
	}
	rexsult, err = lang.RxFromString("2015.62109E+299897596").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-11788916.1"))
	if err != nil {
		lang.SayString("xdiv025")
	} else {
		if !(rexsult.ToString() == "-1.70975947E+299897592") {
			lang.SayString("xdiv025")
		}
	}
	rexsult, err = lang.RxFromString("29.498114").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-26486451"))
	if err != nil {
		lang.SayString("xdiv026")
	} else {
		if !(rexsult.ToString() == "-0.0000011137058") {
			lang.SayString("xdiv026")
		}
	}
	rexsult, err = lang.RxFromString("244375043.E+130840878").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-9.44522029"))
	if err != nil {
		lang.SayString("xdiv027")
	} else {
		if !(rexsult.ToString() == "-2.58728791E+130840885") {
			lang.SayString("xdiv027")
		}
	}
	rexsult, err = lang.RxFromString("-349388.759").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-196215.776"))
	if err != nil {
		lang.SayString("xdiv028")
	} else {
		if !(rexsult.ToString() == "1.78063541") {
			lang.SayString("xdiv028")
		}
	}
	rexsult, err = lang.RxFromString("-70905112.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-91353968.8"))
	if err != nil {
		lang.SayString("xdiv029")
	} else {
		if !(rexsult.ToString() == "0.776157986") {
			lang.SayString("xdiv029")
		}
	}
	rexsult, err = lang.RxFromString("-225094.28").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-88.7723542"))
	if err != nil {
		lang.SayString("xdiv030")
	} else {
		if !(rexsult.ToString() == "2535.63491") {
			lang.SayString("xdiv030")
		}
	}
	rexsult, err = lang.RxFromString("50.4442340").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("82.7952169E+880120759"))
	if err != nil {
		lang.SayString("xdiv031")
	} else {
		if !(rexsult.ToString() == "6.09265075E-880120760") {
			lang.SayString("xdiv031")
		}
	}
	rexsult, err = lang.RxFromString("-32311.9037").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8.36379449"))
	if err != nil {
		lang.SayString("xdiv032")
	} else {
		if !(rexsult.ToString() == "-3863.30675") {
			lang.SayString("xdiv032")
		}
	}
	rexsult, err = lang.RxFromString("615396156.E+549895291").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-29530247.4"))
	if err != nil {
		lang.SayString("xdiv033")
	} else {
		if !(rexsult.ToString() == "-2.08395191E+549895292") {
			lang.SayString("xdiv033")
		}
	}
	rexsult, err = lang.RxFromString("592.142173E-419941416").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3.46079109E-844011845"))
	if err != nil {
		lang.SayString("xdiv034")
	} else {
		if !(rexsult.ToString() == "-1.71100236E+424070431") {
			lang.SayString("xdiv034")
		}
	}
	rexsult, err = lang.RxFromString("849.515993E-878446473").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-1039.08778"))
	if err != nil {
		lang.SayString("xdiv035")
	} else {
		if !(rexsult.ToString() == "-8.17559411E-878446474") {
			lang.SayString("xdiv035")
		}
	}
	rexsult, err = lang.RxFromString("213361789").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-599.644851"))
	if err != nil {
		lang.SayString("xdiv036")
	} else {
		if !(rexsult.ToString() == "-355813.593") {
			lang.SayString("xdiv036")
		}
	}
	rexsult, err = lang.RxFromString("-795522555.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-298.037702"))
	if err != nil {
		lang.SayString("xdiv037")
	} else {
		if !(rexsult.ToString() == "2669201.08") {
			lang.SayString("xdiv037")
		}
	}
	rexsult, err = lang.RxFromString("-501260651.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8761893.0E-689281479"))
	if err != nil {
		lang.SayString("xdiv038")
	} else {
		if !(rexsult.ToString() == "5.72091728E+689281480") {
			lang.SayString("xdiv038")
		}
	}
	rexsult, err = lang.RxFromString("-1.70781105E-848889023").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("36504769.4"))
	if err != nil {
		lang.SayString("xdiv039")
	} else {
		if !(rexsult.ToString() == "-4.67832307E-848889031") {
			lang.SayString("xdiv039")
		}
	}
	rexsult, err = lang.RxFromString("-5290.54984E-490626676").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("842535254"))
	if err != nil {
		lang.SayString("xdiv040")
	} else {
		if !(rexsult.ToString() == "-6.27932162E-490626682") {
			lang.SayString("xdiv040")
		}
	}
	rexsult, err = lang.RxFromString("608.31825E+535268120").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-59609.0993"))
	if err != nil {
		lang.SayString("xdiv041")
	} else {
		if !(rexsult.ToString() == "-1.0205124E+535268118") {
			lang.SayString("xdiv041")
		}
	}
	rexsult, err = lang.RxFromString("-4629035.31").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-167.884398"))
	if err != nil {
		lang.SayString("xdiv042")
	} else {
		if !(rexsult.ToString() == "27572.7546") {
			lang.SayString("xdiv042")
		}
	}
	rexsult, err = lang.RxFromString("-66527378.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-706400268."))
	if err != nil {
		lang.SayString("xdiv043")
	} else {
		if !(rexsult.ToString() == "0.0941780192") {
			lang.SayString("xdiv043")
		}
	}
	rexsult, err = lang.RxFromString("-2510497.53").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("372882462."))
	if err != nil {
		lang.SayString("xdiv044")
	} else {
		if !(rexsult.ToString() == "-0.00673267795") {
			lang.SayString("xdiv044")
		}
	}
	rexsult, err = lang.RxFromString("136.255393E+53329961").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-53494.7201E+720058060"))
	if err != nil {
		lang.SayString("xdiv045")
	} else {
		if !(rexsult.ToString() == "-2.54708115E-666728102") {
			lang.SayString("xdiv045")
		}
	}
	rexsult, err = lang.RxFromString("-876673.100").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-6150.92266"))
	if err != nil {
		lang.SayString("xdiv046")
	} else {
		if !(rexsult.ToString() == "142.527089") {
			lang.SayString("xdiv046")
		}
	}
	rexsult, err = lang.RxFromString("-2.45151797E+911306117").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("27235771"))
	if err != nil {
		lang.SayString("xdiv047")
	} else {
		if !(rexsult.ToString() == "-9.00109628E+911306109") {
			lang.SayString("xdiv047")
		}
	}
	rexsult, err = lang.RxFromString("-9.15117551").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4.95100733E-314511326"))
	if err != nil {
		lang.SayString("xdiv048")
	} else {
		if !(rexsult.ToString() == "1.84834618E+314511326") {
			lang.SayString("xdiv048")
		}
	}
	rexsult, err = lang.RxFromString("3.61890453E-985606128").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("30664416."))
	if err != nil {
		lang.SayString("xdiv049")
	} else {
		if !(rexsult.ToString() == "1.18016418E-985606135") {
			lang.SayString("xdiv049")
		}
	}
	rexsult, err = lang.RxFromString("-257674602E+216723382").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-70820959.4"))
	if err != nil {
		lang.SayString("xdiv050")
	} else {
		if !(rexsult.ToString() == "3.63839468E+216723382") {
			lang.SayString("xdiv050")
		}
	}
	rexsult, err = lang.RxFromString("218699.206").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("556944241."))
	if err != nil {
		lang.SayString("xdiv051")
	} else {
		if !(rexsult.ToString() == "0.000392677022") {
			lang.SayString("xdiv051")
		}
	}
	rexsult, err = lang.RxFromString("106211716.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3456793.74"))
	if err != nil {
		lang.SayString("xdiv052")
	} else {
		if !(rexsult.ToString() == "-30.7255") {
			lang.SayString("xdiv052")
		}
	}
	rexsult, err = lang.RxFromString("1.25018078").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("399856.763E-726816740"))
	if err != nil {
		lang.SayString("xdiv053")
	} else {
		if !(rexsult.ToString() == "3.12657155E+726816734") {
			lang.SayString("xdiv053")
		}
	}
	rexsult, err = lang.RxFromString("364.99811").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-46222.0505"))
	if err != nil {
		lang.SayString("xdiv054")
	} else {
		if !(rexsult.ToString() == "-0.00789662306") {
			lang.SayString("xdiv054")
		}
	}
	rexsult, err = lang.RxFromString("-392217576.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-958364096"))
	if err != nil {
		lang.SayString("xdiv055")
	} else {
		if !(rexsult.ToString() == "0.409257377") {
			lang.SayString("xdiv055")
		}
	}
	rexsult, err = lang.RxFromString("169601285").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("714526.639"))
	if err != nil {
		lang.SayString("xdiv056")
	} else {
		if !(rexsult.ToString() == "237.361738") {
			lang.SayString("xdiv056")
		}
	}
	rexsult, err = lang.RxFromString("-674.094552E+586944319").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("6354.2668E+589657266"))
	if err != nil {
		lang.SayString("xdiv057")
	} else {
		if !(rexsult.ToString() == "-1.0608534E-2712948") {
			lang.SayString("xdiv057")
		}
	}
	rexsult, err = lang.RxFromString("151795163E-371727182").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-488.09788E-738852245"))
	if err != nil {
		lang.SayString("xdiv058")
	} else {
		if !(rexsult.ToString() == "-3.10993285E+367125068") {
			lang.SayString("xdiv058")
		}
	}
	rexsult, err = lang.RxFromString("-746.293386").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("927749.647"))
	if err != nil {
		lang.SayString("xdiv059")
	} else {
		if !(rexsult.ToString() == "-0.000804412471") {
			lang.SayString("xdiv059")
		}
	}
	rexsult, err = lang.RxFromString("888946471E+241331592").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-235739.595"))
	if err != nil {
		lang.SayString("xdiv060")
	} else {
		if !(rexsult.ToString() == "-3.77088317E+241331595") {
			lang.SayString("xdiv060")
		}
	}
	rexsult, err = lang.RxFromString("6.64377249").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("79161.1070E+619453776"))
	if err != nil {
		lang.SayString("xdiv061")
	} else {
		if !(rexsult.ToString() == "8.39272307E-619453781") {
			lang.SayString("xdiv061")
		}
	}
	rexsult, err = lang.RxFromString("3146.66571E-313373366").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("88.5282010"))
	if err != nil {
		lang.SayString("xdiv062")
	} else {
		if !(rexsult.ToString() == "3.55442184E-313373365") {
			lang.SayString("xdiv062")
		}
	}
	rexsult, err = lang.RxFromString("6.44693097").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-87195.8711"))
	if err != nil {
		lang.SayString("xdiv063")
	} else {
		if !(rexsult.ToString() == "-0.0000739361955") {
			lang.SayString("xdiv063")
		}
	}
	rexsult, err = lang.RxFromString("-2113132.56E+577957840").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("981125821"))
	if err != nil {
		lang.SayString("xdiv064")
	} else {
		if !(rexsult.ToString() == "-2.15378345E+577957837") {
			lang.SayString("xdiv064")
		}
	}
	rexsult, err = lang.RxFromString("-7701.42814").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("72667.5181"))
	if err != nil {
		lang.SayString("xdiv065")
	} else {
		if !(rexsult.ToString() == "-0.105981714") {
			lang.SayString("xdiv065")
		}
	}
	rexsult, err = lang.RxFromString("-851.754789").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-582659.149"))
	if err != nil {
		lang.SayString("xdiv066")
	} else {
		if !(rexsult.ToString() == "0.00146184058") {
			lang.SayString("xdiv066")
		}
	}
	rexsult, err = lang.RxFromString("-5.01992943").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7852.16531"))
	if err != nil {
		lang.SayString("xdiv067")
	} else {
		if !(rexsult.ToString() == "-0.000639305113") {
			lang.SayString("xdiv067")
		}
	}
	rexsult, err = lang.RxFromString("-12393257.2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("76803689E+949125770"))
	if err != nil {
		lang.SayString("xdiv068")
	} else {
		if !(rexsult.ToString() == "-1.61362786E-949125771") {
			lang.SayString("xdiv068")
		}
	}
	rexsult, err = lang.RxFromString("-754771634.E+716555026").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-292336.311"))
	if err != nil {
		lang.SayString("xdiv069")
	} else {
		if !(rexsult.ToString() == "2.5818607E+716555029") {
			lang.SayString("xdiv069")
		}
	}
	rexsult, err = lang.RxFromString("-915006.171E+614548652").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-314086965."))
	if err != nil {
		lang.SayString("xdiv070")
	} else {
		if !(rexsult.ToString() == "2.91322555E+614548649") {
			lang.SayString("xdiv070")
		}
	}
	rexsult, err = lang.RxFromString("-296590035").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-481734529"))
	if err != nil {
		lang.SayString("xdiv071")
	} else {
		if !(rexsult.ToString() == "0.615671116") {
			lang.SayString("xdiv071")
		}
	}
	rexsult, err = lang.RxFromString("8.27822605").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9241557.19"))
	if err != nil {
		lang.SayString("xdiv072")
	} else {
		if !(rexsult.ToString() == "8.9576095E-7") {
			lang.SayString("xdiv072")
		}
	}
	rexsult, err = lang.RxFromString("-1.43581098").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7286313.54"))
	if err != nil {
		lang.SayString("xdiv073")
	} else {
		if !(rexsult.ToString() == "-1.97055887E-7") {
			lang.SayString("xdiv073")
		}
	}
	rexsult, err = lang.RxFromString("-699036193.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("759263.509E+533543625"))
	if err != nil {
		lang.SayString("xdiv074")
	} else {
		if !(rexsult.ToString() == "-9.20676662E-533543623") {
			lang.SayString("xdiv074")
		}
	}
	rexsult, err = lang.RxFromString("-83.7273615E-305281957").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-287779593.E+458777774"))
	if err != nil {
		lang.SayString("xdiv075")
	} else {
		if !(rexsult.ToString() == "2.90942664E-764059738") {
			lang.SayString("xdiv075")
		}
	}
	rexsult, err = lang.RxFromString("8.48503224").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("6522.03316"))
	if err != nil {
		lang.SayString("xdiv076")
	} else {
		if !(rexsult.ToString() == "0.00130097962") {
			lang.SayString("xdiv076")
		}
	}
	rexsult, err = lang.RxFromString("527916091").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-809.054070"))
	if err != nil {
		lang.SayString("xdiv077")
	} else {
		if !(rexsult.ToString() == "-652510.272") {
			lang.SayString("xdiv077")
		}
	}
	rexsult, err = lang.RxFromString("3857058.60").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5792997.58E+881077409"))
	if err != nil {
		lang.SayString("xdiv078")
	} else {
		if !(rexsult.ToString() == "6.6581395E-881077410") {
			lang.SayString("xdiv078")
		}
	}
	rexsult, err = lang.RxFromString("-66587363.E+556538173").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-551902402E+357309146"))
	if err != nil {
		lang.SayString("xdiv079")
	} else {
		if !(rexsult.ToString() == "1.20650613E+199229026") {
			lang.SayString("xdiv079")
		}
	}
	rexsult, err = lang.RxFromString("-580.502955").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("38125521.7"))
	if err != nil {
		lang.SayString("xdiv080")
	} else {
		if !(rexsult.ToString() == "-0.0000152260987") {
			lang.SayString("xdiv080")
		}
	}
	rexsult, err = lang.RxFromString("-9627363.00").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-80616885E-749891394"))
	if err != nil {
		lang.SayString("xdiv081")
	} else {
		if !(rexsult.ToString() == "1.19421173E+749891393") {
			lang.SayString("xdiv081")
		}
	}
	rexsult, err = lang.RxFromString("-526.594855E+803110107").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-64.5451639"))
	if err != nil {
		lang.SayString("xdiv082")
	} else {
		if !(rexsult.ToString() == "8.15854858E+803110107") {
			lang.SayString("xdiv082")
		}
	}
	rexsult, err = lang.RxFromString("-8378.55499").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("760.131257"))
	if err != nil {
		lang.SayString("xdiv083")
	} else {
		if !(rexsult.ToString() == "-11.0225108") {
			lang.SayString("xdiv083")
		}
	}
	rexsult, err = lang.RxFromString("-717.697718").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("984304413"))
	if err != nil {
		lang.SayString("xdiv084")
	} else {
		if !(rexsult.ToString() == "-7.2914203E-7") {
			lang.SayString("xdiv084")
		}
	}
	rexsult, err = lang.RxFromString("-76762243.4E-741100094").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-273.706674"))
	if err != nil {
		lang.SayString("xdiv085")
	} else {
		if !(rexsult.ToString() == "2.80454409E-741100089") {
			lang.SayString("xdiv085")
		}
	}
	rexsult, err = lang.RxFromString("-701.518354E+786274918").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8822750.68E+243052107"))
	if err != nil {
		lang.SayString("xdiv086")
	} else {
		if !(rexsult.ToString() == "-7.95124309E+543222806") {
			lang.SayString("xdiv086")
		}
	}
	rexsult, err = lang.RxFromString("-359866845.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4.57434117"))
	if err != nil {
		lang.SayString("xdiv087")
	} else {
		if !(rexsult.ToString() == "78670748.8") {
			lang.SayString("xdiv087")
		}
	}
	rexsult, err = lang.RxFromString("779934536.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-76562645.7"))
	if err != nil {
		lang.SayString("xdiv088")
	} else {
		if !(rexsult.ToString() == "-10.1868807") {
			lang.SayString("xdiv088")
		}
	}
	rexsult, err = lang.RxFromString("-4820.95451").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3516234.99E+303303176"))
	if err != nil {
		lang.SayString("xdiv089")
	} else {
		if !(rexsult.ToString() == "-1.37105584E-303303179") {
			lang.SayString("xdiv089")
		}
	}
	rexsult, err = lang.RxFromString("69355976.9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-9.57838562E+758804984"))
	if err != nil {
		lang.SayString("xdiv090")
	} else {
		if !(rexsult.ToString() == "-7.24088376E-758804978") {
			lang.SayString("xdiv090")
		}
	}
	rexsult, err = lang.RxFromString("-12672093.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8569.78255E-382866025"))
	if err != nil {
		lang.SayString("xdiv091")
	} else {
		if !(rexsult.ToString() == "-1.47869482E+382866028") {
			lang.SayString("xdiv091")
		}
	}
	rexsult, err = lang.RxFromString("-5910750.2").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("66150383E-662459241"))
	if err != nil {
		lang.SayString("xdiv092")
	} else {
		if !(rexsult.ToString() == "-8.93532272E+662459239") {
			lang.SayString("xdiv092")
		}
	}
	rexsult, err = lang.RxFromString("-532577268.E-163806629").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-240650398E-650110558"))
	if err != nil {
		lang.SayString("xdiv093")
	} else {
		if !(rexsult.ToString() == "2.21307454E+486303929") {
			lang.SayString("xdiv093")
		}
	}
	rexsult, err = lang.RxFromString("-671.507198E-908587890").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3057429.32E-555230623"))
	if err != nil {
		lang.SayString("xdiv094")
	} else {
		if !(rexsult.ToString() == "-2.19631307E-353357271") {
			lang.SayString("xdiv094")
		}
	}
	rexsult, err = lang.RxFromString("-294.994352E+346452027").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-6061853.0"))
	if err != nil {
		lang.SayString("xdiv095")
	} else {
		if !(rexsult.ToString() == "4.86640557E+346452022") {
			lang.SayString("xdiv095")
		}
	}
	rexsult, err = lang.RxFromString("329579114").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("146780548."))
	if err != nil {
		lang.SayString("xdiv096")
	} else {
		if !(rexsult.ToString() == "2.24538686") {
			lang.SayString("xdiv096")
		}
	}
	rexsult, err = lang.RxFromString("-789904.686E-217225000").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-1991.07181E-84080059"))
	if err != nil {
		lang.SayString("xdiv097")
	} else {
		if !(rexsult.ToString() == "3.96723354E-133144939") {
			lang.SayString("xdiv097")
		}
	}
	rexsult, err = lang.RxFromString("59893.3544").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-408595868"))
	if err != nil {
		lang.SayString("xdiv098")
	} else {
		if !(rexsult.ToString() == "-0.000146583358") {
			lang.SayString("xdiv098")
		}
	}
	rexsult, err = lang.RxFromString("129.878613").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-54652.7288E-963564940"))
	if err != nil {
		lang.SayString("xdiv099")
	} else {
		if !(rexsult.ToString() == "-2.37643418E+963564937") {
			lang.SayString("xdiv099")
		}
	}
	rexsult, err = lang.RxFromString("9866.99208").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("708756501."))
	if err != nil {
		lang.SayString("xdiv100")
	} else {
		if !(rexsult.ToString() == "0.0000139215543") {
			lang.SayString("xdiv100")
		}
	}
	rexsult, err = lang.RxFromString("-78810.6297").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-399884.68"))
	if err != nil {
		lang.SayString("xdiv101")
	} else {
		if !(rexsult.ToString() == "0.197083393") {
			lang.SayString("xdiv101")
		}
	}
	rexsult, err = lang.RxFromString("409189761").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-771.471460"))
	if err != nil {
		lang.SayString("xdiv102")
	} else {
		if !(rexsult.ToString() == "-530401.683") {
			lang.SayString("xdiv102")
		}
	}
	rexsult, err = lang.RxFromString("-1.68748838").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("460.46924"))
	if err != nil {
		lang.SayString("xdiv103")
	} else {
		if !(rexsult.ToString() == "-0.00366471467") {
			lang.SayString("xdiv103")
		}
	}
	rexsult, err = lang.RxFromString("553527296.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7924.40185"))
	if err != nil {
		lang.SayString("xdiv104")
	} else {
		if !(rexsult.ToString() == "-69850.9877") {
			lang.SayString("xdiv104")
		}
	}
	rexsult, err = lang.RxFromString("-38.7465207").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("64936.2942"))
	if err != nil {
		lang.SayString("xdiv105")
	} else {
		if !(rexsult.ToString() == "-0.000596685123") {
			lang.SayString("xdiv105")
		}
	}
	rexsult, err = lang.RxFromString("-201075.248").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("845.663928"))
	if err != nil {
		lang.SayString("xdiv106")
	} else {
		if !(rexsult.ToString() == "-237.772053") {
			lang.SayString("xdiv106")
		}
	}
	rexsult, err = lang.RxFromString("91048.4559").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("75953609.3"))
	if err != nil {
		lang.SayString("xdiv107")
	} else {
		if !(rexsult.ToString() == "0.00119873771") {
			lang.SayString("xdiv107")
		}
	}
	rexsult, err = lang.RxFromString("6898273.86E-252097460").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("15.3456196"))
	if err != nil {
		lang.SayString("xdiv108")
	} else {
		if !(rexsult.ToString() == "4.49527229E-252097455") {
			lang.SayString("xdiv108")
		}
	}
	rexsult, err = lang.RxFromString("88.4370343").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-980709105E-869899289"))
	if err != nil {
		lang.SayString("xdiv109")
	} else {
		if !(rexsult.ToString() == "-9.0176622E+869899281") {
			lang.SayString("xdiv109")
		}
	}
	rexsult, err = lang.RxFromString("-17643.39").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.0352568E+304871331"))
	if err != nil {
		lang.SayString("xdiv110")
	} else {
		if !(rexsult.ToString() == "-8.66887658E-304871328") {
			lang.SayString("xdiv110")
		}
	}
	rexsult, err = lang.RxFromString("4589785.16").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7459.04237"))
	if err != nil {
		lang.SayString("xdiv111")
	} else {
		if !(rexsult.ToString() == "615.331692") {
			lang.SayString("xdiv111")
		}
	}
	rexsult, err = lang.RxFromString("-51.1632090E-753968082").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8.96207471E-585797887"))
	if err != nil {
		lang.SayString("xdiv112")
	} else {
		if !(rexsult.ToString() == "-5.70885768E-168170195") {
			lang.SayString("xdiv112")
		}
	}
	rexsult, err = lang.RxFromString("982.217817").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7224909.4E-45243816"))
	if err != nil {
		lang.SayString("xdiv113")
	} else {
		if !(rexsult.ToString() == "1.35948807E+45243812") {
			lang.SayString("xdiv113")
		}
	}
	rexsult, err = lang.RxFromString("503712056.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-57490703.5E+924890183"))
	if err != nil {
		lang.SayString("xdiv114")
	} else {
		if !(rexsult.ToString() == "-8.76162623E-924890183") {
			lang.SayString("xdiv114")
		}
	}
	rexsult, err = lang.RxFromString("883.849223").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("249259171"))
	if err != nil {
		lang.SayString("xdiv115")
	} else {
		if !(rexsult.ToString() == "0.00000354590453") {
			lang.SayString("xdiv115")
		}
	}
	rexsult, err = lang.RxFromString("245.092853E+872642874").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("828195.152E+419771532"))
	if err != nil {
		lang.SayString("xdiv116")
	} else {
		if !(rexsult.ToString() == "2.95936112E+452871338") {
			lang.SayString("xdiv116")
		}
	}
	rexsult, err = lang.RxFromString("-83658638.6E+728551928").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2952478.42"))
	if err != nil {
		lang.SayString("xdiv117")
	} else {
		if !(rexsult.ToString() == "-2.83350551E+728551929") {
			lang.SayString("xdiv117")
		}
	}
	rexsult, err = lang.RxFromString("-6291780.97").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("269967.394E-22000817"))
	if err != nil {
		lang.SayString("xdiv118")
	} else {
		if !(rexsult.ToString() == "-2.33057069E+22000818") {
			lang.SayString("xdiv118")
		}
	}
	rexsult, err = lang.RxFromString("978571348.E+222382547").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("6006.19370"))
	if err != nil {
		lang.SayString("xdiv119")
	} else {
		if !(rexsult.ToString() == "1.62927038E+222382552") {
			lang.SayString("xdiv119")
		}
	}
	rexsult, err = lang.RxFromString("14239029.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-36527.2221"))
	if err != nil {
		lang.SayString("xdiv120")
	} else {
		if !(rexsult.ToString() == "-389.819652") {
			lang.SayString("xdiv120")
		}
	}
	rexsult, err = lang.RxFromString("-37721.1567E-115787341").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-828949864E-76251747"))
	if err != nil {
		lang.SayString("xdiv122")
	} else {
		if !(rexsult.ToString() == "4.55047505E-39535599") {
			lang.SayString("xdiv122")
		}
	}
	rexsult, err = lang.RxFromString("-79145.3625").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7718.57307"))
	if err != nil {
		lang.SayString("xdiv124")
	} else {
		if !(rexsult.ToString() == "10.2538852") {
			lang.SayString("xdiv124")
		}
	}
	rexsult, err = lang.RxFromString("2103890.49E+959247237").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("20024.3017"))
	if err != nil {
		lang.SayString("xdiv125")
	} else {
		if !(rexsult.ToString() == "1.05066859E+959247239") {
			lang.SayString("xdiv125")
		}
	}
	rexsult, err = lang.RxFromString("911249557").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("79810804.9"))
	if err != nil {
		lang.SayString("xdiv126")
	} else {
		if !(rexsult.ToString() == "11.4176214") {
			lang.SayString("xdiv126")
		}
	}
	rexsult, err = lang.RxFromString("341134.994").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3.37486292"))
	if err != nil {
		lang.SayString("xdiv127")
	} else {
		if !(rexsult.ToString() == "101081.141") {
			lang.SayString("xdiv127")
		}
	}
	rexsult, err = lang.RxFromString("244.23634").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("512706190E-341459836"))
	if err != nil {
		lang.SayString("xdiv128")
	} else {
		if !(rexsult.ToString() == "4.76367059E+341459829") {
			lang.SayString("xdiv128")
		}
	}
	rexsult, err = lang.RxFromString("-9.22783849E+171585954").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-99.0946800"))
	if err != nil {
		lang.SayString("xdiv129")
	} else {
		if !(rexsult.ToString() == "9.31214318E+171585952") {
			lang.SayString("xdiv129")
		}
	}
	rexsult, err = lang.RxFromString("699631.893").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-226.423958"))
	if err != nil {
		lang.SayString("xdiv130")
	} else {
		if !(rexsult.ToString() == "-3089.9199") {
			lang.SayString("xdiv130")
		}
	}
	rexsult, err = lang.RxFromString("-249350139.E-571793673").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("775732428."))
	if err != nil {
		lang.SayString("xdiv131")
	} else {
		if !(rexsult.ToString() == "-3.21438334E-571793674") {
			lang.SayString("xdiv131")
		}
	}
	rexsult, err = lang.RxFromString("5.11629020").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-480.53194"))
	if err != nil {
		lang.SayString("xdiv132")
	} else {
		if !(rexsult.ToString() == "-0.0106471387") {
			lang.SayString("xdiv132")
		}
	}
	rexsult, err = lang.RxFromString("-8.23352673E-446723147").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-530710.866"))
	if err != nil {
		lang.SayString("xdiv133")
	} else {
		if !(rexsult.ToString() == "1.55141476E-446723152") {
			lang.SayString("xdiv133")
		}
	}
	rexsult, err = lang.RxFromString("7.0598608").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-95908.35"))
	if err != nil {
		lang.SayString("xdiv134")
	} else {
		if !(rexsult.ToString() == "-0.0000736104917") {
			lang.SayString("xdiv134")
		}
	}
	rexsult, err = lang.RxFromString("-7.91189845E+207202706").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1532.71847E+509944335"))
	if err != nil {
		lang.SayString("xdiv135")
	} else {
		if !(rexsult.ToString() == "-5.16200372E-302741632") {
			lang.SayString("xdiv135")
		}
	}
	rexsult, err = lang.RxFromString("208839370.E-215147432").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-75.9420559"))
	if err != nil {
		lang.SayString("xdiv136")
	} else {
		if !(rexsult.ToString() == "-2.7499831E-215147426") {
			lang.SayString("xdiv136")
		}
	}
	rexsult, err = lang.RxFromString("427.754244E-353328369").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4705.0796"))
	if err != nil {
		lang.SayString("xdiv137")
	} else {
		if !(rexsult.ToString() == "9.09132853E-353328371") {
			lang.SayString("xdiv137")
		}
	}
	rexsult, err = lang.RxFromString("44911.089").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-95.1733605E-313081848"))
	if err != nil {
		lang.SayString("xdiv138")
	} else {
		if !(rexsult.ToString() == "-4.71887183E+313081850") {
			lang.SayString("xdiv138")
		}
	}
	rexsult, err = lang.RxFromString("452371821.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4109709.19"))
	if err != nil {
		lang.SayString("xdiv139")
	} else {
		if !(rexsult.ToString() == "-110.073925") {
			lang.SayString("xdiv139")
		}
	}
	rexsult, err = lang.RxFromString("94007.4392").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-9467725.5E+681898234"))
	if err != nil {
		lang.SayString("xdiv140")
	} else {
		if !(rexsult.ToString() == "-9.92925272E-681898237") {
			lang.SayString("xdiv140")
		}
	}
	rexsult, err = lang.RxFromString("99147554.0E-751410586").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("38313.6423"))
	if err != nil {
		lang.SayString("xdiv141")
	} else {
		if !(rexsult.ToString() == "2.58778722E-751410583") {
			lang.SayString("xdiv141")
		}
	}
	rexsult, err = lang.RxFromString("-7919.30254").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-669.607854"))
	if err != nil {
		lang.SayString("xdiv142")
	} else {
		if !(rexsult.ToString() == "11.8267767") {
			lang.SayString("xdiv142")
		}
	}
	rexsult, err = lang.RxFromString("461.58280E+136110821").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("710666052.E-383754231"))
	if err != nil {
		lang.SayString("xdiv143")
	} else {
		if !(rexsult.ToString() == "6.49507316E+519865045") {
			lang.SayString("xdiv143")
		}
	}
	rexsult, err = lang.RxFromString("3455755.47E-112465506").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("771.674306"))
	if err != nil {
		lang.SayString("xdiv144")
	} else {
		if !(rexsult.ToString() == "4.47825649E-112465503") {
			lang.SayString("xdiv144")
		}
	}
	rexsult, err = lang.RxFromString("-477067757.E-961684940").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7.70122608E-741072245"))
	if err != nil {
		lang.SayString("xdiv145")
	} else {
		if !(rexsult.ToString() == "-6.19469877E-220612688") {
			lang.SayString("xdiv145")
		}
	}
	rexsult, err = lang.RxFromString("76482.352").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8237806.8"))
	if err != nil {
		lang.SayString("xdiv146")
	} else {
		if !(rexsult.ToString() == "0.00928430999") {
			lang.SayString("xdiv146")
		}
	}
	rexsult, err = lang.RxFromString("1.21505164E-565556504").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9.26146573"))
	if err != nil {
		lang.SayString("xdiv147")
	} else {
		if !(rexsult.ToString() == "1.31194314E-565556505") {
			lang.SayString("xdiv147")
		}
	}
	rexsult, err = lang.RxFromString("-8303060.25E-169894883").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("901561.985"))
	if err != nil {
		lang.SayString("xdiv148")
	} else {
		if !(rexsult.ToString() == "-9.20963881E-169894883") {
			lang.SayString("xdiv148")
		}
	}
	rexsult, err = lang.RxFromString("-592464.92").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("71445510.7"))
	if err != nil {
		lang.SayString("xdiv149")
	} else {
		if !(rexsult.ToString() == "-0.00829254231") {
			lang.SayString("xdiv149")
		}
	}
	rexsult, err = lang.RxFromString("-73774.4165").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-39.8243027"))
	if err != nil {
		lang.SayString("xdiv150")
	} else {
		if !(rexsult.ToString() == "1852.49738") {
			lang.SayString("xdiv150")
		}
	}
	rexsult, err = lang.RxFromString("-524724715.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-55763.7937"))
	if err != nil {
		lang.SayString("xdiv151")
	} else {
		if !(rexsult.ToString() == "9409.77434") {
			lang.SayString("xdiv151")
		}
	}
	rexsult, err = lang.RxFromString("7.53800427").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("784873768E-9981146"))
	if err != nil {
		lang.SayString("xdiv152")
	} else {
		if !(rexsult.ToString() == "9.6040976E+9981137") {
			lang.SayString("xdiv152")
		}
	}
	rexsult, err = lang.RxFromString("37.6027452").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7.22454233"))
	if err != nil {
		lang.SayString("xdiv153")
	} else {
		if !(rexsult.ToString() == "5.20486191") {
			lang.SayString("xdiv153")
		}
	}
	rexsult, err = lang.RxFromString("2447660.39").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-36981.4253"))
	if err != nil {
		lang.SayString("xdiv154")
	} else {
		if !(rexsult.ToString() == "-66.1862102") {
			lang.SayString("xdiv154")
		}
	}
	rexsult, err = lang.RxFromString("2160.36419").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1418.33574E+656265382"))
	if err != nil {
		lang.SayString("xdiv155")
	} else {
		if !(rexsult.ToString() == "1.52316841E-656265382") {
			lang.SayString("xdiv155")
		}
	}
	rexsult, err = lang.RxFromString("8926.44939").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("54.9430027"))
	if err != nil {
		lang.SayString("xdiv156")
	} else {
		if !(rexsult.ToString() == "162.467447") {
			lang.SayString("xdiv156")
		}
	}
	rexsult, err = lang.RxFromString("861588029").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-41657398E+77955925"))
	if err != nil {
		lang.SayString("xdiv157")
	} else {
		if !(rexsult.ToString() == "-2.06827135E-77955924") {
			lang.SayString("xdiv157")
		}
	}
	rexsult, err = lang.RxFromString("-34.5253062").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("52.6722019"))
	if err != nil {
		lang.SayString("xdiv158")
	} else {
		if !(rexsult.ToString() == "-0.655474899") {
			lang.SayString("xdiv158")
		}
	}
	rexsult, err = lang.RxFromString("-18861647.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("99794586.7"))
	if err != nil {
		lang.SayString("xdiv159")
	} else {
		if !(rexsult.ToString() == "-0.189004711") {
			lang.SayString("xdiv159")
		}
	}
	rexsult, err = lang.RxFromString("322192.407").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("461.67044"))
	if err != nil {
		lang.SayString("xdiv160")
	} else {
		if !(rexsult.ToString() == "697.883986") {
			lang.SayString("xdiv160")
		}
	}
	rexsult, err = lang.RxFromString("-896298518E+61412314").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("78873.8049"))
	if err != nil {
		lang.SayString("xdiv161")
	} else {
		if !(rexsult.ToString() == "-1.13637033E+61412318") {
			lang.SayString("xdiv161")
		}
	}
	rexsult, err = lang.RxFromString("293.773732").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("479899052E+789950177"))
	if err != nil {
		lang.SayString("xdiv162")
	} else {
		if !(rexsult.ToString() == "6.1215735E-789950184") {
			lang.SayString("xdiv162")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("51897955.3"))
	if err != nil {
		lang.SayString("xdiv163")
	} else {
		if !(rexsult.ToString() == "-1.9946713") {
			lang.SayString("xdiv163")
		}
	}
	rexsult, err = lang.RxFromString("37380.7802").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-277719788."))
	if err != nil {
		lang.SayString("xdiv164")
	} else {
		if !(rexsult.ToString() == "-0.000134598908") {
			lang.SayString("xdiv164")
		}
	}
	rexsult, err = lang.RxFromString("320133844.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-977517477"))
	if err != nil {
		lang.SayString("xdiv165")
	} else {
		if !(rexsult.ToString() == "-0.327496798") {
			lang.SayString("xdiv165")
		}
	}
	rexsult, err = lang.RxFromString("721776701E+933646161").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-5689279.64E+669903645"))
	if err != nil {
		lang.SayString("xdiv166")
	} else {
		if !(rexsult.ToString() == "-1.26866097E+263742518") {
			lang.SayString("xdiv166")
		}
	}
	rexsult, err = lang.RxFromString("-5409.00482").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2.16149386"))
	if err != nil {
		lang.SayString("xdiv167")
	} else {
		if !(rexsult.ToString() == "2502.43821") {
			lang.SayString("xdiv167")
		}
	}
	rexsult, err = lang.RxFromString("-957960.367").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("322.858170"))
	if err != nil {
		lang.SayString("xdiv168")
	} else {
		if !(rexsult.ToString() == "-2967.12444") {
			lang.SayString("xdiv168")
		}
	}
	rexsult, err = lang.RxFromString("-11817.8754E+613893442").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3.84735082E+888333249"))
	if err != nil {
		lang.SayString("xdiv169")
	} else {
		if !(rexsult.ToString() == "3.07169165E-274439804") {
			lang.SayString("xdiv169")
		}
	}
	rexsult, err = lang.RxFromString("840258203").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("58363.980E-906584723"))
	if err != nil {
		lang.SayString("xdiv170")
	} else {
		if !(rexsult.ToString() == "1.43968626E+906584727") {
			lang.SayString("xdiv170")
		}
	}
	rexsult, err = lang.RxFromString("-205842096.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-191342.721"))
	if err != nil {
		lang.SayString("xdiv171")
	} else {
		if !(rexsult.ToString() == "1075.77699") {
			lang.SayString("xdiv171")
		}
	}
	rexsult, err = lang.RxFromString("42501124.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("884.938498E+123341480"))
	if err != nil {
		lang.SayString("xdiv172")
	} else {
		if !(rexsult.ToString() == "4.80272065E-123341476") {
			lang.SayString("xdiv172")
		}
	}
	rexsult, err = lang.RxFromString("-57809452.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-620380746"))
	if err != nil {
		lang.SayString("xdiv173")
	} else {
		if !(rexsult.ToString() == "0.0931838268") {
			lang.SayString("xdiv173")
		}
	}
	rexsult, err = lang.RxFromString("-8022370.31").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9858581.6"))
	if err != nil {
		lang.SayString("xdiv174")
	} else {
		if !(rexsult.ToString() == "-0.813744881") {
			lang.SayString("xdiv174")
		}
	}
	rexsult, err = lang.RxFromString("-697273715E-242824870").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3.81757506"))
	if err != nil {
		lang.SayString("xdiv176")
	} else {
		if !(rexsult.ToString() == "1.82648331E-242824862") {
			lang.SayString("xdiv176")
		}
	}
	rexsult, err = lang.RxFromString("-7.42204403E-315716280").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8156111.67E+283261636"))
	if err != nil {
		lang.SayString("xdiv177")
	} else {
		if !(rexsult.ToString() == "9.09997843E-598977923") {
			lang.SayString("xdiv177")
		}
	}
	rexsult, err = lang.RxFromString("738063892").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("89900467.8"))
	if err != nil {
		lang.SayString("xdiv178")
	} else {
		if !(rexsult.ToString() == "8.20978923") {
			lang.SayString("xdiv178")
		}
	}
	rexsult, err = lang.RxFromString("-630309366").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-884783.338E-21595410"))
	if err != nil {
		lang.SayString("xdiv179")
	} else {
		if !(rexsult.ToString() == "7.12388377E+21595412") {
			lang.SayString("xdiv179")
		}
	}
	rexsult, err = lang.RxFromString("613.207774").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3007.78608"))
	if err != nil {
		lang.SayString("xdiv180")
	} else {
		if !(rexsult.ToString() == "-0.203873466") {
			lang.SayString("xdiv180")
		}
	}
	rexsult, err = lang.RxFromString("-93006222.3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3.08964619"))
	if err != nil {
		lang.SayString("xdiv181")
	} else {
		if !(rexsult.ToString() == "30102547.9") {
			lang.SayString("xdiv181")
		}
	}
	rexsult, err = lang.RxFromString("-18116.0621").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("34096.306E-270347092"))
	if err != nil {
		lang.SayString("xdiv182")
	} else {
		if !(rexsult.ToString() == "-5.31320375E+270347091") {
			lang.SayString("xdiv182")
		}
	}
	rexsult, err = lang.RxFromString("19272386.9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-410442379."))
	if err != nil {
		lang.SayString("xdiv183")
	} else {
		if !(rexsult.ToString() == "-0.0469551584") {
			lang.SayString("xdiv183")
		}
	}
	rexsult, err = lang.RxFromString("4180.30821").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-1.6439543E-624759104"))
	if err != nil {
		lang.SayString("xdiv184")
	} else {
		if !(rexsult.ToString() == "-2.54283724E+624759107") {
			lang.SayString("xdiv184")
		}
	}
	rexsult, err = lang.RxFromString("571.536725").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("389.899220"))
	if err != nil {
		lang.SayString("xdiv185")
	} else {
		if !(rexsult.ToString() == "1.46585757") {
			lang.SayString("xdiv185")
		}
	}
	rexsult, err = lang.RxFromString("-622007306.E+159924886").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-126.971745"))
	if err != nil {
		lang.SayString("xdiv186")
	} else {
		if !(rexsult.ToString() == "4.89878521E+159924892") {
			lang.SayString("xdiv186")
		}
	}
	rexsult, err = lang.RxFromString("-29.356551E-282816139").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("37141748E-903397821"))
	if err != nil {
		lang.SayString("xdiv187")
	} else {
		if !(rexsult.ToString() == "-7.90392283E+620581675") {
			lang.SayString("xdiv187")
		}
	}
	rexsult, err = lang.RxFromString("92427442.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("674334898."))
	if err != nil {
		lang.SayString("xdiv188")
	} else {
		if !(rexsult.ToString() == "0.137064599") {
			lang.SayString("xdiv188")
		}
	}
	rexsult, err = lang.RxFromString("44651895.7").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-910508.438"))
	if err != nil {
		lang.SayString("xdiv189")
	} else {
		if !(rexsult.ToString() == "-49.0406171") {
			lang.SayString("xdiv189")
		}
	}
	rexsult, err = lang.RxFromString("647897872.E+374021790").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-467.423029"))
	if err != nil {
		lang.SayString("xdiv190")
	} else {
		if !(rexsult.ToString() == "-1.38610601E+374021796") {
			lang.SayString("xdiv190")
		}
	}
	rexsult, err = lang.RxFromString("25.2592149").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("59.0436981"))
	if err != nil {
		lang.SayString("xdiv191")
	} else {
		if !(rexsult.ToString() == "0.427805434") {
			lang.SayString("xdiv191")
		}
	}
	rexsult, err = lang.RxFromString("-6.850835").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-1273.48240"))
	if err != nil {
		lang.SayString("xdiv192")
	} else {
		if !(rexsult.ToString() == "0.00537960713") {
			lang.SayString("xdiv192")
		}
	}
	rexsult, err = lang.RxFromString("174.272325").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5638.16229"))
	if err != nil {
		lang.SayString("xdiv193")
	} else {
		if !(rexsult.ToString() == "0.0309094198") {
			lang.SayString("xdiv193")
		}
	}
	rexsult, err = lang.RxFromString("3455629.76").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8.27332322"))
	if err != nil {
		lang.SayString("xdiv194")
	} else {
		if !(rexsult.ToString() == "-417683.399") {
			lang.SayString("xdiv194")
		}
	}
	rexsult, err = lang.RxFromString("-924337723E-640771235").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("86639377.1"))
	if err != nil {
		lang.SayString("xdiv195")
	} else {
		if !(rexsult.ToString() == "-1.06687947E-640771234") {
			lang.SayString("xdiv195")
		}
	}
	rexsult, err = lang.RxFromString("-620236932.E+656823969").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3364722.73"))
	if err != nil {
		lang.SayString("xdiv196")
	} else {
		if !(rexsult.ToString() == "-1.84335228E+656823971") {
			lang.SayString("xdiv196")
		}
	}
	rexsult, err = lang.RxFromString("9.10025079").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("702777882E-8192234"))
	if err != nil {
		lang.SayString("xdiv197")
	} else {
		if !(rexsult.ToString() == "1.29489715E+8192226") {
			lang.SayString("xdiv197")
		}
	}
	rexsult, err = lang.RxFromString("-18857539.9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("813013129."))
	if err != nil {
		lang.SayString("xdiv198")
	} else {
		if !(rexsult.ToString() == "-0.0231946315") {
			lang.SayString("xdiv198")
		}
	}
	rexsult, err = lang.RxFromString("-8.29530327").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3243419.57E+35688332"))
	if err != nil {
		lang.SayString("xdiv199")
	} else {
		if !(rexsult.ToString() == "-2.55757946E-35688338") {
			lang.SayString("xdiv199")
		}
	}
	rexsult, err = lang.RxFromString("-57101683.5").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("763551341E+991491712"))
	if err != nil {
		lang.SayString("xdiv200")
	} else {
		if !(rexsult.ToString() == "-7.47843405E-991491714") {
			lang.SayString("xdiv200")
		}
	}
	rexsult, err = lang.RxFromString("-603326.740").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1710.95183"))
	if err != nil {
		lang.SayString("xdiv201")
	} else {
		if !(rexsult.ToString() == "-352.626374") {
			lang.SayString("xdiv201")
		}
	}
	rexsult, err = lang.RxFromString("-48142763.3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-943434114"))
	if err != nil {
		lang.SayString("xdiv202")
	} else {
		if !(rexsult.ToString() == "0.0510292797") {
			lang.SayString("xdiv202")
		}
	}
	rexsult, err = lang.RxFromString("-204.586767").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-235.531847"))
	if err != nil {
		lang.SayString("xdiv203")
	} else {
		if !(rexsult.ToString() == "0.868616154") {
			lang.SayString("xdiv203")
		}
	}
	rexsult, err = lang.RxFromString("-70.3805581").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("830137.913"))
	if err != nil {
		lang.SayString("xdiv204")
	} else {
		if !(rexsult.ToString() == "-0.0000847817658") {
			lang.SayString("xdiv204")
		}
	}
	rexsult, err = lang.RxFromString("-8818.47606").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-60766.4571"))
	if err != nil {
		lang.SayString("xdiv205")
	} else {
		if !(rexsult.ToString() == "0.145120787") {
			lang.SayString("xdiv205")
		}
	}
	rexsult, err = lang.RxFromString("37060929.3E-168439509").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-79576717.1"))
	if err != nil {
		lang.SayString("xdiv206")
	} else {
		if !(rexsult.ToString() == "-4.65725788E-168439510") {
			lang.SayString("xdiv206")
		}
	}
	rexsult, err = lang.RxFromString("-656285310.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-107221462."))
	if err != nil {
		lang.SayString("xdiv207")
	} else {
		if !(rexsult.ToString() == "6.12083904") {
			lang.SayString("xdiv207")
		}
	}
	rexsult, err = lang.RxFromString("653397.125").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7195.30990"))
	if err != nil {
		lang.SayString("xdiv208")
	} else {
		if !(rexsult.ToString() == "90.8087538") {
			lang.SayString("xdiv208")
		}
	}
	rexsult, err = lang.RxFromString("56221910.0E+857909374").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-58.7247929"))
	if err != nil {
		lang.SayString("xdiv209")
	} else {
		if !(rexsult.ToString() == "-9.57379451E+857909379") {
			lang.SayString("xdiv209")
		}
	}
	rexsult, err = lang.RxFromString("809862859E+643769974").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-5.06784016"))
	if err != nil {
		lang.SayString("xdiv210")
	} else {
		if !(rexsult.ToString() == "-1.59804341E+643769982") {
			lang.SayString("xdiv210")
		}
	}
	rexsult, err = lang.RxFromString("-62011.4563E-117563240").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-57.1731586E+115657204"))
	if err != nil {
		lang.SayString("xdiv211")
	} else {
		if !(rexsult.ToString() == "1.08462534E-233220441") {
			lang.SayString("xdiv211")
		}
	}
	rexsult, err = lang.RxFromString("315.33351").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("91588.837E-536020149"))
	if err != nil {
		lang.SayString("xdiv212")
	} else {
		if !(rexsult.ToString() == "3.44292515E+536020146") {
			lang.SayString("xdiv212")
		}
	}
	rexsult, err = lang.RxFromString("739.944710").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("202949.175"))
	if err != nil {
		lang.SayString("xdiv213")
	} else {
		if !(rexsult.ToString() == "0.00364596067") {
			lang.SayString("xdiv213")
		}
	}
	rexsult, err = lang.RxFromString("87686.8016").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4204890.40"))
	if err != nil {
		lang.SayString("xdiv214")
	} else {
		if !(rexsult.ToString() == "0.0208535285") {
			lang.SayString("xdiv214")
		}
	}
	rexsult, err = lang.RxFromString("987126721.E-725794834").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4874166.23"))
	if err != nil {
		lang.SayString("xdiv215")
	} else {
		if !(rexsult.ToString() == "2.0252217E-725794832") {
			lang.SayString("xdiv215")
		}
	}
	rexsult, err = lang.RxFromString("728148726.E-661695938").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("32798.5202"))
	if err != nil {
		lang.SayString("xdiv216")
	} else {
		if !(rexsult.ToString() == "2.22006579E-661695934") {
			lang.SayString("xdiv216")
		}
	}
	rexsult, err = lang.RxFromString("7428219.97").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("667.326760"))
	if err != nil {
		lang.SayString("xdiv217")
	} else {
		if !(rexsult.ToString() == "11131.3084") {
			lang.SayString("xdiv217")
		}
	}
	rexsult, err = lang.RxFromString("-7291.19212").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("209.64966E-588526476"))
	if err != nil {
		lang.SayString("xdiv218")
	} else {
		if !(rexsult.ToString() == "-3.47779821E+588526477") {
			lang.SayString("xdiv218")
		}
	}
	rexsult, err = lang.RxFromString("-358.24550").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4447.78675E+601402509"))
	if err != nil {
		lang.SayString("xdiv219")
	} else {
		if !(rexsult.ToString() == "8.05446664E-601402511") {
			lang.SayString("xdiv219")
		}
	}
	rexsult, err = lang.RxFromString("118.621826").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2.72010038"))
	if err != nil {
		lang.SayString("xdiv220")
	} else {
		if !(rexsult.ToString() == "-43.6093561") {
			lang.SayString("xdiv220")
		}
	}
	rexsult, err = lang.RxFromString("8071961.94").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-135533740.E-102451543"))
	if err != nil {
		lang.SayString("xdiv221")
	} else {
		if !(rexsult.ToString() == "-5.9556845E+102451541") {
			lang.SayString("xdiv221")
		}
	}
	rexsult, err = lang.RxFromString("-35544.4029").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-567830.130"))
	if err != nil {
		lang.SayString("xdiv223")
	} else {
		if !(rexsult.ToString() == "0.0625968948") {
			lang.SayString("xdiv223")
		}
	}
	rexsult, err = lang.RxFromString("-7.16513047E+59297103").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("87767.8211"))
	if err != nil {
		lang.SayString("xdiv224")
	} else {
		if !(rexsult.ToString() == "-8.16373288E+59297098") {
			lang.SayString("xdiv224")
		}
	}
	rexsult, err = lang.RxFromString("-509.483395").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-147242915."))
	if err != nil {
		lang.SayString("xdiv225")
	} else {
		if !(rexsult.ToString() == "0.00000346015559") {
			lang.SayString("xdiv225")
		}
	}
	rexsult, err = lang.RxFromString("-7919047.28E+956041629").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-367667329"))
	if err != nil {
		lang.SayString("xdiv226")
	} else {
		if !(rexsult.ToString() == "2.15386211E+956041627") {
			lang.SayString("xdiv226")
		}
	}
	rexsult, err = lang.RxFromString("895612630.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-36.4104040"))
	if err != nil {
		lang.SayString("xdiv227")
	} else {
		if !(rexsult.ToString() == "-24597712") {
			lang.SayString("xdiv227")
		}
	}
	rexsult, err = lang.RxFromString("25455.4973").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2955.00006E+528196218"))
	if err != nil {
		lang.SayString("xdiv228")
	} else {
		if !(rexsult.ToString() == "8.61438131E-528196218") {
			lang.SayString("xdiv228")
		}
	}
	rexsult, err = lang.RxFromString("-112.294144E+273414172").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-71448007.7"))
	if err != nil {
		lang.SayString("xdiv229")
	} else {
		if !(rexsult.ToString() == "1.57169035E+273414166") {
			lang.SayString("xdiv229")
		}
	}
	rexsult, err = lang.RxFromString("62871.2202").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2484.0382E+211662557"))
	if err != nil {
		lang.SayString("xdiv230")
	} else {
		if !(rexsult.ToString() == "2.53100859E-211662556") {
			lang.SayString("xdiv230")
		}
	}
	rexsult, err = lang.RxFromString("71.9281575").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-9810012.5"))
	if err != nil {
		lang.SayString("xdiv231")
	} else {
		if !(rexsult.ToString() == "-0.0000073321168") {
			lang.SayString("xdiv231")
		}
	}
	rexsult, err = lang.RxFromString("-6388022.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-88.042967"))
	if err != nil {
		lang.SayString("xdiv232")
	} else {
		if !(rexsult.ToString() == "72555.7329") {
			lang.SayString("xdiv232")
		}
	}
	rexsult, err = lang.RxFromString("372567445.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("96.0992141"))
	if err != nil {
		lang.SayString("xdiv233")
	} else {
		if !(rexsult.ToString() == "3876904.18") {
			lang.SayString("xdiv233")
		}
	}
	rexsult, err = lang.RxFromString("802.156517").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-174409310.E-255338020"))
	if err != nil {
		lang.SayString("xdiv234")
	} else {
		if !(rexsult.ToString() == "-4.59927579E+255338014") {
			lang.SayString("xdiv234")
		}
	}
	rexsult, err = lang.RxFromString("-3.65207541").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("74501982.0"))
	if err != nil {
		lang.SayString("xdiv235")
	} else {
		if !(rexsult.ToString() == "-4.90198423E-8") {
			lang.SayString("xdiv235")
		}
	}
	rexsult, err = lang.RxFromString("-5297.76981").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-859.719404"))
	if err != nil {
		lang.SayString("xdiv236")
	} else {
		if !(rexsult.ToString() == "6.16220802") {
			lang.SayString("xdiv236")
		}
	}
	rexsult, err = lang.RxFromString("-684172.592").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("766.448597E+288361959"))
	if err != nil {
		lang.SayString("xdiv237")
	} else {
		if !(rexsult.ToString() == "-8.92652938E-288361957") {
			lang.SayString("xdiv237")
		}
	}
	rexsult, err = lang.RxFromString("626919.219").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("57469.8727E+13188610"))
	if err != nil {
		lang.SayString("xdiv238")
	} else {
		if !(rexsult.ToString() == "1.09086586E-13188609") {
			lang.SayString("xdiv238")
		}
	}
	rexsult, err = lang.RxFromString("-77480.5840").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("893265.594E+287982552"))
	if err != nil {
		lang.SayString("xdiv239")
	} else {
		if !(rexsult.ToString() == "-8.67385742E-287982554") {
			lang.SayString("xdiv239")
		}
	}
	rexsult, err = lang.RxFromString("-7177620.29").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7786343.83"))
	if err != nil {
		lang.SayString("xdiv240")
	} else {
		if !(rexsult.ToString() == "-0.921821647") {
			lang.SayString("xdiv240")
		}
	}
	rexsult, err = lang.RxFromString("9.6224130").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4.50355112"))
	if err != nil {
		lang.SayString("xdiv241")
	} else {
		if !(rexsult.ToString() == "2.13662791") {
			lang.SayString("xdiv241")
		}
	}
	rexsult, err = lang.RxFromString("-66.6337347E-597410086").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-818812885"))
	if err != nil {
		lang.SayString("xdiv242")
	} else {
		if !(rexsult.ToString() == "8.13784638E-597410094") {
			lang.SayString("xdiv242")
		}
	}
	rexsult, err = lang.RxFromString("65587553.7").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("600574.736"))
	if err != nil {
		lang.SayString("xdiv243")
	} else {
		if !(rexsult.ToString() == "109.20798") {
			lang.SayString("xdiv243")
		}
	}
	rexsult, err = lang.RxFromString("-32401.939").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-585200217."))
	if err != nil {
		lang.SayString("xdiv244")
	} else {
		if !(rexsult.ToString() == "0.0000553689798") {
			lang.SayString("xdiv244")
		}
	}
	rexsult, err = lang.RxFromString("69573.988").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-9.77003465E+740933668"))
	if err != nil {
		lang.SayString("xdiv245")
	} else {
		if !(rexsult.ToString() == "-7.12116082E-740933665") {
			lang.SayString("xdiv245")
		}
	}
	rexsult, err = lang.RxFromString("2362.06251").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-433149546.E-152643629"))
	if err != nil {
		lang.SayString("xdiv246")
	} else {
		if !(rexsult.ToString() == "-5.45322633E+152643623") {
			lang.SayString("xdiv246")
		}
	}
	rexsult, err = lang.RxFromString("-615.23488E+249953452").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-21437483.7"))
	if err != nil {
		lang.SayString("xdiv247")
	} else {
		if !(rexsult.ToString() == "2.8699025E+249953447") {
			lang.SayString("xdiv247")
		}
	}
	rexsult, err = lang.RxFromString("216741082.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("250290244"))
	if err != nil {
		lang.SayString("xdiv248")
	} else {
		if !(rexsult.ToString() == "0.86595897") {
			lang.SayString("xdiv248")
		}
	}
	rexsult, err = lang.RxFromString("-6364720.49").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5539245.64"))
	if err != nil {
		lang.SayString("xdiv249")
	} else {
		if !(rexsult.ToString() == "-1.14902297") {
			lang.SayString("xdiv249")
		}
	}
	rexsult, err = lang.RxFromString("-814599.475").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-14.5431191"))
	if err != nil {
		lang.SayString("xdiv250")
	} else {
		if !(rexsult.ToString() == "56012.7074") {
			lang.SayString("xdiv250")
		}
	}
	rexsult, err = lang.RxFromString("-877498.755").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("507408724E-168628106"))
	if err != nil {
		lang.SayString("xdiv251")
	} else {
		if !(rexsult.ToString() == "-1.72937262E+168628103") {
			lang.SayString("xdiv251")
		}
	}
	rexsult, err = lang.RxFromString("10634446.5E+475783861").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("50.7213056E+17807809"))
	if err != nil {
		lang.SayString("xdiv252")
	} else {
		if !(rexsult.ToString() == "2.09664289E+457976057") {
			lang.SayString("xdiv252")
		}
	}
	rexsult, err = lang.RxFromString("-162726.257E-597285918").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4391.54799"))
	if err != nil {
		lang.SayString("xdiv253")
	} else {
		if !(rexsult.ToString() == "3.70544185E-597285917") {
			lang.SayString("xdiv253")
		}
	}
	rexsult, err = lang.RxFromString("700354586.E-99856707").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7198.0493E+436250299"))
	if err != nil {
		lang.SayString("xdiv254")
	} else {
		if !(rexsult.ToString() == "9.72978312E-536107002") {
			lang.SayString("xdiv254")
		}
	}
	rexsult, err = lang.RxFromString("39617663E-463704664").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-895.290346"))
	if err != nil {
		lang.SayString("xdiv255")
	} else {
		if !(rexsult.ToString() == "-4.42511898E-463704660") {
			lang.SayString("xdiv255")
		}
	}
	rexsult, err = lang.RxFromString("5350882.59").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-36329829"))
	if err != nil {
		lang.SayString("xdiv256")
	} else {
		if !(rexsult.ToString() == "-0.147286204") {
			lang.SayString("xdiv256")
		}
	}
	rexsult, err = lang.RxFromString("91966.4084E+210382952").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("166740.46E-42001390"))
	if err != nil {
		lang.SayString("xdiv257")
	} else {
		if !(rexsult.ToString() == "5.51554244E+252384341") {
			lang.SayString("xdiv257")
		}
	}
	rexsult, err = lang.RxFromString("231899031.E-481759076").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("726.337100"))
	if err != nil {
		lang.SayString("xdiv258")
	} else {
		if !(rexsult.ToString() == "3.19271907E-481759071") {
			lang.SayString("xdiv258")
		}
	}
	rexsult, err = lang.RxFromString("-9611312.33").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("22109735.9"))
	if err != nil {
		lang.SayString("xdiv259")
	} else {
		if !(rexsult.ToString() == "-0.434709504") {
			lang.SayString("xdiv259")
		}
	}
	rexsult, err = lang.RxFromString("-5604938.15E-36812542").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("735937577."))
	if err != nil {
		lang.SayString("xdiv260")
	} else {
		if !(rexsult.ToString() == "-7.61605104E-36812545") {
			lang.SayString("xdiv260")
		}
	}
	rexsult, err = lang.RxFromString("693881413.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("260547224E-480281418"))
	if err != nil {
		lang.SayString("xdiv261")
	} else {
		if !(rexsult.ToString() == "2.66316947E+480281418") {
			lang.SayString("xdiv261")
		}
	}
	rexsult, err = lang.RxFromString("-34865.7378E-368768024").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2297117.88"))
	if err != nil {
		lang.SayString("xdiv262")
	} else {
		if !(rexsult.ToString() == "-1.5178036E-368768026") {
			lang.SayString("xdiv262")
		}
	}
	rexsult, err = lang.RxFromString("1123.32456").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7.86747918E+930888796"))
	if err != nil {
		lang.SayString("xdiv263")
	} else {
		if !(rexsult.ToString() == "1.42780748E-930888794") {
			lang.SayString("xdiv263")
		}
	}
	rexsult, err = lang.RxFromString("56.6607465E+467812565").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("909552512E+764516200"))
	if err != nil {
		lang.SayString("xdiv264")
	} else {
		if !(rexsult.ToString() == "6.22951899E-296703643") {
			lang.SayString("xdiv264")
		}
	}
	rexsult, err = lang.RxFromString("-1.85771840E+365552540").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-73028339.7"))
	if err != nil {
		lang.SayString("xdiv265")
	} else {
		if !(rexsult.ToString() == "2.54383217E+365552532") {
			lang.SayString("xdiv265")
		}
	}
	rexsult, err = lang.RxFromString("34.1935525").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-40767.6450"))
	if err != nil {
		lang.SayString("xdiv266")
	} else {
		if !(rexsult.ToString() == "-0.000838742402") {
			lang.SayString("xdiv266")
		}
	}
	rexsult, err = lang.RxFromString("26.0009168E+751618294").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-304019.929"))
	if err != nil {
		lang.SayString("xdiv267")
	} else {
		if !(rexsult.ToString() == "-8.5523725E+751618289") {
			lang.SayString("xdiv267")
		}
	}
	rexsult, err = lang.RxFromString("-58.4853072E+588540055").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4647.3205"))
	if err != nil {
		lang.SayString("xdiv268")
	} else {
		if !(rexsult.ToString() == "1.25847372E+588540053") {
			lang.SayString("xdiv268")
		}
	}
	rexsult, err = lang.RxFromString("51.025101").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4467691.57"))
	if err != nil {
		lang.SayString("xdiv269")
	} else {
		if !(rexsult.ToString() == "-0.0000114209095") {
			lang.SayString("xdiv269")
		}
	}
	rexsult, err = lang.RxFromString("-2214.76582").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("379785372E+223117572"))
	if err != nil {
		lang.SayString("xdiv270")
	} else {
		if !(rexsult.ToString() == "-5.83162487E-223117578") {
			lang.SayString("xdiv270")
		}
	}
	rexsult, err = lang.RxFromString("-2564.75207E-841443929").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-653498187"))
	if err != nil {
		lang.SayString("xdiv271")
	} else {
		if !(rexsult.ToString() == "3.92465063E-841443935") {
			lang.SayString("xdiv271")
		}
	}
	rexsult, err = lang.RxFromString("513115529.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("27775075.6E+217133352"))
	if err != nil {
		lang.SayString("xdiv272")
	} else {
		if !(rexsult.ToString() == "1.84739562E-217133351") {
			lang.SayString("xdiv272")
		}
	}
	rexsult, err = lang.RxFromString("-247157.208").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-532990.453"))
	if err != nil {
		lang.SayString("xdiv273")
	} else {
		if !(rexsult.ToString() == "0.46371789") {
			lang.SayString("xdiv273")
		}
	}
	rexsult, err = lang.RxFromString("40.2490764E-339482253").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7626.85442E+594264540"))
	if err != nil {
		lang.SayString("xdiv274")
	} else {
		if !(rexsult.ToString() == "5.27728395E-933746796") {
			lang.SayString("xdiv274")
		}
	}
	rexsult, err = lang.RxFromString("-1156008.8").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8870382.36"))
	if err != nil {
		lang.SayString("xdiv275")
	} else {
		if !(rexsult.ToString() == "0.130322319") {
			lang.SayString("xdiv275")
		}
	}
	rexsult, err = lang.RxFromString("880097928.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-52455011.1E+204538218"))
	if err != nil {
		lang.SayString("xdiv276")
	} else {
		if !(rexsult.ToString() == "-1.67781478E-204538217") {
			lang.SayString("xdiv276")
		}
	}
	rexsult, err = lang.RxFromString("5796.2524").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("34458329.7E+832129426"))
	if err != nil {
		lang.SayString("xdiv277")
	} else {
		if !(rexsult.ToString() == "1.68210486E-832129430") {
			lang.SayString("xdiv277")
		}
	}
	rexsult, err = lang.RxFromString("27.1000923E-218032223").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-45.0198341"))
	if err != nil {
		lang.SayString("xdiv278")
	} else {
		if !(rexsult.ToString() == "-6.01958955E-218032224") {
			lang.SayString("xdiv278")
		}
	}
	rexsult, err = lang.RxFromString("42643477.8").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("26118465E-730390549"))
	if err != nil {
		lang.SayString("xdiv279")
	} else {
		if !(rexsult.ToString() == "1.63269464E+730390549") {
			lang.SayString("xdiv279")
		}
	}
	rexsult, err = lang.RxFromString("-31918.9176E-163031657").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-21.5422824E-807317258"))
	if err != nil {
		lang.SayString("xdiv280")
	} else {
		if !(rexsult.ToString() == "1.4816869E+644285604") {
			lang.SayString("xdiv280")
		}
	}
	rexsult, err = lang.RxFromString("84224841.0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.62548255E+647087608"))
	if err != nil {
		lang.SayString("xdiv281")
	} else {
		if !(rexsult.ToString() == "3.20797565E-647087601") {
			lang.SayString("xdiv281")
		}
	}
	rexsult, err = lang.RxFromString("-64413698.9").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-6674.1055E-701047852"))
	if err != nil {
		lang.SayString("xdiv282")
	} else {
		if !(rexsult.ToString() == "9.65128569E+701047855") {
			lang.SayString("xdiv282")
		}
	}
	rexsult, err = lang.RxFromString("-62.5059208").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9.5795779E-898350012"))
	if err != nil {
		lang.SayString("xdiv283")
	} else {
		if !(rexsult.ToString() == "-6.52491388E+898350012") {
			lang.SayString("xdiv283")
		}
	}
	rexsult, err = lang.RxFromString("9090950.80").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("436.400932"))
	if err != nil {
		lang.SayString("xdiv284")
	} else {
		if !(rexsult.ToString() == "20831.6485") {
			lang.SayString("xdiv284")
		}
	}
	rexsult, err = lang.RxFromString("-89833825.7E+329205393").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-779430.194"))
	if err != nil {
		lang.SayString("xdiv285")
	} else {
		if !(rexsult.ToString() == "1.15255768E+329205395") {
			lang.SayString("xdiv285")
		}
	}
	rexsult, err = lang.RxFromString("-714562.019E+750205688").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("704079764"))
	if err != nil {
		lang.SayString("xdiv286")
	} else {
		if !(rexsult.ToString() == "-1.01488788E+750205685") {
			lang.SayString("xdiv286")
		}
	}
	rexsult, err = lang.RxFromString("-584537670.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("31139.7737E-146687560"))
	if err != nil {
		lang.SayString("xdiv287")
	} else {
		if !(rexsult.ToString() == "-1.87714168E+146687564") {
			lang.SayString("xdiv287")
		}
	}
	rexsult, err = lang.RxFromString("-4.18074650E-858746879").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("571035.277E-279409165"))
	if err != nil {
		lang.SayString("xdiv288")
	} else {
		if !(rexsult.ToString() == "-7.3213454E-579337720") {
			lang.SayString("xdiv288")
		}
	}
	rexsult, err = lang.RxFromString("5.15309635").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-695649.219E+451948183"))
	if err != nil {
		lang.SayString("xdiv289")
	} else {
		if !(rexsult.ToString() == "-7.40760747E-451948189") {
			lang.SayString("xdiv289")
		}
	}
	rexsult, err = lang.RxFromString("-940030153.E+83797657").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4.11510193"))
	if err != nil {
		lang.SayString("xdiv290")
	} else {
		if !(rexsult.ToString() == "2.28434233E+83797665") {
			lang.SayString("xdiv290")
		}
	}
	rexsult, err = lang.RxFromString("89088.9683E+587739290").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.31932110"))
	if err != nil {
		lang.SayString("xdiv291")
	} else {
		if !(rexsult.ToString() == "6.75263727E+587739294") {
			lang.SayString("xdiv291")
		}
	}
	rexsult, err = lang.RxFromString("3336750").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("6.47961126"))
	if err != nil {
		lang.SayString("xdiv292")
	} else {
		if !(rexsult.ToString() == "514961.448") {
			lang.SayString("xdiv292")
		}
	}
	rexsult, err = lang.RxFromString("904654622.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("692065270.E+329081915"))
	if err != nil {
		lang.SayString("xdiv293")
	} else {
		if !(rexsult.ToString() == "1.30718107E-329081915") {
			lang.SayString("xdiv293")
		}
	}
	rexsult, err = lang.RxFromString("304804380").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4681.23698"))
	if err != nil {
		lang.SayString("xdiv294")
	} else {
		if !(rexsult.ToString() == "-65111.9312") {
			lang.SayString("xdiv294")
		}
	}
	rexsult, err = lang.RxFromString("674.55569").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-82981.2684E+852890752"))
	if err != nil {
		lang.SayString("xdiv295")
	} else {
		if !(rexsult.ToString() == "-8.12901156E-852890755") {
			lang.SayString("xdiv295")
		}
	}
	rexsult, err = lang.RxFromString("-5111.51025E-108006096").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5448870.4E+279212255"))
	if err != nil {
		lang.SayString("xdiv296")
	} else {
		if !(rexsult.ToString() == "-9.38086222E-387218355") {
			lang.SayString("xdiv296")
		}
	}
	rexsult, err = lang.RxFromString("-2623.45068").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-466463938."))
	if err != nil {
		lang.SayString("xdiv297")
	} else {
		if !(rexsult.ToString() == "0.00000562412325") {
			lang.SayString("xdiv297")
		}
	}
	rexsult, err = lang.RxFromString("299350.435").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3373.33551"))
	if err != nil {
		lang.SayString("xdiv298")
	} else {
		if !(rexsult.ToString() == "88.7401903") {
			lang.SayString("xdiv298")
		}
	}
	rexsult, err = lang.RxFromString("-6589947.80").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2448.75933E-591549734"))
	if err != nil {
		lang.SayString("xdiv299")
	} else {
		if !(rexsult.ToString() == "2.69113739E+591549737") {
			lang.SayString("xdiv299")
		}
	}
	rexsult, err = lang.RxFromString("3774.5358E-491090520").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("173.060090"))
	if err != nil {
		lang.SayString("xdiv300")
	} else {
		if !(rexsult.ToString() == "2.18105503E-491090519") {
			lang.SayString("xdiv300")
		}
	}
	rexsult, err = lang.RxFromString("-13.6783690").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-453.610117"))
	if err != nil {
		lang.SayString("xdiv301")
	} else {
		if !(rexsult.ToString() == "0.0301544619") {
			lang.SayString("xdiv301")
		}
	}
	rexsult, err = lang.RxFromString("-990100927.E-615244634").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("223801.421E+247632618"))
	if err != nil {
		lang.SayString("xdiv302")
	} else {
		if !(rexsult.ToString() == "-4.42401537E-862877249") {
			lang.SayString("xdiv302")
		}
	}
	rexsult, err = lang.RxFromString("1275.10292").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-667965353"))
	if err != nil {
		lang.SayString("xdiv303")
	} else {
		if !(rexsult.ToString() == "-0.00000190893572") {
			lang.SayString("xdiv303")
		}
	}
	rexsult, err = lang.RxFromString("-8.76375480E-596792197").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("992.077361"))
	if err != nil {
		lang.SayString("xdiv304")
	} else {
		if !(rexsult.ToString() == "-8.83374134E-596792200") {
			lang.SayString("xdiv304")
		}
	}
	rexsult, err = lang.RxFromString("953.976935E+385444720").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("96503.3378"))
	if err != nil {
		lang.SayString("xdiv305")
	} else {
		if !(rexsult.ToString() == "9.88542942E+385444717") {
			lang.SayString("xdiv305")
		}
	}
	rexsult, err = lang.RxFromString("213577152").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-986710073E+31900046"))
	if err != nil {
		lang.SayString("xdiv306")
	} else {
		if !(rexsult.ToString() == "-2.16453807E-31900047") {
			lang.SayString("xdiv306")
		}
	}
	rexsult, err = lang.RxFromString("91393.9398E-323439228").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-135.701000"))
	if err != nil {
		lang.SayString("xdiv307")
	} else {
		if !(rexsult.ToString() == "-6.73494962E-323439226") {
			lang.SayString("xdiv307")
		}
	}
	rexsult, err = lang.RxFromString("-396.503557").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("45757264.E-254363788"))
	if err != nil {
		lang.SayString("xdiv308")
	} else {
		if !(rexsult.ToString() == "-8.66536856E+254363782") {
			lang.SayString("xdiv308")
		}
	}
	rexsult, err = lang.RxFromString("59807846.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.53345254"))
	if err != nil {
		lang.SayString("xdiv309")
	} else {
		if !(rexsult.ToString() == "39002084.9") {
			lang.SayString("xdiv309")
		}
	}
	rexsult, err = lang.RxFromString("-8046158.45").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8.3635397"))
	if err != nil {
		lang.SayString("xdiv310")
	} else {
		if !(rexsult.ToString() == "-962051.803") {
			lang.SayString("xdiv310")
		}
	}
	rexsult, err = lang.RxFromString("55.1123381E+50627250").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-94.0355047E-162540316"))
	if err != nil {
		lang.SayString("xdiv311")
	} else {
		if !(rexsult.ToString() == "-5.86080101E+213167565") {
			lang.SayString("xdiv311")
		}
	}
	rexsult, err = lang.RxFromString("-948.038054").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3580.84510"))
	if err != nil {
		lang.SayString("xdiv312")
	} else {
		if !(rexsult.ToString() == "-0.264752601") {
			lang.SayString("xdiv312")
		}
	}
	rexsult, err = lang.RxFromString("-6026.42752").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-14.2286406E-334921364"))
	if err != nil {
		lang.SayString("xdiv313")
	} else {
		if !(rexsult.ToString() == "4.23542044E+334921366") {
			lang.SayString("xdiv313")
		}
	}
	rexsult, err = lang.RxFromString("79551.5014").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-538.186229"))
	if err != nil {
		lang.SayString("xdiv314")
	} else {
		if !(rexsult.ToString() == "-147.814078") {
			lang.SayString("xdiv314")
		}
	}
	rexsult, err = lang.RxFromString("42706056.E+623578292").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-690.327745"))
	if err != nil {
		lang.SayString("xdiv315")
	} else {
		if !(rexsult.ToString() == "-6.18634501E+623578296") {
			lang.SayString("xdiv315")
		}
	}
	rexsult, err = lang.RxFromString("2454136.08E+502374077").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("856268.795E-356664934"))
	if err != nil {
		lang.SayString("xdiv316")
	} else {
		if !(rexsult.ToString() == "2.86608142E+859039011") {
			lang.SayString("xdiv316")
		}
	}
	rexsult, err = lang.RxFromString("-3264204.54").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-42704.501"))
	if err != nil {
		lang.SayString("xdiv317")
	} else {
		if !(rexsult.ToString() == "76.437014") {
			lang.SayString("xdiv317")
		}
	}
	rexsult, err = lang.RxFromString("1.21265492").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("44102.6073"))
	if err != nil {
		lang.SayString("xdiv318")
	} else {
		if !(rexsult.ToString() == "0.0000274962183") {
			lang.SayString("xdiv318")
		}
	}
	rexsult, err = lang.RxFromString("-19.054711E+975514652").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-22144.0822"))
	if err != nil {
		lang.SayString("xdiv319")
	} else {
		if !(rexsult.ToString() == "8.60487729E+975514648") {
			lang.SayString("xdiv319")
		}
	}
	rexsult, err = lang.RxFromString("745.78452").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-1922.00670E+375923302"))
	if err != nil {
		lang.SayString("xdiv320")
	} else {
		if !(rexsult.ToString() == "-3.88023892E-375923303") {
			lang.SayString("xdiv320")
		}
	}
	rexsult, err = lang.RxFromString("-963717836").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-823989308"))
	if err != nil {
		lang.SayString("xdiv321")
	} else {
		if !(rexsult.ToString() == "1.16957566") {
			lang.SayString("xdiv321")
		}
	}
	rexsult, err = lang.RxFromString("82.4185291E-321919303").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-215747737.E-995147400"))
	if err != nil {
		lang.SayString("xdiv322")
	} else {
		if !(rexsult.ToString() == "-3.82013412E+673228090") {
			lang.SayString("xdiv322")
		}
	}
	rexsult, err = lang.RxFromString("-808328.607E-790810342").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("53075.7082"))
	if err != nil {
		lang.SayString("xdiv323")
	} else {
		if !(rexsult.ToString() == "-1.52297281E-790810341") {
			lang.SayString("xdiv323")
		}
	}
	rexsult, err = lang.RxFromString("700592.720").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-698485.085"))
	if err != nil {
		lang.SayString("xdiv324")
	} else {
		if !(rexsult.ToString() == "-1.00301744") {
			lang.SayString("xdiv324")
		}
	}
	rexsult, err = lang.RxFromString("-80273928.0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("661346.239"))
	if err != nil {
		lang.SayString("xdiv325")
	} else {
		if !(rexsult.ToString() == "-121.379579") {
			lang.SayString("xdiv325")
		}
	}
	rexsult, err = lang.RxFromString("-24018251.0E+819786764").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("59141.9600E-167165065"))
	if err != nil {
		lang.SayString("xdiv326")
	} else {
		if !(rexsult.ToString() == "-4.06111854E+986951831") {
			lang.SayString("xdiv326")
		}
	}
	rexsult, err = lang.RxFromString("2512953.3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3769170.35E-993621645"))
	if err != nil {
		lang.SayString("xdiv327")
	} else {
		if !(rexsult.ToString() == "-6.66712583E+993621644") {
			lang.SayString("xdiv327")
		}
	}
	rexsult, err = lang.RxFromString("-682.796370").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("71131.0224"))
	if err != nil {
		lang.SayString("xdiv328")
	} else {
		if !(rexsult.ToString() == "-0.00959913617") {
			lang.SayString("xdiv328")
		}
	}
	rexsult, err = lang.RxFromString("89.9997490").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4993.69831"))
	if err != nil {
		lang.SayString("xdiv329")
	} else {
		if !(rexsult.ToString() == "-0.0180226644") {
			lang.SayString("xdiv329")
		}
	}
	rexsult, err = lang.RxFromString("76563354.6E-112338836").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("278271.585E-511481095"))
	if err != nil {
		lang.SayString("xdiv330")
	} else {
		if !(rexsult.ToString() == "2.7513896E+399142261") {
			lang.SayString("xdiv330")
		}
	}
	rexsult, err = lang.RxFromString("-932499.010").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("873.377701E-502190452"))
	if err != nil {
		lang.SayString("xdiv331")
	} else {
		if !(rexsult.ToString() == "-1.06769272E+502190455") {
			lang.SayString("xdiv331")
		}
	}
	rexsult, err = lang.RxFromString("-7735918.21E+799514797").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7748.78023"))
	if err != nil {
		lang.SayString("xdiv332")
	} else {
		if !(rexsult.ToString() == "9.98340123E+799514799") {
			lang.SayString("xdiv332")
		}
	}
	rexsult, err = lang.RxFromString("-5205124.44E-140588661").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-495394029.E-620856313"))
	if err != nil {
		lang.SayString("xdiv334")
	} else {
		if !(rexsult.ToString() == "1.05070391E+480267650") {
			lang.SayString("xdiv334")
		}
	}
	rexsult, err = lang.RxFromString("-8868.72074").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5592399.93"))
	if err != nil {
		lang.SayString("xdiv335")
	} else {
		if !(rexsult.ToString() == "-0.00158585238") {
			lang.SayString("xdiv335")
		}
	}
	rexsult, err = lang.RxFromString("-74.7852037E-175205809").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("4.14316542"))
	if err != nil {
		lang.SayString("xdiv336")
	} else {
		if !(rexsult.ToString() == "-1.80502577E-175205808") {
			lang.SayString("xdiv336")
		}
	}
	rexsult, err = lang.RxFromString("84196.1091E+242628748").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8.07523036E-288231467"))
	if err != nil {
		lang.SayString("xdiv337")
	} else {
		if !(rexsult.ToString() == "1.04264653E+530860219") {
			lang.SayString("xdiv337")
		}
	}
	rexsult, err = lang.RxFromString("38660103.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-6671.73085E+900998477"))
	if err != nil {
		lang.SayString("xdiv338")
	} else {
		if !(rexsult.ToString() == "-5.79461372E-900998474") {
			lang.SayString("xdiv338")
		}
	}
	rexsult, err = lang.RxFromString("-52.2659460").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-296404199E+372050476"))
	if err != nil {
		lang.SayString("xdiv339")
	} else {
		if !(rexsult.ToString() == "1.76333352E-372050483") {
			lang.SayString("xdiv339")
		}
	}
	rexsult, err = lang.RxFromString("6.06625013").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-276.359186"))
	if err != nil {
		lang.SayString("xdiv340")
	} else {
		if !(rexsult.ToString() == "-0.0219506007") {
			lang.SayString("xdiv340")
		}
	}
	rexsult, err = lang.RxFromString("-62971617.5E-241444744").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("46266799.3"))
	if err != nil {
		lang.SayString("xdiv341")
	} else {
		if !(rexsult.ToString() == "-1.36105411E-241444744") {
			lang.SayString("xdiv341")
		}
	}
	rexsult, err = lang.RxFromString("-5.36917800").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-311124593.E-976066491"))
	if err != nil {
		lang.SayString("xdiv342")
	} else {
		if !(rexsult.ToString() == "1.72573243E+976066483") {
			lang.SayString("xdiv342")
		}
	}
	rexsult, err = lang.RxFromString("2467915.01").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-92.5558322"))
	if err != nil {
		lang.SayString("xdiv343")
	} else {
		if !(rexsult.ToString() == "-26664.0681") {
			lang.SayString("xdiv343")
		}
	}
	rexsult, err = lang.RxFromString("187.232671").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-840.469347"))
	if err != nil {
		lang.SayString("xdiv344")
	} else {
		if !(rexsult.ToString() == "-0.222771564") {
			lang.SayString("xdiv344")
		}
	}
	rexsult, err = lang.RxFromString("81233.6823").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-5192.21666E+309315093"))
	if err != nil {
		lang.SayString("xdiv345")
	} else {
		if !(rexsult.ToString() == "-1.56452798E-309315092") {
			lang.SayString("xdiv345")
		}
	}
	rexsult, err = lang.RxFromString("-854.586113").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-79.8715762E-853065103"))
	if err != nil {
		lang.SayString("xdiv346")
	} else {
		if !(rexsult.ToString() == "1.06995023E+853065104") {
			lang.SayString("xdiv346")
		}
	}
	rexsult, err = lang.RxFromString("78872665.3").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("172.102119"))
	if err != nil {
		lang.SayString("xdiv347")
	} else {
		if !(rexsult.ToString() == "458289.914") {
			lang.SayString("xdiv347")
		}
	}
	rexsult, err = lang.RxFromString("328268.1E-436315617").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-204.522245"))
	if err != nil {
		lang.SayString("xdiv348")
	} else {
		if !(rexsult.ToString() == "-1.60504839E-436315614") {
			lang.SayString("xdiv348")
		}
	}
	rexsult, err = lang.RxFromString("-4037911.02E+641367645").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("29.5713010"))
	if err != nil {
		lang.SayString("xdiv349")
	} else {
		if !(rexsult.ToString() == "-1.36548305E+641367650") {
			lang.SayString("xdiv349")
		}
	}
	rexsult, err = lang.RxFromString("-5.47345502").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("59818.7580"))
	if err != nil {
		lang.SayString("xdiv351")
	} else {
		if !(rexsult.ToString() == "-0.0000915006463") {
			lang.SayString("xdiv351")
		}
	}
	rexsult, err = lang.RxFromString("563891620E-361354567").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-845900362."))
	if err != nil {
		lang.SayString("xdiv352")
	} else {
		if !(rexsult.ToString() == "-6.66617069E-361354568") {
			lang.SayString("xdiv352")
		}
	}
	rexsult, err = lang.RxFromString("-69.7231286").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("85773.7504"))
	if err != nil {
		lang.SayString("xdiv353")
	} else {
		if !(rexsult.ToString() == "-0.000812872566") {
			lang.SayString("xdiv353")
		}
	}
	rexsult, err = lang.RxFromString("5125.51188").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("73814638.4E-500934741"))
	if err != nil {
		lang.SayString("xdiv354")
	} else {
		if !(rexsult.ToString() == "6.94376074E+500934736") {
			lang.SayString("xdiv354")
		}
	}
	rexsult, err = lang.RxFromString("-54.6254096").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-332921899."))
	if err != nil {
		lang.SayString("xdiv355")
	} else {
		if !(rexsult.ToString() == "1.64078752E-7") {
			lang.SayString("xdiv355")
		}
	}
	rexsult, err = lang.RxFromString("-9.04778095E-591874079").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8719.40286"))
	if err != nil {
		lang.SayString("xdiv356")
	} else {
		if !(rexsult.ToString() == "-1.03766062E-591874082") {
			lang.SayString("xdiv356")
		}
	}
	rexsult, err = lang.RxFromString("-21006.1733E+884684431").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-48872.9175"))
	if err != nil {
		lang.SayString("xdiv357")
	} else {
		if !(rexsult.ToString() == "4.29812141E+884684430") {
			lang.SayString("xdiv357")
		}
	}
	rexsult, err = lang.RxFromString("-1546783").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-51935370.4"))
	if err != nil {
		lang.SayString("xdiv358")
	} else {
		if !(rexsult.ToString() == "0.0297828433") {
			lang.SayString("xdiv358")
		}
	}
	rexsult, err = lang.RxFromString("61302486.8").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("205.490417"))
	if err != nil {
		lang.SayString("xdiv359")
	} else {
		if !(rexsult.ToString() == "298322.85") {
			lang.SayString("xdiv359")
		}
	}
	rexsult, err = lang.RxFromString("-318180109.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-54008744.6E-170931002"))
	if err != nil {
		lang.SayString("xdiv360")
	} else {
		if !(rexsult.ToString() == "5.89127023E+170931002") {
			lang.SayString("xdiv360")
		}
	}
	rexsult, err = lang.RxFromString("-28486137.1E+901441714").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-42454.940"))
	if err != nil {
		lang.SayString("xdiv361")
	} else {
		if !(rexsult.ToString() == "6.70973439E+901441716") {
			lang.SayString("xdiv361")
		}
	}
	rexsult, err = lang.RxFromString("-546398328.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-27.9149712"))
	if err != nil {
		lang.SayString("xdiv362")
	} else {
		if !(rexsult.ToString() == "19573666.2") {
			lang.SayString("xdiv362")
		}
	}
	rexsult, err = lang.RxFromString("5402066.1E-284978216").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("622.751128"))
	if err != nil {
		lang.SayString("xdiv363")
	} else {
		if !(rexsult.ToString() == "8.67451837E-284978213") {
			lang.SayString("xdiv363")
		}
	}
	rexsult, err = lang.RxFromString("18845620").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3129.43753"))
	if err != nil {
		lang.SayString("xdiv364")
	} else {
		if !(rexsult.ToString() == "6022.04704") {
			lang.SayString("xdiv364")
		}
	}
	rexsult, err = lang.RxFromString("50707.1412E+912475670").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-198098.186E+701407524"))
	if err != nil {
		lang.SayString("xdiv365")
	} else {
		if !(rexsult.ToString() == "-2.5596974E+211068145") {
			lang.SayString("xdiv365")
		}
	}
	rexsult, err = lang.RxFromString("55.8245006E+928885991").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("99170843.9E-47402167"))
	if err != nil {
		lang.SayString("xdiv366")
	} else {
		if !(rexsult.ToString() == "5.62912429E+976288151") {
			lang.SayString("xdiv366")
		}
	}
	rexsult, err = lang.RxFromString("13.8003883E-386224921").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-84126481.9E-296378341"))
	if err != nil {
		lang.SayString("xdiv367")
	} else {
		if !(rexsult.ToString() == "-1.64043331E-89846587") {
			lang.SayString("xdiv367")
		}
	}
	rexsult, err = lang.RxFromString("9820.90457").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("46671.5915"))
	if err != nil {
		lang.SayString("xdiv368")
	} else {
		if !(rexsult.ToString() == "0.210425748") {
			lang.SayString("xdiv368")
		}
	}
	rexsult, err = lang.RxFromString("7.22436006E+831949153").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-11168830E+322331045"))
	if err != nil {
		lang.SayString("xdiv369")
	} else {
		if !(rexsult.ToString() == "-6.46832306E+509618101") {
			lang.SayString("xdiv369")
		}
	}
	rexsult, err = lang.RxFromString("472648900").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-207.784153"))
	if err != nil {
		lang.SayString("xdiv370")
	} else {
		if !(rexsult.ToString() == "-2274711.01") {
			lang.SayString("xdiv370")
		}
	}
	rexsult, err = lang.RxFromString("-8754.49306").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-818.165153E+631475457"))
	if err != nil {
		lang.SayString("xdiv371")
	} else {
		if !(rexsult.ToString() == "1.07001539E-631475456") {
			lang.SayString("xdiv371")
		}
	}
	rexsult, err = lang.RxFromString("98750864").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("191380.551"))
	if err != nil {
		lang.SayString("xdiv372")
	} else {
		if !(rexsult.ToString() == "515.992161") {
			lang.SayString("xdiv372")
		}
	}
	rexsult, err = lang.RxFromString("725292561.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-768963606.E+340762986"))
	if err != nil {
		lang.SayString("xdiv373")
	} else {
		if !(rexsult.ToString() == "-9.43207917E-340762987") {
			lang.SayString("xdiv373")
		}
	}
	rexsult, err = lang.RxFromString("1862.80445").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("648254483."))
	if err != nil {
		lang.SayString("xdiv374")
	} else {
		if !(rexsult.ToString() == "0.00000287356972") {
			lang.SayString("xdiv374")
		}
	}
	rexsult, err = lang.RxFromString("-5549320.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-93580684.1"))
	if err != nil {
		lang.SayString("xdiv375")
	} else {
		if !(rexsult.ToString() == "0.0592998454") {
			lang.SayString("xdiv375")
		}
	}
	rexsult, err = lang.RxFromString("-14677053.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-25784.7358"))
	if err != nil {
		lang.SayString("xdiv376")
	} else {
		if !(rexsult.ToString() == "569.214795") {
			lang.SayString("xdiv376")
		}
	}
	rexsult, err = lang.RxFromString("547402.308E+571687617").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7835797.01E+500067364"))
	if err != nil {
		lang.SayString("xdiv377")
	} else {
		if !(rexsult.ToString() == "-6.98591742E+71620251") {
			lang.SayString("xdiv377")
		}
	}
	rexsult, err = lang.RxFromString("-4131738.09").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7579.07566"))
	if err != nil {
		lang.SayString("xdiv378")
	} else {
		if !(rexsult.ToString() == "-545.150659") {
			lang.SayString("xdiv378")
		}
	}
	rexsult, err = lang.RxFromString("504544.648").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7678.96133E-662143268"))
	if err != nil {
		lang.SayString("xdiv379")
	} else {
		if !(rexsult.ToString() == "-6.57048039E+662143269") {
			lang.SayString("xdiv379")
		}
	}
	rexsult, err = lang.RxFromString("829898241").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8912.99114E+929228149"))
	if err != nil {
		lang.SayString("xdiv380")
	} else {
		if !(rexsult.ToString() == "9.31110811E-929228145") {
			lang.SayString("xdiv380")
		}
	}
	rexsult, err = lang.RxFromString("53.6891691").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-11.2371140"))
	if err != nil {
		lang.SayString("xdiv381")
	} else {
		if !(rexsult.ToString() == "-4.77784323") {
			lang.SayString("xdiv381")
		}
	}
	rexsult, err = lang.RxFromString("-93951823.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-25317.8645"))
	if err != nil {
		lang.SayString("xdiv382")
	} else {
		if !(rexsult.ToString() == "3710.89052") {
			lang.SayString("xdiv382")
		}
	}
	rexsult, err = lang.RxFromString("446919.123").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("951338490."))
	if err != nil {
		lang.SayString("xdiv383")
	} else {
		if !(rexsult.ToString() == "0.000469779293") {
			lang.SayString("xdiv383")
		}
	}
	rexsult, err = lang.RxFromString("-8.01787748").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-88.3076852"))
	if err != nil {
		lang.SayString("xdiv384")
	} else {
		if !(rexsult.ToString() == "0.0907947871") {
			lang.SayString("xdiv384")
		}
	}
	rexsult, err = lang.RxFromString("517458139").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-999731.548"))
	if err != nil {
		lang.SayString("xdiv385")
	} else {
		if !(rexsult.ToString() == "-517.597089") {
			lang.SayString("xdiv385")
		}
	}
	rexsult, err = lang.RxFromString("-405543440").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4013.18295"))
	if err != nil {
		lang.SayString("xdiv386")
	} else {
		if !(rexsult.ToString() == "101052.816") {
			lang.SayString("xdiv386")
		}
	}
	rexsult, err = lang.RxFromString("-49245250.1E+682760825").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-848776.637"))
	if err != nil {
		lang.SayString("xdiv387")
	} else {
		if !(rexsult.ToString() == "5.80190924E+682760826") {
			lang.SayString("xdiv387")
		}
	}
	rexsult, err = lang.RxFromString("-151144455").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-170371.29"))
	if err != nil {
		lang.SayString("xdiv388")
	} else {
		if !(rexsult.ToString() == "887.147447") {
			lang.SayString("xdiv388")
		}
	}
	rexsult, err = lang.RxFromString("-729236746.E+662737067").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9.10823602"))
	if err != nil {
		lang.SayString("xdiv389")
	} else {
		if !(rexsult.ToString() == "-8.00634442E+662737074") {
			lang.SayString("xdiv389")
		}
	}
	rexsult, err = lang.RxFromString("534.394729").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2369839.37"))
	if err != nil {
		lang.SayString("xdiv390")
	} else {
		if !(rexsult.ToString() == "-0.000225498291") {
			lang.SayString("xdiv390")
		}
	}
	rexsult, err = lang.RxFromString("89100.1797").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("224.370309"))
	if err != nil {
		lang.SayString("xdiv391")
	} else {
		if !(rexsult.ToString() == "397.112167") {
			lang.SayString("xdiv391")
		}
	}
	rexsult, err = lang.RxFromString("-821377.777").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("38.552821"))
	if err != nil {
		lang.SayString("xdiv392")
	} else {
		if !(rexsult.ToString() == "-21305.2575") {
			lang.SayString("xdiv392")
		}
	}
	rexsult, err = lang.RxFromString("-392640.782").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2571619.5E+113237865"))
	if err != nil {
		lang.SayString("xdiv393")
	} else {
		if !(rexsult.ToString() == "1.52682301E-113237866") {
			lang.SayString("xdiv393")
		}
	}
	rexsult, err = lang.RxFromString("-651397.712").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-723.59673"))
	if err != nil {
		lang.SayString("xdiv394")
	} else {
		if !(rexsult.ToString() == "900.222023") {
			lang.SayString("xdiv394")
		}
	}
	rexsult, err = lang.RxFromString("86.6890892").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("940380864"))
	if err != nil {
		lang.SayString("xdiv395")
	} else {
		if !(rexsult.ToString() == "9.21850843E-8") {
			lang.SayString("xdiv395")
		}
	}
	rexsult, err = lang.RxFromString("4880.06442E-382222621").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-115627239E-912834031"))
	if err != nil {
		lang.SayString("xdiv396")
	} else {
		if !(rexsult.ToString() == "-4.22051453E+530611405") {
			lang.SayString("xdiv396")
		}
	}
	rexsult, err = lang.RxFromString("173398265E-532383158").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3462.51450E+80986915"))
	if err != nil {
		lang.SayString("xdiv397")
	} else {
		if !(rexsult.ToString() == "5.00787116E-613370069") {
			lang.SayString("xdiv397")
		}
	}
	rexsult, err = lang.RxFromString("-1522176.78").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-6631061.77"))
	if err != nil {
		lang.SayString("xdiv398")
	} else {
		if !(rexsult.ToString() == "0.229552496") {
			lang.SayString("xdiv398")
		}
	}
	rexsult, err = lang.RxFromString("538.10453").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("522934310"))
	if err != nil {
		lang.SayString("xdiv399")
	} else {
		if !(rexsult.ToString() == "0.0000010290098") {
			lang.SayString("xdiv399")
		}
	}
	rexsult, err = lang.RxFromString("880243.444E-750940977").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-354.601578E-204943740"))
	if err != nil {
		lang.SayString("xdiv400")
	} else {
		if !(rexsult.ToString() == "-2.48234497E-545997234") {
			lang.SayString("xdiv400")
		}
	}
	rexsult, err = lang.RxFromString("968370.780").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("6677268.73"))
	if err != nil {
		lang.SayString("xdiv401")
	} else {
		if !(rexsult.ToString() == "0.145024982") {
			lang.SayString("xdiv401")
		}
	}
	rexsult, err = lang.RxFromString("-97.7474945").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("31248241.5"))
	if err != nil {
		lang.SayString("xdiv402")
	} else {
		if !(rexsult.ToString() == "-0.00000312809585") {
			lang.SayString("xdiv402")
		}
	}
	rexsult, err = lang.RxFromString("-187582786.E+369916952").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("957840435E+744365567"))
	if err != nil {
		lang.SayString("xdiv403")
	} else {
		if !(rexsult.ToString() == "-1.95839285E-374448616") {
			lang.SayString("xdiv403")
		}
	}
	rexsult, err = lang.RxFromString("-328026144.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-125499735."))
	if err != nil {
		lang.SayString("xdiv404")
	} else {
		if !(rexsult.ToString() == "2.61375965") {
			lang.SayString("xdiv404")
		}
	}
	rexsult, err = lang.RxFromString("-99424150.2E-523662102").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("3712.35030"))
	if err != nil {
		lang.SayString("xdiv405")
	} else {
		if !(rexsult.ToString() == "-2.67819958E-523662098") {
			lang.SayString("xdiv405")
		}
	}
	rexsult, err = lang.RxFromString("14838.0718").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9489893.28E+830631266"))
	if err != nil {
		lang.SayString("xdiv406")
	} else {
		if !(rexsult.ToString() == "1.56356572E-830631269") {
			lang.SayString("xdiv406")
		}
	}
	rexsult, err = lang.RxFromString("71207472.8").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-31835.0809"))
	if err != nil {
		lang.SayString("xdiv407")
	} else {
		if !(rexsult.ToString() == "-2236.76117") {
			lang.SayString("xdiv407")
		}
	}
	rexsult, err = lang.RxFromString("-20440.4394").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-44.4064328E+511085806"))
	if err != nil {
		lang.SayString("xdiv408")
	} else {
		if !(rexsult.ToString() == "4.60303567E-511085804") {
			lang.SayString("xdiv408")
		}
	}
	rexsult, err = lang.RxFromString("-54.3684171E-807210192").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.04592973E-984041807"))
	if err != nil {
		lang.SayString("xdiv409")
	} else {
		if !(rexsult.ToString() == "-5.19809463E+176831616") {
			lang.SayString("xdiv409")
		}
	}
	rexsult, err = lang.RxFromString("54310060.5E+948159739").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("274320701.E+205880484"))
	if err != nil {
		lang.SayString("xdiv410")
	} else {
		if !(rexsult.ToString() == "1.97980175E+742279254") {
			lang.SayString("xdiv410")
		}
	}
	rexsult, err = lang.RxFromString("-657.186702").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("426844.39"))
	if err != nil {
		lang.SayString("xdiv411")
	} else {
		if !(rexsult.ToString() == "-0.00153964001") {
			lang.SayString("xdiv411")
		}
	}
	rexsult, err = lang.RxFromString("-41593077.0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-688607.564"))
	if err != nil {
		lang.SayString("xdiv412")
	} else {
		if !(rexsult.ToString() == "60.4017138") {
			lang.SayString("xdiv412")
		}
	}
	rexsult, err = lang.RxFromString("-5786.38132").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("190556652.E+177045877"))
	if err != nil {
		lang.SayString("xdiv413")
	} else {
		if !(rexsult.ToString() == "-3.03656748E-177045882") {
			lang.SayString("xdiv413")
		}
	}
	rexsult, err = lang.RxFromString("737622.974").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-241560693E+249506565"))
	if err != nil {
		lang.SayString("xdiv414")
	} else {
		if !(rexsult.ToString() == "-3.05357202E-249506568") {
			lang.SayString("xdiv414")
		}
	}
	rexsult, err = lang.RxFromString("5615373.52").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7.27583808E-949781048"))
	if err != nil {
		lang.SayString("xdiv415")
	} else {
		if !(rexsult.ToString() == "-7.71783739E+949781053") {
			lang.SayString("xdiv415")
		}
	}
	rexsult, err = lang.RxFromString("644136.179").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-835708.103"))
	if err != nil {
		lang.SayString("xdiv416")
	} else {
		if !(rexsult.ToString() == "-0.770766942") {
			lang.SayString("xdiv416")
		}
	}
	rexsult, err = lang.RxFromString("-307.419521E+466861843").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-738689976.E-199032711"))
	if err != nil {
		lang.SayString("xdiv417")
	} else {
		if !(rexsult.ToString() == "4.16168529E+665894547") {
			lang.SayString("xdiv417")
		}
	}
	rexsult, err = lang.RxFromString("-619642.130").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-226740537.E-902590153"))
	if err != nil {
		lang.SayString("xdiv418")
	} else {
		if !(rexsult.ToString() == "2.73282466E+902590150") {
			lang.SayString("xdiv418")
		}
	}
	rexsult, err = lang.RxFromString("-31068.7549").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3.41495042E+86001379"))
	if err != nil {
		lang.SayString("xdiv419")
	} else {
		if !(rexsult.ToString() == "9.09786412E-86001376") {
			lang.SayString("xdiv419")
		}
	}
	rexsult, err = lang.RxFromString("-68951173.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-211804977.E-97318126"))
	if err != nil {
		lang.SayString("xdiv420")
	} else {
		if !(rexsult.ToString() == "3.25540854E+97318125") {
			lang.SayString("xdiv420")
		}
	}
	rexsult, err = lang.RxFromString("-4.09492571E-301749490").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("434.20199E-749390952"))
	if err != nil {
		lang.SayString("xdiv421")
	} else {
		if !(rexsult.ToString() == "-9.43092341E+447641459") {
			lang.SayString("xdiv421")
		}
	}
	rexsult, err = lang.RxFromString("3898.03188").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-82572.615"))
	if err != nil {
		lang.SayString("xdiv422")
	} else {
		if !(rexsult.ToString() == "-0.0472073202") {
			lang.SayString("xdiv422")
		}
	}
	rexsult, err = lang.RxFromString("-1.7619356").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-2546.64043"))
	if err != nil {
		lang.SayString("xdiv423")
	} else {
		if !(rexsult.ToString() == "0.000691866657") {
			lang.SayString("xdiv423")
		}
	}
	rexsult, err = lang.RxFromString("59714.1968").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("29734388.6E-564525525"))
	if err != nil {
		lang.SayString("xdiv424")
	} else {
		if !(rexsult.ToString() == "2.00825373E+564525522") {
			lang.SayString("xdiv424")
		}
	}
	rexsult, err = lang.RxFromString("6.88891136E-935467395").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-785049.562E-741671442"))
	if err != nil {
		lang.SayString("xdiv425")
	} else {
		if !(rexsult.ToString() == "-8.77512923E-193795959") {
			lang.SayString("xdiv425")
		}
	}
	rexsult, err = lang.RxFromString("975566251").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-519.858530"))
	if err != nil {
		lang.SayString("xdiv426")
	} else {
		if !(rexsult.ToString() == "-1876599.49") {
			lang.SayString("xdiv426")
		}
	}
	rexsult, err = lang.RxFromString("307401954").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-231481582."))
	if err != nil {
		lang.SayString("xdiv427")
	} else {
		if !(rexsult.ToString() == "-1.32797586") {
			lang.SayString("xdiv427")
		}
	}
	rexsult, err = lang.RxFromString("-403903.851").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("35.5049687E-72095155"))
	if err != nil {
		lang.SayString("xdiv429")
	} else {
		if !(rexsult.ToString() == "-1.1375981E+72095159") {
			lang.SayString("xdiv429")
		}
	}
	rexsult, err = lang.RxFromString("6.48674979").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-621732.532E+422575800"))
	if err != nil {
		lang.SayString("xdiv430")
	} else {
		if !(rexsult.ToString() == "-1.04333447E-422575805") {
			lang.SayString("xdiv430")
		}
	}
	rexsult, err = lang.RxFromString("-31401.9418").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("36.3960679"))
	if err != nil {
		lang.SayString("xdiv431")
	} else {
		if !(rexsult.ToString() == "-862.783911") {
			lang.SayString("xdiv431")
		}
	}
	rexsult, err = lang.RxFromString("31345321.1").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("51.5482191"))
	if err != nil {
		lang.SayString("xdiv432")
	} else {
		if !(rexsult.ToString() == "608077.673") {
			lang.SayString("xdiv432")
		}
	}
	rexsult, err = lang.RxFromString("-64.172844").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-28506227.2E-767965800"))
	if err != nil {
		lang.SayString("xdiv433")
	} else {
		if !(rexsult.ToString() == "2.25118686E+767965794") {
			lang.SayString("xdiv433")
		}
	}
	rexsult, err = lang.RxFromString("70437.1551").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-62916.1233"))
	if err != nil {
		lang.SayString("xdiv434")
	} else {
		if !(rexsult.ToString() == "-1.11954061") {
			lang.SayString("xdiv434")
		}
	}
	rexsult, err = lang.RxFromString("916260164").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-58.4017325"))
	if err != nil {
		lang.SayString("xdiv435")
	} else {
		if !(rexsult.ToString() == "-15688920.9") {
			lang.SayString("xdiv435")
		}
	}
	rexsult, err = lang.RxFromString("19889085.3E-46816480").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1581683.94"))
	if err != nil {
		lang.SayString("xdiv436")
	} else {
		if !(rexsult.ToString() == "1.25746268E-46816479") {
			lang.SayString("xdiv436")
		}
	}
	rexsult, err = lang.RxFromString("-56312.3383").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("789.466064"))
	if err != nil {
		lang.SayString("xdiv437")
	} else {
		if !(rexsult.ToString() == "-71.3296503") {
			lang.SayString("xdiv437")
		}
	}
	rexsult, err = lang.RxFromString("183442.849").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-925876106"))
	if err != nil {
		lang.SayString("xdiv438")
	} else {
		if !(rexsult.ToString() == "-0.000198128937") {
			lang.SayString("xdiv438")
		}
	}
	rexsult, err = lang.RxFromString("971113.655E-695540249").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-419351120E-977743823"))
	if err != nil {
		lang.SayString("xdiv439")
	} else {
		if !(rexsult.ToString() == "-2.3157531E+282203571") {
			lang.SayString("xdiv439")
		}
	}
	rexsult, err = lang.RxFromString("859658551.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("72338.2054"))
	if err != nil {
		lang.SayString("xdiv440")
	} else {
		if !(rexsult.ToString() == "11883.88") {
			lang.SayString("xdiv440")
		}
	}
	rexsult, err = lang.RxFromString("-3.86446630E+426816068").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-664.534737"))
	if err != nil {
		lang.SayString("xdiv441")
	} else {
		if !(rexsult.ToString() == "5.81529615E+426816065") {
			lang.SayString("xdiv441")
		}
	}
	rexsult, err = lang.RxFromString("-969.881818").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("31170.8555"))
	if err != nil {
		lang.SayString("xdiv442")
	} else {
		if !(rexsult.ToString() == "-0.0311150208") {
			lang.SayString("xdiv442")
		}
	}
	rexsult, err = lang.RxFromString("7980537.27").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("85.4040512"))
	if err != nil {
		lang.SayString("xdiv443")
	} else {
		if !(rexsult.ToString() == "93444.4814") {
			lang.SayString("xdiv443")
		}
	}
	rexsult, err = lang.RxFromString("-114609916.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("7525.14981"))
	if err != nil {
		lang.SayString("xdiv444")
	} else {
		if !(rexsult.ToString() == "-15230.2504") {
			lang.SayString("xdiv444")
		}
	}
	rexsult, err = lang.RxFromString("8.43404682E-500572568").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("474526719"))
	if err != nil {
		lang.SayString("xdiv445")
	} else {
		if !(rexsult.ToString() == "1.77735973E-500572576") {
			lang.SayString("xdiv445")
		}
	}
	rexsult, err = lang.RxFromString("188006433").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2260.17037E-978192525"))
	if err != nil {
		lang.SayString("xdiv446")
	} else {
		if !(rexsult.ToString() == "8.31824165E+978192529") {
			lang.SayString("xdiv446")
		}
	}
	rexsult, err = lang.RxFromString("-9.95836312").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-866466703"))
	if err != nil {
		lang.SayString("xdiv447")
	} else {
		if !(rexsult.ToString() == "1.14930707E-8") {
			lang.SayString("xdiv447")
		}
	}
	rexsult, err = lang.RxFromString("80919339.2E-967231586").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("219.824266"))
	if err != nil {
		lang.SayString("xdiv448")
	} else {
		if !(rexsult.ToString() == "3.6810922E-967231581") {
			lang.SayString("xdiv448")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-89827.5229"))
	if err != nil {
		lang.SayString("xdiv449")
	} else {
		if !(rexsult.ToString() == "-1.77650946") {
			lang.SayString("xdiv449")
		}
	}
	rexsult, err = lang.RxFromString("-4.54000153").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("6966333.74"))
	if err != nil {
		lang.SayString("xdiv450")
	} else {
		if !(rexsult.ToString() == "-6.51706005E-7") {
			lang.SayString("xdiv450")
		}
	}
	rexsult, err = lang.RxFromString("28701538.7E-391015649").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-920999192."))
	if err != nil {
		lang.SayString("xdiv451")
	} else {
		if !(rexsult.ToString() == "-3.11634787E-391015651") {
			lang.SayString("xdiv451")
		}
	}
	rexsult, err = lang.RxFromString("-361382575.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7976.15286E+898491169"))
	if err != nil {
		lang.SayString("xdiv452")
	} else {
		if !(rexsult.ToString() == "4.53078798E-898491165") {
			lang.SayString("xdiv452")
		}
	}
	rexsult, err = lang.RxFromString("7021805.61").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1222952.83"))
	if err != nil {
		lang.SayString("xdiv453")
	} else {
		if !(rexsult.ToString() == "5.74168148") {
			lang.SayString("xdiv453")
		}
	}
	rexsult, err = lang.RxFromString("-40.4811667").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-79655.5635"))
	if err != nil {
		lang.SayString("xdiv454")
	} else {
		if !(rexsult.ToString() == "0.000508202628") {
			lang.SayString("xdiv454")
		}
	}
	rexsult, err = lang.RxFromString("-8755674.38E+117168177").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("148.903404"))
	if err != nil {
		lang.SayString("xdiv455")
	} else {
		if !(rexsult.ToString() == "-5.88010357E+117168181") {
			lang.SayString("xdiv455")
		}
	}
	rexsult, err = lang.RxFromString("34.5329781E+382829392").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-45.2177309"))
	if err != nil {
		lang.SayString("xdiv456")
	} else {
		if !(rexsult.ToString() == "-7.63704357E+382829391") {
			lang.SayString("xdiv456")
		}
	}
	rexsult, err = lang.RxFromString("-37958476.0").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("584367.935"))
	if err != nil {
		lang.SayString("xdiv457")
	} else {
		if !(rexsult.ToString() == "-64.9564662") {
			lang.SayString("xdiv457")
		}
	}
	rexsult, err = lang.RxFromString("495233.553E-414152215").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("62352759.2"))
	if err != nil {
		lang.SayString("xdiv458")
	} else {
		if !(rexsult.ToString() == "7.94244809E-414152218") {
			lang.SayString("xdiv458")
		}
	}
	rexsult, err = lang.RxFromString("-502343060").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-96828.994"))
	if err != nil {
		lang.SayString("xdiv459")
	} else {
		if !(rexsult.ToString() == "5187.9405") {
			lang.SayString("xdiv459")
		}
	}
	rexsult, err = lang.RxFromString("-22.439639E+916362878").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-39.4037681"))
	if err != nil {
		lang.SayString("xdiv460")
	} else {
		if !(rexsult.ToString() == "5.69479521E+916362877") {
			lang.SayString("xdiv460")
		}
	}
	rexsult, err = lang.RxFromString("718180.587E-957473722").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1.66223443"))
	if err != nil {
		lang.SayString("xdiv461")
	} else {
		if !(rexsult.ToString() == "4.3205734E-957473717") {
			lang.SayString("xdiv461")
		}
	}
	rexsult, err = lang.RxFromString("-51592.2698").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-713885.741"))
	if err != nil {
		lang.SayString("xdiv462")
	} else {
		if !(rexsult.ToString() == "0.072269646") {
			lang.SayString("xdiv462")
		}
	}
	rexsult, err = lang.RxFromString("51.2279848E+80439745").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("207.55925E+865165070"))
	if err != nil {
		lang.SayString("xdiv463")
	} else {
		if !(rexsult.ToString() == "2.46811379E-784725326") {
			lang.SayString("xdiv463")
		}
	}
	rexsult, err = lang.RxFromString("-5983.23468").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-39.9544513"))
	if err != nil {
		lang.SayString("xdiv464")
	} else {
		if !(rexsult.ToString() == "149.751392") {
			lang.SayString("xdiv464")
		}
	}
	rexsult, err = lang.RxFromString("921639332.E-917542963").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("287325.891"))
	if err != nil {
		lang.SayString("xdiv465")
	} else {
		if !(rexsult.ToString() == "3.20764456E-917542960") {
			lang.SayString("xdiv465")
		}
	}
	rexsult, err = lang.RxFromString("91095916.8E-787312969").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-58643.418E+58189880"))
	if err != nil {
		lang.SayString("xdiv466")
	} else {
		if !(rexsult.ToString() == "-1.55338689E-845502846") {
			lang.SayString("xdiv466")
		}
	}
	rexsult, err = lang.RxFromString("-6410.5555").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-234964259"))
	if err != nil {
		lang.SayString("xdiv467")
	} else {
		if !(rexsult.ToString() == "0.000027283109") {
			lang.SayString("xdiv467")
		}
	}
	rexsult, err = lang.RxFromString("-5.32711606").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-8447286.21"))
	if err != nil {
		lang.SayString("xdiv468")
	} else {
		if !(rexsult.ToString() == "6.30630468E-7") {
			lang.SayString("xdiv468")
		}
	}
	rexsult, err = lang.RxFromString("-82272171.8").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-776.238587E-372690416"))
	if err != nil {
		lang.SayString("xdiv469")
	} else {
		if !(rexsult.ToString() == "1.05988253E+372690421") {
			lang.SayString("xdiv469")
		}
	}
	rexsult, err = lang.RxFromString("412411244.E-774339264").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("866452.465"))
	if err != nil {
		lang.SayString("xdiv470")
	} else {
		if !(rexsult.ToString() == "4.75976768E-774339262") {
			lang.SayString("xdiv470")
		}
	}
	rexsult, err = lang.RxFromString("-103.474598").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3.01660661E-446661257"))
	if err != nil {
		lang.SayString("xdiv471")
	} else {
		if !(rexsult.ToString() == "3.43016546E+446661258") {
			lang.SayString("xdiv471")
		}
	}
	rexsult, err = lang.RxFromString("-31027.8323").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-475378186."))
	if err != nil {
		lang.SayString("xdiv472")
	} else {
		if !(rexsult.ToString() == "0.0000652697856") {
			lang.SayString("xdiv472")
		}
	}
	rexsult, err = lang.RxFromString("-1199339.72").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-5.73068392E+53774632"))
	if err != nil {
		lang.SayString("xdiv473")
	} else {
		if !(rexsult.ToString() == "2.09283872E-53774627") {
			lang.SayString("xdiv473")
		}
	}
	rexsult, err = lang.RxFromString("-732908.930E+364345433").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-3486146.26"))
	if err != nil {
		lang.SayString("xdiv474")
	} else {
		if !(rexsult.ToString() == "2.10234705E+364345432") {
			lang.SayString("xdiv474")
		}
	}
	rexsult, err = lang.RxFromString("-2376150.83").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-46777583.3"))
	if err != nil {
		lang.SayString("xdiv475")
	} else {
		if !(rexsult.ToString() == "0.0507967847") {
			lang.SayString("xdiv475")
		}
	}
	rexsult, err = lang.RxFromString("6.3664211").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-140854908."))
	if err != nil {
		lang.SayString("xdiv476")
	} else {
		if !(rexsult.ToString() == "-4.51984328E-8") {
			lang.SayString("xdiv476")
		}
	}
	rexsult, err = lang.RxFromString("-15.791522").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("1902.30210E+90741844"))
	if err != nil {
		lang.SayString("xdiv477")
	} else {
		if !(rexsult.ToString() == "-8.30126929E-90741847") {
			lang.SayString("xdiv477")
		}
	}
	rexsult, err = lang.RxFromString("15356.1505E+373950429").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("2.88020400"))
	if err != nil {
		lang.SayString("xdiv478")
	} else {
		if !(rexsult.ToString() == "5.33161905E+373950432") {
			lang.SayString("xdiv478")
		}
	}
	rexsult, err = lang.RxFromString("-3.12001326E+318884762").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("9567.21595"))
	if err != nil {
		lang.SayString("xdiv479")
	} else {
		if !(rexsult.ToString() == "-3.26115066E+318884758") {
			lang.SayString("xdiv479")
		}
	}
	rexsult, err = lang.RxFromString("49436.6528").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("751.919517"))
	if err != nil {
		lang.SayString("xdiv480")
	} else {
		if !(rexsult.ToString() == "65.7472664") {
			lang.SayString("xdiv480")
		}
	}
	rexsult, err = lang.RxFromString("552.669453").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("8.3725760E+16223526"))
	if err != nil {
		lang.SayString("xdiv481")
	} else {
		if !(rexsult.ToString() == "6.60094878E-16223525") {
			lang.SayString("xdiv481")
		}
	}
	rexsult, err = lang.RxFromString("-3266303").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("453741.520"))
	if err != nil {
		lang.SayString("xdiv482")
	} else {
		if !(rexsult.ToString() == "-7.19859844") {
			lang.SayString("xdiv482")
		}
	}
	rexsult, err = lang.RxFromString("12302757.4").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("542922.487E+414443353"))
	if err != nil {
		lang.SayString("xdiv483")
	} else {
		if !(rexsult.ToString() == "2.26602465E-414443352") {
			lang.SayString("xdiv483")
		}
	}
	rexsult, err = lang.RxFromString("-5670757.79E-784754984").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("128144.503"))
	if err != nil {
		lang.SayString("xdiv484")
	} else {
		if !(rexsult.ToString() == "-4.42528369E-784754983") {
			lang.SayString("xdiv484")
		}
	}
	rexsult, err = lang.RxFromString("22.7721968E+842530698").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5223.70462"))
	if err != nil {
		lang.SayString("xdiv485")
	} else {
		if !(rexsult.ToString() == "4.35939596E+842530695") {
			lang.SayString("xdiv485")
		}
	}
	rexsult, err = lang.RxFromString("88.5158199E-980164357").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("325846116"))
	if err != nil {
		lang.SayString("xdiv486")
	} else {
		if !(rexsult.ToString() == "2.71649148E-980164364") {
			lang.SayString("xdiv486")
		}
	}
	rexsult, err = lang.RxFromString("-22881.0408").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("5.63661562"))
	if err != nil {
		lang.SayString("xdiv487")
	} else {
		if !(rexsult.ToString() == "-4059.35802") {
			lang.SayString("xdiv487")
		}
	}
	rexsult, err = lang.RxFromString("-7157.57449").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-76.4455519E-85647047"))
	if err != nil {
		lang.SayString("xdiv488")
	} else {
		if !(rexsult.ToString() == "9.36297052E+85647048") {
			lang.SayString("xdiv488")
		}
	}
	rexsult, err = lang.RxFromString("-503113.801").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-9715149.82E-612184422"))
	if err != nil {
		lang.SayString("xdiv489")
	} else {
		if !(rexsult.ToString() == "5.17865201E+612184420") {
			lang.SayString("xdiv489")
		}
	}
	rexsult, err = lang.RxFromString("-3066962.41").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-55.3096879"))
	if err != nil {
		lang.SayString("xdiv490")
	} else {
		if !(rexsult.ToString() == "55450.7271") {
			lang.SayString("xdiv490")
		}
	}
	rexsult, err = lang.RxFromString("-53311.5738E+156608936").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-7.45890666"))
	if err != nil {
		lang.SayString("xdiv491")
	} else {
		if !(rexsult.ToString() == "7.14737109E+156608939") {
			lang.SayString("xdiv491")
		}
	}
	rexsult, err = lang.RxFromString("998890068.").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-92.057879"))
	if err != nil {
		lang.SayString("xdiv492")
	} else {
		if !(rexsult.ToString() == "-10850674.4") {
			lang.SayString("xdiv492")
		}
	}
	rexsult, err = lang.RxFromString("122.495591").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-407836028."))
	if err != nil {
		lang.SayString("xdiv493")
	} else {
		if !(rexsult.ToString() == "-3.00355002E-7") {
			lang.SayString("xdiv493")
		}
	}
	rexsult, err = lang.RxFromString("187098.488").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("6220.05584E-236541249"))
	if err != nil {
		lang.SayString("xdiv494")
	} else {
		if !(rexsult.ToString() == "3.00798727E+236541250") {
			lang.SayString("xdiv494")
		}
	}
	rexsult, err = lang.RxFromString("4819899.21E+432982550").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-727441917"))
	if err != nil {
		lang.SayString("xdiv495")
	} else {
		if !(rexsult.ToString() == "-6.62582001E+432982547") {
			lang.SayString("xdiv495")
		}
	}
	rexsult, err = lang.RxFromString("5770.01020E+507459752").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-4208339.33E-129766680"))
	if err != nil {
		lang.SayString("xdiv496")
	} else {
		if !(rexsult.ToString() == "-1.37108958E+637226429") {
			lang.SayString("xdiv496")
		}
	}
	rexsult, err = lang.RxFromString("-286.371320").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("710319152"))
	if err != nil {
		lang.SayString("xdiv497")
	} else {
		if !(rexsult.ToString() == "-4.03158664E-7") {
			lang.SayString("xdiv497")
		}
	}
	rexsult, err = lang.RxFromString("-7.27403536").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-481469656E-835183700"))
	if err != nil {
		lang.SayString("xdiv498")
	} else {
		if !(rexsult.ToString() == "1.5107983E+835183692") {
			lang.SayString("xdiv498")
		}
	}
	rexsult, err = lang.RxFromString("-6157.74292").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("-94075286.2E+92555877"))
	if err != nil {
		lang.SayString("xdiv499")
	} else {
		if !(rexsult.ToString() == "6.5455479E-92555882") {
			lang.SayString("xdiv499")
		}
	}
	rexsult, err = lang.RxFromString("-525445087.E+231529167").OpDiv(lang.RxSetWithDigit(9), lang.RxFromString("188227460"))
	if err != nil {
		lang.SayString("xdiv500")
	} else {
		if !(rexsult.ToString() == "-2.79154321E+231529167") {
			lang.SayString("xdiv500")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rdvx101")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rdvx101")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rdvx102")
	} else {
		if !(rexsult.ToString() == "12344") {
			lang.SayString("rdvx102")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rdvx103")
	} else {
		if !(rexsult.ToString() == "12333") {
			lang.SayString("rdvx103")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rdvx104")
	} else {
		if !(rexsult.ToString() == "12223") {
			lang.SayString("rdvx104")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rdvx105")
	} else {
		if !(rexsult.ToString() == "11223") {
			lang.SayString("rdvx105")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx106")
	} else {
		if !(rexsult.ToString() == "3088.8") {
			lang.SayString("rdvx106")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx107")
	} else {
		if !(rexsult.ToString() == "3086.3") {
			lang.SayString("rdvx107")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx108")
	} else {
		if !(rexsult.ToString() == "3088.7") {
			lang.SayString("rdvx108")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx109")
	} else {
		if !(rexsult.ToString() == "3086.2") {
			lang.SayString("rdvx109")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rdvx110")
	} else {
		if !(rexsult.ToString() == "2519.4") {
			lang.SayString("rdvx110")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rdvx111")
	} else {
		if !(rexsult.ToString() == "2473.9") {
			lang.SayString("rdvx111")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rdvx112")
	} else {
		if !(rexsult.ToString() == "2469.5") {
			lang.SayString("rdvx112")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rdvx113")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx113")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rdvx114")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx114")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rdvx115")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx115")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rdvx116")
	} else {
		if !(rexsult.ToString() == "2468.5") {
			lang.SayString("rdvx116")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rdvx117")
	} else {
		if !(rexsult.ToString() == "2464.1") {
			lang.SayString("rdvx117")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("rdvx118")
	} else {
		if !(rexsult.ToString() == "2420.6") {
			lang.SayString("rdvx118")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rdvx201")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rdvx201")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rdvx202")
	} else {
		if !(rexsult.ToString() == "12344") {
			lang.SayString("rdvx202")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rdvx203")
	} else {
		if !(rexsult.ToString() == "12333") {
			lang.SayString("rdvx203")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rdvx204")
	} else {
		if !(rexsult.ToString() == "12223") {
			lang.SayString("rdvx204")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rdvx205")
	} else {
		if !(rexsult.ToString() == "11223") {
			lang.SayString("rdvx205")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx206")
	} else {
		if !(rexsult.ToString() == "3088.8") {
			lang.SayString("rdvx206")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx207")
	} else {
		if !(rexsult.ToString() == "3086.3") {
			lang.SayString("rdvx207")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx208")
	} else {
		if !(rexsult.ToString() == "3088.7") {
			lang.SayString("rdvx208")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx209")
	} else {
		if !(rexsult.ToString() == "3086.2") {
			lang.SayString("rdvx209")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rdvx210")
	} else {
		if !(rexsult.ToString() == "2519.4") {
			lang.SayString("rdvx210")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rdvx211")
	} else {
		if !(rexsult.ToString() == "2473.9") {
			lang.SayString("rdvx211")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rdvx212")
	} else {
		if !(rexsult.ToString() == "2469.5") {
			lang.SayString("rdvx212")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rdvx213")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx213")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rdvx214")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx214")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rdvx215")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx215")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rdvx216")
	} else {
		if !(rexsult.ToString() == "2468.5") {
			lang.SayString("rdvx216")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rdvx217")
	} else {
		if !(rexsult.ToString() == "2464.1") {
			lang.SayString("rdvx217")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("rdvx218")
	} else {
		if !(rexsult.ToString() == "2420.6") {
			lang.SayString("rdvx218")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rdvx301")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rdvx301")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rdvx302")
	} else {
		if !(rexsult.ToString() == "12344") {
			lang.SayString("rdvx302")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rdvx303")
	} else {
		if !(rexsult.ToString() == "12333") {
			lang.SayString("rdvx303")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rdvx304")
	} else {
		if !(rexsult.ToString() == "12223") {
			lang.SayString("rdvx304")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rdvx305")
	} else {
		if !(rexsult.ToString() == "11223") {
			lang.SayString("rdvx305")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx306")
	} else {
		if !(rexsult.ToString() == "3088.8") {
			lang.SayString("rdvx306")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx307")
	} else {
		if !(rexsult.ToString() == "3086.3") {
			lang.SayString("rdvx307")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx308")
	} else {
		if !(rexsult.ToString() == "3088.7") {
			lang.SayString("rdvx308")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx309")
	} else {
		if !(rexsult.ToString() == "3086.2") {
			lang.SayString("rdvx309")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rdvx310")
	} else {
		if !(rexsult.ToString() == "2519.4") {
			lang.SayString("rdvx310")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rdvx311")
	} else {
		if !(rexsult.ToString() == "2473.9") {
			lang.SayString("rdvx311")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rdvx312")
	} else {
		if !(rexsult.ToString() == "2469.5") {
			lang.SayString("rdvx312")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rdvx313")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx313")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rdvx314")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx314")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rdvx315")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx315")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rdvx316")
	} else {
		if !(rexsult.ToString() == "2468.5") {
			lang.SayString("rdvx316")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rdvx317")
	} else {
		if !(rexsult.ToString() == "2464.1") {
			lang.SayString("rdvx317")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("rdvx318")
	} else {
		if !(rexsult.ToString() == "2420.6") {
			lang.SayString("rdvx318")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rdvx401")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rdvx401")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rdvx402")
	} else {
		if !(rexsult.ToString() == "12344") {
			lang.SayString("rdvx402")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rdvx403")
	} else {
		if !(rexsult.ToString() == "12333") {
			lang.SayString("rdvx403")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rdvx404")
	} else {
		if !(rexsult.ToString() == "12223") {
			lang.SayString("rdvx404")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rdvx405")
	} else {
		if !(rexsult.ToString() == "11223") {
			lang.SayString("rdvx405")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx406")
	} else {
		if !(rexsult.ToString() == "3088.8") {
			lang.SayString("rdvx406")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx407")
	} else {
		if !(rexsult.ToString() == "3086.3") {
			lang.SayString("rdvx407")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx408")
	} else {
		if !(rexsult.ToString() == "3088.7") {
			lang.SayString("rdvx408")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx409")
	} else {
		if !(rexsult.ToString() == "3086.2") {
			lang.SayString("rdvx409")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rdvx410")
	} else {
		if !(rexsult.ToString() == "2519.4") {
			lang.SayString("rdvx410")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rdvx411")
	} else {
		if !(rexsult.ToString() == "2473.9") {
			lang.SayString("rdvx411")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rdvx412")
	} else {
		if !(rexsult.ToString() == "2469.5") {
			lang.SayString("rdvx412")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rdvx413")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx413")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rdvx414")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx414")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rdvx415")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx415")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rdvx416")
	} else {
		if !(rexsult.ToString() == "2468.5") {
			lang.SayString("rdvx416")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rdvx417")
	} else {
		if !(rexsult.ToString() == "2464.1") {
			lang.SayString("rdvx417")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("rdvx418")
	} else {
		if !(rexsult.ToString() == "2420.6") {
			lang.SayString("rdvx418")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rdvx501")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rdvx501")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rdvx502")
	} else {
		if !(rexsult.ToString() == "12344") {
			lang.SayString("rdvx502")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rdvx503")
	} else {
		if !(rexsult.ToString() == "12333") {
			lang.SayString("rdvx503")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rdvx504")
	} else {
		if !(rexsult.ToString() == "12223") {
			lang.SayString("rdvx504")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rdvx505")
	} else {
		if !(rexsult.ToString() == "11223") {
			lang.SayString("rdvx505")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx506")
	} else {
		if !(rexsult.ToString() == "3088.8") {
			lang.SayString("rdvx506")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx507")
	} else {
		if !(rexsult.ToString() == "3086.3") {
			lang.SayString("rdvx507")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx508")
	} else {
		if !(rexsult.ToString() == "3088.7") {
			lang.SayString("rdvx508")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx509")
	} else {
		if !(rexsult.ToString() == "3086.2") {
			lang.SayString("rdvx509")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rdvx510")
	} else {
		if !(rexsult.ToString() == "2519.4") {
			lang.SayString("rdvx510")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rdvx511")
	} else {
		if !(rexsult.ToString() == "2473.9") {
			lang.SayString("rdvx511")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rdvx512")
	} else {
		if !(rexsult.ToString() == "2469.5") {
			lang.SayString("rdvx512")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rdvx513")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx513")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rdvx514")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx514")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rdvx515")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx515")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rdvx516")
	} else {
		if !(rexsult.ToString() == "2468.5") {
			lang.SayString("rdvx516")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rdvx517")
	} else {
		if !(rexsult.ToString() == "2464.1") {
			lang.SayString("rdvx517")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("rdvx518")
	} else {
		if !(rexsult.ToString() == "2420.6") {
			lang.SayString("rdvx518")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rdvx601")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rdvx601")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rdvx602")
	} else {
		if !(rexsult.ToString() == "12344") {
			lang.SayString("rdvx602")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rdvx603")
	} else {
		if !(rexsult.ToString() == "12333") {
			lang.SayString("rdvx603")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rdvx604")
	} else {
		if !(rexsult.ToString() == "12223") {
			lang.SayString("rdvx604")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rdvx605")
	} else {
		if !(rexsult.ToString() == "11223") {
			lang.SayString("rdvx605")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx606")
	} else {
		if !(rexsult.ToString() == "3088.8") {
			lang.SayString("rdvx606")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx607")
	} else {
		if !(rexsult.ToString() == "3086.3") {
			lang.SayString("rdvx607")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx608")
	} else {
		if !(rexsult.ToString() == "3088.7") {
			lang.SayString("rdvx608")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx609")
	} else {
		if !(rexsult.ToString() == "3086.2") {
			lang.SayString("rdvx609")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rdvx610")
	} else {
		if !(rexsult.ToString() == "2519.4") {
			lang.SayString("rdvx610")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rdvx611")
	} else {
		if !(rexsult.ToString() == "2473.9") {
			lang.SayString("rdvx611")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rdvx612")
	} else {
		if !(rexsult.ToString() == "2469.5") {
			lang.SayString("rdvx612")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rdvx613")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx613")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rdvx614")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx614")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rdvx615")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx615")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rdvx616")
	} else {
		if !(rexsult.ToString() == "2468.5") {
			lang.SayString("rdvx616")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rdvx617")
	} else {
		if !(rexsult.ToString() == "2464.1") {
			lang.SayString("rdvx617")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("rdvx618")
	} else {
		if !(rexsult.ToString() == "2420.6") {
			lang.SayString("rdvx618")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rdvx701")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rdvx701")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("rdvx702")
	} else {
		if !(rexsult.ToString() == "12344") {
			lang.SayString("rdvx702")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("rdvx703")
	} else {
		if !(rexsult.ToString() == "12333") {
			lang.SayString("rdvx703")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("rdvx704")
	} else {
		if !(rexsult.ToString() == "12223") {
			lang.SayString("rdvx704")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("rdvx705")
	} else {
		if !(rexsult.ToString() == "11223") {
			lang.SayString("rdvx705")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx706")
	} else {
		if !(rexsult.ToString() == "3088.8") {
			lang.SayString("rdvx706")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rdvx707")
	} else {
		if !(rexsult.ToString() == "3086.3") {
			lang.SayString("rdvx707")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx708")
	} else {
		if !(rexsult.ToString() == "3088.7") {
			lang.SayString("rdvx708")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("rdvx709")
	} else {
		if !(rexsult.ToString() == "3086.2") {
			lang.SayString("rdvx709")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("rdvx710")
	} else {
		if !(rexsult.ToString() == "2519.4") {
			lang.SayString("rdvx710")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("rdvx711")
	} else {
		if !(rexsult.ToString() == "2473.9") {
			lang.SayString("rdvx711")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("rdvx712")
	} else {
		if !(rexsult.ToString() == "2469.5") {
			lang.SayString("rdvx712")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("rdvx713")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx713")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rdvx714")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx714")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("rdvx715")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("rdvx715")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("rdvx716")
	} else {
		if !(rexsult.ToString() == "2468.5") {
			lang.SayString("rdvx716")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("rdvx717")
	} else {
		if !(rexsult.ToString() == "2464.1") {
			lang.SayString("rdvx717")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("rdvx718")
	} else {
		if !(rexsult.ToString() == "2420.6") {
			lang.SayString("rdvx718")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("r0dvx101")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("r0dvx101")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.0001"))
	if err != nil {
		lang.SayString("r0dvx102")
	} else {
		if !(rexsult.ToString() == "12344") {
			lang.SayString("r0dvx102")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.001"))
	if err != nil {
		lang.SayString("r0dvx103")
	} else {
		if !(rexsult.ToString() == "12333") {
			lang.SayString("r0dvx103")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.01"))
	if err != nil {
		lang.SayString("r0dvx104")
	} else {
		if !(rexsult.ToString() == "12223") {
			lang.SayString("r0dvx104")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("1.1"))
	if err != nil {
		lang.SayString("r0dvx105")
	} else {
		if !(rexsult.ToString() == "11223") {
			lang.SayString("r0dvx105")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("r0dvx106")
	} else {
		if !(rexsult.ToString() == "3088.8") {
			lang.SayString("r0dvx106")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("r0dvx107")
	} else {
		if !(rexsult.ToString() == "3086.3") {
			lang.SayString("r0dvx107")
		}
	}
	rexsult, err = lang.RxFromString("12355").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("r0dvx108")
	} else {
		if !(rexsult.ToString() == "3088.7") {
			lang.SayString("r0dvx108")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.0001"))
	if err != nil {
		lang.SayString("r0dvx109")
	} else {
		if !(rexsult.ToString() == "3086.2") {
			lang.SayString("r0dvx109")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9"))
	if err != nil {
		lang.SayString("r0dvx110")
	} else {
		if !(rexsult.ToString() == "2519.4") {
			lang.SayString("r0dvx110")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.99"))
	if err != nil {
		lang.SayString("r0dvx111")
	} else {
		if !(rexsult.ToString() == "2473.9") {
			lang.SayString("r0dvx111")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.999"))
	if err != nil {
		lang.SayString("r0dvx112")
	} else {
		if !(rexsult.ToString() == "2469.5") {
			lang.SayString("r0dvx112")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("4.9999"))
	if err != nil {
		lang.SayString("r0dvx113")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("r0dvx113")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("r0dvx114")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("r0dvx114")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.0001"))
	if err != nil {
		lang.SayString("r0dvx115")
	} else {
		if !(rexsult.ToString() == "2469") {
			lang.SayString("r0dvx115")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.001"))
	if err != nil {
		lang.SayString("r0dvx116")
	} else {
		if !(rexsult.ToString() == "2468.5") {
			lang.SayString("r0dvx116")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.01"))
	if err != nil {
		lang.SayString("r0dvx117")
	} else {
		if !(rexsult.ToString() == "2464.1") {
			lang.SayString("r0dvx117")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpDiv(lang.RxSetWithDigit(5), lang.RxFromString("5.1"))
	if err != nil {
		lang.SayString("r0dvx118")
	} else {
		if !(rexsult.ToString() == "2420.6") {
			lang.SayString("r0dvx118")
		}
	}

	return
}
