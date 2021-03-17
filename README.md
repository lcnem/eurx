# EURX

## Install

### Environment setup

This is an example for Ubuntu.

```bash
sudo apt install docker.io
sudo curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

sudo gpasswd -a $(whoami) docker
sudo chgrp docker /var/run/docker.sock
sudo systemctl enable docker
sudo systemctl restart docker
```

### Join mainnet

```bash
git clone https://github.com/lcnem/eurx.git
cd eurx
cp .env.example .env
vi .env
docker run -v ~/.eurxd:/root/.eurxd -v ~/.eurxcli:/root/.eurxcli lcnem/eurx [moniker] --chain-id eurx-1
cp launch/genesis.json ~/.eurxd/config/genesis.json
docker-compose up -d
```

Confirm:

```bash
shasum -a 256 launch/genesis.json
794d8cdcc2495274e29d5305643ed01b9de1895a9c6732287e6638a5e918b449  launch/genesis.json
```

## Deprecated way of Installation

### Environment setup

This is an example for Ubuntu.

```bash
sudo apt update
sudo apt install build-essential
cd ~
wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
echo export PATH='$PATH:/usr/local/go/bin:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Clone

```bash
git clone https://github.com/lcnem/eurx.git
cd eurx
make install
```

### Config daemon

```bash
eurxd init [moniker] --chain-id eurx-1
cp launch/genesis.json ~/.eurxd/config/genesis.json
```

### Config cli

```bash
eurxcli config chain-id eurx-1
eurxcli config trust-node true
```

### Register daemon service

```bash
vi /etc/systemd/system/eurxd.service
```

```txt
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

```bash
systemctl enable eurxd
```

### Register daemon service

```bash
vi /etc/systemd/system/eurxrest.service
```

```txt
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

```bash
systemctl enable eurxrest
```

## License

Forked from [Kava](github.com/Kava-Labs/kava).
Thanks Kava Team.

Copyright © LCNEM, Inc. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
