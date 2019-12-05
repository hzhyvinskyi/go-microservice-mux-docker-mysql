FROM tailor/docker-libvips

ENV PORT 8080

EXPOSE $PORT

COPY . /

CMD ["/bin/app"]