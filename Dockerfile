FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags "-s -w" -o /Pentagon

FROM scratch
COPY --from=builder /Pentagon /
HEALTHCHECK --interval=10s --timeout=10s --start-period=20s --retries=3 CMD ["/hotspots"]
CMD ["/Pentagon"]
