test:
	go1.18beta2 test ./...

bench:
	go1.18beta2 test -benchmem -bench '^Benchmark' ./benchmark/
