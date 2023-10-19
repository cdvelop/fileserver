module github.com/cdvelop/fileserver

go 1.20

require github.com/gabriel-vasile/mimetype v1.4.3

require (
	github.com/cdvelop/model v0.0.56
	golang.org/x/net v0.17.0 // indirect
)

replace github.com/cdvelop/model => ../model
