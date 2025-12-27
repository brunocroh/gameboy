run:
	go run cmd/gameboy/main.go -rom="$(ARGS)"

run-single-step:
	go run cmd/gameboy/main.go -rom=$(ARGS) -single-step

run-watch:
	gow run cmd/gameboy/main.go $(ARGS)

.PHONY: run run-watch run-single-step
