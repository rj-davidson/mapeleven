FROM node:alpine

WORKDIR /app/client/

COPY package.json /app/client/

RUN yarn install

COPY . /app/client/

CMD ["yarn", "run", "dev"]