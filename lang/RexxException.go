package lang

import "errors"

var ExceptionType = []string{"ArgumentNullException", "BadArgumentException", "BadColumnException", "BadNumericException", "DivideException", "ExponentOverflowException", "NoOtherwiseException", "NotCharacterException", "NotLogicException", "NumberFormatException"}

func RxException(kind int, detail string) error {
	message := ExceptionType[kind] + " : " + detail
	return errors.New(message)
}
