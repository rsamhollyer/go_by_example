package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "apple"
	s[1] = "banana"
	s[2] = "pear"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "plum")
	s = append(s, "cherries", "grapes")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("slice 1:", l)

	l = s[:5]
	fmt.Println("slice 2:", l)

	l = s[2:]
	fmt.Println("slice 3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare:", t)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

}
