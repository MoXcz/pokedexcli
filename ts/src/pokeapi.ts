export class PokeAPI {
  private static readonly baseURL = "https://pokeapi.co/api/v2";

  constructor() {}

  async fetchLocations(pageURL?: string): Promise<ShallowLocations> {
    const response = await fetch(pageURL ?? PokeAPI.baseURL + "/location-area");
    return await response.json();
  }

  async fetchLocation(locationName: string): Promise<Location> {
    const url = PokeAPI.baseURL + `/location-area/${locationName}`;
    const response = await fetch(url);
    return await response.json();
  }
}

export type ShallowLocations = {
  count: number;
  next: string;
  previous: any;
  results: Location[];
};

export type Location = {
  name: string;
  url: string;
};
