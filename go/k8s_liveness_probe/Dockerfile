FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
COPY . ./liveness
RUN cd liveness && go build

FROM alpine:latest
COPY --from=builder /go/liveness/liveness .
EXPOSE 8000
ENTRYPOINT [ "./liveness" ]
