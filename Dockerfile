FROM progrium/busybox
RUN opkg-install ca-certificates curl
COPY deflektor /staticbinary
COPY entrypoint /entrypoint
ENTRYPOINT ["sh", "-c", "eval $(cat /entrypoint)"]
EXPOSE 8081
