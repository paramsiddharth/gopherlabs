package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	wait := make(chan int)

	go func() {
		signal := <-signalChan
		switch signal {
		case syscall.SIGINT:
			fmt.Fprintln(os.Stderr, "\rInterrupted!")
		default:
			fmt.Fprintf(os.Stderr, "\rSignal: %d\n", signal)
		}
		wait <- 1
	}()

	code := <-wait
	os.Exit(code)
}
