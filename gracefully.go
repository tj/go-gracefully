package gracefully

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Graceful shutdown utility. Block until SIGINT or SIGQUIT
// is trapped, return for shutdown and force exit on second signal.
func Shutdown() {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT)

	sig := <-ch
	log.Printf("signal received: %s", sig)

	go func() {
		sig := <-ch
		log.Printf("second signal received: %s", sig)
		log.Printf("forcing exit")
		os.Exit(1)
	}()
}
