package bencode

//Hold the string values for a Bencoded dict inside a struct
type BencodedDict struct {
	str_value string
}

//Return the string value of a BencodedDict
func (bd *BencodedDict) StringValue() string {
	return "d" + bd.str_value + "e"
}

//Add an key with an int value to the dict
func (bd *BencodedDict) AppendInt(key string, value int) {
	bd.str_value += BencodeString(key) + BencodeInt(value)
}

//Add an key with an string value to the dict
func (bd *BencodedDict) AppendString(key string, value string) {
	bd.str_value += BencodeString(key) + BencodeString(value)
}

//Add an key with an BencodedList value to the dict
func (bd *BencodedDict) AppendBencodedList(key string, value BencodedList) {
	bd.str_value += BencodeString(key) + value.StringValue()
}

//Add an key with an BencodedDict value to the dict
func (bd *BencodedDict) AppendBencodedDict(key string, value BencodedDict) {
	bd.str_value += BencodeString(key) + value.StringValue()
}
