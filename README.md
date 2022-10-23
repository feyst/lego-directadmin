## Download
```shell
wget -c https://github.com/feyst/lego-directadmin/releases/latest/download/lego-directadmin-$(dpkg --print-architecture).tar.gz -O - | tar -xz
chmod +x lego-directadmin
```
## Test run
```shell
DIRECTADMIN_HOST=https://server.domain.tld:2222 \
    DIRECTADMIN_USERNAME=user-name \
    DIRECTADMIN_PASSWORD='*****************' \
    DIRECTADMIN_DOMAIN=domain.tld \
    go run . \
    present \
    "--" \
    "_acme-challenge.my.example.org." \
    "MsijOYZxqyjGnFGwhjrhfg-Xgbl5r68WPda0J9EgqqI"
```

## Publish
```shell
git tag "v"$(git rev-list --count main)
git push --tags
```
