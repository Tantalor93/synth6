# synth6
tool for DNS64 address synthesis and related utilities, provides this functionality:
* IPv6 address synthesis based on provided IPv6 prefix and IPv4 address according to [RFC6147](https://www.rfc-editor.org/rfc/rfc6147)
* IPv6 prefix discovery using [RFC7050](https://datatracker.ietf.org/doc/rfc7050/) or based on previous IPv6 response of DNS discovery 

# Installation 
using standard go tooling
```
go install github.com/tantalor93/synth6
```

# Usage
```
Usage:
  synth6 [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  prefix      IPv6 prefix discovery for IPv6 address synthesis either using RFC7050 or based on already provided IPv6 address from ipv4only.arpa response
  synth       Synthesis of IPv6 address from IPv4 according to RFC6147

Flags:
  -h, --help   help for synth6

Use "synth6 [command] --help" for more information about a command.
```

# Examples

Discover IPv6 prefix using DNS64 server according to [RFC6147](https://www.rfc-editor.org/rfc/rfc6147), 
**works when you are on IPv6 only network with DNS64 server!**
```
synth6 prefix
64:ff9b::/96
```

Discover IPv6 prefix from previous response of IPv6 discovery using DNS64, 
this is usable when you have only pcap and your computer is not physically on IPv6-only network
```
synth6 prefix 64:ff9b::c000:ab
64:ff9b::/96
```

Synthesis of IPv6 address from IPv6 prefix and IPv4 address
```
synth6 synth 64:ff9b::/96 192.0.0.171
64:ff9b::c000:ab
```
