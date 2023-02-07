package morse_code

// Table 表示一个摩尔斯密码解密表
type Table map[rune]string

// CanEncode 判断这个字符使用这个摩尔斯密码表是否可以编码
func (x Table) CanEncode(character rune) bool {
	_, exists := x[character]
	return exists
}

// ToMorseSearchTree 将摩尔斯加密表转为适合解密使用的摩尔斯查询树
func (x Table) ToMorseSearchTree() (*MorseSearchTree, error) {
	tree := NewMorseSearchTree()
	for character, morseCode := range x {
		err := tree.AddMapping(morseCode, character)
		if err != nil {
			return nil, err
		}
	}
	return tree, nil
}

// DefaultMorseCodeTable 默认的摩尔斯密码表
var DefaultMorseCodeTable Table = map[rune]string{

	// 英文字母
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.--",
	'Z': "--..",

	// 数字
	'0': "-----",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",

	// 标点符号
	'.':  ".-.-.-",
	',':  "--..--",
	'?':  "..--..",
	'\'': ".----.",
	'!':  "-.-.--",
	'/':  "-..-.",
	'(':  "-.--.",
	')':  "-.--.-",
	'&':  ".-...",
	':':  "---...",
	';':  "-.-.-.",
	'=':  "-...-",
	'+':  ".-.-.",
	'-':  "-....-",
	'_':  "..--.-",
	'"':  ".-..-.",
	'$':  "...-..-",
	'@':  ".--.-.",
	'¿':  "..-.-",
	'¡':  "--...-",
}

// DefaultMorseSearchTree 默认的摩尔斯密码表对应的搜索树
var DefaultMorseSearchTree *MorseSearchTree

func init() {
	DefaultMorseSearchTree, _ = DefaultMorseCodeTable.ToMorseSearchTree()
}
