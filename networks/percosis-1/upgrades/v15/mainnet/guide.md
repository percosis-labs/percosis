# v14 to v15 Mainnet Upgrade Guide

Percosis v15 Gov Prop: <https://www.mintscan.io/percosis/proposals/458>

Countdown: <https://www.mintscan.io/percosis/blocks/8732500>

Height: 8732500

## Memory Requirements

This upgrade will **not** be resource intensive. With that being said, we still recommend having 64GB of memory. If having 64GB of physical memory is not possible, the next best thing is to set up swap.

Short version swap setup instructions:

``` {.sh}
sudo swapoff -a
sudo fallocate -l 32G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

To persist swap after restart:

``` {.sh}
sudo cp /etc/fstab /etc/fstab.bak
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

In depth swap setup instructions:
<https://www.digitalocean.com/community/tutorials/how-to-add-swap-space-on-ubuntu-20-04>

## First Time Cpercovisor Setup

If you have never setup Cpercovisor before, follow the following instructions.

If you have already setup Cpercovisor, skip to the next section.

We highly recommend validators use cpercovisor to run their nodes. This
will make low-downtime upgrades smoother, as validators don't have to
manually upgrade binaries during the upgrade, and instead can
pre-install new binaries, and cpercovisor will automatically update them
based on on-chain SoftwareUpgrade proposals.

You should review the docs for cpercovisor located here:
<https://docs.cosmos.network/main/tooling/cpercovisor>

If you choose to use cpercovisor, please continue with these
instructions:

To install Cpercovisor:

``` {.sh}
go install github.com/cosmos/cosmos-sdk/cpercovisor/cmd/cpercovisor@v1.0.0
```

After this, you must make the necessary folders for cosmosvisor in your
daemon home directory (\~/.percosisd).

``` {.sh}
mkdir -p ~/.percosisd
mkdir -p ~/.percosisd/cpercovisor
mkdir -p ~/.percosisd/cpercovisor/genesis
mkdir -p ~/.percosisd/cpercovisor/genesis/bin
mkdir -p ~/.percosisd/cpercovisor/upgrades
```

Copy the current v14 percosisd binary into the
cpercovisor/genesis folder and v14 folder.

```{.sh}
cp $GOPATH/bin/percosisd ~/.percosisd/cpercovisor/genesis/bin
mkdir -p ~/.percosisd/cpercovisor/upgrades/v14/bin
cp $GOPATH/bin/percosisd ~/.percosisd/cpercovisor/upgrades/v14/bin
```

Cpercovisor is now ready to be set up for v15.

Set these environment variables:

```{.sh}
echo "# Setup Cpercovisor" >> ~/.profile
echo "export DAEMON_NAME=percosisd" >> ~/.profile
echo "export DAEMON_HOME=$HOME/.percosisd" >> ~/.profile
echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
echo "export UNSAFE_SKIP_BACKUP=true" >> ~/.profile
source ~/.profile
```

## Cpercovisor Upgrade

Create the v15 folder, make the build, and copy the daemon over to that folder

```{.sh}
mkdir -p ~/.percosisd/cpercovisor/upgrades/v15/bin
cd $HOME/percosis
git pull
git checkout v15.0.0
make build
cp build/percosisd ~/.percosisd/cpercovisor/upgrades/v15/bin
```

Now, at the upgrade height, Cpercovisor will upgrade to the v15 binary

## Manual Option

1. Wait for Percosis to reach the upgrade height 8732500

2. Look for a panic message, followed by endless peer logs. Stop the daemon

3. Run the following commands:

```{.sh}
cd $HOME/percosis
git pull
git checkout v15.0.0
make install
```

4. Start the percosis daemon again, watch the upgrade happen, and then continue to hit blocks

## Further Help

If you need more help, please go to <https://docs.percosis.zone> or join
our discord at <https://discord.gg/pAxjcFnAFH>.