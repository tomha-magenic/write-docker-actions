FROM node:slim

COPY package*.json ./

RUN npm install

COPY . .

CMD ["node", "/scr/index.js"]