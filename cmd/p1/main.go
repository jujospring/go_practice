package main

func getMonthlyPrice(tier string) int {
	var priceInPennies int

	if tier == "basic" {
		priceInPennies = 100 * 100
	} else if tier == "premium" {
		priceInPennies = 150 * 100
	} else if tier == "enterprise" {
		priceInPennies = 500 * 100
	}
	return priceInPennies
}
