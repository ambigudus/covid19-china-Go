FROM golang

WORKDIR /go/src/github.com/ambigudus/covid19-china-Go
ENV GOPROXY=https://goproxy.cn
COPY go.mod .

COPY go.sum .

RUN GO111MODULE=on go mod download

COPY . .
CMD ["go", "run", "main.go"]
