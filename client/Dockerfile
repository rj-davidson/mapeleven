FROM node:alpine
WORKDIR /app
COPY package.json .
RUN yarn install
COPY basicwebsite/simple-vite-react-project .
CMD ["yarn", "run", "dev"]
