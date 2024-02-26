package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var clearScreen map[string]func() //create a map for storing clearScreen funcs

func init() {
	clearScreen = make(map[string]func()) //Initialize it
	clearScreen["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clearScreen["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clearScreen["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

	}
}

func CallClear() {
	macos := runtime.GOOS
	value, ok := clearScreen[macos] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                         //if we defined a clearScreen func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clearScreen terminal screen :(")
	}
}

func ShowClock() {
	height := -1
	for key, val := range fonts {
		valSplit := strings.Split(val, "\n")
		if height == -1 {
			height = len(valSplit)
		} else {
			if len(valSplit) != height {
				panic("Font " + key + " has different height")
			}
		}
	}

	clockStr := time.Now().Format("15:04:05")
	clockSplit := strings.Split(clockStr, "")
	for i := 0; i < height; i++ {
		fmt.Print("  ")
		for _, num := range clockSplit {
			fontItemSplit := strings.Split(fonts[num], "\n")
			fmt.Print(fontItemSplit[i] + "   ")
		}
		fmt.Println()
	}
}

func main() {
	for {
		CallClear()
		ShowClock()
		time.Sleep(100 * time.Millisecond)
	}
}
