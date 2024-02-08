package main

import (
	"fmt"
	"os/exec"
	"time"
	"os"
)

func main() {
	// Start a new process for the previous program
	exec.Command("gnome-terminal", "--", "go", "run", "oppgaveB.go").Run()
	exec.Command("gnome-terminal", "--", "go", "run", "exercise4.go").Run()
}

func A() {}

func B(x int) {

	//Create txt file
	file, err := os.Create("communication.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	//Write messages to the file
	for i:=x; i <= 10; i++ {
		message := fmt.Sprintf("%d", i)
		_, err := file.WriteString(message)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		time.Sleep(1 * time.Second) // Sleep for 1 second between messages
	}
}


/*file, err := os.Open("communication.txt")
if err != nil {
	fmt.Println("File doesnt exists", err)
	file, _ = os.Create("communication.txt")
} */
