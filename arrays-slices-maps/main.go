package main

import "fmt"

type floatMap map[string]float64

func main() {
	userNames := make([]string, 2)

	userNames[0] = "Ena"
	userNames = append(userNames, "Ade")
	userNames = append(userNames, "Dekalor")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4
	courseRatings["node"] = 4

	courseRatings.output()

	for idx, val := range userNames {
		fmt.Println("idx: ", idx)
		fmt.Println("val: ", val)
	}

	for key, val := range courseRatings {
		fmt.Println("key: ", key)
		fmt.Println("val: ", val)
	}
}

func (m floatMap) output() {
	fmt.Println(m)
}
