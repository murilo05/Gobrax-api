FROM golang:latest AS builder

WORKDIR /app

# Copy Source code 
COPY . .
# Download dependencies
RUN go mod download

# Build the Go binary with necessary compiler flags for optimization
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gobrax-api /app/cmd/main.go

# Use a minimal alpine image for the runtime stage
FROM alpine:latest

WORKDIR /root/

# Remove source code - not necessary anymore
RUN rm -rf /app

# Copy the binary from the builder stage
COPY --from=builder /app/gobrax-api .
COPY .env .

# Defina as variáveis de ambiente necessárias
ENV HOST=localhost
ENV PORT=8080
ENV DBHOST=127.0.0.1
ENV DBPORT=5432
ENV DBUSER=postgres
ENV DBPASSWORD=postgres
ENV DBNAME=postgres

# Expose the application on a specific port
EXPOSE 8080

# # Command to run the application
CMD ["./gobrax-api"]