FROM scratch
ADD go-poi go-poi
ADD index.html index.html

COPY data data
VOLUME ["/data"]

# Runtime defaults
EXPOSE 8000
CMD ["/go-poi"]
