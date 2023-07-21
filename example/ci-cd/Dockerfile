FROM golang:1.19-alpine

ARG commit_id

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download && go mod verify

COPY . .
RUN go build -ldflags "-X main.CommitID=$commit_id"  -v -o app  .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /usr/src/app/app ./
EXPOSE 8080
CMD ["./app"]