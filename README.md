# IoT Assignment
Hyperledger Fabric Blockchain Healthcare PoC

## Prerequisites

We assume a working Hyperledger Fabric sample installation and test network operational with CAs set up. We also assume fabric binaries are in `PATH` and `FABRIC_CFG_PATH` is configured. See [Fabric documentation](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html) for details.

## Installation

In the root of repository there is a chaincode package named `vaccine.tar.gz`.
It needs to be installed into Hyperledger Fabric test network.

Place files `vaccine.tar.gz`, and anything under `scripts/` directory in fabric's test-network directory. Run `install_peer1.sh` and `install_peer2.sh` in sequence.

After that, effect can be validated by running `query_installed.sh`. Take note of the package id and save it as an environment variable (your id might be different):

```
export CC_PACKAGE_ID=vaccine_1:8a514eabcaa830105186e7674d87fc683cf334b1f9d4d5ae881fe0047d41427b
```

Next, you need to approve the chaincode as both organizations in the test-net by running `approve_peer1.sh` and `approve_peer2.sh`.

At this point the chaincode is approved and ready to be commited. You can check its status by running `check_commit.sh` and commit it by running `commit.sh`

Check the result with `query_commited.sh`.

You can validate that the chaincode was successfully commited by running `init.sh` and then `test.sh`

## Deployment details

[Fabric documentation](https://hyperledger-fabric.readthedocs.io/en/latest/deploy_chaincode.html)
