import React, { useState } from 'react';
import '../styles/globalui.css';

const LoginPage = () => {
 const [email, setEmail] = useState('');
 const [password, setPassword] = useState('');

 const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Email:', email);
    console.log('Password:', password);
 };

 return (
    <main className="container" style={{ width: '30%' }}>
        <h1 style={{ textAlign: 'center', paddingTop: '50px' }}>Login</h1>
        <form onSubmit={handleSubmit}>
            <input
              type="email"
              id="email"
              placeholder="Email Address"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
            <input
              type="password"
              id="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
            <button type="submit">Login</button>
        </form>
      </main>
 );
};

export default LoginPage;