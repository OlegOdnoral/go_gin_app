cd ../..

@echo build docker image...

docker build -t gin_test .

docker system prune

pause