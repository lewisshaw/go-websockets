import React, { useEffect, useState } from "react";
import { useParams } from "react-router";
import { RacesClient } from "../http/races-client";
import { Race as RaceModel } from "../models/race";
import { TimeMessage } from "../models/time-message";

export function Race() {
    const [race, setRace] = useState<RaceModel | undefined>(undefined)
    const {id} = useParams();
    const [time, setTime] = useState<Number>(0);
    const [phase, setPhase] = useState<"starting" | "started">("starting")
    const [error, setError] = useState<string | undefined>(undefined);
    const [hooting, setHooting] = useState<boolean>(false);
    const client = new RacesClient();

    const idAsNumber = parseInt(id || "0")
    useEffect(() => {
        client.viewRace(idAsNumber).then((race) => {
            setRace(race);
            const conn = new WebSocket(`ws://localhost:8080/races/${idAsNumber}/ws?id=`);
            conn.onclose = function (evt) {
                console.log("Closed");
            };
            conn.onmessage = function (evt) {
                const event = evt as MessageEvent<string>;
                const message = JSON.parse(event.data) as TimeMessage
                console.log(message.time);
                setTime(message.time);
                setPhase(message.phase);
                setHooting(message.hoot);
            };
        }).catch((error: Error) => {
            setError(error.message)
        })
    }, [id])

    function startRace() {
        client.startRace(idAsNumber)
    }

    if (error !== undefined) {
        return <p>{error}</p>
    }
    
    if (race === undefined) {
        return <p>Loading</p>
    }

    return <div>
        <h1>{race.name}</h1>
        <p>{phase === "starting" ? "Countdown" : "Elapsed"} {time}</p>
        {hooting ? <p>"HOOOOOOT"</p> : null}
        <div>
            {race.competitors.map((competitor) => {
                const finished = competitor.finishTime !== 0;
                return <p key={competitor.id}>
                    {competitor.name} {finished ? competitor.finishTime : <button>Finish</button>}
                </p>
            })}
        </div>
        <button onClick={startRace}>Start</button>
    </div>;
}