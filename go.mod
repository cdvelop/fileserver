module github.com/cdvelop/fileserver

go 1.20

require (
	github.com/cdvelop/sqlite v0.0.79
	github.com/cdvelop/strings v0.0.7
	github.com/cdvelop/testools v0.0.41
	github.com/cdvelop/unixid v0.0.21
	github.com/gabriel-vasile/mimetype v1.4.3
)

require (
	github.com/cdvelop/api v0.0.61 // indirect
	github.com/cdvelop/cutkey v0.6.0 // indirect
	github.com/cdvelop/dbtools v0.0.62 // indirect
	github.com/cdvelop/input v0.0.55 // indirect
	github.com/cdvelop/object v0.0.35 // indirect
	github.com/cdvelop/objectdb v0.0.85 // indirect
	github.com/cdvelop/output v0.0.16 // indirect
	github.com/cdvelop/timeserver v0.0.20 // indirect
	github.com/cdvelop/timetools v0.0.21 // indirect
	github.com/cdvelop/wetest v0.0.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.18 // indirect
)

require (
	github.com/cdvelop/fetchserver v0.0.5 // indirect
	github.com/cdvelop/filehandler v0.0.8
	github.com/cdvelop/fileinput v0.0.32
	github.com/cdvelop/logserver v0.0.6 // indirect
	github.com/cdvelop/maps v0.0.7
	github.com/cdvelop/model v0.0.73
	golang.org/x/net v0.18.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/timeserver => ../timeserver

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
