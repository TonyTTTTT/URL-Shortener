# URL Shortener
Honestly saying, I have no experience using Golang or Node.js before, so I [reference](https://intersog.com/blog/how-to-write-a-custom-url-shortener-using-golang-and-redis) a tutorial online to complete this project. 
## To run this project
- First install redis and run the redis server refer to [this](https://redis.io/docs/getting-started/)
```shell
$ cd ROOT/OF/THIS/PROJECT
$ go run main.go
```
## testing the fuctionality
- create short url
	- Notice that \" is for windonws OS to escape the special character, if using Linux, just replace \" as " and the outer " as '
```shell
$ curl -L -X POST localhost:8080/api/v1/urls -H "Content-Type: application/json" -d "{\"url\": \"https://github.com/TonyTTTTT/URL-Shortener\", \"expires\": \"2022-10-04 17:18:00\"}"
```
it will response a json like
```json
{
	"success": true,
	"shortUrl": "http://localhost:8080/OTv0FdGU8Ng",
	"id": "OTv0FdGU8Ng"
}
```
- redirect
```shell
$ curl -L -X GET localhost:8080/OTv0FdGU8Ng
```
then it should sucess redirect to https://github.com/TonyTTTTT/URL-Shortener
## Projcet structure
- base62 is the module for encode the integer into string by base62 encode
- config is the module for loading and applying the configuration.json file
- handler is the module for handle the web request
- storage/redis is the module for interact with redis DB
- log is the dir that I record my progress and findings
- main.go is the entry point of this program
## Package Use
- fasthttp/router for handling the web request
- valyala/fasthttp for create the response object
- redigo for interacting with redsi DB
