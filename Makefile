SOURCES :=	$(shell find . -name "*.go")
GIT_BRANCH ?=	$(shell git rev-parse --abbrev-ref HEAD)
DOCKER_IMAGE ?=	moul/boilergen:$(GIT_BRANCH)

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

.PHONY: docker.build
docker.build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker.push
docker.push: docker.build
	docker push $(DOCKER_IMAGE)
