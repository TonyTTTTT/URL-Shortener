## 2022/4/4
- Done the two requirements
- program flow for initial
	1. main()
	1. configuration.go for config(server config, redis config, etc.)
	1. handler.New() in main.go
- program flow for create shortURL
	1. router.POST in handler.go
	1. encode() in handler.go
	1. Save() in redis.go
		- interact to redis via redigo
	1. Encode() in base62.go
	1. responseHandler() in handler.go
	1. response to client
- program flow for redirect
	1. router.GET in handler.go
	1. redirect() in handler.go
	1. responseHandler() in handler.go
	1. response to client
	


## 2022/4/2
- initial from the [template](https://intersog.com/blog/how-to-write-a-custom-url-shortener-using-golang-and-redis/)
