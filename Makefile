build:
	go get -d -t -v ./...
	go build ./...

generate: swagger.yaml
	$(RM) -r client models
	docker run -it --rm \
		-v $(PWD):/go/src/github.com/mlafeldt/go-launchdarkly \
		-w /go/src/github.com/mlafeldt/go-launchdarkly \
		quay.io/goswagger/swagger generate client --name launchdarkly --spec $<

swagger.yaml:
	curl -fSsL https://launchdarkly.github.io/ld-openapi/swagger.yaml -o $@
