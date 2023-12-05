package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/zhyzhkoslava/ITEA-GO/homework9/dto"
)

func printUserInfo(username string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("User not found.")
		return nil
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	user := new(dto.GitHubUser)
	if err := json.Unmarshal(bodyBytes, user); err != nil {
		return err
	}

	fmt.Printf("ID:\t%d\nName:\t%s\nBio:\t%s\nCreation Year:\t%s\n", user.Id, user.Name, user.Bio, user.CreatedAt)
	os.Exit(0)
	return nil
}

func main() {
	fmt.Println("Ласкаво просимо до консольної програми")
	for {
		fmt.Print("\nВведіть імʼя юзера: ")
		var username string
		if _, err := fmt.Scanln(&username); err != nil {
			fmt.Printf("Помилка введення команди: %s\n", err)
			continue
		}

		err := printUserInfo(username)
		if err != nil {
			fmt.Printf("Помилка: %s\n", err)
		}
	}
}
