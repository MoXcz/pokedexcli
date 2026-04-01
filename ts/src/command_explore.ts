import { State } from "./state.js";

export async function commandExplore(state: State, ...args: string[]) {
  if (args.length != 1) {
    console.log("Invalid argument");
    return;
  }
  const location = args[0];

  const locationResponse = await state.pokeApi.fetchLocation(location);
  const pokemon_encounters = locationResponse.pokemon_encounters;

  console.log(`Exploring ${location}...`);
  console.log("found Pokemon:");
  pokemon_encounters.forEach((enc) => {
    console.log(`  - ${enc.pokemon.name}`);
  });
}
