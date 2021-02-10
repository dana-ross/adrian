package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/fsnotify/fsnotify"
)

// FindFonts loads all the fonts in a directory
func FindFonts(fontsPath string, config Config) {
	filesInfo, err := ioutil.ReadDir(fontsPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range filesInfo {
		if fileInfo.IsDir() == false {
			LoadFont(path.Join(fontsPath, fileInfo.Name()), config)
		}
	}
}

// InstantiateWatcher is a thing
func InstantiateWatcher(path string, config Config) {
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
				case path := <-timerChan:
					LoadFont(path, config)
				}
			}
		}()

		err = watcher.Add(path)

		if err != nil {
			log.Fatal(fmt.Sprintf("Got error adding watcher %s", err))
		}
		<-done
	}()
}
