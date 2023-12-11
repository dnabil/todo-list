import { useRouter } from 'next/router';
import '../styles/globalui.css';

const SigninSignup = () => {
    const router = useRouter();

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