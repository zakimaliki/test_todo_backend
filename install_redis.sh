#!/bin/bash

# Pull Redis image
docker pull redis:latest

# Run Redis container
docker run --name my-redis -p 6379:6379 -d redis:latest

echo "Redis container created and running on localhost:6379"
echo "To use Redis CLI: docker exec -it my-redis redis-cli"