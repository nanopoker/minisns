package utility

import (
	"fmt"
	"testing"
)

func TestGenerateIdentity(t *testing.T) {
	fmt.Println(GenerateIdentity())
}

func TestGenerateSalt(t *testing.T) {
	fmt.Println(GenerateSalt())
}

func TestGenerateUserid(t *testing.T) {
	fmt.Println(GenerateUserid())
}

func TestStringWithCharset(t *testing.T) {
	fmt.Println(StringWithCharset(USERIDLENGTH, USERIDSET))
}

func TestCrypto(t *testing.T) {
	password := "davidchen"
	salt := "VIXRRMNaxh"
	fmt.Println(Crypto(password, salt))
}
