package main

import "fmt"
import "math"
import "reflect"
import "strconv"
import "strings"
import "unicode"

import "agorexx/lang"

var rexx_set = lang.RxSet()
var lower_a = lang.RxFromRune('a')
var upper_a = lang.RxFromRune('A')
var value_a = lang.RxFromInt32(65)
var emptyrex = lang.RxFromEmpty()
var emptystr = lang.RxFromString("")
var blank = lang.RxFromString(" ")
var zero = lang.RxFromInt32(0)
var one = lang.RxFromInt32(1)
var one_center = lang.RxFromString(" 1 ")
var one_leading_zeros = lang.RxFromString("000001")
var one_left = lang.RxFromString("1 ")
var one_negative = lang.RxFromInt32(-1)
var one_over_max = lang.RxFromInt32(lang.MaxArg + 1)
var one_point_zero = lang.RxFromString("1.0")
var one_right = lang.RxFromString(" 1")
var two = lang.RxFromInt32(2)
var three = lang.RxFromInt32(3)
var four = lang.RxFromInt32(4)
var five = lang.RxFromInt32(5)
var six = lang.RxFromInt32(6)
var seven = lang.RxFromInt32(7)
var eight = lang.RxFromInt32(8)
var a = lang.RxFromInt32(50)
var b = lang.RxFromInt32(17)
var c = lang.RxFromInt32(-100)
var e = lang.RxFromString("001001")
var h = lang.RxFromString("0.0000001")
var i = lang.RxFromString("100")
var with_ab = lang.RxFromString("ab")
var x = lang.RxFromString("7.62939453E+28")
var z = lang.RxFromString("-000001.12345")

