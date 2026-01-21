import { PencilSquareIcon } from "@heroicons/react/24/solid";

const Input = ({ 
  type = "text", 
  name, 
  value, 
  className = "",  
  onChange,
  defaultValue,
  min,
  max,
  maxLength
}) => {
  return (
    <div className="relative">
      <input
        type={type}
        name={name}
        value={value}
        onChange={(e) => {
          let newValue = e.target.value;

          if (type === "number") {
            // Allow only numbers and decimal point
            newValue = newValue.replace(/[^0-9.]/g, ""); 

            // Enforce min and max values
            if (min !== undefined && newValue && parseFloat(newValue) < min) {
              newValue = min.toString();
            }
            if (max !== undefined && newValue && parseFloat(newValue) > max) {
              newValue = max.toString();
            }
          }

          // Apply maxLength only if specified
          if (maxLength && newValue.length > maxLength) {
            newValue = newValue.slice(0, maxLength);
          }

          e.target.value = newValue;
          onChange && onChange(e);
        }}
        onInput={(e) => {
          if (type === "number") {
            e.target.value = e.target.value.replace(/[^0-9.]/g, "");
          }
        }}
        defaultValue={defaultValue}
        inputMode={type === "number" ? "decimal" : "text"} // Ensure correct input mode
        className={`w-full border border-gray-400 rounded px-4 py-3 text-lg focus:ring-2 focus:ring-green-500 transition ${className}`}
      />
      <PencilSquareIcon className="absolute top-1/2 right-4 transform -translate-y-1/2 w-6 h-6 text-gray-500" />
    </div>
  );
};

export default Input;
