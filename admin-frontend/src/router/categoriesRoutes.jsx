import { Route } from "react-router-dom";
import CategoryList from "../pages/categories/List";
import AddCategory from "../pages/categories/Add";
import EditCategory from "../pages/categories/Edit";

const CategoriesRoutes = () => {
  return (
    <>
      <Route path="/categories" element={<CategoryList />} />
      <Route path="/categories/add" element={<AddCategory />} />
      <Route path="/categories/edit/:id" element={<EditCategory />} /> 
    </>
  );
};

export default CategoriesRoutes;
