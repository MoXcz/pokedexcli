import { State } from "./state.js";

export async function commandMap(state: State) {
  const locationsResponse = await state.pokeApi.fetchLocations(
    state.nextLocationsURL,
  );

  state.prevLocationsURL = locationsResponse.previous;
  state.nextLocationsURL = locationsResponse.next;

  const locations = locationsResponse.results;

  locations.forEach((loc) => {
    console.log(loc.name);
  });
}

export async function commandMapb(state: State) {
  if (state.prevLocationsURL === undefined) {
    console.log("You're on the first page");
    return;
  }

  const locationsResponse = await state.pokeApi.fetchLocations(
    state.prevLocationsURL,
  );

  if (locationsResponse.previous === null) {
    state.prevLocationsURL = undefined;
  } else {
    state.prevLocationsURL = locationsResponse.previous;
  }

  state.nextLocationsURL = locationsResponse.next;

  const locations = locationsResponse.results;

  locations.forEach((loc) => {
    console.log(loc.name);
  });
}
