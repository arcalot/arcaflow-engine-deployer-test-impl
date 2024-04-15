FROM quay.io/centos/centos:stream8@sha256:cb10b58aa113732193dddbbe4e91d80ca97f8ec742ef66cf4ff8340f5b90faab

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]