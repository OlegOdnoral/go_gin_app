cd ../..

@echo build docker image...

docker build -t go_gin_appt .

docker system prune

pause