package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func appStatus(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Up and running"))
}

func twoSidedPrimesHandler(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["number"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("number param is required"))
		return
	}

	if s, err := strconv.ParseFloat(keys[0], 32); err == nil {
		response := checkForTwoSidePrime(int(s))
		w.Write([]byte(strconv.FormatBool(response)))
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid number passed"))
	}

}

func insertIntoArray(original []int, position int, value int) []int {
	l := len(original)
	target := original
	if cap(original) == l {
		target = make([]int, l+1, l+10)
		copy(target, original[:position])
	} else {
		target = append(target, -1)
	}
	copy(target[position+1:], original[position:])
	target[position] = value
	return target
}
func checkForTwoSidePrime(s int) bool {
	arr := make([]int, 1)
	arr[0] = s
	length := countDigits(s)
	for i := 1; i < length; i++ {
		arr = insertIntoArray(arr, len(arr), s/int(math.Pow(10, float64(i))))
		arr = insertIntoArray(arr, len(arr), s%int(math.Pow(10, float64(i))))
	}
	fmt.Println(arr)
	for _, element := range arr {
		if !isPrime(element) {
			return false
		}
	}

	return true

}

func countDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {

	http.HandleFunc("/", appStatus)
	http.HandleFunc("/checkTwoSidedPrime", twoSidedPrimesHandler)

	http.ListenAndServe(":8090", nil)
}
