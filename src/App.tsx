import React from 'react';
import './App.css';
import { Route, Routes } from 'react-router';
import Index from './pages/Index';
import { Race } from './pages/Race';

function App() {
  return (
    <div className="App">
      <h1>Race Timer</h1>
      <Routes>
        <Route path="/" element={<Index />} />
        <Route path="/races/:id" element={<Race />} />
      </Routes>
    </div>
  );
}

export default App;
