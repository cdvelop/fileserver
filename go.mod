module github.com/cdvelop/fileserver

go 1.20

require (
	github.com/cdvelop/gotools v0.0.57
	github.com/cdvelop/sqlite v0.0.77
	github.com/cdvelop/strings v0.0.7
	github.com/cdvelop/testools v0.0.38
	github.com/cdvelop/unixid v0.0.18
	github.com/gabriel-vasile/mimetype v1.4.3
)

require (
	github.com/cdvelop/api v0.0.55 // indirect
	github.com/cdvelop/cutkey v0.6.0 // indirect
	github.com/cdvelop/dbtools v0.0.59 // indirect
	github.com/cdvelop/input v0.0.52 // indirect
	github.com/cdvelop/object v0.0.27 // indirect
	github.com/cdvelop/objectdb v0.0.82 // indirect
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/timeserver v0.0.17 // indirect
	github.com/cdvelop/timetools v0.0.18 // indirect
	github.com/mattn/go-sqlite3 v1.14.18 // indirect
	golang.org/x/text v0.14.0 // indirect
)

require (
	github.com/cdvelop/fetchserver v0.0.2 // indirect
	github.com/cdvelop/filehandler v0.0.5
	github.com/cdvelop/fileinput v0.0.26
	github.com/cdvelop/logserver v0.0.3 // indirect
	github.com/cdvelop/maps v0.0.2
	github.com/cdvelop/model v0.0.68
	golang.org/x/net v0.18.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/logserver => ../logserver

replace github.com/cdvelop/sqlite => ../sqlite

replace github.com/cdvelop/fetchserver => ../fetchserver

replace github.com/cdvelop/maps => ../maps

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/testools => ../testools

replace github.com/cdvelop/api => ../api

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/fileinput => ../fileinput

replace github.com/cdvelop/objectdb => ../objectdb

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/input => ../input
