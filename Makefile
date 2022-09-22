test:
	go test -count 1 ./...

bench:
	go test -benchmem -bench '^Benchmark' ./benchmark/
