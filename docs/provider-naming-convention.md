# Provider naming convention

There are more Terraform providers developed by JFrog, e.g.:

- [`jfrog/artifactory`](https://registry.terraform.io/providers/jfrog/artifactory)
- [`jfrog/platform`](https://registry.terraform.io/providers/jfrog/platform)
- [`jfrog/project`](https://registry.terraform.io/providers/jfrog/project)
- etc.

So, I decided to use `artifactory.jfrog.crossplane.io` base group for this Crossplane provider.

## Options

1. (*current naming convention*) `artifactory.jfrog.crossplane.io` as base group and resource is `OCIRepository` for each short group (`local`, `remote`, `virtual`, `federated`)

    - :heavy_plus_sign: `ShortGroup` doesn't have to be set for all resources, just for specific resources
    - Auto generated `ShortGroup` for Group of resources like `local.artifactory.jfrog.crossplane.io` for Local Repositories according to Terraform provider (by default)
      - :heavy_plus_sign: Crossplane v2 introduced [managed resource definitions (MRDs)](https://docs.crossplane.io/latest/managed-resources/managed-resource-definitions/) for selective activation of provider resources, reducing [cluster overhead](https://docs.crossplane.io/v2.0/guides/disabling-unused-managed-resources/#the-problem-resource-overhead). It allows to activate only specific CRDs, so it's helpful to have separated groups like `local.artifactory.jfrog.crossplane.io` or `virtual.artifactory.jfrog.crossplane.io` to be able to activate only some of them in case users don't need them all.
      - :heavy_minus_sign: Can be less clear, because there will be the same resources, e.g. `AlpineRepository` for groups `local.artifactory.jfrog.crossplane.io`, `remote.artifactory.jfrog.crossplane.io`, etc. (all resources must have specified `ShortGroup` as `""` (empty string) to override default behavior and have `AlpineRepository` with `artifactory.jfrog.crossplane.io` group)

2. `artifactory.jfrog.crossplane.io` as base group and resource is `OCIRepository`

    - :heavy_plus_sign: `ShortGroup` doesn't have to be set for all resources
    - :heavy_minus_sign: Still need to specify `ShortGroup` for the most of resources
    - Auto generated `ShortGroup` for Group of resources like `local.artifactory.jfrog.crossplane.io` for Local Repositories according to Terraform provider (by default)
      - :heavy_plus_sign: Crossplane v2 introduced [managed resource definitions (MRDs)](https://docs.crossplane.io/latest/managed-resources/managed-resource-definitions/) for selective activation of provider resources, reducing [cluster overhead](https://docs.crossplane.io/v2.0/guides/disabling-unused-managed-resources/#the-problem-resource-overhead). It allows to activate only specific CRDs, so it's helpful to have separated groups like `local.artifactory.jfrog.crossplane.io` or `virtual.artifactory.jfrog.crossplane.io` to be able to activate only some of them in case users don't need them all.
      - :heavy_minus_sign: Can be less clear, because there will be the same resources, e.g. `AlpineRepository` for groups `local.artifactory.jfrog.crossplane.io`, `remote.artifactory.jfrog.crossplane.io`, etc. (all resources must have specified `ShortGroup` as `""` (empty string) to override default behavior and have `AlpineRepository` with `artifactory.jfrog.crossplane.io` group)

3. `jfrog.crossplane.io` as a base group and set `ShortGroup` as `artifactory` to all resources which will produce `artifactory.jfrog.crossplane.io` and K8s kind is just `OCIRepository`

    - :heavy_minus_sign: All resources must have specified `ShortGroup` as `artifactory`

4. `artifactory.crossplane.io` as a base group and resource prefix to be `artifactory`, like `OCIRepository`

    - :heavy_minus_sign: It's not clear for other JFrog providers, e.g. for `jfrog/platform` would be `platform.crossplane.io`

Not sure, which one is the best.
