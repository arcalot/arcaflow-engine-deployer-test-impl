FROM quay.io/centos/centos:stream8@sha256:1338a5c9e232a419978dbb2f6f42a71553e0ed53f3332b618f89a19143070d86

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]