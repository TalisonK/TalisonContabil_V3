FROM node:18-alpine

# Create app directory
WORKDIR /app

# Install app dependencies

COPY package*.json ./

RUN npm install

RUN npm install -g @angular/cli@16.2.7


# Bundle app source

COPY . .

EXPOSE 4200

RUN ["npm", "run", "disable-analytics"]
CMD [ "npm", "start" ]
