# ehs

**WIP**

Ex-HostScan.

Inspired by [Hostscan](https://github.com/cckuailong/hostscan)

This project rebuild origin hostscan, removed other requirements.
Change the output format to grepable, and supported read sub-domain word
from STDIN, it's easier to integrate to other tool chain.

## Install

### `go install`

```
go install github.com/xbol0/ehs@latest
```

## Examples

```
# Basic usage
ehs -f sub_dict.txt -d example.com 1.2.3.4

# Multi ip addresses
ehs -f sub_dict.txt -d example.com 1.2.3.4 2.3.4.5 3.4.5.6

# Use ip list file
cat ip_list.txt | xargs ehs -f sub_dict.txt -d example.com

# Import sub-domain list from STDIN
cat sub_domain.txt | ehs -d example.com 1.2.3.4

# Use online dictionary
curl -fsSL https://some.dictionary/file.txt | ehs -d example.com 1.2.3.4

# Both sub-domain and ip list
curl -fsSL https://dict.com/file.txt | xargs ehs -d example.com

# Use keyword, it will concat with top root domain such as com, net, org
ehs -d example -f dict.txt 1.2.3.4

# Filter by http status code
ehs -f sub.txt -d example.com 1.2.3.4 | grep 200

# Get filtered hostname
ehs -f sub.txt -d example.com 1.2.3.4 | awk '/200/{print $3}'

# Configure timeout, second
ehs -t 5

# Configure threads count
ehs -n 10
```