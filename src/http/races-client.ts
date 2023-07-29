import { Race } from "../models/race";
import { baseApiUrl } from "./constants"

export class RacesClient {
    async getRaces(): Promise<Race[]> {
        const response = await fetch(`${baseApiUrl}/races/view`)
        if (response.status !== 200) {
            throw new Error("Could not get races")
        }
        const body = await response.json()
        return body;
    }

    async addRace(): Promise<Race> {
        const response = await fetch(`${baseApiUrl}/races/create`)
        if (response.status !== 200) {
            throw new Error("Could not create race")
        }
        const body = await response.json()
        return body;
    }

    async viewRace(id: number): Promise<Race> {
        const response = await fetch(`${baseApiUrl}/races/${id}`);
        if (response.status === 404) {
            throw new Error("Could not find race for ID " + id)
        }
        if (response.status !== 200) {
            throw new Error("An error occurred while loading the race");
        }
        const body = await response.json();
        return body;
    }
    async startRace(id: number) {
        await fetch(`${baseApiUrl}/races/${id}/start`);
    }
}