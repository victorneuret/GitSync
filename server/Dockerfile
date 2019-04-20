FROM golang:1.12

WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .
COPY config/config-prod.json config/config.json

# Download all the dependencies
RUN go get -d -v .

# Install the package
RUN go build -v .

# Add universe repos
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 3B4FE6ACC0B21F32
RUN echo "deb http://archive.ubuntu.com/ubuntu bionic main universe" >> /etc/apt/sources.list
RUN apt update

# Install blih
RUN apt -y install python3
RUN git clone https://gitlab.com/EpitechContent/dump.git; cp dump/blih.py /usr/bin/blih

# Run the executable
CMD ["./GitSync"]