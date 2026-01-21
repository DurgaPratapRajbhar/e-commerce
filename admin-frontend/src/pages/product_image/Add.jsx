import { useState } from "react";
import { useParams } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Button from "../../components/Button";
import Input from "../../components/Input";
import Label from "../../components/Label";
import BackButton from "../../components/BackButton";
import { api } from "../../lib/api";

const AddProductImage = () => {
  const {product_id}=useParams();
  const [imagePreview, setImagePreview] = useState(null);
  const [imageFile, setImageFile] = useState(null);
  const [altText, setAltText] = useState("");
  const [isPrimary, setIsPrimary] = useState(false);

  const env=import.meta.env;
  const url=env.VITE_PRODUCT_SERVICE;
  const imageUrl=env.VITE_PRODUCT_IMAGES;
 

  // Handle Image Selection
  const handleImageChange = (event) => {
    const file = event.target.files[0];
    if (!file) return;

    // Validate File Type
    const allowedTypes = ["image/jpeg", "image/png", "image/jpg"];
    if (!allowedTypes.includes(file.type)) {
      toast.error("Invalid file type! Only JPG, JPEG, PNG allowed.");
      return;
    }

    // Validate File Size (Max 2MB)
    if (file.size > 2 * 1024 * 1024) {
      toast.error("File size too large! Max allowed: 2MB.");
      return;
    }

    // Store the image and show preview
    setImageFile(file);
    setImagePreview(URL.createObjectURL(file));
  };

  // Handle Form Submission
  const handleSubmit = async (event) => {
    event.preventDefault();
    if (!imageFile) {
      toast.error("Please select an image!");
      return;
    }

    // Prepare Form Data
    const formData = new FormData();
    formData.append("product_id", product_id);
    formData.append("image", imageFile);
    formData.append("alt_text", altText);
    formData.append("is_primary", isPrimary);

    try {
      await api.post(`${url}${imageUrl}`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      toast.success("Image uploaded successfully!");
      setImagePreview(null);
      setImageFile(null);
      setAltText("");
      setIsPrimary(false);
    } catch (err) {
      toast.error(err.message || "Failed to upload image!");
    }
  };

  return (
    <div className="bg-white shadow-lg rounded-lg p-6">
      <BackButton navigat={`/products/${product_id}/product-images`}>➕ Add Product Image</BackButton>

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

        {/* ✅ Image Preview */}
        {imagePreview && (
          <div className="mt-2">
            <img
              src={imagePreview}
              alt="Preview"
              className="h-32 w-32 object-cover rounded-md shadow-lg"
            />
          </div>
        )}

        {/* ✅ Alt Text */}
        <div>
          <Label>Alt Text:</Label>
          <Input
            type="text"
            name="alt_text"
            value={altText}
            onChange={(e) => setAltText(e.target.value)}
            required
          />
        </div>

        {/* ✅ Primary Image Checkbox */}
        <div className="flex items-center">
          <input
            type="checkbox"
            id="isPrimary"
            checked={isPrimary}
            onChange={() => setIsPrimary(!isPrimary)}
            className="mr-2"
          />
          <Label htmlFor="isPrimary">Set as Primary Image</Label>
        </div>

        <div className="flex justify-between">
          <Button type="button" onClick={() => setImagePreview(null)}>❌ Cancel</Button>
          <Button type="submit">✅ Upload Image</Button>
        </div>
      </form>
    </div>
  );
};

export default AddProductImage;
