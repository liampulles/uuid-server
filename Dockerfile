FROM scratch
COPY uuid-server .
ENTRYPOINT ["/uuid-server"]