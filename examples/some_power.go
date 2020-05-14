package main

import "agorexx/lang"

func main() {

	rexsult, err := lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(4), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("inx301")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("inx301")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(4), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("inx302")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("inx302")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(4), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("inx303")
	} else {
		if !(rexsult.ToString() == "0.003906") {
			lang.SayString("inx303")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(4), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("inx304")
	} else {
		if !(rexsult.ToString() == "0.00001526") {
			lang.SayString("inx304")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(4), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("inx305")
	} else {
		if !(rexsult.ToString() == "2.328E-10") {
			lang.SayString("inx305")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx001")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx001")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx002")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx002")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx003")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx003")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx004")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx004")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx005")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx005")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx006")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx006")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx010")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx010")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx011")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("powx011")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx012")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("powx012")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx013")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("powx013")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx014")
	} else {
		if !(rexsult.ToString() == "16") {
			lang.SayString("powx014")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx015")
	} else {
		if !(rexsult.ToString() == "32") {
			lang.SayString("powx015")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx016")
	} else {
		if !(rexsult.ToString() == "64") {
			lang.SayString("powx016")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx017")
	} else {
		if !(rexsult.ToString() == "128") {
			lang.SayString("powx017")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx018")
	} else {
		if !(rexsult.ToString() == "256") {
			lang.SayString("powx018")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx019")
	} else {
		if !(rexsult.ToString() == "512") {
			lang.SayString("powx019")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx020")
	} else {
		if !(rexsult.ToString() == "1024") {
			lang.SayString("powx020")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("11"))
	if err != nil {
		lang.SayString("powx021")
	} else {
		if !(rexsult.ToString() == "2048") {
			lang.SayString("powx021")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("12"))
	if err != nil {
		lang.SayString("powx022")
	} else {
		if !(rexsult.ToString() == "4096") {
			lang.SayString("powx022")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("15"))
	if err != nil {
		lang.SayString("powx023")
	} else {
		if !(rexsult.ToString() == "32768") {
			lang.SayString("powx023")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("powx024")
	} else {
		if !(rexsult.ToString() == "65536") {
			lang.SayString("powx024")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("31"))
	if err != nil {
		lang.SayString("powx025")
	} else {
		if !(rexsult.ToString() == "2147483648") {
			lang.SayString("powx025")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("powx026")
	} else {
		if !(rexsult.ToString() == "4294967296") {
			lang.SayString("powx026")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("31"))
	if err != nil {
		lang.SayString("powx027")
	} else {
		if !(rexsult.ToString() == "2.14748365E+9") {
			lang.SayString("powx027")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("powx028")
	} else {
		if !(rexsult.ToString() == "4.2949673E+9") {
			lang.SayString("powx028")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(10), lang.RxFromString("31"))
	if err != nil {
		lang.SayString("powx029")
	} else {
		if !(rexsult.ToString() == "2147483648") {
			lang.SayString("powx029")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(10), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("powx030")
	} else {
		if !(rexsult.ToString() == "4294967296") {
			lang.SayString("powx030")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx031")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("powx031")
		}
	}
	rexsult, err = lang.RxFromString("4").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx032")
	} else {
		if !(rexsult.ToString() == "16") {
			lang.SayString("powx032")
		}
	}
	rexsult, err = lang.RxFromString("5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx033")
	} else {
		if !(rexsult.ToString() == "25") {
			lang.SayString("powx033")
		}
	}
	rexsult, err = lang.RxFromString("6").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx034")
	} else {
		if !(rexsult.ToString() == "36") {
			lang.SayString("powx034")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx035")
	} else {
		if !(rexsult.ToString() == "49") {
			lang.SayString("powx035")
		}
	}
	rexsult, err = lang.RxFromString("8").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx036")
	} else {
		if !(rexsult.ToString() == "64") {
			lang.SayString("powx036")
		}
	}
	rexsult, err = lang.RxFromString("9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx037")
	} else {
		if !(rexsult.ToString() == "81") {
			lang.SayString("powx037")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx038")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("powx038")
		}
	}
	rexsult, err = lang.RxFromString("11").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx039")
	} else {
		if !(rexsult.ToString() == "121") {
			lang.SayString("powx039")
		}
	}
	rexsult, err = lang.RxFromString("12").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx040")
	} else {
		if !(rexsult.ToString() == "144") {
			lang.SayString("powx040")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx041")
	} else {
		if !(rexsult.ToString() == "27") {
			lang.SayString("powx041")
		}
	}
	rexsult, err = lang.RxFromString("4").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx042")
	} else {
		if !(rexsult.ToString() == "64") {
			lang.SayString("powx042")
		}
	}
	rexsult, err = lang.RxFromString("5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx043")
	} else {
		if !(rexsult.ToString() == "125") {
			lang.SayString("powx043")
		}
	}
	rexsult, err = lang.RxFromString("6").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx044")
	} else {
		if !(rexsult.ToString() == "216") {
			lang.SayString("powx044")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx045")
	} else {
		if !(rexsult.ToString() == "343") {
			lang.SayString("powx045")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx047")
	} else {
		if !(rexsult.ToString() == "-27") {
			lang.SayString("powx047")
		}
	}
	rexsult, err = lang.RxFromString("-4").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx048")
	} else {
		if !(rexsult.ToString() == "-64") {
			lang.SayString("powx048")
		}
	}
	rexsult, err = lang.RxFromString("-5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx049")
	} else {
		if !(rexsult.ToString() == "-125") {
			lang.SayString("powx049")
		}
	}
	rexsult, err = lang.RxFromString("-6").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx050")
	} else {
		if !(rexsult.ToString() == "-216") {
			lang.SayString("powx050")
		}
	}
	rexsult, err = lang.RxFromString("-7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx051")
	} else {
		if !(rexsult.ToString() == "-343") {
			lang.SayString("powx051")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx052")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx052")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx053")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("powx053")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx054")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("powx054")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx055")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("powx055")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx056")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("powx056")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx057")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("powx057")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx058")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("powx058")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx059")
	} else {
		if !(rexsult.ToString() == "10000000") {
			lang.SayString("powx059")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx060")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("powx060")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx061")
	} else {
		if !(rexsult.ToString() == "1E+9") {
			lang.SayString("powx061")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("22"))
	if err != nil {
		lang.SayString("powx062")
	} else {
		if !(rexsult.ToString() == "1E+22") {
			lang.SayString("powx062")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("77"))
	if err != nil {
		lang.SayString("powx063")
	} else {
		if !(rexsult.ToString() == "1E+77") {
			lang.SayString("powx063")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("99"))
	if err != nil {
		lang.SayString("powx064")
	} else {
		if !(rexsult.ToString() == "1E+99") {
			lang.SayString("powx064")
		}
	}
	rexsult, err = lang.RxFromString("0.3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx070")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx070")
		}
	}
	rexsult, err = lang.RxFromString("0.3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx071")
	} else {
		if !(rexsult.ToString() == "0.3") {
			lang.SayString("powx071")
		}
	}
	rexsult, err = lang.RxFromString("0.3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("powx072")
	} else {
		if !(rexsult.ToString() == "0.3") {
			lang.SayString("powx072")
		}
	}
	rexsult, err = lang.RxFromString("0.3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("powx073")
	} else {
		if !(rexsult.ToString() == "0.09") {
			lang.SayString("powx073")
		}
	}
	rexsult, err = lang.RxFromString("0.3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2.000000000"))
	if err != nil {
		lang.SayString("powx074")
	} else {
		if !(rexsult.ToString() == "0.09") {
			lang.SayString("powx074")
		}
	}
	rexsult, err = lang.RxFromString("6.0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx075")
	} else {
		if !(rexsult.ToString() == "6") {
			lang.SayString("powx075")
		}
	}
	rexsult, err = lang.RxFromString("6.0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx076")
	} else {
		if !(rexsult.ToString() == "36") {
			lang.SayString("powx076")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx077")
	} else {
		if !(rexsult.ToString() == "9") {
			lang.SayString("powx077")
		}
	}
	rexsult, err = lang.RxFromString("4").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx078")
	} else {
		if !(rexsult.ToString() == "64") {
			lang.SayString("powx078")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx080")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx080")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx081")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("powx081")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx082")
	} else {
		if !(rexsult.ToString() == "0.01") {
			lang.SayString("powx082")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx083")
	} else {
		if !(rexsult.ToString() == "0.001") {
			lang.SayString("powx083")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx084")
	} else {
		if !(rexsult.ToString() == "0.0001") {
			lang.SayString("powx084")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx085")
	} else {
		if !(rexsult.ToString() == "0.00001") {
			lang.SayString("powx085")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx086")
	} else {
		if !(rexsult.ToString() == "0.000001") {
			lang.SayString("powx086")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx087")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("powx087")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx088")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("powx088")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx089")
	} else {
		if !(rexsult.ToString() == "1E-9") {
			lang.SayString("powx089")
		}
	}
	rexsult, err = lang.RxFromString("101").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx090")
	} else {
		if !(rexsult.ToString() == "10201") {
			lang.SayString("powx090")
		}
	}
	rexsult, err = lang.RxFromString("101").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx091")
	} else {
		if !(rexsult.ToString() == "1030301") {
			lang.SayString("powx091")
		}
	}
	rexsult, err = lang.RxFromString("101").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx092")
	} else {
		if !(rexsult.ToString() == "104060401") {
			lang.SayString("powx092")
		}
	}
	rexsult, err = lang.RxFromString("101").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx093")
	} else {
		if !(rexsult.ToString() == "1.05101005E+10") {
			lang.SayString("powx093")
		}
	}
	rexsult, err = lang.RxFromString("101").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx094")
	} else {
		if !(rexsult.ToString() == "1.06152015E+12") {
			lang.SayString("powx094")
		}
	}
	rexsult, err = lang.RxFromString("101").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx095")
	} else {
		if !(rexsult.ToString() == "1.07213535E+14") {
			lang.SayString("powx095")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("powx099")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx099")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("powx100")
	} else {
		if !(rexsult.ToString() == "0.333333333") {
			lang.SayString("powx100")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("powx101")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("powx101")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx102")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("powx102")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx103")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("powx103")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("powx104")
	} else {
		if !(rexsult.ToString() == "0.00390625") {
			lang.SayString("powx104")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-16"))
	if err != nil {
		lang.SayString("powx105")
	} else {
		if !(rexsult.ToString() == "0.0000152587891") {
			lang.SayString("powx105")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("powx106")
	} else {
		if !(rexsult.ToString() == "2.32830644E-10") {
			lang.SayString("powx106")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-64"))
	if err != nil {
		lang.SayString("powx108")
	} else {
		if !(rexsult.ToString() == "5.42101086E-20") {
			lang.SayString("powx108")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("powx110")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("powx110")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx111")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("powx111")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx112")
	} else {
		if !(rexsult.ToString() == "0.000001") {
			lang.SayString("powx112")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx113")
	} else {
		if !(rexsult.ToString() == "0.00001") {
			lang.SayString("powx113")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx114")
	} else {
		if !(rexsult.ToString() == "0.0001") {
			lang.SayString("powx114")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx115")
	} else {
		if !(rexsult.ToString() == "0.001") {
			lang.SayString("powx115")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx116")
	} else {
		if !(rexsult.ToString() == "0.01") {
			lang.SayString("powx116")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("powx117")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("powx117")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-77"))
	if err != nil {
		lang.SayString("powx121")
	} else {
		if !(rexsult.ToString() == "1E-77") {
			lang.SayString("powx121")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-22"))
	if err != nil {
		lang.SayString("powx122")
	} else {
		if !(rexsult.ToString() == "1E-22") {
			lang.SayString("powx122")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("powx123")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("powx123")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx124")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("powx124")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx125")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("powx125")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx200")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx200")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx201")
	} else {
		if !(rexsult.ToString() == "0.5") {
			lang.SayString("powx201")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx202")
	} else {
		if !(rexsult.ToString() == "0.25") {
			lang.SayString("powx202")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx203")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("powx203")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx204")
	} else {
		if !(rexsult.ToString() == "0.0625") {
			lang.SayString("powx204")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx205")
	} else {
		if !(rexsult.ToString() == "0.03125") {
			lang.SayString("powx205")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx206")
	} else {
		if !(rexsult.ToString() == "0.015625") {
			lang.SayString("powx206")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx207")
	} else {
		if !(rexsult.ToString() == "0.0078125") {
			lang.SayString("powx207")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx208")
	} else {
		if !(rexsult.ToString() == "0.00390625") {
			lang.SayString("powx208")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx209")
	} else {
		if !(rexsult.ToString() == "0.001953125") {
			lang.SayString("powx209")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx210")
	} else {
		if !(rexsult.ToString() == "0.0009765625") {
			lang.SayString("powx210")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("powx211")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx211")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999998"))
	if err != nil {
		lang.SayString("powx212")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx212")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx213")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx213")
		}
	}
	rexsult, err = lang.RxFromString("123456789E+10").OpPow(lang.RxSetWithDigit(3), lang.RxFromString("-1.23000e+2"))
	if err != nil {
		lang.SayString("powx219")
	} else {
		if !(rexsult.ToString() == "5.86E-2226") {
			lang.SayString("powx219")
		}
	}
	rexsult, err = lang.RxFromString("0E-30").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx223")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx223")
		}
	}
	rexsult, err = lang.RxFromString("0E-10").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx224")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx224")
		}
	}
	rexsult, err = lang.RxFromString("0E-1").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx225")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx225")
		}
	}
	rexsult, err = lang.RxFromString("0E+0").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx226")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx226")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx227")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx227")
		}
	}
	rexsult, err = lang.RxFromString("0E+1").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx228")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx228")
		}
	}
	rexsult, err = lang.RxFromString("0E+10").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx229")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx229")
		}
	}
	rexsult, err = lang.RxFromString("0E+30").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx230")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx230")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E-30"))
	if err != nil {
		lang.SayString("powx231")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx231")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E-10"))
	if err != nil {
		lang.SayString("powx232")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx232")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E-1"))
	if err != nil {
		lang.SayString("powx233")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx233")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E+0"))
	if err != nil {
		lang.SayString("powx234")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx234")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx235")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx235")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E+1"))
	if err != nil {
		lang.SayString("powx236")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx236")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E+10"))
	if err != nil {
		lang.SayString("powx237")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx237")
		}
	}
	rexsult, err = lang.RxFromString("3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E+30"))
	if err != nil {
		lang.SayString("powx238")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx238")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E-30"))
	if err != nil {
		lang.SayString("powx247")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx247")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E-10"))
	if err != nil {
		lang.SayString("powx248")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx248")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E-1"))
	if err != nil {
		lang.SayString("powx249")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx249")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E+0"))
	if err != nil {
		lang.SayString("powx250")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx250")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx251")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx251")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E+1"))
	if err != nil {
		lang.SayString("powx252")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx252")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E+10"))
	if err != nil {
		lang.SayString("powx253")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx253")
		}
	}
	rexsult, err = lang.RxFromString("-3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0E+30"))
	if err != nil {
		lang.SayString("powx254")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx254")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx260")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx260")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx261")
	} else {
		if !(rexsult.ToString() == "-10") {
			lang.SayString("powx261")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx262")
	} else {
		if !(rexsult.ToString() == "100") {
			lang.SayString("powx262")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx263")
	} else {
		if !(rexsult.ToString() == "-1000") {
			lang.SayString("powx263")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx264")
	} else {
		if !(rexsult.ToString() == "10000") {
			lang.SayString("powx264")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx265")
	} else {
		if !(rexsult.ToString() == "-100000") {
			lang.SayString("powx265")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx266")
	} else {
		if !(rexsult.ToString() == "1000000") {
			lang.SayString("powx266")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx267")
	} else {
		if !(rexsult.ToString() == "-10000000") {
			lang.SayString("powx267")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx268")
	} else {
		if !(rexsult.ToString() == "100000000") {
			lang.SayString("powx268")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx269")
	} else {
		if !(rexsult.ToString() == "-1E+9") {
			lang.SayString("powx269")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("22"))
	if err != nil {
		lang.SayString("powx270")
	} else {
		if !(rexsult.ToString() == "1E+22") {
			lang.SayString("powx270")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("77"))
	if err != nil {
		lang.SayString("powx271")
	} else {
		if !(rexsult.ToString() == "-1E+77") {
			lang.SayString("powx271")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("99"))
	if err != nil {
		lang.SayString("powx272")
	} else {
		if !(rexsult.ToString() == "-1E+99") {
			lang.SayString("powx272")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("999"))
	if err != nil {
		lang.SayString("powx391")
	} else {
		if !(rexsult.ToString() == "1E-999") {
			lang.SayString("powx391")
		}
	}
	rexsult, err = lang.RxFromString("0.099").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("999"))
	if err != nil {
		lang.SayString("powx392")
	} else {
		if !(rexsult.ToString() == "4.36073206168265E-1004") {
			lang.SayString("powx392")
		}
	}
	rexsult, err = lang.RxFromString("0.098").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("999"))
	if err != nil {
		lang.SayString("powx393")
	} else {
		if !(rexsult.ToString() == "1.7173136298122E-1008") {
			lang.SayString("powx393")
		}
	}
	rexsult, err = lang.RxFromString("0.097").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("999"))
	if err != nil {
		lang.SayString("powx394")
	} else {
		if !(rexsult.ToString() == "6.09484312744338E-1013") {
			lang.SayString("powx394")
		}
	}
	rexsult, err = lang.RxFromString("0.096").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("999"))
	if err != nil {
		lang.SayString("powx395")
	} else {
		if !(rexsult.ToString() == "1.94518900694812E-1017") {
			lang.SayString("powx395")
		}
	}
	rexsult, err = lang.RxFromString("0.01").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("999"))
	if err != nil {
		lang.SayString("powx396")
	} else {
		if !(rexsult.ToString() == "1E-1998") {
			lang.SayString("powx396")
		}
	}
	rexsult, err = lang.RxFromString("0.02").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("100000000"))
	if err != nil {
		lang.SayString("powx397")
	} else {
		if !(rexsult.ToString() == "3.68466593698046E-169897001") {
			lang.SayString("powx397")
		}
	}
	rexsult, err = lang.RxFromString("1E-502").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx400")
	} else {
		if !(rexsult.ToString() == "1E-1004") {
			lang.SayString("powx400")
		}
	}
	rexsult, err = lang.RxFromString("1E-501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx401")
	} else {
		if !(rexsult.ToString() == "1E-1002") {
			lang.SayString("powx401")
		}
	}
	rexsult, err = lang.RxFromString("2E-501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx402")
	} else {
		if !(rexsult.ToString() == "4E-1002") {
			lang.SayString("powx402")
		}
	}
	rexsult, err = lang.RxFromString("4E-501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx403")
	} else {
		if !(rexsult.ToString() == "1.6E-1001") {
			lang.SayString("powx403")
		}
	}
	rexsult, err = lang.RxFromString("10E-501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx404")
	} else {
		if !(rexsult.ToString() == "1E-1000") {
			lang.SayString("powx404")
		}
	}
	rexsult, err = lang.RxFromString("30E-501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx405")
	} else {
		if !(rexsult.ToString() == "9E-1000") {
			lang.SayString("powx405")
		}
	}
	rexsult, err = lang.RxFromString("40E-501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx406")
	} else {
		if !(rexsult.ToString() == "1.6E-999") {
			lang.SayString("powx406")
		}
	}
	rexsult, err = lang.RxFromString("1E-335").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx410")
	} else {
		if !(rexsult.ToString() == "1E-1005") {
			lang.SayString("powx410")
		}
	}
	rexsult, err = lang.RxFromString("1E-334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx411")
	} else {
		if !(rexsult.ToString() == "1E-1002") {
			lang.SayString("powx411")
		}
	}
	rexsult, err = lang.RxFromString("2E-334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx412")
	} else {
		if !(rexsult.ToString() == "8E-1002") {
			lang.SayString("powx412")
		}
	}
	rexsult, err = lang.RxFromString("3E-334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx413")
	} else {
		if !(rexsult.ToString() == "2.7E-1001") {
			lang.SayString("powx413")
		}
	}
	rexsult, err = lang.RxFromString("4E-334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx414")
	} else {
		if !(rexsult.ToString() == "6.4E-1001") {
			lang.SayString("powx414")
		}
	}
	rexsult, err = lang.RxFromString("5E-334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx415")
	} else {
		if !(rexsult.ToString() == "1.25E-1000") {
			lang.SayString("powx415")
		}
	}
	rexsult, err = lang.RxFromString("10E-334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx416")
	} else {
		if !(rexsult.ToString() == "1E-999") {
			lang.SayString("powx416")
		}
	}
	rexsult, err = lang.RxFromString("2.5E-501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx421")
	} else {
		if !(rexsult.ToString() == "1.6E+1001") {
			lang.SayString("powx421")
		}
	}
	rexsult, err = lang.RxFromString("2.5E-500").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx422")
	} else {
		if !(rexsult.ToString() == "1.6E+999") {
			lang.SayString("powx422")
		}
	}
	rexsult, err = lang.RxFromString("2.5E+499").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx423")
	} else {
		if !(rexsult.ToString() == "1.6E-999") {
			lang.SayString("powx423")
		}
	}
	rexsult, err = lang.RxFromString("2.5E+500").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx424")
	} else {
		if !(rexsult.ToString() == "1.6E-1001") {
			lang.SayString("powx424")
		}
	}
	rexsult, err = lang.RxFromString("2.5E+501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx425")
	} else {
		if !(rexsult.ToString() == "1.6E-1003") {
			lang.SayString("powx425")
		}
	}
	rexsult, err = lang.RxFromString("2.5E+502").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx426")
	} else {
		if !(rexsult.ToString() == "1.6E-1005") {
			lang.SayString("powx426")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+499").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx427")
	} else {
		if !(rexsult.ToString() == "1.6E-997") {
			lang.SayString("powx427")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+500").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx428")
	} else {
		if !(rexsult.ToString() == "1.6E-999") {
			lang.SayString("powx428")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx429")
	} else {
		if !(rexsult.ToString() == "1.6E-1001") {
			lang.SayString("powx429")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+502").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx430")
	} else {
		if !(rexsult.ToString() == "1.6E-1003") {
			lang.SayString("powx430")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+503").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx431")
	} else {
		if !(rexsult.ToString() == "1.6E-1005") {
			lang.SayString("powx431")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+499").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx432")
	} else {
		if !(rexsult.ToString() == "6.25E-996") {
			lang.SayString("powx432")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+500").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx433")
	} else {
		if !(rexsult.ToString() == "6.25E-998") {
			lang.SayString("powx433")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx434")
	} else {
		if !(rexsult.ToString() == "6.25E-1000") {
			lang.SayString("powx434")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+502").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx435")
	} else {
		if !(rexsult.ToString() == "6.25E-1002") {
			lang.SayString("powx435")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+503").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx436")
	} else {
		if !(rexsult.ToString() == "6.25E-1004") {
			lang.SayString("powx436")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+504").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx437")
	} else {
		if !(rexsult.ToString() == "6.25E-1006") {
			lang.SayString("powx437")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx441")
	} else {
		if !(rexsult.ToString() == "1.5625E-998") {
			lang.SayString("powx441")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+335").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx442")
	} else {
		if !(rexsult.ToString() == "1.5625E-1001") {
			lang.SayString("powx442")
		}
	}
	rexsult, err = lang.RxFromString("0.04E+336").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx443")
	} else {
		if !(rexsult.ToString() == "1.5625E-1004") {
			lang.SayString("powx443")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+333").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx444")
	} else {
		if !(rexsult.ToString() == "6.4E-998") {
			lang.SayString("powx444")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx445")
	} else {
		if !(rexsult.ToString() == "6.4E-1001") {
			lang.SayString("powx445")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+335").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx446")
	} else {
		if !(rexsult.ToString() == "6.4E-1004") {
			lang.SayString("powx446")
		}
	}
	rexsult, err = lang.RxFromString("0.25E+336").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx447")
	} else {
		if !(rexsult.ToString() == "6.4E-1007") {
			lang.SayString("powx447")
		}
	}
	rexsult, err = lang.RxFromString("-0.04E+334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx448")
	} else {
		if !(rexsult.ToString() == "-1.5625E-998") {
			lang.SayString("powx448")
		}
	}
	rexsult, err = lang.RxFromString("-0.04E+335").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx449")
	} else {
		if !(rexsult.ToString() == "-1.5625E-1001") {
			lang.SayString("powx449")
		}
	}
	rexsult, err = lang.RxFromString("-0.04E+336").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx450")
	} else {
		if !(rexsult.ToString() == "-1.5625E-1004") {
			lang.SayString("powx450")
		}
	}
	rexsult, err = lang.RxFromString("-0.25E+333").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx451")
	} else {
		if !(rexsult.ToString() == "-6.4E-998") {
			lang.SayString("powx451")
		}
	}
	rexsult, err = lang.RxFromString("-0.25E+334").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx452")
	} else {
		if !(rexsult.ToString() == "-6.4E-1001") {
			lang.SayString("powx452")
		}
	}
	rexsult, err = lang.RxFromString("-0.25E+335").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx453")
	} else {
		if !(rexsult.ToString() == "-6.4E-1004") {
			lang.SayString("powx453")
		}
	}
	rexsult, err = lang.RxFromString("-0.25E+336").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx454")
	} else {
		if !(rexsult.ToString() == "-6.4E-1007") {
			lang.SayString("powx454")
		}
	}
	rexsult, err = lang.RxFromString("-0.04E+499").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx455")
	} else {
		if !(rexsult.ToString() == "6.25E-996") {
			lang.SayString("powx455")
		}
	}
	rexsult, err = lang.RxFromString("-0.04E+500").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx456")
	} else {
		if !(rexsult.ToString() == "6.25E-998") {
			lang.SayString("powx456")
		}
	}
	rexsult, err = lang.RxFromString("-0.04E+501").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx457")
	} else {
		if !(rexsult.ToString() == "6.25E-1000") {
			lang.SayString("powx457")
		}
	}
	rexsult, err = lang.RxFromString("-0.04E+502").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx458")
	} else {
		if !(rexsult.ToString() == "6.25E-1002") {
			lang.SayString("powx458")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx560")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx560")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("powx561")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx561")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx562")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx562")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("powx563")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx563")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx564")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx564")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("powx565")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx565")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx566")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx566")
		}
	}
	rexsult, err = lang.RxFromString("-1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-0"))
	if err != nil {
		lang.SayString("powx567")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx567")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx568")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx568")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx570")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx570")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx572")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx572")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx574")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx574")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx576")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx576")
		}
	}
	rexsult, err = lang.RxFromString("-0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx578")
	} else {
		if !(rexsult.ToString() == "0") {
			lang.SayString("powx578")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx701")
	} else {
		if !(rexsult.ToString() == "1.2345678E+10") {
			lang.SayString("powx701")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx702")
	} else {
		if !(rexsult.ToString() == "1.2345678E+9") {
			lang.SayString("powx702")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx703")
	} else {
		if !(rexsult.ToString() == "1.23456789E+9") {
			lang.SayString("powx703")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx704")
	} else {
		if !(rexsult.ToString() == "1.23456789E+9") {
			lang.SayString("powx704")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx705")
	} else {
		if !(rexsult.ToString() == "1.23456789E+10") {
			lang.SayString("powx705")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx706")
	} else {
		if !(rexsult.ToString() == "1.2345679E+9") {
			lang.SayString("powx706")
		}
	}
	rexsult, err = lang.RxFromString("12345678000").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx741")
	} else {
		if !(rexsult.ToString() == "12345678000") {
			lang.SayString("powx741")
		}
	}
	rexsult, err = lang.RxFromString("1234567800").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx742")
	} else {
		if !(rexsult.ToString() == "1234567800") {
			lang.SayString("powx742")
		}
	}
	rexsult, err = lang.RxFromString("1234567890").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx743")
	} else {
		if !(rexsult.ToString() == "1234567890") {
			lang.SayString("powx743")
		}
	}
	rexsult, err = lang.RxFromString("1234567891").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx744")
	} else {
		if !(rexsult.ToString() == "1234567891") {
			lang.SayString("powx744")
		}
	}
	rexsult, err = lang.RxFromString("12345678901").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx745")
	} else {
		if !(rexsult.ToString() == "12345678901") {
			lang.SayString("powx745")
		}
	}
	rexsult, err = lang.RxFromString("1234567896").OpPow(lang.RxSetWithDigit(15), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx746")
	} else {
		if !(rexsult.ToString() == "1234567896") {
			lang.SayString("powx746")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx163")
	} else {
		if !(rexsult.ToString() == "1E+999999") {
			lang.SayString("powx163")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx164")
	} else {
		if !(rexsult.ToString() == "1E+999998") {
			lang.SayString("powx164")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999997"))
	if err != nil {
		lang.SayString("powx165")
	} else {
		if !(rexsult.ToString() == "1E+999997") {
			lang.SayString("powx165")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("333333"))
	if err != nil {
		lang.SayString("powx166")
	} else {
		if !(rexsult.ToString() == "1E+333333") {
			lang.SayString("powx166")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1000000"))
	if err != nil {
		lang.SayString("powx183")
	} else {
		if !(rexsult.ToString() == "1.09651419E+845098") {
			lang.SayString("powx183")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1000001"))
	if err != nil {
		lang.SayString("powx184")
	} else {
		if !(rexsult.ToString() == "7.67559934E+845098") {
			lang.SayString("powx184")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1000001"))
	if err != nil {
		lang.SayString("powx186")
	} else {
		if !(rexsult.ToString() == "1.30282986E-845099") {
			lang.SayString("powx186")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1000000"))
	if err != nil {
		lang.SayString("powx187")
	} else {
		if !(rexsult.ToString() == "9.11980901E-845099") {
			lang.SayString("powx187")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-333333"))
	if err != nil {
		lang.SayString("powx118")
	} else {
		if !(rexsult.ToString() == "1E-333333") {
			lang.SayString("powx118")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999998"))
	if err != nil {
		lang.SayString("powx119")
	} else {
		if !(rexsult.ToString() == "1E-999998") {
			lang.SayString("powx119")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999999"))
	if err != nil {
		lang.SayString("powx120")
	} else {
		if !(rexsult.ToString() == "1E-999999") {
			lang.SayString("powx120")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx181")
	} else {
		if !(rexsult.ToString() == "2.23778406E+845096") {
			lang.SayString("powx181")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx182")
	} else {
		if !(rexsult.ToString() == "1.56644884E+845097") {
			lang.SayString("powx182")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999999"))
	if err != nil {
		lang.SayString("powx189")
	} else {
		if !(rexsult.ToString() == "6.38386631E-845098") {
			lang.SayString("powx189")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999998"))
	if err != nil {
		lang.SayString("powx190")
	} else {
		if !(rexsult.ToString() == "4.46870641E-845097") {
			lang.SayString("powx190")
		}
	}
	rexsult, err = lang.RxFromString("9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx277")
	} else {
		if !(rexsult.ToString() == "3.59084629E+954241") {
			lang.SayString("powx277")
		}
	}
	rexsult, err = lang.RxFromString("9.99999999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx278")
	} else {
		if !(rexsult.ToString() == "9.99000501E+999998") {
			lang.SayString("powx278")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx279")
	} else {
		if !(rexsult.ToString() == "1E+999999") {
			lang.SayString("powx279")
		}
	}
	rexsult, err = lang.RxFromString("10.0000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx280")
	} else {
		if !(rexsult.ToString() == "1.01005016E+999999") {
			lang.SayString("powx280")
		}
	}
	rexsult, err = lang.RxFromString("10.000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx281")
	} else {
		if !(rexsult.ToString() == "1.1051708E+999999") {
			lang.SayString("powx281")
		}
	}
	rexsult, err = lang.RxFromString("10.00001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx282")
	} else {
		if !(rexsult.ToString() == "2.71827775E+999999") {
			lang.SayString("powx282")
		}
	}
	rexsult, err = lang.RxFromString("10.0001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx283")
	} else {
		if !(rexsult.ToString() == "2.20251443E+1000003") {
			lang.SayString("powx283")
		}
	}
	rexsult, err = lang.RxFromString("11").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx285")
	} else {
		if !(rexsult.ToString() == "4.40317088E+1041391") {
			lang.SayString("powx285")
		}
	}
	rexsult, err = lang.RxFromString("12").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx286")
	} else {
		if !(rexsult.ToString() == "1.4684744E+1079180") {
			lang.SayString("powx286")
		}
	}
	rexsult, err = lang.RxFromString("999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx287")
	} else {
		if !(rexsult.ToString() == "3.08077864E+2999562") {
			lang.SayString("powx287")
		}
	}
	rexsult, err = lang.RxFromString("999999999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx288")
	} else {
		if !(rexsult.ToString() == "9.99000501E+8999990") {
			lang.SayString("powx288")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx290")
	} else {
		if !(rexsult.ToString() == "2.02006812E-301030") {
			lang.SayString("powx290")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx291")
	} else {
		if !(rexsult.ToString() == "1E-999999") {
			lang.SayString("powx291")
		}
	}
	rexsult, err = lang.RxFromString("0.09").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx292")
	} else {
		if !(rexsult.ToString() == "3.59084629E-1045757") {
			lang.SayString("powx292")
		}
	}
	rexsult, err = lang.RxFromString("0.05").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx293")
	} else {
		if !(rexsult.ToString() == "2.02006812E-1301029") {
			lang.SayString("powx293")
		}
	}
	rexsult, err = lang.RxFromString("0.01").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx294")
	} else {
		if !(rexsult.ToString() == "1E-1999998") {
			lang.SayString("powx294")
		}
	}
	rexsult, err = lang.RxFromString("0.0001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx295")
	} else {
		if !(rexsult.ToString() == "1E-3999996") {
			lang.SayString("powx295")
		}
	}
	rexsult, err = lang.RxFromString("0.0000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx297")
	} else {
		if !(rexsult.ToString() == "1E-6999993") {
			lang.SayString("powx297")
		}
	}
	rexsult, err = lang.RxFromString("0.0000000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx298")
	} else {
		if !(rexsult.ToString() == "1E-9999990") {
			lang.SayString("powx298")
		}
	}
	rexsult, err = lang.RxFromString("-9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx310")
	} else {
		if !(rexsult.ToString() == "-3.59084629E+954241") {
			lang.SayString("powx310")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx311")
	} else {
		if !(rexsult.ToString() == "-1E+999999") {
			lang.SayString("powx311")
		}
	}
	rexsult, err = lang.RxFromString("-10.0001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx312")
	} else {
		if !(rexsult.ToString() == "-2.20251443E+1000003") {
			lang.SayString("powx312")
		}
	}
	rexsult, err = lang.RxFromString("-10.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx313")
	} else {
		if !(rexsult.ToString() == "-2.34132266E+1004320") {
			lang.SayString("powx313")
		}
	}
	rexsult, err = lang.RxFromString("-11").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx314")
	} else {
		if !(rexsult.ToString() == "-4.40317088E+1041391") {
			lang.SayString("powx314")
		}
	}
	rexsult, err = lang.RxFromString("-12").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx315")
	} else {
		if !(rexsult.ToString() == "-1.4684744E+1079180") {
			lang.SayString("powx315")
		}
	}
	rexsult, err = lang.RxFromString("-999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx316")
	} else {
		if !(rexsult.ToString() == "-3.08077864E+2999562") {
			lang.SayString("powx316")
		}
	}
	rexsult, err = lang.RxFromString("-999999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx317")
	} else {
		if !(rexsult.ToString() == "-3.67879625E+5999993") {
			lang.SayString("powx317")
		}
	}
	rexsult, err = lang.RxFromString("-999999999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx318")
	} else {
		if !(rexsult.ToString() == "-9.99000501E+8999990") {
			lang.SayString("powx318")
		}
	}
	rexsult, err = lang.RxFromString("-0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx320")
	} else {
		if !(rexsult.ToString() == "-2.02006812E-301030") {
			lang.SayString("powx320")
		}
	}
	rexsult, err = lang.RxFromString("-0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx321")
	} else {
		if !(rexsult.ToString() == "-1E-999999") {
			lang.SayString("powx321")
		}
	}
	rexsult, err = lang.RxFromString("-0.09").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx322")
	} else {
		if !(rexsult.ToString() == "-3.59084629E-1045757") {
			lang.SayString("powx322")
		}
	}
	rexsult, err = lang.RxFromString("-0.05").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx323")
	} else {
		if !(rexsult.ToString() == "-2.02006812E-1301029") {
			lang.SayString("powx323")
		}
	}
	rexsult, err = lang.RxFromString("-0.01").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx324")
	} else {
		if !(rexsult.ToString() == "-1E-1999998") {
			lang.SayString("powx324")
		}
	}
	rexsult, err = lang.RxFromString("-0.0001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx325")
	} else {
		if !(rexsult.ToString() == "-1E-3999996") {
			lang.SayString("powx325")
		}
	}
	rexsult, err = lang.RxFromString("-0.0000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx327")
	} else {
		if !(rexsult.ToString() == "-1E-6999993") {
			lang.SayString("powx327")
		}
	}
	rexsult, err = lang.RxFromString("-0.0000000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx328")
	} else {
		if !(rexsult.ToString() == "-1E-9999990") {
			lang.SayString("powx328")
		}
	}
	rexsult, err = lang.RxFromString("-9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx330")
	} else {
		if !(rexsult.ToString() == "3.98982921E+954240") {
			lang.SayString("powx330")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx331")
	} else {
		if !(rexsult.ToString() == "1E+999998") {
			lang.SayString("powx331")
		}
	}
	rexsult, err = lang.RxFromString("-10.0001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx332")
	} else {
		if !(rexsult.ToString() == "2.2024924E+1000002") {
			lang.SayString("powx332")
		}
	}
	rexsult, err = lang.RxFromString("-10.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx333")
	} else {
		if !(rexsult.ToString() == "2.31814125E+1004319") {
			lang.SayString("powx333")
		}
	}
	rexsult, err = lang.RxFromString("-11").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx334")
	} else {
		if !(rexsult.ToString() == "4.00288262E+1041390") {
			lang.SayString("powx334")
		}
	}
	rexsult, err = lang.RxFromString("-12").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx335")
	} else {
		if !(rexsult.ToString() == "1.22372866E+1079179") {
			lang.SayString("powx335")
		}
	}
	rexsult, err = lang.RxFromString("-999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx336")
	} else {
		if !(rexsult.ToString() == "3.0838625E+2999559") {
			lang.SayString("powx336")
		}
	}
	rexsult, err = lang.RxFromString("-999999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx337")
	} else {
		if !(rexsult.ToString() == "3.67879993E+5999987") {
			lang.SayString("powx337")
		}
	}
	rexsult, err = lang.RxFromString("-999999999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx338")
	} else {
		if !(rexsult.ToString() == "9.99000502E+8999981") {
			lang.SayString("powx338")
		}
	}
	rexsult, err = lang.RxFromString("-0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx340")
	} else {
		if !(rexsult.ToString() == "4.04013624E-301030") {
			lang.SayString("powx340")
		}
	}
	rexsult, err = lang.RxFromString("-0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx341")
	} else {
		if !(rexsult.ToString() == "1E-999998") {
			lang.SayString("powx341")
		}
	}
	rexsult, err = lang.RxFromString("-0.09").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx342")
	} else {
		if !(rexsult.ToString() == "3.98982921E-1045756") {
			lang.SayString("powx342")
		}
	}
	rexsult, err = lang.RxFromString("-0.05").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx343")
	} else {
		if !(rexsult.ToString() == "4.04013624E-1301028") {
			lang.SayString("powx343")
		}
	}
	rexsult, err = lang.RxFromString("-0.01").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx344")
	} else {
		if !(rexsult.ToString() == "1E-1999996") {
			lang.SayString("powx344")
		}
	}
	rexsult, err = lang.RxFromString("-0.0001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx345")
	} else {
		if !(rexsult.ToString() == "1E-3999992") {
			lang.SayString("powx345")
		}
	}
	rexsult, err = lang.RxFromString("-0.0000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx347")
	} else {
		if !(rexsult.ToString() == "1E-6999986") {
			lang.SayString("powx347")
		}
	}
	rexsult, err = lang.RxFromString("-0.0000000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999998"))
	if err != nil {
		lang.SayString("powx348")
	} else {
		if !(rexsult.ToString() == "1E-9999980") {
			lang.SayString("powx348")
		}
	}
	rexsult, err = lang.RxFromString("1e-1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx350")
	} else {
		if !(rexsult.ToString() == "1E-500000") {
			lang.SayString("powx350")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999"))
	if err != nil {
		lang.SayString("powx351")
	} else {
		if !(rexsult.ToString() == "1E-1999998") {
			lang.SayString("powx351")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx352")
	} else {
		if !(rexsult.ToString() == "1E-1000000") {
			lang.SayString("powx352")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500001"))
	if err != nil {
		lang.SayString("powx353")
	} else {
		if !(rexsult.ToString() == "1E-1000002") {
			lang.SayString("powx353")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500002"))
	if err != nil {
		lang.SayString("powx354")
	} else {
		if !(rexsult.ToString() == "1E-1000004") {
			lang.SayString("powx354")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500003"))
	if err != nil {
		lang.SayString("powx355")
	} else {
		if !(rexsult.ToString() == "1E-1000006") {
			lang.SayString("powx355")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500004"))
	if err != nil {
		lang.SayString("powx356")
	} else {
		if !(rexsult.ToString() == "1E-1000008") {
			lang.SayString("powx356")
		}
	}
	rexsult, err = lang.RxFromString("0.010001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx360")
	} else {
		if !(rexsult.ToString() == "5.17176082E-999979") {
			lang.SayString("powx360")
		}
	}
	rexsult, err = lang.RxFromString("0.010000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx361")
	} else {
		if !(rexsult.ToString() == "1.05127109E-1000000") {
			lang.SayString("powx361")
		}
	}
	rexsult, err = lang.RxFromString("0.010000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500001"))
	if err != nil {
		lang.SayString("powx362")
	} else {
		if !(rexsult.ToString() == "1.0512712E-1000002") {
			lang.SayString("powx362")
		}
	}
	rexsult, err = lang.RxFromString("0.0100000009").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx363")
	} else {
		if !(rexsult.ToString() == "1.04602786E-1000000") {
			lang.SayString("powx363")
		}
	}
	rexsult, err = lang.RxFromString("0.0100000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx364")
	} else {
		if !(rexsult.ToString() == "1.00501252E-1000000") {
			lang.SayString("powx364")
		}
	}
	rexsult, err = lang.RxFromString("0.01").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx365")
	} else {
		if !(rexsult.ToString() == "1E-1000000") {
			lang.SayString("powx365")
		}
	}
	rexsult, err = lang.RxFromString("0.0099999999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx366")
	} else {
		if !(rexsult.ToString() == "9.95012479E-1000001") {
			lang.SayString("powx366")
		}
	}
	rexsult, err = lang.RxFromString("0.0099999998").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx367")
	} else {
		if !(rexsult.ToString() == "9.90049834E-1000001") {
			lang.SayString("powx367")
		}
	}
	rexsult, err = lang.RxFromString("0.0099999997").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx368")
	} else {
		if !(rexsult.ToString() == "9.85111939E-1000001") {
			lang.SayString("powx368")
		}
	}
	rexsult, err = lang.RxFromString("0.0099999996").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx369")
	} else {
		if !(rexsult.ToString() == "9.80198673E-1000001") {
			lang.SayString("powx369")
		}
	}
	rexsult, err = lang.RxFromString("0.009").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000"))
	if err != nil {
		lang.SayString("powx370")
	} else {
		if !(rexsult.ToString() == "1.79771012E-1022879") {
			lang.SayString("powx370")
		}
	}
	rexsult, err = lang.RxFromString("1e-1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx371")
	} else {
		if !(rexsult.ToString() == "1E+500000") {
			lang.SayString("powx371")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999999"))
	if err != nil {
		lang.SayString("powx372")
	} else {
		if !(rexsult.ToString() == "1E+1999998") {
			lang.SayString("powx372")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx373")
	} else {
		if !(rexsult.ToString() == "1E+1000000") {
			lang.SayString("powx373")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500001"))
	if err != nil {
		lang.SayString("powx374")
	} else {
		if !(rexsult.ToString() == "1E+1000002") {
			lang.SayString("powx374")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500002"))
	if err != nil {
		lang.SayString("powx375")
	} else {
		if !(rexsult.ToString() == "1E+1000004") {
			lang.SayString("powx375")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500003"))
	if err != nil {
		lang.SayString("powx376")
	} else {
		if !(rexsult.ToString() == "1E+1000006") {
			lang.SayString("powx376")
		}
	}
	rexsult, err = lang.RxFromString("1e-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500004"))
	if err != nil {
		lang.SayString("powx377")
	} else {
		if !(rexsult.ToString() == "1E+1000008") {
			lang.SayString("powx377")
		}
	}
	rexsult, err = lang.RxFromString("0.010001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx381")
	} else {
		if !(rexsult.ToString() == "1.93357743E+999978") {
			lang.SayString("powx381")
		}
	}
	rexsult, err = lang.RxFromString("0.010000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx382")
	} else {
		if !(rexsult.ToString() == "9.51229427E+999999") {
			lang.SayString("powx382")
		}
	}
	rexsult, err = lang.RxFromString("0.010000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500001"))
	if err != nil {
		lang.SayString("powx383")
	} else {
		if !(rexsult.ToString() == "9.51229332E+1000001") {
			lang.SayString("powx383")
		}
	}
	rexsult, err = lang.RxFromString("0.0100000009").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx384")
	} else {
		if !(rexsult.ToString() == "9.55997484E+999999") {
			lang.SayString("powx384")
		}
	}
	rexsult, err = lang.RxFromString("0.0100000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx385")
	} else {
		if !(rexsult.ToString() == "9.95012479E+999999") {
			lang.SayString("powx385")
		}
	}
	rexsult, err = lang.RxFromString("0.01").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx386")
	} else {
		if !(rexsult.ToString() == "1E+1000000") {
			lang.SayString("powx386")
		}
	}
	rexsult, err = lang.RxFromString("0.009999").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx387")
	} else {
		if !(rexsult.ToString() == "5.19768437E+1000021") {
			lang.SayString("powx387")
		}
	}
	rexsult, err = lang.RxFromString("100.000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000"))
	if err != nil {
		lang.SayString("powx388")
	} else {
		if !(rexsult.ToString() == "9.95012479E-1000001") {
			lang.SayString("powx388")
		}
	}
	rexsult, err = lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+1"))
	if err != nil {
		lang.SayString("powx501")
	} else {
		if !(rexsult.ToString() == "1E+20") {
			lang.SayString("powx501")
		}
	}
	rexsult, err = lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+2"))
	if err != nil {
		lang.SayString("powx502")
	} else {
		if !(rexsult.ToString() == "1E+200") {
			lang.SayString("powx502")
		}
	}
	rexsult, err = lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+3"))
	if err != nil {
		lang.SayString("powx503")
	} else {
		if !(rexsult.ToString() == "1E+2000") {
			lang.SayString("powx503")
		}
	}
	rexsult, err = lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+4"))
	if err != nil {
		lang.SayString("powx504")
	} else {
		if !(rexsult.ToString() == "1E+20000") {
			lang.SayString("powx504")
		}
	}
	rexsult, err = lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+5"))
	if err != nil {
		lang.SayString("powx505")
	} else {
		if !(rexsult.ToString() == "1E+200000") {
			lang.SayString("powx505")
		}
	}
	rexsult, err = lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+6"))
	if err != nil {
		lang.SayString("powx506")
	} else {
		if !(rexsult.ToString() == "1E+2000000") {
			lang.SayString("powx506")
		}
	}
	rexsult, err = lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+7"))
	if err != nil {
		lang.SayString("powx507")
	} else {
		if !(rexsult.ToString() == "1E+20000000") {
			lang.SayString("powx507")
		}
	}
	rexsult, err = lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+8"))
	if err != nil {
		lang.SayString("powx508")
	} else {
		if !(rexsult.ToString() == "1E+200000000") {
			lang.SayString("powx508")
		}
	}
	rexsult, err = lang.RxFromString("1.2347E-40").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx750")
	} else {
		if !(rexsult.ToString() == "1.5245E-80") {
			lang.SayString("powx750")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1063")
	} else {
		if !(rexsult.ToString() == "1E+999999999") {
			lang.SayString("powx1063")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999998"))
	if err != nil {
		lang.SayString("powx1064")
	} else {
		if !(rexsult.ToString() == "1E+999999998") {
			lang.SayString("powx1064")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999997"))
	if err != nil {
		lang.SayString("powx1065")
	} else {
		if !(rexsult.ToString() == "1E+999999997") {
			lang.SayString("powx1065")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("333333333"))
	if err != nil {
		lang.SayString("powx1066")
	} else {
		if !(rexsult.ToString() == "1E+333333333") {
			lang.SayString("powx1066")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-333333333"))
	if err != nil {
		lang.SayString("powx1118")
	} else {
		if !(rexsult.ToString() == "1E-333333333") {
			lang.SayString("powx1118")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999999998"))
	if err != nil {
		lang.SayString("powx1119")
	} else {
		if !(rexsult.ToString() == "1E-999999998") {
			lang.SayString("powx1119")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999999999"))
	if err != nil {
		lang.SayString("powx1120")
	} else {
		if !(rexsult.ToString() == "1E-999999999") {
			lang.SayString("powx1120")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999998"))
	if err != nil {
		lang.SayString("powx1181")
	} else {
		if !(rexsult.ToString() == "2.10892313E+845098038") {
			lang.SayString("powx1181")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1182")
	} else {
		if !(rexsult.ToString() == "1.47624619E+845098039") {
			lang.SayString("powx1182")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999999999"))
	if err != nil {
		lang.SayString("powx1189")
	} else {
		if !(rexsult.ToString() == "6.77393787E-845098040") {
			lang.SayString("powx1189")
		}
	}
	rexsult, err = lang.RxFromString("7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999999998"))
	if err != nil {
		lang.SayString("powx1190")
	} else {
		if !(rexsult.ToString() == "4.74175651E-845098039") {
			lang.SayString("powx1190")
		}
	}
	rexsult, err = lang.RxFromString("9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1280")
	} else {
		if !(rexsult.ToString() == "3.05550054E+954242508") {
			lang.SayString("powx1280")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1281")
	} else {
		if !(rexsult.ToString() == "1E+999999999") {
			lang.SayString("powx1281")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1290")
	} else {
		if !(rexsult.ToString() == "4.33559594E-301029996") {
			lang.SayString("powx1290")
		}
	}
	rexsult, err = lang.RxFromString("0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1291")
	} else {
		if !(rexsult.ToString() == "1E-999999999") {
			lang.SayString("powx1291")
		}
	}
	rexsult, err = lang.RxFromString("-9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1310")
	} else {
		if !(rexsult.ToString() == "-3.05550054E+954242508") {
			lang.SayString("powx1310")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1311")
	} else {
		if !(rexsult.ToString() == "-1E+999999999") {
			lang.SayString("powx1311")
		}
	}
	rexsult, err = lang.RxFromString("-0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1320")
	} else {
		if !(rexsult.ToString() == "-4.33559594E-301029996") {
			lang.SayString("powx1320")
		}
	}
	rexsult, err = lang.RxFromString("-0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999999"))
	if err != nil {
		lang.SayString("powx1321")
	} else {
		if !(rexsult.ToString() == "-1E-999999999") {
			lang.SayString("powx1321")
		}
	}
	rexsult, err = lang.RxFromString("-9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999998"))
	if err != nil {
		lang.SayString("powx1330")
	} else {
		if !(rexsult.ToString() == "3.3950006E+954242507") {
			lang.SayString("powx1330")
		}
	}
	rexsult, err = lang.RxFromString("-10").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999998"))
	if err != nil {
		lang.SayString("powx1331")
	} else {
		if !(rexsult.ToString() == "1E+999999998") {
			lang.SayString("powx1331")
		}
	}
	rexsult, err = lang.RxFromString("-0.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999998"))
	if err != nil {
		lang.SayString("powx1340")
	} else {
		if !(rexsult.ToString() == "8.67119187E-301029996") {
			lang.SayString("powx1340")
		}
	}
	rexsult, err = lang.RxFromString("-0.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("999999998"))
	if err != nil {
		lang.SayString("powx1341")
	} else {
		if !(rexsult.ToString() == "1E-999999998") {
			lang.SayString("powx1341")
		}
	}
	rexsult, err = lang.RxFromString("1e-1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000000"))
	if err != nil {
		lang.SayString("powx1350")
	} else {
		if !(rexsult.ToString() == "1E-500000000") {
			lang.SayString("powx1350")
		}
	}
	rexsult, err = lang.RxFromString("0.010001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000000"))
	if err != nil {
		lang.SayString("powx1360")
	} else {
		if !(rexsult.ToString() == "4.34941988E-999978287") {
			lang.SayString("powx1360")
		}
	}
	rexsult, err = lang.RxFromString("0.010000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000000"))
	if err != nil {
		lang.SayString("powx1361")
	} else {
		if !(rexsult.ToString() == "5.18469257E-999999979") {
			lang.SayString("powx1361")
		}
	}
	rexsult, err = lang.RxFromString("0.010000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000001"))
	if err != nil {
		lang.SayString("powx1362")
	} else {
		if !(rexsult.ToString() == "5.18469309E-999999981") {
			lang.SayString("powx1362")
		}
	}
	rexsult, err = lang.RxFromString("0.0100000009").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000000"))
	if err != nil {
		lang.SayString("powx1363")
	} else {
		if !(rexsult.ToString() == "3.49342003E-999999981") {
			lang.SayString("powx1363")
		}
	}
	rexsult, err = lang.RxFromString("0.0100000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("500000000"))
	if err != nil {
		lang.SayString("powx1364")
	} else {
		if !(rexsult.ToString() == "1.48413155E-999999998") {
			lang.SayString("powx1364")
		}
	}
	rexsult, err = lang.RxFromString("1e-1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000000"))
	if err != nil {
		lang.SayString("powx1371")
	} else {
		if !(rexsult.ToString() == "1E+500000000") {
			lang.SayString("powx1371")
		}
	}
	rexsult, err = lang.RxFromString("0.010001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000000"))
	if err != nil {
		lang.SayString("powx1381")
	} else {
		if !(rexsult.ToString() == "2.29915719E+999978286") {
			lang.SayString("powx1381")
		}
	}
	rexsult, err = lang.RxFromString("0.010000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000000"))
	if err != nil {
		lang.SayString("powx1382")
	} else {
		if !(rexsult.ToString() == "1.92875467E+999999978") {
			lang.SayString("powx1382")
		}
	}
	rexsult, err = lang.RxFromString("0.010000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000001"))
	if err != nil {
		lang.SayString("powx1383")
	} else {
		if !(rexsult.ToString() == "1.92875448E+999999980") {
			lang.SayString("powx1383")
		}
	}
	rexsult, err = lang.RxFromString("0.0100000009").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000000"))
	if err != nil {
		lang.SayString("powx1384")
	} else {
		if !(rexsult.ToString() == "2.86252438E+999999980") {
			lang.SayString("powx1384")
		}
	}
	rexsult, err = lang.RxFromString("0.0100000001").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-500000000"))
	if err != nil {
		lang.SayString("powx1385")
	} else {
		if !(rexsult.ToString() == "6.73794717E+999999997") {
			lang.SayString("powx1385")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("2.00000000"))
	if err != nil {
		lang.SayString("powx2002")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("powx2002")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("2.0000000000000001"))
	if err != nil {
		lang.SayString("powx2010")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("powx2010")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1234"))
	if err != nil {
		lang.SayString("powx2011")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2011")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(4), lang.RxFromString("1234"))
	if err != nil {
		lang.SayString("powx2012")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2012")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(3), lang.RxFromString("12.0"))
	if err != nil {
		lang.SayString("powx2016")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2016")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(3), lang.RxFromString("1.00"))
	if err != nil {
		lang.SayString("powx2018")
	} else {
		if !(rexsult.ToString() == "2") {
			lang.SayString("powx2018")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(3), lang.RxFromString("2.00"))
	if err != nil {
		lang.SayString("powx2019")
	} else {
		if !(rexsult.ToString() == "4") {
			lang.SayString("powx2019")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1.0000000001"))
	if err != nil {
		lang.SayString("powx2032")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2032")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1.0000000000001"))
	if err != nil {
		lang.SayString("powx2033")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2033")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1.0000001"))
	if err != nil {
		lang.SayString("powx2035")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2035")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1.0000000001"))
	if err != nil {
		lang.SayString("powx2036")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2036")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1.0000000000001"))
	if err != nil {
		lang.SayString("powx2037")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2037")
		}
	}
	rexsult, err = lang.RxFromString("1").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1.0000000000001"))
	if err != nil {
		lang.SayString("powx2038")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2038")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx2081")
	} else {
		if !(rexsult.ToString() == "8") {
			lang.SayString("powx2081")
		}
	}
	rexsult, err = lang.RxFromString("-2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx2082")
	} else {
		if !(rexsult.ToString() == "-8") {
			lang.SayString("powx2082")
		}
	}
	rexsult, err = lang.RxFromString("2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx2083")
	} else {
		if !(rexsult.ToString() == "0.125") {
			lang.SayString("powx2083")
		}
	}
	rexsult, err = lang.RxFromString("1.7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx2084")
	} else {
		if !(rexsult.ToString() == "69.7575744") {
			lang.SayString("powx2084")
		}
	}
	rexsult, err = lang.RxFromString("0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx2093")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2093")
		}
	}
	rexsult, err = lang.RxFromString("1E-7").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("71"))
	if err != nil {
		lang.SayString("powx2140")
	} else {
		if !(rexsult.ToString() == "1E-497") {
			lang.SayString("powx2140")
		}
	}
	rexsult, err = lang.RxFromString("0.003").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("71"))
	if err != nil {
		lang.SayString("powx2141")
	} else {
		if !(rexsult.ToString() == "7.509466514979725E-180") {
			lang.SayString("powx2141")
		}
	}
	rexsult, err = lang.RxFromString("0.7").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("71"))
	if err != nil {
		lang.SayString("powx2142")
	} else {
		if !(rexsult.ToString() == "1.004525211269079E-11") {
			lang.SayString("powx2142")
		}
	}
	rexsult, err = lang.RxFromString("1.2").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("71"))
	if err != nil {
		lang.SayString("powx2143")
	} else {
		if !(rexsult.ToString() == "418666.7483186515") {
			lang.SayString("powx2143")
		}
	}
	rexsult, err = lang.RxFromString("71").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("71"))
	if err != nil {
		lang.SayString("powx2144")
	} else {
		if !(rexsult.ToString() == "2.750063734834616E+131") {
			lang.SayString("powx2144")
		}
	}
	rexsult, err = lang.RxFromString("9E+9").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("71"))
	if err != nil {
		lang.SayString("powx2145")
	} else {
		if !(rexsult.ToString() == "5.639208733960173E+706") {
			lang.SayString("powx2145")
		}
	}
	rexsult, err = lang.RxFromString("0.003").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("220"))
	if err != nil {
		lang.SayString("powx2228")
	} else {
		if !(rexsult.ToString() == "9.261387130997877E-556") {
			lang.SayString("powx2228")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("385"))
	if err != nil {
		lang.SayString("powx2301")
	} else {
		if !(rexsult.ToString() == "1E+385") {
			lang.SayString("powx2301")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("384"))
	if err != nil {
		lang.SayString("powx2302")
	} else {
		if !(rexsult.ToString() == "1E+384") {
			lang.SayString("powx2302")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx2303")
	} else {
		if !(rexsult.ToString() == "1E+17") {
			lang.SayString("powx2303")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("16"))
	if err != nil {
		lang.SayString("powx2304")
	} else {
		if !(rexsult.ToString() == "1E+16") {
			lang.SayString("powx2304")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("15"))
	if err != nil {
		lang.SayString("powx2305")
	} else {
		if !(rexsult.ToString() == "1000000000000000") {
			lang.SayString("powx2305")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx2306")
	} else {
		if !(rexsult.ToString() == "10000000000") {
			lang.SayString("powx2306")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx2307")
	} else {
		if !(rexsult.ToString() == "100000") {
			lang.SayString("powx2307")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx2308")
	} else {
		if !(rexsult.ToString() == "10") {
			lang.SayString("powx2308")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx2309")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2309")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("powx2310")
	} else {
		if !(rexsult.ToString() == "0.1") {
			lang.SayString("powx2310")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx2311")
	} else {
		if !(rexsult.ToString() == "0.00001") {
			lang.SayString("powx2311")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx2312")
	} else {
		if !(rexsult.ToString() == "0.000001") {
			lang.SayString("powx2312")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx2313")
	} else {
		if !(rexsult.ToString() == "1E-7") {
			lang.SayString("powx2313")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("powx2314")
	} else {
		if !(rexsult.ToString() == "1E-8") {
			lang.SayString("powx2314")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx2315")
	} else {
		if !(rexsult.ToString() == "1E-9") {
			lang.SayString("powx2315")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("powx2316")
	} else {
		if !(rexsult.ToString() == "1E-10") {
			lang.SayString("powx2316")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-383"))
	if err != nil {
		lang.SayString("powx2317")
	} else {
		if !(rexsult.ToString() == "1E-383") {
			lang.SayString("powx2317")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-384"))
	if err != nil {
		lang.SayString("powx2318")
	} else {
		if !(rexsult.ToString() == "1E-384") {
			lang.SayString("powx2318")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-385"))
	if err != nil {
		lang.SayString("powx2319")
	} else {
		if !(rexsult.ToString() == "1E-385") {
			lang.SayString("powx2319")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-397"))
	if err != nil {
		lang.SayString("powx2320")
	} else {
		if !(rexsult.ToString() == "1E-397") {
			lang.SayString("powx2320")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-398"))
	if err != nil {
		lang.SayString("powx2321")
	} else {
		if !(rexsult.ToString() == "1E-398") {
			lang.SayString("powx2321")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-399"))
	if err != nil {
		lang.SayString("powx2322")
	} else {
		if !(rexsult.ToString() == "1E-399") {
			lang.SayString("powx2322")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("-400"))
	if err != nil {
		lang.SayString("powx2323")
	} else {
		if !(rexsult.ToString() == "1E-400") {
			lang.SayString("powx2323")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("0.0000"))
	if err != nil {
		lang.SayString("powx2351")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2351")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(4), lang.RxFromString("0.0000"))
	if err != nil {
		lang.SayString("powx2371")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx2371")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(34), lang.RxFromString("2.9999999999999999999999999999999997"))
	if err != nil {
		lang.SayString("powx2512")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("powx2512")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(34), lang.RxFromString("2.9999999999999999999999999999999998"))
	if err != nil {
		lang.SayString("powx2513")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("powx2513")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(34), lang.RxFromString("2.9999999999999999999999999999999999"))
	if err != nil {
		lang.SayString("powx2514")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("powx2514")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(34), lang.RxFromString("3.0000000000000000000000000000000000"))
	if err != nil {
		lang.SayString("powx2515")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("powx2515")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(34), lang.RxFromString("3.0000000000000000000000000000000001"))
	if err != nil {
		lang.SayString("powx2516")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("powx2516")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(34), lang.RxFromString("3.0000000000000000000000000000000002"))
	if err != nil {
		lang.SayString("powx2517")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("powx2517")
		}
	}
	rexsult, err = lang.RxFromString("10").OpPow(lang.RxSetWithDigit(34), lang.RxFromString("3.0000000000000000000000000000000003"))
	if err != nil {
		lang.SayString("powx2518")
	} else {
		if !(rexsult.ToString() == "1000") {
			lang.SayString("powx2518")
		}
	}
	rexsult, err = lang.RxFromString("-1.2").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4102")
	} else {
		if !(rexsult.ToString() == "-22.18611") {
			lang.SayString("powx4102")
		}
	}
	rexsult, err = lang.RxFromString("-1.3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4103")
	} else {
		if !(rexsult.ToString() == "-86.50416") {
			lang.SayString("powx4103")
		}
	}
	rexsult, err = lang.RxFromString("-1.2").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4122")
	} else {
		if !(rexsult.ToString() == "-22.18611") {
			lang.SayString("powx4122")
		}
	}
	rexsult, err = lang.RxFromString("-1.3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4123")
	} else {
		if !(rexsult.ToString() == "-86.50416") {
			lang.SayString("powx4123")
		}
	}
	rexsult, err = lang.RxFromString("-1.2").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4142")
	} else {
		if !(rexsult.ToString() == "-22.18611") {
			lang.SayString("powx4142")
		}
	}
	rexsult, err = lang.RxFromString("-1.3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4143")
	} else {
		if !(rexsult.ToString() == "-86.50416") {
			lang.SayString("powx4143")
		}
	}
	rexsult, err = lang.RxFromString("-1.2").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4162")
	} else {
		if !(rexsult.ToString() == "-22.18611") {
			lang.SayString("powx4162")
		}
	}
	rexsult, err = lang.RxFromString("-1.3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4163")
	} else {
		if !(rexsult.ToString() == "-86.50416") {
			lang.SayString("powx4163")
		}
	}
	rexsult, err = lang.RxFromString("-1.2").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4182")
	} else {
		if !(rexsult.ToString() == "-22.18611") {
			lang.SayString("powx4182")
		}
	}
	rexsult, err = lang.RxFromString("-1.3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4183")
	} else {
		if !(rexsult.ToString() == "-86.50416") {
			lang.SayString("powx4183")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("11"))
	if err != nil {
		lang.SayString("powx4184")
	} else {
		if !(rexsult.ToString() == "0.0004882813") {
			lang.SayString("powx4184")
		}
	}
	rexsult, err = lang.RxFromString("3.15").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx4185")
	} else {
		if !(rexsult.ToString() == "31.25588") {
			lang.SayString("powx4185")
		}
	}
	rexsult, err = lang.RxFromString("-1.2").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4202")
	} else {
		if !(rexsult.ToString() == "-22.18611") {
			lang.SayString("powx4202")
		}
	}
	rexsult, err = lang.RxFromString("-1.3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4203")
	} else {
		if !(rexsult.ToString() == "-86.50416") {
			lang.SayString("powx4203")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("11"))
	if err != nil {
		lang.SayString("powx4204")
	} else {
		if !(rexsult.ToString() == "0.0004882813") {
			lang.SayString("powx4204")
		}
	}
	rexsult, err = lang.RxFromString("3.15").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx4205")
	} else {
		if !(rexsult.ToString() == "31.25588") {
			lang.SayString("powx4205")
		}
	}
	rexsult, err = lang.RxFromString("-1.2").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4222")
	} else {
		if !(rexsult.ToString() == "-22.18611") {
			lang.SayString("powx4222")
		}
	}
	rexsult, err = lang.RxFromString("-1.3").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("17"))
	if err != nil {
		lang.SayString("powx4223")
	} else {
		if !(rexsult.ToString() == "-86.50416") {
			lang.SayString("powx4223")
		}
	}
	rexsult, err = lang.RxFromString("0.5").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("11"))
	if err != nil {
		lang.SayString("powx4224")
	} else {
		if !(rexsult.ToString() == "0.0004882813") {
			lang.SayString("powx4224")
		}
	}
	rexsult, err = lang.RxFromString("3.15").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx4225")
	} else {
		if !(rexsult.ToString() == "31.25588") {
			lang.SayString("powx4225")
		}
	}
	rexsult, err = lang.RxFromString("-3.15").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx4226")
	} else {
		if !(rexsult.ToString() == "-31.25588") {
			lang.SayString("powx4226")
		}
	}
	rexsult, err = lang.RxFromString("1.000001").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx4301")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx4301")
		}
	}
	rexsult, err = lang.RxFromString("1.000001").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("1.000000"))
	if err != nil {
		lang.SayString("powx4309")
	} else {
		if !(rexsult.ToString() == "1.000001") {
			lang.SayString("powx4309")
		}
	}
	rexsult, err = lang.RxFromString("1.000001").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx4321")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx4321")
		}
	}
	rexsult, err = lang.RxFromString("1.000001").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("1.000000"))
	if err != nil {
		lang.SayString("powx4329")
	} else {
		if !(rexsult.ToString() == "1.000001") {
			lang.SayString("powx4329")
		}
	}
	rexsult, err = lang.RxFromString("0.9999999").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx4341")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx4341")
		}
	}
	rexsult, err = lang.RxFromString("0.9999999").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("1.000000"))
	if err != nil {
		lang.SayString("powx4349")
	} else {
		if !(rexsult.ToString() == "0.9999999") {
			lang.SayString("powx4349")
		}
	}
	rexsult, err = lang.RxFromString("0.9999999").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx4361")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx4361")
		}
	}
	rexsult, err = lang.RxFromString("0.9999999").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("1.000000"))
	if err != nil {
		lang.SayString("powx4369")
	} else {
		if !(rexsult.ToString() == "0.9999999") {
			lang.SayString("powx4369")
		}
	}
	rexsult, err = lang.RxFromString("1e-101").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("powx4382")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("powx4382")
		}
	}
	rexsult, err = lang.RxFromString("1e-101").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("1.000000"))
	if err != nil {
		lang.SayString("powx4384")
	} else {
		if !(rexsult.ToString() == "1E-101") {
			lang.SayString("powx4384")
		}
	}
	rexsult, err = lang.RxFromString("1e-101").OpPow(lang.RxSetWithDigit(7), lang.RxFromString("2.000000"))
	if err != nil {
		lang.SayString("powx4386")
	} else {
		if !(rexsult.ToString() == "1E-202") {
			lang.SayString("powx4386")
		}
	}
	rexsult, err = lang.RxFromString("4953734675913.065314738743322579").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx3001")
	} else {
		if !(rexsult.ToString() == "24539487239343522246155890.99495") {
			lang.SayString("powx3001")
		}
	}
	rexsult, err = lang.RxFromString("9641.684323386955881595490347910E-844").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("powx3002")
	} else {
		if !(rexsult.ToString() == "1.338988152067180337738955757587E+6720") {
			lang.SayString("powx3002")
		}
	}
	rexsult, err = lang.RxFromString("-1.028048571628326871054964307774E+529").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3003")
	} else {
		if !(rexsult.ToString() == "-1.148333858253704284232780819739E+2645") {
			lang.SayString("powx3003")
		}
	}
	rexsult, err = lang.RxFromString("479084.8561808930525417735205519").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3004")
	} else {
		if !(rexsult.ToString() == "2.775233598021235973545933045837E+45") {
			lang.SayString("powx3004")
		}
	}
	rexsult, err = lang.RxFromString("-0363750788.573782205664349562931").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-3172"))
	if err != nil {
		lang.SayString("powx3005")
	} else {
		if !(rexsult.ToString() == "1.348625035575526599245847341566E-27155") {
			lang.SayString("powx3005")
		}
	}
	rexsult, err = lang.RxFromString("1381026551423669919010191878449").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-83"))
	if err != nil {
		lang.SayString("powx3006")
	} else {
		if !(rexsult.ToString() == "2.307977908106564299925193011052E-2502") {
			lang.SayString("powx3006")
		}
	}
	rexsult, err = lang.RxFromString("4627.026960423072127953556635585").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3007")
	} else {
		if !(rexsult.ToString() == "2.181684167222334934221407781701E-15") {
			lang.SayString("powx3007")
		}
	}
	rexsult, err = lang.RxFromString("75353574493.84484153484918212042").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3008")
	} else {
		if !(rexsult.ToString() == "1.276630670287906925570645490707E-98") {
			lang.SayString("powx3008")
		}
	}
	rexsult, err = lang.RxFromString("6907058.216435355874729592373011").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3009")
	} else {
		if !(rexsult.ToString() == "329518156646369505494.897135324") {
			lang.SayString("powx3009")
		}
	}
	rexsult, err = lang.RxFromString("-38949530427253.24030680468677190").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3010")
	} else {
		if !(rexsult.ToString() == "-1.359926959823071332599817363877E+95") {
			lang.SayString("powx3010")
		}
	}
	rexsult, err = lang.RxFromString("-0708069.025667471996378081482549").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-562842"))
	if err != nil {
		lang.SayString("powx3011")
	} else {
		if !(rexsult.ToString() == "5.636548923062838961163666321307E-3292669") {
			lang.SayString("powx3011")
		}
	}
	rexsult, err = lang.RxFromString("4055087.246994644709729942673976").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3012")
	} else {
		if !(rexsult.ToString() == "3.698274893849241116195795515302E-27") {
			lang.SayString("powx3012")
		}
	}
	rexsult, err = lang.RxFromString("4502895892520.396581348110906909E-512").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-816"))
	if err != nil {
		lang.SayString("powx3013")
	} else {
		if !(rexsult.ToString() == "5.63124716816538152028428346101E+407466") {
			lang.SayString("powx3013")
		}
	}
	rexsult, err = lang.RxFromString("467.6721295072628100260239179865").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3014")
	} else {
		if !(rexsult.ToString() == "0.000004572113694193221810609836080931") {
			lang.SayString("powx3014")
		}
	}
	rexsult, err = lang.RxFromString("-7634680140009571846155654339781").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3017")
	} else {
		if !(rexsult.ToString() == "-4.450128382072157170207584847831E+92") {
			lang.SayString("powx3017")
		}
	}
	rexsult, err = lang.RxFromString("262273.0222851186523650889896428E-624").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("74177"))
	if err != nil {
		lang.SayString("powx3018")
	} else {
		if !(rexsult.ToString() == "7.714950095090794263808002842046E-45884502") {
			lang.SayString("powx3018")
		}
	}
	rexsult, err = lang.RxFromString("883429.5928031498103637713570166E+765").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-43979"))
	if err != nil {
		lang.SayString("powx3020")
	} else {
		if !(rexsult.ToString() == "2.014274246041098936308200868298E-33905442") {
			lang.SayString("powx3020")
		}
	}
	rexsult, err = lang.RxFromString("24791301060.37938360567775506973").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3021")
	} else {
		if !(rexsult.ToString() == "4.307289712375673028996126249656E-63") {
			lang.SayString("powx3021")
		}
	}
	rexsult, err = lang.RxFromString("-930711443.9474781586162910776139").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-740"))
	if err != nil {
		lang.SayString("powx3022")
	} else {
		if !(rexsult.ToString() == "1.193603394165051899997226995178E-6637") {
			lang.SayString("powx3022")
		}
	}
	rexsult, err = lang.RxFromString("2358276428765.064191082773385539").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("214"))
	if err != nil {
		lang.SayString("powx3023")
	} else {
		if !(rexsult.ToString() == "5.435856480782850080741276939256E+2647") {
			lang.SayString("powx3023")
		}
	}
	rexsult, err = lang.RxFromString("-3.868744449795653651638308926987E+750").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("8270"))
	if err != nil {
		lang.SayString("powx3024")
	} else {
		if !(rexsult.ToString() == "1.600509322146129374577546618479E+6207359") {
			lang.SayString("powx3024")
		}
	}
	rexsult, err = lang.RxFromString("140422069.5863246490180206814374E-447").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3025")
	} else {
		if !(rexsult.ToString() == "1.304330899731988395473578425854E+2633") {
			lang.SayString("powx3025")
		}
	}
	rexsult, err = lang.RxFromString("75929096475.63450425339472559646E+153").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3026")
	} else {
		if !(rexsult.ToString() == "1.192136299657177324051477375561E-1475") {
			lang.SayString("powx3026")
		}
	}
	rexsult, err = lang.RxFromString("6312318309.142044953357460463732").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3027")
	} else {
		if !(rexsult.ToString() == "1.580762611512787720076533747265E-59") {
			lang.SayString("powx3027")
		}
	}
	rexsult, err = lang.RxFromString("93793652428100.52105928239469937").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx3028")
	} else {
		if !(rexsult.ToString() == "5.617732206663136654187263964365E+125") {
			lang.SayString("powx3028")
		}
	}
	rexsult, err = lang.RxFromString("98471198160.56524417578665886060").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-23994"))
	if err != nil {
		lang.SayString("powx3029")
	} else {
		if !(rexsult.ToString() == "3.455576582520200183817256784427E-263774") {
			lang.SayString("powx3029")
		}
	}
	rexsult, err = lang.RxFromString("329326552.0208398002250836592043").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-2451"))
	if err != nil {
		lang.SayString("powx3030")
	} else {
		if !(rexsult.ToString() == "1.980808996819343876559344394623E-20877") {
			lang.SayString("powx3030")
		}
	}
	rexsult, err = lang.RxFromString("-92980.68431371090354435763218439").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3031")
	} else {
		if !(rexsult.ToString() == "1.156683455371909793870207184337E-10") {
			lang.SayString("powx3031")
		}
	}
	rexsult, err = lang.RxFromString("12135817762.27858606259822256987E+738").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx3032")
	} else {
		if !(rexsult.ToString() == "6.929317520577437720457517499936E+7480") {
			lang.SayString("powx3032")
		}
	}
	rexsult, err = lang.RxFromString("37.27457578793521166809739140081").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-392550"))
	if err != nil {
		lang.SayString("powx3033")
	} else {
		if !(rexsult.ToString() == "8.726922581876932109954698478964E-616859") {
			lang.SayString("powx3033")
		}
	}
	rexsult, err = lang.RxFromString("-2787.980590304199878755265273703").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3034")
	} else {
		if !(rexsult.ToString() == "-1309266999233099220127139.440082") {
			lang.SayString("powx3034")
		}
	}
	rexsult, err = lang.RxFromString("-9890633.854609434943559831911276E+971").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3035")
	} else {
		if !(rexsult.ToString() == "1.022237362667592867768511487814E-1956") {
			lang.SayString("powx3035")
		}
	}
	rexsult, err = lang.RxFromString("3944570323.331121750661920475191").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-17360722"))
	if err != nil {
		lang.SayString("powx3036")
	} else {
		if !(rexsult.ToString() == "6.980067310956329947266905026242E-166593484") {
			lang.SayString("powx3036")
		}
	}
	rexsult, err = lang.RxFromString("19544.14018503427029002552872707").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx3037")
	} else {
		if !(rexsult.ToString() == "381973415.572271400929880255794") {
			lang.SayString("powx3037")
		}
	}
	rexsult, err = lang.RxFromString("-05.75485957937617757983513662981").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3038")
	} else {
		if !(rexsult.ToString() == "36325.23118223611421303238908472") {
			lang.SayString("powx3038")
		}
	}
	rexsult, err = lang.RxFromString("-4208820.898718069194008526302746").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3039")
	} else {
		if !(rexsult.ToString() == "5.558564783291260359142223337994E+39") {
			lang.SayString("powx3039")
		}
	}
	rexsult, err = lang.RxFromString("-70077195478066.30896979085821269E+549").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("4607"))
	if err != nil {
		lang.SayString("powx3040")
	} else {
		if !(rexsult.ToString() == "-3.731780212129336426822373547902E+2593029") {
			lang.SayString("powx3040")
		}
	}
	rexsult, err = lang.RxFromString("-442941.7541811527940918244383454").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-68126768"))
	if err != nil {
		lang.SayString("powx3041")
	} else {
		if !(rexsult.ToString() == "4.854798025822313707140856919099E-384667347") {
			lang.SayString("powx3041")
		}
	}
	rexsult, err = lang.RxFromString("-040726778711.8677615616711676159").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("299692"))
	if err != nil {
		lang.SayString("powx3042")
	} else {
		if !(rexsult.ToString() == "1.496947381410112723573826723795E+3179696") {
			lang.SayString("powx3042")
		}
	}
	rexsult, err = lang.RxFromString("-1934197520.738366912179143085955").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3043")
	} else {
		if !(rexsult.ToString() == "1.399597922275400947497855539475E+37") {
			lang.SayString("powx3043")
		}
	}
	rexsult, err = lang.RxFromString("813262.7723533833038829559646830").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3044")
	} else {
		if !(rexsult.ToString() == "1.859119568310997605545914895133E-18") {
			lang.SayString("powx3044")
		}
	}
	rexsult, err = lang.RxFromString("36105954884.94621434979365589311").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3045")
	} else {
		if !(rexsult.ToString() == "7.999297449713301719582732447386E+73") {
			lang.SayString("powx3045")
		}
	}
	rexsult, err = lang.RxFromString("-075537177538.1814516621962185490").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3046")
	} else {
		if !(rexsult.ToString() == "-4.310049518987988084595264617727E+32") {
			lang.SayString("powx3046")
		}
	}
	rexsult, err = lang.RxFromString("-4223765.415319564898840040697647").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3047")
	} else {
		if !(rexsult.ToString() == "-1.327090775863616939309569791138E-20") {
			lang.SayString("powx3047")
		}
	}
	rexsult, err = lang.RxFromString("-6468.903738522951259063099946195").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("powx3048")
	} else {
		if !(rexsult.ToString() == "3.261027724982089298030362367616E-31") {
			lang.SayString("powx3048")
		}
	}
	rexsult, err = lang.RxFromString("-9567221.183663236817239254783372E-203").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("1650"))
	if err != nil {
		lang.SayString("powx3049")
	} else {
		if !(rexsult.ToString() == "1.979675051437718045518315270002E-323432") {
			lang.SayString("powx3049")
		}
	}
	rexsult, err = lang.RxFromString("8812306098770.200752139142033569E-428").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3050")
	} else {
		if !(rexsult.ToString() == "6.843349527476967274129043949969E-1246") {
			lang.SayString("powx3050")
		}
	}
	rexsult, err = lang.RxFromString("80108033.12724838718736922500904").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3051")
	} else {
		if !(rexsult.ToString() == "4.723539145042336483008674060324E-56") {
			lang.SayString("powx3051")
		}
	}
	rexsult, err = lang.RxFromString("-37942846282.76101663789059003505").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3052")
	} else {
		if !(rexsult.ToString() == "3.351355986382646046773008753885E-64") {
			lang.SayString("powx3052")
		}
	}
	rexsult, err = lang.RxFromString("92659632115305.13735437728445541").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3053")
	} else {
		if !(rexsult.ToString() == "6.329121451953461546696051563323E+83") {
			lang.SayString("powx3053")
		}
	}
	rexsult, err = lang.RxFromString("2838948.589837595494152150647194").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3054")
	} else {
		if !(rexsult.ToString() == "5.235343334986059753096884080673E+38") {
			lang.SayString("powx3054")
		}
	}
	rexsult, err = lang.RxFromString("524995204523.6053307941775794287E+694").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx3055")
	} else {
		if !(rexsult.ToString() == "2.756199647727821911857160230849E+1411") {
			lang.SayString("powx3055")
		}
	}
	rexsult, err = lang.RxFromString("-57131573677452.15449921725097290").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3056")
	} else {
		if !(rexsult.ToString() == "-6.086686503752679375430019503679E+68") {
			lang.SayString("powx3056")
		}
	}
	rexsult, err = lang.RxFromString("90794826.55528018781830463383411").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx3057")
	} else {
		if !(rexsult.ToString() == "1.620669590532856523565742953997E-40") {
			lang.SayString("powx3057")
		}
	}
	rexsult, err = lang.RxFromString("58508794729.35191160840980489138").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-47060867"))
	if err != nil {
		lang.SayString("powx3058")
	} else {
		if !(rexsult.ToString() == "2.599345528648880340572206286149E-506714763") {
			lang.SayString("powx3058")
		}
	}
	rexsult, err = lang.RxFromString("-746104.0768078474426464219416332E+006").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx3059")
	} else {
		if !(rexsult.ToString() == "5.345571346302582882805035996696E+118") {
			lang.SayString("powx3059")
		}
	}
	rexsult, err = lang.RxFromString("55.99427632688387400403789659459E+119").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3060")
	} else {
		if !(rexsult.ToString() == "1.848022584764384077672041056396E-1087") {
			lang.SayString("powx3060")
		}
	}
	rexsult, err = lang.RxFromString("-41214265628.83801241467317270595").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx3061")
	} else {
		if !(rexsult.ToString() == "-41214265628.83801241467317270595") {
			lang.SayString("powx3061")
		}
	}
	rexsult, err = lang.RxFromString("89937.39749201095570357557430822").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3062")
	} else {
		if !(rexsult.ToString() == "4.280776267723913043050100934291E+39") {
			lang.SayString("powx3062")
		}
	}
	rexsult, err = lang.RxFromString("01712661.64677082156284125486943E+359").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3063")
	} else {
		if !(rexsult.ToString() == "2.523651803323047711735501944959E+2191") {
			lang.SayString("powx3063")
		}
	}
	rexsult, err = lang.RxFromString("-2647593306.528617691373470059213").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3064")
	} else {
		if !(rexsult.ToString() == "-1.096581914005902583413810201571E-66") {
			lang.SayString("powx3064")
		}
	}
	rexsult, err = lang.RxFromString("2904078690665765116603253099668E-329").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3065")
	} else {
		if !(rexsult.ToString() == "5.740389208842895561250128407803E+2089") {
			lang.SayString("powx3065")
		}
	}
	rexsult, err = lang.RxFromString("22094338972.39109726522477999515").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3066")
	} else {
		if !(rexsult.ToString() == "4.196391022354122686725315209967E-42") {
			lang.SayString("powx3066")
		}
	}
	rexsult, err = lang.RxFromString("-3374988581607586061255542201048").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3067")
	} else {
		if !(rexsult.ToString() == "1.683365657238878057620634207267E+244") {
			lang.SayString("powx3067")
		}
	}
	rexsult, err = lang.RxFromString("-84172558160661.35863831029352323").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-11272"))
	if err != nil {
		lang.SayString("powx3068")
	} else {
		if !(rexsult.ToString() == "3.004875538850321658947715753143E-156965") {
			lang.SayString("powx3068")
		}
	}
	rexsult, err = lang.RxFromString("-70046932324614.90596396237508541E-568").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3069")
	} else {
		if !(rexsult.ToString() == "-3.436903678302639677280508409829E-1663") {
			lang.SayString("powx3069")
		}
	}
	rexsult, err = lang.RxFromString("0004125384407.053782660115680886").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3070")
	} else {
		if !(rexsult.ToString() == "3.452568541597450106266555783362E-39") {
			lang.SayString("powx3070")
		}
	}
	rexsult, err = lang.RxFromString("-31823131.15691583393820628480997E-440").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx3071")
	} else {
		if !(rexsult.ToString() == "-3.347234803487575870321338308655E-3893") {
			lang.SayString("powx3071")
		}
	}
	rexsult, err = lang.RxFromString("55573867888.91575330563698128150").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("600"))
	if err != nil {
		lang.SayString("powx3072")
	} else {
		if !(rexsult.ToString() == "8.363240671070136278221965616973E+6446") {
			lang.SayString("powx3072")
		}
	}
	rexsult, err = lang.RxFromString("-5447727448431680878699555714796E-800").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3073")
	} else {
		if !(rexsult.ToString() == "-4.798183553278543065204833300725E-3847") {
			lang.SayString("powx3073")
		}
	}
	rexsult, err = lang.RxFromString("0418349404834.547110239542290134").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("9819916"))
	if err != nil {
		lang.SayString("powx3074")
	} else {
		if !(rexsult.ToString() == "1.983582121306690798480936158717E+114122538") {
			lang.SayString("powx3074")
		}
	}
	rexsult, err = lang.RxFromString("-262021.7565194737396448014286436").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("powx3075")
	} else {
		if !(rexsult.ToString() == "4.500918721033033032706782304195E-44") {
			lang.SayString("powx3075")
		}
	}
	rexsult, err = lang.RxFromString("48696050631.42565380301204592392E-505").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3076")
	} else {
		if !(rexsult.ToString() == "8.660017688773759463020340778853E+1482") {
			lang.SayString("powx3076")
		}
	}
	rexsult, err = lang.RxFromString("95316999.19440144356471126680708").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-60791"))
	if err != nil {
		lang.SayString("powx3077")
	} else {
		if !(rexsult.ToString() == "1.797083211825958027634744757793E-485062") {
			lang.SayString("powx3077")
		}
	}
	rexsult, err = lang.RxFromString("-5326702296402708234722215224979E-136").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("8032459"))
	if err != nil {
		lang.SayString("powx3078")
	} else {
		if !(rexsult.ToString() == "-3.279134167884360474119085034554E-845605407") {
			lang.SayString("powx3078")
		}
	}
	rexsult, err = lang.RxFromString("67.18750684079501575335482615780E-281").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3079")
	} else {
		if !(rexsult.ToString() == "6.18044407102311130081751840955E-1955") {
			lang.SayString("powx3079")
		}
	}
	rexsult, err = lang.RxFromString("2454.002078468928665008217727731").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3081")
	} else {
		if !(rexsult.ToString() == "218398452792293853786.926305442") {
			lang.SayString("powx3081")
		}
	}
	rexsult, err = lang.RxFromString("764578.5204849936912066033177429").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("64603"))
	if err != nil {
		lang.SayString("powx3082")
	} else {
		if !(rexsult.ToString() == "5.217717125896850565175391150717E+380086") {
			lang.SayString("powx3082")
		}
	}
	rexsult, err = lang.RxFromString("079203.7330103777716903518367560").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3083")
	} else {
		if !(rexsult.ToString() == "1.548692549503356788115682996756E+39") {
			lang.SayString("powx3083")
		}
	}
	rexsult, err = lang.RxFromString("-4278.581514688669249247007127899E-329").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3084")
	} else {
		if !(rexsult.ToString() == "-1.433834587801771244104676682986E-1627") {
			lang.SayString("powx3084")
		}
	}
	rexsult, err = lang.RxFromString("60867019.81764798845468445196869E+651").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3085")
	} else {
		if !(rexsult.ToString() == "5.085014897388871736767602086646E+3952") {
			lang.SayString("powx3085")
		}
	}
	rexsult, err = lang.RxFromString("18554417738217.62218590965803605E-382").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3086")
	} else {
		if !(rexsult.ToString() == "3.836842998295531899082688721531E+3318") {
			lang.SayString("powx3086")
		}
	}
	rexsult, err = lang.RxFromString("69073355517144.36356688642213839").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx3087")
	} else {
		if !(rexsult.ToString() == "2.472324890841334302628435461516E+138") {
			lang.SayString("powx3087")
		}
	}
	rexsult, err = lang.RxFromString("450282259072.8657099359104277477").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3088")
	} else {
		if !(rexsult.ToString() == "4.932082442194544671633570348838E-24") {
			lang.SayString("powx3088")
		}
	}
	rexsult, err = lang.RxFromString("954678411.7838149266455177850037").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("142989"))
	if err != nil {
		lang.SayString("powx3089")
	} else {
		if !(rexsult.ToString() == "6.125612386394735533471234509777E+1284020") {
			lang.SayString("powx3089")
		}
	}
	rexsult, err = lang.RxFromString("-9244530976.220812127155852389807E+557").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("541089"))
	if err != nil {
		lang.SayString("powx3090")
	} else {
		if !(rexsult.ToString() == "-4.840829839090601816099295547943E+306779003") {
			lang.SayString("powx3090")
		}
	}
	rexsult, err = lang.RxFromString("-75492024.20990197005974241975449").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("powx3091")
	} else {
		if !(rexsult.ToString() == "-1.324643246046162082348970735576E-8") {
			lang.SayString("powx3091")
		}
	}
	rexsult, err = lang.RxFromString("317747.6972215715434186596178036E-452").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx3092")
	} else {
		if !(rexsult.ToString() == "1.009635990896115043331231496209E-893") {
			lang.SayString("powx3092")
		}
	}
	rexsult, err = lang.RxFromString("-8.153334430358647134334545353427").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("powx3093")
	} else {
		if !(rexsult.ToString() == "7.702778966876727056635952801162E-10") {
			lang.SayString("powx3093")
		}
	}
	rexsult, err = lang.RxFromString("7.267345197492967332320456062961E-478").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3094")
	} else {
		if !(rexsult.ToString() == "2.027117616846668568108096583897E-2386") {
			lang.SayString("powx3094")
		}
	}
	rexsult, err = lang.RxFromString("-1223354029.862567054230912271171").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3095")
	} else {
		if !(rexsult.ToString() == "5.016689887192830666848068841227E+72") {
			lang.SayString("powx3095")
		}
	}
	rexsult, err = lang.RxFromString("285397644111.5655679961211349982E+645").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3096")
	} else {
		if !(rexsult.ToString() == "1.227719722087860401233030479451E-1313") {
			lang.SayString("powx3096")
		}
	}
	rexsult, err = lang.RxFromString("-4673112.663442366293812346653429").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("-3430"))
	if err != nil {
		lang.SayString("powx3097")
	} else {
		if !(rexsult.ToString() == "1.780563077186501355629969744956E-22877") {
			lang.SayString("powx3097")
		}
	}
	rexsult, err = lang.RxFromString("88.96492479681278079861456051103").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3098")
	} else {
		if !(rexsult.ToString() == "62643391.7307829022647475885897") {
			lang.SayString("powx3098")
		}
	}
	rexsult, err = lang.RxFromString("064326846.4286437304788069444326E-942").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("92"))
	if err != nil {
		lang.SayString("powx3099")
	} else {
		if !(rexsult.ToString() == "2.355526589209162447302444743229E-85946") {
			lang.SayString("powx3099")
		}
	}
	rexsult, err = lang.RxFromString("504507.0043949324433170405699360").OpPow(lang.RxSetWithDigit(31), lang.RxFromString("605388"))
	if err != nil {
		lang.SayString("powx3100")
	} else {
		if !(rexsult.ToString() == "2.336668965715246723667430661194E+3452447") {
			lang.SayString("powx3100")
		}
	}
	rexsult, err = lang.RxFromString("1.5283550163839789319142430253644").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3201")
	} else {
		if !(rexsult.ToString() == "0.42810618916584924451466691603128") {
			lang.SayString("powx3201")
		}
	}
	rexsult, err = lang.RxFromString("-622903030605.2867503937836507326").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3202")
	} else {
		if !(rexsult.ToString() == "-3.63867365977024043528133080643E+82") {
			lang.SayString("powx3202")
		}
	}
	rexsult, err = lang.RxFromString("-5675915.2465457487632250245209054").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3203")
	} else {
		if !(rexsult.ToString() == "-1.8978038060207777231389234721908E+47") {
			lang.SayString("powx3203")
		}
	}
	rexsult, err = lang.RxFromString("97.647321172555144900685605318635").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3204")
	} else {
		if !(rexsult.ToString() == "8877724578.7935312939231828719842") {
			lang.SayString("powx3204")
		}
	}
	rexsult, err = lang.RxFromString("-9717253267024.5380651435435603552").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3205")
	} else {
		if !(rexsult.ToString() == "-1.089856788008533778004132866133E-39") {
			lang.SayString("powx3205")
		}
	}
	rexsult, err = lang.RxFromString("-4.0817391717190128506083001702335E-767").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("12773"))
	if err != nil {
		lang.SayString("powx3206")
	} else {
		if !(rexsult.ToString() == "-2.1201633803170186724011454907679E-9789089") {
			lang.SayString("powx3206")
		}
	}
	rexsult, err = lang.RxFromString("68625322655934146845003028928647").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-60"))
	if err != nil {
		lang.SayString("powx3207")
	} else {
		if !(rexsult.ToString() == "6.4704731111943370171711131942603E-1911") {
			lang.SayString("powx3207")
		}
	}
	rexsult, err = lang.RxFromString("732515.76532049290815665856727641").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3208")
	} else {
		if !(rexsult.ToString() == "1.6468241050443471359358016585877E-53") {
			lang.SayString("powx3208")
		}
	}
	rexsult, err = lang.RxFromString("-30.458011942978338421676454778733").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx3209")
	} else {
		if !(rexsult.ToString() == "-3.8149797481405136042487643253109E-8") {
			lang.SayString("powx3209")
		}
	}
	rexsult, err = lang.RxFromString("-89640.094149414644660480286430462").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3210")
	} else {
		if !(rexsult.ToString() == "1.9274635591165405888724595165741E-30") {
			lang.SayString("powx3210")
		}
	}
	rexsult, err = lang.RxFromString("458653.1567870081810052917714259").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx3211")
	} else {
		if !(rexsult.ToString() == "210362718230.6879086511745242999") {
			lang.SayString("powx3211")
		}
	}
	rexsult, err = lang.RxFromString("913391.42744224458216174967853722").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3212")
	} else {
		if !(rexsult.ToString() == "1.1986327439971532470297300128074E-12") {
			lang.SayString("powx3212")
		}
	}
	rexsult, err = lang.RxFromString("-917591456829.12109027484399536567").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3213")
	} else {
		if !(rexsult.ToString() == "-1.2943505591853739240003453341911E-36") {
			lang.SayString("powx3213")
		}
	}
	rexsult, err = lang.RxFromString("34938410840645.913399699219228218").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("31"))
	if err != nil {
		lang.SayString("powx3214")
	} else {
		if !(rexsult.ToString() == "6.9566085958798732786509909683267E+419") {
			lang.SayString("powx3214")
		}
	}
	rexsult, err = lang.RxFromString("6034.7374411022598081745006769023E-517").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3215")
	} else {
		if !(rexsult.ToString() == "2.1977340597301840681528507640032E-1540") {
			lang.SayString("powx3215")
		}
	}
	rexsult, err = lang.RxFromString("-5565747671734.1686009705574503152").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-490"))
	if err != nil {
		lang.SayString("powx3216")
	} else {
		if !(rexsult.ToString() == "4.9371745297619526113991728953197E-6246") {
			lang.SayString("powx3216")
		}
	}
	rexsult, err = lang.RxFromString("319545511.89203495546689273564728E+036").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3217")
	} else {
		if !(rexsult.ToString() == "3.0647978448946294457985223953472E-134") {
			lang.SayString("powx3217")
		}
	}
	rexsult, err = lang.RxFromString("-36852134.84194296250843579428931").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-5830630"))
	if err != nil {
		lang.SayString("powx3218")
	} else {
		if !(rexsult.ToString() == "7.4103073243863751288190006633955E-44117245") {
			lang.SayString("powx3218")
		}
	}
	rexsult, err = lang.RxFromString("8.6021905001798578659275880018221E-374").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3219")
	} else {
		if !(rexsult.ToString() == "1.8262649155820433126240754123257E+1492") {
			lang.SayString("powx3219")
		}
	}
	rexsult, err = lang.RxFromString("-54863165.152174109720312887805017").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("736"))
	if err != nil {
		lang.SayString("powx3220")
	} else {
		if !(rexsult.ToString() == "1.2903643981679111625370174573639E+5696") {
			lang.SayString("powx3220")
		}
	}
	rexsult, err = lang.RxFromString("-3263743464517851012531708810307").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("2457206"))
	if err != nil {
		lang.SayString("powx3221")
	} else {
		if !(rexsult.ToString() == "1.1879385107215293840983911703413E+74978486") {
			lang.SayString("powx3221")
		}
	}
	rexsult, err = lang.RxFromString("2856586744.0548637797291151154902E-895").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx3222")
	} else {
		if !(rexsult.ToString() == "3.6180466753307072256807593988336E-8856") {
			lang.SayString("powx3222")
		}
	}
	rexsult, err = lang.RxFromString("5624157233.3433661009203529937625").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3223")
	} else {
		if !(rexsult.ToString() == "3.164788719630326254015832845903E+58") {
			lang.SayString("powx3223")
		}
	}
	rexsult, err = lang.RxFromString("30269.587755640502150977251770554").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3225")
	} else {
		if !(rexsult.ToString() == "25411630481547464128383.220368208") {
			lang.SayString("powx3225")
		}
	}
	rexsult, err = lang.RxFromString("47.525676459351505682005359699680E+704").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3226")
	} else {
		if !(rexsult.ToString() == "8.6782100393941226535150385475464E-4235") {
			lang.SayString("powx3226")
		}
	}
	rexsult, err = lang.RxFromString("67585560.562561229497720955705979").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("827"))
	if err != nil {
		lang.SayString("powx3228")
	} else {
		if !(rexsult.ToString() == "1.9462204583352191108781197788255E+6475") {
			lang.SayString("powx3228")
		}
	}
	rexsult, err = lang.RxFromString("6877386868.9498051860742298735156E-232").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("390"))
	if err != nil {
		lang.SayString("powx3229")
	} else {
		if !(rexsult.ToString() == "3.9368375020172007333542472647804E-86644") {
			lang.SayString("powx3229")
		}
	}
	rexsult, err = lang.RxFromString("-1647335.201144609256134925838937").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3230")
	} else {
		if !(rexsult.ToString() == "3.6849876990439502152784389237891E-13") {
			lang.SayString("powx3230")
		}
	}
	rexsult, err = lang.RxFromString("41407818140948.866630923934138155").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx3231")
	} else {
		if !(rexsult.ToString() == "8.2146348556648547525693528004081E-69") {
			lang.SayString("powx3231")
		}
	}
	rexsult, err = lang.RxFromString("-6.6547424012516834662011706165297").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3232")
	} else {
		if !(rexsult.ToString() == "0.000011513636283388791488128239232906") {
			lang.SayString("powx3232")
		}
	}
	rexsult, err = lang.RxFromString("-27627.758745381267780885067447671").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3233")
	} else {
		if !(rexsult.ToString() == "1.3101128009560812529198521922269E-9") {
			lang.SayString("powx3233")
		}
	}
	rexsult, err = lang.RxFromString("209819.74379099914752963711944307E-228").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3234")
	} else {
		if !(rexsult.ToString() == "5.1595828494111690910650919776705E+890") {
			lang.SayString("powx3234")
		}
	}
	rexsult, err = lang.RxFromString("2.3488457600415474270259330865184").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx3235")
	} else {
		if !(rexsult.ToString() == "2176.1583446147511579113022622255") {
			lang.SayString("powx3235")
		}
	}
	rexsult, err = lang.RxFromString("-5107586300197.9703941034404557409").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3236")
	} else {
		if !(rexsult.ToString() == "1.7753920894188022125919559565029E+76") {
			lang.SayString("powx3236")
		}
	}
	rexsult, err = lang.RxFromString("-70454070095869.70717871212601390").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-6200178"))
	if err != nil {
		lang.SayString("powx3237")
	} else {
		if !(rexsult.ToString() == "2.1499919329830163237787352070977E-85859483") {
			lang.SayString("powx3237")
		}
	}
	rexsult, err = lang.RxFromString("29119.220621511046558757900645228").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3238")
	} else {
		if !(rexsult.ToString() == "718983605328417461.32835984217255") {
			lang.SayString("powx3238")
		}
	}
	rexsult, err = lang.RxFromString("-5168.2214111091132913776042214525").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-5690274"))
	if err != nil {
		lang.SayString("powx3239")
	} else {
		if !(rexsult.ToString() == "4.2003839837606519720780508843429E-21129929") {
			lang.SayString("powx3239")
		}
	}
	rexsult, err = lang.RxFromString("33783.060857197067391462144517964").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2070"))
	if err != nil {
		lang.SayString("powx3240")
	} else {
		if !(rexsult.ToString() == "3.9181336001803008597293818984406E-9375") {
			lang.SayString("powx3240")
		}
	}
	rexsult, err = lang.RxFromString("42207435091050.840296353874733169E-905").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3241")
	} else {
		if !(rexsult.ToString() == "2.3862872940615283599573082966642E-6240") {
			lang.SayString("powx3241")
		}
	}
	rexsult, err = lang.RxFromString("-71800.806700868784841045406679641").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3242")
	} else {
		if !(rexsult.ToString() == "3.7625536850895480882178599428774E-20") {
			lang.SayString("powx3242")
		}
	}
	rexsult, err = lang.RxFromString("53627480801.631504892310016062160").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("328260"))
	if err != nil {
		lang.SayString("powx3243")
	} else {
		if !(rexsult.ToString() == "5.0890352043072920816489012230817E+3522028") {
			lang.SayString("powx3243")
		}
	}
	rexsult, err = lang.RxFromString("-5052601598.5559371338428368438728").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-97855372"))
	if err != nil {
		lang.SayString("powx3244")
	} else {
		if !(rexsult.ToString() == "3.7288295025319405771571918327697E-949541076") {
			lang.SayString("powx3244")
		}
	}
	rexsult, err = lang.RxFromString("4208134320733.7069742988228068191E+146").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3245")
	} else {
		if !(rexsult.ToString() == "3.1358723439830872127129821963857E+634") {
			lang.SayString("powx3245")
		}
	}
	rexsult, err = lang.RxFromString("-8.5077009657942581515590471189084E+308").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx3246")
	} else {
		if !(rexsult.ToString() == "1.9866536812573207868350640760678E+3089") {
			lang.SayString("powx3246")
		}
	}
	rexsult, err = lang.RxFromString("-9504.9703032286960790904181078063E+619").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-86"))
	if err != nil {
		lang.SayString("powx3247")
	} else {
		if !(rexsult.ToString() == "7.8747045500384098303517910143287E-53577") {
			lang.SayString("powx3247")
		}
	}
	rexsult, err = lang.RxFromString("-440220916.66716743026896931194749").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-102725"))
	if err != nil {
		lang.SayString("powx3248")
	} else {
		if !(rexsult.ToString() == "-8.5120729223452845846831015986365E-887922") {
			lang.SayString("powx3248")
		}
	}
	rexsult, err = lang.RxFromString("-46.250735086006350517943464758019").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx3249")
	} else {
		if !(rexsult.ToString() == "-46.250735086006350517943464758019") {
			lang.SayString("powx3249")
		}
	}
	rexsult, err = lang.RxFromString("-61641121299391.316420647102699627E+763").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3250")
	} else {
		if !(rexsult.ToString() == "-7.7833261179975532508748150708605E-6992") {
			lang.SayString("powx3250")
		}
	}
	rexsult, err = lang.RxFromString("96668419802749.555738010239087449E-838").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3251")
	} else {
		if !(rexsult.ToString() == "1.0701157625268896323611633350003E+1648") {
			lang.SayString("powx3251")
		}
	}
	rexsult, err = lang.RxFromString("-8534543911197995906031245719519E+124").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx3252")
	} else {
		if !(rexsult.ToString() == "7.2838439772166785429482995041337E+309") {
			lang.SayString("powx3252")
		}
	}
	rexsult, err = lang.RxFromString("-62663404777.352508979582846931050").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx3253")
	} else {
		if !(rexsult.ToString() == "-1.4897928814133059615670462753825E+97") {
			lang.SayString("powx3253")
		}
	}
	rexsult, err = lang.RxFromString("1.744601214474560992754529320172E-827").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3254")
	} else {
		if !(rexsult.ToString() == "3.285546809961528239499254251598E+1653") {
			lang.SayString("powx3254")
		}
	}
	rexsult, err = lang.RxFromString("0367191549036702224827734853471").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3255")
	} else {
		if !(rexsult.ToString() == "1.8179030119354318182493703269258E+118") {
			lang.SayString("powx3255")
		}
	}
	rexsult, err = lang.RxFromString("097704116.4492566721965710365073").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-96736"))
	if err != nil {
		lang.SayString("powx3256")
	} else {
		if !(rexsult.ToString() == "6.1575766428244048280595151303067E-772913") {
			lang.SayString("powx3256")
		}
	}
	rexsult, err = lang.RxFromString("19533298.147150158931958733807878").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3257")
	} else {
		if !(rexsult.ToString() == "2.1193595047638230427530063654613E+58") {
			lang.SayString("powx3257")
		}
	}
	rexsult, err = lang.RxFromString("-23765003221220177430797028997378").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3258")
	} else {
		if !(rexsult.ToString() == "1.7706154318483481190364979209436E-63") {
			lang.SayString("powx3258")
		}
	}
	rexsult, err = lang.RxFromString("-86863.276249466008289214762980838").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("532"))
	if err != nil {
		lang.SayString("powx3260")
	} else {
		if !(rexsult.ToString() == "2.889757918417383951985971021751E+2627") {
			lang.SayString("powx3260")
		}
	}
	rexsult, err = lang.RxFromString("-40707.169006771111855573524157083").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3261")
	} else {
		if !(rexsult.ToString() == "-5.3988802915897595722440392884051E-33") {
			lang.SayString("powx3261")
		}
	}
	rexsult, err = lang.RxFromString("-90838752568673.728630494658778003E+095").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-738"))
	if err != nil {
		lang.SayString("powx3262")
	} else {
		if !(rexsult.ToString() == "6.2500381480422211807167687339457E-80412") {
			lang.SayString("powx3262")
		}
	}
	rexsult, err = lang.RxFromString("-4245360967593.9206771555839718158E-682").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3263")
	} else {
		if !(rexsult.ToString() == "-1.3069414504933253288042820429894E+2008") {
			lang.SayString("powx3263")
		}
	}
	rexsult, err = lang.RxFromString("-3422145405774.0800213000547612240E-023").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-60811"))
	if err != nil {
		lang.SayString("powx3264")
	} else {
		if !(rexsult.ToString() == "-5.9763069766375402777445242662884E+636429") {
			lang.SayString("powx3264")
		}
	}
	rexsult, err = lang.RxFromString("-24521811.07649485796598387627478E-661").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3265")
	} else {
		if !(rexsult.ToString() == "-3.1190843559949184618590264246586E+5882") {
			lang.SayString("powx3265")
		}
	}
	rexsult, err = lang.RxFromString("-5042529937498.8944492248538951438").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3266")
	} else {
		if !(rexsult.ToString() == "6.4653782991800009492580180960839E+50") {
			lang.SayString("powx3266")
		}
	}
	rexsult, err = lang.RxFromString("71553220259.938950698030519906727E-496").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("754"))
	if err != nil {
		lang.SayString("powx3269")
	} else {
		if !(rexsult.ToString() == "2.4569976333599645534602060057777E-365800") {
			lang.SayString("powx3269")
		}
	}
	rexsult, err = lang.RxFromString("35572.960284795962697740953932508").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3270")
	} else {
		if !(rexsult.ToString() == "56963942247985404337401.149353169") {
			lang.SayString("powx3270")
		}
	}
	rexsult, err = lang.RxFromString("53035405018123760598334895413057E+818").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("powx3271")
	} else {
		if !(rexsult.ToString() == "5.6799053935427267944455848135618E-8498") {
			lang.SayString("powx3271")
		}
	}
	rexsult, err = lang.RxFromString("95.490751127249945886828257312118").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx3272")
	} else {
		if !(rexsult.ToString() == "63039548646186864162.847491534337") {
			lang.SayString("powx3272")
		}
	}
	rexsult, err = lang.RxFromString("69434850287.460788550936730296153").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3273")
	} else {
		if !(rexsult.ToString() == "4.3021939605842038995370443743844E-44") {
			lang.SayString("powx3273")
		}
	}
	rexsult, err = lang.RxFromString("-392.22739924621965621739098725407").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-65551275"))
	if err != nil {
		lang.SayString("powx3274")
	} else {
		if !(rexsult.ToString() == "-1.1015541697515606783291576863564E-170009718") {
			lang.SayString("powx3274")
		}
	}
	rexsult, err = lang.RxFromString("6413265.4423561191792972085539457").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("24514"))
	if err != nil {
		lang.SayString("powx3275")
	} else {
		if !(rexsult.ToString() == "5.4940184639434110212683638736116E+166868") {
			lang.SayString("powx3275")
		}
	}
	rexsult, err = lang.RxFromString("-6.9667706389122107760046184064057E+487").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("32"))
	if err != nil {
		lang.SayString("powx3276")
	} else {
		if !(rexsult.ToString() == "9.4843683271846893633025973101414E+15610") {
			lang.SayString("powx3276")
		}
	}
	rexsult, err = lang.RxFromString("378204716633.40024100602896057615").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3277")
	} else {
		if !(rexsult.ToString() == "1.8484988212401886887562779996737E-35") {
			lang.SayString("powx3277")
		}
	}
	rexsult, err = lang.RxFromString("-3554.5895974968741465654022772100E-073").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("powx3279")
	} else {
		if !(rexsult.ToString() == "3.2202875716019266933215387456197E-695") {
			lang.SayString("powx3279")
		}
	}
	rexsult, err = lang.RxFromString("750187685.63632608923397234391668").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3280")
	} else {
		if !(rexsult.ToString() == "2.3760176068829529745152188798557E+44") {
			lang.SayString("powx3280")
		}
	}
	rexsult, err = lang.RxFromString("30190034242853.251165951457589386E-028").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3281")
	} else {
		if !(rexsult.ToString() == "6.9009494305612589578332690692113E-117") {
			lang.SayString("powx3281")
		}
	}
	rexsult, err = lang.RxFromString("65.537942676774715953400768803539").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx3282")
	} else {
		if !(rexsult.ToString() == "65.537942676774715953400768803539") {
			lang.SayString("powx3282")
		}
	}
	rexsult, err = lang.RxFromString("8015272348677.5489394183881813700").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("949"))
	if err != nil {
		lang.SayString("powx3283")
	} else {
		if !(rexsult.ToString() == "6.5834151309758056859420039675639E+12245") {
			lang.SayString("powx3283")
		}
	}
	rexsult, err = lang.RxFromString("-32595333982.67068622120451804225").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3284")
	} else {
		if !(rexsult.ToString() == "-3.9092014148031739666525606093306E+73") {
			lang.SayString("powx3284")
		}
	}
	rexsult, err = lang.RxFromString("-506650.99395649907139204052441630").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("11"))
	if err != nil {
		lang.SayString("powx3286")
	} else {
		if !(rexsult.ToString() == "-5.6467412678809885333313818078829E+62") {
			lang.SayString("powx3286")
		}
	}
	rexsult, err = lang.RxFromString("4846835159.5922372307656000769395E-241").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-84"))
	if err != nil {
		lang.SayString("powx3287")
	} else {
		if !(rexsult.ToString() == "2.6394110381513074585638599361355E+19430") {
			lang.SayString("powx3287")
		}
	}
	rexsult, err = lang.RxFromString("-35.029311013822259358116556164908").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3288")
	} else {
		if !(rexsult.ToString() == "6.6416138040522124693495178218096E-7") {
			lang.SayString("powx3288")
		}
	}
	rexsult, err = lang.RxFromString("7606663750.6735265233044420887018E+166").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx3289")
	} else {
		if !(rexsult.ToString() == "3.9267106978887346373957314818178E-880") {
			lang.SayString("powx3289")
		}
	}
	rexsult, err = lang.RxFromString("-25677.829660831741274207326697052E-163").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3290")
	} else {
		if !(rexsult.ToString() == "-2.0605121420682764897867221992174E+1427") {
			lang.SayString("powx3290")
		}
	}
	rexsult, err = lang.RxFromString("97271576094.456406729671729224827").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3291")
	} else {
		if !(rexsult.ToString() == "1.0568858765852073181352309401343E-22") {
			lang.SayString("powx3291")
		}
	}
	rexsult, err = lang.RxFromString("41139789894.401826915757391777563").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3292")
	} else {
		if !(rexsult.ToString() == "5.9084812442741091550891451069919E-22") {
			lang.SayString("powx3292")
		}
	}
	rexsult, err = lang.RxFromString("-83310831287241.777598696853498149").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx3293")
	} else {
		if !(rexsult.ToString() == "-2.4916822606682624827900581728387E-70") {
			lang.SayString("powx3293")
		}
	}
	rexsult, err = lang.RxFromString("4506653461.4414974233678331771169").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3294")
	} else {
		if !(rexsult.ToString() == "2.6486272911486461102735412463189E-68") {
			lang.SayString("powx3294")
		}
	}
	rexsult, err = lang.RxFromString("23777.857951114967684767609401626").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("720760"))
	if err != nil {
		lang.SayString("powx3295")
	} else {
		if !(rexsult.ToString() == "1.8014168717277462614499869103226E+3154170") {
			lang.SayString("powx3295")
		}
	}
	rexsult, err = lang.RxFromString("-21337853323980858055292469611948").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3296")
	} else {
		if !(rexsult.ToString() == "9.4385298321304970306180652097874E+187") {
			lang.SayString("powx3296")
		}
	}
	rexsult, err = lang.RxFromString("-818409238.0423893439849743856947").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("756"))
	if err != nil {
		lang.SayString("powx3297")
	} else {
		if !(rexsult.ToString() == "1.6058883946373232750995543573461E+6738") {
			lang.SayString("powx3297")
		}
	}
	rexsult, err = lang.RxFromString("47567380384943.87013600286155046").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("65"))
	if err != nil {
		lang.SayString("powx3298")
	} else {
		if !(rexsult.ToString() == "1.0594982876763214301042437482634E+889") {
			lang.SayString("powx3298")
		}
	}
	rexsult, err = lang.RxFromString("-296615544.05897984545127115290251").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-5416115"))
	if err != nil {
		lang.SayString("powx3299")
	} else {
		if !(rexsult.ToString() == "-3.1678789872980061748117296226033E-45886377") {
			lang.SayString("powx3299")
		}
	}
	rexsult, err = lang.RxFromString("61391705914.046707180652185247584E+739").OpPow(lang.RxSetWithDigit(32), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3300")
	} else {
		if !(rexsult.ToString() == "3.0425105291210947473420999890124E-5249") {
			lang.SayString("powx3300")
		}
	}
	rexsult, err = lang.RxFromString("042.668056830485571428877212944418").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("powx3401")
	} else {
		if !(rexsult.ToString() == "0.0234367363850869744523417227148909") {
			lang.SayString("powx3401")
		}
	}
	rexsult, err = lang.RxFromString("3991.56734635183403814636354392163E-807").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("72"))
	if err != nil {
		lang.SayString("powx3404")
	} else {
		if !(rexsult.ToString() == "1.91570751569471030343614372896318E-57845") {
			lang.SayString("powx3404")
		}
	}
	rexsult, err = lang.RxFromString("-66.3393308595957726456416979163306").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3405")
	} else {
		if !(rexsult.ToString() == "-1284858888.2728582225918489666799") {
			lang.SayString("powx3405")
		}
	}
	rexsult, err = lang.RxFromString("-393606.873703567753255097095208112E+111").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3406")
	} else {
		if !(rexsult.ToString() == "6.45467904123103560528919747688443E-234") {
			lang.SayString("powx3406")
		}
	}
	rexsult, err = lang.RxFromString("465.351982159046525762715549761814").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("240445"))
	if err != nil {
		lang.SayString("powx3408")
	} else {
		if !(rexsult.ToString() == "5.48462986582864118931412123599581E+641454") {
			lang.SayString("powx3408")
		}
	}
	rexsult, err = lang.RxFromString("28066955004783.1076824222873384828").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3409")
	} else {
		if !(rexsult.ToString() == "4.88845689938951583020171325568218E+80") {
			lang.SayString("powx3409")
		}
	}
	rexsult, err = lang.RxFromString("28275236927392.4960902824105246047").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3410")
	} else {
		if !(rexsult.ToString() == "2.26057415546622161347322061281516E+40") {
			lang.SayString("powx3410")
		}
	}
	rexsult, err = lang.RxFromString("11791.8644211874630234271801789996").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("powx3411")
	} else {
		if !(rexsult.ToString() == "2.67510099318723516565332928253711E-33") {
			lang.SayString("powx3411")
		}
	}
	rexsult, err = lang.RxFromString("44.7085340739581668975502342787578").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-9337"))
	if err != nil {
		lang.SayString("powx3412")
	} else {
		if !(rexsult.ToString() == "2.01628956801735709137640528185479E-15410") {
			lang.SayString("powx3412")
		}
	}
	rexsult, err = lang.RxFromString("93354527428804.5458053295581965867E+576").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3413")
	} else {
		if !(rexsult.ToString() == "1.85687015691763406448005521221518E-5310") {
			lang.SayString("powx3413")
		}
	}
	rexsult, err = lang.RxFromString("-2.87155919781024108503670175443740").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx3415")
	} else {
		if !(rexsult.ToString() == "-13275.7774683251354527310820885737") {
			lang.SayString("powx3415")
		}
	}
	rexsult, err = lang.RxFromString("-010.693934338179479652178057279204E+188").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("26485"))
	if err != nil {
		lang.SayString("powx3416")
	} else {
		if !(rexsult.ToString() == "-5.09373179793691945131481396339599E+5006436") {
			lang.SayString("powx3416")
		}
	}
	rexsult, err = lang.RxFromString("611655569568.832698912762075889186").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx3417")
	} else {
		if !(rexsult.ToString() == "611655569568.832698912762075889186") {
			lang.SayString("powx3417")
		}
	}
	rexsult, err = lang.RxFromString("3457947.39062863674882672518304442").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3418")
	} else {
		if !(rexsult.ToString() == "8.36302195229701913376802373659526E-14") {
			lang.SayString("powx3418")
		}
	}
	rexsult, err = lang.RxFromString("-53308666960535.7393391289364591513").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3419")
	} else {
		if !(rexsult.ToString() == "-8.17363502380497033342380498988958E-97") {
			lang.SayString("powx3419")
		}
	}
	rexsult, err = lang.RxFromString("9804385273.49533524416415189990857").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("84"))
	if err != nil {
		lang.SayString("powx3421")
	} else {
		if !(rexsult.ToString() == "1.90244010779692739037080418507909E+839") {
			lang.SayString("powx3421")
		}
	}
	rexsult, err = lang.RxFromString("-5234910986592.18801727046580014273E-547").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3422")
	} else {
		if !(rexsult.ToString() == "4.85896970703117149235935037271084E+3205") {
			lang.SayString("powx3422")
		}
	}
	rexsult, err = lang.RxFromString("698416560151955285929747633786867E-495").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3423")
	} else {
		if !(rexsult.ToString() == "1.66177661007189430761396979787413E-2311") {
			lang.SayString("powx3423")
		}
	}
	rexsult, err = lang.RxFromString("107635.497735316515080720330536027").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3424")
	} else {
		if !(rexsult.ToString() == "7.45037111502910487803432806334714E-21") {
			lang.SayString("powx3424")
		}
	}
	rexsult, err = lang.RxFromString("-32174291345686.5371446616670961807").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3425")
	} else {
		if !(rexsult.ToString() == "1.14834377656109143210058690590666E+108") {
			lang.SayString("powx3425")
		}
	}
	rexsult, err = lang.RxFromString("-8151730494.53190523620899410544099E+688").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3426")
	} else {
		if !(rexsult.ToString() == "-6.291463527748424483752752821837E-6282") {
			lang.SayString("powx3426")
		}
	}
	rexsult, err = lang.RxFromString("1.33649801345976199708341799505220").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3427")
	} else {
		if !(rexsult.ToString() == "0.175464835912284900180305028965188") {
			lang.SayString("powx3427")
		}
	}
	rexsult, err = lang.RxFromString("67762238162788.6551061476018185196").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3428")
	} else {
		if !(rexsult.ToString() == "1.03293631708006509074972764670281E-83") {
			lang.SayString("powx3428")
		}
	}
	rexsult, err = lang.RxFromString("4286562.76568866751577306056498271").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3429")
	} else {
		if !(rexsult.ToString() == "6.20376193064412081058181881805108E+39") {
			lang.SayString("powx3429")
		}
	}
	rexsult, err = lang.RxFromString("-765782.827432642697305644096365566").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("67"))
	if err != nil {
		lang.SayString("powx3430")
	} else {
		if !(rexsult.ToString() == "-1.71821200770749773595473594136582E+394") {
			lang.SayString("powx3430")
		}
	}
	rexsult, err = lang.RxFromString("46.2835931916106252756465724211276").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("59"))
	if err != nil {
		lang.SayString("powx3431")
	} else {
		if !(rexsult.ToString() == "1.82052645780601002671007943923993E+98") {
			lang.SayString("powx3431")
		}
	}
	rexsult, err = lang.RxFromString("-3029555.82298840234029474459694644").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("powx3432")
	} else {
		if !(rexsult.ToString() == "-2.1498622479043130256134010074636E+58") {
			lang.SayString("powx3432")
		}
	}
	rexsult, err = lang.RxFromString("-0138466789523.10694176543700501945E-948").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3433")
	} else {
		if !(rexsult.ToString() == "-5.09012109092637525843636056746667E-4685") {
			lang.SayString("powx3433")
		}
	}
	rexsult, err = lang.RxFromString("-9593566466.96690575714244442109870").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3434")
	} else {
		if !(rexsult.ToString() == "-1.45271091841882960010964421066745E-90") {
			lang.SayString("powx3434")
		}
	}
	rexsult, err = lang.RxFromString("-3189.30765477670526823106100241863E-898").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3435")
	} else {
		if !(rexsult.ToString() == "1.05239431027683904514311527228736E-5367") {
			lang.SayString("powx3435")
		}
	}
	rexsult, err = lang.RxFromString("-17084552395.6714834680088150543965").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3436")
	} else {
		if !(rexsult.ToString() == "4.02141014977177984123011868387622E-62") {
			lang.SayString("powx3436")
		}
	}
	rexsult, err = lang.RxFromString("034956830.349823306815911887469760").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("powx3437")
	} else {
		if !(rexsult.ToString() == "5.48034272566098493462169431762597E-46") {
			lang.SayString("powx3437")
		}
	}
	rexsult, err = lang.RxFromString("-763.440067781256632695791981893608").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("20"))
	if err != nil {
		lang.SayString("powx3438")
	} else {
		if !(rexsult.ToString() == "4.52375407727336769552481661250924E+57") {
			lang.SayString("powx3438")
		}
	}
	rexsult, err = lang.RxFromString("-510472027868440667684575147556654E+789").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("powx3439")
	} else {
		if !(rexsult.ToString() == "4.61079266619522147262600755274182E+6573") {
			lang.SayString("powx3439")
		}
	}
	rexsult, err = lang.RxFromString("070304761.560517086676993503034828E-094").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3440")
	} else {
		if !(rexsult.ToString() == "2.02316135427631488479902919959627E+172") {
			lang.SayString("powx3440")
		}
	}
	rexsult, err = lang.RxFromString("-0970725697662.27605454336231195463").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-4541"))
	if err != nil {
		lang.SayString("powx3441")
	} else {
		if !(rexsult.ToString() == "-3.93253217883893131860607846766852E-54434") {
			lang.SayString("powx3441")
		}
	}
	rexsult, err = lang.RxFromString("-808178238631844268316111259558675").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-598400"))
	if err != nil {
		lang.SayString("powx3442")
	} else {
		if !(rexsult.ToString() == "5.24915585115968730741156502396947E-19691853") {
			lang.SayString("powx3442")
		}
	}
	rexsult, err = lang.RxFromString("-9.90826595069053564311371766315200").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-32"))
	if err != nil {
		lang.SayString("powx3443")
	} else {
		if !(rexsult.ToString() == "1.34299698259038003011439568004625E-32") {
			lang.SayString("powx3443")
		}
	}
	rexsult, err = lang.RxFromString("-196925.469891897719160698483752907").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-41269"))
	if err != nil {
		lang.SayString("powx3444")
	} else {
		if !(rexsult.ToString() == "-2.85288213213712886396300005819329E-218491") {
			lang.SayString("powx3444")
		}
	}
	rexsult, err = lang.RxFromString("421071135212152225162086005824310").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("powx3445")
	} else {
		if !(rexsult.ToString() == "421071135212152225162086005824310") {
			lang.SayString("powx3445")
		}
	}
	rexsult, err = lang.RxFromString("1249441.46421514282301182772247227").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3446")
	} else {
		if !(rexsult.ToString() == "5.12686942572191282348415024932322E-19") {
			lang.SayString("powx3446")
		}
	}
	rexsult, err = lang.RxFromString("74815000.4716875558358937279052903").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3447")
	} else {
		if !(rexsult.ToString() == "7.62218032252683815537906972439985E-56") {
			lang.SayString("powx3447")
		}
	}
	rexsult, err = lang.RxFromString("-1683993.51210241555668790556759021").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3448")
	} else {
		if !(rexsult.ToString() == "-2.60385683509956889000676113860292E-44") {
			lang.SayString("powx3448")
		}
	}
	rexsult, err = lang.RxFromString("-763.148530974741766171756970448158").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("517371"))
	if err != nil {
		lang.SayString("powx3449")
	} else {
		if !(rexsult.ToString() == "-2.18011932265402416040845295078195E+1491378") {
			lang.SayString("powx3449")
		}
	}
	rexsult, err = lang.RxFromString("-77.5841338812312523460591226178754").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3450")
	} else {
		if !(rexsult.ToString() == "-9.81846856873938549466341693997829E-18") {
			lang.SayString("powx3450")
		}
	}
	rexsult, err = lang.RxFromString("5176295309.89943746236102209837813").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-129733"))
	if err != nil {
		lang.SayString("powx3451")
	} else {
		if !(rexsult.ToString() == "1.46969655494208590580439170550344E-1260229") {
			lang.SayString("powx3451")
		}
	}
	rexsult, err = lang.RxFromString("4471634841166.90197229286530307326E+172").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3452")
	} else {
		if !(rexsult.ToString() == "8.94126556389673498386397569249516E+553") {
			lang.SayString("powx3452")
		}
	}
	rexsult, err = lang.RxFromString("-8189130.15945231049670285726774317").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3453")
	} else {
		if !(rexsult.ToString() == "-549178241054875982731.000937875885") {
			lang.SayString("powx3453")
		}
	}
	rexsult, err = lang.RxFromString("-254346232981353293392888785643245").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-764"))
	if err != nil {
		lang.SayString("powx3454")
	} else {
		if !(rexsult.ToString() == "1.79913300198400270448876854118754E-24758") {
			lang.SayString("powx3454")
		}
	}
	rexsult, err = lang.RxFromString("-2875.36745499544177964804277329726").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-13188"))
	if err != nil {
		lang.SayString("powx3455")
	} else {
		if !(rexsult.ToString() == "5.65018522709353395134776077040681E-45614") {
			lang.SayString("powx3455")
		}
	}
	rexsult, err = lang.RxFromString("-7.26927570984219944693602140223103").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx3456")
	} else {
		if !(rexsult.ToString() == "52.8423693457018126451998096211036") {
			lang.SayString("powx3456")
		}
	}
	rexsult, err = lang.RxFromString("-28567151.6868762752241056540028515E+498").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-4470"))
	if err != nil {
		lang.SayString("powx3457")
	} else {
		if !(rexsult.ToString() == "1.88267132619661347172657450614707E-2259388") {
			lang.SayString("powx3457")
		}
	}
	rexsult, err = lang.RxFromString("7191753.79974136447257468282073718").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("81"))
	if err != nil {
		lang.SayString("powx3458")
	} else {
		if !(rexsult.ToString() == "2.5329098313856148261255740414876E+555") {
			lang.SayString("powx3458")
		}
	}
	rexsult, err = lang.RxFromString("502975804.069864536104621700404770").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("685"))
	if err != nil {
		lang.SayString("powx3459")
	} else {
		if !(rexsult.ToString() == "3.62876716573623552761739177592677E+5960") {
			lang.SayString("powx3459")
		}
	}
	rexsult, err = lang.RxFromString("1505368.42063731861590460453659570").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-465243"))
	if err != nil {
		lang.SayString("powx3460")
	} else {
		if !(rexsult.ToString() == "8.51576386928398060125488059890328E-2874106") {
			lang.SayString("powx3460")
		}
	}
	rexsult, err = lang.RxFromString("69225023.2850142784063416137144829").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("8584050"))
	if err != nil {
		lang.SayString("powx3461")
	} else {
		if !(rexsult.ToString() == "3.57838696310427404863081521110295E+67301210") {
			lang.SayString("powx3461")
		}
	}
	rexsult, err = lang.RxFromString("-527566.521273992424649346837337602E-408").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-834663"))
	if err != nil {
		lang.SayString("powx3463")
	} else {
		if !(rexsult.ToString() == "-8.35767913594339088637487996295228E+335766330") {
			lang.SayString("powx3463")
		}
	}
	rexsult, err = lang.RxFromString("69065510.0459653699418083455335366").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3464")
	} else {
		if !(rexsult.ToString() == "7.49598249763416483824919118973567E+54") {
			lang.SayString("powx3464")
		}
	}
	rexsult, err = lang.RxFromString("-2921982.82411285505229122890041841").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3465")
	} else {
		if !(rexsult.ToString() == "-5.49865394870631248479668782154131E-46") {
			lang.SayString("powx3465")
		}
	}
	rexsult, err = lang.RxFromString("4.51117459466491451401835756593793").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("3873385"))
	if err != nil {
		lang.SayString("powx3466")
	} else {
		if !(rexsult.ToString() == "4.58081466040573540591639046524269E+2534315") {
			lang.SayString("powx3466")
		}
	}
	rexsult, err = lang.RxFromString("49553763474698.8118661758811091120").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3467")
	} else {
		if !(rexsult.ToString() == "6.02985091099730236635954801474802E+54") {
			lang.SayString("powx3467")
		}
	}
	rexsult, err = lang.RxFromString("755985583344.379951506071499170749E+956").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3468")
	} else {
		if !(rexsult.ToString() == "1.41121958516547725677142981375469E+6775") {
			lang.SayString("powx3468")
		}
	}
	rexsult, err = lang.RxFromString("5583903.18072100716072011264668631").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3470")
	} else {
		if !(rexsult.ToString() == "5.42861943589418603298670454702901E+33") {
			lang.SayString("powx3470")
		}
	}
	rexsult, err = lang.RxFromString("-955638397975240685017992436116257E+020").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx3471")
	} else {
		if !(rexsult.ToString() == "-1.25467730420304189161768408462414E-265") {
			lang.SayString("powx3471")
		}
	}
	rexsult, err = lang.RxFromString("-136243796098020983814115584402407E+796").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("powx3472")
	} else {
		if !(rexsult.ToString() == "-8.71399185551742199752832286984005E+5796") {
			lang.SayString("powx3472")
		}
	}
	rexsult, err = lang.RxFromString("-808498.482718304598213092937543934E+521").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("48005"))
	if err != nil {
		lang.SayString("powx3473")
	} else {
		if !(rexsult.ToString() == "-1.38177866825839934644922551882306E+25294203") {
			lang.SayString("powx3473")
		}
	}
	rexsult, err = lang.RxFromString("-812.266340554281305985524813074211E+396").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3474")
	} else {
		if !(rexsult.ToString() == "-1.86596988030914616216741808216469E-1197") {
			lang.SayString("powx3474")
		}
	}
	rexsult, err = lang.RxFromString("-929889.720905183397678866648217793E+134").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("powx3475")
	} else {
		if !(rexsult.ToString() == "-1.24367143370300189518778505830181E-420") {
			lang.SayString("powx3475")
		}
	}
	rexsult, err = lang.RxFromString("7812758113817.99135851825223122508").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3477")
	} else {
		if !(rexsult.ToString() == "4.76884421816246896090414314934253E+38") {
			lang.SayString("powx3477")
		}
	}
	rexsult, err = lang.RxFromString("489191747.148674326757767356626016").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("1136942"))
	if err != nil {
		lang.SayString("powx3478")
	} else {
		if !(rexsult.ToString() == "5.91014423217936361945997577207893E+9879433") {
			lang.SayString("powx3478")
		}
	}
	rexsult, err = lang.RxFromString("-599369540.373174482335865567937853E+289").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("powx3479")
	} else {
		if !(rexsult.ToString() == "7.74856580646291366270329165540958E-1192") {
			lang.SayString("powx3479")
		}
	}
	rexsult, err = lang.RxFromString("-3376883870.85961692148022521960139").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3480")
	} else {
		if !(rexsult.ToString() == "-1.9970416371836115312573575617928E-67") {
			lang.SayString("powx3480")
		}
	}
	rexsult, err = lang.RxFromString("58.6776780370008364590621772421025").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("powx3481")
	} else {
		if !(rexsult.ToString() == "3443.0698998139303363200831350523") {
			lang.SayString("powx3481")
		}
	}
	rexsult, err = lang.RxFromString("4099616339.96249499552808575717579").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("291"))
	if err != nil {
		lang.SayString("powx3482")
	} else {
		if !(rexsult.ToString() == "2.03364757877800497409765979877258E+2797") {
			lang.SayString("powx3482")
		}
	}
	rexsult, err = lang.RxFromString("85870777.2282833141709970713739108").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("powx3483")
	} else {
		if !(rexsult.ToString() == "1.35615463448707573424578785973269E-16") {
			lang.SayString("powx3483")
		}
	}
	rexsult, err = lang.RxFromString("20900.9693761555165742010339929779").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-39"))
	if err != nil {
		lang.SayString("powx3484")
	} else {
		if !(rexsult.ToString() == "3.26219014701526335296044439989665E-169") {
			lang.SayString("powx3484")
		}
	}
	rexsult, err = lang.RxFromString("98.4134807921002817357000140482039").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("powx3486")
	} else {
		if !(rexsult.ToString() == "953155.543384739667965055839894682") {
			lang.SayString("powx3486")
		}
	}
	rexsult, err = lang.RxFromString("545746433.649359734136476718176330E-787").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("powx3487")
	} else {
		if !(rexsult.ToString() == "2.06559640092667166976186801348662E+3891") {
			lang.SayString("powx3487")
		}
	}
	rexsult, err = lang.RxFromString("741304513547.273820525801608231737").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3488")
	} else {
		if !(rexsult.ToString() == "3.01985838652892073903194846668712E+47") {
			lang.SayString("powx3488")
		}
	}
	rexsult, err = lang.RxFromString("-706.145005094292315613907254240553").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("4740"))
	if err != nil {
		lang.SayString("powx3489")
	} else {
		if !(rexsult.ToString() == "5.71538696387742688721438158671168E+13503") {
			lang.SayString("powx3489")
		}
	}
	rexsult, err = lang.RxFromString("-769925786.823099083228795187975893").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-31202"))
	if err != nil {
		lang.SayString("powx3490")
	} else {
		if !(rexsult.ToString() == "1.05316836587185980546562763748913E-277275") {
			lang.SayString("powx3490")
		}
	}
	rexsult, err = lang.RxFromString("84438610546049.7256507419289692857E+906").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("powx3491")
	} else {
		if !(rexsult.ToString() == "4.29245144719689283247342866988213E+4599") {
			lang.SayString("powx3491")
		}
	}
	rexsult, err = lang.RxFromString("549760.911304725795164589619286514").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("165"))
	if err != nil {
		lang.SayString("powx3492")
	} else {
		if !(rexsult.ToString() == "1.34488925442386544028875603347654E+947") {
			lang.SayString("powx3492")
		}
	}
	rexsult, err = lang.RxFromString("3650514.18649737956855828939662794").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("8086721"))
	if err != nil {
		lang.SayString("powx3493")
	} else {
		if !(rexsult.ToString() == "1.69397027887946819100277154359204E+53067926") {
			lang.SayString("powx3493")
		}
	}
	rexsult, err = lang.RxFromString("55067723881950.1346958179604099594").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("powx3494")
	} else {
		if !(rexsult.ToString() == "2.14746386538529270173788457887121E-124") {
			lang.SayString("powx3494")
		}
	}
	rexsult, err = lang.RxFromString("868251123.413992653362860637541060E+019").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3495")
	} else {
		if !(rexsult.ToString() == "4.28422354304291884802690733853227E+167") {
			lang.SayString("powx3495")
		}
	}
	rexsult, err = lang.RxFromString("-646.464431574014407536004990059069").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("powx3496")
	} else {
		if !(rexsult.ToString() == "3.27825641569860861774700548035691E-23") {
			lang.SayString("powx3496")
		}
	}
	rexsult, err = lang.RxFromString("354.546679975219753598558273421556").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3497")
	} else {
		if !(rexsult.ToString() == "1.41999246365875617298270414304233E-18") {
			lang.SayString("powx3497")
		}
	}
	rexsult, err = lang.RxFromString("91936087917435.5974889495278215874").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("powx3498")
	} else {
		if !(rexsult.ToString() == "1.8013489993903570871965906508263E-98") {
			lang.SayString("powx3498")
		}
	}
	rexsult, err = lang.RxFromString("-07345.6422518528556136521417259811E-600").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("powx3499")
	} else {
		if !(rexsult.ToString() == "2.91151541552217582082937236255996E-2385") {
			lang.SayString("powx3499")
		}
	}
	rexsult, err = lang.RxFromString("-253280724.939458021588167965038184").OpPow(lang.RxSetWithDigit(33), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("powx3500")
	} else {
		if !(rexsult.ToString() == "2.64005420221406808782284459794424E+50") {
			lang.SayString("powx3500")
		}
	}
	rexsult, err = lang.RxFromString("905.67402").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow001")
	} else {
		if !(rexsult.ToString() == "0.0000012191473") {
			lang.SayString("xpow001")
		}
	}
	rexsult, err = lang.RxFromString("309759261").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("62663"))
	if err != nil {
		lang.SayString("xpow003")
	} else {
		if !(rexsult.ToString() == "1.13679199E+532073") {
			lang.SayString("xpow003")
		}
	}
	rexsult, err = lang.RxFromString("323902.714").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("xpow005")
	} else {
		if !(rexsult.ToString() == "8.65989204E-34") {
			lang.SayString("xpow005")
		}
	}
	rexsult, err = lang.RxFromString("5.11970092").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8807"))
	if err != nil {
		lang.SayString("xpow006")
	} else {
		if !(rexsult.ToString() == "4.81819262E-6247") {
			lang.SayString("xpow006")
		}
	}
	rexsult, err = lang.RxFromString("-7.99874516").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4562"))
	if err != nil {
		lang.SayString("xpow007")
	} else {
		if !(rexsult.ToString() == "3.85236199E+4119") {
			lang.SayString("xpow007")
		}
	}
	rexsult, err = lang.RxFromString("297802878").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-927206"))
	if err != nil {
		lang.SayString("xpow008")
	} else {
		if !(rexsult.ToString() == "1.9460281E-7857078") {
			lang.SayString("xpow008")
		}
	}
	rexsult, err = lang.RxFromString("-766.651824").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("31300"))
	if err != nil {
		lang.SayString("xpow009")
	} else {
		if !(rexsult.ToString() == "8.37189011E+90287") {
			lang.SayString("xpow009")
		}
	}
	rexsult, err = lang.RxFromString("456417160").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-41346"))
	if err != nil {
		lang.SayString("xpow011")
	} else {
		if !(rexsult.ToString() == "1.04766863E-358030") {
			lang.SayString("xpow011")
		}
	}
	rexsult, err = lang.RxFromString("102895.722").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow012")
	} else {
		if !(rexsult.ToString() == "9.17926786E-16") {
			lang.SayString("xpow012")
		}
	}
	rexsult, err = lang.RxFromString("61.3033331E+157644141").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("xpow013")
	} else {
		if !(rexsult.ToString() == "1.88406322E-945864857") {
			lang.SayString("xpow013")
		}
	}
	rexsult, err = lang.RxFromString("80223.3897").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow014")
	} else {
		if !(rexsult.ToString() == "2.13848919E+34") {
			lang.SayString("xpow014")
		}
	}
	rexsult, err = lang.RxFromString("-654645.954").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("xpow015")
	} else {
		if !(rexsult.ToString() == "-4.5283669E-53") {
			lang.SayString("xpow015")
		}
	}
	rexsult, err = lang.RxFromString("-39674.7190").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2490608"))
	if err != nil {
		lang.SayString("xpow017")
	} else {
		if !(rexsult.ToString() == "2.55032329E+11453095") {
			lang.SayString("xpow017")
		}
	}
	rexsult, err = lang.RxFromString("-64138.0578").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow019")
	} else {
		if !(rexsult.ToString() == "-2.63844116E+14") {
			lang.SayString("xpow019")
		}
	}
	rexsult, err = lang.RxFromString("61399.8527").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-64344485"))
	if err != nil {
		lang.SayString("xpow020")
	} else {
		if !(rexsult.ToString() == "1.27378842E-308092161") {
			lang.SayString("xpow020")
		}
	}
	rexsult, err = lang.RxFromString("-722960.204").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-26154600"))
	if err != nil {
		lang.SayString("xpow021")
	} else {
		if !(rexsult.ToString() == "5.34236139E-153242794") {
			lang.SayString("xpow021")
		}
	}
	rexsult, err = lang.RxFromString("43.7456245").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("547441956"))
	if err != nil {
		lang.SayString("xpow023")
	} else {
		if !(rexsult.ToString() == "2.91742391E+898316458") {
			lang.SayString("xpow023")
		}
	}
	rexsult, err = lang.RxFromString("29.498114").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-26486451"))
	if err != nil {
		lang.SayString("xpow026")
	} else {
		if !(rexsult.ToString() == "4.22252513E-38929634") {
			lang.SayString("xpow026")
		}
	}
	rexsult, err = lang.RxFromString("-349388.759").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-196216"))
	if err != nil {
		lang.SayString("xpow028")
	} else {
		if !(rexsult.ToString() == "1.24551752E-1087686") {
			lang.SayString("xpow028")
		}
	}
	rexsult, err = lang.RxFromString("-70905112.4").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-91353969"))
	if err != nil {
		lang.SayString("xpow029")
	} else {
		if !(rexsult.ToString() == "-3.05944741E-717190554") {
			lang.SayString("xpow029")
		}
	}
	rexsult, err = lang.RxFromString("-225094.28").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-89"))
	if err != nil {
		lang.SayString("xpow030")
	} else {
		if !(rexsult.ToString() == "-4.36076965E-477") {
			lang.SayString("xpow030")
		}
	}
	rexsult, err = lang.RxFromString("50.4442340").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow031")
	} else {
		if !(rexsult.ToString() == "4.19268518E+13") {
			lang.SayString("xpow031")
		}
	}
	rexsult, err = lang.RxFromString("-32311.9037").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow032")
	} else {
		if !(rexsult.ToString() == "1.1882296E+36") {
			lang.SayString("xpow032")
		}
	}
	rexsult, err = lang.RxFromString("213361789").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-600"))
	if err != nil {
		lang.SayString("xpow036")
	} else {
		if !(rexsult.ToString() == "3.38854684E-4998") {
			lang.SayString("xpow036")
		}
	}
	rexsult, err = lang.RxFromString("-795522555.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-298"))
	if err != nil {
		lang.SayString("xpow037")
	} else {
		if !(rexsult.ToString() == "4.03232712E-2653") {
			lang.SayString("xpow037")
		}
	}
	rexsult, err = lang.RxFromString("-501260651.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("xpow038")
	} else {
		if !(rexsult.ToString() == "-5.00526961E-79") {
			lang.SayString("xpow038")
		}
	}
	rexsult, err = lang.RxFromString("-4629035.31").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-168"))
	if err != nil {
		lang.SayString("xpow042")
	} else {
		if !(rexsult.ToString() == "1.57614831E-1120") {
			lang.SayString("xpow042")
		}
	}
	rexsult, err = lang.RxFromString("136.255393E+53329961").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow045")
	} else {
		if !(rexsult.ToString() == "2.12927373E-266649816") {
			lang.SayString("xpow045")
		}
	}
	rexsult, err = lang.RxFromString("-876673.100").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6151"))
	if err != nil {
		lang.SayString("xpow046")
	} else {
		if !(rexsult.ToString() == "-4.03111774E-36555") {
			lang.SayString("xpow046")
		}
	}
	rexsult, err = lang.RxFromString("-9.15117551").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow048")
	} else {
		if !(rexsult.ToString() == "-0.0000155817265") {
			lang.SayString("xpow048")
		}
	}
	rexsult, err = lang.RxFromString("106211716.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3456794"))
	if err != nil {
		lang.SayString("xpow052")
	} else {
		if !(rexsult.ToString() == "2.07225581E-27744825") {
			lang.SayString("xpow052")
		}
	}
	rexsult, err = lang.RxFromString("1.25018078").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("xpow053")
	} else {
		if !(rexsult.ToString() == "2.4428189") {
			lang.SayString("xpow053")
		}
	}
	rexsult, err = lang.RxFromString("364.99811").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-46222"))
	if err != nil {
		lang.SayString("xpow054")
	} else {
		if !(rexsult.ToString() == "6.35570856E-118435") {
			lang.SayString("xpow054")
		}
	}
	rexsult, err = lang.RxFromString("169601285").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("714527"))
	if err != nil {
		lang.SayString("xpow056")
	} else {
		if !(rexsult.ToString() == "2.06052444E+5880149") {
			lang.SayString("xpow056")
		}
	}
	rexsult, err = lang.RxFromString("-746.293386").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("927750"))
	if err != nil {
		lang.SayString("xpow059")
	} else {
		if !(rexsult.ToString() == "7.49278741E+2665341") {
			lang.SayString("xpow059")
		}
	}
	rexsult, err = lang.RxFromString("6.64377249").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow061")
	} else {
		if !(rexsult.ToString() == "3795928.44") {
			lang.SayString("xpow061")
		}
	}
	rexsult, err = lang.RxFromString("6.44693097").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-87196"))
	if err != nil {
		lang.SayString("xpow063")
	} else {
		if !(rexsult.ToString() == "4.5088173E-70573") {
			lang.SayString("xpow063")
		}
	}
	rexsult, err = lang.RxFromString("-7701.42814").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("72668"))
	if err != nil {
		lang.SayString("xpow065")
	} else {
		if !(rexsult.ToString() == "2.29543837E+282429") {
			lang.SayString("xpow065")
		}
	}
	rexsult, err = lang.RxFromString("-851.754789").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-582659"))
	if err != nil {
		lang.SayString("xpow066")
	} else {
		if !(rexsult.ToString() == "-6.83532593E-1707375") {
			lang.SayString("xpow066")
		}
	}
	rexsult, err = lang.RxFromString("-5.01992943").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7852"))
	if err != nil {
		lang.SayString("xpow067")
	} else {
		if !(rexsult.ToString() == "7.54481448E+5501") {
			lang.SayString("xpow067")
		}
	}
	rexsult, err = lang.RxFromString("-12393257.2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow068")
	} else {
		if !(rexsult.ToString() == "5.5652375E+56") {
			lang.SayString("xpow068")
		}
	}
	rexsult, err = lang.RxFromString("8.27822605").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9241557"))
	if err != nil {
		lang.SayString("xpow072")
	} else {
		if !(rexsult.ToString() == "5.10219969E+8483169") {
			lang.SayString("xpow072")
		}
	}
	rexsult, err = lang.RxFromString("-1.43581098").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7286314"))
	if err != nil {
		lang.SayString("xpow073")
	} else {
		if !(rexsult.ToString() == "1.09389741E+1144660") {
			lang.SayString("xpow073")
		}
	}
	rexsult, err = lang.RxFromString("-699036193.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow074")
	} else {
		if !(rexsult.ToString() == "5.70160724E+70") {
			lang.SayString("xpow074")
		}
	}
	rexsult, err = lang.RxFromString("-83.7273615E-305281957").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow075")
	} else {
		if !(rexsult.ToString() == "-1.70371828E+915845865") {
			lang.SayString("xpow075")
		}
	}
	rexsult, err = lang.RxFromString("8.48503224").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6522"))
	if err != nil {
		lang.SayString("xpow076")
	} else {
		if !(rexsult.ToString() == "4.76547542E+6056") {
			lang.SayString("xpow076")
		}
	}
	rexsult, err = lang.RxFromString("527916091").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-809"))
	if err != nil {
		lang.SayString("xpow077")
	} else {
		if !(rexsult.ToString() == "2.78609697E-7057") {
			lang.SayString("xpow077")
		}
	}
	rexsult, err = lang.RxFromString("3857058.60").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("xpow078")
	} else {
		if !(rexsult.ToString() == "3.29258824E+39") {
			lang.SayString("xpow078")
		}
	}
	rexsult, err = lang.RxFromString("-580.502955").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("38125522"))
	if err != nil {
		lang.SayString("xpow080")
	} else {
		if !(rexsult.ToString() == "6.07262078E+105371486") {
			lang.SayString("xpow080")
		}
	}
	rexsult, err = lang.RxFromString("-9627363.00").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow081")
	} else {
		if !(rexsult.ToString() == "1.35500601E-56") {
			lang.SayString("xpow081")
		}
	}
	rexsult, err = lang.RxFromString("-8378.55499").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("760"))
	if err != nil {
		lang.SayString("xpow083")
	} else {
		if !(rexsult.ToString() == "4.06007928E+2981") {
			lang.SayString("xpow083")
		}
	}
	rexsult, err = lang.RxFromString("-359866845.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow087")
	} else {
		if !(rexsult.ToString() == "-1.65687909E-43") {
			lang.SayString("xpow087")
		}
	}
	rexsult, err = lang.RxFromString("779934536.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-76562646"))
	if err != nil {
		lang.SayString("xpow088")
	} else {
		if !(rexsult.ToString() == "3.36739063E-680799501") {
			lang.SayString("xpow088")
		}
	}
	rexsult, err = lang.RxFromString("-4820.95451").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("xpow089")
	} else {
		if !(rexsult.ToString() == "5.40172082E+14") {
			lang.SayString("xpow089")
		}
	}
	rexsult, err = lang.RxFromString("69355976.9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("xpow090")
	} else {
		if !(rexsult.ToString() == "3.88294346E-79") {
			lang.SayString("xpow090")
		}
	}
	rexsult, err = lang.RxFromString("-12672093.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("xpow091")
	} else {
		if !(rexsult.ToString() == "-8.42626658E+63") {
			lang.SayString("xpow091")
		}
	}
	rexsult, err = lang.RxFromString("-5910750.2").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow092")
	} else {
		if !(rexsult.ToString() == "-2.52056696E+47") {
			lang.SayString("xpow092")
		}
	}
	rexsult, err = lang.RxFromString("-532577268.E-163806629").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow093")
	} else {
		if !(rexsult.ToString() == "3.52561389E+327613240") {
			lang.SayString("xpow093")
		}
	}
	rexsult, err = lang.RxFromString("-789904.686E-217225000").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow097")
	} else {
		if !(rexsult.ToString() == "1.60269403E+434449988") {
			lang.SayString("xpow097")
		}
	}
	rexsult, err = lang.RxFromString("129.878613").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow099")
	} else {
		if !(rexsult.ToString() == "2.70590029E-11") {
			lang.SayString("xpow099")
		}
	}
	rexsult, err = lang.RxFromString("-78810.6297").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-399885"))
	if err != nil {
		lang.SayString("xpow101")
	} else {
		if !(rexsult.ToString() == "-1.54252408E-1958071") {
			lang.SayString("xpow101")
		}
	}
	rexsult, err = lang.RxFromString("409189761").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-771"))
	if err != nil {
		lang.SayString("xpow102")
	} else {
		if !(rexsult.ToString() == "1.60698414E-6640") {
			lang.SayString("xpow102")
		}
	}
	rexsult, err = lang.RxFromString("-1.68748838").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("460"))
	if err != nil {
		lang.SayString("xpow103")
	} else {
		if !(rexsult.ToString() == "3.39440648E+104") {
			lang.SayString("xpow103")
		}
	}
	rexsult, err = lang.RxFromString("553527296.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-7924"))
	if err != nil {
		lang.SayString("xpow104")
	} else {
		if !(rexsult.ToString() == "2.32397214E-69281") {
			lang.SayString("xpow104")
		}
	}
	rexsult, err = lang.RxFromString("-38.7465207").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("64936"))
	if err != nil {
		lang.SayString("xpow105")
	} else {
		if !(rexsult.ToString() == "3.01500762E+103133") {
			lang.SayString("xpow105")
		}
	}
	rexsult, err = lang.RxFromString("-201075.248").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("846"))
	if err != nil {
		lang.SayString("xpow106")
	} else {
		if !(rexsult.ToString() == "4.37911767E+4486") {
			lang.SayString("xpow106")
		}
	}
	rexsult, err = lang.RxFromString("91048.4559").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("75953609"))
	if err != nil {
		lang.SayString("xpow107")
	} else {
		if !(rexsult.ToString() == "6.94467746E+376674650") {
			lang.SayString("xpow107")
		}
	}
	rexsult, err = lang.RxFromString("88.4370343").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("xpow109")
	} else {
		if !(rexsult.ToString() == "3.41710479E-20") {
			lang.SayString("xpow109")
		}
	}
	rexsult, err = lang.RxFromString("-17643.39").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow110")
	} else {
		if !(rexsult.ToString() == "311289211") {
			lang.SayString("xpow110")
		}
	}
	rexsult, err = lang.RxFromString("4589785.16").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7459"))
	if err != nil {
		lang.SayString("xpow111")
	} else {
		if !(rexsult.ToString() == "2.03795258E+49690") {
			lang.SayString("xpow111")
		}
	}
	rexsult, err = lang.RxFromString("982.217817").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow113")
	} else {
		if !(rexsult.ToString() == "8.81971709E+20") {
			lang.SayString("xpow113")
		}
	}
	rexsult, err = lang.RxFromString("503712056.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("xpow114")
	} else {
		if !(rexsult.ToString() == "6.12217764E-53") {
			lang.SayString("xpow114")
		}
	}
	rexsult, err = lang.RxFromString("883.849223").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("249259171"))
	if err != nil {
		lang.SayString("xpow115")
	} else {
		if !(rexsult.ToString() == "5.15642844E+734411783") {
			lang.SayString("xpow115")
		}
	}
	rexsult, err = lang.RxFromString("-6291780.97").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow118")
	} else {
		if !(rexsult.ToString() == "-2.49069636E+20") {
			lang.SayString("xpow118")
		}
	}
	rexsult, err = lang.RxFromString("14239029.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-36527"))
	if err != nil {
		lang.SayString("xpow120")
	} else {
		if !(rexsult.ToString() == "6.64292731E-261296") {
			lang.SayString("xpow120")
		}
	}
	rexsult, err = lang.RxFromString("-37721.1567E-115787341").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow122")
	} else {
		if !(rexsult.ToString() == "2.43960765E+926298691") {
			lang.SayString("xpow122")
		}
	}
	rexsult, err = lang.RxFromString("-2078852.83E-647080089").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("xpow123")
	} else {
		if !(rexsult.ToString() == "-4.81034533E+647080082") {
			lang.SayString("xpow123")
		}
	}
	rexsult, err = lang.RxFromString("-79145.3625").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-7719"))
	if err != nil {
		lang.SayString("xpow124")
	} else {
		if !(rexsult.ToString() == "-1.13181941E-37811") {
			lang.SayString("xpow124")
		}
	}
	rexsult, err = lang.RxFromString("911249557").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("79810805"))
	if err != nil {
		lang.SayString("xpow126")
	} else {
		if !(rexsult.ToString() == "6.77595741E+715075867") {
			lang.SayString("xpow126")
		}
	}
	rexsult, err = lang.RxFromString("341134.994").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow127")
	} else {
		if !(rexsult.ToString() == "3.96989314E+16") {
			lang.SayString("xpow127")
		}
	}
	rexsult, err = lang.RxFromString("244.23634").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("xpow128")
	} else {
		if !(rexsult.ToString() == "8.69063312E+11") {
			lang.SayString("xpow128")
		}
	}
	rexsult, err = lang.RxFromString("699631.893").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-226"))
	if err != nil {
		lang.SayString("xpow130")
	} else {
		if !(rexsult.ToString() == "1.14675511E-1321") {
			lang.SayString("xpow130")
		}
	}
	rexsult, err = lang.RxFromString("5.11629020").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-481"))
	if err != nil {
		lang.SayString("xpow132")
	} else {
		if !(rexsult.ToString() == "9.83021951E-342") {
			lang.SayString("xpow132")
		}
	}
	rexsult, err = lang.RxFromString("7.0598608").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-95908"))
	if err != nil {
		lang.SayString("xpow134")
	} else {
		if !(rexsult.ToString() == "4.57073877E-81407") {
			lang.SayString("xpow134")
		}
	}
	rexsult, err = lang.RxFromString("-7.91189845E+207202706").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow135")
	} else {
		if !(rexsult.ToString() == "6.25981371E+414405413") {
			lang.SayString("xpow135")
		}
	}
	rexsult, err = lang.RxFromString("44911.089").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("xpow138")
	} else {
		if !(rexsult.ToString() == "2.99546425E-47") {
			lang.SayString("xpow138")
		}
	}
	rexsult, err = lang.RxFromString("452371821.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4109709"))
	if err != nil {
		lang.SayString("xpow139")
	} else {
		if !(rexsult.ToString() == "1.15528807E-35571568") {
			lang.SayString("xpow139")
		}
	}
	rexsult, err = lang.RxFromString("94007.4392").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("xpow140")
	} else {
		if !(rexsult.ToString() == "1.74397397E-45") {
			lang.SayString("xpow140")
		}
	}
	rexsult, err = lang.RxFromString("-7919.30254").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-670"))
	if err != nil {
		lang.SayString("xpow142")
	} else {
		if !(rexsult.ToString() == "7.58147724E-2613") {
			lang.SayString("xpow142")
		}
	}
	rexsult, err = lang.RxFromString("461.58280E+136110821").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow143")
	} else {
		if !(rexsult.ToString() == "4.46423781E+952775765") {
			lang.SayString("xpow143")
		}
	}
	rexsult, err = lang.RxFromString("76482.352").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8237807"))
	if err != nil {
		lang.SayString("xpow146")
	} else {
		if !(rexsult.ToString() == "8.44216559E+40229834") {
			lang.SayString("xpow146")
		}
	}
	rexsult, err = lang.RxFromString("-592464.92").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("71445511"))
	if err != nil {
		lang.SayString("xpow149")
	} else {
		if !(rexsult.ToString() == "-1.58269108E+412430832") {
			lang.SayString("xpow149")
		}
	}
	rexsult, err = lang.RxFromString("-73774.4165").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-40"))
	if err != nil {
		lang.SayString("xpow150")
	} else {
		if !(rexsult.ToString() == "1.92206765E-195") {
			lang.SayString("xpow150")
		}
	}
	rexsult, err = lang.RxFromString("-524724715.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-55764"))
	if err != nil {
		lang.SayString("xpow151")
	} else {
		if !(rexsult.ToString() == "5.47898351E-486259") {
			lang.SayString("xpow151")
		}
	}
	rexsult, err = lang.RxFromString("7.53800427").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow152")
	} else {
		if !(rexsult.ToString() == "10424399.2") {
			lang.SayString("xpow152")
		}
	}
	rexsult, err = lang.RxFromString("37.6027452").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow153")
	} else {
		if !(rexsult.ToString() == "1.06300881E+11") {
			lang.SayString("xpow153")
		}
	}
	rexsult, err = lang.RxFromString("2447660.39").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-36981"))
	if err != nil {
		lang.SayString("xpow154")
	} else {
		if !(rexsult.ToString() == "3.92066064E-236263") {
			lang.SayString("xpow154")
		}
	}
	rexsult, err = lang.RxFromString("2160.36419").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("xpow155")
	} else {
		if !(rexsult.ToString() == "2160.36419") {
			lang.SayString("xpow155")
		}
	}
	rexsult, err = lang.RxFromString("8926.44939").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("55"))
	if err != nil {
		lang.SayString("xpow156")
	} else {
		if !(rexsult.ToString() == "1.93789877E+217") {
			lang.SayString("xpow156")
		}
	}
	rexsult, err = lang.RxFromString("861588029").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("xpow157")
	} else {
		if !(rexsult.ToString() == "1.81468553E-36") {
			lang.SayString("xpow157")
		}
	}
	rexsult, err = lang.RxFromString("-34.5253062").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("53"))
	if err != nil {
		lang.SayString("xpow158")
	} else {
		if !(rexsult.ToString() == "-3.32115821E+81") {
			lang.SayString("xpow158")
		}
	}
	rexsult, err = lang.RxFromString("-18861647.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("99794587"))
	if err != nil {
		lang.SayString("xpow159")
	} else {
		if !(rexsult.ToString() == "-4.2895746E+726063462") {
			lang.SayString("xpow159")
		}
	}
	rexsult, err = lang.RxFromString("322192.407").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("462"))
	if err != nil {
		lang.SayString("xpow160")
	} else {
		if !(rexsult.ToString() == "5.61395873E+2544") {
			lang.SayString("xpow160")
		}
	}
	rexsult, err = lang.RxFromString("293.773732").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("xpow162")
	} else {
		if !(rexsult.ToString() == "2.18808809E+12") {
			lang.SayString("xpow162")
		}
	}
	rexsult, err = lang.RxFromString("-103519362").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("51897955"))
	if err != nil {
		lang.SayString("xpow163")
	} else {
		if !(rexsult.ToString() == "-4.28858229E+415963229") {
			lang.SayString("xpow163")
		}
	}
	rexsult, err = lang.RxFromString("-5409.00482").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow167")
	} else {
		if !(rexsult.ToString() == "3.41794652E-8") {
			lang.SayString("xpow167")
		}
	}
	rexsult, err = lang.RxFromString("-957960.367").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("323"))
	if err != nil {
		lang.SayString("xpow168")
	} else {
		if !(rexsult.ToString() == "-9.4461746E+1931") {
			lang.SayString("xpow168")
		}
	}
	rexsult, err = lang.RxFromString("840258203").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("xpow170")
	} else {
		if !(rexsult.ToString() == "3.51946431E+53") {
			lang.SayString("xpow170")
		}
	}
	rexsult, err = lang.RxFromString("-205842096.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-191343"))
	if err != nil {
		lang.SayString("xpow171")
	} else {
		if !(rexsult.ToString() == "-2.66955553E-1590737") {
			lang.SayString("xpow171")
		}
	}
	rexsult, err = lang.RxFromString("42501124.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("xpow172")
	} else {
		if !(rexsult.ToString() == "4.52484536E+68") {
			lang.SayString("xpow172")
		}
	}
	rexsult, err = lang.RxFromString("-8022370.31").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9858582"))
	if err != nil {
		lang.SayString("xpow174")
	} else {
		if !(rexsult.ToString() == "2.34458249E+68066634") {
			lang.SayString("xpow174")
		}
	}
	rexsult, err = lang.RxFromString("-697273715E-242824870").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("xpow176")
	} else {
		if !(rexsult.ToString() == "4.23045251E+971299444") {
			lang.SayString("xpow176")
		}
	}
	rexsult, err = lang.RxFromString("738063892").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("89900468"))
	if err != nil {
		lang.SayString("xpow178")
	} else {
		if !(rexsult.ToString() == "1.53166723E+797245797") {
			lang.SayString("xpow178")
		}
	}
	rexsult, err = lang.RxFromString("-630309366").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("xpow179")
	} else {
		if !(rexsult.ToString() == "-6.3681921E-80") {
			lang.SayString("xpow179")
		}
	}
	rexsult, err = lang.RxFromString("613.207774").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3008"))
	if err != nil {
		lang.SayString("xpow180")
	} else {
		if !(rexsult.ToString() == "7.5193916E-8386") {
			lang.SayString("xpow180")
		}
	}
	rexsult, err = lang.RxFromString("-93006222.3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow181")
	} else {
		if !(rexsult.ToString() == "-1.24297956E-24") {
			lang.SayString("xpow181")
		}
	}
	rexsult, err = lang.RxFromString("-18116.0621").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow182")
	} else {
		if !(rexsult.ToString() == "-5.94554133E+12") {
			lang.SayString("xpow182")
		}
	}
	rexsult, err = lang.RxFromString("4180.30821").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow184")
	} else {
		if !(rexsult.ToString() == "5.72246828E-8") {
			lang.SayString("xpow184")
		}
	}
	rexsult, err = lang.RxFromString("571.536725").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("390"))
	if err != nil {
		lang.SayString("xpow185")
	} else {
		if !(rexsult.ToString() == "1.76691373E+1075") {
			lang.SayString("xpow185")
		}
	}
	rexsult, err = lang.RxFromString("44651895.7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-910508"))
	if err != nil {
		lang.SayString("xpow189")
	} else {
		if !(rexsult.ToString() == "3.72264277E-6965241") {
			lang.SayString("xpow189")
		}
	}
	rexsult, err = lang.RxFromString("25.2592149").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("59"))
	if err != nil {
		lang.SayString("xpow191")
	} else {
		if !(rexsult.ToString() == "5.53058435E+82") {
			lang.SayString("xpow191")
		}
	}
	rexsult, err = lang.RxFromString("-6.850835").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1273"))
	if err != nil {
		lang.SayString("xpow192")
	} else {
		if !(rexsult.ToString() == "-1.25462678E-1064") {
			lang.SayString("xpow192")
		}
	}
	rexsult, err = lang.RxFromString("174.272325").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5638"))
	if err != nil {
		lang.SayString("xpow193")
	} else {
		if !(rexsult.ToString() == "1.11137724E+12636") {
			lang.SayString("xpow193")
		}
	}
	rexsult, err = lang.RxFromString("3455629.76").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow194")
	} else {
		if !(rexsult.ToString() == "4.91793015E-53") {
			lang.SayString("xpow194")
		}
	}
	rexsult, err = lang.RxFromString("9.10025079").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow197")
	} else {
		if !(rexsult.ToString() == "5168607.19") {
			lang.SayString("xpow197")
		}
	}
	rexsult, err = lang.RxFromString("-8.29530327").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow199")
	} else {
		if !(rexsult.ToString() == "-570.816876") {
			lang.SayString("xpow199")
		}
	}
	rexsult, err = lang.RxFromString("-57101683.5").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow200")
	} else {
		if !(rexsult.ToString() == "1.13029368E+62") {
			lang.SayString("xpow200")
		}
	}
	rexsult, err = lang.RxFromString("-603326.740").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1711"))
	if err != nil {
		lang.SayString("xpow201")
	} else {
		if !(rexsult.ToString() == "-3.35315976E+9890") {
			lang.SayString("xpow201")
		}
	}
	rexsult, err = lang.RxFromString("-204.586767").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-236"))
	if err != nil {
		lang.SayString("xpow203")
	} else {
		if !(rexsult.ToString() == "4.29438222E-546") {
			lang.SayString("xpow203")
		}
	}
	rexsult, err = lang.RxFromString("-70.3805581").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("830138"))
	if err != nil {
		lang.SayString("xpow204")
	} else {
		if !(rexsult.ToString() == "4.95165841E+1533640") {
			lang.SayString("xpow204")
		}
	}
	rexsult, err = lang.RxFromString("-8818.47606").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-60766"))
	if err != nil {
		lang.SayString("xpow205")
	} else {
		if !(rexsult.ToString() == "1.64487755E-239746") {
			lang.SayString("xpow205")
		}
	}
	rexsult, err = lang.RxFromString("-656285310.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-107221462"))
	if err != nil {
		lang.SayString("xpow207")
	} else {
		if !(rexsult.ToString() == "8.0533808E-945381569") {
			lang.SayString("xpow207")
		}
	}
	rexsult, err = lang.RxFromString("653397.125").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7195"))
	if err != nil {
		lang.SayString("xpow208")
	} else {
		if !(rexsult.ToString() == "1.58522983E+41840") {
			lang.SayString("xpow208")
		}
	}
	rexsult, err = lang.RxFromString("-62011.4563E-117563240").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("xpow211")
	} else {
		if !(rexsult.ToString() == "1.75860546E+705379411") {
			lang.SayString("xpow211")
		}
	}
	rexsult, err = lang.RxFromString("315.33351").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("xpow212")
	} else {
		if !(rexsult.ToString() == "3.08269902E+22") {
			lang.SayString("xpow212")
		}
	}
	rexsult, err = lang.RxFromString("739.944710").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("202949"))
	if err != nil {
		lang.SayString("xpow213")
	} else {
		if !(rexsult.ToString() == "1.32611729E+582301") {
			lang.SayString("xpow213")
		}
	}
	rexsult, err = lang.RxFromString("87686.8016").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4204890"))
	if err != nil {
		lang.SayString("xpow214")
	} else {
		if !(rexsult.ToString() == "5.14846981E+20784494") {
			lang.SayString("xpow214")
		}
	}
	rexsult, err = lang.RxFromString("7428219.97").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("667"))
	if err != nil {
		lang.SayString("xpow217")
	} else {
		if !(rexsult.ToString() == "7.5880851E+4582") {
			lang.SayString("xpow217")
		}
	}
	rexsult, err = lang.RxFromString("-7291.19212").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow218")
	} else {
		if !(rexsult.ToString() == "53161482.5") {
			lang.SayString("xpow218")
		}
	}
	rexsult, err = lang.RxFromString("-358.24550").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("xpow219")
	} else {
		if !(rexsult.ToString() == "6.07123474E-11") {
			lang.SayString("xpow219")
		}
	}
	rexsult, err = lang.RxFromString("118.621826").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow220")
	} else {
		if !(rexsult.ToString() == "5.99109471E-7") {
			lang.SayString("xpow220")
		}
	}
	rexsult, err = lang.RxFromString("8071961.94").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("xpow221")
	} else {
		if !(rexsult.ToString() == "1.23885619E-7") {
			lang.SayString("xpow221")
		}
	}
	rexsult, err = lang.RxFromString("-35544.4029").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-567830"))
	if err != nil {
		lang.SayString("xpow223")
	} else {
		if !(rexsult.ToString() == "3.77069368E-2584065") {
			lang.SayString("xpow223")
		}
	}
	rexsult, err = lang.RxFromString("-509.483395").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-147242915"))
	if err != nil {
		lang.SayString("xpow225")
	} else {
		if !(rexsult.ToString() == "-3.10760519E-398605718") {
			lang.SayString("xpow225")
		}
	}
	rexsult, err = lang.RxFromString("895612630.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-36"))
	if err != nil {
		lang.SayString("xpow227")
	} else {
		if !(rexsult.ToString() == "5.2926413E-323") {
			lang.SayString("xpow227")
		}
	}
	rexsult, err = lang.RxFromString("25455.4973").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow228")
	} else {
		if !(rexsult.ToString() == "1.64947128E+13") {
			lang.SayString("xpow228")
		}
	}
	rexsult, err = lang.RxFromString("62871.2202").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow230")
	} else {
		if !(rexsult.ToString() == "3.95279033E+9") {
			lang.SayString("xpow230")
		}
	}
	rexsult, err = lang.RxFromString("71.9281575").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-9810013"))
	if err != nil {
		lang.SayString("xpow231")
	} else {
		if !(rexsult.ToString() == "2.00363798E-18216203") {
			lang.SayString("xpow231")
		}
	}
	rexsult, err = lang.RxFromString("-6388022.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-88"))
	if err != nil {
		lang.SayString("xpow232")
	} else {
		if !(rexsult.ToString() == "1.34201238E-599") {
			lang.SayString("xpow232")
		}
	}
	rexsult, err = lang.RxFromString("372567445.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("96"))
	if err != nil {
		lang.SayString("xpow233")
	} else {
		if !(rexsult.ToString() == "6.84968715E+822") {
			lang.SayString("xpow233")
		}
	}
	rexsult, err = lang.RxFromString("802.156517").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow234")
	} else {
		if !(rexsult.ToString() == "0.00000155411005") {
			lang.SayString("xpow234")
		}
	}
	rexsult, err = lang.RxFromString("-3.65207541").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("74501982"))
	if err != nil {
		lang.SayString("xpow235")
	} else {
		if !(rexsult.ToString() == "2.10339452E+41910325") {
			lang.SayString("xpow235")
		}
	}
	rexsult, err = lang.RxFromString("-5297.76981").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-860"))
	if err != nil {
		lang.SayString("xpow236")
	} else {
		if !(rexsult.ToString() == "1.90523108E-3203") {
			lang.SayString("xpow236")
		}
	}
	rexsult, err = lang.RxFromString("-684172.592").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow237")
	} else {
		if !(rexsult.ToString() == "4.80093005E+46") {
			lang.SayString("xpow237")
		}
	}
	rexsult, err = lang.RxFromString("626919.219").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("xpow238")
	} else {
		if !(rexsult.ToString() == "6.07112959E+34") {
			lang.SayString("xpow238")
		}
	}
	rexsult, err = lang.RxFromString("-77480.5840").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("xpow239")
	} else {
		if !(rexsult.ToString() == "-1.00631969E+44") {
			lang.SayString("xpow239")
		}
	}
	rexsult, err = lang.RxFromString("-7177620.29").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7786344"))
	if err != nil {
		lang.SayString("xpow240")
	} else {
		if !(rexsult.ToString() == "2.96037074E+53383022") {
			lang.SayString("xpow240")
		}
	}
	rexsult, err = lang.RxFromString("9.6224130").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("xpow241")
	} else {
		if !(rexsult.ToString() == "82493.5448") {
			lang.SayString("xpow241")
		}
	}
	rexsult, err = lang.RxFromString("65587553.7").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("600575"))
	if err != nil {
		lang.SayString("xpow243")
	} else {
		if !(rexsult.ToString() == "3.40404817E+4694587") {
			lang.SayString("xpow243")
		}
	}
	rexsult, err = lang.RxFromString("69573.988").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("xpow245")
	} else {
		if !(rexsult.ToString() == "3.76297229E-49") {
			lang.SayString("xpow245")
		}
	}
	rexsult, err = lang.RxFromString("2362.06251").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("xpow246")
	} else {
		if !(rexsult.ToString() == "3.21243577E-14") {
			lang.SayString("xpow246")
		}
	}
	rexsult, err = lang.RxFromString("-6364720.49").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5539246"))
	if err != nil {
		lang.SayString("xpow249")
	} else {
		if !(rexsult.ToString() == "2.96894641E+37687807") {
			lang.SayString("xpow249")
		}
	}
	rexsult, err = lang.RxFromString("-814599.475").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-15"))
	if err != nil {
		lang.SayString("xpow250")
	} else {
		if !(rexsult.ToString() == "-2.16689622E-89") {
			lang.SayString("xpow250")
		}
	}
	rexsult, err = lang.RxFromString("-877498.755").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("xpow251")
	} else {
		if !(rexsult.ToString() == "-5.20274505E+29") {
			lang.SayString("xpow251")
		}
	}
	rexsult, err = lang.RxFromString("700354586.E-99856707").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow254")
	} else {
		if !(rexsult.ToString() == "8.2646761E-698996888") {
			lang.SayString("xpow254")
		}
	}
	rexsult, err = lang.RxFromString("5350882.59").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-36329829"))
	if err != nil {
		lang.SayString("xpow256")
	} else {
		if !(rexsult.ToString() == "9.77006107E-244442546") {
			lang.SayString("xpow256")
		}
	}
	rexsult, err = lang.RxFromString("91966.4084E+210382952").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow257")
	} else {
		if !(rexsult.ToString() == "8.45782027E+420765913") {
			lang.SayString("xpow257")
		}
	}
	rexsult, err = lang.RxFromString("-9611312.33").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("22109736"))
	if err != nil {
		lang.SayString("xpow259")
	} else {
		if !(rexsult.ToString() == "6.74530828E+154387481") {
			lang.SayString("xpow259")
		}
	}
	rexsult, err = lang.RxFromString("693881413.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow261")
	} else {
		if !(rexsult.ToString() == "3.34084066E+26") {
			lang.SayString("xpow261")
		}
	}
	rexsult, err = lang.RxFromString("1123.32456").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow263")
	} else {
		if !(rexsult.ToString() == "2.53537401E+24") {
			lang.SayString("xpow263")
		}
	}
	rexsult, err = lang.RxFromString("34.1935525").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-40768"))
	if err != nil {
		lang.SayString("xpow266")
	} else {
		if !(rexsult.ToString() == "1.4517421E-62536") {
			lang.SayString("xpow266")
		}
	}
	rexsult, err = lang.RxFromString("51.025101").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4467692"))
	if err != nil {
		lang.SayString("xpow269")
	} else {
		if !(rexsult.ToString() == "4.49462589E-7629853") {
			lang.SayString("xpow269")
		}
	}
	rexsult, err = lang.RxFromString("-2214.76582").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("xpow270")
	} else {
		if !(rexsult.ToString() == "2.40608658E+13") {
			lang.SayString("xpow270")
		}
	}
	rexsult, err = lang.RxFromString("513115529.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow272")
	} else {
		if !(rexsult.ToString() == "1.35096929E+26") {
			lang.SayString("xpow272")
		}
	}
	rexsult, err = lang.RxFromString("-247157.208").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-532990"))
	if err != nil {
		lang.SayString("xpow273")
	} else {
		if !(rexsult.ToString() == "1.48314033E-2874401") {
			lang.SayString("xpow273")
		}
	}
	rexsult, err = lang.RxFromString("-1156008.8").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8870382"))
	if err != nil {
		lang.SayString("xpow275")
	} else {
		if !(rexsult.ToString() == "4.32494996E-53780782") {
			lang.SayString("xpow275")
		}
	}
	rexsult, err = lang.RxFromString("880097928.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow276")
	} else {
		if !(rexsult.ToString() == "1.89384751E-45") {
			lang.SayString("xpow276")
		}
	}
	rexsult, err = lang.RxFromString("5796.2524").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow277")
	} else {
		if !(rexsult.ToString() == "1.94734037E+11") {
			lang.SayString("xpow277")
		}
	}
	rexsult, err = lang.RxFromString("42643477.8").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow279")
	} else {
		if !(rexsult.ToString() == "7.7545723E+22") {
			lang.SayString("xpow279")
		}
	}
	rexsult, err = lang.RxFromString("-31918.9176E-163031657").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow280")
	} else {
		if !(rexsult.ToString() == "9.8153025E+326063304") {
			lang.SayString("xpow280")
		}
	}
	rexsult, err = lang.RxFromString("84224841.0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow281")
	} else {
		if !(rexsult.ToString() == "5.97476185E+23") {
			lang.SayString("xpow281")
		}
	}
	rexsult, err = lang.RxFromString("-64413698.9").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("xpow282")
	} else {
		if !(rexsult.ToString() == "-2.17346338E-55") {
			lang.SayString("xpow282")
		}
	}
	rexsult, err = lang.RxFromString("-62.5059208").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("xpow283")
	} else {
		if !(rexsult.ToString() == "9.10356659E+17") {
			lang.SayString("xpow283")
		}
	}
	rexsult, err = lang.RxFromString("9090950.80").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("436"))
	if err != nil {
		lang.SayString("xpow284")
	} else {
		if !(rexsult.ToString() == "8.98789557E+3033") {
			lang.SayString("xpow284")
		}
	}
	rexsult, err = lang.RxFromString("-584537670.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow287")
	} else {
		if !(rexsult.ToString() == "-1.99727337E+26") {
			lang.SayString("xpow287")
		}
	}
	rexsult, err = lang.RxFromString("5.15309635").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("xpow289")
	} else {
		if !(rexsult.ToString() == "0.0000103638749") {
			lang.SayString("xpow289")
		}
	}
	rexsult, err = lang.RxFromString("-940030153.E+83797657").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("xpow290")
	} else {
		if !(rexsult.ToString() == "1.2806571E-335190664") {
			lang.SayString("xpow290")
		}
	}
	rexsult, err = lang.RxFromString("89088.9683E+587739290").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("xpow291")
	} else {
		if !(rexsult.ToString() == "8.90889683E+587739294") {
			lang.SayString("xpow291")
		}
	}
	rexsult, err = lang.RxFromString("3336750").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("xpow292")
	} else {
		if !(rexsult.ToString() == "1.38019997E+39") {
			lang.SayString("xpow292")
		}
	}
	rexsult, err = lang.RxFromString("904654622.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow293")
	} else {
		if !(rexsult.ToString() == "4.95883485E+62") {
			lang.SayString("xpow293")
		}
	}
	rexsult, err = lang.RxFromString("304804380").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4681"))
	if err != nil {
		lang.SayString("xpow294")
	} else {
		if !(rexsult.ToString() == "1.98037102E-39714") {
			lang.SayString("xpow294")
		}
	}
	rexsult, err = lang.RxFromString("674.55569").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow295")
	} else {
		if !(rexsult.ToString() == "2.33269265E-23") {
			lang.SayString("xpow295")
		}
	}
	rexsult, err = lang.RxFromString("-5111.51025E-108006096").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("xpow296")
	} else {
		if !(rexsult.ToString() == "-3.48936323E-540030462") {
			lang.SayString("xpow296")
		}
	}
	rexsult, err = lang.RxFromString("299350.435").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3373"))
	if err != nil {
		lang.SayString("xpow298")
	} else {
		if !(rexsult.ToString() == "1.4281737E+18471") {
			lang.SayString("xpow298")
		}
	}
	rexsult, err = lang.RxFromString("-6589947.80").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow299")
	} else {
		if !(rexsult.ToString() == "2.30269305E-14") {
			lang.SayString("xpow299")
		}
	}
	rexsult, err = lang.RxFromString("-13.6783690").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-454"))
	if err != nil {
		lang.SayString("xpow301")
	} else {
		if !(rexsult.ToString() == "1.73948535E-516") {
			lang.SayString("xpow301")
		}
	}
	rexsult, err = lang.RxFromString("213577152").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("xpow306")
	} else {
		if !(rexsult.ToString() == "5.06351487E-84") {
			lang.SayString("xpow306")
		}
	}
	rexsult, err = lang.RxFromString("-396.503557").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("xpow308")
	} else {
		if !(rexsult.ToString() == "-9.80021128E+12") {
			lang.SayString("xpow308")
		}
	}
	rexsult, err = lang.RxFromString("59807846.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow309")
	} else {
		if !(rexsult.ToString() == "3.57697846E+15") {
			lang.SayString("xpow309")
		}
	}
	rexsult, err = lang.RxFromString("-8046158.45").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow310")
	} else {
		if !(rexsult.ToString() == "1.75674467E+55") {
			lang.SayString("xpow310")
		}
	}
	rexsult, err = lang.RxFromString("55.1123381E+50627250").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("xpow311")
	} else {
		if !(rexsult.ToString() == "2.13186881E-455645266") {
			lang.SayString("xpow311")
		}
	}
	rexsult, err = lang.RxFromString("-948.038054").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3581"))
	if err != nil {
		lang.SayString("xpow312")
	} else {
		if !(rexsult.ToString() == "-1.03058288E+10660") {
			lang.SayString("xpow312")
		}
	}
	rexsult, err = lang.RxFromString("-6026.42752").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("xpow313")
	} else {
		if !(rexsult.ToString() == "-0.000165935788") {
			lang.SayString("xpow313")
		}
	}
	rexsult, err = lang.RxFromString("79551.5014").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-538"))
	if err != nil {
		lang.SayString("xpow314")
	} else {
		if !(rexsult.ToString() == "2.82599389E-2637") {
			lang.SayString("xpow314")
		}
	}
	rexsult, err = lang.RxFromString("-3264204.54").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-42705"))
	if err != nil {
		lang.SayString("xpow317")
	} else {
		if !(rexsult.ToString() == "-1.3729341E-278171") {
			lang.SayString("xpow317")
		}
	}
	rexsult, err = lang.RxFromString("1.21265492").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("44103"))
	if err != nil {
		lang.SayString("xpow318")
	} else {
		if !(rexsult.ToString() == "1.15662573E+3693") {
			lang.SayString("xpow318")
		}
	}
	rexsult, err = lang.RxFromString("745.78452").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow320")
	} else {
		if !(rexsult.ToString() == "0.00000179793204") {
			lang.SayString("xpow320")
		}
	}
	rexsult, err = lang.RxFromString("82.4185291E-321919303").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow322")
	} else {
		if !(rexsult.ToString() == "1.47214396E+643838602") {
			lang.SayString("xpow322")
		}
	}
	rexsult, err = lang.RxFromString("700592.720").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-698485"))
	if err != nil {
		lang.SayString("xpow324")
	} else {
		if !(rexsult.ToString() == "8.83690001E-4082971") {
			lang.SayString("xpow324")
		}
	}
	rexsult, err = lang.RxFromString("-80273928.0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("661346"))
	if err != nil {
		lang.SayString("xpow325")
	} else {
		if !(rexsult.ToString() == "5.45664856E+5227658") {
			lang.SayString("xpow325")
		}
	}
	rexsult, err = lang.RxFromString("2512953.3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("xpow327")
	} else {
		if !(rexsult.ToString() == "2.50762349E-26") {
			lang.SayString("xpow327")
		}
	}
	rexsult, err = lang.RxFromString("-682.796370").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("71131"))
	if err != nil {
		lang.SayString("xpow328")
	} else {
		if !(rexsult.ToString() == "-9.28114741E+201605") {
			lang.SayString("xpow328")
		}
	}
	rexsult, err = lang.RxFromString("89.9997490").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4994"))
	if err != nil {
		lang.SayString("xpow329")
	} else {
		if !(rexsult.ToString() == "3.30336525E-9760") {
			lang.SayString("xpow329")
		}
	}
	rexsult, err = lang.RxFromString("76563354.6E-112338836").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow330")
	} else {
		if !(rexsult.ToString() == "4.48810347E-337016485") {
			lang.SayString("xpow330")
		}
	}
	rexsult, err = lang.RxFromString("-932499.010").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("xpow331")
	} else {
		if !(rexsult.ToString() == "-5.33132815E+53") {
			lang.SayString("xpow331")
		}
	}
	rexsult, err = lang.RxFromString("-5205124.44E-140588661").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow334")
	} else {
		if !(rexsult.ToString() == "-2.61724523E+702943271") {
			lang.SayString("xpow334")
		}
	}
	rexsult, err = lang.RxFromString("-8868.72074").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5592400"))
	if err != nil {
		lang.SayString("xpow335")
	} else {
		if !(rexsult.ToString() == "5.55074142E+22078017") {
			lang.SayString("xpow335")
		}
	}
	rexsult, err = lang.RxFromString("-74.7852037E-175205809").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("xpow336")
	} else {
		if !(rexsult.ToString() == "3.12797104E-700823229") {
			lang.SayString("xpow336")
		}
	}
	rexsult, err = lang.RxFromString("38660103.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("xpow338")
	} else {
		if !(rexsult.ToString() == "7.7474529E-54") {
			lang.SayString("xpow338")
		}
	}
	rexsult, err = lang.RxFromString("-52.2659460").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow339")
	} else {
		if !(rexsult.ToString() == "-0.00000700395833") {
			lang.SayString("xpow339")
		}
	}
	rexsult, err = lang.RxFromString("6.06625013").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-276"))
	if err != nil {
		lang.SayString("xpow340")
	} else {
		if !(rexsult.ToString() == "8.20339149E-217") {
			lang.SayString("xpow340")
		}
	}
	rexsult, err = lang.RxFromString("-5.36917800").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow342")
	} else {
		if !(rexsult.ToString() == "-0.00646065565") {
			lang.SayString("xpow342")
		}
	}
	rexsult, err = lang.RxFromString("2467915.01").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-93"))
	if err != nil {
		lang.SayString("xpow343")
	} else {
		if !(rexsult.ToString() == "3.26055444E-595") {
			lang.SayString("xpow343")
		}
	}
	rexsult, err = lang.RxFromString("187.232671").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-840"))
	if err != nil {
		lang.SayString("xpow344")
	} else {
		if !(rexsult.ToString() == "1.58280862E-1909") {
			lang.SayString("xpow344")
		}
	}
	rexsult, err = lang.RxFromString("81233.6823").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow345")
	} else {
		if !(rexsult.ToString() == "2.82695763E-25") {
			lang.SayString("xpow345")
		}
	}
	rexsult, err = lang.RxFromString("-854.586113").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow346")
	} else {
		if !(rexsult.ToString() == "3.51522679E-24") {
			lang.SayString("xpow346")
		}
	}
	rexsult, err = lang.RxFromString("78872665.3").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("172"))
	if err != nil {
		lang.SayString("xpow347")
	} else {
		if !(rexsult.ToString() == "1.86793137E+1358") {
			lang.SayString("xpow347")
		}
	}
	rexsult, err = lang.RxFromString("-688755561.E-95301699").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("10"))
	if err != nil {
		lang.SayString("xpow350")
	} else {
		if !(rexsult.ToString() == "2.40243244E-953016902") {
			lang.SayString("xpow350")
		}
	}
	rexsult, err = lang.RxFromString("-5.47345502").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("59819"))
	if err != nil {
		lang.SayString("xpow351")
	} else {
		if !(rexsult.ToString() == "-1.16914146E+44162") {
			lang.SayString("xpow351")
		}
	}
	rexsult, err = lang.RxFromString("-69.7231286").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("85774"))
	if err != nil {
		lang.SayString("xpow353")
	} else {
		if !(rexsult.ToString() == "6.41714261E+158113") {
			lang.SayString("xpow353")
		}
	}
	rexsult, err = lang.RxFromString("5125.51188").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7"))
	if err != nil {
		lang.SayString("xpow354")
	} else {
		if !(rexsult.ToString() == "9.29310216E+25") {
			lang.SayString("xpow354")
		}
	}
	rexsult, err = lang.RxFromString("-54.6254096").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-332921899"))
	if err != nil {
		lang.SayString("xpow355")
	} else {
		if !(rexsult.ToString() == "-1.01482569E-578416745") {
			lang.SayString("xpow355")
		}
	}
	rexsult, err = lang.RxFromString("-1546783").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-51935370"))
	if err != nil {
		lang.SayString("xpow358")
	} else {
		if !(rexsult.ToString() == "3.36022461E-321450306") {
			lang.SayString("xpow358")
		}
	}
	rexsult, err = lang.RxFromString("61302486.8").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("205"))
	if err != nil {
		lang.SayString("xpow359")
	} else {
		if !(rexsult.ToString() == "2.71024755E+1596") {
			lang.SayString("xpow359")
		}
	}
	rexsult, err = lang.RxFromString("-318180109.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow360")
	} else {
		if !(rexsult.ToString() == "-3.0664428E-43") {
			lang.SayString("xpow360")
		}
	}
	rexsult, err = lang.RxFromString("-546398328.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-28"))
	if err != nil {
		lang.SayString("xpow362")
	} else {
		if !(rexsult.ToString() == "2.23737032E-245") {
			lang.SayString("xpow362")
		}
	}
	rexsult, err = lang.RxFromString("18845620").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3129"))
	if err != nil {
		lang.SayString("xpow364")
	} else {
		if !(rexsult.ToString() == "1.35967443E+22764") {
			lang.SayString("xpow364")
		}
	}
	rexsult, err = lang.RxFromString("9820.90457").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("46672"))
	if err != nil {
		lang.SayString("xpow368")
	} else {
		if !(rexsult.ToString() == "4.9475307E+186321") {
			lang.SayString("xpow368")
		}
	}
	rexsult, err = lang.RxFromString("7.22436006E+831949153").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("xpow369")
	} else {
		if !(rexsult.ToString() == "1.38420565E-831949154") {
			lang.SayString("xpow369")
		}
	}
	rexsult, err = lang.RxFromString("472648900").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-208"))
	if err != nil {
		lang.SayString("xpow370")
	} else {
		if !(rexsult.ToString() == "4.96547145E-1805") {
			lang.SayString("xpow370")
		}
	}
	rexsult, err = lang.RxFromString("-8754.49306").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow371")
	} else {
		if !(rexsult.ToString() == "2.89835767E-32") {
			lang.SayString("xpow371")
		}
	}
	rexsult, err = lang.RxFromString("98750864").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("191381"))
	if err != nil {
		lang.SayString("xpow372")
	} else {
		if !(rexsult.ToString() == "1.70908809E+1530003") {
			lang.SayString("xpow372")
		}
	}
	rexsult, err = lang.RxFromString("725292561.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow373")
	} else {
		if !(rexsult.ToString() == "1.30585277E-71") {
			lang.SayString("xpow373")
		}
	}
	rexsult, err = lang.RxFromString("-5549320.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-93580684"))
	if err != nil {
		lang.SayString("xpow375")
	} else {
		if !(rexsult.ToString() == "4.20662079E-631130572") {
			lang.SayString("xpow375")
		}
	}
	rexsult, err = lang.RxFromString("-14677053.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-25785"))
	if err != nil {
		lang.SayString("xpow376")
	} else {
		if !(rexsult.ToString() == "-1.64760831E-184792") {
			lang.SayString("xpow376")
		}
	}
	rexsult, err = lang.RxFromString("-4131738.09").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7579"))
	if err != nil {
		lang.SayString("xpow378")
	} else {
		if !(rexsult.ToString() == "-4.68132794E+50143") {
			lang.SayString("xpow378")
		}
	}
	rexsult, err = lang.RxFromString("504544.648").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow379")
	} else {
		if !(rexsult.ToString() == "2.38124001E-46") {
			lang.SayString("xpow379")
		}
	}
	rexsult, err = lang.RxFromString("829898241").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("xpow380")
	} else {
		if !(rexsult.ToString() == "1.86734084E+80") {
			lang.SayString("xpow380")
		}
	}
	rexsult, err = lang.RxFromString("53.6891691").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-11"))
	if err != nil {
		lang.SayString("xpow381")
	} else {
		if !(rexsult.ToString() == "9.35936725E-20") {
			lang.SayString("xpow381")
		}
	}
	rexsult, err = lang.RxFromString("-93951823.4").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-25318"))
	if err != nil {
		lang.SayString("xpow382")
	} else {
		if !(rexsult.ToString() == "9.67857714E-201859") {
			lang.SayString("xpow382")
		}
	}
	rexsult, err = lang.RxFromString("-8.01787748").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-88"))
	if err != nil {
		lang.SayString("xpow384")
	} else {
		if !(rexsult.ToString() == "2.77186088E-80") {
			lang.SayString("xpow384")
		}
	}
	rexsult, err = lang.RxFromString("517458139").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-999732"))
	if err != nil {
		lang.SayString("xpow385")
	} else {
		if !(rexsult.ToString() == "1.24821346E-8711540") {
			lang.SayString("xpow385")
		}
	}
	rexsult, err = lang.RxFromString("-405543440").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4013"))
	if err != nil {
		lang.SayString("xpow386")
	} else {
		if !(rexsult.ToString() == "-8.83061932E-34545") {
			lang.SayString("xpow386")
		}
	}
	rexsult, err = lang.RxFromString("-151144455").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-170371"))
	if err != nil {
		lang.SayString("xpow388")
	} else {
		if !(rexsult.ToString() == "-5.86496369E-1393532") {
			lang.SayString("xpow388")
		}
	}
	rexsult, err = lang.RxFromString("534.394729").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2369839"))
	if err != nil {
		lang.SayString("xpow390")
	} else {
		if !(rexsult.ToString() == "7.12522896E-6464595") {
			lang.SayString("xpow390")
		}
	}
	rexsult, err = lang.RxFromString("89100.1797").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("224"))
	if err != nil {
		lang.SayString("xpow391")
	} else {
		if !(rexsult.ToString() == "5.92654936E+1108") {
			lang.SayString("xpow391")
		}
	}
	rexsult, err = lang.RxFromString("-821377.777").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("39"))
	if err != nil {
		lang.SayString("xpow392")
	} else {
		if !(rexsult.ToString() == "-4.64702482E+230") {
			lang.SayString("xpow392")
		}
	}
	rexsult, err = lang.RxFromString("-392640.782").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow393")
	} else {
		if !(rexsult.ToString() == "-1.65201422E-17") {
			lang.SayString("xpow393")
		}
	}
	rexsult, err = lang.RxFromString("-651397.712").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-724"))
	if err != nil {
		lang.SayString("xpow394")
	} else {
		if !(rexsult.ToString() == "5.96115415E-4210") {
			lang.SayString("xpow394")
		}
	}
	rexsult, err = lang.RxFromString("4880.06442E-382222621").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("xpow396")
	} else {
		if !(rexsult.ToString() == "2.04915328E+382222617") {
			lang.SayString("xpow396")
		}
	}
	rexsult, err = lang.RxFromString("-1522176.78").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6631062"))
	if err != nil {
		lang.SayString("xpow398")
	} else {
		if !(rexsult.ToString() == "4.54268854E-40996310") {
			lang.SayString("xpow398")
		}
	}
	rexsult, err = lang.RxFromString("968370.780").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6677269"))
	if err != nil {
		lang.SayString("xpow401")
	} else {
		if !(rexsult.ToString() == "3.29990931E+39970410") {
			lang.SayString("xpow401")
		}
	}
	rexsult, err = lang.RxFromString("-97.7474945").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("31248242"))
	if err != nil {
		lang.SayString("xpow402")
	} else {
		if !(rexsult.ToString() == "2.90714257E+62187302") {
			lang.SayString("xpow402")
		}
	}
	rexsult, err = lang.RxFromString("14838.0718").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("9"))
	if err != nil {
		lang.SayString("xpow406")
	} else {
		if !(rexsult.ToString() == "3.48656057E+37") {
			lang.SayString("xpow406")
		}
	}
	rexsult, err = lang.RxFromString("71207472.8").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-31835"))
	if err != nil {
		lang.SayString("xpow407")
	} else {
		if !(rexsult.ToString() == "7.05333953E-249986") {
			lang.SayString("xpow407")
		}
	}
	rexsult, err = lang.RxFromString("-20440.4394").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("xpow408")
	} else {
		if !(rexsult.ToString() == "5.7284759E-18") {
			lang.SayString("xpow408")
		}
	}
	rexsult, err = lang.RxFromString("-54.3684171E-807210192").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("xpow409")
	} else {
		if !(rexsult.ToString() == "-5.43684171E-807210191") {
			lang.SayString("xpow409")
		}
	}
	rexsult, err = lang.RxFromString("-657.186702").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("426844"))
	if err != nil {
		lang.SayString("xpow411")
	} else {
		if !(rexsult.ToString() == "3.50000575E+1202713") {
			lang.SayString("xpow411")
		}
	}
	rexsult, err = lang.RxFromString("-41593077.0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-688608"))
	if err != nil {
		lang.SayString("xpow412")
	} else {
		if !(rexsult.ToString() == "1.4215075E-5246519") {
			lang.SayString("xpow412")
		}
	}
	rexsult, err = lang.RxFromString("-5786.38132").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow413")
	} else {
		if !(rexsult.ToString() == "33482208.8") {
			lang.SayString("xpow413")
		}
	}
	rexsult, err = lang.RxFromString("737622.974").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow414")
	} else {
		if !(rexsult.ToString() == "1.83793916E-12") {
			lang.SayString("xpow414")
		}
	}
	rexsult, err = lang.RxFromString("5615373.52").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-7"))
	if err != nil {
		lang.SayString("xpow415")
	} else {
		if !(rexsult.ToString() == "5.6800146E-48") {
			lang.SayString("xpow415")
		}
	}
	rexsult, err = lang.RxFromString("644136.179").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-835708"))
	if err != nil {
		lang.SayString("xpow416")
	} else {
		if !(rexsult.ToString() == "7.41936858E-4854610") {
			lang.SayString("xpow416")
		}
	}
	rexsult, err = lang.RxFromString("-619642.130").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow418")
	} else {
		if !(rexsult.ToString() == "2.60446259E-12") {
			lang.SayString("xpow418")
		}
	}
	rexsult, err = lang.RxFromString("-31068.7549").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow419")
	} else {
		if !(rexsult.ToString() == "-3.33448258E-14") {
			lang.SayString("xpow419")
		}
	}
	rexsult, err = lang.RxFromString("-68951173.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("xpow420")
	} else {
		if !(rexsult.ToString() == "2.10337488E-16") {
			lang.SayString("xpow420")
		}
	}
	rexsult, err = lang.RxFromString("3898.03188").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-82573"))
	if err != nil {
		lang.SayString("xpow422")
	} else {
		if !(rexsult.ToString() == "1.33010737E-296507") {
			lang.SayString("xpow422")
		}
	}
	rexsult, err = lang.RxFromString("-1.7619356").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-2547"))
	if err != nil {
		lang.SayString("xpow423")
	} else {
		if !(rexsult.ToString() == "-2.90664557E-627") {
			lang.SayString("xpow423")
		}
	}
	rexsult, err = lang.RxFromString("59714.1968").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("xpow424")
	} else {
		if !(rexsult.ToString() == "2.12928005E+14") {
			lang.SayString("xpow424")
		}
	}
	rexsult, err = lang.RxFromString("975566251").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-520"))
	if err != nil {
		lang.SayString("xpow426")
	} else {
		if !(rexsult.ToString() == "3.859053E-4675") {
			lang.SayString("xpow426")
		}
	}
	rexsult, err = lang.RxFromString("-403903.851").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("xpow429")
	} else {
		if !(rexsult.ToString() == "2.66141117E+22") {
			lang.SayString("xpow429")
		}
	}
	rexsult, err = lang.RxFromString("6.48674979").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("xpow430")
	} else {
		if !(rexsult.ToString() == "0.0000134226146") {
			lang.SayString("xpow430")
		}
	}
	rexsult, err = lang.RxFromString("-31401.9418").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("36"))
	if err != nil {
		lang.SayString("xpow431")
	} else {
		if !(rexsult.ToString() == "7.77023505E+161") {
			lang.SayString("xpow431")
		}
	}
	rexsult, err = lang.RxFromString("31345321.1").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("52"))
	if err != nil {
		lang.SayString("xpow432")
	} else {
		if !(rexsult.ToString() == "6.32385059E+389") {
			lang.SayString("xpow432")
		}
	}
	rexsult, err = lang.RxFromString("-64.172844").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow433")
	} else {
		if !(rexsult.ToString() == "-0.00000378395654") {
			lang.SayString("xpow433")
		}
	}
	rexsult, err = lang.RxFromString("70437.1551").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-62916"))
	if err != nil {
		lang.SayString("xpow434")
	} else {
		if !(rexsult.ToString() == "5.0294506E-305005") {
			lang.SayString("xpow434")
		}
	}
	rexsult, err = lang.RxFromString("916260164").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-58"))
	if err != nil {
		lang.SayString("xpow435")
	} else {
		if !(rexsult.ToString() == "1.59554587E-520") {
			lang.SayString("xpow435")
		}
	}
	rexsult, err = lang.RxFromString("-56312.3383").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("789"))
	if err != nil {
		lang.SayString("xpow437")
	} else {
		if !(rexsult.ToString() == "-1.68348724E+3748") {
			lang.SayString("xpow437")
		}
	}
	rexsult, err = lang.RxFromString("859658551.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("72338"))
	if err != nil {
		lang.SayString("xpow440")
	} else {
		if !(rexsult.ToString() == "1.8762045E+646291") {
			lang.SayString("xpow440")
		}
	}
	rexsult, err = lang.RxFromString("-969.881818").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("31171"))
	if err != nil {
		lang.SayString("xpow442")
	} else {
		if !(rexsult.ToString() == "-1.02865894E+93099") {
			lang.SayString("xpow442")
		}
	}
	rexsult, err = lang.RxFromString("7980537.27").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("85"))
	if err != nil {
		lang.SayString("xpow443")
	} else {
		if !(rexsult.ToString() == "4.70685763E+586") {
			lang.SayString("xpow443")
		}
	}
	rexsult, err = lang.RxFromString("-114609916.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("7525"))
	if err != nil {
		lang.SayString("xpow444")
	} else {
		if !(rexsult.ToString() == "-4.43620445E+60645") {
			lang.SayString("xpow444")
		}
	}
	rexsult, err = lang.RxFromString("188006433").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow446")
	} else {
		if !(rexsult.ToString() == "3.53464188E+16") {
			lang.SayString("xpow446")
		}
	}
	rexsult, err = lang.RxFromString("-9.95836312").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-866466703"))
	if err != nil {
		lang.SayString("xpow447")
	} else {
		if !(rexsult.ToString() == "-6.71744369E-864896630") {
			lang.SayString("xpow447")
		}
	}
	rexsult, err = lang.RxFromString("159579.444").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-89828"))
	if err != nil {
		lang.SayString("xpow449")
	} else {
		if !(rexsult.ToString() == "9.6995585E-467374") {
			lang.SayString("xpow449")
		}
	}
	rexsult, err = lang.RxFromString("-4.54000153").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6966334"))
	if err != nil {
		lang.SayString("xpow450")
	} else {
		if !(rexsult.ToString() == "3.52568913E+4577271") {
			lang.SayString("xpow450")
		}
	}
	rexsult, err = lang.RxFromString("-361382575.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow452")
	} else {
		if !(rexsult.ToString() == "3.43765537E-69") {
			lang.SayString("xpow452")
		}
	}
	rexsult, err = lang.RxFromString("7021805.61").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("1222953"))
	if err != nil {
		lang.SayString("xpow453")
	} else {
		if !(rexsult.ToString() == "1.26540553E+8372885") {
			lang.SayString("xpow453")
		}
	}
	rexsult, err = lang.RxFromString("-40.4811667").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-79656"))
	if err != nil {
		lang.SayString("xpow454")
	} else {
		if !(rexsult.ToString() == "4.50174275E-128028") {
			lang.SayString("xpow454")
		}
	}
	rexsult, err = lang.RxFromString("-37958476.0").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("584368"))
	if err != nil {
		lang.SayString("xpow457")
	} else {
		if !(rexsult.ToString() == "3.20538268E+4429105") {
			lang.SayString("xpow457")
		}
	}
	rexsult, err = lang.RxFromString("-502343060").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-96829"))
	if err != nil {
		lang.SayString("xpow459")
	} else {
		if !(rexsult.ToString() == "-6.78602119E-842510") {
			lang.SayString("xpow459")
		}
	}
	rexsult, err = lang.RxFromString("-51592.2698").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-713886"))
	if err != nil {
		lang.SayString("xpow462")
	} else {
		if !(rexsult.ToString() == "6.3857692E-3364249") {
			lang.SayString("xpow462")
		}
	}
	rexsult, err = lang.RxFromString("51.2279848E+80439745").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow463")
	} else {
		if !(rexsult.ToString() == "2.62430643E+160879493") {
			lang.SayString("xpow463")
		}
	}
	rexsult, err = lang.RxFromString("-5983.23468").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-40"))
	if err != nil {
		lang.SayString("xpow464")
	} else {
		if !(rexsult.ToString() == "8.36678291E-152") {
			lang.SayString("xpow464")
		}
	}
	rexsult, err = lang.RxFromString("-6410.5555").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-234964259"))
	if err != nil {
		lang.SayString("xpow467")
	} else {
		if !(rexsult.ToString() == "-1.27064467E-894484419") {
			lang.SayString("xpow467")
		}
	}
	rexsult, err = lang.RxFromString("-5.32711606").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8447286"))
	if err != nil {
		lang.SayString("xpow468")
	} else {
		if !(rexsult.ToString() == "9.09138728E-6136888") {
			lang.SayString("xpow468")
		}
	}
	rexsult, err = lang.RxFromString("-82272171.8").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow469")
	} else {
		if !(rexsult.ToString() == "4.76404994E-64") {
			lang.SayString("xpow469")
		}
	}
	rexsult, err = lang.RxFromString("-103.474598").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("xpow471")
	} else {
		if !(rexsult.ToString() == "-9.02607123E-7") {
			lang.SayString("xpow471")
		}
	}
	rexsult, err = lang.RxFromString("-1199339.72").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-6"))
	if err != nil {
		lang.SayString("xpow473")
	} else {
		if !(rexsult.ToString() == "3.36005741E-37") {
			lang.SayString("xpow473")
		}
	}
	rexsult, err = lang.RxFromString("-2376150.83").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-46777583"))
	if err != nil {
		lang.SayString("xpow475")
	} else {
		if !(rexsult.ToString() == "-3.51886193E-298247976") {
			lang.SayString("xpow475")
		}
	}
	rexsult, err = lang.RxFromString("6.3664211").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-140854908"))
	if err != nil {
		lang.SayString("xpow476")
	} else {
		if !(rexsult.ToString() == "7.25432803E-113232608") {
			lang.SayString("xpow476")
		}
	}
	rexsult, err = lang.RxFromString("-15.791522").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("xpow477")
	} else {
		if !(rexsult.ToString() == "249.372167") {
			lang.SayString("xpow477")
		}
	}
	rexsult, err = lang.RxFromString("49436.6528").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("752"))
	if err != nil {
		lang.SayString("xpow480")
	} else {
		if !(rexsult.ToString() == "8.41185718E+3529") {
			lang.SayString("xpow480")
		}
	}
	rexsult, err = lang.RxFromString("552.669453").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("8"))
	if err != nil {
		lang.SayString("xpow481")
	} else {
		if !(rexsult.ToString() == "8.70409632E+21") {
			lang.SayString("xpow481")
		}
	}
	rexsult, err = lang.RxFromString("-3266303").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("453742"))
	if err != nil {
		lang.SayString("xpow482")
	} else {
		if !(rexsult.ToString() == "1.02497315E+2955701") {
			lang.SayString("xpow482")
		}
	}
	rexsult, err = lang.RxFromString("12302757.4").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("xpow483")
	} else {
		if !(rexsult.ToString() == "2.81846276E+35") {
			lang.SayString("xpow483")
		}
	}
	rexsult, err = lang.RxFromString("-22881.0408").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("xpow487")
	} else {
		if !(rexsult.ToString() == "1.43500909E+26") {
			lang.SayString("xpow487")
		}
	}
	rexsult, err = lang.RxFromString("-7157.57449").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-8"))
	if err != nil {
		lang.SayString("xpow488")
	} else {
		if !(rexsult.ToString() == "1.451687E-31") {
			lang.SayString("xpow488")
		}
	}
	rexsult, err = lang.RxFromString("-503113.801").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-10"))
	if err != nil {
		lang.SayString("xpow489")
	} else {
		if !(rexsult.ToString() == "9.62360287E-58") {
			lang.SayString("xpow489")
		}
	}
	rexsult, err = lang.RxFromString("-3066962.41").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-55"))
	if err != nil {
		lang.SayString("xpow490")
	} else {
		if !(rexsult.ToString() == "-1.702296E-357") {
			lang.SayString("xpow490")
		}
	}
	rexsult, err = lang.RxFromString("998890068.").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-92"))
	if err != nil {
		lang.SayString("xpow492")
	} else {
		if !(rexsult.ToString() == "1.10757225E-828") {
			lang.SayString("xpow492")
		}
	}
	rexsult, err = lang.RxFromString("122.495591").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-407836028"))
	if err != nil {
		lang.SayString("xpow493")
	} else {
		if !(rexsult.ToString() == "4.82463773E-851610754") {
			lang.SayString("xpow493")
		}
	}
	rexsult, err = lang.RxFromString("187098.488").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("6"))
	if err != nil {
		lang.SayString("xpow494")
	} else {
		if !(rexsult.ToString() == "4.28964811E+31") {
			lang.SayString("xpow494")
		}
	}
	rexsult, err = lang.RxFromString("-7.27403536").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("xpow498")
	} else {
		if !(rexsult.ToString() == "-0.0000491046885") {
			lang.SayString("xpow498")
		}
	}
	rexsult, err = lang.RxFromString("-6157.74292").OpPow(lang.RxSetWithDigit(9), lang.RxFromString("-9"))
	if err != nil {
		lang.SayString("xpow499")
	} else {
		if !(rexsult.ToString() == "-7.85608218E-35") {
			lang.SayString("xpow499")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("rpox101")
	} else {
		if !(rexsult.ToString() == "3.4877E-21") {
			lang.SayString("rpox101")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("rpox102")
	} else {
		if !(rexsult.ToString() == "4.3056E-17") {
			lang.SayString("rpox102")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("rpox103")
	} else {
		if !(rexsult.ToString() == "5.3153E-13") {
			lang.SayString("rpox103")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("rpox104")
	} else {
		if !(rexsult.ToString() == "6.5617E-9") {
			lang.SayString("rpox104")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("rpox105")
	} else {
		if !(rexsult.ToString() == "0.000081004") {
			lang.SayString("rpox105")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("rpox106")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("rpox106")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rpox107")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rpox107")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox108")
	} else {
		if !(rexsult.ToString() == "1.524E+8") {
			lang.SayString("rpox108")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox109")
	} else {
		if !(rexsult.ToString() == "1.8814E+12") {
			lang.SayString("rpox109")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rpox110")
	} else {
		if !(rexsult.ToString() == "2.3225E+16") {
			lang.SayString("rpox110")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rpox111")
	} else {
		if !(rexsult.ToString() == "2.8672E+20") {
			lang.SayString("rpox111")
		}
	}
	rexsult, err = lang.RxFromString("415").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox112")
	} else {
		if !(rexsult.ToString() == "1.7223E+5") {
			lang.SayString("rpox112")
		}
	}
	rexsult, err = lang.RxFromString("75").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox113")
	} else {
		if !(rexsult.ToString() == "4.2188E+5") {
			lang.SayString("rpox113")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("rpox201")
	} else {
		if !(rexsult.ToString() == "3.4877E-21") {
			lang.SayString("rpox201")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("rpox202")
	} else {
		if !(rexsult.ToString() == "4.3056E-17") {
			lang.SayString("rpox202")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("rpox203")
	} else {
		if !(rexsult.ToString() == "5.3153E-13") {
			lang.SayString("rpox203")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("rpox204")
	} else {
		if !(rexsult.ToString() == "6.5617E-9") {
			lang.SayString("rpox204")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("rpox205")
	} else {
		if !(rexsult.ToString() == "0.000081004") {
			lang.SayString("rpox205")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("rpox206")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("rpox206")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rpox207")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rpox207")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox208")
	} else {
		if !(rexsult.ToString() == "1.524E+8") {
			lang.SayString("rpox208")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox209")
	} else {
		if !(rexsult.ToString() == "1.8814E+12") {
			lang.SayString("rpox209")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rpox210")
	} else {
		if !(rexsult.ToString() == "2.3225E+16") {
			lang.SayString("rpox210")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rpox211")
	} else {
		if !(rexsult.ToString() == "2.8672E+20") {
			lang.SayString("rpox211")
		}
	}
	rexsult, err = lang.RxFromString("415").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox212")
	} else {
		if !(rexsult.ToString() == "1.7223E+5") {
			lang.SayString("rpox212")
		}
	}
	rexsult, err = lang.RxFromString("75").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox213")
	} else {
		if !(rexsult.ToString() == "4.2188E+5") {
			lang.SayString("rpox213")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("rpox301")
	} else {
		if !(rexsult.ToString() == "3.4877E-21") {
			lang.SayString("rpox301")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("rpox302")
	} else {
		if !(rexsult.ToString() == "4.3056E-17") {
			lang.SayString("rpox302")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("rpox303")
	} else {
		if !(rexsult.ToString() == "5.3153E-13") {
			lang.SayString("rpox303")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("rpox304")
	} else {
		if !(rexsult.ToString() == "6.5617E-9") {
			lang.SayString("rpox304")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("rpox305")
	} else {
		if !(rexsult.ToString() == "0.000081004") {
			lang.SayString("rpox305")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("rpox306")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("rpox306")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rpox307")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rpox307")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox308")
	} else {
		if !(rexsult.ToString() == "1.524E+8") {
			lang.SayString("rpox308")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox309")
	} else {
		if !(rexsult.ToString() == "1.8814E+12") {
			lang.SayString("rpox309")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rpox310")
	} else {
		if !(rexsult.ToString() == "2.3225E+16") {
			lang.SayString("rpox310")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rpox311")
	} else {
		if !(rexsult.ToString() == "2.8672E+20") {
			lang.SayString("rpox311")
		}
	}
	rexsult, err = lang.RxFromString("415").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox312")
	} else {
		if !(rexsult.ToString() == "1.7223E+5") {
			lang.SayString("rpox312")
		}
	}
	rexsult, err = lang.RxFromString("75").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox313")
	} else {
		if !(rexsult.ToString() == "4.2188E+5") {
			lang.SayString("rpox313")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("rpox401")
	} else {
		if !(rexsult.ToString() == "3.4877E-21") {
			lang.SayString("rpox401")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("rpox402")
	} else {
		if !(rexsult.ToString() == "4.3056E-17") {
			lang.SayString("rpox402")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("rpox403")
	} else {
		if !(rexsult.ToString() == "5.3153E-13") {
			lang.SayString("rpox403")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("rpox404")
	} else {
		if !(rexsult.ToString() == "6.5617E-9") {
			lang.SayString("rpox404")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("rpox405")
	} else {
		if !(rexsult.ToString() == "0.000081004") {
			lang.SayString("rpox405")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("rpox406")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("rpox406")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rpox407")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rpox407")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox408")
	} else {
		if !(rexsult.ToString() == "1.524E+8") {
			lang.SayString("rpox408")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox409")
	} else {
		if !(rexsult.ToString() == "1.8814E+12") {
			lang.SayString("rpox409")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rpox410")
	} else {
		if !(rexsult.ToString() == "2.3225E+16") {
			lang.SayString("rpox410")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rpox411")
	} else {
		if !(rexsult.ToString() == "2.8672E+20") {
			lang.SayString("rpox411")
		}
	}
	rexsult, err = lang.RxFromString("415").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox412")
	} else {
		if !(rexsult.ToString() == "1.7223E+5") {
			lang.SayString("rpox412")
		}
	}
	rexsult, err = lang.RxFromString("75").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox413")
	} else {
		if !(rexsult.ToString() == "4.2188E+5") {
			lang.SayString("rpox413")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("rpox501")
	} else {
		if !(rexsult.ToString() == "3.4877E-21") {
			lang.SayString("rpox501")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("rpox502")
	} else {
		if !(rexsult.ToString() == "4.3056E-17") {
			lang.SayString("rpox502")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("rpox503")
	} else {
		if !(rexsult.ToString() == "5.3153E-13") {
			lang.SayString("rpox503")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("rpox504")
	} else {
		if !(rexsult.ToString() == "6.5617E-9") {
			lang.SayString("rpox504")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("rpox505")
	} else {
		if !(rexsult.ToString() == "0.000081004") {
			lang.SayString("rpox505")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("rpox506")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("rpox506")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rpox507")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rpox507")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox508")
	} else {
		if !(rexsult.ToString() == "1.524E+8") {
			lang.SayString("rpox508")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox509")
	} else {
		if !(rexsult.ToString() == "1.8814E+12") {
			lang.SayString("rpox509")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rpox510")
	} else {
		if !(rexsult.ToString() == "2.3225E+16") {
			lang.SayString("rpox510")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rpox511")
	} else {
		if !(rexsult.ToString() == "2.8672E+20") {
			lang.SayString("rpox511")
		}
	}
	rexsult, err = lang.RxFromString("415").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox512")
	} else {
		if !(rexsult.ToString() == "1.7223E+5") {
			lang.SayString("rpox512")
		}
	}
	rexsult, err = lang.RxFromString("75").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox513")
	} else {
		if !(rexsult.ToString() == "4.2188E+5") {
			lang.SayString("rpox513")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("rpox601")
	} else {
		if !(rexsult.ToString() == "3.4877E-21") {
			lang.SayString("rpox601")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("rpox602")
	} else {
		if !(rexsult.ToString() == "4.3056E-17") {
			lang.SayString("rpox602")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("rpox603")
	} else {
		if !(rexsult.ToString() == "5.3153E-13") {
			lang.SayString("rpox603")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("rpox604")
	} else {
		if !(rexsult.ToString() == "6.5617E-9") {
			lang.SayString("rpox604")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("rpox605")
	} else {
		if !(rexsult.ToString() == "0.000081004") {
			lang.SayString("rpox605")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("rpox606")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("rpox606")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rpox607")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rpox607")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox608")
	} else {
		if !(rexsult.ToString() == "1.524E+8") {
			lang.SayString("rpox608")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox609")
	} else {
		if !(rexsult.ToString() == "1.8814E+12") {
			lang.SayString("rpox609")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rpox610")
	} else {
		if !(rexsult.ToString() == "2.3225E+16") {
			lang.SayString("rpox610")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rpox611")
	} else {
		if !(rexsult.ToString() == "2.8672E+20") {
			lang.SayString("rpox611")
		}
	}
	rexsult, err = lang.RxFromString("415").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox612")
	} else {
		if !(rexsult.ToString() == "1.7223E+5") {
			lang.SayString("rpox612")
		}
	}
	rexsult, err = lang.RxFromString("75").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox613")
	} else {
		if !(rexsult.ToString() == "4.2188E+5") {
			lang.SayString("rpox613")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("rpox701")
	} else {
		if !(rexsult.ToString() == "3.4877E-21") {
			lang.SayString("rpox701")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("rpox702")
	} else {
		if !(rexsult.ToString() == "4.3056E-17") {
			lang.SayString("rpox702")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("rpox703")
	} else {
		if !(rexsult.ToString() == "5.3153E-13") {
			lang.SayString("rpox703")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("rpox704")
	} else {
		if !(rexsult.ToString() == "6.5617E-9") {
			lang.SayString("rpox704")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("rpox705")
	} else {
		if !(rexsult.ToString() == "0.000081004") {
			lang.SayString("rpox705")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("rpox706")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("rpox706")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("rpox707")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("rpox707")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox708")
	} else {
		if !(rexsult.ToString() == "1.524E+8") {
			lang.SayString("rpox708")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox709")
	} else {
		if !(rexsult.ToString() == "1.8814E+12") {
			lang.SayString("rpox709")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("rpox710")
	} else {
		if !(rexsult.ToString() == "2.3225E+16") {
			lang.SayString("rpox710")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("rpox711")
	} else {
		if !(rexsult.ToString() == "2.8672E+20") {
			lang.SayString("rpox711")
		}
	}
	rexsult, err = lang.RxFromString("415").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("rpox712")
	} else {
		if !(rexsult.ToString() == "1.7223E+5") {
			lang.SayString("rpox712")
		}
	}
	rexsult, err = lang.RxFromString("75").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("rpox713")
	} else {
		if !(rexsult.ToString() == "4.2188E+5") {
			lang.SayString("rpox713")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-5"))
	if err != nil {
		lang.SayString("r0pox101")
	} else {
		if !(rexsult.ToString() == "3.4877E-21") {
			lang.SayString("r0pox101")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-4"))
	if err != nil {
		lang.SayString("r0pox102")
	} else {
		if !(rexsult.ToString() == "4.3056E-17") {
			lang.SayString("r0pox102")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-3"))
	if err != nil {
		lang.SayString("r0pox103")
	} else {
		if !(rexsult.ToString() == "5.3153E-13") {
			lang.SayString("r0pox103")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-2"))
	if err != nil {
		lang.SayString("r0pox104")
	} else {
		if !(rexsult.ToString() == "6.5617E-9") {
			lang.SayString("r0pox104")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("-1"))
	if err != nil {
		lang.SayString("r0pox105")
	} else {
		if !(rexsult.ToString() == "0.000081004") {
			lang.SayString("r0pox105")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("0"))
	if err != nil {
		lang.SayString("r0pox106")
	} else {
		if !(rexsult.ToString() == "1") {
			lang.SayString("r0pox106")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("1"))
	if err != nil {
		lang.SayString("r0pox107")
	} else {
		if !(rexsult.ToString() == "12345") {
			lang.SayString("r0pox107")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("r0pox108")
	} else {
		if !(rexsult.ToString() == "1.524E+8") {
			lang.SayString("r0pox108")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("r0pox109")
	} else {
		if !(rexsult.ToString() == "1.8814E+12") {
			lang.SayString("r0pox109")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("4"))
	if err != nil {
		lang.SayString("r0pox110")
	} else {
		if !(rexsult.ToString() == "2.3225E+16") {
			lang.SayString("r0pox110")
		}
	}
	rexsult, err = lang.RxFromString("12345").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("5"))
	if err != nil {
		lang.SayString("r0pox111")
	} else {
		if !(rexsult.ToString() == "2.8672E+20") {
			lang.SayString("r0pox111")
		}
	}
	rexsult, err = lang.RxFromString("415").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("2"))
	if err != nil {
		lang.SayString("r0pox112")
	} else {
		if !(rexsult.ToString() == "1.7223E+5") {
			lang.SayString("r0pox112")
		}
	}
	rexsult, err = lang.RxFromString("75").OpPow(lang.RxSetWithDigit(5), lang.RxFromString("3"))
	if err != nil {
		lang.SayString("r0pox113")
	} else {
		if !(rexsult.ToString() == "4.2188E+5") {
			lang.SayString("r0pox113")
		}
	}

	return
}
