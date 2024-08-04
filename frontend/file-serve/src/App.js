import './App.css';
import axios from 'axios'
import cors from 'cors'
import { useState } from 'react';

const baseUrl = "http://localhost:3002"

function App() {

  const [imageSrc, setImageSrc] = useState(null)

  function handleDownloadImage(e) {
    const headers = {
      "Content-type": "application/json"
    }

    axios.get(baseUrl + "/files/teste.jpg", {headers, responseType: 'blob'}).then((response) => {
      console.log(response)
      if (response.data instanceof Blob) {
        const url = URL.createObjectURL(response.data);
        console.log(url)
        setImageSrc(url);
      } else {
        throw new Error('Response is not a Blob');
      }
    }).catch((error) => {
      console.error(error)
    })
  }

  return (
    <div className="App">
      <h1>Image serve test!</h1>
      {imageSrc ? <img src={imageSrc} alt='Fetched from server'/> : <p>Press button to download image</p> }
      <button onClick={(e) => handleDownloadImage(e)}>Download image!</button>
    </div>
  );
}

export default App;
