### Initial Setup

```
ssh -i dev.pem ubuntu@ip
```

```
sudo apt-get update
sudo apt-get -y install git
mkdir src
git clone https://github.com/kostya-sh/talk-golangsg-scalability.git src/talk
. ./src/talk/aws/setup.sh
```

### Copy files

```
scp -r -i aws/dev.pem . ubuntu@ip:src/talk
go install talk/...
```

### Slides

http://ip:3999/scalability.slide#1

http://go-talks.appspot.com/github.com/kostya-sh/talk-golangsg-scalability/scalability.slide#1

### PostgreSQL

```
$ sudo -u postgres psql
```

```
\connect d
select count(*) from temperature;
```

```
$ echo "select count(*) from temperature" | sudo -u postgres psql d
```
