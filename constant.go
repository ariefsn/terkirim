package terkirim

type Service string

const (
	DefaultUrl string  = "https://terkirim.cloud/api"
	Whatsapp   Service = "whatsapp"
	Email      Service = "email"
)

var serviceUrlPaths = map[Service]string{
	Whatsapp: "/whatsapp",
	Email:    "/email",
}
