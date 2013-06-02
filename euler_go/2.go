package main

const MAX_VALUE = 4000000

func main(){
	sum := 0
	for p, n := 1, 2; n < MAX_VALUE; p, n = n, p+n {
		if n % 2 == 0 {
			sum += n
		}
	}
	print(sum)
}
