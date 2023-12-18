test:
	go test -count 1 ./...

bench:
	go test -benchmem -bench '^Benchmark' ./benchmark/

cover:
	go test -coverprofile=./bin/cover.out --cover ./...