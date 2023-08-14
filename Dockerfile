ARG GO_VERSION=1.21

# Build stage
FROM golang:${GO_VERSION} AS builder

ENV CGO_ENABLED=0

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN make build
RUN echo "nonroot:x:65534:65534:Non root:/:" > /etc_passwd


# Final stage
FROM scratch

COPY --from=builder /go/bin/go-template-cli /bin/go-template-cli
COPY --from=builder /etc_passwd /etc/passwd

USER nonroot

ENTRYPOINT [ "go-template-cli" ]
CMD [ "version" ]