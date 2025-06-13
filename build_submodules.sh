pushd workflow/compress/qpl
mkdir -p build
mkdir -p install
pushd build
CXXFLAGS="-std=gnu++17" cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=../install ..
CXXFLAGS="-std=gnu++17" cmake --build . --target install
popd
popd
