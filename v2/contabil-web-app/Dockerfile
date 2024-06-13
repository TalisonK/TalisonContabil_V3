FROM node:20.10.0 as build

WORKDIR /app

COPY . .

RUN npm install --loglevel=error

RUN npm run build

FROM nginx:stable

COPY default.conf /etc/nginx/conf.d

COPY --from=build /app/build /usr/share/nginx/html

CMD ["chmod", "777", "/usr/share/nginx/html/logoReto"]

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]