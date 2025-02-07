FROM redis:6.2

EXPOSE 6379

VOLUME /data

ENV ALLOW_EMPTY_PASSWORD=yes

COPY start.sh /start.sh
RUN chmod +x /start.sh

CMD ["/start.sh"]
