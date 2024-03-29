package timers

import (
	"github.com/gogf/gf/g/os/glog"
	"testing"
	"time"
)

func TestDelay(t *testing.T) {
	t.Log(time.Now(), "start")
	Delay(3*time.Second, func(timer *time.Timer) {
		t.Log(time.Now(), "run task")
	})
	time.Sleep(time.Second * 5)
}

func TestAt(t *testing.T) {
	t.Log(time.Now(), "start")
	//timer.Stop()
	At(time.Now().Add(-5*time.Second), func(timer *time.Timer) {
		t.Log(time.Now(), "run task3")
	})
	At(time.Now().Add(5*time.Second), func(timer *time.Timer) {
		t.Log(time.Now(), "run task")
	})
	//timer.Stop()
	At(time.Now().Add(-5*time.Second), func(timer *time.Timer) {
		t.Log(time.Now(), "run task2")
	})
	time.Sleep(time.Second * 6)
}

func TestEvery(t *testing.T) {
	t.Log(time.Now(), "start")
	i := 0
	var ticker *time.Ticker
	ticker = Every(3*time.Second, func(timer *time.Ticker) {
		t.Log(time.Now(), "run task")
		i++
		if i == 2 {
			ticker.Stop()
		}
	})

	time.Sleep(time.Second * 10)
}

func TestLoop(t *testing.T) {
	looper := Loop(1*time.Second, func(looper *Looper) {
		glog.Println(time.Now())
	})

	Delay(5*time.Second, func(timer *time.Timer) {
		looper.Stop()
	})
	looper.Wait()
	t.Log("finished")
}

func TestLoop2(t *testing.T) {
	fromTime := time.Now()
	looper := Loop(2*time.Second, func(looper *Looper) {
		glog.Println(time.Now())

		if time.Since(fromTime) > 3*time.Second {
			looper.Stop()
			glog.Println("stop")
		}
	})

	looper.Wait()
	t.Log("finished")
}
