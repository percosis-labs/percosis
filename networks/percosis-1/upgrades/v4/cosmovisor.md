# Install and setup Cpercovisor

We highly recommend validators use cpercovisor to run their nodes. This
will make low-downtime upgrades smoother, as validators don't have to
manually upgrade binaries during the upgrade, and instead can preinstall
new binaries, and cpercovisor will automatically update them based on
on-chain SoftwareUpgrade proposals.

You should review the docs for cpercovisor located here:
<https://docs.cosmos.network/main/tooling/cpercovisor>

If you choose to use cpercovisor, please continue with these
instructions:

To install Cpercovisor:

    git clone https://github.com/cosmos/cosmos-sdk
    cd cosmos-sdk
    git checkout v0.42.9
    make cpercovisor
    cp cpercovisor/cpercovisor $GOPATH/bin/cpercovisor
    cd $HOME

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
    source ~/.profile

Finally, you should copy the current percosisd binary into the
cpercovisor/genesis folder.

    cp $GOPATH/bin/percosisd ~/.percosisd/cpercovisor/genesis/bin

Prepare for upgrade (v4)
------------------------

To prepare for the upgrade, you need to create some folders, and build
and install the new binary.

    mkdir -p ~/.percosisd/cpercovisor/upgrades/v4/bin
    git clone https://github.com/percosis-labs/percosis
    cd percosis
    git checkout v4.0.0
    make build
    cp build/percosisd ~/.percosisd/cpercovisor/upgrades/v4/bin

Now cpercovisor will run with the current binary, and will automatically
upgrade to this new binary at the appropriate height if run with:

    cpercovisor start

Please note, this does not automatically update your
`$GOPATH/bin/percosisd` binary, to do that after the upgrade, please run
`make install` in the percosis source folder.
