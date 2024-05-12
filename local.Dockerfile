FROM guoxf/golang-run:alpine-3.13.5
LABEL MAINTAINER="guoxf@swtsoft.com"

ENV DUBBO_GO_CONFIG_PATH="./dubbogo.yaml"

WORKDIR /workspace
COPY curd ./
COPY docs/swagger.json ./docs
COPY docs/swagger.yaml ./docs

EXPOSE 20000

ENTRYPOINT ./curd
