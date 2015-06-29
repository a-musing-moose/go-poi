build: clean
	go build *.go

build-static: clean
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo *.go

clean:
	rm -f go-poi

image: build-static
	docker build --rm -t snowballone/go-poi .

run-image: image
	docker run -p 8000:8000 -v /tmp:/data snowballone/go-poi
