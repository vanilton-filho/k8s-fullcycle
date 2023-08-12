FROM golang:1.15
COPY /src .
RUN go build -o server .
CMD [ "./server" ]
