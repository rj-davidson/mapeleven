FROM node:alpine
WORKDIR /app/
COPY client/package.json ./
RUN yarn install
COPY client/ ./
CMD ["yarn", "run", "dev"]
