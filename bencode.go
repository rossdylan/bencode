package bencode

import (
	"strconv"
)

// an interfacing describing the functions needed to be a composite bencoded object (BencodedList, BencodedDict)
type BencodedComposite interface {
	AppendString()
	AppendInt()
	AppendList()
	AppendDict()
	StringValue()
}

// turn a normal string into a Bencoded string
func BencodeString(str string) string {
	return strconv.Itoa(len(str)) + ":" + str
}

// turn a normal int into a Bencoded int
func BencodeInt(i int) string {
	return "i" + strconv.Itoa(i) + "e"
}
