#!/bin/bash

# Script to start Nakama server and test the matchmaking setup

echo "==================================="
echo "Nakama Research - Setup Script"
echo "==================================="
echo ""

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Error: Docker is not installed. Please install Docker first."
    echo "Visit: https://docs.docker.com/get-docker/"
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "Error: Docker Compose is not installed. Please install Docker Compose first."
    echo "Visit: https://docs.docker.com/compose/install/"
    exit 1
fi

echo "✓ Docker and Docker Compose are installed"
echo ""

# Start the services
echo "Starting Nakama and CockroachDB..."
docker-compose up -d

echo ""
echo "Waiting for services to be ready..."
sleep 10

# Check service status
echo ""
echo "Service Status:"
docker-compose ps

echo ""
echo "==================================="
echo "Setup Complete!"
echo "==================================="
echo ""
echo "Nakama Console: http://localhost:7351"
echo "  Username: admin"
echo "  Password: password"
echo ""
echo "Client API: http://localhost:7349"
echo "gRPC API: localhost:7350"
echo ""
echo "To view logs:"
echo "  docker-compose logs -f nakama"
echo ""
echo "To stop services:"
echo "  docker-compose down"
echo ""
echo "The matchmaking function is ready and will match players based on their level!"
echo "Check README.md for usage instructions."
