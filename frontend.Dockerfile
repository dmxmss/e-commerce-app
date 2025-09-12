# builder
FROM node:20-alpine AS builder

WORKDIR /app

COPY ./frontend/package*.json ./

RUN npm ci

COPY ./frontend/ ./

RUN npm run build

# prod
FROM nginx:alpine AS prod

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

# dev
FROM node:20-alpine AS dev

WORKDIR /app

COPY frontend/package*.json ./

RUN npm install

COPY frontend/ .

EXPOSE 5173

CMD ["npm", "run", "dev"]
