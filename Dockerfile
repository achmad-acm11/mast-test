FROM golang:1.21.12-alpine3.20 AS builder
RUN go env -w GO111MODULE=on
WORKDIR /mast-integrator
COPY ./    ./
RUN CGO_ENABLED=0 GOOS=linux go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM rockylinux:8 as golang-app
RUN yum update -y && yum install -y ca-certificates
RUN yum install -y bash git curl wget unzip rpm
RUN yum install -y glibc.i686 zlib.i686 libstdc++.i686
RUN wget https://dl.google.com/android/repository/build-tools_r28.0.2-linux.zip
RUN unzip build-tools_r28.0.2-linux.zip
RUN chmod -R 755 android-9/
RUN cp /android-9/aapt /usr/local/bin/aapt && cp /android-9/lib64/libc++.so /usr/lib64/libc++.so
RUN wget https://vault.centos.org/centos/8/AppStream/x86_64/os/Packages/libplist-2.0.0-10.el8.x86_64.rpm
RUN wget https://vault.centos.org/8.5.2111/AppStream/Source/SPackages/libplist-2.0.0-10.el8.src.rpm
RUN rpm -i libplist-2.0.0-10.el8.src.rpm && rpm -i libplist-2.0.0-10.el8.x86_64.rpm
#RUn apt-get install libplist-utils
WORKDIR /root/
RUN mkdir "tmp" && mkdir "apk_icon"
COPY --from=builder /mast-integrator ./
CMD ["./main"]
