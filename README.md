## Uchiwa Dashboard ##

This project is a module for [Sensu Uchiwa](https://github.com/sensu/uchiwa)
with support of LDAP authentication.

### Docker ###

* How to build a docker image 

```bash
docker build -t uchiwa-ldap:latest .
```

* Environment variables for Docker and LDAP

This image is required LDAP credentials as ENV variables and uchiwa config.json
with a minimal config options.

```bash
docker run --rm -p 8080:8080        \ 
   -e LDAP_BIND_PASSWORD=<password> \
   -e LDAP_BIND_USER=<bind_user>    \
   -e LDAP_FILTER=<ldap_filter>     \
   -e LDAP_BASE_DN=<ldap_base_dn>   \
   -e LDAP_HOST=<ldap_server:389>   \
   -v $PWD:/etc/sensu/              \
   dockerhub.com/xbulat/uchiwa-ldap
```
