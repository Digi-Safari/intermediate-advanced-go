package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	i, err := Slow(ctx, "10")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}

// context should be the first argument passed to function

func Slow(ctx context.Context, input string) (int, error) {
	//sql.DB{}.ExecContext()
	i, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	time.Sleep(1 * time.Second)
	// checking if timeout happened, if yes then we would not return error
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		return i, nil
	}

}
