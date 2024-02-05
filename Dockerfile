FROM quay.io/centos/centos:stream8@sha256:20ef8d90e1bd590f614dccb6e24d612bd4fc85fbe394f2395f103b6aa7140c4d

COPY io_test_script.bash /

ENTRYPOINT [ "bash", "io_test_script.bash" ]