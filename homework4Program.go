package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type Database struct {
	Users []User
}

const (
	CommandHelp   = "help"
	CommandList   = "list"
	CommandAdd    = "add"
	CommandUpdate = "update"
	CommandDelete = "delete"
)

func main() {
	db := Database{}

	fmt.Println("Ласкаво просимо до консольної програми - бази даних користувачів!")
	fmt.Println("Введіть 'help' для перегляду доступних команд.")

	for {
		fmt.Print("\nВведіть команду: ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case CommandHelp:
			printHelp()
		case CommandList:
			listUsers(db)
		case CommandAdd:
			addUser(&db)
		case CommandUpdate:
			updateUser(&db)
		case CommandDelete:
			deleteUser(&db)
		default:
			fmt.Println("Невідома команда. Введіть 'help' для отримання списку команд.")
		}
	}
}

func printHelp() {
	fmt.Println("Доступні команди:")
	fmt.Println("  help   - переглянути список команд")
	fmt.Println("  list   - переглянути список користувачів")
	fmt.Println("  add    - створити нового користувача")
	fmt.Println("  update - оновити дані користувача")
	fmt.Println("  delete - видалити користувача")
}

func listUsers(db Database) {
	if len(db.Users) == 0 {
		fmt.Println("У базі даних немає жодного користувача.")
		return
	}

	fmt.Println("Список користувачів:")
	for _, user := range db.Users {
		fmt.Printf("ID: %d\nІм'я: %s\n\n", user.ID, user.Name)
	}
}

func addUser(db *Database) {

	var id int
	var name string

	fmt.Print("Введіть id нового користувача: ")
	fmt.Scanln(&id)

	for _, user := range db.Users {
		if user.ID == id {
			fmt.Println("Користувач з таким ID вже існує. Введіть унікальний ID.")
			return
		}
	}

	fmt.Print("Введіть ім'я нового користувача: ")
	fmt.Scanln(&name)

	newUser := User{
		ID:   id,
		Name: name,
	}

	db.Users = append(db.Users, newUser)

	fmt.Printf("Користувач %s успішно доданий!\n", name)
}

func updateUser(db *Database) {
	var userID int

	fmt.Print("Введіть ID користувача, якого ви хочете оновити: ")
	fmt.Scanln(&userID)

	foundIndex := -1
	for i, user := range db.Users {
		if user.ID == userID {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("Користувача з таким ID не знайдено.")
		return
	}

	var newName string

	fmt.Print("Введіть нове ім'я користувача: ")
	fmt.Scanln(&newName)

	db.Users[foundIndex].Name = newName

	fmt.Printf("Дані користувача з ID %d успішно оновлені!\n", userID)
}

func deleteUser(db *Database) {
	var userID int

	fmt.Print("Введіть ID користувача, якого ви хочете видалити: ")
	fmt.Scanln(&userID)

	foundIndex := -1
	for i, user := range db.Users {
		if user.ID == userID {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println("Користувача з таким ID не знайдено.")
		return
	}

	db.Users = append(db.Users[:foundIndex], db.Users[foundIndex+1:]...)

	fmt.Printf("Користувача з ID %d успішно видалено!\n", userID)
}
