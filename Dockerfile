FROM ubuntu:latest
LABEL authors="ereng"

ENTRYPOINT ["top", "-b"]