Go?=go

dependence:
	$(Go) mod tidy
	$(Go) mod vendor

.PHONY: test
test: dependence
	$(Go) test ./...
