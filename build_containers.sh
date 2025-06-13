rm -r build
go run wiring/main.go -o build -w docker
cp ./build/.local.env ./build/docker/.env
rm ./build/docker/docker-compose.yml
rm ./build/docker/compress_ctr/Dockerfile
rm ./build/docker/frontend_ctr/Dockerfile
cp ./offload_docker_files/docker-compose.yml ./build/docker/docker-compose.yml
cp ./offload_docker_files/Compress_Dockerfile ./build/docker/compress_ctr/Dockerfile
cp ./offload_docker_files/Frontend_Dockerfile ./build/docker/frontend_ctr/Dockerfile
cp ./workload/xml ./build/wlgen/wlgen_proc/wlgen_proc/xml
