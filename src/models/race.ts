import { Competitor } from "./competitor";

export interface Race {
    id: number,
    name: string,
    competitors: Competitor[]
}