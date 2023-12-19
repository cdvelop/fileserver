module github.com/cdvelop/fileserver

go 1.20

require (
	github.com/cdvelop/filehandler v0.0.32
	github.com/cdvelop/fileinput v0.0.58
	github.com/cdvelop/maps v0.0.8
	github.com/cdvelop/model v0.0.105
	github.com/cdvelop/sqlite v0.0.100
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/testools v0.0.75
	github.com/cdvelop/unixid v0.0.46
	github.com/gabriel-vasile/mimetype v1.4.3
)

require (
	github.com/cdvelop/api v0.0.95 // indirect
	github.com/cdvelop/cutkey v1.0.12 // indirect
	github.com/cdvelop/dbtools v0.0.79 // indirect
	github.com/cdvelop/fetchserver v0.0.24 // indirect
	github.com/cdvelop/input v0.0.77 // indirect
	github.com/cdvelop/logserver v0.0.31 // indirect
	github.com/cdvelop/object v0.0.65 // indirect
	github.com/cdvelop/objectdb v0.0.109 // indirect
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/structs v0.0.1 // indirect
	github.com/cdvelop/timeserver v0.0.31 // indirect
	github.com/cdvelop/timetools v0.0.32 // indirect
	github.com/mattn/go-sqlite3 v1.14.19 // indirect
	golang.org/x/net v0.19.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/fileinput => ../fileinput

replace github.com/cdvelop/testools => ../testools
