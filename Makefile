# execute server
run:
	go run .
	
# Execute tests
test:
	go test -v ./...

# Run docs
doc:
	godoc -http=":3000"

# Run benchmark
bench:
	go test -run="none" -bench=. ./benchmark  -benchtime="3s" -benchmem