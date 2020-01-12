FROM golang:stretch

RUN go get && go build


RUN mkdir -p /app
COPY qovery-go-client-example-postgresql /app/server
WORKDIR /app

# Choose the port to publicly expose to the internet
EXPOSE 8080

# Create a dedicated user to run your app with (for security reasons)
RUN useradd -ms /bin/bash qovery
USER qovery

CMD app/server -bind=":8080"