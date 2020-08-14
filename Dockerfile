FROM ubuntu:18.04

# number of concurrent threads during build
# usage: docker build --build-arg PARALLELISM=8 -t name/name .
ARG PARALLELISM=1

RUN apt-get update && \
    apt-get -y --no-install-recommends install apt-utils software-properties-common wget gpg-agent; \
    apt-get -y clean && \
    rm -rf /var/lib/apt/lists/*

# add repos
RUN set -e; \
    add-apt-repository -y ppa:ubuntu-toolchain-r/test; \
    wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE=1 apt-key add -; \
    echo 'deb http://apt.llvm.org/bionic/ llvm-toolchain-bionic-7 main' >> /etc/apt/sources.list; \
    echo 'deb http://apt.llvm.org/bionic/ llvm-toolchain-bionic-9 main' >> /etc/apt/sources.list; \
    apt-get update

RUN set -e; \
    apt-get -y --no-install-recommends install libtool \
    # compilers (gcc-7, gcc-9)
    build-essential g++-9 ninja-build \
    # CI dependencies
    git ssh tar gzip ca-certificates gnupg \
    # Python3
    python3-pip python3-setuptools\
    # other
    curl file gdb gdbserver ccache python3.6-dev openssl nodejs npm\
    gcovr cppcheck doxygen rsync graphviz graphviz-dev unzip vim zip pkg-config; \
    apt-get -y clean

# compiler clang-7, clang-9 and libc++ only on x86_64, for debug purpose
#RUN set -e; \
#    if [ `uname -m` = "x86_64" ]; then \
#      apt-get -y --no-install-recommends install \
#        clang-7 lldb-7 lld-7 libc++-7-dev libc++abi-7-dev clang-format-7 \
#        clang-9; \
#      apt-get -y clean; \
#    fi
# golang stuff
RUN curl https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz | tar -C /opt -xz
ENV GOPATH=/opt/gopath
RUN mkdir ${GOPATH}
ENV PATH=${PATH}:/opt/go/bin:${GOPATH}/bin
RUN go get github.com/golang/protobuf/protoc-gen-go


# install cmake 3.14.0
RUN set -e; \
    curl -L -o /tmp/cmake.sh https://github.com/Kitware/CMake/releases/download/v3.14.0/cmake-3.14.0-Linux-x86_64.sh; \
    sh /tmp/cmake.sh --prefix=/usr/local --skip-license; \
    rm /tmp/cmake.sh
ADD . /opt/app

CMD ["/bin/bash"]