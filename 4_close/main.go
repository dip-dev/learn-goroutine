// 送り手は、これ以上の送信する値がないことを示すため、チャネルを close できる
// 受け手は、受信の式に2つ目のパラメータを割り当てることで、そのチャネルがcloseされているかどうかを確認できる
//
// v, ok := <-ch
// 受信する値がない、かつ、チャネルが閉じているなら、 ok の変数は、 false
//
// ループの for i := range c は、チャネルが閉じられるまで、チャネルから値を繰り返し受信し続ける
//
// ※ 送り手のチャネルだけをcloseするし、受け手はcloseしてはいけない
// もしcloseしたチャネルへ送信すると、panicになる
//
// ※ チャネルは、ファイルとは異なり、通常は、closeする必要はありません。
// closeするのは、これ以上値が来ないことを受け手が知る必要があるときにだけ
// 例えば、 range ループを終了するという場合
//

package main

import "fmt"

// 引数でchannelへの送信だけを許可するようにしている
// 受信だけ許可するようにすることも可能
// func(<-chan)
func send(n int, ch chan<- int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)

	go send(cap(ch), ch)

	for i := range ch {
		fmt.Println(i)
	}
}
