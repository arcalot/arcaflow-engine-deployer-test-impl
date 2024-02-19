FROM quay.io/centos/centos:stream8@sha256:7b56a6667ca1e57935a055307bca430e1c3d9d328365240c69e21a225f507a5f

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]