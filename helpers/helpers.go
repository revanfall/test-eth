package helpers

import (
	"strconv"
	"strings"
)

const round = float64(100000000000000000)

func hexNumberToInteger(hexaString string) string {
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
func ParseHexToDec(hex string) int64 {

	num, _ := strconv.ParseInt(hexNumberToInteger(hex), 16, 32)

	return num
}

func CountCommission(gasAmount, gasPrice string) float64 {
	gasAmountNum := float64(ParseHexToDec(gasAmount))
	gasPriceNum := float64(ParseHexToDec(gasPrice))

	return gasAmountNum * gasPriceNum / round

}

func CountValue(num string) float64 {
	res, err := strconv.ParseInt(hexNumberToInteger(num), 16, 64)
	if err != nil {
		return 0
	}
	return float64(res) / round
}
