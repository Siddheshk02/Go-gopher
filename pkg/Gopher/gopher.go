package Gopher

//import (
//"strconv"
//)

func Moon(input float64) (result float64) {
	result = (input / 9.81) * 1.622
	return result
}

func Mercury(input float64) (result float64) {
	result = (input / 9.81) * 3.73

	return result
}

func Venus(input float64) (result float64) {
	result = (input / 9.81) * 8.93

	return result
}

func Mars(input float64) (result float64) {
	result = (input / 9.81) * 3.73

	return result
}

func Jupiter(input float64) (result float64) {
	result = (input / 9.81) * 22.96

	return result
}

func Saturn(input float64) (result float64) {
	result = (input / 9.81) * 10.40

	return result
}

func Uranus(input float64) (result float64) {
	result = (input / 9.81) * 9.03

	return result
}

func Neptune(input float64) (result float64) {
	result = (input / 9.81) * 11.67

	return result
}

func Pluto(input float64) (result float64) {
	result = (input / 9.81) * 0.68

	return result
}
