all:

build:
	go install -v

test: example
	go test -v ./...

gen:
	go generate

example: gen build
	tuk-gen model -o ./example -p model ./example/models.yaml

