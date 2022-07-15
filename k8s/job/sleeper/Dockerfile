FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
COPY . ./sleeper
RUN cd sleeper && go build

FROM alpine:latest
COPY --from=builder /go/sleeper/sleeper .
EXPOSE 8000
ENTRYPOINT [ "./sleeper" ]
