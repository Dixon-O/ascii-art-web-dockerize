#!bin/bash
echo "Starting Docker cleanup..."

#Remove stopped containers ...
echo "Removing stopped containers ..."
docker container prune -f


#Remove unused images
echo "Removing unused images..."
docker image prune -f

#Remove dangling images 
DANGLING_IMAGES=$(docker images -f "dangling=true" -q)
if [ -n "$DANGLING_IMAGES" ]; then
    echo "Removing Dangling Images..."
    docker rmi $DANGLING_IMAGES
else
    echo "No Dangling images found."
fi
    echo "Docker cleanup complete"