build:
	go build .
rel:
	go build -ldflags="-w -s" .