package main

import (
	"fmt"
	"time"
	"os"
)

func main() {



	fmt.Println("We goooooooing!")


	

	// Create or open a file for writing
	file, err := os.Create("communication.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write messages to the file
	for i := 1; i <= 100; i++ {
		message := fmt.Sprintf("%d", i)
		_, err := file.WriteString(message)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		time.Sleep(1 * time.Second) // Sleep for 1 second between messages
	}




	fmt.Println("Program terminated.")
}
