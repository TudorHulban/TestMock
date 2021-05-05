.PHONY: gen-mocks
gen-mocks: 
	@mockgen -destination=mocks/mock_doer.go -package=mocks github.com/TudorHulban/TestMock/doer Doer
