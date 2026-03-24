# Pokedexcli

This is `pokedexcli`, a [REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) (Read-Eval-Print Loop) interface for the [PokéAPI](https://pokeapi.co/) that integrates caching for ease of use.

## Motivation

There are a lot of ways to use APIs and to look for Pokémon, so `pokedexcli` uses an API of Pokémon with built-in cache in order to add interactivity to the arduous task of learning the abilities of a Magikarp.

## Quick Start

Install `pokedexcli` using the Go toolchain (information on how to install it [here](https://go.dev/doc/install)):

```sh
go install github.com/MoXcz/pokedexcli@latest
```

Run `pokedexcli` (make sure `$GOPATH` is set correctly and is part of your `$PATH`. More information [here](https://go.dev/wiki/SettingGOPATH)):

```sh
pokedexcli
Pokedex > help
```

## Usage

Available commands:
- `help`. Display a help message with the usage of the commands.
- `exit`. Close `pokedexcli` gracefully.
- `map`. Display 20 Pokémon location areas (the order is as returned by PokéAPI).
- `mapb`. Display previous locations.
- `explore`. Display Pokémon available in a particular location.
- `catch`. Simulate capturing a Pokémon by using a built-in chance based on the base experience of the Pokémon.
- `inspect`. Display information about a caught Pokémon
- `pokedex`. Print caught Pokémon
- `history`. Print previously used commands

## Examples
```sh
pokedexcli
Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught
You may now inspect it with the inspect command
Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  - hp: 40
  - attack: 45
  - defense: 40
  - special-attack: 35
  - special-defense: 35
  - speed: 56
Types:
  - normal
  - flying
```

## Build

### Clone the repo:

```sh
git clone https://github.com/MoXcz/pokedexcli # clone the repository
cd pokedexcli
```

### Build the project:

```sh
go build .
./pokedexcli
```

### Run the tests:

```sh
go test ./...
```
