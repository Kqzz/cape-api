CGO_ENABLED=0 GOOS=linux go build .
sudo docker build -t kqzz/cape-api .

# I'm going to forget the run command so here it is: sudo docker run -it --rm --name cape-api -p 8080:8080 kqzz/cape-api