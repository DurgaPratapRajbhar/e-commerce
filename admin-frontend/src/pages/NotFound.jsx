import { Link, Navigate } from "react-router-dom";
import { useSelector } from "react-redux";

const NotFound = () => {
  const isAuthenticated = useSelector((state) => state.auth.token);
  
  if (isAuthenticated) {
    return <Navigate to="/dashboard" />;
  }

  return <Navigate to="/login" />;
};

export default NotFound;
