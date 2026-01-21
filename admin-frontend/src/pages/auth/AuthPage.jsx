import { useState } from "react";

import LoginForm from "../../components/LoginForm";
import Register from "./Register";
import { motion } from "framer-motion";



const AuthPage = () => {
  const [isRegistering, setIsRegistering] = useState(false);

  const toggleForm = () => {
    setIsRegistering(!isRegistering);
  };

  return (
    <div className="relative w-full h-screen flex justify-center items-center overflow-hidden">

      {/* Background Image */}
      <div
        className="absolute inset-0 bg-cover bg-center"
        style={{ backgroundImage: "url('https://source.unsplash.com/1600x900/?technology,office')" }}
      ></div>

      {/* Overlay for Blur Effect */}
      <div className="absolute inset-0 bg-black bg-opacity-50 backdrop-blur-md"></div>

      {/* Auth Box */}
      <motion.div
        initial={{ opacity: 0, scale: 0.8 }}
        animate={{ opacity: 1, scale: 1 }}
        transition={{ duration: 0.5, ease: "easeOut" }}
        className="relative bg-white bg-opacity-30 backdrop-blur-lg p-8 rounded-2xl shadow-2xl w-96"
      >
        {isRegistering ? <Register /> : <LoginForm />}

        <div className="text-center mt-4">
          <button
            onClick={toggleForm}
            className="text-sm text-blue-600 hover:underline focus:outline-none transition duration-300"
          >
            {isRegistering ? "Already have an account? Login" : "Don't have an account? Register"}
          </button>
        </div>
      </motion.div>
    </div>
  );
};

export default AuthPage;
