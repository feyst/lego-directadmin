package main
/*
DIRECTADMIN_HOST=https://vserver442.axc.eu
DIRECTADMIN_USER=feyst
DIRECTADMIN_PASSWORD='*****************'
DIRECTADMIN_DOMAIN=feyst.nl


*/
import (
//     "bytes"
    "encoding/json"
    "fmt"
    "log"
//     "net/http"
    "os"
    "strings"
    "errors"
)

func main() {

    if len(os.Args) < 4 {
        log.Fatal(errors.New("Too few arguments"))
        return
    }

    action := os.Args[1]
    acmeKey := os.Args[2]
    acmeValue := os.Args[3]

    host := os.Getenv("DIRECTADMIN_HOST")
    user := os.Getenv("DIRECTADMIN_USER")
    password := os.Getenv("DIRECTADMIN_PASSWORD")
    domain := os.Getenv("DIRECTADMIN_DOMAIN")

    subDomain := acmeKey[16:len(acmeKey)-1]
    subDomainEnv := strings.NewReplacer(".", "_", "-", "_").Replace(strings.ToUpper(subDomain))

    if len(os.Getenv("DIRECTADMIN_" + subDomainEnv + "_HOST")) > 0 {
        host = os.Getenv("DIRECTADMIN_" + subDomainEnv + "_HOST")
    }

    if len(os.Getenv("DIRECTADMIN_" + subDomainEnv + "_USER")) > 0 {
        user = os.Getenv("DIRECTADMIN_" + subDomainEnv + "_USER")
    }

    if len(os.Getenv("DIRECTADMIN_" + subDomainEnv + "_PASSWORD")) > 0 {
        password = os.Getenv("DIRECTADMIN_" + subDomainEnv + "_PASSWORD")
    }

    if len(os.Getenv("DIRECTADMIN_" + subDomainEnv + "_DOMAIN")) > 0 {
        domain = os.Getenv("DIRECTADMIN_" + subDomainEnv + "_DOMAIN")
    }

    acmeKeySplit := strings.Split(subDomain, ".")

    if len(domain) == 0 {
        domain = strings.Join(acmeKeySplit[len(acmeKeySplit)-2:], ".")
    }

    if 0 == len(host) {
        log.Fatal(errors.New("No host given"))
        return
    }
    values := map[string]string{
        "domain": domain,
        "type": "TXT",
        "name": acmeKey,
        "value": acmeValue,
        "ttl": "10",
        "affect_pointers": "yes",
        "json": "yes",
        "action": "add",
    }

    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(json_data))
    fmt.Println(os.Getenv("DIRECTADMIN_HOST"))
    fmt.Println(subDomain)
    fmt.Println(host)
    fmt.Println(user)
    fmt.Println(password)
    fmt.Println(subDomainEnv)
    fmt.Println(action)

//     resp, err := http.Post("https://vserver442.axc.eu/CMD_DNS_CONTROL", "application/json",
//         bytes.NewBuffer(json_data))
//
//     if err != nil {
//         log.Fatal(err)
//     }
//
//
//
//     fmt.Println("Hello, World!")
}

