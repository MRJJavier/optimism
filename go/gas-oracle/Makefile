gas-oracle:
	env GO111MODULE=on go build

test:
	go test -v ./...

lint:
	golangci-lint run ./...

binding:
	$(eval tempdir := $(shell mktemp))

	cat abis/OVM_GasPriceOracle.json \
		| jq -r .bin > $(tempdir)

	cat abis/OVM_GasPriceOracle.json \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/gaspriceoracle.go \
		--type GasPriceOracle \
		--bin $(tempdir)

	rm $(tempdir)
