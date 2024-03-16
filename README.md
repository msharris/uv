# UV
A simple command-line app that shows the current UV index for various locations around Australia.

## Usage
To show all locations, run:
```shell
uv
```

Show locations matching the name or ID (comma-separated list):
```shell
uv -l newcastle,syd,'Alice Springs'
```

Sort the output by a field:
```shell
uv -s [field]
```

Reverse the output:
```shell
uv -r
```

Reduce the output (quiet):
```shell
uv -q
```

## Disclaimer
UV observations courtesy of ARPANSA. See [Disclaimer](https://www.arpansa.gov.au/our-services/monitoring/ultraviolet-radiation-monitoring/ultraviolet-radation-data-information#Disclaimer)
