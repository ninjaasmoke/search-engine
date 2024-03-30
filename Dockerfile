# # Use the official Golang image as a base
# FROM golang:latest

# # Set the Current Working Directory inside the container
# WORKDIR /app

# # Copy the source code and JSON files into the container
# COPY go.mod go.sum ./
# COPY main.go ./
# COPY document_info_map.json ./
# COPY final_inverted_index.json ./

# # Fetch dependencies
# RUN go mod download
# RUN go get search-server
# RUN go build -o search_engine .

# # Install Node.js and npm
# RUN apt-get update && apt-get install -y nodejs npm

# WORKDIR /app/front
# RUN npm install
# RUN npm run build
# RUN cp -r /app/front/dist /app/dist

# # Expose port 8080 to the outside world (if your Go server listens on this port)
# EXPOSE 8080

# # Command to run the Go executable
# CMD ["./search_engine"]

# Stage 1: Build the frontend assets
FROM node:latest AS frontend-builder

WORKDIR /app/front

COPY front/package.json front/package-lock.json ./
RUN npm install
COPY front/. .
RUN npm run build

# Stage 2: Build the Go binary
FROM golang:latest AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY document_info_map.json ./
COPY final_inverted_index.json ./

RUN go build -o search_engine .

# Stage 3: Final image
FROM golang:latest

WORKDIR /app

COPY --from=backend-builder /app/search_engine ./
COPY --from=frontend-builder /app/front/dist ./dist
COPY --from=backend-builder /app/document_info_map.json ./
COPY --from=backend-builder /app/final_inverted_index.json ./

EXPOSE 8080

CMD ["./search_engine"]
