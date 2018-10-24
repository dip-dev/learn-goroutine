// goroutineでのエラーを受け取りたい場合は、errgroup.Group{}を使う
//
// eg.Go() に実装したいメソッド（func() error）を渡して、eg.Wait()でWaitする
// 全てのgoroutineが終了したら、eg.Wait()がerrorを返す
// goroutine内で最初に発生したエラーしか受け取れない
//

package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	eg := &errgroup.Group{}
	for i := 1; i <= 10; i++ {
		eg.Go(func() error {
			i := <-ch
			time.Sleep(time.Duration(i) * time.Second)
			if i == 6 {
				return errors.New("error")
			}
			fmt.Println(i)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}
