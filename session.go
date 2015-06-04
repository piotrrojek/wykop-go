package wykop

type Session struct {
	appkey      string
	appsecret   string
	ClearOutput bool
}

func NewSession(appkey, appsecret string) (s Session) {
	s.appkey = appkey
	s.appsecret = appsecret
	s.ClearOutput = false
	return s
}
