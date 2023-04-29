package main

import "fmt"

func main() {
	var sumUsdt float32
	var currencyUsdtTry float32
	var currencyTryUsd float32
	var currencyUsdtUsd float32
	var resultUsdtLiraUsd float32
	var resultUsdtUsd float32

	fmt.Print("Введите сумму USDT (USDT,cent): ")
	fmt.Scan(&sumUsdt)
	fmt.Print("Введите курс binance USDT to TRY : ")
	fmt.Scan(&currencyUsdtTry)
	fmt.Print("Введите банковский курс TRY to USD :")
	fmt.Scan(&currencyTryUsd)
	fmt.Print("Введите курс binance USDT to USD :")
	fmt.Scan(&currencyUsdtUsd)

	resultUsdtLiraUsd = UsdtLiraUsd(sumUsdt, currencyUsdtTry, currencyTryUsd)
	fmt.Println("result USDTtoTRYtoUSD", resultUsdtLiraUsd)

	resultUsdtUsd = UsdtUsd(sumUsdt, currencyUsdtUsd)
	fmt.Println("resultUSDTtoUSD: ", resultUsdtUsd)

	fmt.Println(Benefit(resultUsdtLiraUsd, resultUsdtUsd))

}

func Benefit(resultUsdtLiraUsd float32, resultUsdtUsd float32) float32 {
	benefit := resultUsdtLiraUsd - resultUsdtUsd
	if benefit > 0 {
		fmt.Println("Выгоднее менять через лиру и зират банк выигрываешь: ", benefit)
	} else {
		fmt.Println("Выгоднее менять через бинанс USDT в USD на счет, выигрываешь: ", (-benefit))
		return -benefit
	}
	return benefit
}

func UsdtLiraUsd(usdt float32, currencyUsdtTry float32, currencyTryUsd float32) float32 {

	return (usdt * currencyUsdtTry) / currencyTryUsd
}

func UsdtUsd(usdt float32, currencyUsdtUsd float32) float32 {
	return usdt * currencyUsdtUsd
}

/*
500
20.61 usdt lira 10305
20.982 lira to USD ziraat //// fake rates 18   491.13
0.961 USDT to USD binance		480.5

*/
