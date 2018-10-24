// select ステートメントは、goroutineを複数の通信操作（channelによる処理）ができる
//
// select は、複数ある case のいずれかが準備できるようになるまでブロックし、 準備ができた case を実行
// もし、複数の case の準備ができている場合、 case はランダムに選択される
//
// どの case も準備ができていないのであれば、 select の中の default が実行される
//

package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	boom := time.After(5 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
