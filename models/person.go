package models

import "fmt"

type People map[string]float64

func (p *People) CheckSum() bool {
	var sum float64

	for _, val := range *p {
		sum += val
	}

	fmt.Println(sum)
	return sum == 0
}
