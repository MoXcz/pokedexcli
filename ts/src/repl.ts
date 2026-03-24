import { createInterface } from "node:readline";
import { getCommands } from "./command.js";

const rl = createInterface({
  input: process.stdin,
  output: process.stdout,
  prompt: "> ",
});

export function startREPL() {
  rl.prompt();
  rl.on("line", (input) => {
    if (input == "") {
      rl.prompt();
      return;
    }
    const cleanedInput = cleanInput(input);
    const command = cleanedInput[0];
    // const args = cleanedInput.slice(1)

    const commands = getCommands();
    if (command in commands) {
      try {
        commands[command].callback(commands);
      } catch (err) {
        console.error(err);
      }
    } else {
      console.log("Unknown command");
    }
    rl.prompt();
  });
}

export function cleanInput(str: string): string[] {
  return str
    .trim()
    .toLowerCase()
    .split(" ")
    .filter((x) => x != "");
}
