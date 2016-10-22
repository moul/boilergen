SOURCES :=	$(shell find . -name "*.go")

.PHONY: build
build: boilergen

boilergen: $(SOURCES)
	go build -o $@ ./cmd/boilergen

.PHONY: clean
clean:
	rm -f boilergen

.PHONY: install
install:
	go install -v ./cmd/boilergen
