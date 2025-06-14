#!/bin/bash
set -e

SIZE=${1:-64}  # Default to 64 if not provided
TPUT=${2:-4000}  # Default to 4000 if not provided

. build/.local.env
cd build/cmplx_wlgen/cmplx_wlgen_proc/cmplx_wlgen_proc/
# ./wlgen_proc --frontend.grpc.dial_addr=${FRONTEND_GRPC_DIAL_ADDR} --zipkin.dial_addr=${ZIPKIN_DIAL_ADDR} --sizeKb=64 --numWorkers=10
# ./wlgen_proc --frontend.grpc.dial_addr=${FRONTEND_GRPC_DIAL_ADDR} --sizeKb=64 --numWorkers=10 --duration="15s"
./cmplx_wlgen_proc --frontend.grpc.dial_addr=${FRONTEND_GRPC_DIAL_ADDR} -size=${SIZE} -dur="30s" -tput=${TPUT}
