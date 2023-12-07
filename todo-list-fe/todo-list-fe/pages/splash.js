import Link from 'next/link';
import styles from '../styles/splash.css'; 

const Index = () => {
    return (
        <div>
            <h2>Welcome to the Todo List App!</h2>
            <h3>Sign Up</h3>
            <Link href="/signup">
                <button>Sign Up</button>
            </Link>
            <h3>Sign In</h3>
            <Link href="/signin">
                <button>Sign In</button>
            </Link>
        </div>
    );
};

export default Index;