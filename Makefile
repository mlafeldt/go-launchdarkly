build:
	go get -d -t -v ./...
	go build ./...

generate:
	$(RM) -r client models
	swagger generate client \
		--name launchdarkly_api \
		--spec $(GOPATH)/src/github.com/launchdarkly/ld-openapi/targets/swagger.json
