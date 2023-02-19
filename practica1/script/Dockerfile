FROM ubuntu:22.04
WORKDIR /myapp
COPY script.sh /myapp
RUN apt update
RUN apt install curl -y
RUN apt install jq -y
RUN chmod +x script.sh
CMD ["./script.sh"]