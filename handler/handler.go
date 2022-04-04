package handler

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"net/url"
	"time"
	"urlShorter/storage"
)


func New(schema string, host string, storage storage.Service) *router.Router {
	router := router.New()

	h := handler{schema, host, storage}
	router.POST("/api/v1/urls", responseHandler(h.encode))
	router.GET("/{shortLink}", h.redirect)
	router.GET("/{shortLink}/info", responseHandler(h.decode))
	return router
}


type response struct {
	Success bool `json:"success"`
	Data interface{} `json:"shortUrl"`
	Id string `json:"id"`
}

type handler struct {
	schema string
	host string
	storage storage.Service
}

func responseHandler(h func(ctx *fasthttp.RequestCtx) (string, interface{}, int, error)) fasthttp.RequestHandler {
	// fmt.Printf("in responseHandler\n")
	return func(ctx *fasthttp.RequestCtx) {
		c, data, status, err := h(ctx)
		if err != nil {
			data = err.Error()
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetStatusCode(status)
		err = json.NewEncoder(ctx.Response.BodyWriter()).Encode(response{Data: data, Success: err == nil, Id: c})
		if err != nil {
			log.Printf("could not enocde response to output: %v", err)
		}
	}
}

func (h handler) encode(ctx *fasthttp.RequestCtx) (string, interface{}, int, error) {
	var input struct {
		URL string `json:"url"`
		Expires string `json:"expires"`
	}

	if err := json.Unmarshal(ctx.PostBody(), &input); err != nil {
		return "", nil, http.StatusBadRequest, fmt.Errorf("Unable to decode JSON request body: %v", err)
	}

	uri, err := url.ParseRequestURI(input.URL)
	// fmt.Printf("uri in handler-encode: %s\n", uri)

	if err != nil {
		return "", nil, http.StatusBadRequest, fmt.Errorf("Invalid url")
	}

	layoutISO := "2006-01-02 15:04:05"
	expires, err := time.Parse(layoutISO, input.Expires)
	if err != nil {
		return "", nil, http.StatusBadRequest, fmt.Errorf("Invalid expiration data")
	}

	c, err := h.storage.Save(uri.String(), expires)
	// fmt.Printf("c: %s\n", c)
	if err != nil {
		return "", nil, http.StatusInternalServerError, fmt.Errorf("Could not store in database: %v", err)
	}

	u := url.URL{
		Scheme: h.schema,
		Host: h.host,
		Path: c}

	// fmt.Printf("Generated link: %v \n", u.String())
	// fmt.Printf("u.String(): ", u.String())
	return c, u.String(), http.StatusCreated, nil
}

func (h handler) decode(ctx *fasthttp.RequestCtx) (string, interface{}, int, error) {
	code := ctx.UserValue("shortLink").(string)

	model, err := h.storage.LoadInfo(code)
	if err != nil {
		return "", nil, http.StatusNotFound, fmt.Errorf("URL not found")
	}

	return "", model, http.StatusOK, nil
}

func (h handler) redirect(ctx *fasthttp.RequestCtx) {
	// fmt.Printf("In redirect()\n")
	code := ctx.UserValue("shortLink").(string)

	uri, err := h.storage.Load(code)
	// fmt.Printf("uri: %s\n", uri)
	if err != nil {
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.Response.SetStatusCode(http.StatusNotFound)
		return
	}

	ctx.Redirect(uri, http.StatusMovedPermanently)
}
