FROM quay.io/centos/centos:stream8@sha256:86ba0bf60249e6daee350459e014213742c88341e6e7284695dcf1dfe2c58873

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]