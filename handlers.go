package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

func landingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page(landing()).Render(r.Context(), w)
}

func baselineHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results := getResults()
	page(baseline(results)).Render(r.Context(), w)
}

func chunkedHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Transfer-Encoding", "chunked")
	flusher, _ := w.(http.Flusher)

	fmt.Fprintf(w, `
	<!doctype html><html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Progressive Slow Pages</title>
	</head>
	<body>
	<header>
		<nav><a href="/">Home</a></nav>
		<h1>Chunked</h1>
		<p>Partial html page is sent before sending slow data chunk by chunk. The final chunk completes the HTML page</p>
	</header>
	<section>
		<ul>
	`)
	flusher.Flush()

	for i := range 5 {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(w, "<li>Chunk %d</li>\n", i+1)
		flusher.Flush()
	}

	fmt.Fprintln(w, "</ul><p>Done</p></section></body></html>")
}

func chunkedWithTemplHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	results := make(chan string)

	go func() {
		defer close(results)
		for i := range 5 {
			select {
			case <-r.Context().Done():
				return
			case <-time.After(time.Second):
				results <- fmt.Sprintf("Chunk %d", i+1)
			}
		}
	}()

	templ.Handler(page(chunked(results)), templ.WithStreaming()).ServeHTTP(w, r)
}

func slotsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Transfer-Encoding", "chunked")
	flusher, _ := w.(http.Flusher)

	fmt.Fprintf(w, `
	<!doctype html><html lang="en">
	<head><meta charset="UTF-8"><title>Progressive Slow Pages</title></head>
	<body>
	<template shadowrootmode="open">
		<header>
			<nav><a href="/">Home</a></nav>
			<h1>Slots</h1>
			<p>Full page is sent with template slots. Afterwards slow content is sent to populate the slots.</p>
		</header>
		<section>
			<ul>
			<slot name="1"><li>Loading 1...</li></slot>
			<slot name="2"><li>Loading 2...</li></slot>
			<slot name="3"><li>Loading 3...</li></slot>
			<slot name="4"><li>Loading 4...</li></slot>
			<slot name="5"><li>Loading 5...</li></slot>
			</ul>
		</section>
	</template>
	`)
	flusher.Flush()

	for i := range 5 {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(w, "<li slot=\"%d\">Chunk %d</li>\n", i+1, i+1)
		flusher.Flush()
	}
	fmt.Fprintf(w, "</body></html>")
}

type Chunk struct {
	name    string
	content string
}

func slotsWithTemplHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	results := make(chan Chunk)

	resultCount := 5
	go func() {
		defer close(results)
		for i := range resultCount {
			select {
			case <-r.Context().Done():
				return
			case <-time.After(time.Second):
				results <- Chunk{name: strconv.Itoa(i + 1), content: fmt.Sprintf("Chunk %d", i+1)}
			}
		}
	}()

	templ.Handler(page(slots(resultCount, results)), templ.WithStreaming()).ServeHTTP(w, r)
}

func progressBarHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// resultsAccumulator := []string{}
	resultCount := 5

	percentDone := make(chan int)
	results := make(chan string, resultCount)

	go func() {
		defer close(percentDone)
		defer close(results)
		for i := range resultCount {
			select {
			case <-r.Context().Done():
				return
			case <-time.After(time.Millisecond * 500):
				percentDone <- int(math.Round(float64(i+1) / float64(resultCount) * 100))
				results <- fmt.Sprintf("Result %d", i+1)
			}
		}
	}()

	templ.Handler(page(progressBar(percentDone, results)), templ.WithStreaming()).ServeHTTP(w, r)
}

func getResults() []string {
	results := []string{
		"Chunk 1",
		"Chunk 2",
		"Chunk 3",
		"Chunk 4",
		"Chunk 5",
	}
	time.Sleep(2 * time.Second)
	return results
}
