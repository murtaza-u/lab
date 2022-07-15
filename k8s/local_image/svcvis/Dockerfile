FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
COPY . ./svcvis
RUN cd svcvis && go build

FROM alpine:latest
COPY --from=builder /go/svcvis/svcvis svcvis
EXPOSE 8000
ENTRYPOINT [ "./svcvis" ]
