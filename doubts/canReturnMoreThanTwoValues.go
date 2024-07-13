package main

import "fmt"

func moreThanTwoReturn(a int, b int, c int, d int) (int, int, int, float32, error) {

	var sum int = a + b + c + d
	var subs int = a - b - c - d
	var mul int = a * b * c * d
	var divide float32 = float32(a+b+c+d) / 4

	return sum, subs, mul, divide, nil
}
func moreThanTwoReturn2(a int, b int) (int, int, int, float32, error) {

	var sum int = a + b
	var subs int = a - b
	var mul int = a * b
	var divide float32 = float32(a) / float32(b)

	return sum, subs, mul, divide, nil
}
func main() {
	a, b, c, d, err := moreThanTwoReturn(4, 2, 5, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a, b, c, d)

	a, b, c, d, _ = moreThanTwoReturn2(4, 2)
	fmt.Println(a, b, c, d)

}
