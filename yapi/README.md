### 文件说明
```
-|
 |-config.json：配置文件，镜像打包前需要根据实际环境进行修改
 |-Dockerfile：镜像打包Dockerfile
 |-Makefile：简化本地操作
 |-README.md：说明文档
```

### Dockerfile说明
```dockerfile
FROM node:10.12

EXPOSE 80

RUN mkdir /yapi  && cd yapi
# clone项目源码
RUN git clone https://github.com/YMFE/yapi.git vendors
RUN cd vendors
# 切换到v1.3.23，目前最新的发行版
RUN git checkout v1.3.23
RUN npm i -g node-gyp yapi-cli npm@latest
RUN npm i --production

WORKDIR /yapi/vendors
# 拷贝配置文件
COPY config.json /yapi
# 容器入口命令
CMD node server/app.js
```

### 部署指南
1. 服务依赖于`mongodb`，因此部署之前需要先部署`mongodb`服务
2. 修改`config.json`的管理员账号、数据库和邮件的配置信息
3. 切换到当前目录，执行镜像打包命令
```sh
$ docker build -t yapi:1.0.0 .
```
1. 执行初始化操作，该操作会初始化数据库包括创建管理员账号，不可重复执行
```sh
$ docker run --rm yapi:1.0.0 node server/install.js 
```

5. 启动服务
```sh
$ docker run -d --name yapi -p 8080:80 yapi:1.0.0 
```

### 操作指南
`yapi`支持接口管理、mock服务和自动化测试

详细使用指南，移步[官方使用教程](https://yapi.ymfe.org/documents/index.html)