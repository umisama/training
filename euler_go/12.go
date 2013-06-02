package main

import "./algorithms"

func getNumOfDivisor( primes []int64, triangle int64 ) int {
	// n := X^a * Y^b * Z^c (X,Y,Z : prime number)
	// divisors of n is (a+1)*(b+1)*(c+1)

	divisor := 1
	for {
		for _, prime := range primes {
			n := 1
			for ; triangle % prime == 0; triangle /= prime {
				n++
			}
			divisor *= n
		}
		if triangle == 1 {
			return divisor
		}
	}
	return 0
}

func main(){

	primes := algorithms.Erastosthenes( 1000000 )
	println(primes)

	triangle := int64(0)
	for i:=int64(1);  ; i++ {
		triangle += i

		divisor := getNumOfDivisor( primes, triangle )

		if divisor >= 500 {
			break
		}
	}

	println( triangle )
}
