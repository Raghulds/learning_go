package gofundamentals

/*
Strings are Read only slice of bytes without cap
Zero value for a string is "", not nil

byte -> uint8
rune -> uint32
*/
import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
Banner('Go', 6)

Output: String centered to the width param

	Go

------

Banner('Go🐱', 7)
len(text) is 6!
Why? 🐱 is a UTF-8 encoded which is of 4 byte. So, it's 6 byte
*/

/*
Strings, runes, bytes
*/
func Banner(text string, width int) {
	fmt.Println("Length - ", len(text))
	// padding := (width - len(text)) / 2 // len returns length by utf bytes count
	padding := (width - utf8.RuneCountInString(text)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", width))

	// bytes
	fmt.Println([]byte(text))
	fmt.Println(text[3])
	fmt.Println(text[4])

	fmt.Println(utf8.FullRune([]byte(text)))
}

/*
ASCII - Single byte - 7 bits - 0 to 127
Unicode - Big db - Sequence of bytes
UTF-8 - Encoding format for Unicode - 8 bit - One to Four One byte (8-bit) code units
*/

/*
Strings are UTF-8 encoded
len, s[] accesses as byte (uint8)
for loop accesses as rune (int32)
*/

/*
Unicode → assigns numbers to characters
UTF-8   → encodes those numbers into bytes
string  → UTF-8 encoded bytes
byte    → raw 8-bit chunk
rune    → Unicode code point

Bytes for machines, runes for humans
*/
