FROM node:16-alpine3.14

WORKDIR /src

COPY ./ /src

RUN npm install

ENTRYPOINT [ "npx", "ts-node", "./cli.ts" ]
