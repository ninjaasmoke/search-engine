# Use the official Golang image as a base
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code and JSON files into the container
COPY go.mod go.sum ./
COPY main.go ./
COPY document_info_map.json ./
COPY final_inverted_index.json ./

# Fetch dependencies
RUN go mod download
RUN go get search-server
RUN go build -o search_engine .

# Expose port 8080 to the outside world (if your Go server listens on this port)
EXPOSE 8080

# Command to run the Go executable
CMD ["./search_engine"]
