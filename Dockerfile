FROM quay.io/centos/centos:stream8@sha256:02721b35a33bff24a7bbbd17e366e61b484827f1278cb3046113d1eaa9214f2a

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]