package fonts

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
)

// InstantiateWatcher is a thing
func InstantiateWatcher(path string) {
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		timers := make(map[string]*time.Timer)
		timerChan := make(chan string)

		done := make(chan bool)

		go func() {
			for {
				select {
				case event := <-watcher.Events:
					// log.Printf("inotify event %s", event)

					if event.Op&fsnotify.Write == fsnotify.Write {
						if timers[event.Name] != nil {
							timers[event.Name].Stop()
						}
						timers[event.Name] = time.NewTimer(1 * time.Second)
						go func(path string, c <-chan time.Time) {
							for range c {
								timerChan <- path
							}
						}(event.Name, timers[event.Name].C)
					}

				case err := <-watcher.Errors:
					log.Printf("Got error watching %s, calling watcher func", err)
				case tick := <-timerChan:
					log.Println(tick)
				}
			}
		}()

		log.Printf("Start watching file %s", path)

		err = watcher.Add(path)

		if err != nil {
			log.Fatal(fmt.Sprintf("Got error adding watcher %s", err))
		}
		<-done
	}()
}
