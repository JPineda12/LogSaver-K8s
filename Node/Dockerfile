FROM node:17-alpine3.14 as build

WORKDIR /gRPC-client
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
RUN node api_index.js