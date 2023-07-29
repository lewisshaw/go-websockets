export interface TimeMessage {
    phase: "starting" | "started",
    time: number,
    hoot: boolean,
}