FROM golang as builder

WORKDIR /sekura

RUN mkdir build/

COPY go.mod go.sum .
ADD tools/ tools/

RUN go mod download

ADD pkg/ pkg/
ADD cmd/ cmd/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/gateway-agent ./cmd/agent
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/gateway ./cmd/aio
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/gateway-ui ./cmd/ui
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/gateway-login ./cmd/login


FROM scratch as agent-gateway

WORKDIR /sekura

COPY --from=builder /sekura/build/gateway-agent /sekura/gateway-agent
COPY --from=alpine:3 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./gateway-agent" ]
CMD [ "-h" ]


FROM scratch as ui-gateway

WORKDIR /sekura

COPY --from=builder /sekura/build/gateway-ui /sekura/gateway-ui
COPY --from=alpine:3 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./gateway-ui" ]
CMD [ "-h" ]


FROM scratch as login-gateway

WORKDIR /sekura

COPY --from=builder /sekura/build/gateway-login /sekura/gateway-login
COPY --from=alpine:3 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./gateway-login" ]
CMD [ "-h" ]


FROM scratch as gateway

WORKDIR /sekura

COPY --from=builder /sekura/build/gateway /sekura/gateway
COPY --from=alpine:3 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./gateway" ]
CMD [ "-h" ]
