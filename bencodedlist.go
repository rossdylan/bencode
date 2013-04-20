package bencode

type BencodedList struct {
    str_value string
}

func (bl *BencodedList) StringValue() string {
    return "l" + bl.str_value + "e"
}

func (bl *BencodedList) AppendInt(i int) {
    bl.str_value += BencodeInt(i)
}

func (bl *BencodedList) AppendString(str string) {
    bl.str_value += BencodeString(str)
}

func (bl *BencodedList) AppendBencodedList(l BencodedList) {
    bl.str_value += l.StringValue()
}

func (bl *BencodedList) AppendBencodedDict(d BencodedDict) {
    bl.str_value += d.StringValue()
}
