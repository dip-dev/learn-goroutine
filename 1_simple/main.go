// goroutine (ゴルーチン)は、Goのランタイムに管理される軽量なスレッド
//
// `go f(x, y, z)`
// と書けば、新しいgoroutineが実行される
//
// goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要がある
// 同期する方法は主に2つ
//  - sync パッケージ
//  - channel
//

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	start := time.Now()

	go say("world")
	say("hello")

	fmt.Printf("%f sec\n", (time.Now().Sub(start)).Seconds())
}
