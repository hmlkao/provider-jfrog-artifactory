# Upgrade instructions

Follow these instructions for breaking changes.

## `v0.11.1`

The name of MRDs were changed and need to be removed first because the `spec.names` is immutable.

If you use some of these MRDs, you have to remove created resources first.

```shell
kubectl delete mrds \
terraformbackendrepositories.local.artifactory.jfrog.crossplane.io \
terraformbackendrepositories.local.artifactory.jfrog.m.crossplane.io \
vcsrepositories.remote.artifactory.jfrog.crossplane.io \
vcsrepositories.remote.artifactory.jfrog.m.crossplane.io
```

## `v0.11.0`

The name of MRDs were changed and need to be removed first because the `spec.names` is immutable.

If you use some of these MRDs, you have to remove created resources first.

```shell
kubectl delete mrds \
cocoapodsrepositories.federated.artifactory.jfrog.crossplane.io \
cocoapodsrepositories.federated.artifactory.jfrog.m.crossplane.io \
cranrepositories.federated.artifactory.jfrog.crossplane.io \
cranrepositories.federated.artifactory.jfrog.m.crossplane.io \
gitlfsrepositories.federated.artifactory.jfrog.crossplane.io \
gitlfsrepositories.federated.artifactory.jfrog.m.crossplane.io \
helmocirepositories.federated.artifactory.jfrog.crossplane.io \
helmocirepositories.federated.artifactory.jfrog.m.crossplane.io \
huggingfacemlrepositories.federated.artifactory.jfrog.crossplane.io \
huggingfacemlrepositories.federated.artifactory.jfrog.m.crossplane.io \
npmrepositories.federated.artifactory.jfrog.crossplane.io \
npmrepositories.federated.artifactory.jfrog.m.crossplane.io \
nugetrepositories.federated.artifactory.jfrog.crossplane.io \
nugetrepositories.federated.artifactory.jfrog.m.crossplane.io \
ocirepositories.federated.artifactory.jfrog.crossplane.io \
ocirepositories.federated.artifactory.jfrog.m.crossplane.io \
opkgrepositories.federated.artifactory.jfrog.crossplane.io \
opkgrepositories.federated.artifactory.jfrog.m.crossplane.io \
pypirepositories.federated.artifactory.jfrog.crossplane.io \
pypirepositories.federated.artifactory.jfrog.m.crossplane.io \
rpmrepositories.federated.artifactory.jfrog.crossplane.io \
rpmrepositories.federated.artifactory.jfrog.m.crossplane.io \
sbtrepositories.federated.artifactory.jfrog.crossplane.io \
sbtrepositories.federated.artifactory.jfrog.m.crossplane.io
```

## `v0.9.1`

The scope for `providerconfigusages.artifactory.jfrog.m.crossplane.io` resource was changed from `Cluster` to `Namespaced`.

If you used `providerconfigs.artifactory.jfrog.m.crossplane.io` to set up `ProviderConfig` for the `jfrog-artifactory` provider, follow these instructions:

1. Create `ClusterProviderConfig` with the same configuration as you have in the current `ProviderConfig`
2. Change `spec.providerConfigRef.kind` of all resources from `ProviderConfig` to `ClusterProviderConfig`
3. Upgrade the provider to the `v0.9.1` version
4. Remove all `providerconfigusages.artifactory.jfrog.m.crossplane.io` resources
5. Remove all `providerconfigs.artifactory.jfrog.m.crossplane.io` resources

    If is `ProviderConfig` not removed, check `ProviderConfigUsages` again.

6. Remove CRD `providerconfigusages.artifactory.jfrog.m.crossplane.io`

    ```shell
    kubectl delete crd providerconfigusages.artifactory.jfrog.m.crossplane.io
    ```

7. Remove CRD `providerconfigs.artifactory.jfrog.m.crossplane.io`

    ```shell
    kubectl delete crd providerconfigs.artifactory.jfrog.m.crossplane.io
    ```

8. Check if is the provider healthy (may take up to 5 minutes)

    ```shell
    kubectl get providers provider-jfrog-artifactory
    ```

9. Edit some resource which was already in cluster and check if was the change applied properly
10. Create some resource using this provider to ensure it works properly

## `v0.9.0`

The scope for `providerconfigs.artifactory.jfrog.m.crossplane.io` resource was changed from `Cluster` to `Namespaced`.

If you used `providerconfigs.artifactory.jfrog.m.crossplane.io` to set up `ProviderConfig` for the `jfrog-artifactory` provider, follow these instructions:

1. Create `ClusterProviderConfig` with the same configuration as you have in the current `ProviderConfig`
2. Change `spec.providerConfigRef.kind` of all resources from `ProviderConfig` to `ClusterProviderConfig`
3. Upgrade the provider to the `v0.9.0` version
4. Remove all `providerconfigs.artifactory.jfrog.m.crossplane.io` resources
5. Remove CRD `providerconfigs.artifactory.jfrog.m.crossplane.io`
6. Check if is the provider healthy
7. Create some resource using this provider to it works properly
