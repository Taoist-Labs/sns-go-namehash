package sns_go_namehash

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

// Normalize returns a normalized sns name and a bool value.
//
// Support ASCII:
//
// 0-9: 48~57
// a-z: 97~122
// !  : 33
// $  : 36
// (  : 40
// )  : 41
// *  : 42
// +  : 43
// -  : 45
// .  : 46 (dot)
// _  : 95
func Normalize(name string) (bool, string) {
	lowerCase := strings.ToLower(name)

	chars := []rune(lowerCase)
	for i := 0; i < len(chars); i++ {
		char := lowerCase[i]
		if (char >= 97 && char <= 122) ||
			(char >= 48 && char <= 57) ||
			char == 33 ||
			char == 36 ||
			// char == 40 ||
			// char == 41 ||
			// char == 42 ||
			// char == 43 ||
			(char >= 40 && char <= 43) ||
			char == 45 ||
			char == 46 ||
			char == 95 {
			continue
		}

		return false, ""
	}

	return true, lowerCase
}

func Namehash(name string) string {
	node := "0000000000000000000000000000000000000000000000000000000000000000"

	s := strings.Split(name, ".")
	for i := len(s) - 1; i >= 0; i-- {
		hash := hex.EncodeToString(crypto.Keccak256([]byte(s[i])))
		t, _ := hex.DecodeString(fmt.Sprintf("%s%s", node, hash))
		node = hex.EncodeToString(crypto.Keccak256(t))
	}

	return fmt.Sprintf("0x%s", node)
}
