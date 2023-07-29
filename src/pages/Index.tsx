import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { RacesClient } from '../http/races-client';
import { Race } from '../models/race';

function Index() {
    const [loading, setLoading] = useState<boolean>(true);
    const [races, setRaces] = useState<Race[]>([]);

    function updateRaces(): void {
        const client = new RacesClient();
        client.getRaces().then((races) => {
            setRaces(races);
            setLoading(false);
        });
    }

    function createRace(): void {
        const client = new RacesClient();
        client.addRace().then((race) => {
            setRaces([...races, race])
        });
    }

    useEffect(() => {
        updateRaces();        
    }, [])

    if (loading) {
        return <p>Loading...</p>
    }
    return <div>
        <p>Add Race: <button onClick={createRace}>Add Race</button></p>
        <ul>
            {
                races.map((race) => {
                    return <li key={race.id}>
                        <Link to={`/races/${race.id}`} >{race.name}</Link>
                    </li>
                })
            }
        </ul>
    </div>
}

export default Index;