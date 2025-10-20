# Upgrade instructions

Follow these instructions for breaking changes.

## `v0.9.1`

The scope for `providerconfigusages.artifactory.jfrog.m.crossplane.io` resource was changed from `Cluster` to `Namespaced`.

If you used `providerconfigs.artifactory.jfrog.m.crossplane.io` to set up `ProviderConfig` for the `jfrog-artifactory` provider, follow these instructions:

1. Create `ClusterProviderConfig` with the same configuration as you have in the current `ProviderConfig`
2. Change `spec.providerConfigRef.kind` of all resources from `ProviderConfig` to `ClusterProviderConfig`
3. Remove all `providerconfigusages.artifactory.jfrog.m.crossplane.io` resources
4. Remove all `providerconfigs.artifactory.jfrog.m.crossplane.io` resources
5. Remove CRD `providerconfigusages.artifactory.jfrog.m.crossplane.io`

    ```shell
    kubectl delete crd providerconfigusages.artifactory.jfrog.m.crossplane.io
    ```

6. Remove CRD `providerconfigs.artifactory.jfrog.m.crossplane.io`

    ```shell
    kubectl delete crd providerconfigs.artifactory.jfrog.m.crossplane.io
    ```

7. Upgrade the provider to the `v0.9.1` version
8. Check if is the provider healthy

    ```shell
    kubectl get providers provider-jfrog-artifactory
    ```

9.  Create some resource using this provider to ensure it works properly

## `v0.9.0`

The scope for `providerconfigs.artifactory.jfrog.m.crossplane.io` resource was changed from `Cluster` to `Namespaced`.

If you used `providerconfigs.artifactory.jfrog.m.crossplane.io` to set up `ProviderConfig` for the `jfrog-artifactory` provider, follow these instructions:

1. Create `ClusterProviderConfig` with the same configuration as you have in the current `ProviderConfig`
2. Change `spec.providerConfigRef.kind` of all resources from `ProviderConfig` to `ClusterProviderConfig`
3. Remove all `providerconfigs.artifactory.jfrog.m.crossplane.io` resources
4. Remove CRD `providerconfigs.artifactory.jfrog.m.crossplane.io`
5. Upgrade the provider to the `v0.9.0` version
6. Check if is the provider healthy
7. Create some resource using this provider to it works properly
