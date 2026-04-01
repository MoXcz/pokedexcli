import { State } from "./state.js";

export async function startREPL(state: State) {
  state.rl.prompt();
  state.rl.on("line", async (input) => {
    if (input == "") {
      state.rl.prompt();
      return;
    }
    const cleanedInput = cleanInput(input);
    const command = cleanedInput[0];
    const args = cleanedInput.slice(1);

    if (command in state.commands) {
      try {
        await state.commands[command].callback(state, ...args);
      } catch (err) {
        console.error(err);
      }
    } else {
      console.log("Unknown command");
    }
    state.rl.prompt();
  });
}

export function cleanInput(str: string): string[] {
  return str
    .trim()
    .toLowerCase()
    .split(" ")
    .filter((x) => x != "");
}
