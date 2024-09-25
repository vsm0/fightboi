MAIN = build/main.o

$(MAIN): ./**
	@mkdir -p build
	@go build -ldflags=-checklinkname=0 -o $(MAIN) ./cmd

run: $(MAIN)
	./$(MAIN)

.PHONY: clean
clean: $(MAIN)
	@rm $(MAIN)
