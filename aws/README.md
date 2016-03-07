### Initial Setup

ssh -i dev.pem ubuntu@ip

sudo apt-get update
sudo apt-get -y install git
mkdir src
git clone https://github.com/kostya-sh/talk-golangsg-scalability.git src/talk
. ./src/talk/aws/setup.sh

### PostgreSQL

$ sudo -u postgres psql

```
\connect d
select count(*) from temperature;
```

$ echo "select count(*) from temperature" | sudo -u postgres psql d
