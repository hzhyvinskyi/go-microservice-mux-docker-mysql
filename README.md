# Golang webp converter microservice.
Using C library libvips.

### Docker:
##### docker build -t myapp -f Dockerfile .
##### docker run -p 8080:8080 myapp

bin and tmp folders must be created in root
##### For compiling use:
go build -o ./bin/app ./cmd/webp-converter/main.go
##### Execute program from root:
./bin/app
