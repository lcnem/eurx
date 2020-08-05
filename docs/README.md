# EURX documentation

## Environment

This is an example for Ubuntu.

```shell
apt update
apt install build-essential
cd ~
wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz
echo export PATH="$PATH:/usr/local/go/bin:$HOME/go/bin" >> ~/.bashrc
source ~/.bashrc
```

## Install

```shell
mkdir -p /usr/local/src/github.com/lcnem
cd /usr/local/src/github.com/lcnem
git clone https://github.com/lcnem/eurx.git
cd eurx
git checkout v0.1.0
make install
```

## Setup genesis.json

```shell
eurxd init [moniker] --chain-id eurx-1
cd /usr/local/src/github.com/lcnem/eurx
cp launch/genesis.json ~/.eurxd/config/genesis.json
```

## Setup services

```shell
eurxcli config chain-id eurx-1
eurxcli config trust-node true
```

### Daemon service

```shell
vi /etc/systemd/system/eurxd.service
```

```toml
[Unit]
Description=EURX Node
After=network-online.target

[Service]
User=root
ExecStart=/root/go/bin/eurxd start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

```shell
systemctl enable eurxd
```

### REST service

```shell
vi /etc/systemd/system/eurxrest.service
```

```toml
[Unit]
Description=EURX Rest
After=network-online.target

[Service]
User=root
ExecStart=/root/go/bin/eurxcli rest-server
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

```shell
systemctl enable eurxrest
```
