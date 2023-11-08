# Use an official Go runtime as a parent image
FROM golang:1.16

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go application
RUN go build -o echo-hello-world

# Expose port 1323 (the port your Echo application will run on)
EXPOSE 1323

# Define the command to run your application
CMD ["./echo-hello-world"]
