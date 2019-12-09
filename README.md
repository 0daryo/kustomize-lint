# kustomize-lint

This tool helps you to validate tag in ci
Please post issues if you need more options

In this kind of architecture below, you have to check whether yamls in production directory has sentences related to staging or development by code review.

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

If you use this lint, ci exit with code 1 when production kustomization has wrong tag. for ex)

```user/overlays/production/kustomization.yaml
images:
  - name: nginx-production
    # need to be productionXX, but is mistakenly staging
    newTag: staging
```

# note

## supported tag in kustomization

- image
- newTag

# usage

1. `go get -u github.com/0daryo/kustomize-lint` or get binary from [release page](https://github.com/0daryo/kustomize-lint/releases)
2. write `kustomize-lint.yaml` following the rule in kustomize-lint-ex.yaml

```
files:
  - file:
    name: "*/production/kustomization.yaml"
    sentences:
    # words tag must have
      - name: newTag
        include:
          - production
          - latest
      - name: name
        include:
          - production
```

3. run `kustomize-lint run` with a directory where target yamls are.
