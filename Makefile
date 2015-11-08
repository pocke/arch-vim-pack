all:
	go get github.com/jteeuwen/go-bindata/...
	go-bindata data/
	go build

test:
	go get github.com/jteeuwen/go-bindata/...
	go-bindata data/ testdata/
	go test -v
