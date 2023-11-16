##### For running web ui 
##### Pull the latest image
`docker pull hibiken/asynqmon`

##### Run Asynqmon
docker run --rm \
    --name asynqmon \
    -p 8080:8080 \
    hibiken/asynqmon

#####
`go get -u github.com/hibiken/asynq/tools/asynq`
`asynq stats`    

##### running redis
docker run --name redisps  -d -p 6379:6379 redis