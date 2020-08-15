package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var inp int
	// fmt.Scan(&inp)

	// var wg sync.WaitGroup
	// wg.Add(inp)

	// go func(){
	//     var inp2 int
	//     fmt.Scanf()
	// }
	fmt.Println(add([]int{1, 2, 3, 4, 5, 6}))
}

func add(val []int) int {
	n := len(val) - 1
	if n == 0 {
		return val[n]
	}
	return val[n] + add(val[0:(n)])
}

func toIntArr(strArr []string) intArr []int {
	n := len(val) - 1
	
	if n == 0 {
		y, _ := strconv.Atoi(val[n])
		intArr = append()
	}
	
}

func copy(val []int, strArr []string ) {
	n := len(val) - 1
	if n == 0 {
        
	}
	return val[n] + add(val[0:(n)])
}


// func main() {
// 	// reader := bufio.NewReader(os.Stdin)

// 	// name, _ := reader.ReadString('\n')
// 	// name = strings.Replace(name, "\n", "", -1)
// 	// a := strings.Split(name, " ")
// 	// fmt.Printf("%T", a)
// 	// // fmt.Println(a)
// 	a := []string{"1", "2", "3"}
// 	b := make([]int, len(a))
// 	copy(b, a)
// 	fmt.Println(b)
// }

// func copy(val []int, strArr []string) int {
// 	n := len(strArr) - 1
// 	if n == 0 {
// 		y, _ := strconv.Atoi(strArr[0])
// 		val[0] = y
// 		fmt.Println(y)
// 		return y
// 	}
// 	val[n] = copy(val[0:(n)], strArr[0:(n)])
// 	return copy(val[0:(n)], strArr[0:(n)])
// }
