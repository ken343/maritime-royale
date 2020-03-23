# Need to Install Go, sdl2, & Git (I think git comes with Go install)
FROM golang:1.14.1-buster

RUN apt-get update && apt-get -y install git \
    libsdl2-dev \
    libsdl2-image-dev \
    libsdl2-mixer-dev \
    libsdl2-ttf-dev \
    libsdl2-gfx-dev

RUN go version
RUN git version
