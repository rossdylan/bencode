package bencode

import (
	"fmt"
	"testing"
)

// Test string encoding
func TestBencodeString(t *testing.T) {
	result := BencodeString("testing")
	expectedResult := "7:testing"
	if result != expectedResult {
		t.Error(fmt.Sprintf("Result did not match expected result: %s != %s", result, expectedResult))
	}
	fmt.Println("EncodedString: " + result)
}

// Test int encoding
func TestBencodeInt(t *testing.T) {
	result := BencodeInt(23)
	expectedResult := "i23e"
	if result != expectedResult {
		t.Error(fmt.Sprintf("Result did not match expected result: %s != %s", result, expectedResult))
	}
	fmt.Println("EncodedInt: " + result)
}

// Test BencodedLists
func TestBencodeList(t *testing.T) {
	list := BencodedList{}
	list.AppendInt(20)
	list.AppendString("HerpDerp")
	result := list.StringValue()
	expectedResult := "li20e8:HerpDerpe"
	if result != expectedResult {
		t.Error(fmt.Sprintf("Result did not match expected result: %s != %s", result, expectedResult))
	}
	fmt.Println("EncodedList: " + list.StringValue())
}

// Test BencodedDict
func TestBencodedDict(t *testing.T) {
	dict := BencodedDict{}
	testList := BencodedList{}
	testList.AppendInt(6969)
	dict.AppendInt("testInt", 1337)
	dict.AppendString("testString", "test")
	dict.AppendBencodedList("testDict", testList)
	fmt.Println(dict.StringValue())

}
