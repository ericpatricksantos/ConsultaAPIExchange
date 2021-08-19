package Model

type UrlExchange struct {
	UrlAPI string
}

type ConnectionDb struct {
	Host   string
	Port   int
	User   string
	Passwd string
}

type UrlNoSql struct {
	UrlDb string
}
