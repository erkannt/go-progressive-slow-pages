package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func landingHandler(c echo.Context) error {
	return page(landing()).Render(c.Request().Context(), c.Response().Writer)
}

func baselineHandler(c echo.Context) error {
	results := getResults()
	return page(baseline(results)).Render(c.Request().Context(), c.Response().Writer)
}

func chunkedHandler(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Response().Header().Set("Transfer-Encoding", "chunked")
	flusher, _ := c.Response().Writer.(http.Flusher)

	fmt.Fprintf(c.Response().Writer, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><title>Search | go-progressive-search</title></head><body><header><nav><a href=\"/\">Home</a></nav><h1>Chunked</h1><p>Partial html page is sent before sending slow data chunk by chunk. The final chunk completes the HTML page</p></header><section><ul>")
	flusher.Flush()

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(c.Response().Writer, "<li>Chunk %d</li>\n", i+1)
		flusher.Flush()
	}

	fmt.Fprintln(c.Response().Writer, "</ul><p>Done</p></section></body></html>")
	return nil
}

func chunkedWithTemplHandler(c echo.Context) error {

	results := make(chan string)

	go func() {
		defer close(results)
		for i := range 5 {
			select {
			case <-c.Request().Context().Done():
				return
			case <-time.After(time.Second):
				results <- fmt.Sprintf("Chunk %d", i+1)
			}
		}
	}()

	templ.Handler(page(chunked(results)), templ.WithStreaming()).ServeHTTP(c.Response().Writer, c.Request())
	return nil
}

func slotsHandler(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Response().Header().Set("Transfer-Encoding", "chunked")
	flusher, _ := c.Response().Writer.(http.Flusher)

	fmt.Fprintf(c.Response().Writer, `
	<!doctype html><html lang="en">
	<head><meta charset="UTF-8"><title>Search | go-progressive-search</title></head>
	<body>
	<template shadowrootmode="open">
		<header>
			<nav><a href="/">Home</a></nav>
			<h1>Chunked</h1>
			<p>Full page is sent with template slots. Afterwards slow content is sent to populate the slots.</p>
		</header>
		<section>
			<slot name="1"><div>Loading 1...</div></slot>
			<slot name="2"><div>Loading 2...</div></slot>
			<slot name="3"><div>Loading 3...</div></slot>
			<slot name="4"><div>Loading 4...</div></slot>
			<slot name="5"><div>Loading 5...</div></slot>
		</section>
	</template>
	`)
	flusher.Flush()

	for i := range 5 {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(c.Response().Writer, "<div slot=\"%d\">Chunk %d</div>\n", i+1, i+1)
		flusher.Flush()
	}
	fmt.Fprintf(c.Response().Writer, "</body></html>")

	return nil
}

func getResults() []string {
	results := []string{
		"a result",
		"another result",
		"and yet another result",
	}
	time.Sleep(2 * time.Second)
	return results
}
