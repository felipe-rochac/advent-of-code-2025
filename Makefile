.PHONY: generate-day test clean

generate-day:
	@read -p "Enter day number (1-25): " day; \
	go run cmd/generate-day/main.go -day $$day

test:
	go test ./...

clean:
	rm -f day*/__debug_bin*
