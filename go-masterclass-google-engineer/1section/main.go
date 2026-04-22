package main

import "fmt"

//enums
const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Loglevel int

const (
	LogError Loglevel = iota
	LogWarn
	LogInfo
	LogDebug
	LogFatal
)

//constants

const (
	Host    = "localhost"
	Port    = 8080
	Timeout = 30
	User    = "root"
)

var (
	isRunning bool = true
)

func main() {
	var greeting string // zero value of string is empty string ""
	greeting = "Ram, Ram Bhai saryane !!"
	fmt.Println(greeting)

	var counter int // zero value of int is 0
	counter = 42
	fmt.Println(counter)

	var isActive bool // zero value of bool is false
	isActive = true
	fmt.Println(isActive)

	var fname, lname string // zero value of string is empty string ""
	fname = "Shubham"
	lname = "Joshi"
	fmt.Println(fname, lname)

	email := "shubham.joshi1507@gmail.com" // short variable declaration, type is inferred as string
	fmt.Println(email)

	age := 32 // short variable declaration, type is inferred as int
	fmt.Println(age)

	year := 2026 // short variable declaration, type is inferred as int
	fmt.Println(year)

	// Using constants
	fmt.Printf("Connecting to %s:%d with timeout %d seconds as user %s\n", Host, Port, Timeout, User)

	AppName := "Go Masterclass"
	fmt.Printf("Welcome to %s!\n", AppName)

	const Pi float64 = 3.14159
	fmt.Printf("The value of Pi is approximately %.5f\n", Pi)

	// Using enums
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

}
