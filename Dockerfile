FROM quay.io/centos/centos:stream9

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]