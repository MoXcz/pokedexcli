# Pokedexcli

A [REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) interface for the [PokéAPI](https://pokeapi.co/) that lets you explore locations, catch Pokémon, and build your collection with built-in caching.

## Overview

This project contains two implementations of the same Pokedex CLI application, one in Go (`/go`) and another one in TypeScript (`/ts`).

Both versions provide:

- Explore Pokémon locations and areas
- Catch Pokémon with a chance-based mechanic
- Inspect caught Pokémon details
- View your Pokédex collection
- Location area pagination (map/mapb commands)
- Command history tracking

## Quick Start

### Go Version

See the detailed instructions in [go/README.md](go/README.md):

```sh
go install github.com/MoXcz/pokedexcli@latest
pokedexcli
```

### TypeScript Version

See the detailed instructions in [ts/README.md](ts/README.md):

```sh
cd ts
npm install
npm run build
npm start
```

## Project Structure

```
.
├── go/          # Go implementation
│   ├── main.go  # Entry point
│   ├── *.go
│   └── README.md
└── ts/          # TypeScript implementation
    ├── src/     # Source files
    └── README.md
```
