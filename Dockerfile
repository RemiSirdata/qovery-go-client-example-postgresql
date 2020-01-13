FROM golang:stretch


WORKDIR /app
COPY . /app
RUN go get && go build -o server postgresql.go

# Choose the port to publicly expose to the internet
EXPOSE 8080

# Create a dedicated user to run your app with (for security reasons)
RUN useradd -ms /bin/bash qovery
USER qovery

CMD /app/server -bind=":8080"
