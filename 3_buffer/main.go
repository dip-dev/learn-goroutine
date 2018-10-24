// チャネルは、 bufferとして使える
// バッファを持つチャネルを初期化するには、 make の２つ目の引数にバッファの長さを指定する
//
// ch := make(chan int, 100)
// バッファが詰まった時は、チャネルへの送信をブロックする
// バッファが空の時には、チャネルの受信をブロックする
//

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("send %d\n", i)
		}
	}()

	time.Sleep(5 * time.Second)
	for i := 0; i < 5; i++ {
		n := <-ch
		fmt.Printf("recieve %d\n", n)
	}
}
