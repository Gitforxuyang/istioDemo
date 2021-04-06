FROM golang:1.16
RUN mkdir -p /app
# 以上部分可以弄成一个通用基础镜像
WORKDIR /app
##暴露端口
EXPOSE 50001
EXPOSE 8080
COPY . /app/
#最终运行docker的命令
ENTRYPOINT  ["./bin/server"]