package main

import "fmt"

func main() {
	var users0 []string
	var users = []string{"Tom", "Alice", "Kate"}
	fmt.Println(users0, users)

	// change, for
	users[2] = "Katherine"
	fmt.Println(users)

	for _, value := range users {
		fmt.Println(value)
	}

	// make
	var users2 []string = make([]string, 3)
	fmt.Println(users2)
	users2[0] = "Alice"
	users2[1] = "Bob"
	users2[2] = "Charlie"
	// users2[3] = "David" // err
	fmt.Println(users2)

	// append
	users2 = append(users, "David")
	fmt.Println(users2)
	users22 := []string{"Ee", "Ff"}
	fmt.Println(append(users2, users22...))

	//
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(nums[0:0])
	fmt.Println(nums[0:])
	fmt.Println(nums[:0])
	fmt.Println(nums[2:6])
	fmt.Println(nums[:4])
	fmt.Println(nums[3:])

	// delete
	n := 3
	numsDeleted := append(nums[:n], nums[n+1:]...)
	fmt.Println(numsDeleted)
}
