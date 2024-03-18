package utils

const (
	HOST = "smtp.gmail.com"
	PORT = "587"
)

type Email struct {
	EmailFrom string
	Password  string
	Subject   string
	body      string
	Host      string
	port      string
	Address   string
	EmailTo   []string
}
