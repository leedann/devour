FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD devoursvr devoursvr
EXPOSE 443
ENTRYPOINT ["/devoursvr"]
