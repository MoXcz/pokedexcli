import { State } from "./state";

export function commandExit(_: State) {
  console.log("Closing the Pokedex... Goodbye!");
  process.exit(0);
}
