

## v1.0.0
This is initial release of cshtop

- It will fetch tendermint rpc uri and ldc uri from [cosmos chain registry](https://github.com/cosmos/chain-registry).

- It will fetch price from [coingecko api](https://www.coingecko.com/) every 5 seconds.


**Note**

> You can override the tendermint rpc api and rest uri
```bash
$ cshtop start --app cosmoshub --rest https://lcd-cosmoshub.whispernode.com:443 --rpc https://rpc-cosmoshub.whispernode.com:443
```