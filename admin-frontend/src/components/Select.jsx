const Select = ({ name, className = "", value, data = [], defaultName = "Select" ,onChange}) => {
  return (
    <select
      name={name}
      className={`w-full border border-gray-400 rounded px-4 py-3 text-lg ${className}`}
      value={value}  
      onChange={onChange} 
    >
      <option value="0">{defaultName}</option>
      {data.map((item, index) => (
        <option key={index} value={item.value}>
          {item.name}
        </option>
      ))}
    </select>
  );
};
  
export default Select;
