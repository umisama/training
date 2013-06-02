package main

const TARGET = 10001

func getMax( list []int64 )( max int64 ) {
	for _,i := range list {
		if max < i {
			max = i
	}}
	return
}

func Erastosthenes( max int64 )( ret []int64 ) {
	ret = make([]int64, 0, 1000)

	sub_list := make( []int64, 0, 1000 )

	// 1st step : create table
	for i:=int64(2); i <= max; i++ {
		sub_list = append(sub_list, i)
	}

	p := 0
	for {
		// 2nd step : top is one of prime.
		new_prime := sub_list[p]
		for i := p+1; i < len(sub_list); i++ {
			if( sub_list[i] % new_prime == 0 ) {
				sub_list[i] = -1
			}
		}

		// 3rd step : break loop if new prime^2 is bigger than biggest.
		if (new_prime*new_prime) > getMax( sub_list ) {
			break
		}

		// ( 4th step : point to next prime. )
		for i:=p+1; i<len(sub_list); i++ {
			if( sub_list[i] != -1 ) {
				p = i
				break
			}
		}
	}

	for _,i := range sub_list {
		if i != -1 {
			ret = append(ret, i )
		}
	}

	return
}

func main() {
	primes := Erastosthenes(1000000)
	print(primes[TARGET - 1])
	return
}
