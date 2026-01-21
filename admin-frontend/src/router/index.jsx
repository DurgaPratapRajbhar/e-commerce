import { Route } from "react-router-dom";
import CategoriesRoutes from "./categoriesRoutes";
import ProductsRoutes from "./productsRoutes";
import UserRoutes from "./userRoutes";
// import ImageRoutes from  "./imageRoutes"
const AppRoutes = () => {
  return (
    <>
      {CategoriesRoutes()}
      {ProductsRoutes()}
      {UserRoutes()}
      {/* {ImageRoutes()} */}
    </>
  );
};

export default AppRoutes;
