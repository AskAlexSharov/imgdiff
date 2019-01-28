lint:
	gometalinter.v2 ./... --vendor --deadline=1m --disable-all --enable=gosec
	golangci-lint run --enable-all