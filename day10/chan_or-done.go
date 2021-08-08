package main

import "time"

// or or-done 模式
func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个,1个或2个 chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	case 2:
		select {
		case <-channels[0]:
		case <-channels[1]:
		}
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		if len(channels) > 2 {
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return orDone
}

func sig(after time.Duration) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()

	return c
}

func main() {

}
