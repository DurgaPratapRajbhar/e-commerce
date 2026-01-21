import { Route } from "react-router-dom";
import UserProfile from "../pages/user/UserProfile";
import UserAddresses from "../pages/user/UserAddresses";
import UserList from "../pages/user/UserList";
import AddressList from "../pages/user/AddressList";

const UserRoutes = () => {
  return (
    <>
      <Route path="user/profile" element={<UserProfile />} />
      <Route path="user/addresses" element={<UserAddresses />} />
      <Route path="user/list" element={<UserList />} />
      <Route path="user/addresses/list" element={<AddressList />} />
    </>
  );
};

export default UserRoutes;