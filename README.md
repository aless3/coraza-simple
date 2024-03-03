# IMPORTANT NOTICE
I still have to set up automatic builds, so the examples pulling from ghcr.io still don't work, if you build the image yourself it works.

# coraza-simple

This is a simlpe container based on the [GOLANG base container](https://hub.docker.com/_/golang/) with an http server that uses Coraza as waf to block incoming connections.
Coraza is set to disruptive mode (`SecRuleEngine On` instead of the default`SecRuleEngine DetectionOnly`)

There is a simple docker-compose.yml file that shows how to use this image in traefik.

# How to add rules
## TLDR - but you should really read the long version
The loaded files in are `/etc/coraza/default/coraza.conf`, `/etc/coraza/coreruleset/crs-setup.conf.example` and `.conf` file in the `/etc/coraza/coreruleset/rules/` folder and `.conf` files in the `/etc/coraza/custom/` folder.

## Long version
First of all you should mount the `/etc/coraza/coreruleset/` folder and use an up to date version of the CRS, the one shipped is grabbed when the container is built.

Then the `/etc/coraza/default/` directory has the file `coraza.conf` that is the [reccomended coraza config](https://raw.githubusercontent.com/corazawaf/coraza/v3/dev/coraza.conf-recommended) - you should update that too, you can mount the `/etc/coraza/default/` for easier access to the file (sometimes docker has problems mounting a file instead of a directory).

You can also mount the `/etc/coraza/custom/` folder and place there the `.conf` file, every `.conf` file there will be read at startup, this may be useful when just syncing git with the upstream CRS to avoid adding files in that repo accidentally.

## Simple stupid way to update CRS and coraza default
Replace path/to/xyz to the respective mounted path.
``` shell

# Update coraza.conf
wget https://raw.githubusercontent.com/corazawaf/coraza/v3/dev/coraza.conf-recommended -O path/to/default/coraza.conf

# also set Coraza with the disruptive mode to block requests instead of just logging them
sed -i'.bak' 's|SecRuleEngine DetectionOnly|SecRuleEngine On|' coraza.conf

# Update CRS
cd path/to/coreruleset
git clone pull

```

# Info about why this project
This aims to be an easy replacement for the Modsecurity container with default values, to use while waiting other Coraza implementation matures (like wasm [[1](https://github.com/corazawaf/coraza-proxy-wasm)/[2](https://github.com/jcchavezs/coraza-http-wasm)] or the [yaegi interpreter](https://github.com/traefik/yaegi) gets [full 3rd party compatibility](https://github.com/traefik/yaegi/issues/1612)) and get ready to be used in other software like Traefik.
This is why this project aims to be a temporary fix and hopefully will be replaced by more advanced software in the next few months/years.

I found this useful to use with the [traefik modsecurity plugin](https://plugins.traefik.io/plugins/628c9eadffc0cd18356a9799/modsecurity-plugin) ([source](https://github.com/acouvreur/traefik-modsecurity-plugin)), replacing the [owasp/modsecurity-crs:apache](https://github.com/acouvreur/traefik-modsecurity-plugin/blob/5c33072a479423a8d623cccd3905db1673208acc/docker-compose.yml#L25) image ([docker hub](https://hub.docker.com/r/owasp/modsecurity-crs/)) with this one.
The docker-compose files refer to this plugin.

The server is heavily based on the official [http-server example from Coraza](https://github.com/corazawaf/coraza/tree/main/examples/http-server).

The directory [`testdata`](https://github.com/aless3/coraza-simple/tree/main/http-server/testdata) is a stripped down version form the [official example](https://github.com/corazawaf/coraza/tree/main/examples/http-server/testdata) (I removed the support to change the response body on the fly).
