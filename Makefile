build:
	go get -d -t -v ./...
	go build ./...

generate: swagger.json
	$(RM) -r client models
	docker run -it --rm \
		-v $(PWD):/go/src/github.com/mlafeldt/go-launchdarkly \
		-w /go/src/github.com/mlafeldt/go-launchdarkly \
		quay.io/goswagger/swagger generate client --name launchdarkly --spec $<

swagger.json:
	cp $(GOPATH)/src/github.com/launchdarkly/ld-openapi/targets/swagger.json $@
