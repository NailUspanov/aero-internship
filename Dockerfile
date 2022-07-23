FROM golang:latest
WORKDIR /aero-internship
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["make", "run"]