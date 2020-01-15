FROM golang:1.12-alpine AS development

ENV PROJECT_PATH=/mxprotocol-server
ENV PATH=$PATH:$PROJECT_PATH/build
ENV CGO_ENABLED=1
ENV GO_EXTRA_BUILD_ARGS="-a -installsuffix cgo"

RUN apk add --no-cache ca-certificates make git bash protobuf alpine-sdk nodejs nodejs-npm

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN make dev-requirements clean all

FROM alpine:latest AS production

WORKDIR /root/
RUN apk --no-cache add ca-certificates
RUN mkdir /etc/mxprotocol-server
COPY --from=development /mxprotocol-server/m2m/build/ .
COPY --from=development /mxprotocol-server/configuration/ .
COPY --from=development /mxprotocol-server/m2m/scripts/init .
RUN ["chmod", "+x", "./start"]
ENTRYPOINT ["./start"]
