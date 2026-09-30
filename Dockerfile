# Use Go 1.27 and Alpine Linux as the build stage.
FROM golang:1.27-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy Go module files first so Docker can cache dependency downloads.
COPY go.sum go.mod ./

# Download the project's Go dependencies.
RUN go mod download

# Copy all project files into the container's working directory.
COPY . .

# CGO_ENABLED=0 disables CGO and produces a self-contained binary.
# GOOS=linux builds the binary for Linux.
# -o logradar names the compiled executable "logradar".
RUN CGO_ENABLED=0 GOOS=linux go build -o logradar ./cmd/logradar


# Alpine is a lightweight Linux distribution,
# so it keeps the final Docker image small.
FROM alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy only the files required to run LogRadar.
COPY --from=builder /app/logradar .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/logs ./logs

# Expose port 9000 because LogRadar's /metrics endpoint runs on port 9000.
EXPOSE 9000

# Run the compiled LogRadar executable when the container starts.
CMD ["./logradar"]