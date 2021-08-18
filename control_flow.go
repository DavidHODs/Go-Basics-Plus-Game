package main

import "fmt"

// Covid19stats represnts the COVID19 status
type Covid19stats struct {
	Country string
	NumberOfCases int
	Safety SafetyMeasures
}

// SafetyMeasures are precautions to prevent COVID19 spread
type SafetyMeasures struct {
	First string
	Second string
	Third string
}

var n *int 

func main() {
	// nigeria := Covid19stats{
	// 	Country: "Nigeria",
	// 	NumberOfCases: 88,
	// 	Safety: SafetyMeasures{
	// 		First:  "",
	// 		Second: "",
	// 		Third:  "",
	// 	},

	// }
	// fmt.Println(nigeria)
	// fmt.Println(nigeria.Country)
	// nigeria.NumberOfCases = 12
	// if nigeria.NumberOfCases > 50 {
	// 	fmt.Println("Shut down the whole country")
	// } else {
	// 	fmt.Println("Still thriving")
	// }

	// nigeria := Covid19stats{}
	// nigeria.Country = "Nigeria"
	// nigeria.NumberOfCases = 109
	// fmt.Println(nigeria.Country)
	// SetCountry(&nigeria, "Ghana")
	// fmt.Println(nigeria.Country)
	n = new(int)
	*n = 40
	fmt.Println(*n)
	SetCountry(new(Covid19stats), "Ghana")
	fmt.Println(*n)
}

func SetCountry(cnt *Covid19stats, value string) {
	cnt.Country = value
	*n = 50
}