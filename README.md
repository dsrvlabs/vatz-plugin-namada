# vatz-plugin-namada
Vatz plugin for namada governance monitoring

## Plugins
- node_governance_alarm : monitor the new governance proposal

## Installation and Usage
> Please make sure [Vatz](https://github.com/dsrvlabs/vatz) is running with proper configuration. [Vatz Installation Guide](https://github.com/dsrvlabs/vatz/blob/main/docs/installation.md)

### Install Plugins
- Install with source
```
$ git clone https://github.com/dsrvlabs/vatz-plugin-namada.git
$ cd vatz-plugin-namada
$ make install
```
- Install with Vatz CLI command
```
$ ./vatz plugin install --help
Install new plugin

Usage:
   plugin install [flags]

Examples:
vatz plugin install github.com/dsrvlabs/<somewhere> name

Flags:
  -h, --help   help for install
```
> please make sure install path for the plugins repository URL.
```
$ ./vatz plugin install github.com/dsrvlabs/vatz-plugin-namada/plugins/node_governance_alarm node_governance_alarm
```
- Check plugins list with Vatz CLI command
```
$ vatz plugin list
+-----------------------+------------+---------------------+-------------------------------------------------------------------+---------+
| NAME                  | IS ENABLED | INSTALL DATE        | REPOSITORY                                                        | VERSION |
+-----------------------+------------+---------------------+-------------------------------------------------------------------+---------+
| node_governance_alarm | true       | 2023-09-27 01:15:59 | github.com/dsrvlabs/vatz-plugin-namada/plugins/node_governance_alarm | latest  |
+-----------------------+------------+---------------------+-------------------------------------------------------------------+---------+
```

### Run
> Run as default config or option flags

```
# You can run node_governance_alarm with default config without any argument or with your configuration
$ node_governance_alarm -port 10001 -chainId shielded-expedition.88f17d1d14 -base-dir /mnt/namada -nodeInfo http://127.0.0.1:26657
2024-02-24T13:35:39Z INF Register module=grpc
2024-02-24T13:35:39Z INF Start 127.0.0.1 10001 module=sdk
2024-02-24T13:35:39Z INF Start module=grpc
2024-02-24T13:37:34Z INF Execute module=grpc
2024-02-24T13:37:35Z DBG DEBUG : tmp != proposalId module=plugin
2024-02-24T13:37:35Z INF Last Proposal ID: 194
Type: PGF steward
Author: tnam1qppwtk3sp08gg9j60sqa3gmxn3q5pzn25s5hpccl
Start Epoch: 16
End Epoch: 18
Grace Epoch: 20 module=plugin
```
## Command line arguments

- node_governance_alarm
```
Usage of node_governance_alarm:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -port int
    	Port number (default 10001)
  -chainId string
    	Need to Chain ID (default:shielded-expedition.88f17d1d14)
  -baseDir string
    	Need to Base Dir (default "/mnt/namada")
  -nodeInfo string
    	Need to know Node Info (default "http://0.0.0.0:26657")  	
```

## TroubleShooting
1. Encountered issue related with `Device or Resource Busy` or `Too many open files` error.
 - Check your open file limit and recommended to increase it.
 ```
 $ ulimit -n
 1000000
 ```

## License

`vatz-plugin-namada` is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included in our repository in the `LICENSE` file.
