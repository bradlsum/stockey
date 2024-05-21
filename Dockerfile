FROM golang:1.22.1 AS dev
ARG USER=""
RUN useradd -m ${USER}
USER ${USER}
