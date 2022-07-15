FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
COPY . ./ssd-monitor
RUN cd ssd-monitor && go build

FROM alpine:latest
COPY --from=builder /go/ssd-monitor/ssd-monitor .
EXPOSE 8000
ENTRYPOINT [ "./ssd-monitor" ]
