run:
	go run main.go

build:
	docker build --no-cache -t takecy/go-mongo-sample .

run_d: build
	docker run --rm --name main_go --net golangmongosample_default takecy/go-mongo-sample
