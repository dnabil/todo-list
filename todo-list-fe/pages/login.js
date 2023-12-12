import React, { useState } from 'react';
import Link from 'next/link';
import '../styles/globalui.css';
//import api from '../pages/api';
import { useRouter } from "next/router";
//import backend from "../pages/api";

const LoginPage = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const router = useRouter();

    const handleSubmit = (e) => {
        e.preventDefault();
    
        if (email === "" || password === "") {
          alert("Email dan Password tidak boleh kosong");
          return;
        }
    
        const user = {
          email,
          password,
        };
    
        handleLogin(user);
      };
    
      const handleLogin = async (user) => {
        try {
          const res = await axios.post("http://localhost:5555/api/users/login", user);
    
          if (res.status !== 200) {
            alert(res.data.message);
            return;
          }
    
          //setToken(res.data.token);
          router.push("/todo");
        } catch (error) {
          console.error("Error during login:", error);
        }
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
            <p style={{ textAlign: 'center' }}>
                Don't have an account yet? <Link href="/reg">Register</Link>
            </p>
        </main>
    );
};

export default LoginPage;