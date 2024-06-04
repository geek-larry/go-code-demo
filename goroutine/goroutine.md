go test -run ^TestBadTimeout$ . -v

    PS D:\.code\Golang\src\go-code-demo\goroutine> go test -run ^TestBadTimeout$ . -v
    === RUN   TestBadTimeout
        main_test.go:17: 1002
    --- PASS: TestBadTimeout (17.64s)
    PASS
    ok      go-code-demo/goroutine  17.814s

go test -run ^TestBufferTimeout$ . -v

    PS D:\.code\Golang\src\go-code-demo\goroutine> go test -run ^TestBufferTimeout$ . -v
    === RUN   TestBufferTimeout
        main_test.go:22: 2
    --- PASS: TestBufferTimeout (17.71s)
    PASS
    ok      go-code-demo/goroutine  17.880s

go test -run ^Test2phasesTimeout$ . -v

    PS D:\.code\Golang\src\go-code-demo\goroutine> go test -run ^Test2phasesTimeout$ . -v
    === RUN   Test2phasesTimeout
        main_test.go:33: 2
    --- PASS: Test2phasesTimeout (18.55s)
    PASS
    ok      go-code-demo/goroutine  18.711s

这种场景在实际的业务中更为常见，例如我们将服务端接收请求后的任务拆分为 2 段，一段是执行任务，一段是发送结果。那么就会有两种情况：

- 任务正常执行，向客户端返回执行结果。
- 任务超时执行，向客户端返回超时。

这种情况下，就只能够使用 select，而不能能够设置缓冲区的方式了。因为如果给信道 phase1 设置了缓冲区，phase1 <- true 总能执行成功，那么无论是否超时，都会执行到第二阶段，而没有即时返回，这是我们不愿意看到的。对应到上面的业务，就可能发生一种异常情况，向客户端发送了 2 次响应：

- 任务超时执行，向客户端返回超时，一段时间后，向客户端返回执行结果。

缓冲区不能够区分是否超时了，但是 select 可以（没有接收方，信道发送信号失败，则说明超时了）。