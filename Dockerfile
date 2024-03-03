FROM golang:latest

#
# First we gather the configuration files with the rules and the last CRS available
#

WORKDIR /etc/coraza/default

# Default Coraza config
RUN wget https://raw.githubusercontent.com/corazawaf/coraza/v3/dev/coraza.conf-recommended -O coraza.conf

# The default Coraza does not disrupt connections but only logs them, we change that
RUN sed -i'.bak' 's|SecRuleEngine DetectionOnly|SecRuleEngine On|' coraza.conf

WORKDIR /etc/coraza

# Grab che last CRS from github
RUN git clone https://github.com/coreruleset/coreruleset

#
# Copy the simple http server that will accept and check the requests and build it
#

COPY http-server/* .
RUN go build .

#
# When an OS variable is in use, the http server will use that as the port
#

ENV ENV_PORT=80
expose $ENV_PORT

#
# The command specifies the port, there is an hard-coded default port also in the GO files with the value of 80 just in case
#

CMD ["PORT=$ENV_PORT", "/etc/coraza/http-server"]
