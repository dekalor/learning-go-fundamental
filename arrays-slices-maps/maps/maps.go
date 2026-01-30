package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Print(websites["AWS"])

	websites["LinkedIn"] = "linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
