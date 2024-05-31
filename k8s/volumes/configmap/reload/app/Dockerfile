FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
COPY . ./reload
RUN cd reload && go build -o main

FROM alpine:latest
COPY --from=builder /go/reload/main main
COPY --from=builder /go/reload/env env
ENTRYPOINT [ "./main" ]
