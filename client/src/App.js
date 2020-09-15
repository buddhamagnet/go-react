import React, { useState, useEffect } from 'react';
import axios from 'axios';
import logo from './logo.svg';
import './App.css';

function App() {
  const [data, setData] = useState({});

  useEffect(() => {
    const fetchData = async () => {
      return await axios('http://localhost:8080/shell');
    };

    fetchData().then((result) => setData(result.data));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          {data && (
            <span>
              Hello {data.client}: {data.message}
            </span>
          )}
        </p>
        <a
          className="App-link"
          href="https://github.com/buddhamagnet/cv"
          target="_blank"
          rel="noopener noreferrer"
        >
          check me on Github
        </a>
      </header>
    </div>
  );
}

export default App;
