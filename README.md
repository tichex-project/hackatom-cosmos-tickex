# Cosmos Hub

![banner](./docs/welcome-banner.jpg)

[![CircleCI](https://circleci.com/gh/tichex-project/hackatom-cosmos-tichex/tree/master.svg?style=shield)](https://circleci.com/gh/tichex-project/hackatom-cosmos-tichex/tree/master)
[![codecov](https://codecov.io/gh/tichex-project/hackatom-cosmos-tichex/branch/master/graph/badge.svg)](https://codecov.io/gh/tichex-project/hackatom-cosmos-tichex)
[![Go Report Card](https://goreportcard.com/badge/github.com/tichex-project/hackatom-cosmos-tichex)](https://goreportcard.com/report/github.com/tichex-project/hackatom-cosmos-tichex)
[![license](https://img.shields.io/github/license/tichex-project/hackatom-cosmos-tichex.svg)](https://github.com/tichex-project/hackatom-cosmos-tichex/blob/master/LICENSE)
[![LoC](https://tokei.rs/b1/github/tichex-project/hackatom-cosmos-tichex)](https://github.com/tichex-project/hackatom-cosmos-tichex)
[![GolangCI](https://golangci.com/badges/github.com/tichex-project/hackatom-cosmos-tichex.svg)](https://golangci.com/r/github.com/tichex-project/hackatom-cosmos-tichex)
[![riot.im](https://img.shields.io/badge/riot.im-JOIN%20CHAT-green.svg)](https://riot.im/app/#/room/#cosmos-sdk:matrix.org)

This repository hosts `Gaia`, the first implementation of the Cosmos Hub.

**Note**: Requires [Go 1.12+](https://golang.org/dl/)

**DISCLAIMER**: The current version of Gaia running the Cosmos Hub (v0.34.x) is
__NOT__ maintained in this repository. Gaia and the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/)
have been recently split. All future versions of Gaia, including the next major
upgrade, will be maintained in this repository. However, until the next major upgrade,
Gaia should be fetched and built from the latest [released](https://github.com/cosmos/cosmos-sdk/releases)
__v0.34.x__ version in the SDK repository. In addition, this repository should be
considered unstable until the next major release of Gaia. Please bear with us
while we continue the migration process and update documentation.

## Cosmos Hub Mainnet

To run a full-node for the mainnet of the Cosmos Hub, first [install `gaia`](./docs/installation.md), then follow [the guide](./docs/join-mainnet.md).

For status updates and genesis file, see the [launch repo](https://github.com/cosmos/launch).

## Quick Start

```
make install
```

## Disambiguation

This Cosmos-SDK project is not related to the [React-Cosmos](https://github.com/react-cosmos/react-cosmos) project (yet). Many thanks to Evan Coury and Ovidiu (@skidding) for this Github organization name. As per our agreement, this disambiguation notice will stay here.


