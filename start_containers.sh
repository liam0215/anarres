cd build/docker
docker compose rm -f
docker rmi -f $(docker images -q)
docker compose up
