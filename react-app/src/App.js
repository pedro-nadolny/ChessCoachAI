
import React, { useState } from 'react';
import './App.css';


function App() {
  const [messages, setMessages] = useState([]);
  const [text, setText] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const addMessage = (sent, received) => {
    const date = new Date();
    const time = `${date.getHours() < 10 ? '0' + date.getHours() : date.getHours()}:${date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()}`;
    setMessages([
      ...messages, 
      { text: sent, type: 'sent', detail: `${time} You`}, 
      { text: received, type: 'received', detail: `${time} Chess Coach`}]);
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    fetch('http://localhost:3001', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ text }),
    })
    .then((res) => res.json())
    .then((data) => {
      addMessage(text, data.message);
      setText('');
    });
  };
  
  return (
    <div className="App" style={{padding: '16px'}}>
      <form onSubmit={handleSubmit}>
        <h1>Capablanca</h1>
        <input
          type="text"
          value={text}
          onChange={(e) => setText(e.target.value)}
          style={{ margin: '8px' }}
        />
        <button type="submit" disabled={isLoading}>Send</button>
      </form>
      {
        messages.map((message, i) => (
          <div className={message.type} key={i} style={{ textAlign: message.type === 'sent' ? 'right' : 'left', wordWrap: 'break-word', padding: '8px, 8px'}}>
            {message.text}
            <div style={{color: 'gray', fontSize: '0.85em'}}>{message.detail}</div>
          </div>
        ))}
    </div>
  );
}

export default App;