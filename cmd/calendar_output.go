package cmd

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	_ "math"
	"os"
	_ "os/exec"
	"time"
)

func generateOutput(colorDisable bool) {
	if colorDisable {
		color.NoColor = true
	}

	currentTime := time.Now()
	day := currentTime.Day()
	weekDay := returnWeakDayNumber(string(rune(currentTime.Weekday())))
	currentDay := returnDaysInMont(currentTime.Month().String())
	firstMothWeekday := returnFirstWeakDayNumber(weekDay, currentDay)

	fmt.Print("                ")
	color.Yellow(currentTime.Month().String())
	color.Blue("M     T     W     T     F     S     S")
	fmt.Println("")
	qwe := firstMothWeekday

	blue := color.New(color.FgHiBlue)
	boldBlue := blue.Add(color.Bold)

	for i := 1; i < qwe; i++ {
		fmt.Print("      ")
	}

	//first line
	for i := 1; i < 9-qwe; i++ {

		if i != day {
			fmt.Print(i)
		}
		if i == day {
			boldBlue.Print(i)
		}
		//fmt.Print(i)
		fmt.Print("     ")
	}
	fmt.Println("")

	//second line
	for i := 9 - qwe; i < 16-qwe; i++ {
		if i != day {
			fmt.Print(i)
		}
		if i == day {
			boldBlue.Print(i)
		}
		if i > 9 {
			fmt.Print("    ")
		}

		if i < 10 {
			fmt.Print("     ")
		}
	}
	fmt.Println("")

	//third line
	for i := 16 - qwe; i < 23-qwe; i++ {
		if i != day {
			fmt.Print(i)
		}
		if i == day {
			boldBlue.Print(i)
		}
		if i > 9 {
			fmt.Print("    ")
		}
		if i < 10 {
			fmt.Print("     ")
		}
	}
	fmt.Println("")

	//fourth line
	for i := 23 - qwe; i < 30-qwe; i++ {
		if i != day {
			fmt.Print(i)
		}
		if i == day {
			boldBlue.Print(i)
		}
		fmt.Print("    ")
	}
	fmt.Println("")

	//fifth line
	for i := 30 - qwe; i < 37-qwe; i++ {
		if i > returnDaysInMont(currentTime.Month().String()) {
			break
		}
		if i != day {
			fmt.Print(i)
		}
		if i == day {
			boldBlue.Print(i)
		}
		fmt.Print("    ")
	}

	fmt.Println("")

	//sixth line
	for i := 43 - qwe; i < 50-qwe; i++ {
		if i > returnDaysInMont(currentTime.Month().String()) {
			break
		}
		if i != day {
			fmt.Print(i)

		}
		if i == day {
			boldBlue.Print(i)
		}
		fmt.Print("    ")
	}
	fmt.Println("")

	stopCh := make(chan bool)

	// start a goroutine to wait for keyboard input
	go func() {
		fmt.Println("Press any key to exit...")
		reader := bufio.NewReader(os.Stdin)
		_, _, _ = reader.ReadRune()
		stopCh <- true
	}()

	// start the infinite loop
	for {
		select {
		case <-stopCh:
			fmt.Println("Stopping calendar...")
			return
		default:
			currentTime := time.Now()
			c := currentTime.Format("01-02-2006 15:04:05")

			info := color.New(color.Bold, color.FgYellow).SprintFunc()
			fmt.Printf("%s \r", info(c))

			time.Sleep(1 * time.Second)
			clearScreen()

		}
	}
}
