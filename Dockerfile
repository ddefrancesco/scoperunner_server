FROM alpine:3.8.4 as root-certs
RUN apk add -U --no-cache ca-certificates 
RUN addgroup -g 1001 scope
RUN adduser scope -u 1001 -D -G scope /home/scope

FROM golang:1.21 as builder
WORKDIR /scoperunner-wkdir
RUN mkdir -p /opt/scope/
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./scoperunner-server 

FROM alpine:3.19 as final

COPY --from=root-certs  /etc/passwd /etc/passwd
COPY --from=root-certs  /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --chown=1001:1001 --from=builder /scoperunner-wkdir/scoperunner-server /scoperunner-server
COPY --chown=1001:1001 --from=builder /opt/scope/ /opt/scope
COPY --chown=1001:1001 --from=builder /scoperunner-wkdir/scope-server-config.yaml /opt/scope/scope-server-config.yaml
COPY --chown=1001:1001 --from=builder /scoperunner-wkdir/openngc/csv/NGC.csv /opt/scope/NGC.csv

RUN apk add --no-cache tzdata
RUN apk add --no-cache curl

ENV TZ=Europe/Rome

ENV SCOPE_SERIALPORT=/dev/ttyUSB0
ENV SCOPE_ENVIRONMENTS_FAKESCOPE=false
ENV SCOPE_OPENNGC_CSV_PATH=/opt/scope/NGC.csv

EXPOSE 8000
EXPOSE 9999

USER root

HEALTHCHECK --interval=10s --timeout=30s CMD ./scoperunner-server healthcheck
ENTRYPOINT [ "/scoperunner-server" ]