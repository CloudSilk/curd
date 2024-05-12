FROM registry.cn-shanghai.aliyuncs.com/swtsoft/golang-build:1.20.0-alpine3.17 as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace
COPY . .
ARG ARG_GIT_USERNAME
ARG ARG_GIT_PASSWORD
RUN git config --global url."https://${ARG_GIT_USERNAME}:${ARG_GIT_PASSWORD}@codeup.aliyun.com".insteadOf "https://codeup.aliyun.com"
RUN go env && go build -o curd main.go

FROM registry.cn-shanghai.aliyuncs.com/swtsoft/golang-run:alpine-3.16.0
LABEL MAINTAINER="guoxf@swtsoft.com"

ENV DUBBO_GO_CONFIG_PATH="./dubbogo.yaml"

WORKDIR /workspace
COPY --from=builder /workspace/curd ./

EXPOSE 20000

ENTRYPOINT ./curd
