package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"os/exec"
)





func main() {
	fmt.Println("Program started.")

	time.Sleep(time.Second *1)



	// Open the file for reading
	file, err := os.Open("communication.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	timeout := time.Now()

	// Read messages from the file
	for i:=2; i>1; {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		message := scanner.Text()
		messageInt, _ := strconv.Atoi(message)
		fmt.Println("Program 2 received: %d", messageInt)
		timeout = time.Now()
		timeout = timeout.Add(time.Second * 2)
	}


	if timeout.Before(time.Now()) {
		fmt.Println("B is dead")
		exec.Command("gnome-terminal", "--", "go", "run", "oppgaveB.go").Run()
		time.Sleep(time.Second * 4)
	
	}

	

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}



	//keeps A going
	i := 2
	for i > 1 {

	}


	fmt.Println("Program terminated.")
}