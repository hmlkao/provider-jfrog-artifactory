# Keypair

1. Generate SSH key pair for test purpose

    > [!NOTE]
    > Do NOT use `ssh-keygen` to generate key pair to prevent `Error: RSA private key is of the wrong type.` error message. ([source](https://jfrog.com/help/r/jfrog-artifactory-documentation/setting-up-rsa-keys-pairs))

    ```bash
    # Create pem file with private key
    openssl genrsa -traditional -out crossplane-ssh-key.pem 2048
    # Extract pub part from this file
    openssl rsa -in crossplane-ssh-key.pem -pubout -out crossplane-ssh-key.pub.pem
    ```

2. Generate Kubernetes `Secret` resource

    ```bash
    private_key=$(cat crossplane-ssh-key.pem)
    jq '.stringData += {"private-key":"'"$private_key"'"}' < examples/security/keypair/secret.json.tmpl > examples/security/keypair/secret.json
    ```

3. Create Kubernetes `Secret` resource in cluster

    ```bash
    kubectl apply -f examples/security/keypair/secret.json
    ```

4. Generate Kubernetes `Keypair` resource

    ```bash
    public_key=$(cat crossplane-ssh-key.pub.pem)
    jq '.spec.forProvider += {"publicKey":"'"$public_key"'"}' < examples/security/keypair/keypair.json.tmpl > examples/security/keypair/keypair.json
    ```

5. Create Kubernetes `Keypair` resource in cluster

    ```bash
    kubectl apply -f examples/security/keypair/keypair.json
    ```
