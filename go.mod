module github.com/cdvelop/fileserver

go 1.20

require (
	github.com/cdvelop/api v0.0.34
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/fileinput v0.0.1
	github.com/cdvelop/gotools v0.0.45
	github.com/cdvelop/sqlite v0.0.68
	github.com/cdvelop/timeserver v0.0.6
	github.com/cdvelop/unixid v0.0.7
	github.com/gabriel-vasile/mimetype v1.4.3
)

require (
	github.com/cdvelop/dbtools v0.0.47 // indirect
	github.com/cdvelop/input v0.0.40 // indirect
	github.com/cdvelop/object v0.0.15 // indirect
	github.com/cdvelop/objectdb v0.0.71 // indirect
	github.com/cdvelop/output v0.0.6 // indirect
	github.com/cdvelop/strings v0.0.2 // indirect
	github.com/cdvelop/timetools v0.0.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	golang.org/x/text v0.13.0 // indirect
)

require (
	github.com/cdvelop/model v0.0.57
	github.com/cdvelop/testools v0.0.23
	golang.org/x/net v0.17.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/api => ../api

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/fileinput => ../fileinput

replace github.com/cdvelop/objectdb => ../objectdb

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/input => ../input
