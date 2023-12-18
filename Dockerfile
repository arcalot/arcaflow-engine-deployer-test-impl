FROM quay.io/centos/centos:stream8@sha256:34aaf8788a2467f602c5772884448236bb41dfe1691a78dee33053bb24474395

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]