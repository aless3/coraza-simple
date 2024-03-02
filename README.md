# coraza-simple

This is a simlpe container based on the [GOLANG base container](https://hub.docker.com/_/golang/) with an http server that uses Coraza as waf to block incoming connections.
Coraza is set to disruptive mode (`SecRuleEngine On` instead of the default`SecRuleEngine DetectionOnly`)

This aims to be an easy replacement for the Modsecurity container with default values, to use while waiting other Coraza implementation matures (like wasm [[1](https://github.com/corazawaf/coraza-proxy-wasm)/[2](https://github.com/jcchavezs/coraza-http-wasm)] or the [yaegi interpreter](https://github.com/traefik/yaegi) [gets full 3rd party compatibility](https://github.com/traefik/yaegi/issues/1612)) and get ready to be used in other software like Traefik.
This is why this project aims to be a temporary fix and hopefully will be replaced by more advanced software in the next few months/years.

I found this useful to use with the [traefik modsecurity plugin](https://plugins.traefik.io/plugins/628c9eadffc0cd18356a9799/modsecurity-plugin) ([source](https://github.com/acouvreur/traefik-modsecurity-plugin)), replacing the [owasp/modsecurity-crs:apache](https://github.com/acouvreur/traefik-modsecurity-plugin/blob/5c33072a479423a8d623cccd3905db1673208acc/docker-compose.yml#L25) image ([docker hub](https://hub.docker.com/r/owasp/modsecurity-crs/))with this one.

The server is heavily based on the official [http-server example from Coraza](https://github.com/corazawaf/coraza/tree/main/examples/http-server).
