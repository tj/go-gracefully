package main

import . ".."
import "time"
import "log"

func work(stop, done chan bool) {
	for {
		select {
		case <-stop:
			log.Printf("quitting in 5 seconds")
			time.Sleep(5 * time.Second)
			log.Printf("bye :)")
			done <- true
		default:
			log.Printf("doing work")
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func shutdown(stop chan bool) {
	<-Shutdown()
	stop <- true
}

func main() {
	log.Printf("send SIGINT or SIGQUIT to shutdown")

	done := make(chan bool)
	stop := make(chan bool)

	// register shutdown
	go shutdown(stop)

	// pretend to do some work
	go work(stop, done)
	<-done
}
