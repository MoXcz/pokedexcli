import { Cache } from "./pokecache.js";

export class PokeAPI {
  private static readonly baseURL = "https://pokeapi.co/api/v2";
  private cache: Cache = new Cache(30);

  constructor() {}

  async fetchLocations(pageURL?: string): Promise<ShallowLocations> {
    try {
      const url = pageURL ?? PokeAPI.baseURL + "/location-area";

      const cached = this.cache.get<ShallowLocations>(url);
      if (cached) {
        return cached;
      }

      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`${response.status} ${response.statusText}`);
      }

      const data: ShallowLocations = await response.json();
      this.cache.add(url, data);

      return data;
    } catch (e) {
      throw new Error((e as Error).message);
    }
  }

  async fetchLocation(locationName: string): Promise<Location> {
    const url = PokeAPI.baseURL + `/location-area/${locationName}`;

    const cached = this.cache.get<Location>(url);
    if (cached) {
      return cached;
    }

    try {
      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`${response.status} ${response.statusText}`);
      }

      const data: Location = await response.json();
      this.cache.add(url, data);

      return data;
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
