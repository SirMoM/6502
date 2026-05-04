package computer

import "log"

func ToWord(bytes any) Word {
	var word Word
	switch v := bytes.(type) {
	case Instruction:
		word = Word(v)
	case uint8:
		word = Word(v)
	case int:
		word = Word(v)
	case Word:
		word = v
	default:
		log.Fatalf("unsupported byte type %T", v)
	}
	return word
}
