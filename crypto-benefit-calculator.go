package main

import "fmt"

func main() {
	var sumUsdt float64
	var currencyUsdtTry float64
	var currencyTryUsd float64
	var currencyUsdtUsd float64
	var resultUsdtLiraUsd float64
	var resultUsdtUsd float64
	fmt.Print("Введите сумму USDT (USDT,cent): ")
	fmt.Scan(&sumUsdt)
	fmt.Print("Введите курс binance USDT to TRY : ")
	fmt.Scan(&currencyUsdtTry)
	fmt.Print("Введите банковский курс TRY to USD :")
	fmt.Scan(&currencyTryUsd)
	fmt.Print("Введите курс binance USDT to USD :")
	fmt.Scan(&currencyUsdtUsd)
	resultUsdtLiraUsd = UsdtLiraUsd(sumUsdt, currencyTryUsd, currencyTryUsd)
	resultUsdtUsd = UsdtUsd(sumUsdt, currencyUsdtUsd)
	fmt.Println(Benefit(resultUsdtLiraUsd, resultUsdtUsd))

}

func Benefit(resultUsdtLiraUsd float64, resultUsdtUsd float64) float64 {
	benefit := resultUsdtLiraUsd - resultUsdtUsd
	if benefit > 0 {
		fmt.Println("Выгоднее менять через лиру и зират банк выигрываешь: ", benefit)
	} else {
		fmt.Println("Выгоднее менять через бинанс USDT в USD на счет, выигрываешь: ", (-benefit))
		return -benefit
	}
	return benefit
}

func UsdtLiraUsd(usdt float64, currencyUsdtTry float64, currencyTryUsd float64) float64 {

	binanceExchangeLira := usdt * currencyUsdtTry

	return binanceExchangeLira / currencyTryUsd
}

func UsdtUsd(usdt float64, currencyUsdtUsd float64) float64 {
	return usdt * currencyUsdtUsd
}

/*
500
20.61 usdt lira
20.982 lira to USD ziraat //// fake rates 18
0.961 USDT to USD binance

*/
