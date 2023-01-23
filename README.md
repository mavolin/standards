<div align="center">
<h1>standards</h1>

[![Go Reference](https://pkg.go.dev/badge/github.com/mavolin/standards.svg)](https://pkg.go.dev/github.com/mavolin/standards)
[![Test](https://github.com/mavolin/standards/actions/workflows/test.yml/badge.svg)](https://github.com/mavolin/standards/actions)
[![Code Coverage](https://codecov.io/gh/mavolin/standards/branch/develop/graph/badge.svg?token=ewFEQGgMES)](https://codecov.io/gh/mavolin/standards)
[![Go Report Card](https://goreportcard.com/badge/github.com/mavolin/standards)](https://goreportcard.com/report/github.com/mavolin/standards)
[![License MIT](https://img.shields.io/github/license/mavolin/standards)](https://github.com/mavolin/standards/blob/develop/LICENSE)
</div>

---

## About

Boringly simple parsing and validation for various notations.

## Supported Standards

* 🏦 BICs
* 💰 IBANs with country-specific BBAN validation
* 🏴‍☠️ ISO3166-1 Alpha2 (e.g. `DE`, or `ES`)
* 🚑 German Health Insurance Numbers (Krankenversicherungsnummern)
* 🧓 German Pension Insurance Numbers (Renten-/ Sozialversicherungsnummern)
* 💲 German Tax Identification Numbers (Steuer-IDs)
* ✉ German Postal Codes (Postleitzahlen)

## Each Package Is the Same

Each standard is implemented in its own package,
and each package provides a type for the standard, e.g. `bic.BIC` for BICs.
That type contains all the information that can be extracted from the notation.

```go
package bic

type BIC struct {
	BusinessPartyPrefix string
	CountryCode         string
	BusinessPartySuffix string
	BranchCode          string
}
```

Each type implements:

* `String() string` to get the notation pretty-printed
* `MarshalText() ([]byte, error)` to get the notation in compact form
* `UnmarshalText([]byte) error` to parse the notation

Additionally, each package provides these two functions:

* `Parse(string) (Type, error)` parses the given string, and validates it
* `IsValid(string) bool` simply checks whether the given string is valid in the given standard.

## License

Built with ❤ by [Maximilian von Lindern](https://github.com/mavolin).
Available under the [MIT License](./LICENSE).