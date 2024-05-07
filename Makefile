.PHONY: buf-generate
buf-generate:
	buf generate api -o ./api

.PHONY: build-container
build-container:
	docker build -t alloy-remote-cfg-builder .

.PHONY: buf-generate-container
buf-generate-container:
	docker run -v $(CURDIR):/api alloy-remote-cfg-builder
