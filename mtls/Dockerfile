FROM golang:1.20rc3-alpine3.17 as builder
RUN apk add --update --no-cache ca-certificates git
COPY . /server
WORKDIR /server
RUN mkdir bin -p && go build -o bin/main ./cmd/server/main.go

FROM alpine:3.17.1
RUN adduser --disabled-password --no-create-home antis
COPY --from=builder /server/bin/main /server/main
RUN chown -R antis /server
WORKDIR /server
USER antis
EXPOSE 2023
ENTRYPOINT [ "./main" ]
