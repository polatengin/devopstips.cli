FROM golang:1.16.2-buster AS build

WORKDIR /src

COPY ./ /src

RUN go build -o devopstips .

FROM gcr.io/distroless/static AS runtime

WORKDIR /app

COPY --from=build /src/devopstips /app/devopstips

ENTRYPOINT [ "/app/devopstips" ]
