const Textarea = ({
  name,
  defaultValue,
  className,
  placeholder = "Enter text here",
  rows = 4,
  color = "gray",
}) => {
  return (
    <textarea
      name={name}
      defaultValue={defaultValue}
      placeholder={placeholder}
      rows={rows}
      className={`w-full border border-${color}-400 rounded px-4 py-3 text-lg resize-none focus:ring-2 focus:ring-${color}-500 transition duration-200 ease-in-out hover:border-${color}-500 ${className}`}
    />
  );
};

export default Textarea;