func main() {

	rexx_set = nil

	if !(lang.RxFromInt64(-9223372036854775808).ToString() == "-9223372036854775808") {
		lang.SayString("BUG1001")
	}
	if !(lang.RxFromInt64(9223372036854775807).ToString() == "9223372036854775807") {
		lang.SayString("BUG1002")
	}
	if !(zero.ToString() == "0") {
		lang.SayString("BUG1003")
	}
	// uint64 as string
	if !(lang.RxFromString("18446744073709551615").ToString() == "18446744073709551615") {
		lang.SayString("BUG1004")
	}
	notanumb := fmt.Sprintf("%v", reflect.ValueOf(lang.RxFromString("+++")).Elem().Field(1))
	if !(notanumb == "-2") {
		lang.SayString("BUG1005")
	}
	if !(lang.RxFromRune(0x4F).ToString() == "O") {
		lang.SayString("BUG1006")
	}
	if !(lang.RxFromRune('\u06F1').ToString() == "۱") {
		lang.SayString("BUG1007")
	}
	if !(lang.RxFromRune(0x06F1).ToString() == "۱") {
		lang.SayString("BUG1008")
	}
	if !(lang.RxFromRune(1777).ToString() == "۱") {
		lang.SayString("BUG1009")
	}
	if !(lang.RxFromRune('۱').ToString() == "۱") {
		lang.SayString("BUG1010")
	}
	if !(lang.RxFromString("ᎠᎡᎢᎣᎤᎥᎦᎧᎨᎩᎪᎫᎬᎭᎮᎯᎰᎱᎲᎳᎴᎵᎶᎷᎸᎹᎺᎻᎼᎽᎾᎿᏀᏁᏂᏃᏄᏅᏆᏇᏈᏉᏊᏋᏌᏍᏎᏏᏐᏑᏒᏓᏔᏕᏖᏗᏘᏙᏚᏛᏜᏝᏞᏟᏠᏡᏢᏣᏤᏥᏦᏧᏨᏩᏪᏫᏬᏭᏮᏯᏰᏱᏲᏳᏴ ").ToString() == "ᎠᎡᎢᎣᎤᎥᎦᎧᎨᎩᎪᎫᎬᎭᎮᎯᎰᎱᎲᎳᎴᎵᎶᎷᎸᎹᎺᎻᎼᎽᎾᎿᏀᏁᏂᏃᏄᏅᏆᏇᏈᏉᏊᏋᏌᏍᏎᏏᏐᏑᏒᏓᏔᏕᏖᏗᏘᏙᏚᏛᏜᏝᏞᏟᏠᏡᏢᏣᏤᏥᏦᏧᏨᏩᏪᏫᏬᏭᏮᏯᏰᏱᏲᏳᏴ ") {
		lang.SayString("BUG1011")
	}
	rexx_clone := lang.RxFromClone(a)
	if !rexx_clone.Equals(lang.RxFromString("50")) {
		lang.SayString("BUG1012")
	}
	if !(lang.RxFromBool(true).ToString() == "1") {
		lang.SayString("BUG1013")
	}
	if !(lang.RxFromBool(false).ToString() == "0") {
		lang.SayString("BUG1014")
	}
	rxbool, err := lang.RxFromBool(true).ToBool()
	if !rxbool {
		lang.SayString("BUG1015")
	}
	rxbool, err = lang.RxFromBool(false).ToBool()
	if rxbool {
		lang.SayString("BUG1016")
	}
	rxbool, err = one.ToBool()
	if !rxbool {
		lang.SayString("BUG1017")
	}
	rxbool, err = zero.ToBool()
	if rxbool {
		lang.SayString("BUG1018")
	}
	rxint8, err := lang.RxFromInt8(int8(127)).ToInt8()
	if rxint8 != 127 {
		lang.SayString("BUG1019")
	}
	rxint8, err = lang.RxFromInt8(int8(-128)).ToInt8()
	if rxint8 != -128 {
		lang.SayString("BUG1020")
	}
	rxint8, err = lang.RxFromInt8(int8(0)).ToInt8()
	if rxint8 != 0 {
		lang.SayString("BUG1021")
	}
	rxint8, err = lang.RxFromString("127").ToInt8()
	if rxint8 != 127 {
		lang.SayString("BUG1022")
	}
	rxint8, err = lang.RxFromString("-128").ToInt8()
	if rxint8 != -128 {
		lang.SayString("BUG1023")
	}
	rxint8, err = zero.ToInt8()
	if rxint8 != 0 {
		lang.SayString("BUG1024")
	}
	rxint8, err = lang.RxFromInt32(127).ToInt8()
	if rxint8 != 127 {
		lang.SayString("BUG1025")
	}
	rxint8, err = lang.RxFromInt32(-128).ToInt8()
	if rxint8 != -128 {
		lang.SayString("BUG1026")
	}
	rxint8, err = zero.ToInt8()
	if rxint8 != 0 {
		lang.SayString("BUG1027")
	}
	rxint8, err = lang.RxFromInt16(127).ToInt8()
	if rxint8 != 127 {
		lang.SayString("BUG1028")
	}
	rxint8, err = lang.RxFromInt16(-128).ToInt8()
	if rxint8 != -128 {
		lang.SayString("BUG1029")
	}
	rxint8, err = lang.RxFromInt16(0).ToInt8()
	if rxint8 != 0 {
		lang.SayString("BUG1030")
	}
	rxint8, err = lang.RxFromInt64(127).ToInt8()
	if rxint8 != 127 {
		lang.SayString("BUG1031")
	}
	rxint8, err = lang.RxFromInt64(-128).ToInt8()
	if rxint8 != -128 {
		lang.SayString("BUG1032")
	}
	rxint8, err = lang.RxFromInt64(0).ToInt8()
	if rxint8 != 0 {
		lang.SayString("BUG1033")
	}

	rxint8, err = lang.RxFromFloat32(127.0).ToInt8()
	if rxint8 != 127 {
		lang.SayString("BUG1034")
	}
	rxint8, err = lang.RxFromFloat32(-128.0).ToInt8()
	if rxint8 != -128 {
		lang.SayString("BUG1035")
	}
	rxint8, err = lang.RxFromFloat32(0.0).ToInt8()
	if rxint8 != 0 {
		lang.SayString("BUG1036")
	}
	rxint8, err = lang.RxFromFloat64(127.0).ToInt8()
	if rxint8 != 127 {
		lang.SayString("BUG1037")
	}
	rxint8, err = lang.RxFromFloat64(-128.0).ToInt8()
	if rxint8 != -128 {
		lang.SayString("BUG1038")
	}
	rxint8, err = lang.RxFromFloat64(0.0).ToInt8()
	if rxint8 != 0 {
		lang.SayString("BUG1039")
	}
	// FIXME NULL vs EMPTY?
	if emptyrex.ToString() != "" {
		lang.SayString("BUG1040")
	}
	rxint, err := lang.RxFromInt32(2147483647).ToInt32()
	if rxint != 2147483647 {
		lang.SayString("BUG1041")
	}
	rxint, err = lang.RxFromEmpty().ToInt32()
	if rxint != 0 {
		lang.SayString("BUG5041")
	}
	rxint, err = lang.RxFromInt32(-2147483648).ToInt32()
	if rxint != -2147483648 {
		lang.SayString("BUG1042")
	}
	rxint, err = zero.ToInt32()
	if rxint != 0 {
		lang.SayString("BUG1043")
	}
	if lang.ToString(upper_a) != "A" {
		lang.SayString("BUG1044")
	}
	if lang.ToString(lang.RxFromString("@")) != "@" {
		lang.SayString("BUG1045")
	}
	if lang.ToString(lang.RxFromString("8")) != "8" {
		lang.SayString("BUG1046")
	}
	if lang.ToString(zero) != "0" {
		lang.SayString("BUG1047")
	}
	if lang.ToString(lang.RxFromString("A@80")) != "A@80" {
		lang.SayString("BUG1048")
	}
	//assert ("" == makerexx ("\0").to_str ()) ;
	if lang.ToString(emptystr) != "" {
		lang.SayString("BUG1049")
	}
	if lang.ToString(lang.RxFromString("   ")) != "   " {
		lang.SayString("BUG1050")
	}
	data := "A1 B2 C3 D4 E5 F6 G7"
	test_str_array := strings.Split(data, " ")
	if lang.RxFromStrings(test_str_array).ToString() != data {
		lang.SayString("BUG1051")
	}
	empty := []string{}
	if lang.RxFromStrings(empty).ToString() != "" {
		lang.SayString("BUG1052")
	}
	if lang.RxFromRune('A').ToString() != "A" {
		lang.SayString("BUG1053")
	}
	if lang.RxFromRune('@').ToString() != "@" {
		lang.SayString("BUG1054")
	}
	if lang.RxFromRune('8').ToString() != "8" {
		lang.SayString("BUG1055")
	}
	if lang.RxFromRune('0').ToString() != "0" {
		lang.SayString("BUG1056")
	}
	if lang.RxFromRune(0x0).ToString() != string(rune(0)) {
		lang.SayString("BUG1057")
	}
	default_char_array := []rune{'R', 'E', 'X', 'X', 0x0}
	if !(lang.RxFromRunes(default_char_array).ToString() == string(default_char_array)) {
		lang.SayString("BUG1058")
	}
	num_a, err := z.OpAdd(rexx_set, x)
	num_b, err := x.OpAdd(rexx_set, z)
	if !(num_a.ToString() == num_b.ToString()) {
		lang.SayString("BUG1059")
	}
	if !(num_a.OpEqS(nil, num_b) == num_b.OpEqS(nil, num_a)) {
		lang.SayString("BUG1060")
	}
	opadd, err := lang.RxFromString("abc").OpAdd(rexx_set, emptyrex)
	if err == nil {
		lang.SayString("BUG1061")
		fmt.Println(opadd)
	}
	opadd, err = lang.RxFromInt32(120000000).OpAdd(nil, lang.RxFromString("-999999999999"))
	if !(opadd.ToString() == "-9.99880000E+11") {
		lang.SayString("BUG4060")
	}
	opadd, err = lang.RxFromString("  1E+5").OpAdd(nil, lang.RxFromInt8(int8(-128)))
	if !(opadd.ToString() == "99872") {
		lang.SayString("BUG4060")
	}
	opadd, err = lang.RxFromString("-34359738368").OpAdd(nil, lang.RxFromInt8(int8(-128)))
	if !(opadd.ToString() == "-3.43597385E+10") {
		lang.SayString("BUG4060")
	}
	opadd, err = lang.RxFromString("1.235E+02").OpAdd(nil, lang.RxFromInt8(int8(-128)))
	if !(opadd.ToString() == "-4.5") {
		lang.SayString("BUG4060")
	}
	opadd, err = emptyrex.OpAdd(rexx_set, lang.RxFromString("abc"))
	if err == nil {
		lang.SayString("BUG1062")
	}
	opand, err := one.OpAnd(rexx_set, one)
	if !opand {
		lang.SayString("BUG1063")
	}
	opand, err = one.OpAnd(rexx_set, zero)
	if opand {
		lang.SayString("BUG1064")
	}
	opand, err = zero.OpAnd(rexx_set, one)
	if opand {
		lang.SayString("BUG1065")
	}
	opand, err = zero.OpAnd(rexx_set, zero)
	if opand {
		lang.SayString("BUG1066")
	}
	opand, err = lang.RxFromString("2").OpAnd(rexx_set, emptyrex)
	if err == nil {
		lang.SayString("BUG1067")
	}
	opand, err = zero.OpAnd(rexx_set, lang.RxFromString("2"))
	if err == nil {
		lang.SayString("BUG1068")
	}
	opcc := one.OpCc(rexx_set, zero)
	if !(opcc.ToString() == "10") {
		lang.SayString("BUG1069")
	}

	opccblank := one.OpCcblank(rexx_set, zero)
	if !(opccblank.ToString() == "1 0") {
		lang.SayString("BUG1070")
	}
	opdiv, err := one.OpDiv(rexx_set, two)
	if !(opdiv.ToString() == "0.5") {
		lang.SayString("BUG1071")
	}

	opdiv, err = one.OpDiv(rexx_set, lower_a)

	opdiv, err = lang.RxFromString("-34359738368").OpDiv(nil, lang.RxFromInt8(int8(-128)))
	opdiv, err = lower_a.OpDiv(rexx_set, two)

	opdivi, err := three.OpDivI(rexx_set, two)
	if !(opdivi.ToString() == "1") {
		lang.SayString("BUG1072")
	}
	opdivi, err = h.OpDivI(rexx_set, i)
	if !(opdivi.ToString() == "0") {
		lang.SayString("BUG1073")
	}
	opdivi, err = lang.RxFromRune('\ubEEf').OpDiv(nil, one_negative)
	if err == nil {
		lang.SayString("BUG4073")
	}
	opeq, err := one.OpEq(rexx_set, two)
	if opeq {
		lang.SayString("BUG1074")
	}
	opeq, err = two.OpEq(rexx_set, two)
	if !opeq {
		lang.SayString("BUG1075")
	}
	opeq, err = one.OpEq(rexx_set, one_point_zero)
	if !opeq {
		lang.SayString("BUG1076")
	}
	opeq, err = with_ab.OpEq(rexx_set, upper_a)
	if opeq {
		lang.SayString("BUG1077")
	}
	opeq, err = upper_a.OpEq(rexx_set, lower_a)
	if !opeq {
		lang.SayString("BUG1078")
	}
	opeq, err = one_right.OpEq(rexx_set, one_right)
	if !opeq {
		lang.SayString("BUG1079")
	}
	opeq, err = one_right.OpEq(rexx_set, one_left)
	if !opeq {
		lang.SayString("BUG1080")
	}
	opeq, err = one_right.OpEq(rexx_set, one_center)
	if !opeq {
		lang.SayString("BUG1081")
	}
	opeq, err = one_left.OpEq(rexx_set, one_right)
	if !opeq {
		lang.SayString("BUG1082")
	}
	opeq, err = one_left.OpEq(rexx_set, one_left)
	if !opeq {
		lang.SayString("BUG1083")
	}
	opeq, err = one_left.OpEq(rexx_set, one_center)
	if !opeq {
		lang.SayString("BUG1084")
	}
	opeq, err = one_center.OpEq(rexx_set, one_right)
	if !opeq {
		lang.SayString("BUG1085")
	}
	opeq, err = one_center.OpEq(rexx_set, one_left)
	if !opeq {
		lang.SayString("BUG1086")
	}
	opeq, err = one_center.OpEq(rexx_set, one_center)
	if !opeq {
		lang.SayString("BUG1087")
	}
	opeq, err = value_a.OpEq(rexx_set, one_leading_zeros)
	if opeq {
		lang.SayString("BUG1088")
	}
	opeq, err = value_a.OpEq(rexx_set, e)
	if opeq {
		lang.SayString("BUG1089")
	}
	// rcvr.chars nil defaults to zero
	opeq, err = emptystr.OpEq(rexx_set, emptyrex)
	if err != nil {
		lang.SayString("BUG1090")
	}
	opeqs := one.OpEqS(rexx_set, two)
	if opeqs {
		lang.SayString("BUG1091")
	}
	opeqs = two.OpEqS(rexx_set, two)
	if !opeqs {
		lang.SayString("BUG1092")
	}
	opeqs = one.OpEqS(rexx_set, one_point_zero)
	if opeqs {
		lang.SayString("BUG1093")
	}
	opeqs = one_right.OpEqS(rexx_set, one_right)
	if !opeqs {
		lang.SayString("BUG1094")
	}
	opeqs = one_right.OpEqS(rexx_set, one_left)
	if opeqs {
		lang.SayString("BUG1095")
	}
	opeqs = one_right.OpEqS(rexx_set, one_center)
	if opeqs {
		lang.SayString("BUG1096")
	}
	opeqs = one_left.OpEqS(rexx_set, one_right)
	if opeqs {
		lang.SayString("BUG1097")
	}
	opeqs = one_left.OpEqS(rexx_set, one_left)
	if !opeqs {
		lang.SayString("BUG1098")
	}
	opeqs = one_left.OpEqS(rexx_set, one_center)
	if opeqs {
		lang.SayString("BUG1099")
	}
	opeqs = one_center.OpEqS(rexx_set, one_right)
	if opeqs {
		lang.SayString("BUG1100")
	}
	opeqs = one_center.OpEqS(rexx_set, one_left)
	if opeqs {
		lang.SayString("BUG1101")
	}
	opeqs = one_center.OpEqS(rexx_set, one_center)
	if !opeqs {
		lang.SayString("BUG1102")
	}
	opgt, err := one.OpGt(rexx_set, two)
	if opgt {
		lang.SayString("BUG1103")
	}
	opgt, err = two.OpGt(rexx_set, one)
	if !opgt {
		lang.SayString("BUG1104")
	}
	opgt, err = value_a.OpGt(rexx_set, upper_a)
	if opgt {
		lang.SayString("BUG1105")
	}
	opgt, err = upper_a.OpGt(rexx_set, value_a)
	if !opgt {
		lang.SayString("BUG1106")
	}
	opgteq, err := one.OpGtEq(rexx_set, two)
	if opgteq {
		lang.SayString("BUG1107")
	}
	opgteq, err = two.OpGtEq(rexx_set, one)
	if !opgteq {
		lang.SayString("BUG1108")
	}
	opgteq, err = two.OpGtEq(rexx_set, two)
	if !opgteq {
		lang.SayString("BUG1109")
	}
	opgteqs := one.OpGtEqS(rexx_set, two)
	if opgteqs {
		lang.SayString("BUG1110")
	}
	opgteqs = two.OpGtEqS(rexx_set, one)
	if !opgteqs {
		lang.SayString("BUG1111")
	}
	opgteqs = two.OpGtEqS(rexx_set, two)
	if !opgteqs {
		lang.SayString("BUG1112")
	}
	opgteqs = one.OpGtEqS(rexx_set, one_point_zero)
	if opgteqs {
		lang.SayString("BUG1113")
	}
	opgts := one.OpGtS(rexx_set, two)
	if opgts {
		lang.SayString("BUG1114")
	}
	opgts = two.OpGtS(rexx_set, one)
	if !opgts {
		lang.SayString("BUG1115")
	}
	opgts = one.OpGtS(rexx_set, one_point_zero)
	if opgts {
		lang.SayString("BUG1116")
	}
	oplt, err := one.OpLt(rexx_set, two)
	if !oplt {
		lang.SayString("BUG1117")
	}
	oplt, err = two.OpLt(rexx_set, one)
	if oplt {
		lang.SayString("BUG1118")
	}
	//FIXME
	oplt, err = emptyrex.OpLt(rexx_set, emptystr)
	if err != nil {
		lang.SayString("BUG1119")
	}
	oplteq, err := one.OpLtEq(rexx_set, two)
	if !oplteq {
		lang.SayString("BUG1120")
	}
	oplteq, err = two.OpLtEq(rexx_set, one)
	if oplteq {
		lang.SayString("BUG1121")
	}
	oplteq, err = two.OpLtEq(rexx_set, two)
	if !oplteq {
		lang.SayString("BUG1122")
	}
	oplteqs := one.OpLtEqS(rexx_set, two)
	if !oplteqs {
		lang.SayString("BUG1123")
	}
	oplteqs = two.OpLtEqS(rexx_set, one)
	if oplteqs {
		lang.SayString("BUG1124")
	}
	oplteqs = two.OpLtEqS(rexx_set, two)
	if !oplteqs {
		lang.SayString("BUG1125")
	}
	oplteqs = one.OpLtEqS(rexx_set, one_point_zero)
	if !oplteqs {
		lang.SayString("BUG1126")
	}
	oplteqs = one_point_zero.OpLtEqS(rexx_set, one)
	if oplteqs {
		lang.SayString("BUG1127")
	}
	oplts := one.OpLtS(rexx_set, two)
	if !oplts {
		lang.SayString("BUG1128")
	}
	oplts = two.OpLtS(rexx_set, one)
	if oplts {
		lang.SayString("BUG1129")
	}
	oplts = one.OpLtS(rexx_set, one_point_zero)
	if !oplts {
		lang.SayString("BUG1130")
	}
	oplts = one_point_zero.OpLtS(rexx_set, one)
	if oplts {
		lang.SayString("BUG1131")
	}
	opminus, err := a.OpMinus(rexx_set)
	if !(opminus.ToString() == "-50") {
		lang.SayString("BUG1132")
	}
	opminus, err = zero.OpMinus(rexx_set)
	if !(opminus.ToString() == "0") {
		lang.SayString("BUG1133")
	}
	opmult, err := a.OpMult(rexx_set, b)
	if !(opmult.ToString() == "850") {
		lang.SayString("BUG1134")
	}
	opmult, err = a.OpMult(rexx_set, zero)
	if !(opmult.ToString() == "0") {
		lang.SayString("BUG1135")
	}
	opmult, err = lang.RxFromString("-34359738368").OpMult(nil, lang.RxFromInt8(int8(-128)))
	if !(opmult.ToString() == "4.39804651E+12") {
		lang.SayString("BUG1135")
	}
	opmult, err = lang.RxFromRune(' ').OpMult(nil, lang.RxFromInt8(int8(-128)))
	if err == nil {
		lang.SayString("BUG4135")
	}
	opmult, err = a.OpMult(rexx_set, lower_a)
	if err == nil {
		lang.SayString("BUG5135")
	}
	opmult, err = lang.RxFromInt8(int8(-128)).OpMult(nil, lang.RxFromString("-34359738368"))

	opnot, err := one.OpNot(rexx_set)
	if opnot {
		lang.SayString("BUG1136")
	}
	opnot, err = zero.OpNot(rexx_set)
	if !opnot {
		lang.SayString("BUG1137")
	}

	opnot, err = lower_a.OpNot(rexx_set)

	opnoteq, err := one.OpNotEq(rexx_set, two)
	if !opnoteq {
		lang.SayString("BUG1138")
	}
	opnoteq, err = two.OpNotEq(rexx_set, two)
	if opnoteq {
		lang.SayString("BUG1139")
	}
	opnoteq, err = one.OpNotEq(rexx_set, one_point_zero)
	if opnoteq {
		lang.SayString("BUG1140")
	}
	opnoteqs := one.OpNotEqS(rexx_set, two)
	if !opnoteqs {
		lang.SayString("BUG1141")
	}
	opnoteqs = two.OpNotEqS(rexx_set, two)
	if opnoteqs {
		lang.SayString("BUG1142")
	}
	opnoteqs = one.OpNotEqS(rexx_set, one_point_zero)
	if !opnoteqs {
		lang.SayString("BUG1143")
	}
	opor, err := one.OpOr(rexx_set, one)
	if !opor {
		lang.SayString("BUG1144")
	}
	opor, err = one.OpOr(rexx_set, zero)
	if !opor {
		lang.SayString("BUG1145")
	}
	opor, err = zero.OpOr(rexx_set, one)
	if !opor {
		lang.SayString("BUG1146")
	}
	opor, err = zero.OpOr(rexx_set, zero)
	if opor {
		lang.SayString("BUG1147")
	}
	opor, err = lang.RxFromString("2").OpOr(rexx_set, emptyrex)
	if err == nil {
		lang.SayString("BUG1148")
	}
	opor, err = zero.OpOr(rexx_set, lang.RxFromString("2"))
	if err == nil {
		lang.SayString("BUG1149")
	}
	opplus, err := c.OpPlus(rexx_set)
	if !(opplus.ToString() == "-100") {
		lang.SayString("BUG1150")
	}
	opplus, err = zero.OpPlus(rexx_set)
	if !(opplus.ToString() == "0") {
		lang.SayString("BUG1151")
	}

	oppow, err := a.OpPow(rexx_set, b)
	if !(oppow.ToString() == "7.62939453E+28") {
		lang.SayString("BUG1152")
	}

	oppow, err = a.OpPow(rexx_set, zero)
	if !(oppow.ToString() == "1") {
		lang.SayString("BUG1153")
	}
	oppow, err = lang.RxFromRune('/').OpPow(nil, lang.RxFromInt8(int8(-128)))
	if err == nil {
		lang.SayString("BUG4153")
	}
	oppow, err = lang.RxFromRune('0').OpPow(nil, lang.RxFromInt8(int8(-128)))
	if err == nil {
		lang.SayString("BUG3153")
	}
	oppow, err = lang.RxFromString("50.000000").OpPow(rexx_set, lang.RxFromString("17.000"))
	if err != nil {
		lang.SayString("BUG5153")
	}

	oppow, err = lang.RxFromString("-34359738368").OpPow(nil, lang.RxFromInt8(int8(-128)))
	oppow, err = lang.RxFromInt8(int8(-128)).OpPow(nil, lang.RxFromString("-34359738368"))

	oprem, err := a.OpRem(rexx_set, b)
	if !(oprem.ToString() == "16") {
		lang.SayString("BUG1154")
	}
	oprem, err = zero.OpRem(rexx_set, a)
	if !(oprem.ToString() == "0") {
		lang.SayString("BUG1155")
	}
	oprem, err = one.OpRem(rexx_set, three)
	if !(oprem.ToString() == "1") {
		lang.SayString("BUG1156")
	}
	opsub, err := a.OpSub(rexx_set, b)
	if !(opsub.ToString() == "33") {
		lang.SayString("BUG1157")
	}
	opsub, err = zero.OpSub(rexx_set, a)
	if !(opsub.ToString() == "-50") {
		lang.SayString("BUG1158")
	}
	opsub, err = lang.RxFromString("abc").OpSub(rexx_set, emptyrex)
	if err == nil {
		lang.SayString("BUG1159")
	}
	opsub, err = emptyrex.OpSub(rexx_set, lang.RxFromString("abc"))
	if err == nil {
		lang.SayString("BUG1160")
	}
	opxor, err := one.OpXor(rexx_set, one)
	if opxor {
		lang.SayString("BUG1161")
	}
	opxor, err = one.OpXor(rexx_set, zero)
	if !opxor {
		lang.SayString("BUG1162")
	}
	opxor, err = zero.OpXor(rexx_set, one)
	if !opxor {
		lang.SayString("BUG1163")
	}
	opxor, err = zero.OpXor(rexx_set, zero)
	if opxor {
		lang.SayString("BUG1164")
	}
	opxor, err = lang.RxFromString("2").OpXor(rexx_set, emptyrex)
	if err == nil {
		lang.SayString("BUG1165")
	}
	opxor, err = zero.OpXor(rexx_set, lang.RxFromString("2"))
	if err == nil {
		lang.SayString("BUG1166")
	}
	abbrev, err := lang.RxFromString("Print").Abbrev(lang.RxFromString("Pri"), lang.RxFromString("Pri").Length())
	if !(abbrev.ToString() == "1") {
		lang.SayString("BUG1167")
	}
	abbrev, err = lang.RxFromString("PRINT").Abbrev(lang.RxFromString("Pri"), lang.RxFromString("Pri").Length())
	if !(abbrev.ToString() == "0") {
		lang.SayString("BUG1168")
	}
	abbrev, err = lang.RxFromString("PRINT").Abbrev(lang.RxFromString("PRI"), four)
	if !(abbrev.ToString() == "0") {
		lang.SayString("BUG1169")
	}
	abbrev, err = lang.RxFromString("PRINT").Abbrev(lang.RxFromString("PRY"), lang.RxFromString("PRY").Length())
	if !(abbrev.ToString() == "0") {
		lang.SayString("BUG1170")
	}
	abbrev, err = lang.RxFromString("PRINT").Abbrev(emptystr, emptystr.Length())
	if !(abbrev.ToString() == "1") {
		lang.SayString("BUG1171")
	}
	abbrev, err = lang.RxFromString("PRINT").Abbrev(emptystr, one)
	if !(abbrev.ToString() == "0") {
		lang.SayString("BUG1172")
	}
	abbrev, err = emptyrex.Abbrev(emptystr, one)
	if !(abbrev.ToString() == "0") {
		lang.SayString("BUG1173")
	}
	//FIXME
	abbrev, err = emptyrex.Abbrev(emptyrex, one)
	if err != nil {
		lang.SayString("BUG1174")
	}
	abbrev, err = lower_a.Abbrev(lang.RxFromString("b"), one_over_max)
	if err == nil {
		lang.SayString("BUG1175")
	}
	abs, err := lang.RxFromString("12.3").Abs()
	if !(abs.ToString() == "12.3") {
		lang.SayString("BUG1176")
	}
	abs, err = lang.RxFromString(" -0.307").Abs()
	if !(abs.ToString() == "0.307") {
		lang.SayString("BUG1177")
	}
	abs, err = lang.RxFromString("123.45E+16").Abs()
	if !(abs.ToString() == "1.2345E+18") {
		lang.SayString("BUG1178")
	}
	abs, err = lang.RxFromString("- 1234567.7654321").Abs()
	if !(abs.ToString() == "1234567.7654321") {
		lang.SayString("BUG1179")
	}
	abs, err = lang.RxFromString("xyz").Abs()
	if err == nil {
		lang.SayString("BUG1180")
	}
	b2d, err := lang.RxFromString("01110").B2d(one_negative)
	if !(b2d.ToString() == "14") {
		lang.SayString("BUG1181")
	}
	b2d, err = lang.RxFromString("01110").B2d(one_negative)
	if !(b2d.ToString() == "14") {
		lang.SayString("BUG1182")
	}
	b2d, err = lang.RxFromString("10000001").B2d(one_negative)
	if !(b2d.ToString() == "129") {
		lang.SayString("BUG1183")
	}
	b2d, err = lang.RxFromString("111110000001").B2d(one_negative)
	if !(b2d.ToString() == "3969") {
		lang.SayString("BUG1184")
	}
	b2d, err = lang.RxFromString("1111111110000001").B2d(one_negative)
	if !(b2d.ToString() == "65409") {
		lang.SayString("BUG1185")
	}
	b2d, err = lang.RxFromString("1100011011110000").B2d(one_negative)
	if !(b2d.ToString() == "50928") {
		lang.SayString("BUG1186")
	}
	b2d, err = lang.RxFromString("10000001").B2d(eight)
	if !(b2d.ToString() == "-127") {
		lang.SayString("BUG1187")
	}
	b2d, err = lang.RxFromString("10000001").B2d(lang.RxFromInt32(16))
	if !(b2d.ToString() == "129") {
		lang.SayString("BUG1188")
	}
	b2d, err = lang.RxFromString("1111000010000001").B2d(lang.RxFromInt32(16))
	if !(b2d.ToString() == "-3967") {
		lang.SayString("BUG1189")
	}
	b2d, err = lang.RxFromString("1111000010000001").B2d(lang.RxFromInt32(12))
	if !(b2d.ToString() == "129") {
		lang.SayString("BUG1190")
	}
	b2d, err = lang.RxFromString("1111000010000001").B2d(eight)
	if !(b2d.ToString() == "-127") {
		lang.SayString("BUG1191")
	}
	b2d, err = lang.RxFromString("1111000010000001").B2d(four)
	if !(b2d.ToString() == "1") {
		lang.SayString("BUG1192")
	}
	b2d, err = lang.RxFromString("0000000000110001").B2d(zero)
	if !(b2d.ToString() == "0") {
		lang.SayString("BUG1193")
	}
	b2x, err := lang.RxFromString("11000011").B2x()
	if !(b2x.ToString() == "C3") {
		lang.SayString("BUG1194")
	}
	b2d, err = upper_a.B2d(lower_a)
	if err == nil {
		lang.SayString("BUG1195")
	}
	b2d, err = lang.RxFromString("123").B2d(lang.RxFromInt32(128))
	if err == nil {
		lang.SayString("BUG1196")
	}
	b2d, err = lang.RxFromString("11111111").B2d(lang.RxFromString("9999999999"))
	if err == nil {
		lang.SayString("BUG1197")
	}
	b2x, err = lang.RxFromString("10111").B2x()
	if !(b2x.ToString() == "17") {
		lang.SayString("BUG1198")
	}
	b2x, err = lang.RxFromString("0101").B2x()
	if !(b2x.ToString() == "5") {
		lang.SayString("BUG1199")
	}
	b2x, err = lang.RxFromString("101").B2x()
	if !(b2x.ToString() == "5") {
		lang.SayString("BUG1200")
	}
	b2x, err = lang.RxFromString("111110000").B2x()
	if !(b2x.ToString() == "1F0") {
		lang.SayString("BUG1201")
	}
	b2x, err = emptyrex.B2x()
	if err == nil {
		lang.SayString("BUG1202")
	}
	c2d, err := lang.RxFromRune('M').C2d()
	if !(c2d.ToString() == "77") {
		lang.SayString("BUG1203")
	}
	c2d, err = lang.RxFromRune('\r').C2d()
	if !(c2d.ToString() == "13") {
		lang.SayString("BUG1204")
	}
	c2d, err = lang.RxFromRune(0).C2d()
	if !(c2d.ToString() == "0") {
		lang.SayString("BUG1205")
	}
	c2d, err = lang.RxFromRune('\000').C2d()
	if !(c2d.ToString() == "0") {
		lang.SayString("BUG1206")
	}
	c2d, err = lang.RxFromRune('\x00').C2d()
	if !(c2d.ToString() == "0") {
		lang.SayString("BUG1207")
	}
	c2d, err = lang.RxFromRune('\u0000').C2d()
	if !(c2d.ToString() == "0") {
		lang.SayString("BUG1208")
	}
	c2d, err = lang.Rx([]rune{unicode.MaxRune}, true).C2d()
	if !(c2d.ToString() == "1114111") {
		lang.SayString("BUG1614")
	}
	c2d, err = emptyrex.C2d()
	if err == nil {
		lang.SayString("BUG1209")
	}
	c2x, err := lang.RxFromRune('M').C2x()
	if !(c2x.ToString() == "4D") {
		lang.SayString("BUG1210")
	}
	c2x, err = lang.RxFromRune('\r').C2x()
	if !(c2x.ToString() == "D") {
		lang.SayString("BUG1211")
	}
	c2x, err = lang.RxFromRune(0).C2x()
	if !(c2x.ToString() == "0") {
		lang.SayString("BUG1212")
	}
	c2x, err = lang.RxFromRune(15).C2x()
	if !(c2x.ToString() == "F") {
		lang.SayString("BUG1213")
	}
	c2x, err = lang.RxFromRune(255).C2x()
	if !(c2x.ToString() == "FF") {
		lang.SayString("BUG1214")
	}
	c2x, err = lang.RxFromRune(4095).C2x()
	if !(c2x.ToString() == "FFF") {
		lang.SayString("BUG1215")
	}
	c2x, err = lang.RxFromRune(65535).C2x()
	if !(c2x.ToString() == "FFFF") {
		lang.SayString("BUG1216")
	}
	c2x, err = lang.RxFromRune(-2147483648).C2x()
	if !(c2x.ToString() == "0000") {
		lang.SayString("BUG1217")
	}
	c2x, err = lang.RxFromRune(2147483647).C2x()
	if !(c2x.ToString() == "FFFF") {
		lang.SayString("BUG1218")
	}
	c2x, err = lang.RxFromRune(0x06F1).C2x()
	if !(c2x.ToString() == "6F1") {
		lang.SayString("BUG1219")
	}
	c2x, err = lang.RxFromRune(0x0024).C2x()
	if !(c2x.ToString() == "24") {
		lang.SayString("BUG1220")
	}
	c2x, err = lang.RxFromRune('0').C2x()
	if !(c2x.ToString() == "30") {
		lang.SayString("BUG1221")
	}
	c2x, err = emptyrex.C2x()
	if err == nil {
		lang.SayString("BUG1222")
	}
	center, err := lang.RxFromString("ABC").Center(eight, lang.RxFromString("-"))
	if !(center.ToString() == "--ABC---") {
		lang.SayString("BUG1223")
	}
	center, err = lang.RxFromString("The blue sky").Center(seven, blank)
	if !(center.ToString() == "e blue ") {
		lang.SayString("BUG1224")
	}
	center, err = emptyrex.Center(one_over_max, nil)
	if err == nil {
		lang.SayString("BUG1225")
	}
	centre, err := lang.RxFromString("ABC").Centre(seven, blank)
	if !(centre.ToString() == "  ABC  ") {
		lang.SayString("BUG1226")
	}
	centre, err = lang.RxFromString("The blue sky").Centre(eight, blank)
	if !(centre.ToString() == "e blue s") {
		lang.SayString("BUG1227")
	}
	centre, err = emptyrex.Centre(lang.RxFromInt64(-9223372036854775808), lang.RxFromInt64(-9223372036854775808))
	if err == nil {
		lang.SayString("BUG1228")
	}
	centre, err = emptyrex.Centre(emptyrex, emptyrex)
	if err == nil {
		lang.SayString("BUG1229")
	}
	centre, err = emptyrex.Centre(one, blank)
	if err != nil {
		lang.SayString("BUG1230")
	}
	center, err = emptyrex.Center(emptyrex, emptyrex)
	if err == nil {
		lang.SayString("BUG1231")
	}
	changestr := lang.RxFromString("elephant").ChangeStr(lang.RxFromString("e"), lang.RxFromString("X"))
	if !(changestr.ToString() == "XlXphant") {
		lang.SayString("BUG1232")
	}
	changestr = lang.RxFromString("elephant").ChangeStr(lang.RxFromString("ph"), lang.RxFromString("X"))
	if !(changestr.ToString() == "eleXant") {
		lang.SayString("BUG1233")
	}
	changestr = lang.RxFromString("elephant").ChangeStr(lang.RxFromString("ph"), lang.RxFromString("hph"))
	if !(changestr.ToString() == "elehphant") {
		lang.SayString("BUG1234")
	}
	changestr = lang.RxFromString("elephant").ChangeStr(lang.RxFromString("e"), emptystr)
	if !(changestr.ToString() == "lphant") {
		lang.SayString("BUG1235")
	}
	changestr = lang.RxFromString("elephant").ChangeStr(emptystr, lang.RxFromString("!!"))
	if !(changestr.ToString() == "elephant") {
		lang.SayString("BUG1236")
	}
	changestr = emptyrex.ChangeStr(emptyrex, emptyrex)
	if err == nil {
		lang.SayString("BUG1237")
	}

	charat, err := lang.RxFromString("ABC").CharAt(1)
	if !(charat == 'B') {
		lang.SayString("BUG1238")
	}
	charat, err = lang.RxFromString("ABC").CharAt(0)
	if !(charat == 'A') {
		lang.SayString("BUG1239")
	}
	charat, err = lang.RxFromString("ABC").CharAt(2)
	if !(charat == 'C') {
		lang.SayString("BUG1240")
	}
	charat, err = emptyrex.CharAt(0)
	if err == nil {
		lang.SayString("BUG1240")
	}
	charat, err = lang.RxFromString("ABC").CharAt(-1)
	if err == nil {
		lang.SayString("BUG1240")
	}
	charat, err = lang.RxFromString("ABC").CharAt(3)
	if err == nil {
		lang.SayString("BUG1240")
	}

	compare, err := lang.RxFromString("abc").Compare(lang.RxFromString("abc"), blank)
	if !(compare.ToString() == "0") {
		lang.SayString("BUG1241")
	}
	compare, err = lang.RxFromString("abc").Compare(lang.RxFromString("ak"), blank)
	if !(compare.ToString() == "2") {
		lang.SayString("BUG1242")
	}
	compare, err = lang.RxFromString("ab ").Compare(lang.RxFromString("ab"), blank)
	if !(compare.ToString() == "0") {
		lang.SayString("BUG1243")
	}
	compare, err = lang.RxFromString("ab ").Compare(lang.RxFromString("ab"), blank)
	if !(compare.ToString() == "0") {
		lang.SayString("BUG1244")
	}
	compare, err = lang.RxFromString("ab ").Compare(lang.RxFromString("ab"), lang.RxFromString("x"))
	if !(compare.ToString() == "3") {
		lang.SayString("BUG1245")
	}
	compare, err = lang.RxFromString("ab-- ").Compare(lang.RxFromString("ab"), lang.RxFromString("-"))
	if !(compare.ToString() == "5") {
		lang.SayString("BUG1246")
	}
	compare, err = emptyrex.Compare(emptyrex, emptyrex)
	if err == nil {
		lang.SayString("BUG1247")
	}
	compare, err = emptyrex.Compare(lower_a, blank)
	if err != nil {
		lang.SayString("BUG1248")
	}
	compare, err = lang.RxFromString("qq").Compare(emptyrex, blank)
	if err != nil {
		lang.SayString("BUG1249")
	}
	compareto, err := one.CompareTo("0")
	if !(compareto == 1) {
		lang.SayString("BUG1250")
	}
	compareto, err = one.CompareTo("1")
	if !(compareto == 0) {
		lang.SayString("BUG1251")
	}
	compareto, err = one.CompareTo("2")
	if !(compareto == -1) {
		lang.SayString("BUG1252")
	}
	compareto, err = one.CompareTo("a.+E")
	if !(compareto == -1) {
		lang.SayString("BUG2252")
	}
	compareto, err = lang.RxFromString("a").CompareTo("0")
	if !(compareto == 1) {
		lang.SayString("BUG2253")
	}
	compareto, err = lang.RxFromString("a").CompareTo("")
	if !(compareto == 1) {
		lang.SayString("BUG2254")
	}
	compareto, err = lang.RxFromEmpty().CompareTo("0")
	//if err == nil { //FIXME -NOT IN NETREXX NOW
	if !(compareto == 0) {
		//fmt.Println(compareto)
		lang.SayString("BUG2255")
	}
	copies, err := lang.RxFromString("abc").Copies(three)
	if !(copies.ToString() == "abcabcabc") {
		lang.SayString("BUG1253")
	}
	copies, err = lang.RxFromString("abc").Copies(zero)
	if !(copies.ToString() == "") {
		lang.SayString("BUG1254")
	}
	copies, err = emptystr.Copies(two)
	if !(copies.ToString() == "") {
		lang.SayString("BUG1255")
	}
	copies, err = emptyrex.Copies(one_negative)
	if err == nil {
		lang.SayString("BUG1256")
	}
	copies, err = emptyrex.Copies(one)
	if err != nil {
		lang.SayString("BUG1257")
	}
	foo := lang.ToRxFromString("def")
	foo.GetNode(lang.RxFromRune('a')).Leaf = one
	foo.GetNode(lang.RxFromRune('b')).Leaf = two
	if !(foo.IsIndexed().ToString() == "1") {
		lang.SayString("BUG1258")
	}
	if !(foo.TestNode(lang.RxFromRune('a')) == true) {
		lang.SayString("BUG1259")
	}
	if !(foo.TestNode(lang.RxFromString("IT_IS_NOT_HERE")) == false) {
		lang.SayString("BUG1260")
	}
	bar := lang.ToRxFromString("ghi")
	bar.GetNode(lang.RxFromRune('b')).Leaf = lang.RxFromRune('B')
	bar.GetNode(lang.RxFromRune('c')).Leaf = lang.RxFromRune('C')
	if !(bar.IsIndexed().ToString() == "1") {
		lang.SayString("BUG1261")
	}

	merged := foo.CopyIndexed(bar)

	if !(merged.IsIndexed().ToString() == "1") {
		lang.SayString("BUG1262")
	}
	if !(merged.GetNode(lang.RxFromRune('a')).Leaf.ToString() == "1") {
		lang.SayString("BUG1263")
	}
	if !(merged.GetNode(lang.RxFromRune('b')).Leaf.ToString() == "B") {
		lang.SayString("BUG1264")
	}
	if !(merged.GetNode(lang.RxFromRune('c')).Leaf.ToString() == "C") {
		lang.SayString("BUG1265")
	}
	// Does not contain 'd' should return foo
	if !(merged.GetNode(lang.RxFromRune('d')).Leaf.ToString() == "def") {
		lang.SayString("BUG1266")
	}
	if !(merged.Size(0) == 3) {
		lang.SayString("BUG3336")
	}

	//FIXME REXX RETURNS
	bar.CopyIndexed(foo)
	lang.RxFromEmpty().CopyIndexed(foo)
	too := lang.ToRxFromString("def")
	too.GetNode(lang.RxFromRune('a')).Leaf = one
	too.GetNode(lang.RxFromRune('a')).Leaf = nil
	lang.RxFromString("a").CopyIndexed(too)

	empty_stem := lang.ToRxFromString("stem")
	empty_merge := lang.ToRxFromString("merge")
	nomerge := empty_stem.CopyIndexed(empty_merge)
	if !(nomerge.IsIndexed().ToString() == "0") {
		lang.SayString("BUG1267")
	}
	countstr := lang.RxFromString("elephant").CountStr(lang.RxFromString("e"))
	if !(countstr.ToString() == "2") {
		lang.SayString("BUG1268")
	}
	countstr = lang.RxFromString("elephant").CountStr(lang.RxFromString("ph"))
	if !(countstr.ToString() == "1") {
		lang.SayString("BUG1269")
	}
	countstr = lang.RxFromString("elephant").CountStr(emptystr)
	if !(countstr.ToString() == "0") {
		lang.SayString("BUG1270")
	}
	countstr = emptyrex.CountStr(emptyrex)
	if err != nil {
		lang.SayString("BUG1271")
	}
	d2b, err := zero.D2b(lang.ToRxFromRunes([]rune("zip")))
	if !(d2b.ToString() == "0") {
		lang.SayString("BUG1272")
	}
	d2b, err = lang.RxFromString("9").D2b(lang.ToRxFromRunes([]rune("zip")))
	if !(d2b.ToString() == "1001") {
		lang.SayString("BUG1273")
	}
	d2b, err = lang.RxFromString("19").D2b(lang.ToRxFromRunes([]rune("zip")))
	if !(d2b.ToString() == "10011") {
		lang.SayString("BUG1274")
	}
	d2b, err = lang.RxFromString("129").D2b(lang.ToRxFromRunes([]rune("zip")))
	if !(d2b.ToString() == "10000001") {
		lang.SayString("BUG1275")
	}
	d2b, err = lang.RxFromString("129").D2b(one)
	if !(d2b.ToString() == "1") {
		lang.SayString("BUG1276")
	}
	d2b, err = lang.RxFromString("129").D2b(eight)
	if !(d2b.ToString() == "10000001") {
		lang.SayString("BUG1277")
	}
	d2b, err = lang.RxFromString("127").D2b(lang.RxFromInt32(12))
	if !(d2b.ToString() == "000001111111") {
		lang.SayString("BUG1278")
	}
	d2b, err = lang.RxFromString("129").D2b(lang.RxFromInt32(16))
	if !(d2b.ToString() == "0000000010000001") {
		lang.SayString("BUG1279")
	}
	d2b, err = lang.RxFromString("257").D2b(eight)
	if !(d2b.ToString() == "00000001") {
		lang.SayString("BUG1280")
	}
	d2b, err = lang.RxFromString("-127").D2b(eight)
	if !(d2b.ToString() == "10000001") {
		lang.SayString("BUG1281")
	}
	d2b, err = lang.RxFromString("-127").D2b(lang.RxFromInt32(16))
	if !(d2b.ToString() == "1111111110000001") {
		lang.SayString("BUG1282")
	}
	d2b, err = lang.RxFromString("12").D2b(zero)
	if !(d2b.ToString() == "") {
		lang.SayString("BUG1283")
	}
	d2b, err = emptyrex.D2b(emptyrex)
	if err == nil {
		lang.SayString("BUG1284")
	}
	d2b, err = lang.RxFromString("-1").D2b(lang.RxFromString("ax"))
	if err == nil {
		lang.SayString("BUG1285")
	}
	d2c, err := lang.RxFromString("77").D2c()
	if !(d2c.ToString() == "M") {
		lang.SayString("BUG1286")
	}
	d2c, err = lang.RxFromString("+77").D2c()
	if !(d2c.ToString() == "M") {
		lang.SayString("BUG1287")
	}
	d2c, err = zero.D2c() //FIXME NIL THIS IS HOW TO NULLS \0
	d2cnull, err := d2c.ToRune()
	if !(d2cnull == 0) {
		lang.SayString("BUG1288")
	}
	d2c, err = lang.RxFromString("-1").D2c()
	if err == nil {
		lang.SayString("BUG1289")
	}
	d2c, err = lang.RxFromString("AX").D2c()
	if err == nil {
		lang.SayString("BUG1290")
	}
	d2x, err := lang.RxFromString("9").D2x(one_negative)
	if !(d2x.ToString() == "9") {
		lang.SayString("BUG1291")
	}
	d2x, err = lang.RxFromString("129").D2x(one_negative)
	if !(d2x.ToString() == "81") {
		lang.SayString("BUG1292")
	}
	d2x, err = lang.RxFromString("129").D2x(one)
	if !(d2x.ToString() == "1") {
		lang.SayString("BUG1293")
	}
	d2x, err = lang.RxFromString("129").D2x(two)
	if !(d2x.ToString() == "81") {
		lang.SayString("BUG1294")
	}
	d2x, err = lang.RxFromString("127").D2x(three)
	if !(d2x.ToString() == "07F") {
		lang.SayString("BUG1295")
	}
	d2x, err = lang.RxFromString("129").D2x(four)
	if !(d2x.ToString() == "0081") {
		lang.SayString("BUG1296")
	}
	d2x, err = lang.RxFromString("257").D2x(two)
	if !(d2x.ToString() == "01") {
		lang.SayString("BUG1297")
	}
	d2x, err = lang.RxFromString("-127").D2x(two)
	if !(d2x.ToString() == "81") {
		lang.SayString("BUG1298")
	}
	d2x, err = lang.RxFromString("-127").D2x(four)
	if !(d2x.ToString() == "FF81") {
		lang.SayString("BUG1299")
	}
	d2x, err = lang.RxFromInt32(50928).D2x(one_negative)
	if !(d2x.ToString() == "C6F0") {
		lang.SayString("BUG1299")
	}
	d2x, err = lang.RxFromString("10").D2x(one_negative)
	if !(d2x.ToString() == "A") {
		lang.SayString("BUG1299")
	}
	d2x, err = lang.RxFromString("12").D2x(zero)
	d2xnull, err := d2x.ToRune()
	if !(d2xnull == 0) {
		lang.SayString("BUG1300")
	}
	d2x, err = emptyrex.D2x(emptyrex)
	if err != nil {
		lang.SayString("BUG1301")
	}
	d2x, err = lang.RxFromString("AX").D2x(emptyrex)
	if err == nil {
		lang.SayString("BUG1302")
	}
	d2x, err = lang.RxFromString("129").D2x(lang.RxFromFloat32(-1.9))
	if err == nil {
		lang.SayString("BUG1303")
	}
	d2x, err = lang.RxFromString("1.2345E+18").D2x(one_negative)
	if err == nil {
		lang.SayString("BUG&303")
	}
	datatype, err := lang.RxFromString("101").DataType(lang.RxFromRune('B'))
	if !(datatype.ToString() == "1") {
		lang.SayString("BUG1304")
	}
	datatype, err = lang.RxFromString("12.3").DataType(lang.RxFromRune('D'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1305")
	}
	datatype, err = lang.RxFromString("12.3").DataType(lang.RxFromRune('N'))
	if !(datatype.ToString() == "1") {
		lang.SayString("BUG1306")
	}
	datatype, err = lang.RxFromString("12.3").DataType(lang.RxFromRune('W'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1307")
	}
	datatype, err = lang.RxFromString("LaArca").DataType(lang.RxFromRune('M'))
	if !(datatype.ToString() == "1") {
		lang.SayString("BUG1308")
	}
	datatype, err = emptystr.DataType(lang.RxFromRune('M'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1309")
	}
	datatype, err = lang.RxFromString("Llanes").DataType(lang.RxFromRune('L'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1310")
	}
	datatype, err = lang.RxFromString("3 d").DataType(lang.RxFromRune('s'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1311")
	}
	datatype, err = lang.RxFromString("BCd3").DataType(lang.RxFromRune('X'))
	if !(datatype.ToString() == "1") {
		lang.SayString("BUG1312")
	}
	datatype, err = lang.RxFromString("BCgd3").DataType(lang.RxFromRune('X'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1313")
	}
	// tests for golang identifiers using datatype
	datatype, err = lang.RxFromString("@_is_start").DataType(lang.RxFromRune('S'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1314")
	}
	datatype, err = lang.RxFromString("Alpha_is_start").DataType(lang.RxFromRune('S'))
	if !(datatype.ToString() == "1") {
		lang.SayString("BUG1315")
	}
	datatype, err = lang.RxFromString("_is_start").DataType(lang.RxFromRune('S'))
	if !(datatype.ToString() == "1") {
		lang.SayString("BUG1316")
	}
	datatype, err = lang.RxFromString("8_is_start").DataType(lang.RxFromRune('S'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1317")
	}
	datatype, err = lang.RxFromString("A+_is_part").DataType(lang.RxFromRune('S'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1318")
	}
	datatype, err = lang.RxFromString("A@_is_part").DataType(lang.RxFromRune('S'))
	if !(datatype.ToString() == "0") {
		lang.SayString("BUG1319")
	}
	datatype, err = lang.RxFromString("A8_is_part").DataType(lang.RxFromRune('S'))
	if !(datatype.ToString() == "1") {
		lang.SayString("BUG1320")
	}
	datatype, err = emptyrex.DataType(emptyrex)
	if err == nil {
		lang.SayString("BUG1321")
	}
	//FIXME?
	datatype, err = lang.RxFromString(".........").DataType(lang.RxFromRune('A'))
	if err != nil {
		lang.SayString("BUG1322")
	}
	datatype, err = lang.RxFromRune('Z').DataType(lang.RxFromRune('A'))
	if err != nil {
		lang.SayString("BUG1323")
	}
	datatype, err = lang.RxFromRune('Z').DataType(lang.RxFromRune('B'))
	if err != nil {
		lang.SayString("BUG1324")
	}
	datatype, err = lang.RxFromRune('1').DataType(lang.RxFromRune('D'))
	if err != nil {
		lang.SayString("BUG1325")
	}
	datatype, err = lang.RxFromRune('a').DataType(lang.RxFromRune('L'))
	if err != nil {
		lang.SayString("BUG1326")
	}
	datatype, err = lang.RxFromRune('1').DataType(lang.RxFromRune('M'))
	if err != nil {
		lang.SayString("BUG1327")
	}
	datatype, err = lang.RxFromString("1A").DataType(lang.RxFromRune('N'))
	if err != nil {
		lang.SayString("BUG1328")
	}
	datatype, err = lang.RxFromString("1A").DataType(lang.RxFromRune('U'))
	if err != nil {
		lang.SayString("BUG1329")
	}
	datatype, err = upper_a.DataType(lang.RxFromRune('U'))
	if err != nil {
		lang.SayString("BUG1330")
	}
	datatype, err = upper_a.DataType(lang.RxFromRune('W'))
	if err != nil {
		lang.SayString("BUG1331")
	}
	datatype, err = zero.DataType(lang.RxFromRune('W'))
	if err != nil {
		lang.SayString("BUG1332")
	}
	datatype, err = lang.RxFromString("1.99999999999999999").DataType(lang.RxFromRune('W'))
	if err != nil {
		lang.SayString("BUG1333")
	}
	datatype, err = lang.RxFromString("ABC").DataType(lang.RxFromString("ABDLMNSUWX"))
	if err != nil {
		lang.SayString("BUG1334")
	}
	delstr, err := lang.RxFromString("abcd").DelStr(three, lang.RxFromString("abcd").Length())
	if !(delstr.ToString() == "ab") {
		lang.SayString("BUG1335")
	}
	delstr, err = lang.RxFromString("abcde").DelStr(three, two)
	if !(delstr.ToString() == "abe") {
		lang.SayString("BUG1336")
	}
	delstr, err = lang.RxFromString("abcde").DelStr(six, lang.RxFromString("abcde").Length())
	if !(delstr.ToString() == "abcde") {
		lang.SayString("BUG1337")
	}
	delstr, err = emptyrex.DelStr(emptyrex, emptyrex)
	if err == nil {
		lang.SayString("BUG1338")
	}
	delstr, err = lang.RxFromString("ab").DelStr(lower_a, one_over_max)
	if err == nil {
		lang.SayString("BUG1339")
	}
	delstr, err = lang.RxFromString("ab").DelStr(one, one_over_max)
	if err == nil {
		lang.SayString("BUG1340")
	}
	delstr, err = emptyrex.DelStr(one, zero)
	if err != nil {
		lang.SayString("BUG1341")
	}
	delword, err := lang.RxFromString("Now is the  time").DelWord(two, two)
	if !(delword.ToString() == "Now time") {
		lang.SayString("BUG1342")
	}
	delword, err = emptyrex.DelWord(emptyrex, emptyrex)
	if err == nil {
		lang.SayString("BUG1343")
	}
	delword, err = lang.RxFromString("ab").DelWord(one, one_over_max)
	if err == nil {
		lang.SayString("BUG1344")
	}
	wnum, err := lang.RxFromString("Now is the  time ").Words()
	delword, err = lang.RxFromString("Now is the  time ").DelWord(three, wnum)
	if !(delword.ToString() == "Now is ") {
		lang.SayString("BUG1345")
	}
	wnum, err = lang.RxFromString("Now  time").Words()
	delword, err = lang.RxFromString("Now  time").DelWord(three, wnum)
	if !(delword.ToString() == "Now  time") {
		lang.SayString("BUG1346")
	}
	delword, err = emptyrex.DelWord(three, wnum)
	if err != nil {
		lang.SayString("BUG1347")
	}
	vowel := zero
	vowel.GetNode(lang.RxFromRune('a')).Leaf = one
	// add b
	vowel.GetNode(lang.RxFromRune('b')).Leaf = one
	// remove b using nil
	vowel.GetNode(lang.RxFromRune('b')).Leaf = nil
	vowel.GetNode(lang.RxFromRune('d')).Leaf = lang.RxFromInt32(01)
	if !(vowel.Exists(lang.RxFromRune('a')).ToString() == "1") {
		lang.SayString("BUG1348")
	}
	if !(vowel.Exists(lang.RxFromRune('b')).ToString() == "0") {
		lang.SayString("BUG1349")
	}
	if !(vowel.Exists(lang.RxFromRune('c')).ToString() == "0") {
		lang.SayString("BUG1350")
	}
	vowel.Put(lang.RxFromString("TESTPUT"), lang.RxFromInt32(1000))
	if !vowel.ContainsKey(lang.RxFromString("TESTPUT")) {
		lang.SayString("BUG1351")
	}
	if !vowel.ContainsValue(lang.RxFromInt32(1000)) {
		lang.SayString("BUG1352")
	}
	if lang.RxFromEmpty().ContainsValue(lang.RxFromInt32(1)) {
		lang.SayString("BUG4352")
	}
	vowel.Remove(lang.RxFromRune('d'))

	vowel.Remove(nil)

	vowel.PutAll(nil)

	if vowel.ContainsKey(lang.RxFromRune('d')) {
		lang.SayString("BUG1353")
	}
	if vowel.Size(0) != 2 {
		lang.SayString("BUG1354")
	}
	if vowel.IsEmpty() {
		lang.SayString("BUG1355")
	}
	vowel.Clear()
	if !vowel.IsEmpty() {
		lang.SayString("BUG1356")
	}
	if vowel.Size(0) != 0 {
		lang.SayString("BUG1357")
	}
	format, err := lang.RxFromString(" - 12.73").Format(nil, nil, nil, nil, nil)
	if !(format.ToString() == "-12.73") {
		lang.SayString("BUG1358")
	}
	format, err = lang.RxFromString("0.000").Format(nil, nil, nil, nil, nil)
	if !(format.ToString() == "0") {
		lang.SayString("BUG1359")
	}
	format, err = lang.RxFromRune('3').Format(four, nil, nil, nil, nil)
	if !(format.ToString() == "   3") {
		lang.SayString("BUG1360")
	}
	format, err = lang.RxFromString("1.73").Format(four, zero, nil, nil, nil)
	if !(format.ToString() == "   2") {
		lang.SayString("BUG1361")
	}
	format, err = lang.RxFromString("1.73").Format(four, three, nil, nil, nil)
	if !(format.ToString() == "   1.730") {
		lang.SayString("BUG1362")
	}
	format, err = lang.RxFromString("-.76").Format(four, one, nil, nil, nil)
	if !(format.ToString() == "  -0.8") {
		lang.SayString("BUG1363")
	}
	format, err = lang.RxFromString("3.03").Format(four, nil, nil, nil, nil)
	if !(format.ToString() == "   3.03") {
		lang.SayString("BUG1364")
	}
	format, err = lang.RxFromString(" - 12.73").Format(nil, four, nil, nil, nil)
	if !(format.ToString() == "-12.7300") {
		lang.SayString("BUG1365")
	}
	format, err = lang.RxFromString("12345.73").Format(nil, nil, two, two, nil)
	if !(format.ToString() == "1.234573E+04") {
		lang.SayString("BUG1366")
	}
	format, err = lang.RxFromString("12345.73").Format(nil, three, nil, zero, nil)
	if !(format.ToString() == "1.235E+4") {
		lang.SayString("BUG1367")
	}
	format, err = lang.RxFromString("1.234573").Format(nil, three, nil, zero, nil)
	if !(format.ToString() == "1.235") {
		lang.SayString("BUG1368")
	}
	format, err = lang.RxFromString("123.45").Format(nil, three, two, zero, nil)
	if !(format.ToString() == "1.235E+02") {
		lang.SayString("BUG1369")
	}
	format, err = lang.RxFromString("1.2345").Format(nil, three, two, zero, nil)
	if !(format.ToString() == "1.235    ") {
		lang.SayString("BUG1371")
	}
	format, err = lang.RxFromString("12345.73").Format(nil, nil, three, six, nil)
	if !(format.ToString() == "12345.73     ") {
		lang.SayString("BUG1372")
	}
	format, err = lang.RxFromString("12345e+5").Format(nil, three, nil, nil, nil)
	if !(format.ToString() == "1234500000.000") {
		lang.SayString("BUG1373")
	}
	format, err = upper_a.Format(nil, nil, nil, nil, nil)
	if err == nil {
		lang.SayString("BUG1374")
	}
	format, err = one.Format(one_negative, nil, nil, nil, nil)
	if err == nil {
		lang.SayString("BUG1375")
	}
	format, err = one.Format(nil, one_negative, nil, nil, nil)
	if err == nil {
		lang.SayString("BUG1376")
	}
	format, err = one.Format(nil, nil, one_negative, nil, nil)
	if err == nil {
		lang.SayString("BUG1377")
	}
	format, err = one.Format(nil, nil, nil, one_negative, nil)
	if err == nil {
		lang.SayString("BUG1378")
	}
	format, err = lang.RxFromString("0.1E-510").Format(nil, nil, nil, zero, lang.RxFromRune('e'))
	if !(format.ToString() == "100E-513") {
		lang.SayString("BUG9000")
	}
	format, err = lang.RxFromString("0.1E-510").Format(nil, zero, two, three, lang.RxFromRune('e'))
	if err == nil {
		lang.SayString("BUG9001")
		fmt.Println(err)
	}
	format, err = lang.Rx([]rune("-12345.73"), true).Format(nil, three, nil, zero, nil)
	if !(format.ToString() == "-1.235E+4") {
		lang.SayString("BUG6001")
	}
	format, err = lang.Rx([]rune("-a.bc.G-f"), false).Format(nil, three, nil, zero, nil)
	if err == nil {
		lang.SayString("BUG6002")
		fmt.Println(err)
	}
	format, err = lang.Rx([]rune("-a.bc.G-f"), true).Format(two, nil, two, two, lang.RxFromRune('e'))
	if err == nil {
		lang.SayString("BUG6003")
		fmt.Println(err)
	}

	format, err = lang.RxFromString("+0000000000000000009.00000000100000000").Format(nil, nil, two, zero, lang.RxFromRune('e'))
	if !(format.ToString() == "9.00000000100000000    ") {
		lang.SayString("BUG9002")
	}
	if !(lang.RxFromString("abc").HashCode() == 2345) {
		lang.SayString("BUG1379")
	}
	if !(lang.RxFromString("abcdefghijklmnopqrstuvwxyz").HashCode() == 43396087) {
		lang.SayString("BUG1380")
	}
	if !(emptystr.HashCode() == 0) {
		lang.SayString("BUG1381")
	}
	if !(emptyrex.HashCode() == 0) {
		lang.SayString("BUG1382")
	}
	insert, err := lang.RxFromString("abc").Insert(lang.RxFromString("123"), zero, lang.RxFromString("123").Length(), blank)
	if !(insert.ToString() == "123abc") {
		lang.SayString("BUG1383")
	}
	insert, err = lang.RxFromString("abcdef").Insert(lang.RxFromRune(' '), three, lang.RxFromRune(' ').Length(), blank)
	if !(insert.ToString() == "abc def") {
		lang.SayString("BUG1384")
	}
	insert, err = lang.RxFromString("abc").Insert(lang.RxFromString("123"), five, six, blank)
	if !(insert.ToString() == "abc  123   ") {
		lang.SayString("BUG1385")
	}
	insert, err = lang.RxFromString("abc").Insert(lang.RxFromString("123"), five, six, lang.RxFromRune('+'))
	if !(insert.ToString() == "abc++123+++") {
		lang.SayString("BUG1386")
	}
	insert, err = lang.RxFromString("abc").Insert(lang.RxFromString("123"), zero, five, lang.RxFromRune('-'))
	if !(insert.ToString() == "123--abc") {
		lang.SayString("BUG1387")
	}
	insert, err = lang.RxFromString("abcdef").Insert(lang.RxFromRune(' '), one_over_max, lang.RxFromRune(' ').Length(), blank)
	if err == nil {
		lang.SayString("BUG1388")
	}
	insert, err = lang.RxFromString("abcdef").Insert(lang.RxFromRune(' '), one, one_over_max, blank)
	if err == nil {
		lang.SayString("BUG1389")
	}
	insert, err = lang.RxFromString("abcdef").Insert(lang.RxFromRune(' '), one, one, lang.RxFromString("ABC"))
	if err == nil {
		lang.SayString("BUG1390")
	}
	insert, err = lang.RxFromString("abcdef").Insert(emptyrex, one, one, upper_a)
	if err != nil {
		lang.SayString("BUG1391")
	}
	insert, err = emptyrex.Insert(lower_a, one, one, upper_a)
	if err != nil {
		lang.SayString("BUG1392")
	}
	start, err := lang.RxFromString("abc def ghi").Length().ToInt32()
	if start == 0 {
		start = 1
	}
	lastpos, err := lang.RxFromString("abc def ghi").LastPos(lang.RxFromRune(' '), lang.RxFromInt32(start))
	if !(lastpos.ToString() == "8") {
		lang.SayString("BUG1393")
	}
	lastpos, err = lang.RxFromString("abc def ghi").LastPos(lang.RxFromRune(' '), seven)
	if !(lastpos.ToString() == "4") {
		lang.SayString("BUG1394")
	}
	start, err = lang.RxFromString("abcdefghi").Length().ToInt32()
	if start == 0 {
		start = 1
	}
	lastpos, err = lang.RxFromString("abcdefghi").LastPos(lang.RxFromRune(' '), lang.RxFromInt32(start))
	if !(lastpos.ToString() == "0") {
		lang.SayString("BUG1395")
	}
	start, err = lang.RxFromString("abcdefghi").Length().ToInt32()
	if start == 0 {
		start = 1
	}
	lastpos, err = lang.RxFromString("abcdefghi").LastPos(lang.RxFromString("cd"), lang.RxFromInt32(start))
	if !(lastpos.ToString() == "3") {
		lang.SayString("BUG1396")
	}
	start, err = emptystr.Length().ToInt32()
	if start == 0 {
		start = 1
	}
	lastpos, err = emptystr.LastPos(lang.RxFromRune('?'), lang.RxFromInt32(start))
	if !(lastpos.ToString() == "0") {
		lang.SayString("BUG1397")
	}
	start, err = lang.RxFromString("abcdefghi").Length().ToInt32()
	if start == 0 {
		start = 1
	}
	lastpos, err = lang.RxFromString("abcdefghi").LastPos(emptystr, lang.RxFromInt32(start))
	if !(lastpos.ToString() == "0") {
		lang.SayString("BUG1398")
	}
	lastpos, err = lower_a.LastPos(lower_a, one_over_max)
	if err == nil {
		lang.SayString("BUG1399")
	}
	lastpos, err = emptyrex.LastPos(lower_a, one)
	if err != nil {
		lang.SayString("BUG1400")
	}
	lastpos, err = lower_a.LastPos(emptyrex, one)
	if err != nil {
		lang.SayString("BUG1401")
	}
	left, err := lang.RxFromString("abc d").Left(eight, blank)
	if !(left.ToString() == "abc d   ") {
		lang.SayString("BUG1402")
	}
	left, err = lang.RxFromString("abc d").Left(eight, lang.RxFromRune('.'))
	if !(left.ToString() == "abc d...") {
		lang.SayString("BUG1403")
	}
	left, err = lang.RxFromString("abc defg").Left(six, blank)
	if !(left.ToString() == "abc de") {
		lang.SayString("BUG1404")
	}
	left, err = lower_a.Left(one_over_max, one)
	if err == nil {
		lang.SayString("BUG1405")
	}
	length := lang.RxFromString("abcdefgh").Length()
	if !(length.ToString() == "8") {
		lang.SayString("BUG1406")
	}
	length = emptystr.Length()
	if !(length.ToString() == "0") {
		lang.SayString("BUG1407")
	}
	lower, err := lang.RxFromString("SumA").Lower(one, lang.RxFromString("SumA").Length())
	if !(lower.ToString() == "suma") {
		lang.SayString("BUG1408")
	}
	lower, err = lang.RxFromString("SumA").Lower(two, lang.RxFromString("SumA").Length())
	if !(lower.ToString() == "Suma") {
		lang.SayString("BUG1409")
	}
	lower, err = lang.RxFromString("SuMB").Lower(one, one)
	if !(lower.ToString() == "suMB") {
		lang.SayString("BUG1410")
	}
	lower, err = lang.RxFromString("SUMB").Lower(two, two)
	if !(lower.ToString() == "SumB") {
		lang.SayString("BUG1411")
	}
	lower, err = emptystr.Lower(one, emptystr.Length())
	if !(lower.ToString() == "") {
		lang.SayString("BUG1412")
	}
	lower, err = lang.RxFromString("SumA").Lower(one_over_max, lang.RxFromString("SumA").Length())
	if err == nil {
		lang.SayString("BUG1413")
	}
	lower, err = lang.RxFromString("SumA").Lower(one, one_over_max)
	if err == nil {
		lang.SayString("BUG1414")
	}
	lower, err = emptyrex.Lower(one, one)
	if err != nil {
		lang.SayString("BUG1415")
	}
	max, err := zero.Max(one)
	if !(max.ToString() == "1") {
		lang.SayString("BUG1416")
	}
	max, err = lang.RxFromString("-1").Max(one)
	if !(max.ToString() == "1") {
		lang.SayString("BUG1417")
	}
	oneminus, err := one.OpMinus(nil)
	max, err = lang.RxFromString("+1").Max(oneminus)
	if !(max.ToString() == "1") {
		lang.SayString("BUG1418")
	}
	max, err = lang.RxFromString("1.0").Max(lang.RxFromString("1.00"))
	if !(max.ToString() == "1.0") {
		lang.SayString("BUG1419")
	}
	max, err = lang.RxFromString("1.00").Max(lang.RxFromString("1.0"))
	if !(max.ToString() == "1.00") {
		lang.SayString("BUG1420")
	}
	max, err = lang.RxFromString("123456700000").Max(lang.RxFromString("1234567E+5"))
	if !(max.ToString() == "123456700000") {
		lang.SayString("BUG1421")
	}
	max, err = lang.RxFromString("1234567E+5").Max(lang.RxFromString("123456700000"))
	if !(max.ToString() == "1.234567E+11") {
		lang.SayString("BUG1422")
	}
	max, err = upper_a.Max(one)
	if err == nil {
		lang.SayString("BUG1423")
	}
	max, err = one.Max(upper_a)
	if err == nil {
		lang.SayString("BUG1424")
	}
	min, err := zero.Min(one)
	if !(min.ToString() == "0") {
		lang.SayString("BUG1425")
	}
	min, err = lang.RxFromString("-1").Min(one)
	if !(min.ToString() == "-1") {
		lang.SayString("BUG1426")
	}
	ominus, err := one.OpMinus(nil)
	min, err = lang.RxFromString("+1").Min(ominus)
	if !(min.ToString() == "-1") {
		lang.SayString("BUG1427")
	}
	min, err = lang.RxFromString("1.0").Min(lang.RxFromString("1.00"))
	if !(min.ToString() == "1.0") {
		lang.SayString("BUG1428")
	}
	min, err = lang.RxFromString("1.00").Min(lang.RxFromString("1.0"))
	if !(min.ToString() == "1.00") {
		lang.SayString("BUG1429")
	}
	min, err = lang.RxFromString("123456700000").Min(lang.RxFromString("1234567E+5"))
	if !(min.ToString() == "123456700000") {
		lang.SayString("BUG1430")
	}
	min, err = lang.RxFromString("1234567E+5").Min(lang.RxFromString("123456700000"))
	if !(min.ToString() == "1.234567E+11") {
		lang.SayString("BUG1431")
	}
	min, err = upper_a.Min(one)
	if err == nil {
		lang.SayString("BUG1432")
	}
	min, err = one.Min(upper_a)
	if err == nil {
		lang.SayString("BUG1433")
	}
	overlay, err := lang.RxFromString("abcdef").Overlay(lang.RxFromRune(' '), three, lang.RxFromRune(' ').Length(), blank)
	if !(overlay.ToString() == "ab def") {
		lang.SayString("BUG1434")
	}
	overlay, err = lang.RxFromString("abcdef").Overlay(lang.RxFromRune('.'), three, two, blank)
	if !(overlay.ToString() == "ab. ef") {
		lang.SayString("BUG1435")
	}
	overlay, err = lang.RxFromString("abcd").Overlay(lang.RxFromString("qq"), one, lang.RxFromString("qq").Length(), blank)
	if !(overlay.ToString() == "qqcd") {
		lang.SayString("BUG1436")
	}
	overlay, err = lang.RxFromString("abcd").Overlay(lang.RxFromString("qq"), four, lang.RxFromString("qq").Length(), blank)
	if !(overlay.ToString() == "abcqq") {
		lang.SayString("BUG1437")
	}
	overlay, err = lang.RxFromString("abc").Overlay(lang.RxFromString("123"), five, six, lang.RxFromRune('+'))
	if !(overlay.ToString() == "abc+123+++") {
		lang.SayString("BUG1438")
	}
	overlay, err = lang.RxFromString("abcdef").Overlay(emptyrex, three, two, blank)
	if err != nil {
		lang.SayString("BUG1439")
	}
	overlay, err = lang.RxFromString("abcdef").Overlay(lang.RxFromRune('.'), one_over_max, two, blank)
	if err == nil {
		lang.SayString("BUG1440")
	}
	overlay, err = lang.RxFromString("abcdef").Overlay(lang.RxFromRune('.'), three, one_over_max, blank)
	if err == nil {
		lang.SayString("BUG1441")
	}
	overlay, err = lang.RxFromString("abcdef").Overlay(lang.RxFromRune('.'), three, two, lang.RxFromString("AX"))
	if err == nil {
		lang.SayString("BUG1442")
	}
	overlay, err = emptyrex.Overlay(lang.RxFromRune('.'), three, two, blank)
	if err != nil {
		lang.SayString("BUG1443")
	}
	pos, err := lang.RxFromString("Saturday").Pos(lang.RxFromString("day"), one)
	if !(pos.ToString() == "6") {
		lang.SayString("BUG1444")
	}
	pos, err = lang.RxFromString("abc def ghi").Pos(lang.RxFromRune('x'), one)
	if !(pos.ToString() == "0") {
		lang.SayString("BUG1445")
	}
	pos, err = lang.RxFromString("abc def ghi").Pos(lang.RxFromRune(' '), one)
	if !(pos.ToString() == "4") {
		lang.SayString("BUG1446")
	}
	pos, err = lang.RxFromString("abc def ghi").Pos(lang.RxFromRune(' '), five)
	if !(pos.ToString() == "8") {
		lang.SayString("BUG1447")
	}
	pos, err = lang.RxFromString("Saturday").Pos(emptystr, one)
	if !(pos.ToString() == "0") {
		lang.SayString("BUG1448")
	}
	pos, err = emptyrex.Pos(emptyrex, one)
	if err != nil {
		lang.SayString("BUG1449")
	}
	pos, err = lower_a.Pos(emptyrex, one)
	if err != nil {
		lang.SayString("BUG1450")
	}
	pos, err = lower_a.Pos(one, emptyrex)
	if err == nil {
		lang.SayString("BUG1451")
	}
	reverse, err := lang.RxFromString("ABc.").Reverse()
	if !(reverse.ToString() == ".cBA") {
		lang.SayString("BUG1452")
	}
	reverse, err = lang.RxFromString("XYZ ").Reverse()
	if !(reverse.ToString() == " ZYX") {
		lang.SayString("BUG1453")
	}
	reverse, err = lang.RxFromString("Tranquility").Reverse()
	if !(reverse.ToString() == "ytiliuqnarT") {
		lang.SayString("BUG1454")
	}
	reverse, err = emptystr.Reverse()
	if !(reverse.ToString() == "") {
		lang.SayString("BUG1455")
	}
	reverse, err = emptyrex.Reverse()
	if err != nil {
		lang.SayString("BUG1456")
	}
	right, err := lang.RxFromString("abc  d").Right(eight, blank)
	if !(right.ToString() == "  abc  d") {
		lang.SayString("BUG1457")
	}
	right, err = lang.RxFromString("abc def").Right(five, blank)
	if !(right.ToString() == "c def") {
		lang.SayString("BUG1458")
	}
	right, err = lang.RxFromString("12").Right(five, lang.RxFromRune('0'))
	if !(right.ToString() == "00012") {
		lang.SayString("BUG1459")
	}
	right, err = lang.RxFromString("abc def").Right(one_over_max, blank)
	if err == nil {
		lang.SayString("BUG1460")
	}
	right, err = emptyrex.Right(five, blank)
	if err != nil {
		lang.SayString("BUG1461")
	}
	right, err = lang.RxFromString("abc def").Right(eight, lang.RxFromString("jjj"))
	if err == nil {
		lang.SayString("BUG1462")
	}
	right, err = lang.RxFromString("abc def").Right(zero, blank)
	if err != nil {
		lang.SayString("BUG1463")
	}
	right, err = lang.RxFromString("abc def").Right(lang.RxFromString("-6.7"), blank)
	if err == nil {
		lang.SayString("BUG1464")
	}
	sequence, err := lang.RxFromRune('a').Sequence(lang.RxFromRune('f'))
	if !(sequence.ToString() == "abcdef") {
		lang.SayString("BUG1465")
	}
	sequence, err = lang.RxFromRune(rune(0)).Sequence(lang.RxFromRune('\x03'))
	if !(sequence.ToString() == "\x00\x01\x02\x03") {
		lang.SayString("BUG1466")
	}
	ub := []rune{0xfffe, 0xffff}
	sequence, err = lang.RxFromRune(rune(0xfffe)).Sequence(lang.RxFromRune(0xffff))
	if !(sequence.ToString() == string(ub)) {
		lang.SayString("BUG1467")
	}
	xb := lang.RxFromRunes(ub)
	if !(lang.ToRunesFromRexx(xb)[0] == ub[0]) {
		lang.SayString("BUG1468")
	}
	if !(lang.ToRunesFromRexx(xb)[1] == ub[1]) {
		lang.SayString("BUG1469")
	}
	sequence, err = emptyrex.Sequence(lang.RxFromRune('f'))
	if err == nil {
		lang.SayString("BUG1470")
	}
	sequence, err = lang.RxFromRune('a').Sequence(emptyrex)
	if err == nil {
		lang.SayString("BUG1471")
	}
	sequence, err = lang.RxFromRune('q').Sequence(lang.RxFromRune('k'))
	if err == nil {
		lang.SayString("BUG1472")
	}
	new_set := lang.RxSet()
	get_form := new_set.Formword()
	if !(get_form.ToString() == "scientific") {
		lang.SayString("BUG1473")
	}
	cloned_set := lang.RxSetFromOld(new_set)
	if !(cloned_set.Digits == new_set.Digits) {
		lang.SayString("BUG1474")
	}
	if !(cloned_set.Form == new_set.Form) {
		lang.SayString("BUG1475")
	}
	digit_set := lang.RxSetWithDigit(50)
	if !(digit_set.Digits == 50) {
		lang.SayString("BUG1476")
	}
	digit_form_set := lang.RxSetWithDigitandForm(2, int8(1))
	if !(digit_form_set.Digits == 2) {
		lang.SayString("BUG1477")
	}
	if !(digit_form_set.Form == 1) {
		lang.SayString("BUG1478")
	}
	new_set.SetForm(lang.RxFromString("engineering"))
	get_new_form := new_set.Formword()
	if !(get_new_form.ToString() == "engineering") {
		lang.SayString("BUG1479")
	}
	sign, err := lang.RxFromString("12.3").Sign()
	if !(sign.ToString() == "1") {
		lang.SayString("BUG1480")
	}
	sign, err = lang.RxFromString("0.0").Sign()
	if !(sign.ToString() == "0") {
		lang.SayString("BUG1481")
	}
	sign, err = lang.RxFromString(" -0.307").Sign()
	if !(sign.ToString() == "-1") {
		lang.SayString("BUG1482")
	}
	sign, err = lang.RxFromRune('q').Sign()
	if err == nil {
		lang.SayString("BUG1483")
	}
	significance := lang.RxFromRune('a').Significance()
	if !(significance == 0) {
		lang.SayString("BUG1484")
	}
	space, err := lang.RxFromString("abc  def  ").Space(one, blank)
	if !(space.ToString() == "abc def") {
		lang.SayString("BUG1485")
	}
	space, err = lang.RxFromString("  abc def ").Space(three, blank)
	if !(space.ToString() == "abc   def") {
		lang.SayString("BUG1486")
	}
	space, err = lang.RxFromString("abc  def  ").Space(one, blank)
	if !(space.ToString() == "abc def") {
		lang.SayString("BUG1487")
	}
	space, err = lang.RxFromString("abc  def  ").Space(zero, blank)
	if !(space.ToString() == "abcdef") {
		lang.SayString("BUG1488")
	}
	space, err = lang.RxFromString("abc  def  ").Space(two, lang.RxFromRune('+'))
	if !(space.ToString() == "abc++def") {
		lang.SayString("BUG1489")
	}
	space, err = lang.RxFromRune('a').Space(emptyrex, emptyrex)
	if err == nil {
		lang.SayString("BUG1490")
	}
	space, err = emptyrex.Space(one, blank)
	if err != nil {
		lang.SayString("BUG1491")
	}
	space, err = lang.RxFromRune('a').Space(one_over_max, emptyrex)
	if err == nil {
		lang.SayString("BUG1492")
	}
	strip, err := lang.RxFromString("  ab c  ").Strip(lang.RxFromString("B"), blank)
	if !(strip.ToString() == "ab c") {
		lang.SayString("BUG1493")
	}
	strip, err = lang.RxFromString("  ab c  ").Strip(lang.RxFromRune('L'), blank)
	if !(strip.ToString() == "ab c  ") {
		lang.SayString("BUG1494")
	}
	strip, err = lang.RxFromString("  ab c  ").Strip(lang.RxFromRune('t'), blank)
	if !(strip.ToString() == "  ab c") {
		lang.SayString("BUG1495")
	}
	strip, err = lang.RxFromString("12.70000").Strip(lang.RxFromRune('t'), zero)
	if !(strip.ToString() == "12.7") {
		lang.SayString("BUG1496")
	}
	strip, err = lang.RxFromString("0012.700").Strip(lang.RxFromRune('b'), zero)
	if !(strip.ToString() == "12.7") {
		lang.SayString("BUG1497")
	}
	strip, err = emptystr.Strip(lang.RxFromString("B"), blank)
	if !(strip.ToString() == "") {
		lang.SayString("BUG1498")
	}
	strip, err = lang.RxFromString("0000").Strip(lang.RxFromRune('T'), zero)
	if !(strip.ToString() == "") {
		lang.SayString("BUG1499")
	}
	strip, err = emptyrex.Strip(one, blank)
	if err == nil {
		lang.SayString("BUG1500")
	}
	strip, err = lang.RxFromRune('a').Strip(lang.RxFromString("B"), lang.RxFromString("AXE"))
	if err == nil {
		lang.SayString("BUG1501")
	}
	strip, err = emptyrex.Strip(lang.RxFromString("B"), blank)
	if err != nil {
		lang.SayString("BUG1502")
	}
	// FIX INT 32?
	var numone int32 = 2
	this := lang.RxFromString("abc")
	numtwo, err := this.Length().ToInt32()
	rx01 := lang.RxFromInt32(numtwo + 1 - numone)
	rx02, err := rx01.Max(lang.RxFromInt8(int8(0)))
	substr, err := this.Substr(lang.RxFromInt32(numone), rx02, blank)
	if !(substr.ToString() == "bc") {
		lang.SayString("BUG1503")
	}
	substr, err = lang.RxFromString("abc").Substr(two, four, blank)
	if !(substr.ToString() == "bc  ") {
		lang.SayString("BUG1504")
	}
	substr, err = lang.RxFromString("abc").Substr(five, four, blank)
	if !(substr.ToString() == "    ") {
		lang.SayString("BUG1505")
	}
	substr, err = lang.RxFromString("abc").Substr(two, six, lang.RxFromRune('.'))
	if !(substr.ToString() == "bc....") {
		lang.SayString("BUG1506")
	}
	substr, err = lang.RxFromString("abc").Substr(five, six, lang.RxFromRune('.'))
	if !(substr.ToString() == "......") {
		lang.SayString("BUG1507")
	}
	substr, err = lang.RxFromString("abc").Substr(one_over_max, four, blank)
	if err == nil {
		lang.SayString("BUG1508")
	}
	substr, err = lang.RxFromString("abc").Substr(two, one_over_max, blank)
	if err == nil {
		lang.SayString("BUG1509")
	}
	substr, err = lang.RxFromString("abc").Substr(two, four, lang.RxFromString("AXE"))
	if err == nil {
		lang.SayString("BUG1510")
	}
	substr, err = emptyrex.Substr(one, four, blank)
	if err != nil {
		lang.SayString("BUG1511")
	}
	subword, err := lang.RxFromString("Now is the  time").SubWord(two, two)
	if !(subword.ToString() == "is the") {
		lang.SayString("BUG1512")
	}
	subword, err = lang.RxFromString("Now is the  time").SubWord(three, lang.RxFromString("Now is the  time").Length())
	if !(subword.ToString() == "the  time") {
		lang.SayString("BUG1513")
	}
	subword, err = lang.RxFromString("Now is the  time").SubWord(five, lang.RxFromString("Now is the  time").Length())
	if !(subword.ToString() == "") {
		lang.SayString("BUG1514")
	}
	subword, err = lang.RxFromString("Now is the  time").SubWord(five, lang.RxFromString("Now is the  time").Length())
	if !(subword.ToString() == "") {
		lang.SayString("BUG1515")
	}
	subword, err = lang.RxFromString("abc").SubWord(one_over_max, blank)
	if err == nil {
		lang.SayString("BUG1516")
	}
	subword, err = lang.RxFromString("abc").SubWord(two, lang.RxFromString("AXE"))
	if err == nil {
		lang.SayString("BUG1517")
	}
	subword, err = emptyrex.SubWord(two, four)
	if err != nil {
		lang.SayString("BUG1518")
	}
	translate, err := lang.RxFromString("abbc").Translate(lang.RxFromRune('&'), lang.RxFromRune('b'), blank)
	if !(translate.ToString() == "a&&c") {
		lang.SayString("BUG1519")
	}
	translate, err = lang.RxFromString("abcdef").Translate(lang.RxFromString("12"), lang.RxFromString("ec"), blank)
	if !(translate.ToString() == "ab2d1f") {
		lang.SayString("BUG1520")
	}
	translate, err = lang.RxFromString("abcdef").Translate(lang.RxFromString("12"), lang.RxFromString("abcd"), lang.RxFromRune('.'))
	if !(translate.ToString() == "12..ef") {
		lang.SayString("BUG1521")
	}
	translate, err = lang.RxFromString("4123").Translate(lang.RxFromString("abcd"), lang.RxFromString("1234"), blank)
	if !(translate.ToString() == "dabc") {
		lang.SayString("BUG1522")
	}
	translate, err = lang.RxFromString("4123").Translate(lang.RxFromString("hods"), lang.RxFromString("1234"), blank)
	if !(translate.ToString() == "shod") {
		lang.SayString("BUG1523")
	}
	translate, err = lang.RxFromString("4123").Translate(lang.RxFromString("abcd"), lang.RxFromString("1234"), blank)
	if !(translate.ToString() == "dabc") {
		lang.SayString("BUG1524")
	}
	translate, err = lang.RxFromString("4312").Translate(lang.RxFromString("hods"), lang.RxFromString("1234"), blank)
	if !(translate.ToString() == "sdho") {
		lang.SayString("BUG1525")
	}
	translate, err = lang.RxFromString("abcdef").Translate(lang.RxFromString("12"), lang.RxFromString("ec"), lang.RxFromString("AXE"))
	if err == nil {
		lang.SayString("BUG1526")
	}
	translate, err = lang.RxFromString("abcdef").Translate(emptyrex, lang.RxFromString("ec"), blank)
	if err != nil {
		lang.SayString("BUG1527")
	}
	translate, err = lang.RxFromString("abcdef").Translate(lang.RxFromString("12"), emptyrex, blank)
	if err != nil {
		lang.SayString("BUG1528")
	}
	translate, err = emptyrex.Translate(lang.RxFromString("12"), lang.RxFromString("ec"), blank)
	if err != nil {
		lang.SayString("BUG1529")
	}
	trunc, err := lang.RxFromString("12.3").Trunc(zero)
	if !(trunc.ToString() == "12") {
		lang.SayString("BUG1530")
	}
	trunc, err = lang.RxFromString("127.09782").Trunc(three)
	if !(trunc.ToString() == "127.097") {
		lang.SayString("BUG1531")
	}
	trunc, err = lang.RxFromString("127.1").Trunc(three)
	if !(trunc.ToString() == "127.100") {
		lang.SayString("BUG1532")
	}
	trunc, err = lang.RxFromString("127").Trunc(two)
	if !(trunc.ToString() == "127.00") {
		lang.SayString("BUG1533")
	}
	trunc, err = zero.Trunc(two)
	if !(trunc.ToString() == "0.00") {
		lang.SayString("BUG1534")
	}
	trunc, err = emptyrex.Trunc(one)
	if err != nil {
		lang.SayString("BUG1535")
	}
	trunc, err = lang.RxFromString("A.3").Trunc(lang.RxFromInt32(-11))
	if err == nil {
		lang.SayString("BUG1536")
	}
	trunc, err = lang.RxFromString("12.3").Trunc(lang.RxFromFloat32(1.1))
	if err == nil {
		lang.SayString("BUG1537")
	}
	upper, err := lang.RxFromString("Fou-Baa").Upper(one, lang.RxFromString("Fou-Baa").Length())
	if !(upper.ToString() == "FOU-BAA") {
		lang.SayString("BUG1538")
	}
	upper, err = lang.RxFromString("Mad Sheep").Upper(one, lang.RxFromString("Mad Sheep").Length())
	if !(upper.ToString() == "MAD SHEEP") {
		lang.SayString("BUG1539")
	}
	upper, err = lang.RxFromString("Mad sheep").Upper(five, lang.RxFromString("Mad sheep").Length())
	if !(upper.ToString() == "Mad SHEEP") {
		lang.SayString("BUG1540")
	}
	upper, err = lang.RxFromString("Mad sheep").Upper(five, one)
	if !(upper.ToString() == "Mad Sheep") {
		lang.SayString("BUG1541")
	}
	upper, err = lang.RxFromString("Mad sheep").Upper(five, four)
	if !(upper.ToString() == "Mad SHEEp") {
		lang.SayString("BUG1542")
	}
	upper, err = lang.RxFromString("tinganon").Upper(one, one)
	if !(upper.ToString() == "Tinganon") {
		lang.SayString("BUG1543")
	}
	upper, err = emptystr.Upper(one, emptystr.Length())
	if !(upper.ToString() == "") {
		lang.SayString("BUG1544")
	}
	upper, err = lang.RxFromString("Mad Sheep").Upper(zero, lang.RxFromString("Mad Sheep").Length())
	if err == nil {
		lang.SayString("BUG1545")
	}
	upper, err = lang.RxFromString("Mad Sheep").Upper(one, one_over_max)
	if err == nil {
		lang.SayString("BUG1546")
	}
	upper, err = emptyrex.Upper(one, lang.RxFromString("Mad Sheep").Length())
	if err != nil {
		lang.SayString("BUG1547")
	}
	verify, err := lang.RxFromString("123").Verify(lang.RxFromString("1234567890"), lang.RxFromString("N"), one)
	if !(verify.ToString() == "0") {
		lang.SayString("BUG1548")
	}
	verify, err = lang.RxFromString("1Z3").Verify(lang.RxFromString("1234567890"), lang.RxFromString("N"), one)
	if !(verify.ToString() == "2") {
		lang.SayString("BUG1549")
	}
	verify, err = lang.RxFromString("AB4T").Verify(lang.RxFromString("1234567890"), lang.RxFromRune('M'), one)
	if !(verify.ToString() == "3") {
		lang.SayString("BUG1550")
	}
	verify, err = lang.RxFromString("1P3Q4").Verify(lang.RxFromString("1234567890"), lang.RxFromRune('N'), three)
	if !(verify.ToString() == "4") {
		lang.SayString("BUG1551")
	}
	verify, err = lang.RxFromString("ABCDE").Verify(emptystr, lang.RxFromRune('n'), three)
	if !(verify.ToString() == "3") {
		lang.SayString("BUG1552")
	}
	verify, err = lang.RxFromString("AB3CD5").Verify(lang.RxFromString("1234567890"), lang.RxFromRune('m'), four)
	if !(verify.ToString() == "6") {
		lang.SayString("BUG1553")
	}
	verify, err = lang.RxFromString("1Z3").Verify(lang.RxFromString("1234567890"), emptystr, one)
	if err == nil {
		lang.SayString("BUG1554")
	}
	verify, err = lang.RxFromString("1Z3").Verify(lang.RxFromString("1234567890"), lang.RxFromString("N"), one_negative)
	if err == nil {
		lang.SayString("BUG1555")
	}
	verify, err = emptyrex.Verify(lang.RxFromString("1234567890"), lang.RxFromString("N"), one)
	if err != nil {
		lang.SayString("BUG1556")
	}
	verify, err = lang.RxFromString("1Z3").Verify(emptyrex, lang.RxFromString("N"), one)
	if err != nil {
		lang.SayString("BUG1557")
	}
	word, err := lang.RxFromString("Now is the time").Word(three)
	if !(word.ToString() == "the") {
		lang.SayString("BUG1558")
	}
	word, err = lang.RxFromString("Now is the time").Word(five)
	if !(word.ToString() == "") {
		lang.SayString("BUG1559")
	}
	word, err = lang.RxFromString("Now is the time").Word(five)
	if !(word.ToString() == "") {
		lang.SayString("BUG1560")
	}
	word, err = lang.RxFromString("1Z3").Word(emptyrex)
	if err == nil {
		lang.SayString("BUG1561")
	}
	wordindex, err := lang.RxFromString("Now is the time").WordIndex(three)
	if !(wordindex.ToString() == "8") {
		lang.SayString("BUG1562")
	}
	wordindex, err = lang.RxFromString("Now is the time").WordIndex(six)
	if !(wordindex.ToString() == "0") {
		lang.SayString("BUG1563")
	}
	wordindex, err = lang.RxFromString("1Z3").WordIndex(emptyrex)
	if err == nil {
		lang.SayString("BUG1564")
	}
	wordindex, err = emptyrex.WordIndex(one)
	if err != nil {
		lang.SayString("BUG1565")
	}
	wordlength, err := lang.RxFromString("Now is the time").WordLength(two)
	if !(wordlength.ToString() == "2") {
		lang.SayString("BUG1566")
	}
	wordlength, err = lang.RxFromString("Now comes the time").WordLength(two)
	if !(wordlength.ToString() == "5") {
		lang.SayString("BUG1567")
	}
	wordlength, err = lang.RxFromString("Now is the time").WordLength(six)
	if !(wordlength.ToString() == "0") {
		lang.SayString("BUG1568")
	}
	wordlength, err = lang.RxFromString("1Z3").WordLength(emptyrex)
	if err == nil {
		lang.SayString("BUG1569")
	}
	wordlength, err = emptyrex.WordLength(one)
	if err != nil {
		lang.SayString("BUG1570")
	}
	wordpos, err := lang.RxFromString("Now is the time").WordPos(lang.RxFromString("the"), one)
	if !(wordpos.ToString() == "3") {
		lang.SayString("BUG1571")
	}
	wordpos, err = lang.RxFromString("Now is the time").WordPos(lang.RxFromString("The"), one)
	if !(wordpos.ToString() == "0") {
		lang.SayString("BUG1572")
	}
	wordpos, err = lang.RxFromString("Now is the time").WordPos(lang.RxFromString("is the"), one)
	if !(wordpos.ToString() == "2") {
		lang.SayString("BUG1573")
	}
	wordpos, err = lang.RxFromString("Now is the time").WordPos(lang.RxFromString("is   the"), one)
	if !(wordpos.ToString() == "2") {
		lang.SayString("BUG1574")
	}
	wordpos, err = lang.RxFromString("Now is the time").WordPos(lang.RxFromString("is time"), one)
	if !(wordpos.ToString() == "0") {
		lang.SayString("BUG1575")
	}
	wordpos, err = lang.RxFromString("To be or not to be").WordPos(lang.RxFromString("be"), one)
	if !(wordpos.ToString() == "2") {
		lang.SayString("BUG1576")
	}
	wordpos, err = lang.RxFromString("To be or not to be").WordPos(lang.RxFromString("be"), three)
	if !(wordpos.ToString() == "6") {
		lang.SayString("BUG1577")
	}
	wordpos, err = lang.RxFromString("Now is the time").WordPos(emptyrex, one)
	if err != nil {
		lang.SayString("BUG1578")
	}
	wordpos, err = emptyrex.WordPos(lang.RxFromString("The"), one)
	if err != nil {
		lang.SayString("BUG1579")
	}
	wordpos, err = lang.RxFromString("Now is the time").WordPos(lang.RxFromString("The"), emptyrex)
	if err == nil {
		lang.SayString("BUG1580")
	}
	words, err := lang.RxFromString("Now is the time").Words()
	if !(words.ToString() == "4") {
		lang.SayString("BUG1581")
	}
	words, err = blank.Words()
	if !(words.ToString() == "0") {
		lang.SayString("BUG1582")
	}
	words, err = emptystr.Words()
	if !(words.ToString() == "0") {
		lang.SayString("BUG1583")
	}
	x2b, err := lang.RxFromString("C3").X2b()
	if !(x2b.ToString() == "11000011") {
		lang.SayString("BUG1584")
	}
	x2b, err = lang.RxFromRune('7').X2b()
	if !(x2b.ToString() == "0111") {
		lang.SayString("BUG1585")
	}
	x2b, err = lang.RxFromString("1C1").X2b()
	if !(x2b.ToString() == "000111000001") {
		lang.SayString("BUG1586")
	}
	x2b, err = emptyrex.X2b()
	if err == nil {
		lang.SayString("BUG1587")
	}
	x2b, err = emptystr.X2b()
	if err == nil {
		lang.SayString("BUG1588")
	}
	x2b, err = lang.RxFromString("J").X2b()
	if err == nil {
		lang.SayString("BUG1589")
	}
	x2c, err := lang.RxFromString("004D").X2c()
	if !(x2c.ToString() == "M") {
		lang.SayString("BUG1590")
	}
	x2c, err = lang.RxFromString("4d").X2c()
	if !(x2c.ToString() == "M") {
		lang.SayString("BUG1591")
	}
	x2c, err = lang.RxFromRune('0').X2c()
	x2cnull, err := x2c.ToRune()
	if !(x2cnull == 0) {
		lang.SayString("BUG1592")
	}
	x2c, err = emptyrex.X2c()
	if err == nil {
		lang.SayString("BUG1593")
	}
	x2c, err = emptystr.X2c()
	if err == nil {
		lang.SayString("BUG1594")
	}
	x2c, err = lang.RxFromString("J").X2c()
	if err == nil {
		lang.SayString("BUG1595")
	}
	x2d, err := lang.RxFromString("0E").X2d(one_negative)
	if !(x2d.ToString() == "14") {
		lang.SayString("BUG1596")
	}
	x2d, err = lang.RxFromString("81").X2d(one_negative)
	if !(x2d.ToString() == "129") {
		lang.SayString("BUG1597")
	}
	x2d, err = lang.RxFromString("F81").X2d(one_negative)
	if !(x2d.ToString() == "3969") {
		lang.SayString("BUG1598")
	}
	x2d, err = lang.RxFromString("FF81").X2d(one_negative)
	if !(x2d.ToString() == "65409") {
		lang.SayString("BUG1599")
	}
	x2d, err = lang.RxFromString("c6f0").X2d(one_negative)
	if !(x2d.ToString() == "50928") {
		lang.SayString("BUG1600")
	}
	x2d, err = lang.RxFromString("81").X2d(two)
	if !(x2d.ToString() == "-127") {
		lang.SayString("BUG1601")
	}
	x2d, err = lang.RxFromString("81").X2d(four)
	if !(x2d.ToString() == "129") {
		lang.SayString("BUG1602")
	}
	x2d, err = lang.RxFromString("F081").X2d(four)
	if !(x2d.ToString() == "-3967") {
		lang.SayString("BUG1603")
	}
	x2d, err = lang.RxFromString("F081").X2d(three)
	if !(x2d.ToString() == "129") {
		lang.SayString("BUG1604")
	}
	x2d, err = lang.RxFromString("F081").X2d(two)
	if !(x2d.ToString() == "-127") {
		lang.SayString("BUG1605")
	}
	x2d, err = lang.RxFromString("F081").X2d(one)
	if !(x2d.ToString() == "1") {
		lang.SayString("BUG1606")
	}
	x2d, err = lang.RxFromString("0031").X2d(zero)
	if !(x2d.ToString() == "0") {
		lang.SayString("BUG1607")
	}
	x2d, err = emptyrex.X2d(one)
	if err != nil {
		lang.SayString("BUG1608")
	}
	x2d, err = emptystr.X2d(lang.RxFromInt32(lang.MaxArg + 1))
	if err == nil {
		lang.SayString("BUG1609")
	}
	x2d, err = lang.RxFromString("J").X2d(one)
	if err == nil {
		lang.SayString("BUG1610")
	}
	if !(lang.Rx([]rune{'0', 'A'}, false).ToString() == "0A") {
		lang.SayString("BUG1611")
	}
	if !(lang.Rx([]rune("+1.0E-1"), true).ToString() != "TEST") {
		lang.SayString("BUG1612")
	}
	if !(lang.Rx([]rune("+1.0E++1"), true).ToString() != "TEST") {
		lang.SayString("BUG1613")
	}
	if !(lang.Rx([]rune("+1.0E+۱"), true).ToString() != "TEST") {
		lang.SayString("BUG1614")
	}
	if !(lang.Rx([]rune("-.011111111111111111111111111E-۱011111111111111111111111111"), true).ToString() != "TEST") {
		lang.SayString("BUG1615")
	}
	if !(lang.Rx([]rune("+1.0E+Ꭾ"), true).ToString() != "TEST") {
		lang.SayString("BUG1616")
	}
	if !(lang.Rx([]rune("+1.0E/۱"), true).ToString() != "TEST") {
		lang.SayString("BUG1617")
	}
	re, err := emptyrex.DataType(upper_a)
	if !(re.ToString() != "TEST") {
		lang.SayString("BUG1618")
	}
	we, err := emptyrex.Words()
	if !(we.ToString() != "TEST") {
		lang.SayString("BUG1619")
	}
	run, err := emptyrex.ToRune()
	if !(run != 'j') {
		lang.SayString("BUG1620")
	}
	if !(lang.Rx([]rune("-۱.۱E+۱۱۱۱۱۱۱۱۱۱۱"), true).ToString() != "TEST") {
		lang.SayString("BUG1621")
	}
	run, err = lang.ToRuneFromString("123")
	if err == nil {
		lang.SayString("BUG8073")
	}
	char, err := lang.ToRuneFromString("Ꭳ")
	if !(char == 'Ꭳ') {
		lang.SayString("BUG1646")
	}
	char, err = lang.ToRuneFromString("B")
	if !(char == 'B') {
		lang.SayString("BUG1647")
	}
	run, err = lang.ToRuneFromRunes([]rune("123"))
	if err == nil {
		lang.SayString("BUG8074")
	}
	if !(len(lang.ToRunesFromRexx(emptyrex)) == 0) {
		lang.SayString("BUG1622")
	}
	if !(len(lang.ToRunesFromRexx(nil)) == 0) {
		lang.SayString("BUG1623")
	}
	if !(lang.ToRxFromRunes(nil) == nil) {
		lang.SayString("BUG1624")
	}
	if !(lang.ToString(nil) == "") {
		lang.SayString("BUG1625")
	}
	if !(lang.ToString(emptyrex) == "") {
		lang.SayString("BUG1626")
	}
	rx0I8, err := lang.RxFromInt32(123456).ToInt8()
	if err == nil {
		fmt.Println(rx0I8)
		lang.SayString("BUG8075")
	}
	rx0I8, err = lang.RxFromString("abc").ToInt8()
	if err == nil {
		lang.SayString("BUG8076")
	}
	rx0I6, err := lang.RxFromInt32(123456).ToInt16()
	if err == nil {
		fmt.Println(rx0I6)
		lang.SayString("BUG8077")
	}
	rx0I6, err = lang.RxFromString("abc").ToInt16()
	if err == nil {
		fmt.Println(rx0I6)
		lang.SayString("BUG8078")
	}
	rx0I, err := lang.RxFromString("123.456").ToInt32()
	if err == nil {
		fmt.Println(rx0I)
		lang.SayString("BUG8079")
	}

	rx0I, err = lang.RxFromFloat32(-1.54e+20).ToInt32()
	if err == nil {
		fmt.Println(rx0I)
		lang.SayString("BUG8080")
	}

	rx0I, err = lang.RxFromString("-.1").ToInt32()
	if err == nil {
		fmt.Println(rx0I)
		lang.SayString("BUG8081")
	}
	rx0I, err = lang.RxFromString("0-2").ToInt32()
	if err == nil {
		fmt.Println(rx0I)
		lang.SayString("BUG8082")
	}
	rx0I64, err := lang.RxFromInt32(123456).ToInt64()
	if err != nil {
		fmt.Println(rx0I64)
		lang.SayString("BUG8083")
	}
	rx0I64, err = lang.RxFromEmpty().ToInt64()
	if rx0I64 != 0 {
		lang.SayString("BUG5083")
	}
	rx0I64, err = lang.RxFromString("abc").ToInt64()
	if err == nil {
		fmt.Println(rx0I64)
		lang.SayString("BUG8084")
	}
	rx0I64, err = lang.RxFromString(".1").ToInt64()
	if err == nil {
		fmt.Println(rx0I64)
		lang.SayString("BUG8085")
	}
	rx0I64, err = zero.ToInt64()
	if err != nil {
		fmt.Println(rx0I64)
		lang.SayString("BUG8086")
	}

	rx0I64, err = lang.RxFromFloat32(-0.1234567890).ToInt64()
	if err == nil {
		fmt.Println(rx0I64)
		lang.SayString("BUG8087")
	}

	rx0I64, err = lang.RxFromFloat64(9.1e+40).ToInt64()
	if err == nil {
		fmt.Println(rx0I64)
		lang.SayString("BUG8088")
	}

	rx0I64, err = lang.RxFromFloat32(-0.1234567890).ToInt64()
	if err == nil {
		fmt.Println(rx0I64)
		lang.SayString("BUG8089")
	}

	rx0I32, err := lang.RxFromFloat32(-.1).ToFloat32()
	if err == nil {
		fmt.Println(rx0I32)
		lang.SayString("BUG8090")
	}

	rx0I32, err = lang.RxFromRune(' ').ToFloat32()
	if err == nil {
		fmt.Println(rx0I32)
		lang.SayString("BUG7090")
	}
	if !(lang.Rx([]rune("۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱۱"), true).ToString() != "TEST") {
		lang.SayString("BUG1627")
	}
	if !(lang.Rx([]rune("۱۱.۱۱"), true).ToString() != "TEST") {
		lang.SayString("BUG1628")
	}
	xi, err := lang.RxFromString("000001001101").ToInt32()
	if !(xi != -22) {
		lang.SayString("BUG1629")
	}

	tt, err := lang.Float32Pow(10.753567, 6)
	ftt := lang.RxFromFloat32(tt)
	if !(ftt.ToString() == "1546377") {
		lang.SayString("BUG1631")
	}

	sf, err := lang.Float64Pow(10.753567, 6)
	fsf := lang.RxFromFloat64(sf)
	if !(fsf.ToString() == "1546376.609420583") {
		lang.SayString("BUG1630")
	}

	the_char_array := lang.ToRunesFromRune('B')
	if !(the_char_array[0] == 'B') {
		lang.SayString("BUG1632")
	}
	// throws exception if not a single character.
	backtorune, err := lang.ToRuneFromRunes(the_char_array)
	if !(backtorune == 'B') {
		lang.SayString("BUG1633")
	}

	inFloat64 := lang.RxFromFloat64(math.MaxFloat64)
	// MATCHES NetRexx digits [17]
	if inFloat64.ToString() != "1.797693134862316E+308" {
		lang.SayString("BUG8091")
	}

	// SHOULD OVERFLOW inFloat64 rounded up
	outFloat64, err := inFloat64.ToFloat64()
	if err == nil {
		lang.SayString("BUG8091A")
	}

	// AVOID rounding math.MaxFloat64 as String
	outFloat64, err = lang.RxFromString("1.7976931348623157e+308").ToFloat64()
	if !(outFloat64 == math.MaxFloat64) {
		lang.SayString("BUG1634")
	}

	outFloat64, err = lang.RxFromRune(' ').ToFloat64()
	if err == nil {
		fmt.Println(err)
		lang.SayString("BUG7091")
	}

	inFloat32 := lang.RxFromFloat32(math.MaxFloat32)
	outFloat32, err := inFloat32.ToFloat32()
	// Match NetRexx digits [7]
	if err != nil {
		lang.SayString("BUG8092")
	} else {
		if outFloat32 != 3.402823e+38 {
			lang.SayString("BUG1635")
		}
	}

	INMX16 := lang.RxFromInt32(math.MaxInt16)
	OUTMX16, err := INMX16.ToInt16()
	if err != nil {
		fmt.Println(err)
		lang.SayString("BUG8094")
	}
	if !(OUTMX16 == math.MaxInt16) {
		lang.SayString("BUG1637")
	}
	INMN16 := lang.RxFromInt32(math.MinInt16)
	OUTMN16, err := INMN16.ToInt16()
	if err != nil {
		fmt.Println(err)
		lang.SayString("BUG8095")
	}
	if !(OUTMN16 == math.MinInt16) {
		lang.SayString("BUG1638")
	}
	INMX32 := lang.RxFromInt32(math.MaxInt32)
	OUTMX32, err := INMX32.ToInt32()
	if err != nil {
		fmt.Println(err)
		lang.SayString("BUG8096")
	}
	if !(OUTMX32 == math.MaxInt32) {
		lang.SayString("BUG1639")
	}
	INMN32 := lang.RxFromInt32(math.MinInt32)
	OUTMN32, err := INMN32.ToInt32()
	if err != nil {
		fmt.Println(err)
		lang.SayString("BUG8097")
	}
	if !(OUTMN32 == math.MinInt32) {
		lang.SayString("BUG1640")
	}

	INMX64 := lang.RxFromInt64(math.MaxInt64)
	OUTMX64, err := INMX64.ToInt64()
	if err != nil {
		fmt.Println(err)
		lang.SayString("BUG8098")
	}
	if !(OUTMX64 == math.MaxInt64) {
		lang.SayString("BUG1641")
	}

	INMN64 := lang.RxFromInt64(math.MinInt64)
	OUTMN64, err := INMN64.ToInt64()
	if err != nil {
		fmt.Println(err)
		lang.SayString("BUG8099")
	}
	if !(OUTMN64 == math.MinInt64) {
		lang.SayString("BUG1642")
	}

	// throws exception if not a single character.
	rexx_with_one_char := lang.RxFromString("Ꭳ")
	rx03, err := rexx_with_one_char.ToRune()
	if !(rx03 == 'Ꭳ') {
		lang.SayString("BUG1643")
	}
	if !(lang.ToRxFromString("NetRexxR").ToString() == "NetRexxR") {
		lang.SayString("BUG1644")
	}
	// FIXME NULL
	if !(lang.ToRxFromString(string(rune(0))).ToString() == string(rune(0))) {
		lang.SayString("BUG1645")
	}

	rexx_set = lang.RxSetWithDigitandForm(21, int8(1))
	oppow, err = a.OpPow(rexx_set, b)
	num_a, err = z.OpAdd(rexx_set, x)
	num_b, err = x.OpAdd(rexx_set, z)

	rexx_set = lang.RxSetWithDigitandForm(4, int8(2))
	oppow, err = a.OpPow(rexx_set, b)
	num_a, err = z.OpAdd(rexx_set, x)
	num_b, err = x.OpAdd(rexx_set, z)

	rexx_set = lang.RxSetWithDigitandForm(47, int8(2))
	oppow, err = a.OpPow(rexx_set, b)
	num_a, err = z.OpAdd(rexx_set, x)
	num_b, err = x.OpAdd(rexx_set, z)

	if !(lang.Rx([]rune("۱۱۱۱۱۱۱۱۱S"), true).ToString() == "۱۱۱۱۱۱۱۱۱S") {
		lang.SayString("BUG1628")
	}

	runarray, err := lang.Trunc(lang.RxFromString("   0.0050"), 100)
	if !(string(runarray) == "0.0050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000") {
		lang.SayString("BUG9628")
	}
	if !(len(lang.RxFromEmpty().Keys()) == 0) {
		lang.SayString("BUG3333")
	}
	if lang.RxFromEmpty().TestNode(lang.RxFromInt32(1)) {
		lang.SayString("BUG3334")
	}
	if lang.RxFromString("A").Equals(lang.RxFromEmpty()) {
		lang.SayString("BUG3335")
	}

	lang.RxFromString("A").Equals(nil)

	frun := lang.RxFromRune('२')
	num_a, err = five.OpAdd(rexx_set, frun)

	lang.RxFromString("0E+-9999999+")
	lang.RxFromString("-1.E+-99 999+")

	lang.Rx([]rune("½۱۱Ꭳ.۱۱२"), true)
	lang.Rx([]rune("½.½"), true)

	lang.Rx([]rune("۱۱.۱۱"), true)
	lang.Rx([]rune("۱۱.۱۱E+۱"), true)
	lang.Rx([]rune("۱۱۱۱.۱E-6"), true)
	lang.Rx([]rune("۱۱۱۱.0E+6"), true).ToInt64()

	nan, err := strconv.ParseFloat("NaN", 32)
	inf, err := strconv.ParseFloat("inf", 32)
	infp, err := strconv.ParseFloat("+inf", 32)
	infn, err := strconv.ParseFloat("-inf", 32)

	flea := lang.RxFromFloat64(nan)
	fleb := lang.RxFromFloat64(inf)
	flec := lang.RxFromFloat64(infp)
	fled := lang.RxFromFloat64(infn)

	flee := lang.RxFromFloat32(float32(nan))
	flef := lang.RxFromFloat32(float32(inf))
	fleg := lang.RxFromFloat32(float32(infp))
	fleh := lang.RxFromFloat32(float32(infn))

	flea.ToFloat64()
	fleb.ToFloat64()
	flec.ToFloat64()
	fled.ToFloat64()
	flee.ToFloat32()
	flef.ToFloat32()
	fleg.ToFloat32()
	fleh.ToFloat32()

	lang.RxFromInt32(19).OpEqS(nil, lang.RxFromInt32(18))

	lang.RxFromString("259.9").D2b(eight)
	lang.RxFromString("259.9").D2x(two)
	lang.RxFromString("15").D2b(lang.RxFromString("15"))
	lang.RxFromString("64").D2x(one)

	lang.RxFromString("561").D2b(one)
	lang.RxFromString("261").D2x(one)

	new_set.SetDigits(zero)
	new_set.SetDigits(lower_a)
	new_set.SetDigits(lang.RxFromBool(true))
	new_set.SetForm(lang.RxFromBool(true))
	new_set.SetForm(lang.RxFromString("engineering"))
	new_set.SetForm(lang.RxFromString("scientific"))
	lang.RxUtil()

	x = lang.RxFromString("7.62939453629394536293945362939453E+28")

	// lang.SayString("chars")
	// fmt.Println(fmt.Sprintf("%v", reflect.ValueOf(x).Elem().Field(0)))
	// lang.SayString("ind")
	// fmt.Println(fmt.Sprintf("%v", reflect.ValueOf(x).Elem().Field(1)))
	// lang.SayString("form")
	// fmt.Println(fmt.Sprintf("%v", reflect.ValueOf(x).Elem().Field(2)))
	// lang.SayString("mant")
	// fmt.Println(fmt.Sprintf("%v", reflect.ValueOf(x).Elem().Field(3)))
	// lang.SayString("exp")
	// fmt.Println(fmt.Sprintf("%v", reflect.ValueOf(x).Elem().Field(4)))
	// lang.SayString("dig")
	// fmt.Println(fmt.Sprintf("%v", reflect.ValueOf(x).Elem().Field(5)))
	// lang.SayString("coll")
	// fmt.Println(fmt.Sprintf("%v", reflect.ValueOf(x).Elem().Field(6)))

	lang.D2x(x, -0)
	lang.D2x(z, -100)
	lang.Float64Pow(10.753567, 0)
	lang.Float64Pow(10.753567, -1)
	lang.Trunc(x, 0)
	lang.Trunc(lower_a, 5)

	lang.RxWords()

	lang.Abbrev([]rune{1}, []rune{1, 2}, 0)

	lang.Centre([]rune{1}, 1, 'c')

	lang.DelWord([]rune{1}, 0, 0)

	lang.Verify([]rune{1}, []rune{1}, 'N', 1)

	lang.Verifym([]rune{1}, []rune{1}, 2)

	lang.Verifym([]rune{1}, []rune{}, 1)

	lang.Verifym([]rune{1}, []rune{1}, 1)

	lang.Word([]rune{1}, 1)

	lang.WordLength([]rune{1}, 1)

	lang.WordPos([]rune{32}, []rune{1}, 1)

	lang.WordPos([]rune{65, 32, 32, 66, 32}, []rune{}, -3)

	opadd, err = lang.RxFromString("0.4444444444").OpAdd(nil, lang.RxFromString("0.5555555550"))

	lang.Format(lang.RxFromString(".2111"), 0, 2, 9, 9, 'S')

	// TEST ALL RUNES

	unigroup := [370]rune{
		'\u0030', '\u0031', '\u0032', '\u0033', '\u0034', '\u0035', '\u0036', '\u0037', '\u0038', '\u0039',
		'\u0660', '\u0661', '\u0662', '\u0663', '\u0664', '\u0665', '\u0666', '\u0667', '\u0668', '\u0669',
		'\u06F0', '\u06F1', '\u06F2', '\u06F3', '\u06F4', '\u06F5', '\u06F6', '\u06F7', '\u06F8', '\u06F9',
		'\u07C0', '\u07C1', '\u07C2', '\u07C3', '\u07C4', '\u07C5', '\u07C6', '\u07C7', '\u07C8', '\u07C9',
		'\u0966', '\u0967', '\u0968', '\u0969', '\u096A', '\u096B', '\u096C', '\u096D', '\u096E', '\u096F',
		'\u09E6', '\u09E7', '\u09E8', '\u09E9', '\u09EA', '\u09EB', '\u09EC', '\u09ED', '\u09EE', '\u09EF',
		'\u0A66', '\u0A67', '\u0A68', '\u0A69', '\u0A6A', '\u0A6B', '\u0A6C', '\u0A6D', '\u0A6E', '\u0A6F',
		'\u0AE6', '\u0AE7', '\u0AE8', '\u0AE9', '\u0AEA', '\u0AEB', '\u0AEC', '\u0AED', '\u0AEE', '\u0AEF',
		'\u0B66', '\u0B67', '\u0B68', '\u0B69', '\u0B6A', '\u0B6B', '\u0B6C', '\u0B6D', '\u0B6E', '\u0B6F',
		'\u0BE6', '\u0BE7', '\u0BE8', '\u0BE9', '\u0BEA', '\u0BEB', '\u0BEC', '\u0BED', '\u0BEE', '\u0BEF',
		'\u0C66', '\u0C67', '\u0C68', '\u0C69', '\u0C6A', '\u0C6B', '\u0C6C', '\u0C6D', '\u0C6E', '\u0C6F',
		'\u0CE6', '\u0CE7', '\u0CE8', '\u0CE9', '\u0CEA', '\u0CEB', '\u0CEC', '\u0CED', '\u0CEE', '\u0CEF',
		'\u0D66', '\u0D67', '\u0D68', '\u0D69', '\u0D6A', '\u0D6B', '\u0D6C', '\u0D6D', '\u0D6E', '\u0D6F',
		'\u0DE6', '\u0DE7', '\u0DE8', '\u0DE9', '\u0DEA', '\u0DEB', '\u0DEC', '\u0DED', '\u0DEE', '\u0DEF',
		'\u0E50', '\u0E51', '\u0E52', '\u0E53', '\u0E54', '\u0E55', '\u0E56', '\u0E57', '\u0E58', '\u0E59',
		'\u0ED0', '\u0ED1', '\u0ED2', '\u0ED3', '\u0ED4', '\u0ED5', '\u0ED6', '\u0ED7', '\u0ED8', '\u0ED9',
		'\u0F20', '\u0F21', '\u0F22', '\u0F23', '\u0F24', '\u0F25', '\u0F26', '\u0F27', '\u0F28', '\u0F29',
		'\u1040', '\u1041', '\u1042', '\u1043', '\u1044', '\u1045', '\u1046', '\u1047', '\u1048', '\u1049',
		'\u1090', '\u1091', '\u1092', '\u1093', '\u1094', '\u1095', '\u1096', '\u1097', '\u1098', '\u1099',
		'\u17E0', '\u17E1', '\u17E2', '\u17E3', '\u17E4', '\u17E5', '\u17E6', '\u17E7', '\u17E8', '\u17E9',
		'\u1810', '\u1811', '\u1812', '\u1813', '\u1814', '\u1815', '\u1816', '\u1817', '\u1818', '\u1819',
		'\u1946', '\u1947', '\u1948', '\u1949', '\u194A', '\u194B', '\u194C', '\u194D', '\u194E', '\u194F',
		'\u19D0', '\u19D1', '\u19D2', '\u19D3', '\u19D4', '\u19D5', '\u19D6', '\u19D7', '\u19D8', '\u19D9',
		'\u1A80', '\u1A81', '\u1A82', '\u1A83', '\u1A84', '\u1A85', '\u1A86', '\u1A87', '\u1A88', '\u1A89',
		'\u1A90', '\u1A91', '\u1A92', '\u1A93', '\u1A94', '\u1A95', '\u1A96', '\u1A97', '\u1A98', '\u1A99',
		'\u1B50', '\u1B51', '\u1B52', '\u1B53', '\u1B54', '\u1B55', '\u1B56', '\u1B57', '\u1B58', '\u1B59',
		'\u1BB0', '\u1BB1', '\u1BB2', '\u1BB3', '\u1BB4', '\u1BB5', '\u1BB6', '\u1BB7', '\u1BB8', '\u1BB9',
		'\u1C40', '\u1C41', '\u1C42', '\u1C43', '\u1C44', '\u1C45', '\u1C46', '\u1C47', '\u1C48', '\u1C49',
		'\u1C50', '\u1C51', '\u1C52', '\u1C53', '\u1C54', '\u1C55', '\u1C56', '\u1C57', '\u1C58', '\u1C59',
		'\uA620', '\uA621', '\uA622', '\uA623', '\uA624', '\uA625', '\uA626', '\uA627', '\uA628', '\uA629',
		'\uA8D0', '\uA8D1', '\uA8D2', '\uA8D3', '\uA8D4', '\uA8D5', '\uA8D6', '\uA8D7', '\uA8D8', '\uA8D9',
		'\uA900', '\uA901', '\uA902', '\uA903', '\uA904', '\uA905', '\uA906', '\uA907', '\uA908', '\uA909',
		'\uA9D0', '\uA9D1', '\uA9D2', '\uA9D3', '\uA9D4', '\uA9D5', '\uA9D6', '\uA9D7', '\uA9D8', '\uA9D9',
		'\uA9F0', '\uA9F1', '\uA9F2', '\uA9F3', '\uA9F4', '\uA9F5', '\uA9F6', '\uA9F7', '\uA9F8', '\uA9F9',
		'\uAA50', '\uAA51', '\uAA52', '\uAA53', '\uAA54', '\uAA55', '\uAA56', '\uAA57', '\uAA58', '\uAA59',
		'\uABF0', '\uABF1', '\uABF2', '\uABF3', '\uABF4', '\uABF5', '\uABF6', '\uABF7', '\uABF8', '\uABF9',
		'\uFF10', '\uFF11', '\uFF12', '\uFF13', '\uFF14', '\uFF15', '\uFF16', '\uFF17', '\uFF18', '\uFF19',
	}

	hexgroup := [260]rune{
		0x104A0, 0x104A1, 0x104A2, 0x104A3, 0x104A4, 0x104A5, 0x104A6, 0x104A7, 0x104A8, 0x104A9,
		0x10D30, 0x10D31, 0x10D32, 0x10D33, 0x10D34, 0x10D35, 0x10D36, 0x10D37, 0x10D38, 0x10D39,
		0x11066, 0x11067, 0x11068, 0x11069, 0x1106A, 0x1106B, 0x1106C, 0x1106D, 0x1106E, 0x1106F,
		0x110F0, 0x110F1, 0x110F2, 0x110F3, 0x110F4, 0x110F5, 0x110F6, 0x110F7, 0x110F8, 0x110F9,
		0x11136, 0x11137, 0x11138, 0x11139, 0x1113A, 0x1113B, 0x1113C, 0x1113D, 0x1113E, 0x1113F,
		0x111D0, 0x111D1, 0x111D2, 0x111D3, 0x111D4, 0x111D5, 0x111D6, 0x111D7, 0x111D8, 0x111D9,
		0x112F0, 0x112F1, 0x112F2, 0x112F3, 0x112F4, 0x112F5, 0x112F6, 0x112F7, 0x112F8, 0x112F9,
		0x11450, 0x11451, 0x11452, 0x11453, 0x11454, 0x11455, 0x11456, 0x11457, 0x11458, 0x11459,
		0x114D0, 0x114D1, 0x114D2, 0x114D3, 0x114D4, 0x114D5, 0x114D6, 0x114D7, 0x114D8, 0x114D9,
		0x11650, 0x11651, 0x11652, 0x11653, 0x11654, 0x11655, 0x11656, 0x11657, 0x11658, 0x11659,
		0x116C0, 0x116C1, 0x116C2, 0x116C3, 0x116C4, 0x116C5, 0x116C6, 0x116C7, 0x116C8, 0x116C9,
		0x11730, 0x11731, 0x11732, 0x11733, 0x11734, 0x11735, 0x11736, 0x11737, 0x11738, 0x11739,
		0x118E0, 0x118E1, 0x118E2, 0x118E3, 0x118E4, 0x118E5, 0x118E6, 0x118E7, 0x118E8, 0x118E9,
		0x11C50, 0x11C51, 0x11C52, 0x11C53, 0x11C54, 0x11C55, 0x11C56, 0x11C57, 0x11C58, 0x11C59,
		0x11D50, 0x11D51, 0x11D52, 0x11D53, 0x11D54, 0x11D55, 0x11D56, 0x11D57, 0x11D58, 0x11D59,
		0x11DA0, 0x11DA1, 0x11DA2, 0x11DA3, 0x11DA4, 0x11DA5, 0x11DA6, 0x11DA7, 0x11DA8, 0x11DA9,
		0x16A60, 0x16A61, 0x16A62, 0x16A63, 0x16A64, 0x16A65, 0x16A66, 0x16A67, 0x16A68, 0x16A69,
		0x16B50, 0x16B51, 0x16B52, 0x16B53, 0x16B54, 0x16B55, 0x16B56, 0x16B57, 0x16B58, 0x16B59,
		0x1D7CE, 0x1D7CF, 0x1D7D0, 0x1D7D1, 0x1D7D2, 0x1D7D3, 0x1D7D4, 0x1D7D5, 0x1D7D6, 0x1D7D7,
		0x1D7D8, 0x1D7D9, 0x1D7DA, 0x1D7DB, 0x1D7DC, 0x1D7DD, 0x1D7DE, 0x1D7DF, 0x1D7E0, 0x1D7E1,
		0x1D7E2, 0x1D7E3, 0x1D7E4, 0x1D7E5, 0x1D7E6, 0x1D7E7, 0x1D7E8, 0x1D7E9, 0x1D7EA, 0x1D7EB,
		0x1D7EC, 0x1D7ED, 0x1D7EE, 0x1D7EF, 0x1D7F0, 0x1D7F1, 0x1D7F2, 0x1D7F3, 0x1D7F4, 0x1D7F5,
		0x1D7F6, 0x1D7F7, 0x1D7F8, 0x1D7F9, 0x1D7FA, 0x1D7FB, 0x1D7FC, 0x1D7FD, 0x1D7FE, 0x1D7FF,
		// next two rows : go1.13.1 openbsd[6.6]/amd64 unicode.IsDigit not working for me
		//0x1E140, 0x1E141, 0x1E142, 0x1E143, 0x1E144, 0x1E145, 0x1E146, 0x1E147, 0x1E148, 0x1E149,
		//0x1E2F0, 0x1E2F1, 0x1E2F2, 0x1E2F3, 0x1E2F4, 0x1E2F5, 0x1E2F6, 0x1E2F7, 0x1E2F8, 0x1E2F9,
		0x1E950, 0x1E951, 0x1E952, 0x1E953, 0x1E954, 0x1E955, 0x1E956, 0x1E957, 0x1E958, 0x1E959,
	}

	count := 0
	var col int32 = 0
	for row := 1; row <= 37; row++ {
		for col = 0; col <= 9; col++ {
			number, err := lang.RxFromRune(unigroup[count]).ToInt32()
			if err != nil {
				fmt.Printf("err: row %d col %d\n", row, col)
			} else {
				if number != col {
					fmt.Printf("number: %d row %d col %d\n", number, row, col)
				}
			}
			count++
		}
	}

	count = 0
	for row := 1; row <= 24; row++ {
		for col = 0; col <= 9; col++ {
			number, err := lang.RxFromRune(hexgroup[count]).ToInt32()
			if err != nil {
				fmt.Printf("err: row %d col %d\n", row, col)
			} else {
				if number != col {
					fmt.Printf("number: %d row %d col %d\n", number, row, col)
				}
			}
			count++
		}
	}

	lang.RxIO()

	// UNCOMMENT TO USE SEE KEYBOARD INPUT
	// line := lang.Ask()
	// lang.SayRexx(line)

	// kb := lang.AskOne()
	// lang.SayRexx(kb)

	lang.SayString("you should see the lower case 'a' below")
	lang.SayString("a")
	lang.SayString("another lower case 'a'")
	lang.SayRexx(lower_a)
	lang.SayString("still another lower case 'a'")
	lang.SayRune('a')
	lang.SayString("you should see zero below")
	lang.SayInt64(0)
	lang.SayString("another zero")
	lang.SayFloat32(0.0)
	lang.SayString("still another zero")
	lang.SayFloat64(0.0)
	lang.SayString("you should 1 the value for Bool True")
	lang.SayBool(true)
	lang.SayString("you should see a blank line below")
	lang.Say(nil)
	lang.SayString("another blank line")
	lang.Say(make([]rune, 0))

	vas := make([]rune, 1)
	vas[0] = '\000'

	lang.SayString("still another  blank line")
	lang.Say(vas)

	opminus, err = lower_a.OpMinus(rexx_set)
	if err == nil {
		lang.SayString("BUG1132A")
	}

	lang.RxFromString("1.797693134862315708145274237317043567981e+309").ToFloat64()
	lang.RxFromString("-1.797693134862315708145274237317043567981e+309").ToFloat64()

	//~ 1167
	//Internal error: Rexx: bad dig -10
	_, err = lang.Format(lang.RxFromString("-0"), -10, 50, 19, -10, 'E')
	if err != nil {
		lang.SayString("7766")
	} else {
		lang.SayString("you should see internal error above")
	}

	opmult, err = lang.RxFromString("50.000000").OpPow(rexx_set, lang.RxFromString("17.000"))

	rexsult, err := lang.RxFromString("10.23456785").OpSub(lang.RxSetWithDigit(9), lang.RxFromString("10.23456789"))
	if (err == nil) && !(rexsult.ToString() == "0") {
		lang.SayString("subx121")
	}

	lang.RxFromString("9E+999999999").OpMult(lang.RxSetWithDigit(9), lang.RxFromString("+1.23456789012345E-0"))

	lang.RxFromString("100").OpPow(lang.RxSetWithDigit(16), lang.RxFromString("1E+9"))

	opdiv, err = lang.RxFromInt8(int8(-128)).OpDiv(nil, lang.RxFromString("-34359738368"))

	lang.RxFromString("-0.9E-999999999").OpDiv(nil, z)

	value_a.OpAdd(lang.RxSetWithDigitandForm(0, int8(0)), one_point_zero)

	opcc = emptyrex.OpCc(rexx_set, emptyrex)
	// FIXME ADD IF TEST

	big_set := lang.RxSet()
	big_set.SetForm(lang.RxFromString("engineering"))
	lang.RxFromString("-0.9E-999999999").OpDiv(big_set, z)

	new_set.SetDigits(lang.RxFromString("1.7976931348623157e-308"))

	lang.D2x(z, 1)

	x2c, err = lang.RxFromString("FFFFFFFFFFFFFFFFFFFFFFFF").X2c()

	lang.X2d(x, 1111111110)

	lang.Float32Pow(-1.12345, lang.MaxArg)

	lang.Float32Pow(1.797693134862315708145274237317043567981e+308, 9)

	lang.Trunc(lang.RxFromString(".3E+25"), 10)

	lang.Format(lower_a, 1, 1, 1, 1, 'S')

	lang.Format(one_point_zero, 0, 0, 0, 0, 'E')

	//~ MEMORY drop MAX MINS
	//format, err = lang.RxFromString(".022E-999").Format(five, lang.RxFromInt32(lang.MaxExp-1000), five, five, lang.RxFromRune('P'))

	lang.Format(lang.RxFromString(".1"), 0, 0, lang.MinExp+1, 0, 'E')

	lang.Format(x, 10, 1, -100, 100, 'S')

	lang.Format(zero, 1, -1, 1, 1, 'c')

	lang.Format(zero, 0, 1, 1, 1, 'b')

	format, err = lang.RxFromString("2.73").Format(lang.RxFromInt32(lang.MaxExp), five, nil, nil, nil)

	lang.Format(lang.RxFromString(".1"), 0, 0, lang.MaxExp+1, 0, 'E')

	lang.Space([]rune{32, 65}, 1, 0)

	lang.SubWord([]rune{65}, 1, 0)

	lang.WordPos([]rune{65, 32, 32, 66, 32}, []rune{65, 32, 32, 66, 32}, 1)

	lang.WordPos([]rune{32, 27, 32, 27, 32, 27}, []rune{32, 27}, 1)

	lang.WordPos([]rune{32, 27, 32, 27, 32, 27}, []rune{32, 27, 32}, 1)

	lang.WordPos([]rune{65, 32, 32, 66, 32}, []rune{65, 32, 32, 66, 32}, 2)

	lang.RxParse()

	obj := lang.RxFromString("jan,feb,mar,apr,may,jun")

	var list = []rune{2, 1, 44, 10, 1, 0, 2, 1, 44, 10, 1, 1, 2, 1, 44, 10, 1, 2, 2, 1, 44, 10, 1, 3, 1, 10, 1, 4, 0}

	vars := make([]*lang.Rexx, 5)

	//var a, b, c, d, e *lang.Rexx

	//var newline = lang.RxFromRune('\n')

	lang.Parse(obj, list, vars)

	//a = vars[0]
	//b = vars[1]
	//c = vars[2]
	//d = vars[3]
	//e = vars[4]

	//lang.SayRexx(a.OpCc(nil, newline).OpCc(nil, b).OpCc(nil, newline).OpCc(nil, c).OpCc(nil, newline).OpCc(nil, d).OpCc(nil, newline).OpCc(nil, e))

	//lang.SayRexx(a.OpCc(nil, newline))

	vars = make([]*lang.Rexx, 5)
	vars[1] = x
	vars[3] = z
	lang.Parse(lang.ToRxFromString("Flying pigs have wings"), []rune{7, 1, 10, 1, 0, 9, 3, 10, 1, 2, 1, 10, 1, 4, 0}, vars)

	vars = make([]*lang.Rexx, 5)
	vars[1] = lang.RxFromInt32(lang.MinCol - 1)
	vars[3] = lang.RxFromInt32(lang.MaxCol + 1)
	lang.Parse(lang.ToRxFromString("Flying pigs have wings"), []rune{7, 1, 10, 1, 0, 9, 3, 10, 1, 2, 1, 10, 1, 4, 0}, vars)

	vars = make([]*lang.Rexx, 5)
	vars[1] = five
	vars[3] = five
	lang.Parse(lang.ToRxFromString("Flying pigs have wings"), []rune{7, 1, 10, 1, 0, 8, 3, 10, 1, 2, 1, 10, 1, 4, 0}, vars)

	vars = make([]*lang.Rexx, 5)
	vars[1] = five
	vars[3] = five
	lang.Parse(lang.ToRxFromString("Flying pigs have wings"), []rune{7, 1, 10, 1, 0, 9, 3, 10, 1, 2, 1, 10, 1, 4, 0}, vars)

	vars = make([]*lang.Rexx, 7)
	vars[1] = five
	vars[3] = five
	lang.Parse(lang.ToRxFromString(""), []rune{7, 1, 10, 1, 0, 7, 3, 10, 1, 2, 5, 1, 2, 10, 1, 4, 1, 10, 2, 5, 6, 0}, vars)

	vars = make([]*lang.Rexx, 5)
	lang.Parse(lang.ToRxFromString("This a remark field (up to 39 chars)   44 55 //66    77"), []rune{2, 2, 47, 47, 10, 3, 0, 1, 2, 1, 10, 3, 3, 4, 2, 0}, vars)

	vars = make([]*lang.Rexx, 3)
	lang.Parse(lang.ToRxFromString("This is a sentence."), []rune{1, 10, 3, 0, 1, 2, 0}, vars)

	vars = make([]*lang.Rexx, 2)
	lang.Parse(lang.ToRxFromString("To be, or not to be?"), []rune{2, 1, 44, 10, 1, 0, 1, 10, 1, 1, 0}, vars)

	vars = make([]*lang.Rexx, 4)
	lang.Parse(lang.ToRxFromString("To be, or not to be?"), []rune{2, 1, 44, 10, 1, 0, 1, 10, 3, 1, 2, 3, 0}, vars)

	vars = make([]*lang.Rexx, 2)
	lang.Parse(lang.ToRxFromString("Flying pigs have wings"), []rune{3, 1, 5, 10, 1, 0, 1, 10, 1, 1, 0}, vars)

	vars = make([]*lang.Rexx, 3)
	lang.Parse(lang.ToRxFromString("Flying pigs have wings"), []rune{3, 1, 5, 10, 1, 0, 3, 1, 10, 10, 1, 1, 1, 10, 1, 2, 0}, vars)

	vars = make([]*lang.Rexx, 3)
	lang.Parse(lang.ToRxFromString("Flying pigs have wings"), []rune{3, 1, 5, 10, 1, 0, 4, 1, 5, 10, 1, 1, 1, 10, 1, 2, 0}, vars)

	vars = make([]*lang.Rexx, 3)
	lang.Parse(lang.ToRxFromString("123456789"), []rune{3, 1, 3, 4, 1, 3, 10, 1, 0, 3, 1, 3, 10, 1, 1, 1, 10, 1, 2, 0}, vars)

	vars = make([]*lang.Rexx, 0)
	lang.Parse(lang.ToRxFromString("1234 56789"), []rune{5, 0, 4, 0, 0}, vars)

	vars = make([]*lang.Rexx, 4)
	lang.Parse(lang.ToRxFromString("L/look for/1 10"), []rune{3, 1, 2, 10, 1, 0, 4, 1, 1, 10, 1, 1, 6, 1, 10, 1, 2, 1, 10, 1, 3, 0}, vars)

	vars = make([]*lang.Rexx, 4)
	lang.Parse(lang.ToRxFromString(""), []rune{3, 1, 2, 10, 1, 0, 4, 1, 1, 10, 1, 1, 6, 1, 10, 1, 2, 1, 10, 1, 3, 0}, vars)

	vars = make([]*lang.Rexx, 4)
	lang.Parse(lang.ToRxFromString("MINUS"), []rune{3, 1, 2, 10, 1, 0, 5, 1, 10, 10, 1, 1, 6, 1, 10, 1, 2, 1, 10, 1, 3, 0}, vars)

	vars = make([]*lang.Rexx, 4)
	lang.Parse(lang.RxFromInt32(0), []rune{3, 0, 0}, vars)

	// Working on
	insmallFloat32 := lang.RxFromFloat32(math.SmallestNonzeroFloat32)
	//lang.SayRexx(insmallFloat32)
	outsmallFloat32, err := insmallFloat32.ToFloat32()
	// Overflow
	if err == nil {
		fmt.Println(err)
		lang.SayString("BUG8093")
	}
	//"1.401298464324817e-45"
	//fmt.Println(math.SmallestNonzeroFloat32)
	// if !(outsmallFloat32 == math.SmallestNonzeroFloat32) {
	// 	//fmt.Println(outsmallFloat32)
	// 	lang.SayString("BUG1636")
	// }
	outsmallFloat32, err = lang.RxFromString("1.401298464324817e-45").ToFloat32()
	//fmt.Println(outsmallFloat32)
	//fmt.Println(math.SmallestNonzeroFloat32)
	if outsmallFloat32 != math.SmallestNonzeroFloat32 {
		lang.SayString("BUG1634")
	}

	// OUT OF MEMORY
	// b2d, err = lang.RxFromString("123").B2d(lang.RxFromInt32(lang.MaxExp))

	// OUT OF MEMORY
	// lang.RxFromString("12.3").Trunc(lang.RxFromInt32(lang.MaxArg))

	// OUT OF MEMORY
	// bigcopy, err := one.Copies(lang.RxFromInt32(66846720))
	// if bigcopy == nil {
	// 	lang.SayString("BOMB")
	// }
	// bigcopy.Right(one, zero)

	// OUT OF MEMORY
	// one.Copies(lang.RxFromInt32(99999999))

	// INDEX OUT OF RANGE
	// fmt.Println(lang.Space([]rune{64, 32, 32, 54}, -10000, 32))
	// lang.Space([]rune{32, 65}, 1, 32)

	// FIXME index out of range [4] with length 4 malformed []rune{}
	// vars = make([]*lang.Rexx, 4)
	// vars[1] = lang.RxFromRune(',')
	// lang.Parse(lang.ToRxFromString("To be, or not to be?"), []rune{6, 1, 10, 1, 0, 1, 10, 3, 2, 3, 4, 0}, vars)

	return
}
