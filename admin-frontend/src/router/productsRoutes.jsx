import { Route } from "react-router-dom";
import ProductList from "../pages/products/ProductList";
import AddProduct from "../pages/products/AddProduct";
import ProductEdit from "../pages/products/ProductEdit";
import ImageRoutes from  "./imageRoutes"
const ProductsRoutes = () => {
  return (
    <>
    
      <Route path="products" element={<ProductList />} />
      <Route path="products/add" element={<AddProduct />} />
      <Route path="products/edit/:id" element={<ProductEdit />} />
      <Route path="products/:product_id/product-images/*" element={<ImageRoutes />} />
    </>
  );
};

export default ProductsRoutes;
