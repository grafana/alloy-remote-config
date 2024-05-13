.PHONY: build-container
build-container:
	docker build -t alloy-remote-cfg-builder .

.PHONY: buf-generate
buf-generate:
	docker run -v $(CURDIR):/api alloy-remote-cfg-builder
