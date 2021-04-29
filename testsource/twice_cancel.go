package main

import (
	"context"
	"fmt"
)

func twiceCancel() {
	ctx, ccl := context.WithCancel(context.Background())
	ccl()
	ccl()
	fmt.Println(<-ctx.Done())
}
