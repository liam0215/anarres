rm -r build
go run wiring/main.go -o build -w docker_monolith
cp ./build/.local.env ./build/docker/.env
rm ./build/docker/docker-compose.yml
rm ./build/docker/monolith_ctr/Dockerfile
cp ./offload_docker_files/monolith-docker-compose.yml ./build/docker/docker-compose.yml
cp ./offload_docker_files/Monolith_Dockerfile ./build/docker/monolith_ctr/Dockerfile
cp ./workload/xml ./build/cmplx_wlgen/cmplx_wlgen_proc/cmplx_wlgen_proc/xml
