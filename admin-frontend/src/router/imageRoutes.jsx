import { Routes, Route } from "react-router-dom";
import ImageList from "../pages/product_image/List"; // List product images
import ImageAdd from "../pages/product_image/Add"; // Add new product image
import ImageEdit from "../pages/product_image/Edit"; // Edit an existing product image

const ImageRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<ImageList />} /> {/* List all images for a product */}
      <Route path="add" element={<ImageAdd />} /> {/* Add image for a product */}
      <Route path="edit/:id" element={<ImageEdit />} /> {/* Edit image for a product */}
    </Routes>
  );
};

export default ImageRoutes;
