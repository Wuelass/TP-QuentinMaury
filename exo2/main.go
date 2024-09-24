package main

func Ft_missing(nums []int) int {
	for i := len(nums) - 1; i >= 0; i-- {

	}
}

func main() {
	Ft_missing([]int{3, 1, 2})                   // resultat : 0
	Ft_missing([]int{0, 1})                      // resultat : 2
	Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) // resultat : 8
}
