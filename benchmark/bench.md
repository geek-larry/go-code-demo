go test -bench .
go test -bench='Fib$' .
go test -bench='Fib$' -cpu=2,4 .
go test -bench='Fib$' -benchtime=5s .
go test -bench='Fib$' -benchtime=50x .
go test -bench='Fib$' -benchtime=5s -count=3 .

go test -bench="Fib$" -cpuprofile=cpu.pprof .
go tool pprof -text cpu.pprof