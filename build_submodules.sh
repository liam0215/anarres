pushd deps/qpl
mkdir -p build
mkdir -p install
pushd build
cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=../install ..
cmake --build . --target install
popd
popd
