FROM golang:1.15 as build
LABEL maintainer="117503445"
ENV GO111MODULE=on
WORKDIR /root/project
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag i
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o server_bin
FROM alpine as prod
EXPOSE 80
COPY --from=build /root/project/server_bin /root/server_bin
WORKDIR /root
ENTRYPOINT /root/server_bin