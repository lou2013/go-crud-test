package userService

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id        int32  `json:"id"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
}

var users []User

func init() {
	users = make([]User, 0, 64)
}

func Add(user User) User {
	// var user User
	// json.Unmarshal(c.Body(), &user)
	id := users[len(users)-1].Id + 1
	users = append(users, User{Id: id, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email})

	return user
	// return c.Status(fiber.StatusCreated).JSON(user)
}

func Remove(id int32) (User, error) {
	// uId, err := strconv.Atoi(c.Params("userId"))
	// userId := int32(uId)
	// if err != nil {
	// return User{}, err
	// }
	var user User
	userIndex := -1
	for i, u := range users {
		if u.Id == id {
			user = u
			userIndex = i
		}
	}

	if userIndex == -1 {
		return User{}, fiber.ErrNotFound
		// return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	users = append(users[:userIndex], users[userIndex+1:]...)

	return user, nil
}

func GetById(id int32) (User, error) {

	var user User
	userIndex := -1
	for i, u := range users {
		if u.Id == id {
			user = u
			userIndex = i
		}
	}
	if userIndex == -1 {
		return User{}, fiber.ErrNotFound
	}
	return user, nil
	// return c.Status(fiber.StatusOK).JSON(user)
}

func GetAll() []User {
	return users
}

func Update(id int32, newUser User) (User, error) {

	var user User

	userIndex := -1

	for i, u := range users {
		if u.Id == id {
			user = u
			userIndex = i
		}
	}

	if userIndex == -1 {
		return User{}, fiber.ErrNotFound
	}

	users[userIndex] = User{Id: user.Id, FirstName: newUser.FirstName, LastName: newUser.LastName, Email: newUser.Email}
	newUser.Id = user.Id
	return newUser, nil
}
