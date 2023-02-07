package morse_code

import "errors"

var (
	// ErrInputMustAllLetter 输入的要编码的内容必须都是英文字母，否则会返回此错误
	ErrInputMustAllLetter = errors.New("input must all letter")
)

var (
	// ErrMorseCodeCanNotEmpty 莫斯密码不能为空
	ErrMorseCodeCanNotEmpty = errors.New("morse code can not empty")

	// ErrNotSupportMorseCharacter 不支持的摩尔斯码字符
	ErrNotSupportMorseCharacter = errors.New("not support morse character")

	// ErrMorseCodeNotFound 摩尔斯密码没有找到，可能是输入了非法的摩尔斯密码
	ErrMorseCodeNotFound = errors.New("morse code not found'")
)
