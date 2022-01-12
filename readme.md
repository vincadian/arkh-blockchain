# arkh
**arkh** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/vincadian/arkh-blockchain)

```
cd binary 
tar -xzvf binary.tar.gz
export PATH=$PATH:/workspace/arkh-blockchain/binary
cd
arkhd init validator
cp /workspace/arkh-blockchain/genesis/genesis.json .arkh/config

arkhd start --p2p.seeds 808f01d4a7507bf7478027a08d95c575e1b5fa3c@asc-dataseed.arkhadian.com:26656 --p2p.persistent_peers 808f01d4a7507bf7478027a08d95c575e1b5fa3c@asc-dataseed.arkhadian.com:26656,60c6a8074f75d69c75e7a464163d6652195eedc6@162.55.132.230:26656
```

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/arkhadian/arkh@latest! | sudo bash
```
`arkhadian/arkh` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
