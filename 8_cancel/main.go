// context.Contextを使って、ひとつでもエラーが発生したら全ての実行中のgoroutineを停止することもできる
//

package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	ch := make(chan int)

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)

	for i := 1; i <= 5; i++ {
		i := i
		eg.Go(func() error {
			if err := work(ctx, ch, i); err != nil {
				cancel()
				return err
			}
			return nil
		})
	}

	go func() {
		defer close(ch)
		for i := 1; i <= 100; i++ {
			ch <- i
		}
	}()

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}

func work(ctx context.Context, ch <-chan int, n int) error {
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				return nil
			}
			if i == 60 {
				return errors.New("error")
			}
			fmt.Printf("worker %d: %d\n", n, i)
		case <-ctx.Done():
			return nil
		}
	}
}
