package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "errors"
    "io"
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
    username := os.Getenv("DIRECTADMIN_USERNAME")
    password := os.Getenv("DIRECTADMIN_PASSWORD")
    domain := os.Getenv("DIRECTADMIN_DOMAIN")

    subDomain := acmeKey[16:len(acmeKey)-1]
    subDomainEnv := strings.NewReplacer(".", "_", "-", "_").Replace(strings.ToUpper(subDomain))

    if len(os.Getenv("DIRECTADMIN_" + subDomainEnv + "_HOST")) > 0 {
        host = os.Getenv("DIRECTADMIN_" + subDomainEnv + "_HOST")
    }

    if len(os.Getenv("DIRECTADMIN_" + subDomainEnv + "_USERNAME")) > 0 {
        username = os.Getenv("DIRECTADMIN_" + subDomainEnv + "_USERNAME")
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

    if 0 == len(username) {
        log.Fatal(errors.New("No username given"))
        return
    }

    if 0 == len(password) {
        log.Fatal(errors.New("No password given"))
        return
    }

    if 0 == len(domain) {
        log.Fatal(errors.New("No domain given"))
        return
    }

    var postValues map[string]string;

    if "present" == action {
        postValues = map[string]string{
            "domain": domain,
            "type": "TXT",
            "name": acmeKey,
            "value": acmeValue,
            "ttl": "10",
            "affect_pointers": "yes",
            "json": "yes",
            "action": "add",
        }
    }else if "cleanup" == action {
        postValues = map[string]string{
            "domain": domain,
            "json": "yes",
            "action": "select",
            "delete": "yes",
            "affect_pointers": "yes",
            "txtrecs0": "name=" + acmeKey + "&value=" + acmeValue,
        }
    } else {
        log.Fatal(errors.New("No valid action given"))
        return;
    }

    json_data, err := json.Marshal(postValues)

    if err != nil {
        log.Fatal(err)
        return;
    }

    if host[len(host) -1:] != "/" {
        host = host + "/"
    }

    client := &http.Client{}
    req, err := http.NewRequest("POST", host + "CMD_DNS_CONTROL", bytes.NewBuffer(json_data))
    req.SetBasicAuth(username, password)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode != http.StatusOK {
        log.Fatal(errors.New("Error in Request: " + string(body)))
        return;
    }

    fmt.Println(string(body))
}
