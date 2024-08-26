package xsnowflake

import (
	"fmt"
	"sync"
	"testing"
)

func TestPropsNextId(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			got := Int64()
			fmt.Println(got)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestPropsNextId2(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(NewTradeOrderId())
			//fmt.Println(NewTradeOrderId(), NewTradeOrderId2())
			wg.Done()
		}()
	}
	wg.Wait()
}
