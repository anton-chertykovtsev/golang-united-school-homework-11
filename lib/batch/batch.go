package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var ch = make(chan int64)
	var wg sync.WaitGroup
	var mx sync.Mutex
	var i int64
	for i = 0; i < pool; i++ {
		wg.Add(1)
		go func() {
			for i := range ch {
				user := getOne(i)
				mx.Lock()
				res = append(res, user)
				mx.Unlock()
			}
			wg.Done()
		}()
	}

	for i = 0; i < n; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()
	return
}
