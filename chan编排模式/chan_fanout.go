package main

// fanOut 扇出模式
func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() { // 退出时关闭所有 chan
			for _, i := range out {
				close(i)
			}
		}()

		for v := range ch { // 从 chan 中读取数据
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async { // 异步
					go func() {
						out[i] <- v // 放入到输出 chan 中,异步方式
					}()
				} else {
					out[i] <- v // 放入到输出 chan 中,同步方式
				}
			}
		}
	}()
}
