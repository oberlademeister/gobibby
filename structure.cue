package bib

import (
	"time"
)

#Book: {
	id!: string
	authors!: [...string]
	editors?: [...string]
	title!:     string
	publisher!: string
	year!:      >=1500 & <=2100
	month?:     >=1 & <=12
	isbn?:      =~"([0-9]{10,10}|[0-9]{13,13})"
	edition?:   string
	address?:   string
	abstract?:   string
}

#Internet: {
	id!: string
	authors!: [...string]
	title!:        string
	url!:          string
	accessedDate!: time.Format("2006-01-02")
}

#SAPInternal: {
	id!: string
	authors!: [...string]
	title!:        string
	url!:          string
	accessedDate!: time.Format("2006-01-02")
}

#Wikipedia: {
	id!:           string
	title!:        string
	url!:          string
	accessedDate!: time.Format("2006-01-02")
}
