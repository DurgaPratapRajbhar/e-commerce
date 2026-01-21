import Sidebar from "./Sidebar";
import Header from "./Header";
import Footer from "./Footer";
import { Outlet } from "react-router-dom";

const Layout = () => {
  return (
    <div className="flex">
       <div className="w-64 fixed h-full">
        <Sidebar />
      </div>
      <div className="flex flex-col flex-1 min-h-screen ml-64">
        <Header />
        <main className="flex-1 p-6 pt-20 bg-gradient-to-r from-green-100 via-green-200 to-green-100 min-h-screen">
          <Outlet /> 
        </main>
        <Footer />
      </div>
    </div>
  );
};

export default Layout;
