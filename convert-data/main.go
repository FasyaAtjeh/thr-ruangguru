package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	userData := strings.Split(data, ",")
	age, _ := strconv.Atoi(userData[1])
	return User{
		Name:    userData[0],
		Age:     age,
		Address: userData[2],
	}
}
func main() {
	data := "Edo,20,Jakarta"
	user := ConvertData(data)
	fmt.Printf("{ name: %s, age: %d, address: %s }\n", user.Name, user.Age, user.Address)
}
