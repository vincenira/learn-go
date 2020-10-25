package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/vincenira/learn-go/controllers"

	"github.com/vincenira/learn-go/models"
)

const (
	pi     = 3.14
	first  = iota
	second = iota + 6
	third  = 2 >> iota
)

const (
	fi = iota
	se
	th
)

// HTTPRequest datastructures for http request to demonstrate switch statement
type HTTPRequest struct {
	Method string
}

func main() {
	var i int
	i = 42
	var fl float32 = 42.32
	mobValue := 3
	localPi := 3.1453
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	primes := [5]int{1, 2, 3, 5, 7}
	//panic: assignment to entry in nil map mPanic
	//var mPanic map[string]int
	//mPanic["a"] = 1
	/*
		Solution to panic problem,
		initialize map using make function(or a map literal) before you can add any element
	*/
	m := make(map[string]int)
	m["ab"] = 1
	m["bb"] = 2
	m["cb"] = 3
	m["db"] = 5
	primesSecond := []int{1, 2, 3, 5, 7}
	fmt.Println("Hello from a module, Gophers")
	fmt.Println(pi, first, second, third, fi, se, th)
	fmt.Println(i, fl, mobValue, localPi, mobValue+3, float32(mobValue)+23.44)
	fmt.Println(arr)
	fmt.Println(arr[0:2:5])
	fmt.Println(primes)
	fmt.Println(primes[0:2])
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
	fmt.Println(m)
	fmt.Println(primesSecond)
	primesSecond = append(primesSecond, 11)
	fmt.Print("After Append: ")
	fmt.Println(primesSecond)
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get Request")
		// fallthrough (you used it will executed the below line DELETE)
	case "DELETE":
		println("Delete Request")
	case "POST":
		println("Post Request")
	case "PUT":
		println("Put Request")
	default:
		println("Unhandled method")
	}
	port := 3000
	port, err := startWebServer(port)
	fmt.Println(err)
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return port, errors.New("Something went wrong")
}
