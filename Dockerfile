ARG image=golang
ARG tag=latest

FROM ${image}:${tag} AS dev
ARG USER=""
RUN useradd -m ${USER}
USER ${USER}

FROM ${image}:${tag} AS build
WORKDIR /build
COPY . .
RUN make all

FROM scratch as getfishy
COPY --from=build /build/getfishy /getfishy
ENTRYPOINT [ "/getfishy" ]

FROM scratch as platter
COPY --from=build /build/platter platter/
ENTRYPOINT [ "/platter" ]
