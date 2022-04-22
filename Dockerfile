FROM golang:1.16-alpine3.14

RUN apk add --no-cache build-base \
    mupdf mupdf-dev \
    freetype freetype-dev \
    harfbuzz harfbuzz-dev \
    jbig2dec jbig2dec-dev \
    jpeg jpeg-dev \
    openjpeg openjpeg-dev \
    zlib zlib-dev

WORKDIR /app

COPY . .

RUN export CGO_LDFLAGS="-lmupdf -lm -lmupdf-third -lfreetype -ljbig2dec -lharfbuzz -ljpeg -lopenjp2 -lz" \
    && go mod download \
    && go build -o /pendown-be

EXPOSE 8080

CMD [ "/pendown-be" ]