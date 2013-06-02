package main

const (
	MIN = 1
	MAX = 20
)

func main(){
	// i have no good idea.
	i := 0
big	:for i=1 ; ; i++ {
		for j:=MIN; j<=MAX; j++ {
			if(i%j != 0 ) {
				continue big
			}
			if(j == MAX ) {
				break big
			}
		}
	}
	print(i)
}
