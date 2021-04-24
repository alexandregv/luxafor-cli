# luxafor-cli

A dead simple CLI for Luxafor.

### Installation

```
go install github.com/alexandregv/luxafor-cli@latest
# or
docker pull alexandregv/luxafor-cli
```

### Usage

```
# luxafor-cli <red> <green> <blue>

sudo luxafor-cli 0 255 0
# or
docker run --privileged -v /dev/bus/usb:/dev/bus/usb alexandregv/luxafor-cli 0 255 0
```
