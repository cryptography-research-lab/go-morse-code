package morse_code

// MorseSearchTree 摩尔斯搜索树，我随便编的名字
type MorseSearchTree struct {
	// 摩尔斯密码的根节点
	root *MorseSearchTreeNode
}

func NewMorseSearchTree() *MorseSearchTree {
	return &MorseSearchTree{
		root: &MorseSearchTreeNode{},
	}
}

// AddMapping 添加一个映射关系到查询树上
// morseCode: 摩尔斯密码
// character: 对应的字符
func (x *MorseSearchTree) AddMapping(morseCode string, character rune) error {
	if len(morseCode) == 0 {
		return ErrMorseCodeCanNotEmpty
	}
	currentNode := x.root
	for _, character := range morseCode {
		if character == '.' {
			// left
			if currentNode.Left == nil {
				currentNode.Left = &MorseSearchTreeNode{Character: character}
			}
			currentNode = currentNode.Left
		} else if character == '-' {
			// right
			if currentNode.Right == nil {
				currentNode.Right = &MorseSearchTreeNode{Character: character}
			}
			currentNode = currentNode.Right
		} else {
			// error
			return ErrNotSupportMorseCharacter
		}
	}
	// 如果已经存在的话则直接覆盖
	currentNode.Character = character
	return nil
}

// Query 添加一个映射关系到查询树上
func (x *MorseSearchTree) Query(morseCode string) (rune, error) {
	if len(morseCode) == 0 {
		return 0, ErrMorseCodeCanNotEmpty
	}
	currentNode := x.root
	for _, character := range []rune(morseCode) {
		if character == '.' {
			// left
			if currentNode.Left == nil {
				return 0, nil
			}
			currentNode = currentNode.Left
		} else if character == '-' {
			// right
			if currentNode.Right == nil {
				return 0, nil
			}
			currentNode = currentNode.Right
		} else {
			// error
			return 0, ErrNotSupportMorseCharacter
		}
		if currentNode == nil {
			break
		}
	}
	// 如果没有找到摩尔斯密码对应的字符，则说明树可能创建错了
	if currentNode == nil {
		return 0, ErrMorseCodeNotFound
	} else {
		// 找到了对应的节点，则返回
		return currentNode.Character, nil
	}
}

// MorseSearchTreeNode 用于表示搜索树的上的一个节 点
type MorseSearchTreeNode struct {

	// 当前节点绑定的字符是啥
	Character rune

	// 左右孩子节点
	Left, Right *MorseSearchTreeNode
}
