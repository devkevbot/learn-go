// c := make(chan string)
	// var wg sync.WaitGroup

	// for _, key := range d.Keys {
	// 	wg.Add(1)
	// 	go checkKey(key, c, &wg)
	// }

	// go func() {
	// 	fmt.Println("waiting...")
	// 	wg.Wait()
	// 	close(c)
	// }()

	// // Continue reading messages from each Go routine
	// for msg := range c {
	// 	fmt.Println(msg)
	// }