FROM node:18-alpine AS build

# Create app directory
WORKDIR /app

# Install app dependencies

COPY package*.json ./
RUN npm install
RUN npx ngcc --properties es2023 browser module main --first-only --create-ivy-entry-points

COPY . .

RUN npm run build

FROM nginx:stable

COPY default.conf /etc/nginx/conf.d

COPY --from=build /app/dist/contabil-web-app/ /usr/share/nginx/html

EXPOSE 80









#FROM node:18-alpine AS build
#
## Create app directory
#WORKDIR /app
#
## Install app dependencies
#COPY . .
#
#RUN rm -r dist
#
#RUN npm install
#
#RUN npm run build
#
#
## Bundle app source
#FROM nginx:alpine
#
#COPY ./nginx.conf /etc/nginx/conf/nginx.conf
#
#COPY --from=build /app/dist/contabil-web-app/ /usr/share/nginx/html
#
#EXPOSE 80
