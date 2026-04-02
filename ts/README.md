# Pokedex TypeScript

The TypeScript implementation of the Pokedex API as a CLI.

## Prerequisites

- Node.js
- npm

```sh
npm install    # install dependencies
npm run build  # build the CLI
npm start      # start the CLI
npm run dev    # run CLI
npm test       # run tests
```

## Usage

After running `npm start`, use these commands:

- `help` - Display available commands
- `exit` - Close the application
- `map` - Display 20 Pokémon location areas
- `mapb` - Display previous locations
- `explore <location>` - Display Pokémon in a location
- `catch <pokemon>` - Attempt to catch a Pokémon
- `inspect <pokemon>` - View details of a caught Pokémon
- `pokedex` - View your caught Pokémon
- `history` - View command history

## Example

```
$ npm start
Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught
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
Pokedex >
```
