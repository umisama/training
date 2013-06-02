package main

func isPalindromic( num int )bool {
	// separate with digit
	numlist := make([]int, 0, 1000)
	for ; num > 0; num = int(num/10) { numlist = append(numlist, num%10) }

	// length must be even.
	if len(numlist) % 2 != 0 { return false; }

	// check
	for i:=0; i<int(len(numlist)/2); i++ {
		if numlist[i] != numlist[len(numlist)-1-i] { return false }
	}

	return true
}

func main() {
	max := 0
	for i:=0; i < 1000; i++ {
		for j:=0; j < 1000; j++ {
			if isPalindromic( i*j ) && max < i*j {
				max = i*j
			}
		}
	}
	print(max)
}
