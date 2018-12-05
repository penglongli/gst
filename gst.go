package gst

// For 循环协程的创建与销毁
import (
	"sync"
)

var (
	iMap map[string]chan int
)

func init() {
	iMap = make(map[string]chan int)
}

func NewLoopRoutine(key string, wg *sync.WaitGroup, f func()) {
	quit := make(chan int)
	iMap[key] = quit

	go handle(quit, wg, f)
}

func handle(quit chan int, wg *sync.WaitGroup, f func()) {
	for {
		select {
		case <-quit: {
			wg.Done()
            return
		}
		default:
			f()
		}
	}
}

func StopLoopRoutine(key string) {
	quit := iMap[key]
	if quit == nil {
		return
	}
	quit <- 1
}

