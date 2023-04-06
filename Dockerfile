FROM node:alpine
WORKDIR /app
COPY client/package.json .
RUN yarn install
COPY basicwebsite/simple-vite-react-project .
CMD ["yarn", "run", "dev"]
