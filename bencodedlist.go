package bencode

// Store the string value of a list of bencoded values
type BencodedList struct {
	str_value string
}

// Get the String value of a BencodedList
func (bl *BencodedList) StringValue() string {
	return "l" + bl.str_value + "e"
}

// Append an int to a BencodedList
func (bl *BencodedList) AppendInt(i int) {
	bl.str_value += BencodeInt(i)
}

// Append a String to a BencodedList
func (bl *BencodedList) AppendString(str string) {
	bl.str_value += BencodeString(str)
}

// Append a BencodedList to a BencodedList
func (bl *BencodedList) AppendBencodedList(l BencodedList) {
	bl.str_value += l.StringValue()
}

// Append a BencodedDict to a BencodedList
func (bl *BencodedList) AppendBencodedDict(d BencodedDict) {
	bl.str_value += d.StringValue()
}
