import { useState, useEffect } from "react";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { logout } from "../../features/authSlice";

const Header = () => {
  const [menuOpen, setMenuOpen] = useState(false);
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const dispatch = useDispatch();
  const navigate = useNavigate();

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (!event.target.closest(".dropdown")) {
        setDropdownOpen(false);
      }
    };
    document.addEventListener("click", handleClickOutside);
    return () => document.removeEventListener("click", handleClickOutside);
  }, []);

  const handleLogout = () => {
    dispatch(logout()); // Redux action call
    navigate("/login"); // Redirect to login page
  };

  return ( <></>
    // <header className=" fixed top-0 left-64 right-0 bg-green-400 p-4 shadow-md lg:px-16 px-8 h-16 flex items-center z-50">
    //   <div className="container mx-auto flex flex-wrap items-center">
    //     <div className="flex-1 flex justify-between items-center">
    //       <a href="#" className="text-xl font-semibold text-white hover:text-gray-200">
    //         Admin Dashboard
    //       </a>
    //     </div>

    //     {/* Mobile Menu Button */}
    //     <button onClick={() => setMenuOpen(!menuOpen)} className="md:hidden block focus:outline-none">
    //       <svg className="w-6 h-6 text-white" viewBox="0 0 24 24">
    //         <path fill="currentColor" d="M4 6h16M4 12h16m-7 6h7"></path>
    //       </svg>
    //     </button>

    //     {/* Main Navigation */}
    //     <nav className={`${menuOpen ? "block" : "hidden"} md:flex md:items-center`}>
    //       <ul className="md:flex items-center text-white">
    //         <li><a className="md:px-4 py-2 block hover:bg-green-500" href="#">Home</a></li>
    //         <li><a className="md:px-4 py-2 block hover:bg-green-500" href="#">Services</a></li>
    //         <li><a className="md:px-4 py-2 block hover:bg-green-500" href="#">Contact</a></li>
            
    //         {/* Dropdown */}
    //         <li className="relative dropdown" onMouseEnter={() => setDropdownOpen(true)} onMouseLeave={() => setDropdownOpen(false)}>
    //           <button onClick={() => setDropdownOpen(!dropdownOpen)} className="md:px-4 py-2 flex items-center hover:bg-green-500">
    //             More
    //             <svg className="w-4 h-4 ml-1" viewBox="0 0 20 20">
    //               <path fill="currentColor" d="M5 8l5 5 5-5"></path>
    //             </svg>
    //           </button>
    //           {dropdownOpen && (
    //             <ul className="absolute right-0 bg-white shadow-md text-gray-800">
    //               <li><a className="block px-4 py-2 hover:bg-gray-200" href="#">Settings</a></li>
    //               <li>
    //                 <button onClick={handleLogout} className="block w-full text-left px-4 py-2 hover:bg-gray-200">
    //                   Logout
    //                 </button>
    //               </li>
    //             </ul>
    //           )}
    //         </li>
    //       </ul>
    //     </nav>
    //   </div>
    // </header>
  );
};

export default Header;
