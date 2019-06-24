package main

import (
	"fmt"
)

func ArraysInGolang() {

	a:= [3][3]uint32{{1,2,3},{4,5,6},{7,8,9}}
	b:= [3][3]uint32{{9,8,7},{6,5,4},{3,2,1}}

	fmt.Println("Input Arrays: ")
	fmt.Println( a)
	fmt.Println( b)
	value := sum(a)
	fmt.Println("\nSum of all element in Array 1:",value)
	value = sum(b)
	fmt.Println("\nSum of all element in Array 2:",value)

	ret:= add(a,b)
	fmt.Println("\nAdd both array :",ret)

	ret= mul(a,b)
	fmt.Println("\nMultiply both array: ",ret)

}

func sum(array [3][3]uint32) ( sum uint32){
	sum = 0
	for _,i:= range array{
		for _,j:= range i{
			sum = sum+j
		}
	}
	return
}

func add(array1,array2 [3][3]uint32) (array3 [3][3]uint32) {
	for i:=0; i<3; i++{
		for j:=0;j<3;j++{
			array3[i][j] = array1[i][j] + array2[i][j]
		}
	}

	return
}
func mul(array1,array2 [3][3]uint32) (array3 [3][3]uint32) {
	for i:=0; i<3; i++{
		for j:=0;j<3;j++{
			array3[i][j] = array1[i][j] * array2[i][j]
		}
	}
	return
}