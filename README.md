# Experiments in improving slow page UX without JS

Run `make` and visit [localhost:8080](http://localhost:8080)

Things I've played with:

- [templ](https://templ.guide/)
- [declaritive shadow DOM](https://web.dev/articles/declarative-shadow-dom)
- [streaming HTML](https://dev.to/tigt/the-weirdly-obscure-art-of-streamed-html-4gc2)
- [:has-slotted](https://developer.mozilla.org/en-US/docs/Web/CSS/:has-slotted)

For an example of this in the Node rather than Go ecosystem see this [lamplightdev blog post](https://lamplightdev.com/blog/2024/01/10/streaming-html-out-of-order-without-javascript/).

## Notes

- Not sure how to ensure good a11y when replacing progress bars during loading to communicate progress. I could not get [orca](https://orca.gnome.org/) to read out the page before it had completed loading.
- I was dubious about the code generation aspect of `templ` but once I had set up the tooling ([watching with air](.air.toml), [templ-vscode](https://marketplace.visualstudio.com/items/?itemName=a-h.templ) and [templ go to def](https://marketplace.visualstudio.com/items/?itemName=lsl.vscode-templ-go-to-definition)) it is a good experience.
