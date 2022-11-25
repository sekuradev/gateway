FROM golang as builder

WORKDIR /sekura

RUN mkdir build/

COPY go.mod go.sum .
ADD tools/ tools/

RUN go mod download

ADD pkg/ pkg/
ADD cmd/ cmd/

RUN go build -o ./build/gateway-agent ./cmd/agent
RUN go build -o ./build/gateway ./cmd/aio
RUN go build -o ./build/gateway-ui ./cmd/ui
RUN go build -o ./build/gateway-login ./cmd/login


FROM scratch as agent-gateway

WORKDIR /sekura

COPY --from=builder /sekura/build/gateway /sekura/gateway-agent
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./gateway-agent" ]


FROM scratch as ui-gateway

WORKDIR /sekura

COPY --from=builder /sekura/build/gateway /sekura/gateway-ui
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./gateway-ui" ]


FROM scratch as login-gateway

WORKDIR /sekura

COPY --from=builder /sekura/build/gateway /sekura/gateway-login
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./gateway-login" ]


FROM scratch as gateway

WORKDIR /sekura

COPY --from=builder /sekura/build/gateway /sekura/gateway
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./gateway" ]
