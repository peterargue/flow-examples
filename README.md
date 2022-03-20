# flow-examples
This is a collection of flow related example code.

### Selecting Access Nodes to query
Which hostname you use to connect the an Access Node depends on how you are planning to connect. The details here assume
you want to use the public Access Nodes hosted by Dapper Labs. If you planning to use a different node you trust, some of
the following may not apply.

#### gRPC
Plaintext:

This is the hostname that most people should unless you want to use secure gRPC. It returns the IPs for all public Acess Nodes
and it kept up to date between sporks:
* `access.mainnet.nodes.onflow.org:9000`

You can also connect to individual nodes by their hostname. Each spork, there are a new set of Access nodes created, which will
have different hostnames. In the examples below `mainnet16` would become `mainnet17`, and so on.
* `access-001.mainnet16.nodes.onflow.org:9000`
* `access-002.mainnet16.nodes.onflow.org:9000`
* `access-003.mainnet16.nodes.onflow.org:9000`
* `access-004.mainnet16.nodes.onflow.org:9000`

Secure:

When connecting over TLS, you need to connect to the specific node since TLS certificates are tied to the node. See [secure-grpc](examples/secure-grpc/README.md) for more details.
* `access-001.mainnet16.nodes.onflow.org:9001`
* `access-002.mainnet16.nodes.onflow.org:9001`
* `access-003.mainnet16.nodes.onflow.org:9001`
* `access-004.mainnet16.nodes.onflow.org:9001`

#### gRPC-Web
Secure:
* `https://mainnet.onflow.org`

Plaintext:
* `access.mainnet.nodes.onflow.org:8000`
* `access-001.mainnet16.nodes.onflow.org:8000`
* `access-002.mainnet16.nodes.onflow.org:8000`
* `access-003.mainnet16.nodes.onflow.org:8000`
* `access-004.mainnet16.nodes.onflow.org:8000`

#### REST
Secure:
* `https://rest-mainnet.onflow.org`

Plaintext:
* `access.mainnet.nodes.onflow.org:8070`
* `access-001.mainnet16.nodes.onflow.org:8070`
* `access-002.mainnet16.nodes.onflow.org:8070`
* `access-003.mainnet16.nodes.onflow.org:8070`
* `access-004.mainnet16.nodes.onflow.org:8070`

### Finding Public Access Node for a spork
When using the aggregated domain names, the backend Access Node IPs are updated automatically each spork. If you are connecting to individual nodes, you will need to use the correct
domain name for that node for that spork. Generally, only the middle spork name changes. e.g. `mainnet16` becomes `mainnet17`.

For a more precise way to lookup the nodes, you can fetch the node details from blockchain bootstrapping files.

#### Get the list of nodes for a spork

First, download the spork metadata from github and extract the `node-infos` URL. The `node-infos.pub.json` includes the public keys for Access Nodes participating in the spork.
Note: you will need to change `mainnet16` to the name of the spork you are interested in.
```
wget https://raw.githubusercontent.com/onflow/flow/master/sporks.json
jq -r '.networks.mainnet.mainnet16.stateArtefacts.gcp.nodeInfo' sporks.json
```

#### Get the address and public key for a node

Use the URL from the previous section to download the node-info file from GCP and extract the public access nodes' public keys.
```
wget https://storage.googleapis.com/flow-genesis-bootstrap/mainnet-16-execution/public-root-information/node-infos.pub.json
jq '.[] | select(.Address | test("^access-00[1-4]")?)' node-infos.pub.json
```

Remove the port from the `Address` field, and replace it with the correct port for your connection.
The `NetworkPubKey` field contains the public key used when verifying gRPC TLS connections.
 