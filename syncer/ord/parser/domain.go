package parser

import "strings"

const (
	NameDomain = "text/plain;charset=utf-8"
)

const (
	Domain = "sns"
)

type NameDomainMint struct {
	P    string `json:"p"`
	Op   string `json:"op"`
	Name string `json:"name"`
}

func (m NameDomainMint) Validate() bool {
	if m.P != Domain {
		return false
	}

	if m.Op != "reg" {
		return false
	}
	if m.Name == "" {
		return false
	}

	return true
}

type NameDomainParser struct {
}

func (p *NameDomainParser) Name() string {
	return NameDomain
}

func (p *NameDomainParser) Parse(data []byte) (interface{}, bool, error) {
	var mint NameDomainMint
	err := json.Unmarshal(data, &mint)
	if err != nil {
		return nil, false, err
	}

	count := strings.Count(mint.Name, ".")
	if count != 1 {
		return nil, false, nil
	}

	return strings.ToLower(mint.Name), mint.Validate(), nil
}
