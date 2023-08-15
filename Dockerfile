FROM golang:latest AS builder

RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /Receipts
COPY go.mod .
RUN go mod download
COPY . .
RUN go install

FROM scratch
COPY --from=builder /Receipts .
ENTRYPOINT ["./receipts_by_fcph"]

#CMD ["./main"]
# docker build -t myapp .
# dsudo s