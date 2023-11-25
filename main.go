package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"

	"github.com/gruntwork-io/terratest/modules/random"
)

type User struct {
	Name string
	Age  int
}
type Users []User

func main() {
	// TODO: Implement main
	var u Users
	u = create_user(10, 30, 50)
	l := make_list(u)
	create_csv(l)
}

func create_user(n int, range_min, range_max int) Users {
	// TODO: Implement create_user
	var u Users
	for i := 0; i < n; i++ {
		user := User{
			Name: random_name([]string{"John", "Jane", "Bob", "Alice"}),
			Age:  random_number(range_min, range_max),
		}
		u = append(u, user)
	}

	return u
}

func random_name(s []string) string {
	// TODO: Implement random_name

	len_of_s := len(s)
	// TODO: Implement random_name
	return s[rand.Intn(len_of_s)]
}

func random_number(i, y int) int {
	// TODO: Implement random_number
	// rand 1 -6
	return random.Random(i, y)
}

func make_list(u Users) [][]string {
	l := [][]string{}
	for _, u_tmp := range u {
		l = append(l, []string{u_tmp.Name, fmt.Sprint(u_tmp.Age)})
	}
	return l
}

func create_csv(l [][]string) {
	// TODO: Implement create_csv

	file, err := os.OpenFile("csv.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	defer writer.Flush()

	reader := csv.NewReader(file)

	writer.Write([]string{"Name", "Age"})

	reader.FieldsPerRecord = -1
	records := l
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	fmt.Println("Error reading CSV data:", err)
	// 	return
	// }
	//
	// for _, record := range records {
	// 	fmt.Println(record)
	// }
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
