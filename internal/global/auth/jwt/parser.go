package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/pkg/errors"
)

type Parser interface {
	Parse(token string) (*Token, error)
}

type parser struct {
	secret       []byte
	sigingMethod jwt.SigningMethod
}

func NewParser(secret []byte) Parser {
	return &parser{
		secret:       secret,
		sigingMethod: jwt.GetSigningMethod(constant.JwtSigningMethod),
	}
}

// Parse parses given string token into Token. If validation fails, error is returned.
func (p *parser) Parse(token string) (*Token, error) {
	var tokn Token
	_, err := jwt.ParseWithClaims(token, &tokn, func(t *jwt.Token) (interface{}, error) {
		if err := p.validate(t); err != nil {
			return nil, err
		}

		return p.secret, nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "jwt token validation falied")
	}

	return &tokn, nil
}

func (p *parser) validate(t *jwt.Token) error {
	if t.Method != p.sigingMethod {
		return errors.Errorf("siging method mismatch: expected: %s. actual: %s", p.sigingMethod.Alg(), t.Method.Alg())
	}
	return nil
}
