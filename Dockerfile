# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy dependency file first (cache optimization
COPY go.mod ./

# Copy source code
COPY main.go ./
COPY Artprinter/ ./Artprinter/
COPY handlers/ ./handlers/

# Build the binary
RUN go build -o ascii-art-web .

# Stage 2: Production
FROM alpine:latest

# Add labels
LABEL maintainer="Dixon Osure"
LABEL version="1.0"
LABEL description="ASCII Art Web Generator"
LABEL project="ascii-art-web"

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/ascii-art-web .

# Copy static files (templates and banners)
COPY templates/ ./templates/
COPY banners/ ./banners/

# Expose port
EXPOSE 8080

# Create non-root user
RUN adduser -D appuser
USER appuser

# Run the app
CMD ["./ascii-art-web"]