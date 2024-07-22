package main

import "fmt"

func handleUserInput() {
	for {
		var input int
		fmt.Println("Press 1 and Enter to start publish, Any other key to stop!")
		fmt.Scanln(&input)

		mu.Lock()
		if input == 1 {
			if isPublishing {
				fmt.Println("Already Publishing")
			} else {
				stopCh = make(chan struct{})
				isPublishing = true
				go func() {
					startPublish()
					mu.Lock()
					isPublishing = false
					mu.Unlock()
				}()
			}
		} else {
			if isPublishing {
				close(stopCh)
				isPublishing = false
				fmt.Println("Publish Stop")
			} else {
				fmt.Println("Currently not publishing")
			}
		}
		mu.Unlock()
	}
}
