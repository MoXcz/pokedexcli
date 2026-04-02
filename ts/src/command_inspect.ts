import { State } from "./state.js";

export async function commandInspect(state: State, ...args: string[]) {
  if (args.length != 1) {
    console.log("Invalid argument");
    return;
  }
  const name = args[0];

  const pokemon = state.pokedex[name];

  if (pokemon === undefined) {
    console.log(`you do not have a ${name} in your pokedex`);
    return;
  }

  console.log(`Name: ${pokemon.name}`);
  console.log(`Height: ${pokemon.height}`);
  console.log(`Weight: ${pokemon.weight}`);
  console.log("Stats:");
  pokemon.stats.forEach((stat) => {
    console.log(`  -${stat.stat.name}: ${stat.base_stat}`);
  });
  console.log("Types:");
  pokemon.types.forEach((stat) => {
    console.log(`  - ${stat.type.name}`);
  });
}
