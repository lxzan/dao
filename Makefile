test:
	go test ./...

bench:
	go test -benchmem -bench '^Benchmark' ./benchmark/
