poll nodeos and print actions that match filtering conditions.

```
go install github.com/hayeah/eostools/eos-actions
```

To match all actions dispatched to the account `hello`:

```
> eos-actions --account hello

{"hex_data":"0000000000ea3055","data":{"user":"eosio"},"authorization":[{"permission":"active","actor":"eosio"}],"name":"hi","account":"hello"}
{"hex_data":"0000000000ea3055","data":{"user":"eosio"},"authorization":[{"permission":"active","actor":"eosio"}],"name":"yo","account":"hello"}
```

# DOC

```
eos-actions --help
usage: eos-actions [<flags>]

firehose of eos actions

Flags:
  --help             Show context-sensitive help (also try --help-long and --help-man).
  --account=ACCOUNT  filter action by account name
  --action=ACTION    filter action by action name
```
