package bencode

type BencodedDict struct {
	str_value string
}

func (bd *BencodedDict) StringValue() string {
	return "d" + bd.str_value + "e"
}

func (bd *BencodedDict) AppendInt(key string, value int) {
	bd.str_value += BencodeString(key) + BencodeInt(value)
}

func (bd *BencodedDict) AppendString(key string, value string) {
	bd.str_value += BencodeString(key) + BencodeString(value)
}

func (bd *BencodedDict) AppendBencodedList(key string, value BencodedList) {
	bd.str_value += BencodeString(key) + value.StringValue()
}

func (bd *BencodedDict) AppendBencodedDict(key string, value BencodedDict) {
	bd.str_value += BencodeString(key) + value.StringValue()
}
