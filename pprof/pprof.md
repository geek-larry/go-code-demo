go run main.go > cpu.pprof
go tool pprof -http=:9999 cpu.pprof
go tool pprof cpu.pprof