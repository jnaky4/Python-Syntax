package main

import (
	"fmt"
	"github.com/yourbasic/bit"
	"math"
	"net"
)

func main() {
	//ipv4 is 32bits
	//prints [11111111 11111111 11111111 11111110]
	fmt.Printf("%b\n", net.CIDRMask(31, 32))
	//ipv6 is 128bits
	//prints ffffffffffffffff0000000000000000
	fmt.Printf("%s\n", net.CIDRMask(64, 128))

	println("bit shifting")
	fmt.Printf("%[1]d : %[1]b\n", bitShift(0b0001, 3))
	fmt.Printf("%[1]d : %[1]b\n", bitShift(0b0001, -1))

	println("get nth bit")
	fmt.Printf("%b\n", getNthBit(0b010, 2))
	println("is nth bit set")
	fmt.Printf("%t\n", isNthBitSet(0b101, 2))

	fmt.Println(sieveOfEratosthenes(20))

}

type Bits uint8

const (
	F0 Bits = 1 << iota
	F1
	F2
)

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }

func sieveOfEratosthenes(limit int) *bit.Set {
	sieve := bit.New().AddRange(2, limit)
	sqrtN := int(math.Sqrt(float64(limit)))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < limit; k += p {
			sieve.Delete(k)
		}
	}
	return sieve
}

// & and
// | or
// * XOR
// << left shift
// right shift

func bitShift(b byte, shift int) byte {
	switch {
	case shift < 0:
		return b >> (shift * -1)
	default:
		return b << shift
	}
}

// todo explain big endian & little endian
// todo explain bit mask cidr notation
// func bitMask(b byte, mask byte) byte{
//
// }
func getNthBit(b byte, n int) byte   { return (b>>n - 1) & 0b1 }
func isNthBitSet(b byte, n int) bool { return (b>>n-1)&0b1 == 0b1 }

//func setNthBit(b byte, n int, set bool)byte{ }
