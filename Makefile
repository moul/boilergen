SOURCES :=	$(shell find . -name "*.go")

.PHONY: build
build: boilergen

boilergen: $(SOURCES)
	go build -o $@ ./cmd/boilergen/main.go

clean:
	rm -f boilergen
