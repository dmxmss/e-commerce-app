FROM node:20-alpine AS builder

WORKDIR /app

COPY ./admin/package*.json ./

RUN npm ci

COPY ./admin/ .

RUN npm run build

FROM nginx:alpine AS prod

COPY --from=builder /app/dist/ /usr/share/nginx/html/
COPY ./admin/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

FROM node:20-alpine AS dev

WORKDIR /app

COPY ./admin/package*.json .

RUN npm ci

COPY ./admin .

EXPOSE 5173

CMD ["npm", "run", "dev"]
