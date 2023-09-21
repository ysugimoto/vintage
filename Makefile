.PHONY: test

test:
	go test .
	go test github.com/ysugimoto/vintage/runtime/fastly
