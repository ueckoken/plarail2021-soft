# Build the source code
FROM node:16.13.0-buster-slim AS builder
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
ENV NODE_ENV production
RUN yarn build

# Production image, copy all the files and run nginx
FROM nginx:1.20.1
COPY --from=builder /app/out/ /usr/share/nginx/html/