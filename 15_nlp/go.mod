module github.com/Raghulds/nlp

go 1.25.2

// Risks of using 3rd party packages
// - Security (ignorance, intentional)
// - Bugs
// - Copmatibility (API changes)
// - Legal (licenses)
// - Might be gone (go mod vendor)

replace github.com/Raghulds/stemmer => ../15_stemmer

require github.com/stretchr/testify v1.11.1

require (
	github.com/Raghulds/stemmer v0.0.0-00010101000000-000000000000 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
