module github.com/cdvelop/fileserver

go 1.20

require (
	github.com/cdvelop/filehandler v0.0.34
	github.com/cdvelop/fileinput v0.0.64
	github.com/cdvelop/maps v0.0.8
	github.com/cdvelop/model v0.0.107
	github.com/cdvelop/sqlite v0.0.102
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/testools v0.0.77
	github.com/cdvelop/unixid v0.0.48
	github.com/gabriel-vasile/mimetype v1.4.3
)

require (
	github.com/cdvelop/api v0.0.98 // indirect
	github.com/cdvelop/cutkey v1.0.14 // indirect
	github.com/cdvelop/dbtools v0.0.81 // indirect
	github.com/cdvelop/fetchserver v0.0.26 // indirect
	github.com/cdvelop/input v0.0.81 // indirect
	github.com/cdvelop/logserver v0.0.33 // indirect
	github.com/cdvelop/object v0.0.67 // indirect
	github.com/cdvelop/objectdb v0.0.111 // indirect
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
