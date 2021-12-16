build: bin/main
run: bin/main
dev: run-dev
test: lint unit-test
gotest: go-test


PLATFORM=local

.PHONY: bin/main
bin/main:
	@DOCKER_BUILDKIT=1 docker build . --target bin \
	--output bin/tis-gf-api \
	--platform ${PLATFORM} \

.PHONY: run
run:
	@docker run --rm -ti -p 80:8085 tis-gf-api_app ./main

.PHONY: run-dev
run-dev:
	@docker run -p 8099:8085 tis-gf-api_app ./main

.PHONY: unit-test
unit-test:
	@DOCKER_BUILDKIT=1 docker build . --target unit-test 

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker buildx build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out

.PHONY: lint
lint:
	@docker buildx build . --target lint

.PHONY: go-test 
go-test: 
	go test $$(go list ./... | grep -v ./api/config | grep -v ./models | grep -v ./secrets) -coverprofile .testCoverage.txt