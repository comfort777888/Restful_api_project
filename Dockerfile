FROM golang:1.17-alpine

# Add Maintainer Info
LABEL maintainer="Araya"

RUN mkdir restful-api

# Set the Current Working Directory inside the container
WORKDIR /restful-api

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN go build -o main ./main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/restful-api/main"]