package mailbag

type (
	password string
	userName string
	url      string
	note     string
)

type Secret struct {
	password password `schema.org: "password"`
	userName userName `schema.org: "username"`
	url      url      `schema.org: "url"`
	note     note     `schema.org: "note"`
}

func (s *Secret) Password() string {
	return string(s.password)
}

func (s *Secret) SetPassword(value string) {
	s.password = password(value)
}

func (s *Secret) UserName() string {
	return string(s.userName)
}

func (s *Secret) SetUserName(value string) {
	s.userName = userName(value)
}

func (s *Secret) Url() string {
	return string(s.url)
}
func (s *Secret) SetUrl(value string) {
	s.url = url(value)
}

func (s *Secret) Note() string {
	return string(s.note)
}

func (s *Secret) SetNote(value string) {
	s.note = note(value)
}
