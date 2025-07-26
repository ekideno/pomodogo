package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Print("\033[2J\033[H")

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: pomodogo <work_minutes> <break_minutes>")
		return
	}

	workMin, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Ошибка: work_minutes должно быть числом")
		return
	}

	breakMin, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Ошибка: break_minutes должно быть числом")
		return
	}

	fmt.Println("⏱️ Session info")
	fmt.Println("Now: " + time.Now().Format("15:04"))
	fmt.Printf("Work: %dm\n", workMin)
	fmt.Printf("Break: %dm\n\n", breakMin)

	workDuration := time.Duration(workMin) * time.Second
	breakDuration := time.Duration(breakMin) * time.Second

	fmt.Println("🍅 Working...")
	progressBar(workDuration)

	fmt.Print("\033[1A")
	fmt.Print("\033[2K\r")
	fmt.Print("\033[1A")
	fmt.Print("\033[2K\r")

	fmt.Println("Time to break!")
	fmt.Print("Press Enter to start...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	for i := 0; i < 2; i++ {
		fmt.Print("\033[1A")
		fmt.Print("\033[2K\r")
	}

	fmt.Println("🍏 Break...")
	progressBar(breakDuration)

	fmt.Println("\nSession completed! Good job!")
}

func progressBar(duration time.Duration) {
	total := int(duration.Seconds())
	barWidth := 30

	for elapsed := 0; elapsed <= total; elapsed++ {
		percent := float64(elapsed) / float64(total)
		filled := int(percent * float64(barWidth))
		empty := barWidth - filled

		fmt.Printf("\r[%s%s] %3.0f%%  %02d:%02d left",
			repeat("*", filled),
			repeat(" ", empty),
			percent*100,
			(total-elapsed)/60,
			(total-elapsed)%60,
		)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
