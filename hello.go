package main

import (
	"fmt"
)

func main() {

	/* arreglos */
	arr := []int{5,4,3,2,1}
	arr = append(arr,0)
	arr[1] = 9
	//fmt.Println(arr)
	/****************************/

	/* maps(arreglos asociativos) */
	vertices := make(map[string]int)

	vertices["uno"]=1
	vertices["dos"]=2
	vertices["tres"]=3
	//fmt.Println(vertices["tres"])
	/****************************/


	arreglos()
}

func arreglos(){
	fmt.Println("arreglos")
}
