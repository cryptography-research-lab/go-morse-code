package morse_code

import "strings"

// Encode 对输入的内容进行摩斯密码编码
func Encode(plaintext string) ([]string, error) {
	plaintextRuneSlice := []rune(plaintext)
	result := make([]string, len(plaintextRuneSlice))
	for index, character := range plaintextRuneSlice {
		if to, exists := DefaultMorseCodeTable[character]; exists {
			result[index] = to
		} else {
			return nil, ErrNotSupportMorseCharacter
		}
	}
	return result, nil
}

// EncodeToString 编码为摩尔斯密码的字符串
func EncodeToString(plaintext string) (string, error) {
	encode, err := Encode(plaintext)
	if err != nil {
		return "", err
	}
	return strings.Join(encode, " "), nil
}

// Decode 对输入的内容进行摩斯密码解码，多个摩尔斯密码之间使用空格分割
func Decode(ciphertext string) (string, error) {
	split := strings.Split(ciphertext, " ")
	return DecodeMorseCodeSlice(split)
}

// DecodeMorseCodeSlice 解码摩尔斯密码数组
func DecodeMorseCodeSlice(morseCodeSlice []string) (string, error) {
	resultSlice := make([]rune, len(morseCodeSlice))
	for index, morseCode := range morseCodeSlice {
		character, err := DefaultMorseSearchTree.Query(morseCode)
		if err != nil {
			return "", err
		}
		resultSlice[index] = character
	}
	return string(resultSlice), nil
}
