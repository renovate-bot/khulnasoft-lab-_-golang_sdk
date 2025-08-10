module github.com/khulnasoft-lab/golang_sdk/utils

go 1.23

toolchain go1.24.6

replace github.com/khulnasoft-lab/golang_sdk/client => ../client

require (
	github.com/hashicorp/go-retryablehttp v0.7.8
	github.com/khulnasoft-lab/golang_sdk/client v0.0.0-20240520213426-d989e5f20024
)

require github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
