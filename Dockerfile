FROM quay.io/centos/centos:stream8@sha256:ce6ec049788dd34c9fd99cf6c319a1cc69579977b8433d00cc982df6b75841f6

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]