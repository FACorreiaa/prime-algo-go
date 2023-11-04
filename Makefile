run:
	go run main.go
test:
	go test -run '^Test(day_sort|day_graphs)*' ./...

bench:
	go test -bench='^Benchmark*' ./...

