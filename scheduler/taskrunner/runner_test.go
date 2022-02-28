package taskrunner

// test dispatcher-executor

import (
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 100; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %v", i)
		}

		return nil
	}

	e := func(dc dataChan) error {
		//for d:=range dc{

		//}
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executor received: %v", d)
			default:
				break forloop
			}
		}
		return nil
	}

	runner := NewRunner(100, false, d, e)
	go runner.StartAll()
	time.Sleep(3 * time.Second) // sleep for 3 secs

}
