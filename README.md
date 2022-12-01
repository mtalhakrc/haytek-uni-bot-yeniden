# Cetele

haytek sıtacer üni çetele kontrolü yapar.

### sunucuya dosya gönderme

scp -i ~/.ssh/cetele.pem -r /Users/talha/go/src/github.com/haytek-uni-cetele-kontrol/main ubuntu@3.89.107.238:/home/ubuntu

### linux exe çıktısı

GOOS=linux GOARCH=amd64 go build main.go

### system start

sudo systemctl start cetele.service

### set env var to systemd
sudo systemctl set-environment IS_DEVELOPMENT=true  || makine reboot olursa tekrar girmek gerekiyor.

### systom stop

sudo systemctl stop cetele.service

## baglan

ssh -i ~/.ssh/cetele.pem ubuntu@ec2-3-89-107-238.compute-1.amazonaws.com

## postgres

sudo su postgres --- psql

## last logs

systemctl -l status cetele
