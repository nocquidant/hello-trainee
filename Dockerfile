# Multi stage build
#  - We don't need go installed once our app is compiled
#  - Leaving the build image

FROM golang:1.10 AS build
ADD . /src
WORKDIR /src
RUN go get -d -v -t 
RUN go test --cover -v ./... --run UnitTest 
RUN go build -v -o hello-trainee


FROM alpine:3.4 

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 

EXPOSE 8488  
CMD ["hello-trainee"] 
HEALTHCHECK --interval=10s CMD wget -qO- localhost:8488/hello 

COPY --from=build /src/hello-trainee /usr/local/bin/hello-trainee

RUN chmod +x /usr/local/bin/hello-trainee