FROM golang:1.13-alpine AS build-env
ENV CGO_ENABLED 0
ADD ../../. /go/src/toury_bakcend
WORKDIR /go/src/toury_bakcend
RUN go build -o /bin/toury_api ./web_api/api.go

FROM alpine:3.10
WORKDIR /
COPY --from=build-env /bin/toury_api /toury_api
CMD ["/server"]
