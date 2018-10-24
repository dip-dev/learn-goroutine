// channelを使う以外の同期方法で、sync パッケージがある
//
// sync.WaitGroupは全てのgoroutineの処理が終了するまで、waitする
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("func1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Second)
		fmt.Println("func2")
	}()

	wg.Wait()
	fmt.Printf("%f sec\n", (time.Now().Sub(start)).Seconds())

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("func%d\n", i)
		}(i)
	}
	wg.Wait()
}
