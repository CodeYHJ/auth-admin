FROM node:latest-alpine as NodeBuilder

WORKDIR /app

COPY ./web /app/

RUN cd web && yarn && yarn install && yarn build

FROM golang:1.17-alpine as GoBuilder

WORKDIR /app

RUN apk add --no-cache git

# 配置模块代理
# 国内源打包专用
# ENV GO111MODULE=on \
#     GOPROXY=https://goproxy.cn,direct

# github action 专用
ENV GO111MODULE=on

COPY . /app/

COPY --from=NodeBuilder /app/dist /app/dist

RUN go mod download && go mod verify && go build -v -o go_server



FROM golang:1.17-alpine as runner

ENV PATH /app/bin:$PATH

ENV APP_PATH /app

RUN mkdir -p "$APP_PATH" "$APP_PATH/bin" && chmod -R 777 "$APP_PATH"

COPY --from=GoBuilder /app/go_server /app/bin/go_server

WORKDIR $APP_PATH

CMD ["go_server"]