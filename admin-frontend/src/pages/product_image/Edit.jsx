import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Button from "../../components/Button";
import Input from "../../components/Input";
import Label from "../../components/Label";
import BackButton from "../../components/BackButton";
import { api } from "../../lib/api";

const EditProductImage = () => {
  const { product_id, id } = useParams(); // ✅ Extract product_id & image id
  const [imagePreview, setImagePreview] = useState(null);
  const [imageFile, setImageFile] = useState(null);
  const [altText, setAltText] = useState("");
  const [isPrimary, setIsPrimary] = useState(false);

  const env=import.meta.env;
  const url=env.VITE_PRODUCT_SERVICE;
  const imageUrl=env.VITE_PRODUCT_IMAGES;
  const IMAGE_URL = env.VITE_IMAGE_ROOT;

  useEffect(() => {
    const fetchImageData = async () => {
      try {
        const response = await api.get(`${url}${imageUrl}/by-id/${id}`);
        console.log(response)
        const data = response.data.data;
        setAltText(data.alt_text);
        setIsPrimary(data.is_primary);
        setImagePreview(IMAGE_URL+data.image_url);
      } catch (err) {
        toast.error("Failed to fetch image details!");
      }
    };

    fetchImageData();
  }, [id]);

  const handleImageChange = (event) => {
    const file = event.target.files[0];
    if (!file) return;

    const allowedTypes = ["image/jpeg", "image/png", "image/jpg"];
    if (!allowedTypes.includes(file.type)) {
      toast.error("Invalid file type! Only JPG, JPEG, PNG allowed.");
      return;
    }

    if (file.size > 2 * 1024 * 1024) {
      toast.error("File size too large! Max allowed: 2MB.");
      return;
    }

    setImageFile(file);
    setImagePreview(URL.createObjectURL(file));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    const formData = new FormData();
    formData.append("product_id", product_id);
    if (imageFile) formData.append("image", imageFile);
    formData.append("alt_text", altText);
    formData.append("is_primary", isPrimary);

    try {
      await api.put(`${url}${imageUrl}/${id}`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      toast.success("Image updated successfully!");
    } catch (err) {
      toast.error(err.message || "Failed to update image!");
    }
  };

  return (
    <div className="bg-white shadow-lg rounded-lg p-6">
      <BackButton  navigat={`/products/${product_id}/product-images`}>📝 Edit Product Image</BackButton>
      <ToastContainer />

      <form className="mt-6 space-y-6" onSubmit={handleSubmit}>
        <div>
          <Label>Image:</Label>
          <input
            type="file"
            name="image"
            accept="image/*"
            onChange={handleImageChange}
            className="w-full border border-gray-400 rounded px-4 py-2 text-lg focus:ring-2 focus:ring-green-500 transition"
          />
        </div>

        {imagePreview && (
          <div className="mt-2">
            <img
              src={imagePreview}
              alt="Preview"
              className="h-32 w-32 object-cover rounded-md shadow-lg"
            />
          </div>
        )}

        <div>
          <Label>Alt Text:</Label>
          <Input
            type="text"
            name="alt_text"
            defaultValue={altText}
            onChange={(e) => setAltText(e.target.value)}
            required
          />
        </div>

        <div className="flex justify-between">
          <Button type="button">❌ Cancel</Button>
          <Button type="submit">✅ Update Image</Button>
        </div>
      </form>
    </div>
  );
};

export default EditProductImage;
