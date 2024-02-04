FROM alpine:3.8.4 as root-certs
RUN apk add -U --no-cache ca-certificates 
RUN addgroup -g 1001 scope
RUN adduser scope -u 1001 -D -G scope /home/scope

FROM golang:1.21 as builder
WORKDIR /scoperunner-wkdir
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY ./scope-server-config.yaml /scoperunner-wkdir/scope-server-config.yaml
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./scoperunner-server 

FROM scratch as final
COPY --from=root-certs  /etc/passwd /etc/passwd
COPY --from=root-certs  /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --chown=1001:1001 --from=builder /scoperunner-wkdir/scoperunner-server /scoperunner-server

EXPOSE 8000

USER scope

ENTRYPOINT [ "/scoperunner-server" ]