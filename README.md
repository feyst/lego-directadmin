## Test run
```shell
DIRECTADMIN_HOST=https://server.domain.tld:2222 \
    DIRECTADMIN_USERNAME=user-name \
    DIRECTADMIN_PASSWORD='*****************' \
    DIRECTADMIN_DOMAIN=domain.tld \
    go run . \
    present \
    "_acme-challenge.my.example.org." \
    "MsijOYZxqyjGnFGwhjrhfg-Xgbl5r68WPda0J9EgqqI"
```

## Publish
```shell
git tag "v"$(git rev-list --count main)
git push --tags
```

```shell
DIRECTADMIN_HOST=https://vserver442.axc.eu \
    DIRECTADMIN_USERNAME=feyst \
    DIRECTADMIN_PASSWORD='dH92h!GLNp6BPc!9#wr#L#w@u9%grY' \
    go run . \
    present \
    "_acme-challenge.feyst.nl." \
    "Test-6"
```
