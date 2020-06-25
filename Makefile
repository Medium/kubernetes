.PHONY: gen-matchers test

pkg/test/core/v1/zz_generated.matcher.go: cmd/matcher-gen/generators/generator.go cmd/matcher-gen/generators/matcher.go
	go run ./cmd/matcher-gen -i k8s.io/api/core/v1 -p go.medium.engineering/kubernetes/pkg/test/core/v1 -v5

pkg/test/istio/networking/v1alpha3/zz_generated.matcher.go: cmd/matcher-gen/generators/generator.go cmd/matcher-gen/generators/matcher.go
	go run ./cmd/matcher-gen -i istio.io/client-go/pkg/apis/networking/v1alpha3 -p go.medium.engineering/kubernetes/pkg/test/istio/networking/v1alpha3 -v5

gen-matchers: pkg/test/istio/networking/v1alpha3/zz_generated.matcher.go pkg/test/core/v1/zz_generated.matcher.go

test:
	go test ./...
