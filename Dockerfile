FROM golang:1.18 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

workdir /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/cinemarch -v ./cmd/cinemarch


FROM scratch as prod
COPY --from=builder /app/bin/cinemarch /cinemarch
CMD ["/cinemarch"]
