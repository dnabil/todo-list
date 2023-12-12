import React, { useState } from 'react';
import Link from 'next/link';
import axios from 'axios';
import '../styles/globalui.css';
import { useRouter } from "next/router";

const Register = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const router = useRouter();

    const handleSubmit = (e) => {
        e.preventDefault();
        const values = {
            username,
            email,
            password,
        };

        registerUser(values);
    };

    const registerUser = async (user) => {
        try {
            const res = await axios.post('http://localhost:5555/api/users/register', user);

            if (res.status !== 200) {
                alert(res.data.message);
                return;
            }

            alert(res.data.message);
            //setToken(res.data.token);
            router.push("/todo");
        } catch (error) {
            console.error("Error during register:", error);
        }
    };

    return (
        <div className="container" style={{ width: '30%' }}>
            <h1 style={{ textAlign: 'center', paddingTop: '50px' }}>Register</h1>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    id="name"
                    placeholder='Username'
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <input
                    type="email"
                    id="email"
                    placeholder='Email'
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <input
                    type="password"
                    id="password"
                    placeholder='Password'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button type="submit">Register</button>
            </form>
            <p style={{ textAlign: 'center' }}>
                Already have an account? <Link href="/login">Login</Link>
            </p>
        </div>
    );
};

export default Register;