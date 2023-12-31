# v6 to v7 Upgrade Guide

## Memory Requirements

For this upgrade, nodes will need a total of 64GB of memory. This must
consist of a **minimum** of 32GB of RAM, while the remaining 32GB can be
swap. For best results, use 64GB of physical memory.

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

## Install and setup Cpercovisor

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

Cpercovisor requires some ENVIRONMENT VARIABLES be set in order to
function properly. We recommend setting these in your `.profile` so it
is automatically set in every session.

For validators we recommmend setting

- `DAEMON_ALLOW_DOWNLOAD_BINARIES=false` for security reasons
- `DAEMON_LOG_BUFFER_SIZE=512` to avoid a bug with extra long log
    lines crashing the server.
- `DAEMON_RESTART_AFTER_UPGRADE=true` for unattended upgrades

```{=html}
<!-- -->
```

    echo "# Setup Cpercovisor" >> ~/.profile
    echo "export DAEMON_NAME=percosisd" >> ~/.profile
    echo "export DAEMON_HOME=$HOME/.percosisd" >> ~/.profile
    echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
    echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
    echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
    echo "export UNSAFE_SKIP_BACKUP=true" >> ~/.profile
    source ~/.profile

You may leave out `UNSAFE_SKIP_BACKUP=true`, however the backup takes a
decent amount of time and public snapshots of old states are available.

Finally, you should copy the current percosisd binary into the
cpercovisor/genesis folder.

    cp $GOPATH/bin/percosisd ~/.percosisd/cpercovisor/genesis/bin

## Prepare for upgrade (v7)

To prepare for the upgrade, you need to create some folders, and build
and install the new binary.

    mkdir -p ~/.percosisd/cpercovisor/upgrades/v7/bin
    cd $HOME/percosis
    git pull
    git checkout v7.0.2
    make build
    cp build/percosisd ~/.percosisd/cpercovisor/upgrades/v7/bin

Now cpercovisor will run with the current binary, and will automatically
upgrade to this new binary at the appropriate height if run with:

    cpercovisor start

Please note, this does not automatically update your
`$GOPATH/bin/percosisd` binary, to do that after the upgrade, please run
`make install` in the percosis source folder.

## Further Help

If you need more help, please go to <https://docs.percosis.zone> or join
our discord at <https://discord.gg/pAxjcFnAFH>.
