FROM golang:1.7
WORKDIR /usr/src/app

# Install dependencies
RUN apt-get update -q -y
RUN apt-get install -q -y --no-install-recommends apt-utils

# Install node
RUN apt-get install -q -y --no-install-recommends npm
RUN npm i -g n
RUN n latest

# Intall JQ
RUN apt-get install -q -y jq

# Import the app
COPY . /usr/src/app
CMD ./e2e.sh

# To build
# # cd MontoCarloSim
# # docker build -t monte .

