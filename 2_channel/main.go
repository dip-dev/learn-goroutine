// チャネル( Channel )型は、チャネルオペレータの <- を用いて値の送受信ができる通り道
//
// ch <- v    // v をチャネル ch へ送信する
// v := <-ch  // ch から受信した変数を v へ割り当てる
//
// マップとスライスのように、チャネルは使う前に以下のように生成
// ch := make(chan int)
//
// 通常、片方が準備できるまで送受信はブロックされる
// これにより、明確なロックや条件変数がなくても、goroutineの同期を可能
//

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("func1")
		ch <- struct{}{}
	}()

	<-ch // ここでブロック

	fmt.Printf("%f sec\n", (time.Now().Sub(start)).Seconds())
}
