package validatex

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

// ValidEmail checks if the email is valid
//
// According to RFC 5321:
// Path = Local-part@Domain
//
// 1. Path is up to 256 characters (including the punctuation and element separators)
//
// 2. Local-part is up to 64 characters
//
// 3. Domain is up to 255 characters
func ValidEmail(email string) error {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	if len(addr.Address) > 256 {
		return fmt.Errorf("email address is too long")
	}

	ss := strings.Split(addr.Address, "@")
	if len(ss) != 2 {
		return fmt.Errorf("missing '@' symbol")
	}

	if len(ss[0]) > 64 {
		return fmt.Errorf("part before '@' too long")
	}

	if len(ss[1]) > 255 {
		return fmt.Errorf("part after '@' too long")
	}

	if addr.Address != email {
		return fmt.Errorf("email address is not equal after parsing")
	}

	if !isValidDomain(ss[1]) {
		return fmt.Errorf("invalid domain")
	}

	return nil
}

func isValidDomain(domain string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(domain)
}
