PKG_ML = $(shell go list ./... | sed "s%_$$(pwd)%\.%g" | grep -v -e "vendor*")
PKG = $(shell echo ${PKG_ML} | tr "\n" " ")

compile:
	mkdir -p dist
	go build -o dist/go-parser

testv:
	go test -v ${PKG}
cover:
	mkdir -p reports
	echo "mode: count" > reports/coverage-all.out
	$(foreach pkg,$(PKG_ML), \
		go test -coverprofile=reports/coverage.out -covermode=count $(pkg); \
		tail -n +2 reports/coverage.out >> reports/coverage-all.out; \
	)
	go tool cover -html=reports/coverage-all.out -o reports/coverage.html