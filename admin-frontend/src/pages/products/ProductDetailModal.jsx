import { useState } from "react";
import { XMarkIcon } from "@heroicons/react/24/solid";

const ProductDetailModal = ({ product, onClose }) => {
  const [activeTab, setActiveTab] = useState("details");
  
  const env=import.meta.env;
  const IMAGE_URL = env.VITE_IMAGE_ROOT;

  const formatPrice = (price, discount = 0) => {
    const actualPrice = price - (price * discount / 100);
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'INR' }).format(actualPrice);
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    });
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-lg shadow-xl w-full max-w-4xl max-h-[90vh] overflow-hidden flex flex-col">
        {/* Header */}
        <div className="flex justify-between items-center p-6 border-b">
          <h3 className="text-xl font-bold text-gray-800">{product.name}</h3>
          <button
            onClick={onClose}
            className="text-gray-400 hover:text-gray-500 focus:outline-none"
          >
            <XMarkIcon className="h-6 w-6" />
          </button>
        </div>

        {/* Tabs */}
        <div className="flex border-b">
          {[
            { id: "details", label: "Details" },
            { id: "attributes", label: "Attributes" },
            { id: "variants", label: "Variants" },
            { id: "images", label: "Images" }
          ].map((tab) => (
            <button
              key={tab.id}
              className={`py-3 px-6 focus:outline-none ${
                activeTab === tab.id
                  ? "border-b-2 border-indigo-500 text-indigo-600 font-medium"
                  : "text-gray-500 hover:text-gray-700"
              }`}
              onClick={() => setActiveTab(tab.id)}
            >
              {tab.label}
            </button>
          ))}
        </div>

        {/* Content */}
        <div className="overflow-y-auto p-6 flex-grow">
          {activeTab === "details" && (
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div className="space-y-4">
                {/* Main product info */}
                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">SKU</h4>
                  <p className="text-gray-800">{product.sku}</p>
                </div>
                
                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Price</h4>
                  <div className="flex items-center">
                    <p className="text-lg font-semibold text-gray-800">
                      {formatPrice(product.price, product.discount)}
                    </p>
                    {product.discount > 0 && (
                      <p className="ml-2 text-sm text-gray-500 line-through">
                        {formatPrice(product.price)}
                      </p>
                    )}
                  </div>
                  {product.discount > 0 && (
                    <p className="text-sm text-red-600 mt-1">
                      {product.discount}% discount applied
                    </p>
                  )}
                </div>

                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Stock</h4>
                  <p className={`${
                    product.stock > 10 
                      ? 'text-green-600' 
                      : product.stock > 0 
                        ? 'text-yellow-600' 
                        : 'text-red-600'
                  }`}>
                    {product.stock} 
                    {product.unit && product.quantity_value 
                      ? ` (${product.quantity_value} ${product.unit.symbol})` 
                      : ''}
                  </p>
                </div>

                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Status</h4>
                  <span className={`px-2 py-1 rounded-full text-xs ${
                    product.status === 'active' 
                      ? 'bg-green-100 text-green-800' 
                      : product.status === 'draft' 
                        ? 'bg-yellow-100 text-yellow-800' 
                        : 'bg-red-100 text-red-800'
                  }`}>
                    {product.status}
                  </span>
                </div>

                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Created</h4>
                  <p className="text-gray-800">{formatDate(product.created_at)}</p>
                </div>

                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Last Updated</h4>
                  <p className="text-gray-800">{formatDate(product.updated_at)}</p>
                </div>
              </div>

              <div className="space-y-4">
                {/* Category and brand */}
                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Category</h4>
                  <p className="text-gray-800">{product.category?.name || "N/A"}</p>
                </div>

                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Brand</h4>
                  <p className="text-gray-800">{product.brand || "N/A"}</p>
                </div>

                <div>
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Slug</h4>
                  <p className="text-gray-800">{product.slug}</p>
                </div>

                {/* Description */}
                <div className="col-span-full">
                  <h4 className="text-sm font-medium text-gray-500 mb-1">Description</h4>
                  <div className="border rounded p-3 bg-gray-50 max-h-40 overflow-y-auto">
                    <p className="text-gray-800 whitespace-pre-line">{product.description}</p>
                  </div>
                </div>
              </div>
            </div>
          )}

          {activeTab === "attributes" && (
            <div className="space-y-4">
              {product.attributes && product.attributes.length > 0 ? (
                <div className="border rounded-lg overflow-hidden">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Attribute
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Value
                        </th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {product.attributes.map((attr) => (
                        <tr key={attr.id}>
                          <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                            {attr.key}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {attr.value}
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              ) : (
                <div className="text-center py-12 text-gray-500 bg-gray-50 rounded-lg">
                  No attributes found for this product.
                </div>
              )}
            </div>
          )}

          {activeTab === "variants" && (
            <div className="space-y-4">
              {product.variants && product.variants.length > 0 ? (
                <div className="border rounded-lg overflow-hidden">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Name
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          SKU
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Price
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Stock
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Attributes
                        </th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {product.variants.map((variant) => (
                        <tr key={variant.id}>
                          <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                            {variant.name}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {variant.sku}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {formatPrice(variant.price)}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {variant.stock}
                            {variant.unit ? ` ${variant.unit.symbol}` : ''}
                          </td>
                          <td className="px-6 py-4 text-sm text-gray-500">
                            <div className="space-y-1">
                              {variant.size && <span className="inline-block mr-2">Size: {variant.size}</span>}
                              {variant.color && <span className="inline-block mr-2">Color: {variant.color}</span>}
                              {variant.material && <span className="inline-block mr-2">Material: {variant.material}</span>}
                            </div>
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              ) : (
                <div className="text-center py-12 text-gray-500 bg-gray-50 rounded-lg">
                  No variants found for this product.
                </div>
              )}
            </div>
          )}

          {activeTab === "images" && (
            <div className="space-y-4">
              {product.images && product.images.length > 0 ? (
                <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
                  {product.images.map((image, index) => (
                    <div key={index} className="border rounded-lg overflow-hidden">
                      <img
                        src={IMAGE_URL+image.image_url}
                        alt={`Product image ${index + 1}`}
                        className="w-full h-48 object-cover"
                      />
                    </div>
                  ))}
                </div>
              ) : (
                <div className="text-center py-12 text-gray-500 bg-gray-50 rounded-lg">
                  No images found for this product.
                </div>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default ProductDetailModal;