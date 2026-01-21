const Label = ({ className = "", children }) => {
    return (
      <label className={`block text-lg text-gray-700 font-medium ${className}`}>
        {children}
      </label>
    );
  };
  
  export default Label;
  