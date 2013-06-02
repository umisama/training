package main

import "math/big"

func main(){
	// i have no good idea.

	parts := big.NewInt(0)

	for i:=1; i<=100; i++ {
		bigi := big.NewInt(int64(i))
		bigi.Mul( bigi, bigi )
		parts.Add(parts, bigi)
	}

	sum := 0
	for i:=1; i<=100; i++ {
		sum += i;
	}
	combine := big.NewInt(int64(sum))
	combine.Mul( combine, combine )

	answer := big.NewInt(0)
	answer.Sub( combine, parts )

	println(answer.String())
}
