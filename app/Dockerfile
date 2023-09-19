FROM golang:1.20 AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o sre-hiring .

FROM scratch
COPY --from=build /app/sre-hiring /sre-hiring

EXPOSE 8080
CMD ["/sre-hiring"]

