FROM centos:7
WORKDIR /
COPY ./epel-release-latest-7.noarch.rpm .
RUN  yum install -y epel-release-latest-7.noarch.rpm && yum install -y telnet bind-utils nslookup iperf iproute net-tools
ADD http .
EXPOSE 8000
CMD ["./http"]
