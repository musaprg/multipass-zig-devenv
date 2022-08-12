# multipass-zig-devenv
Zig devenv built on top of multipass VM

## Usage

### Launch VM with multipass

```
$ multipass-zig-devenv -name hoge -cpus 2 -mem 4G -disk 20G -image 20.04 launch
```

### Just want to generate cloud-config

```
$ multipass-zig-devenv gen > cloud-config.yaml
```
