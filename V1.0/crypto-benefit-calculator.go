package main

import "fmt"

func main() {
	sumUsdt, currencyUsdtTry, currencyTryUsd, currencyUsdtUsd := getUserInput()
	resultUsdtLiraUsd := convertUsdtToLiraToUsd(sumUsdt, currencyUsdtTry, currencyTryUsd)
	resultUsdtUsd := convertUsdtToUsd(sumUsdt, currencyUsdtUsd)

	displayResults(resultUsdtLiraUsd, resultUsdtUsd)
	calculateBenefit(resultUsdtLiraUsd, resultUsdtUsd)
}

func getUserInput() (float32, float32, float32, float32) {
	var sumUsdt, currencyUsdtTry, currencyTryUsd, currencyUsdtUsd float32

	fmt.Print("Введите сумму USDT (USDT,cent): ")
	fmt.Scan(&sumUsdt)
	fmt.Print("Введите курс binance USDT to TRY : ")
	fmt.Scan(&currencyUsdtTry)
	fmt.Print("Введите банковский курс TRY to USD :")
	fmt.Scan(&currencyTryUsd)
	fmt.Print("Введите курс binance USDT to USD :")
	fmt.Scan(&currencyUsdtUsd)

	return sumUsdt, currencyUsdtTry, currencyTryUsd, currencyUsdtUsd
}

func convertUsdtToLiraToUsd(usdt, currencyUsdtTry, currencyTryUsd float32) float32 {
	return (usdt * currencyUsdtTry) / currencyTryUsd
}

func convertUsdtToUsd(usdt, currencyUsdtUsd float32) float32 {
	return usdt * currencyUsdtUsd
}

func displayResults(resultUsdtLiraUsd, resultUsdtUsd float32) {
	fmt.Println("result USDTtoTRYtoUSD", resultUsdtLiraUsd)
	fmt.Println("resultUSDTtoUSD: ", resultUsdtUsd)
}

func calculateBenefit(resultUsdtLiraUsd, resultUsdtUsd float32) {
	benefit := resultUsdtLiraUsd - resultUsdtUsd
	if benefit > 0 {
		fmt.Println("Выгоднее менять через лиру и зират банк выигрываешь: ", benefit)
	} else {
		fmt.Println("Выгоднее менять через бинанс USDT в USD на счет, выигрываешь: ", (-benefit))
	}
}
