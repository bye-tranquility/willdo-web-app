FROM nginx:alpine

RUN mkdir -p /var/cache/nginx /var/log/nginx /var/run && \
    chown -R nginx:nginx /var/cache/nginx /var/log/nginx /var/run /etc/nginx && \
    rm -rf /var/cache/apk/*

USER nginx

EXPOSE 8080

ENTRYPOINT ["nginx", "-g", "daemon off;"]