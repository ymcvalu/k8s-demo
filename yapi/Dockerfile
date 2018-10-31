FROM node:10.12

EXPOSE 80

RUN mkdir /yapi \
  && cd yapi \
  && git clone https://github.com/YMFE/yapi.git vendors \
  && cd vendors \
  && git checkout v1.3.23 \
  && npm i -g node-gyp yapi-cli npm@latest \
  && npm i --production

WORKDIR /yapi/vendors

COPY config.json /yapi

CMD node server/app.js
