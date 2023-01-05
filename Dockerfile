FROM golang:1.19.3-alpine3.16 AS go-builder

ENV APPNAME=platined

RUN ARCH=`uname -m`; echo ${ARCH};
RUN apk add --no-cache ca-certificates build-base git linux-headers;

WORKDIR /code

COPY . /code/

RUN BUILD_TAGS=muslc LINK_STATICALLY=true make build

RUN echo "Ensuring binary is statically linked ..." \
  && (file /code/build/$APPNAME | grep "statically linked")

FROM alpine:3.16

WORKDIR /chain

ENV APPNAME=platined

COPY --from=go-builder /code/build/$APPNAME /usr/bin/$APPNAME

COPY ./scripts/docker/* /opt/

# Fix windows issue
RUN dos2unix /opt/*.sh

RUN chmod +x /opt/*.sh

# rest server
EXPOSE 1317
# grpc
EXPOSE 9090
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["sh", "-c", "/usr/bin/$APPNAME", "version"]

