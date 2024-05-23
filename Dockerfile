FROM golang:latest AS dev
ARG USER=""
RUN useradd -m ${USER}
USER ${USER}
