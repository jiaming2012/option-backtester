# Use an official Ubuntu base image
FROM ubuntu:latest

# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive

# Install OpenJDK-17
RUN apt-get update && \
    apt-get install -y openjdk-17-jdk openjdk-17-jre && \
    mkdir /etc/thetadata

# Download the latest version of ThetaTerminal
RUN apt-get install -y wget && \
    wget https://download-latest.thetadata.us -O /etc/thetadata/ThetaTerminal.jar

# Set the working directory
WORKDIR /etc/thetadata

# ENV USERNAME=$THETA_DATA_PASSWORD
# ENV THETA_DATA_PASSWORD=defaultPassword

# Run your Java application
# Use environment variables for username and password
CMD java -jar ThetaTerminal.jar $THETA_DATA_USERNAME $THETA_DATA_PASSWORD