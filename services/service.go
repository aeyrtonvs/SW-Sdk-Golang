package services

type Service struct {
	Url       string
	User      string
	Password  string
	Token     string
	ProxyPort int
	Proxy     string
}

func (service *Service) Init(url string, token string, proxyPort int, proxy string) {
	service.Url = url
	service.Token = token
	service.ProxyPort = proxyPort
	service.Proxy = proxy
}

func (service *Service) InitAuth(url string, user string, password string, proxyPort int, proxy string) {
	service.Url = url
	service.User = user
	service.Password = password
	service.ProxyPort = proxyPort
	service.Proxy = proxy
}
