package main

import "fmt"

func Ft_profit(prices []int) int {

	min := prices[0]
	maxP := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if maxP < prices[i]-min {
			maxP = prices[i] - min
		}
	}
	fmt.Println(maxP)
	return 0
}

func main() {
	Ft_profit([]int{7, 1, 5, 3, 6, 4}) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	Ft_profit([]int{7, 6, 4, 3, 1}) // resultat : 0
}
