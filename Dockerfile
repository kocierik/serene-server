FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y yt-dlp libimage-exiftool-perl id3v2
RUN wget https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux && \
    chmod +x yt-dlp_linux && mv yt-dlp_linux /usr/bin/yt-dlp

RUN go build cmd/main.go

EXPOSE 4000

CMD ["./main"]
