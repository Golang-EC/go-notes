package main

type Practice struct {
	name         string
	initDate     string
	endDate      string
	timeDuration string
}

func New(name string, initDate string, endDate string, timeDuration string) Practice {
	p := Practice{name, initDate, endDate, timeDuration}
	return p
}
