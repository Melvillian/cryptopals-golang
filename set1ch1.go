package main

import (
	"github.com/Melvillian/cryptopals/set1"
	"fmt"
)

func main() {
	str := set1.Hex2base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d");
	fmt.Println(str)
}
