# kustomize-lint

This tool helps you to validate tag in ci
Please post issues if you need more options

In this kind of architecture below, you have to check whether yamls in production directory has sentences related to staging or development by code review.
If you use this lint, ci exit with code 1 when production or staging has mistaken sentences.

```
├── user
│   ├── base
│   │   ├── deployment.yaml
│   │   ├── kustomization.yaml
│   │   └── service.yaml
│   └── overlays
│       ├── development
│       │   ├── deployment.yaml
│       │   ├── kustomization.yaml
│       │   ├── service.yaml
│       ├── production
│       │   ├── deployment.yaml
│       │   ├── kustomization.yaml
│       │   ├── service.yaml
│       └── staging
│           ├── deployment.yaml
│           ├── kustomization.yaml
│           ├── service.yaml
```

# note
## functions
- if newTag section in production/kustomization.yaml is productionXX or latest
- if newTag section in staging/kustomization.yaml is stagingXX or latest

# usage

1. `go get -u github.com/0daryo/kustomize-lint`
2.
3. run `kustomize-lint run` with a directory where target yamls are
