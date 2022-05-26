FROM golang:alpine AS build-env
WORKDIR /provider-pattern
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY go.mod /provider-pattern/go.mod
COPY go.sum /provider-pattern/go.sum
RUN go mod download
COPY . /provider-pattern
RUN CGO_ENABLED=0 GOOS=linux go build -o build/provider-pattern ./src/cmd/main


FROM alpine
COPY ./src/cmd/main/.env /
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /provider-pattern/build/provider-pattern /

EXPOSE 3000

ENTRYPOINT ["/provider-pattern"]