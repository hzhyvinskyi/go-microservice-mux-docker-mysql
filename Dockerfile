FROM scratch

ENV PORT 8080

EXPOSE $PORT

COPY ./bin/app /

COPY .env /

CMD ["/app"]