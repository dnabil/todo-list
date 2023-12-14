import { useContext, useState, useEffect } from "react";
import Link from 'next/link';
import axios from 'axios';
import '../styles/globalui.css';
import { useRouter } from "next/router";
import { AuthContext } from "../utils/AuthContext";

const LoginPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();
  const { setToken } = useContext(AuthContext);
  const [error, setError] = useState('');

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

      const token = res.data.token;
      // Set token to localStorage or cookie
      localStorage.setItem('token', token);
      //setToken(res.data.token);
      router.push('/todo');
    } catch (error) {
      setError(error.res.data.message);
    }
  };

  useEffect(() => {
    const token = localStorage.getItem('token');

    if (token) {
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
    }
  }, []);

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