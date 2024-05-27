package research_errgroup

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func Case() {
	ctx := context.Background()
	ctx, parentCancel := context.WithCancel(ctx)

	errGroup, groupCtx := errgroup.WithContext(ctx)
	groupCtx, cancel := context.WithCancel(groupCtx)
	_ = cancel

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Parent ctx is done")
		}
	}()

	errGroup.Go(func() error {
		select {
		case <-groupCtx.Done():
			fmt.Println("#1 is done")
			return errors.New("#1 ctx is done")
		}
	})

	errGroup.Go(func() error {
		select {
		case <-groupCtx.Done():
			fmt.Println("#2 is done")
			return errors.New("#2 ctx is done")
		}
	})

	errGroup.Go(func() error {
		select {
		case <-groupCtx.Done():
			fmt.Println("#3 is done")
			return errors.New("#3 ctx is done")
		}
	})

	errGroup.Go(func() error {
		time.Sleep(1 * time.Second)

		return errors.New("#last is finished")
	})

	cancel()

	if err := errGroup.Wait(); err != nil {
		fmt.Println("Result error: ", err.Error())
	}

	fmt.Println("parent cancel")
	parentCancel()

	time.Sleep(1 * time.Second)
}
