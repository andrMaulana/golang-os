package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	helper "go-exercise/helper"
	"os"
	"strconv"
	"strings"
)

type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

var menuItems []MenuItem
var nextID = 1
var filename = "menu.json"

func main() {
	loadMenu()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
		saveMenu()
	}()

	for {
		helper.ClearScreen()
		fmt.Println("\nMenu Options:")
		fmt.Println("1. Add Menu Item")
		fmt.Println("2. View Menu")
		fmt.Println("3. Update Menu Item")
		fmt.Println("4. Delete Menu Item")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		choice := getUserInput()

		switch choice {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Tambah Menu ===")
			addMenuItem()
		case "2":
			viewMenu()
		case "3":
			updateMenuItem()
		case "4":
			deleteMenuItem()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addMenuItem() {
	fmt.Println("\nAdd Menu Item")
	fmt.Print("Enter name of the item: ")
	name := getUserInput()
	fmt.Print("Enter price of the item: ")
	priceStr := getUserInput()
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Invalid price. Please enter a valid number.")
		return
	}

	menuItem := MenuItem{
		ID:    nextID,
		Name:  name,
		Price: price,
	}
	nextID++
	menuItems = append(menuItems, menuItem)
	fmt.Println("Menu item added successfully.")
}

func viewMenu() {
	fmt.Println("\nMenu:")
	if len(menuItems) == 0 {
		fmt.Println("No items in the menu.")
		return
	}
	for _, item := range menuItems {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f\n", item.ID, item.Name, item.Price)
	}
}

func updateMenuItem() {
	fmt.Println("\nUpdate Menu Item")
	fmt.Print("Enter ID of the item to update: ")
	idStr := getUserInput()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}

	index := findMenuItemIndexByID(id)
	if index == -1 {
		fmt.Println("Item not found.")
		return
	}

	fmt.Print("Enter new name of the item: ")
	newName := getUserInput()
	fmt.Print("Enter new price of the item: ")
	newPriceStr := getUserInput()
	newPrice, err := strconv.ParseFloat(newPriceStr, 64)
	if err != nil {
		fmt.Println("Invalid price. Please enter a valid number.")
		return
	}

	menuItems[index].Name = newName
	menuItems[index].Price = newPrice
	fmt.Println("Menu item updated successfully.")
}

func deleteMenuItem() {
	fmt.Println("\nDelete Menu Item")
	fmt.Print("Enter ID of the item to delete: ")
	idStr := getUserInput()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a valid number.")
		return
	}

	index := findMenuItemIndexByID(id)
	if index == -1 {
		fmt.Println("Item not found.")
		return
	}

	menuItems = append(menuItems[:index], menuItems[index+1:]...)
	fmt.Println("Menu item deleted successfully.")
}

func findMenuItemIndexByID(id int) int {
	for i, item := range menuItems {
		if item.ID == id {
			return i
		}
	}
	return -1
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func loadMenu() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&menuItems); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Find the next ID
	for _, item := range menuItems {
		if item.ID >= nextID {
			nextID = item.ID + 1
		}
	}
}

func saveMenu() {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(menuItems); err != nil {
		panic(err)
	}
}
