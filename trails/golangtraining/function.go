package main

import (
	"fmt"
)

func mult(x , y ,z int ) (int,int,error) {
	return x+y+z,x/z,nil
}
 
func main() {
	fmt.Println("Hello, playground")

	a,b,c:=mult(42, 11,13)

	fmt.Println(a,b,c)	
}

