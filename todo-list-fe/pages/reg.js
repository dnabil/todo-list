import { useState } from 'react';
import Link from 'next/link';
import '../styles/globalui.css';

const Register = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log('Form Submitted: ', { name, email, password });
    };

    return (
        <div className="container" style={{ width: '30%' }}>
            <h1 style={{ textAlign: 'center', paddingTop: '50px' }}>Register</h1>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    id="name"
                    placeholder='Name'
                    value={name}
                    onChange={(e) => setName(e.target.value)}
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
            <p style={{ textAlign: 'center'}}>
                Already have an account? <Link href="/login">Login</Link>
            </p>
        </div>
    );
};

export default Register;