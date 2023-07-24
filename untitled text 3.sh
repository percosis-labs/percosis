#! /bin/bash
set -e

validator1Mnemonic="frequent arch plastic color expose solution arrive truly impulse address treat gold muscle marble entire holiday wave shiver senior orphan pelican shift luxury cry"
#        fury1eqajecv88hlauzuwung4rc56hkgxde75geargz
# furyvaloper1eqajecv88hlauzuwung4rc56hkgxde75ynlkan

validator2Mnemonic="spy flight nuclear cram post smoke comic cliff swear either vessel regular strike orange force spin sketch hip ankle clutch focus zebra warrior below"
#        fury1pnxypvwtq2sal7yxt0zmrh3224r2syr0xzkna5
# furyvaloper1pnxypvwtq2sal7yxt0zmrh3224r2syr02g5xg9

validator3Mnemonic="cage unique solution cry wall grit payment voyage amount address category museum fly suit chief buddy river pink property run sunny breeze own napkin"
#        fury13y5p0hrtkpasypvzpcydts3pyruqm9jhe59axv
# furyvaloper13y5p0hrtkpasypvzpcydts3pyruqm9jh478gna

validator4Mnemonic="solution scan voyage exercise hip funny noble antique state episode burden pride boss unit room off fire index butter fatigue celery myself spatial egg"
#        fury1yup9sz7r62yx42csfckx0tkmayxm4mjulwqupj
# furyvaloper1yup9sz7r62yx42csfckx0tkmayxm4mjunyzf5r

validator5Mnemonic="scan radio neither girl evidence alert method sport twenty rotate citizen fringe client uphold replace correct scrap crawl amount unlock soup ahead next ankle"
#        fury1ph85un2mcs56qmftzraz9425qqkvanxpg8vp69
# furyvaloper1ph85un2mcs56qmftzraz9425qqkvanxpydw505

validator6Mnemonic="rhythm scrub angry olympic tower around clarify rice relax soon tuition toward shove utility wish grief glimpse seed educate lock vendor bracket hour maple"
#        fury15cxhnxx4plmll7acdfpsxtjrkdrct3xfp5ltg0
# furyvaloper1eqajecv88hlauzuwung4rc56hkgxde75ynlkan

validator7Mnemonic="rhythm scrub angry olympic tower around clarify rice relax soon tuition toward shove utility wish grief glimpse seed educate lock vendor bracket hour maple"
#        fury1zg2qcg446daapvz48lxpfrcsh7f33kwmdu20ye
# furyvaloper1eqajecv88hlauzuwung4rc56hkgxde75ynlkan

validator8Mnemonic="price viable rug deny merry rich damage ecology work chest require purchase arrive oak average cover million tank indicate deposit notice material brick actor"
#        fury1ujf4x2hyyqjsvew97s4vxsjnhzmvxka73jeevy
# furyvaloper1eqajecv88hlauzuwung4rc56hkgxde75ynlkan

validator9Mnemonic="ring monkey roof twelve wrist lonely actor gift turkey deposit shoulder fly swift sniff biology swear spirit second subway already shoe cricket only spring"
#        fury156nj3ck4xw3g85v8emh39082ty9g9a6wggk0zs
# furyvaloper1eqajecv88hlauzuwung4rc56hkgxde75ynlkan

validator10Mnemonic="alley trouble conduct flush pencil layer resource dance possible welcome move one barrel faint collect deputy abandon pluck patient manual entry neglect topple ostrich"
# fury1zduvuqgjhdfpgrsfjdnpqk32qcz4424rdulycm

faucetMnemonic="crash sort dwarf disease change advice attract clump avoid mobile clump right junior axis book fresh mask tube front require until face effort vault"
# fury1adkm6svtzjsxxvg7g6rshg6kj9qwej8gwqadqd

evmFaucetMnemonic="hundred flash cattle inquiry gorilla quick enact lazy galaxy apple bitter liberty print sun hurdle oak town cash because round chalk marriage response success"
# 0x3C854F92F726A7897C8B23F55B2D6E2C482EF3E0
# fury18jz5lyhhy6ncjlyty064kttw93yzaulq7rlptu

adrianMnemonic="taxi letter right spend concert brush cave album cactus chapter divert blossom select become lazy decade novel orange shove artist between despair brush sort"
# 0x416A6FE1D676ECBBBFCC42E6C80025062F6A1FF5
# fury1g94xlcwkwmkth07vgtnvsqp9qchk58l44wtxw6

