package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// low case plz
var targetStart = "11cbd"
var targetEnd = "893d4"

const threads = 16

var areacodelist = []int{130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 145, 147, 149, 150,
	151, 152, 153, 155, 156, 157, 158,
	159, 162, 165, 166, 167, 170, 171,
	172, 173, 174, 175, 176, 177, 178,
	180, 181, 182, 183, 184, 185, 186,
	187, 188, 189, 190, 191, 192, 193,
	195, 196, 197, 198, 199}

func main() {
	c := make(chan struct{}, threads)
	for _, areacode := range areacodelist {
		c <- struct{}{}
		go func(areacode int) {
			fmt.Println("start", areacode)
			defer func() { <-c }()
			for i := 0; i < 100000000; i++ {
				phone := fmt.Sprintf("86%d%08d", areacode, i)
				targetTest := sha256.Sum256([]byte(phone))
				var hexTest = make([]byte, hex.EncodedLen(len(targetTest)))
				hex.Encode(hexTest, targetTest[:])
				if targetStart == string(hexTest[0:5]) && targetEnd == string(hexTest[len(hexTest)-5:]) {
					fmt.Println(phone)
				}
			}
		}(areacode)
	}
}
