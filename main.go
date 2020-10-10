package main

import (
    "errors"
    "flag"
    "fmt"
    "github.com/nerney/dappy"
    "github.com/sensu/uchiwa/uchiwa"
    "github.com/sensu/uchiwa/uchiwa/audit"
    "github.com/sensu/uchiwa/uchiwa/authentication"
    "github.com/sensu/uchiwa/uchiwa/authorization"
    "github.com/sensu/uchiwa/uchiwa/config"
    "github.com/sensu/uchiwa/uchiwa/filters"
    "os"
)

func CheckUser(username, password string) (*authentication.User, error)  {
    ldap_password  := os.Getenv("LDAP_BIND_PASSWORD")
    ldap_bind_user := os.Getenv("LDAP_BIND_USER")

    ldap_host      := os.Getenv("LDAP_HOST")
    ldap_base_dn   := os.Getenv("LDAP_BASE_DN")
    ldap_filter    := os.Getenv("LDAP_FILTER")

    var client dappy.Client

    client, err := dappy.New(dappy.Config{
        BaseDN: ldap_base_dn,
        Filter: ldap_filter,
        ROUser: dappy.User{
            Name: ldap_bind_user,
            Pass: ldap_password,
        },
        Host:   ldap_host,
    })

    if err != nil {
        return &authentication.User{}, errors.New("Cannot connect to ldap")
    }

    connect := client.Auth(username, password)
    if connect != nil {
        return &authentication.User{}, fmt.Errorf("invalid user '%s' or invalid password", username)
    }

    return &authentication.User{Username:username}, nil
}

func main() {
    configFile := flag.String("c", "/etc/sensu/config.json", "Full or relative path to the configuration file")
    configDir  := flag.String("d", "/etc/sensu/", "Full or relative path to the configuration directory, or comma delimited directories")
    publicPath := flag.String("p", "/opt/uchiwa/public", "Full or relative path to the public directory")
    flag.Parse()

    config := config.Load(*configFile, *configDir)

    u := uchiwa.Init(config)

    auth := authentication.New(config.Uchiwa.Auth)

    auth.Advanced(CheckUser, string("ldap"))

    // Audit
    audit.Log = audit.LogMock

    // Authorization
    uchiwa.Authorization = &authorization.Uchiwa{}

    // Filters
    uchiwa.Filters = &filters.Uchiwa{}

    u.WebServer(publicPath, auth)
}

