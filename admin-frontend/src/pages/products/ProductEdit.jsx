import React, { useState, useEffect } from "react";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import CategorySearch from "../../components/CategorySearch";
import { productService } from "../../services/productService";
import { useNavigate, useParams } from "react-router-dom";

const EditProduct = () => {
  const { id } = useParams();
  const navigate = useNavigate();



  // State for product data
  const [productData, setProductData] = useState({
    name: "",
    category_id: "",
    description: "",
    price: "",
    discount: "",
    stock: "",
    brand: "",
    status: "active",
    sku: "",
    uom_id: "",
    quantity_value: "",
    variants: [],
    attributes: []
  });

  // UI state
  const [activeTab, setActiveTab] = useState("basic");
  const [isLoading, setIsLoading] = useState(true);
  const [showCategoryModal, setShowCategoryModal] = useState(false);
  const [errors, setErrors] = useState({});
  const [selectedCategory, setSelectedCategory] = useState(null);
  const [units, setUnits] = useState([]);

  // Fetch initial product data and units
  useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true);
      try {
        // Fetch units
        const unitsResponse = await productService.getProductUnits();
        const processedUnits = unitsResponse.data.map((unit) => ({
          id: unit.id,
          name: unit.name || unit.Name,
        }));
        setUnits(processedUnits);

        // Fetch product data
        const productResponse = await productService.getProductById(id);
        const data = productResponse.data;
        
        setProductData({
          name: data.name || "",
          category_id: data.category_id || "",
          description: data.description || "",
          price: data.price || "",
          discount: data.discount || "",
          stock: data.stock || "",
          brand: data.brand || "",
          status: data.status || "active",
          sku: data.sku || "",
          uom_id: data.uom_id || "",
          quantity_value: data.quantity_value || "",
          variants: data.variants || [],
          attributes: data.attributes || []
        });

        // Set selected category if available
        if (data.category_id) {
          const categoryResponse = await productService.getCategoryById(data.category_id);
          setSelectedCategory(categoryResponse.data);
        }

      } catch (err) {
        console.error("Error loading data:", err);
        toast.error("Failed to load product data");
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, [id]);

  // Handle input change
  const handleChange = (e) => {
    const { name, value } = e.target;
    setProductData(prev => ({ ...prev, [name]: value }));
    setErrors(prev => ({ ...prev, [name]: "" }));
  };

  // Handle category selection
  const handleSelectCategory = (category) => {
    setSelectedCategory(category);
    setProductData(prev => ({ ...prev, category_id: category.id }));
    setErrors(prev => ({ ...prev, category_id: "" }));
    setShowCategoryModal(false);
    toast.success(`Selected: ${category.name}`);
  };

  // Variant and attribute handlers (same as ProductAdmin)
  const addVariant = () => {
    setProductData(prev => ({
      ...prev,
      variants: [...prev.variants, { 
        name: "", 
        price: "", 
        stock: "", 
        sku: "", 
        uom_id: "", 
        quantity_value: "", 
        size: "", 
        color: "" 
      }]
    }));
  };

  const updateVariant = (index, field, value) => {
    setProductData(prev => {
      const updatedVariants = [...prev.variants];
      updatedVariants[index][field] = value;
      return { ...prev, variants: updatedVariants };
    });
  };

  const addAttribute = () => {
    setProductData(prev => ({
      ...prev,
      attributes: [...prev.attributes, { key: "", value: "" }]
    }));
  };

  const updateAttribute = (index, field, value) => {
    setProductData(prev => {
      const updatedAttributes = [...prev.attributes];
      updatedAttributes[index][field] = value;
      return { ...prev, attributes: updatedAttributes };
    });
  };

  // Validation (same as ProductAdmin)
  const validateForm = () => {
    const newErrors = {};
    
    if (!productData.name.trim()) newErrors.name = "Product name is required";
    if (!productData.category_id) newErrors.category_id = "Category is required";
    if (!productData.price || productData.price <= 0) newErrors.price = "Valid price is required";
    if (!productData.sku.trim()) newErrors.sku = "SKU is required";
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const normalizeProductData = (data) => {
    return {
      ...data,
      price: data.price ? Number(data.price) : 0,
      discount: data.discount ? Number(data.discount) : 0,
      stock: data.stock ? Number(data.stock) : 0,
      category_id: data.category_id ? Number(data.category_id) : undefined,
      uom_id: data.uom_id ? Number(data.uom_id) : undefined,
      quantity_value: data.quantity_value ? Number(data.quantity_value) : undefined,
      variants: data.variants.map(variant => ({
        ...variant,
        price: variant.price ? Number(variant.price) : 0,
        stock: variant.stock ? Number(variant.stock) : 0,
        uom_id: variant.uom_id ? Number(variant.uom_id) : undefined,
        quantity_value: variant.quantity_value ? Number(variant.quantity_value) : 0 // or null if reverted to *float64
      })),
      attributes: data.attributes,
    };
  };

 

  // Form submission
  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!validateForm()) {
      toast.error("Please validate all required fields");
      return;
    }
    alert(productData.quantity_value)
    setIsLoading(true);
    try {
      const normalizedData = normalizeProductData(productData);
      const response = await productService.updateProduct(id, normalizedData);
      toast.success(response.message || "Product updated successfully!");
      navigate("/products");
    } catch (err) {
      console.error("Error updating product:", err);
      toast.error("Failed to update product");
    } finally {
      setIsLoading(false);
    }
  };

  const CategoryModal = () => (
    <CategorySearch
      onSelectCategory={handleSelectCategory}
      selectedCategoryId={selectedCategory?.id}
      setShowCategoryModal={setShowCategoryModal}
    />
  );

  // Render same UI as ProductAdmin with edit functionality
  return (
    <div className="bg-gray-50 min-h-screen pb-10">
      <ToastContainer position="top-right" autoClose={3000} />
      
      {showCategoryModal && <CategoryModal />}
      
      <div className="bg-white border-b shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-4">
            <h1 className="text-2xl font-bold text-gray-900">Edit Product</h1>
            <div className="flex space-x-2">
              <button 
                onClick={() => navigate("/products")}
                className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 bg-white hover:bg-gray-50"
              >
                Cancel
              </button>
              <button
                onClick={handleSubmit}
                className="px-4 py-2 bg-blue-600 rounded-md text-white hover:bg-blue-700 flex items-center"
                disabled={isLoading}
              >
                {isLoading ? (
                  <>
                    <svg className="animate-spin h-4 w-4 mr-2" viewBox="0 0 24 24">
                      <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                      <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Updating...
                  </>
                ) : "Update Product"}
              </button>
            </div>
          </div>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-6">
        <div className="bg-white rounded-t-lg border-b">
          <nav className="flex">
            {["basic", "pricing", "inventory", "variants", "attributes"].map((tab) => (
              <button
                key={tab}
                onClick={() => setActiveTab(tab)}
                className={`px-6 py-3 text-sm font-medium border-b-2 -mb-px ${
                  activeTab === tab 
                    ? "border-blue-500 text-blue-600" 
                    : "border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300"
                }`}
              >
                {tab.charAt(0).toUpperCase() + tab.slice(1)}
              </button>
            ))}
          </nav>
        </div>

        <div className="bg-white rounded-b-lg shadow p-6">
          <form>
            {activeTab === "basic" && (
              <div className="space-y-6">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Product Name <span className="text-red-500">*</span>
                    </label>
                    <input
                      type="text"
                      name="name"
                      value={productData.name}
                      onChange={handleChange}
                      placeholder="Enter product name"
                      className={`w-full p-2 border rounded ${errors.name ? "border-red-500" : "border-gray-300"}`}
                    />
                    {errors.name && <p className="mt-1 text-sm text-red-500">{errors.name}</p>}
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Category <span className="text-red-500">*</span>
                    </label>
                    <div className="flex">
                      <input
                        type="text"
                        readOnly
                        value={selectedCategory ? selectedCategory.name : ""}
                        placeholder="Select a category"
                        className={`w-full p-2 border rounded-l ${errors.category_id ? "border-red-500" : "border-gray-300"}`}
                      />
                      <button
                        type="button"
                        onClick={() => setShowCategoryModal(true)}
                        className="px-4 py-2 bg-gray-100 border border-gray-300 rounded-r hover:bg-gray-200"
                      >
                        Browse
                      </button>
                    </div>
                    {errors.category_id && <p className="mt-1 text-sm text-red-500">{errors.category_id}</p>}
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Brand
                    </label>
                    <input
                      type="text"
                      name="brand"
                      value={productData.brand}
                      onChange={handleChange}
                      placeholder="Enter brand name"
                      className="w-full p-2 border border-gray-300 rounded"
                    />
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Status
                    </label>
                    <select
                      name="status"
                      value={productData.status}
                      onChange={handleChange}
                      className="w-full p-2 border border-gray-300 rounded"
                    >
                      <option value="active">Active</option>
                      <option value="inactive">Inactive</option>
                      <option value="draft">Draft</option>
                    </select>
                  </div>
                </div>
                
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    Description
                  </label>
                  <textarea
                    name="description"
                    value={productData.description}
                    onChange={handleChange}
                    rows="4"
                    placeholder="Enter product description"
                    className="w-full p-2 border border-gray-300 rounded"
                  ></textarea>
                </div>
              </div>
            )}
            
            {activeTab === "pricing" && (
              <div className="space-y-6">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Price <span className="text-red-500">*</span>
                    </label>
                    <div className="relative">
                      <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <span className="text-gray-500">₹</span>
                      </div>
                      <input
                        type="number"
                        name="price"
                        value={productData.price}
                        onChange={handleChange}
                        placeholder="0.00"
                        min="0"
                        step="0.01"
                        className={`w-full p-2 pl-8 border rounded ${errors.price ? "border-red-500" : "border-gray-300"}`}
                      />
                    </div>
                    {errors.price && <p className="mt-1 text-sm text-red-500">{errors.price}</p>}
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Discount (%)
                    </label>
                    <input
                      type="number"
                      name="discount"
                      value={productData.discount}
                      onChange={handleChange}
                      placeholder="0"
                      min="0"
                      max="100"
                      className="w-full p-2 border border-gray-300 rounded"
                    />
                  </div>
                </div>
              </div>
            )}
            
            {activeTab === "inventory" && (
              <div className="space-y-6">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      SKU <span className="text-red-500">*</span>
                    </label>
                    <input
                      type="text"
                      name="sku"
                      value={productData.sku}
                      onChange={handleChange}
                      placeholder="Enter unique SKU"
                      className={`w-full p-2 border rounded ${errors.sku ? "border-red-500" : "border-gray-300"}`}
                    />
                    {errors.sku && <p className="mt-1 text-sm text-red-500">{errors.sku}</p>}
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Stock Quantity
                    </label>
                    <input
                      type="number"
                      name="stock"
                      value={productData.stock}
                      onChange={handleChange}
                      placeholder="0"
                      min="0"
                      className="w-full p-2 border border-gray-300 rounded"
                    />
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Unit of Measurement
                    </label>
                    <select
                      name="uom_id"
                      value={productData.uom_id}
                      onChange={handleChange}
                      className="w-full p-2 border border-gray-300 rounded"
                    >
                      <option value="">Select Unit</option>
                      {units.map(unit => (
                        <option key={unit.id} value={unit.id}>{unit.name}</option>
                      ))}
                    </select>
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Quantity Value
                    </label>
                    <input
                      type="number"
                      name="quantity_value"
                      value={productData.quantity_value}
                      onChange={handleChange}
                      placeholder="e.g., 1.5"
                      min="0"
                      step="0.01"
                      className="w-full p-2 border border-gray-300 rounded"
                    />
                  </div>
                </div>
              </div>
            )}
            
            {/* {activeTab === "variants" && (
              <div className="space-y-6">
                <div className="flex justify-between items-center">
                  <h3 className="text-lg font-medium">Product Variants</h3>
                  <button
                    type="button"
                    onClick={addVariant}
                    className="px-3 py-1 bg-blue-50 text-blue-600 rounded border border-blue-200 hover:bg-blue-100"
                  >
                    + Add Variant
                  </button>
                </div>
                
                {productData.variants.length === 0 ? (
                  <div className="bg-gray-50 border border-dashed border-gray-300 rounded-lg p-6 text-center">
                    <p className="text-gray-500">No variants added yet.</p>
                  </div>
                ) : (
                  <div className="space-y-4">
                    {productData.variants.map((variant, index) => (
                      <div key={index} className="border rounded-lg overflow-hidden">
                        <div className="bg-gray-50 px-4 py-2 flex justify-between items-center">
                          <h4 className="font-medium">
                            Variant #{index + 1}: {variant.name || "Unnamed"}
                          </h4>
                          <button
                            type="button"
                            onClick={() => {
                              setProductData(prev => ({
                                ...prev,
                                variants: prev.variants.filter((_, i) => i !== index)
                              }));
                            }}
                            className="text-red-600 text-sm hover:text-red-800"
                          >
                            Remove
                          </button>
                        </div>
                        <div className="p-4 grid grid-cols-1 md:grid-cols-3 gap-4">
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">Name</label>
                            <input
                              type="text"
                              value={variant.name}
                              onChange={(e) => updateVariant(index, "name", e.target.value)}
                              className="w-full p-2 border border-gray-300 rounded"
                            />
                          </div>
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">SKU</label>
                            <input
                              type="text"
                              value={variant.sku}
                              onChange={(e) => updateVariant(index, "sku", e.target.value)}
                              className="w-full p-2 border border-gray-300 rounded"
                            />
                          </div>
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">Price</label>
                            <input
                              type="number"
                              value={variant.price}
                              onChange={(e) => updateVariant(index, "price", e.target.value)}
                              className="w-full p-2 border border-gray-300 rounded"
                            />
                          </div>
                          
                        </div>
                      </div>
                    ))}
                  </div>
                )}
              </div>
            )} */}


            {/* Variants */}
            {activeTab === "variants" && (
              <div className="space-y-6">
                <div className="flex justify-between items-center">
                  <h3 className="text-lg font-medium">Product Variants</h3>
                  <button
                    type="button"
                    onClick={addVariant}
                    className="px-3 py-1 bg-blue-50 text-blue-600 rounded border border-blue-200 hover:bg-blue-100"
                  >
                    + Add Variant
                  </button>
                </div>
                {productData.variants.length === 0 ? (
                  <div className="bg-gray-50 border border-dashed border-gray-300 rounded-lg p-6 text-center">
                    <p className="text-gray-500">No variants added yet.</p>
                    <button
                      type="button"
                      onClick={addVariant}
                      className="mt-3 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
                    >
                      Add Your First Variant
                    </button>
                  </div>
                ) : (
                  <div className="space-y-4">
                    {productData.variants.map((variant, index) => (
                      <div key={index} className="border rounded-lg overflow-hidden">
                        <div className="bg-gray-50 px-4 py-2 flex justify-between items-center">
                          <h4 className="font-medium">Variant #{index + 1}: {variant.name || "Unnamed"}</h4>
                          <button
                            type="button"
                            onClick={() => setProductData(prev => ({ ...prev, variants: prev.variants.filter((_, i) => i !== index) }))}
                            className="text-red-600 text-sm hover:text-red-800"
                          >
                            Remove
                          </button>
                        </div>
                        <div className="p-4 grid grid-cols-1 md:grid-cols-3 gap-4">
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                              Name <span className="text-red-500">*</span>
                            </label>
                            <input
                              type="text"
                              value={variant.name}
                              onChange={(e) => updateVariant(index, "name", e.target.value)}
                              placeholder="e.g., Large Size"
                              className={`w-full p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500 ${
                                errors[`variant_${index}_name`] ? "border-red-500" : "border-gray-300"
                              }`}
                            />
                            {errors[`variant_${index}_name`] && <p className="mt-1 text-sm text-red-500">{errors[`variant_${index}_name`]}</p>}
                          </div>
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                              Price <span className="text-red-500">*</span>
                            </label>
                            <div className="relative">
                              <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                                <span className="text-gray-500">₹</span>
                              </div>
                              <input
                                type="number"
                                value={variant.price}
                                onChange={(e) => updateVariant(index, "price", e.target.value)}
                                placeholder="0.00"
                                min="0"
                                step="0.01"
                                className={`w-full p-2 pl-8 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500 ${
                                  errors[`variant_${index}_price`] ? "border-red-500" : "border-gray-300"
                                }`}
                              />
                            </div>
                            {errors[`variant_${index}_price`] && <p className="mt-1 text-sm text-red-500">{errors[`variant_${index}_price`]}</p>}
                          </div>
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">Stock</label>
                            <input
                              type="number"
                              value={variant.stock}
                              onChange={(e) => updateVariant(index, "stock", e.target.value)}
                              placeholder="0"
                              min="0"
                              className={`w-full p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500 ${
                                errors[`variant_${index}_stock`] ? "border-red-500" : "border-gray-300"
                              }`}
                            />
                            {errors[`variant_${index}_stock`] && <p className="mt-1 text-sm text-red-500">{errors[`variant_${index}_stock`]}</p>}
                          </div>
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">Unit of Measurement</label>
                            <select
                              value={variant.uom_id}
                              onChange={(e) => updateVariant(index, "uom_id", e.target.value)}
                              className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                            >
                              <option value="">Select Unit</option>
                              {units.map(unit => (
                                <option key={unit.id} value={unit.id}>{unit.name}</option>
                              ))}
                            </select>
                          </div>
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">Quantity Value</label>
                            <input
                              type="number"
                              value={variant.quantity_value}
                              onChange={(e) => updateVariant(index, "quantity_value", e.target.value)}
                              placeholder="e.g., 1.5"
                              min="0"
                              step="0.01"
                              className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                            />
                          </div>
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">Size</label>
                            <input
                              type="text"
                              value={variant.size}
                              onChange={(e) => updateVariant(index, "size", e.target.value)}
                              placeholder="e.g., XL"
                              className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                            />
                          </div>
                          <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">Color</label>
                            <input
                              type="text"
                              value={variant.color}
                              onChange={(e) => updateVariant(index, "color", e.target.value)}
                              placeholder="e.g., Red"
                              className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                            />
                          </div>

                           <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">Weight</label>
                            <input
                              type="text"
                              value={variant.weight}
                              onChange={(e) => updateVariant(index, "weight", e.target.value)}
                              placeholder="e.g., 500g"
                              className="w-full p-2 border border-gray-300 rounded"
                            />
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                )}
              </div>
            )}

            
            {activeTab === "attributes" && (
              <div className="space-y-6">
                <div className="flex justify-between items-center">
                  <h3 className="text-lg font-medium">Product Attributes</h3>
                  <button
                    type="button"
                    onClick={addAttribute}
                    className="px-3 py-1 bg-blue-50 text-blue-600 rounded border border-blue-200 hover:bg-blue-100"
                  >
                    + Add Attribute
                  </button>
                </div>
                
                {productData.attributes.length === 0 ? (
                  <div className="bg-gray-50 border border-dashed border-gray-300 rounded-lg p-6 text-center">
                    <p className="text-gray-500">No attributes added yet.</p>
                  </div>
                ) : (
                  <div className="space-y-4">
                    {productData.attributes.map((attr, index) => (
                      <div key={index} className="flex gap-4 items-start">
                        <div className="flex-1">
                          <input
                            type="text"
                            value={attr.key}
                            onChange={(e) => updateAttribute(index, "key", e.target.value)}
                            placeholder="Attribute name"
                            className="w-full p-2 border border-gray-300 rounded"
                          />
                        </div>
                        <div className="flex-1">
                          <input
                            type="text"
                            value={attr.value}
                            onChange={(e) => updateAttribute(index, "value", e.target.value)}
                            placeholder="Value"
                            className="w-full p-2 border border-gray-300 rounded"
                          />
                        </div>
                        <button
                          type="button"
                          onClick={() => {
                            setProductData(prev => ({
                              ...prev,
                              attributes: prev.attributes.filter((_, i) => i !== index)
                            }));
                          }}
                          className="px-2 py-2 text-red-600 hover:text-red-800"
                        >
                          ×
                        </button>
                      </div>
                    ))}
                  </div>
                )}
              </div>
            )}
          </form>
        </div>

        <div className="mt-6 flex justify-end space-x-3">
          <button
            type="button"
            onClick={() => navigate("/products")}
            className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 bg-white hover:bg-gray-50"
          >
            Cancel
          </button>
          <button
            onClick={handleSubmit}
            className="px-6 py-2 bg-blue-600 border border-transparent rounded-md text-white font-medium hover:bg-blue-700"
            disabled={isLoading}
          >
            {isLoading ? "Updating..." : "Update Product"}
          </button>
        </div>
      </div>
    </div>
  );
};

export default EditProduct;







