build:
	go get -d -t -v ./...
	go build ./...

generate: swagger.json
	$(RM) -r client models
	swagger generate client --name launchdarkly_api --spec $<

swagger.json:
	cp $(GOPATH)/src/github.com/launchdarkly/ld-openapi/targets/swagger.json $@
