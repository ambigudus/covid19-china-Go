FROM golang as build

WORKDIR /go/src/github.com/ambigudus/covid19-china-Go
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

ADD go.mod .
ADD go.sum .
RUN go mod download


COPY . .
RUN  GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go
FROM scratch as prod

CMD ["go", "run", "main.go"]