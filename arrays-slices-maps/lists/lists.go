package lists

import "fmt"

type product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Badminton", "Gym", "Sleep"}
	fmt.Println(hobbies)
	fmt.Println("First hobby: ", hobbies[0])
	optional_hobbies := hobbies[1:3]
	fmt.Println("Optional hobby: ", optional_hobbies)

	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	courseGoals := []string{"Go Developer", "Mastering Go"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Mastering Go Exactly"
	courseGoals = append(courseGoals, "Go jos")
	fmt.Println(courseGoals)

	products := []product{
		{
			id:    "1",
			title: "first product",
			price: 100000,
		},
		{
			id:    "2",
			title: "second product",
			price: 50000,
		},
	}

	fmt.Println(products)

	newProduct := product{
		id:    "3",
		title: "third product",
		price: 55000,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

// func main() {
// 	prices := []float64{10.99, 20.0}
// 	fmt.Println(prices[0:1])

// 	prices[1] = 9.99

// 	prices = append(prices, 5.99)
// 	fmt.Println(prices)
// }

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}

// 	prices := [4]float64{10.99, 4.55, 30.99, 20.0}
// 	fmt.Println(prices[3])
// 	fmt.Println(productNames)
// }
