run:
	go run main.go

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

bench:
	go test -bench='^Benchmark*' ./...

