import { useState } from 'react';
import { Login } from './BackendConnections/Login';
import './App.css';

function App() {
  const [inputData, setInputData] = useState('');

  const sendData = () => {
    const dataCollection = new FormData();
    dataCollection.append('email', inputData);
    Login(dataCollection).then((data) => {
      console.log(data);
    });
    setInputData('');
  };
  return (
    <>
      <input value={inputData} onChange={(e) => setInputData(e.target.value)} />
      <div onClick={sendData}>Send</div>
    </>
  );
}

export default App;
