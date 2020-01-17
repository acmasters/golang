package matematica

import (
	"fmt"
	"strconv"
)

func Media(numeros ...float64) float64 {
	total := 0.000
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArrendodada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArrendodada
}
