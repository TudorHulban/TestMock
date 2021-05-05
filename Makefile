## -destination=mocks/mock_doer.go: put the generated mocks in the file mocks/mock_doer.go.
## -package=mocks: put the generated mocks in the package mocks.
## github.com/TudorHulban/TestMock/doer: generate mocks for this package.
## Doer: generate mocks for this interface. This argument is required — we need to specify the interfaces to generate mocks for explicitly 
## or use a comma-separated list (e.g. Doer1,Doer2).

.PHONY: gen-doer
gen-doer: 
	@mockgen -destination=mocks/mock_doer.go -package=mocks github.com/TudorHulban/TestMock/doer Doer

.PHONY: gen-mocks
gen-mocks:
	@go generate ./...