FROM node:8.4.0

COPY . /var/www/node/vue-frontend

RUN cd /var/www/node/vue-frontend && npm install

EXPOSE 1337

WORKDIR /var/www/node/vue-frontend

CMD ["npm", "run", "dev"]
