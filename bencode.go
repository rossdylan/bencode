package bencode

import (
	"strconv"
)

type BencodedComposite interface {
	AppendString()
	AppendInt()
	AppendList()
	AppendDict()
	StringValue()
}

func BencodeString(str string) string {
	return strconv.Itoa(len(str)) + ":" + str
}

func BencodeInt(i int) string {
	return "i" + strconv.Itoa(i) + "e"
}
