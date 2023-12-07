import React, { useState } from 'react';
import styles from '../styles/global.css';
import { Box, Button, Flex, FormControl, Heading, Input, InputGroup, InputLeftElement, InputRightElement, Stack, VStack, useColorModeValue, FormLabel, Text, Link, Icon } from "@chakra-ui/react";

const LoginPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Email:', email);
    console.log('Password:', password);
  };

  return (
    <div>
      <Heading textAlign="center" fontSize="24px">
          Sign In to Your Account
          </Heading>
          <Box rounded="lg" bg={useColorModeValue("white", "gray.700")} p={8}>
          <form onSubmit={handleSubmit}>
            <Stack spacing={1}>
              <FormControl>
                <InputGroup>
                  <Input type="text" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
                </InputGroup>
              </FormControl>
              <FormControl>
                <InputGroup>
                  <Input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
                </InputGroup>
              </FormControl>
              <Stack spacing={5} pt={2}>
                <Button
                  type="submit"
                  value="submit"
                  size="lg"
                  bg={"blue.400"}
                  color={"white"}
                  height={"45px"}
                  _hover={{
                    bg: "blue.500",
                  }}
                >
                  Sign in
                </Button>
              </Stack>
              <Stack pt={6}>
                <Text align="center">
                  Don't have an account?{" "}
                  <Link color="blue.400" href="./register">
                    register
                  </Link>
                </Text>
              </Stack>
            </Stack>
          </form>
        </Box>
    </div>
  );
};

export default LoginPage;