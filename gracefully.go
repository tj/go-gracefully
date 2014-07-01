package gracefully

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Graceful shutdown utility. Block on the signal channel
// returned to perform the initial graceful shutdown. If a
// second INT or QUIT signal is trapped then exit will be
// forced with the status of 1.
func Shutdown() chan os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		sig := <-ch
		log.Printf("signal received: %s", sig)

		go func() {
			sig := <-ch
			log.Printf("second signal received: %s", sig)
			log.Printf("forcing exit")
			os.Exit(1)
		}()

		ch <- sig
	}()

	return ch
}
