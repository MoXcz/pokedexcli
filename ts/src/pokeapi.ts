export class PokeAPI {
  private static readonly baseURL = "https://pokeapi.co/api/v2";

  constructor() {}

  async fetchLocations(pageURL?: string): Promise<ShallowLocations> {
    try {
      const response = await fetch(
        pageURL ?? PokeAPI.baseURL + "/location-area",
      );
      if (!response.ok) {
        throw new Error(`${response.status} ${response.statusText}`);
      }

      return await response.json();
    } catch (e) {
      throw new Error((e as Error).message);
    }
  }

  async fetchLocation(locationName: string): Promise<Location> {
    try {
      const url = PokeAPI.baseURL + `/location-area/${locationName}`;
      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`${response.status} ${response.statusText}`);
      }

      return await response.json();
    } catch (e) {
      throw new Error((e as Error).message);
    }
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
