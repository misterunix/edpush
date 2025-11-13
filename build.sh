rm bin/edpush-linux-amd64
GOOS=linux GOARCH=amd64 go build -o bin/edpush-linux-amd64 -ldflags '-s -w'

scp bin/edpush-linux-amd64 bjones@10.0.0.205:/home/bjones/easy-diffusion/

