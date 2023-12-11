package main

import (
	"errors"
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

func program_main() {
	db := Database{}

	fmt.Println("Ласкаво просимо до консольної програми - бази даних користувачів!")
	fmt.Println("Введіть 'help' для перегляду доступних команд.")

	for {
		fmt.Print("\nВведіть команду: ")
		var command string
		if _, err := fmt.Scanln(&command); err != nil {
			fmt.Printf("Помилка введення команди: %s\n", err)
			continue
		}

		err := processCommand(command, &db)
		if err != nil {
			fmt.Printf("Помилка: %s\n", err)
		}
	}
}

func processCommand(command string, db *Database) error {
	switch command {
	case CommandHelp:
		printHelp()
	case CommandList:
		listUsers(*db)
	case CommandAdd:
		err := addUser(db)
		if err != nil {
			return err
		}
	case CommandUpdate:
		err := updateUser(db)
		if err != nil {
			return err
		}
	case CommandDelete:
		err := deleteUser(db)
		if err != nil {
			return err
		}
	default:
		return errors.New("Невідома команда. Введіть 'help' для отримання списку команд.")
	}

	return nil
}

func printHelp() {
	fmt.Println("Доступні команди:")
	fmt.Printf("  %s - переглянути список команд\n", CommandHelp)
	fmt.Printf("  %s - переглянути список користувачів\n", CommandList)
	fmt.Printf("  %s - створити нового користувача\n", CommandAdd)
	fmt.Printf("  %s - оновити дані користувача\n", CommandUpdate)
	fmt.Printf("  %s - видалити користувача\n", CommandDelete)
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

func addUser(db *Database) error {
	var id int
	var name string

	fmt.Print("Введіть id нового користувача: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		return fmt.Errorf("Помилка введення id: %s", err)
	}

	for _, user := range db.Users {
		if user.ID == id {
			return errors.New("Користувач з таким ID вже існує. Введіть унікальний ID.")
		}
	}

	fmt.Print("Введіть ім'я нового користувача: ")
	_, err = fmt.Scanln(&name)
	if err != nil {
		return fmt.Errorf("Помилка введення імені: %s", err)
	}

	newUser := User{
		ID:   id,
		Name: name,
	}

	db.Users = append(db.Users, newUser)

	fmt.Printf("Користувач %s успішно доданий!\n", name)
	return nil
}

func updateUser(db *Database) error {
	var userID int

	fmt.Print("Введіть ID користувача, якого ви хочете оновити: ")
	_, err := fmt.Scanln(&userID)
	if err != nil {
		return fmt.Errorf("Помилка введення ID: %s", err)
	}

	foundIndex := findUserIndexByID(*db, userID)

	if foundIndex == -1 {
		return errors.New("Користувача з таким ID не знайдено.")
	}

	var newName string

	fmt.Print("Введіть нове ім'я користувача: ")
	_, err = fmt.Scanln(&newName)
	if err != nil {
		return fmt.Errorf("Помилка введення нового імені: %s", err)
	}

	db.Users[foundIndex].Name = newName

	fmt.Printf("Дані користувача з ID %d успішно оновлені!\n", userID)
	return nil
}

func deleteUser(db *Database) error {
	var userID int

	fmt.Print("Введіть ID користувача, якого ви хочете видалити: ")
	_, err := fmt.Scanln(&userID)
	if err != nil {
		return fmt.Errorf("Помилка введення ID: %s", err)
	}

	foundIndex := findUserIndexByID(*db, userID)

	if foundIndex == -1 {
		return errors.New("Користувача з таким ID не знайдено.")
	}

	db.Users = append(db.Users[:foundIndex], db.Users[foundIndex+1:]...)

	fmt.Printf("Користувача з ID %d успішно видалено!\n", userID)
	return nil
}

func findUserIndexByID(db Database, userID int) int {
	for i, user := range db.Users {
		if user.ID == userID {
			return i
		}
	}
	return -1
}
