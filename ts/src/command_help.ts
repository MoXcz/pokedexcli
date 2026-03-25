import { State } from "./state.js";

export async function commandHelp(state: State) {
  console.log("Welcome to the Pokedex!\nUsage:\n");

  for (const c in state.commands) {
    console.log(`${state.commands[c].name}: ${state.commands[c].description}`);
  }
}
