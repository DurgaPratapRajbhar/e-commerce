import { useState, useEffect } from "react";
import { useNavigate ,useParams} from "react-router-dom";
import { Squares2X2Icon , TrashIcon } from "@heroicons/react/24/solid"; 
import { api } from "../../lib/api";
import { API_ENDPOINTS } from "../../config/apiConfig";

const List = () => {
 
  const env=import.meta.env;
  const url=env.VITE_PRODUCT_SERVICE;
  const imageUrl=env.VITE_PRODUCT_IMAGES;
  const IMAGE_URL = env.VITE_IMAGE_ROOT;
  
  alert(imageUrl)
  const {product_id}=useParams();
  const [images, setImages] = useState([]);
  const navigate = useNavigate();  

  // Fetch the list of product images when the component mounts
  useEffect(() => {
    getProductImages();
  }, []);

  // Get product images from the server
  const getProductImages = async () => {
    try {
      const response = await api.get(`${url}${imageUrl}/by-product/${product_id}`);
      setImages(response.data.data); // Assuming the response contains the image data
    } catch (err) {
      console.log(err);
    }
  };

  // Handle the deletion of a product image
  const deleteProductImage = async (id) => {
    if (!window.confirm(`Do you really want to delete this image?`)) {
      return;
    }
    try {
      await api.delete(`${url}${imageUrl}/${id}`);
      setImages(images.filter((image) => image.id !== id)); // Remove the deleted image from the state
      
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className="bg-white shadow-lg rounded-lg p-6">
      <h2 className="text-2xl font-semibold text-gray-800">📸 Product Images</h2>

      <button
        className="bg-green-400 text-white px-4 py-2 rounded mt-4 hover:bg-green-500 transition flex items-center gap-2"
        onClick={() => navigate(`/products/${product_id}/product-images/add`)}

      >
        ➕ Add Product Image
      </button>

      <table className="w-full mt-4 border-collapse border border-gray-300">
        <thead className="bg-gray-200">
          <tr>
            <th className="border px-4 py-2">#</th>
            <th className="border px-4 py-2">Image</th>
            <th className="border px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {images.map((image, index) => (
            <tr key={image.id} className="text-center">
              <td className="border px-4 py-2">{index + 1}</td>
              <td className="border px-4 py-2">
                <img
                  src={IMAGE_URL + image.image_url}
                  alt="Product"
                  className="w-20 h-20 object-cover rounded-lg shadow-md border border-gray-300"
                />
              </td>
              <td className="border px-6 py-2">
              <div className="flex flex-row items-center gap-2">
                <button
                  className="text-blue-500 hover:text-blue-700"
                  onClick={() => navigate(`/products/${product_id}/product-images/edit/${image.id}`)}
                >
                  <Squares2X2Icon  className="h-6 w-6" />
                </button>

                <button
                  className="text-red-500 hover:text-red-700"
                  onClick={() => deleteProductImage(image.id)}
                >
                  <TrashIcon className="h-6 w-6" />
                </button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default List;
