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
    <main class="container">
      <section id="preview">
        <h2>Preview</h2>
        <p>
          Sed ultricies dolor non ante vulputate hendrerit. Vivamus sit amet suscipit sapien. Nulla
          iaculis eros a elit pharetra egestas.
        </p>
        <form>
          <div class="grid">
            <input
              type="text"
              name="firstname"
              placeholder="First name"
              aria-label="First name"
              required
            />
            <input
              type="email"
              name="email"
              placeholder="Email address"
              aria-label="Email address"
              required
            />
            <button type="submit">Subscribe</button>
          </div>
          <fieldset>
            <label for="terms">
              <input type="checkbox" role="switch" id="terms" name="terms" />
              I agree to the
              <a href="#" onclick="event.preventDefault()">Privacy Policy</a>
            </label>
          </fieldset>
        </form>
      </section>
      </main>
  );
};

export default LoginPage;