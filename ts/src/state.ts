import { createInterface, type Interface } from "node:readline";
import { commandExit } from "./command_exit.js";
import { commandHelp } from "./command_help.js";
import { PokeAPI } from "./pokeapi.js";
import { commandMap, commandMapb } from "./command_map.js";
import { commandExplore } from "./command_explore.js";

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
      name: "explore",
      description: "Explore pokemon in a location",
      callback: commandExplore,
    },
  };
  const pokeApi = new PokeAPI();

  const state: State = {
    rl: rl,
    commands: commands,
    pokeApi: pokeApi,
    nextLocationsURL: undefined,
    prevLocationsURL: undefined,
  };

  return state;
}
