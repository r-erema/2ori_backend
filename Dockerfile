FROM golang:1.13-alpine AS build-env
ENV CGO_ENABLED 0
ADD . /go/src/toury_bakcend
ENV GOPATH "/go/src/toury_bakcend"
WORKDIR /go/src/toury_bakcend
RUN go build -gcflags "all=-N -l" -o /bin/toury_api ./web_api/api.go

# Compile Delve
RUN apk add --no-cache git
RUN go get github.com/derekparker/delve/cmd/dlv

FROM alpine:3.10

RUN apk add --no-cache libc6-compat

WORKDIR /

COPY --from=build-env /go/src/toury_bakcend/bin/dlv /
COPY --from=build-env ./bin/toury_api /

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/toury_api"]
#CMD ["/server"]
