import React, { useState, useEffect, useRef } from 'react';

const Document = () => {
  const [sendNums] = useState([3, 4, 7, 1, 2]);
  const [receivedNums, setReceivedNums] = useState([]);
  const requestSent = useRef(false);

  useEffect(() => {
    const handleNumbersReceived = (numbers) => {
      console.log('Received numbers:', numbers);
      setReceivedNums(numbers);
    };

    const removeListener = window.api.onNumbersReceived(handleNumbersReceived);

    const sendAndReceiveNumbers = async () => {
      if (requestSent.current) return;
      requestSent.current = true;

      try {
        console.log('Sending numbers:', sendNums);
        const result = await window.api.sendRecvNumbers(sendNums);
        console.log('Stream completed:', result);
      } catch (error) {
        console.error('Error in bidirectional streaming:', error);
      }
    };

    sendAndReceiveNumbers();

    return removeListener;
  }, [sendNums]);

  return (
    <div>
      <h2>Sent Numbers:</h2>
      <p>{sendNums.join(', ')}</p>
      <h2>Received Numbers:</h2>
      <ul>
        {receivedNums.map((num, index) => (
          <li key={index}>Received: {num}</li>
        ))}
      </ul>
    </div>
  );
};

export default Document;