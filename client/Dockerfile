FROM node:alpine

WORKDIR /app/client/

COPY package.json /app/client

RUN yarn install

COPY . /app/frontend/

CMD ["yarn", "run", "dev"]