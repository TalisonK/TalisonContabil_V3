
build:
	go build -o ./bin/output ./cmd/talison_contabil/main.go

start: build
	./bin/output

docker:
	docker build -t backend -f Dockerfile .


deploy: docker
	docker run --rm -d -p 3033:3033/tcp backend:latest