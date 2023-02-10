package main

import "fmt"

func main() {
	// Channels
	// c := make(chan int)
	// go func() {
	// 	for i := 0; i < 1000000; i++ {
	// 		c <- i
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		fmt.Println(<-c)
	// 	}
	// }()
	// time.Sleep(4 * time.Second) // the program wait 4 seconds for the execution to complete

	// Range Clause
	// c := make(chan int)
	// go func() {
	// 	for i := 0; i < 10000; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()

	// for n := range c {
	// 	fmt.Println(n)
	// }

	// N to 1
	// c := make(chan int)

	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	fmt.Println("One")
	// 	for i := 0; i < 10; i++ {
	// 		c <- i
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	fmt.Println("Two")
	// 	for i := 0; i < 10; i++ {
	// 		c <- i
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	wg.Wait()
	// 	close(c)
	// }()

	// for n := range c {
	// 	fmt.Println(n)
	// }

	// Semaphore
	// n := 2
	// c := make(chan int)
	// done := make(chan bool)

	// for j := 0; j < n; j++ {
	// 	go func() {
	// 		for i := 0; i < 10; i++ {
	// 			c <- i
	// 			// time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	// 		}
	// 		done <- true
	// 	}()
	// }
	// // go func() {
	// // 	for i := 0; i < 10; i++ {
	// // 		c <- i
	// // 		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
	// // 	}
	// // 	done <- true
	// // }()

	// go func() {
	// 	for i := 0; i < n; i++ {
	// 		<-done
	// 	}
	// 	// <-done
	// 	close(c)
	// }()

	// for n := range c {
	// 	fmt.Println(n)
	// }

	// 1 to N
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}
	// go func() {
	// 	// fmt.Println("Two")
	// 	for n := range c {
	// 		fmt.Println(n)
	// 	}
	// 	done <- true
	// }()

	// <-done
	for i := 0; i < 10; i++ {
		<-done
	}

}
