FROM golang:1.18-alpine as builder
RUN apk add --no-cache tzdata
WORKDIR $GOPATH/src/go.k6.io/k6
ADD . .
RUN apk --no-cache add build-base git
RUN go install go.k6.io/xk6/cmd/xk6@latest
RUN CGO_ENABLED=0 xk6 build \
    --with github.com/TylerVolfgant/xk6-basboy@latest --output /tmp/k6
RUN cp /tmp/k6 /usr/bin/k6
WORKDIR /home/k6
ENTRYPOINT ["tail", "-f", "/dev/null"]
