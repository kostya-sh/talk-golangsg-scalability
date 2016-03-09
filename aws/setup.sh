d=$HOME/src/talk/aws

mkdir -pv $HOME/bin

cat <<EOF > ~/.bash_profile
export GOPATH=$HOME
export GOROOT=$HOME/go
export PATH=$HOME/bin:$HOME/go/bin:$PATH
EOF

. ~/.bash_profile

# Go
if [ ! -d $GOROOT ] ; then
    wget -c https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
    tar xf go1.6.linux-amd64.tar.gz
    rm go1.6.linux-amd64.tar.gz
fi

# postgresql
sudo apt-get -y install postgresql
sudo cp $d/pg_hba.conf /etc/postgresql/9.3/main/
sudo service postgresql restart
sudo -u postgres psql template1 < $d/create.sql

# Talk
go get github.com/lib/pq
go install talk/...
gen-temp > ~/t.csv

# Present
go get golang.org/x/tools/cmd/present
(cd $HOME/src/talk; present -play=false -http ':3999' 1>/dev/null 2>/dev/null) &

