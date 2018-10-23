package main

import (
	"fmt"
)

func mult(x , y ,z int ) (int,int,error) {
	return x+y+z,x/z,nil
}
 
func main() {

	t:= []string{"g","h","i"}
	fmt.Println("dcl:",t)
	s:=t[0:2]
	
	fmt.Println("dcls",s,len(s),cap(s))
	s=append(s,"w")
	fmt.Println("dcls",s)



//	fmt.Println("Hello, playground")

//	a,b,c:=mult(42, 11,13)

//	fmt.Println(a,b,c)	
}

