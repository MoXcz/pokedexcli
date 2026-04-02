import { State } from "./state.js";

export async function commandCatch(state: State, ...args: string[]) {
  if (args.length != 1) {
    console.log("Invalid argument");
    return;
  }
  const pokemon = args[0];

  const pokemonResponse = await state.pokeApi.fetchPokemon(pokemon);

  console.log(`Throwing a Pokeball at ${pokemon}...`);

  if (Math.random() > 0.5) {
    state.pokedex[pokemon] = pokemonResponse;
    console.log(`${pokemon} was caught!`);
    return;
  }
  console.log(`${pokemon} escaped!`);
}
