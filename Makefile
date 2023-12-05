GO111MODULE=on

.PHONY: all test clean build docker

deploy:
	cdk deploy

lint: 
	golint -set_exit_status lib/... lamda-go/...
