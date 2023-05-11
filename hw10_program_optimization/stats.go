package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/goccy/go-json"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	var user User

	domains := make(DomainStat)
	reader := bufio.NewReader(r)

	for {
		line, isPrefix, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		if isPrefix {
			continue
		}

		if err = json.Unmarshal(line, &user); err != nil {
			return nil, err
		}

		if strings.HasSuffix(user.Email, domain) {
			d := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			domains[d]++
		}
	}
	return domains, nil
}
