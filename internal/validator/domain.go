package validator

import (
	"net"
	"strings"

	"github.com/coinbase/baseca/internal/config"
)

var valid_domains []string
var valid_certificate_authorities []string

func IsValidateDomain(fully_qualified_domain_name string) bool {

	arr := strings.Split(fully_qualified_domain_name, ".")

	if len(arr) < 2 {
		return false
	}

	domain_slice := arr[len(arr)-2:]
	domain := strings.Join(domain_slice, ".")

	for _, valid_domain := range valid_domains {
		if domain == valid_domain {
			return true
		}
	}

	// Fallback Check IP Address for CN/SAN
	return net.ParseIP(fully_qualified_domain_name) != nil
}

func IsSupportedCertificateAuthority(certificate_authority string) bool {
	for _, ca := range valid_certificate_authorities {
		if ca == certificate_authority {
			return true
		}
	}
	return false
}

func SupportedConfig(cfg *config.Config) {
	valid_domains = cfg.Domains

	for certificate_authority := range cfg.ACMPCA {
		valid_certificate_authorities = append(valid_certificate_authorities, certificate_authority)
	}
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func SanitizeInput(input []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range input {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
