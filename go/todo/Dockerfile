FROM golang:1.20rc3-alpine3.17 as builder
RUN apk add --update --no-cache ca-certificates git
COPY . /todo
WORKDIR /todo
RUN mkdir bin -p && go build -o bin/main ./cmd/api/main.go

FROM alpine:3.17.1
RUN adduser --disabled-password --no-create-home antis
COPY --from=builder /todo/bin/main /todo/main
RUN chown -R antis /todo
WORKDIR /todo
USER antis
EXPOSE 5000
ENTRYPOINT [ "./main" ]
