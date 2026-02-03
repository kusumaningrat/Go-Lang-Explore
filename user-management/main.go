/*

This is simple user management system using go with in-memory as database. There are several functions here such as:
AddUser()
ListAllUsers()
GetUserByName()
UpdateUser()
DeleteUser()

*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	ID      string
	Name    string
	Age     int
	Address string
}

type UserUpdateInput struct {
	Name    *string
	Age     *int
	Address *string
}

type APIResponse struct {
	Message string
	Data    interface{} `json:"data"`
}

var users []User

func PrettyJson(data interface{}) string {
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(result)
}

func AddUser(name string, age int, address string) (User, error) {
	// Validation if user has empty input name
	if name == "" {
		// return errors.New("name cannot be empty"), nil
		fmt.Println("name cannot be empty")
	}

	// Generate ID with uuid
	userId := uuid.New()
	formatedId := strings.Split(userId.String(), "-")[0]

	// Declare data structure that will be store to slice
	user := User{
		ID:      formatedId,
		Name:    name,
		Age:     age,
		Address: address,
	}

	// Save data to slice
	users = append(users, user)
	// result := UserResponse{
	// 	Message: "User created successfully",
	// 	Data:    users,
	// }
	// fmt.Println(PrettyJson(result))
	// Return created user otherwise nill
	return user, nil
}

func ListAllUsers() ([]User, error) {
	return users, nil
}

func GetUserByName(name string) (*User, error) {

	//  Loop through users data
	for i, user := range users {
		// Check if parameter name is match with  user.Name in database (users slice)
		if name == user.Name {
			// Return matching user
			return &users[i], nil
		}
	}

	return nil, nil
}

func UpdateUser(userId string, input UserUpdateInput) (User, error) {
	// found variable for checking if name is match or not
	for i := range users {
		// Check if parameter userId is match with users.id in database
		if users[i].ID == userId {
			if input.Name != nil {
				users[i].Name = *input.Name
			}
			if input.Age != nil {
				users[i].Age = *input.Age
			}
			if input.Name != nil {
				users[i].Name = *input.Name
			}

			return users[i], nil
		}
	}

	return User{}, fmt.Errorf("user not found")
}

func DeleteUser(userId string) error {
	for i, user := range users {
		if user.ID == userId {
			// Delete using slice index, [:i] is all records before i and [i+1] is all record after i. (i) will be ignored and merge selected
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User not found")
}

func FormattedResponse(message string, data any) APIResponse {
	return APIResponse{
		Message: message,
		Data:    data,
	}
}

func main() {

	// Add user

	fmt.Println("==========Add users==========")

	// Add sample users
	aliceUser, err := AddUser("Alice", 24, "New York")
	if err != nil {
		fmt.Println("Error while adding new user: ", err)
	}
	fmt.Println(
		PrettyJson(FormattedResponse("User added successfully", aliceUser)),
	)

	bobUser, err := AddUser("Bob", 24, "Chichago")
	if err != nil {
		fmt.Println("Error while adding new user: ", err)
	}
	fmt.Println(
		PrettyJson(FormattedResponse("User added successfully", bobUser)),
	)

	// if err := AddUser("Alice", 24, "New York"); err != nil {
	// 	fmt.Println("Error while adding new user: ", err)
	// }
	// if err := AddUser("Bob", 24, "Chichago"); err != nil {
	// 	fmt.Println("Error while adding new user: ", err)
	// }

	fmt.Println()
	fmt.Println("==========List all users==========")

	// List all users
	user, err := ListAllUsers()
	if err != nil {
		fmt.Println("Error while loading users data", err)
	}
	fmt.Println(
		PrettyJson(FormattedResponse("User loaded successfully", user)),
	)
	// fmt.Println(PrettyJson(result))

	// Get one user
	fmt.Println()
	fmt.Println("==========Get One user==========")

	// Get user bob
	// if err := GetUserByName("Bob"); err != nil {
	// 	fmt.Println("Cannot get user", err)
	// }

	bob, err := GetUserByName("Bob")
	if err != nil {
		fmt.Println("Error while loading user data", err)
		return
	}

	if bob == nil {
		fmt.Println("User not found")
		return
	}

	// result2 := UserResponse{
	// 	Message: "User loaded successfully",
	// 	Data:    getOneUser,
	// }
	// fmt.Println(PrettyJson(result2))
	fmt.Println(
		PrettyJson(FormattedResponse("User loaded successfully", bob)),
	)

	// Update user bob

	fmt.Println()
	fmt.Println("==========Update user==========")
	id := bob.ID
	name := "Bob Update"
	age := 25

	input := UserUpdateInput{
		Name: &name,
		Age:  &age,
	}

	updateUser, err := UpdateUser(id, input)
	if err != nil {
		fmt.Println("Error while updating user", err)
	}

	fmt.Println(
		PrettyJson(FormattedResponse("User updated successfully", updateUser)),
	)

	fmt.Println(
		PrettyJson(FormattedResponse("User loaded successfully", user)),
	)

	fmt.Println()
	fmt.Println("==========Delete user==========")
	aliceId := aliceUser.ID

	isDeleted := DeleteUser(aliceId)
	if isDeleted != nil {
		fmt.Println("Error while deleting user", err)
	}
	fmt.Println("User deleted successfully")

	fmt.Println(
		PrettyJson(FormattedResponse("New user record", users)),
	)
}
