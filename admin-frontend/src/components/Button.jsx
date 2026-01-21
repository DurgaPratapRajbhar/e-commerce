const Button = ({ 
  type = "button", 
  onClick, 
  children, 
  className = "", 
  bgColor = "bg-green-600", 
  hoverColor = "hover:bg-green-700", 
  focusRing = "focus:ring-green-300",
  isLoading = false  
}) => {
  return (
    <button
      disabled={isLoading}  
      type={type}
      onClick={onClick}
      className={`px-6 py-3 rounded-xl shadow-lg text-lg font-semibold tracking-wide transition-all duration-300
        ${bgColor} text-white ${hoverColor} active:scale-95 focus:ring-4 ${focusRing} ${className} 
        ${isLoading ? "opacity-50 cursor-not-allowed" : ""}`}  
    >
      {isLoading ? "⏳ Processing..." : children}
    </button>
  );
};

export default Button;
