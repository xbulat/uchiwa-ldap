FROM golang:1.9.2-alpine AS build-env

WORKDIR /go/src/uchiwa

RUN apk add --no-cache nodejs-npm git

COPY . .

RUN go get                              && \
    go build -o /build/uchiwa .         && \
    cd /go/src/github.com/sensu/uchiwa/ && \
    npm install --production --unsafe-perm

FROM alpine

LABEL version="1.7.0" \
      description="Uchiwa LDAP - dashboard for the Sensu monitoring framework" \
      maintainer="xbulat at github.com"

COPY --from=build-env /build/uchiwa /bin
COPY --from=build-env /go/src/github.com/sensu/uchiwa/public /opt/uchiwa/public/

EXPOSE 3000

CMD ["/bin/uchiwa"]
