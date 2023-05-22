FROM ubuntu:latest

EXPOSE 8080

COPY start.sh start.sh

CMD ["/bin/bash","start.sh"]
