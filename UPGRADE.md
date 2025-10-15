# Upgrade instructions

Follow these instructions for breaking changes

## `v0.9.0`

Scope for `providerconfigs.artifactory.jfrog.m.crossplane.io` resource was changed from `Cluster` to `Namespaced`.

If you used `providerconfigs.artifactory.jfrog.m.crossplane.io` to set up `ProviderConfig` for `jfrog-artifactory` provider, follow these instructions:

1. Create `ClusterProviderConfig` with the same configuration as you have in current `ProviderConfig`
2. Change `providerConfigRef.kind` of all resources from `ProviderConfig` to `ClusterProviderConfig`
3. Remove all `providerconfigs.artifactory.jfrog.m.crossplane.io` resources
4. Remove CRD `providerconfigs.artifactory.jfrog.m.crossplane.io`
5. Upgrade provider to the `v0.9.0` version
6. Check if provider started (is in `Running` state)
7. Create some resource using this provider to it works properly
