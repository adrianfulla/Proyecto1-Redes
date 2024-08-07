# syntax=docker/dockerfile:1

# Stage 1: Build the Go server
FROM golang:latest as server-build

# Set the working directory
WORKDIR /app/

COPY . .

RUN go mod download

# Stage 2: Build the frontend
FROM node:16 as frontend-build

# Set the working directory
WORKDIR /app/frontend

# Copy package.json and package-lock.json files
COPY frontend/package*.json ./

# Install dependencies and build the frontend
RUN npm cache clean --force
RUN npm install
COPY frontend/ ./
RUN npm run build

# Stage 3: Final stage - Copy built artifacts and run the server
FROM golang:latest

RUN apt-get update && apt-get install -y bash


# Set the working directory
WORKDIR /app

# Copy the Go server build from the first stage
COPY --from=server-build /app /app

# Copy the built frontend from the second stage
COPY --from=frontend-build /app/frontend/dist /app/frontend/dist

RUN curl -sL https://deb.nodesource.com/setup_16.x | bash - && \
    apt-get install -y nodejs


# Expose ports
EXPOSE 3000
EXPOSE 8080

# Set environment variables if needed
ENV GO_ENV=production

# Command to run the Go server
CMD ["go", "run", "server/server.go"]
