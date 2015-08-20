package robots

import (
	"net/http"
	"net/url"
	"bufio"
	"strings"
	"log"

	"github.com/a2n/alu"
)

type Robots struct {
	// Host > User-Agent > Allow | Disallow
	hosts map[string]map[string]map[string]bool
	logger *log.Logger
}

func NewRobots() *Robots {
	return &Robots{
		hosts: make(map[string]map[string]map[string]bool),
		logger: alu.NewLogger("robots.log"),
	}
}

func (r *Robots) IsAllowURLString(ua string, u string) bool {
	u1, err := url.Parse(u)
	if err != nil {
		log.Panicf("%s has error, %s.", alu.Caller(), err.Error())
	}
	return r.IsAllowURL(ua, u1)
}

func (r *Robots) IsAllowURL(ua string, u *url.URL) bool {
	if u == nil {
		return false
	}

	if len(u.String()) == 0 {
		return false
	}

	host := u.Scheme + "://" + u.Host
	if r.hosts[host] == nil {
		if r.get(host) != nil {
			return false
		}
	}

	// ok menas key is existing in map.
	_, ok := r.hosts[host]["*"]
	if ok {
		ua = "*"
	}

	for k, v := range r.hosts[host][ua] {
		if strings.Index(u.Path, k) == 0 {
			return v
		}
	}
	return true
}

func (r *Robots) get(host string) error {
	resp, err := http.DefaultClient.Get(host + "/robots.txt")
	if err != nil {
		return err
	}

	r.hosts[host] = make(map[string]map[string]bool)
	reader := bufio.NewReader(resp.Body)
	ua := ""
	for  {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			return nil
		}

		s := strings.Split(line, " ")
		if len(s) != 2 {
			// A newline.
			continue
		}

		s[1] = strings.Trim(s[1], "\n")
		switch s[0] {
		case "User-agent:":
			ua = s[1]
			r.hosts[host][s[1]] = make(map[string]bool)
		default:
			r.hosts[host][ua][s[1]] = s[0] == "Allow"
		}
	}

	resp.Body.Close()
	return nil
}
