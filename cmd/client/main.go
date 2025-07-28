package main

func main() {
	Execute()
}

// func main() {
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer watcher.Close()

// 	// Recursively add all subdirectories
// 	err = filepath.Walk("/home/athena/me/Athena", func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if info.IsDir() {
// 			log.Println("Watching:", path)
// 			return watcher.Add(path)
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case event, ok := <-watcher.Events:
// 				if !ok {
// 					return
// 				}
// 				log.Println("event:", event)

// 				log.Println(event.Name)

// 				// If a new directory is created, add it to the watcher
// 				if event.Op&fsnotify.Create == fsnotify.Create {
// 					fi, err := os.Stat(event.Name)
// 					if err == nil && fi.IsDir() {
// 						log.Println("New directory detected, watching:", event.Name)
// 						watcher.Add(event.Name)
// 					}
// 				}
// 			case err, ok := <-watcher.Errors:
// 				if !ok {
// 					return
// 				}
// 				log.Println("error:", err)
// 			}
// 		}
// 	}()

// 	// Block forever
// 	select {}
// }
