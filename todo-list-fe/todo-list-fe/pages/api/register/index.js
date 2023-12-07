import { v4 as uuidv4 } from 'uuid';

const handler = async (req, res) => {
  if (req.method === 'POST') {
    try {
      const { name, email, password } = req.body;

      // You can use a database like MongoDB to store user data
      // For now, let's just simulate a successful registration by creating a unique user ID
      const userId = uuidv4();

      // Save userId, name, email, and password to your database here

      res.status(200).json({ success: true, message: 'User registered successfully', userId });
    } catch (err) {
      res.status(500).json({ success: false, message: 'Something went wrong', error: err.message });
    }
  } else {
    res.status(405).json({ success: false, message: 'Method not allowed' });
  }
};

export default handler;