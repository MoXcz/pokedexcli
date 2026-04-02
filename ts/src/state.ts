import { createInterface, type Interface } from "node:readline";
import { commandExit } from "./command_exit.js";
import { commandHelp } from "./command_help.js";
import { PokeAPI, Pokemon } from "./pokeapi.js";
import { commandMap, commandMapb } from "./command_map.js";
import { commandExplore } from "./command_explore.js";
import { commandCatch } from "./command_catch.js";
import { commandInspect } from "./command_inspect.js";

export type CLICommand = {
  name: string;
  description: string;
  callback: (state: State, ...args: string[]) => Promise<void>;
};

export type State = {
  rl: Interface;
  commands: Record<string, CLICommand>;
  pokeApi: PokeAPI;
  nextLocationsURL: string | undefined;
  prevLocationsURL: string | undefined;
  pokedex: Record<string, Pokemon>;
};

export function initState(): State {
  const rl = createInterface({
    input: process.stdin,
    output: process.stdout,
    prompt: "> ",
  });

  const commands: Record<string, CLICommand> = {
    exit: {
      name: "exit",
      description: "Exits the Pokedex",
      callback: commandExit,
    },
    help: {
      name: "help",
      description: "Displays a help message",
      callback: commandHelp,
    },
    map: {
      name: "map",
      description: "Call map locations",
      callback: commandMap,
    },
    mapb: {
      name: "mapb",
      description: "Call previous map locations",
      callback: commandMapb,
    },
    explore: {
      name: "explore <location_name>",
      description: "Explore pokemon in a location",
      callback: commandExplore,
    },
    catch: {
      name: "catch <pokemon",
      description: "Try to catch a pokemon",
      callback: commandCatch,
    },
  };
  const pokeApi = new PokeAPI();
  const pokedex: Record<string, Pokemon> = {};

  const state: State = {
    rl: rl,
    commands: commands,
    pokeApi: pokeApi,
    nextLocationsURL: undefined,
    prevLocationsURL: undefined,
    pokedex: pokedex,
  };

  return state;
}
