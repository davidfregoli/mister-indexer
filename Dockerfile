FROM alpine
WORKDIR /app
COPY bin/mr-indexer ./
CMD ["/app/mr-indexer"]