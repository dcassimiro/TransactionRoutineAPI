.PHONY: run  mock  test

VERSION = $(shell git branch --show-current)

# commands to execute
run:
	VERSION=$(VERSION) go run main.go


# test commands
test:
	go test -coverprofile=coverage.out ./app/... ./api/... ./model/... ./store/...

# command to generate mocks
mock: 
	rm -rf ./mocks

	mockgen -source=./store/account/account.go -destination=./mocks/account_store_mock.go -package=mocks -mock_names=Store=MockAccountStore
	mockgen -source=./store/transaction/transaction.go -destination=./mocks/transaction_store_mock.go -package=mocks -mock_names=Store=MockTransactionStore
	
	mockgen -source=./app/account/account.go -destination=./mocks/account_app_mock.go -package=mocks -mock_names=App=MockAccountApp
	mockgen -source=./app/transaction/transaction.go -destination=./mocks/transaction_app_mock.go -package=mocks -mock_names=App=MockTransactionApp