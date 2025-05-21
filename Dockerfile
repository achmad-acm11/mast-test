FROM golang:1.21.12-alpine3.20 AS builder
RUN go env -w GO111MODULE=on
WORKDIR /mast-integrator
COPY ./    ./
RUN CGO_ENABLED=0 GOOS=linux go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.20 as aapt-downloader
RUN apk update && apk add --no-cache \
    bash \
    curl \
    wget \
    unzip \
    libc6-compat \
    zlib \
    libstdc++ \
    ca-certificates \
    openjdk17-jdk
WORKDIR /opt
RUN wget https://dl.google.com/android/repository/build-tools_r28.0.2-linux.zip && \
    unzip build-tools_r28.0.2-linux.zip && \
    chmod -R 755 android-9 && \
    mv android-9 /opt/build-tools

FROM alpine:3.20 as golang-app
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
RUN apk add --no-cache git
COPY --from=aapt-downloader /opt/build-tools/aapt ./usr/local/bin/aapt
RUN chmod 755 /usr/local/bin/aapt
# ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
# ENV ZONEINFO /zoneinfo.zip
WORKDIR /root/
COPY --from=builder /mast-integrator ./
#COPY --from=builder /sca-integrator/_public_key.pem ./
#RUN mkdir "_scanned-project-files"
#RUN mkdir "_project-repository"
CMD ["./main"]
