FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /receipts_by_fcph
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o receipts_by_fcph .
#RUN go build -o receipts_by_fcph
FROM scratch
COPY --from=builder /receipts_by_fcph .

EXPOSE 9095

ENTRYPOINT ["./receipts_by_fcph"]

#CMD ["./main"]
# docker build -t myapp .
# dsudo s