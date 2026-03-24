import { getCommands } from "./command.js";

export function commandHelp() {
  console.log("Welcome to the Pokedex!\nUsage:\n");
  const commands = getCommands();

  for (const c in commands) {
    console.log(`${commands[c].name}: ${commands[c].description}`);
  }
}
