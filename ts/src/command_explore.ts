import { State } from "./state.js";

export async function commandExplore(state: State, location: string) {
  const locationResponse = await state.pokeApi.fetchLocation(location);
  const pokemon_encounters = locationResponse.pokemon_encounters;

  console.log("Exploring pastoria-city-area...");
  console.log("found Pokemon:");
  pokemon_encounters.forEach((enc) => {
    console.log(`  - ${enc.pokemon.name}`);
  });
}
