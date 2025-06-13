. build/.local.env
cd build/wlgen/wlgen_proc/wlgen_proc/
# ./wlgen_proc --frontend.grpc.dial_addr=${FRONTEND_GRPC_DIAL_ADDR} --zipkin.dial_addr=${ZIPKIN_DIAL_ADDR} --sizeKb=64 --numWorkers=10
./wlgen_proc --frontend.grpc.dial_addr=${FRONTEND_GRPC_DIAL_ADDR} --sizeKb=64 --numWorkers=10