jordanMnemonic="exclude name enlist alter hundred skate midnight point strike put spend purchase learn subject cute tourist annual chief outer invite art neglect catalog tower"
# 0x6578967C69F881FFD959B0B7F4FF5CF2E03545A3
# fury1v4ufvlrflzqllk2ekzmlfl6u7tsr23drwh0xhd

vivekMnemonic="taxi letter right spend concert brush cave album cactus chapter divert blossom select become lazy decade novel orange shove artist between despair brush sort"
# 0xD29F1F332209267D184B46F50AB3A404C61D5EED
# fury162037vezpyn86xztgm6s4vayqnrp6hhd3gpv06

relayerMnemonic="visual predict emerge alone clean dirt demand sing offer output zone morning ordinary similar palace spawn during future cruel define female stool primary alone"
# 0xa2F728F997f62F47D4262a70947F6c36885dF9fa
# fury15tmj37vh7ch504px9fcfglmvx6y9m70646ev8t

DATA=~/.fury
# remove any old state and config
rm -rf $DATA

BINARY=fury

# Create new data directory, overwriting any that alread existed
chainID="highbury_710-1"
$BINARY init validator --chain-id $chainID 

# hacky enable of rest api
sed -in-place='' 's/enable = false/enable = true/g' $DATA/config/app.toml

# Set evm tracer to json
sed -in-place='' 's/tracer = ""/tracer = "json"/g' $DATA/config/app.toml

# Enable full error trace to be returned on tx failure
sed -in-place='' '/iavl-cache-size/a\
trace = true' $DATA/config/app.toml

# Set client chain id
sed -in-place='' 's/chain-id = ""/chain-id = "clockend_4200-1"/g' $DATA/config/client.toml

# avoid having to use password for keys
$BINARY config keyring-backend test

# Create validator keys and add account to genesis
validatorKeyName="argentina"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

validatorKeyName="arsenal"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

validatorKeyName="brazil"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

validatorKeyName="barcelona"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

validatorKeyName="brooklyn"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

validatorKeyName="india"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

validatorKeyName="giants"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

validatorKeyName="lakers"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

validatorKeyName="yankees"
$BINARY keys add $validatorKeyName 
$BINARY add-genesis-account $validatorKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

evmFaucetKeyName="evm_faucet"
$BINARY keys add $evmFaucetKeyName 
$BINARY add-genesis-account $evmFaucetKeyName 20000000000ufury,1000000000bnb,20000000000000000000000afury

# Create a delegation tx for the validator and add to genesis
$BINARY gentx $validatorKeyName 1000000000ufury --keyring-backend test --chain-id $chainID
$BINARY collect-gentxs

# Replace stake with ufury
sed -in-place='' 's/stake/ufury/g' $DATA/config/genesis.json

# Replace the default evm denom of aphoton with ufury
sed -in-place='' 's/aphoton/afury/g' $DATA/config/genesis.json

# Zero out the total supply so it gets recalculated during InitGenesis
jq '.app_state.bank.supply = []' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

# Disable fee market
jq '.app_state.feemarket.params.no_base_fee = true' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

# Disable london fork
jq '.app_state.evm.params.chain_config.london_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
jq '.app_state.evm.params.chain_config.arrow_glacier_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
jq '.app_state.evm.params.chain_config.gray_glacier_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
jq '.app_state.evm.params.chain_config.merge_netsplit_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
jq '.app_state.evm.params.chain_config.shanghai_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
jq '.app_state.evm.params.chain_config.cancun_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

# Add earn vault
jq '.app_state.earn.params.allowed_vaults =  [
    {
        denom: "usdf",
        strategies: ["STRATEGY_TYPE_JINX"],
    },
    {
        denom: "bfury",
        strategies: ["STRATEGY_TYPE_SAVINGS"],
    }]' $DATA/config/genesis.json | sponge $DATA/config/genesis.json

jq '.app_state.savings.params.supported_denoms = ["bfury-furyvaloper1ffv7nhd3z6sych2qpqkk03ec6hzkmufyz4scd0"]' $DATA/config/genesis.json | sponge $DATA/config/genesis.json


$BINARY config broadcast-mode block