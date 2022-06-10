# Stage 0 : Build the C library 

FROM debian AS lib_builder

WORKDIR /foundry

RUN apt-get update -y && \
  #DEBIAN_FRONTEND=noninteractive TZ=Etc/UTC apt-get -y install tzdata -y \
  apt-get install -y \
  build-essential \
  cmake \
  git

RUN git clone https://github.com/jgarff/rpi_ws281x.git \
  && cd rpi_ws281x \ 
  && mkdir build \
  && cd build \ 
  && cmake -D BUILD_SHARED=OFF -D BUILD_TEST=OFF .. \
  && cmake --build . \
  && make install

# Stage 1 : Build the go executable with the ws2811 library

FROM golang:latest AS go_builder

COPY --from=lib_builder /usr/local/lib/libws2811.a /usr/local/lib/
COPY --from=lib_builder /usr/local/include/ws2811 /usr/local/include/ws2811

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY pkg/ ./pkg/

RUN go build -o ./spectrum ./cmd/spectrum/

# Stage 2 : Build the final image with the go executable

FROM debian:stable-slim
WORKDIR /app

COPY --from=go_builder app/spectrum spectrum

CMD [ "./spectrum" ]