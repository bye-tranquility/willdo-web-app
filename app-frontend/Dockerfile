FROM node:20-alpine

RUN addgroup -S app_user_group && \
    adduser -S app_user -G app_user_group
WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install && \
    npm cache clean --force
COPY . .

RUN chown -R app_user:app_user_group /app
USER app_user

EXPOSE 3000
ENTRYPOINT ["npm", "run", "dev"]