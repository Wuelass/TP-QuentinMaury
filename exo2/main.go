package main

import "fmt"

func Ft_missing(nums []int) int {

	n := len(nums)
	total := (n * (n + 1)) / 2
	somme := 0

	for _, num := range nums {
		somme += num
	}

	fmt.Println(total - somme)
	return total - somme
}

func main() {
	Ft_missing([]int{3, 1, 2})                   // resultat : 0
	Ft_missing([]int{0, 1})                      // resultat : 2
	Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) // resultat : 8
}
