import React, { useState, FC } from 'react';
import { HelloRequest } from './generated/hello_messages_pb';
import { HelloServicesClient } from './generated/hello_services_pb_service';
import { grpc } from '@improbable-eng/grpc-web';

const Button: FC = () => {
  const [text, setText] = useState('');

  function handleHello() {
    const client = new HelloServicesClient('http://34.85.40.127:30000', {});
    const req = new HelloRequest();

    const meta = new grpc.Metadata();
    const deadline = new Date();
    deadline.setSeconds(deadline.getSeconds() + 3);
    meta.set('deadline', deadline.getTime().toString());

    client.hello(req, meta, (err, res) => {
      if (err || res == null) {
        throw err;
      }
      setText(res.getText());
    });
  }

  return (
    <>
      <p>
        <button onClick={handleHello}>ON</button>
      </p>
      <p>
        <button onClick={() => setText('')}>OFF</button>
      </p>
      <p>{text}</p>
    </>
  );
};

export default Button;
