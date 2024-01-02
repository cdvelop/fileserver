module github.com/cdvelop/fileserver

go 1.20

require (
	github.com/cdvelop/filehandler v0.0.38
	github.com/cdvelop/fileinput v0.0.70
	github.com/cdvelop/maps v0.0.8
	github.com/cdvelop/model v0.0.109
	github.com/cdvelop/sqlite v0.0.104
	github.com/cdvelop/strings v0.0.9
	github.com/cdvelop/testools v0.0.82
	github.com/cdvelop/unixid v0.0.49
	github.com/gabriel-vasile/mimetype v1.4.3
)

require (
	github.com/cdvelop/api v0.0.103 // indirect
	github.com/cdvelop/cutkey v1.0.14 // indirect
	github.com/cdvelop/dbtools v0.0.81 // indirect
	github.com/cdvelop/fetchserver v0.0.26 // indirect
	github.com/cdvelop/input v0.0.83 // indirect
	github.com/cdvelop/logserver v0.0.35 // indirect
	github.com/cdvelop/object v0.0.70 // indirect
	github.com/cdvelop/objectdb v0.0.113 // indirect
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/structs v0.0.1 // indirect
	github.com/cdvelop/timeserver v0.0.32 // indirect
	github.com/cdvelop/timetools v0.0.34 // indirect
	github.com/mattn/go-sqlite3 v1.14.19 // indirect
	golang.org/x/net v0.19.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/fileinput => ../fileinput

replace github.com/cdvelop/testools => ../testools
