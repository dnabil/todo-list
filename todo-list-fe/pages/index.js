import React, { useState } from 'react';
import Link from 'next/link';
import axios from 'axios';
import '../styles/globalui.css';
import { useRouter } from "next/router";

const SigninSignup = () => {
    const router = useRouter();

    // const handleRegister = async () => {
    //     try {
    //         const res = await axios.post('http://localhost:5555/api/users/register');
    //         if (res.status === 200) {
    //             router.push('/reg');
    //         }
    //     } catch (error) {
    //         console.error(error);
    //     }
    // };

    // const handleLogin = async () => {
    //     try {
    //         const res = await axios.post('http://localhost:5555/api/users/login');
    //         if (res.status === 200) {
    //             router.push('/login');
    //         }
    //     } catch (error) {
    //         console.error(error);
    //     }
    // };

    const handleRegister = () => {
        router.push('/reg');
    };
    
    const handleLogin = () => {
        router.push('/login');
    };
    
    return (
        <main class="container">
            <h1 style={{ textAlign: 'center', paddingTop: '50px' }}>Sign in or Sign up</h1>
            <div class="grid">
                <button onClick={handleRegister}>Register</button>
                <button onClick={handleLogin}>Login</button>
            </div>
        </main>
    );
};

export default SigninSignup;