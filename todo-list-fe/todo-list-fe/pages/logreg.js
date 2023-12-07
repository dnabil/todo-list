// pages/index.js

import { Box, Button, SimpleGrid, Flex, FormControl, Heading, Input, InputGroup, InputLeftElement, InputRightElement, Stack, VStack, useColorModeValue, FormLabel, Text, Link, Icon, Container } from "@chakra-ui/react";

import styles from '../styles/splash.css';

const Index = () => {
  return (
    <Container>
    <Heading textAlign="center" fontSize="30px">
        Welcome to the Todo List App!
      </Heading>

    <Box h="100vh" p={4}>
      <SimpleGrid columns={2} spacing={10}>
        <VStack spacing={4} align="center">
          <Text fontSize="4xl">Register</Text>
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
        </VStack>
        <VStack spacing={4} align="center">
          <Text fontSize="4xl">Login</Text>
          <Link href="/login">
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
                </Link>
        </VStack>
      </SimpleGrid>
    </Box>
    </Container>
  );
};

export default Index;