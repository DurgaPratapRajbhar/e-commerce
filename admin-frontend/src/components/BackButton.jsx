import { useNavigate } from 'react-router-dom';

const BackButton = ({ children ,navigat }) => {
  const navigate = useNavigate();  

  return (
    <>
      <button
        onClick={() => navigate(navigat)} 
        className="text-gray-600 hover:text-gray-900 flex items-center gap-2 text-lg"
      >
        ⬅ <span className="font-medium">Back</span>
      </button>

      <h2 className="text-2xl font-semibold text-gray-800 text-center mt-4">
        {children}
      </h2>
    </>
  );
};

export default BackButton;
