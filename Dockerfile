# build stage
FROM golang:1.13

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
 && echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /app
ADD ./go.mod /app
ADD ./go.sum /app

RUN export GOPROXY=https://goproxy.cn && go mod download

ADD . /app

RUN chmod a+w /app/var

RUN CGO_ENABLED=0 go build -o culture

EXPOSE 9000
EXPOSE 8080
CMD /app/culture