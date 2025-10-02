run:
	go run cmd/gameboy/main.go $(ARGS)

run-watch:
	gow run cmd/gameboy/main.go $(ARGS)

.PHONY: run run-watch
