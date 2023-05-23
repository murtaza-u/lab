FROM golang:1.19.3-alpine3.16 as builder
RUN apk add --update --no-cache ca-certificates git
COPY . /datastore
WORKDIR /datastore
RUN go build -o main

FROM alpine:3.16.2
LABEL maintainer="murtaza@murtazau.xyz"
RUN adduser --disabled-password --no-create-home murtaza
COPY --from=builder /datastore/main /datastore/main
RUN chown -R murtaza /datastore
WORKDIR /datastore
USER murtaza
EXPOSE 8080
ENTRYPOINT [ "./main" ]
