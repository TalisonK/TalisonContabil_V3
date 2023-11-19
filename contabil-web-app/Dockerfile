FROM node:18-alpine AS build

# Create app directory
WORKDIR /app

# Install app dependencies
COPY . .

RUN npm install

RUN npm run build


# Bundle app source
FROM nginx:alpine

COPY --from=build /app/dist/contabil-web-app/ /usr/share/nginx/html

EXPOSE 80
