CGO_ENABLED=0 GOOS=linux go build .
sudo docker build -t kqzz/cape-api .