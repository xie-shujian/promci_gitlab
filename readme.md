# promci with gitea
use go-git, gitea

check the .git folder, if not find, use git clone, or use git pull

not use git cmd from shell

# flow
```mermaid
flowchart LR
    A[gitea] --> B[promci] --> C[linux server]

```

# config
## gitea access token
read_api, read_repository, read_registry
## repository
* test46
## directory
linux server repository directory

# build
```
cd C:/workplace1/promci
set GOOS=linux
set GOARCH=amd64
go build promci
```

# config file
/etc/promci/promci.yml
# log file
/var/log/promci.log
# run
192.168.1.101
```
/app/promci --config.file=c:/test/promci.yml &
```

# gitea config
## webhook
http://192.168.1.101:8866/promci

HTTP Method: POST

HTTP Headers: Authrization=access token

# test
update pc folder
```
cd C:/git/test46
git add .
git commit -m "kafka"
git push http://gitlab.test.com/monitor/test46.git
```
linux server result
```
tree /etc/test46/
```