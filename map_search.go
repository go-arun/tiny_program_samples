package main

import "fmt"

func main(){
	m := make(map[int]string)
	m[1] = "Arun"
	m[2] = "Kumar"
	m[3] = "lal"
	m[4] = "Kiran"
	m[5] = "Gowri"
	m[6] = "Hari"
	
	fmt.Println(m)
//Search with key
	v,found := m[4]
	fmt.Println(v,found)
//No such direct option to search with Value 
	search_val := "Kiran"
	for key,val := range m {
		if val == search_val {
			fmt.Println("Yes this values is in key postion :",key)
		}
		
	}
}
