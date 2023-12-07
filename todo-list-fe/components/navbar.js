import Link from 'next/link';
import { useRouter } from 'next/router';
import '../styles/globalui.css';

const Navbar = () => {
    const router = useRouter();

    const handleLogout = () => {
        // Fungsi untuk mengeluarkan pengguna dari aplikasi Anda
        // Contoh: hapus token autentikasi dari penyimpanan lokal
        localStorage.removeItem('token');

        // Alihkan ke halaman login atau halaman yang ingin Anda tampilkan setelah logout
        router.push('/login');
    };

    return (
            <div className="navbar-right">
                <span className="user-name">Nama User</span>
                <button onClick={handleLogout}>Logout</button>
            </div>
        
    );
};

export default Navbar;