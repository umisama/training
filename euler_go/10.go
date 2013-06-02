package main

import "./algorithms"

func main() {
	sum			:= int64(0)
	prime_list	:= algorithms.Erastosthenes(2000000)
	for _,i := range prime_list { sum += i }
	println(sum)
}
