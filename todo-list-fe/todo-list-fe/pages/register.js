import { useState } from 'react';
import styles from '../styles/global.css'; 

import { Box, Button, Flex, FormControl, Heading, Input, InputGroup, InputLeftElement, InputRightElement, Stack, VStack, useColorModeValue, FormLabel, Text, Link, Icon } from "@chakra-ui/react";


const Register = () => {
  const [formData, setFormData] = useState({ name: '', email: '', password: '' });
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(null);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setSuccess(false);
    setError(null);

    try {
      const response = await fetch('http://localhost:3000/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Something went wrong');
      }

      const data = await response.json();
      setSuccess(true);
      setFormData({ name: '', email: '', password: '' });
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <Heading textAlign="center" fontSize="24px">
          Register Your Self
          </Heading>
          <Stack spacing={4} pt={2}>
      <form onSubmit={handleSubmit}>
        <input type="text" placeholder="Name" id="name" name="name" value={formData.name} onChange={handleChange} required />

        <input type="email" placeholder="Email" id="email" name="email" value={formData.email} onChange={handleChange} required />

        <input type="password" placeholder="Password" id="password" name="password" value={formData.password} onChange={handleChange} required />

        <Button
                  type="submit"
                  value="submit"
                  size="lg"
                  bg={"blue.400"}
                  color={"white"}
                  _hover={{
                    bg: "blue.500",
                  }}
                >
                  Register
                </Button>
      </form>
      </Stack>
      {success && <p>Registration successful!</p>}
      {error && <p>Error: {error}</p>}
    </div>
  );
};

export default Register;