package main

func main() {

aloop:
	for b:=1; b<1000; b++ {
		for a:=1; a<b; a++ {
			c := 1000 - a - b
			if c < 0 { continue aloop }
			if ( c*c ) == ( a*a + b*b ) {
				println( a*b*c)
			}
		}
	}
}
