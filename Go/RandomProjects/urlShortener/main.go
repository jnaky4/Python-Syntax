package main

import (
	"fmt"
	"math"
	"math/rand"
)

//https://leetcode.com/problems/encode-and-decode-tinyurl/solutions/3479149/hold-my-beer-solution-0ms-beats-100/

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var collisionCount int
const codecLen = 4
var mapSize int

func main(){
	ac := Constructor()

	mapSize = int(math.Pow(float64(len(symbols)), float64(codecLen)))
	//numCodecs := mapSize/2
	numCodecs := 500_000

	for i := 0; i < numCodecs; i ++{
		ac.encode("https://www.google.com/" + string(rune(i)))
	}

	//print sorted keys
	//keys := make([]string, 0, len(ac.decodeMap))
	//for k := range ac.decodeMap {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys)
	//
	//for _, k := range keys {
	//	fmt.Println(k[len(k)-2:])
	//}

	//count collisions
	println("Map size: ", mapSize)
	fmt.Printf("collisions/total codecs: %d/%d=%.3f\n", collisionCount, numCodecs, float64(collisionCount)/float64(numCodecs))
}

type Codec struct {
	encodeMap map[string]string
	decodeMap map[string]string
}

func Constructor() Codec {
	return Codec{make(map[string]string, mapSize), make(map[string]string, mapSize)}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	res := "http://tinyurl.com/"
	exists := true

	for exists{
		exists = false
		res = "http://tinyurl.com/"
		if _, ok := c.encodeMap[longUrl]; !ok {
			for i := 0; i <= codecLen-1; i++ {
				res += string(symbols[rand.Intn(len(symbols))])
			}

			if _, ok := c.decodeMap[res]; !ok {
				//println(res)
				c.encodeMap[longUrl] = res
				c.decodeMap[res] = longUrl
				return c.encodeMap[longUrl]
			} else {
				//println("hash already exists " + res[len(res)-2:])
				collisionCount++
				exists = true
			}
		}
	}
	return c.encodeMap[longUrl]
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.decodeMap[shortUrl]
}
