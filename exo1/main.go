package main

import "fmt"

func Ft_coin(coins []int, amount int) int {
	reverse_coins := []int{}
	resulte := 0

	for i := len(coins) - 1; i >= 0; i-- {
		reverse_coins = append(reverse_coins, coins[i])
	}

	for i := 0; i < len(reverse_coins); i++ {
		if amount >= reverse_coins[i] {
			resulte += amount / reverse_coins[i]
			amount = amount % reverse_coins[i]
		}
	}

	if amount != 0 {
		fmt.Println("-1")
		return -1
	}
	fmt.Println(resulte)
	return resulte
}

func main() {
	Ft_coin([]int{1, 2, 5}, 11) // resultat : 3 car (11 == 5 + 5 + 1)
	Ft_coin([]int{2}, 3)        // resultat : -1
	Ft_coin([]int{1}, 0)        // resultat : 0
}
