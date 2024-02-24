# vatz-plugin-namda
Vatz plugin for namada governance monitoring

## Plugins
- node_governance_alarm : monitor the new governance proposal and whether or not to vote

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
# Your node have to enable API configuration ({HOME_DIR}/config/app.toml)
$ node_governance_alarm -apiPort <API server port{default is 1317}> -voterAddr <Account Address>
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10005 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z DBG DEBUG : tmp == proposalId module=plugin
2023-05-31T07:08:10Z INF Lastest proposal is #51
 module=plugin
```
## Command line arguments

- node_governance_alarm
```
Usage of node_governance_alarm:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -apiPort uint
    	Need to know proposal id (default 1317)
  -port int
    	Port number (default 10005)
  -proposalId uint
    	Need to know last proposal id
  -voterAddr string
    	Need to voter address (default "address")
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
