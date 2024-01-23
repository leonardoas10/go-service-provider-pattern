FROM golang:1.19-alpine

RUN apk --no-cache add git
RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

# Copy the rest of the application files
COPY ./ ./

ENTRYPOINT CompileDaemon --build="go build -o build/goapp ./src/cmd/main" -command="./build/goapp" -build-dir=/app
