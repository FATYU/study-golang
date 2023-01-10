package main

import (
	"fmt"
	"math/big"
)

func main() {

	// big1, _ := new(big.Int).SetString("1837947885749375974239759283758473295742395000", 10)
	// fmt.Println("big1 is: ", big1.BitLen())

	// big2 := big1.Uint64()
	// fmt.Println("big2 is: ", big2)
	//
	// bigbig := big.NewInt(0).Lsh(big.NewInt(0xFFFFFFFFFFFFFF), 16)
	// fmt.Println(bigbig)
	//
	// bigbig =
	// 	big.NewInt(0xFFFF)
	// fmt.Println(bigbig)
	//
	// bigbig = new(big.Int).Add(
	// 	big.NewInt(0).Lsh(big.NewInt(0xFFFFFFFFFFFFFF), 16),
	// 	big.NewInt(0xFFFF),
	// )
	// fmt.Println(bigbig)

	// 4722366482869645148160
	// 72057594037993470

	// fmt.Println(133)
	// for n := 1; n < 1000; n++ {
	// 	fmt.Println((((n + 7) & -8) >> 3) == ((n + 7) / 8))
	// }

	fmt.Println(uint64(^big.Word(0)))
	fmt.Println(uint64(big.Word(0)))

	const wordBytes = (32 << (uint64(^big.Word(0)) >> 63)) / 8

	fmt.Println(32 << (uint64(^big.Word(0)) >> 63))
	bigword := big.Word(281474976710417)
	fmt.Println(byte(bigword))
}
